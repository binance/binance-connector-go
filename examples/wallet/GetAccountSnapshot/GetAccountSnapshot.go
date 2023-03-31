package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	AccountSnapshot()
}

func AccountSnapshot() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// GetAccountSnapshotService get all orders from account - /sapi/v1/accountSnapshot
	accountSnapshot, err := client.NewGetAccountSnapshotService().MarketType("SPOT").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(accountSnapshot))
}
