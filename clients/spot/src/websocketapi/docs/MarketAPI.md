# \MarketAPI

All URIs are relative to *http://localhost*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AvgPrice**](MarketAPI.md#AvgPrice) | /avgPrice | Current average price
[**BlockTradesHistorical**](MarketAPI.md#BlockTradesHistorical) | /blockTrades.historical | Historical Block Trades
[**Depth**](MarketAPI.md#Depth) | /depth | Order book
[**Klines**](MarketAPI.md#Klines) | /klines | Klines
[**ReferencePrice**](MarketAPI.md#ReferencePrice) | /referencePrice | Query Reference Price
[**ReferencePriceCalculation**](MarketAPI.md#ReferencePriceCalculation) | /referencePrice.calculation | Query Reference Price Calculation
[**Ticker**](MarketAPI.md#Ticker) | /ticker | Rolling window price change statistics
[**Ticker24hr**](MarketAPI.md#Ticker24hr) | /ticker.24hr | 24hr ticker price change statistics
[**TickerBook**](MarketAPI.md#TickerBook) | /ticker.book | Symbol order book ticker
[**TickerPrice**](MarketAPI.md#TickerPrice) | /ticker.price | Symbol price ticker
[**TickerTradingDay**](MarketAPI.md#TickerTradingDay) | /ticker.tradingDay | Trading Day Ticker
[**TradesAggregate**](MarketAPI.md#TradesAggregate) | /trades.aggregate | Aggregate trades
[**TradesHistorical**](MarketAPI.md#TradesHistorical) | /trades.historical | Historical trades
[**TradesRecent**](MarketAPI.md#TradesRecent) | /trades.recent | Recent trades
[**UiKlines**](MarketAPI.md#UiKlines) | /uiKlines | UI Klines


## AvgPrice

> AvgPriceResponse AvgPrice().Symbol(symbol).Id(id).Execute()

Current average price


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.MarketAPI.AvgPrice().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketAPI.AvgPrice``: %v\n", err)
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
 **id** | **string** | Client-generated request identifier. | 

### Return type

[**AvgPriceResponse**](AvgPriceResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## BlockTradesHistorical

> BlockTradesHistoricalResponse BlockTradesHistorical().Symbol(symbol).FromId(fromId).Id(id).Limit(limit).Execute()

Historical Block Trades


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BNBBTC" // string | 
	fromId := int64(582) // int64 | Block trade ID to fetch from
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	limit := int64(500) // int64 | Default: 500; Maximum: 1000 (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.MarketAPI.BlockTradesHistorical().Symbol(symbol).FromId(fromId).Id(id).Limit(limit).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketAPI.BlockTradesHistorical``: %v\n", err)
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
 **fromId** | **int64** | Block trade ID to fetch from | 
 **id** | **string** | Client-generated request identifier. | 
 **limit** | **int64** | Default: 500; Maximum: 1000 | 

### Return type

[**BlockTradesHistoricalResponse**](BlockTradesHistoricalResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## Depth

> DepthResponse Depth().Symbol(symbol).Id(id).Limit(limit).SymbolStatus(symbolStatus).Execute()

Order book


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	limit := int32(1) // int32 |  (optional)
	symbolStatus := models.ExchangeInfoSymbolStatusParameterTrading // ExchangeInfoSymbolStatusParameter | Filters for symbols that have this `tradingStatus`. A status mismatch returns error `-1220 SYMBOL_DOES_NOT_MATCH_STATUS`. (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.MarketAPI.Depth().Symbol(symbol).Id(id).Limit(limit).SymbolStatus(symbolStatus).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketAPI.Depth``: %v\n", err)
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
 **id** | **string** | Client-generated request identifier. | 
 **limit** | **int32** |  | 
 **symbolStatus** | [**ExchangeInfoSymbolStatusParameter**](ExchangeInfoSymbolStatusParameter.md) | Filters for symbols that have this &#x60;tradingStatus&#x60;. A status mismatch returns error &#x60;-1220 SYMBOL_DOES_NOT_MATCH_STATUS&#x60;. | 

### Return type

[**DepthResponse**](DepthResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## Klines

> KlinesResponse Klines().Symbol(symbol).Interval(interval).Id(id).StartTime(startTime).EndTime(endTime).TimeZone(timeZone).Limit(limit).Execute()

Klines


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	interval := models.KlinesIntervalParameterInterval1s // KlinesIntervalParameter | 
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	startTime := int64(1735693200000) // int64 |  (optional)
	endTime := int64(1735693200000) // int64 |  (optional)
	timeZone := "0" // string | Default: 0 (UTC) (optional)
	limit := int32(1) // int32 |  (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.MarketAPI.Klines().Symbol(symbol).Interval(interval).Id(id).StartTime(startTime).EndTime(endTime).TimeZone(timeZone).Limit(limit).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketAPI.Klines``: %v\n", err)
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
 **interval** | [**KlinesIntervalParameter**](KlinesIntervalParameter.md) |  | 
 **id** | **string** | Client-generated request identifier. | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **timeZone** | **string** | Default: 0 (UTC) | 
 **limit** | **int32** |  | 

### Return type

[**KlinesResponse**](KlinesResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## ReferencePrice

> ReferencePriceResponse ReferencePrice().Symbol(symbol).Id(id).Execute()

Query Reference Price


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BAZUSD" // string | 
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.MarketAPI.ReferencePrice().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketAPI.ReferencePrice``: %v\n", err)
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
 **id** | **string** | Client-generated request identifier. | 

### Return type

[**ReferencePriceResponse**](ReferencePriceResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## ReferencePriceCalculation

> ReferencePriceCalculationResponse ReferencePriceCalculation().Symbol(symbol).Id(id).SymbolStatus(symbolStatus).Execute()

Query Reference Price Calculation


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BAZUSD" // string | 
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	symbolStatus := models.ReferencePriceCalculationSymbolStatusParameterTrading // ReferencePriceCalculationSymbolStatusParameter |  (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.MarketAPI.ReferencePriceCalculation().Symbol(symbol).Id(id).SymbolStatus(symbolStatus).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketAPI.ReferencePriceCalculation``: %v\n", err)
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
 **id** | **string** | Client-generated request identifier. | 
 **symbolStatus** | [**ReferencePriceCalculationSymbolStatusParameter**](ReferencePriceCalculationSymbolStatusParameter.md) |  | 

### Return type

[**ReferencePriceCalculationResponse**](ReferencePriceCalculationResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## Ticker

> TickerResponse Ticker().Id(id).Symbol(symbol).Symbols(symbols).Type(type_).WindowSize(windowSize).SymbolStatus(symbolStatus).Execute()

Rolling window price change statistics


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	symbol := "BNBUSDT" // string | Query ticker of a single symbol (optional)
	symbols := []string{"BTCUSDT"} // []string | Query ticker for multiple symbols (optional)
	type_ := models.TickerTypeParameterFull // TickerTypeParameter | Ticker type. Supported values: FULL (default) or MINI (optional)
	windowSize := models.TickerWindowSizeParameterWindowSize1m // TickerWindowSizeParameter | Defaults to 1d if no parameter provided. (optional)
	symbolStatus := models.ExchangeInfoSymbolStatusParameterTrading // ExchangeInfoSymbolStatusParameter | Filters for symbols that have this `tradingStatus`. For a single symbol, a status mismatch returns error `-1220 SYMBOL_DOES_NOT_MATCH_STATUS`. For multiple or all symbols, non-matching ones are simply excluded from the response. Valid values: `TRADING`, `HALT`, `BREAK` (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.MarketAPI.Ticker().Id(id).Symbol(symbol).Symbols(symbols).Type(type_).WindowSize(windowSize).SymbolStatus(symbolStatus).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketAPI.Ticker``: %v\n", err)
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
 **id** | **string** | Client-generated request identifier. | 
 **symbol** | **string** | Query ticker of a single symbol | 
 **symbols** | **[]string** | Query ticker for multiple symbols | 
 **type_** | [**TickerTypeParameter**](TickerTypeParameter.md) | Ticker type. Supported values: FULL (default) or MINI | 
 **windowSize** | [**TickerWindowSizeParameter**](TickerWindowSizeParameter.md) | Defaults to 1d if no parameter provided. | 
 **symbolStatus** | [**ExchangeInfoSymbolStatusParameter**](ExchangeInfoSymbolStatusParameter.md) | Filters for symbols that have this &#x60;tradingStatus&#x60;. For a single symbol, a status mismatch returns error &#x60;-1220 SYMBOL_DOES_NOT_MATCH_STATUS&#x60;. For multiple or all symbols, non-matching ones are simply excluded from the response. Valid values: &#x60;TRADING&#x60;, &#x60;HALT&#x60;, &#x60;BREAK&#x60; | 

### Return type

[**TickerResponse**](TickerResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## Ticker24hr

> Ticker24hrResponse Ticker24hr().Id(id).Symbol(symbol).Symbols(symbols).Type(type_).SymbolStatus(symbolStatus).Execute()

24hr ticker price change statistics


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	symbol := "BNBUSDT" // string |  (optional)
	symbols := []string{"BTCUSDT"} // []string |  (optional)
	type_ := models.TickerTypeParameterFull // TickerTypeParameter | Ticker type. Supported values: FULL (default) or MINI (optional)
	symbolStatus := models.ExchangeInfoSymbolStatusParameterTrading // ExchangeInfoSymbolStatusParameter | Filters for symbols that have this `tradingStatus`. For a single symbol, a status mismatch returns error `-1220 SYMBOL_DOES_NOT_MATCH_STATUS`. For multiple or all symbols, non-matching ones are simply excluded from the response. Valid values: `TRADING`, `HALT`, `BREAK` (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.MarketAPI.Ticker24hr().Id(id).Symbol(symbol).Symbols(symbols).Type(type_).SymbolStatus(symbolStatus).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketAPI.Ticker24hr``: %v\n", err)
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
 **id** | **string** | Client-generated request identifier. | 
 **symbol** | **string** |  | 
 **symbols** | **[]string** |  | 
 **type_** | [**TickerTypeParameter**](TickerTypeParameter.md) | Ticker type. Supported values: FULL (default) or MINI | 
 **symbolStatus** | [**ExchangeInfoSymbolStatusParameter**](ExchangeInfoSymbolStatusParameter.md) | Filters for symbols that have this &#x60;tradingStatus&#x60;. For a single symbol, a status mismatch returns error &#x60;-1220 SYMBOL_DOES_NOT_MATCH_STATUS&#x60;. For multiple or all symbols, non-matching ones are simply excluded from the response. Valid values: &#x60;TRADING&#x60;, &#x60;HALT&#x60;, &#x60;BREAK&#x60; | 

### Return type

[**Ticker24hrResponse**](Ticker24hrResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## TickerBook

> TickerBookResponse TickerBook().Id(id).Symbol(symbol).Symbols(symbols).SymbolStatus(symbolStatus).Execute()

Symbol order book ticker


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	symbol := "BNBUSDT" // string | Query ticker for a single symbol (optional)
	symbols := []string{"BTCUSDT"} // []string | Query ticker for multiple symbols (optional)
	symbolStatus := models.ExchangeInfoSymbolStatusParameterTrading // ExchangeInfoSymbolStatusParameter | Filters for symbols that have this `tradingStatus`. For a single symbol, a status mismatch returns error `-1220 SYMBOL_DOES_NOT_MATCH_STATUS`. For multiple or all symbols, non-matching ones are simply excluded from the response. Valid values: `TRADING`, `HALT`, `BREAK` (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.MarketAPI.TickerBook().Id(id).Symbol(symbol).Symbols(symbols).SymbolStatus(symbolStatus).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketAPI.TickerBook``: %v\n", err)
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
 **id** | **string** | Client-generated request identifier. | 
 **symbol** | **string** | Query ticker for a single symbol | 
 **symbols** | **[]string** | Query ticker for multiple symbols | 
 **symbolStatus** | [**ExchangeInfoSymbolStatusParameter**](ExchangeInfoSymbolStatusParameter.md) | Filters for symbols that have this &#x60;tradingStatus&#x60;. For a single symbol, a status mismatch returns error &#x60;-1220 SYMBOL_DOES_NOT_MATCH_STATUS&#x60;. For multiple or all symbols, non-matching ones are simply excluded from the response. Valid values: &#x60;TRADING&#x60;, &#x60;HALT&#x60;, &#x60;BREAK&#x60; | 

### Return type

[**TickerBookResponse**](TickerBookResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## TickerPrice

> TickerPriceResponse TickerPrice().Id(id).Symbol(symbol).Symbols(symbols).SymbolStatus(symbolStatus).Execute()

Symbol price ticker


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	symbol := "BNBUSDT" // string | Query price for a single symbol (optional)
	symbols := []string{"BTCUSDT"} // []string | Query price for multiple symbols (optional)
	symbolStatus := models.ExchangeInfoSymbolStatusParameterTrading // ExchangeInfoSymbolStatusParameter | Filters for symbols that have this `tradingStatus`. For a single symbol, a status mismatch returns error `-1220 SYMBOL_DOES_NOT_MATCH_STATUS`. For multiple or all symbols, non-matching ones are simply excluded from the response. Valid values: `TRADING`, `HALT`, `BREAK` (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.MarketAPI.TickerPrice().Id(id).Symbol(symbol).Symbols(symbols).SymbolStatus(symbolStatus).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketAPI.TickerPrice``: %v\n", err)
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
 **id** | **string** | Client-generated request identifier. | 
 **symbol** | **string** | Query price for a single symbol | 
 **symbols** | **[]string** | Query price for multiple symbols | 
 **symbolStatus** | [**ExchangeInfoSymbolStatusParameter**](ExchangeInfoSymbolStatusParameter.md) | Filters for symbols that have this &#x60;tradingStatus&#x60;. For a single symbol, a status mismatch returns error &#x60;-1220 SYMBOL_DOES_NOT_MATCH_STATUS&#x60;. For multiple or all symbols, non-matching ones are simply excluded from the response. Valid values: &#x60;TRADING&#x60;, &#x60;HALT&#x60;, &#x60;BREAK&#x60; | 

### Return type

[**TickerPriceResponse**](TickerPriceResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## TickerTradingDay

> TickerTradingDayResponse TickerTradingDay().Id(id).Symbol(symbol).Symbols(symbols).TimeZone(timeZone).Type(type_).SymbolStatus(symbolStatus).Execute()

Trading Day Ticker


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	symbol := "BNBUSDT" // string |  (optional)
	symbols := []string{"BTCUSDT"} // []string |  (optional)
	timeZone := "0" // string | Default: 0 (UTC) (optional)
	type_ := models.TickerTypeParameterFull // TickerTypeParameter | Ticker type. Supported values: FULL (default) or MINI (optional)
	symbolStatus := models.ExchangeInfoSymbolStatusParameterTrading // ExchangeInfoSymbolStatusParameter | Filters for symbols that have this `tradingStatus`. For a single symbol, a status mismatch returns error `-1220 SYMBOL_DOES_NOT_MATCH_STATUS`. For multiple or all symbols, non-matching ones are simply excluded from the response. Valid values: `TRADING`, `HALT`, `BREAK` (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.MarketAPI.TickerTradingDay().Id(id).Symbol(symbol).Symbols(symbols).TimeZone(timeZone).Type(type_).SymbolStatus(symbolStatus).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketAPI.TickerTradingDay``: %v\n", err)
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
 **id** | **string** | Client-generated request identifier. | 
 **symbol** | **string** |  | 
 **symbols** | **[]string** |  | 
 **timeZone** | **string** | Default: 0 (UTC) | 
 **type_** | [**TickerTypeParameter**](TickerTypeParameter.md) | Ticker type. Supported values: FULL (default) or MINI | 
 **symbolStatus** | [**ExchangeInfoSymbolStatusParameter**](ExchangeInfoSymbolStatusParameter.md) | Filters for symbols that have this &#x60;tradingStatus&#x60;. For a single symbol, a status mismatch returns error &#x60;-1220 SYMBOL_DOES_NOT_MATCH_STATUS&#x60;. For multiple or all symbols, non-matching ones are simply excluded from the response. Valid values: &#x60;TRADING&#x60;, &#x60;HALT&#x60;, &#x60;BREAK&#x60; | 

### Return type

[**TickerTradingDayResponse**](TickerTradingDayResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## TradesAggregate

> TradesAggregateResponse TradesAggregate().Symbol(symbol).Id(id).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Aggregate trades


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	fromId := int64(1) // int64 | Aggregate trade ID to begin at (optional)
	startTime := int64(1735693200000) // int64 | Timestamp in ms to get aggregate trades from INCLUSIVE. (optional)
	endTime := int64(1735693200000) // int64 | Timestamp in ms to get aggregate trades until INCLUSIVE. (optional)
	limit := int32(1) // int32 |  (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.MarketAPI.TradesAggregate().Symbol(symbol).Id(id).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketAPI.TradesAggregate``: %v\n", err)
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
 **id** | **string** | Client-generated request identifier. | 
 **fromId** | **int64** | Aggregate trade ID to begin at | 
 **startTime** | **int64** | Timestamp in ms to get aggregate trades from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get aggregate trades until INCLUSIVE. | 
 **limit** | **int32** |  | 

### Return type

[**TradesAggregateResponse**](TradesAggregateResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## TradesHistorical

> TradesHistoricalResponse TradesHistorical().Symbol(symbol).Id(id).FromId(fromId).Limit(limit).Execute()

Historical trades


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	fromId := int64(1) // int64 | Trade ID to begin at (optional)
	limit := int32(1) // int32 |  (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.MarketAPI.TradesHistorical().Symbol(symbol).Id(id).FromId(fromId).Limit(limit).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketAPI.TradesHistorical``: %v\n", err)
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
 **id** | **string** | Client-generated request identifier. | 
 **fromId** | **int64** | Trade ID to begin at | 
 **limit** | **int32** |  | 

### Return type

[**TradesHistoricalResponse**](TradesHistoricalResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## TradesRecent

> TradesRecentResponse TradesRecent().Symbol(symbol).Id(id).Limit(limit).Execute()

Recent trades


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	limit := int32(1) // int32 |  (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.MarketAPI.TradesRecent().Symbol(symbol).Id(id).Limit(limit).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketAPI.TradesRecent``: %v\n", err)
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
 **id** | **string** | Client-generated request identifier. | 
 **limit** | **int32** |  | 

### Return type

[**TradesRecentResponse**](TradesRecentResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## UiKlines

> UiKlinesResponse UiKlines().Symbol(symbol).Interval(interval).Id(id).StartTime(startTime).EndTime(endTime).TimeZone(timeZone).Limit(limit).Execute()

UI Klines


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	interval := models.KlinesIntervalParameterInterval1s // KlinesIntervalParameter | 
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	startTime := int64(1735693200000) // int64 |  (optional)
	endTime := int64(1735693200000) // int64 |  (optional)
	timeZone := "0" // string | Default: 0 (UTC) (optional)
	limit := int32(1) // int32 |  (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.MarketAPI.UiKlines().Symbol(symbol).Interval(interval).Id(id).StartTime(startTime).EndTime(endTime).TimeZone(timeZone).Limit(limit).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketAPI.UiKlines``: %v\n", err)
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
 **interval** | [**KlinesIntervalParameter**](KlinesIntervalParameter.md) |  | 
 **id** | **string** | Client-generated request identifier. | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **timeZone** | **string** | Default: 0 (UTC) | 
 **limit** | **int32** |  | 

### Return type

[**UiKlinesResponse**](UiKlinesResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)

