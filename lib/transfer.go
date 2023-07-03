package lib

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func SignETHTx1(c *ethclient.Client, prv string, to string, amount string, data []byte) (
	tx *types.Transaction, chainID *big.Int, privateKey *ecdsa.PrivateKey, err error) {

	gasLimit := uint64(15_000_000)

	amountB, b := ParseAmount(amount)
	if !b {
		return tx, chainID, privateKey, fmt.Errorf("ParseAmount return false")
	}

	privateKey, _, fromAddress, err := AnalysePrivateKey(prv)
	if err != nil {
		return nil, common.Big0, nil, fmt.Errorf("AnalysePrivateKey err: %v", err)
	}

	nonce, gasPrice, chainID, err := QBasic(c, fromAddress)

	if err != nil {
		return nil, common.Big0, nil, fmt.Errorf("QBasic err: %v", err)
	}

	toAddress := common.HexToAddress(to)
	tx = types.NewTransaction(nonce, toAddress, amountB, gasLimit, gasPrice, data)

	return
}

func SignETHTx2(tx *types.Transaction, chainID *big.Int, prv *ecdsa.PrivateKey) (
	signedTx *types.Transaction, err error) {

	signedTx, err = types.SignTx(tx, types.NewEIP155Signer(chainID), prv)
	return
}

func SendTransaction(c *ethclient.Client, signedTx *types.Transaction) (txHash string, err error) {
	err = c.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return
	}

	txHash = signedTx.Hash().Hex()
	return
}

func CheckReceiptStatus(c *ethclient.Client, txHash common.Hash) error {

	fmt.Println("will to check Txhash: ", txHash.Hex())
	maxRetries := 3

OuterLoop:
	for {
		// 最大重试次数为maxRetries
		for retries := 0; retries <= maxRetries; retries++ {
			_, _, err := c.TransactionByHash(context.Background(), txHash)
			if err == ethereum.NotFound {
				// 请求失败，打印错误信息，并等待一段时间后重试
				fmt.Printf("TransactionByHash 重试 %d: %v\n", retries+1, err)
				time.Sleep(2 * time.Second)

				continue
			} else if err != nil {
				return fmt.Errorf("TransactionByHash err: %v", err)
			}

			break OuterLoop
		}

		fmt.Println("ReceiptStatus checking: TransactionByHash……")
		time.Sleep(time.Second * time.Duration(2))
	}

	for i := 0; ; i++ {
		re, err := c.TransactionReceipt(context.Background(), txHash)
		if err == ethereum.NotFound {
			// 请求失败，打印错误信息，并等待一段时间后重试
			fmt.Printf("TransactionReceipt 重试 %d: %v\n", i+1, err)
			time.Sleep(5 * time.Second)

			continue
		} else if err != nil {
			return fmt.Errorf("TransactionReceipt err: %v", err)
		}

		if re.Status == 1 {
			break
		}
		if re.Status != 1 && i < 5 {

			fmt.Println("ReceiptStatus checking: TransactionReceipt……")
			time.Sleep(time.Second * time.Duration(2))
			continue
		} else {
			return fmt.Errorf("re.Status is not 1")
		}
	}

	return nil
}

func TransferERC20(c *ethclient.Client, tokenAddress common.Address,
	pri string, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {

	privateKey, _, fromAddress, err := AnalysePrivateKey(pri)
	if err != nil {
		return nil, err
	}

	transferFnSignature := []byte("transfer(address,uint256)")
	h := sha3.NewLegacyKeccak256()
	h.Write(transferFnSignature)
	methodID := h.Sum(nil)[:4]

	// fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	// fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	// fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := c.EstimateGas(context.Background(), ethereum.CallMsg{
		From: fromAddress,
		To:   &tokenAddress,
		Data: data,
	})
	if err != nil {
		return nil, err
	}

	nonce, gasPrice, chainID, err := QBasic(c, fromAddress)
	if err != nil {
		return nil, err
	}
	tx := types.NewTransaction(nonce, tokenAddress, big.NewInt(0), gasLimit, gasPrice, data)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, err
	}

	err = c.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func ApproveERC20(c *ethclient.Client, tokenAddress common.Address,
	pri string, _spender common.Address, amount *big.Int) (*types.Transaction, error) {

	privateKey, _, fromAddress, err := AnalysePrivateKey(pri)
	if err != nil {
		return nil, err
	}

	transferFnSignature := []byte("approve(address,uint256)")
	h := sha3.NewLegacyKeccak256()
	h.Write(transferFnSignature)
	methodID := h.Sum(nil)[:4]

	// fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb
	paddedAddress := common.LeftPadBytes(_spender.Bytes(), 32)
	// fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	// fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := c.EstimateGas(context.Background(), ethereum.CallMsg{
		From: fromAddress,
		To:   &tokenAddress,
		Data: data,
	})
	if err != nil {
		return nil, err
	}

	nonce, gasPrice, chainID, err := QBasic(c, fromAddress)
	if err != nil {
		return nil, err
	}
	tx := types.NewTransaction(nonce, tokenAddress, big.NewInt(0), gasLimit, gasPrice, data)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, err
	}

	err = c.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func TransferNT(c *ethclient.Client, prv string, to string,
	amount string, data []byte) (tx *types.Transaction, err error) {
	tx, cid, privateKey, err := SignETHTx1(c, prv, to, amount, data)
	if err != nil {
		return tx, err
	}

	signedTx, err := SignETHTx2(tx, cid, privateKey)
	if err != nil {
		return tx, err
	}

	_, err = SendTransaction(c, signedTx)
	if err != nil {
		return tx, err
	}

	return signedTx, err
}

func Transfers(c *ethclient.Client, tokenAddress common.Address,
	pri string, contractInput []interface{}) (*types.Transaction, error) {

	privateKey, _, fromAddress, err := AnalysePrivateKey(pri)
	if err != nil {
		return nil, err
	}

	// 构建合约调用参数
	contractABI := `[{"constant":true,"inputs":[{"name":"_token","type":"address"},{"name":"to","type":"address[]"}],"name":"transfers","outputs":[],"payable":true,"stateMutability":"payable","type":"function"}]`
	contractMethod := "transfers"

	// 构建合约调用数据
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		return nil, err
	}

	data, err := parsedABI.Pack(contractMethod, contractInput...)
	if err != nil {
		return nil, err
	}

	// 构建交易参数
	gasLimit := uint64(3000000)

	nonce, gasPrice, chainID, err := QBasic(c, fromAddress)
	if err != nil {
		return nil, err
	}

	// 创建交易
	tx := types.NewTransaction(nonce, tokenAddress, big.NewInt(0), gasLimit, gasPrice, data)

	// 使用私钥签名交易
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, err
	}

	// 发送交易
	err = c.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, err
	}

	return signedTx, nil

}
