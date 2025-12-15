# \BfusdAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**GetBfusdAccount**](BfusdAPI.md#GetBfusdAccount) | **Get** /sapi/v1/bfusd/account | Get BFUSD Account (USER_DATA)
[**GetBfusdQuotaDetails**](BfusdAPI.md#GetBfusdQuotaDetails) | **Get** /sapi/v1/bfusd/quota | Get BFUSD Quota Details (USER_DATA)
[**GetBfusdRateHistory**](BfusdAPI.md#GetBfusdRateHistory) | **Get** /sapi/v1/bfusd/history/rateHistory | Get BFUSD Rate History (USER_DATA)
[**GetBfusdRedemptionHistory**](BfusdAPI.md#GetBfusdRedemptionHistory) | **Get** /sapi/v1/bfusd/history/redemptionHistory | Get BFUSD Redemption History (USER_DATA)
[**GetBfusdRewardsHistory**](BfusdAPI.md#GetBfusdRewardsHistory) | **Get** /sapi/v1/bfusd/history/rewardsHistory | Get BFUSD Rewards History (USER_DATA)
[**GetBfusdSubscriptionHistory**](BfusdAPI.md#GetBfusdSubscriptionHistory) | **Get** /sapi/v1/bfusd/history/subscriptionHistory | Get BFUSD subscription history(USER_DATA)
[**RedeemBfusd**](BfusdAPI.md#RedeemBfusd) | **Post** /sapi/v1/bfusd/redeem | Redeem BFUSD(TRADE)
[**SubscribeBfusd**](BfusdAPI.md#SubscribeBfusd) | **Post** /sapi/v1/bfusd/subscribe | Subscribe BFUSD(TRADE)


## GetBfusdAccount

> GetBfusdAccountResponse GetBfusdAccount(ctx).RecvWindow(recvWindow).Execute()

Get BFUSD Account (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/simpleearn"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (ms) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.BfusdAPI.GetBfusdAccount(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `BfusdAPI.GetBfusdAccount``: %v\n", err)
		return
	}

	// response from `GetBfusdAccount`: GetBfusdAccountResponse
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

[**GetBfusdAccountResponse**](GetBfusdAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetBfusdQuotaDetails

> GetBfusdQuotaDetailsResponse GetBfusdQuotaDetails(ctx).RecvWindow(recvWindow).Execute()

Get BFUSD Quota Details (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/simpleearn"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (ms) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.BfusdAPI.GetBfusdQuotaDetails(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `BfusdAPI.GetBfusdQuotaDetails``: %v\n", err)
		return
	}

	// response from `GetBfusdQuotaDetails`: GetBfusdQuotaDetailsResponse
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

[**GetBfusdQuotaDetailsResponse**](GetBfusdQuotaDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetBfusdRateHistory

> GetBfusdRateHistoryResponse GetBfusdRateHistory(ctx).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get BFUSD Rate History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/simpleearn"
	"github.com/binance/binance-connector-go/common/common"
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

	resp, err := apiClient.RestApi.BfusdAPI.GetBfusdRateHistory(context.Background()).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `BfusdAPI.GetBfusdRateHistory``: %v\n", err)
		return
	}

	// response from `GetBfusdRateHistory`: GetBfusdRateHistoryResponse
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

[**GetBfusdRateHistoryResponse**](GetBfusdRateHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetBfusdRedemptionHistory

> GetBfusdRedemptionHistoryResponse GetBfusdRedemptionHistory(ctx).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get BFUSD Redemption History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/simpleearn"
	"github.com/binance/binance-connector-go/common/common"
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

	resp, err := apiClient.RestApi.BfusdAPI.GetBfusdRedemptionHistory(context.Background()).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `BfusdAPI.GetBfusdRedemptionHistory``: %v\n", err)
		return
	}

	// response from `GetBfusdRedemptionHistory`: GetBfusdRedemptionHistoryResponse
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

[**GetBfusdRedemptionHistoryResponse**](GetBfusdRedemptionHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetBfusdRewardsHistory

> GetBfusdRewardsHistoryResponse GetBfusdRewardsHistory(ctx).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get BFUSD Rewards History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/simpleearn"
	"github.com/binance/binance-connector-go/common/common"
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

	resp, err := apiClient.RestApi.BfusdAPI.GetBfusdRewardsHistory(context.Background()).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `BfusdAPI.GetBfusdRewardsHistory``: %v\n", err)
		return
	}

	// response from `GetBfusdRewardsHistory`: GetBfusdRewardsHistoryResponse
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

[**GetBfusdRewardsHistoryResponse**](GetBfusdRewardsHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetBfusdSubscriptionHistory

> GetBfusdSubscriptionHistoryResponse GetBfusdSubscriptionHistory(ctx).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get BFUSD subscription history(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/simpleearn"
	"github.com/binance/binance-connector-go/common/common"
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

	resp, err := apiClient.RestApi.BfusdAPI.GetBfusdSubscriptionHistory(context.Background()).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `BfusdAPI.GetBfusdSubscriptionHistory``: %v\n", err)
		return
	}

	// response from `GetBfusdSubscriptionHistory`: GetBfusdSubscriptionHistoryResponse
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

[**GetBfusdSubscriptionHistoryResponse**](GetBfusdSubscriptionHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## RedeemBfusd

> RedeemBfusdResponse RedeemBfusd(ctx).Amount(amount).Type(type_).RecvWindow(recvWindow).Execute()

Redeem BFUSD(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/simpleearn"
	"github.com/binance/binance-connector-go/common/common"
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

	resp, err := apiClient.RestApi.BfusdAPI.RedeemBfusd(context.Background()).Amount(amount).Type(type_).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `BfusdAPI.RedeemBfusd``: %v\n", err)
		return
	}

	// response from `RedeemBfusd`: RedeemBfusdResponse
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

[**RedeemBfusdResponse**](RedeemBfusdResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SubscribeBfusd

> SubscribeBfusdResponse SubscribeBfusd(ctx).Asset(asset).Amount(amount).RecvWindow(recvWindow).Execute()

Subscribe BFUSD(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/simpleearn"
	"github.com/binance/binance-connector-go/common/common"
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

	resp, err := apiClient.RestApi.BfusdAPI.SubscribeBfusd(context.Background()).Asset(asset).Amount(amount).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `BfusdAPI.SubscribeBfusd``: %v\n", err)
		return
	}

	// response from `SubscribeBfusd`: SubscribeBfusdResponse
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

[**SubscribeBfusdResponse**](SubscribeBfusdResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

