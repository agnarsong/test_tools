package cmd

import (
	"context"
	"fmt"
	"log"
	"mantle/test/lib/layer2"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {

	viper.SetConfigName("stressConf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}

	mantleCmd.AddCommand(toolsCmd)
	toolsCmd.AddCommand(getStateBatchAppendedCmd)

}

var toolsCmd = &cobra.Command{
	Use:     "t",
	Aliases: []string{"tools"},
	Short:   "测试工具",
	Long:    "测试工具",
	RunE: func(cmd *cobra.Command, args []string) error {

		// 打印l1 l2 的高度
		cid1, err := mc.L1Client.ChainID(context.Background())
		fmt.Println("l1 ChainID: ", cid1)
		if err != nil {
			return err
		}

		cid2, err := mc.L2Client.ChainID(context.Background())
		fmt.Println("l2 ChainID: ", cid2)
		if err != nil {
			return err
		}

		return nil
	},
	PreRunE: stressCmd.PersistentPreRunE,
}

// StateBatchAppended
var getStateBatchAppendedCmd = &cobra.Command{
	Use:     "gsba",
	Aliases: []string{"getStateBatchAppended"},
	Short:   "遍历StateBatchAppended",
	Long:    "遍历StateBatchAppended",
	RunE: func(cmd *cobra.Command, args []string) error {

		if int64(endBlock) == int64(0) {
			endBlock, _ = mc.L1Client.BlockNumber(context.Background())
		}

		for i := int64(startBlock); i < int64(endBlock); i++ {
			layer2.GetAppendEvent(
				mc.L1Client,
				mc.Env.AddressList.Proxy__Rollup,
				mc.Env.AddressList.StateCommitmentChain,
				i)
		}

		return nil
	},
	PreRunE: stressCmd.PersistentPreRunE,
}
