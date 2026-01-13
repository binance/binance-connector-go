# Keep-Alive Configuration

```go
package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/giftcard"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	CreateASingleTokenGiftCard()
}

func CreateASingleTokenGiftCard() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.GiftCardRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
		common.WithKeepAlive(false),
	)
	apiClient := client.NewBinanceGiftCardClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.MarketDataAPI.CreateASingleTokenGiftCard(context.Background()).Token("token_example").Amount(1.0).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```
