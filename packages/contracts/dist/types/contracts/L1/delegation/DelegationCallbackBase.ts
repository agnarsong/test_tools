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
  PayableOverrides,
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

export interface DelegationCallbackBaseInterface extends utils.Interface {
  functions: {
    "delegation()": FunctionFragment;
    "onDelegationReceived(address,address[],uint256[])": FunctionFragment;
    "onDelegationReceived(address,address,address[],uint256[])": FunctionFragment;
    "onDelegationWithdrawn(address,address,address[],uint256[])": FunctionFragment;
    "onDelegationWithdrawn(address,address[],uint256[])": FunctionFragment;
    "paused()": FunctionFragment;
    "payForService(address,uint256)": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic:
      | "delegation"
      | "onDelegationReceived(address,address[],uint256[])"
      | "onDelegationReceived(address,address,address[],uint256[])"
      | "onDelegationWithdrawn(address,address,address[],uint256[])"
      | "onDelegationWithdrawn(address,address[],uint256[])"
      | "paused"
      | "payForService"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "delegation",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "onDelegationReceived(address,address[],uint256[])",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<string>[],
      PromiseOrValue<BigNumberish>[]
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "onDelegationReceived(address,address,address[],uint256[])",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<string>,
      PromiseOrValue<string>[],
      PromiseOrValue<BigNumberish>[]
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "onDelegationWithdrawn(address,address,address[],uint256[])",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<string>,
      PromiseOrValue<string>[],
      PromiseOrValue<BigNumberish>[]
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "onDelegationWithdrawn(address,address[],uint256[])",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<string>[],
      PromiseOrValue<BigNumberish>[]
    ]
  ): string;
  encodeFunctionData(functionFragment: "paused", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "payForService",
    values: [PromiseOrValue<string>, PromiseOrValue<BigNumberish>]
  ): string;

  decodeFunctionResult(functionFragment: "delegation", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "onDelegationReceived(address,address[],uint256[])",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "onDelegationReceived(address,address,address[],uint256[])",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "onDelegationWithdrawn(address,address,address[],uint256[])",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "onDelegationWithdrawn(address,address[],uint256[])",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "paused", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "payForService",
    data: BytesLike
  ): Result;

  events: {
    "Initialized(uint8)": EventFragment;
    "Paused(address)": EventFragment;
    "Unpaused(address)": EventFragment;
  };

  getEvent(nameOrSignatureOrTopic: "Initialized"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "Paused"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "Unpaused"): EventFragment;
}

export interface InitializedEventObject {
  version: number;
}
export type InitializedEvent = TypedEvent<[number], InitializedEventObject>;

export type InitializedEventFilter = TypedEventFilter<InitializedEvent>;

export interface PausedEventObject {
  account: string;
}
export type PausedEvent = TypedEvent<[string], PausedEventObject>;

export type PausedEventFilter = TypedEventFilter<PausedEvent>;

export interface UnpausedEventObject {
  account: string;
}
export type UnpausedEvent = TypedEvent<[string], UnpausedEventObject>;

export type UnpausedEventFilter = TypedEventFilter<UnpausedEvent>;

export interface DelegationCallbackBase extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: DelegationCallbackBaseInterface;

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
    delegation(overrides?: CallOverrides): Promise<[string]>;

    "onDelegationReceived(address,address[],uint256[])"(
      delegator: PromiseOrValue<string>,
      investorDelegationShares: PromiseOrValue<string>[],
      investorShares: PromiseOrValue<BigNumberish>[],
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "onDelegationReceived(address,address,address[],uint256[])"(
      delegator: PromiseOrValue<string>,
      operator: PromiseOrValue<string>,
      delegationShares: PromiseOrValue<string>[],
      investorShares: PromiseOrValue<BigNumberish>[],
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "onDelegationWithdrawn(address,address,address[],uint256[])"(
      delegator: PromiseOrValue<string>,
      operator: PromiseOrValue<string>,
      delegationShares: PromiseOrValue<string>[],
      investorShares: PromiseOrValue<BigNumberish>[],
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "onDelegationWithdrawn(address,address[],uint256[])"(
      delegator: PromiseOrValue<string>,
      investorDelegationShares: PromiseOrValue<string>[],
      investorShares: PromiseOrValue<BigNumberish>[],
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    paused(overrides?: CallOverrides): Promise<[boolean]>;

    payForService(
      token: PromiseOrValue<string>,
      amount: PromiseOrValue<BigNumberish>,
      overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;
  };

  delegation(overrides?: CallOverrides): Promise<string>;

  "onDelegationReceived(address,address[],uint256[])"(
    delegator: PromiseOrValue<string>,
    investorDelegationShares: PromiseOrValue<string>[],
    investorShares: PromiseOrValue<BigNumberish>[],
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "onDelegationReceived(address,address,address[],uint256[])"(
    delegator: PromiseOrValue<string>,
    operator: PromiseOrValue<string>,
    delegationShares: PromiseOrValue<string>[],
    investorShares: PromiseOrValue<BigNumberish>[],
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "onDelegationWithdrawn(address,address,address[],uint256[])"(
    delegator: PromiseOrValue<string>,
    operator: PromiseOrValue<string>,
    delegationShares: PromiseOrValue<string>[],
    investorShares: PromiseOrValue<BigNumberish>[],
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "onDelegationWithdrawn(address,address[],uint256[])"(
    delegator: PromiseOrValue<string>,
    investorDelegationShares: PromiseOrValue<string>[],
    investorShares: PromiseOrValue<BigNumberish>[],
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  paused(overrides?: CallOverrides): Promise<boolean>;

  payForService(
    token: PromiseOrValue<string>,
    amount: PromiseOrValue<BigNumberish>,
    overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  callStatic: {
    delegation(overrides?: CallOverrides): Promise<string>;

    "onDelegationReceived(address,address[],uint256[])"(
      delegator: PromiseOrValue<string>,
      investorDelegationShares: PromiseOrValue<string>[],
      investorShares: PromiseOrValue<BigNumberish>[],
      overrides?: CallOverrides
    ): Promise<void>;

    "onDelegationReceived(address,address,address[],uint256[])"(
      delegator: PromiseOrValue<string>,
      operator: PromiseOrValue<string>,
      delegationShares: PromiseOrValue<string>[],
      investorShares: PromiseOrValue<BigNumberish>[],
      overrides?: CallOverrides
    ): Promise<void>;

    "onDelegationWithdrawn(address,address,address[],uint256[])"(
      delegator: PromiseOrValue<string>,
      operator: PromiseOrValue<string>,
      delegationShares: PromiseOrValue<string>[],
      investorShares: PromiseOrValue<BigNumberish>[],
      overrides?: CallOverrides
    ): Promise<void>;

    "onDelegationWithdrawn(address,address[],uint256[])"(
      delegator: PromiseOrValue<string>,
      investorDelegationShares: PromiseOrValue<string>[],
      investorShares: PromiseOrValue<BigNumberish>[],
      overrides?: CallOverrides
    ): Promise<void>;

    paused(overrides?: CallOverrides): Promise<boolean>;

    payForService(
      token: PromiseOrValue<string>,
      amount: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<void>;
  };

  filters: {
    "Initialized(uint8)"(version?: null): InitializedEventFilter;
    Initialized(version?: null): InitializedEventFilter;

    "Paused(address)"(account?: null): PausedEventFilter;
    Paused(account?: null): PausedEventFilter;

    "Unpaused(address)"(account?: null): UnpausedEventFilter;
    Unpaused(account?: null): UnpausedEventFilter;
  };

  estimateGas: {
    delegation(overrides?: CallOverrides): Promise<BigNumber>;

    "onDelegationReceived(address,address[],uint256[])"(
      delegator: PromiseOrValue<string>,
      investorDelegationShares: PromiseOrValue<string>[],
      investorShares: PromiseOrValue<BigNumberish>[],
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "onDelegationReceived(address,address,address[],uint256[])"(
      delegator: PromiseOrValue<string>,
      operator: PromiseOrValue<string>,
      delegationShares: PromiseOrValue<string>[],
      investorShares: PromiseOrValue<BigNumberish>[],
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "onDelegationWithdrawn(address,address,address[],uint256[])"(
      delegator: PromiseOrValue<string>,
      operator: PromiseOrValue<string>,
      delegationShares: PromiseOrValue<string>[],
      investorShares: PromiseOrValue<BigNumberish>[],
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "onDelegationWithdrawn(address,address[],uint256[])"(
      delegator: PromiseOrValue<string>,
      investorDelegationShares: PromiseOrValue<string>[],
      investorShares: PromiseOrValue<BigNumberish>[],
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    paused(overrides?: CallOverrides): Promise<BigNumber>;

    payForService(
      token: PromiseOrValue<string>,
      amount: PromiseOrValue<BigNumberish>,
      overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    delegation(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    "onDelegationReceived(address,address[],uint256[])"(
      delegator: PromiseOrValue<string>,
      investorDelegationShares: PromiseOrValue<string>[],
      investorShares: PromiseOrValue<BigNumberish>[],
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "onDelegationReceived(address,address,address[],uint256[])"(
      delegator: PromiseOrValue<string>,
      operator: PromiseOrValue<string>,
      delegationShares: PromiseOrValue<string>[],
      investorShares: PromiseOrValue<BigNumberish>[],
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "onDelegationWithdrawn(address,address,address[],uint256[])"(
      delegator: PromiseOrValue<string>,
      operator: PromiseOrValue<string>,
      delegationShares: PromiseOrValue<string>[],
      investorShares: PromiseOrValue<BigNumberish>[],
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "onDelegationWithdrawn(address,address[],uint256[])"(
      delegator: PromiseOrValue<string>,
      investorDelegationShares: PromiseOrValue<string>[],
      investorShares: PromiseOrValue<BigNumberish>[],
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    paused(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    payForService(
      token: PromiseOrValue<string>,
      amount: PromiseOrValue<BigNumberish>,
      overrides?: PayableOverrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;
  };
}
