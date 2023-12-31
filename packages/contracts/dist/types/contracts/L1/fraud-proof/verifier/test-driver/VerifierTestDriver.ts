/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
  BigNumber,
  BigNumberish,
  BytesLike,
  CallOverrides,
  PopulatedTransaction,
  Signer,
  utils,
} from "ethers";
import type { FunctionFragment, Result } from "@ethersproject/abi";
import type { Listener, Provider } from "@ethersproject/providers";
import type {
  TypedEventFilter,
  TypedEvent,
  TypedListener,
  OnEvent,
  PromiseOrValue,
} from "../../../../../common";

export declare namespace EVMTypesLib {
  export type TransactionStruct = {
    nonce: PromiseOrValue<BigNumberish>;
    gasPrice: PromiseOrValue<BigNumberish>;
    gas: PromiseOrValue<BigNumberish>;
    to: PromiseOrValue<string>;
    value: PromiseOrValue<BigNumberish>;
    data: PromiseOrValue<BytesLike>;
    v: PromiseOrValue<BigNumberish>;
    r: PromiseOrValue<BigNumberish>;
    s: PromiseOrValue<BigNumberish>;
  };

  export type TransactionStructOutput = [
    BigNumber,
    BigNumber,
    BigNumber,
    string,
    BigNumber,
    string,
    BigNumber,
    BigNumber,
    BigNumber
  ] & {
    nonce: BigNumber;
    gasPrice: BigNumber;
    gas: BigNumber;
    to: string;
    value: BigNumber;
    data: string;
    v: BigNumber;
    r: BigNumber;
    s: BigNumber;
  };
}

export interface VerifierTestDriverInterface extends utils.Interface {
  functions: {
    "verifyProof(address,uint256,uint256,address,bytes32,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),uint8,bytes32,bytes)": FunctionFragment;
  };

  getFunction(nameOrSignatureOrTopic: "verifyProof"): FunctionFragment;

  encodeFunctionData(
    functionFragment: "verifyProof",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<string>,
      PromiseOrValue<BytesLike>,
      EVMTypesLib.TransactionStruct,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BytesLike>,
      PromiseOrValue<BytesLike>
    ]
  ): string;

  decodeFunctionResult(
    functionFragment: "verifyProof",
    data: BytesLike
  ): Result;

  events: {};
}

export interface VerifierTestDriver extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: VerifierTestDriverInterface;

  queryFilter<TEvent extends TypedEvent>(
    event: TypedEventFilter<TEvent>,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TEvent>>;

  listeners<TEvent extends TypedEvent>(
    eventFilter?: TypedEventFilter<TEvent>
  ): Array<TypedListener<TEvent>>;
  listeners(eventName?: string): Array<Listener>;
  removeAllListeners<TEvent extends TypedEvent>(
    eventFilter: TypedEventFilter<TEvent>
  ): this;
  removeAllListeners(eventName?: string): this;
  off: OnEvent<this>;
  on: OnEvent<this>;
  once: OnEvent<this>;
  removeListener: OnEvent<this>;

  functions: {
    verifyProof(
      sequencerAddress: PromiseOrValue<string>,
      timestamp: PromiseOrValue<BigNumberish>,
      number: PromiseOrValue<BigNumberish>,
      origin: PromiseOrValue<string>,
      txHash: PromiseOrValue<BytesLike>,
      transaction: EVMTypesLib.TransactionStruct,
      verifier: PromiseOrValue<BigNumberish>,
      currStateHash: PromiseOrValue<BytesLike>,
      proof: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<[string]>;
  };

  verifyProof(
    sequencerAddress: PromiseOrValue<string>,
    timestamp: PromiseOrValue<BigNumberish>,
    number: PromiseOrValue<BigNumberish>,
    origin: PromiseOrValue<string>,
    txHash: PromiseOrValue<BytesLike>,
    transaction: EVMTypesLib.TransactionStruct,
    verifier: PromiseOrValue<BigNumberish>,
    currStateHash: PromiseOrValue<BytesLike>,
    proof: PromiseOrValue<BytesLike>,
    overrides?: CallOverrides
  ): Promise<string>;

  callStatic: {
    verifyProof(
      sequencerAddress: PromiseOrValue<string>,
      timestamp: PromiseOrValue<BigNumberish>,
      number: PromiseOrValue<BigNumberish>,
      origin: PromiseOrValue<string>,
      txHash: PromiseOrValue<BytesLike>,
      transaction: EVMTypesLib.TransactionStruct,
      verifier: PromiseOrValue<BigNumberish>,
      currStateHash: PromiseOrValue<BytesLike>,
      proof: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<string>;
  };

  filters: {};

  estimateGas: {
    verifyProof(
      sequencerAddress: PromiseOrValue<string>,
      timestamp: PromiseOrValue<BigNumberish>,
      number: PromiseOrValue<BigNumberish>,
      origin: PromiseOrValue<string>,
      txHash: PromiseOrValue<BytesLike>,
      transaction: EVMTypesLib.TransactionStruct,
      verifier: PromiseOrValue<BigNumberish>,
      currStateHash: PromiseOrValue<BytesLike>,
      proof: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    verifyProof(
      sequencerAddress: PromiseOrValue<string>,
      timestamp: PromiseOrValue<BigNumberish>,
      number: PromiseOrValue<BigNumberish>,
      origin: PromiseOrValue<string>,
      txHash: PromiseOrValue<BytesLike>,
      transaction: EVMTypesLib.TransactionStruct,
      verifier: PromiseOrValue<BigNumberish>,
      currStateHash: PromiseOrValue<BytesLike>,
      proof: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;
  };
}
