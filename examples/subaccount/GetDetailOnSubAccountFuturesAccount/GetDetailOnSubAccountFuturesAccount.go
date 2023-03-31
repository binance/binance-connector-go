package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	GetDetailOnSubAccountFuturesAccount()
}

func GetDetailOnSubAccountFuturesAccount() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Get Detail on Sub-account's Futures Account (For Master Account) - /sapi/v1/sub-account/futures/account
	getDetailOnSubAccountFuturesAccount, err := client.NewGetDetailOnSubAccountFuturesAccountService().Email("from@email.com").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(getDetailOnSubAccountFuturesAccount))
}
