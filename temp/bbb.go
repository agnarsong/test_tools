package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 连接以太坊节点
	client, err := ethclient.Dial("https://goerli.infura.io/v3/3f31ba1b9fc54dfd92e1e26a64fba7e0")
	if err != nil {
		log.Fatal(err)
	}

	// 合约ABI和地址
	contractABI := "[{\"inputs\":[],\"name\":\"getTotalElements\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalElements\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"
	// CanonicalTransactionChain
	contractAddress := common.HexToAddress("0x654e6dF111F98374d9e5d908D7a5392C308aA18D")

	// 创建ABI实例
	contractAbi, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		log.Fatal(err)
	}

	// 打包合约函数参数
	input, err := contractAbi.Pack("getTotalElements")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("input: ", input)
	// 调用合约函数
	output, err := client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contractAddress,
		Data: input,
	}, big.NewInt(8506706))
	if err != nil {
		log.Fatal(err)
	}

	// 解析合约函数返回值
	result, err := contractAbi.Unpack("getTotalElements", output)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
