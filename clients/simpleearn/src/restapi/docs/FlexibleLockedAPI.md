# \FlexibleLockedAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**GetCollateralRecord**](FlexibleLockedAPI.md#GetCollateralRecord) | **Get** /sapi/v1/simple-earn/flexible/history/collateralRecord | Get Collateral Record (USER_DATA)
[**GetFlexiblePersonalLeftQuota**](FlexibleLockedAPI.md#GetFlexiblePersonalLeftQuota) | **Get** /sapi/v1/simple-earn/flexible/personalLeftQuota | Get Flexible Personal Left Quota (USER_DATA)
[**GetFlexibleProductPosition**](FlexibleLockedAPI.md#GetFlexibleProductPosition) | **Get** /sapi/v1/simple-earn/flexible/position | Get Flexible Product Position (USER_DATA)
[**GetFlexibleRedemptionRecord**](FlexibleLockedAPI.md#GetFlexibleRedemptionRecord) | **Get** /sapi/v1/simple-earn/flexible/history/redemptionRecord | Get Flexible Redemption Record (USER_DATA)
[**GetFlexibleRewardsHistory**](FlexibleLockedAPI.md#GetFlexibleRewardsHistory) | **Get** /sapi/v1/simple-earn/flexible/history/rewardsRecord | Get Flexible Rewards History (USER_DATA)
[**GetFlexibleSubscriptionPreview**](FlexibleLockedAPI.md#GetFlexibleSubscriptionPreview) | **Get** /sapi/v1/simple-earn/flexible/subscriptionPreview | Get Flexible Subscription Preview (USER_DATA)
[**GetFlexibleSubscriptionRecord**](FlexibleLockedAPI.md#GetFlexibleSubscriptionRecord) | **Get** /sapi/v1/simple-earn/flexible/history/subscriptionRecord | Get Flexible Subscription Record (USER_DATA)
[**GetLockedPersonalLeftQuota**](FlexibleLockedAPI.md#GetLockedPersonalLeftQuota) | **Get** /sapi/v1/simple-earn/locked/personalLeftQuota | Get Locked Personal Left Quota (USER_DATA)
[**GetLockedProductPosition**](FlexibleLockedAPI.md#GetLockedProductPosition) | **Get** /sapi/v1/simple-earn/locked/position | Get Locked Product Position (USER_DATA)
[**GetLockedRedemptionRecord**](FlexibleLockedAPI.md#GetLockedRedemptionRecord) | **Get** /sapi/v1/simple-earn/locked/history/redemptionRecord | Get Locked Redemption Record (USER_DATA)
[**GetLockedRewardsHistory**](FlexibleLockedAPI.md#GetLockedRewardsHistory) | **Get** /sapi/v1/simple-earn/locked/history/rewardsRecord | Get Locked Rewards History (USER_DATA)
[**GetLockedSubscriptionPreview**](FlexibleLockedAPI.md#GetLockedSubscriptionPreview) | **Get** /sapi/v1/simple-earn/locked/subscriptionPreview | Get Locked Subscription Preview (USER_DATA)
[**GetLockedSubscriptionRecord**](FlexibleLockedAPI.md#GetLockedSubscriptionRecord) | **Get** /sapi/v1/simple-earn/locked/history/subscriptionRecord | Get Locked Subscription Record (USER_DATA)
[**GetRateHistory**](FlexibleLockedAPI.md#GetRateHistory) | **Get** /sapi/v1/simple-earn/flexible/history/rateHistory | Get Rate History (USER_DATA)
[**GetSimpleEarnFlexibleProductList**](FlexibleLockedAPI.md#GetSimpleEarnFlexibleProductList) | **Get** /sapi/v1/simple-earn/flexible/list | Get Simple Earn Flexible Product List (USER_DATA)
[**GetSimpleEarnLockedProductList**](FlexibleLockedAPI.md#GetSimpleEarnLockedProductList) | **Get** /sapi/v1/simple-earn/locked/list | Get Simple Earn Locked Product List (USER_DATA)
[**RedeemFlexibleProduct**](FlexibleLockedAPI.md#RedeemFlexibleProduct) | **Post** /sapi/v1/simple-earn/flexible/redeem | Redeem Flexible Product (TRADE)
[**RedeemLockedProduct**](FlexibleLockedAPI.md#RedeemLockedProduct) | **Post** /sapi/v1/simple-earn/locked/redeem | Redeem Locked Product (TRADE)
[**SetFlexibleAutoSubscribe**](FlexibleLockedAPI.md#SetFlexibleAutoSubscribe) | **Post** /sapi/v1/simple-earn/flexible/setAutoSubscribe | Set Flexible Auto Subscribe (USER_DATA)
[**SetLockedAutoSubscribe**](FlexibleLockedAPI.md#SetLockedAutoSubscribe) | **Post** /sapi/v1/simple-earn/locked/setAutoSubscribe | Set Locked Auto Subscribe (USER_DATA)
[**SetLockedProductRedeemOption**](FlexibleLockedAPI.md#SetLockedProductRedeemOption) | **Post** /sapi/v1/simple-earn/locked/setRedeemOption | Set Locked Product Redeem Option (USER_DATA)
[**SimpleAccount**](FlexibleLockedAPI.md#SimpleAccount) | **Get** /sapi/v1/simple-earn/account | Simple Account (USER_DATA)
[**SubscribeFlexibleProduct**](FlexibleLockedAPI.md#SubscribeFlexibleProduct) | **Post** /sapi/v1/simple-earn/flexible/subscribe | Subscribe Flexible Product (TRADE)
[**SubscribeLockedProduct**](FlexibleLockedAPI.md#SubscribeLockedProduct) | **Post** /sapi/v1/simple-earn/locked/subscribe | Subscribe Locked Product (TRADE)


## GetCollateralRecord

> GetCollateralRecordResponse GetCollateralRecord(ctx).ProductId(productId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Collateral Record (USER_DATA)


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
	productId := "1" // string |  (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	current := int64(1) // int64 | Currently querying page. (optional)
	size := int64(10) // int64 | Number of results per page. (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (ms) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleLockedAPI.GetCollateralRecord(context.Background()).ProductId(productId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleLockedAPI.GetCollateralRecord``: %v\n", err)
		return
	}

	// response from `GetCollateralRecord`: GetCollateralRecordResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. | 
 **size** | **int64** | Number of results per page. | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 (ms) | 

### Return type

[**GetCollateralRecordResponse**](GetCollateralRecordResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetFlexiblePersonalLeftQuota

> GetFlexiblePersonalLeftQuotaResponse GetFlexiblePersonalLeftQuota(ctx).ProductId(productId).RecvWindow(recvWindow).Execute()

Get Flexible Personal Left Quota (USER_DATA)


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
	productId := "1" // string | 
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (ms) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleLockedAPI.GetFlexiblePersonalLeftQuota(context.Background()).ProductId(productId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleLockedAPI.GetFlexiblePersonalLeftQuota``: %v\n", err)
		return
	}

	// response from `GetFlexiblePersonalLeftQuota`: GetFlexiblePersonalLeftQuotaResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** |  | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 (ms) | 

### Return type

[**GetFlexiblePersonalLeftQuotaResponse**](GetFlexiblePersonalLeftQuotaResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetFlexibleProductPosition

> GetFlexibleProductPositionResponse GetFlexibleProductPosition(ctx).Asset(asset).ProductId(productId).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Flexible Product Position (USER_DATA)


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
	asset := "USDC" // string |  (optional)
	productId := "1" // string |  (optional)
	current := int64(1) // int64 | Currently querying page. Starts from 1. (optional)
	size := int64(10) // int64 | Number of results per page. (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (ms) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleLockedAPI.GetFlexibleProductPosition(context.Background()).Asset(asset).ProductId(productId).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleLockedAPI.GetFlexibleProductPosition``: %v\n", err)
		return
	}

	// response from `GetFlexibleProductPosition`: GetFlexibleProductPositionResponse
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
 **productId** | **string** |  | 
 **current** | **int64** | Currently querying page. Starts from 1. | 
 **size** | **int64** | Number of results per page. | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 (ms) | 

### Return type

[**GetFlexibleProductPositionResponse**](GetFlexibleProductPositionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetFlexibleRedemptionRecord

> GetFlexibleRedemptionRecordResponse GetFlexibleRedemptionRecord(ctx).ProductId(productId).RedeemId(redeemId).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Flexible Redemption Record (USER_DATA)


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
	productId := "1" // string |  (optional)
	redeemId := "1" // string |  (optional)
	asset := "USDC" // string |  (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	current := int64(1) // int64 | Currently querying page. Starts from 1. (optional)
	size := int64(10) // int64 | Number of results per page. (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (ms) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleLockedAPI.GetFlexibleRedemptionRecord(context.Background()).ProductId(productId).RedeemId(redeemId).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleLockedAPI.GetFlexibleRedemptionRecord``: %v\n", err)
		return
	}

	// response from `GetFlexibleRedemptionRecord`: GetFlexibleRedemptionRecordResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** |  | 
 **redeemId** | **string** |  | 
 **asset** | **string** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Starts from 1. | 
 **size** | **int64** | Number of results per page. | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 (ms) | 

### Return type

[**GetFlexibleRedemptionRecordResponse**](GetFlexibleRedemptionRecordResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetFlexibleRewardsHistory

> GetFlexibleRewardsHistoryResponse GetFlexibleRewardsHistory(ctx).ProductId(productId).Asset(asset).StartTime(startTime).EndTime(endTime).Type(type_).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Flexible Rewards History (USER_DATA)


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
	productId := "1" // string |  (optional)
	asset := "USDC" // string |  (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	type_ := models.GetFlexibleRewardsHistoryTypeParameterBonus // GetFlexibleRewardsHistoryTypeParameter | `BONUS` - Bonus tiered APR, `REALTIME` - Real-time APR, `REWARDS` - Historical rewards, `ALL` - All types. Default: `ALL` (optional)
	current := int64(1) // int64 | Currently querying page. Starts from 1. (optional)
	size := int64(10) // int64 | Number of results per page. (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (ms) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleLockedAPI.GetFlexibleRewardsHistory(context.Background()).ProductId(productId).Asset(asset).StartTime(startTime).EndTime(endTime).Type(type_).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleLockedAPI.GetFlexibleRewardsHistory``: %v\n", err)
		return
	}

	// response from `GetFlexibleRewardsHistory`: GetFlexibleRewardsHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** |  | 
 **asset** | **string** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **type_** | [**GetFlexibleRewardsHistoryTypeParameter**](GetFlexibleRewardsHistoryTypeParameter.md) | &#x60;BONUS&#x60; - Bonus tiered APR, &#x60;REALTIME&#x60; - Real-time APR, &#x60;REWARDS&#x60; - Historical rewards, &#x60;ALL&#x60; - All types. Default: &#x60;ALL&#x60; | 
 **current** | **int64** | Currently querying page. Starts from 1. | 
 **size** | **int64** | Number of results per page. | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 (ms) | 

### Return type

[**GetFlexibleRewardsHistoryResponse**](GetFlexibleRewardsHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetFlexibleSubscriptionPreview

> GetFlexibleSubscriptionPreviewResponse GetFlexibleSubscriptionPreview(ctx).ProductId(productId).Amount(amount).RecvWindow(recvWindow).Execute()

Get Flexible Subscription Preview (USER_DATA)


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
	productId := "1" // string | 
	amount := float64(1.0) // float64 | 
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (ms) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleLockedAPI.GetFlexibleSubscriptionPreview(context.Background()).ProductId(productId).Amount(amount).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleLockedAPI.GetFlexibleSubscriptionPreview``: %v\n", err)
		return
	}

	// response from `GetFlexibleSubscriptionPreview`: GetFlexibleSubscriptionPreviewResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** |  | 
 **amount** | **float64** |  | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 (ms) | 

### Return type

[**GetFlexibleSubscriptionPreviewResponse**](GetFlexibleSubscriptionPreviewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetFlexibleSubscriptionRecord

> GetFlexibleSubscriptionRecordResponse GetFlexibleSubscriptionRecord(ctx).ProductId(productId).PurchaseId(purchaseId).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Flexible Subscription Record (USER_DATA)


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
	productId := "1" // string |  (optional)
	purchaseId := "1" // string |  (optional)
	asset := "USDC" // string |  (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	current := int64(1) // int64 | Currently querying page. Starts from 1. (optional)
	size := int64(10) // int64 | Number of results per page. (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (ms) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleLockedAPI.GetFlexibleSubscriptionRecord(context.Background()).ProductId(productId).PurchaseId(purchaseId).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleLockedAPI.GetFlexibleSubscriptionRecord``: %v\n", err)
		return
	}

	// response from `GetFlexibleSubscriptionRecord`: GetFlexibleSubscriptionRecordResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** |  | 
 **purchaseId** | **string** |  | 
 **asset** | **string** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Starts from 1. | 
 **size** | **int64** | Number of results per page. | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 (ms) | 

### Return type

[**GetFlexibleSubscriptionRecordResponse**](GetFlexibleSubscriptionRecordResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetLockedPersonalLeftQuota

> GetLockedPersonalLeftQuotaResponse GetLockedPersonalLeftQuota(ctx).ProjectId(projectId).RecvWindow(recvWindow).Execute()

Get Locked Personal Left Quota (USER_DATA)


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
	projectId := "1" // string | 
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (ms) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleLockedAPI.GetLockedPersonalLeftQuota(context.Background()).ProjectId(projectId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleLockedAPI.GetLockedPersonalLeftQuota``: %v\n", err)
		return
	}

	// response from `GetLockedPersonalLeftQuota`: GetLockedPersonalLeftQuotaResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** |  | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 (ms) | 

### Return type

[**GetLockedPersonalLeftQuotaResponse**](GetLockedPersonalLeftQuotaResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetLockedProductPosition

> GetLockedProductPositionResponse GetLockedProductPosition(ctx).Asset(asset).PositionId(positionId).ProjectId(projectId).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Locked Product Position (USER_DATA)


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
	asset := "USDC" // string |  (optional)
	positionId := "1" // string |  (optional)
	projectId := "1" // string |  (optional)
	current := int64(1) // int64 | Currently querying page. Starts from 1. (optional)
	size := int64(10) // int64 | Number of results per page. (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (ms) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleLockedAPI.GetLockedProductPosition(context.Background()).Asset(asset).PositionId(positionId).ProjectId(projectId).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleLockedAPI.GetLockedProductPosition``: %v\n", err)
		return
	}

	// response from `GetLockedProductPosition`: GetLockedProductPositionResponse
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
 **positionId** | **string** |  | 
 **projectId** | **string** |  | 
 **current** | **int64** | Currently querying page. Starts from 1. | 
 **size** | **int64** | Number of results per page. | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 (ms) | 

### Return type

[**GetLockedProductPositionResponse**](GetLockedProductPositionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetLockedRedemptionRecord

> GetLockedRedemptionRecordResponse GetLockedRedemptionRecord(ctx).PositionId(positionId).RedeemId(redeemId).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Locked Redemption Record (USER_DATA)


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
	positionId := "1" // string |  (optional)
	redeemId := "1" // string |  (optional)
	asset := "USDC" // string |  (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	current := int64(1) // int64 | Currently querying page. Starts from 1. (optional)
	size := int64(10) // int64 | Number of results per page. (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (ms) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleLockedAPI.GetLockedRedemptionRecord(context.Background()).PositionId(positionId).RedeemId(redeemId).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleLockedAPI.GetLockedRedemptionRecord``: %v\n", err)
		return
	}

	// response from `GetLockedRedemptionRecord`: GetLockedRedemptionRecordResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **positionId** | **string** |  | 
 **redeemId** | **string** |  | 
 **asset** | **string** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Starts from 1. | 
 **size** | **int64** | Number of results per page. | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 (ms) | 

### Return type

[**GetLockedRedemptionRecordResponse**](GetLockedRedemptionRecordResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetLockedRewardsHistory

> GetLockedRewardsHistoryResponse GetLockedRewardsHistory(ctx).PositionId(positionId).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Locked Rewards History (USER_DATA)


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
	positionId := "1" // string |  (optional)
	asset := "USDC" // string |  (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	current := int64(1) // int64 | Currently querying page. Starts from 1. (optional)
	size := int64(10) // int64 | Number of results per page. (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (ms) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleLockedAPI.GetLockedRewardsHistory(context.Background()).PositionId(positionId).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleLockedAPI.GetLockedRewardsHistory``: %v\n", err)
		return
	}

	// response from `GetLockedRewardsHistory`: GetLockedRewardsHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **positionId** | **string** |  | 
 **asset** | **string** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Starts from 1. | 
 **size** | **int64** | Number of results per page. | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 (ms) | 

### Return type

[**GetLockedRewardsHistoryResponse**](GetLockedRewardsHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetLockedSubscriptionPreview

> GetLockedSubscriptionPreviewResponse GetLockedSubscriptionPreview(ctx).ProjectId(projectId).Amount(amount).AutoSubscribe(autoSubscribe).RecvWindow(recvWindow).Execute()

Get Locked Subscription Preview (USER_DATA)


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
	projectId := "1" // string | 
	amount := float64(1.0) // float64 | 
	autoSubscribe := true // bool | default true. (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (ms) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleLockedAPI.GetLockedSubscriptionPreview(context.Background()).ProjectId(projectId).Amount(amount).AutoSubscribe(autoSubscribe).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleLockedAPI.GetLockedSubscriptionPreview``: %v\n", err)
		return
	}

	// response from `GetLockedSubscriptionPreview`: GetLockedSubscriptionPreviewResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** |  | 
 **amount** | **float64** |  | 
 **autoSubscribe** | **bool** | default true. | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 (ms) | 

### Return type

[**GetLockedSubscriptionPreviewResponse**](GetLockedSubscriptionPreviewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetLockedSubscriptionRecord

> GetLockedSubscriptionRecordResponse GetLockedSubscriptionRecord(ctx).PurchaseId(purchaseId).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Locked Subscription Record (USER_DATA)


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
	purchaseId := "1" // string |  (optional)
	asset := "USDC" // string |  (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	current := int64(1) // int64 | Currently querying page. Starts from 1. (optional)
	size := int64(10) // int64 | Number of results per page. (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (ms) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleLockedAPI.GetLockedSubscriptionRecord(context.Background()).PurchaseId(purchaseId).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleLockedAPI.GetLockedSubscriptionRecord``: %v\n", err)
		return
	}

	// response from `GetLockedSubscriptionRecord`: GetLockedSubscriptionRecordResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **purchaseId** | **string** |  | 
 **asset** | **string** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Starts from 1. | 
 **size** | **int64** | Number of results per page. | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 (ms) | 

### Return type

[**GetLockedSubscriptionRecordResponse**](GetLockedSubscriptionRecordResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetRateHistory

> GetRateHistoryResponse GetRateHistory(ctx).ProductId(productId).AprPeriod(aprPeriod).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Rate History (USER_DATA)


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
	productId := "1" // string | 
	aprPeriod := models.GetRateHistoryAprPeriodParameterDay // GetRateHistoryAprPeriodParameter |  (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	current := int64(1) // int64 | Currently querying page (optional)
	size := int64(10) // int64 | Number of results per page (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (ms) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleLockedAPI.GetRateHistory(context.Background()).ProductId(productId).AprPeriod(aprPeriod).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleLockedAPI.GetRateHistory``: %v\n", err)
		return
	}

	// response from `GetRateHistory`: GetRateHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** |  | 
 **aprPeriod** | [**GetRateHistoryAprPeriodParameter**](GetRateHistoryAprPeriodParameter.md) |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page | 
 **size** | **int64** | Number of results per page | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 (ms) | 

### Return type

[**GetRateHistoryResponse**](GetRateHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetSimpleEarnFlexibleProductList

> GetSimpleEarnFlexibleProductListResponse GetSimpleEarnFlexibleProductList(ctx).Asset(asset).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Simple Earn Flexible Product List (USER_DATA)


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
	asset := "USDC" // string |  (optional)
	current := int64(1) // int64 | Currently querying page. Starts from 1. (optional)
	size := int64(10) // int64 | Number of results per page. (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (ms) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleLockedAPI.GetSimpleEarnFlexibleProductList(context.Background()).Asset(asset).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleLockedAPI.GetSimpleEarnFlexibleProductList``: %v\n", err)
		return
	}

	// response from `GetSimpleEarnFlexibleProductList`: GetSimpleEarnFlexibleProductListResponse
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
 **current** | **int64** | Currently querying page. Starts from 1. | 
 **size** | **int64** | Number of results per page. | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 (ms) | 

### Return type

[**GetSimpleEarnFlexibleProductListResponse**](GetSimpleEarnFlexibleProductListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetSimpleEarnLockedProductList

> GetSimpleEarnLockedProductListResponse GetSimpleEarnLockedProductList(ctx).Asset(asset).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Simple Earn Locked Product List (USER_DATA)


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
	asset := "USDC" // string |  (optional)
	current := int64(1) // int64 | Currently querying page. Starts from 1. (optional)
	size := int64(10) // int64 | Number of results per page. (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (ms) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleLockedAPI.GetSimpleEarnLockedProductList(context.Background()).Asset(asset).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleLockedAPI.GetSimpleEarnLockedProductList``: %v\n", err)
		return
	}

	// response from `GetSimpleEarnLockedProductList`: GetSimpleEarnLockedProductListResponse
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
 **current** | **int64** | Currently querying page. Starts from 1. | 
 **size** | **int64** | Number of results per page. | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 (ms) | 

### Return type

[**GetSimpleEarnLockedProductListResponse**](GetSimpleEarnLockedProductListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## RedeemFlexibleProduct

> RedeemFlexibleProductResponse RedeemFlexibleProduct(ctx).ProductId(productId).RedeemAll(redeemAll).Amount(amount).DestAccount(destAccount).RecvWindow(recvWindow).Execute()

Redeem Flexible Product (TRADE)


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
	productId := "1" // string | 
	redeemAll := false // bool |  (optional)
	amount := float64(1.0) // float64 | if redeemAll is false, amount is mandatory (optional)
	destAccount := models.RedeemFlexibleProductDestAccountParameterSpot // RedeemFlexibleProductDestAccountParameter |  (optional)
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleLockedAPI.RedeemFlexibleProduct(context.Background()).ProductId(productId).RedeemAll(redeemAll).Amount(amount).DestAccount(destAccount).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleLockedAPI.RedeemFlexibleProduct``: %v\n", err)
		return
	}

	// response from `RedeemFlexibleProduct`: RedeemFlexibleProductResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** |  | 
 **redeemAll** | **bool** |  | 
 **amount** | **float64** | if redeemAll is false, amount is mandatory | 
 **destAccount** | [**RedeemFlexibleProductDestAccountParameter**](RedeemFlexibleProductDestAccountParameter.md) |  | 
 **recvWindow** | **int64** | Request validity window in milliseconds. | 

### Return type

[**RedeemFlexibleProductResponse**](RedeemFlexibleProductResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## RedeemLockedProduct

> RedeemLockedProductResponse RedeemLockedProduct(ctx).PositionId(positionId).RecvWindow(recvWindow).Execute()

Redeem Locked Product (TRADE)


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
	positionId := "1" // string | Locked product position ID
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleLockedAPI.RedeemLockedProduct(context.Background()).PositionId(positionId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleLockedAPI.RedeemLockedProduct``: %v\n", err)
		return
	}

	// response from `RedeemLockedProduct`: RedeemLockedProductResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **positionId** | **string** | Locked product position ID | 
 **recvWindow** | **int64** | Request validity window in milliseconds. | 

### Return type

[**RedeemLockedProductResponse**](RedeemLockedProductResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SetFlexibleAutoSubscribe

> SetFlexibleAutoSubscribeResponse SetFlexibleAutoSubscribe(ctx).ProductId(productId).AutoSubscribe(autoSubscribe).RecvWindow(recvWindow).Execute()

Set Flexible Auto Subscribe (USER_DATA)


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
	productId := "1" // string | 
	autoSubscribe := true // bool | 
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleLockedAPI.SetFlexibleAutoSubscribe(context.Background()).ProductId(productId).AutoSubscribe(autoSubscribe).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleLockedAPI.SetFlexibleAutoSubscribe``: %v\n", err)
		return
	}

	// response from `SetFlexibleAutoSubscribe`: SetFlexibleAutoSubscribeResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** |  | 
 **autoSubscribe** | **bool** |  | 
 **recvWindow** | **int64** | Request validity window in milliseconds. | 

### Return type

[**SetFlexibleAutoSubscribeResponse**](SetFlexibleAutoSubscribeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SetLockedAutoSubscribe

> SetLockedAutoSubscribeResponse SetLockedAutoSubscribe(ctx).PositionId(positionId).AutoSubscribe(autoSubscribe).RecvWindow(recvWindow).Execute()

Set Locked Auto Subscribe (USER_DATA)


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
	positionId := "1" // string | 
	autoSubscribe := true // bool | 
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleLockedAPI.SetLockedAutoSubscribe(context.Background()).PositionId(positionId).AutoSubscribe(autoSubscribe).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleLockedAPI.SetLockedAutoSubscribe``: %v\n", err)
		return
	}

	// response from `SetLockedAutoSubscribe`: SetLockedAutoSubscribeResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **positionId** | **string** |  | 
 **autoSubscribe** | **bool** |  | 
 **recvWindow** | **int64** | Request validity window in milliseconds. | 

### Return type

[**SetLockedAutoSubscribeResponse**](SetLockedAutoSubscribeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SetLockedProductRedeemOption

> SetLockedProductRedeemOptionResponse SetLockedProductRedeemOption(ctx).PositionId(positionId).RedeemTo(redeemTo).RecvWindow(recvWindow).Execute()

Set Locked Product Redeem Option (USER_DATA)


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
	positionId := "1" // string | 
	redeemTo := models.SetLockedProductRedeemOptionRedeemToParameterSpot // SetLockedProductRedeemOptionRedeemToParameter | 
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleLockedAPI.SetLockedProductRedeemOption(context.Background()).PositionId(positionId).RedeemTo(redeemTo).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleLockedAPI.SetLockedProductRedeemOption``: %v\n", err)
		return
	}

	// response from `SetLockedProductRedeemOption`: SetLockedProductRedeemOptionResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **positionId** | **string** |  | 
 **redeemTo** | [**SetLockedProductRedeemOptionRedeemToParameter**](SetLockedProductRedeemOptionRedeemToParameter.md) |  | 
 **recvWindow** | **int64** | Request validity window in milliseconds. | 

### Return type

[**SetLockedProductRedeemOptionResponse**](SetLockedProductRedeemOptionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SimpleAccount

> SimpleAccountResponse SimpleAccount(ctx).RecvWindow(recvWindow).Execute()

Simple Account (USER_DATA)


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

	resp, err := apiClient.RestApi.FlexibleLockedAPI.SimpleAccount(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleLockedAPI.SimpleAccount``: %v\n", err)
		return
	}

	// response from `SimpleAccount`: SimpleAccountResponse
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

[**SimpleAccountResponse**](SimpleAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SubscribeFlexibleProduct

> SubscribeFlexibleProductResponse SubscribeFlexibleProduct(ctx).ProductId(productId).Amount(amount).AutoSubscribe(autoSubscribe).SourceAccount(sourceAccount).RecvWindow(recvWindow).Execute()

Subscribe Flexible Product (TRADE)


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
	productId := "1" // string | 
	amount := float64(1.0) // float64 | 
	autoSubscribe := true // bool |  (optional)
	sourceAccount := models.SubscribeFlexibleProductSourceAccountParameterSpot // SubscribeFlexibleProductSourceAccountParameter |  (optional)
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleLockedAPI.SubscribeFlexibleProduct(context.Background()).ProductId(productId).Amount(amount).AutoSubscribe(autoSubscribe).SourceAccount(sourceAccount).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleLockedAPI.SubscribeFlexibleProduct``: %v\n", err)
		return
	}

	// response from `SubscribeFlexibleProduct`: SubscribeFlexibleProductResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** |  | 
 **amount** | **float64** |  | 
 **autoSubscribe** | **bool** |  | 
 **sourceAccount** | [**SubscribeFlexibleProductSourceAccountParameter**](SubscribeFlexibleProductSourceAccountParameter.md) |  | 
 **recvWindow** | **int64** | Request validity window in milliseconds. | 

### Return type

[**SubscribeFlexibleProductResponse**](SubscribeFlexibleProductResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SubscribeLockedProduct

> SubscribeLockedProductResponse SubscribeLockedProduct(ctx).ProjectId(projectId).Amount(amount).AutoSubscribe(autoSubscribe).SourceAccount(sourceAccount).RedeemTo(redeemTo).RecvWindow(recvWindow).Execute()

Subscribe Locked Product (TRADE)


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
	projectId := "1" // string | 
	amount := float64(1.0) // float64 | 
	autoSubscribe := false // bool |  (optional)
	sourceAccount := models.SubscribeFlexibleProductSourceAccountParameterSpot // SubscribeFlexibleProductSourceAccountParameter |  (optional)
	redeemTo := models.SetLockedProductRedeemOptionRedeemToParameterSpot // SetLockedProductRedeemOptionRedeemToParameter |  (optional)
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleLockedAPI.SubscribeLockedProduct(context.Background()).ProjectId(projectId).Amount(amount).AutoSubscribe(autoSubscribe).SourceAccount(sourceAccount).RedeemTo(redeemTo).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleLockedAPI.SubscribeLockedProduct``: %v\n", err)
		return
	}

	// response from `SubscribeLockedProduct`: SubscribeLockedProductResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** |  | 
 **amount** | **float64** |  | 
 **autoSubscribe** | **bool** |  | 
 **sourceAccount** | [**SubscribeFlexibleProductSourceAccountParameter**](SubscribeFlexibleProductSourceAccountParameter.md) |  | 
 **redeemTo** | [**SetLockedProductRedeemOptionRedeemToParameter**](SetLockedProductRedeemOptionRedeemToParameter.md) |  | 
 **recvWindow** | **int64** | Request validity window in milliseconds. | 

### Return type

[**SubscribeLockedProductResponse**](SubscribeLockedProductResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

