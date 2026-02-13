# Keep-Alive Configuration

```go
package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/mining"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	AcquiringAlgorithm()
}

func AcquiringAlgorithm() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.MiningRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
		common.WithKeepAlive(false),
	)
	apiClient := client.NewBinanceMiningClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.MiningAPI.AcquiringAlgorithm(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```
