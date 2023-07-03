package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"mantle/test/config"
	"mantle/test/lib"
)

var (
	prv        string
	to         string
	value      string
	dataString string
)

func init() {
	clientCmd.AddCommand(ctCmd)
	ctCmd.AddCommand(ethCmd)
	ctCmd.AddCommand(erc20Cmd)

	ctCmd.PersistentFlags().StringP("toAddress", "t", "", "receiver address")

	ctCmd.PersistentFlags().StringVarP(&value, "value", "v", "1", "transfer value")
	ctCmd.PersistentFlags().StringVarP(&dataString, "dataString", "d", "", "transfer message")

	//将配置文件和入参绑定，使用入参替换默认配置
	config.AppConfig.BindPFlag("to", ctCmd.PersistentFlags().Lookup("to"))
}

var ctCmd = &cobra.Command{
	Use:     "t",
	Aliases: []string{"transfer"},
	Short:   "执行transfer相关操作",
	Long:    "执行transfer相关操作",
	Example: "mt client transfer SendETHTransaction",
	RunE: func(cmd *cobra.Command, args []string) error {

		fmt.Println("nothing todo……")
		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {

		prv = config.AppConfig.GetString("privateKey")
		to = config.AppConfig.GetString("to")

		return nil
	},
}

var ethCmd = &cobra.Command{
	Use:     "e",
	Aliases: []string{"SendETHTransaction"},
	Short:   "执行 ETH transfer",
	Long:    "执行 ETH transfer",
	RunE: func(cmd *cobra.Command, args []string) error {

		tx, cid, privateKey, err := lib.SignETHTx1(c, prv, to, value, []byte(dataString))
		if err != nil {
			return err
		}

		signedTx, err := lib.SignETHTx2(tx, cid, privateKey)
		if err != nil {
			return err
		}

		txHash, err := lib.SendTransaction(c, signedTx)
		if err != nil {
			return err
		}

		fmt.Println("txHash: ", txHash)
		return nil
	},
	PreRunE: ctCmd.PreRunE,
}

// dataString 为空, 并不能transferERC20
var erc20Cmd = &cobra.Command{
	Use:     "20",
	Aliases: []string{"SendERC20Transaction"},
	Short:   "执行 ERC20 transfer",
	Long:    "执行 ERC20 transfer",
	RunE: func(cmd *cobra.Command, args []string) error {

		tx, cid, privateKey, err := lib.SignETHTx1(c, prv, to, value, []byte(dataString))
		if err != nil {
			return err
		}

		signedTx, err := lib.SignETHTx2(tx, cid, privateKey)
		if err != nil {
			return err
		}

		txHash, err := lib.SendTransaction(c, signedTx)
		if err != nil {
			return err
		}

		fmt.Println("txHash: ", txHash)
		return nil
	},
	PreRunE: ctCmd.PreRunE,
}
