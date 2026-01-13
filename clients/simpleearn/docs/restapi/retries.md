# Retries Configuration

```go
package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/simpleearn"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	ClaimBoostRewards()
}

func ClaimBoostRewards() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SimpleEarnRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
		common.WithRetries(2),
	)
	apiClient := client.NewBinanceSimpleEarnClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.FlexibleLockedAPI.GetSimpleEarnFlexibleProductList(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```
