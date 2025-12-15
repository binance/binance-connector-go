# Timeout

```go
package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/convert"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	ListAllConvertPairs()
}

func ListAllConvertPairs() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.ConvertRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
		common.WithTimeout(2000),
	)
	apiClient := client.NewBinanceConvertClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.MarketDataAPI.ListAllConvertPairs(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```
