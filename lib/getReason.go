package lib

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
)

var web3 = map[string]string{
	"r":   "https://rinkeby.infura.io/v3/023f2af0f670457d9c4ea9cb524f0810",
	"a":   "https://rinkeby.arbitrum.io/rpc",
	"b":   "https://data-seed-prebsc-2-s2.binance.org:8545",
	"t":   "https://evm-rpc.testnet.teleport.network",
	"q":   "https://teleport-localvalidator.qa.davionlabs.com/",
	"l1":  "http://127.0.0.1:9545",
	"l2":  "http://127.0.0.1:8545",
	"dev": "https://mantle-l2geth.dev.davionlabs.com",
}

func reason(url string, txHash common.Hash) {
	client, err := ethclient.Dial(web3[url])
	if err != nil {
		fmt.Println("Error dialing:", err)
		return
	}

	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		fmt.Println("Error getting transaction:", err)
		return
	}

	if isPending {
		fmt.Println("Transaction is pending")
		return
	}

	rec, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		fmt.Println("Error getting TransactionReceipt:", err)
		return
	}
	if rec.Status != 1 {
		code := hexutil.Encode(tx.Data())
		codeStr := strings.TrimPrefix(string(code), "0x")
		fmt.Println("errCodeStr:", codeStr)
	}
}

func main() {
	url := "l1"
	hash := "0xdcba6074fda0870d60c791fefd5f252cf3b613940ee64eaa5890e4ede050666b"

	reason(url, common.HexToHash(hash))
}
