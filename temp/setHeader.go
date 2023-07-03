package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	// 创建一个新的 Geth 客户端
	client, err := rpc.DialContext(context.Background(), "http://127.0.0.1:9545")
	if err != nil {
		log.Fatal(err)
	}

	// 设置要更改的区块哈希
	blockHash := "0x96"

	// 调用 debug_setHead 方法
	var result interface{}
	err = client.CallContext(context.Background(), &result, "debug_setHead", blockHash)
	if err != nil {
		log.Fatal(err)
	}

	// 打印结果
	fmt.Printf("Result: %v\n", result)
}
