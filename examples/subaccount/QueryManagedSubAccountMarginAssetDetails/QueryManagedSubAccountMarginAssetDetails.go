package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	QueryManagedSubAccountMarginAssetDetails()
}

func QueryManagedSubAccountMarginAssetDetails() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Query Managed Sub-account Margin Asset Details (For Investor Master Account) (USER_DATA)
	queryManagedSubAccountMarginAssetDetails, err := client.NewQueryManagedSubAccountMarginAssetDetailsService().Email("email@email.com").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(queryManagedSubAccountMarginAssetDetails))
}
