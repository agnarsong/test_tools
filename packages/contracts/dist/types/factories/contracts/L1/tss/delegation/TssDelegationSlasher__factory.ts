/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import { Signer, utils, Contract, ContractFactory, Overrides } from "ethers";
import type { Provider, TransactionRequest } from "@ethersproject/providers";
import type { PromiseOrValue } from "../../../../../common";
import type {
  TssDelegationSlasher,
  TssDelegationSlasherInterface,
} from "../../../../../contracts/L1/tss/delegation/TssDelegationSlasher";

const _abi = [
  {
    inputs: [
      {
        internalType: "contract IDelegationManager",
        name: "_delegationManager",
        type: "address",
      },
      {
        internalType: "contract IDelegation",
        name: "_delegation",
        type: "address",
      },
    ],
    stateMutability: "nonpayable",
    type: "constructor",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "previouslySlashedAddress",
        type: "address",
      },
    ],
    name: "FrozenStatusReset",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "contractAdded",
        type: "address",
      },
    ],
    name: "GloballyPermissionedContractAdded",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "contractRemoved",
        type: "address",
      },
    ],
    name: "GloballyPermissionedContractRemoved",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "uint8",
        name: "version",
        type: "uint8",
      },
    ],
    name: "Initialized",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "slashedOperator",
        type: "address",
      },
      {
        indexed: true,
        internalType: "address",
        name: "slashingContract",
        type: "address",
      },
    ],
    name: "OperatorSlashed",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "operator",
        type: "address",
      },
      {
        indexed: true,
        internalType: "address",
        name: "contractAddress",
        type: "address",
      },
    ],
    name: "OptedIntoSlashing",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "previousOwner",
        type: "address",
      },
      {
        indexed: true,
        internalType: "address",
        name: "newOwner",
        type: "address",
      },
    ],
    name: "OwnershipTransferred",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "address",
        name: "account",
        type: "address",
      },
    ],
    name: "Paused",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "operator",
        type: "address",
      },
      {
        indexed: true,
        internalType: "address",
        name: "contractAddress",
        type: "address",
      },
      {
        indexed: false,
        internalType: "uint32",
        name: "unbondedAfter",
        type: "uint32",
      },
    ],
    name: "SlashingAbilityRevoked",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "address",
        name: "account",
        type: "address",
      },
    ],
    name: "Unpaused",
    type: "event",
  },
  {
    inputs: [
      {
        internalType: "address[]",
        name: "contracts",
        type: "address[]",
      },
    ],
    name: "addGloballyPermissionedContracts",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "contractAddress",
        type: "address",
      },
    ],
    name: "allowToSlash",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    name: "bondedUntil",
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
  {
    inputs: [
      {
        internalType: "address",
        name: "toBeSlashed",
        type: "address",
      },
      {
        internalType: "address",
        name: "slashingContract",
        type: "address",
      },
    ],
    name: "canSlash",
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
    name: "delegation",
    outputs: [
      {
        internalType: "contract IDelegation",
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
        internalType: "address",
        name: "toBeFrozen",
        type: "address",
      },
    ],
    name: "freezeOperator",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    name: "globallyPermissionedContracts",
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
    inputs: [
      {
        internalType: "address",
        name: "initialOwner",
        type: "address",
      },
    ],
    name: "initialize",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "investmentManager",
    outputs: [
      {
        internalType: "contract IDelegationManager",
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
        internalType: "address",
        name: "staker",
        type: "address",
      },
    ],
    name: "isFrozen",
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
    name: "owner",
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
  {
    inputs: [],
    name: "paused",
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
    inputs: [
      {
        internalType: "address[]",
        name: "contracts",
        type: "address[]",
      },
    ],
    name: "removeGloballyPermissionedContracts",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "renounceOwnership",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address[]",
        name: "frozenAddresses",
        type: "address[]",
      },
    ],
    name: "resetFrozenStatus",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "operator",
        type: "address",
      },
      {
        internalType: "uint32",
        name: "unbondedAfter",
        type: "uint32",
      },
    ],
    name: "revokeSlashingAbility",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "newOwner",
        type: "address",
      },
    ],
    name: "transferOwnership",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
] as const;

const _bytecode =
  "0x60c060405234801561001057600080fd5b50604051610f76380380610f7683398101604081905261002f91610136565b6001600160a01b03808316608052811660a052818161004c61005e565b50610057905061005e565b5050610170565b600054610100900460ff16156100ca5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101561011c576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b038116811461013357600080fd5b50565b6000806040838503121561014957600080fd5b82516101548161011e565b60208401519092506101658161011e565b809150509250929050565b60805160a051610dc56101b16000396000818161026a015281816105370152818161065b01526106fc015260008181610162015261050e0152610dc56000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c80639f114f1c116100a2578063df5cf72311610071578063df5cf72314610265578063e58398361461028c578063ec73f8081461029f578063f2fde38b146102b2578063fb3f2922146102c557600080fd5b80639f114f1c14610219578063bcdf8d141461022c578063c4d66de81461023f578063d98128c01461025257600080fd5b8063715018a6116100de578063715018a6146101a7578063755defe4146101af5780637cf72bba146101f55780638da5cb5b1461020857600080fd5b806328648d791461011057806338c8ee64146101485780634b31bb101461015d5780635c975abb1461019c575b600080fd5b61013361011e366004610c31565b60976020526000908152604090205460ff1681565b60405190151581526020015b60405180910390f35b61015b610156366004610c31565b6102d8565b005b6101847f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200161013f565b60655460ff16610133565b61015b610389565b6101e06101bd366004610c55565b609860209081526000928352604080842090915290825290205463ffffffff1681565b60405163ffffffff909116815260200161013f565b61015b610203366004610c8e565b61039d565b6033546001600160a01b0316610184565b61015b610227366004610c31565b6103ec565b61015b61023a366004610c8e565b6103f6565b61015b61024d366004610c31565b610440565b610133610260366004610c55565b6105a5565b6101847f000000000000000000000000000000000000000000000000000000000000000081565b61013361029a366004610c31565b610613565b61015b6102ad366004610c8e565b6107a2565b61015b6102c0366004610c31565b6107ec565b61015b6102d3366004610d03565b610862565b6102e061086d565b6102ea81336105a5565b61037c5760405162461bcd60e51b815260206004820152605260248201527f536c61736865722e667265657a654f70657261746f723a206d73672e73656e6460448201527f657220646f6573206e6f742068617665207065726d697373696f6e20746f20736064820152713630b9b4103a3434b99037b832b930ba37b960711b608482015260a4015b60405180910390fd5b61038681336108b3565b50565b610391610925565b61039b600061097f565b565b6103a5610925565b60005b818110156103e7576103df8383838181106103c5576103c5610d3a565b90506020020160208101906103da9190610c31565b6109d1565b6001016103a8565b505050565b6103863382610a3b565b6103fe610925565b60005b818110156103e75761043883838381811061041e5761041e610d3a565b90506020020160208101906104339190610c31565b610a9e565b600101610401565b600054610100900460ff16158080156104605750600054600160ff909116105b8061047a5750303b15801561047a575060005460ff166001145b6104dd5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610373565b6000805460ff191660011790558015610500576000805461ff0019166101001790555b6105098261097f565b6105327f0000000000000000000000000000000000000000000000000000000000000000610b08565b61055b7f0000000000000000000000000000000000000000000000000000000000000000610b08565b80156105a1576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b6001600160a01b03811660009081526097602052604081205460ff16156105ce5750600161060d565b6001600160a01b0380841660009081526098602090815260408083209386168352929052205463ffffffff164210156106095750600161060d565b5060005b92915050565b6001600160a01b03811660009081526099602052604081205460ff161561063c57506001919050565b604051633e28391d60e01b81526001600160a01b0383811660048301527f00000000000000000000000000000000000000000000000000000000000000001690633e28391d9060240160206040518083038186803b15801561069d57600080fd5b505afa1580156106b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106d59190610d50565b1561079a57604051631976849960e21b81526001600160a01b0383811660048301526000917f0000000000000000000000000000000000000000000000000000000000000000909116906365da12649060240160206040518083038186803b15801561074057600080fd5b505afa158015610754573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107789190610d72565b6001600160a01b031660009081526099602052604090205460ff169392505050565b506000919050565b6107aa610925565b60005b818110156103e7576107e48383838181106107ca576107ca610d3a565b90506020020160208101906107df9190610c31565b610b08565b6001016107ad565b6107f4610925565b6001600160a01b0381166108595760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610373565b6103868161097f565b6105a1823383610b74565b60655460ff161561039b5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610373565b6001600160a01b03821660009081526099602052604090205460ff166105a1576001600160a01b03808316600081815260996020526040808220805460ff1916600117905551928416927ff4acf64f79205932079da891a26db757fc3c07818c0d1b1f9cdab0dcc45941539190a35050565b6033546001600160a01b0316331461039b5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610373565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600160a01b03811660009081526099602052604090205460ff1615610386576001600160a01b038116600081815260996020526040808220805460ff19169055517fd4cef0af27800d466fcacd85779857378b85cb61569005ff1464fa6e5ced69d89190a250565b6001600160a01b03808316600081815260986020908152604080832094861680845294909152808220805463ffffffff191663ffffffff179055517fefa9fb38e813d53c15edf501e03852843a3fed691960523391d71a092b3627d89190a35050565b6001600160a01b03811660009081526097602052604090205460ff1615610386576001600160a01b038116600081815260976020526040808220805460ff19169055517f7a536e216e18c96837fc3462941e45c8537eecb729fe74b2c16b3e322e17d0f79190a250565b6001600160a01b03811660009081526097602052604090205460ff16610386576001600160a01b038116600081815260976020526040808220805460ff19166001179055517f255c4db16a16879fcabc9798880abb8d43877ae0c53d36e34d27409e0a6577499190a250565b6001600160a01b0380841660009081526098602090815260408083209386168352929052205463ffffffff90811614156103e7576001600160a01b03838116600081815260986020908152604080832094871680845294825291829020805463ffffffff191663ffffffff871690811790915591519182527f9aa1b1391f35c672ed1f3b7ece632f4513e618366bef7a2f67b7c6bc1f2d2b14910160405180910390a3505050565b6001600160a01b038116811461038657600080fd5b600060208284031215610c4357600080fd5b8135610c4e81610c1c565b9392505050565b60008060408385031215610c6857600080fd5b8235610c7381610c1c565b91506020830135610c8381610c1c565b809150509250929050565b60008060208385031215610ca157600080fd5b823567ffffffffffffffff80821115610cb957600080fd5b818501915085601f830112610ccd57600080fd5b813581811115610cdc57600080fd5b8660208260051b8501011115610cf157600080fd5b60209290920196919550909350505050565b60008060408385031215610d1657600080fd5b8235610d2181610c1c565b9150602083013563ffffffff81168114610c8357600080fd5b634e487b7160e01b600052603260045260246000fd5b600060208284031215610d6257600080fd5b81518015158114610c4e57600080fd5b600060208284031215610d8457600080fd5b8151610c4e81610c1c56fea26469706673582212205de67160131229859eb4c4e7370a7d295ef12cebea1c9b0dfefb51dcde9c9d1c64736f6c63430008090033";

type TssDelegationSlasherConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: TssDelegationSlasherConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class TssDelegationSlasher__factory extends ContractFactory {
  constructor(...args: TssDelegationSlasherConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override deploy(
    _delegationManager: PromiseOrValue<string>,
    _delegation: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<TssDelegationSlasher> {
    return super.deploy(
      _delegationManager,
      _delegation,
      overrides || {}
    ) as Promise<TssDelegationSlasher>;
  }
  override getDeployTransaction(
    _delegationManager: PromiseOrValue<string>,
    _delegation: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(
      _delegationManager,
      _delegation,
      overrides || {}
    );
  }
  override attach(address: string): TssDelegationSlasher {
    return super.attach(address) as TssDelegationSlasher;
  }
  override connect(signer: Signer): TssDelegationSlasher__factory {
    return super.connect(signer) as TssDelegationSlasher__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): TssDelegationSlasherInterface {
    return new utils.Interface(_abi) as TssDelegationSlasherInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): TssDelegationSlasher {
    return new Contract(
      address,
      _abi,
      signerOrProvider
    ) as TssDelegationSlasher;
  }
}