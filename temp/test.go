package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	client, err := rpc.DialHTTP("http://127.0.0.1:12002")
	if err != nil {
		log.Fatal(err)
	}

	var peers []interface{}
	err = client.CallContext(context.Background(), &peers, "admin_peers")
	if err != nil {
		log.Fatal(err)
	}

	for _, peer := range peers {
		peerInfo, ok := peer.(map[string]interface{})
		if !ok {
			continue
		}

		if _, ok := peerInfo["enode"]; !ok {
			continue
		}

		enode := peerInfo["enode"].(string)
		nodeID := common.HexToHash(enode[len("enode://") : len("enode://")+64])
		nodePublicKey := nodeID.Hex()
		fmt.Println("Node public key:", nodePublicKey)
	}
}
