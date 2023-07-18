package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type TX struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	TransactionIndex  string `json:"transactionIndex"`
	From              string `json:"from"`
	To                string `json:"to"`
	Value             string `json:"value"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	IsError           string `json:"isError"`
	TxReceiptStatus   string `json:"txreceipt_status"`
	Input             string `json:"input"`
	ContractAddress   string `json:"contractAddress"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	GasUsed           string `json:"gasUsed"`
	Confirmations     string `json:"confirmations"`
	MethodId          string `json:"methodId"`
	FunctionName      string `json:"functionName"`
}

type RESP struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []TX   `json:"result"`
}

func main() {

	startBlockNum := 8814900
	offset := 500
	sort := "asc"
	// mtbatcher
	// address := "0x1409f6cb6dc13b2a557271750ccc1524be098eba"
	// sequencer
	// address := "0xc4AaE221f1C62E8CBC657Af5b051eA573914cFc7"
	// proposer
	address := "0x88EC574e2ef0EcF9043373139099f7E535F94dBC"
	tmpURL := "https://api-goerli.etherscan.io/api?module=account&action=txlist&" +
		"address=%v&startblock=%v&endblock=99999999&page=1&offset=%v&sort=%v&" +
		"apikey=B1XAN986315AME96W9QK7X1RGQ6WJMWEPW"

	for {

		url := fmt.Sprintf(tmpURL, address, startBlockNum, offset, sort)
		fmt.Println(url)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("http.NewRequest err: ", err)
			break
		}

		req.Header = http.Header{
			"Host":          {"www.host.com"},
			"Content-Type":  {"application/json"},
			"Authorization": {"Bearer Token"},
			"user-agent":    {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36 Edg/109.0.1518.78"},
		}

		// 请求JSON返回的数据
		client := &http.Client{}
		response, err := client.Do(req)
		if err != nil {
			fmt.Println("client.Do err: ", err)
			break
		}

		// 延时关闭
		defer response.Body.Close()

		// 返回的结果resp.Body
		body, _ := io.ReadAll(response.Body)

		var resp RESP
		// 把请求到的数据Unmarshal到res中
		err = json.Unmarshal(body, &resp)
		if err != nil {
			fmt.Println("json.Unmarshal err: ", err)
			break
		}
		if len(resp.Result) == 0 {
			fmt.Println("未查询到更新的数据！")
			break
		}
		for _, res := range resp.Result {

			input, _ := hex.DecodeString(res.Input[2:])
			fmt.Println(res.BlockNumber, res.Hash, len(input), res.TimeStamp, "", "", "", res.GasPrice, res.GasUsed, "", "", "", "", "")
		}

		tmpBlock, err := strconv.Atoi(resp.Result[len(resp.Result)-1].BlockNumber)

		if err != nil {
			fmt.Println("strconv.Atoi err: ", err)
			break
		}
		startBlockNum = tmpBlock + 1

		time.Sleep(time.Duration(5) * time.Second)
	}
}
