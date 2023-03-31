package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	GetDetailOnSubAccountMarginAccount()
}

func GetDetailOnSubAccountMarginAccount() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Get Detail on Sub-account's Margin Account (For Master Account) - /sapi/v1/sub-account/margin/account
	getDetailOnSubAccountMarginAccount, err := client.NewGetDetailOnSubAccountMarginAccountService().Email("from@email.com").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(getDetailOnSubAccountMarginAccount))
}
