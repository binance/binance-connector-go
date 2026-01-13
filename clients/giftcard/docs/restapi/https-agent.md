# HTTPS Agent Configuration

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	client "github.com/binance/binance-connector-go/clients/giftcard"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	CreateASingleTokenGiftCard()
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

func CreateASingleTokenGiftCard() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.GiftCardRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
		common.WithHTTPSAgent(&LoggingTransport{Base: http.DefaultTransport}),
	)
	apiClient := client.NewBinanceGiftCardClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.MarketDataAPI.CreateASingleTokenGiftCard(context.Background()).Token("token_example").Amount(1.0).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```
