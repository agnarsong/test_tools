package cmd

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"

	"mantle/test/lib"
)

func init() {
	clientCmd.AddCommand(cbCmd)
	clientCmd.AddCommand(bnCmd)
	cbCmd.AddCommand(bhCmd)
	cbCmd.AddCommand(tcCmd)

	bhCmd.AddCommand(bsCmd)

	// bs.
}

var cbCmd = &cobra.Command{
	Use:     "b",
	Aliases: []string{"block"},
	Args:    cobra.ExactArgs(1),
	Short:   "执行block相关操作",
	Long:    "执行block相关操作",
	Example: "mt client block <blockNum | blockHash>",
	RunE: func(cmd *cobra.Command, args []string) error {

		if !strings.HasPrefix(strings.ToLower(args[0]), "0x") {

			blockNum, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return err
			}

			block, err := lib.BlockByNumber(c, blockNum)
			if err != nil {
				return err
			}

			// 调用 CallMethodByName 函数
			// d, err := lib.CallMethodByName("Difficulty", block)
			// if err != nil {
			// 	return err
			// }
			// fmt.Println("difficulty: ", d)

			// 调用 CallAllMethods 函数
			results, err := lib.CallAllMethods(block)
			fmt.Println("======", results)
			if err != nil {
				fmt.Println(err)
			} else {
				// 输出返回值
				for methodName, methodResult := range results {
					fmt.Printf("%s: %v\n", methodName, methodResult)
				}
			}

			fmt.Println("baseFeePerGas: ", block.BaseFee())
			fmt.Println("difficulty: ", block.Difficulty())

		} else {
			block, err := lib.BlockByHash(c, args[0])
			if err != nil {
				return err
			}

			fmt.Printf("根据blockhash: %v 查询到了区块, block.timestamp: %v\n", args[0], block.Time())
		}

		return nil
	},
}

var bnCmd = &cobra.Command{
	Use:     "bn",
	Aliases: []string{"blockNum"},
	Short:   "查询最新区块高度",
	Long:    "查询最新区块高度",
	RunE: func(cmd *cobra.Command, args []string) error {

		bn, err := lib.BlockNumber(c)
		if err != nil {
			return err
		}

		fmt.Println("最新区块高度为: ", bn)
		return nil
	},
}

var bhCmd = &cobra.Command{
	Use:     "h",
	Aliases: []string{"header"},
	Args:    cobra.ExactArgs(1),
	Short:   "查询指定区块的header",
	Long:    "查询指定区块的header",
	Example: "mt client block header <blockNum | blockHash>",
	RunE: func(cmd *cobra.Command, args []string) error {

		var header *types.Header
		var err error

		if !strings.HasPrefix(strings.ToLower(args[0]), "0x") {

			blockNum, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return err
			}
			header, err = lib.HeaderByNumber(c, blockNum)
			if err != nil {
				return err
			}

		} else {
			header, err = lib.HeaderByHash(c, args[0])
			if err != nil {
				return err
			}
		}

		fmt.Println("header.BaseFee: ", header.BaseFee)
		fmt.Println("header.Bloom: ", header.Bloom)
		fmt.Println("header.Coinbase: ", header.Coinbase)
		fmt.Println("header.Difficulty: ", header.Difficulty)
		fmt.Println("header.Extra: ", header.Extra)
		fmt.Println("header.GasLimit: ", header.GasLimit)
		fmt.Println("header.GasUsed: ", header.GasUsed)
		fmt.Println("header.MixDigest: ", header.MixDigest)
		fmt.Println("header.Nonce: ", header.Nonce)
		fmt.Println("header.Number: ", header.Number)
		fmt.Println("header.ParentHash: ", header.ParentHash)
		fmt.Println("header.ReceiptHash: ", header.ReceiptHash)
		fmt.Println("header.Root: ", header.Root)
		fmt.Println("header.Time: ", header.Time)
		fmt.Println("header.TxHash: ", header.TxHash)
		fmt.Println("header.UncleHash: ", header.UncleHash)

		return nil
	},
}

var (
	targetStateRoot common.Hash
	err             error
)
var bsCmd = &cobra.Command{
	Use:     "bs",
	Aliases: []string{"bystateroot"},
	Args:    cobra.ExactArgs(1),
	Short:   "遍历区块,根据bystateroot查询指定区块的header",
	Long:    "遍历区块,根据bystateroot查询指定区块的header",
	Example: "mt client block header bystateroot <stateroot>",
	RunE: func(cmd *cobra.Command, args []string) error {

		b, err := lib.FindBlockNumberByStateRoot(c, startBlock, endBlock, targetStateRoot)
		if err != nil {
			return fmt.Errorf("FindBlockNumberByStateRoot err: %v", err)
		}

		fmt.Println("block.baseFeePerGas: ", b.BaseFee())
		fmt.Println("block.Difficulty: ", b.Difficulty())
		fmt.Println("block.Extra: ", common.Bytes2Hex(b.Extra()))
		fmt.Println("block.GasLimit: ", b.GasLimit())
		fmt.Println("block.GasUsed: ", b.GasUsed())
		fmt.Println("block.Hash: ", b.Hash())
		fmt.Println("block.Bloom: ", common.Bytes2Hex(b.Bloom().Bytes()))
		fmt.Println("block.Miner: ", b.Coinbase())
		fmt.Println("block.mixHash: ", b.MixDigest())
		fmt.Println("block.Nonce: ", b.Nonce())
		fmt.Println("block.Number: ", b.Number())
		fmt.Println("block.ParentHash: ", b.ParentHash())
		fmt.Println("block.ReceiptHash: ", b.ReceiptHash())
		fmt.Println("block.UncleHash: ", b.UncleHash())
		fmt.Println("block.Size: ", b.Size())
		fmt.Println("block.Root: ", b.Root())
		fmt.Println("block.Time: ", b.Time())
		fmt.Println("block.TxHashs: ", b.Transactions())

		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {

		if endBlock == 0 {
			endBlock, err = c.BlockNumber(context.Background())
			if err != nil {
				return fmt.Errorf("BlockNumber err: %v", err)
			}
		}
		targetStateRoot = common.HexToHash(args[0])
		return nil
	},
}

var tcCmd = &cobra.Command{
	Use:     "tc",
	Aliases: []string{"transactionCount"},
	Args:    cobra.ExactArgs(1),
	Short:   "统计区块包含的交易数量",
	Long:    "统计区块包含的交易数量",
	Example: "mt client block transactionCount <blockHash>",
	RunE: func(cmd *cobra.Command, args []string) error {

		count, err := lib.TransactionCount(c, args[0])
		if err != nil {
			return err
		}

		fmt.Printf("根据blockHash: %v 查询到该区块下包含 %v 笔txs。", args[0], count)
		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if !strings.HasPrefix(strings.ToLower(args[0]), "0x") {

			return errors.New("blockHash非法……")
		}
		return nil
	},
}
