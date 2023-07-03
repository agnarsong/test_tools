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
} from "../../common";

export interface MockStateCommitmentChainInterface extends utils.Interface {
  functions: {
    "FRAUD_PROOF_WINDOW()": FunctionFragment;
    "appendStateBatch(bytes32[],uint256,bytes)": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic: "FRAUD_PROOF_WINDOW" | "appendStateBatch"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "FRAUD_PROOF_WINDOW",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "appendStateBatch",
    values: [
      PromiseOrValue<BytesLike>[],
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BytesLike>
    ]
  ): string;

  decodeFunctionResult(
    functionFragment: "FRAUD_PROOF_WINDOW",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "appendStateBatch",
    data: BytesLike
  ): Result;

  events: {};
}

export interface MockStateCommitmentChain extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: MockStateCommitmentChainInterface;

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
    FRAUD_PROOF_WINDOW(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    appendStateBatch(
      _batch: PromiseOrValue<BytesLike>[],
      _shouldStartAtElement: PromiseOrValue<BigNumberish>,
      _signature: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;
  };

  FRAUD_PROOF_WINDOW(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  appendStateBatch(
    _batch: PromiseOrValue<BytesLike>[],
    _shouldStartAtElement: PromiseOrValue<BigNumberish>,
    _signature: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  callStatic: {
    FRAUD_PROOF_WINDOW(overrides?: CallOverrides): Promise<BigNumber>;

    appendStateBatch(
      _batch: PromiseOrValue<BytesLike>[],
      _shouldStartAtElement: PromiseOrValue<BigNumberish>,
      _signature: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;
  };

  filters: {};

  estimateGas: {
    FRAUD_PROOF_WINDOW(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    appendStateBatch(
      _batch: PromiseOrValue<BytesLike>[],
      _shouldStartAtElement: PromiseOrValue<BigNumberish>,
      _signature: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    FRAUD_PROOF_WINDOW(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    appendStateBatch(
      _batch: PromiseOrValue<BytesLike>[],
      _shouldStartAtElement: PromiseOrValue<BigNumberish>,
      _signature: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;
  };
}