# \FiatAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**Deposit**](FiatAPI.md#Deposit) | **Post** /sapi/v1/fiat/deposit | Deposit(TRADE)
[**FiatWithdraw**](FiatAPI.md#FiatWithdraw) | **Post** /sapi/v2/fiat/withdraw | Fiat Withdraw(WITHDRAW)
[**GetFiatDepositWithdrawHistory**](FiatAPI.md#GetFiatDepositWithdrawHistory) | **Get** /sapi/v1/fiat/orders | Get Fiat Deposit/Withdraw History (USER_DATA)
[**GetFiatPaymentsHistory**](FiatAPI.md#GetFiatPaymentsHistory) | **Get** /sapi/v1/fiat/payments | Get Fiat Payments History (USER_DATA)
[**GetOrderDetail**](FiatAPI.md#GetOrderDetail) | **Get** /sapi/v1/fiat/get-order-detail | Get Order Detail(USER_DATA)


## Deposit

> DepositResponse Deposit(ctx).Currency(currency).ApiPaymentMethod(apiPaymentMethod).Amount(amount).RecvWindow(recvWindow).Ext(ext).Execute()

Deposit(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/fiat"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	currency := "currency_example" // string | 
	apiPaymentMethod := "apiPaymentMethod_example" // string | 
	amount := int64(789) // int64 | 
	recvWindow := int64(5000) // int64 |  (optional)
	ext := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceFiatClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FiatAPI.Deposit(context.Background()).Currency(currency).ApiPaymentMethod(apiPaymentMethod).Amount(amount).RecvWindow(recvWindow).Ext(ext).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FiatAPI.Deposit``: %v\n", err)
		return
	}

	// response from `Deposit`: DepositResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **currency** | **string** |  | 
 **apiPaymentMethod** | **string** |  | 
 **amount** | **int64** |  | 
 **recvWindow** | **int64** |  | 
 **ext** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

### Return type

[**DepositResponse**](DepositResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## FiatWithdraw

> FiatWithdrawResponse FiatWithdraw(ctx).Currency(currency).ApiPaymentMethod(apiPaymentMethod).Amount(amount).AccountInfo(accountInfo).RecvWindow(recvWindow).Ext(ext).Execute()

Fiat Withdraw(WITHDRAW)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/fiat"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	currency := "currency_example" // string | 
	apiPaymentMethod := "apiPaymentMethod_example" // string | 
	amount := int64(789) // int64 | 
	accountInfo := *models.NewAccountInfo() // AccountInfo | 
	recvWindow := int64(5000) // int64 |  (optional)
	ext := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceFiatClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FiatAPI.FiatWithdraw(context.Background()).Currency(currency).ApiPaymentMethod(apiPaymentMethod).Amount(amount).AccountInfo(accountInfo).RecvWindow(recvWindow).Ext(ext).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FiatAPI.FiatWithdraw``: %v\n", err)
		return
	}

	// response from `FiatWithdraw`: FiatWithdrawResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **currency** | **string** |  | 
 **apiPaymentMethod** | **string** |  | 
 **amount** | **int64** |  | 
 **accountInfo** | [**AccountInfo**](AccountInfo.md) |  | 
 **recvWindow** | **int64** |  | 
 **ext** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

### Return type

[**FiatWithdrawResponse**](FiatWithdrawResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetFiatDepositWithdrawHistory

> GetFiatDepositWithdrawHistoryResponse GetFiatDepositWithdrawHistory(ctx).TransactionType(transactionType).BeginTime(beginTime).EndTime(endTime).Page(page).Rows(rows).RecvWindow(recvWindow).Execute()

Get Fiat Deposit/Withdraw History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/fiat"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	transactionType := "transactionType_example" // string | 0-buy,1-sell
	beginTime := int64(789) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	page := int64(1) // int64 | default 1 (optional)
	rows := int64(100) // int64 | default 100, max 500 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceFiatClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FiatAPI.GetFiatDepositWithdrawHistory(context.Background()).TransactionType(transactionType).BeginTime(beginTime).EndTime(endTime).Page(page).Rows(rows).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FiatAPI.GetFiatDepositWithdrawHistory``: %v\n", err)
		return
	}

	// response from `GetFiatDepositWithdrawHistory`: GetFiatDepositWithdrawHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **transactionType** | **string** | 0-buy,1-sell | 
 **beginTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **page** | **int64** | default 1 | 
 **rows** | **int64** | default 100, max 500 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetFiatDepositWithdrawHistoryResponse**](GetFiatDepositWithdrawHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetFiatPaymentsHistory

> GetFiatPaymentsHistoryResponse GetFiatPaymentsHistory(ctx).TransactionType(transactionType).BeginTime(beginTime).EndTime(endTime).Page(page).Rows(rows).RecvWindow(recvWindow).Execute()

Get Fiat Payments History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/fiat"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	transactionType := "transactionType_example" // string | 0-buy,1-sell
	beginTime := int64(789) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	page := int64(1) // int64 | default 1 (optional)
	rows := int64(100) // int64 | default 100, max 500 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceFiatClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FiatAPI.GetFiatPaymentsHistory(context.Background()).TransactionType(transactionType).BeginTime(beginTime).EndTime(endTime).Page(page).Rows(rows).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FiatAPI.GetFiatPaymentsHistory``: %v\n", err)
		return
	}

	// response from `GetFiatPaymentsHistory`: GetFiatPaymentsHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **transactionType** | **string** | 0-buy,1-sell | 
 **beginTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **page** | **int64** | default 1 | 
 **rows** | **int64** | default 100, max 500 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetFiatPaymentsHistoryResponse**](GetFiatPaymentsHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetOrderDetail

> GetOrderDetailResponse GetOrderDetail(ctx).OrderNo(orderNo).RecvWindow(recvWindow).Execute()

Get Order Detail(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/fiat"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	orderNo := "orderNo_example" // string | order id retrieved from the api call of withdrawal
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceFiatClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FiatAPI.GetOrderDetail(context.Background()).OrderNo(orderNo).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FiatAPI.GetOrderDetail``: %v\n", err)
		return
	}

	// response from `GetOrderDetail`: GetOrderDetailResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **orderNo** | **string** | order id retrieved from the api call of withdrawal | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetOrderDetailResponse**](GetOrderDetailResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

