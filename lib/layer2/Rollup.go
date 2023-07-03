package layer2

import (
	"context"
	"fmt"

	"mantle/test/lib"
	"mantle/test/lib/layer2/bindings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func RejectLatestCreatedAssertionWithBatch(c *ethclient.Client, proxy_rollupAddr string, sccAddr string,
	height int64, proKeyStr string) {

	rollupContract, err := bindings.NewRollup(common.HexToAddress(proxy_rollupAddr), c)
	if err != nil {
		panic(err.Error())
	}

	appendEvent := GetAppendEvent(c, proxy_rollupAddr, sccAddr, height)

	batch := bindings.LibBVMCodecChainBatchHeader{
		BatchIndex:        appendEvent.BatchIndex,
		BatchRoot:         appendEvent.BatchRoot,
		BatchSize:         appendEvent.BatchSize,
		PrevTotalElements: appendEvent.PrevTotalElements,
		Signature:         appendEvent.Signature,
		ExtraData:         appendEvent.ExtraData,
	}

	auth, _ := lib.GetAuth(c, proKeyStr)
	// opts.NoSend = true

	tx, err := rollupContract.RejectLatestCreatedAssertionWithBatch(auth, batch)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("=================send transaction begin==========================")
	c.SendTransaction(context.Background(), tx)
	fmt.Println("=================send transaction end==========================")

}
