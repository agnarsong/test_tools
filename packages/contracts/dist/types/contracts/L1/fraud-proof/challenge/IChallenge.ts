/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
  BigNumber,
  BigNumberish,
  BytesLike,
  CallOverrides,
  ContractTransaction,
  Overrides,
  PopulatedTransaction,
  Signer,
  utils,
} from "ethers";
import type {
  FunctionFragment,
  Result,
  EventFragment,
} from "@ethersproject/abi";
import type { Listener, Provider } from "@ethersproject/providers";
import type {
  TypedEventFilter,
  TypedEvent,
  TypedListener,
  OnEvent,
  PromiseOrValue,
} from "../../../../common";

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

export interface IChallengeInterface extends utils.Interface {
  functions: {
    "bisectExecution(bytes32[3],uint256,uint256,uint256,uint256,uint256)": FunctionFragment;
    "completeChallenge(bool)": FunctionFragment;
    "currentResponder()": FunctionFragment;
    "currentResponderTimeLeft()": FunctionFragment;
    "initialize(address,address,address,address,uint256,bytes32,bytes32)": FunctionFragment;
    "initializeChallengeLength(bytes32,uint256)": FunctionFragment;
    "setRollback()": FunctionFragment;
    "timeout()": FunctionFragment;
    "verifyOneStepProof((address,uint256,uint256,address,(uint64,uint256,uint64,address,uint256,bytes,uint256,uint256,uint256),bytes32,bytes32),uint8,bytes,uint256,uint256,uint256)": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic:
      | "bisectExecution"
      | "completeChallenge"
      | "currentResponder"
      | "currentResponderTimeLeft"
      | "initialize"
      | "initializeChallengeLength"
      | "setRollback"
      | "timeout"
      | "verifyOneStepProof"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "bisectExecution",
    values: [
      [
        PromiseOrValue<BytesLike>,
        PromiseOrValue<BytesLike>,
        PromiseOrValue<BytesLike>
      ],
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BigNumberish>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "completeChallenge",
    values: [PromiseOrValue<boolean>]
  ): string;
  encodeFunctionData(
    functionFragment: "currentResponder",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "currentResponderTimeLeft",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "initialize",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<string>,
      PromiseOrValue<string>,
      PromiseOrValue<string>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BytesLike>,
      PromiseOrValue<BytesLike>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "initializeChallengeLength",
    values: [PromiseOrValue<BytesLike>, PromiseOrValue<BigNumberish>]
  ): string;
  encodeFunctionData(
    functionFragment: "setRollback",
    values?: undefined
  ): string;
  encodeFunctionData(functionFragment: "timeout", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "verifyOneStepProof",
    values: [
      VerificationContext.ContextStruct,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BytesLike>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BigNumberish>
    ]
  ): string;

  decodeFunctionResult(
    functionFragment: "bisectExecution",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "completeChallenge",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "currentResponder",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "currentResponderTimeLeft",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "initialize", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "initializeChallengeLength",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "setRollback",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "timeout", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "verifyOneStepProof",
    data: BytesLike
  ): Result;

  events: {
    "Bisected(bytes32,bytes32,bytes32,uint256,uint256,uint256,uint256)": EventFragment;
    "ChallengeCompleted(address,address,uint8)": EventFragment;
  };

  getEvent(nameOrSignatureOrTopic: "Bisected"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "ChallengeCompleted"): EventFragment;
}

export interface BisectedEventObject {
  startState: string;
  midState: string;
  endState: string;
  blockNum: BigNumber;
  blockTime: BigNumber;
  challengedSegmentStart: BigNumber;
  challengedSegmentLength: BigNumber;
}
export type BisectedEvent = TypedEvent<
  [string, string, string, BigNumber, BigNumber, BigNumber, BigNumber],
  BisectedEventObject
>;

export type BisectedEventFilter = TypedEventFilter<BisectedEvent>;

export interface ChallengeCompletedEventObject {
  winner: string;
  loser: string;
  reason: number;
}
export type ChallengeCompletedEvent = TypedEvent<
  [string, string, number],
  ChallengeCompletedEventObject
>;

export type ChallengeCompletedEventFilter =
  TypedEventFilter<ChallengeCompletedEvent>;

export interface IChallenge extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: IChallengeInterface;

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
    bisectExecution(
      bisection: [
        PromiseOrValue<BytesLike>,
        PromiseOrValue<BytesLike>,
        PromiseOrValue<BytesLike>
      ],
      challengedSegmentIndex: PromiseOrValue<BigNumberish>,
      challengedSegmentStart: PromiseOrValue<BigNumberish>,
      challengedSegmentLength: PromiseOrValue<BigNumberish>,
      prevChallengedSegmentStart: PromiseOrValue<BigNumberish>,
      prevChallengedSegmentLength: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    completeChallenge(
      arg0: PromiseOrValue<boolean>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    currentResponder(overrides?: CallOverrides): Promise<[string]>;

    currentResponderTimeLeft(overrides?: CallOverrides): Promise<[BigNumber]>;

    initialize(
      _defender: PromiseOrValue<string>,
      _challenger: PromiseOrValue<string>,
      _verifier: PromiseOrValue<string>,
      _resultReceiver: PromiseOrValue<string>,
      _startInboxSize: PromiseOrValue<BigNumberish>,
      _startStateHash: PromiseOrValue<BytesLike>,
      _endStateHash: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    initializeChallengeLength(
      checkStateHash: PromiseOrValue<BytesLike>,
      _numSteps: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    setRollback(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    timeout(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    verifyOneStepProof(
      ctx: VerificationContext.ContextStruct,
      verifyType: PromiseOrValue<BigNumberish>,
      proof: PromiseOrValue<BytesLike>,
      challengedStepIndex: PromiseOrValue<BigNumberish>,
      prevChallengedSegmentStart: PromiseOrValue<BigNumberish>,
      prevChallengedSegmentLength: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;
  };

  bisectExecution(
    bisection: [
      PromiseOrValue<BytesLike>,
      PromiseOrValue<BytesLike>,
      PromiseOrValue<BytesLike>
    ],
    challengedSegmentIndex: PromiseOrValue<BigNumberish>,
    challengedSegmentStart: PromiseOrValue<BigNumberish>,
    challengedSegmentLength: PromiseOrValue<BigNumberish>,
    prevChallengedSegmentStart: PromiseOrValue<BigNumberish>,
    prevChallengedSegmentLength: PromiseOrValue<BigNumberish>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  completeChallenge(
    arg0: PromiseOrValue<boolean>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  currentResponder(overrides?: CallOverrides): Promise<string>;

  currentResponderTimeLeft(overrides?: CallOverrides): Promise<BigNumber>;

  initialize(
    _defender: PromiseOrValue<string>,
    _challenger: PromiseOrValue<string>,
    _verifier: PromiseOrValue<string>,
    _resultReceiver: PromiseOrValue<string>,
    _startInboxSize: PromiseOrValue<BigNumberish>,
    _startStateHash: PromiseOrValue<BytesLike>,
    _endStateHash: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  initializeChallengeLength(
    checkStateHash: PromiseOrValue<BytesLike>,
    _numSteps: PromiseOrValue<BigNumberish>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  setRollback(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  timeout(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  verifyOneStepProof(
    ctx: VerificationContext.ContextStruct,
    verifyType: PromiseOrValue<BigNumberish>,
    proof: PromiseOrValue<BytesLike>,
    challengedStepIndex: PromiseOrValue<BigNumberish>,
    prevChallengedSegmentStart: PromiseOrValue<BigNumberish>,
    prevChallengedSegmentLength: PromiseOrValue<BigNumberish>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  callStatic: {
    bisectExecution(
      bisection: [
        PromiseOrValue<BytesLike>,
        PromiseOrValue<BytesLike>,
        PromiseOrValue<BytesLike>
      ],
      challengedSegmentIndex: PromiseOrValue<BigNumberish>,
      challengedSegmentStart: PromiseOrValue<BigNumberish>,
      challengedSegmentLength: PromiseOrValue<BigNumberish>,
      prevChallengedSegmentStart: PromiseOrValue<BigNumberish>,
      prevChallengedSegmentLength: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<void>;

    completeChallenge(
      arg0: PromiseOrValue<boolean>,
      overrides?: CallOverrides
    ): Promise<void>;

    currentResponder(overrides?: CallOverrides): Promise<string>;

    currentResponderTimeLeft(overrides?: CallOverrides): Promise<BigNumber>;

    initialize(
      _defender: PromiseOrValue<string>,
      _challenger: PromiseOrValue<string>,
      _verifier: PromiseOrValue<string>,
      _resultReceiver: PromiseOrValue<string>,
      _startInboxSize: PromiseOrValue<BigNumberish>,
      _startStateHash: PromiseOrValue<BytesLike>,
      _endStateHash: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;

    initializeChallengeLength(
      checkStateHash: PromiseOrValue<BytesLike>,
      _numSteps: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<void>;

    setRollback(overrides?: CallOverrides): Promise<void>;

    timeout(overrides?: CallOverrides): Promise<void>;

    verifyOneStepProof(
      ctx: VerificationContext.ContextStruct,
      verifyType: PromiseOrValue<BigNumberish>,
      proof: PromiseOrValue<BytesLike>,
      challengedStepIndex: PromiseOrValue<BigNumberish>,
      prevChallengedSegmentStart: PromiseOrValue<BigNumberish>,
      prevChallengedSegmentLength: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<void>;
  };

  filters: {
    "Bisected(bytes32,bytes32,bytes32,uint256,uint256,uint256,uint256)"(
      startState?: null,
      midState?: null,
      endState?: null,
      blockNum?: null,
      blockTime?: null,
      challengedSegmentStart?: null,
      challengedSegmentLength?: null
    ): BisectedEventFilter;
    Bisected(
      startState?: null,
      midState?: null,
      endState?: null,
      blockNum?: null,
      blockTime?: null,
      challengedSegmentStart?: null,
      challengedSegmentLength?: null
    ): BisectedEventFilter;

    "ChallengeCompleted(address,address,uint8)"(
      winner?: null,
      loser?: null,
      reason?: null
    ): ChallengeCompletedEventFilter;
    ChallengeCompleted(
      winner?: null,
      loser?: null,
      reason?: null
    ): ChallengeCompletedEventFilter;
  };

  estimateGas: {
    bisectExecution(
      bisection: [
        PromiseOrValue<BytesLike>,
        PromiseOrValue<BytesLike>,
        PromiseOrValue<BytesLike>
      ],
      challengedSegmentIndex: PromiseOrValue<BigNumberish>,
      challengedSegmentStart: PromiseOrValue<BigNumberish>,
      challengedSegmentLength: PromiseOrValue<BigNumberish>,
      prevChallengedSegmentStart: PromiseOrValue<BigNumberish>,
      prevChallengedSegmentLength: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    completeChallenge(
      arg0: PromiseOrValue<boolean>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    currentResponder(overrides?: CallOverrides): Promise<BigNumber>;

    currentResponderTimeLeft(overrides?: CallOverrides): Promise<BigNumber>;

    initialize(
      _defender: PromiseOrValue<string>,
      _challenger: PromiseOrValue<string>,
      _verifier: PromiseOrValue<string>,
      _resultReceiver: PromiseOrValue<string>,
      _startInboxSize: PromiseOrValue<BigNumberish>,
      _startStateHash: PromiseOrValue<BytesLike>,
      _endStateHash: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    initializeChallengeLength(
      checkStateHash: PromiseOrValue<BytesLike>,
      _numSteps: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    setRollback(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    timeout(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    verifyOneStepProof(
      ctx: VerificationContext.ContextStruct,
      verifyType: PromiseOrValue<BigNumberish>,
      proof: PromiseOrValue<BytesLike>,
      challengedStepIndex: PromiseOrValue<BigNumberish>,
      prevChallengedSegmentStart: PromiseOrValue<BigNumberish>,
      prevChallengedSegmentLength: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    bisectExecution(
      bisection: [
        PromiseOrValue<BytesLike>,
        PromiseOrValue<BytesLike>,
        PromiseOrValue<BytesLike>
      ],
      challengedSegmentIndex: PromiseOrValue<BigNumberish>,
      challengedSegmentStart: PromiseOrValue<BigNumberish>,
      challengedSegmentLength: PromiseOrValue<BigNumberish>,
      prevChallengedSegmentStart: PromiseOrValue<BigNumberish>,
      prevChallengedSegmentLength: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    completeChallenge(
      arg0: PromiseOrValue<boolean>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    currentResponder(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    currentResponderTimeLeft(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    initialize(
      _defender: PromiseOrValue<string>,
      _challenger: PromiseOrValue<string>,
      _verifier: PromiseOrValue<string>,
      _resultReceiver: PromiseOrValue<string>,
      _startInboxSize: PromiseOrValue<BigNumberish>,
      _startStateHash: PromiseOrValue<BytesLike>,
      _endStateHash: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    initializeChallengeLength(
      checkStateHash: PromiseOrValue<BytesLike>,
      _numSteps: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    setRollback(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    timeout(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    verifyOneStepProof(
      ctx: VerificationContext.ContextStruct,
      verifyType: PromiseOrValue<BigNumberish>,
      proof: PromiseOrValue<BytesLike>,
      challengedStepIndex: PromiseOrValue<BigNumberish>,
      prevChallengedSegmentStart: PromiseOrValue<BigNumberish>,
      prevChallengedSegmentLength: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;
  };
}
