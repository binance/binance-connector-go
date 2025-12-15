# Compression Configuration

```go
package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/copytrading"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	GetFuturesLeadTraderStatus()
}

func GetFuturesLeadTraderStatus() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.CopyTradingRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
		common.WithCompression(false),
	)
	apiClient := client.NewBinanceCopyTradingClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.FutureCopyTradingAPI.GetFuturesLeadTraderStatus(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```
