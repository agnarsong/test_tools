package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// Define the HTTP request payload
	payload := map[string]interface{}{
		"query": `
			mutation {
				createUser(
					id: "1",
					name: "John Smith"
				) {
					id
					name
				}
			}
		`,
	}

	// Serialize the payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	// Send the HTTP request to the Subgraph endpoint
	resp, err := http.Post("https://graph.testnet.mantle.xyz/subgraphs/name/agnar/test1", "application/json", bytes.NewReader(jsonPayload))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected response status code: %d\n", resp.StatusCode)
		return
	}

	// Print the response body
	var responseBody struct {
		Data json.RawMessage `json:"data"`
	}
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response body: %s\n", responseBody.Data)
}
