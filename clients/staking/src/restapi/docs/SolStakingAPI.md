# \SolStakingAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**ClaimBoostRewards**](SolStakingAPI.md#ClaimBoostRewards) | **Post** /sapi/v1/sol-staking/sol/claim | Claim Boost Rewards(TRADE)
[**GetBnsolRateHistory**](SolStakingAPI.md#GetBnsolRateHistory) | **Get** /sapi/v1/sol-staking/sol/history/rateHistory | Get BNSOL Rate History(USER_DATA)
[**GetBnsolRewardsHistory**](SolStakingAPI.md#GetBnsolRewardsHistory) | **Get** /sapi/v1/sol-staking/sol/history/bnsolRewardsHistory | Get BNSOL rewards history(USER_DATA)
[**GetBoostRewardsHistory**](SolStakingAPI.md#GetBoostRewardsHistory) | **Get** /sapi/v1/sol-staking/sol/history/boostRewardsHistory | Get Boost Rewards History(USER_DATA)
[**GetSolRedemptionHistory**](SolStakingAPI.md#GetSolRedemptionHistory) | **Get** /sapi/v1/sol-staking/sol/history/redemptionHistory | Get SOL redemption history(USER_DATA)
[**GetSolStakingHistory**](SolStakingAPI.md#GetSolStakingHistory) | **Get** /sapi/v1/sol-staking/sol/history/stakingHistory | Get SOL staking history(USER_DATA)
[**GetSolStakingQuotaDetails**](SolStakingAPI.md#GetSolStakingQuotaDetails) | **Get** /sapi/v1/sol-staking/sol/quota | Get SOL staking quota details(USER_DATA)
[**GetUnclaimedRewards**](SolStakingAPI.md#GetUnclaimedRewards) | **Get** /sapi/v1/sol-staking/sol/history/unclaimedRewards | Get Unclaimed Rewards(USER_DATA)
[**RedeemSol**](SolStakingAPI.md#RedeemSol) | **Post** /sapi/v1/sol-staking/sol/redeem | Redeem SOL(TRADE)
[**SolStakingAccount**](SolStakingAPI.md#SolStakingAccount) | **Get** /sapi/v1/sol-staking/account | SOL Staking account(USER_DATA)
[**SubscribeSolStaking**](SolStakingAPI.md#SubscribeSolStaking) | **Post** /sapi/v1/sol-staking/sol/stake | Subscribe SOL Staking(TRADE)


## ClaimBoostRewards

> ClaimBoostRewardsResponse ClaimBoostRewards(ctx).RecvWindow(recvWindow).Execute()

Claim Boost Rewards(TRADE)


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
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStakingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.SolStakingAPI.ClaimBoostRewards(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `SolStakingAPI.ClaimBoostRewards``: %v\n", err)
		return
	}

	// response from `ClaimBoostRewards`: ClaimBoostRewardsResponse
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

[**ClaimBoostRewardsResponse**](ClaimBoostRewardsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetBnsolRateHistory

> GetBnsolRateHistoryResponse GetBnsolRateHistory(ctx).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get BNSOL Rate History(USER_DATA)


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

	resp, err := apiClient.RestApi.SolStakingAPI.GetBnsolRateHistory(context.Background()).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `SolStakingAPI.GetBnsolRateHistory``: %v\n", err)
		return
	}

	// response from `GetBnsolRateHistory`: GetBnsolRateHistoryResponse
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
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10, Max:100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetBnsolRateHistoryResponse**](GetBnsolRateHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetBnsolRewardsHistory

> GetBnsolRewardsHistoryResponse GetBnsolRewardsHistory(ctx).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get BNSOL rewards history(USER_DATA)


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

	resp, err := apiClient.RestApi.SolStakingAPI.GetBnsolRewardsHistory(context.Background()).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `SolStakingAPI.GetBnsolRewardsHistory``: %v\n", err)
		return
	}

	// response from `GetBnsolRewardsHistory`: GetBnsolRewardsHistoryResponse
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
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10, Max:100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetBnsolRewardsHistoryResponse**](GetBnsolRewardsHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetBoostRewardsHistory

> GetBoostRewardsHistoryResponse GetBoostRewardsHistory(ctx).Type(type_).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Boost Rewards History(USER_DATA)


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
	type_ := "CLAIM" // string | \"CLAIM\", \"DISTRIBUTE\", default \"CLAIM\"
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

	resp, err := apiClient.RestApi.SolStakingAPI.GetBoostRewardsHistory(context.Background()).Type(type_).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `SolStakingAPI.GetBoostRewardsHistory``: %v\n", err)
		return
	}

	// response from `GetBoostRewardsHistory`: GetBoostRewardsHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | \&quot;CLAIM\&quot;, \&quot;DISTRIBUTE\&quot;, default \&quot;CLAIM\&quot; | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10, Max:100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetBoostRewardsHistoryResponse**](GetBoostRewardsHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetSolRedemptionHistory

> GetSolRedemptionHistoryResponse GetSolRedemptionHistory(ctx).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get SOL redemption history(USER_DATA)


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

	resp, err := apiClient.RestApi.SolStakingAPI.GetSolRedemptionHistory(context.Background()).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `SolStakingAPI.GetSolRedemptionHistory``: %v\n", err)
		return
	}

	// response from `GetSolRedemptionHistory`: GetSolRedemptionHistoryResponse
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
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10, Max:100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSolRedemptionHistoryResponse**](GetSolRedemptionHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetSolStakingHistory

> GetSolStakingHistoryResponse GetSolStakingHistory(ctx).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get SOL staking history(USER_DATA)


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

	resp, err := apiClient.RestApi.SolStakingAPI.GetSolStakingHistory(context.Background()).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `SolStakingAPI.GetSolStakingHistory``: %v\n", err)
		return
	}

	// response from `GetSolStakingHistory`: GetSolStakingHistoryResponse
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
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10, Max:100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSolStakingHistoryResponse**](GetSolStakingHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetSolStakingQuotaDetails

> GetSolStakingQuotaDetailsResponse GetSolStakingQuotaDetails(ctx).RecvWindow(recvWindow).Execute()

Get SOL staking quota details(USER_DATA)


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
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStakingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.SolStakingAPI.GetSolStakingQuotaDetails(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `SolStakingAPI.GetSolStakingQuotaDetails``: %v\n", err)
		return
	}

	// response from `GetSolStakingQuotaDetails`: GetSolStakingQuotaDetailsResponse
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

[**GetSolStakingQuotaDetailsResponse**](GetSolStakingQuotaDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetUnclaimedRewards

> GetUnclaimedRewardsResponse GetUnclaimedRewards(ctx).RecvWindow(recvWindow).Execute()

Get Unclaimed Rewards(USER_DATA)


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
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStakingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.SolStakingAPI.GetUnclaimedRewards(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `SolStakingAPI.GetUnclaimedRewards``: %v\n", err)
		return
	}

	// response from `GetUnclaimedRewards`: GetUnclaimedRewardsResponse
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

[**GetUnclaimedRewardsResponse**](GetUnclaimedRewardsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## RedeemSol

> RedeemSolResponse RedeemSol(ctx).Amount(amount).RecvWindow(recvWindow).Execute()

Redeem SOL(TRADE)


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
	amount := float32(1.0) // float32 | Amount in SOL.
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStakingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.SolStakingAPI.RedeemSol(context.Background()).Amount(amount).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `SolStakingAPI.RedeemSol``: %v\n", err)
		return
	}

	// response from `RedeemSol`: RedeemSolResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **float32** | Amount in SOL. | 
 **recvWindow** | **int64** |  | 

### Return type

[**RedeemSolResponse**](RedeemSolResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SolStakingAccount

> SolStakingAccountResponse SolStakingAccount(ctx).RecvWindow(recvWindow).Execute()

SOL Staking account(USER_DATA)


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
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStakingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.SolStakingAPI.SolStakingAccount(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `SolStakingAPI.SolStakingAccount``: %v\n", err)
		return
	}

	// response from `SolStakingAccount`: SolStakingAccountResponse
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

[**SolStakingAccountResponse**](SolStakingAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SubscribeSolStaking

> SubscribeSolStakingResponse SubscribeSolStaking(ctx).Amount(amount).RecvWindow(recvWindow).Execute()

Subscribe SOL Staking(TRADE)


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
	amount := float32(1.0) // float32 | Amount in SOL.
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStakingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.SolStakingAPI.SubscribeSolStaking(context.Background()).Amount(amount).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `SolStakingAPI.SubscribeSolStaking``: %v\n", err)
		return
	}

	// response from `SubscribeSolStaking`: SubscribeSolStakingResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **float32** | Amount in SOL. | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubscribeSolStakingResponse**](SubscribeSolStakingResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

