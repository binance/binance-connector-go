# Key Pair Based Authentication

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	client "github.com/binance/binance-connector-go/clients/pay"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	GetPayTradeHistory()
}

func GetPayTradeHistory() {
	privateKeyPath := "/path/to/private_key.pem"

	privateKeyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		panic(err)
	}
	privateKey := string(privateKeyBytes)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.PayRestApiProdUrl),
		common.WithApiKey("your-api-key"),
		common.WithPrivateKey(privateKey),
		common.WithPrivateKeyPassphrase("your-passphrase"),
	)
	apiClient := client.NewBinancePayClient(client.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.PayAPI.GetPayTradeHistory(context.Background()).Execute()
	if err != nil {
		fmt.Println(err)
		return
	}

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Printf("Response: %s\n", string(dataValue))
}
```
