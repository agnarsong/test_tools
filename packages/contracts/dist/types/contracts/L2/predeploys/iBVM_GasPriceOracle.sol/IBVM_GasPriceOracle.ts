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
} from "../../../../common";

export interface IBVM_GasPriceOracleInterface extends utils.Interface {
  functions: {
    "IsBurning()": FunctionFragment;
  };

  getFunction(nameOrSignatureOrTopic: "IsBurning"): FunctionFragment;

  encodeFunctionData(functionFragment: "IsBurning", values?: undefined): string;

  decodeFunctionResult(functionFragment: "IsBurning", data: BytesLike): Result;

  events: {};
}

export interface IBVM_GasPriceOracle extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: IBVM_GasPriceOracleInterface;

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
    IsBurning(overrides?: CallOverrides): Promise<[BigNumber]>;
  };

  IsBurning(overrides?: CallOverrides): Promise<BigNumber>;

  callStatic: {
    IsBurning(overrides?: CallOverrides): Promise<BigNumber>;
  };

  filters: {};

  estimateGas: {
    IsBurning(overrides?: CallOverrides): Promise<BigNumber>;
  };

  populateTransaction: {
    IsBurning(overrides?: CallOverrides): Promise<PopulatedTransaction>;
  };
}
