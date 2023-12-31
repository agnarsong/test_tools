/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
  BigNumber,
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
} from "../../../common";

export interface TestLib_AddressAliasHelperInterface extends utils.Interface {
  functions: {
    "applyL1ToL2Alias(address)": FunctionFragment;
    "undoL1ToL2Alias(address)": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic: "applyL1ToL2Alias" | "undoL1ToL2Alias"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "applyL1ToL2Alias",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "undoL1ToL2Alias",
    values: [PromiseOrValue<string>]
  ): string;

  decodeFunctionResult(
    functionFragment: "applyL1ToL2Alias",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "undoL1ToL2Alias",
    data: BytesLike
  ): Result;

  events: {};
}

export interface TestLib_AddressAliasHelper extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: TestLib_AddressAliasHelperInterface;

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
    applyL1ToL2Alias(
      _address: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<[string]>;

    undoL1ToL2Alias(
      _address: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<[string]>;
  };

  applyL1ToL2Alias(
    _address: PromiseOrValue<string>,
    overrides?: CallOverrides
  ): Promise<string>;

  undoL1ToL2Alias(
    _address: PromiseOrValue<string>,
    overrides?: CallOverrides
  ): Promise<string>;

  callStatic: {
    applyL1ToL2Alias(
      _address: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<string>;

    undoL1ToL2Alias(
      _address: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<string>;
  };

  filters: {};

  estimateGas: {
    applyL1ToL2Alias(
      _address: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    undoL1ToL2Alias(
      _address: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    applyL1ToL2Alias(
      _address: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    undoL1ToL2Alias(
      _address: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;
  };
}
