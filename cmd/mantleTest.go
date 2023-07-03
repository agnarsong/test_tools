package cmd

import (
	"fmt"
	"mantle/test/config"

	"github.com/spf13/cobra"
)

var mtCmd = &cobra.Command{
	Use:   "mt",
	Short: "mt short",
	Long:  "mantle test tools",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("author: ", config.AppConfig.GetString("author"))
		fmt.Println("license: ", config.AppConfig.GetString("license"))
		fmt.Println("nothing to do ……")
	},
}

func Execute() {
	mtCmd.Execute()
}

var (
	conf       string
	contract   string
	startBlock uint64
	endBlock   uint64
	block      uint64
)

func init() {
	cobra.OnInitialize(initConfig)
	mtCmd.PersistentFlags().StringVar(&conf, "config", "conf.yaml", "")

	mtCmd.PersistentFlags().StringVarP(&url, "url", "u", "http://localhost:8545", "将执行交易的rpc地址")
	mtCmd.PersistentFlags().StringVarP(&contract, "contract", "c", "", "将执行交易的contract地址")
	mtCmd.PersistentFlags().StringVarP(&toAddress, "toAddress", "t", "", "目标地址: eoa address | contract address")

	mtCmd.PersistentFlags().Uint64VarP(&startBlock, "startBlock", "s", 0, "start Block number")
	mtCmd.PersistentFlags().Uint64VarP(&endBlock, "endBlock", "e", 0, "end Block number")
	mtCmd.PersistentFlags().Uint64VarP(&block, "block", "b", 0, "将执行交的 BlockNumber")
}

func initConfig() {
	config.InitConfig(conf)
}
