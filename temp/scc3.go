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
	client, err := ethclient.Dial("https://goerli.infura.io/v3/3f31ba1b9fc54dfd92e1e26a64fba7e0")
	// client, err := ethclient.Dial("https://eth-goerli.g.alchemy.com/v2/vX39fRJe5UvWip7sAj_nmSO9p7Z5Y2KI")
	if err != nil {
		log.Fatal(err)
	}

	// 智能合约地址
	contractAddress := common.HexToAddress("0x5a94Dc6cc85fdA49d8E9A8b85DDE8629025C42be")

	// 部分ABI
	abiJSON := `
        [
			{
				"inputs": [
				  {
					"internalType": "address",
					"name": "recipient",
					"type": "address"
				  },
				  {
					"internalType": "uint256",
					"name": "amount",
					"type": "uint256"
				  }
				],
				"name": "transfer",
				"outputs": [
				  {
					"internalType": "bool",
					"name": "",
					"type": "bool"
				  }
				],
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
	inputData, err := contractAbi.Pack("transfer", common.HexToAddress("0xD9f577a9Bce09Ec29bD661231e874A95fcDdB827"), big.NewInt(10))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inputData: ", common.Bytes2Hex(inputData))

	// 交易发送者的地址和私钥
	privateKey, err := crypto.HexToECDSA("e7bbfa2e0a24248701d7de1745495d8c7e08a835782006db5dddc37442f2ccb2")
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
	gasLimit := uint64(30000000) // 填写你认为合适的gasLimit
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("SuggestGasPrice: ", err)
	}
	// gasPrice.Add(gasPrice, big.NewInt(77927815949))
	fmt.Println("gasPrice: ", gasPrice)
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

	fmt.Println("Transaction sent: ", signedTx.Hash().Hex())
}
