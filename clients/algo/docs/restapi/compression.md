# Compression Configuration

```go
package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/algo"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	QueryHistoricalAlgoOrdersSpotAlgo()
}

func QueryHistoricalAlgoOrdersSpotAlgo() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.AlgoRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
		common.WithCompression(false),
	)
	apiClient := client.NewBinanceAlgoClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.SpotAlgoAPI.QueryHistoricalAlgoOrdersSpotAlgo(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```
