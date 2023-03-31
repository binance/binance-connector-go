package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	GetAllMarginAssets()
}

func GetAllMarginAssets() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// GetAllMarginAssetsService - /sapi/v1/margin/allAssets
	getAllMarginAssets, err := client.NewGetAllMarginAssetsService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(getAllMarginAssets))
}
