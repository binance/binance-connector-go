package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	GetFiatDepositWithdrawHistory()
}

func GetFiatDepositWithdrawHistory() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// GetFiatDepositWithdrawHistoryService - /sapi/v1/fiat/orders
	getFiatDepositWithdrawHistory, err := client.NewGetFiatDepositWithdrawHistoryService().
		TransactionType("0").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(getFiatDepositWithdrawHistory))
}
