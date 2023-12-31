/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import { Signer, utils, Contract, ContractFactory, Overrides } from "ethers";
import type { Provider, TransactionRequest } from "@ethersproject/providers";
import type { PromiseOrValue } from "../../../../common";
import type {
  TestLib_CrossDomainUtils,
  TestLib_CrossDomainUtilsInterface,
} from "../../../../contracts/test-libraries/bridge/TestLib_CrossDomainUtils";

const _abi = [
  {
    inputs: [
      {
        internalType: "address",
        name: "_target",
        type: "address",
      },
      {
        internalType: "address",
        name: "_sender",
        type: "address",
      },
      {
        internalType: "bytes",
        name: "_message",
        type: "bytes",
      },
      {
        internalType: "uint256",
        name: "_messageNonce",
        type: "uint256",
      },
    ],
    name: "encodeXDomainCalldata",
    outputs: [
      {
        internalType: "bytes",
        name: "",
        type: "bytes",
      },
    ],
    stateMutability: "pure",
    type: "function",
  },
] as const;

const _bytecode =
  "0x6102ad61003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c8063053156471461003a575b600080fd5b61004d6100483660046100f9565b610063565b60405161005a9190610220565b60405180910390f35b60606100718585858561007a565b95945050505050565b606084848484604051602401610093949392919061023a565b60408051601f198184030181529190526020810180516001600160e01b031663cbd4ece960e01b1790529050949350505050565b80356001600160a01b03811681146100de57600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b6000806000806080858703121561010f57600080fd5b610118856100c7565b9350610126602086016100c7565b9250604085013567ffffffffffffffff8082111561014357600080fd5b818701915087601f83011261015757600080fd5b813581811115610169576101696100e3565b604051601f8201601f19908116603f01168101908382118183101715610191576101916100e3565b816040528281528a60208487010111156101aa57600080fd5b826020860160208301376000928101602001929092525095989497509495606001359450505050565b6000815180845260005b818110156101f9576020818501810151868301820152016101dd565b8181111561020b576000602083870101525b50601f01601f19169290920160200192915050565b60208152600061023360208301846101d3565b9392505050565b6001600160a01b03858116825284166020820152608060408201819052600090610266908301856101d3565b90508260608301529594505050505056fea2646970667358221220d89a00b89e0e296fec9705046db2ceaba537f3a66938c9ccaa4290a35828ec6b64736f6c63430008090033";

type TestLib_CrossDomainUtilsConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: TestLib_CrossDomainUtilsConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class TestLib_CrossDomainUtils__factory extends ContractFactory {
  constructor(...args: TestLib_CrossDomainUtilsConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override deploy(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<TestLib_CrossDomainUtils> {
    return super.deploy(overrides || {}) as Promise<TestLib_CrossDomainUtils>;
  }
  override getDeployTransaction(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(overrides || {});
  }
  override attach(address: string): TestLib_CrossDomainUtils {
    return super.attach(address) as TestLib_CrossDomainUtils;
  }
  override connect(signer: Signer): TestLib_CrossDomainUtils__factory {
    return super.connect(signer) as TestLib_CrossDomainUtils__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): TestLib_CrossDomainUtilsInterface {
    return new utils.Interface(_abi) as TestLib_CrossDomainUtilsInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): TestLib_CrossDomainUtils {
    return new Contract(
      address,
      _abi,
      signerOrProvider
    ) as TestLib_CrossDomainUtils;
  }
}
