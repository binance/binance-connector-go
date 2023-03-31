package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	WithdrawAssetsFromTheManagedSubAccount()
}

func WithdrawAssetsFromTheManagedSubAccount() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	withdrawAssetsFromTheManagedSubAccount, err := client.NewWithdrawAssetsFromTheManagedSubAccountService().FromEmail("email@email.com").
		Asset("BTC").Amount(1.5).TransferDate(123132123).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(withdrawAssetsFromTheManagedSubAccount))
}
