package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	GetAllMarginPairs()
}

func GetAllMarginPairs() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// GetAllMarginPairsService - /sapi/v1/margin/allPairs
	getAllMarginPairs, err := client.NewGetAllMarginPairsService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(getAllMarginPairs))
}
