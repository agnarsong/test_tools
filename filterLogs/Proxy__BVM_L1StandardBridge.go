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

	contractAddress := common.HexToAddress("0x1e51308584567A501d6dA2c3527aC1583ABB920A")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(9253868),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("abi/contracts/L1/messaging/L1StandardBridge.sol/L1StandardBridge.json")
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(file)
	if err != nil {
		log.Fatal(err)
	}

	// method := "ETHDepositInitiated"
	// dSig := []byte("ETHDepositInitiated(address,address,uint256,bytes)")
	// method := "ERC20DepositInitiated"
	// dSig := []byte("ERC20DepositInitiated(address,address,address,address,uint256,bytes)")
	// method := "ETHWithdrawalFinalized"
	// dSig := []byte("ETHWithdrawalFinalized(address,address,uint256,bytes)")
	method := "ERC20WithdrawalFinalized"
	dSig := []byte("ERC20WithdrawalFinalized(address,address,address,address,uint256,bytes)")

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
			// lib.PrintMap(output)
			fmt.Println("hash.len: ", len(common.Bytes2Hex(output["_data"].([]byte))))
			fmt.Println("string.len: ", len(string(output["_data"].([]byte))))

		default:
			fmt.Println("default: ", dHash.Hex())
		}

		fmt.Printf("\n\n")
	}
}
