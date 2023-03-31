package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	QuerySubAccountAssetsForMasterAccount()
}

func QuerySubAccountAssetsForMasterAccount() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Query Sub-account Assets (For Master Account)(USER_DATA)
	querySubAccountAssetsForMasterAccount, err := client.NewQuerySubAccountAssetsService().Email("from@email.com").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(querySubAccountAssetsForMasterAccount))
}
