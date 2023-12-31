/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import { Signer, utils, Contract, ContractFactory, Overrides } from "ethers";
import type { Provider, TransactionRequest } from "@ethersproject/providers";
import type { PromiseOrValue } from "../../../common";
import type {
  EmptyContract,
  EmptyContractInterface,
} from "../../../contracts/test-helpers/EmptyContract";

const _abi = [
  {
    inputs: [],
    name: "foo",
    outputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    stateMutability: "pure",
    type: "function",
  },
] as const;

const _bytecode =
  "0x6080604052348015600f57600080fd5b50607780601d6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063c298557814602d575b600080fd5b600060405190815260200160405180910390f3fea2646970667358221220b72a61a19498172c16e3dd47e923813a716c2b1749272efe87dc903ee958638864736f6c63430008090033";

type EmptyContractConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: EmptyContractConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class EmptyContract__factory extends ContractFactory {
  constructor(...args: EmptyContractConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override deploy(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<EmptyContract> {
    return super.deploy(overrides || {}) as Promise<EmptyContract>;
  }
  override getDeployTransaction(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(overrides || {});
  }
  override attach(address: string): EmptyContract {
    return super.attach(address) as EmptyContract;
  }
  override connect(signer: Signer): EmptyContract__factory {
    return super.connect(signer) as EmptyContract__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): EmptyContractInterface {
    return new utils.Interface(_abi) as EmptyContractInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): EmptyContract {
    return new Contract(address, _abi, signerOrProvider) as EmptyContract;
  }
}
