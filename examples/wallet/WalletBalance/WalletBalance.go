package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	WalletBalance()
}

func WalletBalance() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// WalletBalanceService - /sapi/v1/asset/wallet/balance
	walletBalance, err := client.NewWalletBalanceService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(walletBalance))
}
