package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	client, err := rpc.DialContext(context.Background(), "https://goerli.infura.io/v3/3f31ba1b9fc54dfd92e1e26a64fba7e0")
	if err != nil {
		log.Fatal(err)
	}

	var accounts []string
	err = client.CallContext(context.Background(), &accounts, "eth_accounts")
	if err != nil {
		log.Fatal(err)
	}

	finalizedheader := types.Header{}
	client.CallContext(context.Background(), &finalizedheader, "eth_getBlockByNumber", "finalized", false)
	fmt.Println("finalized header: ", finalizedheader.Number)

	safeheader := types.Header{}
	client.CallContext(context.Background(), &safeheader, "eth_getBlockByNumber", "safe", false)
	fmt.Println("safe header:      ", safeheader.Number)

	latestheader := types.Header{}
	client.CallContext(context.Background(), &latestheader, "eth_getBlockByNumber", "latest", false)
	fmt.Println("latest header:    ", latestheader.Number)

	fmt.Println("AccountNum:", len(accounts))
	for _, account := range accounts {
		fmt.Println(account)
	}
}
