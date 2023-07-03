package stress

import (
	"encoding/csv"
	"fmt"

	"math/big"
	"os"
	"time"

	"mantle/test/lib"
	"mantle/test/lib/layer2"
	"mantle/test/lib/layer2/bindings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Distribute Native Token
func DNT(mc *layer2.MantleCenter, layer string, num int) error {

	c := mc.L1Client
	if layer == "l2" {
		c = mc.L2Client
	}

	//文件读取
	f, err := os.Open("./wallets.csv")
	if err != nil {
		return err
	}
	reader := csv.NewReader(f)
	preData, err := reader.ReadAll()
	if err != nil {
		return err
	}

	fmt.Printf("开始执行Distribute Native Token, layer: %v\n", layer)
	// 多协程初始化数据
	for i := 0; i < min(len(preData), num); i++ {

		tx, err := lib.TransferNT(c,
			mc.Env.PrivateKeyList[0][0],
			preData[i][1],
			mc.Env.Amount,
			[]byte(""),
		)
		if err != nil {
			return err
		}

		fmt.Printf("第 %v 次transfer,txHash: %v\n", i+1, tx.Hash())
		if err := lib.CheckReceiptStatus(c, tx.Hash()); err != nil {
			return err
		}

		time.Sleep(time.Microsecond * time.Duration(400))

	}

	fmt.Printf("完成执行Distribute Native Token, layer: %v\n", layer)
	return nil
}

func DeployL1CustomERC20(mc *layer2.MantleCenter) error {

	auth, err := lib.GetAuth(mc.L1Client, mc.Env.PrivateKeyList[0][0])
	if err != nil {
		return fmt.Errorf("lib.GetAuth err: %v", err)
	}
	// fmt.Println(auth.GasPrice)
	// auth.GasPrice.Mul(auth.GasPrice, big.NewInt(2))
	// fmt.Println(auth.GasPrice)

	var tx *types.Transaction
	mc.L1ERC20Address, tx, mc.L1ERC20, err = bindings.DeployL1CustomERC20(
		auth, mc.L1Client, "L1CustomERC20", "L1E")
	if err != nil {
		return fmt.Errorf("bindings.DeployL1CustomERC20 err: %v", err)
	}
	fmt.Println("deployHash: ", tx.Hash())
	if err := lib.CheckReceiptStatus(mc.L1Client, tx.Hash()); err != nil {
		return fmt.Errorf("lib.CheckReceiptStatus err: %v", err)
	}

	mc.Env.L1ERC20Address = mc.L1ERC20Address.Hex()

	return nil
}

func DeployL2CustomERC20(mc *layer2.MantleCenter) error {

	auth, err := lib.GetAuth(mc.L2Client, mc.Env.PrivateKeyList[0][0])
	if err != nil {
		return fmt.Errorf("lib.GetAuth err: %v", err)
	}

	var tx *types.Transaction
	mc.L2ERC20Address, tx, mc.L2ERC20, err = bindings.DeployL2CustomERC20(
		auth,
		mc.L2Client,
		mc.L2BridgeAddress,
		mc.L1ERC20Address,
	)
	if err != nil {
		return fmt.Errorf("bindings.DeployL2CustomERC20 err: %v", err)
	}

	if err := lib.CheckReceiptStatus(mc.L2Client, tx.Hash()); err != nil {
		return fmt.Errorf("lib.CheckReceiptStatus err: %v", err)
	}

	mc.Env.L2ERC20Address = mc.L2ERC20Address.Hex()

	return nil
}

func DL2CERC20(mc *layer2.MantleCenter, pri string) (string, error) {

	auth, err := lib.GetAuth(mc.L2Client, pri)
	if err != nil {
		return "", fmt.Errorf("lib.GetAuth err: %v", err)
	}

	var tx *types.Transaction
	L2ERC20Address, tx, _, err := bindings.DeployL2CustomERC20(
		auth,
		mc.L2Client,
		mc.L2BridgeAddress,
		common.HexToAddress("0xE01f935bE5033DD6A82830c2Acf73Eb78F9A12EC"),
	)

	if err != nil {
		return "", fmt.Errorf("bindings.DeployL2CustomERC20 err: %v", err)
	}

	if err := lib.CheckReceiptStatus(mc.L2Client, tx.Hash()); err != nil {
		return "", fmt.Errorf("lib.CheckReceiptStatus err: %v", err)
	}
	fmt.Println("deploy txHash: ", tx.Hash())
	return L2ERC20Address.Hex(), nil
}

// Distribute ERC20
func D20(c *ethclient.Client, tokenAddress common.Address, pri string, amount string) error {

	//文件读取
	f, err := os.Open("./wallets.csv")
	if err != nil {
		return err
	}
	reader := csv.NewReader(f)
	preData, err := reader.ReadAll()
	if err != nil {
		return err
	}

	amountB := new(big.Int)
	amountB.SetString(amount, 10)
	if amount == "-1" {
		// 1 << 90
		amountB = new(big.Int).Lsh(big.NewInt(1), 90)
	}

	// 多协程初始化数据
	for i := 0; i < len(preData); i++ {

		tx, err := lib.TransferERC20(c,
			tokenAddress,
			pri,
			common.HexToAddress(preData[i][1]),
			amountB,
		)
		if err != nil {
			return err
		}

		fmt.Printf("第 %v 次transfer,txHash: %v\n", i+1, tx.Hash())

		time.Sleep(time.Microsecond * time.Duration(400))
	}

	return nil
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
