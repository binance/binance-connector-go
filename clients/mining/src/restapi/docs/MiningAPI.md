# \MiningAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AccountList**](MiningAPI.md#AccountList) | **Get** /sapi/v1/mining/statistics/user/list | Account List(USER_DATA)
[**AcquiringAlgorithm**](MiningAPI.md#AcquiringAlgorithm) | **Get** /sapi/v1/mining/pub/algoList | Acquiring Algorithm(MARKET_DATA)
[**AcquiringCoinname**](MiningAPI.md#AcquiringCoinname) | **Get** /sapi/v1/mining/pub/coinList | Acquiring CoinName(MARKET_DATA)
[**CancelHashrateResaleConfiguration**](MiningAPI.md#CancelHashrateResaleConfiguration) | **Post** /sapi/v1/mining/hash-transfer/config/cancel | Cancel hashrate resale configuration(USER_DATA)
[**EarningsList**](MiningAPI.md#EarningsList) | **Get** /sapi/v1/mining/payment/list | Earnings List(USER_DATA)
[**ExtraBonusList**](MiningAPI.md#ExtraBonusList) | **Get** /sapi/v1/mining/payment/other | Extra Bonus List(USER_DATA)
[**HashrateResaleDetail**](MiningAPI.md#HashrateResaleDetail) | **Get** /sapi/v1/mining/hash-transfer/profit/details | Hashrate Resale Detail(USER_DATA)
[**HashrateResaleList**](MiningAPI.md#HashrateResaleList) | **Get** /sapi/v1/mining/hash-transfer/config/details/list | Hashrate Resale List
[**HashrateResaleRequest**](MiningAPI.md#HashrateResaleRequest) | **Post** /sapi/v1/mining/hash-transfer/config | Hashrate Resale Request(USER_DATA)
[**MiningAccountEarning**](MiningAPI.md#MiningAccountEarning) | **Get** /sapi/v1/mining/payment/uid | Mining Account Earning(USER_DATA)
[**RequestForDetailMinerList**](MiningAPI.md#RequestForDetailMinerList) | **Get** /sapi/v1/mining/worker/detail | Request for Detail Miner List(USER_DATA)
[**RequestForMinerList**](MiningAPI.md#RequestForMinerList) | **Get** /sapi/v1/mining/worker/list | Request for Miner List(USER_DATA)
[**StatisticList**](MiningAPI.md#StatisticList) | **Get** /sapi/v1/mining/statistics/user/status | Statistic List(USER_DATA)


## AccountList

> AccountListResponse AccountList(ctx).Algo(algo).UserName(userName).RecvWindow(recvWindow).Execute()

Account List(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/mining"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	algo := "algo_example" // string | Algorithm(sha256) sha256
	userName := "userName_example" // string | Mining account test
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMiningClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MiningAPI.AccountList(context.Background()).Algo(algo).UserName(userName).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MiningAPI.AccountList``: %v\n", err)
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
 **algo** | **string** | Algorithm(sha256) sha256 | 
 **userName** | **string** | Mining account test | 
 **recvWindow** | **int64** |  | 

### Return type

[**AccountListResponse**](AccountListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## AcquiringAlgorithm

> AcquiringAlgorithmResponse AcquiringAlgorithm(ctx).Execute()

Acquiring Algorithm(MARKET_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/mining"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMiningClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MiningAPI.AcquiringAlgorithm(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MiningAPI.AcquiringAlgorithm``: %v\n", err)
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

Acquiring CoinName(MARKET_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/mining"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMiningClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MiningAPI.AcquiringCoinname(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MiningAPI.AcquiringCoinname``: %v\n", err)
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

Cancel hashrate resale configuration(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/mining"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	configId := int64(1) // int64 | Mining ID 168
	userName := "userName_example" // string | Mining account test
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMiningClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MiningAPI.CancelHashrateResaleConfiguration(context.Background()).ConfigId(configId).UserName(userName).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MiningAPI.CancelHashrateResaleConfiguration``: %v\n", err)
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
 **configId** | **int64** | Mining ID 168 | 
 **userName** | **string** | Mining account test | 
 **recvWindow** | **int64** |  | 

### Return type

[**CancelHashrateResaleConfigurationResponse**](CancelHashrateResaleConfigurationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## EarningsList

> EarningsListResponse EarningsList(ctx).Algo(algo).UserName(userName).Coin(coin).StartDate(startDate).EndDate(endDate).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()

Earnings List(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/mining"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	algo := "algo_example" // string | Algorithm(sha256) sha256
	userName := "userName_example" // string | Mining account test
	coin := "coin_example" // string | Coin Name  (optional)
	startDate := int64(789) // int64 | Millisecond timestamp  (optional)
	endDate := int64(789) // int64 | Millisecond timestamp  (optional)
	pageIndex := int64(1) // int64 | Page number, empty default first page, starting from 1  (optional)
	pageSize := int64(789) // int64 | Min 10,Max 200  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMiningClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MiningAPI.EarningsList(context.Background()).Algo(algo).UserName(userName).Coin(coin).StartDate(startDate).EndDate(endDate).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MiningAPI.EarningsList``: %v\n", err)
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
 **algo** | **string** | Algorithm(sha256) sha256 | 
 **userName** | **string** | Mining account test | 
 **coin** | **string** | Coin Name  | 
 **startDate** | **int64** | Millisecond timestamp  | 
 **endDate** | **int64** | Millisecond timestamp  | 
 **pageIndex** | **int64** | Page number, empty default first page, starting from 1  | 
 **pageSize** | **int64** | Min 10,Max 200  | 
 **recvWindow** | **int64** |  | 

### Return type

[**EarningsListResponse**](EarningsListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## ExtraBonusList

> ExtraBonusListResponse ExtraBonusList(ctx).Algo(algo).UserName(userName).Coin(coin).StartDate(startDate).EndDate(endDate).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()

Extra Bonus List(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/mining"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	algo := "algo_example" // string | Algorithm(sha256) sha256
	userName := "userName_example" // string | Mining account test
	coin := "coin_example" // string | Coin Name  (optional)
	startDate := int64(789) // int64 | Millisecond timestamp  (optional)
	endDate := int64(789) // int64 | Millisecond timestamp  (optional)
	pageIndex := int64(1) // int64 | Page number, empty default first page, starting from 1  (optional)
	pageSize := int64(789) // int64 | Min 10,Max 200  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMiningClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MiningAPI.ExtraBonusList(context.Background()).Algo(algo).UserName(userName).Coin(coin).StartDate(startDate).EndDate(endDate).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MiningAPI.ExtraBonusList``: %v\n", err)
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
 **algo** | **string** | Algorithm(sha256) sha256 | 
 **userName** | **string** | Mining account test | 
 **coin** | **string** | Coin Name  | 
 **startDate** | **int64** | Millisecond timestamp  | 
 **endDate** | **int64** | Millisecond timestamp  | 
 **pageIndex** | **int64** | Page number, empty default first page, starting from 1  | 
 **pageSize** | **int64** | Min 10,Max 200  | 
 **recvWindow** | **int64** |  | 

### Return type

[**ExtraBonusListResponse**](ExtraBonusListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## HashrateResaleDetail

> HashrateResaleDetailResponse HashrateResaleDetail(ctx).ConfigId(configId).UserName(userName).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()

Hashrate Resale Detail(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/mining"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	configId := int64(1) // int64 | Mining ID 168
	userName := "userName_example" // string | Mining account test
	pageIndex := int64(1) // int64 | Page number, empty default first page, starting from 1  (optional)
	pageSize := int64(789) // int64 | Min 10,Max 200  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMiningClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MiningAPI.HashrateResaleDetail(context.Background()).ConfigId(configId).UserName(userName).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MiningAPI.HashrateResaleDetail``: %v\n", err)
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
 **configId** | **int64** | Mining ID 168 | 
 **userName** | **string** | Mining account test | 
 **pageIndex** | **int64** | Page number, empty default first page, starting from 1  | 
 **pageSize** | **int64** | Min 10,Max 200  | 
 **recvWindow** | **int64** |  | 

### Return type

[**HashrateResaleDetailResponse**](HashrateResaleDetailResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## HashrateResaleList

> HashrateResaleListResponse HashrateResaleList(ctx).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()

Hashrate Resale List


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/mining"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	pageIndex := int64(1) // int64 | Page number, empty default first page, starting from 1  (optional)
	pageSize := int64(789) // int64 | Min 10,Max 200  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMiningClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MiningAPI.HashrateResaleList(context.Background()).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MiningAPI.HashrateResaleList``: %v\n", err)
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
 **pageIndex** | **int64** | Page number, empty default first page, starting from 1  | 
 **pageSize** | **int64** | Min 10,Max 200  | 
 **recvWindow** | **int64** |  | 

### Return type

[**HashrateResaleListResponse**](HashrateResaleListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## HashrateResaleRequest

> HashrateResaleRequestResponse HashrateResaleRequest(ctx).UserName(userName).Algo(algo).EndDate(endDate).StartDate(startDate).ToPoolUser(toPoolUser).HashRate(hashRate).RecvWindow(recvWindow).Execute()

Hashrate Resale Request(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/mining"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	userName := "userName_example" // string | Mining account test
	algo := "algo_example" // string | Algorithm(sha256) sha256
	endDate := int64(789) // int64 | Resale End Time (Millisecond timestamp) 1617659086000
	startDate := int64(789) // int64 | Resale Start Time(Millisecond timestamp) 1607659086000
	toPoolUser := "toPoolUser_example" // string | Mining Account S19pro
	hashRate := int64(789) // int64 | Resale hashrate h/s must be transferred (BTC is greater than 500000000000 ETH is greater than 500000) 100000000
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMiningClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MiningAPI.HashrateResaleRequest(context.Background()).UserName(userName).Algo(algo).EndDate(endDate).StartDate(startDate).ToPoolUser(toPoolUser).HashRate(hashRate).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MiningAPI.HashrateResaleRequest``: %v\n", err)
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
 **userName** | **string** | Mining account test | 
 **algo** | **string** | Algorithm(sha256) sha256 | 
 **endDate** | **int64** | Resale End Time (Millisecond timestamp) 1617659086000 | 
 **startDate** | **int64** | Resale Start Time(Millisecond timestamp) 1607659086000 | 
 **toPoolUser** | **string** | Mining Account S19pro | 
 **hashRate** | **int64** | Resale hashrate h/s must be transferred (BTC is greater than 500000000000 ETH is greater than 500000) 100000000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**HashrateResaleRequestResponse**](HashrateResaleRequestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MiningAccountEarning

> MiningAccountEarningResponse MiningAccountEarning(ctx).Algo(algo).StartDate(startDate).EndDate(endDate).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()

Mining Account Earning(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/mining"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	algo := "algo_example" // string | Algorithm(sha256) sha256
	startDate := int64(789) // int64 | Millisecond timestamp  (optional)
	endDate := int64(789) // int64 | Millisecond timestamp  (optional)
	pageIndex := int64(1) // int64 | Page number, empty default first page, starting from 1  (optional)
	pageSize := int64(789) // int64 | Min 10,Max 200  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMiningClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MiningAPI.MiningAccountEarning(context.Background()).Algo(algo).StartDate(startDate).EndDate(endDate).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MiningAPI.MiningAccountEarning``: %v\n", err)
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
 **algo** | **string** | Algorithm(sha256) sha256 | 
 **startDate** | **int64** | Millisecond timestamp  | 
 **endDate** | **int64** | Millisecond timestamp  | 
 **pageIndex** | **int64** | Page number, empty default first page, starting from 1  | 
 **pageSize** | **int64** | Min 10,Max 200  | 
 **recvWindow** | **int64** |  | 

### Return type

[**MiningAccountEarningResponse**](MiningAccountEarningResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## RequestForDetailMinerList

> RequestForDetailMinerListResponse RequestForDetailMinerList(ctx).Algo(algo).UserName(userName).WorkerName(workerName).RecvWindow(recvWindow).Execute()

Request for Detail Miner List(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/mining"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	algo := "algo_example" // string | Algorithm(sha256) sha256
	userName := "userName_example" // string | Mining account test
	workerName := "workerName_example" // string | Miner’s name(required) bhdc1.16A10404B
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMiningClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MiningAPI.RequestForDetailMinerList(context.Background()).Algo(algo).UserName(userName).WorkerName(workerName).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MiningAPI.RequestForDetailMinerList``: %v\n", err)
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
 **algo** | **string** | Algorithm(sha256) sha256 | 
 **userName** | **string** | Mining account test | 
 **workerName** | **string** | Miner’s name(required) bhdc1.16A10404B | 
 **recvWindow** | **int64** |  | 

### Return type

[**RequestForDetailMinerListResponse**](RequestForDetailMinerListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## RequestForMinerList

> RequestForMinerListResponse RequestForMinerList(ctx).Algo(algo).UserName(userName).PageIndex(pageIndex).Sort(sort).SortColumn(sortColumn).WorkerStatus(workerStatus).RecvWindow(recvWindow).Execute()

Request for Miner List(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/mining"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	algo := "algo_example" // string | Algorithm(sha256) sha256
	userName := "userName_example" // string | Mining account test
	pageIndex := int64(1) // int64 | Page number, empty default first page, starting from 1  (optional)
	sort := int64(0) // int64 | sort sequence(default=0)0 positive sequence，1 negative sequence (optional)
	sortColumn := int64(1) // int64 | Sort by( default 1): <br></br>1: miner name, <br></br>2: real-time computing power, <br></br>3: daily average computing power, <br></br>4: real-time rejection rate, <br></br>5: last submission time (optional)
	workerStatus := int64(0) // int64 | miners status(default=0),0 all，1 valid，2 invalid，3 failure (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMiningClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MiningAPI.RequestForMinerList(context.Background()).Algo(algo).UserName(userName).PageIndex(pageIndex).Sort(sort).SortColumn(sortColumn).WorkerStatus(workerStatus).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MiningAPI.RequestForMinerList``: %v\n", err)
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
 **algo** | **string** | Algorithm(sha256) sha256 | 
 **userName** | **string** | Mining account test | 
 **pageIndex** | **int64** | Page number, empty default first page, starting from 1  | 
 **sort** | **int64** | sort sequence(default&#x3D;0)0 positive sequence，1 negative sequence | 
 **sortColumn** | **int64** | Sort by( default 1): &lt;br&gt;&lt;/br&gt;1: miner name, &lt;br&gt;&lt;/br&gt;2: real-time computing power, &lt;br&gt;&lt;/br&gt;3: daily average computing power, &lt;br&gt;&lt;/br&gt;4: real-time rejection rate, &lt;br&gt;&lt;/br&gt;5: last submission time | 
 **workerStatus** | **int64** | miners status(default&#x3D;0),0 all，1 valid，2 invalid，3 failure | 
 **recvWindow** | **int64** |  | 

### Return type

[**RequestForMinerListResponse**](RequestForMinerListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## StatisticList

> StatisticListResponse StatisticList(ctx).Algo(algo).UserName(userName).RecvWindow(recvWindow).Execute()

Statistic List(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/mining"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	algo := "algo_example" // string | Algorithm(sha256) sha256
	userName := "userName_example" // string | Mining account test
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMiningClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MiningAPI.StatisticList(context.Background()).Algo(algo).UserName(userName).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MiningAPI.StatisticList``: %v\n", err)
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
 **algo** | **string** | Algorithm(sha256) sha256 | 
 **userName** | **string** | Mining account test | 
 **recvWindow** | **int64** |  | 

### Return type

[**StatisticListResponse**](StatisticListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

