# Proxy Configuration

```go
package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/alpha"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	GetExchangeInfo()
}

func GetExchangeInfo() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.AlphaRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
		common.WithProxy(common.ProxyConfig{
			Host:     "127.0.0.1",
			Port:     8080,
			Protocol: "http",
		}),
	)
	apiClient := client.NewBinanceAlphaClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.MarketDataAPI.GetExchangeInfo(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```
