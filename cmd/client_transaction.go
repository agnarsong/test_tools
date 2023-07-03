package cmd

import (
	"errors"
	"fmt"
	"mantle/test/lib"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

func init() {
	clientCmd.AddCommand(txCmd)
	txCmd.AddCommand(reCmd)
	txCmd.AddCommand(btCmd)
	txCmd.AddCommand(bfCmd)
	txCmd.AddCommand(atCmd)
	txCmd.AddCommand(fdCmd)
}

var txCmd = &cobra.Command{
	Use:     "tx",
	Aliases: []string{"transaction"},
	Args:    cobra.ExactArgs(1),
	Short:   "根据txHash查询交易",
	Long:    "根据txHash查询交易",
	Example: "mt client transaction <txHash>",
	RunE: func(cmd *cobra.Command, args []string) error {

		tx, _, err := lib.TransactionByHash(c, args[0])
		if err != nil {
			return err
		}
		fmt.Printf("根据txHash: %v 查询到txGas: %v 查询到inputData.len: %v\n", args[0], tx.Gas(), len(tx.Data()))
		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {

		if !strings.HasPrefix(strings.ToLower(args[0]), "0x") {
			return errors.New("txHash非法……")
		}
		return nil
	},
}

var reCmd = &cobra.Command{
	Use:     "re",
	Aliases: []string{"receipt"},
	Args:    cobra.ExactArgs(1),
	Short:   "根据txHash查询Receipt",
	Long:    "根据txHash查询Receipt",
	Example: "mt client transaction receipt <txHash>",
	RunE: func(cmd *cobra.Command, args []string) error {

		tr, err := lib.TransactionReceipt(c, args[0])
		if err != nil {
			return err
		}
		fmt.Printf("根据txHash: %v 查询到Receipt Status: %v\n", args[0], tr.Status)
		return nil
	},
	PreRunE: txCmd.PreRunE,
}

var btCmd = &cobra.Command{
	Use:     "bt",
	Aliases: []string{"byTo"},
	Args:    cobra.ExactArgs(1),
	Short:   "根据tx的to查询交易",
	Long:    "根据tx的to查询交易",
	Example: "mt client transaction byTo <toAddress>",
	RunE: func(cmd *cobra.Command, args []string) error {

		rs, err := lib.TransactionByTo(c, startBlock, endBlock, args[0])
		if err != nil {
			return err
		}

		for _, r := range rs {
			fmt.Println("============")
			fmt.Println("Block number: ", r.Block.Number())
			fmt.Println("txHash: ", r.Tx.Hash().Hex())
			fmt.Println("txFrom: ", r.From)
			fmt.Println("txTo: ", r.Tx.To())
			fmt.Println("txGas: ", r.Tx.Gas())
			fmt.Println("time: ", r.Block.Time())
			fmt.Println("txGasPrice: ", r.Tx.GasPrice().Uint64())
			fmt.Println("txInputData: ", common.Bytes2Hex(r.Tx.Data()))
			fmt.Println("reGasUsed: ", r.Re.GasUsed)
			fmt.Println("V: ", r.V)
			fmt.Println("R: ", r.R)
			fmt.Println("S: ", r.S)
		}
		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {

		_, isAddress = lib.CheckAddress(args[0])

		if !isAddress {
			return errors.New("输入的address非法,请确认……")
		}
		return nil
	},
}

var bfCmd = &cobra.Command{
	Use:     "bf",
	Aliases: []string{"byFrom"},
	Args:    cobra.ExactArgs(1),
	Short:   "根据tx的from查询交易",
	Long:    "根据tx的from查询交易",
	Example: "mt client transaction byFrom <fromAddress>",
	RunE: func(cmd *cobra.Command, args []string) error {

		rs, err := lib.TransactionByFrom(c, startBlock, endBlock, args[0])
		if err != nil {
			return err
		}

		for _, r := range rs {
			fmt.Println("============")
			fmt.Println("Block number: ", r.Block.Number())
			fmt.Println("txHash: ", r.Tx.Hash().Hex())
			fmt.Println("txFrom: ", r.From)
			fmt.Println("txTo: ", r.Tx.To())
			fmt.Println("txGas: ", r.Tx.Gas())
			fmt.Println("txGasPrice: ", r.Tx.GasPrice().Uint64())
			fmt.Println("txInputData: ", common.Bytes2Hex(r.Tx.Data()))
			fmt.Println("reGasUsed: ", r.Re.GasUsed)
			fmt.Println("V: ", r.V)
			fmt.Println("R: ", r.R)
			fmt.Println("S: ", r.S)
		}
		return nil
	},
	PreRunE: btCmd.PreRunE,
}

var atCmd = &cobra.Command{
	Use:     "at",
	Aliases: []string{"allTx"},
	Short:   "根据tx的to查询交易",
	Long:    "根据tx的to查询交易",
	Example: "mt client transaction allTx",
	RunE: func(cmd *cobra.Command, args []string) error {

		rs, err := lib.AllTransactions(c, startBlock, endBlock)
		if err != nil {
			return err
		}

		for _, r := range rs {
			fmt.Println("============")
			fmt.Println("Block number: ", r.Block.Number())
			fmt.Println("txHash: ", r.Tx.Hash().Hex())
			fmt.Println("txFrom: ", r.From)
			fmt.Println("txTo: ", r.Tx.To())
			fmt.Println("txGas: ", r.Tx.Gas())
			fmt.Println("txGasPrice: ", r.Tx.GasPrice().Uint64())
			fmt.Println("txInputData: ", common.Bytes2Hex(r.Tx.Data()))
			fmt.Println("reGasUsed: ", r.Re.GasUsed)
			fmt.Println("V: ", r.V)
			fmt.Println("R: ", r.R)
			fmt.Println("S: ", r.S)
		}
		return nil
	},
}

var fdCmd = &cobra.Command{
	Use:     "fd",
	Aliases: []string{"formatData"},
	Args:    cobra.ExactArgs(1),
	Short:   "遍历区块,根据条件解析交易的数据",
	Long:    "遍历区块,根据条件解析交易的数据",
	Example: "mt client transaction formatData <jsonfile>",
	RunE: func(cmd *cobra.Command, args []string) error {

		// 解析abi文件
		// abiJson的优先级比文件高
		if fromAddress != "" && toAddress != "" {
			// from and to
		} else if fromAddress == "" {
			// to
		} else if toAddress == "" {
			// from
		} else {
			// all
		}
		return nil
		pabi, err := lib.GetABI(filePath, abiJson)
		if err != nil {
			return err
		}

		lib.PrintAbi(pabi)

		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {

		if fromAddress != "" {
			_, isAddress = lib.CheckAddress(fromAddress)

			if !isAddress {
				return errors.New("输入的fromAddress非法,请确认……")
			}
		}
		if toAddress != "" {
			_, isAddress = lib.CheckAddress(toAddress)

			if !isAddress {
				return errors.New("输入的toAddress非法,请确认……")
			}
		}

		return nil
	},
}
