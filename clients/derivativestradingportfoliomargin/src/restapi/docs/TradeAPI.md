# \TradeAPI

All URIs are relative to *https://papi.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**CancelAllCmOpenConditionalOrders**](TradeAPI.md#CancelAllCmOpenConditionalOrders) | **Delete** /papi/v1/cm/conditional/allOpenOrders | Cancel All CM Open Conditional Orders(TRADE)
[**CancelAllCmOpenOrders**](TradeAPI.md#CancelAllCmOpenOrders) | **Delete** /papi/v1/cm/allOpenOrders | Cancel All CM Open Orders(TRADE)
[**CancelAllUmOpenConditionalOrders**](TradeAPI.md#CancelAllUmOpenConditionalOrders) | **Delete** /papi/v1/um/conditional/allOpenOrders | Cancel All UM Open Conditional Orders (TRADE)
[**CancelAllUmOpenOrders**](TradeAPI.md#CancelAllUmOpenOrders) | **Delete** /papi/v1/um/allOpenOrders | Cancel All UM Open Orders(TRADE)
[**CancelCmConditionalOrder**](TradeAPI.md#CancelCmConditionalOrder) | **Delete** /papi/v1/cm/conditional/order | Cancel CM Conditional Order(TRADE)
[**CancelCmOrder**](TradeAPI.md#CancelCmOrder) | **Delete** /papi/v1/cm/order | Cancel CM Order(TRADE)
[**CancelMarginAccountAllOpenOrdersOnASymbol**](TradeAPI.md#CancelMarginAccountAllOpenOrdersOnASymbol) | **Delete** /papi/v1/margin/allOpenOrders | Cancel Margin Account All Open Orders on a Symbol(TRADE)
[**CancelMarginAccountOcoOrders**](TradeAPI.md#CancelMarginAccountOcoOrders) | **Delete** /papi/v1/margin/orderList | Cancel Margin Account OCO Orders(TRADE)
[**CancelMarginAccountOrder**](TradeAPI.md#CancelMarginAccountOrder) | **Delete** /papi/v1/margin/order | Cancel Margin Account Order(TRADE)
[**CancelUmConditionalOrder**](TradeAPI.md#CancelUmConditionalOrder) | **Delete** /papi/v1/um/conditional/order | Cancel UM Conditional Order(TRADE)
[**CancelUmOrder**](TradeAPI.md#CancelUmOrder) | **Delete** /papi/v1/um/order | Cancel UM Order(TRADE)
[**CmAccountTradeList**](TradeAPI.md#CmAccountTradeList) | **Get** /papi/v1/cm/userTrades | CM Account Trade List(USER_DATA)
[**CmPositionAdlQuantileEstimation**](TradeAPI.md#CmPositionAdlQuantileEstimation) | **Get** /papi/v1/cm/adlQuantile | CM Position ADL Quantile Estimation(USER_DATA)
[**GetUmFuturesBnbBurnStatus**](TradeAPI.md#GetUmFuturesBnbBurnStatus) | **Get** /papi/v1/um/feeBurn | Get UM Futures BNB Burn Status (USER_DATA)
[**MarginAccountBorrow**](TradeAPI.md#MarginAccountBorrow) | **Post** /papi/v1/marginLoan | Margin Account Borrow(MARGIN)
[**MarginAccountNewOco**](TradeAPI.md#MarginAccountNewOco) | **Post** /papi/v1/margin/order/oco | Margin Account New OCO(TRADE)
[**MarginAccountRepay**](TradeAPI.md#MarginAccountRepay) | **Post** /papi/v1/repayLoan | Margin Account Repay(MARGIN)
[**MarginAccountRepayDebt**](TradeAPI.md#MarginAccountRepayDebt) | **Post** /papi/v1/margin/repay-debt | Margin Account Repay Debt(TRADE)
[**MarginAccountTradeList**](TradeAPI.md#MarginAccountTradeList) | **Get** /papi/v1/margin/myTrades | Margin Account Trade List (USER_DATA)
[**ModifyCmOrder**](TradeAPI.md#ModifyCmOrder) | **Put** /papi/v1/cm/order | Modify CM Order(TRADE)
[**ModifyUmOrder**](TradeAPI.md#ModifyUmOrder) | **Put** /papi/v1/um/order | Modify UM Order(TRADE)
[**NewCmConditionalOrder**](TradeAPI.md#NewCmConditionalOrder) | **Post** /papi/v1/cm/conditional/order | New CM Conditional Order(TRADE)
[**NewCmOrder**](TradeAPI.md#NewCmOrder) | **Post** /papi/v1/cm/order | New CM Order(TRADE)
[**NewMarginOrder**](TradeAPI.md#NewMarginOrder) | **Post** /papi/v1/margin/order | New Margin Order(TRADE)
[**NewUmConditionalOrder**](TradeAPI.md#NewUmConditionalOrder) | **Post** /papi/v1/um/conditional/order | New UM Conditional Order (TRADE)
[**NewUmOrder**](TradeAPI.md#NewUmOrder) | **Post** /papi/v1/um/order | New UM Order (TRADE)
[**QueryAllCmConditionalOrders**](TradeAPI.md#QueryAllCmConditionalOrders) | **Get** /papi/v1/cm/conditional/allOrders | Query All CM Conditional Orders(USER_DATA)
[**QueryAllCmOrders**](TradeAPI.md#QueryAllCmOrders) | **Get** /papi/v1/cm/allOrders | Query All CM Orders (USER_DATA)
[**QueryAllCurrentCmOpenConditionalOrders**](TradeAPI.md#QueryAllCurrentCmOpenConditionalOrders) | **Get** /papi/v1/cm/conditional/openOrders | Query All Current CM Open Conditional Orders (USER_DATA)
[**QueryAllCurrentCmOpenOrders**](TradeAPI.md#QueryAllCurrentCmOpenOrders) | **Get** /papi/v1/cm/openOrders | Query All Current CM Open Orders(USER_DATA)
[**QueryAllCurrentUmOpenConditionalOrders**](TradeAPI.md#QueryAllCurrentUmOpenConditionalOrders) | **Get** /papi/v1/um/conditional/openOrders | Query All Current UM Open Conditional Orders(USER_DATA)
[**QueryAllCurrentUmOpenOrders**](TradeAPI.md#QueryAllCurrentUmOpenOrders) | **Get** /papi/v1/um/openOrders | Query All Current UM Open Orders(USER_DATA)
[**QueryAllMarginAccountOrders**](TradeAPI.md#QueryAllMarginAccountOrders) | **Get** /papi/v1/margin/allOrders | Query All Margin Account Orders (USER_DATA)
[**QueryAllUmConditionalOrders**](TradeAPI.md#QueryAllUmConditionalOrders) | **Get** /papi/v1/um/conditional/allOrders | Query All UM Conditional Orders(USER_DATA)
[**QueryAllUmOrders**](TradeAPI.md#QueryAllUmOrders) | **Get** /papi/v1/um/allOrders | Query All UM Orders(USER_DATA)
[**QueryCmConditionalOrderHistory**](TradeAPI.md#QueryCmConditionalOrderHistory) | **Get** /papi/v1/cm/conditional/orderHistory | Query CM Conditional Order History(USER_DATA)
[**QueryCmModifyOrderHistory**](TradeAPI.md#QueryCmModifyOrderHistory) | **Get** /papi/v1/cm/orderAmendment | Query CM Modify Order History(TRADE)
[**QueryCmOrder**](TradeAPI.md#QueryCmOrder) | **Get** /papi/v1/cm/order | Query CM Order(USER_DATA)
[**QueryCurrentCmOpenConditionalOrder**](TradeAPI.md#QueryCurrentCmOpenConditionalOrder) | **Get** /papi/v1/cm/conditional/openOrder | Query Current CM Open Conditional Order(USER_DATA)
[**QueryCurrentCmOpenOrder**](TradeAPI.md#QueryCurrentCmOpenOrder) | **Get** /papi/v1/cm/openOrder | Query Current CM Open Order (USER_DATA)
[**QueryCurrentMarginOpenOrder**](TradeAPI.md#QueryCurrentMarginOpenOrder) | **Get** /papi/v1/margin/openOrders | Query Current Margin Open Order (USER_DATA)
[**QueryCurrentUmOpenConditionalOrder**](TradeAPI.md#QueryCurrentUmOpenConditionalOrder) | **Get** /papi/v1/um/conditional/openOrder | Query Current UM Open Conditional Order(USER_DATA)
[**QueryCurrentUmOpenOrder**](TradeAPI.md#QueryCurrentUmOpenOrder) | **Get** /papi/v1/um/openOrder | Query Current UM Open Order(USER_DATA)
[**QueryMarginAccountOrder**](TradeAPI.md#QueryMarginAccountOrder) | **Get** /papi/v1/margin/order | Query Margin Account Order (USER_DATA)
[**QueryMarginAccountsAllOco**](TradeAPI.md#QueryMarginAccountsAllOco) | **Get** /papi/v1/margin/allOrderList | Query Margin Account&#39;s all OCO (USER_DATA)
[**QueryMarginAccountsOco**](TradeAPI.md#QueryMarginAccountsOco) | **Get** /papi/v1/margin/orderList | Query Margin Account&#39;s OCO (USER_DATA)
[**QueryMarginAccountsOpenOco**](TradeAPI.md#QueryMarginAccountsOpenOco) | **Get** /papi/v1/margin/openOrderList | Query Margin Account&#39;s Open OCO (USER_DATA)
[**QueryUmConditionalOrderHistory**](TradeAPI.md#QueryUmConditionalOrderHistory) | **Get** /papi/v1/um/conditional/orderHistory | Query UM Conditional Order History(USER_DATA)
[**QueryUmModifyOrderHistory**](TradeAPI.md#QueryUmModifyOrderHistory) | **Get** /papi/v1/um/orderAmendment | Query UM Modify Order History(TRADE)
[**QueryUmOrder**](TradeAPI.md#QueryUmOrder) | **Get** /papi/v1/um/order | Query UM Order (USER_DATA)
[**QueryUsersCmForceOrders**](TradeAPI.md#QueryUsersCmForceOrders) | **Get** /papi/v1/cm/forceOrders | Query User&#39;s CM Force Orders(USER_DATA)
[**QueryUsersMarginForceOrders**](TradeAPI.md#QueryUsersMarginForceOrders) | **Get** /papi/v1/margin/forceOrders | Query User&#39;s Margin Force Orders(USER_DATA)
[**QueryUsersUmForceOrders**](TradeAPI.md#QueryUsersUmForceOrders) | **Get** /papi/v1/um/forceOrders | Query User&#39;s UM Force Orders (USER_DATA)
[**ToggleBnbBurnOnUmFuturesTrade**](TradeAPI.md#ToggleBnbBurnOnUmFuturesTrade) | **Post** /papi/v1/um/feeBurn | Toggle BNB Burn On UM Futures Trade (TRADE)
[**UmAccountTradeList**](TradeAPI.md#UmAccountTradeList) | **Get** /papi/v1/um/userTrades | UM Account Trade List(USER_DATA)
[**UmPositionAdlQuantileEstimation**](TradeAPI.md#UmPositionAdlQuantileEstimation) | **Get** /papi/v1/um/adlQuantile | UM Position ADL Quantile Estimation(USER_DATA)


## CancelAllCmOpenConditionalOrders

> CancelAllCmOpenConditionalOrdersResponse CancelAllCmOpenConditionalOrders(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Cancel All CM Open Conditional Orders(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
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
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CancelAllCmOpenConditionalOrders(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CancelAllCmOpenConditionalOrders``: %v\n", err)
		return
	}

	// response from `CancelAllCmOpenConditionalOrders`: CancelAllCmOpenConditionalOrdersResponse
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

[**CancelAllCmOpenConditionalOrdersResponse**](CancelAllCmOpenConditionalOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CancelAllCmOpenOrders

> CancelAllCmOpenOrdersResponse CancelAllCmOpenOrders(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Cancel All CM Open Orders(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
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
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CancelAllCmOpenOrders(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CancelAllCmOpenOrders``: %v\n", err)
		return
	}

	// response from `CancelAllCmOpenOrders`: CancelAllCmOpenOrdersResponse
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

[**CancelAllCmOpenOrdersResponse**](CancelAllCmOpenOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CancelAllUmOpenConditionalOrders

> CancelAllUmOpenConditionalOrdersResponse CancelAllUmOpenConditionalOrders(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Cancel All UM Open Conditional Orders (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
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
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CancelAllUmOpenConditionalOrders(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CancelAllUmOpenConditionalOrders``: %v\n", err)
		return
	}

	// response from `CancelAllUmOpenConditionalOrders`: CancelAllUmOpenConditionalOrdersResponse
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

[**CancelAllUmOpenConditionalOrdersResponse**](CancelAllUmOpenConditionalOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CancelAllUmOpenOrders

> CancelAllUmOpenOrdersResponse CancelAllUmOpenOrders(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Cancel All UM Open Orders(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
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
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CancelAllUmOpenOrders(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CancelAllUmOpenOrders``: %v\n", err)
		return
	}

	// response from `CancelAllUmOpenOrders`: CancelAllUmOpenOrdersResponse
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

[**CancelAllUmOpenOrdersResponse**](CancelAllUmOpenOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CancelCmConditionalOrder

> CancelCmConditionalOrderResponse CancelCmConditionalOrder(ctx).Symbol(symbol).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()

Cancel CM Conditional Order(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	strategyId := int64(1) // int64 |  (optional)
	newClientStrategyId := "1" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CancelCmConditionalOrder(context.Background()).Symbol(symbol).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CancelCmConditionalOrder``: %v\n", err)
		return
	}

	// response from `CancelCmConditionalOrder`: CancelCmConditionalOrderResponse
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
 **strategyId** | **int64** |  | 
 **newClientStrategyId** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CancelCmConditionalOrderResponse**](CancelCmConditionalOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CancelCmOrder

> CancelCmOrderResponse CancelCmOrder(ctx).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Cancel CM Order(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
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
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CancelCmOrder(context.Background()).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CancelCmOrder``: %v\n", err)
		return
	}

	// response from `CancelCmOrder`: CancelCmOrderResponse
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

[**CancelCmOrderResponse**](CancelCmOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CancelMarginAccountAllOpenOrdersOnASymbol

> CancelMarginAccountAllOpenOrdersOnASymbolResponse CancelMarginAccountAllOpenOrdersOnASymbol(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Cancel Margin Account All Open Orders on a Symbol(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
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
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CancelMarginAccountAllOpenOrdersOnASymbol(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CancelMarginAccountAllOpenOrdersOnASymbol``: %v\n", err)
		return
	}

	// response from `CancelMarginAccountAllOpenOrdersOnASymbol`: CancelMarginAccountAllOpenOrdersOnASymbolResponse
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

[**CancelMarginAccountAllOpenOrdersOnASymbolResponse**](CancelMarginAccountAllOpenOrdersOnASymbolResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CancelMarginAccountOcoOrders

> CancelMarginAccountOcoOrdersResponse CancelMarginAccountOcoOrders(ctx).Symbol(symbol).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()

Cancel Margin Account OCO Orders(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	orderListId := int64(1) // int64 | Either `orderListId` or `listClientOrderId` must be provided (optional)
	listClientOrderId := "1" // string | Either `orderListId` or `listClientOrderId` must be provided (optional)
	newClientOrderId := "1" // string | Used to uniquely identify this cancel. Automatically generated by default (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CancelMarginAccountOcoOrders(context.Background()).Symbol(symbol).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CancelMarginAccountOcoOrders``: %v\n", err)
		return
	}

	// response from `CancelMarginAccountOcoOrders`: CancelMarginAccountOcoOrdersResponse
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
 **orderListId** | **int64** | Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided | 
 **listClientOrderId** | **string** | Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided | 
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default | 
 **recvWindow** | **int64** |  | 

### Return type

[**CancelMarginAccountOcoOrdersResponse**](CancelMarginAccountOcoOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CancelMarginAccountOrder

> CancelMarginAccountOrderResponse CancelMarginAccountOrder(ctx).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()

Cancel Margin Account Order(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	orderId := int64(1) // int64 |  (optional)
	origClientOrderId := "1" // string |  (optional)
	newClientOrderId := "1" // string | Used to uniquely identify this cancel. Automatically generated by default (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CancelMarginAccountOrder(context.Background()).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CancelMarginAccountOrder``: %v\n", err)
		return
	}

	// response from `CancelMarginAccountOrder`: CancelMarginAccountOrderResponse
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
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default | 
 **recvWindow** | **int64** |  | 

### Return type

[**CancelMarginAccountOrderResponse**](CancelMarginAccountOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CancelUmConditionalOrder

> CancelUmConditionalOrderResponse CancelUmConditionalOrder(ctx).Symbol(symbol).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()

Cancel UM Conditional Order(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	strategyId := int64(1) // int64 |  (optional)
	newClientStrategyId := "1" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CancelUmConditionalOrder(context.Background()).Symbol(symbol).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CancelUmConditionalOrder``: %v\n", err)
		return
	}

	// response from `CancelUmConditionalOrder`: CancelUmConditionalOrderResponse
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
 **strategyId** | **int64** |  | 
 **newClientStrategyId** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CancelUmConditionalOrderResponse**](CancelUmConditionalOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CancelUmOrder

> CancelUmOrderResponse CancelUmOrder(ctx).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Cancel UM Order(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
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
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CancelUmOrder(context.Background()).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CancelUmOrder``: %v\n", err)
		return
	}

	// response from `CancelUmOrder`: CancelUmOrderResponse
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

[**CancelUmOrderResponse**](CancelUmOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CmAccountTradeList

> CmAccountTradeListResponse CmAccountTradeList(ctx).Symbol(symbol).Pair(pair).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()

CM Account Trade List(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string |  (optional)
	pair := "pair_example" // string |  (optional)
	startTime := int64(1623319461670) // int64 | Timestamp in ms to get funding from INCLUSIVE. (optional)
	endTime := int64(1641782889000) // int64 | Timestamp in ms to get funding until INCLUSIVE. (optional)
	fromId := int64(1) // int64 | Trade id to fetch from. Default gets most recent trades. (optional)
	limit := int64(100) // int64 | Default 100; max 1000 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CmAccountTradeList(context.Background()).Symbol(symbol).Pair(pair).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CmAccountTradeList``: %v\n", err)
		return
	}

	// response from `CmAccountTradeList`: CmAccountTradeListResponse
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
 **startTime** | **int64** | Timestamp in ms to get funding from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding until INCLUSIVE. | 
 **fromId** | **int64** | Trade id to fetch from. Default gets most recent trades. | 
 **limit** | **int64** | Default 100; max 1000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**CmAccountTradeListResponse**](CmAccountTradeListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CmPositionAdlQuantileEstimation

> CmPositionAdlQuantileEstimationResponse CmPositionAdlQuantileEstimation(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

CM Position ADL Quantile Estimation(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
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
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CmPositionAdlQuantileEstimation(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CmPositionAdlQuantileEstimation``: %v\n", err)
		return
	}

	// response from `CmPositionAdlQuantileEstimation`: CmPositionAdlQuantileEstimationResponse
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

[**CmPositionAdlQuantileEstimationResponse**](CmPositionAdlQuantileEstimationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetUmFuturesBnbBurnStatus

> GetUmFuturesBnbBurnStatusResponse GetUmFuturesBnbBurnStatus(ctx).RecvWindow(recvWindow).Execute()

Get UM Futures BNB Burn Status (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.GetUmFuturesBnbBurnStatus(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.GetUmFuturesBnbBurnStatus``: %v\n", err)
		return
	}

	// response from `GetUmFuturesBnbBurnStatus`: GetUmFuturesBnbBurnStatusResponse
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

[**GetUmFuturesBnbBurnStatusResponse**](GetUmFuturesBnbBurnStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MarginAccountBorrow

> MarginAccountBorrowResponse MarginAccountBorrow(ctx).Asset(asset).Amount(amount).RecvWindow(recvWindow).Execute()

Margin Account Borrow(MARGIN)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	asset := "asset_example" // string | 
	amount := float32(1.0) // float32 | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.MarginAccountBorrow(context.Background()).Asset(asset).Amount(amount).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.MarginAccountBorrow``: %v\n", err)
		return
	}

	// response from `MarginAccountBorrow`: MarginAccountBorrowResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | 
 **amount** | **float32** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**MarginAccountBorrowResponse**](MarginAccountBorrowResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MarginAccountNewOco

> MarginAccountNewOcoResponse MarginAccountNewOco(ctx).Symbol(symbol).Side(side).Quantity(quantity).Price(price).StopPrice(stopPrice).ListClientOrderId(listClientOrderId).LimitClientOrderId(limitClientOrderId).LimitIcebergQty(limitIcebergQty).StopClientOrderId(stopClientOrderId).StopLimitPrice(stopLimitPrice).StopIcebergQty(stopIcebergQty).StopLimitTimeInForce(stopLimitTimeInForce).NewOrderRespType(newOrderRespType).SideEffectType(sideEffectType).RecvWindow(recvWindow).Execute()

Margin Account New OCO(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	side := models.NewCmConditionalOrderSideParameterBuy // NewCmConditionalOrderSideParameter | 
	quantity := float32(1.0) // float32 | Order quantity
	price := float32(1.0) // float32 | 
	stopPrice := float32(1.0) // float32 | 
	listClientOrderId := "1" // string | Either `orderListId` or `listClientOrderId` must be provided (optional)
	limitClientOrderId := "1" // string | A unique Id for the limit order (optional)
	limitIcebergQty := float32(1.0) // float32 |  (optional)
	stopClientOrderId := "1" // string | A unique Id for the stop loss/stop loss limit leg (optional)
	stopLimitPrice := float32(1.0) // float32 | If provided, stopLimitTimeInForce is required. (optional)
	stopIcebergQty := float32(1.0) // float32 |  (optional)
	stopLimitTimeInForce := models.MarginAccountNewOcoStopLimitTimeInForceParameterGtc // MarginAccountNewOcoStopLimitTimeInForceParameter | Valid values are `GTC/FOK/IOC` (optional)
	newOrderRespType := models.NewCmOrderNewOrderRespTypeParameterAck // NewCmOrderNewOrderRespTypeParameter | \"ACK\", \"RESULT\", default \"ACK\" (optional)
	sideEffectType := models.NewMarginOrderSideEffectTypeParameterNoSideEffect // NewMarginOrderSideEffectTypeParameter | NO_SIDE_EFFECT, MARGIN_BUY, AUTO_REPAY; default NO_SIDE_EFFECT. (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.MarginAccountNewOco(context.Background()).Symbol(symbol).Side(side).Quantity(quantity).Price(price).StopPrice(stopPrice).ListClientOrderId(listClientOrderId).LimitClientOrderId(limitClientOrderId).LimitIcebergQty(limitIcebergQty).StopClientOrderId(stopClientOrderId).StopLimitPrice(stopLimitPrice).StopIcebergQty(stopIcebergQty).StopLimitTimeInForce(stopLimitTimeInForce).NewOrderRespType(newOrderRespType).SideEffectType(sideEffectType).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.MarginAccountNewOco``: %v\n", err)
		return
	}

	// response from `MarginAccountNewOco`: MarginAccountNewOcoResponse
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
 **side** | [**NewCmConditionalOrderSideParameter**](NewCmConditionalOrderSideParameter.md) |  | 
 **quantity** | **float32** | Order quantity | 
 **price** | **float32** |  | 
 **stopPrice** | **float32** |  | 
 **listClientOrderId** | **string** | Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided | 
 **limitClientOrderId** | **string** | A unique Id for the limit order | 
 **limitIcebergQty** | **float32** |  | 
 **stopClientOrderId** | **string** | A unique Id for the stop loss/stop loss limit leg | 
 **stopLimitPrice** | **float32** | If provided, stopLimitTimeInForce is required. | 
 **stopIcebergQty** | **float32** |  | 
 **stopLimitTimeInForce** | [**MarginAccountNewOcoStopLimitTimeInForceParameter**](MarginAccountNewOcoStopLimitTimeInForceParameter.md) | Valid values are &#x60;GTC/FOK/IOC&#x60; | 
 **newOrderRespType** | [**NewCmOrderNewOrderRespTypeParameter**](NewCmOrderNewOrderRespTypeParameter.md) | \&quot;ACK\&quot;, \&quot;RESULT\&quot;, default \&quot;ACK\&quot; | 
 **sideEffectType** | [**NewMarginOrderSideEffectTypeParameter**](NewMarginOrderSideEffectTypeParameter.md) | NO_SIDE_EFFECT, MARGIN_BUY, AUTO_REPAY; default NO_SIDE_EFFECT. | 
 **recvWindow** | **int64** |  | 

### Return type

[**MarginAccountNewOcoResponse**](MarginAccountNewOcoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MarginAccountRepay

> MarginAccountRepayResponse MarginAccountRepay(ctx).Asset(asset).Amount(amount).RecvWindow(recvWindow).Execute()

Margin Account Repay(MARGIN)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	asset := "asset_example" // string | 
	amount := float32(1.0) // float32 | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.MarginAccountRepay(context.Background()).Asset(asset).Amount(amount).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.MarginAccountRepay``: %v\n", err)
		return
	}

	// response from `MarginAccountRepay`: MarginAccountRepayResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | 
 **amount** | **float32** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**MarginAccountRepayResponse**](MarginAccountRepayResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MarginAccountRepayDebt

> MarginAccountRepayDebtResponse MarginAccountRepayDebt(ctx).Asset(asset).Amount(amount).SpecifyRepayAssets(specifyRepayAssets).RecvWindow(recvWindow).Execute()

Margin Account Repay Debt(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	asset := "asset_example" // string | 
	amount := "amount_example" // string |  (optional)
	specifyRepayAssets := "specifyRepayAssets_example" // string | Specific asset list to repay debt; Can be added in batch, separated by commas (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.MarginAccountRepayDebt(context.Background()).Asset(asset).Amount(amount).SpecifyRepayAssets(specifyRepayAssets).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.MarginAccountRepayDebt``: %v\n", err)
		return
	}

	// response from `MarginAccountRepayDebt`: MarginAccountRepayDebtResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | 
 **amount** | **string** |  | 
 **specifyRepayAssets** | **string** | Specific asset list to repay debt; Can be added in batch, separated by commas | 
 **recvWindow** | **int64** |  | 

### Return type

[**MarginAccountRepayDebtResponse**](MarginAccountRepayDebtResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MarginAccountTradeList

> MarginAccountTradeListResponse MarginAccountTradeList(ctx).Symbol(symbol).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()

Margin Account Trade List (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	orderId := int64(1) // int64 |  (optional)
	startTime := int64(1623319461670) // int64 | Timestamp in ms to get funding from INCLUSIVE. (optional)
	endTime := int64(1641782889000) // int64 | Timestamp in ms to get funding until INCLUSIVE. (optional)
	fromId := int64(1) // int64 | Trade id to fetch from. Default gets most recent trades. (optional)
	limit := int64(100) // int64 | Default 100; max 1000 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.MarginAccountTradeList(context.Background()).Symbol(symbol).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.MarginAccountTradeList``: %v\n", err)
		return
	}

	// response from `MarginAccountTradeList`: MarginAccountTradeListResponse
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
 **startTime** | **int64** | Timestamp in ms to get funding from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding until INCLUSIVE. | 
 **fromId** | **int64** | Trade id to fetch from. Default gets most recent trades. | 
 **limit** | **int64** | Default 100; max 1000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**MarginAccountTradeListResponse**](MarginAccountTradeListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## ModifyCmOrder

> ModifyCmOrderResponse ModifyCmOrder(ctx).Symbol(symbol).Side(side).Quantity(quantity).Price(price).OrderId(orderId).OrigClientOrderId(origClientOrderId).PriceMatch(priceMatch).RecvWindow(recvWindow).Execute()

Modify CM Order(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	side := models.NewCmConditionalOrderSideParameterBuy // NewCmConditionalOrderSideParameter | 
	quantity := float32(1.0) // float32 | Order quantity
	price := float32(1.0) // float32 | 
	orderId := int64(1) // int64 |  (optional)
	origClientOrderId := "1" // string |  (optional)
	priceMatch := models.ModifyCmOrderPriceMatchParameterNone // ModifyCmOrderPriceMatchParameter | only avaliable for `LIMIT`/`STOP`/`TAKE_PROFIT` order; can be set to `OPPONENT`/ `OPPONENT_5`/ `OPPONENT_10`/ `OPPONENT_20`: /`QUEUE`/ `QUEUE_5`/ `QUEUE_10`/ `QUEUE_20`; Can't be passed together with `price` (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.ModifyCmOrder(context.Background()).Symbol(symbol).Side(side).Quantity(quantity).Price(price).OrderId(orderId).OrigClientOrderId(origClientOrderId).PriceMatch(priceMatch).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.ModifyCmOrder``: %v\n", err)
		return
	}

	// response from `ModifyCmOrder`: ModifyCmOrderResponse
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
 **side** | [**NewCmConditionalOrderSideParameter**](NewCmConditionalOrderSideParameter.md) |  | 
 **quantity** | **float32** | Order quantity | 
 **price** | **float32** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | 
 **priceMatch** | [**ModifyCmOrderPriceMatchParameter**](ModifyCmOrderPriceMatchParameter.md) | only avaliable for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; can be set to &#x60;OPPONENT&#x60;/ &#x60;OPPONENT_5&#x60;/ &#x60;OPPONENT_10&#x60;/ &#x60;OPPONENT_20&#x60;: /&#x60;QUEUE&#x60;/ &#x60;QUEUE_5&#x60;/ &#x60;QUEUE_10&#x60;/ &#x60;QUEUE_20&#x60;; Can&#39;t be passed together with &#x60;price&#x60; | 
 **recvWindow** | **int64** |  | 

### Return type

[**ModifyCmOrderResponse**](ModifyCmOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## ModifyUmOrder

> ModifyUmOrderResponse ModifyUmOrder(ctx).Symbol(symbol).Side(side).Quantity(quantity).Price(price).OrderId(orderId).OrigClientOrderId(origClientOrderId).PriceMatch(priceMatch).RecvWindow(recvWindow).Execute()

Modify UM Order(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	side := models.NewCmConditionalOrderSideParameterBuy // NewCmConditionalOrderSideParameter | 
	quantity := float32(1.0) // float32 | Order quantity
	price := float32(1.0) // float32 | 
	orderId := int64(1) // int64 |  (optional)
	origClientOrderId := "1" // string |  (optional)
	priceMatch := models.ModifyCmOrderPriceMatchParameterNone // ModifyCmOrderPriceMatchParameter | only avaliable for `LIMIT`/`STOP`/`TAKE_PROFIT` order; can be set to `OPPONENT`/ `OPPONENT_5`/ `OPPONENT_10`/ `OPPONENT_20`: /`QUEUE`/ `QUEUE_5`/ `QUEUE_10`/ `QUEUE_20`; Can't be passed together with `price` (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.ModifyUmOrder(context.Background()).Symbol(symbol).Side(side).Quantity(quantity).Price(price).OrderId(orderId).OrigClientOrderId(origClientOrderId).PriceMatch(priceMatch).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.ModifyUmOrder``: %v\n", err)
		return
	}

	// response from `ModifyUmOrder`: ModifyUmOrderResponse
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
 **side** | [**NewCmConditionalOrderSideParameter**](NewCmConditionalOrderSideParameter.md) |  | 
 **quantity** | **float32** | Order quantity | 
 **price** | **float32** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | 
 **priceMatch** | [**ModifyCmOrderPriceMatchParameter**](ModifyCmOrderPriceMatchParameter.md) | only avaliable for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; can be set to &#x60;OPPONENT&#x60;/ &#x60;OPPONENT_5&#x60;/ &#x60;OPPONENT_10&#x60;/ &#x60;OPPONENT_20&#x60;: /&#x60;QUEUE&#x60;/ &#x60;QUEUE_5&#x60;/ &#x60;QUEUE_10&#x60;/ &#x60;QUEUE_20&#x60;; Can&#39;t be passed together with &#x60;price&#x60; | 
 **recvWindow** | **int64** |  | 

### Return type

[**ModifyUmOrderResponse**](ModifyUmOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## NewCmConditionalOrder

> NewCmConditionalOrderResponse NewCmConditionalOrder(ctx).Symbol(symbol).Side(side).StrategyType(strategyType).PositionSide(positionSide).TimeInForce(timeInForce).Quantity(quantity).ReduceOnly(reduceOnly).Price(price).WorkingType(workingType).PriceProtect(priceProtect).NewClientStrategyId(newClientStrategyId).StopPrice(stopPrice).ActivationPrice(activationPrice).CallbackRate(callbackRate).RecvWindow(recvWindow).Execute()

New CM Conditional Order(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	side := models.NewCmConditionalOrderSideParameterBuy // NewCmConditionalOrderSideParameter | 
	strategyType := models.NewCmConditionalOrderStrategyTypeParameterStop // NewCmConditionalOrderStrategyTypeParameter | \"STOP\", \"STOP_MARKET\", \"TAKE_PROFIT\", \"TAKE_PROFIT_MARKET\", and \"TRAILING_STOP_MARKET\"
	positionSide := models.NewCmConditionalOrderPositionSideParameterBoth // NewCmConditionalOrderPositionSideParameter | Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent in Hedge Mode. (optional)
	timeInForce := models.NewCmConditionalOrderTimeInForceParameterGtc // NewCmConditionalOrderTimeInForceParameter |  (optional)
	quantity := float32(1.0) // float32 |  (optional)
	reduceOnly := "false" // string | \"true\" or \"false\". default \"false\". Cannot be sent in Hedge Mode . (optional)
	price := float32(1.0) // float32 |  (optional)
	workingType := models.NewCmConditionalOrderWorkingTypeParameterMarkPrice // NewCmConditionalOrderWorkingTypeParameter | stopPrice triggered by: \"MARK_PRICE\", \"CONTRACT_PRICE\". Default \"CONTRACT_PRICE\" (optional)
	priceProtect := "false" // string | \"TRUE\" or \"FALSE\", default \"FALSE\". Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders (optional)
	newClientStrategyId := "1" // string |  (optional)
	stopPrice := float32(1.0) // float32 | Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders. (optional)
	activationPrice := float32(1.0) // float32 | Used with `TRAILING_STOP_MARKET` orders, default as the mark price (optional)
	callbackRate := float32(1.0) // float32 | Used with `TRAILING_STOP_MARKET` orders, min 0.1, max 5 where 1 for 1% (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.NewCmConditionalOrder(context.Background()).Symbol(symbol).Side(side).StrategyType(strategyType).PositionSide(positionSide).TimeInForce(timeInForce).Quantity(quantity).ReduceOnly(reduceOnly).Price(price).WorkingType(workingType).PriceProtect(priceProtect).NewClientStrategyId(newClientStrategyId).StopPrice(stopPrice).ActivationPrice(activationPrice).CallbackRate(callbackRate).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.NewCmConditionalOrder``: %v\n", err)
		return
	}

	// response from `NewCmConditionalOrder`: NewCmConditionalOrderResponse
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
 **side** | [**NewCmConditionalOrderSideParameter**](NewCmConditionalOrderSideParameter.md) |  | 
 **strategyType** | [**NewCmConditionalOrderStrategyTypeParameter**](NewCmConditionalOrderStrategyTypeParameter.md) | \&quot;STOP\&quot;, \&quot;STOP_MARKET\&quot;, \&quot;TAKE_PROFIT\&quot;, \&quot;TAKE_PROFIT_MARKET\&quot;, and \&quot;TRAILING_STOP_MARKET\&quot; | 
 **positionSide** | [**NewCmConditionalOrderPositionSideParameter**](NewCmConditionalOrderPositionSideParameter.md) | Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent in Hedge Mode. | 
 **timeInForce** | [**NewCmConditionalOrderTimeInForceParameter**](NewCmConditionalOrderTimeInForceParameter.md) |  | 
 **quantity** | **float32** |  | 
 **reduceOnly** | **string** | \&quot;true\&quot; or \&quot;false\&quot;. default \&quot;false\&quot;. Cannot be sent in Hedge Mode . | 
 **price** | **float32** |  | 
 **workingType** | [**NewCmConditionalOrderWorkingTypeParameter**](NewCmConditionalOrderWorkingTypeParameter.md) | stopPrice triggered by: \&quot;MARK_PRICE\&quot;, \&quot;CONTRACT_PRICE\&quot;. Default \&quot;CONTRACT_PRICE\&quot; | 
 **priceProtect** | **string** | \&quot;TRUE\&quot; or \&quot;FALSE\&quot;, default \&quot;FALSE\&quot;. Used with &#x60;STOP/STOP_MARKET&#x60; or &#x60;TAKE_PROFIT/TAKE_PROFIT_MARKET&#x60; orders | 
 **newClientStrategyId** | **string** |  | 
 **stopPrice** | **float32** | Used with &#x60;STOP/STOP_MARKET&#x60; or &#x60;TAKE_PROFIT/TAKE_PROFIT_MARKET&#x60; orders. | 
 **activationPrice** | **float32** | Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, default as the mark price | 
 **callbackRate** | **float32** | Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, min 0.1, max 5 where 1 for 1% | 
 **recvWindow** | **int64** |  | 

### Return type

[**NewCmConditionalOrderResponse**](NewCmConditionalOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## NewCmOrder

> NewCmOrderResponse NewCmOrder(ctx).Symbol(symbol).Side(side).Type(type_).PositionSide(positionSide).TimeInForce(timeInForce).Quantity(quantity).ReduceOnly(reduceOnly).Price(price).PriceMatch(priceMatch).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).RecvWindow(recvWindow).Execute()

New CM Order(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	side := models.NewCmConditionalOrderSideParameterBuy // NewCmConditionalOrderSideParameter | 
	type_ := models.NewCmOrderTypeParameterLimit // NewCmOrderTypeParameter | `LIMIT`, `MARKET`
	positionSide := models.NewCmConditionalOrderPositionSideParameterBoth // NewCmConditionalOrderPositionSideParameter | Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent in Hedge Mode. (optional)
	timeInForce := models.NewCmConditionalOrderTimeInForceParameterGtc // NewCmConditionalOrderTimeInForceParameter |  (optional)
	quantity := float32(1.0) // float32 |  (optional)
	reduceOnly := "false" // string | \"true\" or \"false\". default \"false\". Cannot be sent in Hedge Mode . (optional)
	price := float32(1.0) // float32 |  (optional)
	priceMatch := models.ModifyCmOrderPriceMatchParameterNone // ModifyCmOrderPriceMatchParameter | only avaliable for `LIMIT`/`STOP`/`TAKE_PROFIT` order; can be set to `OPPONENT`/ `OPPONENT_5`/ `OPPONENT_10`/ `OPPONENT_20`: /`QUEUE`/ `QUEUE_5`/ `QUEUE_10`/ `QUEUE_20`; Can't be passed together with `price` (optional)
	newClientOrderId := "1" // string | Used to uniquely identify this cancel. Automatically generated by default (optional)
	newOrderRespType := models.NewCmOrderNewOrderRespTypeParameterAck // NewCmOrderNewOrderRespTypeParameter | \"ACK\", \"RESULT\", default \"ACK\" (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.NewCmOrder(context.Background()).Symbol(symbol).Side(side).Type(type_).PositionSide(positionSide).TimeInForce(timeInForce).Quantity(quantity).ReduceOnly(reduceOnly).Price(price).PriceMatch(priceMatch).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.NewCmOrder``: %v\n", err)
		return
	}

	// response from `NewCmOrder`: NewCmOrderResponse
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
 **side** | [**NewCmConditionalOrderSideParameter**](NewCmConditionalOrderSideParameter.md) |  | 
 **type_** | [**NewCmOrderTypeParameter**](NewCmOrderTypeParameter.md) | &#x60;LIMIT&#x60;, &#x60;MARKET&#x60; | 
 **positionSide** | [**NewCmConditionalOrderPositionSideParameter**](NewCmConditionalOrderPositionSideParameter.md) | Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent in Hedge Mode. | 
 **timeInForce** | [**NewCmConditionalOrderTimeInForceParameter**](NewCmConditionalOrderTimeInForceParameter.md) |  | 
 **quantity** | **float32** |  | 
 **reduceOnly** | **string** | \&quot;true\&quot; or \&quot;false\&quot;. default \&quot;false\&quot;. Cannot be sent in Hedge Mode . | 
 **price** | **float32** |  | 
 **priceMatch** | [**ModifyCmOrderPriceMatchParameter**](ModifyCmOrderPriceMatchParameter.md) | only avaliable for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; can be set to &#x60;OPPONENT&#x60;/ &#x60;OPPONENT_5&#x60;/ &#x60;OPPONENT_10&#x60;/ &#x60;OPPONENT_20&#x60;: /&#x60;QUEUE&#x60;/ &#x60;QUEUE_5&#x60;/ &#x60;QUEUE_10&#x60;/ &#x60;QUEUE_20&#x60;; Can&#39;t be passed together with &#x60;price&#x60; | 
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default | 
 **newOrderRespType** | [**NewCmOrderNewOrderRespTypeParameter**](NewCmOrderNewOrderRespTypeParameter.md) | \&quot;ACK\&quot;, \&quot;RESULT\&quot;, default \&quot;ACK\&quot; | 
 **recvWindow** | **int64** |  | 

### Return type

[**NewCmOrderResponse**](NewCmOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## NewMarginOrder

> NewMarginOrderResponse NewMarginOrder(ctx).Symbol(symbol).Side(side).Type(type_).Quantity(quantity).QuoteOrderQty(quoteOrderQty).Price(price).StopPrice(stopPrice).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).IcebergQty(icebergQty).SideEffectType(sideEffectType).TimeInForce(timeInForce).SelfTradePreventionMode(selfTradePreventionMode).AutoRepayAtCancel(autoRepayAtCancel).RecvWindow(recvWindow).Execute()

New Margin Order(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	side := models.NewCmConditionalOrderSideParameterBuy // NewCmConditionalOrderSideParameter | 
	type_ := models.NewCmOrderTypeParameterLimit // NewCmOrderTypeParameter | `LIMIT`, `MARKET`
	quantity := float32(1.0) // float32 |  (optional)
	quoteOrderQty := float32(1.0) // float32 |  (optional)
	price := float32(1.0) // float32 |  (optional)
	stopPrice := float32(1.0) // float32 | Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders. (optional)
	newClientOrderId := "1" // string | Used to uniquely identify this cancel. Automatically generated by default (optional)
	newOrderRespType := models.NewCmOrderNewOrderRespTypeParameterAck // NewCmOrderNewOrderRespTypeParameter | \"ACK\", \"RESULT\", default \"ACK\" (optional)
	icebergQty := float32(1.0) // float32 | Used with `LIMIT`, `STOP_LOSS_LIMIT`, and `TAKE_PROFIT_LIMIT` to create an iceberg order (optional)
	sideEffectType := models.NewMarginOrderSideEffectTypeParameterNoSideEffect // NewMarginOrderSideEffectTypeParameter | NO_SIDE_EFFECT, MARGIN_BUY, AUTO_REPAY; default NO_SIDE_EFFECT. (optional)
	timeInForce := models.NewCmConditionalOrderTimeInForceParameterGtc // NewCmConditionalOrderTimeInForceParameter |  (optional)
	selfTradePreventionMode := models.NewMarginOrderSelfTradePreventionModeParameterNone // NewMarginOrderSelfTradePreventionModeParameter | `NONE`:No STP / `EXPIRE_TAKER`:expire taker order when STP triggers/ `EXPIRE_MAKER`:expire taker order when STP triggers/ `EXPIRE_BOTH`:expire both orders when STP triggers (optional)
	autoRepayAtCancel := true // bool | Only when MARGIN_BUY or AUTO_BORROW_REPAY order takes effect, true means that the debt generated by the order needs to be repay after the order is cancelled. The default is true (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.NewMarginOrder(context.Background()).Symbol(symbol).Side(side).Type(type_).Quantity(quantity).QuoteOrderQty(quoteOrderQty).Price(price).StopPrice(stopPrice).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).IcebergQty(icebergQty).SideEffectType(sideEffectType).TimeInForce(timeInForce).SelfTradePreventionMode(selfTradePreventionMode).AutoRepayAtCancel(autoRepayAtCancel).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.NewMarginOrder``: %v\n", err)
		return
	}

	// response from `NewMarginOrder`: NewMarginOrderResponse
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
 **side** | [**NewCmConditionalOrderSideParameter**](NewCmConditionalOrderSideParameter.md) |  | 
 **type_** | [**NewCmOrderTypeParameter**](NewCmOrderTypeParameter.md) | &#x60;LIMIT&#x60;, &#x60;MARKET&#x60; | 
 **quantity** | **float32** |  | 
 **quoteOrderQty** | **float32** |  | 
 **price** | **float32** |  | 
 **stopPrice** | **float32** | Used with &#x60;STOP/STOP_MARKET&#x60; or &#x60;TAKE_PROFIT/TAKE_PROFIT_MARKET&#x60; orders. | 
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default | 
 **newOrderRespType** | [**NewCmOrderNewOrderRespTypeParameter**](NewCmOrderNewOrderRespTypeParameter.md) | \&quot;ACK\&quot;, \&quot;RESULT\&quot;, default \&quot;ACK\&quot; | 
 **icebergQty** | **float32** | Used with &#x60;LIMIT&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; to create an iceberg order | 
 **sideEffectType** | [**NewMarginOrderSideEffectTypeParameter**](NewMarginOrderSideEffectTypeParameter.md) | NO_SIDE_EFFECT, MARGIN_BUY, AUTO_REPAY; default NO_SIDE_EFFECT. | 
 **timeInForce** | [**NewCmConditionalOrderTimeInForceParameter**](NewCmConditionalOrderTimeInForceParameter.md) |  | 
 **selfTradePreventionMode** | [**NewMarginOrderSelfTradePreventionModeParameter**](NewMarginOrderSelfTradePreventionModeParameter.md) | &#x60;NONE&#x60;:No STP / &#x60;EXPIRE_TAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_MAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_BOTH&#x60;:expire both orders when STP triggers | 
 **autoRepayAtCancel** | **bool** | Only when MARGIN_BUY or AUTO_BORROW_REPAY order takes effect, true means that the debt generated by the order needs to be repay after the order is cancelled. The default is true | 
 **recvWindow** | **int64** |  | 

### Return type

[**NewMarginOrderResponse**](NewMarginOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## NewUmConditionalOrder

> NewUmConditionalOrderResponse NewUmConditionalOrder(ctx).Symbol(symbol).Side(side).StrategyType(strategyType).PositionSide(positionSide).TimeInForce(timeInForce).Quantity(quantity).ReduceOnly(reduceOnly).Price(price).WorkingType(workingType).PriceProtect(priceProtect).NewClientStrategyId(newClientStrategyId).StopPrice(stopPrice).ActivationPrice(activationPrice).CallbackRate(callbackRate).PriceMatch(priceMatch).SelfTradePreventionMode(selfTradePreventionMode).GoodTillDate(goodTillDate).RecvWindow(recvWindow).Execute()

New UM Conditional Order (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	side := models.NewCmConditionalOrderSideParameterBuy // NewCmConditionalOrderSideParameter | 
	strategyType := models.NewCmConditionalOrderStrategyTypeParameterStop // NewCmConditionalOrderStrategyTypeParameter | \"STOP\", \"STOP_MARKET\", \"TAKE_PROFIT\", \"TAKE_PROFIT_MARKET\", and \"TRAILING_STOP_MARKET\"
	positionSide := models.NewCmConditionalOrderPositionSideParameterBoth // NewCmConditionalOrderPositionSideParameter | Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent in Hedge Mode. (optional)
	timeInForce := models.NewCmConditionalOrderTimeInForceParameterGtc // NewCmConditionalOrderTimeInForceParameter |  (optional)
	quantity := float32(1.0) // float32 |  (optional)
	reduceOnly := "false" // string | \"true\" or \"false\". default \"false\". Cannot be sent in Hedge Mode . (optional)
	price := float32(1.0) // float32 |  (optional)
	workingType := models.NewCmConditionalOrderWorkingTypeParameterMarkPrice // NewCmConditionalOrderWorkingTypeParameter | stopPrice triggered by: \"MARK_PRICE\", \"CONTRACT_PRICE\". Default \"CONTRACT_PRICE\" (optional)
	priceProtect := "false" // string | \"TRUE\" or \"FALSE\", default \"FALSE\". Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders (optional)
	newClientStrategyId := "1" // string |  (optional)
	stopPrice := float32(1.0) // float32 | Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders. (optional)
	activationPrice := float32(1.0) // float32 | Used with `TRAILING_STOP_MARKET` orders, default as the mark price (optional)
	callbackRate := float32(1.0) // float32 | Used with `TRAILING_STOP_MARKET` orders, min 0.1, max 5 where 1 for 1% (optional)
	priceMatch := models.ModifyCmOrderPriceMatchParameterNone // ModifyCmOrderPriceMatchParameter | only avaliable for `LIMIT`/`STOP`/`TAKE_PROFIT` order; can be set to `OPPONENT`/ `OPPONENT_5`/ `OPPONENT_10`/ `OPPONENT_20`: /`QUEUE`/ `QUEUE_5`/ `QUEUE_10`/ `QUEUE_20`; Can't be passed together with `price` (optional)
	selfTradePreventionMode := models.NewMarginOrderSelfTradePreventionModeParameterNone // NewMarginOrderSelfTradePreventionModeParameter | `NONE`:No STP / `EXPIRE_TAKER`:expire taker order when STP triggers/ `EXPIRE_MAKER`:expire taker order when STP triggers/ `EXPIRE_BOTH`:expire both orders when STP triggers (optional)
	goodTillDate := int64(789) // int64 | order cancel time for timeInForce `GTD`, mandatory when `timeInforce` set to `GTD`; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000Mode. It must be sent in Hedge Mode. (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.NewUmConditionalOrder(context.Background()).Symbol(symbol).Side(side).StrategyType(strategyType).PositionSide(positionSide).TimeInForce(timeInForce).Quantity(quantity).ReduceOnly(reduceOnly).Price(price).WorkingType(workingType).PriceProtect(priceProtect).NewClientStrategyId(newClientStrategyId).StopPrice(stopPrice).ActivationPrice(activationPrice).CallbackRate(callbackRate).PriceMatch(priceMatch).SelfTradePreventionMode(selfTradePreventionMode).GoodTillDate(goodTillDate).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.NewUmConditionalOrder``: %v\n", err)
		return
	}

	// response from `NewUmConditionalOrder`: NewUmConditionalOrderResponse
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
 **side** | [**NewCmConditionalOrderSideParameter**](NewCmConditionalOrderSideParameter.md) |  | 
 **strategyType** | [**NewCmConditionalOrderStrategyTypeParameter**](NewCmConditionalOrderStrategyTypeParameter.md) | \&quot;STOP\&quot;, \&quot;STOP_MARKET\&quot;, \&quot;TAKE_PROFIT\&quot;, \&quot;TAKE_PROFIT_MARKET\&quot;, and \&quot;TRAILING_STOP_MARKET\&quot; | 
 **positionSide** | [**NewCmConditionalOrderPositionSideParameter**](NewCmConditionalOrderPositionSideParameter.md) | Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent in Hedge Mode. | 
 **timeInForce** | [**NewCmConditionalOrderTimeInForceParameter**](NewCmConditionalOrderTimeInForceParameter.md) |  | 
 **quantity** | **float32** |  | 
 **reduceOnly** | **string** | \&quot;true\&quot; or \&quot;false\&quot;. default \&quot;false\&quot;. Cannot be sent in Hedge Mode . | 
 **price** | **float32** |  | 
 **workingType** | [**NewCmConditionalOrderWorkingTypeParameter**](NewCmConditionalOrderWorkingTypeParameter.md) | stopPrice triggered by: \&quot;MARK_PRICE\&quot;, \&quot;CONTRACT_PRICE\&quot;. Default \&quot;CONTRACT_PRICE\&quot; | 
 **priceProtect** | **string** | \&quot;TRUE\&quot; or \&quot;FALSE\&quot;, default \&quot;FALSE\&quot;. Used with &#x60;STOP/STOP_MARKET&#x60; or &#x60;TAKE_PROFIT/TAKE_PROFIT_MARKET&#x60; orders | 
 **newClientStrategyId** | **string** |  | 
 **stopPrice** | **float32** | Used with &#x60;STOP/STOP_MARKET&#x60; or &#x60;TAKE_PROFIT/TAKE_PROFIT_MARKET&#x60; orders. | 
 **activationPrice** | **float32** | Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, default as the mark price | 
 **callbackRate** | **float32** | Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, min 0.1, max 5 where 1 for 1% | 
 **priceMatch** | [**ModifyCmOrderPriceMatchParameter**](ModifyCmOrderPriceMatchParameter.md) | only avaliable for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; can be set to &#x60;OPPONENT&#x60;/ &#x60;OPPONENT_5&#x60;/ &#x60;OPPONENT_10&#x60;/ &#x60;OPPONENT_20&#x60;: /&#x60;QUEUE&#x60;/ &#x60;QUEUE_5&#x60;/ &#x60;QUEUE_10&#x60;/ &#x60;QUEUE_20&#x60;; Can&#39;t be passed together with &#x60;price&#x60; | 
 **selfTradePreventionMode** | [**NewMarginOrderSelfTradePreventionModeParameter**](NewMarginOrderSelfTradePreventionModeParameter.md) | &#x60;NONE&#x60;:No STP / &#x60;EXPIRE_TAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_MAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_BOTH&#x60;:expire both orders when STP triggers | 
 **goodTillDate** | **int64** | order cancel time for timeInForce &#x60;GTD&#x60;, mandatory when &#x60;timeInforce&#x60; set to &#x60;GTD&#x60;; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000Mode. It must be sent in Hedge Mode. | 
 **recvWindow** | **int64** |  | 

### Return type

[**NewUmConditionalOrderResponse**](NewUmConditionalOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## NewUmOrder

> NewUmOrderResponse NewUmOrder(ctx).Symbol(symbol).Side(side).Type(type_).PositionSide(positionSide).TimeInForce(timeInForce).Quantity(quantity).ReduceOnly(reduceOnly).Price(price).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).PriceMatch(priceMatch).SelfTradePreventionMode(selfTradePreventionMode).GoodTillDate(goodTillDate).RecvWindow(recvWindow).Execute()

New UM Order (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	side := models.NewCmConditionalOrderSideParameterBuy // NewCmConditionalOrderSideParameter | 
	type_ := models.NewCmOrderTypeParameterLimit // NewCmOrderTypeParameter | `LIMIT`, `MARKET`
	positionSide := models.NewCmConditionalOrderPositionSideParameterBoth // NewCmConditionalOrderPositionSideParameter | Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent in Hedge Mode. (optional)
	timeInForce := models.NewCmConditionalOrderTimeInForceParameterGtc // NewCmConditionalOrderTimeInForceParameter |  (optional)
	quantity := float32(1.0) // float32 |  (optional)
	reduceOnly := "false" // string | \"true\" or \"false\". default \"false\". Cannot be sent in Hedge Mode . (optional)
	price := float32(1.0) // float32 |  (optional)
	newClientOrderId := "1" // string | Used to uniquely identify this cancel. Automatically generated by default (optional)
	newOrderRespType := models.NewCmOrderNewOrderRespTypeParameterAck // NewCmOrderNewOrderRespTypeParameter | \"ACK\", \"RESULT\", default \"ACK\" (optional)
	priceMatch := models.ModifyCmOrderPriceMatchParameterNone // ModifyCmOrderPriceMatchParameter | only avaliable for `LIMIT`/`STOP`/`TAKE_PROFIT` order; can be set to `OPPONENT`/ `OPPONENT_5`/ `OPPONENT_10`/ `OPPONENT_20`: /`QUEUE`/ `QUEUE_5`/ `QUEUE_10`/ `QUEUE_20`; Can't be passed together with `price` (optional)
	selfTradePreventionMode := models.NewMarginOrderSelfTradePreventionModeParameterNone // NewMarginOrderSelfTradePreventionModeParameter | `NONE`:No STP / `EXPIRE_TAKER`:expire taker order when STP triggers/ `EXPIRE_MAKER`:expire taker order when STP triggers/ `EXPIRE_BOTH`:expire both orders when STP triggers (optional)
	goodTillDate := int64(789) // int64 | order cancel time for timeInForce `GTD`, mandatory when `timeInforce` set to `GTD`; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000Mode. It must be sent in Hedge Mode. (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.NewUmOrder(context.Background()).Symbol(symbol).Side(side).Type(type_).PositionSide(positionSide).TimeInForce(timeInForce).Quantity(quantity).ReduceOnly(reduceOnly).Price(price).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).PriceMatch(priceMatch).SelfTradePreventionMode(selfTradePreventionMode).GoodTillDate(goodTillDate).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.NewUmOrder``: %v\n", err)
		return
	}

	// response from `NewUmOrder`: NewUmOrderResponse
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
 **side** | [**NewCmConditionalOrderSideParameter**](NewCmConditionalOrderSideParameter.md) |  | 
 **type_** | [**NewCmOrderTypeParameter**](NewCmOrderTypeParameter.md) | &#x60;LIMIT&#x60;, &#x60;MARKET&#x60; | 
 **positionSide** | [**NewCmConditionalOrderPositionSideParameter**](NewCmConditionalOrderPositionSideParameter.md) | Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent in Hedge Mode. | 
 **timeInForce** | [**NewCmConditionalOrderTimeInForceParameter**](NewCmConditionalOrderTimeInForceParameter.md) |  | 
 **quantity** | **float32** |  | 
 **reduceOnly** | **string** | \&quot;true\&quot; or \&quot;false\&quot;. default \&quot;false\&quot;. Cannot be sent in Hedge Mode . | 
 **price** | **float32** |  | 
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default | 
 **newOrderRespType** | [**NewCmOrderNewOrderRespTypeParameter**](NewCmOrderNewOrderRespTypeParameter.md) | \&quot;ACK\&quot;, \&quot;RESULT\&quot;, default \&quot;ACK\&quot; | 
 **priceMatch** | [**ModifyCmOrderPriceMatchParameter**](ModifyCmOrderPriceMatchParameter.md) | only avaliable for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; can be set to &#x60;OPPONENT&#x60;/ &#x60;OPPONENT_5&#x60;/ &#x60;OPPONENT_10&#x60;/ &#x60;OPPONENT_20&#x60;: /&#x60;QUEUE&#x60;/ &#x60;QUEUE_5&#x60;/ &#x60;QUEUE_10&#x60;/ &#x60;QUEUE_20&#x60;; Can&#39;t be passed together with &#x60;price&#x60; | 
 **selfTradePreventionMode** | [**NewMarginOrderSelfTradePreventionModeParameter**](NewMarginOrderSelfTradePreventionModeParameter.md) | &#x60;NONE&#x60;:No STP / &#x60;EXPIRE_TAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_MAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_BOTH&#x60;:expire both orders when STP triggers | 
 **goodTillDate** | **int64** | order cancel time for timeInForce &#x60;GTD&#x60;, mandatory when &#x60;timeInforce&#x60; set to &#x60;GTD&#x60;; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000Mode. It must be sent in Hedge Mode. | 
 **recvWindow** | **int64** |  | 

### Return type

[**NewUmOrderResponse**](NewUmOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryAllCmConditionalOrders

> QueryAllCmConditionalOrdersResponse QueryAllCmConditionalOrders(ctx).Symbol(symbol).StrategyId(strategyId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query All CM Conditional Orders(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string |  (optional)
	strategyId := int64(1) // int64 |  (optional)
	startTime := int64(1623319461670) // int64 | Timestamp in ms to get funding from INCLUSIVE. (optional)
	endTime := int64(1641782889000) // int64 | Timestamp in ms to get funding until INCLUSIVE. (optional)
	limit := int64(100) // int64 | Default 100; max 1000 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryAllCmConditionalOrders(context.Background()).Symbol(symbol).StrategyId(strategyId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryAllCmConditionalOrders``: %v\n", err)
		return
	}

	// response from `QueryAllCmConditionalOrders`: QueryAllCmConditionalOrdersResponse
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
 **strategyId** | **int64** |  | 
 **startTime** | **int64** | Timestamp in ms to get funding from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding until INCLUSIVE. | 
 **limit** | **int64** | Default 100; max 1000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryAllCmConditionalOrdersResponse**](QueryAllCmConditionalOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryAllCmOrders

> QueryAllCmOrdersResponse QueryAllCmOrders(ctx).Symbol(symbol).Pair(pair).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query All CM Orders (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	pair := "pair_example" // string |  (optional)
	orderId := int64(1) // int64 |  (optional)
	startTime := int64(1623319461670) // int64 | Timestamp in ms to get funding from INCLUSIVE. (optional)
	endTime := int64(1641782889000) // int64 | Timestamp in ms to get funding until INCLUSIVE. (optional)
	limit := int64(100) // int64 | Default 100; max 1000 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryAllCmOrders(context.Background()).Symbol(symbol).Pair(pair).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryAllCmOrders``: %v\n", err)
		return
	}

	// response from `QueryAllCmOrders`: QueryAllCmOrdersResponse
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
 **startTime** | **int64** | Timestamp in ms to get funding from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding until INCLUSIVE. | 
 **limit** | **int64** | Default 100; max 1000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryAllCmOrdersResponse**](QueryAllCmOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryAllCurrentCmOpenConditionalOrders

> QueryAllCurrentCmOpenConditionalOrdersResponse QueryAllCurrentCmOpenConditionalOrders(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query All Current CM Open Conditional Orders (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
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
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryAllCurrentCmOpenConditionalOrders(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryAllCurrentCmOpenConditionalOrders``: %v\n", err)
		return
	}

	// response from `QueryAllCurrentCmOpenConditionalOrders`: QueryAllCurrentCmOpenConditionalOrdersResponse
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

[**QueryAllCurrentCmOpenConditionalOrdersResponse**](QueryAllCurrentCmOpenConditionalOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryAllCurrentCmOpenOrders

> QueryAllCurrentCmOpenOrdersResponse QueryAllCurrentCmOpenOrders(ctx).Symbol(symbol).Pair(pair).RecvWindow(recvWindow).Execute()

Query All Current CM Open Orders(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
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
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryAllCurrentCmOpenOrders(context.Background()).Symbol(symbol).Pair(pair).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryAllCurrentCmOpenOrders``: %v\n", err)
		return
	}

	// response from `QueryAllCurrentCmOpenOrders`: QueryAllCurrentCmOpenOrdersResponse
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

[**QueryAllCurrentCmOpenOrdersResponse**](QueryAllCurrentCmOpenOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryAllCurrentUmOpenConditionalOrders

> QueryAllCurrentUmOpenConditionalOrdersResponse QueryAllCurrentUmOpenConditionalOrders(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query All Current UM Open Conditional Orders(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
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
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryAllCurrentUmOpenConditionalOrders(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryAllCurrentUmOpenConditionalOrders``: %v\n", err)
		return
	}

	// response from `QueryAllCurrentUmOpenConditionalOrders`: QueryAllCurrentUmOpenConditionalOrdersResponse
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

[**QueryAllCurrentUmOpenConditionalOrdersResponse**](QueryAllCurrentUmOpenConditionalOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryAllCurrentUmOpenOrders

> QueryAllCurrentUmOpenOrdersResponse QueryAllCurrentUmOpenOrders(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query All Current UM Open Orders(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
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
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryAllCurrentUmOpenOrders(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryAllCurrentUmOpenOrders``: %v\n", err)
		return
	}

	// response from `QueryAllCurrentUmOpenOrders`: QueryAllCurrentUmOpenOrdersResponse
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

[**QueryAllCurrentUmOpenOrdersResponse**](QueryAllCurrentUmOpenOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryAllMarginAccountOrders

> QueryAllMarginAccountOrdersResponse QueryAllMarginAccountOrders(ctx).Symbol(symbol).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query All Margin Account Orders (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	orderId := int64(1) // int64 |  (optional)
	startTime := int64(1623319461670) // int64 | Timestamp in ms to get funding from INCLUSIVE. (optional)
	endTime := int64(1641782889000) // int64 | Timestamp in ms to get funding until INCLUSIVE. (optional)
	limit := int64(100) // int64 | Default 100; max 1000 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryAllMarginAccountOrders(context.Background()).Symbol(symbol).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryAllMarginAccountOrders``: %v\n", err)
		return
	}

	// response from `QueryAllMarginAccountOrders`: QueryAllMarginAccountOrdersResponse
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
 **startTime** | **int64** | Timestamp in ms to get funding from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding until INCLUSIVE. | 
 **limit** | **int64** | Default 100; max 1000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryAllMarginAccountOrdersResponse**](QueryAllMarginAccountOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryAllUmConditionalOrders

> QueryAllUmConditionalOrdersResponse QueryAllUmConditionalOrders(ctx).Symbol(symbol).StrategyId(strategyId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query All UM Conditional Orders(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string |  (optional)
	strategyId := int64(1) // int64 |  (optional)
	startTime := int64(1623319461670) // int64 | Timestamp in ms to get funding from INCLUSIVE. (optional)
	endTime := int64(1641782889000) // int64 | Timestamp in ms to get funding until INCLUSIVE. (optional)
	limit := int64(100) // int64 | Default 100; max 1000 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryAllUmConditionalOrders(context.Background()).Symbol(symbol).StrategyId(strategyId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryAllUmConditionalOrders``: %v\n", err)
		return
	}

	// response from `QueryAllUmConditionalOrders`: QueryAllUmConditionalOrdersResponse
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
 **strategyId** | **int64** |  | 
 **startTime** | **int64** | Timestamp in ms to get funding from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding until INCLUSIVE. | 
 **limit** | **int64** | Default 100; max 1000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryAllUmConditionalOrdersResponse**](QueryAllUmConditionalOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryAllUmOrders

> QueryAllUmOrdersResponse QueryAllUmOrders(ctx).Symbol(symbol).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query All UM Orders(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	orderId := int64(1) // int64 |  (optional)
	startTime := int64(1623319461670) // int64 | Timestamp in ms to get funding from INCLUSIVE. (optional)
	endTime := int64(1641782889000) // int64 | Timestamp in ms to get funding until INCLUSIVE. (optional)
	limit := int64(100) // int64 | Default 100; max 1000 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryAllUmOrders(context.Background()).Symbol(symbol).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryAllUmOrders``: %v\n", err)
		return
	}

	// response from `QueryAllUmOrders`: QueryAllUmOrdersResponse
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
 **startTime** | **int64** | Timestamp in ms to get funding from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding until INCLUSIVE. | 
 **limit** | **int64** | Default 100; max 1000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryAllUmOrdersResponse**](QueryAllUmOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryCmConditionalOrderHistory

> QueryCmConditionalOrderHistoryResponse QueryCmConditionalOrderHistory(ctx).Symbol(symbol).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()

Query CM Conditional Order History(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	strategyId := int64(1) // int64 |  (optional)
	newClientStrategyId := "1" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryCmConditionalOrderHistory(context.Background()).Symbol(symbol).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryCmConditionalOrderHistory``: %v\n", err)
		return
	}

	// response from `QueryCmConditionalOrderHistory`: QueryCmConditionalOrderHistoryResponse
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
 **strategyId** | **int64** |  | 
 **newClientStrategyId** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryCmConditionalOrderHistoryResponse**](QueryCmConditionalOrderHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryCmModifyOrderHistory

> QueryCmModifyOrderHistoryResponse QueryCmModifyOrderHistory(ctx).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query CM Modify Order History(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	orderId := int64(1) // int64 |  (optional)
	origClientOrderId := "1" // string |  (optional)
	startTime := int64(1623319461670) // int64 | Timestamp in ms to get funding from INCLUSIVE. (optional)
	endTime := int64(1641782889000) // int64 | Timestamp in ms to get funding until INCLUSIVE. (optional)
	limit := int64(100) // int64 | Default 100; max 1000 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryCmModifyOrderHistory(context.Background()).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryCmModifyOrderHistory``: %v\n", err)
		return
	}

	// response from `QueryCmModifyOrderHistory`: QueryCmModifyOrderHistoryResponse
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
 **startTime** | **int64** | Timestamp in ms to get funding from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding until INCLUSIVE. | 
 **limit** | **int64** | Default 100; max 1000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryCmModifyOrderHistoryResponse**](QueryCmModifyOrderHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryCmOrder

> QueryCmOrderResponse QueryCmOrder(ctx).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query CM Order(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
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
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryCmOrder(context.Background()).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryCmOrder``: %v\n", err)
		return
	}

	// response from `QueryCmOrder`: QueryCmOrderResponse
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

[**QueryCmOrderResponse**](QueryCmOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryCurrentCmOpenConditionalOrder

> QueryCurrentCmOpenConditionalOrderResponse QueryCurrentCmOpenConditionalOrder(ctx).Symbol(symbol).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()

Query Current CM Open Conditional Order(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	strategyId := int64(1) // int64 |  (optional)
	newClientStrategyId := "1" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryCurrentCmOpenConditionalOrder(context.Background()).Symbol(symbol).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryCurrentCmOpenConditionalOrder``: %v\n", err)
		return
	}

	// response from `QueryCurrentCmOpenConditionalOrder`: QueryCurrentCmOpenConditionalOrderResponse
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
 **strategyId** | **int64** |  | 
 **newClientStrategyId** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryCurrentCmOpenConditionalOrderResponse**](QueryCurrentCmOpenConditionalOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryCurrentCmOpenOrder

> QueryCurrentCmOpenOrderResponse QueryCurrentCmOpenOrder(ctx).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Current CM Open Order (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
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
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryCurrentCmOpenOrder(context.Background()).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryCurrentCmOpenOrder``: %v\n", err)
		return
	}

	// response from `QueryCurrentCmOpenOrder`: QueryCurrentCmOpenOrderResponse
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

[**QueryCurrentCmOpenOrderResponse**](QueryCurrentCmOpenOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryCurrentMarginOpenOrder

> QueryCurrentMarginOpenOrderResponse QueryCurrentMarginOpenOrder(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query Current Margin Open Order (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
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
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryCurrentMarginOpenOrder(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryCurrentMarginOpenOrder``: %v\n", err)
		return
	}

	// response from `QueryCurrentMarginOpenOrder`: QueryCurrentMarginOpenOrderResponse
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

[**QueryCurrentMarginOpenOrderResponse**](QueryCurrentMarginOpenOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryCurrentUmOpenConditionalOrder

> QueryCurrentUmOpenConditionalOrderResponse QueryCurrentUmOpenConditionalOrder(ctx).Symbol(symbol).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()

Query Current UM Open Conditional Order(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	strategyId := int64(1) // int64 |  (optional)
	newClientStrategyId := "1" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryCurrentUmOpenConditionalOrder(context.Background()).Symbol(symbol).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryCurrentUmOpenConditionalOrder``: %v\n", err)
		return
	}

	// response from `QueryCurrentUmOpenConditionalOrder`: QueryCurrentUmOpenConditionalOrderResponse
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
 **strategyId** | **int64** |  | 
 **newClientStrategyId** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryCurrentUmOpenConditionalOrderResponse**](QueryCurrentUmOpenConditionalOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryCurrentUmOpenOrder

> QueryCurrentUmOpenOrderResponse QueryCurrentUmOpenOrder(ctx).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Current UM Open Order(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
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
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryCurrentUmOpenOrder(context.Background()).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryCurrentUmOpenOrder``: %v\n", err)
		return
	}

	// response from `QueryCurrentUmOpenOrder`: QueryCurrentUmOpenOrderResponse
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

[**QueryCurrentUmOpenOrderResponse**](QueryCurrentUmOpenOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryMarginAccountOrder

> QueryMarginAccountOrderResponse QueryMarginAccountOrder(ctx).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Margin Account Order (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
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
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountOrder(context.Background()).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryMarginAccountOrder``: %v\n", err)
		return
	}

	// response from `QueryMarginAccountOrder`: QueryMarginAccountOrderResponse
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

[**QueryMarginAccountOrderResponse**](QueryMarginAccountOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryMarginAccountsAllOco

> QueryMarginAccountsAllOcoResponse QueryMarginAccountsAllOco(ctx).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query Margin Account's all OCO (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	fromId := int64(1) // int64 | Trade id to fetch from. Default gets most recent trades. (optional)
	startTime := int64(1623319461670) // int64 | Timestamp in ms to get funding from INCLUSIVE. (optional)
	endTime := int64(1641782889000) // int64 | Timestamp in ms to get funding until INCLUSIVE. (optional)
	limit := int64(100) // int64 | Default 100; max 1000 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsAllOco(context.Background()).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryMarginAccountsAllOco``: %v\n", err)
		return
	}

	// response from `QueryMarginAccountsAllOco`: QueryMarginAccountsAllOcoResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **fromId** | **int64** | Trade id to fetch from. Default gets most recent trades. | 
 **startTime** | **int64** | Timestamp in ms to get funding from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding until INCLUSIVE. | 
 **limit** | **int64** | Default 100; max 1000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryMarginAccountsAllOcoResponse**](QueryMarginAccountsAllOcoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryMarginAccountsOco

> QueryMarginAccountsOcoResponse QueryMarginAccountsOco(ctx).OrderListId(orderListId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Margin Account's OCO (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	orderListId := int64(1) // int64 | Either `orderListId` or `listClientOrderId` must be provided (optional)
	origClientOrderId := "1" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsOco(context.Background()).OrderListId(orderListId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryMarginAccountsOco``: %v\n", err)
		return
	}

	// response from `QueryMarginAccountsOco`: QueryMarginAccountsOcoResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **orderListId** | **int64** | Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided | 
 **origClientOrderId** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryMarginAccountsOcoResponse**](QueryMarginAccountsOcoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryMarginAccountsOpenOco

> QueryMarginAccountsOpenOcoResponse QueryMarginAccountsOpenOco(ctx).RecvWindow(recvWindow).Execute()

Query Margin Account's Open OCO (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsOpenOco(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryMarginAccountsOpenOco``: %v\n", err)
		return
	}

	// response from `QueryMarginAccountsOpenOco`: QueryMarginAccountsOpenOcoResponse
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

[**QueryMarginAccountsOpenOcoResponse**](QueryMarginAccountsOpenOcoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryUmConditionalOrderHistory

> QueryUmConditionalOrderHistoryResponse QueryUmConditionalOrderHistory(ctx).Symbol(symbol).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()

Query UM Conditional Order History(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	strategyId := int64(1) // int64 |  (optional)
	newClientStrategyId := "1" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryUmConditionalOrderHistory(context.Background()).Symbol(symbol).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryUmConditionalOrderHistory``: %v\n", err)
		return
	}

	// response from `QueryUmConditionalOrderHistory`: QueryUmConditionalOrderHistoryResponse
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
 **strategyId** | **int64** |  | 
 **newClientStrategyId** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryUmConditionalOrderHistoryResponse**](QueryUmConditionalOrderHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryUmModifyOrderHistory

> QueryUmModifyOrderHistoryResponse QueryUmModifyOrderHistory(ctx).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query UM Modify Order History(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	orderId := int64(1) // int64 |  (optional)
	origClientOrderId := "1" // string |  (optional)
	startTime := int64(1623319461670) // int64 | Timestamp in ms to get funding from INCLUSIVE. (optional)
	endTime := int64(1641782889000) // int64 | Timestamp in ms to get funding until INCLUSIVE. (optional)
	limit := int64(100) // int64 | Default 100; max 1000 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryUmModifyOrderHistory(context.Background()).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryUmModifyOrderHistory``: %v\n", err)
		return
	}

	// response from `QueryUmModifyOrderHistory`: QueryUmModifyOrderHistoryResponse
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
 **startTime** | **int64** | Timestamp in ms to get funding from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding until INCLUSIVE. | 
 **limit** | **int64** | Default 100; max 1000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryUmModifyOrderHistoryResponse**](QueryUmModifyOrderHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryUmOrder

> QueryUmOrderResponse QueryUmOrder(ctx).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query UM Order (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
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
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryUmOrder(context.Background()).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryUmOrder``: %v\n", err)
		return
	}

	// response from `QueryUmOrder`: QueryUmOrderResponse
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

[**QueryUmOrderResponse**](QueryUmOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryUsersCmForceOrders

> QueryUsersCmForceOrdersResponse QueryUsersCmForceOrders(ctx).Symbol(symbol).AutoCloseType(autoCloseType).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query User's CM Force Orders(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string |  (optional)
	autoCloseType := models.QueryUsersCmForceOrdersAutoCloseTypeParameterLiquidation // QueryUsersCmForceOrdersAutoCloseTypeParameter | `LIQUIDATION` for liquidation orders, `ADL` for ADL orders. (optional)
	startTime := int64(1623319461670) // int64 | Timestamp in ms to get funding from INCLUSIVE. (optional)
	endTime := int64(1641782889000) // int64 | Timestamp in ms to get funding until INCLUSIVE. (optional)
	limit := int64(100) // int64 | Default 100; max 1000 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryUsersCmForceOrders(context.Background()).Symbol(symbol).AutoCloseType(autoCloseType).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryUsersCmForceOrders``: %v\n", err)
		return
	}

	// response from `QueryUsersCmForceOrders`: QueryUsersCmForceOrdersResponse
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
 **autoCloseType** | [**QueryUsersCmForceOrdersAutoCloseTypeParameter**](QueryUsersCmForceOrdersAutoCloseTypeParameter.md) | &#x60;LIQUIDATION&#x60; for liquidation orders, &#x60;ADL&#x60; for ADL orders. | 
 **startTime** | **int64** | Timestamp in ms to get funding from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding until INCLUSIVE. | 
 **limit** | **int64** | Default 100; max 1000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryUsersCmForceOrdersResponse**](QueryUsersCmForceOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryUsersMarginForceOrders

> QueryUsersMarginForceOrdersResponse QueryUsersMarginForceOrders(ctx).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Query User's Margin Force Orders(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	startTime := int64(1623319461670) // int64 | Timestamp in ms to get funding from INCLUSIVE. (optional)
	endTime := int64(1641782889000) // int64 | Timestamp in ms to get funding until INCLUSIVE. (optional)
	current := int64(1) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(10) // int64 | Default:10 Max:100 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryUsersMarginForceOrders(context.Background()).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryUsersMarginForceOrders``: %v\n", err)
		return
	}

	// response from `QueryUsersMarginForceOrders`: QueryUsersMarginForceOrdersResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** | Timestamp in ms to get funding from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding until INCLUSIVE. | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10 Max:100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryUsersMarginForceOrdersResponse**](QueryUsersMarginForceOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryUsersUmForceOrders

> QueryUsersUmForceOrdersResponse QueryUsersUmForceOrders(ctx).Symbol(symbol).AutoCloseType(autoCloseType).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query User's UM Force Orders (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string |  (optional)
	autoCloseType := models.QueryUsersCmForceOrdersAutoCloseTypeParameterLiquidation // QueryUsersCmForceOrdersAutoCloseTypeParameter | `LIQUIDATION` for liquidation orders, `ADL` for ADL orders. (optional)
	startTime := int64(1623319461670) // int64 | Timestamp in ms to get funding from INCLUSIVE. (optional)
	endTime := int64(1641782889000) // int64 | Timestamp in ms to get funding until INCLUSIVE. (optional)
	limit := int64(100) // int64 | Default 100; max 1000 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryUsersUmForceOrders(context.Background()).Symbol(symbol).AutoCloseType(autoCloseType).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryUsersUmForceOrders``: %v\n", err)
		return
	}

	// response from `QueryUsersUmForceOrders`: QueryUsersUmForceOrdersResponse
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
 **autoCloseType** | [**QueryUsersCmForceOrdersAutoCloseTypeParameter**](QueryUsersCmForceOrdersAutoCloseTypeParameter.md) | &#x60;LIQUIDATION&#x60; for liquidation orders, &#x60;ADL&#x60; for ADL orders. | 
 **startTime** | **int64** | Timestamp in ms to get funding from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding until INCLUSIVE. | 
 **limit** | **int64** | Default 100; max 1000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryUsersUmForceOrdersResponse**](QueryUsersUmForceOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## ToggleBnbBurnOnUmFuturesTrade

> ToggleBnbBurnOnUmFuturesTradeResponse ToggleBnbBurnOnUmFuturesTrade(ctx).FeeBurn(feeBurn).RecvWindow(recvWindow).Execute()

Toggle BNB Burn On UM Futures Trade (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	feeBurn := "feeBurn_example" // string | \"true\": Fee Discount On; \"false\": Fee Discount Off
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.ToggleBnbBurnOnUmFuturesTrade(context.Background()).FeeBurn(feeBurn).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.ToggleBnbBurnOnUmFuturesTrade``: %v\n", err)
		return
	}

	// response from `ToggleBnbBurnOnUmFuturesTrade`: ToggleBnbBurnOnUmFuturesTradeResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **feeBurn** | **string** | \&quot;true\&quot;: Fee Discount On; \&quot;false\&quot;: Fee Discount Off | 
 **recvWindow** | **int64** |  | 

### Return type

[**ToggleBnbBurnOnUmFuturesTradeResponse**](ToggleBnbBurnOnUmFuturesTradeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## UmAccountTradeList

> UmAccountTradeListResponse UmAccountTradeList(ctx).Symbol(symbol).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()

UM Account Trade List(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	startTime := int64(1623319461670) // int64 | Timestamp in ms to get funding from INCLUSIVE. (optional)
	endTime := int64(1641782889000) // int64 | Timestamp in ms to get funding until INCLUSIVE. (optional)
	fromId := int64(1) // int64 | Trade id to fetch from. Default gets most recent trades. (optional)
	limit := int64(100) // int64 | Default 100; max 1000 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.UmAccountTradeList(context.Background()).Symbol(symbol).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.UmAccountTradeList``: %v\n", err)
		return
	}

	// response from `UmAccountTradeList`: UmAccountTradeListResponse
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
 **startTime** | **int64** | Timestamp in ms to get funding from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding until INCLUSIVE. | 
 **fromId** | **int64** | Trade id to fetch from. Default gets most recent trades. | 
 **limit** | **int64** | Default 100; max 1000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**UmAccountTradeListResponse**](UmAccountTradeListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## UmPositionAdlQuantileEstimation

> UmPositionAdlQuantileEstimationResponse UmPositionAdlQuantileEstimation(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

UM Position ADL Quantile Estimation(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
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
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.UmPositionAdlQuantileEstimation(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.UmPositionAdlQuantileEstimation``: %v\n", err)
		return
	}

	// response from `UmPositionAdlQuantileEstimation`: UmPositionAdlQuantileEstimationResponse
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

[**UmPositionAdlQuantileEstimationResponse**](UmPositionAdlQuantileEstimationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

