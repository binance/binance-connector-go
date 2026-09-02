# \TradeAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**CreateSpecialKey**](TradeAPI.md#CreateSpecialKey) | **Post** /sapi/v1/margin/apiKey | Create Special Key(Low-Latency Trading) (TRADE)
[**DeleteSpecialKey**](TradeAPI.md#DeleteSpecialKey) | **Delete** /sapi/v1/margin/apiKey | Delete Special Key(Low-Latency Trading) (TRADE)
[**EditIpForSpecialKey**](TradeAPI.md#EditIpForSpecialKey) | **Put** /sapi/v1/margin/apiKey/ip | Edit ip for Special Key(Low-Latency Trading) (TRADE)
[**ExitSpecialKeyMode**](TradeAPI.md#ExitSpecialKeyMode) | **Post** /sapi/v1/margin/exit-special-key-mode | Exit Special Key Mode (TRADE)
[**GetForceLiquidationRecord**](TradeAPI.md#GetForceLiquidationRecord) | **Get** /sapi/v1/margin/forceLiquidationRec | Get Force Liquidation Record (USER_DATA)
[**GetSmallLiabilityExchangeCoinList**](TradeAPI.md#GetSmallLiabilityExchangeCoinList) | **Get** /sapi/v1/margin/exchange-small-liability | Get Small Liability Exchange Coin List (USER_DATA)
[**GetSmallLiabilityExchangeHistory**](TradeAPI.md#GetSmallLiabilityExchangeHistory) | **Get** /sapi/v1/margin/exchange-small-liability-history | Get Small Liability Exchange History (USER_DATA)
[**LiquidationLoanRepay**](TradeAPI.md#LiquidationLoanRepay) | **Post** /sapi/v1/margin/liquidation-loan/repay | Liquidation Loan Repay (MARGIN)
[**MarginAccountCancelAllOpenOrdersOnASymbol**](TradeAPI.md#MarginAccountCancelAllOpenOrdersOnASymbol) | **Delete** /sapi/v1/margin/openOrders | Margin Account Cancel all Open Orders on a Symbol (TRADE)
[**MarginAccountCancelOco**](TradeAPI.md#MarginAccountCancelOco) | **Delete** /sapi/v1/margin/orderList | Margin Account Cancel OCO (TRADE)
[**MarginAccountCancelOrder**](TradeAPI.md#MarginAccountCancelOrder) | **Delete** /sapi/v1/margin/order | Margin Account Cancel Order (TRADE)
[**MarginAccountNewOco**](TradeAPI.md#MarginAccountNewOco) | **Post** /sapi/v1/margin/order/oco | Margin Account New OCO (TRADE)
[**MarginAccountNewOrder**](TradeAPI.md#MarginAccountNewOrder) | **Post** /sapi/v1/margin/order | Margin Account New Order (TRADE)
[**MarginAccountNewOto**](TradeAPI.md#MarginAccountNewOto) | **Post** /sapi/v1/margin/order/oto | Margin Account New OTO (TRADE)
[**MarginAccountNewOtoco**](TradeAPI.md#MarginAccountNewOtoco) | **Post** /sapi/v1/margin/order/otoco | Margin Account New OTOCO (TRADE)
[**MarginManualLiquidation**](TradeAPI.md#MarginManualLiquidation) | **Post** /sapi/v1/margin/manual-liquidation | Margin Manual Liquidation (TRADE)
[**QueryCurrentMarginOrderCountUsage**](TradeAPI.md#QueryCurrentMarginOrderCountUsage) | **Get** /sapi/v1/margin/rateLimit/order | Query Current Margin Order Count Usage (TRADE)
[**QueryLiquidationLoan**](TradeAPI.md#QueryLiquidationLoan) | **Get** /sapi/v1/margin/liquidation-loan | Query Liquidation Loan (USER_DATA)
[**QueryLiquidationLoanRepayHistory**](TradeAPI.md#QueryLiquidationLoanRepayHistory) | **Get** /sapi/v1/margin/liquidation-loan/repay-history | Query Liquidation Loan Repay History (USER_DATA)
[**QueryMarginAccountsAllOco**](TradeAPI.md#QueryMarginAccountsAllOco) | **Get** /sapi/v1/margin/allOrderList | Query Margin Account&#39;s all OCO (USER_DATA)
[**QueryMarginAccountsAllOrders**](TradeAPI.md#QueryMarginAccountsAllOrders) | **Get** /sapi/v1/margin/allOrders | Query Margin Account&#39;s All Orders (USER_DATA)
[**QueryMarginAccountsOco**](TradeAPI.md#QueryMarginAccountsOco) | **Get** /sapi/v1/margin/orderList | Query Margin Account&#39;s OCO (USER_DATA)
[**QueryMarginAccountsOpenOco**](TradeAPI.md#QueryMarginAccountsOpenOco) | **Get** /sapi/v1/margin/openOrderList | Query Margin Account&#39;s Open OCO (USER_DATA)
[**QueryMarginAccountsOpenOrders**](TradeAPI.md#QueryMarginAccountsOpenOrders) | **Get** /sapi/v1/margin/openOrders | Query Margin Account&#39;s Open Orders (USER_DATA)
[**QueryMarginAccountsOrder**](TradeAPI.md#QueryMarginAccountsOrder) | **Get** /sapi/v1/margin/order | Query Margin Account&#39;s Order (USER_DATA)
[**QueryMarginAccountsTradeList**](TradeAPI.md#QueryMarginAccountsTradeList) | **Get** /sapi/v1/margin/myTrades | Query Margin Account&#39;s Trade List (USER_DATA)
[**QueryPreventedMatches**](TradeAPI.md#QueryPreventedMatches) | **Get** /sapi/v1/margin/myPreventedMatches | Query Prevented Matches (USER_DATA)
[**QuerySpecialKey**](TradeAPI.md#QuerySpecialKey) | **Get** /sapi/v1/margin/apiKey | Query Special key(Low Latency Trading) (TRADE)
[**QuerySpecialKeyList**](TradeAPI.md#QuerySpecialKeyList) | **Get** /sapi/v1/margin/api-key-list | Query Special key List(Low Latency Trading) (TRADE)
[**SmallLiabilityExchange**](TradeAPI.md#SmallLiabilityExchange) | **Post** /sapi/v1/margin/exchange-small-liability | Small Liability Exchange (MARGIN)


## CreateSpecialKey

> CreateSpecialKeyResponse CreateSpecialKey(ctx).ApiName(apiName).Symbol(symbol).Ip(ip).PublicKey(publicKey).PermissionMode(permissionMode).RecvWindow(recvWindow).Execute()

Create Special Key(Low-Latency Trading) (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	apiName := "apiName" // string | 
	symbol := "BTCUSDT" // string |  (optional)
	ip := "69.210.67.14,69.210.67.15" // string | Can be added in batches, separated by commas. Max 30 for an API key (optional)
	publicKey := "publicKey" // string | 1. If publicKey is inputted it will create an RSA or Ed25519 key.  2. Need to be encoded to URL-encoded format (optional)
	permissionMode := models.CreateSpecialKeyPermissionModeParameterTrade // CreateSpecialKeyPermissionModeParameter | This parameter is only for the Ed25519 API key, and does not effact for other encryption methods. The value can be TRADE (TRADE for all permissions) or READ (READ for USER_DATA, FIX_API_READ_ONLY). The default value is TRADE. (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CreateSpecialKey(context.Background()).ApiName(apiName).Symbol(symbol).Ip(ip).PublicKey(publicKey).PermissionMode(permissionMode).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CreateSpecialKey``: %v\n", err)
		return
	}

	// response from `CreateSpecialKey`: CreateSpecialKeyResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **apiName** | **string** |  | 
 **symbol** | **string** |  | 
 **ip** | **string** | Can be added in batches, separated by commas. Max 30 for an API key | 
 **publicKey** | **string** | 1. If publicKey is inputted it will create an RSA or Ed25519 key.  2. Need to be encoded to URL-encoded format | 
 **permissionMode** | [**CreateSpecialKeyPermissionModeParameter**](CreateSpecialKeyPermissionModeParameter.md) | This parameter is only for the Ed25519 API key, and does not effact for other encryption methods. The value can be TRADE (TRADE for all permissions) or READ (READ for USER_DATA, FIX_API_READ_ONLY). The default value is TRADE. | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateSpecialKeyResponse**](CreateSpecialKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## DeleteSpecialKey

> DeleteSpecialKey(ctx).ApiName(apiName).Symbol(symbol).RecvWindow(recvWindow).Execute()

Delete Special Key(Low-Latency Trading) (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	apiName := "apiName" // string |  (optional)
	symbol := "BTCUSDT" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	 err := apiClient.RestApi.TradeAPI.DeleteSpecialKey(context.Background()).ApiName(apiName).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.DeleteSpecialKey``: %v\n", err)
		return
	}

}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **apiName** | **string** |  | 
 **symbol** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: Not defined

[[Back to README]](../../../README.md)


## EditIpForSpecialKey

> EditIpForSpecialKey(ctx).Ip(ip).Symbol(symbol).RecvWindow(recvWindow).Execute()

Edit ip for Special Key(Low-Latency Trading) (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	ip := "24.156.99.202" // string | Can be added in batches, separated by commas. Max 30 for an API key
	symbol := "BTCUSDT" // string | isolated margin pair (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	 err := apiClient.RestApi.TradeAPI.EditIpForSpecialKey(context.Background()).Ip(ip).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.EditIpForSpecialKey``: %v\n", err)
		return
	}

}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **ip** | **string** | Can be added in batches, separated by commas. Max 30 for an API key | 
 **symbol** | **string** | isolated margin pair | 
 **recvWindow** | **int64** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: Not defined

[[Back to README]](../../../README.md)


## ExitSpecialKeyMode

> map[string]interface{} ExitSpecialKeyMode(ctx).RecvWindow(recvWindow).Execute()

Exit Special Key Mode (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 | The value cannot be greater than `60000` (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.ExitSpecialKeyMode(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.ExitSpecialKeyMode``: %v\n", err)
		return
	}

	// response from `ExitSpecialKeyMode`: map[string]interface{}
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetForceLiquidationRecord

> GetForceLiquidationRecordResponse GetForceLiquidationRecord(ctx).StartTime(startTime).EndTime(endTime).IsolatedSymbol(isolatedSymbol).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Force Liquidation Record (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	isolatedSymbol := "BTCUSDT" // string |  (optional)
	current := int64(1) // int64 |  (optional)
	size := int64(10) // int64 |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.GetForceLiquidationRecord(context.Background()).StartTime(startTime).EndTime(endTime).IsolatedSymbol(isolatedSymbol).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.GetForceLiquidationRecord``: %v\n", err)
		return
	}

	// response from `GetForceLiquidationRecord`: GetForceLiquidationRecordResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **isolatedSymbol** | **string** |  | 
 **current** | **int64** |  | 
 **size** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetForceLiquidationRecordResponse**](GetForceLiquidationRecordResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetSmallLiabilityExchangeCoinList

> GetSmallLiabilityExchangeCoinListResponse GetSmallLiabilityExchangeCoinList(ctx).RecvWindow(recvWindow).Execute()

Get Small Liability Exchange Coin List (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.GetSmallLiabilityExchangeCoinList(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.GetSmallLiabilityExchangeCoinList``: %v\n", err)
		return
	}

	// response from `GetSmallLiabilityExchangeCoinList`: GetSmallLiabilityExchangeCoinListResponse
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

[**GetSmallLiabilityExchangeCoinListResponse**](GetSmallLiabilityExchangeCoinListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetSmallLiabilityExchangeHistory

> GetSmallLiabilityExchangeHistoryResponse GetSmallLiabilityExchangeHistory(ctx).Current(current).Size(size).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

Get Small Liability Exchange History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	current := int64(1) // int64 | 
	size := int64(10) // int64 | 
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.GetSmallLiabilityExchangeHistory(context.Background()).Current(current).Size(size).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.GetSmallLiabilityExchangeHistory``: %v\n", err)
		return
	}

	// response from `GetSmallLiabilityExchangeHistory`: GetSmallLiabilityExchangeHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **current** | **int64** |  | 
 **size** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSmallLiabilityExchangeHistoryResponse**](GetSmallLiabilityExchangeHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## LiquidationLoanRepay

> LiquidationLoanRepayResponse LiquidationLoanRepay(ctx).Asset(asset).Amount(amount).RecvWindow(recvWindow).Execute()

Liquidation Loan Repay (MARGIN)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	asset := "USDT" // string | The asset to repay (e.g. USDT, USDC)
	amount := float64(300.00) // float64 | Repayment amount, must be greater than 0
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.LiquidationLoanRepay(context.Background()).Asset(asset).Amount(amount).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.LiquidationLoanRepay``: %v\n", err)
		return
	}

	// response from `LiquidationLoanRepay`: LiquidationLoanRepayResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** | The asset to repay (e.g. USDT, USDC) | 
 **amount** | **float64** | Repayment amount, must be greater than 0 | 
 **recvWindow** | **int64** |  | 

### Return type

[**LiquidationLoanRepayResponse**](LiquidationLoanRepayResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MarginAccountCancelAllOpenOrdersOnASymbol

> MarginAccountCancelAllOpenOrdersOnASymbolResponse MarginAccountCancelAllOpenOrdersOnASymbol(ctx).Symbol(symbol).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()

Margin Account Cancel all Open Orders on a Symbol (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSDT" // string | 
	isIsolated := models.QueryMarginAccountsOpenOrdersIsIsolatedParameterTrue // QueryMarginAccountsOpenOrdersIsIsolatedParameter |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.MarginAccountCancelAllOpenOrdersOnASymbol(context.Background()).Symbol(symbol).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.MarginAccountCancelAllOpenOrdersOnASymbol``: %v\n", err)
		return
	}

	// response from `MarginAccountCancelAllOpenOrdersOnASymbol`: MarginAccountCancelAllOpenOrdersOnASymbolResponse
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
 **isIsolated** | [**QueryMarginAccountsOpenOrdersIsIsolatedParameter**](QueryMarginAccountsOpenOrdersIsIsolatedParameter.md) |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**MarginAccountCancelAllOpenOrdersOnASymbolResponse**](MarginAccountCancelAllOpenOrdersOnASymbolResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MarginAccountCancelOco

> MarginAccountCancelOcoResponse MarginAccountCancelOco(ctx).Symbol(symbol).IsIsolated(isIsolated).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()

Margin Account Cancel OCO (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSDT" // string | 
	isIsolated := models.QueryMarginAccountsOpenOrdersIsIsolatedParameterTrue // QueryMarginAccountsOpenOrdersIsIsolatedParameter |  (optional)
	orderListId := int64(1) // int64 |  (optional)
	listClientOrderId := "1" // string |  (optional)
	newClientOrderId := "1" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.MarginAccountCancelOco(context.Background()).Symbol(symbol).IsIsolated(isIsolated).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.MarginAccountCancelOco``: %v\n", err)
		return
	}

	// response from `MarginAccountCancelOco`: MarginAccountCancelOcoResponse
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
 **isIsolated** | [**QueryMarginAccountsOpenOrdersIsIsolatedParameter**](QueryMarginAccountsOpenOrdersIsIsolatedParameter.md) |  | 
 **orderListId** | **int64** |  | 
 **listClientOrderId** | **string** |  | 
 **newClientOrderId** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**MarginAccountCancelOcoResponse**](MarginAccountCancelOcoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MarginAccountCancelOrder

> MarginAccountCancelOrderResponse MarginAccountCancelOrder(ctx).Symbol(symbol).IsIsolated(isIsolated).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()

Margin Account Cancel Order (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "LTCBTC" // string | 
	isIsolated := models.QueryMarginAccountsOpenOrdersIsIsolatedParameterTrue // QueryMarginAccountsOpenOrdersIsIsolatedParameter |  (optional)
	orderId := int64(1) // int64 |  (optional)
	origClientOrderId := "1" // string |  (optional)
	newClientOrderId := "1" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.MarginAccountCancelOrder(context.Background()).Symbol(symbol).IsIsolated(isIsolated).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.MarginAccountCancelOrder``: %v\n", err)
		return
	}

	// response from `MarginAccountCancelOrder`: MarginAccountCancelOrderResponse
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
 **isIsolated** | [**QueryMarginAccountsOpenOrdersIsIsolatedParameter**](QueryMarginAccountsOpenOrdersIsIsolatedParameter.md) |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | 
 **newClientOrderId** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**MarginAccountCancelOrderResponse**](MarginAccountCancelOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MarginAccountNewOco

> MarginAccountNewOcoResponse MarginAccountNewOco(ctx).Symbol(symbol).Side(side).Quantity(quantity).Price(price).StopPrice(stopPrice).IsIsolated(isIsolated).ListClientOrderId(listClientOrderId).LimitClientOrderId(limitClientOrderId).LimitIcebergQty(limitIcebergQty).StopClientOrderId(stopClientOrderId).StopLimitPrice(stopLimitPrice).StopIcebergQty(stopIcebergQty).StopLimitTimeInForce(stopLimitTimeInForce).NewOrderRespType(newOrderRespType).SideEffectType(sideEffectType).SelfTradePreventionMode(selfTradePreventionMode).AutoRepayAtCancel(autoRepayAtCancel).RecvWindow(recvWindow).Execute()

Margin Account New OCO (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "LTCBTC" // string | 
	side := models.MarginAccountNewOrderSideParameterBuy // MarginAccountNewOrderSideParameter | 
	quantity := float64(1.0) // float64 | 
	price := float64(1.0) // float64 | 
	stopPrice := float64(1.0) // float64 | 
	isIsolated := models.QueryMarginAccountsOpenOrdersIsIsolatedParameterTrue // QueryMarginAccountsOpenOrdersIsIsolatedParameter |  (optional)
	listClientOrderId := "1" // string | A unique Id for the entire orderList (optional)
	limitClientOrderId := "1" // string | A unique Id for the limit order (optional)
	limitIcebergQty := float64(1.0) // float64 |  (optional)
	stopClientOrderId := "1" // string | A unique Id for the stop loss/stop loss limit leg (optional)
	stopLimitPrice := float64(1.0) // float64 | If provided, `stopLimitTimeInForce` is required. (optional)
	stopIcebergQty := float64(1.0) // float64 |  (optional)
	stopLimitTimeInForce := models.MarginAccountNewOcoStopLimitTimeInForceParameterGtc // MarginAccountNewOcoStopLimitTimeInForceParameter |  (optional)
	newOrderRespType := models.MarginAccountNewOrderNewOrderRespTypeParameterAck // MarginAccountNewOrderNewOrderRespTypeParameter |  (optional)
	sideEffectType := models.MarginAccountNewOrderSideEffectTypeParameterNoSideEffect // MarginAccountNewOrderSideEffectTypeParameter |  (optional)
	selfTradePreventionMode := models.MarginAccountNewOrderSelfTradePreventionModeParameterExpireTaker // MarginAccountNewOrderSelfTradePreventionModeParameter |  (optional)
	autoRepayAtCancel := false // bool | Only when MARGIN_BUY or AUTO_BORROW_REPAY order takes effect, true means that the debt generated by the order needs to be repay after the order is cancelled. (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.MarginAccountNewOco(context.Background()).Symbol(symbol).Side(side).Quantity(quantity).Price(price).StopPrice(stopPrice).IsIsolated(isIsolated).ListClientOrderId(listClientOrderId).LimitClientOrderId(limitClientOrderId).LimitIcebergQty(limitIcebergQty).StopClientOrderId(stopClientOrderId).StopLimitPrice(stopLimitPrice).StopIcebergQty(stopIcebergQty).StopLimitTimeInForce(stopLimitTimeInForce).NewOrderRespType(newOrderRespType).SideEffectType(sideEffectType).SelfTradePreventionMode(selfTradePreventionMode).AutoRepayAtCancel(autoRepayAtCancel).RecvWindow(recvWindow).Execute()
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
 **side** | [**MarginAccountNewOrderSideParameter**](MarginAccountNewOrderSideParameter.md) |  | 
 **quantity** | **float64** |  | 
 **price** | **float64** |  | 
 **stopPrice** | **float64** |  | 
 **isIsolated** | [**QueryMarginAccountsOpenOrdersIsIsolatedParameter**](QueryMarginAccountsOpenOrdersIsIsolatedParameter.md) |  | 
 **listClientOrderId** | **string** | A unique Id for the entire orderList | 
 **limitClientOrderId** | **string** | A unique Id for the limit order | 
 **limitIcebergQty** | **float64** |  | 
 **stopClientOrderId** | **string** | A unique Id for the stop loss/stop loss limit leg | 
 **stopLimitPrice** | **float64** | If provided, &#x60;stopLimitTimeInForce&#x60; is required. | 
 **stopIcebergQty** | **float64** |  | 
 **stopLimitTimeInForce** | [**MarginAccountNewOcoStopLimitTimeInForceParameter**](MarginAccountNewOcoStopLimitTimeInForceParameter.md) |  | 
 **newOrderRespType** | [**MarginAccountNewOrderNewOrderRespTypeParameter**](MarginAccountNewOrderNewOrderRespTypeParameter.md) |  | 
 **sideEffectType** | [**MarginAccountNewOrderSideEffectTypeParameter**](MarginAccountNewOrderSideEffectTypeParameter.md) |  | 
 **selfTradePreventionMode** | [**MarginAccountNewOrderSelfTradePreventionModeParameter**](MarginAccountNewOrderSelfTradePreventionModeParameter.md) |  | 
 **autoRepayAtCancel** | **bool** | Only when MARGIN_BUY or AUTO_BORROW_REPAY order takes effect, true means that the debt generated by the order needs to be repay after the order is cancelled. | 
 **recvWindow** | **int64** |  | 

### Return type

[**MarginAccountNewOcoResponse**](MarginAccountNewOcoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MarginAccountNewOrder

> MarginAccountNewOrderResponse MarginAccountNewOrder(ctx).Symbol(symbol).Side(side).Type(type_).IsIsolated(isIsolated).Quantity(quantity).QuoteOrderQty(quoteOrderQty).Price(price).StopPrice(stopPrice).NewClientOrderId(newClientOrderId).IcebergQty(icebergQty).NewOrderRespType(newOrderRespType).SideEffectType(sideEffectType).TimeInForce(timeInForce).SelfTradePreventionMode(selfTradePreventionMode).TrailingDelta(trailingDelta).AutoRepayAtCancel(autoRepayAtCancel).RecvWindow(recvWindow).Execute()

Margin Account New Order (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSDT" // string | 
	side := models.MarginAccountNewOrderSideParameterBuy // MarginAccountNewOrderSideParameter | 
	type_ := models.MarginAccountNewOrderTypeParameterLimit // MarginAccountNewOrderTypeParameter | 
	isIsolated := models.QueryMarginAccountsOpenOrdersIsIsolatedParameterTrue // QueryMarginAccountsOpenOrdersIsIsolatedParameter |  (optional)
	quantity := float64(1.0) // float64 |  (optional)
	quoteOrderQty := float64(1.0) // float64 |  (optional)
	price := float64(1.0) // float64 |  (optional)
	stopPrice := float64(1.0) // float64 | Used with `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, and `TAKE_PROFIT_LIMIT` orders. (optional)
	newClientOrderId := "1" // string | A unique id among open orders. Automatically generated if not sent. (optional)
	icebergQty := float64(1.0) // float64 | Used with `LIMIT`, `STOP_LOSS_LIMIT`, and `TAKE_PROFIT_LIMIT` to create an iceberg order. (optional)
	newOrderRespType := models.MarginAccountNewOrderNewOrderRespTypeParameterAck // MarginAccountNewOrderNewOrderRespTypeParameter | MARKET and LIMIT order types default to FULL, all other orders default to ACK. (optional)
	sideEffectType := models.MarginAccountNewOrderSideEffectTypeParameterNoSideEffect // MarginAccountNewOrderSideEffectTypeParameter |  (optional)
	timeInForce := models.MarginAccountNewOrderTimeInForceParameterGtc // MarginAccountNewOrderTimeInForceParameter |  (optional)
	selfTradePreventionMode := models.MarginAccountNewOrderSelfTradePreventionModeParameterExpireTaker // MarginAccountNewOrderSelfTradePreventionModeParameter |  (optional)
	trailingDelta := int64(100) // int64 | Used with `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, and `TAKE_PROFIT_LIMIT` orders. (optional)
	autoRepayAtCancel := true // bool | Only when MARGIN_BUY or AUTO_BORROW_REPAY order takes effect, true means that the debt generated by the order needs to be repaid after the order is cancelled. (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.MarginAccountNewOrder(context.Background()).Symbol(symbol).Side(side).Type(type_).IsIsolated(isIsolated).Quantity(quantity).QuoteOrderQty(quoteOrderQty).Price(price).StopPrice(stopPrice).NewClientOrderId(newClientOrderId).IcebergQty(icebergQty).NewOrderRespType(newOrderRespType).SideEffectType(sideEffectType).TimeInForce(timeInForce).SelfTradePreventionMode(selfTradePreventionMode).TrailingDelta(trailingDelta).AutoRepayAtCancel(autoRepayAtCancel).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.MarginAccountNewOrder``: %v\n", err)
		return
	}

	// response from `MarginAccountNewOrder`: MarginAccountNewOrderResponse
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
 **side** | [**MarginAccountNewOrderSideParameter**](MarginAccountNewOrderSideParameter.md) |  | 
 **type_** | [**MarginAccountNewOrderTypeParameter**](MarginAccountNewOrderTypeParameter.md) |  | 
 **isIsolated** | [**QueryMarginAccountsOpenOrdersIsIsolatedParameter**](QueryMarginAccountsOpenOrdersIsIsolatedParameter.md) |  | 
 **quantity** | **float64** |  | 
 **quoteOrderQty** | **float64** |  | 
 **price** | **float64** |  | 
 **stopPrice** | **float64** | Used with &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; orders. | 
 **newClientOrderId** | **string** | A unique id among open orders. Automatically generated if not sent. | 
 **icebergQty** | **float64** | Used with &#x60;LIMIT&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; to create an iceberg order. | 
 **newOrderRespType** | [**MarginAccountNewOrderNewOrderRespTypeParameter**](MarginAccountNewOrderNewOrderRespTypeParameter.md) | MARKET and LIMIT order types default to FULL, all other orders default to ACK. | 
 **sideEffectType** | [**MarginAccountNewOrderSideEffectTypeParameter**](MarginAccountNewOrderSideEffectTypeParameter.md) |  | 
 **timeInForce** | [**MarginAccountNewOrderTimeInForceParameter**](MarginAccountNewOrderTimeInForceParameter.md) |  | 
 **selfTradePreventionMode** | [**MarginAccountNewOrderSelfTradePreventionModeParameter**](MarginAccountNewOrderSelfTradePreventionModeParameter.md) |  | 
 **trailingDelta** | **int64** | Used with &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; orders. | 
 **autoRepayAtCancel** | **bool** | Only when MARGIN_BUY or AUTO_BORROW_REPAY order takes effect, true means that the debt generated by the order needs to be repaid after the order is cancelled. | 
 **recvWindow** | **int64** |  | 

### Return type

[**MarginAccountNewOrderResponse**](MarginAccountNewOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MarginAccountNewOto

> MarginAccountNewOtoResponse MarginAccountNewOto(ctx).Symbol(symbol).WorkingType(workingType).WorkingSide(workingSide).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).WorkingIcebergQty(workingIcebergQty).PendingType(pendingType).PendingSide(pendingSide).PendingQuantity(pendingQuantity).IsIsolated(isIsolated).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).SideEffectType(sideEffectType).SelfTradePreventionMode(selfTradePreventionMode).AutoRepayAtCancel(autoRepayAtCancel).WorkingClientOrderId(workingClientOrderId).WorkingTimeInForce(workingTimeInForce).PendingClientOrderId(pendingClientOrderId).PendingPrice(pendingPrice).PendingStopPrice(pendingStopPrice).PendingTrailingDelta(pendingTrailingDelta).PendingIcebergQty(pendingIcebergQty).PendingTimeInForce(pendingTimeInForce).Execute()

Margin Account New OTO (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSDT" // string | 
	workingType := models.MarginAccountNewOtoWorkingTypeParameterLimit // MarginAccountNewOtoWorkingTypeParameter | 
	workingSide := models.MarginAccountNewOtoWorkingSideParameterBuy // MarginAccountNewOtoWorkingSideParameter | 
	workingPrice := float64(1.0) // float64 | 
	workingQuantity := float64(1.0) // float64 | Sets the quantity for the working order.
	workingIcebergQty := float64(1.0) // float64 | This can only be used if `workingTimeInForce` is `GTC`.
	pendingType := models.MarginAccountNewOrderTypeParameterLimit // MarginAccountNewOrderTypeParameter | 
	pendingSide := models.MarginAccountNewOrderSideParameterBuy // MarginAccountNewOrderSideParameter | 
	pendingQuantity := float64(1.0) // float64 | Sets the quantity for the pending order.
	isIsolated := models.QueryMarginAccountsOpenOrdersIsIsolatedParameterTrue // QueryMarginAccountsOpenOrdersIsIsolatedParameter |  (optional)
	listClientOrderId := "1" // string | Arbitrary unique ID among open order lists. Automatically generated if not sent.<br/>A new order list with the same listClientOrderId is accepted only when the previous one is filled or completely expired.<br/>`listClientOrderId` is distinct from the `workingClientOrderId` and the `pendingClientOrderId`. (optional)
	newOrderRespType := models.MarginAccountNewOrderNewOrderRespTypeParameterAck // MarginAccountNewOrderNewOrderRespTypeParameter | MARKET and LIMIT order types default to FULL, all other orders default to ACK. (optional)
	sideEffectType := models.MarginAccountNewOtoSideEffectTypeParameterNoSideEffect // MarginAccountNewOtoSideEffectTypeParameter |  (optional)
	selfTradePreventionMode := models.MarginAccountNewOrderSelfTradePreventionModeParameterExpireTaker // MarginAccountNewOrderSelfTradePreventionModeParameter |  (optional)
	autoRepayAtCancel := true // bool | Only when MARGIN_BUY order takes effect, true means that the debt generated by the order needs to be repaid after the order is cancelled. (optional)
	workingClientOrderId := "1" // string | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent. (optional)
	workingTimeInForce := models.MarginAccountNewOrderTimeInForceParameterGtc // MarginAccountNewOrderTimeInForceParameter |  (optional)
	pendingClientOrderId := "1" // string | Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent. (optional)
	pendingPrice := float64(1.0) // float64 |  (optional)
	pendingStopPrice := float64(1.0) // float64 |  (optional)
	pendingTrailingDelta := float64(1.0) // float64 |  (optional)
	pendingIcebergQty := float64(1.0) // float64 | This can only be used if `pendingTimeInForce` is `GTC`. (optional)
	pendingTimeInForce := models.MarginAccountNewOrderTimeInForceParameterGtc // MarginAccountNewOrderTimeInForceParameter |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.MarginAccountNewOto(context.Background()).Symbol(symbol).WorkingType(workingType).WorkingSide(workingSide).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).WorkingIcebergQty(workingIcebergQty).PendingType(pendingType).PendingSide(pendingSide).PendingQuantity(pendingQuantity).IsIsolated(isIsolated).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).SideEffectType(sideEffectType).SelfTradePreventionMode(selfTradePreventionMode).AutoRepayAtCancel(autoRepayAtCancel).WorkingClientOrderId(workingClientOrderId).WorkingTimeInForce(workingTimeInForce).PendingClientOrderId(pendingClientOrderId).PendingPrice(pendingPrice).PendingStopPrice(pendingStopPrice).PendingTrailingDelta(pendingTrailingDelta).PendingIcebergQty(pendingIcebergQty).PendingTimeInForce(pendingTimeInForce).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.MarginAccountNewOto``: %v\n", err)
		return
	}

	// response from `MarginAccountNewOto`: MarginAccountNewOtoResponse
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
 **workingType** | [**MarginAccountNewOtoWorkingTypeParameter**](MarginAccountNewOtoWorkingTypeParameter.md) |  | 
 **workingSide** | [**MarginAccountNewOtoWorkingSideParameter**](MarginAccountNewOtoWorkingSideParameter.md) |  | 
 **workingPrice** | **float64** |  | 
 **workingQuantity** | **float64** | Sets the quantity for the working order. | 
 **workingIcebergQty** | **float64** | This can only be used if &#x60;workingTimeInForce&#x60; is &#x60;GTC&#x60;. | 
 **pendingType** | [**MarginAccountNewOrderTypeParameter**](MarginAccountNewOrderTypeParameter.md) |  | 
 **pendingSide** | [**MarginAccountNewOrderSideParameter**](MarginAccountNewOrderSideParameter.md) |  | 
 **pendingQuantity** | **float64** | Sets the quantity for the pending order. | 
 **isIsolated** | [**QueryMarginAccountsOpenOrdersIsIsolatedParameter**](QueryMarginAccountsOpenOrdersIsIsolatedParameter.md) |  | 
 **listClientOrderId** | **string** | Arbitrary unique ID among open order lists. Automatically generated if not sent.&lt;br/&gt;A new order list with the same listClientOrderId is accepted only when the previous one is filled or completely expired.&lt;br/&gt;&#x60;listClientOrderId&#x60; is distinct from the &#x60;workingClientOrderId&#x60; and the &#x60;pendingClientOrderId&#x60;. | 
 **newOrderRespType** | [**MarginAccountNewOrderNewOrderRespTypeParameter**](MarginAccountNewOrderNewOrderRespTypeParameter.md) | MARKET and LIMIT order types default to FULL, all other orders default to ACK. | 
 **sideEffectType** | [**MarginAccountNewOtoSideEffectTypeParameter**](MarginAccountNewOtoSideEffectTypeParameter.md) |  | 
 **selfTradePreventionMode** | [**MarginAccountNewOrderSelfTradePreventionModeParameter**](MarginAccountNewOrderSelfTradePreventionModeParameter.md) |  | 
 **autoRepayAtCancel** | **bool** | Only when MARGIN_BUY order takes effect, true means that the debt generated by the order needs to be repaid after the order is cancelled. | 
 **workingClientOrderId** | **string** | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent. | 
 **workingTimeInForce** | [**MarginAccountNewOrderTimeInForceParameter**](MarginAccountNewOrderTimeInForceParameter.md) |  | 
 **pendingClientOrderId** | **string** | Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent. | 
 **pendingPrice** | **float64** |  | 
 **pendingStopPrice** | **float64** |  | 
 **pendingTrailingDelta** | **float64** |  | 
 **pendingIcebergQty** | **float64** | This can only be used if &#x60;pendingTimeInForce&#x60; is &#x60;GTC&#x60;. | 
 **pendingTimeInForce** | [**MarginAccountNewOrderTimeInForceParameter**](MarginAccountNewOrderTimeInForceParameter.md) |  | 

### Return type

[**MarginAccountNewOtoResponse**](MarginAccountNewOtoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MarginAccountNewOtoco

> MarginAccountNewOtocoResponse MarginAccountNewOtoco(ctx).Symbol(symbol).WorkingType(workingType).WorkingSide(workingSide).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).PendingSide(pendingSide).PendingQuantity(pendingQuantity).PendingAboveType(pendingAboveType).IsIsolated(isIsolated).SideEffectType(sideEffectType).AutoRepayAtCancel(autoRepayAtCancel).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingTimeInForce(workingTimeInForce).PendingAboveClientOrderId(pendingAboveClientOrderId).PendingAbovePrice(pendingAbovePrice).PendingAboveStopPrice(pendingAboveStopPrice).PendingAboveTrailingDelta(pendingAboveTrailingDelta).PendingAboveIcebergQty(pendingAboveIcebergQty).PendingAboveTimeInForce(pendingAboveTimeInForce).PendingBelowType(pendingBelowType).PendingBelowClientOrderId(pendingBelowClientOrderId).PendingBelowPrice(pendingBelowPrice).PendingBelowStopPrice(pendingBelowStopPrice).PendingBelowTrailingDelta(pendingBelowTrailingDelta).PendingBelowIcebergQty(pendingBelowIcebergQty).PendingBelowTimeInForce(pendingBelowTimeInForce).Execute()

Margin Account New OTOCO (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSDT" // string | 
	workingType := models.MarginAccountNewOtoWorkingTypeParameterLimit // MarginAccountNewOtoWorkingTypeParameter | 
	workingSide := models.MarginAccountNewOtoWorkingSideParameterBuy // MarginAccountNewOtoWorkingSideParameter | 
	workingPrice := float64(1.0) // float64 | 
	workingQuantity := float64(1.0) // float64 | 
	pendingSide := models.MarginAccountNewOrderSideParameterBuy // MarginAccountNewOrderSideParameter | 
	pendingQuantity := float64(1.0) // float64 | 
	pendingAboveType := models.MarginAccountNewOtocoPendingAboveTypeParameterLimitMaker // MarginAccountNewOtocoPendingAboveTypeParameter | 
	isIsolated := models.QueryMarginAccountsOpenOrdersIsIsolatedParameterTrue // QueryMarginAccountsOpenOrdersIsIsolatedParameter |  (optional)
	sideEffectType := models.MarginAccountNewOtoSideEffectTypeParameterNoSideEffect // MarginAccountNewOtoSideEffectTypeParameter |  (optional)
	autoRepayAtCancel := true // bool | Only when MARGIN_BUY order takes effect, true means that the debt generated by the order needs to be repaid after the order is cancelled. (optional)
	listClientOrderId := "1" // string | Arbitrary unique ID among open order lists. Automatically generated if not sent. A new order list with the same listClientOrderId is accepted only when the previous one is filled or completely expired. `listClientOrderId` is distinct from the `workingClientOrderId`, `pendingAboveClientOrderId`, and the `pendingBelowClientOrderId`. (optional)
	newOrderRespType := models.MarginAccountNewOrderNewOrderRespTypeParameterAck // MarginAccountNewOrderNewOrderRespTypeParameter |  (optional)
	selfTradePreventionMode := models.MarginAccountNewOrderSelfTradePreventionModeParameterExpireTaker // MarginAccountNewOrderSelfTradePreventionModeParameter |  (optional)
	workingClientOrderId := "1" // string | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent. (optional)
	workingIcebergQty := float64(1.0) // float64 | This can only be used if `workingTimeInForce` is `GTC`. (optional)
	workingTimeInForce := models.MarginAccountNewOtocoWorkingTimeInForceParameterGtc // MarginAccountNewOtocoWorkingTimeInForceParameter |  (optional)
	pendingAboveClientOrderId := "1" // string | Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent. (optional)
	pendingAbovePrice := float64(1.0) // float64 |  (optional)
	pendingAboveStopPrice := float64(1.0) // float64 |  (optional)
	pendingAboveTrailingDelta := float64(1.0) // float64 |  (optional)
	pendingAboveIcebergQty := float64(1.0) // float64 | This can only be used if `pendingAboveTimeInForce` is `GTC`. (optional)
	pendingAboveTimeInForce := models.MarginAccountNewOtocoWorkingTimeInForceParameterGtc // MarginAccountNewOtocoWorkingTimeInForceParameter |  (optional)
	pendingBelowType := models.MarginAccountNewOtocoPendingAboveTypeParameterLimitMaker // MarginAccountNewOtocoPendingAboveTypeParameter |  (optional)
	pendingBelowClientOrderId := "1" // string | Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent. (optional)
	pendingBelowPrice := float64(1.0) // float64 |  (optional)
	pendingBelowStopPrice := float64(1.0) // float64 |  (optional)
	pendingBelowTrailingDelta := float64(1.0) // float64 |  (optional)
	pendingBelowIcebergQty := float64(1.0) // float64 | This can only be used if `pendingBelowTimeInForce` is `GTC`. (optional)
	pendingBelowTimeInForce := models.MarginAccountNewOtocoWorkingTimeInForceParameterGtc // MarginAccountNewOtocoWorkingTimeInForceParameter |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.MarginAccountNewOtoco(context.Background()).Symbol(symbol).WorkingType(workingType).WorkingSide(workingSide).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).PendingSide(pendingSide).PendingQuantity(pendingQuantity).PendingAboveType(pendingAboveType).IsIsolated(isIsolated).SideEffectType(sideEffectType).AutoRepayAtCancel(autoRepayAtCancel).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingTimeInForce(workingTimeInForce).PendingAboveClientOrderId(pendingAboveClientOrderId).PendingAbovePrice(pendingAbovePrice).PendingAboveStopPrice(pendingAboveStopPrice).PendingAboveTrailingDelta(pendingAboveTrailingDelta).PendingAboveIcebergQty(pendingAboveIcebergQty).PendingAboveTimeInForce(pendingAboveTimeInForce).PendingBelowType(pendingBelowType).PendingBelowClientOrderId(pendingBelowClientOrderId).PendingBelowPrice(pendingBelowPrice).PendingBelowStopPrice(pendingBelowStopPrice).PendingBelowTrailingDelta(pendingBelowTrailingDelta).PendingBelowIcebergQty(pendingBelowIcebergQty).PendingBelowTimeInForce(pendingBelowTimeInForce).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.MarginAccountNewOtoco``: %v\n", err)
		return
	}

	// response from `MarginAccountNewOtoco`: MarginAccountNewOtocoResponse
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
 **workingType** | [**MarginAccountNewOtoWorkingTypeParameter**](MarginAccountNewOtoWorkingTypeParameter.md) |  | 
 **workingSide** | [**MarginAccountNewOtoWorkingSideParameter**](MarginAccountNewOtoWorkingSideParameter.md) |  | 
 **workingPrice** | **float64** |  | 
 **workingQuantity** | **float64** |  | 
 **pendingSide** | [**MarginAccountNewOrderSideParameter**](MarginAccountNewOrderSideParameter.md) |  | 
 **pendingQuantity** | **float64** |  | 
 **pendingAboveType** | [**MarginAccountNewOtocoPendingAboveTypeParameter**](MarginAccountNewOtocoPendingAboveTypeParameter.md) |  | 
 **isIsolated** | [**QueryMarginAccountsOpenOrdersIsIsolatedParameter**](QueryMarginAccountsOpenOrdersIsIsolatedParameter.md) |  | 
 **sideEffectType** | [**MarginAccountNewOtoSideEffectTypeParameter**](MarginAccountNewOtoSideEffectTypeParameter.md) |  | 
 **autoRepayAtCancel** | **bool** | Only when MARGIN_BUY order takes effect, true means that the debt generated by the order needs to be repaid after the order is cancelled. | 
 **listClientOrderId** | **string** | Arbitrary unique ID among open order lists. Automatically generated if not sent. A new order list with the same listClientOrderId is accepted only when the previous one is filled or completely expired. &#x60;listClientOrderId&#x60; is distinct from the &#x60;workingClientOrderId&#x60;, &#x60;pendingAboveClientOrderId&#x60;, and the &#x60;pendingBelowClientOrderId&#x60;. | 
 **newOrderRespType** | [**MarginAccountNewOrderNewOrderRespTypeParameter**](MarginAccountNewOrderNewOrderRespTypeParameter.md) |  | 
 **selfTradePreventionMode** | [**MarginAccountNewOrderSelfTradePreventionModeParameter**](MarginAccountNewOrderSelfTradePreventionModeParameter.md) |  | 
 **workingClientOrderId** | **string** | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent. | 
 **workingIcebergQty** | **float64** | This can only be used if &#x60;workingTimeInForce&#x60; is &#x60;GTC&#x60;. | 
 **workingTimeInForce** | [**MarginAccountNewOtocoWorkingTimeInForceParameter**](MarginAccountNewOtocoWorkingTimeInForceParameter.md) |  | 
 **pendingAboveClientOrderId** | **string** | Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent. | 
 **pendingAbovePrice** | **float64** |  | 
 **pendingAboveStopPrice** | **float64** |  | 
 **pendingAboveTrailingDelta** | **float64** |  | 
 **pendingAboveIcebergQty** | **float64** | This can only be used if &#x60;pendingAboveTimeInForce&#x60; is &#x60;GTC&#x60;. | 
 **pendingAboveTimeInForce** | [**MarginAccountNewOtocoWorkingTimeInForceParameter**](MarginAccountNewOtocoWorkingTimeInForceParameter.md) |  | 
 **pendingBelowType** | [**MarginAccountNewOtocoPendingAboveTypeParameter**](MarginAccountNewOtocoPendingAboveTypeParameter.md) |  | 
 **pendingBelowClientOrderId** | **string** | Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent. | 
 **pendingBelowPrice** | **float64** |  | 
 **pendingBelowStopPrice** | **float64** |  | 
 **pendingBelowTrailingDelta** | **float64** |  | 
 **pendingBelowIcebergQty** | **float64** | This can only be used if &#x60;pendingBelowTimeInForce&#x60; is &#x60;GTC&#x60;. | 
 **pendingBelowTimeInForce** | [**MarginAccountNewOtocoWorkingTimeInForceParameter**](MarginAccountNewOtocoWorkingTimeInForceParameter.md) |  | 

### Return type

[**MarginAccountNewOtocoResponse**](MarginAccountNewOtocoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MarginManualLiquidation

> MarginManualLiquidationResponse MarginManualLiquidation(ctx).Type(type_).Symbol(symbol).RecvWindow(recvWindow).Execute()

Margin Manual Liquidation (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	type_ := models.QueryMarginAvailableInventoryTypeParameterMargin // QueryMarginAvailableInventoryTypeParameter | 
	symbol := "ETHUSDT" // string | When type selects `ISOLATED`, `symbol` must be filled in (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.MarginManualLiquidation(context.Background()).Type(type_).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.MarginManualLiquidation``: %v\n", err)
		return
	}

	// response from `MarginManualLiquidation`: MarginManualLiquidationResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**QueryMarginAvailableInventoryTypeParameter**](QueryMarginAvailableInventoryTypeParameter.md) |  | 
 **symbol** | **string** | When type selects &#x60;ISOLATED&#x60;, &#x60;symbol&#x60; must be filled in | 
 **recvWindow** | **int64** |  | 

### Return type

[**MarginManualLiquidationResponse**](MarginManualLiquidationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryCurrentMarginOrderCountUsage

> QueryCurrentMarginOrderCountUsageResponse QueryCurrentMarginOrderCountUsage(ctx).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query Current Margin Order Count Usage (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	isIsolated := models.QueryMarginAccountsOpenOrdersIsIsolatedParameterTrue // QueryMarginAccountsOpenOrdersIsIsolatedParameter |  (optional)
	symbol := "BTCUSDT" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryCurrentMarginOrderCountUsage(context.Background()).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryCurrentMarginOrderCountUsage``: %v\n", err)
		return
	}

	// response from `QueryCurrentMarginOrderCountUsage`: QueryCurrentMarginOrderCountUsageResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **isIsolated** | [**QueryMarginAccountsOpenOrdersIsIsolatedParameter**](QueryMarginAccountsOpenOrdersIsIsolatedParameter.md) |  | 
 **symbol** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryCurrentMarginOrderCountUsageResponse**](QueryCurrentMarginOrderCountUsageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryLiquidationLoan

> QueryLiquidationLoanResponse QueryLiquidationLoan(ctx).RecvWindow(recvWindow).Execute()

Query Liquidation Loan (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryLiquidationLoan(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryLiquidationLoan``: %v\n", err)
		return
	}

	// response from `QueryLiquidationLoan`: QueryLiquidationLoanResponse
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

[**QueryLiquidationLoanResponse**](QueryLiquidationLoanResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryLiquidationLoanRepayHistory

> QueryLiquidationLoanRepayHistoryResponse QueryLiquidationLoanRepayHistory(ctx).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Query Liquidation Loan Repay History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	startTime := int64(1714492800000) // int64 | Start time in Unix timestamp (milliseconds). Defaults to 7 days ago if not specified (optional)
	endTime := int64(1714579200000) // int64 | End time in Unix timestamp (milliseconds). Defaults to now if not specified (optional)
	current := int64(1) // int64 | Current page number, default `1` (optional)
	size := int64(50) // int64 | Page size, default `50` (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryLiquidationLoanRepayHistory(context.Background()).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryLiquidationLoanRepayHistory``: %v\n", err)
		return
	}

	// response from `QueryLiquidationLoanRepayHistory`: QueryLiquidationLoanRepayHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** | Start time in Unix timestamp (milliseconds). Defaults to 7 days ago if not specified | 
 **endTime** | **int64** | End time in Unix timestamp (milliseconds). Defaults to now if not specified | 
 **current** | **int64** | Current page number, default &#x60;1&#x60; | 
 **size** | **int64** | Page size, default &#x60;50&#x60; | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryLiquidationLoanRepayHistoryResponse**](QueryLiquidationLoanRepayHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryMarginAccountsAllOco

> QueryMarginAccountsAllOcoResponse QueryMarginAccountsAllOco(ctx).IsIsolated(isIsolated).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query Margin Account's all OCO (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	isIsolated := models.QueryMarginAccountsOpenOrdersIsIsolatedParameterTrue // QueryMarginAccountsOpenOrdersIsIsolatedParameter |  (optional)
	symbol := "LTCBTC" // string |  (optional)
	fromId := int64(1) // int64 |  (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	limit := int64(100) // int64 |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsAllOco(context.Background()).IsIsolated(isIsolated).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
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
 **isIsolated** | [**QueryMarginAccountsOpenOrdersIsIsolatedParameter**](QueryMarginAccountsOpenOrdersIsIsolatedParameter.md) |  | 
 **symbol** | **string** |  | 
 **fromId** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryMarginAccountsAllOcoResponse**](QueryMarginAccountsAllOcoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryMarginAccountsAllOrders

> QueryMarginAccountsAllOrdersResponse QueryMarginAccountsAllOrders(ctx).Symbol(symbol).IsIsolated(isIsolated).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query Margin Account's All Orders (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BNBBTC" // string | 
	isIsolated := models.QueryMarginAccountsOpenOrdersIsIsolatedParameterTrue // QueryMarginAccountsOpenOrdersIsIsolatedParameter |  (optional)
	orderId := int64(1) // int64 |  (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	limit := int64(100) // int64 |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsAllOrders(context.Background()).Symbol(symbol).IsIsolated(isIsolated).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryMarginAccountsAllOrders``: %v\n", err)
		return
	}

	// response from `QueryMarginAccountsAllOrders`: QueryMarginAccountsAllOrdersResponse
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
 **isIsolated** | [**QueryMarginAccountsOpenOrdersIsIsolatedParameter**](QueryMarginAccountsOpenOrdersIsIsolatedParameter.md) |  | 
 **orderId** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryMarginAccountsAllOrdersResponse**](QueryMarginAccountsAllOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryMarginAccountsOco

> QueryMarginAccountsOcoResponse QueryMarginAccountsOco(ctx).IsIsolated(isIsolated).Symbol(symbol).OrderListId(orderListId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Margin Account's OCO (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	isIsolated := models.QueryMarginAccountsOpenOrdersIsIsolatedParameterTrue // QueryMarginAccountsOpenOrdersIsIsolatedParameter |  (optional)
	symbol := "LTCBTC" // string |  (optional)
	orderListId := int64(1) // int64 |  (optional)
	origClientOrderId := "1" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsOco(context.Background()).IsIsolated(isIsolated).Symbol(symbol).OrderListId(orderListId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
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
 **isIsolated** | [**QueryMarginAccountsOpenOrdersIsIsolatedParameter**](QueryMarginAccountsOpenOrdersIsIsolatedParameter.md) |  | 
 **symbol** | **string** |  | 
 **orderListId** | **int64** |  | 
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

> QueryMarginAccountsOpenOcoResponse QueryMarginAccountsOpenOco(ctx).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query Margin Account's Open OCO (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	isIsolated := models.QueryMarginAccountsOpenOrdersIsIsolatedParameterTrue // QueryMarginAccountsOpenOrdersIsIsolatedParameter |  (optional)
	symbol := "LTCBTC" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsOpenOco(context.Background()).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()
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
 **isIsolated** | [**QueryMarginAccountsOpenOrdersIsIsolatedParameter**](QueryMarginAccountsOpenOrdersIsIsolatedParameter.md) |  | 
 **symbol** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryMarginAccountsOpenOcoResponse**](QueryMarginAccountsOpenOcoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryMarginAccountsOpenOrders

> QueryMarginAccountsOpenOrdersResponse QueryMarginAccountsOpenOrders(ctx).Symbol(symbol).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()

Query Margin Account's Open Orders (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BNBBTC" // string | isolated margin pair (optional)
	isIsolated := models.QueryMarginAccountsOpenOrdersIsIsolatedParameterTrue // QueryMarginAccountsOpenOrdersIsIsolatedParameter |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsOpenOrders(context.Background()).Symbol(symbol).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryMarginAccountsOpenOrders``: %v\n", err)
		return
	}

	// response from `QueryMarginAccountsOpenOrders`: QueryMarginAccountsOpenOrdersResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | isolated margin pair | 
 **isIsolated** | [**QueryMarginAccountsOpenOrdersIsIsolatedParameter**](QueryMarginAccountsOpenOrdersIsIsolatedParameter.md) |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryMarginAccountsOpenOrdersResponse**](QueryMarginAccountsOpenOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryMarginAccountsOrder

> QueryMarginAccountsOrderResponse QueryMarginAccountsOrder(ctx).Symbol(symbol).IsIsolated(isIsolated).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Margin Account's Order (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BNBBTC" // string | 
	isIsolated := models.QueryMarginAccountsOpenOrdersIsIsolatedParameterTrue // QueryMarginAccountsOpenOrdersIsIsolatedParameter |  (optional)
	orderId := int64(1) // int64 |  (optional)
	origClientOrderId := "1" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsOrder(context.Background()).Symbol(symbol).IsIsolated(isIsolated).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryMarginAccountsOrder``: %v\n", err)
		return
	}

	// response from `QueryMarginAccountsOrder`: QueryMarginAccountsOrderResponse
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
 **isIsolated** | [**QueryMarginAccountsOpenOrdersIsIsolatedParameter**](QueryMarginAccountsOpenOrdersIsIsolatedParameter.md) |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryMarginAccountsOrderResponse**](QueryMarginAccountsOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryMarginAccountsTradeList

> QueryMarginAccountsTradeListResponse QueryMarginAccountsTradeList(ctx).Symbol(symbol).IsIsolated(isIsolated).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()

Query Margin Account's Trade List (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BNBBTC" // string | 
	isIsolated := models.QueryMarginAccountsOpenOrdersIsIsolatedParameterTrue // QueryMarginAccountsOpenOrdersIsIsolatedParameter |  (optional)
	orderId := int64(1) // int64 |  (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	fromId := int64(1) // int64 |  (optional)
	limit := int64(500) // int64 |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsTradeList(context.Background()).Symbol(symbol).IsIsolated(isIsolated).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryMarginAccountsTradeList``: %v\n", err)
		return
	}

	// response from `QueryMarginAccountsTradeList`: QueryMarginAccountsTradeListResponse
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
 **isIsolated** | [**QueryMarginAccountsOpenOrdersIsIsolatedParameter**](QueryMarginAccountsOpenOrdersIsIsolatedParameter.md) |  | 
 **orderId** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **fromId** | **int64** |  | 
 **limit** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryMarginAccountsTradeListResponse**](QueryMarginAccountsTradeListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryPreventedMatches

> QueryPreventedMatchesResponse QueryPreventedMatches(ctx).Symbol(symbol).PreventedMatchId(preventedMatchId).OrderId(orderId).FromPreventedMatchId(fromPreventedMatchId).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()

Query Prevented Matches (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSDT" // string | 
	preventedMatchId := int64(1) // int64 |  (optional)
	orderId := int64(1) // int64 |  (optional)
	fromPreventedMatchId := int64(1) // int64 |  (optional)
	isIsolated := models.QueryMarginAccountsOpenOrdersIsIsolatedParameterTrue // QueryMarginAccountsOpenOrdersIsIsolatedParameter |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryPreventedMatches(context.Background()).Symbol(symbol).PreventedMatchId(preventedMatchId).OrderId(orderId).FromPreventedMatchId(fromPreventedMatchId).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryPreventedMatches``: %v\n", err)
		return
	}

	// response from `QueryPreventedMatches`: QueryPreventedMatchesResponse
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
 **preventedMatchId** | **int64** |  | 
 **orderId** | **int64** |  | 
 **fromPreventedMatchId** | **int64** |  | 
 **isIsolated** | [**QueryMarginAccountsOpenOrdersIsIsolatedParameter**](QueryMarginAccountsOpenOrdersIsIsolatedParameter.md) |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryPreventedMatchesResponse**](QueryPreventedMatchesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QuerySpecialKey

> QuerySpecialKeyResponse QuerySpecialKey(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query Special key(Low Latency Trading) (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
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
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QuerySpecialKey(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QuerySpecialKey``: %v\n", err)
		return
	}

	// response from `QuerySpecialKey`: QuerySpecialKeyResponse
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

[**QuerySpecialKeyResponse**](QuerySpecialKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QuerySpecialKeyList

> QuerySpecialKeyListResponse QuerySpecialKeyList(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query Special key List(Low Latency Trading) (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
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
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QuerySpecialKeyList(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QuerySpecialKeyList``: %v\n", err)
		return
	}

	// response from `QuerySpecialKeyList`: QuerySpecialKeyListResponse
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

[**QuerySpecialKeyListResponse**](QuerySpecialKeyListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SmallLiabilityExchange

> SmallLiabilityExchange(ctx).AssetNames(assetNames).RecvWindow(recvWindow).Execute()

Small Liability Exchange (MARGIN)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	assetNames := "BTC,ETH" // string | The assets list of small liability exchange
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	 err := apiClient.RestApi.TradeAPI.SmallLiabilityExchange(context.Background()).AssetNames(assetNames).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.SmallLiabilityExchange``: %v\n", err)
		return
	}

}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **assetNames** | **string** | The assets list of small liability exchange | 
 **recvWindow** | **int64** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: Not defined

[[Back to README]](../../../README.md)

