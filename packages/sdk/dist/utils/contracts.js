"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.getBridgeAdapters = exports.getAllOEContracts = exports.getOEContract = exports.BRIDGE_ADAPTER_DATA = exports.CONTRACT_ADDRESSES = exports.DEFAULT_L2_CONTRACT_ADDRESSES = void 0;
const contracts_1 = require("@mantleio/contracts");
const ethers_1 = require("ethers");
const coercion_1 = require("./coercion");
const interfaces_1 = require("../interfaces");
const adapters_1 = require("../adapters");
exports.DEFAULT_L2_CONTRACT_ADDRESSES = {
    L2CrossDomainMessenger: contracts_1.predeploys.L2CrossDomainMessenger,
    L2StandardBridge: contracts_1.predeploys.L2StandardBridge,
    BVM_L1BlockNumber: contracts_1.predeploys.BVM_L1BlockNumber,
    BVM_L2ToL1MessagePasser: contracts_1.predeploys.BVM_L2ToL1MessagePasser,
    BVM_DeployerWhitelist: contracts_1.predeploys.BVM_DeployerWhitelist,
    BVM_ETH: contracts_1.predeploys.BVM_ETH,
    BVM_MANTLE: contracts_1.predeploys.BVM_MANTLE,
    BVM_GasPriceOracle: contracts_1.predeploys.BVM_GasPriceOracle,
    BVM_SequencerFeeVault: contracts_1.predeploys.BVM_SequencerFeeVault,
    WETH: contracts_1.predeploys.WETH9,
    TssRewardContract: contracts_1.predeploys.TssRewardContract,
};
const NAME_REMAPPING = {
    AddressManager: 'Lib_AddressManager',
    BVM_L1BlockNumber: 'iBVM_L1BlockNumber',
    WETH: 'WETH9',
};
exports.CONTRACT_ADDRESSES = {
    [interfaces_1.L2ChainID.MANTLE]: {
        l1: {
            AddressManager: '0x6968f3F16C3e64003F02E121cf0D5CCBf5625a42',
            L1CrossDomainMessenger: '0x676A795fe6E43C17c668de16730c3F690FEB7120',
            L1StandardBridge: '0x95fC37A27a2f68e3A647CDc081F0A89bb47c3012',
            StateCommitmentChain: '0x89E9D387555AF0cDE22cb98833Bae40d640AD7fa',
            CanonicalTransactionChain: '0x291dc3819b863e19b0a9b9809F8025d2EB4aaE93',
            BondManager: '0x31aBe1c466C2A8b95fd84258dD1471472979B650',
            Rollup: process.env.Rollup ||
                '0xD1328C9167e0693B689b5aa5a024379d4e437858',
        },
        l2: exports.DEFAULT_L2_CONTRACT_ADDRESSES,
    },
    [interfaces_1.L2ChainID.MANTLE_KOVAN]: {
        l1: {
            AddressManager: '0x100Dd3b414Df5BbA2B542864fF94aF8024aFdf3a',
            L1CrossDomainMessenger: '0x4361d0F75A0186C05f971c566dC6bEa5957483fD',
            L1StandardBridge: '0x22F24361D548e5FaAfb36d1437839f080363982B',
            StateCommitmentChain: '0xD7754711773489F31A0602635f3F167826ce53C5',
            CanonicalTransactionChain: '0xf7B88A133202d41Fe5E2Ab22e6309a1A4D50AF74',
            BondManager: '0xc5a603d273E28185c18Ba4d26A0024B2d2F42740',
            Rollup: process.env.Rollup ||
                '0x9faB987C9C469EB23Da31B7848B28aCf30905eA8',
        },
        l2: exports.DEFAULT_L2_CONTRACT_ADDRESSES,
    },
    [interfaces_1.L2ChainID.MANTLE_GOERLIQA]: {
        l1: {
            AddressManager: process.env.ADDRESS_MANAGER_ADDRESS ||
                '0x327903410307971Ca7Ba8A6CB2291D3b8825d7F5',
            L1CrossDomainMessenger: process.env.L1_CROSS_DOMAIN_MESSENGER_ADDRESS ||
                '0x3f41DAcb2dB659e45826126d004ad3E0C8eA680e',
            L1StandardBridge: process.env.L1_STANDARD_BRIDGE_ADDRESS ||
                '0x4cf99b9BC9B2Da64033D1Fb65146Ea60fbe8AD4B',
            StateCommitmentChain: process.env.STATE_COMMITMENT_CHAIN_ADDRESS ||
                '0x88EC574e2ef0EcF9043373139099f7E535F94dBC',
            CanonicalTransactionChain: process.env.CANONICAL_TRANSACTION_CHAIN_ADDRESS ||
                '0x258e80D5371fD7fFdDFE29E60b366f9FC44844c8',
            BondManager: process.env.BOND_MANAGER_ADDRESS ||
                '0xc723Cb5f3337c2F6Eab9b29E78CE42a28B8661d1',
            Rollup: process.env.Rollup ||
                '0x9faB987C9C469EB23Da31B7848B28aCf30905eA8',
        },
        l2: exports.DEFAULT_L2_CONTRACT_ADDRESSES,
    },
    [interfaces_1.L2ChainID.MANTLE_TESTNET]: {
        l1: {
            AddressManager: '0xA647F5947C50248bc4b2eF773791c9C2bc01C65A',
            L1CrossDomainMessenger: '0x7Bfe603647d5380ED3909F6f87580D0Af1B228B4',
            L1StandardBridge: '0xc92470D7Ffa21473611ab6c6e2FcFB8637c8f330',
            StateCommitmentChain: '0x91A5D806BA73d0AA4bFA9B318126dDE60582e92a',
            CanonicalTransactionChain: '0x654e6dF111F98374d9e5d908D7a5392C308aA18D',
            BondManager: '0xeBE3f28BbFa7bB8f2C066C1A792073203B985e27',
            Rollup: process.env.Rollup ||
                '0x9faB987C9C469EB23Da31B7848B28aCf30905eA8',
        },
        l2: exports.DEFAULT_L2_CONTRACT_ADDRESSES,
    },
    [interfaces_1.L2ChainID.MANTLE_HARDHAT_LOCAL]: {
        l1: {
            AddressManager: process.env.ADDRESS_MANAGER_ADDRESS ||
                '0x92aBAD50368175785e4270ca9eFd169c949C4ce1',
            L1CrossDomainMessenger: process.env.L1_CROSS_DOMAIN_MESSENGER_ADDRESS ||
                '0x7959CF3b8ffC87Faca8aD8a1B5D95c0f58C0BEf8',
            L1StandardBridge: process.env.L1_STANDARD_BRIDGE_ADDRESS ||
                '0x8BAccFF561FDe61D6bC8B6f299fFBa561d2189B9',
            StateCommitmentChain: process.env.STATE_COMMITMENT_CHAIN_ADDRESS ||
                '0xd9e2F450525079e1e29fB23Bc7Caca6F61f8fD4a',
            CanonicalTransactionChain: process.env.CANONICAL_TRANSACTION_CHAIN_ADDRESS ||
                '0x0090171f848B2aa86918E5Ef2406Ab3d424fdd83',
            BondManager: process.env.BOND_MANAGER_ADDRESS ||
                '0x9faB987C9C469EB23Da31B7848B28aCf30905eA8',
            Rollup: process.env.Rollup ||
                '0x9faB987C9C469EB23Da31B7848B28aCf30905eA8',
        },
        l2: exports.DEFAULT_L2_CONTRACT_ADDRESSES,
    },
    [interfaces_1.L2ChainID.MANTLE_HARDHAT_DEVNET]: {
        l1: {
            AddressManager: process.env.ADDRESS_MANAGER_ADDRESS ||
                '0x92aBAD50368175785e4270ca9eFd169c949C4ce1',
            L1CrossDomainMessenger: process.env.L1_CROSS_DOMAIN_MESSENGER_ADDRESS ||
                '0x7959CF3b8ffC87Faca8aD8a1B5D95c0f58C0BEf8',
            L1StandardBridge: process.env.L1_STANDARD_BRIDGE_ADDRESS ||
                '0x8BAccFF561FDe61D6bC8B6f299fFBa561d2189B9',
            StateCommitmentChain: process.env.STATE_COMMITMENT_CHAIN_ADDRESS ||
                '0xd9e2F450525079e1e29fB23Bc7Caca6F61f8fD4a',
            CanonicalTransactionChain: process.env.CANONICAL_TRANSACTION_CHAIN_ADDRESS ||
                '0x0090171f848B2aa86918E5Ef2406Ab3d424fdd83',
            BondManager: process.env.BOND_MANAGER_ADDRESS ||
                '0x9faB987C9C469EB23Da31B7848B28aCf30905eA8',
            Rollup: process.env.Rollup ||
                '0x9faB987C9C469EB23Da31B7848B28aCf30905eA8',
        },
        l2: exports.DEFAULT_L2_CONTRACT_ADDRESSES,
    },
};
exports.BRIDGE_ADAPTER_DATA = {
    [interfaces_1.L2ChainID.MANTLE]: {
        BitBTC: {
            Adapter: adapters_1.StandardBridgeAdapter,
            l1Bridge: '0xaBA2c5F108F7E820C049D5Af70B16ac266c8f128',
            l2Bridge: '0x158F513096923fF2d3aab2BcF4478536de6725e2',
        },
        DAI: {
            Adapter: adapters_1.ERC20BridgeAdapter,
            l1Bridge: '0x10E6593CDda8c58a1d0f14C5164B376352a55f2F',
            l2Bridge: '0x467194771dAe2967Aef3ECbEDD3Bf9a310C76C65',
        },
    },
    [interfaces_1.L2ChainID.MANTLE_KOVAN]: {
        wstETH: {
            Adapter: adapters_1.ERC20BridgeAdapter,
            l1Bridge: '0xa88751C0a08623E11ff38c6B70F2BbEe7865C17c',
            l2Bridge: '0xF9C842dE4381a70eB265d10CF8D43DceFF5bA935',
        },
        BitBTC: {
            Adapter: adapters_1.StandardBridgeAdapter,
            l1Bridge: '0x0b651A42F32069d62d5ECf4f2a7e5Bd3E9438746',
            l2Bridge: '0x0CFb46528a7002a7D8877a5F7a69b9AaF1A9058e',
        },
        USX: {
            Adapter: adapters_1.StandardBridgeAdapter,
            l1Bridge: '0x40E862341b2416345F02c41Ac70df08525150dC7',
            l2Bridge: '0xB4d37826b14Cd3CB7257A2A5094507d701fe715f',
        },
        DAI: {
            Adapter: adapters_1.ERC20BridgeAdapter,
            l1Bridge: '0xb415e822C4983ecD6B1c1596e8a5f976cf6CD9e3',
            l2Bridge: '0x467194771dAe2967Aef3ECbEDD3Bf9a310C76C65',
        },
    },
};
const getOEContract = (contractName, l2ChainId, opts = {}) => {
    const addresses = exports.CONTRACT_ADDRESSES[l2ChainId];
    if (addresses === undefined && opts.address === undefined) {
        throw new Error(`cannot get contract ${contractName} for unknown L2 chain ID ${l2ChainId}, you must provide an address`);
    }
    const name = NAME_REMAPPING[contractName] || contractName;
    let iface;
    iface = (0, contracts_1.getContractInterface)(name);
    return new ethers_1.Contract((0, coercion_1.toAddress)(opts.address || addresses.l1[contractName] || addresses.l2[contractName]), iface, opts.signerOrProvider);
};
exports.getOEContract = getOEContract;
const getAllOEContracts = (l2ChainId, opts = {}) => {
    var _a, _b, _c, _d;
    const addresses = exports.CONTRACT_ADDRESSES[l2ChainId] || {
        l1: {
            AddressManager: undefined,
            L1CrossDomainMessenger: undefined,
            L1StandardBridge: undefined,
            StateCommitmentChain: undefined,
            CanonicalTransactionChain: undefined,
            BondManager: undefined,
        },
        l2: exports.DEFAULT_L2_CONTRACT_ADDRESSES,
    };
    const l1Contracts = {};
    for (const [contractName, contractAddress] of Object.entries(addresses.l1)) {
        l1Contracts[contractName] = (0, exports.getOEContract)(contractName, l2ChainId, {
            address: ((_b = (_a = opts.overrides) === null || _a === void 0 ? void 0 : _a.l1) === null || _b === void 0 ? void 0 : _b[contractName]) || contractAddress,
            signerOrProvider: opts.l1SignerOrProvider,
        });
    }
    const l2Contracts = {};
    for (const [contractName, contractAddress] of Object.entries(addresses.l2)) {
        l2Contracts[contractName] = (0, exports.getOEContract)(contractName, l2ChainId, {
            address: ((_d = (_c = opts.overrides) === null || _c === void 0 ? void 0 : _c.l2) === null || _d === void 0 ? void 0 : _d[contractName]) || contractAddress,
            signerOrProvider: opts.l2SignerOrProvider,
        });
    }
    return {
        l1: l1Contracts,
        l2: l2Contracts,
    };
};
exports.getAllOEContracts = getAllOEContracts;
const getBridgeAdapters = (l2ChainId, messenger, opts) => {
    const adapterData = Object.assign(Object.assign(Object.assign({}, (exports.CONTRACT_ADDRESSES[l2ChainId]
        ? {
            Standard: {
                Adapter: adapters_1.StandardBridgeAdapter,
                l1Bridge: exports.CONTRACT_ADDRESSES[l2ChainId].l1.L1StandardBridge,
                l2Bridge: contracts_1.predeploys.L2StandardBridge,
            },
            ETH: {
                Adapter: adapters_1.ETHBridgeAdapter,
                l1Bridge: exports.CONTRACT_ADDRESSES[l2ChainId].l1.L1StandardBridge,
                l2Bridge: contracts_1.predeploys.L2StandardBridge,
            },
        }
        : {})), (exports.BRIDGE_ADAPTER_DATA[l2ChainId] || {})), ((opts === null || opts === void 0 ? void 0 : opts.overrides) || {}));
    const adapters = {};
    for (const [bridgeName, bridgeData] of Object.entries(adapterData)) {
        adapters[bridgeName] = new bridgeData.Adapter({
            messenger,
            l1Bridge: bridgeData.l1Bridge,
            l2Bridge: bridgeData.l2Bridge,
        });
    }
    return adapters;
};
exports.getBridgeAdapters = getBridgeAdapters;
//# sourceMappingURL=contracts.js.map