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
} from "../../../common";

export declare namespace Lib_BVMCodec {
  export type ChainBatchHeaderStruct = {
    batchIndex: PromiseOrValue<BigNumberish>;
    batchRoot: PromiseOrValue<BytesLike>;
    batchSize: PromiseOrValue<BigNumberish>;
    prevTotalElements: PromiseOrValue<BigNumberish>;
    signature: PromiseOrValue<BytesLike>;
    extraData: PromiseOrValue<BytesLike>;
  };

  export type ChainBatchHeaderStructOutput = [
    BigNumber,
    string,
    BigNumber,
    BigNumber,
    string,
    string
  ] & {
    batchIndex: BigNumber;
    batchRoot: string;
    batchSize: BigNumber;
    prevTotalElements: BigNumber;
    signature: string;
    extraData: string;
  };

  export type ChainInclusionProofStruct = {
    index: PromiseOrValue<BigNumberish>;
    siblings: PromiseOrValue<BytesLike>[];
  };

  export type ChainInclusionProofStructOutput = [BigNumber, string[]] & {
    index: BigNumber;
    siblings: string[];
  };
}

export declare namespace IL1CrossDomainMessenger {
  export type L2MessageInclusionProofStruct = {
    stateRoot: PromiseOrValue<BytesLike>;
    stateRootBatchHeader: Lib_BVMCodec.ChainBatchHeaderStruct;
    stateRootProof: Lib_BVMCodec.ChainInclusionProofStruct;
    stateTrieWitness: PromiseOrValue<BytesLike>;
    storageTrieWitness: PromiseOrValue<BytesLike>;
  };

  export type L2MessageInclusionProofStructOutput = [
    string,
    Lib_BVMCodec.ChainBatchHeaderStructOutput,
    Lib_BVMCodec.ChainInclusionProofStructOutput,
    string,
    string
  ] & {
    stateRoot: string;
    stateRootBatchHeader: Lib_BVMCodec.ChainBatchHeaderStructOutput;
    stateRootProof: Lib_BVMCodec.ChainInclusionProofStructOutput;
    stateTrieWitness: string;
    storageTrieWitness: string;
  };
}

export interface L1CrossDomainMessengerInterface extends utils.Interface {
  functions: {
    "allowMessage(bytes32)": FunctionFragment;
    "blockMessage(bytes32)": FunctionFragment;
    "blockedMessages(bytes32)": FunctionFragment;
    "getPauseOwner()": FunctionFragment;
    "initialize(address)": FunctionFragment;
    "libAddressManager()": FunctionFragment;
    "owner()": FunctionFragment;
    "pause()": FunctionFragment;
    "pauseByPOwner()": FunctionFragment;
    "paused()": FunctionFragment;
    "relayMessage(address,address,bytes,uint256,(bytes32,(uint256,bytes32,uint256,uint256,bytes,bytes),(uint256,bytes32[]),bytes,bytes))": FunctionFragment;
    "relayedMessages(bytes32)": FunctionFragment;
    "renounceOwnership()": FunctionFragment;
    "replayMessage(address,address,bytes,uint256,uint32,uint32)": FunctionFragment;
    "resolve(string)": FunctionFragment;
    "sendMessage(address,bytes,uint32)": FunctionFragment;
    "setPauseOwner(address)": FunctionFragment;
    "successfulMessages(bytes32)": FunctionFragment;
    "transferOwnership(address)": FunctionFragment;
    "unpause()": FunctionFragment;
    "xDomainMessageSender()": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic:
      | "allowMessage"
      | "blockMessage"
      | "blockedMessages"
      | "getPauseOwner"
      | "initialize"
      | "libAddressManager"
      | "owner"
      | "pause"
      | "pauseByPOwner"
      | "paused"
      | "relayMessage"
      | "relayedMessages"
      | "renounceOwnership"
      | "replayMessage"
      | "resolve"
      | "sendMessage"
      | "setPauseOwner"
      | "successfulMessages"
      | "transferOwnership"
      | "unpause"
      | "xDomainMessageSender"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "allowMessage",
    values: [PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "blockMessage",
    values: [PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "blockedMessages",
    values: [PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "getPauseOwner",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "initialize",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "libAddressManager",
    values?: undefined
  ): string;
  encodeFunctionData(functionFragment: "owner", values?: undefined): string;
  encodeFunctionData(functionFragment: "pause", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "pauseByPOwner",
    values?: undefined
  ): string;
  encodeFunctionData(functionFragment: "paused", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "relayMessage",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<string>,
      PromiseOrValue<BytesLike>,
      PromiseOrValue<BigNumberish>,
      IL1CrossDomainMessenger.L2MessageInclusionProofStruct
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "relayedMessages",
    values: [PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "renounceOwnership",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "replayMessage",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<string>,
      PromiseOrValue<BytesLike>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BigNumberish>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "resolve",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "sendMessage",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<BytesLike>,
      PromiseOrValue<BigNumberish>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "setPauseOwner",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "successfulMessages",
    values: [PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "transferOwnership",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(functionFragment: "unpause", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "xDomainMessageSender",
    values?: undefined
  ): string;

  decodeFunctionResult(
    functionFragment: "allowMessage",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "blockMessage",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "blockedMessages",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getPauseOwner",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "initialize", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "libAddressManager",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "owner", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "pause", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "pauseByPOwner",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "paused", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "relayMessage",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "relayedMessages",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "renounceOwnership",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "replayMessage",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "resolve", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "sendMessage",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "setPauseOwner",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "successfulMessages",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "transferOwnership",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "unpause", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "xDomainMessageSender",
    data: BytesLike
  ): Result;

  events: {
    "FailedRelayedMessage(bytes32)": EventFragment;
    "Initialized(uint8)": EventFragment;
    "MessageAllowed(bytes32)": EventFragment;
    "MessageBlocked(bytes32)": EventFragment;
    "OwnershipTransferred(address,address)": EventFragment;
    "Paused(address)": EventFragment;
    "RelayedMessage(bytes32)": EventFragment;
    "SentMessage(address,address,bytes,uint256,uint256)": EventFragment;
    "Unpaused(address)": EventFragment;
  };

  getEvent(nameOrSignatureOrTopic: "FailedRelayedMessage"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "Initialized"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "MessageAllowed"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "MessageBlocked"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "OwnershipTransferred"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "Paused"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "RelayedMessage"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "SentMessage"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "Unpaused"): EventFragment;
}

export interface FailedRelayedMessageEventObject {
  msgHash: string;
}
export type FailedRelayedMessageEvent = TypedEvent<
  [string],
  FailedRelayedMessageEventObject
>;

export type FailedRelayedMessageEventFilter =
  TypedEventFilter<FailedRelayedMessageEvent>;

export interface InitializedEventObject {
  version: number;
}
export type InitializedEvent = TypedEvent<[number], InitializedEventObject>;

export type InitializedEventFilter = TypedEventFilter<InitializedEvent>;

export interface MessageAllowedEventObject {
  _xDomainCalldataHash: string;
}
export type MessageAllowedEvent = TypedEvent<
  [string],
  MessageAllowedEventObject
>;

export type MessageAllowedEventFilter = TypedEventFilter<MessageAllowedEvent>;

export interface MessageBlockedEventObject {
  _xDomainCalldataHash: string;
}
export type MessageBlockedEvent = TypedEvent<
  [string],
  MessageBlockedEventObject
>;

export type MessageBlockedEventFilter = TypedEventFilter<MessageBlockedEvent>;

export interface OwnershipTransferredEventObject {
  previousOwner: string;
  newOwner: string;
}
export type OwnershipTransferredEvent = TypedEvent<
  [string, string],
  OwnershipTransferredEventObject
>;

export type OwnershipTransferredEventFilter =
  TypedEventFilter<OwnershipTransferredEvent>;

export interface PausedEventObject {
  account: string;
}
export type PausedEvent = TypedEvent<[string], PausedEventObject>;

export type PausedEventFilter = TypedEventFilter<PausedEvent>;

export interface RelayedMessageEventObject {
  msgHash: string;
}
export type RelayedMessageEvent = TypedEvent<
  [string],
  RelayedMessageEventObject
>;

export type RelayedMessageEventFilter = TypedEventFilter<RelayedMessageEvent>;

export interface SentMessageEventObject {
  target: string;
  sender: string;
  message: string;
  messageNonce: BigNumber;
  gasLimit: BigNumber;
}
export type SentMessageEvent = TypedEvent<
  [string, string, string, BigNumber, BigNumber],
  SentMessageEventObject
>;

export type SentMessageEventFilter = TypedEventFilter<SentMessageEvent>;

export interface UnpausedEventObject {
  account: string;
}
export type UnpausedEvent = TypedEvent<[string], UnpausedEventObject>;

export type UnpausedEventFilter = TypedEventFilter<UnpausedEvent>;

export interface L1CrossDomainMessenger extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: L1CrossDomainMessengerInterface;

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
    allowMessage(
      _xDomainCalldataHash: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    blockMessage(
      _xDomainCalldataHash: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    blockedMessages(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<[boolean]>;

    getPauseOwner(overrides?: CallOverrides): Promise<[string]>;

    initialize(
      _libAddressManager: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    libAddressManager(overrides?: CallOverrides): Promise<[string]>;

    owner(overrides?: CallOverrides): Promise<[string]>;

    pause(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    pauseByPOwner(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    paused(overrides?: CallOverrides): Promise<[boolean]>;

    relayMessage(
      _target: PromiseOrValue<string>,
      _sender: PromiseOrValue<string>,
      _message: PromiseOrValue<BytesLike>,
      _messageNonce: PromiseOrValue<BigNumberish>,
      _proof: IL1CrossDomainMessenger.L2MessageInclusionProofStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    relayedMessages(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<[boolean]>;

    renounceOwnership(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    replayMessage(
      _target: PromiseOrValue<string>,
      _sender: PromiseOrValue<string>,
      _message: PromiseOrValue<BytesLike>,
      _queueIndex: PromiseOrValue<BigNumberish>,
      _oldGasLimit: PromiseOrValue<BigNumberish>,
      _newGasLimit: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    resolve(
      _name: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<[string]>;

    sendMessage(
      _target: PromiseOrValue<string>,
      _message: PromiseOrValue<BytesLike>,
      _gasLimit: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    setPauseOwner(
      _pauseOwner: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    successfulMessages(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<[boolean]>;

    transferOwnership(
      newOwner: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    unpause(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    xDomainMessageSender(overrides?: CallOverrides): Promise<[string]>;
  };

  allowMessage(
    _xDomainCalldataHash: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  blockMessage(
    _xDomainCalldataHash: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  blockedMessages(
    arg0: PromiseOrValue<BytesLike>,
    overrides?: CallOverrides
  ): Promise<boolean>;

  getPauseOwner(overrides?: CallOverrides): Promise<string>;

  initialize(
    _libAddressManager: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  libAddressManager(overrides?: CallOverrides): Promise<string>;

  owner(overrides?: CallOverrides): Promise<string>;

  pause(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  pauseByPOwner(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  paused(overrides?: CallOverrides): Promise<boolean>;

  relayMessage(
    _target: PromiseOrValue<string>,
    _sender: PromiseOrValue<string>,
    _message: PromiseOrValue<BytesLike>,
    _messageNonce: PromiseOrValue<BigNumberish>,
    _proof: IL1CrossDomainMessenger.L2MessageInclusionProofStruct,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  relayedMessages(
    arg0: PromiseOrValue<BytesLike>,
    overrides?: CallOverrides
  ): Promise<boolean>;

  renounceOwnership(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  replayMessage(
    _target: PromiseOrValue<string>,
    _sender: PromiseOrValue<string>,
    _message: PromiseOrValue<BytesLike>,
    _queueIndex: PromiseOrValue<BigNumberish>,
    _oldGasLimit: PromiseOrValue<BigNumberish>,
    _newGasLimit: PromiseOrValue<BigNumberish>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  resolve(
    _name: PromiseOrValue<string>,
    overrides?: CallOverrides
  ): Promise<string>;

  sendMessage(
    _target: PromiseOrValue<string>,
    _message: PromiseOrValue<BytesLike>,
    _gasLimit: PromiseOrValue<BigNumberish>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  setPauseOwner(
    _pauseOwner: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  successfulMessages(
    arg0: PromiseOrValue<BytesLike>,
    overrides?: CallOverrides
  ): Promise<boolean>;

  transferOwnership(
    newOwner: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  unpause(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  xDomainMessageSender(overrides?: CallOverrides): Promise<string>;

  callStatic: {
    allowMessage(
      _xDomainCalldataHash: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;

    blockMessage(
      _xDomainCalldataHash: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;

    blockedMessages(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<boolean>;

    getPauseOwner(overrides?: CallOverrides): Promise<string>;

    initialize(
      _libAddressManager: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;

    libAddressManager(overrides?: CallOverrides): Promise<string>;

    owner(overrides?: CallOverrides): Promise<string>;

    pause(overrides?: CallOverrides): Promise<void>;

    pauseByPOwner(overrides?: CallOverrides): Promise<void>;

    paused(overrides?: CallOverrides): Promise<boolean>;

    relayMessage(
      _target: PromiseOrValue<string>,
      _sender: PromiseOrValue<string>,
      _message: PromiseOrValue<BytesLike>,
      _messageNonce: PromiseOrValue<BigNumberish>,
      _proof: IL1CrossDomainMessenger.L2MessageInclusionProofStruct,
      overrides?: CallOverrides
    ): Promise<void>;

    relayedMessages(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<boolean>;

    renounceOwnership(overrides?: CallOverrides): Promise<void>;

    replayMessage(
      _target: PromiseOrValue<string>,
      _sender: PromiseOrValue<string>,
      _message: PromiseOrValue<BytesLike>,
      _queueIndex: PromiseOrValue<BigNumberish>,
      _oldGasLimit: PromiseOrValue<BigNumberish>,
      _newGasLimit: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<void>;

    resolve(
      _name: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<string>;

    sendMessage(
      _target: PromiseOrValue<string>,
      _message: PromiseOrValue<BytesLike>,
      _gasLimit: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<void>;

    setPauseOwner(
      _pauseOwner: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;

    successfulMessages(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<boolean>;

    transferOwnership(
      newOwner: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;

    unpause(overrides?: CallOverrides): Promise<void>;

    xDomainMessageSender(overrides?: CallOverrides): Promise<string>;
  };

  filters: {
    "FailedRelayedMessage(bytes32)"(
      msgHash?: PromiseOrValue<BytesLike> | null
    ): FailedRelayedMessageEventFilter;
    FailedRelayedMessage(
      msgHash?: PromiseOrValue<BytesLike> | null
    ): FailedRelayedMessageEventFilter;

    "Initialized(uint8)"(version?: null): InitializedEventFilter;
    Initialized(version?: null): InitializedEventFilter;

    "MessageAllowed(bytes32)"(
      _xDomainCalldataHash?: PromiseOrValue<BytesLike> | null
    ): MessageAllowedEventFilter;
    MessageAllowed(
      _xDomainCalldataHash?: PromiseOrValue<BytesLike> | null
    ): MessageAllowedEventFilter;

    "MessageBlocked(bytes32)"(
      _xDomainCalldataHash?: PromiseOrValue<BytesLike> | null
    ): MessageBlockedEventFilter;
    MessageBlocked(
      _xDomainCalldataHash?: PromiseOrValue<BytesLike> | null
    ): MessageBlockedEventFilter;

    "OwnershipTransferred(address,address)"(
      previousOwner?: PromiseOrValue<string> | null,
      newOwner?: PromiseOrValue<string> | null
    ): OwnershipTransferredEventFilter;
    OwnershipTransferred(
      previousOwner?: PromiseOrValue<string> | null,
      newOwner?: PromiseOrValue<string> | null
    ): OwnershipTransferredEventFilter;

    "Paused(address)"(account?: null): PausedEventFilter;
    Paused(account?: null): PausedEventFilter;

    "RelayedMessage(bytes32)"(
      msgHash?: PromiseOrValue<BytesLike> | null
    ): RelayedMessageEventFilter;
    RelayedMessage(
      msgHash?: PromiseOrValue<BytesLike> | null
    ): RelayedMessageEventFilter;

    "SentMessage(address,address,bytes,uint256,uint256)"(
      target?: PromiseOrValue<string> | null,
      sender?: null,
      message?: null,
      messageNonce?: null,
      gasLimit?: null
    ): SentMessageEventFilter;
    SentMessage(
      target?: PromiseOrValue<string> | null,
      sender?: null,
      message?: null,
      messageNonce?: null,
      gasLimit?: null
    ): SentMessageEventFilter;

    "Unpaused(address)"(account?: null): UnpausedEventFilter;
    Unpaused(account?: null): UnpausedEventFilter;
  };

  estimateGas: {
    allowMessage(
      _xDomainCalldataHash: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    blockMessage(
      _xDomainCalldataHash: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    blockedMessages(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    getPauseOwner(overrides?: CallOverrides): Promise<BigNumber>;

    initialize(
      _libAddressManager: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    libAddressManager(overrides?: CallOverrides): Promise<BigNumber>;

    owner(overrides?: CallOverrides): Promise<BigNumber>;

    pause(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    pauseByPOwner(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    paused(overrides?: CallOverrides): Promise<BigNumber>;

    relayMessage(
      _target: PromiseOrValue<string>,
      _sender: PromiseOrValue<string>,
      _message: PromiseOrValue<BytesLike>,
      _messageNonce: PromiseOrValue<BigNumberish>,
      _proof: IL1CrossDomainMessenger.L2MessageInclusionProofStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    relayedMessages(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    renounceOwnership(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    replayMessage(
      _target: PromiseOrValue<string>,
      _sender: PromiseOrValue<string>,
      _message: PromiseOrValue<BytesLike>,
      _queueIndex: PromiseOrValue<BigNumberish>,
      _oldGasLimit: PromiseOrValue<BigNumberish>,
      _newGasLimit: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    resolve(
      _name: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    sendMessage(
      _target: PromiseOrValue<string>,
      _message: PromiseOrValue<BytesLike>,
      _gasLimit: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    setPauseOwner(
      _pauseOwner: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    successfulMessages(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    transferOwnership(
      newOwner: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    unpause(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    xDomainMessageSender(overrides?: CallOverrides): Promise<BigNumber>;
  };

  populateTransaction: {
    allowMessage(
      _xDomainCalldataHash: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    blockMessage(
      _xDomainCalldataHash: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    blockedMessages(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    getPauseOwner(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    initialize(
      _libAddressManager: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    libAddressManager(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    owner(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    pause(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    pauseByPOwner(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    paused(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    relayMessage(
      _target: PromiseOrValue<string>,
      _sender: PromiseOrValue<string>,
      _message: PromiseOrValue<BytesLike>,
      _messageNonce: PromiseOrValue<BigNumberish>,
      _proof: IL1CrossDomainMessenger.L2MessageInclusionProofStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    relayedMessages(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    renounceOwnership(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    replayMessage(
      _target: PromiseOrValue<string>,
      _sender: PromiseOrValue<string>,
      _message: PromiseOrValue<BytesLike>,
      _queueIndex: PromiseOrValue<BigNumberish>,
      _oldGasLimit: PromiseOrValue<BigNumberish>,
      _newGasLimit: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    resolve(
      _name: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    sendMessage(
      _target: PromiseOrValue<string>,
      _message: PromiseOrValue<BytesLike>,
      _gasLimit: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    setPauseOwner(
      _pauseOwner: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    successfulMessages(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    transferOwnership(
      newOwner: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    unpause(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    xDomainMessageSender(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;
  };
}
