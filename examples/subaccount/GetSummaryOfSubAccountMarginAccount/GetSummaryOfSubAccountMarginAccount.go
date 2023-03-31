package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	GetSummaryOfSubAccountMarginAccount()
}

func GetSummaryOfSubAccountMarginAccount() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Get Summary of Sub-account's Margin Account (For Master Account) - /sapi/v1/sub-account/margin/accountSummary
	getSummaryOfSubAccountMarginAccount, err := client.NewGetSummaryOfSubAccountMarginAccountService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(getSummaryOfSubAccountMarginAccount))
}
