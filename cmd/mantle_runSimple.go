package cmd

import (
	"context"
	"fmt"
	"log"
	"sync"

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

	mantleCmd.AddCommand(rsCmd)
}

var rsCmd = &cobra.Command{
	Use:     "rs",
	Aliases: []string{"runSimple"},
	Short:   "日常回归测试工具",
	Long:    "日常回归测试工具",
	RunE: func(cmd *cobra.Command, args []string) error {

		waitGroup := new(sync.WaitGroup)

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
		// a, b := lib.ParseAmount(amount)
		// if !b {
		// 	return fmt.Errorf("ParseAmount return false")
		// }

		// // deposit ETH
		// if strings.Contains(args[0], "a") {
		// 	waitGroup.Add(1)
		// 	go func(wg *sync.WaitGroup) {
		// 		defer wg.Done()
		// 		for {
		// 			tx, err := layer2.Deposit(&mc, mc.Env.PrivateKeyList[1][0], common.HexToAddress("0"),
		// 				common.HexToAddress(mc.Env.L2EthAddress), a)
		// 			if err != nil {
		// 				fmt.Println("DepositETH err: ", err)
		// 			}
		// 			fmt.Println("DepositETH txHash: ", tx.Hash())
		// 			time.Sleep(time.Duration(2) * time.Second)

		// 			_, l2b, err := layer2.BETH(&mc, common.HexToAddress(mc.Env.PrivateKeyList[1][1]), nil)
		// 			if err != nil {
		// 				fmt.Println("layer2.BETH err: ", err)
		// 			}
		// 			fmt.Println("l2 ETH balance: ", l2b)
		// 		}
		// 	}(waitGroup)
		// }

		// // deposit ERC20
		// if strings.Contains(args[0], "b") {
		// 	waitGroup.Add(1)
		// 	go func(wg *sync.WaitGroup) {
		// 		defer wg.Done()
		// 		for {
		// 			tx, err := layer2.Deposit(&mc, mc.Env.PrivateKeyList[2][0], mc.L1ERC20Address,
		// 				mc.L2ERC20Address, big.NewInt(1))
		// 			if err != nil {
		// 				fmt.Println("DepositERC20 err: ", err)
		// 			}
		// 			fmt.Println("DepositERC20 txHash: ", tx.Hash())
		// 			time.Sleep(time.Duration(2) * time.Second)

		// 			_, l2b, err := layer2.BERC20(&mc, common.HexToAddress(mc.Env.PrivateKeyList[2][1]), mc.L1ERC20Address,
		// 				mc.L2ERC20Address, nil)
		// 			if err != nil {
		// 				fmt.Println("layer2.BERC20 err: ", err)
		// 			}

		// 			fmt.Println("l2 ERC20 balance: ", l2b)
		// 		}
		// 	}(waitGroup)
		// }

		// // deposit MNT
		// if strings.Contains(args[0], "c") {
		// 	waitGroup.Add(1)
		// 	go func(wg *sync.WaitGroup) {
		// 		defer wg.Done()
		// 		for {
		// 			tx, err := layer2.Deposit(&mc, mc.Env.PrivateKeyList[3][0], mc.L1MNTAddress,
		// 				mc.L2MNTAddress, big.NewInt(1))
		// 			if err != nil {
		// 				fmt.Println("DepositMNT err: ", err)
		// 			}
		// 			fmt.Println("DepositMNT txHash: ", tx.Hash())
		// 			time.Sleep(time.Duration(2) * time.Second)

		// 			_, l2b, err := layer2.BERC20(&mc, common.HexToAddress(mc.Env.PrivateKeyList[3][1]), mc.L1MNTAddress,
		// 				mc.L2MNTAddress, nil)
		// 			if err != nil {
		// 				fmt.Println("MNT Balance err: ", err)
		// 			}

		// 			fmt.Println("l2 MNT balance: ", l2b)
		// 		}
		// 	}(waitGroup)
		// }

		// // l2 一直部署erc20的合约
		// if strings.Contains(args[0], "d") {
		// 	waitGroup.Add(1)
		// 	go func(wg *sync.WaitGroup) {
		// 		defer wg.Done()

		// 		for {
		// 			address, err := stress.DL2CERC20(&mc, mc.Env.PrivateKeyList[4][0])
		// 			if err != nil {
		// 				fmt.Println("DL2CERC20 err: ", err)
		// 			}
		// 			time.Sleep(time.Second * time.Duration(5))
		// 			fmt.Println("ERC20 address: ", address)
		// 		}
		// 	}(waitGroup)
		// }

		// // Withdraw MNT
		// if strings.Contains(args[0], "e") {
		// 	waitGroup.Add(1)
		// 	go func(wg *sync.WaitGroup) {
		// 		defer wg.Done()
		// 		for {
		// 			tx, err := layer2.Withdraw(&mc, mc.Env.PrivateKeyList[5][0],
		// 				mc.L2MNTAddress, big.NewInt(1))
		// 			if err != nil {
		// 				fmt.Println("Withdraw MNT err: ", err)
		// 			}
		// 			fmt.Println("Withdraw MNT txHash: ", tx.Hash())
		// 			time.Sleep(time.Duration(2) * time.Second)

		// 			layer2.FinalizeMessage(&mc, tx.Hash())

		// 			l1b, _, err := layer2.BERC20(&mc, common.HexToAddress(mc.Env.PrivateKeyList[5][1]), mc.L1MNTAddress,
		// 				mc.L2MNTAddress, nil)
		// 			if err != nil {
		// 				fmt.Println("MNT Balance err: ", err)
		// 			}

		// 			fmt.Println("l1 MNT balance: ", l1b)
		// 		}
		// 	}(waitGroup)
		// }

		// // Withdraw ERC20
		// if strings.Contains(args[0], "f") {
		// 	waitGroup.Add(1)
		// 	go func(wg *sync.WaitGroup) {
		// 		defer wg.Done()
		// 		for {
		// 			tx, err := layer2.Withdraw(&mc, mc.Env.PrivateKeyList[6][0],
		// 				mc.L2ERC20Address, big.NewInt(1))
		// 			if err != nil {
		// 				fmt.Println("Withdraw ERC20 err: ", err)
		// 			}
		// 			fmt.Println("Withdraw ERC20 txHash: ", tx.Hash())
		// 			time.Sleep(time.Duration(2) * time.Second)

		// 			layer2.FinalizeMessage(&mc, tx.Hash())

		// 			l1b, _, err := layer2.BERC20(&mc, common.HexToAddress(mc.Env.PrivateKeyList[6][1]), mc.L1ERC20Address,
		// 				mc.L2ERC20Address, nil)
		// 			if err != nil {
		// 				fmt.Println("ERC20 Balance err: ", err)
		// 			}

		// 			fmt.Println("l1 ERC20 balance: ", l1b)
		// 		}
		// 	}(waitGroup)
		// }

		// // Withdraw ETH
		// if strings.Contains(args[0], "g") {
		// 	waitGroup.Add(1)
		// 	go func(wg *sync.WaitGroup) {
		// 		defer wg.Done()
		// 		for {
		// 			tx, err := layer2.Withdraw(&mc, mc.Env.PrivateKeyList[7][0],
		// 				mc.L2ETHAddress, big.NewInt(1))
		// 			if err != nil {
		// 				fmt.Println("Withdraw ETH err: ", err)
		// 			}
		// 			fmt.Println("Withdraw ETH txHash: ", tx.Hash())
		// 			time.Sleep(time.Duration(2) * time.Second)

		// 			layer2.FinalizeMessage(&mc, tx.Hash())

		// 			l1b, _, err := layer2.BETH(&mc, common.HexToAddress(mc.Env.PrivateKeyList[7][1]), nil)
		// 			if err != nil {
		// 				fmt.Println("layer2.BETH err: ", err)
		// 			}

		// 			fmt.Println("l1 ETH balance: ", l1b)
		// 		}
		// 	}(waitGroup)
		// }

		// // ERC20 transfer
		// if strings.Contains(args[0], "h") {
		// 	waitGroup.Add(1)
		// 	go func(wg *sync.WaitGroup) {
		// 		defer wg.Done()
		// 		for {
		// 			for retry := 0; retry < maxRetries; retry++ {
		// 				tx, err := lib.TransferERC20(mc.L2Client,
		// 					mc.L2ERC20Address,
		// 					mc.Env.PrivateKeyList[8][0],
		// 					common.HexToAddress("1"),
		// 					big.NewInt(1),
		// 				)
		// 				if err != nil {
		// 					if retry < maxRetries-1 {
		// 						// 间隔一段时间后重试
		// 						fmt.Println("TransferERC20 Retry", retry+1)
		// 						time.Sleep(time.Second * 5)
		// 						continue
		// 					}
		// 					fmt.Println("l2 TransferERC20 err: ", err)
		// 				}
		// 				fmt.Println("TransferERC20 txHash: ", tx.Hash())
		// 				break
		// 			}
		// 			time.Sleep(time.Duration(2) * time.Second)

		// 			_, l2b, err := layer2.BERC20(&mc, common.HexToAddress("1"), mc.L1ERC20Address, mc.L2ERC20Address, nil)
		// 			if err != nil {
		// 				fmt.Println("layer2.BERC20 err: ", err)
		// 			}
		// 			fmt.Println("l2 ERC20 balance: ", l2b)
		// 		}
		// 	}(waitGroup)
		// }

		// // MNT transfer
		// if strings.Contains(args[0], "i") {
		// 	waitGroup.Add(1)
		// 	go func(wg *sync.WaitGroup) {
		// 		defer wg.Done()
		// 		for {
		// 			for retry := 0; retry < maxRetries; retry++ {
		// 				tx, err := lib.TransferNT(mc.L2Client,
		// 					mc.Env.PrivateKeyList[9][0],
		// 					"2",
		// 					"1",
		// 					[]byte(""),
		// 				)
		// 				if err != nil {
		// 					if retry < maxRetries-1 {
		// 						// 间隔一段时间后重试
		// 						fmt.Println("TransferNT Retry", retry+1)
		// 						time.Sleep(time.Second * 5)
		// 						continue
		// 					}

		// 					fmt.Println("l2 Transfer MNT err: ", err)

		// 				}
		// 				if err := lib.CheckReceiptStatus(mc.L2Client, tx.Hash()); err != nil {
		// 					fmt.Println("l2 Transfer MNT CheckReceiptStatus err: ", err)
		// 				}
		// 				fmt.Println("TransferMNT txHash: ", tx.Hash())
		// 				break
		// 			}
		// 			time.Sleep(time.Duration(2) * time.Second)

		// 			_, l2b, err := layer2.BERC20(&mc, common.HexToAddress("2"), mc.L1MNTAddress, mc.L2MNTAddress, nil)
		// 			if err != nil {
		// 				fmt.Println("layer2.BERC20 err: ", err)
		// 			}
		// 			fmt.Println("l2 MNT balance: ", l2b)
		// 		}
		// 	}(waitGroup)
		// }

		// // transfer嵌套
		// if strings.Contains(args[0], "j") {
		// 	waitGroup.Add(1)
		// 	// 需要先给合约转账
		// 	amount, b := lib.ParseAmount("1234567890123456")
		// 	if !b {
		// 		return fmt.Errorf("ParseAmount return false")
		// 	}
		// 	_, err := lib.TransferERC20(mc.L2Client, mc.L2ERC20Address, mc.Env.PrivateKeyList[0][0], mc.L2ERC20Address, amount)
		// 	if err != nil {
		// 		return fmt.Errorf("TransferERC20 err: %v", err)
		// 	}
		// 	go func(wg *sync.WaitGroup) {
		// 		defer wg.Done()

		// 		contractInput := []interface{}{
		// 			mc.L2ERC20Address,
		// 			[]common.Address{
		// 				common.HexToAddress("3"),
		// 				common.HexToAddress("4"),
		// 				common.HexToAddress("5"),
		// 				common.HexToAddress("6"),
		// 				common.HexToAddress("7"),
		// 				common.HexToAddress("8"),
		// 				common.HexToAddress("9"),
		// 				common.HexToAddress("10"),
		// 				common.HexToAddress("11"),
		// 				common.HexToAddress("12"),
		// 			},
		// 		}
		// 		for {
		// 			tx, err := lib.Transfers(mc.L2Client,
		// 				mc.L2ERC20Address,
		// 				mc.Env.PrivateKeyList[10][0], contractInput)
		// 			if err != nil {
		// 				fmt.Println("Transfers err: ", err)
		// 			}
		// 			fmt.Println("L2ERC20.Transfers txHash: ", tx.Hash())
		// 			time.Sleep(time.Duration(2) * time.Second)

		// 			for _, to := range contractInput[1].([]common.Address) {
		// 				_, l2b, err := layer2.BERC20(&mc, to, mc.L1ERC20Address, mc.L2ERC20Address, nil)
		// 				if err != nil {
		// 					fmt.Println("layer2.BERC20 err: ", err)
		// 				}
		// 				fmt.Println("l2 ERC20 balance: ", l2b)
		// 			}

		// 		}
		// 	}(waitGroup)
		// }

		// // l2 SubscribeNewHead
		// if strings.Contains(args[0], "k") {
		// 	waitGroup.Add(1)

		// 	go func(wg *sync.WaitGroup) {
		// 		defer wg.Done()

		// 		layer2.SubscribeNewHead(*mc.L2WSClient)
		// 	}(waitGroup)
		// }

		// // SubscribeFilterLogs
		// if strings.Contains(args[0], "l") {
		// 	waitGroup.Add(1)

		// 	go func(wg *sync.WaitGroup) {
		// 		defer wg.Done()

		// 		layer2.SubscribeFilterLogs(*mc.L2WSClient, mc.Env.L2cdmAddress)
		// 	}(waitGroup)
		// }

		waitGroup.Wait()
		return nil
	},
	PreRunE: stressCmd.PersistentPreRunE,
}
