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
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// client, err := ethclient.Dial("http://127.0.0.1:9545")
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	// client, err := ethclient.Dial("https://rpc.testnet.mantle.xyz/")
	// client, err := ethclient.Dial("https://mantle-l2geth.qa.davionlabs.com")
	// client, err := ethclient.Dial("https://goerli.infura.io/v3/3f31ba1b9fc54dfd92e1e26a64fba7e0")
	// client, err := ethclient.Dial("http://dev-bitnetwork-eth-goerli.node.bybchain.io")

	if err != nil {
		log.Fatal(err)
	}

	// 0x Protocol (ZRX) token address
	contractAddress := common.HexToAddress("0x4200000000000000000000000000000000000020")
	// contractAddress := common.HexToAddress("0xe6cd9e7b620964bECd42c7Ad41e56724f515E284")
	// contractAddress := common.HexToAddress("0x7A8B94a9fA2bb0581D2EEA2fEd875FCA97494612")
	// contractAddress := common.HexToAddress("0xa647f5947c50248bc4b2ef773791c9c2bc01c65a")

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string("[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastBatchTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tssMembers\",\"type\":\"address[]\"}],\"name\":\"DistributeTssReward\",\"type\":\"event\"}]")))
	// contractAbi, err := abi.JSON(strings.NewReader(string("[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contract IDelegationCallback\",\"name\":\"delegationTerms\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"returnData\",\"type\":\"bytes32\"}],\"name\":\"OnDelegationWithdrawnCallFailure\",\"type\":\"event\"}]")))
	// contractAbi, err := abi.JSON(strings.NewReader(string("[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegatior\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"DelegateTo\",\"type\":\"event\"}]")))
	// contractAbi, err := abi.JSON(strings.NewReader(string("[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldAddress\",\"type\":\"address\"}],\"name\":\"AddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldAddress\",\"type\":\"address\"}],\"name\":\"AddressSet\",\"type\":\"event\"}]")))

	if err != nil {
		log.Fatal(err)
	}

	method := "DistributeTssReward"
	dSig := []byte("DistributeTssReward(uint256,uint256,uint256,address[])")
	// method := "OnDelegationWithdrawnCallFailure"
	// dSig := []byte("OnDelegationWithdrawnCallFailure(address,bytes32)")
	// method := "DelegateTo"
	// dSig := []byte("DelegateTo(address,address)")
	// method := "AddressSet"
	// dSig := []byte("AddressSet(string,address,address)")

	dHash := crypto.Keccak256Hash(dSig)
	fmt.Println("dHash: ", dHash)

	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)

		switch vLog.Topics[0].Hex() {
		case dHash.Hex():
			fmt.Printf("Log Name: %v\n", method)
			logdata, err := contractAbi.Unpack(method, vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(logdata)

			fmt.Println(common.Bytes2Hex(vLog.Data))

		default:
			fmt.Println("default: ", dHash.Hex())
		}

		fmt.Printf("\n\n")
	}
}
