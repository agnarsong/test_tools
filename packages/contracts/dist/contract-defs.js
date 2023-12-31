"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.getContractFactory = exports.getContractInterface = exports.getContractDefinition = void 0;
const ethers_1 = require("ethers");
const getContractDefinition = (name) => {
    const { getContractArtifact } = require('./contract-artifacts');
    const artifact = getContractArtifact(name);
    if (artifact === undefined) {
        throw new Error(`Unable to find artifact for contract: ${name}`);
    }
    return artifact;
};
exports.getContractDefinition = getContractDefinition;
const getContractInterface = (name) => {
    const definition = (0, exports.getContractDefinition)(name);
    return new ethers_1.ethers.utils.Interface(definition.abi);
};
exports.getContractInterface = getContractInterface;
const getContractFactory = (name, signer) => {
    const definition = (0, exports.getContractDefinition)(name);
    const contractInterface = (0, exports.getContractInterface)(name);
    return new ethers_1.ethers.ContractFactory(contractInterface, definition.bytecode, signer);
};
exports.getContractFactory = getContractFactory;
//# sourceMappingURL=contract-defs.js.map