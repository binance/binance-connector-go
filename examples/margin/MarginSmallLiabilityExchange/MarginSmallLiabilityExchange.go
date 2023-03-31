package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	MarginSmallLiabilityExchange()
}

func MarginSmallLiabilityExchange() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// MarginSmallLiabilityExchangeService - /sapi/v1/margin/exchange-small-liability
	marginSmallLiabilityExchange, err := client.NewMarginSmallLiabilityExchangeService().
		AssetNames("BTC,ETH,BNB").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginSmallLiabilityExchange))
}
