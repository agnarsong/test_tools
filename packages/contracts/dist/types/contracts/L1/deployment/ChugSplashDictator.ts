/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
  BigNumber,
  BytesLike,
  CallOverrides,
  ContractTransaction,
  Overrides,
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
} from "../../../common";

export interface ChugSplashDictatorInterface extends utils.Interface {
  functions: {
    "bridgeSlotKey()": FunctionFragment;
    "bridgeSlotVal()": FunctionFragment;
    "codeHash()": FunctionFragment;
    "doActions(bytes)": FunctionFragment;
    "finalOwner()": FunctionFragment;
    "isUpgrading()": FunctionFragment;
    "mantleAddressSlotKey()": FunctionFragment;
    "mantleAddressSlotVal()": FunctionFragment;
    "messengerSlotKey()": FunctionFragment;
    "messengerSlotVal()": FunctionFragment;
    "returnOwnership()": FunctionFragment;
    "target()": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic:
      | "bridgeSlotKey"
      | "bridgeSlotVal"
      | "codeHash"
      | "doActions"
      | "finalOwner"
      | "isUpgrading"
      | "mantleAddressSlotKey"
      | "mantleAddressSlotVal"
      | "messengerSlotKey"
      | "messengerSlotVal"
      | "returnOwnership"
      | "target"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "bridgeSlotKey",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "bridgeSlotVal",
    values?: undefined
  ): string;
  encodeFunctionData(functionFragment: "codeHash", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "doActions",
    values: [PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "finalOwner",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "isUpgrading",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "mantleAddressSlotKey",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "mantleAddressSlotVal",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "messengerSlotKey",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "messengerSlotVal",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "returnOwnership",
    values?: undefined
  ): string;
  encodeFunctionData(functionFragment: "target", values?: undefined): string;

  decodeFunctionResult(
    functionFragment: "bridgeSlotKey",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "bridgeSlotVal",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "codeHash", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "doActions", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "finalOwner", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "isUpgrading",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "mantleAddressSlotKey",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "mantleAddressSlotVal",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "messengerSlotKey",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "messengerSlotVal",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "returnOwnership",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "target", data: BytesLike): Result;

  events: {};
}

export interface ChugSplashDictator extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: ChugSplashDictatorInterface;

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
    bridgeSlotKey(overrides?: CallOverrides): Promise<[string]>;

    bridgeSlotVal(overrides?: CallOverrides): Promise<[string]>;

    codeHash(overrides?: CallOverrides): Promise<[string]>;

    doActions(
      _code: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    finalOwner(overrides?: CallOverrides): Promise<[string]>;

    isUpgrading(overrides?: CallOverrides): Promise<[boolean]>;

    mantleAddressSlotKey(overrides?: CallOverrides): Promise<[string]>;

    mantleAddressSlotVal(overrides?: CallOverrides): Promise<[string]>;

    messengerSlotKey(overrides?: CallOverrides): Promise<[string]>;

    messengerSlotVal(overrides?: CallOverrides): Promise<[string]>;

    returnOwnership(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    target(overrides?: CallOverrides): Promise<[string]>;
  };

  bridgeSlotKey(overrides?: CallOverrides): Promise<string>;

  bridgeSlotVal(overrides?: CallOverrides): Promise<string>;

  codeHash(overrides?: CallOverrides): Promise<string>;

  doActions(
    _code: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  finalOwner(overrides?: CallOverrides): Promise<string>;

  isUpgrading(overrides?: CallOverrides): Promise<boolean>;

  mantleAddressSlotKey(overrides?: CallOverrides): Promise<string>;

  mantleAddressSlotVal(overrides?: CallOverrides): Promise<string>;

  messengerSlotKey(overrides?: CallOverrides): Promise<string>;

  messengerSlotVal(overrides?: CallOverrides): Promise<string>;

  returnOwnership(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  target(overrides?: CallOverrides): Promise<string>;

  callStatic: {
    bridgeSlotKey(overrides?: CallOverrides): Promise<string>;

    bridgeSlotVal(overrides?: CallOverrides): Promise<string>;

    codeHash(overrides?: CallOverrides): Promise<string>;

    doActions(
      _code: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;

    finalOwner(overrides?: CallOverrides): Promise<string>;

    isUpgrading(overrides?: CallOverrides): Promise<boolean>;

    mantleAddressSlotKey(overrides?: CallOverrides): Promise<string>;

    mantleAddressSlotVal(overrides?: CallOverrides): Promise<string>;

    messengerSlotKey(overrides?: CallOverrides): Promise<string>;

    messengerSlotVal(overrides?: CallOverrides): Promise<string>;

    returnOwnership(overrides?: CallOverrides): Promise<void>;

    target(overrides?: CallOverrides): Promise<string>;
  };

  filters: {};

  estimateGas: {
    bridgeSlotKey(overrides?: CallOverrides): Promise<BigNumber>;

    bridgeSlotVal(overrides?: CallOverrides): Promise<BigNumber>;

    codeHash(overrides?: CallOverrides): Promise<BigNumber>;

    doActions(
      _code: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    finalOwner(overrides?: CallOverrides): Promise<BigNumber>;

    isUpgrading(overrides?: CallOverrides): Promise<BigNumber>;

    mantleAddressSlotKey(overrides?: CallOverrides): Promise<BigNumber>;

    mantleAddressSlotVal(overrides?: CallOverrides): Promise<BigNumber>;

    messengerSlotKey(overrides?: CallOverrides): Promise<BigNumber>;

    messengerSlotVal(overrides?: CallOverrides): Promise<BigNumber>;

    returnOwnership(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    target(overrides?: CallOverrides): Promise<BigNumber>;
  };

  populateTransaction: {
    bridgeSlotKey(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    bridgeSlotVal(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    codeHash(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    doActions(
      _code: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    finalOwner(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    isUpgrading(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    mantleAddressSlotKey(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    mantleAddressSlotVal(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    messengerSlotKey(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    messengerSlotVal(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    returnOwnership(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    target(overrides?: CallOverrides): Promise<PopulatedTransaction>;
  };
}
