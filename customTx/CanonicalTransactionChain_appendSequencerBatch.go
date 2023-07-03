package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Connect to an Ethereum client (replace with your Ethereum client URL)
	client, err := ethclient.Dial("https://goerli.infura.io/v3/3f31ba1b9fc54dfd92e1e26a64fba7e0")
	if err != nil {
		fmt.Errorf("Dial err: %v", err)
	}

	// Replace with the address of the contract you want to interact with
	contractAddress := common.HexToAddress("0x4022243F55675F33C8159A16Be7a22264ACBB545")

	// Replace with the private key of the sender's Ethereum account
	privateKey, err := crypto.HexToECDSA("e7bbfa2e0a24248701d7de1745495d8c7e08a835782006db5dddc37442f2ccb2")
	if err != nil {
		fmt.Errorf("HexToECDSA err: %v", err)
	}

	// Get the current nonce for the sender's account
	// (to avoid replay attacks and ensure the transaction is unique)
	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(privateKey.PublicKey))
	if err != nil {
		fmt.Errorf("PendingNonceAt err: %v", err)
	}

	// Create the function signature and input data for the appendSequencerBatch function
	// We need to right-pad the data to match the required length for each parameter
	functionSignature := []byte("appendSequencerBatch()")
	shouldStartAtElement := make([]byte, 32)  // uint40 (8 bytes), right-padded with zeros
	totalElementsToAppend := make([]byte, 32) // uint24 (3 bytes), right-padded with zeros
	numContexts := make([]byte, 32)           // uint24 (3 bytes), right-padded with zeros
	inputData := append(functionSignature, shouldStartAtElement...)
	inputData = append(inputData, totalElementsToAppend...)
	inputData = append(inputData, numContexts...)

	// Create the transaction
	gasLimit := uint64(300000)         // You may need to adjust the gas limit accordingly
	gasPrice := big.NewInt(1000000000) // 1 Gwei
	recipient := contractAddress
	value := big.NewInt(0) // No Ether value sent with the transaction

	tx := ethereum.CallMsg{
		To:   &recipient,
		Data: inputData,
	}

	h := types.NewEIP155Signer(big.NewInt(5)).Hash(types.NewTransaction(nonce, recipient, value, gasLimit, gasPrice, tx.Data))
	sig, err := crypto.Sign(h[:], privateKey)
	if err != nil {
		fmt.Errorf("Sign err: %v", err)
	}
	fmt.Println("sig: ", common.Bytes2Hex(sig))

	// Sign the transaction
	signedTx, err := types.SignTx(types.NewTransaction(nonce, recipient, value, gasLimit, gasPrice, tx.Data), types.NewEIP155Signer(big.NewInt(5)), privateKey)
	if err != nil {
		fmt.Errorf("SignTx err: %v", err)
	}

	// Send the transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Errorf("SendTransaction err: %v", err)
	}

	fmt.Printf("Transaction hash: 0x%x\n", signedTx.Hash())
}
