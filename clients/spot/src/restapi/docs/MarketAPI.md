# \MarketAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AggTrades**](MarketAPI.md#AggTrades) | **Get** /api/v3/aggTrades | Compressed/Aggregate trades list
[**AvgPrice**](MarketAPI.md#AvgPrice) | **Get** /api/v3/avgPrice | Current average price
[**Depth**](MarketAPI.md#Depth) | **Get** /api/v3/depth | Order book
[**GetTrades**](MarketAPI.md#GetTrades) | **Get** /api/v3/trades | Recent trades list
[**HistoricalTrades**](MarketAPI.md#HistoricalTrades) | **Get** /api/v3/historicalTrades | Old trade lookup
[**Klines**](MarketAPI.md#Klines) | **Get** /api/v3/klines | Kline/Candlestick data
[**Ticker**](MarketAPI.md#Ticker) | **Get** /api/v3/ticker | Rolling window price change statistics
[**Ticker24hr**](MarketAPI.md#Ticker24hr) | **Get** /api/v3/ticker/24hr | 24hr ticker price change statistics
[**TickerBookTicker**](MarketAPI.md#TickerBookTicker) | **Get** /api/v3/ticker/bookTicker | Symbol order book ticker
[**TickerPrice**](MarketAPI.md#TickerPrice) | **Get** /api/v3/ticker/price | Symbol price ticker
[**TickerTradingDay**](MarketAPI.md#TickerTradingDay) | **Get** /api/v3/ticker/tradingDay | Trading Day Ticker
[**UiKlines**](MarketAPI.md#UiKlines) | **Get** /api/v3/uiKlines | UIKlines


## AggTrades

> AggTradesResponse AggTrades(ctx).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Compressed/Aggregate trades list


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	fromId := int64(1) // int64 | ID to get aggregate trades from INCLUSIVE. (optional)
	startTime := int64(1735693200000) // int64 | Timestamp in ms to get aggregate trades from INCLUSIVE. (optional)
	endTime := int64(1735693200000) // int64 | Timestamp in ms to get aggregate trades until INCLUSIVE. (optional)
	limit := int32(500) // int32 | Default: 500; Maximum: 1000. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketAPI.AggTrades(context.Background()).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketAPI.AggTrades``: %v\n", err)
		return
	}

	// response from `AggTrades`: AggTradesResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | 
 **fromId** | **int64** | ID to get aggregate trades from INCLUSIVE. | 
 **startTime** | **int64** | Timestamp in ms to get aggregate trades from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get aggregate trades until INCLUSIVE. | 
 **limit** | **int32** | Default: 500; Maximum: 1000. | 

### Return type

[**AggTradesResponse**](AggTradesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## AvgPrice

> AvgPriceResponse AvgPrice(ctx).Symbol(symbol).Execute()

Current average price


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BNBUSDT" // string | 

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketAPI.AvgPrice(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketAPI.AvgPrice``: %v\n", err)
		return
	}

	// response from `AvgPrice`: AvgPriceResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | 

### Return type

[**AvgPriceResponse**](AvgPriceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## Depth

> DepthResponse Depth(ctx).Symbol(symbol).Limit(limit).SymbolStatus(symbolStatus).Execute()

Order book


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	limit := int32(500) // int32 | Default: 500; Maximum: 1000. (optional)
	symbolStatus := models.ExchangeInfoSymbolStatusParameterTrading // ExchangeInfoSymbolStatusParameter |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketAPI.Depth(context.Background()).Symbol(symbol).Limit(limit).SymbolStatus(symbolStatus).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketAPI.Depth``: %v\n", err)
		return
	}

	// response from `Depth`: DepthResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | 
 **limit** | **int32** | Default: 500; Maximum: 1000. | 
 **symbolStatus** | [**ExchangeInfoSymbolStatusParameter**](ExchangeInfoSymbolStatusParameter.md) |  | 

### Return type

[**DepthResponse**](DepthResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetTrades

> GetTradesResponse GetTrades(ctx).Symbol(symbol).Limit(limit).Execute()

Recent trades list


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	limit := int32(500) // int32 | Default: 500; Maximum: 1000. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketAPI.GetTrades(context.Background()).Symbol(symbol).Limit(limit).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketAPI.GetTrades``: %v\n", err)
		return
	}

	// response from `GetTrades`: GetTradesResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | 
 **limit** | **int32** | Default: 500; Maximum: 1000. | 

### Return type

[**GetTradesResponse**](GetTradesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## HistoricalTrades

> HistoricalTradesResponse HistoricalTrades(ctx).Symbol(symbol).Limit(limit).FromId(fromId).Execute()

Old trade lookup


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	limit := int32(500) // int32 | Default: 500; Maximum: 1000. (optional)
	fromId := int64(1) // int64 | ID to get aggregate trades from INCLUSIVE. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketAPI.HistoricalTrades(context.Background()).Symbol(symbol).Limit(limit).FromId(fromId).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketAPI.HistoricalTrades``: %v\n", err)
		return
	}

	// response from `HistoricalTrades`: HistoricalTradesResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | 
 **limit** | **int32** | Default: 500; Maximum: 1000. | 
 **fromId** | **int64** | ID to get aggregate trades from INCLUSIVE. | 

### Return type

[**HistoricalTradesResponse**](HistoricalTradesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## Klines

> KlinesResponse Klines(ctx).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).TimeZone(timeZone).Limit(limit).Execute()

Kline/Candlestick data


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	interval := models.KlinesIntervalParameterInterval1s // KlinesIntervalParameter | 
	startTime := int64(1735693200000) // int64 | Timestamp in ms to get aggregate trades from INCLUSIVE. (optional)
	endTime := int64(1735693200000) // int64 | Timestamp in ms to get aggregate trades until INCLUSIVE. (optional)
	timeZone := "timeZone_example" // string | Default: 0 (UTC) (optional)
	limit := int32(500) // int32 | Default: 500; Maximum: 1000. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketAPI.Klines(context.Background()).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).TimeZone(timeZone).Limit(limit).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketAPI.Klines``: %v\n", err)
		return
	}

	// response from `Klines`: KlinesResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | 
 **interval** | [**KlinesIntervalParameter**](KlinesIntervalParameter.md) |  | 
 **startTime** | **int64** | Timestamp in ms to get aggregate trades from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get aggregate trades until INCLUSIVE. | 
 **timeZone** | **string** | Default: 0 (UTC) | 
 **limit** | **int32** | Default: 500; Maximum: 1000. | 

### Return type

[**KlinesResponse**](KlinesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## Ticker

> TickerResponse Ticker(ctx).Symbol(symbol).Symbols(symbols).WindowSize(windowSize).Type(type_).SymbolStatus(symbolStatus).Execute()

Rolling window price change statistics


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BNBUSDT" // string | Symbol to query (optional)
	symbols := []string{"Inner_example"} // []string | List of symbols to query (optional)
	windowSize := models.TickerWindowSizeParameterWindowSize1m // TickerWindowSizeParameter |  (optional)
	type_ := models.TickerTypeParameterFull // TickerTypeParameter |  (optional)
	symbolStatus := models.ExchangeInfoSymbolStatusParameterTrading // ExchangeInfoSymbolStatusParameter |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketAPI.Ticker(context.Background()).Symbol(symbol).Symbols(symbols).WindowSize(windowSize).Type(type_).SymbolStatus(symbolStatus).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketAPI.Ticker``: %v\n", err)
		return
	}

	// response from `Ticker`: TickerResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol to query | 
 **symbols** | **[]string** | List of symbols to query | 
 **windowSize** | [**TickerWindowSizeParameter**](TickerWindowSizeParameter.md) |  | 
 **type_** | [**TickerTypeParameter**](TickerTypeParameter.md) |  | 
 **symbolStatus** | [**ExchangeInfoSymbolStatusParameter**](ExchangeInfoSymbolStatusParameter.md) |  | 

### Return type

[**TickerResponse**](TickerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## Ticker24hr

> Ticker24hrResponse Ticker24hr(ctx).Symbol(symbol).Symbols(symbols).Type(type_).SymbolStatus(symbolStatus).Execute()

24hr ticker price change statistics


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BNBUSDT" // string | Symbol to query (optional)
	symbols := []string{"Inner_example"} // []string | List of symbols to query (optional)
	type_ := models.TickerTypeParameterFull // TickerTypeParameter |  (optional)
	symbolStatus := models.ExchangeInfoSymbolStatusParameterTrading // ExchangeInfoSymbolStatusParameter |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketAPI.Ticker24hr(context.Background()).Symbol(symbol).Symbols(symbols).Type(type_).SymbolStatus(symbolStatus).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketAPI.Ticker24hr``: %v\n", err)
		return
	}

	// response from `Ticker24hr`: Ticker24hrResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol to query | 
 **symbols** | **[]string** | List of symbols to query | 
 **type_** | [**TickerTypeParameter**](TickerTypeParameter.md) |  | 
 **symbolStatus** | [**ExchangeInfoSymbolStatusParameter**](ExchangeInfoSymbolStatusParameter.md) |  | 

### Return type

[**Ticker24hrResponse**](Ticker24hrResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## TickerBookTicker

> TickerBookTickerResponse TickerBookTicker(ctx).Symbol(symbol).Symbols(symbols).SymbolStatus(symbolStatus).Execute()

Symbol order book ticker


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BNBUSDT" // string | Symbol to query (optional)
	symbols := []string{"Inner_example"} // []string | List of symbols to query (optional)
	symbolStatus := models.ExchangeInfoSymbolStatusParameterTrading // ExchangeInfoSymbolStatusParameter |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketAPI.TickerBookTicker(context.Background()).Symbol(symbol).Symbols(symbols).SymbolStatus(symbolStatus).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketAPI.TickerBookTicker``: %v\n", err)
		return
	}

	// response from `TickerBookTicker`: TickerBookTickerResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol to query | 
 **symbols** | **[]string** | List of symbols to query | 
 **symbolStatus** | [**ExchangeInfoSymbolStatusParameter**](ExchangeInfoSymbolStatusParameter.md) |  | 

### Return type

[**TickerBookTickerResponse**](TickerBookTickerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## TickerPrice

> TickerPriceResponse TickerPrice(ctx).Symbol(symbol).Symbols(symbols).SymbolStatus(symbolStatus).Execute()

Symbol price ticker


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BNBUSDT" // string | Symbol to query (optional)
	symbols := []string{"Inner_example"} // []string | List of symbols to query (optional)
	symbolStatus := models.ExchangeInfoSymbolStatusParameterTrading // ExchangeInfoSymbolStatusParameter |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketAPI.TickerPrice(context.Background()).Symbol(symbol).Symbols(symbols).SymbolStatus(symbolStatus).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketAPI.TickerPrice``: %v\n", err)
		return
	}

	// response from `TickerPrice`: TickerPriceResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol to query | 
 **symbols** | **[]string** | List of symbols to query | 
 **symbolStatus** | [**ExchangeInfoSymbolStatusParameter**](ExchangeInfoSymbolStatusParameter.md) |  | 

### Return type

[**TickerPriceResponse**](TickerPriceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## TickerTradingDay

> TickerTradingDayResponse TickerTradingDay(ctx).Symbol(symbol).Symbols(symbols).TimeZone(timeZone).Type(type_).SymbolStatus(symbolStatus).Execute()

Trading Day Ticker


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BNBUSDT" // string | Symbol to query (optional)
	symbols := []string{"Inner_example"} // []string | List of symbols to query (optional)
	timeZone := "timeZone_example" // string | Default: 0 (UTC) (optional)
	type_ := models.TickerTypeParameterFull // TickerTypeParameter |  (optional)
	symbolStatus := models.ExchangeInfoSymbolStatusParameterTrading // ExchangeInfoSymbolStatusParameter |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketAPI.TickerTradingDay(context.Background()).Symbol(symbol).Symbols(symbols).TimeZone(timeZone).Type(type_).SymbolStatus(symbolStatus).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketAPI.TickerTradingDay``: %v\n", err)
		return
	}

	// response from `TickerTradingDay`: TickerTradingDayResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol to query | 
 **symbols** | **[]string** | List of symbols to query | 
 **timeZone** | **string** | Default: 0 (UTC) | 
 **type_** | [**TickerTypeParameter**](TickerTypeParameter.md) |  | 
 **symbolStatus** | [**ExchangeInfoSymbolStatusParameter**](ExchangeInfoSymbolStatusParameter.md) |  | 

### Return type

[**TickerTradingDayResponse**](TickerTradingDayResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## UiKlines

> UiKlinesResponse UiKlines(ctx).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).TimeZone(timeZone).Limit(limit).Execute()

UIKlines


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	interval := models.KlinesIntervalParameterInterval1s // KlinesIntervalParameter | 
	startTime := int64(1735693200000) // int64 | Timestamp in ms to get aggregate trades from INCLUSIVE. (optional)
	endTime := int64(1735693200000) // int64 | Timestamp in ms to get aggregate trades until INCLUSIVE. (optional)
	timeZone := "timeZone_example" // string | Default: 0 (UTC) (optional)
	limit := int32(500) // int32 | Default: 500; Maximum: 1000. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketAPI.UiKlines(context.Background()).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).TimeZone(timeZone).Limit(limit).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketAPI.UiKlines``: %v\n", err)
		return
	}

	// response from `UiKlines`: UiKlinesResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | 
 **interval** | [**KlinesIntervalParameter**](KlinesIntervalParameter.md) |  | 
 **startTime** | **int64** | Timestamp in ms to get aggregate trades from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get aggregate trades until INCLUSIVE. | 
 **timeZone** | **string** | Default: 0 (UTC) | 
 **limit** | **int32** | Default: 500; Maximum: 1000. | 

### Return type

[**UiKlinesResponse**](UiKlinesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

