package lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func GetABI(filePath string, abiJson string) (abiDef abi.ABI, err error) {

	data := []byte(abiJson)

	if abiJson == "" {
		data, err = ioutil.ReadFile(filePath)
		if err != nil {
			return abiDef, fmt.Errorf("ioutil.ReadFile err: %v", err)
		}
	}

	if strings.HasPrefix(string(data), "{") {
		// 解析 JSON 文件
		var contractJSON map[string]interface{}
		if err := json.Unmarshal(data, &contractJSON); err != nil {

			return abiDef, fmt.Errorf("json.Unmarshal err: %v", err)
		}

		// 将 interface{} 转换为 []byte
		data, err = json.Marshal(contractJSON["abi"])
		if err != nil {
			return abiDef, fmt.Errorf("json.Marshal err: %v", err)
		}

	} else if strings.HasPrefix(string(data), "[") {

	} else {
		err = errors.New("err abi json file")
		return
	}

	abiDef, err = abi.JSON(bytes.NewReader(data))
	if err != nil {
		return abiDef, fmt.Errorf("abi.JSON err: %v", err)
	}
	return
}

func ParseContractFunctionInputs(abiJSON string, inputDataBytes []byte) (string, map[string]interface{}, error) {
	// 解析 ABI
	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return "", nil, err
	}

	// 解析方法名（函数选择器）
	if len(inputDataBytes) < 4 {
		return "", nil, fmt.Errorf("input data is too short")
	}
	method, err := parsedABI.MethodById(inputDataBytes[:4])
	if err != nil {
		return "", nil, err
	}

	// 解析函数参数
	args := map[string]interface{}{}
	if len(inputDataBytes) < 4+len(method.Inputs)*32 {
		return "", nil, fmt.Errorf("input data is not long enough for the expected arguments")
	}

	err = method.Inputs.UnpackIntoMap(args, inputDataBytes[4:])
	if err != nil {
		return "", nil, fmt.Errorf("failed to unpack arguments: %v", err)
	}

	return method.Name, args, nil
}

func ParseContractFunctionOutputs(abiJSON string, funcName string, data []byte) (map[string]interface{}, error) {
	// 解析 ABI
	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return nil, fmt.Errorf("JSON err: %v", err)
	}

	return ParseContractFunctionOutputs1(parsedABI, funcName, data)
}

func ParseContractFunctionOutputs1(parsedABI abi.ABI, funcName string, data []byte) (map[string]interface{}, error) {

	// 解析方法名（函数选择器）
	if len(data)%32 != 0 {
		return nil, fmt.Errorf("len(data) is %d, want a multiple of 32", len(data))
	}

	// 解析函数参数
	args := map[string]interface{}{}
	if err := parsedABI.UnpackIntoMap(args, funcName, data); err != nil {
		return nil, fmt.Errorf("UnpackIntoMap err: %v", err)
	}
	return args, nil
}

// todo…… ParseContractEventOutputs

func PrintAbi(abiDef abi.ABI) {
	for _, m := range abiDef.Methods {
		fmt.Println("Function name:", m.Name)

		for _, input := range m.Inputs {
			fmt.Printf("Input name: %s, type: %s\n", input.Name, input.Type.String())
		}

		for _, output := range m.Outputs {
			fmt.Printf("Output name: %s, type: %s\n", output.Name, output.Type.String())
		}

		fmt.Println("State Mutability:", m.StateMutability)
		fmt.Println()
	}
	fmt.Println("======================================")

	for _, e := range abiDef.Events {
		fmt.Println("Event name:", e.Name)

		for _, input := range e.Inputs {
			fmt.Printf("Input name: %s, type: %s\n", input.Name, input.Type.String())
		}

		fmt.Println("RawName:", e.RawName)

		fmt.Println("Sig:", e.Sig)
		fmt.Println()
	}

	fmt.Println("======================================")
	for _, er := range abiDef.Errors {
		fmt.Println("Error name:", er.Name)

		for _, input := range er.Inputs {
			fmt.Printf("Input name: %s, type: %s\n", input.Name, input.Type.String())
		}

		fmt.Println("Sig:", er.Sig)
	}
}
