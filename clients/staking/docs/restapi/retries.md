# Retries Configuration

```go
package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/staking"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	ClaimBoostRewards()
}

func ClaimBoostRewards() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.StakingRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
		common.WithRetries(2),
	)
	apiClient := client.NewBinanceStakingClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.SolStakingAPI.ClaimBoostRewards(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```
