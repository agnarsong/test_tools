"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.HexToBytes = exports.BIG_BALANCE = exports.isHardhatNode = exports.getContractFromArtifact = exports.sendImpersonatedTx = exports.fundAccount = exports.getAdvancedContract = exports.deployAndVerifyAndThen = void 0;
const ethers_1 = require("ethers");
const core_utils_1 = require("@mantleio/core-utils");
const deployAndVerifyAndThen = async ({ hre, name, args, contract, iface, postDeployAction, }) => {
    const { deploy } = hre.deployments;
    const { deployer } = await hre.getNamedAccounts();
    const result = await deploy(name, {
        contract,
        from: deployer,
        args,
        log: true,
        waitConfirmations: hre.deployConfig.numDeployConfirmations,
    });
    await hre.ethers.provider.waitForTransaction(result.transactionHash);
    if (result.newlyDeployed) {
        if (!(await (0, exports.isHardhatNode)(hre))) {
            try {
                console.log('Verifying on Etherscan...');
                await hre.run('verify:verify', {
                    address: result.address,
                    constructorArguments: args,
                });
                console.log('Successfully verified on Etherscan');
            }
            catch (error) {
                console.log('Error when verifying bytecode on Etherscan:');
                console.log(error);
            }
            try {
                console.log('Verifying on Sourcify...');
                await hre.run('sourcify');
                console.log('Successfully verified on Sourcify');
            }
            catch (error) {
                console.log('Error when verifying bytecode on Sourcify:');
                console.log(error);
            }
        }
        if (postDeployAction) {
            const signer = hre.ethers.provider.getSigner(deployer);
            let abi = result.abi;
            if (iface !== undefined) {
                const factory = await hre.ethers.getContractFactory(iface);
                abi = factory.interface;
            }
            await postDeployAction((0, exports.getAdvancedContract)({
                hre,
                contract: new ethers_1.Contract(result.address, abi, signer),
            }));
        }
    }
};
exports.deployAndVerifyAndThen = deployAndVerifyAndThen;
const getAdvancedContract = (opts) => {
    const def = Object.defineProperty;
    Object.defineProperty = (obj, propName, prop) => {
        prop.writable = true;
        return def(obj, propName, prop);
    };
    const contract = new ethers_1.Contract(opts.contract.address, opts.contract.interface, opts.contract.signer || opts.contract.provider);
    Object.defineProperty = def;
    for (const fnName of Object.keys(contract.functions)) {
        const fn = contract[fnName].bind(contract);
        contract[fnName] = async (...args) => {
            let gasPrice = opts.hre.deployConfig.gasPrice || undefined;
            if (contract.interface.getFunction(fnName).constant) {
                gasPrice = 0;
            }
            const tx = await fn(...args, {
                gasPrice,
            });
            if (typeof tx !== 'object' || typeof tx.wait !== 'function') {
                return tx;
            }
            const maxTimeout = 120;
            let timeout = 0;
            while (true) {
                await (0, core_utils_1.sleep)(1000);
                const receipt = await contract.provider.getTransactionReceipt(tx.hash);
                if (receipt === null) {
                    timeout++;
                    if (timeout > maxTimeout && opts.hre.network.name === 'kovan') {
                        console.log(`WARNING: Exceeded max timeout on transaction. Attempting to submit transaction again...`);
                        return contract[fnName](...args);
                    }
                }
                else if (receipt.confirmations >= opts.hre.deployConfig.numDeployConfirmations) {
                    return tx;
                }
            }
        };
    }
    return contract;
};
exports.getAdvancedContract = getAdvancedContract;
const fundAccount = async (hre, address, amount) => {
    if (!hre.deployConfig.isForkedNetwork) {
        throw new Error('this method can only be used against a forked network');
    }
    console.log(`Funding account ${address}...`);
    await hre.ethers.provider.send('hardhat_setBalance', [
        address,
        amount.toHexString(),
    ]);
    console.log(`Waiting for balance to reflect...`);
    await (0, core_utils_1.awaitCondition)(async () => {
        const balance = await hre.ethers.provider.getBalance(address);
        return balance.gte(amount);
    }, 5000, 100);
    console.log(`Account successfully funded.`);
};
exports.fundAccount = fundAccount;
const sendImpersonatedTx = async (opts) => {
    if (!opts.hre.deployConfig.isForkedNetwork) {
        throw new Error('this method can only be used against a forked network');
    }
    console.log(`Impersonating account ${opts.from}...`);
    await opts.hre.ethers.provider.send('hardhat_impersonateAccount', [opts.from]);
    console.log(`Funding account ${opts.from}...`);
    await (0, exports.fundAccount)(opts.hre, opts.from, exports.BIG_BALANCE);
    console.log(`Sending impersonated transaction...`);
    const tx = await opts.contract.populateTransaction[opts.fn](...opts.args);
    const provider = new opts.hre.ethers.providers.JsonRpcProvider(opts.hre.network.config.url);
    await provider.send('eth_sendTransaction', [
        Object.assign(Object.assign({}, tx), { from: opts.from, gas: opts.gas }),
    ]);
    console.log(`Stopping impersonation of account ${opts.from}...`);
    await opts.hre.ethers.provider.send('hardhat_stopImpersonatingAccount', [
        opts.from,
    ]);
};
exports.sendImpersonatedTx = sendImpersonatedTx;
const getContractFromArtifact = async (hre, name, options = {}) => {
    const artifact = await hre.deployments.get(name);
    await hre.ethers.provider.waitForTransaction(artifact.receipt.transactionHash);
    let iface = new hre.ethers.utils.Interface(artifact.abi);
    if (options.iface) {
        const factory = await hre.ethers.getContractFactory(options.iface);
        iface = factory.interface;
    }
    let signerOrProvider = hre.ethers.provider;
    if (options.signerOrProvider) {
        if (typeof options.signerOrProvider === 'string') {
            signerOrProvider = hre.ethers.provider.getSigner(options.signerOrProvider);
        }
        else {
            signerOrProvider = options.signerOrProvider;
        }
    }
    return (0, exports.getAdvancedContract)({
        hre,
        contract: new hre.ethers.Contract(artifact.address, iface, signerOrProvider),
    });
};
exports.getContractFromArtifact = getContractFromArtifact;
const isHardhatNode = async (hre) => {
    return (await (0, core_utils_1.getChainId)(hre.ethers.provider)) === 31337;
};
exports.isHardhatNode = isHardhatNode;
exports.BIG_BALANCE = ethers_1.ethers.BigNumber.from(`0xFFFFFFFFFFFFFFFFFFFF`);
const HexToBytes = async (hex) => {
    const bytes = [];
    for (let c = 0; c < hex.length; c += 2) {
        bytes.push(parseInt(hex.substr(c, 2), 16));
    }
    return bytes;
};
exports.HexToBytes = HexToBytes;
//# sourceMappingURL=deploy-utils.js.map