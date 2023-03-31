package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	GetIPRestrictionForSubAccountAPIKey()
}

func GetIPRestrictionForSubAccountAPIKey() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Get IP Restriction for a Sub-account API Key (For Master Account) - /sapi/v1/sub-account/subaccountApi/ipRestriction
	getIPRestrictionForSubAccountAPIKey, err := client.NewGetIPRestrictionForSubAccountAPIKeyService().Email("email@email.com").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(getIPRestrictionForSubAccountAPIKey))
}
