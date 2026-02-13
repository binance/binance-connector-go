# Proxy Configuration

```go
package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/rebate"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	GetSpotRebateHistoryRecords()
}

func GetSpotRebateHistoryRecords() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.RebateRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
		common.WithProxy(common.ProxyConfig{
			Host:     "127.0.0.1",
			Port:     8080,
			Protocol: "http",
		}),
	)
	apiClient := client.NewBinanceRebateClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.RebateAPI.GetSpotRebateHistoryRecords(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```
