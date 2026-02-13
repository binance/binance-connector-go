# Keep-Alive Configuration

```go
package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	OptionAccountInformation()
}

func OptionAccountInformation() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.DerivativesTradingOptionsRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
		common.WithKeepAlive(false),
	)
	apiClient := client.NewBinanceDerivativesTradingOptionsClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.AccountAPI.OptionAccountInformation(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```
