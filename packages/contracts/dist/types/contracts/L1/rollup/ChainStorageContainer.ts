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
} from "../../../common";

export interface ChainStorageContainerInterface extends utils.Interface {
  functions: {
    "deleteElementsAfterInclusive(uint256,bytes27)": FunctionFragment;
    "deleteElementsAfterInclusive(uint256)": FunctionFragment;
    "get(uint256)": FunctionFragment;
    "getGlobalMetadata()": FunctionFragment;
    "length()": FunctionFragment;
    "libAddressManager()": FunctionFragment;
    "owner()": FunctionFragment;
    "push(bytes32,bytes27)": FunctionFragment;
    "push(bytes32)": FunctionFragment;
    "resolve(string)": FunctionFragment;
    "setGlobalMetadata(bytes27)": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic:
      | "deleteElementsAfterInclusive(uint256,bytes27)"
      | "deleteElementsAfterInclusive(uint256)"
      | "get"
      | "getGlobalMetadata"
      | "length"
      | "libAddressManager"
      | "owner"
      | "push(bytes32,bytes27)"
      | "push(bytes32)"
      | "resolve"
      | "setGlobalMetadata"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "deleteElementsAfterInclusive(uint256,bytes27)",
    values: [PromiseOrValue<BigNumberish>, PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "deleteElementsAfterInclusive(uint256)",
    values: [PromiseOrValue<BigNumberish>]
  ): string;
  encodeFunctionData(
    functionFragment: "get",
    values: [PromiseOrValue<BigNumberish>]
  ): string;
  encodeFunctionData(
    functionFragment: "getGlobalMetadata",
    values?: undefined
  ): string;
  encodeFunctionData(functionFragment: "length", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "libAddressManager",
    values?: undefined
  ): string;
  encodeFunctionData(functionFragment: "owner", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "push(bytes32,bytes27)",
    values: [PromiseOrValue<BytesLike>, PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "push(bytes32)",
    values: [PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "resolve",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "setGlobalMetadata",
    values: [PromiseOrValue<BytesLike>]
  ): string;

  decodeFunctionResult(
    functionFragment: "deleteElementsAfterInclusive(uint256,bytes27)",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "deleteElementsAfterInclusive(uint256)",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "get", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "getGlobalMetadata",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "length", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "libAddressManager",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "owner", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "push(bytes32,bytes27)",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "push(bytes32)",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "resolve", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "setGlobalMetadata",
    data: BytesLike
  ): Result;

  events: {};
}

export interface ChainStorageContainer extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: ChainStorageContainerInterface;

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
    "deleteElementsAfterInclusive(uint256,bytes27)"(
      _index: PromiseOrValue<BigNumberish>,
      _globalMetadata: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "deleteElementsAfterInclusive(uint256)"(
      _index: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    get(
      _index: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<[string]>;

    getGlobalMetadata(overrides?: CallOverrides): Promise<[string]>;

    length(overrides?: CallOverrides): Promise<[BigNumber]>;

    libAddressManager(overrides?: CallOverrides): Promise<[string]>;

    owner(overrides?: CallOverrides): Promise<[string]>;

    "push(bytes32,bytes27)"(
      _object: PromiseOrValue<BytesLike>,
      _globalMetadata: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    "push(bytes32)"(
      _object: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    resolve(
      _name: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<[string]>;

    setGlobalMetadata(
      _globalMetadata: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;
  };

  "deleteElementsAfterInclusive(uint256,bytes27)"(
    _index: PromiseOrValue<BigNumberish>,
    _globalMetadata: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "deleteElementsAfterInclusive(uint256)"(
    _index: PromiseOrValue<BigNumberish>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  get(
    _index: PromiseOrValue<BigNumberish>,
    overrides?: CallOverrides
  ): Promise<string>;

  getGlobalMetadata(overrides?: CallOverrides): Promise<string>;

  length(overrides?: CallOverrides): Promise<BigNumber>;

  libAddressManager(overrides?: CallOverrides): Promise<string>;

  owner(overrides?: CallOverrides): Promise<string>;

  "push(bytes32,bytes27)"(
    _object: PromiseOrValue<BytesLike>,
    _globalMetadata: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  "push(bytes32)"(
    _object: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  resolve(
    _name: PromiseOrValue<string>,
    overrides?: CallOverrides
  ): Promise<string>;

  setGlobalMetadata(
    _globalMetadata: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  callStatic: {
    "deleteElementsAfterInclusive(uint256,bytes27)"(
      _index: PromiseOrValue<BigNumberish>,
      _globalMetadata: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;

    "deleteElementsAfterInclusive(uint256)"(
      _index: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<void>;

    get(
      _index: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<string>;

    getGlobalMetadata(overrides?: CallOverrides): Promise<string>;

    length(overrides?: CallOverrides): Promise<BigNumber>;

    libAddressManager(overrides?: CallOverrides): Promise<string>;

    owner(overrides?: CallOverrides): Promise<string>;

    "push(bytes32,bytes27)"(
      _object: PromiseOrValue<BytesLike>,
      _globalMetadata: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;

    "push(bytes32)"(
      _object: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;

    resolve(
      _name: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<string>;

    setGlobalMetadata(
      _globalMetadata: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;
  };

  filters: {};

  estimateGas: {
    "deleteElementsAfterInclusive(uint256,bytes27)"(
      _index: PromiseOrValue<BigNumberish>,
      _globalMetadata: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "deleteElementsAfterInclusive(uint256)"(
      _index: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    get(
      _index: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    getGlobalMetadata(overrides?: CallOverrides): Promise<BigNumber>;

    length(overrides?: CallOverrides): Promise<BigNumber>;

    libAddressManager(overrides?: CallOverrides): Promise<BigNumber>;

    owner(overrides?: CallOverrides): Promise<BigNumber>;

    "push(bytes32,bytes27)"(
      _object: PromiseOrValue<BytesLike>,
      _globalMetadata: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    "push(bytes32)"(
      _object: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    resolve(
      _name: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    setGlobalMetadata(
      _globalMetadata: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    "deleteElementsAfterInclusive(uint256,bytes27)"(
      _index: PromiseOrValue<BigNumberish>,
      _globalMetadata: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "deleteElementsAfterInclusive(uint256)"(
      _index: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    get(
      _index: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    getGlobalMetadata(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    length(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    libAddressManager(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    owner(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    "push(bytes32,bytes27)"(
      _object: PromiseOrValue<BytesLike>,
      _globalMetadata: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    "push(bytes32)"(
      _object: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    resolve(
      _name: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    setGlobalMetadata(
      _globalMetadata: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;
  };
}