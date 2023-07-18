package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func executeHTTPRequest(method, url, body, contentType string) ([]byte, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go 'curl_command'")
		return
	}

	curlCommand := os.Args[1]

	// 解析 curl 命令
	// 提取方法、URL、请求体和内容类型等信息
	method := "GET"
	url := ""
	body := ""
	contentType := "application/json"

	// 提取方法
	methodIndex := strings.Index(curlCommand, " -X ")
	if methodIndex != -1 {
		endIndex := strings.Index(curlCommand[methodIndex+4:], " ")
		method = curlCommand[methodIndex+4 : methodIndex+4+endIndex]
	}
	fmt.Println(method)
	// 提取 URL
	urlIndex := strings.Index(curlCommand, " '")
	if urlIndex != -1 {
		endIndex := strings.Index(curlCommand[urlIndex+2:], "'")
		url = curlCommand[urlIndex+2 : urlIndex+2+endIndex]
	}
	fmt.Println(url)

	// 提取请求体
	dataIndex := strings.Index(curlCommand, " --data ")
	if dataIndex != -1 {
		endIndex := strings.LastIndex(curlCommand, "'")
		body = curlCommand[dataIndex+8 : endIndex]
	}

	responseBody, err := executeHTTPRequest(method, url, body, contentType)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response:", string(responseBody))
}
