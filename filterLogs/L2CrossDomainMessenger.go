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
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	runType := flag.String("t", "fl", "the type of run, 'other' or 'filterLogs'")
	flag.Parse()

	file, err := os.Open("abi/contracts/L2/messaging/L2CrossDomainMessenger.sol/L2CrossDomainMessenger.json")
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(file)
	if err != nil {
		log.Fatal(err)
	}

	if *runType == "fl" {
		client, err := ethclient.Dial("https://rpc4blockscout.mainnet.gomantle.org")
		if err != nil {
			log.Fatal(err)
		}

		contractAddress := common.HexToAddress("0x4200000000000000000000000000000000000007")
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

		// method := "ETHDepositInitiated"
		// dSig := []byte("ETHDepositInitiated(address,address,uint256,bytes)")
		// method := "ERC20DepositInitiated"
		// dSig := []byte("ERC20DepositInitiated(address,address,address,address,uint256,bytes)")
		// method := "WithdrawalInitiated"
		// dSig := []byte("WithdrawalInitiated(address,address,address,address,uint256,bytes)")
		method := "FailedRelayedMessage"
		dSig := []byte("FailedRelayedMessage(bytes32)")
		// method := "RelayedMessage"
		// dSig := []byte("RelayedMessage(bytes32)")
		// method := "SentMessage"
		// dSig := []byte("SentMessage(address,address,bytes,uint256,uint256)")

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
	} else {

		_l1Token := common.HexToAddress("0x0000000000000000000000000000000000000000")
		BVM_ETH := common.HexToAddress("0xdEAddEaDdeadDEadDEADDEAddEADDEAddead1111")
		_from := common.HexToAddress("0xa42D235E91C378c982e31877868511DCdaA352E8")
		_to := common.HexToAddress("0xa42D235E91C378c982e31877868511DCdaA352E8")
		value := big.NewInt(1)             // replace with msg.value
		_data, err := hexutil.Decode("0x") // replace with _data
		if err != nil {
			log.Fatal(err)
		}

		callData, err := contractAbi.Pack("finalizeDeposit", _l1Token, BVM_ETH, _from, _to, value, _data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(hexutil.Encode(callData))
	}
}
