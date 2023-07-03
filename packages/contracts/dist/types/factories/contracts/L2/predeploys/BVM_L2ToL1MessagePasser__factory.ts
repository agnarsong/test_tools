/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import { Signer, utils, Contract, ContractFactory, Overrides } from "ethers";
import type { Provider, TransactionRequest } from "@ethersproject/providers";
import type { PromiseOrValue } from "../../../../common";
import type {
  BVM_L2ToL1MessagePasser,
  BVM_L2ToL1MessagePasserInterface,
} from "../../../../contracts/L2/predeploys/BVM_L2ToL1MessagePasser";

const _abi = [
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "uint256",
        name: "_nonce",
        type: "uint256",
      },
      {
        indexed: false,
        internalType: "address",
        name: "_sender",
        type: "address",
      },
      {
        indexed: false,
        internalType: "bytes",
        name: "_data",
        type: "bytes",
      },
    ],
    name: "L2ToL1Message",
    type: "event",
  },
  {
    inputs: [
      {
        internalType: "bytes",
        name: "_message",
        type: "bytes",
      },
    ],
    name: "passMessageToL1",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "bytes32",
        name: "",
        type: "bytes32",
      },
    ],
    name: "sentMessages",
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
] as const;

const _bytecode =
  "0x608060405234801561001057600080fd5b50610242806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806382e3702d1461003b578063cafa81dc14610072575b600080fd5b61005e6100493660046100d6565b60006020819052908152604090205460ff1681565b604051901515815260200160405180910390f35b610085610080366004610105565b610087565b005b6001600080833360405160200161009f9291906101b6565b60408051808303601f19018152918152815160209283012083529082019290925201600020805460ff191691151591909117905550565b6000602082840312156100e857600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b60006020828403121561011757600080fd5b813567ffffffffffffffff8082111561012f57600080fd5b818401915084601f83011261014357600080fd5b813581811115610155576101556100ef565b604051601f8201601f19908116603f0116810190838211818310171561017d5761017d6100ef565b8160405282815287602084870101111561019657600080fd5b826020860160208301376000928101602001929092525095945050505050565b6000835160005b818110156101d757602081870181015185830152016101bd565b818111156101e6576000828501525b5060609390931b6bffffffffffffffffffffffff1916919092019081526014019291505056fea2646970667358221220a5c338b6c39fb5f5d42de33bacb945608db4a507585b3f40c0324e4402967d6c64736f6c63430008090033";

type BVM_L2ToL1MessagePasserConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: BVM_L2ToL1MessagePasserConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class BVM_L2ToL1MessagePasser__factory extends ContractFactory {
  constructor(...args: BVM_L2ToL1MessagePasserConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override deploy(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<BVM_L2ToL1MessagePasser> {
    return super.deploy(overrides || {}) as Promise<BVM_L2ToL1MessagePasser>;
  }
  override getDeployTransaction(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(overrides || {});
  }
  override attach(address: string): BVM_L2ToL1MessagePasser {
    return super.attach(address) as BVM_L2ToL1MessagePasser;
  }
  override connect(signer: Signer): BVM_L2ToL1MessagePasser__factory {
    return super.connect(signer) as BVM_L2ToL1MessagePasser__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): BVM_L2ToL1MessagePasserInterface {
    return new utils.Interface(_abi) as BVM_L2ToL1MessagePasserInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): BVM_L2ToL1MessagePasser {
    return new Contract(
      address,
      _abi,
      signerOrProvider
    ) as BVM_L2ToL1MessagePasser;
  }
}