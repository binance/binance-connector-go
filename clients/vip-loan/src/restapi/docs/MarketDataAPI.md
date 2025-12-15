# \MarketDataAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**GetBorrowInterestRate**](MarketDataAPI.md#GetBorrowInterestRate) | **Get** /sapi/v1/loan/vip/request/interestRate | Get Borrow Interest Rate(USER_DATA)
[**GetCollateralAssetData**](MarketDataAPI.md#GetCollateralAssetData) | **Get** /sapi/v1/loan/vip/collateral/data | Get Collateral Asset Data(USER_DATA)
[**GetLoanableAssetsData**](MarketDataAPI.md#GetLoanableAssetsData) | **Get** /sapi/v1/loan/vip/loanable/data | Get Loanable Assets Data(USER_DATA)


## GetBorrowInterestRate

> GetBorrowInterestRateResponse GetBorrowInterestRate(ctx).LoanCoin(loanCoin).RecvWindow(recvWindow).Execute()

Get Borrow Interest Rate(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/viploan"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	loanCoin := "loanCoin_example" // string | 
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
 **loanCoin** | **string** |  | 
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

Get Collateral Asset Data(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/viploan"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	collateralCoin := "collateralCoin_example" // string |  (optional)
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

Get Loanable Assets Data(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/viploan"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	loanCoin := "loanCoin_example" // string |  (optional)
	vipLevel := int64(1) // int64 | default:user's vip level (optional)
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
 **vipLevel** | **int64** | default:user&#39;s vip level | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetLoanableAssetsDataResponse**](GetLoanableAssetsDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

