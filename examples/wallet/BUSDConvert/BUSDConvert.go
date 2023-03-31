package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	BUSDConvert()
}

func BUSDConvert() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// BUSDConvertService - /sapi/v1/asset/convert-transfer
	bUSDConvert, err := client.NewBUSDConvertService().ClientTranId("118263407119").
		Asset("BUSD").Amount(20.0).AccountType("MAIN").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(bUSDConvert))
}
