import { L1ChainID, L2ChainID } from '../interfaces';
export declare const DEPOSIT_CONFIRMATION_BLOCKS: {
    [ChainID in L2ChainID]: number;
};
export declare const CHAIN_BLOCK_TIMES: {
    [ChainID in L1ChainID]: number;
};
