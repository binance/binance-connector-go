package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	CloudMiningPaymentHistory()
}

func CloudMiningPaymentHistory() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// CloudMiningPaymentHistoryService - /sapi/v1/asset/ledger-transfer/cloud-mining/queryByPage
	cloudMiningPaymentHistory, err := client.NewCloudMiningPaymentHistoryService().
		StartTime(1664442061000).EndTime(1664442078000).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(cloudMiningPaymentHistory))
}
