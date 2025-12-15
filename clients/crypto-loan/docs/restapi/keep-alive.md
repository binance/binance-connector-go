# Keep-Alive Configuration

```go
package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/cryptoloan"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	GetFlexibleLoanBorrowHistory()
}

func GetFlexibleLoanBorrowHistory() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.CryptoLoanRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
		common.WithKeepAlive(false),
	)
	apiClient := client.NewBinanceCryptoLoanClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.FlexibleRateAPI.GetFlexibleLoanBorrowHistory(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```
