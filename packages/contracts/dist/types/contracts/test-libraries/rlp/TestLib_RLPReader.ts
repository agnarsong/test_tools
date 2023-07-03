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

export interface TestLib_RLPReaderInterface extends utils.Interface {
  functions: {
    "readAddress(bytes)": FunctionFragment;
    "readBool(bytes)": FunctionFragment;
    "readBytes(bytes)": FunctionFragment;
    "readBytes32(bytes)": FunctionFragment;
    "readList(bytes)": FunctionFragment;
    "readString(bytes)": FunctionFragment;
    "readUint256(bytes)": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic:
      | "readAddress"
      | "readBool"
      | "readBytes"
      | "readBytes32"
      | "readList"
      | "readString"
      | "readUint256"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "readAddress",
    values: [PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "readBool",
    values: [PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "readBytes",
    values: [PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "readBytes32",
    values: [PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "readList",
    values: [PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "readString",
    values: [PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "readUint256",
    values: [PromiseOrValue<BytesLike>]
  ): string;

  decodeFunctionResult(
    functionFragment: "readAddress",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "readBool", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "readBytes", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "readBytes32",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "readList", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "readString", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "readUint256",
    data: BytesLike
  ): Result;

  events: {};
}

export interface TestLib_RLPReader extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: TestLib_RLPReaderInterface;

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
    readAddress(
      _in: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<[string]>;

    readBool(
      _in: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<[boolean]>;

    readBytes(
      _in: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<[string]>;

    readBytes32(
      _in: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<[string]>;

    readList(
      _in: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<[string[]]>;

    readString(
      _in: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<[string]>;

    readUint256(
      _in: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<[BigNumber]>;
  };

  readAddress(
    _in: PromiseOrValue<BytesLike>,
    overrides?: CallOverrides
  ): Promise<string>;

  readBool(
    _in: PromiseOrValue<BytesLike>,
    overrides?: CallOverrides
  ): Promise<boolean>;

  readBytes(
    _in: PromiseOrValue<BytesLike>,
    overrides?: CallOverrides
  ): Promise<string>;

  readBytes32(
    _in: PromiseOrValue<BytesLike>,
    overrides?: CallOverrides
  ): Promise<string>;

  readList(
    _in: PromiseOrValue<BytesLike>,
    overrides?: CallOverrides
  ): Promise<string[]>;

  readString(
    _in: PromiseOrValue<BytesLike>,
    overrides?: CallOverrides
  ): Promise<string>;

  readUint256(
    _in: PromiseOrValue<BytesLike>,
    overrides?: CallOverrides
  ): Promise<BigNumber>;

  callStatic: {
    readAddress(
      _in: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<string>;

    readBool(
      _in: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<boolean>;

    readBytes(
      _in: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<string>;

    readBytes32(
      _in: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<string>;

    readList(
      _in: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<string[]>;

    readString(
      _in: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<string>;

    readUint256(
      _in: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;
  };

  filters: {};

  estimateGas: {
    readAddress(
      _in: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    readBool(
      _in: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    readBytes(
      _in: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    readBytes32(
      _in: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    readList(
      _in: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    readString(
      _in: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    readUint256(
      _in: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    readAddress(
      _in: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    readBool(
      _in: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    readBytes(
      _in: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    readBytes32(
      _in: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    readList(
      _in: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    readString(
      _in: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    readUint256(
      _in: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;
  };
}
