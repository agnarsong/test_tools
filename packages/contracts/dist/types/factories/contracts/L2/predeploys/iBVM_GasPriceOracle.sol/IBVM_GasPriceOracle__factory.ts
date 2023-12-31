/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Signer, utils } from "ethers";
import type { Provider } from "@ethersproject/providers";
import type {
  IBVM_GasPriceOracle,
  IBVM_GasPriceOracleInterface,
} from "../../../../../contracts/L2/predeploys/iBVM_GasPriceOracle.sol/IBVM_GasPriceOracle";

const _abi = [
  {
    inputs: [],
    name: "IsBurning",
    outputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
] as const;

export class IBVM_GasPriceOracle__factory {
  static readonly abi = _abi;
  static createInterface(): IBVM_GasPriceOracleInterface {
    return new utils.Interface(_abi) as IBVM_GasPriceOracleInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): IBVM_GasPriceOracle {
    return new Contract(address, _abi, signerOrProvider) as IBVM_GasPriceOracle;
  }
}
