package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	AccountApiTradingStatus()
}

func AccountApiTradingStatus() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// AccountApiTradingStatusService - /sapi/v1/account/apiTradingStatus
	accountApiTradingStatus, err := client.NewAccountApiTradingStatusService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(accountApiTradingStatus))
}
