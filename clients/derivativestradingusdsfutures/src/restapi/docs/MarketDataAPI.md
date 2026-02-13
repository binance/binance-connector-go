# \MarketDataAPI

All URIs are relative to *https://fapi.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AdlRisk**](MarketDataAPI.md#AdlRisk) | **Get** /fapi/v1/symbolAdlRisk | ADL Risk
[**Basis**](MarketDataAPI.md#Basis) | **Get** /futures/data/basis | Basis
[**CheckServerTime**](MarketDataAPI.md#CheckServerTime) | **Get** /fapi/v1/time | Check Server Time
[**CompositeIndexSymbolInformation**](MarketDataAPI.md#CompositeIndexSymbolInformation) | **Get** /fapi/v1/indexInfo | Composite Index Symbol Information
[**CompressedAggregateTradesList**](MarketDataAPI.md#CompressedAggregateTradesList) | **Get** /fapi/v1/aggTrades | Compressed/Aggregate Trades List
[**ContinuousContractKlineCandlestickData**](MarketDataAPI.md#ContinuousContractKlineCandlestickData) | **Get** /fapi/v1/continuousKlines | Continuous Contract Kline/Candlestick Data
[**ExchangeInformation**](MarketDataAPI.md#ExchangeInformation) | **Get** /fapi/v1/exchangeInfo | Exchange Information
[**GetFundingRateHistory**](MarketDataAPI.md#GetFundingRateHistory) | **Get** /fapi/v1/fundingRate | Get Funding Rate History
[**GetFundingRateInfo**](MarketDataAPI.md#GetFundingRateInfo) | **Get** /fapi/v1/fundingInfo | Get Funding Rate Info
[**IndexPriceKlineCandlestickData**](MarketDataAPI.md#IndexPriceKlineCandlestickData) | **Get** /fapi/v1/indexPriceKlines | Index Price Kline/Candlestick Data
[**KlineCandlestickData**](MarketDataAPI.md#KlineCandlestickData) | **Get** /fapi/v1/klines | Kline/Candlestick Data
[**LongShortRatio**](MarketDataAPI.md#LongShortRatio) | **Get** /futures/data/globalLongShortAccountRatio | Long/Short Ratio
[**MarkPrice**](MarketDataAPI.md#MarkPrice) | **Get** /fapi/v1/premiumIndex | Mark Price
[**MarkPriceKlineCandlestickData**](MarketDataAPI.md#MarkPriceKlineCandlestickData) | **Get** /fapi/v1/markPriceKlines | Mark Price Kline/Candlestick Data
[**MultiAssetsModeAssetIndex**](MarketDataAPI.md#MultiAssetsModeAssetIndex) | **Get** /fapi/v1/assetIndex | Multi-Assets Mode Asset Index
[**OldTradesLookup**](MarketDataAPI.md#OldTradesLookup) | **Get** /fapi/v1/historicalTrades | Old Trades Lookup (MARKET_DATA)
[**OpenInterest**](MarketDataAPI.md#OpenInterest) | **Get** /fapi/v1/openInterest | Open Interest
[**OpenInterestStatistics**](MarketDataAPI.md#OpenInterestStatistics) | **Get** /futures/data/openInterestHist | Open Interest Statistics
[**OrderBook**](MarketDataAPI.md#OrderBook) | **Get** /fapi/v1/depth | Order Book
[**PremiumIndexKlineData**](MarketDataAPI.md#PremiumIndexKlineData) | **Get** /fapi/v1/premiumIndexKlines | Premium index Kline Data
[**QuarterlyContractSettlementPrice**](MarketDataAPI.md#QuarterlyContractSettlementPrice) | **Get** /futures/data/delivery-price | Quarterly Contract Settlement Price
[**QueryIndexPriceConstituents**](MarketDataAPI.md#QueryIndexPriceConstituents) | **Get** /fapi/v1/constituents | Query Index Price Constituents
[**QueryInsuranceFundBalanceSnapshot**](MarketDataAPI.md#QueryInsuranceFundBalanceSnapshot) | **Get** /fapi/v1/insuranceBalance | Query Insurance Fund Balance Snapshot
[**RecentTradesList**](MarketDataAPI.md#RecentTradesList) | **Get** /fapi/v1/trades | Recent Trades List
[**RpiOrderBook**](MarketDataAPI.md#RpiOrderBook) | **Get** /fapi/v1/rpiDepth | RPI Order Book
[**SymbolOrderBookTicker**](MarketDataAPI.md#SymbolOrderBookTicker) | **Get** /fapi/v1/ticker/bookTicker | Symbol Order Book Ticker
[**SymbolPriceTicker**](MarketDataAPI.md#SymbolPriceTicker) | **Get** /fapi/v1/ticker/price | Symbol Price Ticker
[**SymbolPriceTickerV2**](MarketDataAPI.md#SymbolPriceTickerV2) | **Get** /fapi/v2/ticker/price | Symbol Price Ticker V2
[**TakerBuySellVolume**](MarketDataAPI.md#TakerBuySellVolume) | **Get** /futures/data/takerlongshortRatio | Taker Buy/Sell Volume
[**TestConnectivity**](MarketDataAPI.md#TestConnectivity) | **Get** /fapi/v1/ping | Test Connectivity
[**Ticker24hrPriceChangeStatistics**](MarketDataAPI.md#Ticker24hrPriceChangeStatistics) | **Get** /fapi/v1/ticker/24hr | 24hr Ticker Price Change Statistics
[**TopTraderLongShortRatioAccounts**](MarketDataAPI.md#TopTraderLongShortRatioAccounts) | **Get** /futures/data/topLongShortAccountRatio | Top Trader Long/Short Ratio (Accounts)
[**TopTraderLongShortRatioPositions**](MarketDataAPI.md#TopTraderLongShortRatioPositions) | **Get** /futures/data/topLongShortPositionRatio | Top Trader Long/Short Ratio (Positions)
[**TradingSchedule**](MarketDataAPI.md#TradingSchedule) | **Get** /fapi/v1/tradingSchedule | Trading Schedule


## AdlRisk

> AdlRiskResponse AdlRisk(ctx).Symbol(symbol).Execute()

ADL Risk


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.AdlRisk(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.AdlRisk``: %v\n", err)
		return
	}

	// response from `AdlRisk`: AdlRiskResponse
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

[**AdlRiskResponse**](AdlRiskResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


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

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	pair := "pair_example" // string | 
	contractType := models.BasisContractTypeParameterPerpetual // BasisContractTypeParameter | 
	period := models.BasisPeriodParameterPeriod5m // BasisPeriodParameter | \"5m\",\"15m\",\"30m\",\"1h\",\"2h\",\"4h\",\"6h\",\"12h\",\"1d\"
	limit := int64(30) // int64 | Default 30,Max 500
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

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
 **pair** | **string** |  | 
 **contractType** | [**BasisContractTypeParameter**](BasisContractTypeParameter.md) |  | 
 **period** | [**BasisPeriodParameter**](BasisPeriodParameter.md) | \&quot;5m\&quot;,\&quot;15m\&quot;,\&quot;30m\&quot;,\&quot;1h\&quot;,\&quot;2h\&quot;,\&quot;4h\&quot;,\&quot;6h\&quot;,\&quot;12h\&quot;,\&quot;1d\&quot; | 
 **limit** | **int64** | Default 30,Max 500 | 
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

Check Server Time


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

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


## CompositeIndexSymbolInformation

> CompositeIndexSymbolInformationResponse CompositeIndexSymbolInformation(ctx).Symbol(symbol).Execute()

Composite Index Symbol Information


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.CompositeIndexSymbolInformation(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.CompositeIndexSymbolInformation``: %v\n", err)
		return
	}

	// response from `CompositeIndexSymbolInformation`: CompositeIndexSymbolInformationResponse
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

[**CompositeIndexSymbolInformationResponse**](CompositeIndexSymbolInformationResponse.md)

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | 
	fromId := int64(1) // int64 | ID to get aggregate trades from INCLUSIVE. (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	limit := int64(100) // int64 | Default 100; max 1000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

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
 **symbol** | **string** |  | 
 **fromId** | **int64** | ID to get aggregate trades from INCLUSIVE. | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int64** | Default 100; max 1000 | 

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	pair := "pair_example" // string | 
	contractType := models.BasisContractTypeParameterPerpetual // BasisContractTypeParameter | 
	interval := models.ContinuousContractKlineCandlestickDataIntervalParameterInterval1m // ContinuousContractKlineCandlestickDataIntervalParameter | 
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	limit := int64(100) // int64 | Default 100; max 1000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

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
 **pair** | **string** |  | 
 **contractType** | [**BasisContractTypeParameter**](BasisContractTypeParameter.md) |  | 
 **interval** | [**ContinuousContractKlineCandlestickDataIntervalParameter**](ContinuousContractKlineCandlestickDataIntervalParameter.md) |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int64** | Default 100; max 1000 | 

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

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


## GetFundingRateHistory

> GetFundingRateHistoryResponse GetFundingRateHistory(ctx).Symbol(symbol).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Get Funding Rate History


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string |  (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	limit := int64(100) // int64 | Default 100; max 1000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.GetFundingRateHistory(context.Background()).Symbol(symbol).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.GetFundingRateHistory``: %v\n", err)
		return
	}

	// response from `GetFundingRateHistory`: GetFundingRateHistoryResponse
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
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int64** | Default 100; max 1000 | 

### Return type

[**GetFundingRateHistoryResponse**](GetFundingRateHistoryResponse.md)

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	pair := "pair_example" // string | 
	interval := models.ContinuousContractKlineCandlestickDataIntervalParameterInterval1m // ContinuousContractKlineCandlestickDataIntervalParameter | 
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	limit := int64(100) // int64 | Default 100; max 1000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

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
 **pair** | **string** |  | 
 **interval** | [**ContinuousContractKlineCandlestickDataIntervalParameter**](ContinuousContractKlineCandlestickDataIntervalParameter.md) |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int64** | Default 100; max 1000 | 

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | 
	interval := models.ContinuousContractKlineCandlestickDataIntervalParameterInterval1m // ContinuousContractKlineCandlestickDataIntervalParameter | 
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	limit := int64(100) // int64 | Default 100; max 1000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

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
 **symbol** | **string** |  | 
 **interval** | [**ContinuousContractKlineCandlestickDataIntervalParameter**](ContinuousContractKlineCandlestickDataIntervalParameter.md) |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int64** | Default 100; max 1000 | 

### Return type

[**KlineCandlestickDataResponse**](KlineCandlestickDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## LongShortRatio

> LongShortRatioResponse LongShortRatio(ctx).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

Long/Short Ratio


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | 
	period := models.BasisPeriodParameterPeriod5m // BasisPeriodParameter | \"5m\",\"15m\",\"30m\",\"1h\",\"2h\",\"4h\",\"6h\",\"12h\",\"1d\"
	limit := int64(100) // int64 | Default 100; max 1000 (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.LongShortRatio(context.Background()).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
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
 **symbol** | **string** |  | 
 **period** | [**BasisPeriodParameter**](BasisPeriodParameter.md) | \&quot;5m\&quot;,\&quot;15m\&quot;,\&quot;30m\&quot;,\&quot;1h\&quot;,\&quot;2h\&quot;,\&quot;4h\&quot;,\&quot;6h\&quot;,\&quot;12h\&quot;,\&quot;1d\&quot; | 
 **limit** | **int64** | Default 100; max 1000 | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**LongShortRatioResponse**](LongShortRatioResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MarkPrice

> MarkPriceResponse MarkPrice(ctx).Symbol(symbol).Execute()

Mark Price


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.MarkPrice(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.MarkPrice``: %v\n", err)
		return
	}

	// response from `MarkPrice`: MarkPriceResponse
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

[**MarkPriceResponse**](MarkPriceResponse.md)

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | 
	interval := models.ContinuousContractKlineCandlestickDataIntervalParameterInterval1m // ContinuousContractKlineCandlestickDataIntervalParameter | 
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	limit := int64(100) // int64 | Default 100; max 1000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

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
 **symbol** | **string** |  | 
 **interval** | [**ContinuousContractKlineCandlestickDataIntervalParameter**](ContinuousContractKlineCandlestickDataIntervalParameter.md) |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int64** | Default 100; max 1000 | 

### Return type

[**MarkPriceKlineCandlestickDataResponse**](MarkPriceKlineCandlestickDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MultiAssetsModeAssetIndex

> MultiAssetsModeAssetIndexResponse MultiAssetsModeAssetIndex(ctx).Symbol(symbol).Execute()

Multi-Assets Mode Asset Index


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.MultiAssetsModeAssetIndex(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.MultiAssetsModeAssetIndex``: %v\n", err)
		return
	}

	// response from `MultiAssetsModeAssetIndex`: MultiAssetsModeAssetIndexResponse
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

[**MultiAssetsModeAssetIndexResponse**](MultiAssetsModeAssetIndexResponse.md)

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | 
	limit := int64(100) // int64 | Default 100; max 1000 (optional)
	fromId := int64(1) // int64 | ID to get aggregate trades from INCLUSIVE. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

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
 **symbol** | **string** |  | 
 **limit** | **int64** | Default 100; max 1000 | 
 **fromId** | **int64** | ID to get aggregate trades from INCLUSIVE. | 

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | 

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

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
 **symbol** | **string** |  | 

### Return type

[**OpenInterestResponse**](OpenInterestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## OpenInterestStatistics

> OpenInterestStatisticsResponse OpenInterestStatistics(ctx).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

Open Interest Statistics


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | 
	period := models.BasisPeriodParameterPeriod5m // BasisPeriodParameter | \"5m\",\"15m\",\"30m\",\"1h\",\"2h\",\"4h\",\"6h\",\"12h\",\"1d\"
	limit := int64(100) // int64 | Default 100; max 1000 (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.OpenInterestStatistics(context.Background()).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
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
 **symbol** | **string** |  | 
 **period** | [**BasisPeriodParameter**](BasisPeriodParameter.md) | \&quot;5m\&quot;,\&quot;15m\&quot;,\&quot;30m\&quot;,\&quot;1h\&quot;,\&quot;2h\&quot;,\&quot;4h\&quot;,\&quot;6h\&quot;,\&quot;12h\&quot;,\&quot;1d\&quot; | 
 **limit** | **int64** | Default 100; max 1000 | 
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

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | 
	limit := int64(100) // int64 | Default 100; max 1000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

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
 **symbol** | **string** |  | 
 **limit** | **int64** | Default 100; max 1000 | 

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | 
	interval := models.ContinuousContractKlineCandlestickDataIntervalParameterInterval1m // ContinuousContractKlineCandlestickDataIntervalParameter | 
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	limit := int64(100) // int64 | Default 100; max 1000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

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
 **symbol** | **string** |  | 
 **interval** | [**ContinuousContractKlineCandlestickDataIntervalParameter**](ContinuousContractKlineCandlestickDataIntervalParameter.md) |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int64** | Default 100; max 1000 | 

### Return type

[**PremiumIndexKlineDataResponse**](PremiumIndexKlineDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QuarterlyContractSettlementPrice

> QuarterlyContractSettlementPriceResponse QuarterlyContractSettlementPrice(ctx).Pair(pair).Execute()

Quarterly Contract Settlement Price


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	pair := "pair_example" // string | 

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.QuarterlyContractSettlementPrice(context.Background()).Pair(pair).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.QuarterlyContractSettlementPrice``: %v\n", err)
		return
	}

	// response from `QuarterlyContractSettlementPrice`: QuarterlyContractSettlementPriceResponse
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

### Return type

[**QuarterlyContractSettlementPriceResponse**](QuarterlyContractSettlementPriceResponse.md)

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | 

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

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
 **symbol** | **string** |  | 

### Return type

[**QueryIndexPriceConstituentsResponse**](QueryIndexPriceConstituentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryInsuranceFundBalanceSnapshot

> QueryInsuranceFundBalanceSnapshotResponse QueryInsuranceFundBalanceSnapshot(ctx).Symbol(symbol).Execute()

Query Insurance Fund Balance Snapshot


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.QueryInsuranceFundBalanceSnapshot(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.QueryInsuranceFundBalanceSnapshot``: %v\n", err)
		return
	}

	// response from `QueryInsuranceFundBalanceSnapshot`: QueryInsuranceFundBalanceSnapshotResponse
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

[**QueryInsuranceFundBalanceSnapshotResponse**](QueryInsuranceFundBalanceSnapshotResponse.md)

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | 
	limit := int64(100) // int64 | Default 100; max 1000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

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
 **symbol** | **string** |  | 
 **limit** | **int64** | Default 100; max 1000 | 

### Return type

[**RecentTradesListResponse**](RecentTradesListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## RpiOrderBook

> RpiOrderBookResponse RpiOrderBook(ctx).Symbol(symbol).Limit(limit).Execute()

RPI Order Book


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | 
	limit := int64(100) // int64 | Default 100; max 1000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.RpiOrderBook(context.Background()).Symbol(symbol).Limit(limit).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.RpiOrderBook``: %v\n", err)
		return
	}

	// response from `RpiOrderBook`: RpiOrderBookResponse
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
 **limit** | **int64** | Default 100; max 1000 | 

### Return type

[**RpiOrderBookResponse**](RpiOrderBookResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SymbolOrderBookTicker

> SymbolOrderBookTickerResponse SymbolOrderBookTicker(ctx).Symbol(symbol).Execute()

Symbol Order Book Ticker


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.SymbolOrderBookTicker(context.Background()).Symbol(symbol).Execute()
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
 **symbol** | **string** |  | 

### Return type

[**SymbolOrderBookTickerResponse**](SymbolOrderBookTickerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SymbolPriceTicker

> SymbolPriceTickerResponse SymbolPriceTicker(ctx).Symbol(symbol).Execute()

Symbol Price Ticker


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.SymbolPriceTicker(context.Background()).Symbol(symbol).Execute()
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
 **symbol** | **string** |  | 

### Return type

[**SymbolPriceTickerResponse**](SymbolPriceTickerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SymbolPriceTickerV2

> SymbolPriceTickerV2Response SymbolPriceTickerV2(ctx).Symbol(symbol).Execute()

Symbol Price Ticker V2


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.SymbolPriceTickerV2(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.SymbolPriceTickerV2``: %v\n", err)
		return
	}

	// response from `SymbolPriceTickerV2`: SymbolPriceTickerV2Response
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

[**SymbolPriceTickerV2Response**](SymbolPriceTickerV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## TakerBuySellVolume

> TakerBuySellVolumeResponse TakerBuySellVolume(ctx).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

Taker Buy/Sell Volume


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | 
	period := models.BasisPeriodParameterPeriod5m // BasisPeriodParameter | \"5m\",\"15m\",\"30m\",\"1h\",\"2h\",\"4h\",\"6h\",\"12h\",\"1d\"
	limit := int64(100) // int64 | Default 100; max 1000 (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.TakerBuySellVolume(context.Background()).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
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
 **symbol** | **string** |  | 
 **period** | [**BasisPeriodParameter**](BasisPeriodParameter.md) | \&quot;5m\&quot;,\&quot;15m\&quot;,\&quot;30m\&quot;,\&quot;1h\&quot;,\&quot;2h\&quot;,\&quot;4h\&quot;,\&quot;6h\&quot;,\&quot;12h\&quot;,\&quot;1d\&quot; | 
 **limit** | **int64** | Default 100; max 1000 | 
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

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

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
 **symbol** | **string** |  | 

### Return type

[**Ticker24hrPriceChangeStatisticsResponse**](Ticker24hrPriceChangeStatisticsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## TopTraderLongShortRatioAccounts

> TopTraderLongShortRatioAccountsResponse TopTraderLongShortRatioAccounts(ctx).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

Top Trader Long/Short Ratio (Accounts)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | 
	period := models.BasisPeriodParameterPeriod5m // BasisPeriodParameter | \"5m\",\"15m\",\"30m\",\"1h\",\"2h\",\"4h\",\"6h\",\"12h\",\"1d\"
	limit := int64(100) // int64 | Default 100; max 1000 (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.TopTraderLongShortRatioAccounts(context.Background()).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
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
 **symbol** | **string** |  | 
 **period** | [**BasisPeriodParameter**](BasisPeriodParameter.md) | \&quot;5m\&quot;,\&quot;15m\&quot;,\&quot;30m\&quot;,\&quot;1h\&quot;,\&quot;2h\&quot;,\&quot;4h\&quot;,\&quot;6h\&quot;,\&quot;12h\&quot;,\&quot;1d\&quot; | 
 **limit** | **int64** | Default 100; max 1000 | 
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

> TopTraderLongShortRatioPositionsResponse TopTraderLongShortRatioPositions(ctx).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()

Top Trader Long/Short Ratio (Positions)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | 
	period := models.BasisPeriodParameterPeriod5m // BasisPeriodParameter | \"5m\",\"15m\",\"30m\",\"1h\",\"2h\",\"4h\",\"6h\",\"12h\",\"1d\"
	limit := int64(100) // int64 | Default 100; max 1000 (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.TopTraderLongShortRatioPositions(context.Background()).Symbol(symbol).Period(period).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
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
 **symbol** | **string** |  | 
 **period** | [**BasisPeriodParameter**](BasisPeriodParameter.md) | \&quot;5m\&quot;,\&quot;15m\&quot;,\&quot;30m\&quot;,\&quot;1h\&quot;,\&quot;2h\&quot;,\&quot;4h\&quot;,\&quot;6h\&quot;,\&quot;12h\&quot;,\&quot;1d\&quot; | 
 **limit** | **int64** | Default 100; max 1000 | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

[**TopTraderLongShortRatioPositionsResponse**](TopTraderLongShortRatioPositionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## TradingSchedule

> TradingScheduleResponse TradingSchedule(ctx).Execute()

Trading Schedule


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.TradingSchedule(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.TradingSchedule``: %v\n", err)
		return
	}

	// response from `TradingSchedule`: TradingScheduleResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

This endpoint does not need any parameter.

### Return type

[**TradingScheduleResponse**](TradingScheduleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

