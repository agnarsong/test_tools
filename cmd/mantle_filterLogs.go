package cmd

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"mantle/test/lib/layer2"
	"mantle/test/lib/layer2/bindings"
)

func init() {

	mantleCmd.AddCommand(filterLogsCmd)
}

var filterLogsCmd = &cobra.Command{
	Use:     "fl",
	Aliases: []string{"filterLogs"},
	Short:   "filterLogs",
	Long:    "filterLogs",
	RunE: func(cmd *cobra.Command, args []string) error {

		contractAbi, err := abi.JSON(strings.NewReader(bindings.L1CrossDomainMessengerABI))
		if err != nil {
			return fmt.Errorf("abi.JSON err: %v", err)
		}

		v := map[string]interface{}{}
		bn, li, err := layer2.GetMapLogs(
			mc.L1Client,
			common.HexToAddress("0x52753615226F8aC8a464bfecb11Ef798CFF3793f"),
			contractAbi,
			"SentMessage(address,address,bytes,uint256,uint256)",
			v,
		)
		if err != nil {
			return err
		}

		fmt.Println("blockNumber: ", bn)
		fmt.Println("logIndex: ", li)
		fmt.Println("target", v["target"])
		fmt.Println("sender", v["sender"])
		// Convert interface{} to []byte
		message, ok := v["message"].([]byte)
		if !ok {
			// Handle the case when the value is not []byte
			fmt.Println("Value is not []byte")
		}
		fmt.Println("message", common.Bytes2Hex(message))
		fmt.Println("messageNonce", v["messageNonce"])
		fmt.Println("gasLimit", v["gasLimit"])
		return nil
	},
	PersistentPreRunE: mantleCmd.PersistentPreRunE,
}
