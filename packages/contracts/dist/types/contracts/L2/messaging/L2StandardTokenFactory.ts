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

export interface L2StandardTokenFactoryInterface extends utils.Interface {
  functions: {
    "createStandardL2Token(address,string,string,uint8)": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic: "createStandardL2Token"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "createStandardL2Token",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<string>,
      PromiseOrValue<string>,
      PromiseOrValue<BigNumberish>
    ]
  ): string;

  decodeFunctionResult(
    functionFragment: "createStandardL2Token",
    data: BytesLike
  ): Result;

  events: {
    "StandardL2TokenCreated(address,address,uint8)": EventFragment;
  };

  getEvent(nameOrSignatureOrTopic: "StandardL2TokenCreated"): EventFragment;
}

export interface StandardL2TokenCreatedEventObject {
  _l1Token: string;
  _l2Token: string;
  decimal: number;
}
export type StandardL2TokenCreatedEvent = TypedEvent<
  [string, string, number],
  StandardL2TokenCreatedEventObject
>;

export type StandardL2TokenCreatedEventFilter =
  TypedEventFilter<StandardL2TokenCreatedEvent>;

export interface L2StandardTokenFactory extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: L2StandardTokenFactoryInterface;

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
    createStandardL2Token(
      _l1Token: PromiseOrValue<string>,
      _name: PromiseOrValue<string>,
      _symbol: PromiseOrValue<string>,
      _decimal: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;
  };

  createStandardL2Token(
    _l1Token: PromiseOrValue<string>,
    _name: PromiseOrValue<string>,
    _symbol: PromiseOrValue<string>,
    _decimal: PromiseOrValue<BigNumberish>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  callStatic: {
    createStandardL2Token(
      _l1Token: PromiseOrValue<string>,
      _name: PromiseOrValue<string>,
      _symbol: PromiseOrValue<string>,
      _decimal: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<void>;
  };

  filters: {
    "StandardL2TokenCreated(address,address,uint8)"(
      _l1Token?: PromiseOrValue<string> | null,
      _l2Token?: PromiseOrValue<string> | null,
      decimal?: null
    ): StandardL2TokenCreatedEventFilter;
    StandardL2TokenCreated(
      _l1Token?: PromiseOrValue<string> | null,
      _l2Token?: PromiseOrValue<string> | null,
      decimal?: null
    ): StandardL2TokenCreatedEventFilter;
  };

  estimateGas: {
    createStandardL2Token(
      _l1Token: PromiseOrValue<string>,
      _name: PromiseOrValue<string>,
      _symbol: PromiseOrValue<string>,
      _decimal: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    createStandardL2Token(
      _l1Token: PromiseOrValue<string>,
      _name: PromiseOrValue<string>,
      _symbol: PromiseOrValue<string>,
      _decimal: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;
  };
}
