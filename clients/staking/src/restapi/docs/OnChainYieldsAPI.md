# \OnChainYieldsAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**GetOnChainYieldsLockedPersonalLeftQuota**](OnChainYieldsAPI.md#GetOnChainYieldsLockedPersonalLeftQuota) | **Get** /sapi/v1/onchain-yields/locked/personalLeftQuota | Get On-chain Yields Locked Personal Left Quota (USER_DATA)
[**GetOnChainYieldsLockedProductList**](OnChainYieldsAPI.md#GetOnChainYieldsLockedProductList) | **Get** /sapi/v1/onchain-yields/locked/list | Get On-chain Yields Locked Product List (USER_DATA)
[**GetOnChainYieldsLockedProductPosition**](OnChainYieldsAPI.md#GetOnChainYieldsLockedProductPosition) | **Get** /sapi/v1/onchain-yields/locked/position | Get On-chain Yields Locked Product Position (USER_DATA)
[**GetOnChainYieldsLockedRedemptionRecord**](OnChainYieldsAPI.md#GetOnChainYieldsLockedRedemptionRecord) | **Get** /sapi/v1/onchain-yields/locked/history/redemptionRecord | Get On-chain Yields Locked Redemption Record (USER_DATA)
[**GetOnChainYieldsLockedRewardsHistory**](OnChainYieldsAPI.md#GetOnChainYieldsLockedRewardsHistory) | **Get** /sapi/v1/onchain-yields/locked/history/rewardsRecord | Get On-chain Yields Locked Rewards History (USER_DATA)
[**GetOnChainYieldsLockedSubscriptionPreview**](OnChainYieldsAPI.md#GetOnChainYieldsLockedSubscriptionPreview) | **Get** /sapi/v1/onchain-yields/locked/subscriptionPreview | Get On-chain Yields Locked Subscription Preview (USER_DATA)
[**GetOnChainYieldsLockedSubscriptionRecord**](OnChainYieldsAPI.md#GetOnChainYieldsLockedSubscriptionRecord) | **Get** /sapi/v1/onchain-yields/locked/history/subscriptionRecord | Get On-chain Yields Locked Subscription Record (USER_DATA)
[**OnChainYieldsAccount**](OnChainYieldsAPI.md#OnChainYieldsAccount) | **Get** /sapi/v1/onchain-yields/account | On-chain Yields Account (USER_DATA)
[**RedeemOnChainYieldsLockedProduct**](OnChainYieldsAPI.md#RedeemOnChainYieldsLockedProduct) | **Post** /sapi/v1/onchain-yields/locked/redeem | Redeem On-chain Yields Locked Product (TRADE)
[**SetOnChainYieldsLockedAutoSubscribe**](OnChainYieldsAPI.md#SetOnChainYieldsLockedAutoSubscribe) | **Post** /sapi/v1/onchain-yields/locked/setAutoSubscribe | Set On-chain Yields Locked Auto Subscribe(USER_DATA)
[**SetOnChainYieldsLockedProductRedeemOption**](OnChainYieldsAPI.md#SetOnChainYieldsLockedProductRedeemOption) | **Post** /sapi/v1/onchain-yields/locked/setRedeemOption | Set On-chain Yields Locked Product Redeem Option(USER_DATA)
[**SubscribeOnChainYieldsLockedProduct**](OnChainYieldsAPI.md#SubscribeOnChainYieldsLockedProduct) | **Post** /sapi/v1/onchain-yields/locked/subscribe | Subscribe On-chain Yields Locked Product(TRADE)


## GetOnChainYieldsLockedPersonalLeftQuota

> GetOnChainYieldsLockedPersonalLeftQuotaResponse GetOnChainYieldsLockedPersonalLeftQuota(ctx).ProjectId(projectId).RecvWindow(recvWindow).Execute()

Get On-chain Yields Locked Personal Left Quota (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/staking"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	projectId := "1" // string | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStakingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.OnChainYieldsAPI.GetOnChainYieldsLockedPersonalLeftQuota(context.Background()).ProjectId(projectId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `OnChainYieldsAPI.GetOnChainYieldsLockedPersonalLeftQuota``: %v\n", err)
		return
	}

	// response from `GetOnChainYieldsLockedPersonalLeftQuota`: GetOnChainYieldsLockedPersonalLeftQuotaResponse
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
 **recvWindow** | **int64** |  | 

### Return type

[**GetOnChainYieldsLockedPersonalLeftQuotaResponse**](GetOnChainYieldsLockedPersonalLeftQuotaResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetOnChainYieldsLockedProductList

> GetOnChainYieldsLockedProductListResponse GetOnChainYieldsLockedProductList(ctx).Asset(asset).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get On-chain Yields Locked Product List (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/staking"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	asset := "BETH" // string | WBETH or BETH, default to BETH (optional)
	current := int64(1) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(10) // int64 | Default:10, Max:100 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStakingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.OnChainYieldsAPI.GetOnChainYieldsLockedProductList(context.Background()).Asset(asset).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `OnChainYieldsAPI.GetOnChainYieldsLockedProductList``: %v\n", err)
		return
	}

	// response from `GetOnChainYieldsLockedProductList`: GetOnChainYieldsLockedProductListResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** | WBETH or BETH, default to BETH | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10, Max:100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetOnChainYieldsLockedProductListResponse**](GetOnChainYieldsLockedProductListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetOnChainYieldsLockedProductPosition

> GetOnChainYieldsLockedProductPositionResponse GetOnChainYieldsLockedProductPosition(ctx).Asset(asset).PositionId(positionId).ProjectId(projectId).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get On-chain Yields Locked Product Position (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/staking"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	asset := "BETH" // string | WBETH or BETH, default to BETH (optional)
	positionId := int64(1) // int64 |  (optional)
	projectId := "1" // string |  (optional)
	current := int64(1) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(10) // int64 | Default:10, Max:100 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStakingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.OnChainYieldsAPI.GetOnChainYieldsLockedProductPosition(context.Background()).Asset(asset).PositionId(positionId).ProjectId(projectId).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `OnChainYieldsAPI.GetOnChainYieldsLockedProductPosition``: %v\n", err)
		return
	}

	// response from `GetOnChainYieldsLockedProductPosition`: GetOnChainYieldsLockedProductPositionResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** | WBETH or BETH, default to BETH | 
 **positionId** | **int64** |  | 
 **projectId** | **string** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10, Max:100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetOnChainYieldsLockedProductPositionResponse**](GetOnChainYieldsLockedProductPositionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetOnChainYieldsLockedRedemptionRecord

> GetOnChainYieldsLockedRedemptionRecordResponse GetOnChainYieldsLockedRedemptionRecord(ctx).PositionId(positionId).RedeemId(redeemId).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get On-chain Yields Locked Redemption Record (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/staking"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	positionId := int64(1) // int64 |  (optional)
	redeemId := "1" // string |  (optional)
	asset := "BETH" // string | WBETH or BETH, default to BETH (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	current := int64(1) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(10) // int64 | Default:10, Max:100 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStakingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.OnChainYieldsAPI.GetOnChainYieldsLockedRedemptionRecord(context.Background()).PositionId(positionId).RedeemId(redeemId).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `OnChainYieldsAPI.GetOnChainYieldsLockedRedemptionRecord``: %v\n", err)
		return
	}

	// response from `GetOnChainYieldsLockedRedemptionRecord`: GetOnChainYieldsLockedRedemptionRecordResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **positionId** | **int64** |  | 
 **redeemId** | **string** |  | 
 **asset** | **string** | WBETH or BETH, default to BETH | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10, Max:100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetOnChainYieldsLockedRedemptionRecordResponse**](GetOnChainYieldsLockedRedemptionRecordResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetOnChainYieldsLockedRewardsHistory

> GetOnChainYieldsLockedRewardsHistoryResponse GetOnChainYieldsLockedRewardsHistory(ctx).PositionId(positionId).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get On-chain Yields Locked Rewards History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/staking"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	positionId := "1" // string |  (optional)
	asset := "BETH" // string | WBETH or BETH, default to BETH (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	current := int64(1) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(10) // int64 | Default:10, Max:100 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStakingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.OnChainYieldsAPI.GetOnChainYieldsLockedRewardsHistory(context.Background()).PositionId(positionId).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `OnChainYieldsAPI.GetOnChainYieldsLockedRewardsHistory``: %v\n", err)
		return
	}

	// response from `GetOnChainYieldsLockedRewardsHistory`: GetOnChainYieldsLockedRewardsHistoryResponse
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
 **asset** | **string** | WBETH or BETH, default to BETH | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10, Max:100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetOnChainYieldsLockedRewardsHistoryResponse**](GetOnChainYieldsLockedRewardsHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetOnChainYieldsLockedSubscriptionPreview

> GetOnChainYieldsLockedSubscriptionPreviewResponse GetOnChainYieldsLockedSubscriptionPreview(ctx).ProjectId(projectId).Amount(amount).AutoSubscribe(autoSubscribe).RecvWindow(recvWindow).Execute()

Get On-chain Yields Locked Subscription Preview (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/staking"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	projectId := "1" // string | 
	amount := float32(1.0) // float32 | Amount in SOL.
	autoSubscribe := true // bool | true or false, default true. (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStakingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.OnChainYieldsAPI.GetOnChainYieldsLockedSubscriptionPreview(context.Background()).ProjectId(projectId).Amount(amount).AutoSubscribe(autoSubscribe).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `OnChainYieldsAPI.GetOnChainYieldsLockedSubscriptionPreview``: %v\n", err)
		return
	}

	// response from `GetOnChainYieldsLockedSubscriptionPreview`: GetOnChainYieldsLockedSubscriptionPreviewResponse
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
 **amount** | **float32** | Amount in SOL. | 
 **autoSubscribe** | **bool** | true or false, default true. | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetOnChainYieldsLockedSubscriptionPreviewResponse**](GetOnChainYieldsLockedSubscriptionPreviewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetOnChainYieldsLockedSubscriptionRecord

> GetOnChainYieldsLockedSubscriptionRecordResponse GetOnChainYieldsLockedSubscriptionRecord(ctx).PurchaseId(purchaseId).ClientId(clientId).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get On-chain Yields Locked Subscription Record (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/staking"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	purchaseId := "1" // string |  (optional)
	clientId := "1" // string |  (optional)
	asset := "BETH" // string | WBETH or BETH, default to BETH (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	current := int64(1) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(10) // int64 | Default:10, Max:100 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStakingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.OnChainYieldsAPI.GetOnChainYieldsLockedSubscriptionRecord(context.Background()).PurchaseId(purchaseId).ClientId(clientId).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `OnChainYieldsAPI.GetOnChainYieldsLockedSubscriptionRecord``: %v\n", err)
		return
	}

	// response from `GetOnChainYieldsLockedSubscriptionRecord`: GetOnChainYieldsLockedSubscriptionRecordResponse
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
 **clientId** | **string** |  | 
 **asset** | **string** | WBETH or BETH, default to BETH | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10, Max:100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetOnChainYieldsLockedSubscriptionRecordResponse**](GetOnChainYieldsLockedSubscriptionRecordResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## OnChainYieldsAccount

> OnChainYieldsAccountResponse OnChainYieldsAccount(ctx).RecvWindow(recvWindow).Execute()

On-chain Yields Account (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/staking"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStakingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.OnChainYieldsAPI.OnChainYieldsAccount(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `OnChainYieldsAPI.OnChainYieldsAccount``: %v\n", err)
		return
	}

	// response from `OnChainYieldsAccount`: OnChainYieldsAccountResponse
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

[**OnChainYieldsAccountResponse**](OnChainYieldsAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## RedeemOnChainYieldsLockedProduct

> RedeemOnChainYieldsLockedProductResponse RedeemOnChainYieldsLockedProduct(ctx).PositionId(positionId).ChannelId(channelId).RecvWindow(recvWindow).Execute()

Redeem On-chain Yields Locked Product (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/staking"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	positionId := "1" // string | 
	channelId := "1" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStakingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.OnChainYieldsAPI.RedeemOnChainYieldsLockedProduct(context.Background()).PositionId(positionId).ChannelId(channelId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `OnChainYieldsAPI.RedeemOnChainYieldsLockedProduct``: %v\n", err)
		return
	}

	// response from `RedeemOnChainYieldsLockedProduct`: RedeemOnChainYieldsLockedProductResponse
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
 **channelId** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**RedeemOnChainYieldsLockedProductResponse**](RedeemOnChainYieldsLockedProductResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SetOnChainYieldsLockedAutoSubscribe

> SetOnChainYieldsLockedAutoSubscribeResponse SetOnChainYieldsLockedAutoSubscribe(ctx).PositionId(positionId).AutoSubscribe(autoSubscribe).RecvWindow(recvWindow).Execute()

Set On-chain Yields Locked Auto Subscribe(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/staking"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	positionId := "1" // string | 
	autoSubscribe := true // bool | true or false
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStakingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.OnChainYieldsAPI.SetOnChainYieldsLockedAutoSubscribe(context.Background()).PositionId(positionId).AutoSubscribe(autoSubscribe).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `OnChainYieldsAPI.SetOnChainYieldsLockedAutoSubscribe``: %v\n", err)
		return
	}

	// response from `SetOnChainYieldsLockedAutoSubscribe`: SetOnChainYieldsLockedAutoSubscribeResponse
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
 **autoSubscribe** | **bool** | true or false | 
 **recvWindow** | **int64** |  | 

### Return type

[**SetOnChainYieldsLockedAutoSubscribeResponse**](SetOnChainYieldsLockedAutoSubscribeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SetOnChainYieldsLockedProductRedeemOption

> SetOnChainYieldsLockedProductRedeemOptionResponse SetOnChainYieldsLockedProductRedeemOption(ctx).PositionId(positionId).RedeemTo(redeemTo).RecvWindow(recvWindow).Execute()

Set On-chain Yields Locked Product Redeem Option(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/staking"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	positionId := "1" // string | 
	redeemTo := "redeemTo_example" // string | 'SPOT','FLEXIBLE'
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStakingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.OnChainYieldsAPI.SetOnChainYieldsLockedProductRedeemOption(context.Background()).PositionId(positionId).RedeemTo(redeemTo).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `OnChainYieldsAPI.SetOnChainYieldsLockedProductRedeemOption``: %v\n", err)
		return
	}

	// response from `SetOnChainYieldsLockedProductRedeemOption`: SetOnChainYieldsLockedProductRedeemOptionResponse
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
 **redeemTo** | **string** | &#39;SPOT&#39;,&#39;FLEXIBLE&#39; | 
 **recvWindow** | **int64** |  | 

### Return type

[**SetOnChainYieldsLockedProductRedeemOptionResponse**](SetOnChainYieldsLockedProductRedeemOptionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SubscribeOnChainYieldsLockedProduct

> SubscribeOnChainYieldsLockedProductResponse SubscribeOnChainYieldsLockedProduct(ctx).ProjectId(projectId).Amount(amount).AutoSubscribe(autoSubscribe).SourceAccount(sourceAccount).RedeemTo(redeemTo).ChannelId(channelId).ClientId(clientId).RecvWindow(recvWindow).Execute()

Subscribe On-chain Yields Locked Product(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/staking"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	projectId := "1" // string | 
	amount := float32(1.0) // float32 | Amount in SOL.
	autoSubscribe := true // bool | true or false, default true. (optional)
	sourceAccount := "SPOT" // string | `SPOT`,`FUND`,`ALL`, default `SPOT` (optional)
	redeemTo := "redeemTo_example" // string | `SPOT`,`FLEXIBLE`, default `FLEXIBLE` Takes effect when Auto Subscribe is false (optional)
	channelId := "1" // string |  (optional)
	clientId := "1" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStakingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.OnChainYieldsAPI.SubscribeOnChainYieldsLockedProduct(context.Background()).ProjectId(projectId).Amount(amount).AutoSubscribe(autoSubscribe).SourceAccount(sourceAccount).RedeemTo(redeemTo).ChannelId(channelId).ClientId(clientId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `OnChainYieldsAPI.SubscribeOnChainYieldsLockedProduct``: %v\n", err)
		return
	}

	// response from `SubscribeOnChainYieldsLockedProduct`: SubscribeOnChainYieldsLockedProductResponse
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
 **amount** | **float32** | Amount in SOL. | 
 **autoSubscribe** | **bool** | true or false, default true. | 
 **sourceAccount** | **string** | &#x60;SPOT&#x60;,&#x60;FUND&#x60;,&#x60;ALL&#x60;, default &#x60;SPOT&#x60; | 
 **redeemTo** | **string** | &#x60;SPOT&#x60;,&#x60;FLEXIBLE&#x60;, default &#x60;FLEXIBLE&#x60; Takes effect when Auto Subscribe is false | 
 **channelId** | **string** |  | 
 **clientId** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubscribeOnChainYieldsLockedProductResponse**](SubscribeOnChainYieldsLockedProductResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

