# \AccountManagementAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**CreateAVirtualSubAccount**](AccountManagementAPI.md#CreateAVirtualSubAccount) | **Post** /sapi/v1/sub-account/virtualSubAccount | Create a Virtual Sub-account (For Master Account) (USER_DATA)
[**EnableFuturesForSubAccount**](AccountManagementAPI.md#EnableFuturesForSubAccount) | **Post** /sapi/v1/sub-account/futures/enable | Enable Futures for Sub-account (For Master Account) (USER_DATA)
[**EnableOptionsForSubAccount**](AccountManagementAPI.md#EnableOptionsForSubAccount) | **Post** /sapi/v1/sub-account/eoptions/enable | Enable Options for Sub-account (For Master Account) (USER_DATA)
[**GetFuturesPositionRiskOfSubAccount**](AccountManagementAPI.md#GetFuturesPositionRiskOfSubAccount) | **Get** /sapi/v1/sub-account/futures/positionRisk | Get Futures Position-Risk of Sub-account (For Master Account) (USER_DATA)
[**GetFuturesPositionRiskOfSubAccountV2**](AccountManagementAPI.md#GetFuturesPositionRiskOfSubAccountV2) | **Get** /sapi/v2/sub-account/futures/positionRisk | Get Futures Position-Risk of Sub-account V2 (For Master Account) (USER_DATA)
[**GetSubAccountsStatusOnMarginOrFutures**](AccountManagementAPI.md#GetSubAccountsStatusOnMarginOrFutures) | **Get** /sapi/v1/sub-account/status | Get Sub-account&#39;s Status on Margin Or Futures (For Master Account) (USER_DATA)
[**QuerySubAccountList**](AccountManagementAPI.md#QuerySubAccountList) | **Get** /sapi/v1/sub-account/list | Query Sub-account List (For Master Account) (USER_DATA)
[**QuerySubAccountTransactionStatistics**](AccountManagementAPI.md#QuerySubAccountTransactionStatistics) | **Get** /sapi/v1/sub-account/transaction-statistics | Query Sub-account Transaction Statistics (For Master Account) (USER_DATA)


## CreateAVirtualSubAccount

> CreateAVirtualSubAccountResponse CreateAVirtualSubAccount(ctx).SubAccountString(subAccountString).RecvWindow(recvWindow).Execute()

Create a Virtual Sub-account (For Master Account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	subAccountString := "subAccountString_example" // string | Please input a string. We will create a virtual email using that string for you to register
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountManagementAPI.CreateAVirtualSubAccount(context.Background()).SubAccountString(subAccountString).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountManagementAPI.CreateAVirtualSubAccount``: %v\n", err)
		return
	}

	// response from `CreateAVirtualSubAccount`: CreateAVirtualSubAccountResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **subAccountString** | **string** | Please input a string. We will create a virtual email using that string for you to register | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateAVirtualSubAccountResponse**](CreateAVirtualSubAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## EnableFuturesForSubAccount

> EnableFuturesForSubAccountResponse EnableFuturesForSubAccount(ctx).Email(email).RecvWindow(recvWindow).Execute()

Enable Futures for Sub-account (For Master Account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	email := "sub-account-email@email.com" // string | [Sub-account email](#email-address)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountManagementAPI.EnableFuturesForSubAccount(context.Background()).Email(email).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountManagementAPI.EnableFuturesForSubAccount``: %v\n", err)
		return
	}

	// response from `EnableFuturesForSubAccount`: EnableFuturesForSubAccountResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | [Sub-account email](#email-address) | 
 **recvWindow** | **int64** |  | 

### Return type

[**EnableFuturesForSubAccountResponse**](EnableFuturesForSubAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## EnableOptionsForSubAccount

> EnableOptionsForSubAccountResponse EnableOptionsForSubAccount(ctx).Email(email).RecvWindow(recvWindow).Execute()

Enable Options for Sub-account (For Master Account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	email := "sub-account-email@email.com" // string | [Sub-account email](#email-address)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountManagementAPI.EnableOptionsForSubAccount(context.Background()).Email(email).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountManagementAPI.EnableOptionsForSubAccount``: %v\n", err)
		return
	}

	// response from `EnableOptionsForSubAccount`: EnableOptionsForSubAccountResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | [Sub-account email](#email-address) | 
 **recvWindow** | **int64** |  | 

### Return type

[**EnableOptionsForSubAccountResponse**](EnableOptionsForSubAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetFuturesPositionRiskOfSubAccount

> GetFuturesPositionRiskOfSubAccountResponse GetFuturesPositionRiskOfSubAccount(ctx).Email(email).RecvWindow(recvWindow).Execute()

Get Futures Position-Risk of Sub-account (For Master Account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	email := "sub-account-email@email.com" // string | [Sub-account email](#email-address)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountManagementAPI.GetFuturesPositionRiskOfSubAccount(context.Background()).Email(email).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountManagementAPI.GetFuturesPositionRiskOfSubAccount``: %v\n", err)
		return
	}

	// response from `GetFuturesPositionRiskOfSubAccount`: GetFuturesPositionRiskOfSubAccountResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | [Sub-account email](#email-address) | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetFuturesPositionRiskOfSubAccountResponse**](GetFuturesPositionRiskOfSubAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetFuturesPositionRiskOfSubAccountV2

> GetFuturesPositionRiskOfSubAccountV2Response GetFuturesPositionRiskOfSubAccountV2(ctx).Email(email).FuturesType(futuresType).RecvWindow(recvWindow).Execute()

Get Futures Position-Risk of Sub-account V2 (For Master Account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	email := "sub-account-email@email.com" // string | [Sub-account email](#email-address)
	futuresType := int64(789) // int64 | 1:USDT-margined Futures，2: Coin-margined Futures
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountManagementAPI.GetFuturesPositionRiskOfSubAccountV2(context.Background()).Email(email).FuturesType(futuresType).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountManagementAPI.GetFuturesPositionRiskOfSubAccountV2``: %v\n", err)
		return
	}

	// response from `GetFuturesPositionRiskOfSubAccountV2`: GetFuturesPositionRiskOfSubAccountV2Response
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | [Sub-account email](#email-address) | 
 **futuresType** | **int64** | 1:USDT-margined Futures，2: Coin-margined Futures | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetFuturesPositionRiskOfSubAccountV2Response**](GetFuturesPositionRiskOfSubAccountV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetSubAccountsStatusOnMarginOrFutures

> GetSubAccountsStatusOnMarginOrFuturesResponse GetSubAccountsStatusOnMarginOrFutures(ctx).Email(email).RecvWindow(recvWindow).Execute()

Get Sub-account's Status on Margin Or Futures (For Master Account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	email := "email_example" // string | Managed sub-account email (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountManagementAPI.GetSubAccountsStatusOnMarginOrFutures(context.Background()).Email(email).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountManagementAPI.GetSubAccountsStatusOnMarginOrFutures``: %v\n", err)
		return
	}

	// response from `GetSubAccountsStatusOnMarginOrFutures`: GetSubAccountsStatusOnMarginOrFuturesResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Managed sub-account email | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSubAccountsStatusOnMarginOrFuturesResponse**](GetSubAccountsStatusOnMarginOrFuturesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QuerySubAccountList

> QuerySubAccountListResponse QuerySubAccountList(ctx).Email(email).IsFreeze(isFreeze).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Query Sub-account List (For Master Account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	email := "email_example" // string | Managed sub-account email (optional)
	isFreeze := "isFreeze_example" // string | true or false (optional)
	page := int64(1) // int64 | Default value: 1 (optional)
	limit := int64(1) // int64 | Default value: 1, Max value: 200 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountManagementAPI.QuerySubAccountList(context.Background()).Email(email).IsFreeze(isFreeze).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountManagementAPI.QuerySubAccountList``: %v\n", err)
		return
	}

	// response from `QuerySubAccountList`: QuerySubAccountListResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Managed sub-account email | 
 **isFreeze** | **string** | true or false | 
 **page** | **int64** | Default value: 1 | 
 **limit** | **int64** | Default value: 1, Max value: 200 | 
 **recvWindow** | **int64** |  | 

### Return type

[**QuerySubAccountListResponse**](QuerySubAccountListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QuerySubAccountTransactionStatistics

> QuerySubAccountTransactionStatisticsResponse QuerySubAccountTransactionStatistics(ctx).Email(email).RecvWindow(recvWindow).Execute()

Query Sub-account Transaction Statistics (For Master Account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	email := "email_example" // string | Managed sub-account email (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountManagementAPI.QuerySubAccountTransactionStatistics(context.Background()).Email(email).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountManagementAPI.QuerySubAccountTransactionStatistics``: %v\n", err)
		return
	}

	// response from `QuerySubAccountTransactionStatistics`: QuerySubAccountTransactionStatisticsResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Managed sub-account email | 
 **recvWindow** | **int64** |  | 

### Return type

[**QuerySubAccountTransactionStatisticsResponse**](QuerySubAccountTransactionStatisticsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

