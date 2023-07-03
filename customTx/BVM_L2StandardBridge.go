package main

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"mantle/test/lib"
	"mantle/test/lib/layer2/bindings"
)

func main() {
	// Connect to an Ethereum client (replace with your Ethereum client URL)
	client, err := ethclient.Dial("http://k8s-qa2-mantlel2-3c6b2e255f-b5a20c5d9c301f22.elb.ap-southeast-1.amazonaws.com")
	if err != nil {
		fmt.Errorf("Dial err: %v", err)
	}

	// Replace with the address of the contract you want to interact with
	contractAddress := common.HexToAddress("0x4200000000000000000000000000000000000010")
	l2StandardBridge, err := bindings.NewL2StandardBridge(contractAddress, client)
	if err != nil {
		fmt.Errorf("NewL2StandardBridge err: %v", err)
	}

	auth, err := lib.GetAuth(client, "e7bbfa2e0a24248701d7de1745495d8c7e08a835782006db5dddc37442f2ccb2")
	if err != nil {
		fmt.Errorf("lib.GetAuth err: %v", err)
	}
	fmt.Println(auth)

	// Prepare the function parameters
	l2TokenAddress := common.HexToAddress("0xDeadDeAddeAddEAddeadDEaDDEAdDeaDDeAD0000") // Replace with the address of the _l2Token
	amount := big.NewInt(1)                                                             // Replace with the value of _amount (in wei)
	l1Gas := uint32(2000000)                                                            // Replace with the value of _l1Gas
	data := []byte(strings.Repeat("大麦", 10400))                                         // Replace with the value of _data

	tx, err := l2StandardBridge.Withdraw(auth, l2TokenAddress, amount, l1Gas, data)
	if err != nil {
		fmt.Errorf("l2StandardBridge.Withdraw err: %v", err)
	}

	// Send the transaction
	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		fmt.Errorf("SendTransaction err: %v", err)
	}

	fmt.Printf("Transaction hash: 0x%x\n", tx.Hash())
}
