# \BorrowRepayAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**GetFutureHourlyInterestRate**](BorrowRepayAPI.md#GetFutureHourlyInterestRate) | **Get** /sapi/v1/margin/next-hourly-interest-rate | Get future hourly interest rate (USER_DATA)
[**GetInterestHistory**](BorrowRepayAPI.md#GetInterestHistory) | **Get** /sapi/v1/margin/interestHistory | Get Interest History (USER_DATA)
[**MarginAccountBorrowRepay**](BorrowRepayAPI.md#MarginAccountBorrowRepay) | **Post** /sapi/v1/margin/borrow-repay | Margin account borrow/repay(MARGIN)
[**QueryBorrowRepayRecordsInMarginAccount**](BorrowRepayAPI.md#QueryBorrowRepayRecordsInMarginAccount) | **Get** /sapi/v1/margin/borrow-repay | Query borrow/repay records in Margin account(USER_DATA)
[**QueryMarginInterestRateHistory**](BorrowRepayAPI.md#QueryMarginInterestRateHistory) | **Get** /sapi/v1/margin/interestRateHistory | Query Margin Interest Rate History (USER_DATA)
[**QueryMaxBorrow**](BorrowRepayAPI.md#QueryMaxBorrow) | **Get** /sapi/v1/margin/maxBorrowable | Query Max Borrow (USER_DATA)


## GetFutureHourlyInterestRate

> GetFutureHourlyInterestRateResponse GetFutureHourlyInterestRate(ctx).Assets(assets).IsIsolated(isIsolated).Execute()

Get future hourly interest rate (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	assets := "assets_example" // string | List of assets, separated by commas, up to 20
	isIsolated := false // bool | for isolated margin or not, \"TRUE\", \"FALSE\"

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.BorrowRepayAPI.GetFutureHourlyInterestRate(context.Background()).Assets(assets).IsIsolated(isIsolated).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `BorrowRepayAPI.GetFutureHourlyInterestRate``: %v\n", err)
		return
	}

	// response from `GetFutureHourlyInterestRate`: GetFutureHourlyInterestRateResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **assets** | **string** | List of assets, separated by commas, up to 20 | 
 **isIsolated** | **bool** | for isolated margin or not, \&quot;TRUE\&quot;, \&quot;FALSE\&quot; | 

### Return type

[**GetFutureHourlyInterestRateResponse**](GetFutureHourlyInterestRateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetInterestHistory

> GetInterestHistoryResponse GetInterestHistory(ctx).Asset(asset).IsolatedSymbol(isolatedSymbol).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Interest History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	asset := "asset_example" // string |  (optional)
	isolatedSymbol := "isolatedSymbol_example" // string | isolated symbol (optional)
	startTime := int64(1623319461670) // int64 | 只支持查询最近90天的数据 (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	current := int64(1) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(10) // int64 | Default:10 Max:100 (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.BorrowRepayAPI.GetInterestHistory(context.Background()).Asset(asset).IsolatedSymbol(isolatedSymbol).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `BorrowRepayAPI.GetInterestHistory``: %v\n", err)
		return
	}

	// response from `GetInterestHistory`: GetInterestHistoryResponse
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
 **isolatedSymbol** | **string** | isolated symbol | 
 **startTime** | **int64** | 只支持查询最近90天的数据 | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10 Max:100 | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**GetInterestHistoryResponse**](GetInterestHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MarginAccountBorrowRepay

> MarginAccountBorrowRepayResponse MarginAccountBorrowRepay(ctx).Asset(asset).IsIsolated(isIsolated).Symbol(symbol).Amount(amount).Type(type_).RecvWindow(recvWindow).Execute()

Margin account borrow/repay(MARGIN)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	asset := "asset_example" // string | 
	isIsolated := "FALSE" // string | `TRUE` for Isolated Margin, `FALSE` for Cross Margin, Default `FALSE`
	symbol := "symbol_example" // string | 
	amount := "amount_example" // string | 
	type_ := "type__example" // string | `MARGIN`,`ISOLATED`
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.BorrowRepayAPI.MarginAccountBorrowRepay(context.Background()).Asset(asset).IsIsolated(isIsolated).Symbol(symbol).Amount(amount).Type(type_).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `BorrowRepayAPI.MarginAccountBorrowRepay``: %v\n", err)
		return
	}

	// response from `MarginAccountBorrowRepay`: MarginAccountBorrowRepayResponse
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
 **isIsolated** | **string** | &#x60;TRUE&#x60; for Isolated Margin, &#x60;FALSE&#x60; for Cross Margin, Default &#x60;FALSE&#x60; | 
 **symbol** | **string** |  | 
 **amount** | **string** |  | 
 **type_** | **string** | &#x60;MARGIN&#x60;,&#x60;ISOLATED&#x60; | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**MarginAccountBorrowRepayResponse**](MarginAccountBorrowRepayResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryBorrowRepayRecordsInMarginAccount

> QueryBorrowRepayRecordsInMarginAccountResponse QueryBorrowRepayRecordsInMarginAccount(ctx).Type(type_).Asset(asset).IsolatedSymbol(isolatedSymbol).TxId(txId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Query borrow/repay records in Margin account(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	type_ := "type__example" // string | `MARGIN`,`ISOLATED`
	asset := "asset_example" // string |  (optional)
	isolatedSymbol := "isolatedSymbol_example" // string | isolated symbol (optional)
	txId := int64(1) // int64 | `tranId` in `POST /sapi/v1/margin/loan` (optional)
	startTime := int64(1623319461670) // int64 | 只支持查询最近90天的数据 (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	current := int64(1) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(10) // int64 | Default:10 Max:100 (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.BorrowRepayAPI.QueryBorrowRepayRecordsInMarginAccount(context.Background()).Type(type_).Asset(asset).IsolatedSymbol(isolatedSymbol).TxId(txId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `BorrowRepayAPI.QueryBorrowRepayRecordsInMarginAccount``: %v\n", err)
		return
	}

	// response from `QueryBorrowRepayRecordsInMarginAccount`: QueryBorrowRepayRecordsInMarginAccountResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | &#x60;MARGIN&#x60;,&#x60;ISOLATED&#x60; | 
 **asset** | **string** |  | 
 **isolatedSymbol** | **string** | isolated symbol | 
 **txId** | **int64** | &#x60;tranId&#x60; in &#x60;POST /sapi/v1/margin/loan&#x60; | 
 **startTime** | **int64** | 只支持查询最近90天的数据 | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10 Max:100 | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**QueryBorrowRepayRecordsInMarginAccountResponse**](QueryBorrowRepayRecordsInMarginAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryMarginInterestRateHistory

> QueryMarginInterestRateHistoryResponse QueryMarginInterestRateHistory(ctx).Asset(asset).VipLevel(vipLevel).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

Query Margin Interest Rate History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	asset := "asset_example" // string | 
	vipLevel := int64(1) // int64 | User's current specific margin data will be returned if vipLevel is omitted (optional)
	startTime := int64(1623319461670) // int64 | 只支持查询最近90天的数据 (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.BorrowRepayAPI.QueryMarginInterestRateHistory(context.Background()).Asset(asset).VipLevel(vipLevel).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `BorrowRepayAPI.QueryMarginInterestRateHistory``: %v\n", err)
		return
	}

	// response from `QueryMarginInterestRateHistory`: QueryMarginInterestRateHistoryResponse
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
 **vipLevel** | **int64** | User&#39;s current specific margin data will be returned if vipLevel is omitted | 
 **startTime** | **int64** | 只支持查询最近90天的数据 | 
 **endTime** | **int64** |  | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**QueryMarginInterestRateHistoryResponse**](QueryMarginInterestRateHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryMaxBorrow

> QueryMaxBorrowResponse QueryMaxBorrow(ctx).Asset(asset).IsolatedSymbol(isolatedSymbol).RecvWindow(recvWindow).Execute()

Query Max Borrow (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	asset := "asset_example" // string | 
	isolatedSymbol := "isolatedSymbol_example" // string | isolated symbol (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.BorrowRepayAPI.QueryMaxBorrow(context.Background()).Asset(asset).IsolatedSymbol(isolatedSymbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `BorrowRepayAPI.QueryMaxBorrow``: %v\n", err)
		return
	}

	// response from `QueryMaxBorrow`: QueryMaxBorrowResponse
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
 **isolatedSymbol** | **string** | isolated symbol | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**QueryMaxBorrowResponse**](QueryMaxBorrowResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

