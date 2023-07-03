package main

import (
	"context"
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
	client, err := ethclient.Dial("https://rpc.ankr.com/eth")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x6968f3F16C3e64003F02E121cf0D5CCBf5625a42")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(17577534),
		ToBlock:   big.NewInt(17579534),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("abi/contracts/libraries/resolver/Lib_AddressManager.sol/Lib_AddressManager.json")
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(file)
	if err != nil {
		log.Fatal(err)
	}

	// method := "AddressSet"
	// dSig := []byte("AddressSet(string,address,address)")
	method := "OwnershipTransferred"
	dSig := []byte("OwnershipTransferred(address,address)")

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

			logdata, err := contractAbi.Unpack(method, vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(logdata)

		default:
			fmt.Println("default: ", dHash.Hex())
		}

		fmt.Printf("\n\n")
	}
}
