# \TradeAPI

All URIs are relative to *https://dapi.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AccountTradeList**](TradeAPI.md#AccountTradeList) | **Get** /dapi/v1/userTrades | Account Trade List (USER_DATA)
[**AllOrders**](TradeAPI.md#AllOrders) | **Get** /dapi/v1/allOrders | All Orders (USER_DATA)
[**AutoCancelAllOpenOrders**](TradeAPI.md#AutoCancelAllOpenOrders) | **Post** /dapi/v1/countdownCancelAll | Auto-Cancel All Open Orders (TRADE)
[**CancelAllOpenOrders**](TradeAPI.md#CancelAllOpenOrders) | **Delete** /dapi/v1/allOpenOrders | Cancel All Open Orders(TRADE)
[**CancelMultipleOrders**](TradeAPI.md#CancelMultipleOrders) | **Delete** /dapi/v1/batchOrders | Cancel Multiple Orders(TRADE)
[**CancelOrder**](TradeAPI.md#CancelOrder) | **Delete** /dapi/v1/order | Cancel Order (TRADE)
[**ChangeInitialLeverage**](TradeAPI.md#ChangeInitialLeverage) | **Post** /dapi/v1/leverage | Change Initial Leverage (TRADE)
[**ChangeMarginType**](TradeAPI.md#ChangeMarginType) | **Post** /dapi/v1/marginType | Change Margin Type (TRADE)
[**ChangePositionMode**](TradeAPI.md#ChangePositionMode) | **Post** /dapi/v1/positionSide/dual | Change Position Mode(TRADE)
[**CurrentAllOpenOrders**](TradeAPI.md#CurrentAllOpenOrders) | **Get** /dapi/v1/openOrders | Current All Open Orders (USER_DATA)
[**GetOrderModifyHistory**](TradeAPI.md#GetOrderModifyHistory) | **Get** /dapi/v1/orderAmendment | Get Order Modify History (USER_DATA)
[**GetPositionMarginChangeHistory**](TradeAPI.md#GetPositionMarginChangeHistory) | **Get** /dapi/v1/positionMargin/history | Get Position Margin Change History(TRADE)
[**ModifyIsolatedPositionMargin**](TradeAPI.md#ModifyIsolatedPositionMargin) | **Post** /dapi/v1/positionMargin | Modify Isolated Position Margin(TRADE)
[**ModifyMultipleOrders**](TradeAPI.md#ModifyMultipleOrders) | **Put** /dapi/v1/batchOrders | Modify Multiple Orders(TRADE)
[**ModifyOrder**](TradeAPI.md#ModifyOrder) | **Put** /dapi/v1/order | Modify Order (TRADE)
[**NewOrder**](TradeAPI.md#NewOrder) | **Post** /dapi/v1/order | New Order (TRADE)
[**PlaceMultipleOrders**](TradeAPI.md#PlaceMultipleOrders) | **Post** /dapi/v1/batchOrders | Place Multiple Orders(TRADE)
[**PositionAdlQuantileEstimation**](TradeAPI.md#PositionAdlQuantileEstimation) | **Get** /dapi/v1/adlQuantile | Position ADL Quantile Estimation(USER_DATA)
[**PositionInformation**](TradeAPI.md#PositionInformation) | **Get** /dapi/v1/positionRisk | Position Information(USER_DATA)
[**QueryCurrentOpenOrder**](TradeAPI.md#QueryCurrentOpenOrder) | **Get** /dapi/v1/openOrder | Query Current Open Order(USER_DATA)
[**QueryOrder**](TradeAPI.md#QueryOrder) | **Get** /dapi/v1/order | Query Order (USER_DATA)
[**UsersForceOrders**](TradeAPI.md#UsersForceOrders) | **Get** /dapi/v1/forceOrders | User&#39;s Force Orders(USER_DATA)


## AccountTradeList

> AccountTradeListResponse AccountTradeList(ctx).Symbol(symbol).Pair(pair).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()

Account Trade List (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string |  (optional)
	pair := "pair_example" // string |  (optional)
	orderId := "1" // string |  (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	fromId := int64(1) // int64 | ID to get aggregate trades from INCLUSIVE. (optional)
	limit := int64(100) // int64 | Default 100; max 1000 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.AccountTradeList(context.Background()).Symbol(symbol).Pair(pair).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()
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
 **pair** | **string** |  | 
 **orderId** | **string** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **fromId** | **int64** | ID to get aggregate trades from INCLUSIVE. | 
 **limit** | **int64** | Default 100; max 1000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**AccountTradeListResponse**](AccountTradeListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## AllOrders

> AllOrdersResponse AllOrders(ctx).Symbol(symbol).Pair(pair).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

All Orders (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string |  (optional)
	pair := "pair_example" // string |  (optional)
	orderId := int64(1) // int64 |  (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	limit := int64(100) // int64 | Default 100; max 1000 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.AllOrders(context.Background()).Symbol(symbol).Pair(pair).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
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
 **pair** | **string** |  | 
 **orderId** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int64** | Default 100; max 1000 | 
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

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	countdownTime := int64(789) // int64 | countdown time, 1000 for 1 second. 0 to cancel the timer
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

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
 **countdownTime** | **int64** | countdown time, 1000 for 1 second. 0 to cancel the timer | 
 **recvWindow** | **int64** |  | 

### Return type

[**AutoCancelAllOpenOrdersResponse**](AutoCancelAllOpenOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CancelAllOpenOrders

> CancelAllOpenOrdersResponse CancelAllOpenOrders(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Cancel All Open Orders(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

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

Cancel Multiple Orders(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	orderIdList := []int64{int64(1234567)} // []int64 | max length 10 <br /> e.g. [1234567,2345678] (optional)
	origClientOrderIdList := []string{"my_id_1"} // []string | max length 10<br /> e.g. [\"my_id_1\",\"my_id_2\"], encode the double quotes. No space after comma. (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

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
 **orderIdList** | **[]int64** | max length 10 &lt;br /&gt; e.g. [1234567,2345678] | 
 **origClientOrderIdList** | **[]string** | max length 10&lt;br /&gt; e.g. [\&quot;my_id_1\&quot;,\&quot;my_id_2\&quot;], encode the double quotes. No space after comma. | 
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

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	orderId := int64(1) // int64 |  (optional)
	origClientOrderId := "1" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	leverage := int64(789) // int64 | target initial leverage: int from 1 to 125
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

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
 **leverage** | **int64** | target initial leverage: int from 1 to 125 | 
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

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	marginType := models.ChangeMarginTypeMarginTypeParameterIsolated // ChangeMarginTypeMarginTypeParameter | ISOLATED, CROSSED
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

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
 **marginType** | [**ChangeMarginTypeMarginTypeParameter**](ChangeMarginTypeMarginTypeParameter.md) | ISOLATED, CROSSED | 
 **recvWindow** | **int64** |  | 

### Return type

[**ChangeMarginTypeResponse**](ChangeMarginTypeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## ChangePositionMode

> ChangePositionModeResponse ChangePositionMode(ctx).DualSidePosition(dualSidePosition).RecvWindow(recvWindow).Execute()

Change Position Mode(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	dualSidePosition := "dualSidePosition_example" // string | \"true\": Hedge Mode; \"false\": One-way Mode
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

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


## CurrentAllOpenOrders

> CurrentAllOpenOrdersResponse CurrentAllOpenOrders(ctx).Symbol(symbol).Pair(pair).RecvWindow(recvWindow).Execute()

Current All Open Orders (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string |  (optional)
	pair := "pair_example" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CurrentAllOpenOrders(context.Background()).Symbol(symbol).Pair(pair).RecvWindow(recvWindow).Execute()
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
 **pair** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CurrentAllOpenOrdersResponse**](CurrentAllOpenOrdersResponse.md)

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	orderId := int64(1) // int64 |  (optional)
	origClientOrderId := "1" // string |  (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	limit := int64(100) // int64 | Default 100; max 1000 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

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
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int64** | Default 100; max 1000 | 
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

Get Position Margin Change History(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	type_ := int64(789) // int64 | 1: Add position margin,2: Reduce position margin (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	limit := int64(100) // int64 | Default 100; max 1000 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

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
 **type_** | **int64** | 1: Add position margin,2: Reduce position margin | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int64** | Default 100; max 1000 | 
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

Modify Isolated Position Margin(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	amount := float32(1.0) // float32 | 
	type_ := models.PlaceMultipleOrdersBatchOrdersParameterInnerTypeLimit // PlaceMultipleOrdersBatchOrdersParameterInnerType | 
	positionSide := models.PlaceMultipleOrdersBatchOrdersParameterInnerPositionSideBoth // PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide | Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent with Hedge Mode. (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

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
 **amount** | **float32** |  | 
 **type_** | [**PlaceMultipleOrdersBatchOrdersParameterInnerType**](PlaceMultipleOrdersBatchOrdersParameterInnerType.md) |  | 
 **positionSide** | [**PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide**](PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide.md) | Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent with Hedge Mode. | 
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

Modify Multiple Orders(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	batchOrders := []models.ModifyMultipleOrdersBatchOrdersParameterInner{*models.NewModifyMultipleOrdersBatchOrdersParameterInner()} // []ModifyMultipleOrdersBatchOrdersParameterInner | order list. Max 5 orders
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

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

> ModifyOrderResponse ModifyOrder(ctx).Symbol(symbol).Side(side).OrderId(orderId).OrigClientOrderId(origClientOrderId).Quantity(quantity).Price(price).PriceMatch(priceMatch).RecvWindow(recvWindow).Execute()

Modify Order (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	side := models.ModifyMultipleOrdersBatchOrdersParameterInnerSideBuy // ModifyMultipleOrdersBatchOrdersParameterInnerSide | `SELL`, `BUY`
	orderId := int64(1) // int64 |  (optional)
	origClientOrderId := "1" // string |  (optional)
	quantity := float32(1.0) // float32 | quantity measured by contract number, Cannot be sent with `closePosition`=`true` (optional)
	price := float32(1.0) // float32 |  (optional)
	priceMatch := models.PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatchNone // PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch | only avaliable for `LIMIT`/`STOP`/`TAKE_PROFIT` order; can be set to `OPPONENT`/ `OPPONENT_5`/ `OPPONENT_10`/ `OPPONENT_20`: /`QUEUE`/ `QUEUE_5`/ `QUEUE_10`/ `QUEUE_20`; Can't be passed together with `price` (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.ModifyOrder(context.Background()).Symbol(symbol).Side(side).OrderId(orderId).OrigClientOrderId(origClientOrderId).Quantity(quantity).Price(price).PriceMatch(priceMatch).RecvWindow(recvWindow).Execute()
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
 **side** | [**ModifyMultipleOrdersBatchOrdersParameterInnerSide**](ModifyMultipleOrdersBatchOrdersParameterInnerSide.md) | &#x60;SELL&#x60;, &#x60;BUY&#x60; | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | 
 **quantity** | **float32** | quantity measured by contract number, Cannot be sent with &#x60;closePosition&#x60;&#x3D;&#x60;true&#x60; | 
 **price** | **float32** |  | 
 **priceMatch** | [**PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch**](PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch.md) | only avaliable for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; can be set to &#x60;OPPONENT&#x60;/ &#x60;OPPONENT_5&#x60;/ &#x60;OPPONENT_10&#x60;/ &#x60;OPPONENT_20&#x60;: /&#x60;QUEUE&#x60;/ &#x60;QUEUE_5&#x60;/ &#x60;QUEUE_10&#x60;/ &#x60;QUEUE_20&#x60;; Can&#39;t be passed together with &#x60;price&#x60; | 
 **recvWindow** | **int64** |  | 

### Return type

[**ModifyOrderResponse**](ModifyOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## NewOrder

> NewOrderResponse NewOrder(ctx).Symbol(symbol).Side(side).Type(type_).PositionSide(positionSide).TimeInForce(timeInForce).Quantity(quantity).ReduceOnly(reduceOnly).Price(price).NewClientOrderId(newClientOrderId).StopPrice(stopPrice).ClosePosition(closePosition).ActivationPrice(activationPrice).CallbackRate(callbackRate).WorkingType(workingType).PriceProtect(priceProtect).NewOrderRespType(newOrderRespType).PriceMatch(priceMatch).SelfTradePreventionMode(selfTradePreventionMode).RecvWindow(recvWindow).Execute()

New Order (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	side := models.ModifyMultipleOrdersBatchOrdersParameterInnerSideBuy // ModifyMultipleOrdersBatchOrdersParameterInnerSide | `SELL`, `BUY`
	type_ := models.PlaceMultipleOrdersBatchOrdersParameterInnerTypeLimit // PlaceMultipleOrdersBatchOrdersParameterInnerType | 
	positionSide := models.PlaceMultipleOrdersBatchOrdersParameterInnerPositionSideBoth // PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide | Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent with Hedge Mode. (optional)
	timeInForce := models.PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForceGtc // PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce |  (optional)
	quantity := float32(1.0) // float32 | quantity measured by contract number, Cannot be sent with `closePosition`=`true` (optional)
	reduceOnly := "false" // string | \"true\" or \"false\". default \"false\". Cannot be sent in Hedge Mode; cannot be sent with `closePosition`=`true`(Close-All) (optional)
	price := float32(1.0) // float32 |  (optional)
	newClientOrderId := "1" // string | A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: `^[\\.A-Z\\:/a-z0-9_-]{1,36}$` (optional)
	stopPrice := float32(1.0) // float32 | Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders. (optional)
	closePosition := "closePosition_example" // string | `true`, `false`；Close-All,used with `STOP_MARKET` or `TAKE_PROFIT_MARKET`. (optional)
	activationPrice := float32(1.0) // float32 | Used with `TRAILING_STOP_MARKET` orders, default as the latest price(supporting different `workingType`) (optional)
	callbackRate := float32(1.0) // float32 | Used with `TRAILING_STOP_MARKET` orders, min 0.1, max 10 where 1 for 1% (optional)
	workingType := models.PlaceMultipleOrdersBatchOrdersParameterInnerWorkingTypeMarkPrice // PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType | stopPrice triggered by: \"MARK_PRICE\", \"CONTRACT_PRICE\". Default \"CONTRACT_PRICE\" (optional)
	priceProtect := "false" // string | \"TRUE\" or \"FALSE\", default \"FALSE\". Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders. (optional)
	newOrderRespType := models.PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespTypeAck // PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType | \"ACK\", \"RESULT\", default \"ACK\" (optional)
	priceMatch := models.PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatchNone // PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch | only avaliable for `LIMIT`/`STOP`/`TAKE_PROFIT` order; can be set to `OPPONENT`/ `OPPONENT_5`/ `OPPONENT_10`/ `OPPONENT_20`: /`QUEUE`/ `QUEUE_5`/ `QUEUE_10`/ `QUEUE_20`; Can't be passed together with `price` (optional)
	selfTradePreventionMode := models.PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionModeNone // PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode | `EXPIRE_TAKER`:expire taker order when STP triggers/ `EXPIRE_MAKER`:expire taker order when STP triggers/ `EXPIRE_BOTH`:expire both orders when STP triggers; default `EXPIRE_MAKER` (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.NewOrder(context.Background()).Symbol(symbol).Side(side).Type(type_).PositionSide(positionSide).TimeInForce(timeInForce).Quantity(quantity).ReduceOnly(reduceOnly).Price(price).NewClientOrderId(newClientOrderId).StopPrice(stopPrice).ClosePosition(closePosition).ActivationPrice(activationPrice).CallbackRate(callbackRate).WorkingType(workingType).PriceProtect(priceProtect).NewOrderRespType(newOrderRespType).PriceMatch(priceMatch).SelfTradePreventionMode(selfTradePreventionMode).RecvWindow(recvWindow).Execute()
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
 **side** | [**ModifyMultipleOrdersBatchOrdersParameterInnerSide**](ModifyMultipleOrdersBatchOrdersParameterInnerSide.md) | &#x60;SELL&#x60;, &#x60;BUY&#x60; | 
 **type_** | [**PlaceMultipleOrdersBatchOrdersParameterInnerType**](PlaceMultipleOrdersBatchOrdersParameterInnerType.md) |  | 
 **positionSide** | [**PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide**](PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide.md) | Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent with Hedge Mode. | 
 **timeInForce** | [**PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce**](PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce.md) |  | 
 **quantity** | **float32** | quantity measured by contract number, Cannot be sent with &#x60;closePosition&#x60;&#x3D;&#x60;true&#x60; | 
 **reduceOnly** | **string** | \&quot;true\&quot; or \&quot;false\&quot;. default \&quot;false\&quot;. Cannot be sent in Hedge Mode; cannot be sent with &#x60;closePosition&#x60;&#x3D;&#x60;true&#x60;(Close-All) | 
 **price** | **float32** |  | 
 **newClientOrderId** | **string** | A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: &#x60;^[\\.A-Z\\:/a-z0-9_-]{1,36}$&#x60; | 
 **stopPrice** | **float32** | Used with &#x60;STOP/STOP_MARKET&#x60; or &#x60;TAKE_PROFIT/TAKE_PROFIT_MARKET&#x60; orders. | 
 **closePosition** | **string** | &#x60;true&#x60;, &#x60;false&#x60;；Close-All,used with &#x60;STOP_MARKET&#x60; or &#x60;TAKE_PROFIT_MARKET&#x60;. | 
 **activationPrice** | **float32** | Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, default as the latest price(supporting different &#x60;workingType&#x60;) | 
 **callbackRate** | **float32** | Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, min 0.1, max 10 where 1 for 1% | 
 **workingType** | [**PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType**](PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType.md) | stopPrice triggered by: \&quot;MARK_PRICE\&quot;, \&quot;CONTRACT_PRICE\&quot;. Default \&quot;CONTRACT_PRICE\&quot; | 
 **priceProtect** | **string** | \&quot;TRUE\&quot; or \&quot;FALSE\&quot;, default \&quot;FALSE\&quot;. Used with &#x60;STOP/STOP_MARKET&#x60; or &#x60;TAKE_PROFIT/TAKE_PROFIT_MARKET&#x60; orders. | 
 **newOrderRespType** | [**PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType**](PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType.md) | \&quot;ACK\&quot;, \&quot;RESULT\&quot;, default \&quot;ACK\&quot; | 
 **priceMatch** | [**PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch**](PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch.md) | only avaliable for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; can be set to &#x60;OPPONENT&#x60;/ &#x60;OPPONENT_5&#x60;/ &#x60;OPPONENT_10&#x60;/ &#x60;OPPONENT_20&#x60;: /&#x60;QUEUE&#x60;/ &#x60;QUEUE_5&#x60;/ &#x60;QUEUE_10&#x60;/ &#x60;QUEUE_20&#x60;; Can&#39;t be passed together with &#x60;price&#x60; | 
 **selfTradePreventionMode** | [**PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode**](PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode.md) | &#x60;EXPIRE_TAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_MAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_BOTH&#x60;:expire both orders when STP triggers; default &#x60;EXPIRE_MAKER&#x60; | 
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

Place Multiple Orders(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	batchOrders := []models.PlaceMultipleOrdersBatchOrdersParameterInner{*models.NewPlaceMultipleOrdersBatchOrdersParameterInner()} // []PlaceMultipleOrdersBatchOrdersParameterInner | order list. Max 5 orders
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

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

Position ADL Quantile Estimation(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

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


## PositionInformation

> PositionInformationResponse PositionInformation(ctx).MarginAsset(marginAsset).Pair(pair).RecvWindow(recvWindow).Execute()

Position Information(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	marginAsset := "marginAsset_example" // string |  (optional)
	pair := "pair_example" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.PositionInformation(context.Background()).MarginAsset(marginAsset).Pair(pair).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.PositionInformation``: %v\n", err)
		return
	}

	// response from `PositionInformation`: PositionInformationResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **marginAsset** | **string** |  | 
 **pair** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PositionInformationResponse**](PositionInformationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryCurrentOpenOrder

> QueryCurrentOpenOrderResponse QueryCurrentOpenOrder(ctx).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Current Open Order(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	orderId := int64(1) // int64 |  (optional)
	origClientOrderId := "1" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	orderId := int64(1) // int64 |  (optional)
	origClientOrderId := "1" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

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


## UsersForceOrders

> UsersForceOrdersResponse UsersForceOrders(ctx).Symbol(symbol).AutoCloseType(autoCloseType).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

User's Force Orders(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string |  (optional)
	autoCloseType := models.UsersForceOrdersAutoCloseTypeParameterLiquidation // UsersForceOrdersAutoCloseTypeParameter | \"LIQUIDATION\" for liquidation orders, \"ADL\" for ADL orders. (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	limit := int64(100) // int64 | Default 100; max 1000 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

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
 **limit** | **int64** | Default 100; max 1000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**UsersForceOrdersResponse**](UsersForceOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

