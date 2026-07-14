# \DefaultAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**Deposit**](DefaultAPI.md#Deposit) | **Post** /sapi/v1/fiat/deposit | Deposit (TRADE)
[**FiatWithdraw**](DefaultAPI.md#FiatWithdraw) | **Post** /sapi/v2/fiat/withdraw | Fiat Withdraw (TRADE)
[**GetFiatDepositWithdrawHistory**](DefaultAPI.md#GetFiatDepositWithdrawHistory) | **Get** /sapi/v1/fiat/orders | Get Fiat Deposit/Withdraw History (USER_DATA)
[**GetFiatPaymentsHistory**](DefaultAPI.md#GetFiatPaymentsHistory) | **Get** /sapi/v1/fiat/payments | Get Fiat Payments History (USER_DATA)
[**GetOrderDetail**](DefaultAPI.md#GetOrderDetail) | **Get** /sapi/v1/fiat/get-order-detail | Get Order Detail (USER_DATA)


## Deposit

> DepositResponse Deposit(ctx).Currency(currency).ApiPaymentMethod(apiPaymentMethod).Amount(amount).RecvWindow(recvWindow).Ext(ext).Execute()

Deposit (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/fiat"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	currency := "currency_example" // string | 
	apiPaymentMethod := models.DepositRequestApiPaymentMethodPix // DepositRequestApiPaymentMethod | 
	amount := "amount_example" // string | deposit amount
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)
	ext := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceFiatClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.DefaultAPI.Deposit(context.Background()).Currency(currency).ApiPaymentMethod(apiPaymentMethod).Amount(amount).RecvWindow(recvWindow).Ext(ext).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.Deposit``: %v\n", err)
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
 **apiPaymentMethod** | [**DepositRequestApiPaymentMethod**](DepositRequestApiPaymentMethod.md) |  | 
 **amount** | **string** | deposit amount | 
 **recvWindow** | **int64** | Request validity window in milliseconds | 
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

Fiat Withdraw (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/fiat"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	currency := "currency_example" // string | Fiat currency, such as BRL, ARS, MXN
	apiPaymentMethod := models.FiatWithdrawRequestApiPaymentMethodBankTransfer // FiatWithdrawRequestApiPaymentMethod | 
	amount := int64(789) // int64 | withdraw amount
	accountInfo := *models.NewFiatWithdrawRequestAccountInfo("1056894222") // FiatWithdrawRequestAccountInfo | 
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)
	ext := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceFiatClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.DefaultAPI.FiatWithdraw(context.Background()).Currency(currency).ApiPaymentMethod(apiPaymentMethod).Amount(amount).AccountInfo(accountInfo).RecvWindow(recvWindow).Ext(ext).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.FiatWithdraw``: %v\n", err)
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
 **currency** | **string** | Fiat currency, such as BRL, ARS, MXN | 
 **apiPaymentMethod** | [**FiatWithdrawRequestApiPaymentMethod**](FiatWithdrawRequestApiPaymentMethod.md) |  | 
 **amount** | **int64** | withdraw amount | 
 **accountInfo** | [**FiatWithdrawRequestAccountInfo**](FiatWithdrawRequestAccountInfo.md) |  | 
 **recvWindow** | **int64** | Request validity window in milliseconds | 
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
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	transactionType := "0" // string | 0: deposit, 1: withdraw
	beginTime := int64(1641782889000) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	page := int64(1) // int64 |  (optional)
	rows := int64(100) // int64 |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceFiatClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.DefaultAPI.GetFiatDepositWithdrawHistory(context.Background()).TransactionType(transactionType).BeginTime(beginTime).EndTime(endTime).Page(page).Rows(rows).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.GetFiatDepositWithdrawHistory``: %v\n", err)
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
 **transactionType** | **string** | 0: deposit, 1: withdraw | 
 **beginTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **page** | **int64** |  | 
 **rows** | **int64** |  | 
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
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	transactionType := "0" // string | 0: buy, 1: sell
	beginTime := int64(1641782889000) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	page := int64(1) // int64 |  (optional)
	rows := int64(100) // int64 |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceFiatClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.DefaultAPI.GetFiatPaymentsHistory(context.Background()).TransactionType(transactionType).BeginTime(beginTime).EndTime(endTime).Page(page).Rows(rows).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.GetFiatPaymentsHistory``: %v\n", err)
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
 **transactionType** | **string** | 0: buy, 1: sell | 
 **beginTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **page** | **int64** |  | 
 **rows** | **int64** |  | 
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

Get Order Detail (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/fiat"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	orderNo := "036752*678" // string | Order ID retrieved from the withdrawal API
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceFiatClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.DefaultAPI.GetOrderDetail(context.Background()).OrderNo(orderNo).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.GetOrderDetail``: %v\n", err)
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
 **orderNo** | **string** | Order ID retrieved from the withdrawal API | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetOrderDetailResponse**](GetOrderDetailResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

