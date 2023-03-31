package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	AggTradesList()
}

func AggTradesList() {
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient("", "", baseURL)

	// AggTradesList
	aggTradesList, err := client.NewAggTradesListService().
		Symbol("BTCUSDT").Limit(20).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(aggTradesList))
}
