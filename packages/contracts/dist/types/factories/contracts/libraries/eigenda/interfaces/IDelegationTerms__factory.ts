/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Signer, utils } from "ethers";
import type { Provider } from "@ethersproject/providers";
import type {
  IDelegationTerms,
  IDelegationTermsInterface,
} from "../../../../../contracts/libraries/eigenda/interfaces/IDelegationTerms";

const _abi = [
  {
    inputs: [
      {
        internalType: "address",
        name: "delegator",
        type: "address",
      },
      {
        internalType: "contract IInvestmentStrategy[]",
        name: "investorStrats",
        type: "address[]",
      },
      {
        internalType: "uint256[]",
        name: "investorShares",
        type: "uint256[]",
      },
    ],
    name: "onDelegationReceived",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "delegator",
        type: "address",
      },
      {
        internalType: "contract IInvestmentStrategy[]",
        name: "investorStrats",
        type: "address[]",
      },
      {
        internalType: "uint256[]",
        name: "investorShares",
        type: "uint256[]",
      },
    ],
    name: "onDelegationWithdrawn",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "contract IERC20",
        name: "token",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "amount",
        type: "uint256",
      },
    ],
    name: "payForService",
    outputs: [],
    stateMutability: "payable",
    type: "function",
  },
] as const;

export class IDelegationTerms__factory {
  static readonly abi = _abi;
  static createInterface(): IDelegationTermsInterface {
    return new utils.Interface(_abi) as IDelegationTermsInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): IDelegationTerms {
    return new Contract(address, _abi, signerOrProvider) as IDelegationTerms;
  }
}
