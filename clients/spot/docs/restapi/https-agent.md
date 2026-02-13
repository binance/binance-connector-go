# HTTPS Agent Configuration

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	client "github.com/binance/binance-connector-go/clients/spot/src"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	ExchangeInfo()
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

func ExchangeInfo() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithHTTPSAgent(&LoggingTransport{Base: http.DefaultTransport}),
	)
	apiClient := client.NewBinanceSpotClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.GeneralAPI.ExchangeInfo(context.Background()).Symbol("BTCUSDT").Execute()
	if err != nil {
		log.Println(err)
		return
	}

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```
