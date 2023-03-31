package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	MarginSmallLiabilityExchangeCoinList()
}

func MarginSmallLiabilityExchangeCoinList() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// MarginSmallLiabilityExchangeCoinListService - /sapi/v1/margin/exchange-small-liability
	marginSmallLiabilityExchangeCoinList, err := client.NewMarginSmallLiabilityExchangeCoinListService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginSmallLiabilityExchangeCoinList))
}
