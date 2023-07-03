package main

import (
	"context"
	"fmt"
	"log"
	"mantle/test/lib"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://goerli.infura.io/v3/3f31ba1b9fc54dfd92e1e26a64fba7e0")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x476F6F17cAeA1C9035e569789F22e798742d26Ab")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(9267309),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("abi/contracts/L1/messaging/L1CrossDomainMessenger.sol/L1CrossDomainMessenger.json")
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(file)
	if err != nil {
		log.Fatal(err)
	}

	// method := "SentMessage"
	// dSig := []byte("SentMessage(address,address,bytes,uint256,uint256)")
	// method := "RelayedMessage"
	// dSig := []byte("RelayedMessage(bytes32)")
	method := "FailedRelayedMessage"
	dSig := []byte("FailedRelayedMessage(bytes32)")
	// method := "MessageBlocked"
	// dSig := []byte("MessageBlocked(bytes32)")
	// method := "MessageAllowed"
	// dSig := []byte("MessageAllowed(bytes32)")
	// method := "Paused"
	// dSig := []byte("Paused(address)")

	dHash := crypto.Keccak256Hash(dSig)
	fmt.Println("dHash: ", dHash)

	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)

		switch vLog.Topics[0].Hex() {
		case dHash.Hex():
			fmt.Printf("Log Name: %v\n", method)

			for i, top := range vLog.Topics {
				fmt.Println(fmt.Sprintf("Topics[%v]", i), top)
			}

			output, err := lib.ParseContractFunctionOutputs1(contractAbi, method, vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("output: ")
			lib.PrintMap(output)

		default:
			fmt.Println("default: ", dHash.Hex())
		}

		fmt.Printf("\n\n")
	}
}
