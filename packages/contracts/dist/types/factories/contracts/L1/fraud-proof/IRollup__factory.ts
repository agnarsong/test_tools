/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Signer, utils } from "ethers";
import type { Provider } from "@ethersproject/providers";
import type {
  IRollup,
  IRollupInterface,
} from "../../../../contracts/L1/fraud-proof/IRollup";

const _abi = [
  {
    inputs: [],
    name: "AssertionAlreadyResolved",
    type: "error",
  },
  {
    inputs: [],
    name: "AssertionOutOfRange",
    type: "error",
  },
  {
    inputs: [],
    name: "ChallengePeriodPending",
    type: "error",
  },
  {
    inputs: [],
    name: "ChallengedStaker",
    type: "error",
  },
  {
    inputs: [],
    name: "DifferentParent",
    type: "error",
  },
  {
    inputs: [],
    name: "EmptyAssertion",
    type: "error",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "staker1Challenge",
        type: "address",
      },
      {
        internalType: "address",
        name: "staker2Challenge",
        type: "address",
      },
    ],
    name: "InDifferentChallenge",
    type: "error",
  },
  {
    inputs: [],
    name: "InboxReadLimitExceeded",
    type: "error",
  },
  {
    inputs: [],
    name: "InsufficientStake",
    type: "error",
  },
  {
    inputs: [],
    name: "InvalidParent",
    type: "error",
  },
  {
    inputs: [],
    name: "MinimumAssertionPeriodNotPassed",
    type: "error",
  },
  {
    inputs: [],
    name: "NoStaker",
    type: "error",
  },
  {
    inputs: [],
    name: "NoUnresolvedAssertion",
    type: "error",
  },
  {
    inputs: [],
    name: "NotAllStaked",
    type: "error",
  },
  {
    inputs: [],
    name: "NotInChallenge",
    type: "error",
  },
  {
    inputs: [],
    name: "NotStaked",
    type: "error",
  },
  {
    inputs: [],
    name: "ParentAssertionUnstaked",
    type: "error",
  },
  {
    inputs: [],
    name: "PreviousStateHash",
    type: "error",
  },
  {
    inputs: [],
    name: "RedundantInitialized",
    type: "error",
  },
  {
    inputs: [],
    name: "StakedOnUnconfirmedAssertion",
    type: "error",
  },
  {
    inputs: [],
    name: "StakerStakedOnTarget",
    type: "error",
  },
  {
    inputs: [],
    name: "StakersPresent",
    type: "error",
  },
  {
    inputs: [],
    name: "TransferFailed",
    type: "error",
  },
  {
    inputs: [],
    name: "UnproposedAssertion",
    type: "error",
  },
  {
    inputs: [],
    name: "WrongOrder",
    type: "error",
  },
  {
    inputs: [],
    name: "ZeroAddress",
    type: "error",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "uint256",
        name: "assertionID",
        type: "uint256",
      },
      {
        indexed: false,
        internalType: "address",
        name: "challengeAddr",
        type: "address",
      },
    ],
    name: "AssertionChallenged",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "uint256",
        name: "assertionID",
        type: "uint256",
      },
    ],
    name: "AssertionConfirmed",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "uint256",
        name: "assertionID",
        type: "uint256",
      },
      {
        indexed: false,
        internalType: "address",
        name: "asserterAddr",
        type: "address",
      },
      {
        indexed: false,
        internalType: "bytes32",
        name: "vmHash",
        type: "bytes32",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "inboxSize",
        type: "uint256",
      },
    ],
    name: "AssertionCreated",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "uint256",
        name: "assertionID",
        type: "uint256",
      },
    ],
    name: "AssertionRejected",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "address",
        name: "stakerAddr",
        type: "address",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "assertionID",
        type: "uint256",
      },
    ],
    name: "StakerStaked",
    type: "event",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "assertionID",
        type: "uint256",
      },
    ],
    name: "advanceStake",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "assertions",
    outputs: [
      {
        internalType: "contract AssertionMap",
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
        internalType: "address[2]",
        name: "players",
        type: "address[2]",
      },
      {
        internalType: "uint256[2]",
        name: "assertionIDs",
        type: "uint256[2]",
      },
    ],
    name: "challengeAssertion",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "winner",
        type: "address",
      },
      {
        internalType: "address",
        name: "loser",
        type: "address",
      },
    ],
    name: "completeChallenge",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "confirmFirstUnresolvedAssertion",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "confirmedInboxSize",
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
  {
    inputs: [
      {
        internalType: "bytes32",
        name: "vmHash",
        type: "bytes32",
      },
      {
        internalType: "uint256",
        name: "inboxSize",
        type: "uint256",
      },
    ],
    name: "createAssertion",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "bytes32",
        name: "vmHash",
        type: "bytes32",
      },
      {
        internalType: "uint256",
        name: "inboxSize",
        type: "uint256",
      },
      {
        internalType: "bytes32[]",
        name: "_batch",
        type: "bytes32[]",
      },
      {
        internalType: "uint256",
        name: "_shouldStartAtElement",
        type: "uint256",
      },
      {
        internalType: "bytes",
        name: "_signature",
        type: "bytes",
      },
    ],
    name: "createAssertionWithStateBatch",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "currentRequiredStake",
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
  {
    inputs: [
      {
        internalType: "address",
        name: "addr",
        type: "address",
      },
    ],
    name: "isStaked",
    outputs: [
      {
        internalType: "bool",
        name: "",
        type: "bool",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "rejectFirstUnresolvedAssertion",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        components: [
          {
            internalType: "uint256",
            name: "batchIndex",
            type: "uint256",
          },
          {
            internalType: "bytes32",
            name: "batchRoot",
            type: "bytes32",
          },
          {
            internalType: "uint256",
            name: "batchSize",
            type: "uint256",
          },
          {
            internalType: "uint256",
            name: "prevTotalElements",
            type: "uint256",
          },
          {
            internalType: "bytes",
            name: "signature",
            type: "bytes",
          },
          {
            internalType: "bytes",
            name: "extraData",
            type: "bytes",
          },
        ],
        internalType: "struct Lib_BVMCodec.ChainBatchHeader",
        name: "_batchHeader",
        type: "tuple",
      },
    ],
    name: "rejectLatestCreatedAssertionWithBatch",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "stakerAddress",
        type: "address",
      },
    ],
    name: "removeStake",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "stakeAmount",
        type: "uint256",
      },
      {
        internalType: "address",
        name: "operator",
        type: "address",
      },
    ],
    name: "stake",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "stakeAmount",
        type: "uint256",
      },
    ],
    name: "unstake",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "withdraw",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
] as const;

export class IRollup__factory {
  static readonly abi = _abi;
  static createInterface(): IRollupInterface {
    return new utils.Interface(_abi) as IRollupInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): IRollup {
    return new Contract(address, _abi, signerOrProvider) as IRollup;
  }
}
