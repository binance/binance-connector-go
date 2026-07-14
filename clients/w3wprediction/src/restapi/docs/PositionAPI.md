# \PositionAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**GetPositionByToken**](PositionAPI.md#GetPositionByToken) | **Get** /sapi/v1/w3w/wallet/prediction/position/token | Get Position by Token (USER_DATA)
[**QueryPnL**](PositionAPI.md#QueryPnL) | **Get** /sapi/v1/w3w/wallet/prediction/pnl/query | Query PnL (USER_DATA)
[**QueryPositions**](PositionAPI.md#QueryPositions) | **Get** /sapi/v1/w3w/wallet/prediction/position/list | Query Positions (USER_DATA)
[**QueryPositionsByFilter**](PositionAPI.md#QueryPositionsByFilter) | **Get** /sapi/v1/w3w/wallet/prediction/position/filter | Query Positions by Filter (USER_DATA)
[**QuerySettledPositionHistory**](PositionAPI.md#QuerySettledPositionHistory) | **Get** /sapi/v1/w3w/wallet/prediction/position/settled-history | Query Settled Position History (USER_DATA)


## GetPositionByToken

> GetPositionByTokenResponse GetPositionByToken(ctx).WalletAddress(walletAddress).TokenId(tokenId).RecvWindow(recvWindow).Execute()

Get Position by Token (USER_DATA)


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
	tokenId := "112233" // string | Prediction outcome token ID
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.PositionAPI.GetPositionByToken(context.Background()).WalletAddress(walletAddress).TokenId(tokenId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `PositionAPI.GetPositionByToken``: %v\n", err)
		return
	}

	// response from `GetPositionByToken`: GetPositionByTokenResponse
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
 **tokenId** | **string** | Prediction outcome token ID | 
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**GetPositionByTokenResponse**](GetPositionByTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryPnL

> QueryPnLResponse QueryPnL(ctx).WalletAddress(walletAddress).TokenId(tokenId).MarketId(marketId).MarketTopicId(marketTopicId).ActiveOnly(activeOnly).RecvWindow(recvWindow).Execute()

Query PnL (USER_DATA)


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

	resp, err := apiClient.RestApi.PositionAPI.QueryPnL(context.Background()).WalletAddress(walletAddress).TokenId(tokenId).MarketId(marketId).MarketTopicId(marketTopicId).ActiveOnly(activeOnly).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `PositionAPI.QueryPnL``: %v\n", err)
		return
	}

	// response from `QueryPnL`: QueryPnLResponse
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

[**QueryPnLResponse**](QueryPnLResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryPositions

> QueryPositionsResponse QueryPositions(ctx).WalletAddress(walletAddress).Tab(tab).Offset(offset).Limit(limit).RecvWindow(recvWindow).Execute()

Query Positions (USER_DATA)


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
	tab := "ONGOING" // string | Position status tab. Values from `PositionQueryType`. Default `ONGOING` (optional)
	offset := int32(0) // int32 | Pagination offset. Default `0` (optional)
	limit := int32(20) // int32 | Page size. Default `20`, range 1–100 (optional)
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.PositionAPI.QueryPositions(context.Background()).WalletAddress(walletAddress).Tab(tab).Offset(offset).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `PositionAPI.QueryPositions``: %v\n", err)
		return
	}

	// response from `QueryPositions`: QueryPositionsResponse
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
 **tab** | **string** | Position status tab. Values from &#x60;PositionQueryType&#x60;. Default &#x60;ONGOING&#x60; | 
 **offset** | **int32** | Pagination offset. Default &#x60;0&#x60; | 
 **limit** | **int32** | Page size. Default &#x60;20&#x60;, range 1–100 | 
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**QueryPositionsResponse**](QueryPositionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryPositionsByFilter

> QueryPositionsByFilterResponse QueryPositionsByFilter(ctx).WalletAddress(walletAddress).MarketTopicId(marketTopicId).RecvWindow(recvWindow).Execute()

Query Positions by Filter (USER_DATA)


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
	walletAddress := "0x12e32db8817e292508c34111cbc4b23340df542c" // string | User's prediction wallet address (optional)
	marketTopicId := int64(4229564) // int64 | Filter by market topic ID (optional)
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.PositionAPI.QueryPositionsByFilter(context.Background()).WalletAddress(walletAddress).MarketTopicId(marketTopicId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `PositionAPI.QueryPositionsByFilter``: %v\n", err)
		return
	}

	// response from `QueryPositionsByFilter`: QueryPositionsByFilterResponse
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
 **marketTopicId** | **int64** | Filter by market topic ID | 
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**QueryPositionsByFilterResponse**](QueryPositionsByFilterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QuerySettledPositionHistory

> QuerySettledPositionHistoryResponse QuerySettledPositionHistory(ctx).WalletAddress(walletAddress).L1Category(l1Category).Result(result).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).RecvWindow(recvWindow).Execute()

Query Settled Position History (USER_DATA)


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
	l1Category := "crypto" // string | Filter by level-1 category (optional)
	result := int32(1) // int32 | Settlement result filter (optional)
	startDate := "2026-05-01" // string | Start date. Format: `yyyy-MM-dd`. Must be ≤ `endDate` (optional)
	endDate := "2026-05-25" // string | End date. Format: `yyyy-MM-dd`. Must be ≥ `startDate` (optional)
	offset := int32(0) // int32 | Pagination offset. Default `0` (optional)
	limit := int32(20) // int32 | Page size. Default `20`, range 1–100 (optional)
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.PositionAPI.QuerySettledPositionHistory(context.Background()).WalletAddress(walletAddress).L1Category(l1Category).Result(result).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `PositionAPI.QuerySettledPositionHistory``: %v\n", err)
		return
	}

	// response from `QuerySettledPositionHistory`: QuerySettledPositionHistoryResponse
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
 **l1Category** | **string** | Filter by level-1 category | 
 **result** | **int32** | Settlement result filter | 
 **startDate** | **string** | Start date. Format: &#x60;yyyy-MM-dd&#x60;. Must be ≤ &#x60;endDate&#x60; | 
 **endDate** | **string** | End date. Format: &#x60;yyyy-MM-dd&#x60;. Must be ≥ &#x60;startDate&#x60; | 
 **offset** | **int32** | Pagination offset. Default &#x60;0&#x60; | 
 **limit** | **int32** | Page size. Default &#x60;20&#x60;, range 1–100 | 
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**QuerySettledPositionHistoryResponse**](QuerySettledPositionHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

