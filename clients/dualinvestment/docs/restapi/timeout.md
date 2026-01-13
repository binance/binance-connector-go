# Timeout

```go
package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/dualinvestment"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	GetDualInvestmentPositions()
}

func GetDualInvestmentPositions() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.DualInvestmentRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
		common.WithTimeout(2000),
	)
	apiClient := client.NewBinanceDualInvestmentClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.TradeAPI.GetDualInvestmentPositions(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```
