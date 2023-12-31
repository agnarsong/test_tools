"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.printComparison = exports.printSectionHead = exports.getEtherscanUrl = exports.getArtifactFromManagedName = exports.color = exports.getInput = void 0;
const readline_1 = require("readline");
const core_utils_1 = require("@mantleio/core-utils");
const getInput = (query) => {
    const rl = (0, readline_1.createInterface)({
        input: process.stdin,
        output: process.stdout,
    });
    return new Promise((resolve) => rl.question(query, (ans) => {
        rl.close();
        resolve(ans);
    }));
};
exports.getInput = getInput;
const codes = {
    reset: '\x1b[0m',
    red: '\x1b[0;31m',
    green: '\x1b[0;32m',
    cyan: '\x1b[0;36m',
    yellow: '\x1b[1;33m',
};
exports.color = Object.entries(codes).reduce((obj, [k]) => {
    obj[k] = (msg) => `${codes[k]}${msg}${codes.reset}`;
    return obj;
}, {});
const locateArtifact = (name) => {
    return {
        'ChainStorageContainer-CTC-batches': 'L1/rollup/ChainStorageContainer.sol/ChainStorageContainer.json',
        'ChainStorageContainer-SCC-batches': 'L1/rollup/ChainStorageContainer.sol/ChainStorageContainer.json',
        CanonicalTransactionChain: 'L1/rollup/CanonicalTransactionChain.sol/CanonicalTransactionChain.json',
        StateCommitmentChain: 'L1/rollup/StateCommitmentChain.sol/StateCommitmentChain.json',
        BondManager: 'L1/verification/BondManager.sol/BondManager.json',
        BVM_L1CrossDomainMessenger: 'L1/messaging/L1CrossDomainMessenger.sol/L1CrossDomainMessenger.json',
        Proxy__BVM_L1CrossDomainMessenger: 'libraries/resolver/Lib_ResolvedDelegateProxy.sol/Lib_ResolvedDelegateProxy.json',
        Proxy__BVM_L1StandardBridge: 'chugsplash/L1ChugSplashProxy.sol/L1ChugSplashProxy.json',
    }[name];
};
const getArtifactFromManagedName = (name) => {
    return require(`../artifacts/contracts/${locateArtifact(name)}`);
};
exports.getArtifactFromManagedName = getArtifactFromManagedName;
const getEtherscanUrl = (network, address) => {
    const escPrefix = network.chainId !== 1 ? `${network.name}.` : '';
    return `https://${escPrefix}etherscan.io/address/${address}`;
};
exports.getEtherscanUrl = getEtherscanUrl;
const truncateLongString = (value) => {
    return value.length > 66 ? `${value.slice(0, 66)}...` : value;
};
const printSectionHead = (msg) => {
    console.log(exports.color.cyan(msg));
    console.log(exports.color.cyan('='.repeat(Math.max(...msg.split('\n').map((s) => s.length)))));
};
exports.printSectionHead = printSectionHead;
const printComparison = (action, description, expected, deployed) => {
    console.log(`\n${action}:`);
    if ((0, core_utils_1.hexStringEquals)(expected.value, deployed.value)) {
        console.log(exports.color.green(`${expected.name}: ${truncateLongString(expected.value)}`));
        console.log('matches');
        console.log(exports.color.green(`${deployed.name}: ${truncateLongString(deployed.value)}`));
        console.log(exports.color.green(`${description} looks good! 😎`));
    }
    else {
        throw new Error(`${description} looks wrong. ${expected.value}\ndoes not match\n${deployed.value}.`);
    }
};
exports.printComparison = printComparison;
//# sourceMappingURL=validation-utils.js.map