package main

import (
	"context"
	"flag"
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

	runType := flag.String("t", "fl", "the type of run, 'other' or 'filterLogs'")
	flag.Parse()

	file, err := os.Open("abi/contracts/L1/rollup/StateCommitmentChain.sol/StateCommitmentChain.json")
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(file)
	if err != nil {
		log.Fatal(err)
	}

	if *runType == "fl" {
		client, err := ethclient.Dial("https://goerli.infura.io/v3/3f31ba1b9fc54dfd92e1e26a64fba7e0")
		if err != nil {
			log.Fatal(err)
		}

		contractAddress := common.HexToAddress("0x9256B84db2425973E582E31989c1E820dCeCE716")
		query := ethereum.FilterQuery{
			// FromBlock: big.NewInt(9253868),
			FromBlock: big.NewInt(9267309),
			Addresses: []common.Address{
				contractAddress,
			},
		}

		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			log.Fatal(err)
		}

		// method := "StateBatchAppended"
		// dSig := []byte("StateBatchAppended(uint256,bytes32,uint256,uint256,bytes,bytes)")
		// method := "DistributeTssReward"
		// dSig := []byte("DistributeTssReward(uint256,uint256,uint256,address[])")
		method := "RollBackL2Chain"
		dSig := []byte("RollBackL2Chain(uint256)")

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

				for i, top := range vLog.Topics {
					fmt.Println(fmt.Sprintf("Topics[%v]", i), top)
				}
				fmt.Println(common.Bytes2Hex(vLog.Data))

				output, err := lib.ParseContractFunctionOutputs1(contractAbi, method, vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("output: ")
				lib.PrintMap(output)
				fmt.Println(common.Bytes2Hex(output["_extraData"].([]byte)))

			default:
				fmt.Println("default: ", dHash.Hex())
			}

			fmt.Printf("\n\n")
		}
	} else {

	}
}
