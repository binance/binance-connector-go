# \NFTAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**GetNFTAsset**](NFTAPI.md#GetNFTAsset) | **Get** /sapi/v1/nft/user/getAsset | Get NFT Asset(USER_DATA)
[**GetNFTDepositHistory**](NFTAPI.md#GetNFTDepositHistory) | **Get** /sapi/v1/nft/history/deposit | Get NFT Deposit History(USER_DATA)
[**GetNFTTransactionHistory**](NFTAPI.md#GetNFTTransactionHistory) | **Get** /sapi/v1/nft/history/transactions | Get NFT Transaction History(USER_DATA)
[**GetNFTWithdrawHistory**](NFTAPI.md#GetNFTWithdrawHistory) | **Get** /sapi/v1/nft/history/withdraw | Get NFT Withdraw History(USER_DATA)


## GetNFTAsset

> GetNFTAssetResponse GetNFTAsset(ctx).Limit(limit).Page(page).RecvWindow(recvWindow).Execute()

Get NFT Asset(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/nft"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	limit := int64(50) // int64 | Default 50, Max 50 (optional)
	page := int64(1) // int64 | Default 1 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceNFTClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.NFTAPI.GetNFTAsset(context.Background()).Limit(limit).Page(page).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `NFTAPI.GetNFTAsset``: %v\n", err)
		return
	}

	// response from `GetNFTAsset`: GetNFTAssetResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | Default 50, Max 50 | 
 **page** | **int64** | Default 1 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetNFTAssetResponse**](GetNFTAssetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetNFTDepositHistory

> GetNFTDepositHistoryResponse GetNFTDepositHistory(ctx).StartTime(startTime).EndTime(endTime).Limit(limit).Page(page).RecvWindow(recvWindow).Execute()

Get NFT Deposit History(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/nft"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	limit := int64(50) // int64 | Default 50, Max 50 (optional)
	page := int64(1) // int64 | Default 1 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceNFTClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.NFTAPI.GetNFTDepositHistory(context.Background()).StartTime(startTime).EndTime(endTime).Limit(limit).Page(page).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `NFTAPI.GetNFTDepositHistory``: %v\n", err)
		return
	}

	// response from `GetNFTDepositHistory`: GetNFTDepositHistoryResponse
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
 **limit** | **int64** | Default 50, Max 50 | 
 **page** | **int64** | Default 1 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetNFTDepositHistoryResponse**](GetNFTDepositHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetNFTTransactionHistory

> GetNFTTransactionHistoryResponse GetNFTTransactionHistory(ctx).OrderType(orderType).StartTime(startTime).EndTime(endTime).Limit(limit).Page(page).RecvWindow(recvWindow).Execute()

Get NFT Transaction History(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/nft"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	orderType := int64(789) // int64 | 0: purchase order, 1: sell order, 2: royalty income, 3: primary market order, 4: mint fee
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	limit := int64(50) // int64 | Default 50, Max 50 (optional)
	page := int64(1) // int64 | Default 1 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceNFTClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.NFTAPI.GetNFTTransactionHistory(context.Background()).OrderType(orderType).StartTime(startTime).EndTime(endTime).Limit(limit).Page(page).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `NFTAPI.GetNFTTransactionHistory``: %v\n", err)
		return
	}

	// response from `GetNFTTransactionHistory`: GetNFTTransactionHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **orderType** | **int64** | 0: purchase order, 1: sell order, 2: royalty income, 3: primary market order, 4: mint fee | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int64** | Default 50, Max 50 | 
 **page** | **int64** | Default 1 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetNFTTransactionHistoryResponse**](GetNFTTransactionHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetNFTWithdrawHistory

> GetNFTWithdrawHistoryResponse GetNFTWithdrawHistory(ctx).StartTime(startTime).EndTime(endTime).Limit(limit).Page(page).RecvWindow(recvWindow).Execute()

Get NFT Withdraw History(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/nft"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	limit := int64(50) // int64 | Default 50, Max 50 (optional)
	page := int64(1) // int64 | Default 1 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceNFTClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.NFTAPI.GetNFTWithdrawHistory(context.Background()).StartTime(startTime).EndTime(endTime).Limit(limit).Page(page).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `NFTAPI.GetNFTWithdrawHistory``: %v\n", err)
		return
	}

	// response from `GetNFTWithdrawHistory`: GetNFTWithdrawHistoryResponse
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
 **limit** | **int64** | Default 50, Max 50 | 
 **page** | **int64** | Default 1 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetNFTWithdrawHistoryResponse**](GetNFTWithdrawHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

