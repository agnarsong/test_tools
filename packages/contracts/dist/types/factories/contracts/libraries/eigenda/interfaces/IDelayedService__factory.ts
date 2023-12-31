/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Signer, utils } from "ethers";
import type { Provider } from "@ethersproject/providers";
import type {
  IDelayedService,
  IDelayedServiceInterface,
} from "../../../../../contracts/libraries/eigenda/interfaces/IDelayedService";

const _abi = [
  {
    inputs: [],
    name: "BLOCK_STALE_MEASURE",
    outputs: [
      {
        internalType: "uint32",
        name: "",
        type: "uint32",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
] as const;

export class IDelayedService__factory {
  static readonly abi = _abi;
  static createInterface(): IDelayedServiceInterface {
    return new utils.Interface(_abi) as IDelayedServiceInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): IDelayedService {
    return new Contract(address, _abi, signerOrProvider) as IDelayedService;
  }
}
