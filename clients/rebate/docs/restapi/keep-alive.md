# Keep-Alive Configuration

```go
package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/rebate"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	GetSpotRebateHistoryRecords()
}

func GetSpotRebateHistoryRecords() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.RebateRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
		common.WithKeepAlive(false),
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
