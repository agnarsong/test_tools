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

export declare namespace Lib_BVMCodec {
  export type QueueElementStruct = {
    transactionHash: PromiseOrValue<BytesLike>;
    timestamp: PromiseOrValue<BigNumberish>;
    blockNumber: PromiseOrValue<BigNumberish>;
  };

  export type QueueElementStructOutput = [string, number, number] & {
    transactionHash: string;
    timestamp: number;
    blockNumber: number;
  };
}

export interface CanonicalTransactionChainInterface extends utils.Interface {
  functions: {
    "MAX_ROLLUP_TX_SIZE()": FunctionFragment;
    "MIN_ROLLUP_TX_GAS()": FunctionFragment;
    "appendSequencerBatch()": FunctionFragment;
    "batches()": FunctionFragment;
    "enqueue(address,uint256,bytes)": FunctionFragment;
    "enqueueGasCost()": FunctionFragment;
    "enqueueL2GasPrepaid()": FunctionFragment;
    "getLastBlockNumber()": FunctionFragment;
    "getLastTimestamp()": FunctionFragment;
    "getNextQueueIndex()": FunctionFragment;
    "getNumPendingQueueElements()": FunctionFragment;
    "getQueueElement(uint256)": FunctionFragment;
    "getQueueLength()": FunctionFragment;
    "getTotalBatches()": FunctionFragment;
    "getTotalElements()": FunctionFragment;
    "l2GasDiscountDivisor()": FunctionFragment;
    "libAddressManager()": FunctionFragment;
    "maxTransactionGasLimit()": FunctionFragment;
    "resolve(string)": FunctionFragment;
    "setGasParams(uint256,uint256)": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic:
      | "MAX_ROLLUP_TX_SIZE"
      | "MIN_ROLLUP_TX_GAS"
      | "appendSequencerBatch"
      | "batches"
      | "enqueue"
      | "enqueueGasCost"
      | "enqueueL2GasPrepaid"
      | "getLastBlockNumber"
      | "getLastTimestamp"
      | "getNextQueueIndex"
      | "getNumPendingQueueElements"
      | "getQueueElement"
      | "getQueueLength"
      | "getTotalBatches"
      | "getTotalElements"
      | "l2GasDiscountDivisor"
      | "libAddressManager"
      | "maxTransactionGasLimit"
      | "resolve"
      | "setGasParams"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "MAX_ROLLUP_TX_SIZE",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "MIN_ROLLUP_TX_GAS",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "appendSequencerBatch",
    values?: undefined
  ): string;
  encodeFunctionData(functionFragment: "batches", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "enqueue",
    values: [
      PromiseOrValue<string>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BytesLike>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "enqueueGasCost",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "enqueueL2GasPrepaid",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "getLastBlockNumber",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "getLastTimestamp",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "getNextQueueIndex",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "getNumPendingQueueElements",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "getQueueElement",
    values: [PromiseOrValue<BigNumberish>]
  ): string;
  encodeFunctionData(
    functionFragment: "getQueueLength",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "getTotalBatches",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "getTotalElements",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "l2GasDiscountDivisor",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "libAddressManager",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "maxTransactionGasLimit",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "resolve",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "setGasParams",
    values: [PromiseOrValue<BigNumberish>, PromiseOrValue<BigNumberish>]
  ): string;

  decodeFunctionResult(
    functionFragment: "MAX_ROLLUP_TX_SIZE",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "MIN_ROLLUP_TX_GAS",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "appendSequencerBatch",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "batches", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "enqueue", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "enqueueGasCost",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "enqueueL2GasPrepaid",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getLastBlockNumber",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getLastTimestamp",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getNextQueueIndex",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getNumPendingQueueElements",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getQueueElement",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getQueueLength",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getTotalBatches",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getTotalElements",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "l2GasDiscountDivisor",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "libAddressManager",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "maxTransactionGasLimit",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "resolve", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "setGasParams",
    data: BytesLike
  ): Result;

  events: {
    "L2GasParamsUpdated(uint256,uint256,uint256)": EventFragment;
    "QueueBatchAppended(uint256,uint256,uint256)": EventFragment;
    "SequencerBatchAppended(uint256,uint256,uint256)": EventFragment;
    "TransactionBatchAppended(uint256,bytes32,uint256,uint256,bytes,bytes)": EventFragment;
    "TransactionEnqueued(address,address,uint256,bytes,uint256,uint256)": EventFragment;
  };

  getEvent(nameOrSignatureOrTopic: "L2GasParamsUpdated"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "QueueBatchAppended"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "SequencerBatchAppended"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "TransactionBatchAppended"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "TransactionEnqueued"): EventFragment;
}

export interface L2GasParamsUpdatedEventObject {
  l2GasDiscountDivisor: BigNumber;
  enqueueGasCost: BigNumber;
  enqueueL2GasPrepaid: BigNumber;
}
export type L2GasParamsUpdatedEvent = TypedEvent<
  [BigNumber, BigNumber, BigNumber],
  L2GasParamsUpdatedEventObject
>;

export type L2GasParamsUpdatedEventFilter =
  TypedEventFilter<L2GasParamsUpdatedEvent>;

export interface QueueBatchAppendedEventObject {
  _startingQueueIndex: BigNumber;
  _numQueueElements: BigNumber;
  _totalElements: BigNumber;
}
export type QueueBatchAppendedEvent = TypedEvent<
  [BigNumber, BigNumber, BigNumber],
  QueueBatchAppendedEventObject
>;

export type QueueBatchAppendedEventFilter =
  TypedEventFilter<QueueBatchAppendedEvent>;

export interface SequencerBatchAppendedEventObject {
  _startingQueueIndex: BigNumber;
  _numQueueElements: BigNumber;
  _totalElements: BigNumber;
}
export type SequencerBatchAppendedEvent = TypedEvent<
  [BigNumber, BigNumber, BigNumber],
  SequencerBatchAppendedEventObject
>;

export type SequencerBatchAppendedEventFilter =
  TypedEventFilter<SequencerBatchAppendedEvent>;

export interface TransactionBatchAppendedEventObject {
  _batchIndex: BigNumber;
  _batchRoot: string;
  _batchSize: BigNumber;
  _prevTotalElements: BigNumber;
  _signature: string;
  _extraData: string;
}
export type TransactionBatchAppendedEvent = TypedEvent<
  [BigNumber, string, BigNumber, BigNumber, string, string],
  TransactionBatchAppendedEventObject
>;

export type TransactionBatchAppendedEventFilter =
  TypedEventFilter<TransactionBatchAppendedEvent>;

export interface TransactionEnqueuedEventObject {
  _l1TxOrigin: string;
  _target: string;
  _gasLimit: BigNumber;
  _data: string;
  _queueIndex: BigNumber;
  _timestamp: BigNumber;
}
export type TransactionEnqueuedEvent = TypedEvent<
  [string, string, BigNumber, string, BigNumber, BigNumber],
  TransactionEnqueuedEventObject
>;

export type TransactionEnqueuedEventFilter =
  TypedEventFilter<TransactionEnqueuedEvent>;

export interface CanonicalTransactionChain extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: CanonicalTransactionChainInterface;

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
    MAX_ROLLUP_TX_SIZE(overrides?: CallOverrides): Promise<[BigNumber]>;

    MIN_ROLLUP_TX_GAS(overrides?: CallOverrides): Promise<[BigNumber]>;

    appendSequencerBatch(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    batches(overrides?: CallOverrides): Promise<[string]>;

    enqueue(
      _target: PromiseOrValue<string>,
      _gasLimit: PromiseOrValue<BigNumberish>,
      _data: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    enqueueGasCost(overrides?: CallOverrides): Promise<[BigNumber]>;

    enqueueL2GasPrepaid(overrides?: CallOverrides): Promise<[BigNumber]>;

    getLastBlockNumber(overrides?: CallOverrides): Promise<[number]>;

    getLastTimestamp(overrides?: CallOverrides): Promise<[number]>;

    getNextQueueIndex(overrides?: CallOverrides): Promise<[number]>;

    getNumPendingQueueElements(overrides?: CallOverrides): Promise<[number]>;

    getQueueElement(
      _index: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<
      [Lib_BVMCodec.QueueElementStructOutput] & {
        _element: Lib_BVMCodec.QueueElementStructOutput;
      }
    >;

    getQueueLength(overrides?: CallOverrides): Promise<[number]>;

    getTotalBatches(
      overrides?: CallOverrides
    ): Promise<[BigNumber] & { _totalBatches: BigNumber }>;

    getTotalElements(
      overrides?: CallOverrides
    ): Promise<[BigNumber] & { _totalElements: BigNumber }>;

    l2GasDiscountDivisor(overrides?: CallOverrides): Promise<[BigNumber]>;

    libAddressManager(overrides?: CallOverrides): Promise<[string]>;

    maxTransactionGasLimit(overrides?: CallOverrides): Promise<[BigNumber]>;

    resolve(
      _name: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<[string]>;

    setGasParams(
      _l2GasDiscountDivisor: PromiseOrValue<BigNumberish>,
      _enqueueGasCost: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;
  };

  MAX_ROLLUP_TX_SIZE(overrides?: CallOverrides): Promise<BigNumber>;

  MIN_ROLLUP_TX_GAS(overrides?: CallOverrides): Promise<BigNumber>;

  appendSequencerBatch(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  batches(overrides?: CallOverrides): Promise<string>;

  enqueue(
    _target: PromiseOrValue<string>,
    _gasLimit: PromiseOrValue<BigNumberish>,
    _data: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  enqueueGasCost(overrides?: CallOverrides): Promise<BigNumber>;

  enqueueL2GasPrepaid(overrides?: CallOverrides): Promise<BigNumber>;

  getLastBlockNumber(overrides?: CallOverrides): Promise<number>;

  getLastTimestamp(overrides?: CallOverrides): Promise<number>;

  getNextQueueIndex(overrides?: CallOverrides): Promise<number>;

  getNumPendingQueueElements(overrides?: CallOverrides): Promise<number>;

  getQueueElement(
    _index: PromiseOrValue<BigNumberish>,
    overrides?: CallOverrides
  ): Promise<Lib_BVMCodec.QueueElementStructOutput>;

  getQueueLength(overrides?: CallOverrides): Promise<number>;

  getTotalBatches(overrides?: CallOverrides): Promise<BigNumber>;

  getTotalElements(overrides?: CallOverrides): Promise<BigNumber>;

  l2GasDiscountDivisor(overrides?: CallOverrides): Promise<BigNumber>;

  libAddressManager(overrides?: CallOverrides): Promise<string>;

  maxTransactionGasLimit(overrides?: CallOverrides): Promise<BigNumber>;

  resolve(
    _name: PromiseOrValue<string>,
    overrides?: CallOverrides
  ): Promise<string>;

  setGasParams(
    _l2GasDiscountDivisor: PromiseOrValue<BigNumberish>,
    _enqueueGasCost: PromiseOrValue<BigNumberish>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  callStatic: {
    MAX_ROLLUP_TX_SIZE(overrides?: CallOverrides): Promise<BigNumber>;

    MIN_ROLLUP_TX_GAS(overrides?: CallOverrides): Promise<BigNumber>;

    appendSequencerBatch(overrides?: CallOverrides): Promise<void>;

    batches(overrides?: CallOverrides): Promise<string>;

    enqueue(
      _target: PromiseOrValue<string>,
      _gasLimit: PromiseOrValue<BigNumberish>,
      _data: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;

    enqueueGasCost(overrides?: CallOverrides): Promise<BigNumber>;

    enqueueL2GasPrepaid(overrides?: CallOverrides): Promise<BigNumber>;

    getLastBlockNumber(overrides?: CallOverrides): Promise<number>;

    getLastTimestamp(overrides?: CallOverrides): Promise<number>;

    getNextQueueIndex(overrides?: CallOverrides): Promise<number>;

    getNumPendingQueueElements(overrides?: CallOverrides): Promise<number>;

    getQueueElement(
      _index: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<Lib_BVMCodec.QueueElementStructOutput>;

    getQueueLength(overrides?: CallOverrides): Promise<number>;

    getTotalBatches(overrides?: CallOverrides): Promise<BigNumber>;

    getTotalElements(overrides?: CallOverrides): Promise<BigNumber>;

    l2GasDiscountDivisor(overrides?: CallOverrides): Promise<BigNumber>;

    libAddressManager(overrides?: CallOverrides): Promise<string>;

    maxTransactionGasLimit(overrides?: CallOverrides): Promise<BigNumber>;

    resolve(
      _name: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<string>;

    setGasParams(
      _l2GasDiscountDivisor: PromiseOrValue<BigNumberish>,
      _enqueueGasCost: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<void>;
  };

  filters: {
    "L2GasParamsUpdated(uint256,uint256,uint256)"(
      l2GasDiscountDivisor?: null,
      enqueueGasCost?: null,
      enqueueL2GasPrepaid?: null
    ): L2GasParamsUpdatedEventFilter;
    L2GasParamsUpdated(
      l2GasDiscountDivisor?: null,
      enqueueGasCost?: null,
      enqueueL2GasPrepaid?: null
    ): L2GasParamsUpdatedEventFilter;

    "QueueBatchAppended(uint256,uint256,uint256)"(
      _startingQueueIndex?: null,
      _numQueueElements?: null,
      _totalElements?: null
    ): QueueBatchAppendedEventFilter;
    QueueBatchAppended(
      _startingQueueIndex?: null,
      _numQueueElements?: null,
      _totalElements?: null
    ): QueueBatchAppendedEventFilter;

    "SequencerBatchAppended(uint256,uint256,uint256)"(
      _startingQueueIndex?: null,
      _numQueueElements?: null,
      _totalElements?: null
    ): SequencerBatchAppendedEventFilter;
    SequencerBatchAppended(
      _startingQueueIndex?: null,
      _numQueueElements?: null,
      _totalElements?: null
    ): SequencerBatchAppendedEventFilter;

    "TransactionBatchAppended(uint256,bytes32,uint256,uint256,bytes,bytes)"(
      _batchIndex?: PromiseOrValue<BigNumberish> | null,
      _batchRoot?: null,
      _batchSize?: null,
      _prevTotalElements?: null,
      _signature?: null,
      _extraData?: null
    ): TransactionBatchAppendedEventFilter;
    TransactionBatchAppended(
      _batchIndex?: PromiseOrValue<BigNumberish> | null,
      _batchRoot?: null,
      _batchSize?: null,
      _prevTotalElements?: null,
      _signature?: null,
      _extraData?: null
    ): TransactionBatchAppendedEventFilter;

    "TransactionEnqueued(address,address,uint256,bytes,uint256,uint256)"(
      _l1TxOrigin?: PromiseOrValue<string> | null,
      _target?: PromiseOrValue<string> | null,
      _gasLimit?: null,
      _data?: null,
      _queueIndex?: PromiseOrValue<BigNumberish> | null,
      _timestamp?: null
    ): TransactionEnqueuedEventFilter;
    TransactionEnqueued(
      _l1TxOrigin?: PromiseOrValue<string> | null,
      _target?: PromiseOrValue<string> | null,
      _gasLimit?: null,
      _data?: null,
      _queueIndex?: PromiseOrValue<BigNumberish> | null,
      _timestamp?: null
    ): TransactionEnqueuedEventFilter;
  };

  estimateGas: {
    MAX_ROLLUP_TX_SIZE(overrides?: CallOverrides): Promise<BigNumber>;

    MIN_ROLLUP_TX_GAS(overrides?: CallOverrides): Promise<BigNumber>;

    appendSequencerBatch(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    batches(overrides?: CallOverrides): Promise<BigNumber>;

    enqueue(
      _target: PromiseOrValue<string>,
      _gasLimit: PromiseOrValue<BigNumberish>,
      _data: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    enqueueGasCost(overrides?: CallOverrides): Promise<BigNumber>;

    enqueueL2GasPrepaid(overrides?: CallOverrides): Promise<BigNumber>;

    getLastBlockNumber(overrides?: CallOverrides): Promise<BigNumber>;

    getLastTimestamp(overrides?: CallOverrides): Promise<BigNumber>;

    getNextQueueIndex(overrides?: CallOverrides): Promise<BigNumber>;

    getNumPendingQueueElements(overrides?: CallOverrides): Promise<BigNumber>;

    getQueueElement(
      _index: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    getQueueLength(overrides?: CallOverrides): Promise<BigNumber>;

    getTotalBatches(overrides?: CallOverrides): Promise<BigNumber>;

    getTotalElements(overrides?: CallOverrides): Promise<BigNumber>;

    l2GasDiscountDivisor(overrides?: CallOverrides): Promise<BigNumber>;

    libAddressManager(overrides?: CallOverrides): Promise<BigNumber>;

    maxTransactionGasLimit(overrides?: CallOverrides): Promise<BigNumber>;

    resolve(
      _name: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    setGasParams(
      _l2GasDiscountDivisor: PromiseOrValue<BigNumberish>,
      _enqueueGasCost: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    MAX_ROLLUP_TX_SIZE(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    MIN_ROLLUP_TX_GAS(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    appendSequencerBatch(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    batches(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    enqueue(
      _target: PromiseOrValue<string>,
      _gasLimit: PromiseOrValue<BigNumberish>,
      _data: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    enqueueGasCost(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    enqueueL2GasPrepaid(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    getLastBlockNumber(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    getLastTimestamp(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    getNextQueueIndex(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    getNumPendingQueueElements(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    getQueueElement(
      _index: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    getQueueLength(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    getTotalBatches(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    getTotalElements(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    l2GasDiscountDivisor(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    libAddressManager(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    maxTransactionGasLimit(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    resolve(
      _name: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    setGasParams(
      _l2GasDiscountDivisor: PromiseOrValue<BigNumberish>,
      _enqueueGasCost: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;
  };
}
