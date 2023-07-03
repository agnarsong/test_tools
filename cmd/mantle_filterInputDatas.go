package cmd

import (
	"fmt"
	"mantle/test/lib"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

func init() {

	mantleCmd.AddCommand(filterInputDatasCmd)
}

var filterInputDatasCmd = &cobra.Command{
	Use:     "fid",
	Aliases: []string{"filterInputDatas"},
	Short:   "filterInputDatas",
	Long:    "filterInputDatas",
	RunE: func(cmd *cobra.Command, args []string) error {

		client := mc.L1Client
		startBlock := uint64(0)
		endBlock := uint64(0)
		toAddress := "0x8013ac56e4c4A562b72a6D8B39D60cDa7aE173A9"
		rs, err := lib.TransactionByTo(client, startBlock, endBlock, toAddress)
		if err != nil {
			return err
		}

		// strings.TrimPrefix(str, "0x")
		// // 判断是否以"Hello"开头
		// if strings.HasPrefix(str, "Hello") {
		// 	fmt.Println("字符串以Hello开头")
		// } else {
		// 	fmt.Println("字符串不以Hello开头")
		// }
		for _, r := range rs {
			fmt.Println("inputData: ", common.Bytes2Hex(r.Tx.Data()))
		}
		return nil
	},
	PersistentPreRunE: mantleCmd.PersistentPreRunE,
}
