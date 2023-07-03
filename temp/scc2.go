package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 连接以太坊节点
	client, err := ethclient.Dial("http://localhost:9545")
	if err != nil {
		log.Fatal(err)
	}

	// 智能合约地址
	contractAddress := common.HexToAddress("0x8BAccFF561FDe61D6bC8B6f299fFBa561d2189B9")

	// 部分ABI
	abiJSON := `
        [
            {
                "inputs": [
                    {
                        "internalType": "uint256",
                        "name": "windowSize",
                        "type": "uint256"
                    }
                ],
                "name": "setFraudProofWindow",
                "outputs": [],
                "stateMutability": "nonpayable",
                "type": "function"
            }
        ]
    `

	// 解析部分ABI
	contractAbi, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		log.Fatal(err)
	}

	// 调用合约函数
	inputData, err := contractAbi.Pack("setFraudProofWindow", big.NewInt(10))
	if err != nil {
		log.Fatal(err)
	}

	// 交易发送者的地址和私钥
	privateKey, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// 创建调用合约函数的交易
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("PendingNonceAt: ", err)
	}
	gasLimit := uint64(1000000000) // 填写你认为合适的gasLimit
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("SuggestGasPrice: ", err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal("ChainID: ", err)
	}
	// 创建交易
	tx := types.NewTransaction(nonce, contractAddress, big.NewInt(0), gasLimit, gasPrice, inputData)

	// 签名交易
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transaction sent: %s", signedTx.Hash().Hex())
}
