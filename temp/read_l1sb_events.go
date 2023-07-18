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

type L1cdmFailedRelayedMessage struct {
	MsgHash [32]byte
}

type L1cdmInitialized struct {
	Version uint8
}

type L1cdmMessageAllowed struct {
	XDomainCalldataHash [32]byte
}

type L1cdmMessageBlocked struct {
	XDomainCalldataHash [32]byte
}

type L1cdmOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
}

type L1cdmPaused struct {
	Account common.Address
}

type L1cdmRelayedMessage struct {
	MsgHash [32]byte
}

type L1cdmSentMessage struct {
	Target       common.Address
	Sender       common.Address
	Message      []byte
	MessageNonce *big.Int
	GasLimit     *big.Int
}

type L1cdmUnpaused struct {
	Account common.Address
}

func main() {
	client, err := ethclient.Dial("https://goerli.infura.io/v3/3f31ba1b9fc54dfd92e1e26a64fba7e0")
	if err != nil {
		log.Fatal(err)
	}

	// 0x Protocol (ZRX) token address
	contractAddress := common.HexToAddress("0x3917B2c9258F3398D69cF2C231f0F446e3ad50E2")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(9332145),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	// contractAbi, err := abi.JSON(strings.NewReader(string(l1cdm.L1cdmABI)))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// FailedRelayedMessageSig := []byte("FailedRelayedMessage(bytes32)")
	// FailedRelayedMessageSigHash := crypto.Keccak256Hash(FailedRelayedMessageSig)

	// InitializedSig := []byte("Initialized(uint8)")
	// InitializedSigHash := crypto.Keccak256Hash(InitializedSig)

	// MessageAllowedSig := []byte("MessageAllowed(bytes32)")
	// MessageAllowedSigHash := crypto.Keccak256Hash(MessageAllowedSig)

	// MessageBlockedSig := []byte("MessageBlocked(bytes32)")
	// MessageBlockedSigHash := crypto.Keccak256Hash(MessageBlockedSig)

	// OwnershipTransferredSig := []byte("OwnershipTransferred(address,address)")
	// OwnershipTransferredSigHash := crypto.Keccak256Hash(OwnershipTransferredSig)

	// PausedSig := []byte("Paused(address)")
	// PausedSigHash := crypto.Keccak256Hash(PausedSig)

	// RelayedMessageSig := []byte("RelayedMessage(bytes32)")
	// RelayedMessageSigHash := crypto.Keccak256Hash(RelayedMessageSig)

	// SentMessageSig := []byte("SentMessage(address,address,bytes,uint256,uint256)")
	// SentMessageSigHash := crypto.Keccak256Hash(SentMessageSig)

	// UnpausedSig := []byte("Unpaused(address)")
	// UnpausedSigHash := crypto.Keccak256Hash(UnpausedSig)

	// // 获取合约 ABI
	// contractAbi, err := getContractABI(client, contractAddress)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	contractAbi, err := abi.JSON(strings.NewReader(string("[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]")))
	if err != nil {
		log.Fatal(err)
	}

	// SequencerBatchAppendedSig := []byte("SequencerBatchAppended(uint256,uint256,uint256)")
	// SequencerBatchAppendedSigHash := crypto.Keccak256Hash(SequencerBatchAppendedSig)

	// TransactionBatchAppendedSig := []byte("TransactionBatchAppended(uint256,bytes32,uint256,uint256,bytes,bytes)")
	// TransactionBatchAppendedSigHash := crypto.Keccak256Hash(TransactionBatchAppendedSig)
	OwnershipTransferredSig := []byte("OwnershipTransferred(address)")
	OwnershipTransferredHash := crypto.Keccak256Hash(OwnershipTransferredSig)

	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)

		switch vLog.Topics[0].Hex() {
		// case FailedRelayedMessageSigHash.Hex():
		// 	fmt.Printf("Log Name: FailedRelayedMessage\n")
		// 	logdata, err := contractAbi.Unpack("FailedRelayedMessage", vLog.Data)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	fmt.Println(logdata)
		// case InitializedSigHash.Hex():
		// 	fmt.Printf("Log Name: Initialized\n")
		// 	logdata, err := contractAbi.Unpack("Initialized", vLog.Data)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	fmt.Println(logdata)
		// case MessageAllowedSigHash.Hex():
		// 	fmt.Printf("Log Name: MessageAllowed\n")
		// 	logdata, err := contractAbi.Unpack("MessageAllowed", vLog.Data)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	fmt.Println(logdata)
		// case MessageBlockedSigHash.Hex():
		// 	fmt.Printf("Log Name: MessageBlocked\n")
		// 	logdata, err := contractAbi.Unpack("MessageBlocked", vLog.Data)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	fmt.Println(logdata)
		// case OwnershipTransferredSigHash.Hex():
		// 	fmt.Printf("Log Name: OwnershipTransferred\n")
		// 	logdata, err := contractAbi.Unpack("OwnershipTransferred", vLog.Data)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	fmt.Println(logdata)
		// case PausedSigHash.Hex():
		// 	fmt.Printf("Log Name: Paused\n")
		// 	logdata, err := contractAbi.Unpack("Paused", vLog.Data)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	fmt.Println(logdata)
		// case RelayedMessageSigHash.Hex():
		// 	fmt.Printf("Log Name: RelayedMessage\n")
		// 	logdata, err := contractAbi.Unpack("RelayedMessage", vLog.Data)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	fmt.Println(logdata)
		// case SentMessageSigHash.Hex():
		// 	fmt.Printf("Log Name: SentMessage\n")
		// 	logdata, err := contractAbi.Unpack("SentMessage", vLog.Data)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	fmt.Println(logdata)
		// case UnpausedSigHash.Hex():
		// 	fmt.Printf("Log Name: Unpaused\n")
		// 	logdata, err := contractAbi.Unpack("Unpaused", vLog.Data)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	fmt.Println(logdata)

		// case SequencerBatchAppendedSigHash.Hex():
		// 	fmt.Printf("Log Name: SequencerBatchAppended\n")
		// 	logdata, err := contractAbi.Unpack("SequencerBatchAppended", vLog.Data)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	fmt.Println(logdata)

		// case TransactionBatchAppendedSigHash.Hex():
		// 	fmt.Printf("Log Name: TransactionBatchAppended\n")
		// 	logdata, err := contractAbi.Unpack("TransactionBatchAppended", vLog.Data)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	fmt.Println(logdata)

		// 	fmt.Println(common.Bytes2Hex(vLog.Data))

		case OwnershipTransferredHash.Hex():
			fmt.Printf("Log Name: SentMessage\n")
			logdata, err := contractAbi.Unpack("SentMessage", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(logdata)

			fmt.Println(common.Bytes2Hex(vLog.Data))
		}

		fmt.Printf("\n\n")
	}
}
