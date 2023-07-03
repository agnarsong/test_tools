package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func main() {
	// 合约函数的 ABI
	abiJSON := `[{"inputs":[{"internalType":"bytes32","name":"vmHash","type":"bytes32"},{"internalType":"uint256","name":"inboxSize","type":"uint256"},{"internalType":"bytes32[]","name":"_batch","type":"bytes32[]"},{"internalType":"uint256","name":"_shouldStartAtElement","type":"uint256"},{"internalType":"bytes","name":"_signature","type":"bytes"}],"name":"createAssertionWithStateBatch","outputs":[],"stateMutability":"nonpayable","type":"function"}]`

	abiJSON = `[{"constant":false,"inputs":[{"name":"memo","type":"bytes"}],"name":"receive","outputs":[],"payable":true,"stateMutability":"payable","type":"function"},{"constant":false,"inputs":[],"name":"send","outputs":[{"name":"amount","type":"uint256"}],"payable":true,"stateMutability":"payable","type":"function"},{"constant":false,"inputs":[{"name":"addr","type":"address"}],"name":"get","outputs":[{"name":"hash","type":"bytes"}],"payable":true,"stateMutability":"payable","type":"function"}]`

	// 已知的函数输入数据（例如，通过事务获得）
	inputData := "49cd30042efb73518663a6e4c6343bbc7ca67f6e01b8be60510052ca4cf73f4cb01cde02000000000000000000000000000000000000000000000000000000000000005700000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000004e00000000000000000000000000000000000000000000000000000000000001e00000000000000000000000000000000000000000000000000000000000000009c6e4fee5d8dd9102294918312955269ff962bf90ea144537bc7a5ef0cfa718330d7e2e63bc5cfbc6b4c6ba448aa4acf85aa205f39b76c97c7cef8be04aba2f22db003319f41566b3b3b72ac0974005893fa13153c9119649ff8931196d1028515a6004038f6d07bbcba1941614a3737118cc9667198b32b3c8ba17d518464ed30e6bf3d3ca3f0cb31b4299fea14260ae1c4fa35db13e6c6751909309ff78ea7af8edfca3e379ac3f00aa03c1fba6aa66b0f3ec9c674ad92b8be370493dd86bc7f63df9bcb6078f9b16365f0e627176e0d7810022e053905367c8bf4c650415dc25cc3f01dd1c802184abadff125c0baa0832b61d8a2758e1c068079a6988cb822efb73518663a6e4c6343bbc7ca67f6e01b8be60510052ca4cf73f4cb01cde020000000000000000000000000000000000000000000000000000000000000041a297e05b22a1fdbf699247a65ea16f7148ac781c5ea287c7ec3e134bfdb610165902a8117b858bfd9b31bd6f9cfe18358d4209ae07f8b7f11253dc0f78c10a4b0000000000000000000000000000000000000000000000000000000000000000"

	inputData = `00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000015800000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000158000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000001580000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000015800000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000158`

	args, err := parseContractFunctionInputs1(abiJSON, "createAssertionWithStateBatch", inputData)
	args, err = parseContractFunctionInputs1(abiJSON, "send", inputData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Function arguments: %v\n", args)
}

func parseContractFunctionInputs1(abiJSON string, funcName string, inputData string) (map[string]interface{}, error) {
	// 解析 ABI
	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return nil, fmt.Errorf("JSON err: ", err)
	}

	// 解析方法名（函数选择器）
	data, err := hex.DecodeString(inputData)
	if err != nil {
		return nil, fmt.Errorf("DecodeString err: ", err)
	}
	if len(data)%32 != 0 {
		return nil, fmt.Errorf("len(data) is %d, want a multiple of 32", len(data))
	}

	// 解析函数参数
	args := map[string]interface{}{}
	if err := parsedABI.UnpackIntoMap(args, funcName, data); err != nil {
		return nil, fmt.Errorf("UnpackIntoMap err: ", err)
	}
	return args, nil
}
