package main

import (
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/clients/spot/src/websocketapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	OrderListPlaceOtoco()
}

func OrderListPlaceOtoco() {
	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)

	wsClient := client.NewBinanceSpotClient(
		client.WithWebsocketAPI(configuration),
	)
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	responseChan, errorChan, err := wsClient.WebsocketAPI.TradeAPI.OrderListPlaceOtoco().Symbol("BNBUSDT").WorkingType(models.OrderListPlaceOpoWorkingTypeParameterLimit).WorkingSide(models.OrderCancelReplaceSideParameterBuy).WorkingPrice(1.0).WorkingQuantity(1.0).PendingSide(models.OrderCancelReplaceSideParameterBuy).PendingQuantity(1.0).PendingAboveType(models.OrderListPlaceOcoAboveTypeParameterStopLossLimit).ExecuteAsync()
	if err != nil {
		log.Printf("Error executing OrderListPlaceOtoco request: %v\n", err)
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
