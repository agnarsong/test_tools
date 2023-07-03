package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mantle/test/lib"
	"mantle/test/lib/layer2"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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

	mtCmd.AddCommand(mantleCmd)
	mantleCmd.AddCommand(ERC20Cmd)
	mantleCmd.AddCommand(d20Cmd)
	mantleCmd.AddCommand(stCmd)
	mantleCmd.AddCommand(rCmd)
	mantleCmd.AddCommand(toCTCCmd)
	mantleCmd.AddCommand(toCTCWithTxHashCmd)
	mantleCmd.AddCommand(statusCmd)

	mantleCmd.PersistentFlags().StringVarP(&layer, "layer", "l", "l1", "layer: l1 or l2")
	mantleCmd.PersistentFlags().StringVar(&l1token, "l1token", "0", "l1 token address")
	mantleCmd.PersistentFlags().StringVar(&l2token, "l2token", "0", "l2 token address")
	mantleCmd.PersistentFlags().StringVarP(&amount, "amount", "a", "1", "transfer amount")
	mantleCmd.PersistentFlags().StringVarP(&privateKey, "privateKey", "p", "", "privateKey")
	mantleCmd.PersistentFlags().IntVar(&accountNum, "accountNum", 10, "init accounts num")

	// 初始化
	if err := viper.Unmarshal(&mc.Env); err != nil {
		fmt.Println("viper.Unmarshal err: ", err)
		return
	}

	if err = layer2.InitSc(&mc); err != nil {
		fmt.Println("viper.Unmarshal err: ", err)
		return
	}

	// 更新配置文件
	if viper.Get("addresslist.Proxy__L1MantleToken") != mc.Env.AddressList.Proxy__L1MantleToken {
		viper.Set("addresslist", mc.Env.AddressList)
		viper.Set("userprivatekeylist", mc.Env.PrivateKeyList)

		err := viper.WriteConfig()
		if err != nil {
			fmt.Println("write config failed: ", err)
			return
		}
	}
}

var (
	mc         layer2.MantleCenter
	l1token    string
	l2token    string
	layer      string
	amount     string
	privateKey string
	accountNum int
	isMNT      bool
	isETH      bool
	maxRetries = 10
)

type IntoConf struct {
	Url        string
	Layer      string
	Contract   string
	PriKey     string
	ToAddress  string
	Amount     string
	IsMNT      bool
	IsETH      bool
	StartBlock uint64
	EndBlock   uint64
	Block      uint64
}

var mantleCmd = &cobra.Command{
	Use:     "m",
	Aliases: []string{"mantle"},
	Short:   "mantle相关的测试工具, 主要服务于layer2",
	Long:    "mantle相关的测试工具, 主要服务于layer2",
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
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
}

// 检查状态的后续还能继续补充，争取一个命令就能检查整条链的状态
var statusCmd = &cobra.Command{
	Use:       "status",
	Aliases:   []string{"status"},
	ValidArgs: []string{"scc", "ctc"},
	Args:      cobra.OnlyValidArgs,
	Short:     "查看 mantle 特定合约的状态",
	Long:      "查看 mantle 特定合约的状态",
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) < 1 || len(args) > 1 {

			return errors.New("只能输入一个参数! ")
		}

		switch args[0] {
		case "scc":
			totalBatches, err := mc.SCC.GetTotalBatches(&bind.CallOpts{})
			if err != nil {
				return err
			}
			totalElements, err := mc.SCC.GetTotalElements(&bind.CallOpts{})
			if err != nil {
				return err
			}
			fraudProofWindow, err := mc.SCC.FRAUDPROOFWINDOW(&bind.CallOpts{})
			if err != nil {
				return err
			}
			sequencerPublishWindow, err := mc.SCC.SEQUENCERPUBLISHWINDOW(&bind.CallOpts{})
			if err != nil {
				return err
			}
			fmt.Println("SCC Status: ")
			fmt.Println("getTotalBatches: ", totalBatches)
			fmt.Println("getTotalElements: ", totalElements)
			fmt.Println("FRAUD_PROOF_WINDOW: ", fraudProofWindow)
			fmt.Println("SEQUENCER_PUBLISH_WINDOW: ", sequencerPublishWindow)
		case "ctc":
			totalBatches, err := mc.CTC.GetTotalBatches(&bind.CallOpts{})
			if err != nil {
				return err
			}
			totalElements, err := mc.CTC.GetTotalElements(&bind.CallOpts{})
			if err != nil {
				return err
			}
			queueLength, err := mc.CTC.GetQueueLength(&bind.CallOpts{})
			if err != nil {
				return err
			}
			nextQueueIndex, err := mc.CTC.GetNextQueueIndex(&bind.CallOpts{})
			if err != nil {
				return err
			}
			numPendingQueueElements, err := mc.CTC.GetNumPendingQueueElements(&bind.CallOpts{})
			if err != nil {
				return err
			}
			maxTransactionGasLimit, err := mc.CTC.MaxTransactionGasLimit(&bind.CallOpts{})
			if err != nil {
				return err
			}
			maxRollupTxSize, err := mc.CTC.MAXROLLUPTXSIZE(&bind.CallOpts{})
			if err != nil {
				return err
			}
			fmt.Println("CTC Status: ")
			fmt.Println("getTotalBatches: ", totalBatches)
			fmt.Println("getTotalElements: ", totalElements)
			fmt.Println("getQueueLength: ", queueLength)
			fmt.Println("getNumPendingQueueElements: ", numPendingQueueElements)
			fmt.Println("getNextQueueIndex: ", nextQueueIndex)
			fmt.Println("MAX_ROLLUP_TX_SIZE: ", maxRollupTxSize)
			fmt.Println("MaxTransactionGasLimit: ", maxTransactionGasLimit)
		}

		return nil
	},
	PreRunE: mantleCmd.PersistentPreRunE,
}

var toCTCCmd = &cobra.Command{
	Use:     "tc",
	Aliases: []string{"toCTC"},
	Short:   "解析发送到CTC的请求",
	Long:    "解析发送到CTC的请求",
	RunE: func(cmd *cobra.Command, args []string) error {

		rs, err := lib.TransactionByTo(mc.L1Client, 0, 0, mc.CTCAddress.Hex())
		if err != nil {
			return err
		}

		for _, r := range rs {
			inputData := common.Bytes2Hex(r.Tx.Data())
			ps, err := layer2.DecodeAppendSequencerBatchParams(inputData)
			if err != nil {
				return err
			}
			fmt.Println("ShouldStartAtElement: ", ps.ShouldStartAtElement)
			fmt.Println("TotalElementsToAppend: ", ps.TotalElementsToAppend)
		}
		return nil
	},
	PreRunE: mantleCmd.PersistentPreRunE,
}

var toCTCWithTxHashCmd = &cobra.Command{
	Use:     "tcwth",
	Aliases: []string{"toCTCWithTxHash"},
	Args:    cobra.ExactArgs(1),
	Short:   "根据发送到CTC的txHash解析inputdata",
	Long:    "根据发送到CTC的txHash解析inputdata",
	Example: "mt tcwth <txHash>",
	RunE: func(cmd *cobra.Command, args []string) error {

		tx, _, err := lib.TransactionByHash(mc.L1Client, args[0])
		if err != nil {
			return err
		}

		inputData := common.Bytes2Hex(tx.Data())
		ps, err := layer2.DecodeAppendSequencerBatchParams(inputData)
		if err != nil {
			return err
		}

		fmt.Println("ShouldStartAtElement: ", ps.ShouldStartAtElement)
		fmt.Println("TotalElementsToAppend: ", ps.TotalElementsToAppend)

		return nil
	},
	PreRunE: mantleCmd.PersistentPreRunE,
}
