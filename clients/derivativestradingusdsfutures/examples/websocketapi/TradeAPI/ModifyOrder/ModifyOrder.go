package main

import (
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures/src/websocketapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	ModifyOrder()
}

func ModifyOrder() {
	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.DerivativesTradingUsdsFuturesWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)

	wsClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
		client.WithWebsocketAPI(configuration),
	)
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	responseChan, errorChan, err := wsClient.WebsocketAPI.TradeAPI.ModifyOrder().Symbol("symbol_example").Side(models.ModifyOrderSideParameterBuy).Quantity(1.0).Price(1.0).ExecuteAsync()
	if err != nil {
		log.Printf("Error executing ModifyOrder request: %v\n", err)
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
