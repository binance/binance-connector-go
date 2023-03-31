package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	QueryManagedSubAccountAssetDetails()
}

func QueryManagedSubAccountAssetDetails() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Query Managed Sub-account Asset Details（For Investor Master Account）- /sapi/v1/sub-account/managed-subaccount/asset
	queryManagedSubAccountAssetDetails, err := client.NewQueryManagedSubAccountAssetDetailsService().Email("email@email.com").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(queryManagedSubAccountAssetDetails))
}
