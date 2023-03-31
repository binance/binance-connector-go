package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	MarginInterestRateHistory()
}

func MarginInterestRateHistory() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// MarginInterestRateHistoryService - /sapi/v1/margin/interestRateHistory
	marginInterestRateHistory, err := client.NewMarginInterestRateHistoryService().Asset("USDT").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginInterestRateHistory))
}
