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

export interface BVM_SequencerFeeVaultInterface extends utils.Interface {
  functions: {
    "L1Gas()": FunctionFragment;
    "burner()": FunctionFragment;
    "bvmGasPriceOracleAddress()": FunctionFragment;
    "l1FeeWallet()": FunctionFragment;
    "minWithdrawalAmount()": FunctionFragment;
    "owner()": FunctionFragment;
    "renounceOwnership()": FunctionFragment;
    "setBurner(address)": FunctionFragment;
    "setL1FeeWallet(address)": FunctionFragment;
    "setMinWithdrawalAmount(uint256)": FunctionFragment;
    "transferOwnership(address)": FunctionFragment;
    "withdraw()": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic:
      | "L1Gas"
      | "burner"
      | "bvmGasPriceOracleAddress"
      | "l1FeeWallet"
      | "minWithdrawalAmount"
      | "owner"
      | "renounceOwnership"
      | "setBurner"
      | "setL1FeeWallet"
      | "setMinWithdrawalAmount"
      | "transferOwnership"
      | "withdraw"
  ): FunctionFragment;

  encodeFunctionData(functionFragment: "L1Gas", values?: undefined): string;
  encodeFunctionData(functionFragment: "burner", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "bvmGasPriceOracleAddress",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "l1FeeWallet",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "minWithdrawalAmount",
    values?: undefined
  ): string;
  encodeFunctionData(functionFragment: "owner", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "renounceOwnership",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "setBurner",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "setL1FeeWallet",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "setMinWithdrawalAmount",
    values: [PromiseOrValue<BigNumberish>]
  ): string;
  encodeFunctionData(
    functionFragment: "transferOwnership",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(functionFragment: "withdraw", values?: undefined): string;

  decodeFunctionResult(functionFragment: "L1Gas", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "burner", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "bvmGasPriceOracleAddress",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "l1FeeWallet",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "minWithdrawalAmount",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "owner", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "renounceOwnership",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "setBurner", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "setL1FeeWallet",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "setMinWithdrawalAmount",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "transferOwnership",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "withdraw", data: BytesLike): Result;

  events: {
    "OwnershipTransferred(address,address)": EventFragment;
  };

  getEvent(nameOrSignatureOrTopic: "OwnershipTransferred"): EventFragment;
}

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

export interface BVM_SequencerFeeVault extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: BVM_SequencerFeeVaultInterface;

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
    L1Gas(overrides?: CallOverrides): Promise<[BigNumber]>;

    burner(overrides?: CallOverrides): Promise<[string]>;

    bvmGasPriceOracleAddress(overrides?: CallOverrides): Promise<[string]>;

    l1FeeWallet(overrides?: CallOverrides): Promise<[string]>;

    minWithdrawalAmount(overrides?: CallOverrides): Promise<[BigNumber]>;

    owner(overrides?: CallOverrides): Promise<[string]>;

    renounceOwnership(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    setBurner(
      _burner: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    setL1FeeWallet(
      _l1FeeWallet: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    setMinWithdrawalAmount(
      _minWithdrawalAmount: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    transferOwnership(
      newOwner: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    withdraw(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;
  };

  L1Gas(overrides?: CallOverrides): Promise<BigNumber>;

  burner(overrides?: CallOverrides): Promise<string>;

  bvmGasPriceOracleAddress(overrides?: CallOverrides): Promise<string>;

  l1FeeWallet(overrides?: CallOverrides): Promise<string>;

  minWithdrawalAmount(overrides?: CallOverrides): Promise<BigNumber>;

  owner(overrides?: CallOverrides): Promise<string>;

  renounceOwnership(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  setBurner(
    _burner: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  setL1FeeWallet(
    _l1FeeWallet: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  setMinWithdrawalAmount(
    _minWithdrawalAmount: PromiseOrValue<BigNumberish>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  transferOwnership(
    newOwner: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  withdraw(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  callStatic: {
    L1Gas(overrides?: CallOverrides): Promise<BigNumber>;

    burner(overrides?: CallOverrides): Promise<string>;

    bvmGasPriceOracleAddress(overrides?: CallOverrides): Promise<string>;

    l1FeeWallet(overrides?: CallOverrides): Promise<string>;

    minWithdrawalAmount(overrides?: CallOverrides): Promise<BigNumber>;

    owner(overrides?: CallOverrides): Promise<string>;

    renounceOwnership(overrides?: CallOverrides): Promise<void>;

    setBurner(
      _burner: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;

    setL1FeeWallet(
      _l1FeeWallet: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;

    setMinWithdrawalAmount(
      _minWithdrawalAmount: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<void>;

    transferOwnership(
      newOwner: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;

    withdraw(overrides?: CallOverrides): Promise<void>;
  };

  filters: {
    "OwnershipTransferred(address,address)"(
      previousOwner?: PromiseOrValue<string> | null,
      newOwner?: PromiseOrValue<string> | null
    ): OwnershipTransferredEventFilter;
    OwnershipTransferred(
      previousOwner?: PromiseOrValue<string> | null,
      newOwner?: PromiseOrValue<string> | null
    ): OwnershipTransferredEventFilter;
  };

  estimateGas: {
    L1Gas(overrides?: CallOverrides): Promise<BigNumber>;

    burner(overrides?: CallOverrides): Promise<BigNumber>;

    bvmGasPriceOracleAddress(overrides?: CallOverrides): Promise<BigNumber>;

    l1FeeWallet(overrides?: CallOverrides): Promise<BigNumber>;

    minWithdrawalAmount(overrides?: CallOverrides): Promise<BigNumber>;

    owner(overrides?: CallOverrides): Promise<BigNumber>;

    renounceOwnership(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    setBurner(
      _burner: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    setL1FeeWallet(
      _l1FeeWallet: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    setMinWithdrawalAmount(
      _minWithdrawalAmount: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    transferOwnership(
      newOwner: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    withdraw(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    L1Gas(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    burner(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    bvmGasPriceOracleAddress(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    l1FeeWallet(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    minWithdrawalAmount(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    owner(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    renounceOwnership(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    setBurner(
      _burner: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    setL1FeeWallet(
      _l1FeeWallet: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    setMinWithdrawalAmount(
      _minWithdrawalAmount: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    transferOwnership(
      newOwner: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    withdraw(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;
  };
}
