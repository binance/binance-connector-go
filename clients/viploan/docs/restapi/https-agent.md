# HTTPS Agent Configuration

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	client "github.com/binance/binance-connector-go/clients/viploan"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	GetCollateralAssetData()
}

type LoggingTransport struct {
	Base http.RoundTripper
}

func (t *LoggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	log.Printf("→ %s %s", req.Method, req.URL)
	resp, err := t.Base.RoundTrip(req)
	if err != nil {
		log.Printf("✗ %v", err)
	} else {
		log.Printf("← %d %s", resp.StatusCode, req.URL)
	}
	return resp, err
}

func GetCollateralAssetData() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.VipLoanRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
		common.WithHTTPSAgent(&LoggingTransport{Base: http.DefaultTransport}),
	)
	apiClient := client.NewBinanceVipLoanClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.MarketDataAPI.GetCollateralAssetData(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```
