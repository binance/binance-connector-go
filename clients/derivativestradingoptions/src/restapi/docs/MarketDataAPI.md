# \MarketDataAPI

All URIs are relative to *https://eapi.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**CheckServerTime**](MarketDataAPI.md#CheckServerTime) | **Get** /eapi/v1/time | Check Server Time
[**ExchangeInformation**](MarketDataAPI.md#ExchangeInformation) | **Get** /eapi/v1/exchangeInfo | Exchange Information
[**HistoricalExerciseRecords**](MarketDataAPI.md#HistoricalExerciseRecords) | **Get** /eapi/v1/exerciseHistory | Historical Exercise Records
[**IndexPrice**](MarketDataAPI.md#IndexPrice) | **Get** /eapi/v1/index | Index Price
[**KlineCandlestickData**](MarketDataAPI.md#KlineCandlestickData) | **Get** /eapi/v1/klines | Kline/Candlestick Data
[**OpenInterest**](MarketDataAPI.md#OpenInterest) | **Get** /eapi/v1/openInterest | Open Interest
[**OptionMarkPrice**](MarketDataAPI.md#OptionMarkPrice) | **Get** /eapi/v1/mark | Option Mark Price
[**OrderBook**](MarketDataAPI.md#OrderBook) | **Get** /eapi/v1/depth | Order Book
[**RecentBlockTradesList**](MarketDataAPI.md#RecentBlockTradesList) | **Get** /eapi/v1/blockTrades | Recent Block Trades List
[**RecentTradesList**](MarketDataAPI.md#RecentTradesList) | **Get** /eapi/v1/trades | Recent Trades List
[**TestConnectivity**](MarketDataAPI.md#TestConnectivity) | **Get** /eapi/v1/ping | Test Connectivity
[**Ticker24hrPriceChangeStatistics**](MarketDataAPI.md#Ticker24hrPriceChangeStatistics) | **Get** /eapi/v1/ticker | 24hr Ticker Price Change Statistics


## CheckServerTime

> CheckServerTimeResponse CheckServerTime(ctx).Execute()

Check Server Time


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.CheckServerTime(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.CheckServerTime``: %v\n", err)
		return
	}

	// response from `CheckServerTime`: CheckServerTimeResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

This endpoint does not need any parameter.

### Return type

[**CheckServerTimeResponse**](CheckServerTimeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## ExchangeInformation

> ExchangeInformationResponse ExchangeInformation(ctx).Execute()

Exchange Information


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.ExchangeInformation(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.ExchangeInformation``: %v\n", err)
		return
	}

	// response from `ExchangeInformation`: ExchangeInformationResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

This endpoint does not need any parameter.

### Return type

[**ExchangeInformationResponse**](ExchangeInformationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## HistoricalExerciseRecords

> HistoricalExerciseRecordsResponse HistoricalExerciseRecords(ctx).Underlying(underlying).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Historical Exercise Records


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	underlying := "underlying_example" // string | underlying, e.g BTCUSDT (optional)
	startTime := int64(1623319461670) // int64 | Start Time, e.g 1593511200000 (optional)
	endTime := int64(1641782889000) // int64 | End Time, e.g 1593512200000 (optional)
	limit := int64(100) // int64 | Number of result sets returned Default:100 Max:1000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.HistoricalExerciseRecords(context.Background()).Underlying(underlying).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.HistoricalExerciseRecords``: %v\n", err)
		return
	}

	// response from `HistoricalExerciseRecords`: HistoricalExerciseRecordsResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **underlying** | **string** | underlying, e.g BTCUSDT | 
 **startTime** | **int64** | Start Time, e.g 1593511200000 | 
 **endTime** | **int64** | End Time, e.g 1593512200000 | 
 **limit** | **int64** | Number of result sets returned Default:100 Max:1000 | 

### Return type

[**HistoricalExerciseRecordsResponse**](HistoricalExerciseRecordsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## IndexPrice

> IndexPriceResponse IndexPrice(ctx).Underlying(underlying).Execute()

Index Price


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	underlying := "underlying_example" // string | Option underlying, e.g BTCUSDT

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.IndexPrice(context.Background()).Underlying(underlying).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.IndexPrice``: %v\n", err)
		return
	}

	// response from `IndexPrice`: IndexPriceResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **underlying** | **string** | Option underlying, e.g BTCUSDT | 

### Return type

[**IndexPriceResponse**](IndexPriceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## KlineCandlestickData

> KlineCandlestickDataResponse KlineCandlestickData(ctx).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Kline/Candlestick Data


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | Option trading pair, e.g BTC-200730-9000-C
	interval := "interval_example" // string | Time interval
	startTime := int64(1623319461670) // int64 | Start Time, e.g 1593511200000 (optional)
	endTime := int64(1641782889000) // int64 | End Time, e.g 1593512200000 (optional)
	limit := int64(100) // int64 | Number of result sets returned Default:100 Max:1000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.KlineCandlestickData(context.Background()).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.KlineCandlestickData``: %v\n", err)
		return
	}

	// response from `KlineCandlestickData`: KlineCandlestickDataResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | 
 **interval** | **string** | Time interval | 
 **startTime** | **int64** | Start Time, e.g 1593511200000 | 
 **endTime** | **int64** | End Time, e.g 1593512200000 | 
 **limit** | **int64** | Number of result sets returned Default:100 Max:1000 | 

### Return type

[**KlineCandlestickDataResponse**](KlineCandlestickDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## OpenInterest

> OpenInterestResponse OpenInterest(ctx).UnderlyingAsset(underlyingAsset).Expiration(expiration).Execute()

Open Interest


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	underlyingAsset := "underlyingAsset_example" // string | underlying asset, e.g ETH/BTC
	expiration := "expiration_example" // string | expiration date, e.g 221225

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.OpenInterest(context.Background()).UnderlyingAsset(underlyingAsset).Expiration(expiration).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.OpenInterest``: %v\n", err)
		return
	}

	// response from `OpenInterest`: OpenInterestResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **underlyingAsset** | **string** | underlying asset, e.g ETH/BTC | 
 **expiration** | **string** | expiration date, e.g 221225 | 

### Return type

[**OpenInterestResponse**](OpenInterestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## OptionMarkPrice

> OptionMarkPriceResponse OptionMarkPrice(ctx).Symbol(symbol).Execute()

Option Mark Price


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | Option trading pair, e.g BTC-200730-9000-C (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.OptionMarkPrice(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.OptionMarkPrice``: %v\n", err)
		return
	}

	// response from `OptionMarkPrice`: OptionMarkPriceResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | 

### Return type

[**OptionMarkPriceResponse**](OptionMarkPriceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## OrderBook

> OrderBookResponse OrderBook(ctx).Symbol(symbol).Limit(limit).Execute()

Order Book


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | Option trading pair, e.g BTC-200730-9000-C
	limit := int64(100) // int64 | Number of result sets returned Default:100 Max:1000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.OrderBook(context.Background()).Symbol(symbol).Limit(limit).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.OrderBook``: %v\n", err)
		return
	}

	// response from `OrderBook`: OrderBookResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | 
 **limit** | **int64** | Number of result sets returned Default:100 Max:1000 | 

### Return type

[**OrderBookResponse**](OrderBookResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## RecentBlockTradesList

> RecentBlockTradesListResponse RecentBlockTradesList(ctx).Symbol(symbol).Limit(limit).Execute()

Recent Block Trades List


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | Option trading pair, e.g BTC-200730-9000-C (optional)
	limit := int64(100) // int64 | Number of result sets returned Default:100 Max:1000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.RecentBlockTradesList(context.Background()).Symbol(symbol).Limit(limit).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.RecentBlockTradesList``: %v\n", err)
		return
	}

	// response from `RecentBlockTradesList`: RecentBlockTradesListResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | 
 **limit** | **int64** | Number of result sets returned Default:100 Max:1000 | 

### Return type

[**RecentBlockTradesListResponse**](RecentBlockTradesListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## RecentTradesList

> RecentTradesListResponse RecentTradesList(ctx).Symbol(symbol).Limit(limit).Execute()

Recent Trades List


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | Option trading pair, e.g BTC-200730-9000-C
	limit := int64(100) // int64 | Number of result sets returned Default:100 Max:1000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.RecentTradesList(context.Background()).Symbol(symbol).Limit(limit).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.RecentTradesList``: %v\n", err)
		return
	}

	// response from `RecentTradesList`: RecentTradesListResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | 
 **limit** | **int64** | Number of result sets returned Default:100 Max:1000 | 

### Return type

[**RecentTradesListResponse**](RecentTradesListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## TestConnectivity

> TestConnectivity(ctx).Execute()

Test Connectivity


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	 err := apiClient.RestApi.MarketDataAPI.TestConnectivity(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.TestConnectivity``: %v\n", err)
		return
	}

}
```

### Path Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: Not defined

[[Back to README]](../../../README.md)


## Ticker24hrPriceChangeStatistics

> Ticker24hrPriceChangeStatisticsResponse Ticker24hrPriceChangeStatistics(ctx).Symbol(symbol).Execute()

24hr Ticker Price Change Statistics


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | Option trading pair, e.g BTC-200730-9000-C (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.Ticker24hrPriceChangeStatistics(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.Ticker24hrPriceChangeStatistics``: %v\n", err)
		return
	}

	// response from `Ticker24hrPriceChangeStatistics`: Ticker24hrPriceChangeStatisticsResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | 

### Return type

[**Ticker24hrPriceChangeStatisticsResponse**](Ticker24hrPriceChangeStatisticsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

