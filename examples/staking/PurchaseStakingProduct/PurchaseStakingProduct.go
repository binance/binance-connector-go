package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	PurchaseStakingProduct()
}

func PurchaseStakingProduct() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	purchaseStakingProduct, err := client.NewPurchaseStakingProductService().Product("STAKING").
		ProductId("AXS*90").Amount(90.00).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(purchaseStakingProduct))
}
