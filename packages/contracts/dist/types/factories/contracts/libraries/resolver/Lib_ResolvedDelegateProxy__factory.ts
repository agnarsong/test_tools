/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import { Signer, utils, Contract, ContractFactory, Overrides } from "ethers";
import type { Provider, TransactionRequest } from "@ethersproject/providers";
import type { PromiseOrValue } from "../../../../common";
import type {
  Lib_ResolvedDelegateProxy,
  Lib_ResolvedDelegateProxyInterface,
} from "../../../../contracts/libraries/resolver/Lib_ResolvedDelegateProxy";

const _abi = [
  {
    inputs: [
      {
        internalType: "address",
        name: "_libAddressManager",
        type: "address",
      },
      {
        internalType: "string",
        name: "_implementationName",
        type: "string",
      },
    ],
    stateMutability: "nonpayable",
    type: "constructor",
  },
  {
    stateMutability: "payable",
    type: "fallback",
  },
] as const;

const _bytecode =
  "0x608060405234801561001057600080fd5b506040516104fd3803806104fd83398101604081905261002f91610125565b30600090815260016020908152604080832080546001600160a01b0319166001600160a01b038716179055828252909120825161006e92840190610076565b505050610252565b82805461008290610217565b90600052602060002090601f0160209004810192826100a457600085556100ea565b82601f106100bd57805160ff19168380011785556100ea565b828001600101855582156100ea579182015b828111156100ea5782518255916020019190600101906100cf565b506100f69291506100fa565b5090565b5b808211156100f657600081556001016100fb565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561013857600080fd5b82516001600160a01b038116811461014f57600080fd5b602084810151919350906001600160401b038082111561016e57600080fd5b818601915086601f83011261018257600080fd5b8151818111156101945761019461010f565b604051601f8201601f19908116603f011681019083821181831017156101bc576101bc61010f565b8160405282815289868487010111156101d457600080fd5b600093505b828410156101f657848401860151818501870152928501926101d9565b828411156102075760008684830101525b8096505050505050509250929050565b600181811c9082168061022b57607f821691505b6020821081141561024c57634e487b7160e01b600052602260045260246000fd5b50919050565b61029c806102616000396000f3fe6080604081815230600090815260016020908152828220549082905291812063bf40fac160e01b909352916001600160a01b039091169063bf40fac19061004790608461017e565b60206040518083038186803b15801561005f57600080fd5b505afa158015610073573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100979190610226565b90506001600160a01b0381166100ff5760405162461bcd60e51b815260206004820152602360248201527f5461726765742061646472657373206d75737420626520696e697469616c697a60448201526232b21760e91b606482015260840160405180910390fd5b600080826001600160a01b031660003660405161011d929190610256565b600060405180830381855af49150503d8060008114610158576040519150601f19603f3d011682016040523d82523d6000602084013e61015d565b606091505b5090925090506001821515141561017657805160208201f35b805160208201fd5b600060208083526000845481600182811c9150808316806101a057607f831692505b8583108114156101be57634e487b7160e01b85526022600452602485fd5b8786018381526020018180156101db57600181146101ec57610217565b60ff19861682528782019650610217565b60008b81526020902060005b86811015610211578154848201529085019089016101f8565b83019750505b50949998505050505050505050565b60006020828403121561023857600080fd5b81516001600160a01b038116811461024f57600080fd5b9392505050565b818382376000910190815291905056fea2646970667358221220edb025fb8656bfe68e1f6fdcfa8ab0e0c6a5e51eccebff0b181072c10fc6096564736f6c63430008090033";

type Lib_ResolvedDelegateProxyConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: Lib_ResolvedDelegateProxyConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class Lib_ResolvedDelegateProxy__factory extends ContractFactory {
  constructor(...args: Lib_ResolvedDelegateProxyConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override deploy(
    _libAddressManager: PromiseOrValue<string>,
    _implementationName: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<Lib_ResolvedDelegateProxy> {
    return super.deploy(
      _libAddressManager,
      _implementationName,
      overrides || {}
    ) as Promise<Lib_ResolvedDelegateProxy>;
  }
  override getDeployTransaction(
    _libAddressManager: PromiseOrValue<string>,
    _implementationName: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(
      _libAddressManager,
      _implementationName,
      overrides || {}
    );
  }
  override attach(address: string): Lib_ResolvedDelegateProxy {
    return super.attach(address) as Lib_ResolvedDelegateProxy;
  }
  override connect(signer: Signer): Lib_ResolvedDelegateProxy__factory {
    return super.connect(signer) as Lib_ResolvedDelegateProxy__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): Lib_ResolvedDelegateProxyInterface {
    return new utils.Interface(_abi) as Lib_ResolvedDelegateProxyInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): Lib_ResolvedDelegateProxy {
    return new Contract(
      address,
      _abi,
      signerOrProvider
    ) as Lib_ResolvedDelegateProxy;
  }
}
