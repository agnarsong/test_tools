"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.CHAIN_BLOCK_TIMES = exports.DEPOSIT_CONFIRMATION_BLOCKS = void 0;
const interfaces_1 = require("../interfaces");
exports.DEPOSIT_CONFIRMATION_BLOCKS = {
    [interfaces_1.L2ChainID.MANTLE]: 50,
    [interfaces_1.L2ChainID.MANTLE_KOVAN]: 12,
    [interfaces_1.L2ChainID.MANTLE_GOERLIQA]: 12,
    [interfaces_1.L2ChainID.MANTLE_TESTNET]: 12,
    [interfaces_1.L2ChainID.MANTLE_HARDHAT_LOCAL]: 2,
    [interfaces_1.L2ChainID.MANTLE_HARDHAT_DEVNET]: 2,
};
exports.CHAIN_BLOCK_TIMES = {
    [interfaces_1.L1ChainID.MAINNET]: 13,
    [interfaces_1.L1ChainID.GOERLI]: 15,
    [interfaces_1.L1ChainID.KOVAN]: 4,
    [interfaces_1.L1ChainID.HARDHAT_LOCAL]: 1,
};
//# sourceMappingURL=chain-constants.js.map