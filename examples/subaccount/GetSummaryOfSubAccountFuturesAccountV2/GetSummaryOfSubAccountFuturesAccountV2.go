package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	GetSummaryOfSubAccountFuturesAccountV2()
}

func GetSummaryOfSubAccountFuturesAccountV2() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Get Summary of Sub-account's Futures Account V2 (For Master Account) - /sapi/v1/sub-account/futures/accountSummary
	getSummaryOfSubAccountFuturesAccountV2, err := client.NewGetSummaryOfSubAccountFuturesAccountV2Service().FuturesType(1).
		Page(1).Limit(10).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(getSummaryOfSubAccountFuturesAccountV2))
}
