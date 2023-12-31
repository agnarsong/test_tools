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

export interface AssertionMapInterface extends utils.Interface {
  functions: {
    "assertions(uint256)": FunctionFragment;
    "createAssertion(uint256,bytes32,uint256,uint256,uint256)": FunctionFragment;
    "deleteAssertion(uint256)": FunctionFragment;
    "getDeadline(uint256)": FunctionFragment;
    "getInboxSize(uint256)": FunctionFragment;
    "getNumStakers(uint256)": FunctionFragment;
    "getParentID(uint256)": FunctionFragment;
    "getProposalTime(uint256)": FunctionFragment;
    "getStateHash(uint256)": FunctionFragment;
    "initialize()": FunctionFragment;
    "isStaker(uint256,address)": FunctionFragment;
    "rollupAddress()": FunctionFragment;
    "setRollupAddress(address)": FunctionFragment;
    "stakeOnAssertion(uint256,address)": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic:
      | "assertions"
      | "createAssertion"
      | "deleteAssertion"
      | "getDeadline"
      | "getInboxSize"
      | "getNumStakers"
      | "getParentID"
      | "getProposalTime"
      | "getStateHash"
      | "initialize"
      | "isStaker"
      | "rollupAddress"
      | "setRollupAddress"
      | "stakeOnAssertion"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "assertions",
    values: [PromiseOrValue<BigNumberish>]
  ): string;
  encodeFunctionData(
    functionFragment: "createAssertion",
    values: [
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BytesLike>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BigNumberish>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "deleteAssertion",
    values: [PromiseOrValue<BigNumberish>]
  ): string;
  encodeFunctionData(
    functionFragment: "getDeadline",
    values: [PromiseOrValue<BigNumberish>]
  ): string;
  encodeFunctionData(
    functionFragment: "getInboxSize",
    values: [PromiseOrValue<BigNumberish>]
  ): string;
  encodeFunctionData(
    functionFragment: "getNumStakers",
    values: [PromiseOrValue<BigNumberish>]
  ): string;
  encodeFunctionData(
    functionFragment: "getParentID",
    values: [PromiseOrValue<BigNumberish>]
  ): string;
  encodeFunctionData(
    functionFragment: "getProposalTime",
    values: [PromiseOrValue<BigNumberish>]
  ): string;
  encodeFunctionData(
    functionFragment: "getStateHash",
    values: [PromiseOrValue<BigNumberish>]
  ): string;
  encodeFunctionData(
    functionFragment: "initialize",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "isStaker",
    values: [PromiseOrValue<BigNumberish>, PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "rollupAddress",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "setRollupAddress",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "stakeOnAssertion",
    values: [PromiseOrValue<BigNumberish>, PromiseOrValue<string>]
  ): string;

  decodeFunctionResult(functionFragment: "assertions", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "createAssertion",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "deleteAssertion",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getDeadline",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getInboxSize",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getNumStakers",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getParentID",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getProposalTime",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getStateHash",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "initialize", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "isStaker", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "rollupAddress",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "setRollupAddress",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "stakeOnAssertion",
    data: BytesLike
  ): Result;

  events: {
    "Initialized(uint8)": EventFragment;
  };

  getEvent(nameOrSignatureOrTopic: "Initialized"): EventFragment;
}

export interface InitializedEventObject {
  version: number;
}
export type InitializedEvent = TypedEvent<[number], InitializedEventObject>;

export type InitializedEventFilter = TypedEventFilter<InitializedEvent>;

export interface AssertionMap extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: AssertionMapInterface;

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
    assertions(
      arg0: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<
      [
        string,
        BigNumber,
        BigNumber,
        BigNumber,
        BigNumber,
        BigNumber,
        BigNumber
      ] & {
        stateHash: string;
        inboxSize: BigNumber;
        parent: BigNumber;
        deadline: BigNumber;
        proposalTime: BigNumber;
        numStakers: BigNumber;
        childInboxSize: BigNumber;
      }
    >;

    createAssertion(
      assertionID: PromiseOrValue<BigNumberish>,
      stateHash: PromiseOrValue<BytesLike>,
      inboxSize: PromiseOrValue<BigNumberish>,
      parentID: PromiseOrValue<BigNumberish>,
      deadline: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    deleteAssertion(
      assertionID: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    getDeadline(
      assertionID: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<[BigNumber]>;

    getInboxSize(
      assertionID: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<[BigNumber]>;

    getNumStakers(
      assertionID: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<[BigNumber]>;

    getParentID(
      assertionID: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<[BigNumber]>;

    getProposalTime(
      assertionID: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<[BigNumber]>;

    getStateHash(
      assertionID: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<[string]>;

    initialize(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    isStaker(
      assertionID: PromiseOrValue<BigNumberish>,
      stakerAddress: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<[boolean]>;

    rollupAddress(overrides?: CallOverrides): Promise<[string]>;

    setRollupAddress(
      _rollupAddress: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    stakeOnAssertion(
      assertionID: PromiseOrValue<BigNumberish>,
      stakerAddress: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;
  };

  assertions(
    arg0: PromiseOrValue<BigNumberish>,
    overrides?: CallOverrides
  ): Promise<
    [
      string,
      BigNumber,
      BigNumber,
      BigNumber,
      BigNumber,
      BigNumber,
      BigNumber
    ] & {
      stateHash: string;
      inboxSize: BigNumber;
      parent: BigNumber;
      deadline: BigNumber;
      proposalTime: BigNumber;
      numStakers: BigNumber;
      childInboxSize: BigNumber;
    }
  >;

  createAssertion(
    assertionID: PromiseOrValue<BigNumberish>,
    stateHash: PromiseOrValue<BytesLike>,
    inboxSize: PromiseOrValue<BigNumberish>,
    parentID: PromiseOrValue<BigNumberish>,
    deadline: PromiseOrValue<BigNumberish>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  deleteAssertion(
    assertionID: PromiseOrValue<BigNumberish>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  getDeadline(
    assertionID: PromiseOrValue<BigNumberish>,
    overrides?: CallOverrides
  ): Promise<BigNumber>;

  getInboxSize(
    assertionID: PromiseOrValue<BigNumberish>,
    overrides?: CallOverrides
  ): Promise<BigNumber>;

  getNumStakers(
    assertionID: PromiseOrValue<BigNumberish>,
    overrides?: CallOverrides
  ): Promise<BigNumber>;

  getParentID(
    assertionID: PromiseOrValue<BigNumberish>,
    overrides?: CallOverrides
  ): Promise<BigNumber>;

  getProposalTime(
    assertionID: PromiseOrValue<BigNumberish>,
    overrides?: CallOverrides
  ): Promise<BigNumber>;

  getStateHash(
    assertionID: PromiseOrValue<BigNumberish>,
    overrides?: CallOverrides
  ): Promise<string>;

  initialize(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  isStaker(
    assertionID: PromiseOrValue<BigNumberish>,
    stakerAddress: PromiseOrValue<string>,
    overrides?: CallOverrides
  ): Promise<boolean>;

  rollupAddress(overrides?: CallOverrides): Promise<string>;

  setRollupAddress(
    _rollupAddress: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  stakeOnAssertion(
    assertionID: PromiseOrValue<BigNumberish>,
    stakerAddress: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  callStatic: {
    assertions(
      arg0: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<
      [
        string,
        BigNumber,
        BigNumber,
        BigNumber,
        BigNumber,
        BigNumber,
        BigNumber
      ] & {
        stateHash: string;
        inboxSize: BigNumber;
        parent: BigNumber;
        deadline: BigNumber;
        proposalTime: BigNumber;
        numStakers: BigNumber;
        childInboxSize: BigNumber;
      }
    >;

    createAssertion(
      assertionID: PromiseOrValue<BigNumberish>,
      stateHash: PromiseOrValue<BytesLike>,
      inboxSize: PromiseOrValue<BigNumberish>,
      parentID: PromiseOrValue<BigNumberish>,
      deadline: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<void>;

    deleteAssertion(
      assertionID: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<void>;

    getDeadline(
      assertionID: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    getInboxSize(
      assertionID: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    getNumStakers(
      assertionID: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    getParentID(
      assertionID: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    getProposalTime(
      assertionID: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    getStateHash(
      assertionID: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<string>;

    initialize(overrides?: CallOverrides): Promise<void>;

    isStaker(
      assertionID: PromiseOrValue<BigNumberish>,
      stakerAddress: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<boolean>;

    rollupAddress(overrides?: CallOverrides): Promise<string>;

    setRollupAddress(
      _rollupAddress: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;

    stakeOnAssertion(
      assertionID: PromiseOrValue<BigNumberish>,
      stakerAddress: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;
  };

  filters: {
    "Initialized(uint8)"(version?: null): InitializedEventFilter;
    Initialized(version?: null): InitializedEventFilter;
  };

  estimateGas: {
    assertions(
      arg0: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    createAssertion(
      assertionID: PromiseOrValue<BigNumberish>,
      stateHash: PromiseOrValue<BytesLike>,
      inboxSize: PromiseOrValue<BigNumberish>,
      parentID: PromiseOrValue<BigNumberish>,
      deadline: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    deleteAssertion(
      assertionID: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    getDeadline(
      assertionID: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    getInboxSize(
      assertionID: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    getNumStakers(
      assertionID: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    getParentID(
      assertionID: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    getProposalTime(
      assertionID: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    getStateHash(
      assertionID: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    initialize(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    isStaker(
      assertionID: PromiseOrValue<BigNumberish>,
      stakerAddress: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    rollupAddress(overrides?: CallOverrides): Promise<BigNumber>;

    setRollupAddress(
      _rollupAddress: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    stakeOnAssertion(
      assertionID: PromiseOrValue<BigNumberish>,
      stakerAddress: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    assertions(
      arg0: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    createAssertion(
      assertionID: PromiseOrValue<BigNumberish>,
      stateHash: PromiseOrValue<BytesLike>,
      inboxSize: PromiseOrValue<BigNumberish>,
      parentID: PromiseOrValue<BigNumberish>,
      deadline: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    deleteAssertion(
      assertionID: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    getDeadline(
      assertionID: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    getInboxSize(
      assertionID: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    getNumStakers(
      assertionID: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    getParentID(
      assertionID: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    getProposalTime(
      assertionID: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    getStateHash(
      assertionID: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    initialize(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    isStaker(
      assertionID: PromiseOrValue<BigNumberish>,
      stakerAddress: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    rollupAddress(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    setRollupAddress(
      _rollupAddress: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    stakeOnAssertion(
      assertionID: PromiseOrValue<BigNumberish>,
      stakerAddress: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;
  };
}
