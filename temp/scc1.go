package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 连接以太坊节点
	// client, err := ethclient.Dial("https://goerli.infura.io/v3/3f31ba1b9fc54dfd92e1e26a64fba7e0")
	client, err := ethclient.Dial("http://localhost:9545")
	if err != nil {
		log.Fatal(err)
	}

	// ERC20 代币合约地址和要查询余额的地址
	// contractAddress := common.HexToAddress("0x91A5D806BA73d0AA4bFA9B318126dDE60582e92a") // 合约地址
	contractAddress := common.HexToAddress("0x8BAccFF561FDe61D6bC8B6f299fFBa561d2189B9") // 合约地址

	// 计算 balanceOf 方法的函数选择器
	// methodID := crypto.Keccak256([]byte("FRAUD_PROOF_WINDOW()"))[:4]
	// methodID := crypto.Keccak256([]byte("SEQUENCER_PUBLISH_WINDOW()"))[:4]
	// methodID := crypto.Keccak256([]byte("getLastSequencerTimestamp()"))[:4]
	// methodID := crypto.Keccak256([]byte("getTotalBatches()"))[:4]
	// methodID := crypto.Keccak256([]byte("getTotalElements()"))[:4]
	methodID := crypto.Keccak256([]byte("libAddressManager()"))[:4]

	// 拼接参数，将地址左侧补零
	var data []byte
	data = append(data, methodID...)
	// data = append(data, common.LeftPadBytes([]byte{}, 32)...)

	// fmt.Println("data: ", common.Bytes2Hex(data))
	// 调用合约方法获取余额
	result, err := client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 将余额转换为字符串并输出
	fmt.Println(result)
	fmt.Println(common.Bytes2Hex(result))
}
