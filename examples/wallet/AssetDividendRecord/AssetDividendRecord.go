package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	AssetDividendRecord()
}

func AssetDividendRecord() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// AssetDividendRecordService - /sapi/v1/asset/assetDividend
	assetDividendRecord, err := client.NewAssetDividendRecordService().Asset("BTC").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(assetDividendRecord))
}
