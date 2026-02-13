# Keep-Alive Configuration

```go
package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/spot/src"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	ExchangeInfo()
}

func ExchangeInfo() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithKeepAlive(false),
	)
	apiClient := client.NewBinanceSpotClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.GeneralAPI.ExchangeInfo(context.Background()).Symbol("BTCUSDT").Execute()
	if err != nil {
		log.Println(err)
		return
	}

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```
