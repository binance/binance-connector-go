# \MarketAPI

All URIs are relative to *http://localhost*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AvgPrice**](MarketAPI.md#AvgPrice) | /avgPrice | WebSocket Current average price
[**Depth**](MarketAPI.md#Depth) | /depth | WebSocket Order book
[**Klines**](MarketAPI.md#Klines) | /klines | WebSocket Klines
[**Ticker**](MarketAPI.md#Ticker) | /ticker | WebSocket Rolling window price change statistics
[**Ticker24hr**](MarketAPI.md#Ticker24hr) | /ticker.24hr | WebSocket 24hr ticker price change statistics
[**TickerBook**](MarketAPI.md#TickerBook) | /ticker.book | WebSocket Symbol order book ticker
[**TickerPrice**](MarketAPI.md#TickerPrice) | /ticker.price | WebSocket Symbol price ticker
[**TickerTradingDay**](MarketAPI.md#TickerTradingDay) | /ticker.tradingDay | WebSocket Trading Day Ticker
[**TradesAggregate**](MarketAPI.md#TradesAggregate) | /trades.aggregate | WebSocket Aggregate trades
[**TradesHistorical**](MarketAPI.md#TradesHistorical) | /trades.historical | WebSocket Historical trades
[**TradesRecent**](MarketAPI.md#TradesRecent) | /trades.recent | WebSocket Recent trades
[**UiKlines**](MarketAPI.md#UiKlines) | /uiKlines | WebSocket UI Klines


## AvgPrice

> AvgPriceResponse AvgPrice().Symbol(symbol).Id(id).Execute()

WebSocket Current average price


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

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
 **id** | **string** | Unique WebSocket request ID. | 

### Return type

[**AvgPriceResponse**](AvgPriceResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## Depth

> DepthResponse Depth().Symbol(symbol).Id(id).Limit(limit).SymbolStatus(symbolStatus).Execute()

WebSocket Order book


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	limit := int32(100) // int32 | Default: 100; Maximum: 5000 (optional)
	symbolStatus := models.ExchangeInfoSymbolStatusParameterTrading // ExchangeInfoSymbolStatusParameter |  (optional)

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
 **id** | **string** | Unique WebSocket request ID. | 
 **limit** | **int32** | Default: 100; Maximum: 5000 | 
 **symbolStatus** | [**ExchangeInfoSymbolStatusParameter**](ExchangeInfoSymbolStatusParameter.md) |  | 

### Return type

[**DepthResponse**](DepthResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## Klines

> KlinesResponse Klines().Symbol(symbol).Interval(interval).Id(id).StartTime(startTime).EndTime(endTime).TimeZone(timeZone).Limit(limit).Execute()

WebSocket Klines


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	interval := models.KlinesIntervalParameterInterval1s // KlinesIntervalParameter | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	startTime := int64(1735693200000) // int64 |  (optional)
	endTime := int64(1735693200000) // int64 |  (optional)
	timeZone := "timeZone_example" // string | Default: 0 (UTC) (optional)
	limit := int32(100) // int32 | Default: 100; Maximum: 5000 (optional)

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
 **id** | **string** | Unique WebSocket request ID. | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **timeZone** | **string** | Default: 0 (UTC) | 
 **limit** | **int32** | Default: 100; Maximum: 5000 | 

### Return type

[**KlinesResponse**](KlinesResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## Ticker

> TickerResponse Ticker().Id(id).Symbol(symbol).Symbols(symbols).Type(type_).WindowSize(windowSize).SymbolStatus(symbolStatus).Execute()

WebSocket Rolling window price change statistics


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	symbol := "BNBUSDT" // string | Describe a single symbol (optional)
	symbols := []string{"Inner_example"} // []string | List of symbols to query (optional)
	type_ := models.TickerTypeParameterFull // TickerTypeParameter |  (optional)
	windowSize := models.TickerWindowSizeParameterWindowSize1m // TickerWindowSizeParameter |  (optional)
	symbolStatus := models.ExchangeInfoSymbolStatusParameterTrading // ExchangeInfoSymbolStatusParameter |  (optional)

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
 **id** | **string** | Unique WebSocket request ID. | 
 **symbol** | **string** | Describe a single symbol | 
 **symbols** | **[]string** | List of symbols to query | 
 **type_** | [**TickerTypeParameter**](TickerTypeParameter.md) |  | 
 **windowSize** | [**TickerWindowSizeParameter**](TickerWindowSizeParameter.md) |  | 
 **symbolStatus** | [**ExchangeInfoSymbolStatusParameter**](ExchangeInfoSymbolStatusParameter.md) |  | 

### Return type

[**TickerResponse**](TickerResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## Ticker24hr

> Ticker24hrResponse Ticker24hr().Id(id).Symbol(symbol).Symbols(symbols).Type(type_).SymbolStatus(symbolStatus).Execute()

WebSocket 24hr ticker price change statistics


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	symbol := "BNBUSDT" // string | Describe a single symbol (optional)
	symbols := []string{"Inner_example"} // []string | List of symbols to query (optional)
	type_ := models.TickerTypeParameterFull // TickerTypeParameter |  (optional)
	symbolStatus := models.ExchangeInfoSymbolStatusParameterTrading // ExchangeInfoSymbolStatusParameter |  (optional)

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
 **id** | **string** | Unique WebSocket request ID. | 
 **symbol** | **string** | Describe a single symbol | 
 **symbols** | **[]string** | List of symbols to query | 
 **type_** | [**TickerTypeParameter**](TickerTypeParameter.md) |  | 
 **symbolStatus** | [**ExchangeInfoSymbolStatusParameter**](ExchangeInfoSymbolStatusParameter.md) |  | 

### Return type

[**Ticker24hrResponse**](Ticker24hrResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## TickerBook

> TickerBookResponse TickerBook().Id(id).Symbol(symbol).Symbols(symbols).SymbolStatus(symbolStatus).Execute()

WebSocket Symbol order book ticker


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	symbol := "BNBUSDT" // string | Describe a single symbol (optional)
	symbols := []string{"Inner_example"} // []string | List of symbols to query (optional)
	symbolStatus := models.ExchangeInfoSymbolStatusParameterTrading // ExchangeInfoSymbolStatusParameter |  (optional)

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
 **id** | **string** | Unique WebSocket request ID. | 
 **symbol** | **string** | Describe a single symbol | 
 **symbols** | **[]string** | List of symbols to query | 
 **symbolStatus** | [**ExchangeInfoSymbolStatusParameter**](ExchangeInfoSymbolStatusParameter.md) |  | 

### Return type

[**TickerBookResponse**](TickerBookResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## TickerPrice

> TickerPriceResponse TickerPrice().Id(id).Symbol(symbol).Symbols(symbols).SymbolStatus(symbolStatus).Execute()

WebSocket Symbol price ticker


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	symbol := "BNBUSDT" // string | Describe a single symbol (optional)
	symbols := []string{"Inner_example"} // []string | List of symbols to query (optional)
	symbolStatus := models.ExchangeInfoSymbolStatusParameterTrading // ExchangeInfoSymbolStatusParameter |  (optional)

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
 **id** | **string** | Unique WebSocket request ID. | 
 **symbol** | **string** | Describe a single symbol | 
 **symbols** | **[]string** | List of symbols to query | 
 **symbolStatus** | [**ExchangeInfoSymbolStatusParameter**](ExchangeInfoSymbolStatusParameter.md) |  | 

### Return type

[**TickerPriceResponse**](TickerPriceResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## TickerTradingDay

> TickerTradingDayResponse TickerTradingDay().Id(id).Symbol(symbol).Symbols(symbols).TimeZone(timeZone).Type(type_).SymbolStatus(symbolStatus).Execute()

WebSocket Trading Day Ticker


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	symbol := "BNBUSDT" // string | Describe a single symbol (optional)
	symbols := []string{"Inner_example"} // []string | List of symbols to query (optional)
	timeZone := "timeZone_example" // string | Default: 0 (UTC) (optional)
	type_ := models.TickerTypeParameterFull // TickerTypeParameter |  (optional)
	symbolStatus := models.ExchangeInfoSymbolStatusParameterTrading // ExchangeInfoSymbolStatusParameter |  (optional)

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
 **id** | **string** | Unique WebSocket request ID. | 
 **symbol** | **string** | Describe a single symbol | 
 **symbols** | **[]string** | List of symbols to query | 
 **timeZone** | **string** | Default: 0 (UTC) | 
 **type_** | [**TickerTypeParameter**](TickerTypeParameter.md) |  | 
 **symbolStatus** | [**ExchangeInfoSymbolStatusParameter**](ExchangeInfoSymbolStatusParameter.md) |  | 

### Return type

[**TickerTradingDayResponse**](TickerTradingDayResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## TradesAggregate

> TradesAggregateResponse TradesAggregate().Symbol(symbol).Id(id).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

WebSocket Aggregate trades


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	fromId := int32(1) // int32 | Aggregate trade ID to begin at (optional)
	startTime := int64(1735693200000) // int64 |  (optional)
	endTime := int64(1735693200000) // int64 |  (optional)
	limit := int32(100) // int32 | Default: 100; Maximum: 5000 (optional)

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
 **id** | **string** | Unique WebSocket request ID. | 
 **fromId** | **int32** | Aggregate trade ID to begin at | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default: 100; Maximum: 5000 | 

### Return type

[**TradesAggregateResponse**](TradesAggregateResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## TradesHistorical

> TradesHistoricalResponse TradesHistorical().Symbol(symbol).Id(id).FromId(fromId).Limit(limit).Execute()

WebSocket Historical trades


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	fromId := int32(1) // int32 | Aggregate trade ID to begin at (optional)
	limit := int32(100) // int32 | Default: 100; Maximum: 5000 (optional)

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
 **id** | **string** | Unique WebSocket request ID. | 
 **fromId** | **int32** | Aggregate trade ID to begin at | 
 **limit** | **int32** | Default: 100; Maximum: 5000 | 

### Return type

[**TradesHistoricalResponse**](TradesHistoricalResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## TradesRecent

> TradesRecentResponse TradesRecent().Symbol(symbol).Id(id).Limit(limit).Execute()

WebSocket Recent trades


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	limit := int32(100) // int32 | Default: 100; Maximum: 5000 (optional)

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
 **id** | **string** | Unique WebSocket request ID. | 
 **limit** | **int32** | Default: 100; Maximum: 5000 | 

### Return type

[**TradesRecentResponse**](TradesRecentResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## UiKlines

> UiKlinesResponse UiKlines().Symbol(symbol).Interval(interval).Id(id).StartTime(startTime).EndTime(endTime).TimeZone(timeZone).Limit(limit).Execute()

WebSocket UI Klines


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	interval := models.KlinesIntervalParameterInterval1s // KlinesIntervalParameter | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	startTime := int64(1735693200000) // int64 |  (optional)
	endTime := int64(1735693200000) // int64 |  (optional)
	timeZone := "timeZone_example" // string | Default: 0 (UTC) (optional)
	limit := int32(100) // int32 | Default: 100; Maximum: 5000 (optional)

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
 **id** | **string** | Unique WebSocket request ID. | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **timeZone** | **string** | Default: 0 (UTC) | 
 **limit** | **int32** | Default: 100; Maximum: 5000 | 

### Return type

[**UiKlinesResponse**](UiKlinesResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)

