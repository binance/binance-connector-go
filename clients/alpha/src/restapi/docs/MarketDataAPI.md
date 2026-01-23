# \MarketDataAPI

All URIs are relative to *https://www.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AggregatedTrades**](MarketDataAPI.md#AggregatedTrades) | **Get** /bapi/defi/v1/public/alpha-trade/agg-trades | Aggregated Trades
[**GetExchangeInfo**](MarketDataAPI.md#GetExchangeInfo) | **Get** /bapi/defi/v1/public/alpha-trade/get-exchange-info | Get Exchange Info
[**Klines**](MarketDataAPI.md#Klines) | **Get** /bapi/defi/v1/public/alpha-trade/klines | Klines (Candlestick Data)
[**Ticker**](MarketDataAPI.md#Ticker) | **Get** /bapi/defi/v1/public/alpha-trade/ticker | Ticker (24hr Price Statistics)
[**TokenList**](MarketDataAPI.md#TokenList) | **Get** /bapi/defi/v1/public/wallet-direct/buw/wallet/cex/alpha/all/token/list | Token List


## AggregatedTrades

> AggregatedTradesResponse AggregatedTrades(ctx).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Aggregated Trades


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/alpha"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | e.g., \"ALPHA_175USDT\" – use token ID from Token List
	fromId := int64(1) // int64 | starting trade ID to fetch from (optional)
	startTime := int64(1623319461670) // int64 | start timestamp (milliseconds) (optional)
	endTime := int64(1641782889000) // int64 | end timestamp (milliseconds) (optional)
	limit := int64(500) // int64 | number of results to return (default 500, max 1000) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceAlphaClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.AggregatedTrades(context.Background()).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.AggregatedTrades``: %v\n", err)
		return
	}

	// response from `AggregatedTrades`: AggregatedTradesResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | e.g., \&quot;ALPHA_175USDT\&quot; – use token ID from Token List | 
 **fromId** | **int64** | starting trade ID to fetch from | 
 **startTime** | **int64** | start timestamp (milliseconds) | 
 **endTime** | **int64** | end timestamp (milliseconds) | 
 **limit** | **int64** | number of results to return (default 500, max 1000) | 

### Return type

[**AggregatedTradesResponse**](AggregatedTradesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetExchangeInfo

> GetExchangeInfoResponse GetExchangeInfo(ctx).Execute()

Get Exchange Info


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/alpha"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceAlphaClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.GetExchangeInfo(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.GetExchangeInfo``: %v\n", err)
		return
	}

	// response from `GetExchangeInfo`: GetExchangeInfoResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

This endpoint does not need any parameter.

### Return type

[**GetExchangeInfoResponse**](GetExchangeInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## Klines

> KlinesResponse Klines(ctx).Symbol(symbol).Interval(interval).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

Klines (Candlestick Data)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/alpha"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | e.g., \"ALPHA_175USDT\" – use token ID from Token List
	interval := "interval_example" // string | e.g., \"1h\" – supported intervals: 1s, 15s, 1m, 3m, 5m, 15m, 30m, 1h, 2h, 4h, 6h, 8h, 12h, 1d, 3d, 1w, 1M
	limit := int64(500) // int64 | number of results to return (default 500, max 1000) (optional)
	startTime := int64(1623319461670) // int64 | start timestamp (milliseconds) (optional)
	endTime := int64(1641782889000) // int64 | end timestamp (milliseconds) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceAlphaClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.Klines(context.Background()).Symbol(symbol).Interval(interval).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.Klines``: %v\n", err)
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
 **symbol** | **string** | e.g., \&quot;ALPHA_175USDT\&quot; – use token ID from Token List | 
 **interval** | **string** | e.g., \&quot;1h\&quot; – supported intervals: 1s, 15s, 1m, 3m, 5m, 15m, 30m, 1h, 2h, 4h, 6h, 8h, 12h, 1d, 3d, 1w, 1M | 
 **limit** | **int64** | number of results to return (default 500, max 1000) | 
 **startTime** | **int64** | start timestamp (milliseconds) | 
 **endTime** | **int64** | end timestamp (milliseconds) | 

### Return type

[**KlinesResponse**](KlinesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## Ticker

> TickerResponse Ticker(ctx).Symbol(symbol).Execute()

Ticker (24hr Price Statistics)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/alpha"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | e.g., \"ALPHA_175USDT\" – use token ID from Token List

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceAlphaClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.Ticker(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.Ticker``: %v\n", err)
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
 **symbol** | **string** | e.g., \&quot;ALPHA_175USDT\&quot; – use token ID from Token List | 

### Return type

[**TickerResponse**](TickerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## TokenList

> TokenListResponse TokenList(ctx).Execute()

Token List


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/alpha"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceAlphaClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.TokenList(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.TokenList``: %v\n", err)
		return
	}

	// response from `TokenList`: TokenListResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

This endpoint does not need any parameter.

### Return type

[**TokenListResponse**](TokenListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

