package main

import (
	"context"
	"fmt"
	"time"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	GetC2CTradeHistory()
}

func GetC2CTradeHistory() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// GetC2CTradeHistoryService - /sapi/v1/c2c/orderMatch/listUserOrderHistory
	getC2CTradeHistory, err := client.NewGetC2CTradeHistoryService().
		Timestamp(uint64(time.Now().UnixMilli())).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(binance_connector.PrettyPrint(getC2CTradeHistory))
}
