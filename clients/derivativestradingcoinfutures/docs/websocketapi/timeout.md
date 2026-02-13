# Timeout

```go
package main

import (
	"encoding/json"
	"log"
	"time"

	client "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	PositionInformation()
}

func PositionInformation() {
	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.DerivativesTradingCoinFuturesWebsocketApiProdUrl),
		common.WithWsTimeout(3000*time.Millisecond),
	)

	wsClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
		client.WithWebsocketAPI(configuration),
	)
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	responseChan, errorChan, err := wsClient.WebsocketAPI.TradeAPI.PositionInformation().ExecuteAsync()
	if err != nil {
		log.Printf("Error executing position information request: %v\n", err)
		return
	}

	select {
	case resp := <-responseChan:
		result, _ := json.MarshalIndent(resp.Typed, "", "  ")
		log.Printf("Result: %s\n", result)
	case err := <-errorChan:
		log.Printf("Error: %v\n", err)
	}

	err = wsClient.WebsocketAPI.CloseWebSocketConnection()
	if err != nil {
		log.Printf("Error closing WebSocket connection: %v\n", err)
		return
	}
}
```
