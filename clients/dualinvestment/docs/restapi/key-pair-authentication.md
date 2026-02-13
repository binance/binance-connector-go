# Key Pair Based Authentication

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	client "github.com/binance/binance-connector-go/clients/dualinvestment"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	GetDualInvestmentPositions()
}

func GetDualInvestmentPositions() {
	privateKeyPath := "/path/to/private_key.pem"

	privateKeyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		panic(err)
	}
	privateKey := string(privateKeyBytes)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.DualInvestmentRestApiProdUrl),
		common.WithApiKey("your-api-key"),
		common.WithPrivateKey(privateKey),
		common.WithPrivateKeyPassphrase("your-passphrase"),
	)
	apiClient := client.NewBinanceDualInvestmentClient(client.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.GetDualInvestmentPositions(context.Background()).Execute()
	if err != nil {
		fmt.Println(err)
		return
	}

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Printf("Response: %s\n", string(dataValue))
}
```
