package lib

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func BlockNumber(c *ethclient.Client) (bn uint64, err error) {

	bn, err = c.BlockNumber(context.Background())
	if err != nil {
		return bn, err
	}

	return
}

func HeaderByNumber(c *ethclient.Client, bn int64) (header *types.Header, err error) {

	blockNumber := big.NewInt(bn)
	if bn == -1 {
		blockNumber = nil
	}

	header, err = c.HeaderByNumber(context.Background(), blockNumber)
	if err != nil {
		return header, err
	}

	return
}

func HeaderByHash(c *ethclient.Client, bh string) (header *types.Header, err error) {

	blockHash := common.HexToHash(bh)
	header, err = c.HeaderByHash(context.Background(), blockHash)
	if err != nil {
		return header, err
	}

	return
}

func BlockByNumber(c *ethclient.Client, bn int64) (block *types.Block, err error) {

	if bn == -1 {
		bnt, err := BlockNumber(c)
		if err != nil {
			return block, err
		}
		bn = int64(bnt)
	}

	blockNumber := big.NewInt(bn)
	block, err = c.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		return block, err
	}

	return
}

func BlockByHash(c *ethclient.Client, bh string) (block *types.Block, err error) {

	blockHash := common.HexToHash(bh)
	block, err = c.BlockByHash(context.Background(), blockHash)
	if err != nil {
		return block, err
	}

	return
}

func TransactionCount(c *ethclient.Client, bh string) (count uint, err error) {

	block, err := BlockByHash(c, bh)
	if err != nil {
		return count, err
	}

	count, err = c.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		return count, err
	}

	return
}

func FindBlockNumberByStateRoot(c *ethclient.Client, startBlock, endBlock uint64, targetStateRoot common.Hash) (*types.Block, error) {
	for blockNumber := startBlock; blockNumber <= endBlock; blockNumber++ {

		block, err := c.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
		if err != nil {
			return nil, err
		}

		if block.Root() == targetStateRoot {
			// // header 的stateRoot 和 block的 Root() 返回的内容一致
			// h, err := HeaderByHash(c, block.Hash().Hex())
			// if err != nil {
			// 	return 0, nil
			// }
			// fmt.Println(targetStateRoot)
			// fmt.Println(h.Root.Hex())
			// fmt.Println(block.Root())

			return block, nil

		}
	}
	return nil, fmt.Errorf("target stateRoot not found")
}
