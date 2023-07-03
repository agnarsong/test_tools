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

export interface IBVM_L2ToL1MessagePasserInterface extends utils.Interface {
  functions: {
    "passMessageToL1(bytes)": FunctionFragment;
  };

  getFunction(nameOrSignatureOrTopic: "passMessageToL1"): FunctionFragment;

  encodeFunctionData(
    functionFragment: "passMessageToL1",
    values: [PromiseOrValue<BytesLike>]
  ): string;

  decodeFunctionResult(
    functionFragment: "passMessageToL1",
    data: BytesLike
  ): Result;

  events: {
    "L2ToL1Message(uint256,address,bytes)": EventFragment;
  };

  getEvent(nameOrSignatureOrTopic: "L2ToL1Message"): EventFragment;
}

export interface L2ToL1MessageEventObject {
  _nonce: BigNumber;
  _sender: string;
  _data: string;
}
export type L2ToL1MessageEvent = TypedEvent<
  [BigNumber, string, string],
  L2ToL1MessageEventObject
>;

export type L2ToL1MessageEventFilter = TypedEventFilter<L2ToL1MessageEvent>;

export interface IBVM_L2ToL1MessagePasser extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: IBVM_L2ToL1MessagePasserInterface;

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
    passMessageToL1(
      _message: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;
  };

  passMessageToL1(
    _message: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  callStatic: {
    passMessageToL1(
      _message: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;
  };

  filters: {
    "L2ToL1Message(uint256,address,bytes)"(
      _nonce?: null,
      _sender?: null,
      _data?: null
    ): L2ToL1MessageEventFilter;
    L2ToL1Message(
      _nonce?: null,
      _sender?: null,
      _data?: null
    ): L2ToL1MessageEventFilter;
  };

  estimateGas: {
    passMessageToL1(
      _message: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    passMessageToL1(
      _message: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;
  };
}
