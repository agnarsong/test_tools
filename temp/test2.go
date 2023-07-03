package main

import (
	"context"
	"encoding/hex"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type TransactionInfo struct {
	Hash             string   `json:"hash"`
	Nonce            uint64   `json:"nonce"`
	BlockHash        string   `json:"blockHash"`
	BlockNumber      uint64   `json:"blockNumber"`
	TransactionIndex uint64   `json:"transactionIndex"`
	From             string   `json:"from"`
	To               string   `json:"to"`
	Value            *big.Int `json:"value"`
	GasPrice         *big.Int `json:"gasPrice"`
	Gas              uint64   `json:"gas"`
	Input            string   `json:"input"`
}

type NestedTransactionInfo struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Value    string `json:"value"`
	Gas      uint64 `json:"gas"`
	GasPrice string `json:"gasPrice"`
	Input    string `json:"input"`
	Nonce    uint64 `json:"nonce"`
}

func main() {
	// 创建以太坊客户端
	client, err := ethclient.Dial("http://127.0.0.1:9545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// 创建RPC客户端
	rpcClient, err := rpc.Dial("http://127.0.0.1:9545")
	if err != nil {
		log.Fatalf("Failed to connect to the RPC client: %v", err)
	}

	// 获取最新的块头
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to get block header: %v", err)
	}

	// 从最新的块开始向前遍历
	blockNumber := header.Number.Uint64()
	for blockNumber > 0 {
		// 获取块信息
		block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
		if err != nil {
			log.Fatalf("Failed to get block: %v", err)
		}

		// 遍历块中的所有交易
		for _, tx := range block.Transactions() {
			txInfo, err := getTransactionInfo(tx, client, rpcClient)
			if err != nil {
				log.Printf("Failed to get transaction info for tx %v: %v", tx.Hash().Hex(), err)
				continue
			}

			// 检查交易是否有内部交易
			if len(txInfo.Input) > 2 {
				log.Printf("Tx %v has nested transaction(s)", tx.Hash().Hex())

				// 解析交易数据
				var inputData []byte
				inputData, err = hex.DecodeString(txInfo.Input[2:])
				if err != nil {
					log.Printf("Failed to decode tx input data for tx %v: %v", tx.Hash().Hex(), err)
					continue
				}

				if len(inputData) < 4 {
					log.Printf("Invalid input data length for tx %v", tx.Hash().Hex())
					continue
				}

				// 获取内部交易信息
				var nestedTxHash common.Hash
				copy(nestedTxHash[:], inputData[:32])
				nestedTx, isPending, err := client.TransactionByHash(context.Background(), nestedTxHash)
				if err != nil {
					log.Printf("Failed to get nested transaction for tx %v: %v", tx.Hash().Hex(), err)
					continue
				}

				// 获取内部交易的发送方和接收方
				var nestedFrom, nestedTo string
				if nestedTx.To() == nil {
					nestedFrom = nestedTx.From().Hex()
				} else {
					nestedTo = nestedTx.To().Hex()
				}
				log.Printf("Nested transaction in tx %v: From: %v, To: %v, Value: %v", tx.Hash().Hex(), nestedFrom, nestedTo, nestedTx.Value().String())

				// 获取内部交易的内部交易
				if len(nestedTx.Input()) > 2 {
					log.Printf("Nested transaction in tx %v has nested transaction(s)", tx.Hash().Hex())

					// 解析交易数据
					inputData := nestedTx.Input()[2:]
					if len(inputData) < 4 {
						log.Printf("Invalid input data length for tx %v", nestedTx.Hash().Hex())
						continue
					}

					// 获取内部交易信息
					var nestedTxHash common.Hash
					copy(nestedTxHash[:], inputData[:32])
					nestedTx, isPending, err = client.TransactionByHash(context.Background(), nestedTxHash)
					if err != nil {
						log.Printf("Failed to get nested transaction for tx %v: %v", tx.Hash().Hex(), err)
						continue
					}

					// 获取内部交易的发送方和接收方
					if nestedTx.To() == nil {
						nestedFrom = nestedTx.From().Hex()
					} else {
						nestedTo = nestedTx.To().Hex()
					}
					log.Printf("Nested transaction in tx %v: From: %v, To: %v, Value: %v", tx.Hash().Hex(), nestedFrom, nestedTo, nestedTx.Value().String())
				}
			}
		}
	}
}
