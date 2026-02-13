# \MarketMakerEndpointsAPI

All URIs are relative to *https://eapi.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AutoCancelAllOpenOrders**](MarketMakerEndpointsAPI.md#AutoCancelAllOpenOrders) | **Post** /eapi/v1/countdownCancelAllHeartBeat | Auto-Cancel All Open Orders (Kill-Switch) Heartbeat (TRADE)
[**GetAutoCancelAllOpenOrders**](MarketMakerEndpointsAPI.md#GetAutoCancelAllOpenOrders) | **Get** /eapi/v1/countdownCancelAll | Get Auto-Cancel All Open Orders (Kill-Switch) Config (TRADE)
[**GetMarketMakerProtectionConfig**](MarketMakerEndpointsAPI.md#GetMarketMakerProtectionConfig) | **Get** /eapi/v1/mmp | Get Market Maker Protection Config (TRADE)
[**ResetMarketMakerProtectionConfig**](MarketMakerEndpointsAPI.md#ResetMarketMakerProtectionConfig) | **Post** /eapi/v1/mmpReset | Reset Market Maker Protection Config (TRADE)
[**SetAutoCancelAllOpenOrders**](MarketMakerEndpointsAPI.md#SetAutoCancelAllOpenOrders) | **Post** /eapi/v1/countdownCancelAll | Set Auto-Cancel All Open Orders (Kill-Switch) Config (TRADE)
[**SetMarketMakerProtectionConfig**](MarketMakerEndpointsAPI.md#SetMarketMakerProtectionConfig) | **Post** /eapi/v1/mmpSet | Set Market Maker Protection Config (TRADE)


## AutoCancelAllOpenOrders

> AutoCancelAllOpenOrdersResponse AutoCancelAllOpenOrders(ctx).Underlyings(underlyings).RecvWindow(recvWindow).Execute()

Auto-Cancel All Open Orders (Kill-Switch) Heartbeat (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	underlyings := "underlyings_example" // string | Option Underlying Symbols, e.g BTCUSDT,ETHUSDT
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketMakerEndpointsAPI.AutoCancelAllOpenOrders(context.Background()).Underlyings(underlyings).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketMakerEndpointsAPI.AutoCancelAllOpenOrders``: %v\n", err)
		return
	}

	// response from `AutoCancelAllOpenOrders`: AutoCancelAllOpenOrdersResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **underlyings** | **string** | Option Underlying Symbols, e.g BTCUSDT,ETHUSDT | 
 **recvWindow** | **int64** |  | 

### Return type

[**AutoCancelAllOpenOrdersResponse**](AutoCancelAllOpenOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetAutoCancelAllOpenOrders

> GetAutoCancelAllOpenOrdersResponse GetAutoCancelAllOpenOrders(ctx).Underlying(underlying).RecvWindow(recvWindow).Execute()

Get Auto-Cancel All Open Orders (Kill-Switch) Config (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	underlying := "underlying_example" // string | underlying, e.g BTCUSDT (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketMakerEndpointsAPI.GetAutoCancelAllOpenOrders(context.Background()).Underlying(underlying).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketMakerEndpointsAPI.GetAutoCancelAllOpenOrders``: %v\n", err)
		return
	}

	// response from `GetAutoCancelAllOpenOrders`: GetAutoCancelAllOpenOrdersResponse
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
 **recvWindow** | **int64** |  | 

### Return type

[**GetAutoCancelAllOpenOrdersResponse**](GetAutoCancelAllOpenOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetMarketMakerProtectionConfig

> GetMarketMakerProtectionConfigResponse GetMarketMakerProtectionConfig(ctx).Underlying(underlying).RecvWindow(recvWindow).Execute()

Get Market Maker Protection Config (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	underlying := "underlying_example" // string | underlying, e.g BTCUSDT (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketMakerEndpointsAPI.GetMarketMakerProtectionConfig(context.Background()).Underlying(underlying).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketMakerEndpointsAPI.GetMarketMakerProtectionConfig``: %v\n", err)
		return
	}

	// response from `GetMarketMakerProtectionConfig`: GetMarketMakerProtectionConfigResponse
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
 **recvWindow** | **int64** |  | 

### Return type

[**GetMarketMakerProtectionConfigResponse**](GetMarketMakerProtectionConfigResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## ResetMarketMakerProtectionConfig

> ResetMarketMakerProtectionConfigResponse ResetMarketMakerProtectionConfig(ctx).Underlying(underlying).RecvWindow(recvWindow).Execute()

Reset Market Maker Protection Config (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	underlying := "underlying_example" // string | underlying, e.g BTCUSDT (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketMakerEndpointsAPI.ResetMarketMakerProtectionConfig(context.Background()).Underlying(underlying).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketMakerEndpointsAPI.ResetMarketMakerProtectionConfig``: %v\n", err)
		return
	}

	// response from `ResetMarketMakerProtectionConfig`: ResetMarketMakerProtectionConfigResponse
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
 **recvWindow** | **int64** |  | 

### Return type

[**ResetMarketMakerProtectionConfigResponse**](ResetMarketMakerProtectionConfigResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SetAutoCancelAllOpenOrders

> SetAutoCancelAllOpenOrdersResponse SetAutoCancelAllOpenOrders(ctx).Underlying(underlying).CountdownTime(countdownTime).RecvWindow(recvWindow).Execute()

Set Auto-Cancel All Open Orders (Kill-Switch) Config (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	underlying := "underlying_example" // string | Option underlying, e.g BTCUSDT
	countdownTime := int64(789) // int64 | Countdown time in milliseconds (ex. 1,000 for 1 second). 0 to disable the timer. Negative values (ex. -10000) are not accepted. Minimum acceptable value is 5,000
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketMakerEndpointsAPI.SetAutoCancelAllOpenOrders(context.Background()).Underlying(underlying).CountdownTime(countdownTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketMakerEndpointsAPI.SetAutoCancelAllOpenOrders``: %v\n", err)
		return
	}

	// response from `SetAutoCancelAllOpenOrders`: SetAutoCancelAllOpenOrdersResponse
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
 **countdownTime** | **int64** | Countdown time in milliseconds (ex. 1,000 for 1 second). 0 to disable the timer. Negative values (ex. -10000) are not accepted. Minimum acceptable value is 5,000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**SetAutoCancelAllOpenOrdersResponse**](SetAutoCancelAllOpenOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SetMarketMakerProtectionConfig

> SetMarketMakerProtectionConfigResponse SetMarketMakerProtectionConfig(ctx).Underlying(underlying).WindowTimeInMilliseconds(windowTimeInMilliseconds).FrozenTimeInMilliseconds(frozenTimeInMilliseconds).QtyLimit(qtyLimit).DeltaLimit(deltaLimit).RecvWindow(recvWindow).Execute()

Set Market Maker Protection Config (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	underlying := "underlying_example" // string | underlying, e.g BTCUSDT (optional)
	windowTimeInMilliseconds := int64(789) // int64 | MMP Interval in milliseconds; Range (0,5000] (optional)
	frozenTimeInMilliseconds := int64(789) // int64 | MMP frozen time in milliseconds, if set to 0 manual reset is required (optional)
	qtyLimit := float32(1.0) // float32 | quantity limit (optional)
	deltaLimit := float32(1.0) // float32 | net delta limit (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketMakerEndpointsAPI.SetMarketMakerProtectionConfig(context.Background()).Underlying(underlying).WindowTimeInMilliseconds(windowTimeInMilliseconds).FrozenTimeInMilliseconds(frozenTimeInMilliseconds).QtyLimit(qtyLimit).DeltaLimit(deltaLimit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketMakerEndpointsAPI.SetMarketMakerProtectionConfig``: %v\n", err)
		return
	}

	// response from `SetMarketMakerProtectionConfig`: SetMarketMakerProtectionConfigResponse
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
 **windowTimeInMilliseconds** | **int64** | MMP Interval in milliseconds; Range (0,5000] | 
 **frozenTimeInMilliseconds** | **int64** | MMP frozen time in milliseconds, if set to 0 manual reset is required | 
 **qtyLimit** | **float32** | quantity limit | 
 **deltaLimit** | **float32** | net delta limit | 
 **recvWindow** | **int64** |  | 

### Return type

[**SetMarketMakerProtectionConfigResponse**](SetMarketMakerProtectionConfigResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

