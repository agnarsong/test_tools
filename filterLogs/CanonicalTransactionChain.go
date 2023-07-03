package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	runType := flag.String("t", "fl", "the type of run, 'other' or 'filterLogs'")
	flag.Parse()

	file, err := os.Open("abi/contracts/L1/rollup/CanonicalTransactionChain.sol/CanonicalTransactionChain.json")
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(file)
	if err != nil {
		log.Fatal(err)
	}

	if *runType == "fl" {
		client, err := ethclient.Dial("https://rpc.ankr.com/eth")
		if err != nil {
			log.Fatal(err)
		}

		contractAddress := common.HexToAddress("0x291dc3819b863e19b0a9b9809F8025d2EB4aaE93")
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(17607410),
			ToBlock:   big.NewInt(17608410),
			Addresses: []common.Address{
				contractAddress,
			},
		}

		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			log.Fatal(err)
		}

		// method := "L2GasParamsUpdated"
		// dSig := []byte("L2GasParamsUpdated(uint256,uint256,uint256)")
		method := "TransactionEnqueued"
		dSig := []byte("TransactionEnqueued(address,address,uint256,bytes,uint256,uint256)")
		// method := "CTCBatchReset"
		// dSig := []byte("CTCBatchReset(uint256,uint40,uint40,uint40,uint40,uint40,uint40)")
		// method := "TransactionBatchAppended"
		// dSig := []byte("TransactionBatchAppended(uint256,bytes32,uint256,uint256,bytes,bytes)")
		// method := "SequencerBatchAppended"
		// dSig := []byte("SequencerBatchAppended(uint256,uint256,uint256)")

		dHash := crypto.Keccak256Hash(dSig)
		fmt.Println("dHash: ", dHash)

		for _, vLog := range logs {
			fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
			fmt.Printf("Log Index: %d\n", vLog.Index)

			switch vLog.Topics[0].Hex() {
			case dHash.Hex():
				fmt.Printf("Log Name: %v\n", method)
				fmt.Println("Topics[0]: ", vLog.Topics[0])
				for i, top := range vLog.Topics {
					fmt.Println(fmt.Sprintf("Topics[%v]", i), top)
				}

				logdata, err := contractAbi.Unpack(method, vLog.Data)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(logdata)

			default:
				fmt.Println("other Topics[0]: ", vLog.Topics[0].Hex())
			}

			fmt.Printf("\n\n")
		}
	} else {

	}
}
