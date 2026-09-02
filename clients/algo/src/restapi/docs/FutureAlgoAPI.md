# \FutureAlgoAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**CancelAlgoOrderFutureAlgo**](FutureAlgoAPI.md#CancelAlgoOrderFutureAlgo) | **Delete** /sapi/v1/algo/futures/order | Cancel Futures Algo Order (TRADE)
[**QueryCurrentAlgoOpenOrdersFutureAlgo**](FutureAlgoAPI.md#QueryCurrentAlgoOpenOrdersFutureAlgo) | **Get** /sapi/v1/algo/futures/openOrders | Query Current Futures Algo Open Orders (USER_DATA)
[**QueryHistoricalAlgoOrdersFutureAlgo**](FutureAlgoAPI.md#QueryHistoricalAlgoOrdersFutureAlgo) | **Get** /sapi/v1/algo/futures/historicalOrders | Query Historical Futures Algo Orders (USER_DATA)
[**QuerySubOrdersFutureAlgo**](FutureAlgoAPI.md#QuerySubOrdersFutureAlgo) | **Get** /sapi/v1/algo/futures/subOrders | Query Futures Sub Orders (USER_DATA)
[**TimeWeightedAveragePriceFutureAlgo**](FutureAlgoAPI.md#TimeWeightedAveragePriceFutureAlgo) | **Post** /sapi/v1/algo/futures/newOrderTwap | Time-Weighted Futures Average Price (Twap) New Order (TRADE)
[**VolumeParticipationFutureAlgo**](FutureAlgoAPI.md#VolumeParticipationFutureAlgo) | **Post** /sapi/v1/algo/futures/newOrderVp | Volume Participation (VP) New Order (TRADE)


## CancelAlgoOrderFutureAlgo

> CancelAlgoOrderFutureAlgoResponse CancelAlgoOrderFutureAlgo(ctx).AlgoId(algoId).ClientAlgoId(clientAlgoId).RecvWindow(recvWindow).Execute()

Cancel Futures Algo Order (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/algo"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	algoId := int64(1) // int64 | eg. 14511 (optional)
	clientAlgoId := "65ce1630101a480b85915d7e11fd5078" // string | eg. \"65ce1630101a480b85915d7e11fd5078\" (optional)
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceAlgoClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FutureAlgoAPI.CancelAlgoOrderFutureAlgo(context.Background()).AlgoId(algoId).ClientAlgoId(clientAlgoId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FutureAlgoAPI.CancelAlgoOrderFutureAlgo``: %v\n", err)
		return
	}

	// response from `CancelAlgoOrderFutureAlgo`: CancelAlgoOrderFutureAlgoResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **algoId** | **int64** | eg. 14511 | 
 **clientAlgoId** | **string** | eg. \&quot;65ce1630101a480b85915d7e11fd5078\&quot; | 
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**CancelAlgoOrderFutureAlgoResponse**](CancelAlgoOrderFutureAlgoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryCurrentAlgoOpenOrdersFutureAlgo

> QueryCurrentAlgoOpenOrdersFutureAlgoResponse QueryCurrentAlgoOpenOrdersFutureAlgo(ctx).RecvWindow(recvWindow).Execute()

Query Current Futures Algo Open Orders (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/algo"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceAlgoClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FutureAlgoAPI.QueryCurrentAlgoOpenOrdersFutureAlgo(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FutureAlgoAPI.QueryCurrentAlgoOpenOrdersFutureAlgo``: %v\n", err)
		return
	}

	// response from `QueryCurrentAlgoOpenOrdersFutureAlgo`: QueryCurrentAlgoOpenOrdersFutureAlgoResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**QueryCurrentAlgoOpenOrdersFutureAlgoResponse**](QueryCurrentAlgoOpenOrdersFutureAlgoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryHistoricalAlgoOrdersFutureAlgo

> QueryHistoricalAlgoOrdersFutureAlgoResponse QueryHistoricalAlgoOrdersFutureAlgo(ctx).Symbol(symbol).Side(side).StartTime(startTime).EndTime(endTime).Page(page).PageSize(pageSize).RecvWindow(recvWindow).Execute()

Query Historical Futures Algo Orders (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/algo"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSDT" // string | Trading symbol eg. BTCUSDT (optional)
	side := models.QueryHistoricalAlgoOrdersFutureAlgoSideParameterBuy // QueryHistoricalAlgoOrdersFutureAlgoSideParameter | BUY or SELL (optional)
	startTime := int64(1623319461670) // int64 | in milliseconds  eg.1641522717552 (optional)
	endTime := int64(1641782889000) // int64 | in milliseconds  eg.1641522526562 (optional)
	page := int64(1) // int64 | Page number (optional)
	pageSize := int64(100) // int64 | Records per page (optional)
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceAlgoClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FutureAlgoAPI.QueryHistoricalAlgoOrdersFutureAlgo(context.Background()).Symbol(symbol).Side(side).StartTime(startTime).EndTime(endTime).Page(page).PageSize(pageSize).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FutureAlgoAPI.QueryHistoricalAlgoOrdersFutureAlgo``: %v\n", err)
		return
	}

	// response from `QueryHistoricalAlgoOrdersFutureAlgo`: QueryHistoricalAlgoOrdersFutureAlgoResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol eg. BTCUSDT | 
 **side** | [**QueryHistoricalAlgoOrdersFutureAlgoSideParameter**](QueryHistoricalAlgoOrdersFutureAlgoSideParameter.md) | BUY or SELL | 
 **startTime** | **int64** | in milliseconds  eg.1641522717552 | 
 **endTime** | **int64** | in milliseconds  eg.1641522526562 | 
 **page** | **int64** | Page number | 
 **pageSize** | **int64** | Records per page | 
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**QueryHistoricalAlgoOrdersFutureAlgoResponse**](QueryHistoricalAlgoOrdersFutureAlgoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QuerySubOrdersFutureAlgo

> QuerySubOrdersFutureAlgoResponse QuerySubOrdersFutureAlgo(ctx).AlgoId(algoId).Page(page).PageSize(pageSize).RecvWindow(recvWindow).Execute()

Query Futures Sub Orders (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/algo"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	algoId := int64(1) // int64 | eg. 14511
	page := int64(1) // int64 | Page number (optional)
	pageSize := int64(100) // int64 | Records per page (optional)
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceAlgoClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FutureAlgoAPI.QuerySubOrdersFutureAlgo(context.Background()).AlgoId(algoId).Page(page).PageSize(pageSize).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FutureAlgoAPI.QuerySubOrdersFutureAlgo``: %v\n", err)
		return
	}

	// response from `QuerySubOrdersFutureAlgo`: QuerySubOrdersFutureAlgoResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **algoId** | **int64** | eg. 14511 | 
 **page** | **int64** | Page number | 
 **pageSize** | **int64** | Records per page | 
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**QuerySubOrdersFutureAlgoResponse**](QuerySubOrdersFutureAlgoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## TimeWeightedAveragePriceFutureAlgo

> TimeWeightedAveragePriceFutureAlgoResponse TimeWeightedAveragePriceFutureAlgo(ctx).Symbol(symbol).Side(side).Quantity(quantity).Duration(duration).PositionSide(positionSide).ClientAlgoId(clientAlgoId).ReduceOnly(reduceOnly).LimitPrice(limitPrice).RecvWindow(recvWindow).Execute()

Time-Weighted Futures Average Price (Twap) New Order (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/algo"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSDT" // string | Trading symbol eg. BTCUSDT
	side := models.QueryHistoricalAlgoOrdersFutureAlgoSideParameterBuy // QueryHistoricalAlgoOrdersFutureAlgoSideParameter | Trading side ( BUY or SELL )
	quantity := float64(1) // float64 | Quantity of base asset; The notional (`quantity` * `mark price(base asset)`) must be more than the equivalent of 1,000 USDT and less than the equivalent of 1,000,000 USDT
	duration := int64(5000) // int64 | Duration for TWAP orders in seconds
	positionSide := models.TimeWeightedAveragePriceFutureAlgoPositionSideParameterBoth // TimeWeightedAveragePriceFutureAlgoPositionSideParameter | Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent in Hedge Mode. (optional)
	clientAlgoId := "1" // string | A unique id among Algo orders (length should be 32 characters)， If it is not sent, we will give default value (optional)
	reduceOnly := false // bool | \"true\" or \"false\". Default \"false\"; Cannot be sent in Hedge Mode; Cannot be sent when you open a position (optional)
	limitPrice := float64(1) // float64 | Limit price of the order; If it is not sent, will place order by market price by default (optional)
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceAlgoClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FutureAlgoAPI.TimeWeightedAveragePriceFutureAlgo(context.Background()).Symbol(symbol).Side(side).Quantity(quantity).Duration(duration).PositionSide(positionSide).ClientAlgoId(clientAlgoId).ReduceOnly(reduceOnly).LimitPrice(limitPrice).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FutureAlgoAPI.TimeWeightedAveragePriceFutureAlgo``: %v\n", err)
		return
	}

	// response from `TimeWeightedAveragePriceFutureAlgo`: TimeWeightedAveragePriceFutureAlgoResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol eg. BTCUSDT | 
 **side** | [**QueryHistoricalAlgoOrdersFutureAlgoSideParameter**](QueryHistoricalAlgoOrdersFutureAlgoSideParameter.md) | Trading side ( BUY or SELL ) | 
 **quantity** | **float64** | Quantity of base asset; The notional (&#x60;quantity&#x60; * &#x60;mark price(base asset)&#x60;) must be more than the equivalent of 1,000 USDT and less than the equivalent of 1,000,000 USDT | 
 **duration** | **int64** | Duration for TWAP orders in seconds | 
 **positionSide** | [**TimeWeightedAveragePriceFutureAlgoPositionSideParameter**](TimeWeightedAveragePriceFutureAlgoPositionSideParameter.md) | Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent in Hedge Mode. | 
 **clientAlgoId** | **string** | A unique id among Algo orders (length should be 32 characters)， If it is not sent, we will give default value | 
 **reduceOnly** | **bool** | \&quot;true\&quot; or \&quot;false\&quot;. Default \&quot;false\&quot;; Cannot be sent in Hedge Mode; Cannot be sent when you open a position | 
 **limitPrice** | **float64** | Limit price of the order; If it is not sent, will place order by market price by default | 
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**TimeWeightedAveragePriceFutureAlgoResponse**](TimeWeightedAveragePriceFutureAlgoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## VolumeParticipationFutureAlgo

> VolumeParticipationFutureAlgoResponse VolumeParticipationFutureAlgo(ctx).Symbol(symbol).Side(side).Quantity(quantity).Urgency(urgency).PositionSide(positionSide).ClientAlgoId(clientAlgoId).ReduceOnly(reduceOnly).LimitPrice(limitPrice).RecvWindow(recvWindow).Execute()

Volume Participation (VP) New Order (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/algo"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSDT" // string | Trading symbol eg. BTCUSDT
	side := models.QueryHistoricalAlgoOrdersFutureAlgoSideParameterBuy // QueryHistoricalAlgoOrdersFutureAlgoSideParameter | Trading side ( BUY or SELL )
	quantity := float64(1) // float64 | Quantity of base asset; The notional (`quantity` * `mark price(base asset)`) must be more than the equivalent of 10,000 USDT and less than the equivalent of 1,000,000 USDT
	urgency := models.VolumeParticipationFutureAlgoUrgencyParameterLow // VolumeParticipationFutureAlgoUrgencyParameter | Represent the relative speed of the current execution; ENUM: LOW, MEDIUM, HIGH
	positionSide := models.TimeWeightedAveragePriceFutureAlgoPositionSideParameterBoth // TimeWeightedAveragePriceFutureAlgoPositionSideParameter | Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent in Hedge Mode. (optional)
	clientAlgoId := "1" // string | A unique id among Algo orders (length should be 32 characters)， If it is not sent, we will give default value (optional)
	reduceOnly := false // bool | \"true\" or \"false\". Default \"false\"; Cannot be sent in Hedge Mode; Cannot be sent when you open a position (optional)
	limitPrice := float64(1) // float64 | Limit price of the order; If it is not sent, will place order by market price by default (optional)
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceAlgoClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FutureAlgoAPI.VolumeParticipationFutureAlgo(context.Background()).Symbol(symbol).Side(side).Quantity(quantity).Urgency(urgency).PositionSide(positionSide).ClientAlgoId(clientAlgoId).ReduceOnly(reduceOnly).LimitPrice(limitPrice).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FutureAlgoAPI.VolumeParticipationFutureAlgo``: %v\n", err)
		return
	}

	// response from `VolumeParticipationFutureAlgo`: VolumeParticipationFutureAlgoResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Trading symbol eg. BTCUSDT | 
 **side** | [**QueryHistoricalAlgoOrdersFutureAlgoSideParameter**](QueryHistoricalAlgoOrdersFutureAlgoSideParameter.md) | Trading side ( BUY or SELL ) | 
 **quantity** | **float64** | Quantity of base asset; The notional (&#x60;quantity&#x60; * &#x60;mark price(base asset)&#x60;) must be more than the equivalent of 10,000 USDT and less than the equivalent of 1,000,000 USDT | 
 **urgency** | [**VolumeParticipationFutureAlgoUrgencyParameter**](VolumeParticipationFutureAlgoUrgencyParameter.md) | Represent the relative speed of the current execution; ENUM: LOW, MEDIUM, HIGH | 
 **positionSide** | [**TimeWeightedAveragePriceFutureAlgoPositionSideParameter**](TimeWeightedAveragePriceFutureAlgoPositionSideParameter.md) | Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent in Hedge Mode. | 
 **clientAlgoId** | **string** | A unique id among Algo orders (length should be 32 characters)， If it is not sent, we will give default value | 
 **reduceOnly** | **bool** | \&quot;true\&quot; or \&quot;false\&quot;. Default \&quot;false\&quot;; Cannot be sent in Hedge Mode; Cannot be sent when you open a position | 
 **limitPrice** | **float64** | Limit price of the order; If it is not sent, will place order by market price by default | 
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**VolumeParticipationFutureAlgoResponse**](VolumeParticipationFutureAlgoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

