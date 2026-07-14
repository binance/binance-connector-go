# \DefaultAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AccountList**](DefaultAPI.md#AccountList) | **Get** /sapi/v1/mining/statistics/user/list | Account List (USER_DATA)
[**AcquiringAlgorithm**](DefaultAPI.md#AcquiringAlgorithm) | **Get** /sapi/v1/mining/pub/algoList | Acquiring Algorithm (MARKET_DATA)
[**AcquiringCoinname**](DefaultAPI.md#AcquiringCoinname) | **Get** /sapi/v1/mining/pub/coinList | Acquiring CoinName (MARKET_DATA)
[**CancelHashrateResaleConfiguration**](DefaultAPI.md#CancelHashrateResaleConfiguration) | **Post** /sapi/v1/mining/hash-transfer/config/cancel | Cancel hashrate resale configuration (USER_DATA)
[**EarningsList**](DefaultAPI.md#EarningsList) | **Get** /sapi/v1/mining/payment/list | Earnings List (USER_DATA)
[**ExtraBonusList**](DefaultAPI.md#ExtraBonusList) | **Get** /sapi/v1/mining/payment/other | Extra Bonus List (USER_DATA)
[**HashrateResaleDetail**](DefaultAPI.md#HashrateResaleDetail) | **Get** /sapi/v1/mining/hash-transfer/profit/details | Hashrate Resale Detail (USER_DATA)
[**HashrateResaleList**](DefaultAPI.md#HashrateResaleList) | **Get** /sapi/v1/mining/hash-transfer/config/details/list | Hashrate Resale List (USER_DATA)
[**HashrateResaleRequest**](DefaultAPI.md#HashrateResaleRequest) | **Post** /sapi/v1/mining/hash-transfer/config | Hashrate Resale Request (USER_DATA)
[**MiningAccountEarning**](DefaultAPI.md#MiningAccountEarning) | **Get** /sapi/v1/mining/payment/uid | Mining Account Earning (USER_DATA)
[**RequestForDetailMinerList**](DefaultAPI.md#RequestForDetailMinerList) | **Get** /sapi/v1/mining/worker/detail | Request for Detail Miner List (USER_DATA)
[**RequestForMinerList**](DefaultAPI.md#RequestForMinerList) | **Get** /sapi/v1/mining/worker/list | Request for Miner List (USER_DATA)
[**StatisticList**](DefaultAPI.md#StatisticList) | **Get** /sapi/v1/mining/statistics/user/status | Statistic List (USER_DATA)


## AccountList

> AccountListResponse AccountList(ctx).Algo(algo).UserName(userName).RecvWindow(recvWindow).Execute()

Account List (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/mining"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	algo := "sha256" // string | Algorithm name.
	userName := "test" // string | Mining account
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMiningClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.DefaultAPI.AccountList(context.Background()).Algo(algo).UserName(userName).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.AccountList``: %v\n", err)
		return
	}

	// response from `AccountList`: AccountListResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **algo** | **string** | Algorithm name. | 
 **userName** | **string** | Mining account | 
 **recvWindow** | **int64** | Request validity window in milliseconds. | 

### Return type

[**AccountListResponse**](AccountListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## AcquiringAlgorithm

> AcquiringAlgorithmResponse AcquiringAlgorithm(ctx).Execute()

Acquiring Algorithm (MARKET_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/mining"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMiningClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.DefaultAPI.AcquiringAlgorithm(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.AcquiringAlgorithm``: %v\n", err)
		return
	}

	// response from `AcquiringAlgorithm`: AcquiringAlgorithmResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

This endpoint does not need any parameter.

### Return type

[**AcquiringAlgorithmResponse**](AcquiringAlgorithmResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## AcquiringCoinname

> AcquiringCoinnameResponse AcquiringCoinname(ctx).Execute()

Acquiring CoinName (MARKET_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/mining"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMiningClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.DefaultAPI.AcquiringCoinname(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.AcquiringCoinname``: %v\n", err)
		return
	}

	// response from `AcquiringCoinname`: AcquiringCoinnameResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

This endpoint does not need any parameter.

### Return type

[**AcquiringCoinnameResponse**](AcquiringCoinnameResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CancelHashrateResaleConfiguration

> CancelHashrateResaleConfigurationResponse CancelHashrateResaleConfiguration(ctx).ConfigId(configId).UserName(userName).RecvWindow(recvWindow).Execute()

Cancel hashrate resale configuration (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/mining"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	configId := int64(168) // int64 | Mining ID
	userName := "test" // string | Mining Account
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMiningClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.DefaultAPI.CancelHashrateResaleConfiguration(context.Background()).ConfigId(configId).UserName(userName).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.CancelHashrateResaleConfiguration``: %v\n", err)
		return
	}

	// response from `CancelHashrateResaleConfiguration`: CancelHashrateResaleConfigurationResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **configId** | **int64** | Mining ID | 
 **userName** | **string** | Mining Account | 
 **recvWindow** | **int64** | Request validity window in milliseconds. | 

### Return type

[**CancelHashrateResaleConfigurationResponse**](CancelHashrateResaleConfigurationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## EarningsList

> EarningsListResponse EarningsList(ctx).Algo(algo).UserName(userName).Coin(coin).StartDate(startDate).EndDate(endDate).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()

Earnings List (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/mining"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	algo := "sha256" // string | Algorithm name.
	userName := "test" // string | Mining account.
	coin := "BTC" // string | Coin name (optional)
	startDate := int64(1770736694138) // int64 | Search start time in milliseconds. (optional)
	endDate := int64(1770736694138) // int64 | Search end time in milliseconds. (optional)
	pageIndex := int64(1) // int64 | Page number, starting from 1. (optional)
	pageSize := int64(10) // int64 | Number of rows per page. (optional)
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMiningClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.DefaultAPI.EarningsList(context.Background()).Algo(algo).UserName(userName).Coin(coin).StartDate(startDate).EndDate(endDate).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.EarningsList``: %v\n", err)
		return
	}

	// response from `EarningsList`: EarningsListResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **algo** | **string** | Algorithm name. | 
 **userName** | **string** | Mining account. | 
 **coin** | **string** | Coin name | 
 **startDate** | **int64** | Search start time in milliseconds. | 
 **endDate** | **int64** | Search end time in milliseconds. | 
 **pageIndex** | **int64** | Page number, starting from 1. | 
 **pageSize** | **int64** | Number of rows per page. | 
 **recvWindow** | **int64** | Request validity window in milliseconds. | 

### Return type

[**EarningsListResponse**](EarningsListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## ExtraBonusList

> ExtraBonusListResponse ExtraBonusList(ctx).Algo(algo).UserName(userName).Coin(coin).StartDate(startDate).EndDate(endDate).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()

Extra Bonus List (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/mining"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	algo := "sha256" // string | Transfer algorithm
	userName := "test" // string | Mining account
	coin := "BTC" // string | Coin name (optional)
	startDate := int64(1770736694138) // int64 | Search start time in milliseconds. (optional)
	endDate := int64(1770736694138) // int64 | Search end time in milliseconds. (optional)
	pageIndex := int64(1) // int64 | Page number, starting from 1. (optional)
	pageSize := int64(10) // int64 | Number of rows per page. (optional)
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMiningClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.DefaultAPI.ExtraBonusList(context.Background()).Algo(algo).UserName(userName).Coin(coin).StartDate(startDate).EndDate(endDate).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.ExtraBonusList``: %v\n", err)
		return
	}

	// response from `ExtraBonusList`: ExtraBonusListResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **algo** | **string** | Transfer algorithm | 
 **userName** | **string** | Mining account | 
 **coin** | **string** | Coin name | 
 **startDate** | **int64** | Search start time in milliseconds. | 
 **endDate** | **int64** | Search end time in milliseconds. | 
 **pageIndex** | **int64** | Page number, starting from 1. | 
 **pageSize** | **int64** | Number of rows per page. | 
 **recvWindow** | **int64** | Request validity window in milliseconds. | 

### Return type

[**ExtraBonusListResponse**](ExtraBonusListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## HashrateResaleDetail

> HashrateResaleDetailResponse HashrateResaleDetail(ctx).ConfigId(configId).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()

Hashrate Resale Detail (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/mining"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	configId := int64(168) // int64 | Configuration ID.
	pageIndex := int64(1) // int64 | Page number, starting from 1. (optional)
	pageSize := int64(10) // int64 | Number of rows per page. (optional)
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMiningClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.DefaultAPI.HashrateResaleDetail(context.Background()).ConfigId(configId).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.HashrateResaleDetail``: %v\n", err)
		return
	}

	// response from `HashrateResaleDetail`: HashrateResaleDetailResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **configId** | **int64** | Configuration ID. | 
 **pageIndex** | **int64** | Page number, starting from 1. | 
 **pageSize** | **int64** | Number of rows per page. | 
 **recvWindow** | **int64** | Request validity window in milliseconds. | 

### Return type

[**HashrateResaleDetailResponse**](HashrateResaleDetailResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## HashrateResaleList

> HashrateResaleListResponse HashrateResaleList(ctx).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()

Hashrate Resale List (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/mining"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	pageIndex := int64(1) // int64 | Page number, starting from 1. (optional)
	pageSize := int64(10) // int64 | Number of rows per page. (optional)
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMiningClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.DefaultAPI.HashrateResaleList(context.Background()).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.HashrateResaleList``: %v\n", err)
		return
	}

	// response from `HashrateResaleList`: HashrateResaleListResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **pageIndex** | **int64** | Page number, starting from 1. | 
 **pageSize** | **int64** | Number of rows per page. | 
 **recvWindow** | **int64** | Request validity window in milliseconds. | 

### Return type

[**HashrateResaleListResponse**](HashrateResaleListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## HashrateResaleRequest

> HashrateResaleRequestResponse HashrateResaleRequest(ctx).UserName(userName).Algo(algo).EndDate(endDate).StartDate(startDate).ToPoolUser(toPoolUser).HashRate(hashRate).RecvWindow(recvWindow).Execute()

Hashrate Resale Request (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/mining"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	userName := "test" // string | Mining Account
	algo := "sha256" // string | Transfer algorithm
	endDate := int64(1770736694138) // int64 | Resale End Time (Millisecond timestamp)
	startDate := int64(1770736694138) // int64 | Resale Start Time(Millisecond timestamp)
	toPoolUser := "S19pro" // string | Mining Account
	hashRate := int64(100000000) // int64 | Resale hashrate h/s must be transferred (BTC is greater than 500000000000 ETH is greater than 500000)
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMiningClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.DefaultAPI.HashrateResaleRequest(context.Background()).UserName(userName).Algo(algo).EndDate(endDate).StartDate(startDate).ToPoolUser(toPoolUser).HashRate(hashRate).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.HashrateResaleRequest``: %v\n", err)
		return
	}

	// response from `HashrateResaleRequest`: HashrateResaleRequestResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **userName** | **string** | Mining Account | 
 **algo** | **string** | Transfer algorithm | 
 **endDate** | **int64** | Resale End Time (Millisecond timestamp) | 
 **startDate** | **int64** | Resale Start Time(Millisecond timestamp) | 
 **toPoolUser** | **string** | Mining Account | 
 **hashRate** | **int64** | Resale hashrate h/s must be transferred (BTC is greater than 500000000000 ETH is greater than 500000) | 
 **recvWindow** | **int64** | Request validity window in milliseconds. | 

### Return type

[**HashrateResaleRequestResponse**](HashrateResaleRequestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MiningAccountEarning

> MiningAccountEarningResponse MiningAccountEarning(ctx).Algo(algo).StartDate(startDate).EndDate(endDate).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()

Mining Account Earning (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/mining"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	algo := "sha256" // string | Algorithm
	startDate := int64(1770736694138) // int64 | Millisecond timestamp (optional)
	endDate := int64(1770736694138) // int64 | Millisecond timestamp (optional)
	pageIndex := int64(1) // int64 | Page number, starting from 1. (optional)
	pageSize := int64(10) // int64 | Number of rows per page. (optional)
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMiningClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.DefaultAPI.MiningAccountEarning(context.Background()).Algo(algo).StartDate(startDate).EndDate(endDate).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.MiningAccountEarning``: %v\n", err)
		return
	}

	// response from `MiningAccountEarning`: MiningAccountEarningResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **algo** | **string** | Algorithm | 
 **startDate** | **int64** | Millisecond timestamp | 
 **endDate** | **int64** | Millisecond timestamp | 
 **pageIndex** | **int64** | Page number, starting from 1. | 
 **pageSize** | **int64** | Number of rows per page. | 
 **recvWindow** | **int64** | Request validity window in milliseconds. | 

### Return type

[**MiningAccountEarningResponse**](MiningAccountEarningResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## RequestForDetailMinerList

> RequestForDetailMinerListResponse RequestForDetailMinerList(ctx).Algo(algo).UserName(userName).WorkerName(workerName).RecvWindow(recvWindow).Execute()

Request for Detail Miner List (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/mining"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	algo := "sha256" // string | Algorithm
	userName := "test" // string | Mining account
	workerName := "bhdc1.16A10404B" // string | Miner name.
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMiningClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.DefaultAPI.RequestForDetailMinerList(context.Background()).Algo(algo).UserName(userName).WorkerName(workerName).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.RequestForDetailMinerList``: %v\n", err)
		return
	}

	// response from `RequestForDetailMinerList`: RequestForDetailMinerListResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **algo** | **string** | Algorithm | 
 **userName** | **string** | Mining account | 
 **workerName** | **string** | Miner name. | 
 **recvWindow** | **int64** | Request validity window in milliseconds. | 

### Return type

[**RequestForDetailMinerListResponse**](RequestForDetailMinerListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## RequestForMinerList

> RequestForMinerListResponse RequestForMinerList(ctx).Algo(algo).UserName(userName).PageIndex(pageIndex).Sort(sort).SortColumn(sortColumn).WorkerStatus(workerStatus).RecvWindow(recvWindow).Execute()

Request for Miner List (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/mining"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	algo := "sha256" // string | Algorithm
	userName := "test" // string | Mining account
	pageIndex := int64(1) // int64 | Page number, starting from 1. (optional)
	sort := int64(0) // int64 | Sort order. 0 for ascending, 1 for descending. (optional)
	sortColumn := int64(1) // int64 | Sort by: 1 miner name, 2 real-time hashrate, 3 daily average hashrate, 4 real-time rejection rate, 5 last submission time (optional)
	workerStatus := int64(0) // int64 | Miner status. 0 all, 1 valid, 2 invalid, 3 failure. (optional)
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMiningClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.DefaultAPI.RequestForMinerList(context.Background()).Algo(algo).UserName(userName).PageIndex(pageIndex).Sort(sort).SortColumn(sortColumn).WorkerStatus(workerStatus).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.RequestForMinerList``: %v\n", err)
		return
	}

	// response from `RequestForMinerList`: RequestForMinerListResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **algo** | **string** | Algorithm | 
 **userName** | **string** | Mining account | 
 **pageIndex** | **int64** | Page number, starting from 1. | 
 **sort** | **int64** | Sort order. 0 for ascending, 1 for descending. | 
 **sortColumn** | **int64** | Sort by: 1 miner name, 2 real-time hashrate, 3 daily average hashrate, 4 real-time rejection rate, 5 last submission time | 
 **workerStatus** | **int64** | Miner status. 0 all, 1 valid, 2 invalid, 3 failure. | 
 **recvWindow** | **int64** | Request validity window in milliseconds. | 

### Return type

[**RequestForMinerListResponse**](RequestForMinerListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## StatisticList

> StatisticListResponse StatisticList(ctx).Algo(algo).UserName(userName).RecvWindow(recvWindow).Execute()

Statistic List (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/mining"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	algo := "sha256" // string | Algorithm
	userName := "test" // string | Mining account
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMiningClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.DefaultAPI.StatisticList(context.Background()).Algo(algo).UserName(userName).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.StatisticList``: %v\n", err)
		return
	}

	// response from `StatisticList`: StatisticListResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **algo** | **string** | Algorithm | 
 **userName** | **string** | Mining account | 
 **recvWindow** | **int64** | Request validity window in milliseconds. | 

### Return type

[**StatisticListResponse**](StatisticListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

