package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	QueryMarginAvailableInventory()
}

func QueryMarginAvailableInventory() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// QueryMarginAvailableInventoryService - /sapi/v1/margin/available-inventory
	queryMarginAvailableInventory, err := client.NewQueryMarginAvailableInventoryService().
		MarginType("MARGIN").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(queryMarginAvailableInventory))
}
