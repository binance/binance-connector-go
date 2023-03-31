package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	GetSubAccountDepositAddress()
}

func GetSubAccountDepositAddress() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Get Sub-account Deposit Address (For Master Account) - /sapi/v1/capital/deposit/subAddress
	getSubAccountDepositAddress, err := client.NewGetSubAccountDepositAddressService().Email("from@email.com").
		Coin("BTC").Network("BTC").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(getSubAccountDepositAddress))
}
