"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.MessageReceiptStatus = exports.MessageDirection = exports.MessageStatus = exports.L2ChainID = exports.L1ChainID = void 0;
var L1ChainID;
(function (L1ChainID) {
    L1ChainID[L1ChainID["MAINNET"] = 10] = "MAINNET";
    L1ChainID[L1ChainID["GOERLI"] = 5] = "GOERLI";
    L1ChainID[L1ChainID["KOVAN"] = 42] = "KOVAN";
    L1ChainID[L1ChainID["HARDHAT_LOCAL"] = 31337] = "HARDHAT_LOCAL";
})(L1ChainID = exports.L1ChainID || (exports.L1ChainID = {}));
var L2ChainID;
(function (L2ChainID) {
    L2ChainID[L2ChainID["MANTLE"] = 10] = "MANTLE";
    L2ChainID[L2ChainID["MANTLE_TESTNET"] = 5001] = "MANTLE_TESTNET";
    L2ChainID[L2ChainID["MANTLE_GOERLIQA"] = 1705003] = "MANTLE_GOERLIQA";
    L2ChainID[L2ChainID["MANTLE_KOVAN"] = 69] = "MANTLE_KOVAN";
    L2ChainID[L2ChainID["MANTLE_HARDHAT_LOCAL"] = 31337] = "MANTLE_HARDHAT_LOCAL";
    L2ChainID[L2ChainID["MANTLE_HARDHAT_DEVNET"] = 17] = "MANTLE_HARDHAT_DEVNET";
})(L2ChainID = exports.L2ChainID || (exports.L2ChainID = {}));
var MessageStatus;
(function (MessageStatus) {
    MessageStatus[MessageStatus["UNCONFIRMED_L1_TO_L2_MESSAGE"] = 0] = "UNCONFIRMED_L1_TO_L2_MESSAGE";
    MessageStatus[MessageStatus["FAILED_L1_TO_L2_MESSAGE"] = 1] = "FAILED_L1_TO_L2_MESSAGE";
    MessageStatus[MessageStatus["STATE_ROOT_NOT_PUBLISHED"] = 2] = "STATE_ROOT_NOT_PUBLISHED";
    MessageStatus[MessageStatus["IN_CHALLENGE_PERIOD"] = 3] = "IN_CHALLENGE_PERIOD";
    MessageStatus[MessageStatus["READY_FOR_RELAY"] = 4] = "READY_FOR_RELAY";
    MessageStatus[MessageStatus["RELAYED"] = 5] = "RELAYED";
})(MessageStatus = exports.MessageStatus || (exports.MessageStatus = {}));
var MessageDirection;
(function (MessageDirection) {
    MessageDirection[MessageDirection["L1_TO_L2"] = 0] = "L1_TO_L2";
    MessageDirection[MessageDirection["L2_TO_L1"] = 1] = "L2_TO_L1";
})(MessageDirection = exports.MessageDirection || (exports.MessageDirection = {}));
var MessageReceiptStatus;
(function (MessageReceiptStatus) {
    MessageReceiptStatus[MessageReceiptStatus["RELAYED_FAILED"] = 0] = "RELAYED_FAILED";
    MessageReceiptStatus[MessageReceiptStatus["RELAYED_SUCCEEDED"] = 1] = "RELAYED_SUCCEEDED";
})(MessageReceiptStatus = exports.MessageReceiptStatus || (exports.MessageReceiptStatus = {}));
//# sourceMappingURL=types.js.map