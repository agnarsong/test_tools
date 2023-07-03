package cmd

import (
	"encoding/hex"
	"fmt"
	"mantle/test/lib"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var (
	filePath string
	abiJson  string
	data     string
	funcName string
)

type PAbi struct {
	AbiInput  string `json:"abiInput"`
	InputData string `json:"inputData"`
	AbiOutput string `json:"abiOutput"`
	OutData   string `json:"outData"`
	AbiEvent  string `json:"abiEvent"`
	EventData string `json:"eventData"`
}

func init() {
	mtCmd.AddCommand(abiCmd)
	abiCmd.AddCommand(piCmd)
	abiCmd.AddCommand(poCmd)

	abiCmd.PersistentFlags().StringVarP(&filePath, "filePath", "f", "Pabi.json", "存储abi和data的jason文件")
	abiCmd.PersistentFlags().StringVarP(&abiJson, "abiJson", "a", "", "abi")
	abiCmd.PersistentFlags().StringVarP(&data, "data", "d", "", "InputData or OutputData")
	abiCmd.PersistentFlags().StringVarP(&funcName, "funcName", "n", "", "function name")
}

var abiCmd = &cobra.Command{
	Use:     "a",
	Aliases: []string{"abi"},
	Short:   "加强的abi工具",
	Long:    "加强的abi工具",
	RunE: func(cmd *cobra.Command, args []string) error {

		// tssnode0: 8e09231cd4e10460b4b4b0b291055ba5d2c900daed32642b47bca8c034c895ef
		// tssnode1: 00f75cacb3ff9b4ed18e093dbbf6f7e7188060fabfb352d2253ea0e26f96a3eb
		// tssnode2: de87f566d458c0e53eb866cadd4e8e39a83cb8e20801e57e1fb55a3e4107027b
		// tssnode3: 32a79d56c34105915ad71b6e0fc52b4cd6e7212dfa7b4dd0b24602282ee605ef
		// tssnode4: 6f900861d9dc69be839ab01f4dfc1c7c2076545880ff310c385ee49087ec47d1
		// tssnode5: 813a6d06a23daed107fa1a6f2cacb2c88488667889afa1d607199c4624273c51
		// tssnode6: e8f33f1fb0edb7a2fd4150dab6486bb0ecb29bf5aab27deceee16597c45efaf1

		// tssnode0: 0xd5750ebe91654ab6e345fd1c6f97348265e5ef9f
		// tssnode1: 0xd5751caac4cc34f9147fd2d856abef1c54e8b22b
		// tssnode2: 0xd5752dbebc3fbdee41f2f8dd7286d471517de7e9
		// tssnode3: 0xd5753cf7f8a55bb03e24841af6b7e98644c28836
		// tssnode4: 0xd575435e464252d8297bd79dd35eb57b14b62e81
		// tssnode5: 0xd57559b6edca44ebdbb2b618c3574ef3077afb5e
		// tssnode6: 0xd575667cf7b120fdb6662a083ddb75624f21ad5a

		// a, b, c, _ := lib.AnalysePrivateKey("8e09231cd4e10460b4b4b0b291055ba5d2c900daed32642b47bca8c034c895ef")
		// fmt.Println(common.Bytes2Hex(crypto.FromECDSA(a)), common.Bytes2Hex(crypto.FromECDSAPub(b)), c)
		// fmt.Println(common.Bytes2Hex(crypto.FromECDSA(a)), common.Bytes2Hex(crypto.CompressPubkey(b)), c)
		// 解析abi文件
		// abiJson的优先级比文件高
		pabi, err := lib.GetABI(filePath, abiJson)
		if err != nil {
			return err
		}

		lib.PrintAbi(pabi)

		return nil
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

		if abiJson == "" || data == "" {
			_, err := os.Stat(filePath)
			if os.IsNotExist(err) {
				return fmt.Errorf("file %s does not exist", filePath)
			}
		}

		return nil
	},
}

var piCmd = &cobra.Command{
	Use:     "pi",
	Aliases: []string{"ParseInputs"},
	Short:   "解析tx的inputData",
	Long:    "解析tx的inputData",
	Example: "mt client abi ParseInputs ",
	RunE: func(cmd *cobra.Command, args []string) error {

		if abiJson == "" || data == "" {

			var pabi PAbi
			err := lib.ReadJSONFile(filePath, &pabi)
			if err != nil {
				return fmt.Errorf("ReadJSONFile err: %v", err)
			}
			abiJson = pabi.AbiInput
			data = pabi.InputData
		}

		inputDataBytes := common.FromHex(data)
		// {\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vmHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"inboxSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_batch\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"_shouldStartAtElement\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"createAssertionWithStateBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}
		funcName, inputArgs, err := lib.ParseContractFunctionInputs(abiJson, inputDataBytes)
		if err != nil {
			return fmt.Errorf("ParseContractFunctionInputs err: %v", err)
		}
		fmt.Println(inputArgs)
		fmt.Println("funcName: ", funcName)
		fmt.Println("inputArgs: ")
		lib.PrintMap(inputArgs)

		return nil
	},
}

var poCmd = &cobra.Command{
	Use:     "po",
	Aliases: []string{"ParseOutputs"},
	Short:   "解析tx的outData",
	Long:    "解析tx的outData",
	Example: "mt client abi ParseOutputs ",
	RunE: func(cmd *cobra.Command, args []string) error {

		if abiJson == "" || data == "" {

			var pabi PAbi
			err := lib.ReadJSONFile(filePath, &pabi)
			if err != nil {
				return fmt.Errorf("ReadJSONFile err: %v", err)
			}
			abiJson = pabi.AbiEvent
			data = pabi.EventData
		}

		outputData, err := hex.DecodeString(data)
		if err != nil {
			return fmt.Errorf("DecodeString err: %v", err)
		}
		output, err := lib.ParseContractFunctionOutputs(abiJson, funcName, outputData)
		if err != nil {
			return fmt.Errorf("ParseContractFunctionInputs err: %v", err)
		}

		fmt.Println("output: ")
		lib.PrintMap(output)

		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {

		if funcName == "" {
			return fmt.Errorf("function name isn't nil")
		}

		return nil
	},
}
