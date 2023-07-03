package lib

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TransactionByHash(c *ethclient.Client, th string) (tx *types.Transaction, isPending bool, err error) {

	txHash := common.HexToHash(th)
	tx, isPending, err = c.TransactionByHash(context.Background(), txHash)
	if err != nil {
		return
	}

	return
}

func TransactionReceipt(c *ethclient.Client, th string) (tr *types.Receipt, err error) {

	txHash := common.HexToHash(th)
	tr, err = c.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return
	}

	return
}

func L2TransactionReceipt(c *ethclient.Client, th string) (tr *types.Receipt, err error) {

	txHash := common.HexToHash(th)
	tr, err = c.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return
	}

	return
}

type Res struct {
	Block *types.Block
	Tx    *types.Transaction
	Re    *types.Receipt
	From  string
	V     *big.Int
	R     *big.Int
	S     *big.Int
}

func TransactionByTo(c *ethclient.Client, startBlock uint64, endBlock uint64, address string) (rs []Res, err error) {

	// 设定要遍历的区块范围
	if endBlock == 0 {
		endBlock, err = getBlockNumberWithRetry(c)
		// fmt.Println("endBlock blockNumber: ", endBlock)
		if err != nil {
			return nil, fmt.Errorf("BlockNumber err: %v", err)
		}
	}

	if startBlock > endBlock {
		return rs, errors.New("startBlock > endBlock ")
	}

	for blockNumber := startBlock; blockNumber <= endBlock; blockNumber++ {
		block, err := getBlockByNumberWithRetry(c, blockNumber)
		if err != nil {
			return nil, fmt.Errorf("BlockByNumber err: %v", err)
		}

		// 遍历区块中的交易
		for _, tx := range block.Transactions() {
			if tx.To() != nil && tx.To().Hex() == address {
				re, err := getTransactionReceiptWithRetry(c, tx.Hash())
				if err != nil {
					return nil, fmt.Errorf("TransactionReceipt err: %v", err)
				}

				// 获取 v, r, s 值
				v, r, s := tx.RawSignatureValues()

				var fromAddress common.Address

				if r.String() != "0" {
					signer := types.LatestSignerForChainID(tx.ChainId())
					fromAddress, err = types.Sender(signer, tx)
					if err != nil {
						return nil, fmt.Errorf("types Sender: %v", err)
					}
				}

				rs = append(rs, Res{
					Block: block,
					Tx:    tx,
					Re:    re,
					From:  fromAddress.Hex(),
					V:     v,
					R:     r,
					S:     s,
				})
			}
		}
	}

	return rs, nil
}

func TransactionByFrom(c *ethclient.Client, startBlock uint64, endBlock uint64, address string) (rs []Res, err error) {

	// 设定要遍历的区块范围
	if endBlock == 0 {
		endBlock, err = c.BlockNumber(context.Background())
		// fmt.Println("endBlock blockNumber: ", endBlock)
		if err != nil {
			return nil, err
		}
	}

	for blockNumber := startBlock; blockNumber <= endBlock; blockNumber++ {
		block, err := c.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
		if err != nil {
			return nil, err
		}

		// 遍历区块中的交易
		for _, tx := range block.Transactions() {
			// 获取 'from' 地址
			// 注释掉的方法不支持 tx.Type() == 2
			// fromAddress, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)
			// if err != nil {
			// 	return nil, fmt.Errorf("types Sender: %v", err)
			// }
			signer := types.LatestSignerForChainID(tx.ChainId())
			fromAddress, err := types.Sender(signer, tx)
			if err != nil {
				return nil, fmt.Errorf("types Sender: %v", err)
			}

			if fromAddress.Hex() == address {
				re, err := c.TransactionReceipt(context.Background(), tx.Hash())
				if err != nil {
					return nil, err
				}

				// 获取 v, r, s 值
				v, r, s := tx.RawSignatureValues()
				rs = append(rs, Res{
					Block: block,
					Tx:    tx,
					Re:    re,
					From:  address,
					V:     v,
					R:     r,
					S:     s,
				})
			}
		}
	}

	return rs, nil
}

func AllTransactions(c *ethclient.Client, startBlock uint64, endBlock uint64) (rs []Res, err error) {

	// 设定要遍历的区块范围
	if endBlock == 0 {
		endBlock, err = c.BlockNumber(context.Background())
		// fmt.Println("endBlock blockNumber: ", endBlock)
		if err != nil {
			return nil, err
		}
	}

	for blockNumber := startBlock; blockNumber <= endBlock; blockNumber++ {
		block, err := c.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
		if err != nil {
			return nil, err
		}

		// 遍历区块中的交易
		for _, tx := range block.Transactions() {

			re, err := c.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				return nil, err
			}

			// fromAddress, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)
			// if err != nil {
			// 	return nil, fmt.Errorf("types Sender: %v", err)
			// }

			// 获取 v, r, s 值
			v, r, s := tx.RawSignatureValues()

			var fromAddress common.Address

			if r.String() != "0" {
				signer := types.LatestSignerForChainID(tx.ChainId())
				fromAddress, err = types.Sender(signer, tx)
				if err != nil {
					return nil, fmt.Errorf("types Sender: %v", err)
				}
			}

			rs = append(rs, Res{
				Block: block,
				Tx:    tx,
				Re:    re,
				From:  fromAddress.Hex(),
				V:     v,
				R:     r,
				S:     s,
			})

		}
	}

	return rs, nil
}

func getBlockNumberWithRetry(c *ethclient.Client) (uint64, error) {
	var bn uint64
	var err error
	maxRetries := 5

	for retry := 0; retry < maxRetries; retry++ {
		bn, err = c.BlockNumber(context.Background())
		if err != nil {
			if retry < maxRetries-1 {
				// 间隔一段时间后重试
				fmt.Println("getBlockNumberWithRetry")
				time.Sleep(time.Second * 5)
				continue
			}
			return 0, fmt.Errorf("BlockNumber err: %v", err)
		}
		break
	}

	return bn, nil
}

func getBlockByNumberWithRetry(c *ethclient.Client, blockNumber uint64) (*types.Block, error) {
	var block *types.Block
	var err error
	maxRetries := 5

	for retry := 0; retry < maxRetries; retry++ {
		block, err = c.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
		if err != nil {
			if retry < maxRetries-1 {
				// 间隔一段时间后重试
				fmt.Println("getBlockByNumberWithRetry")
				time.Sleep(time.Second * 5)
				continue
			}
			return nil, fmt.Errorf("BlockByNumber err: %v", err)
		}
		break
	}

	return block, nil
}

func getTransactionReceiptWithRetry(c *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	var receipt *types.Receipt
	var err error
	maxRetries := 5

	for retry := 0; retry < maxRetries; retry++ {
		receipt, err = c.TransactionReceipt(context.Background(), txHash)
		if err != nil {
			if retry < maxRetries-1 {
				// 间隔一段时间后重试
				fmt.Println("getTransactionReceiptWithRetry")
				time.Sleep(time.Second * 5)
				continue
			}
			return nil, fmt.Errorf("TransactionReceipt err: %v", err)
		}
		break
	}

	return receipt, nil
}
