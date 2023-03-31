package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	BUSDConvertHistory()
}

func BUSDConvertHistory() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// BUSDConvertHistoryService - /sapi/v1/asset/convert-transfer/queryByPage
	bUSDConvertHistory, err := client.NewBUSDConvertHistoryService().
		StartTime(1664442061000).EndTime(1664442078000).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(bUSDConvertHistory))
}
