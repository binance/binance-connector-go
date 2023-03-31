package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	QueryCrossMarginPair()
}

func QueryCrossMarginPair() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// QueryCrossMarginPairService - /sapi/v1/margin/pair
	queryCrossMarginPair, err := client.NewQueryCrossMarginPairService().Symbol("BTCUSDT").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(queryCrossMarginPair))
}
