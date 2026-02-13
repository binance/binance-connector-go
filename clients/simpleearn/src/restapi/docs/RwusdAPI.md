# \RwusdAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**GetRwusdAccount**](RwusdAPI.md#GetRwusdAccount) | **Get** /sapi/v1/rwusd/account | Get RWUSD Account (USER_DATA)
[**GetRwusdQuotaDetails**](RwusdAPI.md#GetRwusdQuotaDetails) | **Get** /sapi/v1/rwusd/quota | Get RWUSD Quota Details (USER_DATA)
[**GetRwusdRateHistory**](RwusdAPI.md#GetRwusdRateHistory) | **Get** /sapi/v1/rwusd/history/rateHistory | Get RWUSD Rate History (USER_DATA)
[**GetRwusdRedemptionHistory**](RwusdAPI.md#GetRwusdRedemptionHistory) | **Get** /sapi/v1/rwusd/history/redemptionHistory | Get RWUSD Redemption History (USER_DATA)
[**GetRwusdRewardsHistory**](RwusdAPI.md#GetRwusdRewardsHistory) | **Get** /sapi/v1/rwusd/history/rewardsHistory | Get RWUSD Rewards History (USER_DATA)
[**GetRwusdSubscriptionHistory**](RwusdAPI.md#GetRwusdSubscriptionHistory) | **Get** /sapi/v1/rwusd/history/subscriptionHistory | Get RWUSD subscription history(USER_DATA)
[**RedeemRwusd**](RwusdAPI.md#RedeemRwusd) | **Post** /sapi/v1/rwusd/redeem | Redeem RWUSD(TRADE)
[**SubscribeRwusd**](RwusdAPI.md#SubscribeRwusd) | **Post** /sapi/v1/rwusd/subscribe | Subscribe RWUSD(TRADE)


## GetRwusdAccount

> GetRwusdAccountResponse GetRwusdAccount(ctx).RecvWindow(recvWindow).Execute()

Get RWUSD Account (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/simpleearn"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (ms) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.RwusdAPI.GetRwusdAccount(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `RwusdAPI.GetRwusdAccount``: %v\n", err)
		return
	}

	// response from `GetRwusdAccount`: GetRwusdAccountResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** | The value cannot be greater than 60000 (ms) | 

### Return type

[**GetRwusdAccountResponse**](GetRwusdAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetRwusdQuotaDetails

> GetRwusdQuotaDetailsResponse GetRwusdQuotaDetails(ctx).RecvWindow(recvWindow).Execute()

Get RWUSD Quota Details (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/simpleearn"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (ms) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.RwusdAPI.GetRwusdQuotaDetails(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `RwusdAPI.GetRwusdQuotaDetails``: %v\n", err)
		return
	}

	// response from `GetRwusdQuotaDetails`: GetRwusdQuotaDetailsResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** | The value cannot be greater than 60000 (ms) | 

### Return type

[**GetRwusdQuotaDetailsResponse**](GetRwusdQuotaDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetRwusdRateHistory

> GetRwusdRateHistoryResponse GetRwusdRateHistory(ctx).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get RWUSD Rate History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/simpleearn"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	current := int64(1) // int64 | Currently querying page. Starts from 1. Default: 1 (optional)
	size := int64(10) // int64 | Number of results per page. Default: 10, Max: 100 (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (ms) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.RwusdAPI.GetRwusdRateHistory(context.Background()).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `RwusdAPI.GetRwusdRateHistory``: %v\n", err)
		return
	}

	// response from `GetRwusdRateHistory`: GetRwusdRateHistoryResponse
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
 **current** | **int64** | Currently querying page. Starts from 1. Default: 1 | 
 **size** | **int64** | Number of results per page. Default: 10, Max: 100 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 (ms) | 

### Return type

[**GetRwusdRateHistoryResponse**](GetRwusdRateHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetRwusdRedemptionHistory

> GetRwusdRedemptionHistoryResponse GetRwusdRedemptionHistory(ctx).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get RWUSD Redemption History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/simpleearn"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	current := int64(1) // int64 | Currently querying page. Starts from 1. Default: 1 (optional)
	size := int64(10) // int64 | Number of results per page. Default: 10, Max: 100 (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (ms) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.RwusdAPI.GetRwusdRedemptionHistory(context.Background()).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `RwusdAPI.GetRwusdRedemptionHistory``: %v\n", err)
		return
	}

	// response from `GetRwusdRedemptionHistory`: GetRwusdRedemptionHistoryResponse
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
 **current** | **int64** | Currently querying page. Starts from 1. Default: 1 | 
 **size** | **int64** | Number of results per page. Default: 10, Max: 100 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 (ms) | 

### Return type

[**GetRwusdRedemptionHistoryResponse**](GetRwusdRedemptionHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetRwusdRewardsHistory

> GetRwusdRewardsHistoryResponse GetRwusdRewardsHistory(ctx).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get RWUSD Rewards History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/simpleearn"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	current := int64(1) // int64 | Currently querying page. Starts from 1. Default: 1 (optional)
	size := int64(10) // int64 | Number of results per page. Default: 10, Max: 100 (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (ms) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.RwusdAPI.GetRwusdRewardsHistory(context.Background()).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `RwusdAPI.GetRwusdRewardsHistory``: %v\n", err)
		return
	}

	// response from `GetRwusdRewardsHistory`: GetRwusdRewardsHistoryResponse
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
 **current** | **int64** | Currently querying page. Starts from 1. Default: 1 | 
 **size** | **int64** | Number of results per page. Default: 10, Max: 100 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 (ms) | 

### Return type

[**GetRwusdRewardsHistoryResponse**](GetRwusdRewardsHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetRwusdSubscriptionHistory

> GetRwusdSubscriptionHistoryResponse GetRwusdSubscriptionHistory(ctx).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get RWUSD subscription history(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/simpleearn"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	asset := "asset_example" // string | USDC or USDT (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	current := int64(1) // int64 | Currently querying page. Starts from 1. Default: 1 (optional)
	size := int64(10) // int64 | Number of results per page. Default: 10, Max: 100 (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (ms) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.RwusdAPI.GetRwusdSubscriptionHistory(context.Background()).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `RwusdAPI.GetRwusdSubscriptionHistory``: %v\n", err)
		return
	}

	// response from `GetRwusdSubscriptionHistory`: GetRwusdSubscriptionHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** | USDC or USDT | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Starts from 1. Default: 1 | 
 **size** | **int64** | Number of results per page. Default: 10, Max: 100 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 (ms) | 

### Return type

[**GetRwusdSubscriptionHistoryResponse**](GetRwusdSubscriptionHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## RedeemRwusd

> RedeemRwusdResponse RedeemRwusd(ctx).Amount(amount).Type(type_).RecvWindow(recvWindow).Execute()

Redeem RWUSD(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/simpleearn"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	amount := float32(1.0) // float32 | Amount
	type_ := "s" // string | FAST or STANDARD, defaults to STANDARD
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (ms) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.RwusdAPI.RedeemRwusd(context.Background()).Amount(amount).Type(type_).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `RwusdAPI.RedeemRwusd``: %v\n", err)
		return
	}

	// response from `RedeemRwusd`: RedeemRwusdResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **float32** | Amount | 
 **type_** | **string** | FAST or STANDARD, defaults to STANDARD | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 (ms) | 

### Return type

[**RedeemRwusdResponse**](RedeemRwusdResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SubscribeRwusd

> SubscribeRwusdResponse SubscribeRwusd(ctx).Asset(asset).Amount(amount).RecvWindow(recvWindow).Execute()

Subscribe RWUSD(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/simpleearn"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	asset := "asset_example" // string | USDT or USDC (whichever is eligible)
	amount := float32(1.0) // float32 | Amount
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (ms) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.RwusdAPI.SubscribeRwusd(context.Background()).Asset(asset).Amount(amount).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `RwusdAPI.SubscribeRwusd``: %v\n", err)
		return
	}

	// response from `SubscribeRwusd`: SubscribeRwusdResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** | USDT or USDC (whichever is eligible) | 
 **amount** | **float32** | Amount | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 (ms) | 

### Return type

[**SubscribeRwusdResponse**](SubscribeRwusdResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

