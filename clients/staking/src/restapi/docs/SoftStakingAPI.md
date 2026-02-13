# \SoftStakingAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**GetSoftStakingProductList**](SoftStakingAPI.md#GetSoftStakingProductList) | **Get** /sapi/v1/soft-staking/list | Get Soft Staking Product List (USER_DATA)
[**GetSoftStakingRewardsHistory**](SoftStakingAPI.md#GetSoftStakingRewardsHistory) | **Get** /sapi/v1/soft-staking/history/rewardsRecord | Get Soft Staking Rewards History(USER_DATA)
[**SetSoftStaking**](SoftStakingAPI.md#SetSoftStaking) | **Get** /sapi/v1/soft-staking/set | Set Soft Staking (USER_DATA)


## GetSoftStakingProductList

> GetSoftStakingProductListResponse GetSoftStakingProductList(ctx).Asset(asset).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Soft Staking Product List (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/staking"
	"github.com/binance/binance-connector-go/common/v2/common"
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

	resp, err := apiClient.RestApi.SoftStakingAPI.GetSoftStakingProductList(context.Background()).Asset(asset).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `SoftStakingAPI.GetSoftStakingProductList``: %v\n", err)
		return
	}

	// response from `GetSoftStakingProductList`: GetSoftStakingProductListResponse
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

[**GetSoftStakingProductListResponse**](GetSoftStakingProductListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetSoftStakingRewardsHistory

> GetSoftStakingRewardsHistoryResponse GetSoftStakingRewardsHistory(ctx).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Soft Staking Rewards History(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/staking"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
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

	resp, err := apiClient.RestApi.SoftStakingAPI.GetSoftStakingRewardsHistory(context.Background()).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `SoftStakingAPI.GetSoftStakingRewardsHistory``: %v\n", err)
		return
	}

	// response from `GetSoftStakingRewardsHistory`: GetSoftStakingRewardsHistoryResponse
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
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10, Max:100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSoftStakingRewardsHistoryResponse**](GetSoftStakingRewardsHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SetSoftStaking

> SetSoftStakingResponse SetSoftStaking(ctx).SoftStaking(softStaking).RecvWindow(recvWindow).Execute()

Set Soft Staking (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/staking"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	softStaking := true // bool | true or false
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStakingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.SoftStakingAPI.SetSoftStaking(context.Background()).SoftStaking(softStaking).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `SoftStakingAPI.SetSoftStaking``: %v\n", err)
		return
	}

	// response from `SetSoftStaking`: SetSoftStakingResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **softStaking** | **bool** | true or false | 
 **recvWindow** | **int64** |  | 

### Return type

[**SetSoftStakingResponse**](SetSoftStakingResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

