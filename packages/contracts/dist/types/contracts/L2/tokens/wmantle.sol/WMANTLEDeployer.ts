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
} from "../../../../common";

export interface WMANTLEDeployerInterface extends utils.Interface {
  functions: {
    "calculateAddr()": FunctionFragment;
    "deploy()": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic: "calculateAddr" | "deploy"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "calculateAddr",
    values?: undefined
  ): string;
  encodeFunctionData(functionFragment: "deploy", values?: undefined): string;

  decodeFunctionResult(
    functionFragment: "calculateAddr",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "deploy", data: BytesLike): Result;

  events: {};
}

export interface WMANTLEDeployer extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: WMANTLEDeployerInterface;

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
    calculateAddr(
      overrides?: CallOverrides
    ): Promise<[string] & { predictedAddress: string }>;

    deploy(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;
  };

  calculateAddr(overrides?: CallOverrides): Promise<string>;

  deploy(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  callStatic: {
    calculateAddr(overrides?: CallOverrides): Promise<string>;

    deploy(overrides?: CallOverrides): Promise<string>;
  };

  filters: {};

  estimateGas: {
    calculateAddr(overrides?: CallOverrides): Promise<BigNumber>;

    deploy(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    calculateAddr(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    deploy(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;
  };
}
