package generate

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"mantle/test/lib"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

func GRawTransactions(c *ethclient.Client, numberOfTransactions int, wFileName string, tFileName string) {
	// 打开包含私钥和地址的文件
	file, err := os.Open(wFileName)
	if err != nil {
		fmt.Println("os.Open err: ", err)
	}
	defer file.Close()

	// 创建名为 RawTransactions.csv 的文件，并写入表头
	transactionsFile, err := os.Create(tFileName)
	if err != nil {
		fmt.Println("os.Create err: ", err)
	}
	defer transactionsFile.Close()

	transactionsWriter := csv.NewWriter(transactionsFile)
	defer transactionsWriter.Flush()

	// 从文件中读取每个钱包的私钥和地址，并为每个钱包创建一笔交易
	reader := csv.NewReader(file)
	for i := 0; i < numberOfTransactions; i++ {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("reader.Read err: ", err)
		}
		privateKey, err := crypto.HexToECDSA(record[0])
		if err != nil {
			fmt.Println(record[0])
			fmt.Println("crypto.HexToECDSA err: ", err)
			continue
		}

		address := common.HexToAddress(record[1])
		nonce, gasPrice, chainID, err := lib.QBasic(c, address)
		if err != nil {
			fmt.Println("lib.QBasic err: ", err)
		}

		// 构造交易数据
		gasLimit := uint64(21000) // 固定的 gasLimit

		to := common.HexToAddress("0x1234567890123456789012345678901234567890") // 接收方地址

		var data []byte
		tx := types.NewTransaction(nonce, to, big.NewInt(1), gasLimit, gasPrice, data)

		// 使用钱包的私钥进行签名
		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
		if err != nil {
			fmt.Println("types.SignTx err: ", err)
		}

		txb, err := rlp.EncodeToBytes(signedTx)
		if err != nil {
			fmt.Println("rlp.EncodeToBytes err: ", err)
		}

		if err := transactionsWriter.Write([]string{common.Bytes2Hex(txb)}); err != nil {
			fmt.Println("transactionsWriter.Writ err: ", err)
		}
	}

	log.Printf("%d 笔交易已生成并写入文件！", numberOfTransactions)
}
