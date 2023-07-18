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
	client, err := ethclient.Dial("http://127.0.0.1:9545")
	if err != nil {
		log.Fatal(err)
	}

	// ERC20 代币合约地址和要查询余额的地址
	contractAddress := common.HexToAddress("0x92aBAD50368175785e4270ca9eFd169c949C4ce1") // 合约地址
	address := common.HexToAddress("0x52753615226F8aC8a464bfecb11Ef798CFF3793f")         // 要查询余额的地址

	// 计算 balanceOf 方法的函数选择器
	methodID := crypto.Keccak256([]byte("balanceOf(address)"))[:4]

	// 拼接参数，将地址左侧补零
	var data []byte
	data = append(data, methodID...)
	data = append(data, common.LeftPadBytes(address.Bytes(), 32)...)

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
	fmt.Printf("Balance of %s: %s USDT\n", address.Hex(), common.BytesToHash(result).Big())
}
