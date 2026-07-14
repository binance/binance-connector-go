# \MarketDataAPI

All URIs are relative to *https://www.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AggregatedTrades**](MarketDataAPI.md#AggregatedTrades) | **Get** /bapi/defi/v1/public/alpha-trade/agg-trades | Aggregated Trades
[**FullDepth**](MarketDataAPI.md#FullDepth) | **Get** /bapi/defi/v1/public/alpha-trade/fullDepth | Full Depth
[**GetExchangeInfo**](MarketDataAPI.md#GetExchangeInfo) | **Get** /bapi/defi/v1/public/alpha-trade/get-exchange-info | Get Exchange Info
[**Klines**](MarketDataAPI.md#Klines) | **Get** /bapi/defi/v1/public/alpha-trade/klines | Klines
[**Ticker**](MarketDataAPI.md#Ticker) | **Get** /bapi/defi/v1/public/alpha-trade/ticker | Ticker
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
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "ALPHA_118USDC" // string | Trading pair symbol, e.g. ALPHA_118USDC (use token ID from Token List).
	fromId := int64(58470) // int64 | Starting aggregate trade ID to fetch from. (optional)
	startTime := int64(1752568680000) // int64 | Start timestamp in milliseconds. (optional)
	endTime := int64(1752572280000) // int64 | End timestamp in milliseconds. (optional)
	limit := int64(500) // int64 | Number of results to return. (optional)

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
 **symbol** | **string** | Trading pair symbol, e.g. ALPHA_118USDC (use token ID from Token List). | 
 **fromId** | **int64** | Starting aggregate trade ID to fetch from. | 
 **startTime** | **int64** | Start timestamp in milliseconds. | 
 **endTime** | **int64** | End timestamp in milliseconds. | 
 **limit** | **int64** | Number of results to return. | 

### Return type

[**AggregatedTradesResponse**](AggregatedTradesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## FullDepth

> FullDepthResponse FullDepth(ctx).Symbol(symbol).Limit(limit).Execute()

Full Depth


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/alpha"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "ALPHA_175USDT" // string | Trading pair symbol, e.g. ALPHA_175USDT (use token ID from Token List).
	limit := models.FullDepthLimitParameterLimit5 // FullDepthLimitParameter | Number of price levels to return. Valid values: 5, 10, 20, 50, 100, 500, 1000. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceAlphaClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.FullDepth(context.Background()).Symbol(symbol).Limit(limit).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.FullDepth``: %v\n", err)
		return
	}

	// response from `FullDepth`: FullDepthResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading pair symbol, e.g. ALPHA_175USDT (use token ID from Token List). | 
 **limit** | [**FullDepthLimitParameter**](FullDepthLimitParameter.md) | Number of price levels to return. Valid values: 5, 10, 20, 50, 100, 500, 1000. | 

### Return type

[**FullDepthResponse**](FullDepthResponse.md)

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
	"github.com/binance/binance-connector-go/common/v2/common"
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

Klines


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/alpha"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "ALPHA_175USDT" // string | Trading pair symbol, e.g. ALPHA_175USDT (use token ID from Token List).
	interval := models.KlinesIntervalParameterInterval1s // KlinesIntervalParameter | Kline interval.
	limit := int64(500) // int64 | Number of klines to return. (optional)
	startTime := int64(1752642000000) // int64 | Start timestamp in milliseconds. (optional)
	endTime := int64(1752645599999) // int64 | End timestamp in milliseconds. (optional)

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
 **symbol** | **string** | Trading pair symbol, e.g. ALPHA_175USDT (use token ID from Token List). | 
 **interval** | [**KlinesIntervalParameter**](KlinesIntervalParameter.md) | Kline interval. | 
 **limit** | **int64** | Number of klines to return. | 
 **startTime** | **int64** | Start timestamp in milliseconds. | 
 **endTime** | **int64** | End timestamp in milliseconds. | 

### Return type

[**KlinesResponse**](KlinesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## Ticker

> TickerResponse Ticker(ctx).Symbol(symbol).Execute()

Ticker


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/alpha"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "ALPHA_175USDT" // string | Trading pair symbol, e.g. ALPHA_175USDT (use token ID from Token List).

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
 **symbol** | **string** | Trading pair symbol, e.g. ALPHA_175USDT (use token ID from Token List). | 

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
	"github.com/binance/binance-connector-go/common/v2/common"
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

