/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Signer, utils } from "ethers";
import type { Provider } from "@ethersproject/providers";
import type {
  Lib_AddressResolver,
  Lib_AddressResolverInterface,
} from "../../../../contracts/libraries/resolver/Lib_AddressResolver";

const _abi = [
  {
    inputs: [],
    name: "libAddressManager",
    outputs: [
      {
        internalType: "contract Lib_AddressManager",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "string",
        name: "_name",
        type: "string",
      },
    ],
    name: "resolve",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
] as const;

export class Lib_AddressResolver__factory {
  static readonly abi = _abi;
  static createInterface(): Lib_AddressResolverInterface {
    return new utils.Interface(_abi) as Lib_AddressResolverInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): Lib_AddressResolver {
    return new Contract(address, _abi, signerOrProvider) as Lib_AddressResolver;
  }
}
