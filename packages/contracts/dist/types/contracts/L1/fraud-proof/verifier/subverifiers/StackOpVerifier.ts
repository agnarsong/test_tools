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

export declare namespace VerificationContext {
  export type ContextStruct = {
    coinbase: PromiseOrValue<string>;
    timestamp: PromiseOrValue<BigNumberish>;
    number: PromiseOrValue<BigNumberish>;
    origin: PromiseOrValue<string>;
    transaction: EVMTypesLib.TransactionStruct;
    inputRoot: PromiseOrValue<BytesLike>;
    txHash: PromiseOrValue<BytesLike>;
  };

  export type ContextStructOutput = [
    string,
    BigNumber,
    BigNumber,
    string,
    EVMTypesLib.TransactionStructOutput,
    string,
    string
  ] & {
    coinbase: string;
    timestamp: BigNumber;
    number: BigNumber;
    origin: string;
    transaction: EVMTypesLib.TransactionStructOutput;
    inputRoot: string;
    txHash: string;
  };
}

export declare namespace OneStepProof {
  export type StateProofStruct = {
    blockNumber: PromiseOrValue<BigNumberish>;
    transactionIdx: PromiseOrValue<BigNumberish>;
    depth: PromiseOrValue<BigNumberish>;
    gas: PromiseOrValue<BigNumberish>;
    refund: PromiseOrValue<BigNumberish>;
    lastDepthHash: PromiseOrValue<BytesLike>;
    contractAddress: PromiseOrValue<string>;
    caller: PromiseOrValue<string>;
    value: PromiseOrValue<BigNumberish>;
    callFlag: PromiseOrValue<BigNumberish>;
    out: PromiseOrValue<BigNumberish>;
    outSize: PromiseOrValue<BigNumberish>;
    pc: PromiseOrValue<BigNumberish>;
    opCode: PromiseOrValue<BigNumberish>;
    codeHash: PromiseOrValue<BytesLike>;
    stackSize: PromiseOrValue<BigNumberish>;
    stackHash: PromiseOrValue<BytesLike>;
    memSize: PromiseOrValue<BigNumberish>;
    memRoot: PromiseOrValue<BytesLike>;
    inputDataSize: PromiseOrValue<BigNumberish>;
    inputDataRoot: PromiseOrValue<BytesLike>;
    returnDataSize: PromiseOrValue<BigNumberish>;
    returnDataRoot: PromiseOrValue<BytesLike>;
    committedGlobalStateRoot: PromiseOrValue<BytesLike>;
    globalStateRoot: PromiseOrValue<BytesLike>;
    selfDestructAcc: PromiseOrValue<BytesLike>;
    logAcc: PromiseOrValue<BytesLike>;
    blockHashRoot: PromiseOrValue<BytesLike>;
    accessListRoot: PromiseOrValue<BytesLike>;
  };

  export type StateProofStructOutput = [
    BigNumber,
    BigNumber,
    number,
    BigNumber,
    BigNumber,
    string,
    string,
    string,
    BigNumber,
    number,
    BigNumber,
    BigNumber,
    BigNumber,
    number,
    string,
    BigNumber,
    string,
    BigNumber,
    string,
    BigNumber,
    string,
    BigNumber,
    string,
    string,
    string,
    string,
    string,
    string,
    string
  ] & {
    blockNumber: BigNumber;
    transactionIdx: BigNumber;
    depth: number;
    gas: BigNumber;
    refund: BigNumber;
    lastDepthHash: string;
    contractAddress: string;
    caller: string;
    value: BigNumber;
    callFlag: number;
    out: BigNumber;
    outSize: BigNumber;
    pc: BigNumber;
    opCode: number;
    codeHash: string;
    stackSize: BigNumber;
    stackHash: string;
    memSize: BigNumber;
    memRoot: string;
    inputDataSize: BigNumber;
    inputDataRoot: string;
    returnDataSize: BigNumber;
    returnDataRoot: string;
    committedGlobalStateRoot: string;
    globalStateRoot: string;
    selfDestructAcc: string;
    logAcc: string;
    blockHashRoot: string;
    accessListRoot: string;
  };
}

export interface StackOpVerifierInterface extends utils.Interface {
  functions: {
    "executeOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32),bytes32,bytes)": FunctionFragment;
    "verifyOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32),bytes32,bytes)": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic: "executeOneStepProof" | "verifyOneStepProof"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "executeOneStepProof",
    values: [
      VerificationContext.ContextStruct,
      PromiseOrValue<BytesLike>,
      PromiseOrValue<BytesLike>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "verifyOneStepProof",
    values: [
      VerificationContext.ContextStruct,
      PromiseOrValue<BytesLike>,
      PromiseOrValue<BytesLike>
    ]
  ): string;

  decodeFunctionResult(
    functionFragment: "executeOneStepProof",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "verifyOneStepProof",
    data: BytesLike
  ): Result;

  events: {};
}

export interface StackOpVerifier extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: StackOpVerifierInterface;

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
    executeOneStepProof(
      ctx: VerificationContext.ContextStruct,
      currStateHash: PromiseOrValue<BytesLike>,
      encoded: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<[OneStepProof.StateProofStructOutput]>;

    verifyOneStepProof(
      ctx: VerificationContext.ContextStruct,
      currStateHash: PromiseOrValue<BytesLike>,
      encoded: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<[string]>;
  };

  executeOneStepProof(
    ctx: VerificationContext.ContextStruct,
    currStateHash: PromiseOrValue<BytesLike>,
    encoded: PromiseOrValue<BytesLike>,
    overrides?: CallOverrides
  ): Promise<OneStepProof.StateProofStructOutput>;

  verifyOneStepProof(
    ctx: VerificationContext.ContextStruct,
    currStateHash: PromiseOrValue<BytesLike>,
    encoded: PromiseOrValue<BytesLike>,
    overrides?: CallOverrides
  ): Promise<string>;

  callStatic: {
    executeOneStepProof(
      ctx: VerificationContext.ContextStruct,
      currStateHash: PromiseOrValue<BytesLike>,
      encoded: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<OneStepProof.StateProofStructOutput>;

    verifyOneStepProof(
      ctx: VerificationContext.ContextStruct,
      currStateHash: PromiseOrValue<BytesLike>,
      encoded: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<string>;
  };

  filters: {};

  estimateGas: {
    executeOneStepProof(
      ctx: VerificationContext.ContextStruct,
      currStateHash: PromiseOrValue<BytesLike>,
      encoded: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    verifyOneStepProof(
      ctx: VerificationContext.ContextStruct,
      currStateHash: PromiseOrValue<BytesLike>,
      encoded: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    executeOneStepProof(
      ctx: VerificationContext.ContextStruct,
      currStateHash: PromiseOrValue<BytesLike>,
      encoded: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    verifyOneStepProof(
      ctx: VerificationContext.ContextStruct,
      currStateHash: PromiseOrValue<BytesLike>,
      encoded: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;
  };
}
