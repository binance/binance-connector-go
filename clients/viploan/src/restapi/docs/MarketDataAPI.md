# \MarketDataAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**GetBorrowInterestRate**](MarketDataAPI.md#GetBorrowInterestRate) | **Get** /sapi/v1/loan/vip/request/interestRate | Get Borrow Interest Rate (USER_DATA)
[**GetCollateralAssetData**](MarketDataAPI.md#GetCollateralAssetData) | **Get** /sapi/v1/loan/vip/collateral/data | Get Collateral Asset Data (USER_DATA)
[**GetLoanableAssetsData**](MarketDataAPI.md#GetLoanableAssetsData) | **Get** /sapi/v1/loan/vip/loanable/data | Get Loanable Assets Data (USER_DATA)
[**GetVIPLoanInterestRateHistory**](MarketDataAPI.md#GetVIPLoanInterestRateHistory) | **Get** /sapi/v1/loan/vip/interestRateHistory | Get VIP Loan Interest Rate History (USER_DATA)
[**QueryVIPLoanFixedRateMarket**](MarketDataAPI.md#QueryVIPLoanFixedRateMarket) | **Get** /sapi/v1/loan/vip/fixed/market | Query VIP Loan Fixed Rate Market (USER_DATA)


## GetBorrowInterestRate

> GetBorrowInterestRateResponse GetBorrowInterestRate(ctx).LoanCoin(loanCoin).RecvWindow(recvWindow).Execute()

Get Borrow Interest Rate (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/viploan"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	loanCoin := "BTC" // string | Max 10 assets, Multiple split by \",\"
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceVipLoanClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.GetBorrowInterestRate(context.Background()).LoanCoin(loanCoin).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.GetBorrowInterestRate``: %v\n", err)
		return
	}

	// response from `GetBorrowInterestRate`: GetBorrowInterestRateResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **loanCoin** | **string** | Max 10 assets, Multiple split by \&quot;,\&quot; | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetBorrowInterestRateResponse**](GetBorrowInterestRateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetCollateralAssetData

> GetCollateralAssetDataResponse GetCollateralAssetData(ctx).CollateralCoin(collateralCoin).RecvWindow(recvWindow).Execute()

Get Collateral Asset Data (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/viploan"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	collateralCoin := "BUSD" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceVipLoanClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.GetCollateralAssetData(context.Background()).CollateralCoin(collateralCoin).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.GetCollateralAssetData``: %v\n", err)
		return
	}

	// response from `GetCollateralAssetData`: GetCollateralAssetDataResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **collateralCoin** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetCollateralAssetDataResponse**](GetCollateralAssetDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetLoanableAssetsData

> GetLoanableAssetsDataResponse GetLoanableAssetsData(ctx).LoanCoin(loanCoin).VipLevel(vipLevel).RecvWindow(recvWindow).Execute()

Get Loanable Assets Data (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/viploan"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	loanCoin := "BUSD" // string |  (optional)
	vipLevel := int64(1) // int64 | Defaults to the user's VIP level. (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceVipLoanClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.GetLoanableAssetsData(context.Background()).LoanCoin(loanCoin).VipLevel(vipLevel).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.GetLoanableAssetsData``: %v\n", err)
		return
	}

	// response from `GetLoanableAssetsData`: GetLoanableAssetsDataResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **loanCoin** | **string** |  | 
 **vipLevel** | **int64** | Defaults to the user&#39;s VIP level. | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetLoanableAssetsDataResponse**](GetLoanableAssetsDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetVIPLoanInterestRateHistory

> GetVIPLoanInterestRateHistoryResponse GetVIPLoanInterestRateHistory(ctx).Coin(coin).RecvWindow(recvWindow).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).Execute()

Get VIP Loan Interest Rate History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/viploan"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	coin := "USDT" // string | 
	recvWindow := int64(5000) // int64 | 
	startTime := int64(1623319461670) // int64 | If both startTime and endTime are omitted, the most recent 90 days are returned. (optional)
	endTime := int64(1641782889000) // int64 | Maximum interval between startTime and endTime is 180 days. Time is based on UTC+0. (optional)
	current := int64(1) // int64 | Current page number, starting from 1. (optional)
	limit := int64(10) // int64 | Number of records per page. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceVipLoanClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.GetVIPLoanInterestRateHistory(context.Background()).Coin(coin).RecvWindow(recvWindow).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.GetVIPLoanInterestRateHistory``: %v\n", err)
		return
	}

	// response from `GetVIPLoanInterestRateHistory`: GetVIPLoanInterestRateHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **coin** | **string** |  | 
 **recvWindow** | **int64** |  | 
 **startTime** | **int64** | If both startTime and endTime are omitted, the most recent 90 days are returned. | 
 **endTime** | **int64** | Maximum interval between startTime and endTime is 180 days. Time is based on UTC+0. | 
 **current** | **int64** | Current page number, starting from 1. | 
 **limit** | **int64** | Number of records per page. | 

### Return type

[**GetVIPLoanInterestRateHistoryResponse**](GetVIPLoanInterestRateHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryVIPLoanFixedRateMarket

> QueryVIPLoanFixedRateMarketResponse QueryVIPLoanFixedRateMarket(ctx).LoanCoin(loanCoin).Duration(duration).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Query VIP Loan Fixed Rate Market (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/viploan"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	loanCoin := "USDT" // string | Loan coin
	duration := int64(30) // int64 | Duration in days, minimum 1 (optional)
	current := int64(1) // int64 | Page number, default 1, minimum 1 (optional)
	size := int64(10) // int64 | Page size, default 10, range [1, 100] (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than `60000` (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceVipLoanClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.QueryVIPLoanFixedRateMarket(context.Background()).LoanCoin(loanCoin).Duration(duration).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.QueryVIPLoanFixedRateMarket``: %v\n", err)
		return
	}

	// response from `QueryVIPLoanFixedRateMarket`: QueryVIPLoanFixedRateMarketResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **loanCoin** | **string** | Loan coin | 
 **duration** | **int64** | Duration in days, minimum 1 | 
 **current** | **int64** | Page number, default 1, minimum 1 | 
 **size** | **int64** | Page size, default 10, range [1, 100] | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**QueryVIPLoanFixedRateMarketResponse**](QueryVIPLoanFixedRateMarketResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

