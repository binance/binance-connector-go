# Proxy Configuration

```go
package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	AccountInfo()
}

func AccountInfo() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.WalletRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
		common.WithProxy(common.ProxyConfig{
			Host:     "127.0.0.1",
			Port:     8080,
			Protocol: "http",
		}),
	)
	apiClient := client.NewBinanceWalletClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.AccountAPI.AccountInfo(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```
