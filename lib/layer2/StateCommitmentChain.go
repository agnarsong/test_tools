package layer2

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"mantle/test/lib/layer2/bindings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// StateCommitmentChainStateBatchAppended
// StateCommitmentChainStateBatchDeleted
// StateCommitmentChainDistributeTssReward
// StateCommitmentChainRollBackL2Chain
func GetAppendEvent(c *ethclient.Client, proxy_rollupAddr string, sccAddr string,
	height int64) *bindings.StateCommitmentChainStateBatchAppended {
	sccContract, err := bindings.NewStateCommitmentChain(
		common.HexToAddress(sccAddr), c,
	)
	if err != nil {
		panic(err.Error())
	}

	block, err := c.BlockByNumber(context.Background(), big.NewInt(height))
	if err != nil {
		panic(err.Error())
	}
	var appendEvent *bindings.StateCommitmentChainStateBatchAppended

	if block.Transactions().Len() != 0 {
		for _, v := range block.Transactions() {
			receipt, _ := c.TransactionReceipt(context.Background(), v.Hash())
			appendEvent = parseSCCAppendEvent(proxy_rollupAddr, v, receipt, sccContract)
			if appendEvent == nil {
			} else {
				fmt.Println("==========appendEvent=============")
				fmt.Println("block number: ", block.Number())
				fmt.Println("appendEvent.BatchIndex: ", appendEvent.BatchIndex)
				fmt.Println("appendEvent.BatchRoot: ", appendEvent.BatchRoot)
				fmt.Println("appendEvent.BatchRoot: ", hex.EncodeToString(appendEvent.BatchRoot[:]))
				fmt.Println("appendEvent.Signature: ", common.Bytes2Hex(appendEvent.Signature))
				fmt.Println("appendEvent.BatchSize: ", appendEvent.BatchSize)
				fmt.Println("appendEvent.ExtraData: ", common.Bytes2Hex(appendEvent.ExtraData))
				fmt.Println("appendEvent.PrevTotalElements: ", appendEvent.PrevTotalElements)
				fmt.Println("==========appendEvent=============")
			}
		}
		return appendEvent
	}
	return nil
}

func parseSCCAppendEvent(proxy_rollup string, tx *types.Transaction, receipt *types.Receipt, scc *bindings.StateCommitmentChain) *bindings.StateCommitmentChainStateBatchAppended {
	if tx.To() != nil {
		if strings.Compare(tx.To().String(), proxy_rollup) == 0 {
			for _, vlog := range receipt.Logs {
				if append, err := scc.ParseStateBatchAppended(*vlog); err == nil {
					return append
				} else {
				}
			}
		}
	}
	return nil
}
