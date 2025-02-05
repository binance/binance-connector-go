package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	GetFuturesLeadTraderStatus()
}

func GetFuturesLeadTraderStatus() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Get Futures Lead Trader Status - /sapi/v1/copyTrading/futures/userStatus
	getFuturesLeadTraderStatus, err := client.NewGetFuturesLeadTraderStatusService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(getFuturesLeadTraderStatus))
}
