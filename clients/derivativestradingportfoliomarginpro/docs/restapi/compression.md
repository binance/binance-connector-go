# Compression Configuration

```go
package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomarginpro"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	GetPortfolioMarginProAccountInfo()
}

func GetPortfolioMarginProAccountInfo() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.DerivativesTradingPortfolioMarginProRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
		common.WithCompression(false),
	)
	apiClient := client.NewBinanceDerivativesTradingPortfolioMarginProClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.AccountAPI.GetPortfolioMarginProAccountInfo(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```
