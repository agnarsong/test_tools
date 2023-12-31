"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.parseDeployConfig = exports.getDeployConfig = void 0;
const ethers_1 = require("ethers");
const configSpec = {
    isForkedNetwork: {
        type: 'boolean',
        default: false,
    },
    numDeployConfirmations: {
        type: 'number',
        default: 0,
    },
    gasPrice: {
        type: 'number',
        default: undefined,
    },
    l1BlockTimeSeconds: {
        type: 'number',
    },
    l2BlockGasLimit: {
        type: 'number',
    },
    l2ChainId: {
        type: 'number',
    },
    ctcL2GasDiscountDivisor: {
        type: 'number',
    },
    ctcEnqueueGasCost: {
        type: 'number',
    },
    sccFaultProofWindowSeconds: {
        type: 'number',
    },
    sccSequencerPublishWindowSeconds: {
        type: 'number',
    },
    blockStaleMeasure: {
        type: 'number',
    },
    daFraudProofPeriod: {
        type: 'number',
    },
    l2SubmittedBlockNumber: {
        type: 'number',
    },
    bvmSequencerAddress: {
        type: 'address',
    },
    bvmProposerAddress: {
        type: 'address',
    },
    bvmRolluperAddress: {
        type: 'address',
    },
    bvmBlockSignerAddress: {
        type: 'address',
    },
    bvmFeeWalletAddress: {
        type: 'address',
    },
    bvmAddressManagerOwner: {
        type: 'address',
    },
    bvmGasPriceOracleOwner: {
        type: 'address',
    },
    bvmFeeWalletOwner: {
        type: 'address',
    },
    bvmWhitelistOwner: {
        type: 'address',
        default: ethers_1.ethers.constants.AddressZero,
    },
    dataManagerAddress: {
        type: 'address',
    },
    bvmEigenSequencerAddress: {
        type: 'address',
    },
    sccAddress: {
        type: 'address',
    },
    gasPriceOracleOverhead: {
        type: 'number',
        default: 2750,
    },
    gasPriceOracleScalar: {
        type: 'number',
        default: 1500000,
    },
    gasPriceOracleDecimals: {
        type: 'number',
        default: 6,
    },
    gasPriceOracleIsBurning: {
        type: 'number',
        default: 1,
    },
    gasPriceOracleCharge: {
        type: 'number',
        default: 0,
    },
    gasPriceOracleL1BaseFee: {
        type: 'number',
        default: 1,
    },
    gasPriceOracleL2GasPrice: {
        type: 'number',
        default: 1,
    },
    hfBerlinBlock: {
        type: 'number',
        default: 0,
    },
    contractsDeployerKey: {
        type: 'string',
    },
    contractsRpcUrl: {
        type: 'string',
    },
};
const getDeployConfig = (network) => {
    let config;
    try {
        config = require(`../deploy-config/${network}.ts`).default;
    }
    catch (err) {
        throw new Error(`error while loading deploy config for network: ${network}, ${err}`);
    }
    return (0, exports.parseDeployConfig)(config);
};
exports.getDeployConfig = getDeployConfig;
const parseDeployConfig = (config) => {
    const parsed = Object.assign({}, config);
    for (const [key, spec] of Object.entries(configSpec)) {
        if (parsed[key] === undefined) {
            if ('default' in spec) {
                parsed[key] = spec.default;
            }
            else {
                throw new Error(`deploy config is missing required field: ${key} (${spec.type})`);
            }
        }
        else {
            if (spec.type === 'address') {
                if (!ethers_1.ethers.utils.isAddress(parsed[key])) {
                    throw new Error(`deploy config field: ${key} is not of type ${spec.type}: ${parsed[key]}`);
                }
            }
            else if (typeof parsed[key] !== spec.type) {
                throw new Error(`deploy config field: ${key} is not of type ${spec.type}: ${parsed[key]}`);
            }
        }
    }
    return parsed;
};
exports.parseDeployConfig = parseDeployConfig;
//# sourceMappingURL=deploy-config.js.map