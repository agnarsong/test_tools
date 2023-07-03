package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	// 合约函数的 ABI
	abiJSON := `[{"inputs":[{"internalType":"bytes32","name":"vmHash","type":"bytes32"},{"internalType":"uint256","name":"inboxSize","type":"uint256"},{"internalType":"bytes32[]","name":"_batch","type":"bytes32[]"},{"internalType":"uint256","name":"_shouldStartAtElement","type":"uint256"},{"internalType":"bytes","name":"_signature","type":"bytes"}],"name":"createAssertionWithStateBatch","outputs":[],"stateMutability":"nonpayable","type":"function"}]`

	// 已知的函数输入数据（例如，通过事务获得）
	inputData := "0x49cd30042efb73518663a6e4c6343bbc7ca67f6e01b8be60510052ca4cf73f4cb01cde02000000000000000000000000000000000000000000000000000000000000005700000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000004e00000000000000000000000000000000000000000000000000000000000001e00000000000000000000000000000000000000000000000000000000000000009c6e4fee5d8dd9102294918312955269ff962bf90ea144537bc7a5ef0cfa718330d7e2e63bc5cfbc6b4c6ba448aa4acf85aa205f39b76c97c7cef8be04aba2f22db003319f41566b3b3b72ac0974005893fa13153c9119649ff8931196d1028515a6004038f6d07bbcba1941614a3737118cc9667198b32b3c8ba17d518464ed30e6bf3d3ca3f0cb31b4299fea14260ae1c4fa35db13e6c6751909309ff78ea7af8edfca3e379ac3f00aa03c1fba6aa66b0f3ec9c674ad92b8be370493dd86bc7f63df9bcb6078f9b16365f0e627176e0d7810022e053905367c8bf4c650415dc25cc3f01dd1c802184abadff125c0baa0832b61d8a2758e1c068079a6988cb822efb73518663a6e4c6343bbc7ca67f6e01b8be60510052ca4cf73f4cb01cde020000000000000000000000000000000000000000000000000000000000000041a297e05b22a1fdbf699247a65ea16f7148ac781c5ea287c7ec3e134bfdb610165902a8117b858bfd9b31bd6f9cfe18358d4209ae07f8b7f11253dc0f78c10a4b0000000000000000000000000000000000000000000000000000000000000000"

	// // 合约函数的 ABI
	abiJSON = `[{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"}]`

	// 已知的函数输入数据（例如，通过事务获得）
	inputData = "0x70a082310000000000000000000000009607Dd1f7F415A90D8E7C8575Ae888AC7C785D22"

	functionName, args, err := parseContractFunctionInputs2(abiJSON, inputData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Function name: %s\n", functionName)
	fmt.Printf("Function arguments: %v\n", args)
}

func parseContractFunctionInputs2(abiJSON string, inputData string) (string, map[string]interface{}, error) {
	// 解析 ABI
	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return "", nil, err
	}

	// 解析方法名（函数选择器）
	inputDataBytes := common.FromHex(inputData)
	if len(inputDataBytes) < 4 {
		return "", nil, fmt.Errorf("input data is too short")
	}
	method, err := parsedABI.MethodById(inputDataBytes[:4])
	if err != nil {
		return "", nil, err
	}

	// // 解析函数参数并将其存储在 map 中
	// args := make(map[string]struct {
	// 	VmHash               common.Hash   `json:vmHash`
	// 	Batch                []common.Hash `json:_batch`
	// 	InboxSize            *big.Int      `json:inboxSize`
	// 	ShouldStartAtElement *big.Int      `json:_shouldStartAtElement`
	// 	Signature            common.Hash   `json:_signature`
	// })
	args := map[string]interface{}{}

	err = parsedABI.UnpackIntoMap(args, method.Name, inputDataBytes[4:])
	if err != nil {
		return "", nil, fmt.Errorf("failed to unpack arguments: %w", err)
	}

	return method.Name, args, nil
}
