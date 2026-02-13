# Retries Configuration

```go
package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/pay"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	GetPayTradeHistory()
}

func GetPayTradeHistory() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.PayRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
		common.WithRetries(2),
	)
	apiClient := client.NewBinancePayClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.PayAPI.GetPayTradeHistory(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```
