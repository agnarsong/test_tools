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
import type { FunctionFragment, Result } from "@ethersproject/abi";
import type { Listener, Provider } from "@ethersproject/providers";
import type {
  TypedEventFilter,
  TypedEvent,
  TypedListener,
  OnEvent,
  PromiseOrValue,
} from "../../../../common";

export interface IEigenLayrDelegationInterface extends utils.Interface {
  functions: {
    "decreaseDelegatedShares(address,address[],uint256[])": FunctionFragment;
    "delegateTo(address)": FunctionFragment;
    "delegateToBySignature(address,address,uint256,bytes32,bytes32)": FunctionFragment;
    "delegatedTo(address)": FunctionFragment;
    "delegationTerms(address)": FunctionFragment;
    "increaseDelegatedShares(address,address,uint256)": FunctionFragment;
    "isDelegated(address)": FunctionFragment;
    "isNotDelegated(address)": FunctionFragment;
    "isOperator(address)": FunctionFragment;
    "operatorShares(address,address)": FunctionFragment;
    "registerAsOperator(address)": FunctionFragment;
    "undelegate(address)": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic:
      | "decreaseDelegatedShares"
      | "delegateTo"
      | "delegateToBySignature"
      | "delegatedTo"
      | "delegationTerms"
      | "increaseDelegatedShares"
      | "isDelegated"
      | "isNotDelegated"
      | "isOperator"
      | "operatorShares"
      | "registerAsOperator"
      | "undelegate"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "decreaseDelegatedShares",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<string>[],
      PromiseOrValue<BigNumberish>[]
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "delegateTo",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "delegateToBySignature",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<string>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BytesLike>,
      PromiseOrValue<BytesLike>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "delegatedTo",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "delegationTerms",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "increaseDelegatedShares",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<string>,
      PromiseOrValue<BigNumberish>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "isDelegated",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "isNotDelegated",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "isOperator",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "operatorShares",
    values: [PromiseOrValue<string>, PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "registerAsOperator",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "undelegate",
    values: [PromiseOrValue<string>]
  ): string;

  decodeFunctionResult(
    functionFragment: "decreaseDelegatedShares",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "delegateTo", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "delegateToBySignature",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "delegatedTo",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "delegationTerms",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "increaseDelegatedShares",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "isDelegated",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "isNotDelegated",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "isOperator", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "operatorShares",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "registerAsOperator",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "undelegate", data: BytesLike): Result;

  events: {};
}

export interface IEigenLayrDelegation extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: IEigenLayrDelegationInterface;

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
    decreaseDelegatedShares(
      staker: PromiseOrValue<string>,
      strategies: PromiseOrValue<string>[],
      shares: PromiseOrValue<BigNumberish>[],
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    delegateTo(
      operator: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    delegateToBySignature(
      staker: PromiseOrValue<string>,
      operator: PromiseOrValue<string>,
      expiry: PromiseOrValue<BigNumberish>,
      r: PromiseOrValue<BytesLike>,
      vs: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    delegatedTo(
      staker: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<[string]>;

    delegationTerms(
      operator: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<[string]>;

    increaseDelegatedShares(
      staker: PromiseOrValue<string>,
      strategy: PromiseOrValue<string>,
      shares: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    isDelegated(
      staker: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<[boolean]>;

    isNotDelegated(
      staker: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    isOperator(
      operator: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<[boolean]>;

    operatorShares(
      operator: PromiseOrValue<string>,
      strategy: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<[BigNumber]>;

    registerAsOperator(
      dt: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    undelegate(
      staker: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;
  };

  decreaseDelegatedShares(
    staker: PromiseOrValue<string>,
    strategies: PromiseOrValue<string>[],
    shares: PromiseOrValue<BigNumberish>[],
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  delegateTo(
    operator: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  delegateToBySignature(
    staker: PromiseOrValue<string>,
    operator: PromiseOrValue<string>,
    expiry: PromiseOrValue<BigNumberish>,
    r: PromiseOrValue<BytesLike>,
    vs: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  delegatedTo(
    staker: PromiseOrValue<string>,
    overrides?: CallOverrides
  ): Promise<string>;

  delegationTerms(
    operator: PromiseOrValue<string>,
    overrides?: CallOverrides
  ): Promise<string>;

  increaseDelegatedShares(
    staker: PromiseOrValue<string>,
    strategy: PromiseOrValue<string>,
    shares: PromiseOrValue<BigNumberish>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  isDelegated(
    staker: PromiseOrValue<string>,
    overrides?: CallOverrides
  ): Promise<boolean>;

  isNotDelegated(
    staker: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  isOperator(
    operator: PromiseOrValue<string>,
    overrides?: CallOverrides
  ): Promise<boolean>;

  operatorShares(
    operator: PromiseOrValue<string>,
    strategy: PromiseOrValue<string>,
    overrides?: CallOverrides
  ): Promise<BigNumber>;

  registerAsOperator(
    dt: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  undelegate(
    staker: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  callStatic: {
    decreaseDelegatedShares(
      staker: PromiseOrValue<string>,
      strategies: PromiseOrValue<string>[],
      shares: PromiseOrValue<BigNumberish>[],
      overrides?: CallOverrides
    ): Promise<void>;

    delegateTo(
      operator: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;

    delegateToBySignature(
      staker: PromiseOrValue<string>,
      operator: PromiseOrValue<string>,
      expiry: PromiseOrValue<BigNumberish>,
      r: PromiseOrValue<BytesLike>,
      vs: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;

    delegatedTo(
      staker: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<string>;

    delegationTerms(
      operator: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<string>;

    increaseDelegatedShares(
      staker: PromiseOrValue<string>,
      strategy: PromiseOrValue<string>,
      shares: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<void>;

    isDelegated(
      staker: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<boolean>;

    isNotDelegated(
      staker: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<boolean>;

    isOperator(
      operator: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<boolean>;

    operatorShares(
      operator: PromiseOrValue<string>,
      strategy: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    registerAsOperator(
      dt: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;

    undelegate(
      staker: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;
  };

  filters: {};

  estimateGas: {
    decreaseDelegatedShares(
      staker: PromiseOrValue<string>,
      strategies: PromiseOrValue<string>[],
      shares: PromiseOrValue<BigNumberish>[],
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    delegateTo(
      operator: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    delegateToBySignature(
      staker: PromiseOrValue<string>,
      operator: PromiseOrValue<string>,
      expiry: PromiseOrValue<BigNumberish>,
      r: PromiseOrValue<BytesLike>,
      vs: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    delegatedTo(
      staker: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    delegationTerms(
      operator: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    increaseDelegatedShares(
      staker: PromiseOrValue<string>,
      strategy: PromiseOrValue<string>,
      shares: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    isDelegated(
      staker: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    isNotDelegated(
      staker: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    isOperator(
      operator: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    operatorShares(
      operator: PromiseOrValue<string>,
      strategy: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    registerAsOperator(
      dt: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    undelegate(
      staker: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    decreaseDelegatedShares(
      staker: PromiseOrValue<string>,
      strategies: PromiseOrValue<string>[],
      shares: PromiseOrValue<BigNumberish>[],
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    delegateTo(
      operator: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    delegateToBySignature(
      staker: PromiseOrValue<string>,
      operator: PromiseOrValue<string>,
      expiry: PromiseOrValue<BigNumberish>,
      r: PromiseOrValue<BytesLike>,
      vs: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    delegatedTo(
      staker: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    delegationTerms(
      operator: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    increaseDelegatedShares(
      staker: PromiseOrValue<string>,
      strategy: PromiseOrValue<string>,
      shares: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    isDelegated(
      staker: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    isNotDelegated(
      staker: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    isOperator(
      operator: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    operatorShares(
      operator: PromiseOrValue<string>,
      strategy: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    registerAsOperator(
      dt: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    undelegate(
      staker: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;
  };
}
