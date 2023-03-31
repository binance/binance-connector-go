package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	APIKeyPermission()
}

func APIKeyPermission() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// APIKeyPermissionService - /sapi/v1/account/apiRestrictions
	apiKeyPermission, err := client.NewAPIKeyPermissionService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(apiKeyPermission))
}
