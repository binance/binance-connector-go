package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	GetSymbolsDelistSchedule()
}

func GetSymbolsDelistSchedule() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, binance_connector.SIGNATURE_HMAC_SHA256, baseURL)

	// GetSymbolsDelistScheduleService - /sapi/v1/spot/delist-schedule
	getSymbolsDelistSchedule, err := client.NewGetSymbolsDelistScheduleService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(getSymbolsDelistSchedule))
}
