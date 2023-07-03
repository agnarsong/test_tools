package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		jsonData := `{"name":"Mantle","logoURI":"https://token-list.mantle.xyz/mantle_logo.svg","keywords":["scaling","layer2","infrastructure"],"timestamp":"2023-02-08T09:38:45.488Z","tokens":[{"chainId":31337,"address":"0x92aBAD50368175785e4270ca9eFd169c949C4ce1","name":"BitDAO","symbol":"BIT","decimals":18,"logoURI":"https://token-list.mantle.xyz/data/BitDAO/logo.svg","extensions":{"optimismBridgeAddress":"0x52753615226F8aC8a464bfecb11Ef798CFF3793f"}}],"version":{"major":1,"minor":0,"patch":0}}`
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(jsonData))
	})

	http.ListenAndServe(":9081", nil)
}
