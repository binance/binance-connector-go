# \WalletAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**GetPortfolio**](WalletAPI.md#GetPortfolio) | **Get** /sapi/v1/w3w/wallet/prediction/pnl/portfolio | Get Portfolio (USER_DATA)
[**GetQuotaStatus**](WalletAPI.md#GetQuotaStatus) | **Get** /sapi/v1/w3w/wallet/prediction/quota/limit/status | Get Quota Status (USER_DATA)
[**ListPredictionWallets**](WalletAPI.md#ListPredictionWallets) | **Get** /sapi/v1/w3w/wallet/prediction/wallet/list | List Prediction Wallets (USER_DATA)
[**QueryPaymentOptionBalances**](WalletAPI.md#QueryPaymentOptionBalances) | **Get** /sapi/v1/w3w/wallet/prediction/balance/payment-options | Query Payment Option Balances (USER_DATA)


## GetPortfolio

> GetPortfolioResponse GetPortfolio(ctx).WalletAddress(walletAddress).TokenId(tokenId).MarketId(marketId).MarketTopicId(marketTopicId).ActiveOnly(activeOnly).RecvWindow(recvWindow).Execute()

Get Portfolio (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/w3wprediction"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	walletAddress := "0x12e32db8817e292508c34111cbc4b23340df542c" // string | User's prediction wallet address
	tokenId := "112233" // string | Filter by prediction token ID (optional)
	marketId := int64(5567895) // int64 | Filter by market ID. Must be > 0 (optional)
	marketTopicId := int64(4229564) // int64 | Filter by market topic ID. Must be > 0 (optional)
	activeOnly := false // bool | If `true`, return only active (unresolved) positions (optional)
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.WalletAPI.GetPortfolio(context.Background()).WalletAddress(walletAddress).TokenId(tokenId).MarketId(marketId).MarketTopicId(marketTopicId).ActiveOnly(activeOnly).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WalletAPI.GetPortfolio``: %v\n", err)
		return
	}

	// response from `GetPortfolio`: GetPortfolioResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **walletAddress** | **string** | User&#39;s prediction wallet address | 
 **tokenId** | **string** | Filter by prediction token ID | 
 **marketId** | **int64** | Filter by market ID. Must be &gt; 0 | 
 **marketTopicId** | **int64** | Filter by market topic ID. Must be &gt; 0 | 
 **activeOnly** | **bool** | If &#x60;true&#x60;, return only active (unresolved) positions | 
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**GetPortfolioResponse**](GetPortfolioResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetQuotaStatus

> GetQuotaStatusResponse GetQuotaStatus(ctx).RecvWindow(recvWindow).Execute()

Get Quota Status (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/w3wprediction"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.WalletAPI.GetQuotaStatus(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WalletAPI.GetQuotaStatus``: %v\n", err)
		return
	}

	// response from `GetQuotaStatus`: GetQuotaStatusResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**GetQuotaStatusResponse**](GetQuotaStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## ListPredictionWallets

> ListPredictionWalletsResponse ListPredictionWallets(ctx).RecvWindow(recvWindow).Execute()

List Prediction Wallets (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/w3wprediction"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.WalletAPI.ListPredictionWallets(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WalletAPI.ListPredictionWallets``: %v\n", err)
		return
	}

	// response from `ListPredictionWallets`: ListPredictionWalletsResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**ListPredictionWalletsResponse**](ListPredictionWalletsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryPaymentOptionBalances

> QueryPaymentOptionBalancesResponse QueryPaymentOptionBalances(ctx).RecvWindow(recvWindow).Execute()

Query Payment Option Balances (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/w3wprediction"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.WalletAPI.QueryPaymentOptionBalances(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WalletAPI.QueryPaymentOptionBalances``: %v\n", err)
		return
	}

	// response from `QueryPaymentOptionBalances`: QueryPaymentOptionBalancesResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**QueryPaymentOptionBalancesResponse**](QueryPaymentOptionBalancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

