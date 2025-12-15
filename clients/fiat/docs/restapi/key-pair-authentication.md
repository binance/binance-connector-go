# Key Pair Based Authentication

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	client "github.com/binance/binance-connector-go/clients/fiat"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	GetFiatDepositWithdrawHistory()
}

func GetFiatDepositWithdrawHistory() {
	privateKeyPath := "/path/to/private_key.pem"

	privateKeyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		panic(err)
	}
	privateKey := string(privateKeyBytes)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.FiatRestApiProdUrl),
		common.WithApiKey("your-api-key"),
		common.WithPrivateKey(privateKey),
		common.WithPrivateKeyPassphrase("your-passphrase"),
	)
	apiClient := client.NewBinanceFiatClient(client.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FiatAPI.GetFiatDepositWithdrawHistory(context.Background()).TransactionType("transactionType_example").Execute()
	if err != nil {
		fmt.Println(err)
		return
	}

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Printf("Response: %s\n", string(dataValue))
}
```
