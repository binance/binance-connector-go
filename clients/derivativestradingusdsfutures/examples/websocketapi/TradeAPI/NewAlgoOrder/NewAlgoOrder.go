package main

import (
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures/src/websocketapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	NewAlgoOrder()
}

func NewAlgoOrder() {
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

	responseChan, errorChan, err := wsClient.WebsocketAPI.TradeAPI.NewAlgoOrder().AlgoType("algoType_example").Symbol("symbol_example").Side(models.ModifyOrderSideParameterBuy).Type("type__example").ExecuteAsync()
	if err != nil {
		log.Printf("Error executing NewAlgoOrder request: %v\n", err)
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
