# \TradeAPI

All URIs are relative to *https://fapi.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AccountTradeList**](TradeAPI.md#AccountTradeList) | **Get** /fapi/v1/userTrades | Account Trade List (USER_DATA)
[**AllOrders**](TradeAPI.md#AllOrders) | **Get** /fapi/v1/allOrders | All Orders (USER_DATA)
[**AutoCancelAllOpenOrders**](TradeAPI.md#AutoCancelAllOpenOrders) | **Post** /fapi/v1/countdownCancelAll | Auto-Cancel All Open Orders (TRADE)
[**CancelAlgoOrder**](TradeAPI.md#CancelAlgoOrder) | **Delete** /fapi/v1/algoOrder | Cancel Algo Order (TRADE)
[**CancelAllAlgoOpenOrders**](TradeAPI.md#CancelAllAlgoOpenOrders) | **Delete** /fapi/v1/algoOpenOrders | Cancel All Algo Open Orders (TRADE)
[**CancelAllOpenOrders**](TradeAPI.md#CancelAllOpenOrders) | **Delete** /fapi/v1/allOpenOrders | Cancel All Open Orders (TRADE)
[**CancelMultipleOrders**](TradeAPI.md#CancelMultipleOrders) | **Delete** /fapi/v1/batchOrders | Cancel Multiple Orders (TRADE)
[**CancelOrder**](TradeAPI.md#CancelOrder) | **Delete** /fapi/v1/order | Cancel Order (TRADE)
[**ChangeInitialLeverage**](TradeAPI.md#ChangeInitialLeverage) | **Post** /fapi/v1/leverage | Change Initial Leverage (TRADE)
[**ChangeMarginType**](TradeAPI.md#ChangeMarginType) | **Post** /fapi/v1/marginType | Change Margin Type (TRADE)
[**ChangeMultiAssetsMode**](TradeAPI.md#ChangeMultiAssetsMode) | **Post** /fapi/v1/multiAssetsMargin | Change Multi-Assets Mode (TRADE)
[**ChangePositionMode**](TradeAPI.md#ChangePositionMode) | **Post** /fapi/v1/positionSide/dual | Change Position Mode (TRADE)
[**CurrentAllAlgoOpenOrders**](TradeAPI.md#CurrentAllAlgoOpenOrders) | **Get** /fapi/v1/openAlgoOrders | Current All Algo Open Orders (USER_DATA)
[**CurrentAllOpenOrders**](TradeAPI.md#CurrentAllOpenOrders) | **Get** /fapi/v1/openOrders | Current All Open Orders (USER_DATA)
[**FuturesTradfiPerpsContract**](TradeAPI.md#FuturesTradfiPerpsContract) | **Post** /fapi/v1/stock/contract | Futures TradFi Perps Contract (USER_DATA)
[**GetOrderModifyHistory**](TradeAPI.md#GetOrderModifyHistory) | **Get** /fapi/v1/orderAmendment | Get Order Modify History (USER_DATA)
[**GetPositionMarginChangeHistory**](TradeAPI.md#GetPositionMarginChangeHistory) | **Get** /fapi/v1/positionMargin/history | Get Position Margin Change History (TRADE)
[**ModifyIsolatedPositionMargin**](TradeAPI.md#ModifyIsolatedPositionMargin) | **Post** /fapi/v1/positionMargin | Modify Isolated Position Margin (TRADE)
[**ModifyMultipleOrders**](TradeAPI.md#ModifyMultipleOrders) | **Put** /fapi/v1/batchOrders | Modify Multiple Orders (TRADE)
[**ModifyOrder**](TradeAPI.md#ModifyOrder) | **Put** /fapi/v1/order | Modify Order (TRADE)
[**NewAlgoOrder**](TradeAPI.md#NewAlgoOrder) | **Post** /fapi/v1/algoOrder | New Algo Order (TRADE)
[**NewOrder**](TradeAPI.md#NewOrder) | **Post** /fapi/v1/order | New Order (TRADE)
[**PlaceMultipleOrders**](TradeAPI.md#PlaceMultipleOrders) | **Post** /fapi/v1/batchOrders | Place Multiple Orders (TRADE)
[**PositionAdlQuantileEstimation**](TradeAPI.md#PositionAdlQuantileEstimation) | **Get** /fapi/v1/adlQuantile | Position ADL Quantile Estimation (USER_DATA)
[**PositionInformationV2**](TradeAPI.md#PositionInformationV2) | **Get** /fapi/v2/positionRisk | Position Information V2 (USER_DATA)
[**PositionInformationV3**](TradeAPI.md#PositionInformationV3) | **Get** /fapi/v3/positionRisk | Position Information V3 (USER_DATA)
[**QueryAlgoOrder**](TradeAPI.md#QueryAlgoOrder) | **Get** /fapi/v1/algoOrder | Query Algo Order (USER_DATA)
[**QueryAllAlgoOrders**](TradeAPI.md#QueryAllAlgoOrders) | **Get** /fapi/v1/allAlgoOrders | Query All Algo Orders (USER_DATA)
[**QueryCurrentOpenOrder**](TradeAPI.md#QueryCurrentOpenOrder) | **Get** /fapi/v1/openOrder | Query Current Open Order (USER_DATA)
[**QueryOrder**](TradeAPI.md#QueryOrder) | **Get** /fapi/v1/order | Query Order (USER_DATA)
[**TestOrder**](TradeAPI.md#TestOrder) | **Post** /fapi/v1/order/test | Test Order (TRADE)
[**UsersForceOrders**](TradeAPI.md#UsersForceOrders) | **Get** /fapi/v1/forceOrders | User&#39;s Force Orders (USER_DATA)


## AccountTradeList

> AccountTradeListResponse AccountTradeList(ctx).Symbol(symbol).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()

Account Trade List (USER_DATA)


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
	symbol := "BTCUSDT" // string | 
	orderId := int64(25851813) // int64 | Must be used together with parameter `symbol`. (optional)
	startTime := int64(1623319461670) // int64 | Start time (optional)
	endTime := int64(1641782889000) // int64 | End time (optional)
	fromId := int64(1) // int64 | Trade id to fetch from. Default gets most recent trades. (optional)
	limit := int64(50) // int64 |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.AccountTradeList(context.Background()).Symbol(symbol).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.AccountTradeList``: %v\n", err)
		return
	}

	// response from `AccountTradeList`: AccountTradeListResponse
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
 **orderId** | **int64** | Must be used together with parameter &#x60;symbol&#x60;. | 
 **startTime** | **int64** | Start time | 
 **endTime** | **int64** | End time | 
 **fromId** | **int64** | Trade id to fetch from. Default gets most recent trades. | 
 **limit** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**AccountTradeListResponse**](AccountTradeListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## AllOrders

> AllOrdersResponse AllOrders(ctx).Symbol(symbol).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

All Orders (USER_DATA)


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
	symbol := "BTCUSDT" // string | 
	orderId := int64(1917641) // int64 |  (optional)
	startTime := int64(1623319461670) // int64 | Start time (optional)
	endTime := int64(1641782889000) // int64 | End time (optional)
	limit := int64(50) // int64 |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.AllOrders(context.Background()).Symbol(symbol).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.AllOrders``: %v\n", err)
		return
	}

	// response from `AllOrders`: AllOrdersResponse
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
 **orderId** | **int64** |  | 
 **startTime** | **int64** | Start time | 
 **endTime** | **int64** | End time | 
 **limit** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**AllOrdersResponse**](AllOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## AutoCancelAllOpenOrders

> AutoCancelAllOpenOrdersResponse AutoCancelAllOpenOrders(ctx).Symbol(symbol).CountdownTime(countdownTime).RecvWindow(recvWindow).Execute()

Auto-Cancel All Open Orders (TRADE)


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
	symbol := "BTCUSDT" // string | 
	countdownTime := int64(1000) // int64 | Countdown in milliseconds. `1000` means 1 second; `0` disables countdown cancel-all.
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.AutoCancelAllOpenOrders(context.Background()).Symbol(symbol).CountdownTime(countdownTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.AutoCancelAllOpenOrders``: %v\n", err)
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
 **symbol** | **string** |  | 
 **countdownTime** | **int64** | Countdown in milliseconds. &#x60;1000&#x60; means 1 second; &#x60;0&#x60; disables countdown cancel-all. | 
 **recvWindow** | **int64** |  | 

### Return type

[**AutoCancelAllOpenOrdersResponse**](AutoCancelAllOpenOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CancelAlgoOrder

> CancelAlgoOrderResponse CancelAlgoOrder(ctx).AlgoId(algoId).ClientAlgoId(clientAlgoId).RecvWindow(recvWindow).Execute()

Cancel Algo Order (TRADE)


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
	algoId := int64(2146760) // int64 |  (optional)
	clientAlgoId := "6B2I9XVcJpCjqPAJ4YoFX7" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CancelAlgoOrder(context.Background()).AlgoId(algoId).ClientAlgoId(clientAlgoId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CancelAlgoOrder``: %v\n", err)
		return
	}

	// response from `CancelAlgoOrder`: CancelAlgoOrderResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **algoId** | **int64** |  | 
 **clientAlgoId** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CancelAlgoOrderResponse**](CancelAlgoOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CancelAllAlgoOpenOrders

> CancelAllAlgoOpenOrdersResponse CancelAllAlgoOpenOrders(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Cancel All Algo Open Orders (TRADE)


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
	symbol := "BTCUSDT" // string | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CancelAllAlgoOpenOrders(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CancelAllAlgoOpenOrders``: %v\n", err)
		return
	}

	// response from `CancelAllAlgoOpenOrders`: CancelAllAlgoOpenOrdersResponse
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
 **recvWindow** | **int64** |  | 

### Return type

[**CancelAllAlgoOpenOrdersResponse**](CancelAllAlgoOpenOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CancelAllOpenOrders

> CancelAllOpenOrdersResponse CancelAllOpenOrders(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Cancel All Open Orders (TRADE)


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
	symbol := "BTCUSDT" // string | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CancelAllOpenOrders(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CancelAllOpenOrders``: %v\n", err)
		return
	}

	// response from `CancelAllOpenOrders`: CancelAllOpenOrdersResponse
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
 **recvWindow** | **int64** |  | 

### Return type

[**CancelAllOpenOrdersResponse**](CancelAllOpenOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CancelMultipleOrders

> CancelMultipleOrdersResponse CancelMultipleOrders(ctx).Symbol(symbol).OrderIdList(orderIdList).OrigClientOrderIdList(origClientOrderIdList).RecvWindow(recvWindow).Execute()

Cancel Multiple Orders (TRADE)


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
	symbol := "BTCUSDT" // string | 
	orderIdList := []int64{int64(1234567)} // []int64 |  (optional)
	origClientOrderIdList := []string{"my_id_1"} // []string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CancelMultipleOrders(context.Background()).Symbol(symbol).OrderIdList(orderIdList).OrigClientOrderIdList(origClientOrderIdList).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CancelMultipleOrders``: %v\n", err)
		return
	}

	// response from `CancelMultipleOrders`: CancelMultipleOrdersResponse
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
 **orderIdList** | **[]int64** |  | 
 **origClientOrderIdList** | **[]string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CancelMultipleOrdersResponse**](CancelMultipleOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CancelOrder

> CancelOrderResponse CancelOrder(ctx).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Cancel Order (TRADE)


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
	symbol := "BTCUSDT" // string | 
	orderId := int64(283194212) // int64 |  (optional)
	origClientOrderId := "myOrder1" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CancelOrder(context.Background()).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CancelOrder``: %v\n", err)
		return
	}

	// response from `CancelOrder`: CancelOrderResponse
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
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CancelOrderResponse**](CancelOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## ChangeInitialLeverage

> ChangeInitialLeverageResponse ChangeInitialLeverage(ctx).Symbol(symbol).Leverage(leverage).RecvWindow(recvWindow).Execute()

Change Initial Leverage (TRADE)


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
	symbol := "BTCUSDT" // string | 
	leverage := int64(1) // int64 | target initial leverage
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.ChangeInitialLeverage(context.Background()).Symbol(symbol).Leverage(leverage).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.ChangeInitialLeverage``: %v\n", err)
		return
	}

	// response from `ChangeInitialLeverage`: ChangeInitialLeverageResponse
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
 **leverage** | **int64** | target initial leverage | 
 **recvWindow** | **int64** |  | 

### Return type

[**ChangeInitialLeverageResponse**](ChangeInitialLeverageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## ChangeMarginType

> ChangeMarginTypeResponse ChangeMarginType(ctx).Symbol(symbol).MarginType(marginType).RecvWindow(recvWindow).Execute()

Change Margin Type (TRADE)


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
	symbol := "BTCUSDT" // string | 
	marginType := models.ChangeMarginTypeMarginTypeParameterIsolated // ChangeMarginTypeMarginTypeParameter | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.ChangeMarginType(context.Background()).Symbol(symbol).MarginType(marginType).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.ChangeMarginType``: %v\n", err)
		return
	}

	// response from `ChangeMarginType`: ChangeMarginTypeResponse
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
 **marginType** | [**ChangeMarginTypeMarginTypeParameter**](ChangeMarginTypeMarginTypeParameter.md) |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**ChangeMarginTypeResponse**](ChangeMarginTypeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## ChangeMultiAssetsMode

> ChangeMultiAssetsModeResponse ChangeMultiAssetsMode(ctx).MultiAssetsMargin(multiAssetsMargin).RecvWindow(recvWindow).Execute()

Change Multi-Assets Mode (TRADE)


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
	multiAssetsMargin := "true" // string | \"true\": Multi-Assets Mode; \"false\": Single-Asset Mode
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.ChangeMultiAssetsMode(context.Background()).MultiAssetsMargin(multiAssetsMargin).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.ChangeMultiAssetsMode``: %v\n", err)
		return
	}

	// response from `ChangeMultiAssetsMode`: ChangeMultiAssetsModeResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **multiAssetsMargin** | **string** | \&quot;true\&quot;: Multi-Assets Mode; \&quot;false\&quot;: Single-Asset Mode | 
 **recvWindow** | **int64** |  | 

### Return type

[**ChangeMultiAssetsModeResponse**](ChangeMultiAssetsModeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## ChangePositionMode

> ChangePositionModeResponse ChangePositionMode(ctx).DualSidePosition(dualSidePosition).RecvWindow(recvWindow).Execute()

Change Position Mode (TRADE)


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
	dualSidePosition := "true" // string | \"true\": Hedge Mode; \"false\": One-way Mode
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.ChangePositionMode(context.Background()).DualSidePosition(dualSidePosition).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.ChangePositionMode``: %v\n", err)
		return
	}

	// response from `ChangePositionMode`: ChangePositionModeResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **dualSidePosition** | **string** | \&quot;true\&quot;: Hedge Mode; \&quot;false\&quot;: One-way Mode | 
 **recvWindow** | **int64** |  | 

### Return type

[**ChangePositionModeResponse**](ChangePositionModeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CurrentAllAlgoOpenOrders

> CurrentAllAlgoOpenOrdersResponse CurrentAllAlgoOpenOrders(ctx).AlgoType(algoType).Symbol(symbol).AlgoId(algoId).RecvWindow(recvWindow).Execute()

Current All Algo Open Orders (USER_DATA)


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
	algoType := "CONDITIONAL" // string |  (optional)
	symbol := "BTCUSDT" // string |  (optional)
	algoId := int64(2148627) // int64 |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CurrentAllAlgoOpenOrders(context.Background()).AlgoType(algoType).Symbol(symbol).AlgoId(algoId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CurrentAllAlgoOpenOrders``: %v\n", err)
		return
	}

	// response from `CurrentAllAlgoOpenOrders`: CurrentAllAlgoOpenOrdersResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **algoType** | **string** |  | 
 **symbol** | **string** |  | 
 **algoId** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CurrentAllAlgoOpenOrdersResponse**](CurrentAllAlgoOpenOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CurrentAllOpenOrders

> CurrentAllOpenOrdersResponse CurrentAllOpenOrders(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Current All Open Orders (USER_DATA)


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
	symbol := "BTCUSDT" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CurrentAllOpenOrders(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CurrentAllOpenOrders``: %v\n", err)
		return
	}

	// response from `CurrentAllOpenOrders`: CurrentAllOpenOrdersResponse
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
 **recvWindow** | **int64** |  | 

### Return type

[**CurrentAllOpenOrdersResponse**](CurrentAllOpenOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## FuturesTradfiPerpsContract

> FuturesTradfiPerpsContractResponse FuturesTradfiPerpsContract(ctx).RecvWindow(recvWindow).Execute()

Futures TradFi Perps Contract (USER_DATA)


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
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.FuturesTradfiPerpsContract(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.FuturesTradfiPerpsContract``: %v\n", err)
		return
	}

	// response from `FuturesTradfiPerpsContract`: FuturesTradfiPerpsContractResponse
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

[**FuturesTradfiPerpsContractResponse**](FuturesTradfiPerpsContractResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetOrderModifyHistory

> GetOrderModifyHistoryResponse GetOrderModifyHistory(ctx).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Get Order Modify History (USER_DATA)


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
	symbol := "BTCUSDT" // string | 
	orderId := int64(20072994037) // int64 |  (optional)
	origClientOrderId := "LJ9R4QZDihCaS8UAOOLpgW" // string |  (optional)
	startTime := int64(1623319461670) // int64 | Timestamp in ms to get modification history from INCLUSIVE (optional)
	endTime := int64(1641782889000) // int64 | Timestamp in ms to get modification history until INCLUSIVE (optional)
	limit := int64(50) // int64 |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.GetOrderModifyHistory(context.Background()).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.GetOrderModifyHistory``: %v\n", err)
		return
	}

	// response from `GetOrderModifyHistory`: GetOrderModifyHistoryResponse
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
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | 
 **startTime** | **int64** | Timestamp in ms to get modification history from INCLUSIVE | 
 **endTime** | **int64** | Timestamp in ms to get modification history until INCLUSIVE | 
 **limit** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetOrderModifyHistoryResponse**](GetOrderModifyHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetPositionMarginChangeHistory

> GetPositionMarginChangeHistoryResponse GetPositionMarginChangeHistory(ctx).Symbol(symbol).Type(type_).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Get Position Margin Change History (TRADE)


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
	symbol := "BTCUSDT" // string | 
	type_ := "1" // string | 1: Add position margin，2: Reduce position margin (optional)
	startTime := int64(1623319461670) // int64 | Start time (optional)
	endTime := int64(1641782889000) // int64 | time if not pass (optional)
	limit := int64(50) // int64 |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.GetPositionMarginChangeHistory(context.Background()).Symbol(symbol).Type(type_).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.GetPositionMarginChangeHistory``: %v\n", err)
		return
	}

	// response from `GetPositionMarginChangeHistory`: GetPositionMarginChangeHistoryResponse
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
 **type_** | **string** | 1: Add position margin，2: Reduce position margin | 
 **startTime** | **int64** | Start time | 
 **endTime** | **int64** | time if not pass | 
 **limit** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetPositionMarginChangeHistoryResponse**](GetPositionMarginChangeHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## ModifyIsolatedPositionMargin

> ModifyIsolatedPositionMarginResponse ModifyIsolatedPositionMargin(ctx).Symbol(symbol).Amount(amount).Type(type_).PositionSide(positionSide).RecvWindow(recvWindow).Execute()

Modify Isolated Position Margin (TRADE)


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
	symbol := "BTCUSDT" // string | 
	amount := float32(1.0) // float32 | Margin asset
	type_ := int32(1) // int32 | 1: Add position margin，2: Reduce position margin
	positionSide := "BOTH" // string | Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent with Hedge Mode. (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.ModifyIsolatedPositionMargin(context.Background()).Symbol(symbol).Amount(amount).Type(type_).PositionSide(positionSide).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.ModifyIsolatedPositionMargin``: %v\n", err)
		return
	}

	// response from `ModifyIsolatedPositionMargin`: ModifyIsolatedPositionMarginResponse
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
 **amount** | **float32** | Margin asset | 
 **type_** | **int32** | 1: Add position margin，2: Reduce position margin | 
 **positionSide** | **string** | Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent with Hedge Mode. | 
 **recvWindow** | **int64** |  | 

### Return type

[**ModifyIsolatedPositionMarginResponse**](ModifyIsolatedPositionMarginResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## ModifyMultipleOrders

> ModifyMultipleOrdersResponse ModifyMultipleOrders(ctx).BatchOrders(batchOrders).RecvWindow(recvWindow).Execute()

Modify Multiple Orders (TRADE)


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
	batchOrders := []models.ModifyMultipleOrdersBatchOrdersParameterInner{*models.NewModifyMultipleOrdersBatchOrdersParameterInner()} // []ModifyMultipleOrdersBatchOrdersParameterInner | order list. Max 5 orders
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.ModifyMultipleOrders(context.Background()).BatchOrders(batchOrders).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.ModifyMultipleOrders``: %v\n", err)
		return
	}

	// response from `ModifyMultipleOrders`: ModifyMultipleOrdersResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **batchOrders** | [**[]ModifyMultipleOrdersBatchOrdersParameterInner**](ModifyMultipleOrdersBatchOrdersParameterInner.md) | order list. Max 5 orders | 
 **recvWindow** | **int64** |  | 

### Return type

[**ModifyMultipleOrdersResponse**](ModifyMultipleOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## ModifyOrder

> ModifyOrderResponse ModifyOrder(ctx).Symbol(symbol).Side(side).Quantity(quantity).Price(price).OrderId(orderId).OrigClientOrderId(origClientOrderId).PriceMatch(priceMatch).ModifyId(modifyId).RecvWindow(recvWindow).Execute()

Modify Order (TRADE)


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
	symbol := "BTCUSDT" // string | 
	side := models.NewAlgoOrderSideParameterBuy // NewAlgoOrderSideParameter | 
	quantity := float32(1.0) // float32 | Order quantity, cannot be sent with `closePosition=true`
	price := float32(30005) // float32 | 
	orderId := int64(20072994037) // int64 |  (optional)
	origClientOrderId := "LJ9R4QZDihCaS8UAOOLpgW" // string |  (optional)
	priceMatch := models.NewAlgoOrderPriceMatchParameterOpponent // NewAlgoOrderPriceMatchParameter | only avaliable for `LIMIT`/`STOP`/`TAKE_PROFIT` order; Can't be passed together with `price` (optional)
	modifyId := int64(1) // int64 | User-defined modification identifier, returned as-is in the response. Optional; not validated for uniqueness. (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.ModifyOrder(context.Background()).Symbol(symbol).Side(side).Quantity(quantity).Price(price).OrderId(orderId).OrigClientOrderId(origClientOrderId).PriceMatch(priceMatch).ModifyId(modifyId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.ModifyOrder``: %v\n", err)
		return
	}

	// response from `ModifyOrder`: ModifyOrderResponse
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
 **side** | [**NewAlgoOrderSideParameter**](NewAlgoOrderSideParameter.md) |  | 
 **quantity** | **float32** | Order quantity, cannot be sent with &#x60;closePosition&#x3D;true&#x60; | 
 **price** | **float32** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | 
 **priceMatch** | [**NewAlgoOrderPriceMatchParameter**](NewAlgoOrderPriceMatchParameter.md) | only avaliable for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; Can&#39;t be passed together with &#x60;price&#x60; | 
 **modifyId** | **int64** | User-defined modification identifier, returned as-is in the response. Optional; not validated for uniqueness. | 
 **recvWindow** | **int64** |  | 

### Return type

[**ModifyOrderResponse**](ModifyOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## NewAlgoOrder

> NewAlgoOrderResponse NewAlgoOrder(ctx).AlgoType(algoType).Symbol(symbol).Side(side).Type(type_).PositionSide(positionSide).TimeInForce(timeInForce).Quantity(quantity).Price(price).TriggerPrice(triggerPrice).WorkingType(workingType).PriceMatch(priceMatch).ClosePosition(closePosition).PriceProtect(priceProtect).ReduceOnly(reduceOnly).ActivatePrice(activatePrice).CallbackRate(callbackRate).ClientAlgoId(clientAlgoId).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).GoodTillDate(goodTillDate).RecvWindow(recvWindow).Execute()

New Algo Order (TRADE)


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
	algoType := models.NewAlgoOrderAlgoTypeParameterConditional // NewAlgoOrderAlgoTypeParameter | 
	symbol := "BNBUSDT" // string | 
	side := models.NewAlgoOrderSideParameterBuy // NewAlgoOrderSideParameter | 
	type_ := models.NewAlgoOrderTypeParameterLimit // NewAlgoOrderTypeParameter | 
	positionSide := "BOTH" // string | Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent in Hedge Mode. (optional)
	timeInForce := models.NewAlgoOrderTimeInForceParameterGtc // NewAlgoOrderTimeInForceParameter |  (optional)
	quantity := float32(1.0) // float32 | Cannot be sent with `closePosition`=`true`(Close-All) (optional)
	price := float32(1.0) // float32 |  (optional)
	triggerPrice := float32(1.0) // float32 | Trigger price (optional)
	workingType := models.NewAlgoOrderWorkingTypeParameterMarkPrice // NewAlgoOrderWorkingTypeParameter |  (optional)
	priceMatch := models.NewAlgoOrderPriceMatchParameterOpponent // NewAlgoOrderPriceMatchParameter | only avaliable for `LIMIT`/`STOP`/`TAKE_PROFIT` order; Can't be passed together with `price` (optional)
	closePosition := models.NewAlgoOrderClosePositionParameterTrue // NewAlgoOrderClosePositionParameter | Close-All，used with `STOP_MARKET` or `TAKE_PROFIT_MARKET`. (optional)
	priceProtect := models.NewAlgoOrderClosePositionParameterTrue // NewAlgoOrderClosePositionParameter | Used with `STOP_MARKET` or `TAKE_PROFIT_MARKET` order. when price reaches the triggerPrice ，the difference rate between \"MARK_PRICE\" and \"CONTRACT_PRICE\" cannot be larger than the Price Protection Threshold of the symbol.' (optional)
	reduceOnly := models.NewAlgoOrderClosePositionParameterTrue // NewAlgoOrderClosePositionParameter | Cannot be sent in Hedge Mode; cannot be sent with `closePosition`=`true`' (optional)
	activatePrice := float32(1.0) // float32 | Used with `TRAILING_STOP_MARKET` orders, default as the latest price(supporting different `workingType`) (optional)
	callbackRate := float32(1) // float32 | Used with `TRAILING_STOP_MARKET` orders (optional)
	clientAlgoId := "1" // string | A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: `^[\\.A-Z\\:/a-z0-9_-]{1,36}$` (optional)
	newOrderRespType := models.NewAlgoOrderNewOrderRespTypeParameterAck // NewAlgoOrderNewOrderRespTypeParameter |  (optional)
	selfTradePreventionMode := models.NewAlgoOrderSelfTradePreventionModeParameterNone // NewAlgoOrderSelfTradePreventionModeParameter | `EXPIRE_TAKER`:expire taker order when STP triggers / `EXPIRE_MAKER`:expire taker order when STP triggers/ `EXPIRE_BOTH`:expire both orders when STP triggers; default `NONE` (optional)
	goodTillDate := int64(1770736694138) // int64 | order cancel time for timeInForce `GTD`, mandatory when `timeInforce` set to `GTD`; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.NewAlgoOrder(context.Background()).AlgoType(algoType).Symbol(symbol).Side(side).Type(type_).PositionSide(positionSide).TimeInForce(timeInForce).Quantity(quantity).Price(price).TriggerPrice(triggerPrice).WorkingType(workingType).PriceMatch(priceMatch).ClosePosition(closePosition).PriceProtect(priceProtect).ReduceOnly(reduceOnly).ActivatePrice(activatePrice).CallbackRate(callbackRate).ClientAlgoId(clientAlgoId).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).GoodTillDate(goodTillDate).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.NewAlgoOrder``: %v\n", err)
		return
	}

	// response from `NewAlgoOrder`: NewAlgoOrderResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **algoType** | [**NewAlgoOrderAlgoTypeParameter**](NewAlgoOrderAlgoTypeParameter.md) |  | 
 **symbol** | **string** |  | 
 **side** | [**NewAlgoOrderSideParameter**](NewAlgoOrderSideParameter.md) |  | 
 **type_** | [**NewAlgoOrderTypeParameter**](NewAlgoOrderTypeParameter.md) |  | 
 **positionSide** | **string** | Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent in Hedge Mode. | 
 **timeInForce** | [**NewAlgoOrderTimeInForceParameter**](NewAlgoOrderTimeInForceParameter.md) |  | 
 **quantity** | **float32** | Cannot be sent with &#x60;closePosition&#x60;&#x3D;&#x60;true&#x60;(Close-All) | 
 **price** | **float32** |  | 
 **triggerPrice** | **float32** | Trigger price | 
 **workingType** | [**NewAlgoOrderWorkingTypeParameter**](NewAlgoOrderWorkingTypeParameter.md) |  | 
 **priceMatch** | [**NewAlgoOrderPriceMatchParameter**](NewAlgoOrderPriceMatchParameter.md) | only avaliable for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; Can&#39;t be passed together with &#x60;price&#x60; | 
 **closePosition** | [**NewAlgoOrderClosePositionParameter**](NewAlgoOrderClosePositionParameter.md) | Close-All，used with &#x60;STOP_MARKET&#x60; or &#x60;TAKE_PROFIT_MARKET&#x60;. | 
 **priceProtect** | [**NewAlgoOrderClosePositionParameter**](NewAlgoOrderClosePositionParameter.md) | Used with &#x60;STOP_MARKET&#x60; or &#x60;TAKE_PROFIT_MARKET&#x60; order. when price reaches the triggerPrice ，the difference rate between \&quot;MARK_PRICE\&quot; and \&quot;CONTRACT_PRICE\&quot; cannot be larger than the Price Protection Threshold of the symbol.&#39; | 
 **reduceOnly** | [**NewAlgoOrderClosePositionParameter**](NewAlgoOrderClosePositionParameter.md) | Cannot be sent in Hedge Mode; cannot be sent with &#x60;closePosition&#x60;&#x3D;&#x60;true&#x60;&#39; | 
 **activatePrice** | **float32** | Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, default as the latest price(supporting different &#x60;workingType&#x60;) | 
 **callbackRate** | **float32** | Used with &#x60;TRAILING_STOP_MARKET&#x60; orders | 
 **clientAlgoId** | **string** | A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: &#x60;^[\\.A-Z\\:/a-z0-9_-]{1,36}$&#x60; | 
 **newOrderRespType** | [**NewAlgoOrderNewOrderRespTypeParameter**](NewAlgoOrderNewOrderRespTypeParameter.md) |  | 
 **selfTradePreventionMode** | [**NewAlgoOrderSelfTradePreventionModeParameter**](NewAlgoOrderSelfTradePreventionModeParameter.md) | &#x60;EXPIRE_TAKER&#x60;:expire taker order when STP triggers / &#x60;EXPIRE_MAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_BOTH&#x60;:expire both orders when STP triggers; default &#x60;NONE&#x60; | 
 **goodTillDate** | **int64** | order cancel time for timeInForce &#x60;GTD&#x60;, mandatory when &#x60;timeInforce&#x60; set to &#x60;GTD&#x60;; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**NewAlgoOrderResponse**](NewAlgoOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## NewOrder

> NewOrderResponse NewOrder(ctx).Symbol(symbol).Side(side).Type(type_).PositionSide(positionSide).TimeInForce(timeInForce).ReduceOnly(reduceOnly).Quantity(quantity).Price(price).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).PriceMatch(priceMatch).SelfTradePreventionMode(selfTradePreventionMode).GoodTillDate(goodTillDate).RecvWindow(recvWindow).Execute()

New Order (TRADE)


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
	symbol := "BTCUSDT" // string | 
	side := models.NewAlgoOrderSideParameterBuy // NewAlgoOrderSideParameter | 
	type_ := models.PlaceMultipleOrdersBatchOrdersParameterInnerTypeLimit // PlaceMultipleOrdersBatchOrdersParameterInnerType | Order type
	positionSide := "BOTH" // string | Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent in Hedge Mode. (optional)
	timeInForce := models.NewAlgoOrderTimeInForceParameterGtc // NewAlgoOrderTimeInForceParameter |  (optional)
	reduceOnly := models.NewAlgoOrderClosePositionParameterTrue // NewAlgoOrderClosePositionParameter | Cannot be sent in Hedge Mode (optional)
	quantity := float32(1.0) // float32 |  (optional)
	price := float32(1.0) // float32 |  (optional)
	newClientOrderId := "1" // string | A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: `^[\\.A-Z\\:/a-z0-9_-]{1,36}$` (optional)
	newOrderRespType := models.NewAlgoOrderNewOrderRespTypeParameterAck // NewAlgoOrderNewOrderRespTypeParameter |  (optional)
	priceMatch := models.NewAlgoOrderPriceMatchParameterOpponent // NewAlgoOrderPriceMatchParameter | only avaliable for `LIMIT`/`STOP`/`TAKE_PROFIT` order; Can't be passed together with `price` (optional)
	selfTradePreventionMode := models.NewOrderSelfTradePreventionModeParameterExpireTaker // NewOrderSelfTradePreventionModeParameter | `EXPIRE_TAKER`:expire taker order when STP triggers/ `EXPIRE_MAKER`:expire taker order when STP triggers/ `EXPIRE_BOTH`:expire both orders when STP triggers; default `EXPIRE_MAKER` (optional)
	goodTillDate := int64(1770736694138) // int64 | order cancel time for timeInForce `GTD`, mandatory when `timeInforce` set to `GTD`; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.NewOrder(context.Background()).Symbol(symbol).Side(side).Type(type_).PositionSide(positionSide).TimeInForce(timeInForce).ReduceOnly(reduceOnly).Quantity(quantity).Price(price).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).PriceMatch(priceMatch).SelfTradePreventionMode(selfTradePreventionMode).GoodTillDate(goodTillDate).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.NewOrder``: %v\n", err)
		return
	}

	// response from `NewOrder`: NewOrderResponse
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
 **side** | [**NewAlgoOrderSideParameter**](NewAlgoOrderSideParameter.md) |  | 
 **type_** | [**PlaceMultipleOrdersBatchOrdersParameterInnerType**](PlaceMultipleOrdersBatchOrdersParameterInnerType.md) | Order type | 
 **positionSide** | **string** | Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent in Hedge Mode. | 
 **timeInForce** | [**NewAlgoOrderTimeInForceParameter**](NewAlgoOrderTimeInForceParameter.md) |  | 
 **reduceOnly** | [**NewAlgoOrderClosePositionParameter**](NewAlgoOrderClosePositionParameter.md) | Cannot be sent in Hedge Mode | 
 **quantity** | **float32** |  | 
 **price** | **float32** |  | 
 **newClientOrderId** | **string** | A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: &#x60;^[\\.A-Z\\:/a-z0-9_-]{1,36}$&#x60; | 
 **newOrderRespType** | [**NewAlgoOrderNewOrderRespTypeParameter**](NewAlgoOrderNewOrderRespTypeParameter.md) |  | 
 **priceMatch** | [**NewAlgoOrderPriceMatchParameter**](NewAlgoOrderPriceMatchParameter.md) | only avaliable for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; Can&#39;t be passed together with &#x60;price&#x60; | 
 **selfTradePreventionMode** | [**NewOrderSelfTradePreventionModeParameter**](NewOrderSelfTradePreventionModeParameter.md) | &#x60;EXPIRE_TAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_MAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_BOTH&#x60;:expire both orders when STP triggers; default &#x60;EXPIRE_MAKER&#x60; | 
 **goodTillDate** | **int64** | order cancel time for timeInForce &#x60;GTD&#x60;, mandatory when &#x60;timeInforce&#x60; set to &#x60;GTD&#x60;; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**NewOrderResponse**](NewOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## PlaceMultipleOrders

> PlaceMultipleOrdersResponse PlaceMultipleOrders(ctx).BatchOrders(batchOrders).RecvWindow(recvWindow).Execute()

Place Multiple Orders (TRADE)


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
	batchOrders := []models.PlaceMultipleOrdersBatchOrdersParameterInner{*models.NewPlaceMultipleOrdersBatchOrdersParameterInner()} // []PlaceMultipleOrdersBatchOrdersParameterInner | order list. Max 5 orders
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.PlaceMultipleOrders(context.Background()).BatchOrders(batchOrders).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.PlaceMultipleOrders``: %v\n", err)
		return
	}

	// response from `PlaceMultipleOrders`: PlaceMultipleOrdersResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **batchOrders** | [**[]PlaceMultipleOrdersBatchOrdersParameterInner**](PlaceMultipleOrdersBatchOrdersParameterInner.md) | order list. Max 5 orders | 
 **recvWindow** | **int64** |  | 

### Return type

[**PlaceMultipleOrdersResponse**](PlaceMultipleOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## PositionAdlQuantileEstimation

> PositionAdlQuantileEstimationResponse PositionAdlQuantileEstimation(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Position ADL Quantile Estimation (USER_DATA)


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
	symbol := "BTCUSDT" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.PositionAdlQuantileEstimation(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.PositionAdlQuantileEstimation``: %v\n", err)
		return
	}

	// response from `PositionAdlQuantileEstimation`: PositionAdlQuantileEstimationResponse
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
 **recvWindow** | **int64** |  | 

### Return type

[**PositionAdlQuantileEstimationResponse**](PositionAdlQuantileEstimationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## PositionInformationV2

> PositionInformationV2Response PositionInformationV2(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Position Information V2 (USER_DATA)


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
	symbol := "BTCUSDT" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.PositionInformationV2(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.PositionInformationV2``: %v\n", err)
		return
	}

	// response from `PositionInformationV2`: PositionInformationV2Response
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
 **recvWindow** | **int64** |  | 

### Return type

[**PositionInformationV2Response**](PositionInformationV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## PositionInformationV3

> PositionInformationV3Response PositionInformationV3(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Position Information V3 (USER_DATA)


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
	symbol := "BTCUSDT" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.PositionInformationV3(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.PositionInformationV3``: %v\n", err)
		return
	}

	// response from `PositionInformationV3`: PositionInformationV3Response
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
 **recvWindow** | **int64** |  | 

### Return type

[**PositionInformationV3Response**](PositionInformationV3Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryAlgoOrder

> QueryAlgoOrderResponse QueryAlgoOrder(ctx).AlgoId(algoId).ClientAlgoId(clientAlgoId).RecvWindow(recvWindow).Execute()

Query Algo Order (USER_DATA)


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
	algoId := int64(1) // int64 | Order ID (optional)
	clientAlgoId := "1" // string | Client order ID (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryAlgoOrder(context.Background()).AlgoId(algoId).ClientAlgoId(clientAlgoId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryAlgoOrder``: %v\n", err)
		return
	}

	// response from `QueryAlgoOrder`: QueryAlgoOrderResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **algoId** | **int64** | Order ID | 
 **clientAlgoId** | **string** | Client order ID | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryAlgoOrderResponse**](QueryAlgoOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryAllAlgoOrders

> QueryAllAlgoOrdersResponse QueryAllAlgoOrders(ctx).Symbol(symbol).AlgoId(algoId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query All Algo Orders (USER_DATA)


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
	symbol := "BTCUSDT" // string | Symbol
	algoId := int64(2146760) // int64 |  (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	limit := int64(50) // int64 |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryAllAlgoOrders(context.Background()).Symbol(symbol).AlgoId(algoId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryAllAlgoOrders``: %v\n", err)
		return
	}

	// response from `QueryAllAlgoOrders`: QueryAllAlgoOrdersResponse
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
 **algoId** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryAllAlgoOrdersResponse**](QueryAllAlgoOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryCurrentOpenOrder

> QueryCurrentOpenOrderResponse QueryCurrentOpenOrder(ctx).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Current Open Order (USER_DATA)


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
	symbol := "BTCUSDT" // string | 
	orderId := int64(1917641) // int64 |  (optional)
	origClientOrderId := "abc" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryCurrentOpenOrder(context.Background()).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryCurrentOpenOrder``: %v\n", err)
		return
	}

	// response from `QueryCurrentOpenOrder`: QueryCurrentOpenOrderResponse
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
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryCurrentOpenOrderResponse**](QueryCurrentOpenOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryOrder

> QueryOrderResponse QueryOrder(ctx).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Order (USER_DATA)


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
	symbol := "BTCUSDT" // string | 
	orderId := int64(1917641) // int64 |  (optional)
	origClientOrderId := "abc" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryOrder(context.Background()).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryOrder``: %v\n", err)
		return
	}

	// response from `QueryOrder`: QueryOrderResponse
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
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryOrderResponse**](QueryOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## TestOrder

> TestOrderResponse TestOrder(ctx).Symbol(symbol).Side(side).Type(type_).PositionSide(positionSide).ReduceOnly(reduceOnly).Quantity(quantity).Price(price).NewClientOrderId(newClientOrderId).StopPrice(stopPrice).ClosePosition(closePosition).ActivationPrice(activationPrice).CallbackRate(callbackRate).TimeInForce(timeInForce).WorkingType(workingType).PriceProtect(priceProtect).NewOrderRespType(newOrderRespType).PriceMatch(priceMatch).SelfTradePreventionMode(selfTradePreventionMode).GoodTillDate(goodTillDate).RecvWindow(recvWindow).Execute()

Test Order (TRADE)


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
	symbol := "BTCUSDT" // string | 
	side := models.NewAlgoOrderSideParameterBuy // NewAlgoOrderSideParameter | 
	type_ := models.PlaceMultipleOrdersBatchOrdersParameterInnerTypeLimit // PlaceMultipleOrdersBatchOrdersParameterInnerType | 
	positionSide := models.TestOrderPositionSideParameterBoth // TestOrderPositionSideParameter | Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent in Hedge Mode. (optional)
	reduceOnly := models.NewAlgoOrderClosePositionParameterTrue // NewAlgoOrderClosePositionParameter | Cannot be sent in Hedge Mode; cannot be sent with `closePosition`=`true` (optional)
	quantity := float32(1.0) // float32 | Cannot be sent with `closePosition`=`true`(Close-All) (optional)
	price := float32(1.0) // float32 |  (optional)
	newClientOrderId := "1" // string | A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: `^[\\.A-Z\\:/a-z0-9_-]{1,36}$` (optional)
	stopPrice := float32(1.0) // float32 | Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders. (optional)
	closePosition := models.NewAlgoOrderClosePositionParameterTrue // NewAlgoOrderClosePositionParameter | Close-All，used with `STOP_MARKET` or `TAKE_PROFIT_MARKET`.\" (optional)
	activationPrice := float32(1.0) // float32 | Used with `TRAILING_STOP_MARKET` orders, default as the latest price(supporting different `workingType`) (optional)
	callbackRate := float32(1) // float32 | Used with `TRAILING_STOP_MARKET` orders (optional)
	timeInForce := models.NewAlgoOrderTimeInForceParameterGtc // NewAlgoOrderTimeInForceParameter |  (optional)
	workingType := models.NewAlgoOrderWorkingTypeParameterMarkPrice // NewAlgoOrderWorkingTypeParameter |  (optional)
	priceProtect := models.NewAlgoOrderClosePositionParameterTrue // NewAlgoOrderClosePositionParameter |  (optional)
	newOrderRespType := models.NewAlgoOrderNewOrderRespTypeParameterAck // NewAlgoOrderNewOrderRespTypeParameter |  (optional)
	priceMatch := models.NewAlgoOrderPriceMatchParameterOpponent // NewAlgoOrderPriceMatchParameter | only avaliable for `LIMIT`/`STOP`/`TAKE_PROFIT` order; Can't be passed together with `price` (optional)
	selfTradePreventionMode := models.NewAlgoOrderSelfTradePreventionModeParameterNone // NewAlgoOrderSelfTradePreventionModeParameter | `NONE`:No STP / `EXPIRE_TAKER`:expire taker order when STP triggers/ `EXPIRE_MAKER`:expire taker order when STP triggers/ `EXPIRE_BOTH`:expire both orders when STP triggers; default `NONE` (optional)
	goodTillDate := int64(1770736694138) // int64 | order cancel time for timeInForce `GTD`, mandatory when `timeInforce` set to `GTD`; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.TestOrder(context.Background()).Symbol(symbol).Side(side).Type(type_).PositionSide(positionSide).ReduceOnly(reduceOnly).Quantity(quantity).Price(price).NewClientOrderId(newClientOrderId).StopPrice(stopPrice).ClosePosition(closePosition).ActivationPrice(activationPrice).CallbackRate(callbackRate).TimeInForce(timeInForce).WorkingType(workingType).PriceProtect(priceProtect).NewOrderRespType(newOrderRespType).PriceMatch(priceMatch).SelfTradePreventionMode(selfTradePreventionMode).GoodTillDate(goodTillDate).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.TestOrder``: %v\n", err)
		return
	}

	// response from `TestOrder`: TestOrderResponse
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
 **side** | [**NewAlgoOrderSideParameter**](NewAlgoOrderSideParameter.md) |  | 
 **type_** | [**PlaceMultipleOrdersBatchOrdersParameterInnerType**](PlaceMultipleOrdersBatchOrdersParameterInnerType.md) |  | 
 **positionSide** | [**TestOrderPositionSideParameter**](TestOrderPositionSideParameter.md) | Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent in Hedge Mode. | 
 **reduceOnly** | [**NewAlgoOrderClosePositionParameter**](NewAlgoOrderClosePositionParameter.md) | Cannot be sent in Hedge Mode; cannot be sent with &#x60;closePosition&#x60;&#x3D;&#x60;true&#x60; | 
 **quantity** | **float32** | Cannot be sent with &#x60;closePosition&#x60;&#x3D;&#x60;true&#x60;(Close-All) | 
 **price** | **float32** |  | 
 **newClientOrderId** | **string** | A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: &#x60;^[\\.A-Z\\:/a-z0-9_-]{1,36}$&#x60; | 
 **stopPrice** | **float32** | Used with &#x60;STOP/STOP_MARKET&#x60; or &#x60;TAKE_PROFIT/TAKE_PROFIT_MARKET&#x60; orders. | 
 **closePosition** | [**NewAlgoOrderClosePositionParameter**](NewAlgoOrderClosePositionParameter.md) | Close-All，used with &#x60;STOP_MARKET&#x60; or &#x60;TAKE_PROFIT_MARKET&#x60;.\&quot; | 
 **activationPrice** | **float32** | Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, default as the latest price(supporting different &#x60;workingType&#x60;) | 
 **callbackRate** | **float32** | Used with &#x60;TRAILING_STOP_MARKET&#x60; orders | 
 **timeInForce** | [**NewAlgoOrderTimeInForceParameter**](NewAlgoOrderTimeInForceParameter.md) |  | 
 **workingType** | [**NewAlgoOrderWorkingTypeParameter**](NewAlgoOrderWorkingTypeParameter.md) |  | 
 **priceProtect** | [**NewAlgoOrderClosePositionParameter**](NewAlgoOrderClosePositionParameter.md) |  | 
 **newOrderRespType** | [**NewAlgoOrderNewOrderRespTypeParameter**](NewAlgoOrderNewOrderRespTypeParameter.md) |  | 
 **priceMatch** | [**NewAlgoOrderPriceMatchParameter**](NewAlgoOrderPriceMatchParameter.md) | only avaliable for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; Can&#39;t be passed together with &#x60;price&#x60; | 
 **selfTradePreventionMode** | [**NewAlgoOrderSelfTradePreventionModeParameter**](NewAlgoOrderSelfTradePreventionModeParameter.md) | &#x60;NONE&#x60;:No STP / &#x60;EXPIRE_TAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_MAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_BOTH&#x60;:expire both orders when STP triggers; default &#x60;NONE&#x60; | 
 **goodTillDate** | **int64** | order cancel time for timeInForce &#x60;GTD&#x60;, mandatory when &#x60;timeInforce&#x60; set to &#x60;GTD&#x60;; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**TestOrderResponse**](TestOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## UsersForceOrders

> UsersForceOrdersResponse UsersForceOrders(ctx).Symbol(symbol).AutoCloseType(autoCloseType).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

User's Force Orders (USER_DATA)


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
	symbol := "BTCUSDT" // string |  (optional)
	autoCloseType := models.UsersForceOrdersAutoCloseTypeParameterLiquidation // UsersForceOrdersAutoCloseTypeParameter | \"LIQUIDATION\" for liquidation orders, \"ADL\" for ADL orders. (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	limit := int64(50) // int64 |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.UsersForceOrders(context.Background()).Symbol(symbol).AutoCloseType(autoCloseType).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.UsersForceOrders``: %v\n", err)
		return
	}

	// response from `UsersForceOrders`: UsersForceOrdersResponse
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
 **autoCloseType** | [**UsersForceOrdersAutoCloseTypeParameter**](UsersForceOrdersAutoCloseTypeParameter.md) | \&quot;LIQUIDATION\&quot; for liquidation orders, \&quot;ADL\&quot; for ADL orders. | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UsersForceOrdersResponse**](UsersForceOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

