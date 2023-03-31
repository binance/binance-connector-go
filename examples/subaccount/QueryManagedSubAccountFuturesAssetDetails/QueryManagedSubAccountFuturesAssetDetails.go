package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	QueryManagedSubAccountFuturesAssetDetails()
}

func QueryManagedSubAccountFuturesAssetDetails() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Query Managed Sub-account Futures Asset Details（For Investor Master Account）(USER_DATA)
	queryManagedSubAccountFuturesAssetDetails, err := client.NewQueryManagedSubAccountFuturesAssetDetailsService().Email("email@email.com").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(queryManagedSubAccountFuturesAssetDetails))
}
