package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"mantle/test/contracts/Scc"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://localhost:9545")
	if err != nil {
		log.Fatal(err)
	}

	// privateKey, err := crypto.HexToECDSA("041deb3563e965bce6e803b88b9d25005cb1414c4cdade04181363e87ca9e259")
	privateKey, err := crypto.HexToECDSA("61cb1c6e53252dc3358c8b2bd07e329b8475dab24783f933434c43a327b1b5ae")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress("0x8BAccFF561FDe61D6bC8B6f299fFBa561d2189B9")
	scc, err := Scc.NewContracts(address, client)
	if err != nil {
		log.Fatal(err)
	}

	result, err := scc.FRAUDPROOFWINDOW(nil)
	if err != nil {
		log.Fatal("FRAUDPROOFWINDOW: ", err)
	}

	fmt.Println("FRAUDPROOFWINDOW: ", result)
	fmt.Println(big.NewInt(5).String())
	tx, err := scc.SetFraudProofWindow(auth, big.NewInt(5))
	if err != nil {
		log.Fatal("SetFraudProofWindow:", err)
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())

	tx, isPending, err := client.TransactionByHash(context.Background(), tx.Hash())
	if err != nil {
		log.Fatal("TransactionByHash: ", err)
	}

	times := 0
	for {
		fmt.Println(times)
		if isPending == true && times < 5 {
			time.Sleep(time.Duration(5) * time.Second)
			times++
			continue
		} else {
			break
		}
	}

	fmt.Println("isPending: ", isPending)

	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Fatal("TransactionReceipt: ", err)

	}
	fmt.Println("receipt status: ", receipt.Status)

	result, err = scc.FRAUDPROOFWINDOW(nil)
	if err != nil {
		log.Fatal("FRAUDPROOFWINDOW: ", err)
	}

	fmt.Println("FRAUDPROOFWINDOW: ", result)

}
