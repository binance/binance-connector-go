# \OtcAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**CreateOtcBlocktrade**](OtcAPI.md#CreateOtcBlocktrade) | **Post** /sapi/v1/w3w/wallet/prediction/otc/blocktrade/create | Create OTC Blocktrade (PREDICTION_TRADE)
[**FulfilOtcBlocktrade**](OtcAPI.md#FulfilOtcBlocktrade) | **Post** /sapi/v1/w3w/wallet/prediction/otc/blocktrade/fulfil | Fulfil OTC Blocktrade (PREDICTION_TRADE)
[**GetOtcBlocktradeDetail**](OtcAPI.md#GetOtcBlocktradeDetail) | **Post** /sapi/v1/w3w/wallet/prediction/otc/blocktrade/detail | Get OTC Blocktrade Detail (PREDICTION_TRADE)
[**GetOtcBlocktradeEvents**](OtcAPI.md#GetOtcBlocktradeEvents) | **Post** /sapi/v1/w3w/wallet/prediction/otc/blocktrade/events | Get OTC Blocktrade Events (PREDICTION_TRADE)
[**GetOtcReservedBalances**](OtcAPI.md#GetOtcReservedBalances) | **Post** /sapi/v1/w3w/wallet/prediction/otc/blocktrade/reserved-balances | Get OTC Reserved Balances (PREDICTION_TRADE)
[**ListOtcBlocktrades**](OtcAPI.md#ListOtcBlocktrades) | **Post** /sapi/v1/w3w/wallet/prediction/otc/blocktrade/list | List OTC Blocktrades (PREDICTION_TRADE)
[**PreviewOtcBlocktrade**](OtcAPI.md#PreviewOtcBlocktrade) | **Post** /sapi/v1/w3w/wallet/prediction/otc/blocktrade/preview | Preview OTC Blocktrade (PREDICTION_TRADE)
[**RemoveOtcBlocktrades**](OtcAPI.md#RemoveOtcBlocktrades) | **Post** /sapi/v1/w3w/wallet/prediction/otc/blocktrade/remove | Remove OTC Blocktrades (PREDICTION_TRADE)


## CreateOtcBlocktrade

> CreateOtcBlocktradeResponse CreateOtcBlocktrade(ctx).MarketId(marketId).TokenId(tokenId).Side(side).MakerAmount(makerAmount).TakerAmount(takerAmount).PricePerShare(pricePerShare).Expiration(expiration).Execute()

Create OTC Blocktrade (PREDICTION_TRADE)


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
	marketId := "123" // string | PredictFun market id
	tokenId := "71321045679252212594626385532706912750332728571942532289631379312455583992563" // string | ERC-1155 outcome token id
	side := models.GetQuoteSideParameterBuy // GetQuoteSideParameter | Trade side. Enum: `BUY` (BID), `SELL` (ASK)
	makerAmount := "600000000000000000000" // string | Maker amount in wei. BID：USDT; ASK：shares
	takerAmount := "1000000000000000000000" // string | Taker amount in wei. BID：shares; ASK：USDT
	pricePerShare := "0.65" // string | Price per share (decimal ether, e.g. `0.65`)
	expiration := int64(1790000000) // int64 | Expiration timestamp in seconds (GTD order)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.OtcAPI.CreateOtcBlocktrade(context.Background()).MarketId(marketId).TokenId(tokenId).Side(side).MakerAmount(makerAmount).TakerAmount(takerAmount).PricePerShare(pricePerShare).Expiration(expiration).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `OtcAPI.CreateOtcBlocktrade``: %v\n", err)
		return
	}

	// response from `CreateOtcBlocktrade`: CreateOtcBlocktradeResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **marketId** | **string** | PredictFun market id | 
 **tokenId** | **string** | ERC-1155 outcome token id | 
 **side** | [**GetQuoteSideParameter**](GetQuoteSideParameter.md) | Trade side. Enum: &#x60;BUY&#x60; (BID), &#x60;SELL&#x60; (ASK) | 
 **makerAmount** | **string** | Maker amount in wei. BID：USDT; ASK：shares | 
 **takerAmount** | **string** | Taker amount in wei. BID：shares; ASK：USDT | 
 **pricePerShare** | **string** | Price per share (decimal ether, e.g. &#x60;0.65&#x60;) | 
 **expiration** | **int64** | Expiration timestamp in seconds (GTD order) | 

### Return type

[**CreateOtcBlocktradeResponse**](CreateOtcBlocktradeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## FulfilOtcBlocktrade

> FulfilOtcBlocktradeResponse FulfilOtcBlocktrade(ctx).OrderId(orderId).SecretToken(secretToken).Execute()

Fulfil OTC Blocktrade (PREDICTION_TRADE)


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
	orderId := "26080500000001234567" // string | Maker order id (returned by maker create)
	secretToken := "a1b2c3d4-e5f6-7890-abcd-ef1234567890" // string | One-time fulfilment token the maker shared out-of-band

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.OtcAPI.FulfilOtcBlocktrade(context.Background()).OrderId(orderId).SecretToken(secretToken).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `OtcAPI.FulfilOtcBlocktrade``: %v\n", err)
		return
	}

	// response from `FulfilOtcBlocktrade`: FulfilOtcBlocktradeResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **string** | Maker order id (returned by maker create) | 
 **secretToken** | **string** | One-time fulfilment token the maker shared out-of-band | 

### Return type

[**FulfilOtcBlocktradeResponse**](FulfilOtcBlocktradeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetOtcBlocktradeDetail

> GetOtcBlocktradeDetailResponse GetOtcBlocktradeDetail(ctx).OrderId(orderId).Execute()

Get OTC Blocktrade Detail (PREDICTION_TRADE)


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
	orderId := "26080500000001234567" // string | Maker order id (returned by create)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.OtcAPI.GetOtcBlocktradeDetail(context.Background()).OrderId(orderId).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `OtcAPI.GetOtcBlocktradeDetail``: %v\n", err)
		return
	}

	// response from `GetOtcBlocktradeDetail`: GetOtcBlocktradeDetailResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **string** | Maker order id (returned by create) | 

### Return type

[**GetOtcBlocktradeDetailResponse**](GetOtcBlocktradeDetailResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetOtcBlocktradeEvents

> GetOtcBlocktradeEventsResponse GetOtcBlocktradeEvents(ctx).First(first).After(after).EventTypes(eventTypes).MarketId(marketId).Execute()

Get OTC Blocktrade Events (PREDICTION_TRADE)


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
	first := int32(20) // int32 | Page size (optional)
	after := "next-cursor-string" // string | Pagination cursor (optional)
	eventTypes := []string{"Inner_example"} // []string | Filter by event types (e.g. `[\"MATCH_SUCCESS\",\"EXPIRE\"]`) (optional)
	marketId := int64(123) // int64 | Filter by market id (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.OtcAPI.GetOtcBlocktradeEvents(context.Background()).First(first).After(after).EventTypes(eventTypes).MarketId(marketId).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `OtcAPI.GetOtcBlocktradeEvents``: %v\n", err)
		return
	}

	// response from `GetOtcBlocktradeEvents`: GetOtcBlocktradeEventsResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **first** | **int32** | Page size | 
 **after** | **string** | Pagination cursor | 
 **eventTypes** | **[]string** | Filter by event types (e.g. &#x60;[\&quot;MATCH_SUCCESS\&quot;,\&quot;EXPIRE\&quot;]&#x60;) | 
 **marketId** | **int64** | Filter by market id | 

### Return type

[**GetOtcBlocktradeEventsResponse**](GetOtcBlocktradeEventsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetOtcReservedBalances

> GetOtcReservedBalancesResponse GetOtcReservedBalances(ctx).Assets(assets).Execute()

Get OTC Reserved Balances (PREDICTION_TRADE)


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
	assets := []models.GetOtcReservedBalancesAssetsParameterInner{*models.NewGetOtcReservedBalancesAssetsParameterInner()} // []GetOtcReservedBalancesAssetsParameterInner | Assets to query (max 50)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.OtcAPI.GetOtcReservedBalances(context.Background()).Assets(assets).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `OtcAPI.GetOtcReservedBalances``: %v\n", err)
		return
	}

	// response from `GetOtcReservedBalances`: GetOtcReservedBalancesResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **assets** | [**[]GetOtcReservedBalancesAssetsParameterInner**](GetOtcReservedBalancesAssetsParameterInner.md) | Assets to query (max 50) | 

### Return type

[**GetOtcReservedBalancesResponse**](GetOtcReservedBalancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## ListOtcBlocktrades

> ListOtcBlocktradesResponse ListOtcBlocktrades(ctx).First(first).After(after).Status(status).Execute()

List OTC Blocktrades (PREDICTION_TRADE)


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
	first := int32(20) // int32 | Page size (optional)
	after := "next-cursor-string" // string | Pagination cursor (from previous response) (optional)
	status := models.ListOtcBlocktradesStatusParameterOpen // ListOtcBlocktradesStatusParameter | Filter by status. Enum: `OPEN`, `FULFILLED`, `MATCHED`, `CANCELLED`, `EXPIRED`, `FAILED` (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.OtcAPI.ListOtcBlocktrades(context.Background()).First(first).After(after).Status(status).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `OtcAPI.ListOtcBlocktrades``: %v\n", err)
		return
	}

	// response from `ListOtcBlocktrades`: ListOtcBlocktradesResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **first** | **int32** | Page size | 
 **after** | **string** | Pagination cursor (from previous response) | 
 **status** | [**ListOtcBlocktradesStatusParameter**](ListOtcBlocktradesStatusParameter.md) | Filter by status. Enum: &#x60;OPEN&#x60;, &#x60;FULFILLED&#x60;, &#x60;MATCHED&#x60;, &#x60;CANCELLED&#x60;, &#x60;EXPIRED&#x60;, &#x60;FAILED&#x60; | 

### Return type

[**ListOtcBlocktradesResponse**](ListOtcBlocktradesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## PreviewOtcBlocktrade

> PreviewOtcBlocktradeResponse PreviewOtcBlocktrade(ctx).SecretToken(secretToken).Execute()

Preview OTC Blocktrade (PREDICTION_TRADE)


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
	secretToken := "a1b2c3d4-e5f6-7890-abcd-ef1234567890" // string | One-time fulfilment token the maker shared out-of-band

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.OtcAPI.PreviewOtcBlocktrade(context.Background()).SecretToken(secretToken).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `OtcAPI.PreviewOtcBlocktrade``: %v\n", err)
		return
	}

	// response from `PreviewOtcBlocktrade`: PreviewOtcBlocktradeResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **secretToken** | **string** | One-time fulfilment token the maker shared out-of-band | 

### Return type

[**PreviewOtcBlocktradeResponse**](PreviewOtcBlocktradeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## RemoveOtcBlocktrades

> RemoveOtcBlocktradesResponse RemoveOtcBlocktrades(ctx).OrderIds(orderIds).Execute()

Remove OTC Blocktrades (PREDICTION_TRADE)


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
	orderIds := []string{"Inner_example"} // []string | Order ids to remove (max 100). Must be the `orderId` returned by Create OTC Blocktrade

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.OtcAPI.RemoveOtcBlocktrades(context.Background()).OrderIds(orderIds).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `OtcAPI.RemoveOtcBlocktrades``: %v\n", err)
		return
	}

	// response from `RemoveOtcBlocktrades`: RemoveOtcBlocktradesResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **orderIds** | **[]string** | Order ids to remove (max 100). Must be the &#x60;orderId&#x60; returned by Create OTC Blocktrade | 

### Return type

[**RemoveOtcBlocktradesResponse**](RemoveOtcBlocktradesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

