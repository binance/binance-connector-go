# Time Unit

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"

	client "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	AggTrades()
}

func AggTrades() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithTimeUnit(common.MICROSECOND),
	)
	apiClient := client.NewBinanceSpotClient(client.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketAPI.AggTrades(context.Background()).Symbol("BNBUSDT").Limit(1).Execute()
	if err != nil {
		fmt.Println(err)
		return
	}

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Printf("Response: %s\n", string(dataValue))
}
```
