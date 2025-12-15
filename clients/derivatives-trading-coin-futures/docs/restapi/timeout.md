# Timeout

```go
package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	ExchangeInformation()
}

func ExchangeInformation() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.DerivativesTradingCoinFuturesRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
		common.WithTimeout(2000),
	)
	apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.MarketDataAPI.ExchangeInformation(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```
