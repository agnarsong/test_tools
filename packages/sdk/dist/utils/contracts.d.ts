import { ethers, Contract } from 'ethers';
import { DeepPartial } from './type-utils';
import { OEContracts, OEL1Contracts, OEL2Contracts, OEContractsLike, OEL2ContractsLike, AddressLike, BridgeAdapters, BridgeAdapterData, ICrossChainMessenger, L2ChainID } from '../interfaces';
export declare const DEFAULT_L2_CONTRACT_ADDRESSES: OEL2ContractsLike;
export declare const CONTRACT_ADDRESSES: {
    [ChainID in L2ChainID]: OEContractsLike;
};
export declare const BRIDGE_ADAPTER_DATA: {
    [ChainID in L2ChainID]?: BridgeAdapterData;
};
export declare const getOEContract: (contractName: keyof OEL1Contracts | keyof OEL2Contracts, l2ChainId: number, opts?: {
    address?: AddressLike;
    signerOrProvider?: ethers.Signer | ethers.providers.Provider;
}) => Contract;
export declare const getAllOEContracts: (l2ChainId: number, opts?: {
    l1SignerOrProvider?: ethers.Signer | ethers.providers.Provider;
    l2SignerOrProvider?: ethers.Signer | ethers.providers.Provider;
    overrides?: DeepPartial<OEContractsLike>;
}) => OEContracts;
export declare const getBridgeAdapters: (l2ChainId: number, messenger: ICrossChainMessenger, opts?: {
    overrides?: BridgeAdapterData;
}) => BridgeAdapters;
