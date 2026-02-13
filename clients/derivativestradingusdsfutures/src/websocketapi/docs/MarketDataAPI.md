# \MarketDataAPI

All URIs are relative to *http://localhost*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**OrderBook**](MarketDataAPI.md#OrderBook) | /depth | Order Book
[**SymbolOrderBookTicker**](MarketDataAPI.md#SymbolOrderBookTicker) | /ticker.book | Symbol Order Book Ticker
[**SymbolPriceTicker**](MarketDataAPI.md#SymbolPriceTicker) | /ticker.price | Symbol Price Ticker


## OrderBook

> OrderBookResponse OrderBook().Symbol(symbol).Id(id).Limit(limit).Execute()

Order Book


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	limit := int64(10) // int64 | Default 500; Valid limits:[5, 10, 20, 50, 100, 500, 1000] (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.MarketDataAPI.OrderBook().Symbol(symbol).Id(id).Limit(limit).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.OrderBook``: %v\n", err)
		return
	}

	result, _ := json.MarshalIndent(resp.Typed, "", "  ")
	log.Printf("Result: %s\n", result)

	err = wsClient.WebsocketAPI.CloseWebSocketConnection()
	if err != nil {
		log.Printf("Error closing WebSocket connection: %v\n", err)
		return
	}
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | 
 **id** | **string** | Unique WebSocket request ID. | 
 **limit** | **int64** | Default 500; Valid limits:[5, 10, 20, 50, 100, 500, 1000] | 

### Return type

[**OrderBookResponse**](OrderBookResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## SymbolOrderBookTicker

> SymbolOrderBookTickerResponse SymbolOrderBookTicker().Id(id).Symbol(symbol).Execute()

Symbol Order Book Ticker


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	symbol := "symbol_example" // string |  (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.MarketDataAPI.SymbolOrderBookTicker().Id(id).Symbol(symbol).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.SymbolOrderBookTicker``: %v\n", err)
		return
	}

	result, _ := json.MarshalIndent(resp.Typed, "", "  ")
	log.Printf("Result: %s\n", result)

	err = wsClient.WebsocketAPI.CloseWebSocketConnection()
	if err != nil {
		log.Printf("Error closing WebSocket connection: %v\n", err)
		return
	}
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Unique WebSocket request ID. | 
 **symbol** | **string** |  | 

### Return type

[**SymbolOrderBookTickerResponse**](SymbolOrderBookTickerResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## SymbolPriceTicker

> SymbolPriceTickerResponse SymbolPriceTicker().Id(id).Symbol(symbol).Execute()

Symbol Price Ticker


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	symbol := "symbol_example" // string |  (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.MarketDataAPI.SymbolPriceTicker().Id(id).Symbol(symbol).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.SymbolPriceTicker``: %v\n", err)
		return
	}

	result, _ := json.MarshalIndent(resp.Typed, "", "  ")
	log.Printf("Result: %s\n", result)

	err = wsClient.WebsocketAPI.CloseWebSocketConnection()
	if err != nil {
		log.Printf("Error closing WebSocket connection: %v\n", err)
		return
	}
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Unique WebSocket request ID. | 
 **symbol** | **string** |  | 

### Return type

[**SymbolPriceTickerResponse**](SymbolPriceTickerResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)

