# Key Pair Based Authentication

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	client "github.com/binance/binance-connector-go/clients/stocks"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	ExchangeInfo()
}

func ExchangeInfo() {
	privateKeyPath := "/path/to/private_key.pem"

	privateKeyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		panic(err)
	}
	privateKey := string(privateKeyBytes)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.StocksRestApiProdUrl),
		common.WithApiKey("your-api-key"),
		common.WithPrivateKey(privateKey),
		common.WithPrivateKeyPassphrase("your-passphrase"),
	)
	apiClient := client.NewBinanceStocksClient(client.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.ExchangeInfo(context.Background()).Execute()
	if err != nil {
		fmt.Println(err)
		return
	}

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Printf("Response: %s\n", string(dataValue))
}
```
