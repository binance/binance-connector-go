# Key Pair Based Authentication

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	client "github.com/binance/binance-connector-go/clients/alpha"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	GetExchangeInfo()
}

func GetExchangeInfo() {
	privateKeyPath := "/path/to/private_key.pem"

	privateKeyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		panic(err)
	}
	privateKey := string(privateKeyBytes)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.AlphaRestApiProdUrl),
		common.WithApiKey("your-api-key"),
		common.WithPrivateKey(privateKey),
		common.WithPrivateKeyPassphrase("your-passphrase"),
	)
	apiClient := client.NewBinanceAlphaClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.MarketDataAPI.GetExchangeInfo(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```
