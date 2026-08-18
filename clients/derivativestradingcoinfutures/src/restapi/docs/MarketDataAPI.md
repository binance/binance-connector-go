# \MarketDataAPI

All URIs are relative to *https://dapi.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**Basis**](MarketDataAPI.md#Basis) | **Get** /futures/data/basis | Basis
[**CheckServerTime**](MarketDataAPI.md#CheckServerTime) | **Get** /dapi/v1/time | Check Server time
[**CompressedAggregateTradesList**](MarketDataAPI.md#CompressedAggregateTradesList) | **Get** /dapi/v1/aggTrades | Compressed/Aggregate Trades List
[**ContinuousContractKlineCandlestickData**](MarketDataAPI.md#ContinuousContractKlineCandlestickData) | **Get** /dapi/v1/continuousKlines | Continuous Contract Kline/Candlestick Data
[**ExchangeInformation**](MarketDataAPI.md#ExchangeInformation) | **Get** /dapi/v1/exchangeInfo | Exchange Information
[**GetFundingRateHistoryOfPerpetualFutures**](MarketDataAPI.md#GetFundingRateHistoryOfPerpetualFutures) | **Get** /dapi/v1/fundingRate | Get Funding Rate History of Perpetual Futures
[**GetFundingRateInfo**](MarketDataAPI.md#GetFundingRateInfo) | **Get** /dapi/v1/fundingInfo | Get Funding Rate Info
[**IndexPriceAndMarkPrice**](MarketDataAPI.md#IndexPriceAndMarkPrice) | **Get** /dapi/v1/premiumIndex | Index Price and Mark Price
[**IndexPriceKlineCandlestickData**](MarketDataAPI.md#IndexPriceKlineCandlestickData) | **Get** /dapi/v1/indexPriceKlines | Index Price Kline/Candlestick Data
[**KlineCandlestickData**](MarketDataAPI.md#KlineCandlestickData) | **Get** /dapi/v1/klines | Kline/Candlestick Data
[**LongShortRatio**](MarketDataAPI.md#LongShortRatio) | **Get** /futures/data/globalLongShortAccountRatio | Long/Short Ratio
[**MarkPriceKlineCandlestickData**](MarketDataAPI.md#MarkPriceKlineCandlestickData) | **Get** /dapi/v1/markPriceKlines | Mark Price Kline/Candlestick Data
[**OldTradesLookup**](MarketDataAPI.md#OldTradesLookup) | **Get** /dapi/v1/historicalTrades | Old Trades Lookup (MARKET_DATA)
[**OpenInterest**](MarketDataAPI.md#OpenInterest) | **Get** /dapi/v1/openInterest | Open Interest
[**OpenInterestStatistics**](MarketDataAPI.md#OpenInterestStatistics) | **Get** /futures/data/openInterestHist | Open Interest Statistics
[**OrderBook**](MarketDataAPI.md#OrderBook) | **Get** /dapi/v1/depth | Order Book
[**PremiumIndexKlineData**](MarketDataAPI.md#PremiumIndexKlineData) | **Get** /dapi/v1/premiumIndexKlines | Premium index Kline Data
[**QueryIndexPriceConstituents**](MarketDataAPI.md#QueryIndexPriceConstituents) | **Get** /dapi/v1/constituents | Query Index Price Constituents
[**RecentTradesList**](MarketDataAPI.md#RecentTradesList) | **Get** /dapi/v1/trades | Recent Trades List
[**SymbolOrderBookTicker**](MarketDataAPI.md#SymbolOrderBookTicker) | **Get** /dapi/v1/ticker/bookTicker | Symbol Order Book Ticker
[**SymbolPriceTicker**](MarketDataAPI.md#SymbolPriceTicker) | **Get** /dapi/v1/ticker/price | Symbol Price Ticker
[**TakerBuySellVolume**](MarketDataAPI.md#TakerBuySellVolume) | **Get** /futures/data/takerBuySellVol | Taker Buy/Sell Volume
[**TestConnectivity**](MarketDataAPI.md#TestConnectivity) | **Get** /dapi/v1/ping | Test Connectivity
[**Ticker24hrPriceChangeStatistics**](MarketDataAPI.md#Ticker24hrPriceChangeStatistics) | **Get** /dapi/v1/ticker/24hr | 24hr Ticker Price Change Statistics
[**TopTraderLongShortRatioAccounts**](MarketDataAPI.md#TopTraderLongShortRatioAccounts) | **Get** /futures/data/topLongShortAccountRatio | Top Trader Long/Short Account Ratio
[**TopTraderLongShortRatioPositions**](MarketDataAPI.md#TopTraderLongShortRatioPositions) | **Get** /futures/data/topLongShortPositionRatio | Top Trader Long/Short Position Ratio


## Basis

> BasisResponse Basis(ctx).Pair(pair).ContractType(contractType).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

Basis


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	pair := "pair_example" // string | Pair.
	contractType := models.BasisContractTypeParameterPerpetual // BasisContractTypeParameter | Contract type.
	period := models.BasisPeriodParameterPeriod5m // BasisPeriodParameter | Period interval.
	limit := int64(30) // int64 | Maximum number of records to return. (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.Basis(context.Background()).Pair(pair).ContractType(contractType).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.Basis``: %v\n", err)
		return
	}

	// response from `Basis`: BasisResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **pair** | **string** | Pair. | 
 **contractType** | [**BasisContractTypeParameter**](BasisContractTypeParameter.md) | Contract type. | 
 **period** | [**BasisPeriodParameter**](BasisPeriodParameter.md) | Period interval. | 
 **limit** | **int64** | Maximum number of records to return. | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**BasisResponse**](BasisResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CheckServerTime

> CheckServerTimeResponse CheckServerTime(ctx).Execute()

Check Server time


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

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


## CompressedAggregateTradesList

> CompressedAggregateTradesListResponse CompressedAggregateTradesList(ctx).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Compressed/Aggregate Trades List


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | Symbol
	fromId := int64(1) // int64 | ID to get aggregate trades from INCLUSIVE. (optional)
	startTime := int64(1623319461670) // int64 | Timestamp in ms to get aggregate trades from INCLUSIVE. (optional)
	endTime := int64(1641782889000) // int64 | Timestamp in ms to get aggregate trades until INCLUSIVE. (optional)
	limit := int64(30) // int64 | Maximum number of records to return. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.CompressedAggregateTradesList(context.Background()).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.CompressedAggregateTradesList``: %v\n", err)
		return
	}

	// response from `CompressedAggregateTradesList`: CompressedAggregateTradesListResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol | 
 **fromId** | **int64** | ID to get aggregate trades from INCLUSIVE. | 
 **startTime** | **int64** | Timestamp in ms to get aggregate trades from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get aggregate trades until INCLUSIVE. | 
 **limit** | **int64** | Maximum number of records to return. | 

### Return type

[**CompressedAggregateTradesListResponse**](CompressedAggregateTradesListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## ContinuousContractKlineCandlestickData

> ContinuousContractKlineCandlestickDataResponse ContinuousContractKlineCandlestickData(ctx).Pair(pair).ContractType(contractType).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Continuous Contract Kline/Candlestick Data


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	pair := "BTCUSD" // string | After CM migration, accepts both CM and UM pair values.
	contractType := models.BasisContractTypeParameterPerpetual // BasisContractTypeParameter | 
	interval := models.ContinuousContractKlineCandlestickDataIntervalParameterInterval1m // ContinuousContractKlineCandlestickDataIntervalParameter | Interval
	startTime := int64(1623319461670) // int64 | Start time (optional)
	endTime := int64(1641782889000) // int64 | End time (optional)
	limit := int64(30) // int64 | Maximum number of records to return. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.ContinuousContractKlineCandlestickData(context.Background()).Pair(pair).ContractType(contractType).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.ContinuousContractKlineCandlestickData``: %v\n", err)
		return
	}

	// response from `ContinuousContractKlineCandlestickData`: ContinuousContractKlineCandlestickDataResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **pair** | **string** | After CM migration, accepts both CM and UM pair values. | 
 **contractType** | [**BasisContractTypeParameter**](BasisContractTypeParameter.md) |  | 
 **interval** | [**ContinuousContractKlineCandlestickDataIntervalParameter**](ContinuousContractKlineCandlestickDataIntervalParameter.md) | Interval | 
 **startTime** | **int64** | Start time | 
 **endTime** | **int64** | End time | 
 **limit** | **int64** | Maximum number of records to return. | 

### Return type

[**ContinuousContractKlineCandlestickDataResponse**](ContinuousContractKlineCandlestickDataResponse.md)

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

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


## GetFundingRateHistoryOfPerpetualFutures

> GetFundingRateHistoryOfPerpetualFuturesResponse GetFundingRateHistoryOfPerpetualFutures(ctx).Symbol(symbol).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Get Funding Rate History of Perpetual Futures


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | Symbol
	startTime := int64(1623319461670) // int64 | Timestamp in ms to get funding rate from INCLUSIVE. (optional)
	endTime := int64(1641782889000) // int64 | Timestamp in ms to get funding rate until INCLUSIVE. (optional)
	limit := int64(30) // int64 | Maximum number of records to return. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.GetFundingRateHistoryOfPerpetualFutures(context.Background()).Symbol(symbol).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.GetFundingRateHistoryOfPerpetualFutures``: %v\n", err)
		return
	}

	// response from `GetFundingRateHistoryOfPerpetualFutures`: GetFundingRateHistoryOfPerpetualFuturesResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol | 
 **startTime** | **int64** | Timestamp in ms to get funding rate from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding rate until INCLUSIVE. | 
 **limit** | **int64** | Maximum number of records to return. | 

### Return type

[**GetFundingRateHistoryOfPerpetualFuturesResponse**](GetFundingRateHistoryOfPerpetualFuturesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetFundingRateInfo

> GetFundingRateInfoResponse GetFundingRateInfo(ctx).Execute()

Get Funding Rate Info


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.GetFundingRateInfo(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.GetFundingRateInfo``: %v\n", err)
		return
	}

	// response from `GetFundingRateInfo`: GetFundingRateInfoResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

This endpoint does not need any parameter.

### Return type

[**GetFundingRateInfoResponse**](GetFundingRateInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## IndexPriceAndMarkPrice

> IndexPriceAndMarkPriceResponse IndexPriceAndMarkPrice(ctx).Symbol(symbol).Pair(pair).Execute()

Index Price and Mark Price


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSD_PERP" // string |  (optional)
	pair := "BTCUSD" // string |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.IndexPriceAndMarkPrice(context.Background()).Symbol(symbol).Pair(pair).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.IndexPriceAndMarkPrice``: %v\n", err)
		return
	}

	// response from `IndexPriceAndMarkPrice`: IndexPriceAndMarkPriceResponse
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
 **pair** | **string** |  | 

### Return type

[**IndexPriceAndMarkPriceResponse**](IndexPriceAndMarkPriceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## IndexPriceKlineCandlestickData

> IndexPriceKlineCandlestickDataResponse IndexPriceKlineCandlestickData(ctx).Pair(pair).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Index Price Kline/Candlestick Data


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	pair := "BTCUSD" // string | After CM migration, accepts both CM and UM pair values.
	interval := models.ContinuousContractKlineCandlestickDataIntervalParameterInterval1m // ContinuousContractKlineCandlestickDataIntervalParameter | Interval
	startTime := int64(1623319461670) // int64 | Start time (optional)
	endTime := int64(1641782889000) // int64 | End time (optional)
	limit := int64(30) // int64 | Maximum number of records to return. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.IndexPriceKlineCandlestickData(context.Background()).Pair(pair).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.IndexPriceKlineCandlestickData``: %v\n", err)
		return
	}

	// response from `IndexPriceKlineCandlestickData`: IndexPriceKlineCandlestickDataResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **pair** | **string** | After CM migration, accepts both CM and UM pair values. | 
 **interval** | [**ContinuousContractKlineCandlestickDataIntervalParameter**](ContinuousContractKlineCandlestickDataIntervalParameter.md) | Interval | 
 **startTime** | **int64** | Start time | 
 **endTime** | **int64** | End time | 
 **limit** | **int64** | Maximum number of records to return. | 

### Return type

[**IndexPriceKlineCandlestickDataResponse**](IndexPriceKlineCandlestickDataResponse.md)

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSD" // string | After CM migration, accepts both CM and UM symbols.
	interval := models.ContinuousContractKlineCandlestickDataIntervalParameterInterval1m // ContinuousContractKlineCandlestickDataIntervalParameter | Interval
	startTime := int64(1623319461670) // int64 | Start time (optional)
	endTime := int64(1641782889000) // int64 | End time (optional)
	limit := int64(30) // int64 | Maximum number of records to return. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

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
 **symbol** | **string** | After CM migration, accepts both CM and UM symbols. | 
 **interval** | [**ContinuousContractKlineCandlestickDataIntervalParameter**](ContinuousContractKlineCandlestickDataIntervalParameter.md) | Interval | 
 **startTime** | **int64** | Start time | 
 **endTime** | **int64** | End time | 
 **limit** | **int64** | Maximum number of records to return. | 

### Return type

[**KlineCandlestickDataResponse**](KlineCandlestickDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## LongShortRatio

> LongShortRatioResponse LongShortRatio(ctx).Pair(pair).Period(period).ContractType(contractType).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

Long/Short Ratio


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	pair := "pair_example" // string | BTCUSD
	period := models.BasisPeriodParameterPeriod5m // BasisPeriodParameter | 
	contractType := models.BasisContractTypeParameterPerpetual // BasisContractTypeParameter | Contract type filter. If omitted, returns aggregated data across all contract types. (optional)
	limit := int64(30) // int64 | Maximum number of records to return. (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.LongShortRatio(context.Background()).Pair(pair).Period(period).ContractType(contractType).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.LongShortRatio``: %v\n", err)
		return
	}

	// response from `LongShortRatio`: LongShortRatioResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **pair** | **string** | BTCUSD | 
 **period** | [**BasisPeriodParameter**](BasisPeriodParameter.md) |  | 
 **contractType** | [**BasisContractTypeParameter**](BasisContractTypeParameter.md) | Contract type filter. If omitted, returns aggregated data across all contract types. | 
 **limit** | **int64** | Maximum number of records to return. | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**LongShortRatioResponse**](LongShortRatioResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MarkPriceKlineCandlestickData

> MarkPriceKlineCandlestickDataResponse MarkPriceKlineCandlestickData(ctx).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Mark Price Kline/Candlestick Data


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSD" // string | After CM migration, accepts both CM and UM symbols.
	interval := models.ContinuousContractKlineCandlestickDataIntervalParameterInterval1m // ContinuousContractKlineCandlestickDataIntervalParameter | Interval
	startTime := int64(1623319461670) // int64 | Start time (optional)
	endTime := int64(1641782889000) // int64 | End time (optional)
	limit := int64(30) // int64 | Maximum number of records to return. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.MarkPriceKlineCandlestickData(context.Background()).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.MarkPriceKlineCandlestickData``: %v\n", err)
		return
	}

	// response from `MarkPriceKlineCandlestickData`: MarkPriceKlineCandlestickDataResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | After CM migration, accepts both CM and UM symbols. | 
 **interval** | [**ContinuousContractKlineCandlestickDataIntervalParameter**](ContinuousContractKlineCandlestickDataIntervalParameter.md) | Interval | 
 **startTime** | **int64** | Start time | 
 **endTime** | **int64** | End time | 
 **limit** | **int64** | Maximum number of records to return. | 

### Return type

[**MarkPriceKlineCandlestickDataResponse**](MarkPriceKlineCandlestickDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## OldTradesLookup

> OldTradesLookupResponse OldTradesLookup(ctx).Symbol(symbol).Limit(limit).FromId(fromId).Execute()

Old Trades Lookup (MARKET_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | Symbol
	limit := int64(30) // int64 | Maximum number of records to return. (optional)
	fromId := int64(595103) // int64 | TradeId to fetch from. Default gets most recent trades. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.OldTradesLookup(context.Background()).Symbol(symbol).Limit(limit).FromId(fromId).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.OldTradesLookup``: %v\n", err)
		return
	}

	// response from `OldTradesLookup`: OldTradesLookupResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol | 
 **limit** | **int64** | Maximum number of records to return. | 
 **fromId** | **int64** | TradeId to fetch from. Default gets most recent trades. | 

### Return type

[**OldTradesLookupResponse**](OldTradesLookupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## OpenInterest

> OpenInterestResponse OpenInterest(ctx).Symbol(symbol).Execute()

Open Interest


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSD_200626" // string | Symbol

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.OpenInterest(context.Background()).Symbol(symbol).Execute()
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
 **symbol** | **string** | Symbol | 

### Return type

[**OpenInterestResponse**](OpenInterestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## OpenInterestStatistics

> OpenInterestStatisticsResponse OpenInterestStatistics(ctx).Pair(pair).Period(period).ContractType(contractType).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

Open Interest Statistics


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	pair := "BTCUSD" // string | 
	period := models.BasisPeriodParameterPeriod5m // BasisPeriodParameter | 
	contractType := models.BasisContractTypeParameterPerpetual // BasisContractTypeParameter | Contract type filter. If omitted, returns aggregated data across all contract types. (optional)
	limit := int64(30) // int64 | Maximum number of records to return. (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.OpenInterestStatistics(context.Background()).Pair(pair).Period(period).ContractType(contractType).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.OpenInterestStatistics``: %v\n", err)
		return
	}

	// response from `OpenInterestStatistics`: OpenInterestStatisticsResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **pair** | **string** |  | 
 **period** | [**BasisPeriodParameter**](BasisPeriodParameter.md) |  | 
 **contractType** | [**BasisContractTypeParameter**](BasisContractTypeParameter.md) | Contract type filter. If omitted, returns aggregated data across all contract types. | 
 **limit** | **int64** | Maximum number of records to return. | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**OpenInterestStatisticsResponse**](OpenInterestStatisticsResponse.md)

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSD_PERP" // string | Symbol
	limit := int64(500) // int64 | Valid limits:[5, 10, 20, 50, 100, 500, 1000]. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

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
 **symbol** | **string** | Symbol | 
 **limit** | **int64** | Valid limits:[5, 10, 20, 50, 100, 500, 1000]. | 

### Return type

[**OrderBookResponse**](OrderBookResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## PremiumIndexKlineData

> PremiumIndexKlineDataResponse PremiumIndexKlineData(ctx).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Premium index Kline Data


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSD" // string | After CM migration, accepts both CM and UM symbols.
	interval := models.ContinuousContractKlineCandlestickDataIntervalParameterInterval1m // ContinuousContractKlineCandlestickDataIntervalParameter | Interval
	startTime := int64(1623319461670) // int64 | Start time (optional)
	endTime := int64(1641782889000) // int64 | End time (optional)
	limit := int64(30) // int64 | Maximum number of records to return. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.PremiumIndexKlineData(context.Background()).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.PremiumIndexKlineData``: %v\n", err)
		return
	}

	// response from `PremiumIndexKlineData`: PremiumIndexKlineDataResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | After CM migration, accepts both CM and UM symbols. | 
 **interval** | [**ContinuousContractKlineCandlestickDataIntervalParameter**](ContinuousContractKlineCandlestickDataIntervalParameter.md) | Interval | 
 **startTime** | **int64** | Start time | 
 **endTime** | **int64** | End time | 
 **limit** | **int64** | Maximum number of records to return. | 

### Return type

[**PremiumIndexKlineDataResponse**](PremiumIndexKlineDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryIndexPriceConstituents

> QueryIndexPriceConstituentsResponse QueryIndexPriceConstituents(ctx).Symbol(symbol).Execute()

Query Index Price Constituents


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSD" // string | Symbol

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.QueryIndexPriceConstituents(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.QueryIndexPriceConstituents``: %v\n", err)
		return
	}

	// response from `QueryIndexPriceConstituents`: QueryIndexPriceConstituentsResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol | 

### Return type

[**QueryIndexPriceConstituentsResponse**](QueryIndexPriceConstituentsResponse.md)

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSD" // string | Symbol
	limit := int64(30) // int64 | Maximum number of records to return. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

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
 **symbol** | **string** | Symbol | 
 **limit** | **int64** | Maximum number of records to return. | 

### Return type

[**RecentTradesListResponse**](RecentTradesListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SymbolOrderBookTicker

> SymbolOrderBookTickerResponse SymbolOrderBookTicker(ctx).Symbol(symbol).Pair(pair).Execute()

Symbol Order Book Ticker


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSD_200626" // string | Symbol (optional)
	pair := "BTCUSD" // string | Symbol (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.SymbolOrderBookTicker(context.Background()).Symbol(symbol).Pair(pair).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.SymbolOrderBookTicker``: %v\n", err)
		return
	}

	// response from `SymbolOrderBookTicker`: SymbolOrderBookTickerResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol | 
 **pair** | **string** | Symbol | 

### Return type

[**SymbolOrderBookTickerResponse**](SymbolOrderBookTickerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SymbolPriceTicker

> SymbolPriceTickerResponse SymbolPriceTicker(ctx).Symbol(symbol).Pair(pair).Execute()

Symbol Price Ticker


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSD_200626" // string | Symbol (optional)
	pair := "BTCUSD" // string | Pair (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.SymbolPriceTicker(context.Background()).Symbol(symbol).Pair(pair).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.SymbolPriceTicker``: %v\n", err)
		return
	}

	// response from `SymbolPriceTicker`: SymbolPriceTickerResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol | 
 **pair** | **string** | Pair | 

### Return type

[**SymbolPriceTickerResponse**](SymbolPriceTickerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## TakerBuySellVolume

> TakerBuySellVolumeResponse TakerBuySellVolume(ctx).Pair(pair).ContractType(contractType).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

Taker Buy/Sell Volume


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	pair := "BTCUSD" // string | 
	contractType := models.TakerBuySellVolumeContractTypeParameterAll // TakerBuySellVolumeContractTypeParameter | 
	period := models.BasisPeriodParameterPeriod5m // BasisPeriodParameter | 
	limit := int64(30) // int64 | Maximum number of records to return. (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.TakerBuySellVolume(context.Background()).Pair(pair).ContractType(contractType).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.TakerBuySellVolume``: %v\n", err)
		return
	}

	// response from `TakerBuySellVolume`: TakerBuySellVolumeResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **pair** | **string** |  | 
 **contractType** | [**TakerBuySellVolumeContractTypeParameter**](TakerBuySellVolumeContractTypeParameter.md) |  | 
 **period** | [**BasisPeriodParameter**](BasisPeriodParameter.md) |  | 
 **limit** | **int64** | Maximum number of records to return. | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**TakerBuySellVolumeResponse**](TakerBuySellVolumeResponse.md)

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

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

> Ticker24hrPriceChangeStatisticsResponse Ticker24hrPriceChangeStatistics(ctx).Symbol(symbol).Pair(pair).Execute()

24hr Ticker Price Change Statistics


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSD_200925" // string | Symbol (optional)
	pair := "BTCUSD" // string | Pair (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.Ticker24hrPriceChangeStatistics(context.Background()).Symbol(symbol).Pair(pair).Execute()
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
 **symbol** | **string** | Symbol | 
 **pair** | **string** | Pair | 

### Return type

[**Ticker24hrPriceChangeStatisticsResponse**](Ticker24hrPriceChangeStatisticsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## TopTraderLongShortRatioAccounts

> TopTraderLongShortRatioAccountsResponse TopTraderLongShortRatioAccounts(ctx).Pair(pair).Period(period).ContractType(contractType).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

Top Trader Long/Short Account Ratio


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	pair := "BTCUSD" // string | Pair
	period := models.BasisPeriodParameterPeriod5m // BasisPeriodParameter | 
	contractType := models.BasisContractTypeParameterPerpetual // BasisContractTypeParameter | Contract type filter. If omitted, returns aggregated data across all contract types. (optional)
	limit := int64(30) // int64 | Maximum number of records to return. (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.TopTraderLongShortRatioAccounts(context.Background()).Pair(pair).Period(period).ContractType(contractType).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.TopTraderLongShortRatioAccounts``: %v\n", err)
		return
	}

	// response from `TopTraderLongShortRatioAccounts`: TopTraderLongShortRatioAccountsResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **pair** | **string** | Pair | 
 **period** | [**BasisPeriodParameter**](BasisPeriodParameter.md) |  | 
 **contractType** | [**BasisContractTypeParameter**](BasisContractTypeParameter.md) | Contract type filter. If omitted, returns aggregated data across all contract types. | 
 **limit** | **int64** | Maximum number of records to return. | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**TopTraderLongShortRatioAccountsResponse**](TopTraderLongShortRatioAccountsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## TopTraderLongShortRatioPositions

> TopTraderLongShortRatioPositionsResponse TopTraderLongShortRatioPositions(ctx).Pair(pair).Period(period).ContractType(contractType).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

Top Trader Long/Short Position Ratio


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	pair := "BTCUSD" // string | 
	period := models.BasisPeriodParameterPeriod5m // BasisPeriodParameter | 
	contractType := models.BasisContractTypeParameterPerpetual // BasisContractTypeParameter | Contract type filter. If omitted, returns aggregated data across all contract types. (optional)
	limit := int64(30) // int64 | Maximum number of records to return. (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.TopTraderLongShortRatioPositions(context.Background()).Pair(pair).Period(period).ContractType(contractType).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.TopTraderLongShortRatioPositions``: %v\n", err)
		return
	}

	// response from `TopTraderLongShortRatioPositions`: TopTraderLongShortRatioPositionsResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **pair** | **string** |  | 
 **period** | [**BasisPeriodParameter**](BasisPeriodParameter.md) |  | 
 **contractType** | [**BasisContractTypeParameter**](BasisContractTypeParameter.md) | Contract type filter. If omitted, returns aggregated data across all contract types. | 
 **limit** | **int64** | Maximum number of records to return. | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**TopTraderLongShortRatioPositionsResponse**](TopTraderLongShortRatioPositionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

