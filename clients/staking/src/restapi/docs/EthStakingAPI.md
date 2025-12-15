# \EthStakingAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**EthStakingAccount**](EthStakingAPI.md#EthStakingAccount) | **Get** /sapi/v2/eth-staking/account | ETH Staking account(USER_DATA)
[**GetCurrentEthStakingQuota**](EthStakingAPI.md#GetCurrentEthStakingQuota) | **Get** /sapi/v1/eth-staking/eth/quota | Get current ETH staking quota(USER_DATA)
[**GetEthRedemptionHistory**](EthStakingAPI.md#GetEthRedemptionHistory) | **Get** /sapi/v1/eth-staking/eth/history/redemptionHistory | Get ETH redemption history(USER_DATA)
[**GetEthStakingHistory**](EthStakingAPI.md#GetEthStakingHistory) | **Get** /sapi/v1/eth-staking/eth/history/stakingHistory | Get ETH staking history(USER_DATA)
[**GetWbethRateHistory**](EthStakingAPI.md#GetWbethRateHistory) | **Get** /sapi/v1/eth-staking/eth/history/rateHistory | Get WBETH Rate History(USER_DATA)
[**GetWbethRewardsHistory**](EthStakingAPI.md#GetWbethRewardsHistory) | **Get** /sapi/v1/eth-staking/eth/history/wbethRewardsHistory | Get WBETH rewards history(USER_DATA)
[**GetWbethUnwrapHistory**](EthStakingAPI.md#GetWbethUnwrapHistory) | **Get** /sapi/v1/eth-staking/wbeth/history/unwrapHistory | Get WBETH unwrap history(USER_DATA)
[**GetWbethWrapHistory**](EthStakingAPI.md#GetWbethWrapHistory) | **Get** /sapi/v1/eth-staking/wbeth/history/wrapHistory | Get WBETH wrap history(USER_DATA)
[**RedeemEth**](EthStakingAPI.md#RedeemEth) | **Post** /sapi/v1/eth-staking/eth/redeem | Redeem ETH(TRADE)
[**SubscribeEthStaking**](EthStakingAPI.md#SubscribeEthStaking) | **Post** /sapi/v2/eth-staking/eth/stake | Subscribe ETH Staking(TRADE)
[**WrapBeth**](EthStakingAPI.md#WrapBeth) | **Post** /sapi/v1/eth-staking/wbeth/wrap | Wrap BETH(TRADE)


## EthStakingAccount

> EthStakingAccountResponse EthStakingAccount(ctx).RecvWindow(recvWindow).Execute()

ETH Staking account(USER_DATA)


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

	resp, err := apiClient.RestApi.EthStakingAPI.EthStakingAccount(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `EthStakingAPI.EthStakingAccount``: %v\n", err)
		return
	}

	// response from `EthStakingAccount`: EthStakingAccountResponse
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

[**EthStakingAccountResponse**](EthStakingAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetCurrentEthStakingQuota

> GetCurrentEthStakingQuotaResponse GetCurrentEthStakingQuota(ctx).RecvWindow(recvWindow).Execute()

Get current ETH staking quota(USER_DATA)


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

	resp, err := apiClient.RestApi.EthStakingAPI.GetCurrentEthStakingQuota(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `EthStakingAPI.GetCurrentEthStakingQuota``: %v\n", err)
		return
	}

	// response from `GetCurrentEthStakingQuota`: GetCurrentEthStakingQuotaResponse
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

[**GetCurrentEthStakingQuotaResponse**](GetCurrentEthStakingQuotaResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetEthRedemptionHistory

> GetEthRedemptionHistoryResponse GetEthRedemptionHistory(ctx).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get ETH redemption history(USER_DATA)


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

	resp, err := apiClient.RestApi.EthStakingAPI.GetEthRedemptionHistory(context.Background()).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `EthStakingAPI.GetEthRedemptionHistory``: %v\n", err)
		return
	}

	// response from `GetEthRedemptionHistory`: GetEthRedemptionHistoryResponse
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

[**GetEthRedemptionHistoryResponse**](GetEthRedemptionHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetEthStakingHistory

> GetEthStakingHistoryResponse GetEthStakingHistory(ctx).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get ETH staking history(USER_DATA)


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

	resp, err := apiClient.RestApi.EthStakingAPI.GetEthStakingHistory(context.Background()).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `EthStakingAPI.GetEthStakingHistory``: %v\n", err)
		return
	}

	// response from `GetEthStakingHistory`: GetEthStakingHistoryResponse
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

[**GetEthStakingHistoryResponse**](GetEthStakingHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetWbethRateHistory

> GetWbethRateHistoryResponse GetWbethRateHistory(ctx).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get WBETH Rate History(USER_DATA)


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

	resp, err := apiClient.RestApi.EthStakingAPI.GetWbethRateHistory(context.Background()).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `EthStakingAPI.GetWbethRateHistory``: %v\n", err)
		return
	}

	// response from `GetWbethRateHistory`: GetWbethRateHistoryResponse
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

[**GetWbethRateHistoryResponse**](GetWbethRateHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetWbethRewardsHistory

> GetWbethRewardsHistoryResponse GetWbethRewardsHistory(ctx).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get WBETH rewards history(USER_DATA)


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

	resp, err := apiClient.RestApi.EthStakingAPI.GetWbethRewardsHistory(context.Background()).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `EthStakingAPI.GetWbethRewardsHistory``: %v\n", err)
		return
	}

	// response from `GetWbethRewardsHistory`: GetWbethRewardsHistoryResponse
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

[**GetWbethRewardsHistoryResponse**](GetWbethRewardsHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetWbethUnwrapHistory

> GetWbethUnwrapHistoryResponse GetWbethUnwrapHistory(ctx).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get WBETH unwrap history(USER_DATA)


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

	resp, err := apiClient.RestApi.EthStakingAPI.GetWbethUnwrapHistory(context.Background()).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `EthStakingAPI.GetWbethUnwrapHistory``: %v\n", err)
		return
	}

	// response from `GetWbethUnwrapHistory`: GetWbethUnwrapHistoryResponse
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

[**GetWbethUnwrapHistoryResponse**](GetWbethUnwrapHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetWbethWrapHistory

> GetWbethWrapHistoryResponse GetWbethWrapHistory(ctx).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get WBETH wrap history(USER_DATA)


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

	resp, err := apiClient.RestApi.EthStakingAPI.GetWbethWrapHistory(context.Background()).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `EthStakingAPI.GetWbethWrapHistory``: %v\n", err)
		return
	}

	// response from `GetWbethWrapHistory`: GetWbethWrapHistoryResponse
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

[**GetWbethWrapHistoryResponse**](GetWbethWrapHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## RedeemEth

> RedeemEthResponse RedeemEth(ctx).Amount(amount).Asset(asset).RecvWindow(recvWindow).Execute()

Redeem ETH(TRADE)


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
	amount := float32(1.0) // float32 | Amount in SOL.
	asset := "BETH" // string | WBETH or BETH, default to BETH (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStakingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.EthStakingAPI.RedeemEth(context.Background()).Amount(amount).Asset(asset).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `EthStakingAPI.RedeemEth``: %v\n", err)
		return
	}

	// response from `RedeemEth`: RedeemEthResponse
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
 **asset** | **string** | WBETH or BETH, default to BETH | 
 **recvWindow** | **int64** |  | 

### Return type

[**RedeemEthResponse**](RedeemEthResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SubscribeEthStaking

> SubscribeEthStakingResponse SubscribeEthStaking(ctx).Amount(amount).RecvWindow(recvWindow).Execute()

Subscribe ETH Staking(TRADE)


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
	amount := float32(1.0) // float32 | Amount in SOL.
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStakingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.EthStakingAPI.SubscribeEthStaking(context.Background()).Amount(amount).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `EthStakingAPI.SubscribeEthStaking``: %v\n", err)
		return
	}

	// response from `SubscribeEthStaking`: SubscribeEthStakingResponse
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

[**SubscribeEthStakingResponse**](SubscribeEthStakingResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## WrapBeth

> WrapBethResponse WrapBeth(ctx).Amount(amount).RecvWindow(recvWindow).Execute()

Wrap BETH(TRADE)


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
	amount := float32(1.0) // float32 | Amount in SOL.
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStakingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.EthStakingAPI.WrapBeth(context.Background()).Amount(amount).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `EthStakingAPI.WrapBeth``: %v\n", err)
		return
	}

	// response from `WrapBeth`: WrapBethResponse
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

[**WrapBethResponse**](WrapBethResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

