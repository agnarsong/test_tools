package lib

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"os"
	"reflect"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func CheckAddress(address string) (a common.Address, b bool) {

	if !common.IsHexAddress(address) {
		return a, b
	}

	return common.HexToAddress(address), true
}

func AnalysePrivateKey(prv string) (privateKey *ecdsa.PrivateKey, publicKeyECDSA *ecdsa.PublicKey,
	fromAddress common.Address, err error) {

	privateKey, err = crypto.HexToECDSA(prv)
	if err != nil {
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err = errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return
	}

	fromAddress = crypto.PubkeyToAddress(*publicKeyECDSA)
	return
}

func PrintKeyValue(v interface{}) {
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)

		// Skip over unexported fields or methods
		if field.PkgPath != "" {
			continue
		}
		fmt.Printf("%v: %v\n", field.Name, value.Interface())
	}
	fmt.Println()
}

type ExportedBlock struct {
	Header       *types.Header
	Uncles       []*types.Header
	Transactions types.Transactions

	// caches
	Hash common.Hash
	Size uint64

	// These fields are used by package eth to track
	// inter-peer block relay.
	ReceivedAt   time.Time
	ReceivedFrom interface{}
}

func NewExportedBlock(b *types.Block) ExportedBlock {
	return ExportedBlock{
		Header:       b.Header(),
		Uncles:       b.Uncles(),
		Transactions: b.Transactions(),
		Hash:         b.Hash(),
		Size:         b.Size(),
		ReceivedAt:   b.ReceivedAt,
		ReceivedFrom: b.ReceivedFrom,
	}
}

func CallMethodByName(name string, receiver interface{}, args ...interface{}) ([]interface{}, error) {
	method := reflect.ValueOf(receiver).MethodByName(name)

	if !method.IsValid() {
		return nil, fmt.Errorf("method %s not found", name)
	}

	// 构造参数列表
	var input []reflect.Value
	for _, arg := range args {
		input = append(input, reflect.ValueOf(arg))
	}

	// 调用方法
	result := method.Call(input)

	// 处理返回值
	var output []interface{}
	for _, value := range result {
		output = append(output, value.Interface())
	}

	return output, nil
}

// 对象的所有方法都不需要入参的时候，可以使用
func CallAllMethods(receiver interface{}) (map[string][]interface{}, error) {
	// 创建返回值映射
	results := make(map[string][]interface{})

	// 遍历所有方法并调用它们
	value := reflect.ValueOf(receiver)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	structType := value.Type()
	fmt.Println("structType", structType.NumMethod())
	for i := 0; i < structType.NumMethod(); i++ {
		method := structType.Method(i)

		// 调用方法
		result, err := CallMethodByName(method.Name, receiver)
		if err != nil {
			return nil, err
		}

		// 保存返回值
		results[method.Name] = result
	}

	return results, nil
}

func PrintMap(m map[string]interface{}) {
	value := reflect.ValueOf(m)

	fmt.Printf("Type of myMap: %s\n", value.Type())

	for _, key := range value.MapKeys() {
		fmt.Printf("%s: %v\n", key, value.MapIndex(key))
	}
}

func ReadJSONFile(path string, v interface{}) error {
	// Open the file
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read the file contents
	contents, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	// Unmarshal the JSON data into a Go struct
	err = json.Unmarshal(contents, v)
	if err != nil {
		return err
	}

	return nil
}

func ParseAmount(amount string) (*big.Int, bool) {
	a := new(big.Int)
	_, b := a.SetString(amount, 10)
	if amount == "-1" {
		// 1 << 78
		// 302231454903657293676544 / 1e18 = 302231.4549036573
		a = new(big.Int).Lsh(big.NewInt(1), 80)
	}

	return a, b
}

func BalanceOf(c *ethclient.Client, contractAddress common.Address,
	holderAddress common.Address, bl *big.Int) (*big.Int, error) {

	// 计算 balanceOf 方法的函数选择器
	methodID := crypto.Keccak256([]byte("balanceOf(address)"))[:4]

	// 拼接参数，将地址左侧补零
	var data []byte
	data = append(data, methodID...)
	data = append(data, common.LeftPadBytes(holderAddress.Bytes(), 32)...)

	// fmt.Println("data: ", common.Bytes2Hex(data))
	// 调用合约方法获取余额
	res, err := c.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}, nil)

	if err != nil {
		return common.Big0, fmt.Errorf("CallContract err: %v", err)
	}

	return new(big.Int).SetBytes(res), err
}
