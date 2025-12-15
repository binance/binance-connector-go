# Keep-Alive Configuration

```go
package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/fiat"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	GetFiatDepositWithdrawHistory()
}

func GetFiatDepositWithdrawHistory() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.FiatRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
		common.WithKeepAlive(false),
	)
	apiClient := client.NewBinanceFiatClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.FiatAPI.GetFiatDepositWithdrawHistory(context.Background()).TransactionType("transactionType_example").Execute()
	if err != nil {
		log.Println(err)
		return
	}

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```
