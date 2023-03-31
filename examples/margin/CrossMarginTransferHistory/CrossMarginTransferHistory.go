package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	CrossMarginTransferHistory()
}

func CrossMarginTransferHistory() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// CrossMarginTransferHistoryService - /sapi/v1/margin/transfer
	crossMarginTransferHistory, err := client.NewCrossMarginTransferHistoryService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(crossMarginTransferHistory))
}
