# \SpotAlgoAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**CancelAlgoOrderSpotAlgo**](SpotAlgoAPI.md#CancelAlgoOrderSpotAlgo) | **Delete** /sapi/v1/algo/spot/order | Cancel Algo Order(TRADE)
[**QueryCurrentAlgoOpenOrdersSpotAlgo**](SpotAlgoAPI.md#QueryCurrentAlgoOpenOrdersSpotAlgo) | **Get** /sapi/v1/algo/spot/openOrders | Query Current Algo Open Orders(USER_DATA)
[**QueryHistoricalAlgoOrdersSpotAlgo**](SpotAlgoAPI.md#QueryHistoricalAlgoOrdersSpotAlgo) | **Get** /sapi/v1/algo/spot/historicalOrders | Query Historical Algo Orders(USER_DATA)
[**QuerySubOrdersSpotAlgo**](SpotAlgoAPI.md#QuerySubOrdersSpotAlgo) | **Get** /sapi/v1/algo/spot/subOrders | Query Sub Orders(USER_DATA)
[**TimeWeightedAveragePriceSpotAlgo**](SpotAlgoAPI.md#TimeWeightedAveragePriceSpotAlgo) | **Post** /sapi/v1/algo/spot/newOrderTwap | Time-Weighted Average Price(Twap) New Order(TRADE)


## CancelAlgoOrderSpotAlgo

> CancelAlgoOrderSpotAlgoResponse CancelAlgoOrderSpotAlgo(ctx).AlgoId(algoId).RecvWindow(recvWindow).Execute()

Cancel Algo Order(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/algo"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	algoId := int64(1) // int64 | eg. 14511
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceAlgoClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.SpotAlgoAPI.CancelAlgoOrderSpotAlgo(context.Background()).AlgoId(algoId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `SpotAlgoAPI.CancelAlgoOrderSpotAlgo``: %v\n", err)
		return
	}

	// response from `CancelAlgoOrderSpotAlgo`: CancelAlgoOrderSpotAlgoResponse
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
 **recvWindow** | **int64** |  | 

### Return type

[**CancelAlgoOrderSpotAlgoResponse**](CancelAlgoOrderSpotAlgoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryCurrentAlgoOpenOrdersSpotAlgo

> QueryCurrentAlgoOpenOrdersSpotAlgoResponse QueryCurrentAlgoOpenOrdersSpotAlgo(ctx).RecvWindow(recvWindow).Execute()

Query Current Algo Open Orders(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/algo"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceAlgoClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.SpotAlgoAPI.QueryCurrentAlgoOpenOrdersSpotAlgo(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `SpotAlgoAPI.QueryCurrentAlgoOpenOrdersSpotAlgo``: %v\n", err)
		return
	}

	// response from `QueryCurrentAlgoOpenOrdersSpotAlgo`: QueryCurrentAlgoOpenOrdersSpotAlgoResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** |  | 

### Return type

[**QueryCurrentAlgoOpenOrdersSpotAlgoResponse**](QueryCurrentAlgoOpenOrdersSpotAlgoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryHistoricalAlgoOrdersSpotAlgo

> QueryHistoricalAlgoOrdersSpotAlgoResponse QueryHistoricalAlgoOrdersSpotAlgo(ctx).Symbol(symbol).Side(side).StartTime(startTime).EndTime(endTime).Page(page).PageSize(pageSize).RecvWindow(recvWindow).Execute()

Query Historical Algo Orders(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/algo"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "BTCUSDT" // string | Trading symbol eg. BTCUSDT (optional)
	side := "BUY" // string | BUY or SELL (optional)
	startTime := int64(1623319461670) // int64 | in milliseconds  eg.1641522717552 (optional)
	endTime := int64(1641782889000) // int64 | in milliseconds  eg.1641522526562 (optional)
	page := int64(1) // int64 | Default is 1 (optional)
	pageSize := int64(100) // int64 | MIN 1, MAX 100; Default 100 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceAlgoClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.SpotAlgoAPI.QueryHistoricalAlgoOrdersSpotAlgo(context.Background()).Symbol(symbol).Side(side).StartTime(startTime).EndTime(endTime).Page(page).PageSize(pageSize).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `SpotAlgoAPI.QueryHistoricalAlgoOrdersSpotAlgo``: %v\n", err)
		return
	}

	// response from `QueryHistoricalAlgoOrdersSpotAlgo`: QueryHistoricalAlgoOrdersSpotAlgoResponse
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
 **side** | **string** | BUY or SELL | 
 **startTime** | **int64** | in milliseconds  eg.1641522717552 | 
 **endTime** | **int64** | in milliseconds  eg.1641522526562 | 
 **page** | **int64** | Default is 1 | 
 **pageSize** | **int64** | MIN 1, MAX 100; Default 100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryHistoricalAlgoOrdersSpotAlgoResponse**](QueryHistoricalAlgoOrdersSpotAlgoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QuerySubOrdersSpotAlgo

> QuerySubOrdersSpotAlgoResponse QuerySubOrdersSpotAlgo(ctx).AlgoId(algoId).Page(page).PageSize(pageSize).RecvWindow(recvWindow).Execute()

Query Sub Orders(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/algo"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	algoId := int64(1) // int64 | eg. 14511
	page := int64(1) // int64 | Default is 1 (optional)
	pageSize := int64(100) // int64 | MIN 1, MAX 100; Default 100 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceAlgoClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.SpotAlgoAPI.QuerySubOrdersSpotAlgo(context.Background()).AlgoId(algoId).Page(page).PageSize(pageSize).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `SpotAlgoAPI.QuerySubOrdersSpotAlgo``: %v\n", err)
		return
	}

	// response from `QuerySubOrdersSpotAlgo`: QuerySubOrdersSpotAlgoResponse
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
 **page** | **int64** | Default is 1 | 
 **pageSize** | **int64** | MIN 1, MAX 100; Default 100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**QuerySubOrdersSpotAlgoResponse**](QuerySubOrdersSpotAlgoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## TimeWeightedAveragePriceSpotAlgo

> TimeWeightedAveragePriceSpotAlgoResponse TimeWeightedAveragePriceSpotAlgo(ctx).Symbol(symbol).Side(side).Quantity(quantity).Duration(duration).ClientAlgoId(clientAlgoId).LimitPrice(limitPrice).Execute()

Time-Weighted Average Price(Twap) New Order(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/algo"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "BTCUSDT" // string | Trading symbol eg. BTCUSDT
	side := "BUY" // string | Trading side ( BUY or SELL )
	quantity := float32(1.0) // float32 | Quantity of base asset; Maximum notional per order is 200k, 2mm or 10mm, depending on symbol. Please reduce your size if you order is above the maximum notional per order.
	duration := int64(5000) // int64 | Duration for TWAP orders in seconds. [300, 86400]
	clientAlgoId := "1" // string | A unique id among Algo orders (length should be 32 characters)， If it is not sent, we will give default value (optional)
	limitPrice := float32(1.0) // float32 | Limit price of the order; If it is not sent, will place order by market price by default (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceAlgoClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.SpotAlgoAPI.TimeWeightedAveragePriceSpotAlgo(context.Background()).Symbol(symbol).Side(side).Quantity(quantity).Duration(duration).ClientAlgoId(clientAlgoId).LimitPrice(limitPrice).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `SpotAlgoAPI.TimeWeightedAveragePriceSpotAlgo``: %v\n", err)
		return
	}

	// response from `TimeWeightedAveragePriceSpotAlgo`: TimeWeightedAveragePriceSpotAlgoResponse
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
 **side** | **string** | Trading side ( BUY or SELL ) | 
 **quantity** | **float32** | Quantity of base asset; Maximum notional per order is 200k, 2mm or 10mm, depending on symbol. Please reduce your size if you order is above the maximum notional per order. | 
 **duration** | **int64** | Duration for TWAP orders in seconds. [300, 86400] | 
 **clientAlgoId** | **string** | A unique id among Algo orders (length should be 32 characters)， If it is not sent, we will give default value | 
 **limitPrice** | **float32** | Limit price of the order; If it is not sent, will place order by market price by default | 

### Return type

[**TimeWeightedAveragePriceSpotAlgoResponse**](TimeWeightedAveragePriceSpotAlgoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

