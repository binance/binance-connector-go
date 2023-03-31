package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	GetSummaryOfSubAccountFuturesAccount()
}

func GetSummaryOfSubAccountFuturesAccount() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Get Summary of Sub-account's Futures Account (For Master Account) - /sapi/v1/sub-account/futures/accountSummary
	getSummaryOfSubAccountFuturesAccount, err := client.NewGetSummaryOfSubAccountFuturesAccountService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(getSummaryOfSubAccountFuturesAccount))
}
