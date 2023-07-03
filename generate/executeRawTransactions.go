package generate

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

func ExecuteRawTransactions(c *ethclient.Client, txFile string) {
	// 打开包含原始交易数据的文件
	file, err := os.Open(txFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 从文件中读取每条原始交易数据，并提交交易
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		var signedTx types.Transaction
		err = rlp.DecodeBytes(common.Hex2Bytes(record[0]), &signedTx)
		if err != nil {
			panic(fmt.Sprintf("rlp.DecodeBytes err: %v", err))
		}

		// 下面这句有问题
		// from, err := types.Sender(types.HomesteadSigner{}, tx)
		from, err := types.Sender(types.LatestSignerForChainID(signedTx.ChainId()), &signedTx)
		if err != nil {
			log.Fatal(err)
		}

		// 提交交易
		if err := SubmitTransaction(c, &signedTx); err != nil {
			log.Printf("交易 %s 提交失败：%v", signedTx.Hash().Hex(), err)
		} else {
			log.Printf("%v 提交的交易 %s 提交成功！", from, signedTx.Hash().Hex())
		}

	}
}

func SubmitTransaction(c *ethclient.Client, tx *types.Transaction) error {

	err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		return err
	}

	// // 获取交易的签名
	// chainID, err := client.NetworkID(context.Background())
	// if err != nil {
	// 	return err
	// }

	// signer := types.NewEIP155Signer(chainID)
	// signedTx, err := tx.WithSignature(signer, crypto.PubkeyToAddress(tx.From()))
	// if err != nil {
	// 	return err
	// }

	// // 将交易发送到以太坊网络
	// err = client.SendTransaction(context.Background(), signedTx)
	// if err != nil {
	// 	return err
	// }

	return nil
}
