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

export interface L2CrossDomainMessengerInterface extends utils.Interface {
  functions: {
    "l1CrossDomainMessenger()": FunctionFragment;
    "messageNonce()": FunctionFragment;
    "relayMessage(address,address,bytes,uint256)": FunctionFragment;
    "relayedMessages(bytes32)": FunctionFragment;
    "sendMessage(address,bytes,uint32)": FunctionFragment;
    "sentMessages(bytes32)": FunctionFragment;
    "successfulMessages(bytes32)": FunctionFragment;
    "xDomainMessageSender()": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic:
      | "l1CrossDomainMessenger"
      | "messageNonce"
      | "relayMessage"
      | "relayedMessages"
      | "sendMessage"
      | "sentMessages"
      | "successfulMessages"
      | "xDomainMessageSender"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "l1CrossDomainMessenger",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "messageNonce",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "relayMessage",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<string>,
      PromiseOrValue<BytesLike>,
      PromiseOrValue<BigNumberish>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "relayedMessages",
    values: [PromiseOrValue<BytesLike>]
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
    functionFragment: "sentMessages",
    values: [PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "successfulMessages",
    values: [PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "xDomainMessageSender",
    values?: undefined
  ): string;

  decodeFunctionResult(
    functionFragment: "l1CrossDomainMessenger",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "messageNonce",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "relayMessage",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "relayedMessages",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "sendMessage",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "sentMessages",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "successfulMessages",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "xDomainMessageSender",
    data: BytesLike
  ): Result;

  events: {
    "FailedRelayedMessage(bytes32)": EventFragment;
    "RelayedMessage(bytes32)": EventFragment;
    "SentMessage(address,address,bytes,uint256,uint256)": EventFragment;
  };

  getEvent(nameOrSignatureOrTopic: "FailedRelayedMessage"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "RelayedMessage"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "SentMessage"): EventFragment;
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

export interface L2CrossDomainMessenger extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: L2CrossDomainMessengerInterface;

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
    l1CrossDomainMessenger(overrides?: CallOverrides): Promise<[string]>;

    messageNonce(overrides?: CallOverrides): Promise<[BigNumber]>;

    relayMessage(
      _target: PromiseOrValue<string>,
      _sender: PromiseOrValue<string>,
      _message: PromiseOrValue<BytesLike>,
      _messageNonce: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    relayedMessages(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<[boolean]>;

    sendMessage(
      _target: PromiseOrValue<string>,
      _message: PromiseOrValue<BytesLike>,
      _gasLimit: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    sentMessages(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<[boolean]>;

    successfulMessages(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<[boolean]>;

    xDomainMessageSender(overrides?: CallOverrides): Promise<[string]>;
  };

  l1CrossDomainMessenger(overrides?: CallOverrides): Promise<string>;

  messageNonce(overrides?: CallOverrides): Promise<BigNumber>;

  relayMessage(
    _target: PromiseOrValue<string>,
    _sender: PromiseOrValue<string>,
    _message: PromiseOrValue<BytesLike>,
    _messageNonce: PromiseOrValue<BigNumberish>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  relayedMessages(
    arg0: PromiseOrValue<BytesLike>,
    overrides?: CallOverrides
  ): Promise<boolean>;

  sendMessage(
    _target: PromiseOrValue<string>,
    _message: PromiseOrValue<BytesLike>,
    _gasLimit: PromiseOrValue<BigNumberish>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  sentMessages(
    arg0: PromiseOrValue<BytesLike>,
    overrides?: CallOverrides
  ): Promise<boolean>;

  successfulMessages(
    arg0: PromiseOrValue<BytesLike>,
    overrides?: CallOverrides
  ): Promise<boolean>;

  xDomainMessageSender(overrides?: CallOverrides): Promise<string>;

  callStatic: {
    l1CrossDomainMessenger(overrides?: CallOverrides): Promise<string>;

    messageNonce(overrides?: CallOverrides): Promise<BigNumber>;

    relayMessage(
      _target: PromiseOrValue<string>,
      _sender: PromiseOrValue<string>,
      _message: PromiseOrValue<BytesLike>,
      _messageNonce: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<void>;

    relayedMessages(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<boolean>;

    sendMessage(
      _target: PromiseOrValue<string>,
      _message: PromiseOrValue<BytesLike>,
      _gasLimit: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<void>;

    sentMessages(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<boolean>;

    successfulMessages(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<boolean>;

    xDomainMessageSender(overrides?: CallOverrides): Promise<string>;
  };

  filters: {
    "FailedRelayedMessage(bytes32)"(
      msgHash?: PromiseOrValue<BytesLike> | null
    ): FailedRelayedMessageEventFilter;
    FailedRelayedMessage(
      msgHash?: PromiseOrValue<BytesLike> | null
    ): FailedRelayedMessageEventFilter;

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
  };

  estimateGas: {
    l1CrossDomainMessenger(overrides?: CallOverrides): Promise<BigNumber>;

    messageNonce(overrides?: CallOverrides): Promise<BigNumber>;

    relayMessage(
      _target: PromiseOrValue<string>,
      _sender: PromiseOrValue<string>,
      _message: PromiseOrValue<BytesLike>,
      _messageNonce: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    relayedMessages(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    sendMessage(
      _target: PromiseOrValue<string>,
      _message: PromiseOrValue<BytesLike>,
      _gasLimit: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    sentMessages(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    successfulMessages(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    xDomainMessageSender(overrides?: CallOverrides): Promise<BigNumber>;
  };

  populateTransaction: {
    l1CrossDomainMessenger(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    messageNonce(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    relayMessage(
      _target: PromiseOrValue<string>,
      _sender: PromiseOrValue<string>,
      _message: PromiseOrValue<BytesLike>,
      _messageNonce: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    relayedMessages(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    sendMessage(
      _target: PromiseOrValue<string>,
      _message: PromiseOrValue<BytesLike>,
      _gasLimit: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    sentMessages(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    successfulMessages(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    xDomainMessageSender(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;
  };
}
