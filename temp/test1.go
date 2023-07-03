// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"math/big"

// 	"github.com/ethereum/go-ethereum"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/ethclient"
// )

// func main() {
// 	client, err := ethclient.Dial("http://127.0.0.1:9545")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 构造要查询的过滤条件
// 	query := ethereum.FilterQuery{
// 		FromBlock: big.NewInt(0),
// 		ToBlock:   big.NewInt(160),
// 		Addresses: []common.Address{common.HexToAddress("0x19C22f181280dF6Ad1d97285cdD430173Df91C12")},
// 	}

// 	// 查询符合条件的日志
// 	logs, err := client.FilterLogs(context.Background(), query)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 处理查询到的日志
// 	for _, vLog := range logs {
// 		fmt.Println("vLog: ", vLog, "\n") // 处理日志
// 	}
// }

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
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:9545")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x19C22f181280dF6Ad1d97285cdD430173Df91C12")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   big.NewInt(160),
		Addresses: []common.Address{contractAddress},
	}

	// 设置 ABI，需要根据实际情况替换下面的代码
	abiString := "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldAddress\",\"type\":\"address\"}],\"name\":\"AddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	contractAbi, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		log.Fatal(err)
	}
	contractAbi.Methods()

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	for _, l := range logs {
		fmt.Println("Log Block Number:", l.BlockNumber)
		fmt.Println("Log Index:", l.Index)
		fmt.Println("Log Data:", l.Data)

		// // 解析事件参数
		// eventAbi, err := contractAbi.EventByID(l.Topics[0])
		// if err != nil {
		// 	log.Fatal(err)
		// }

		if len(l.Data) > 0 {
			// 将日志数据解码成结构体
			event := map[string]interface{}{}
			err = contractAbi.UnpackIntoMap(event, "AddressSet", l.Data)
			if err != nil {
				log.Fatal(err)
			}

			// 打印事件参数
			fmt.Println("Event:", event)
		}

	}
}
