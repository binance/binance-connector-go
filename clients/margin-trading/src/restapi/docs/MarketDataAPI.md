# \MarketDataAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**CrossMarginCollateralRatio**](MarketDataAPI.md#CrossMarginCollateralRatio) | **Get** /sapi/v1/margin/crossMarginCollateralRatio | Cross margin collateral ratio (MARKET_DATA)
[**GetAllCrossMarginPairs**](MarketDataAPI.md#GetAllCrossMarginPairs) | **Get** /sapi/v1/margin/allPairs | Get All Cross Margin Pairs (MARKET_DATA)
[**GetAllIsolatedMarginSymbol**](MarketDataAPI.md#GetAllIsolatedMarginSymbol) | **Get** /sapi/v1/margin/isolated/allPairs | Get All Isolated Margin Symbol(MARKET_DATA)
[**GetAllMarginAssets**](MarketDataAPI.md#GetAllMarginAssets) | **Get** /sapi/v1/margin/allAssets | Get All Margin Assets (MARKET_DATA)
[**GetDelistSchedule**](MarketDataAPI.md#GetDelistSchedule) | **Get** /sapi/v1/margin/delist-schedule | Get Delist Schedule (MARKET_DATA)
[**GetLimitPricePairs**](MarketDataAPI.md#GetLimitPricePairs) | **Get** /sapi/v1/margin/limit-price-pairs | Get Limit Price Pairs(MARKET_DATA)
[**GetListSchedule**](MarketDataAPI.md#GetListSchedule) | **Get** /sapi/v1/margin/list-schedule | Get list Schedule (MARKET_DATA)
[**QueryIsolatedMarginTierData**](MarketDataAPI.md#QueryIsolatedMarginTierData) | **Get** /sapi/v1/margin/isolatedMarginTier | Query Isolated Margin Tier Data (USER_DATA)
[**QueryLiabilityCoinLeverageBracketInCrossMarginProMode**](MarketDataAPI.md#QueryLiabilityCoinLeverageBracketInCrossMarginProMode) | **Get** /sapi/v1/margin/leverageBracket | Query Liability Coin Leverage Bracket in Cross Margin Pro Mode(MARKET_DATA)
[**QueryMarginAvailableInventory**](MarketDataAPI.md#QueryMarginAvailableInventory) | **Get** /sapi/v1/margin/available-inventory | Query Margin Available Inventory(USER_DATA)
[**QueryMarginPriceindex**](MarketDataAPI.md#QueryMarginPriceindex) | **Get** /sapi/v1/margin/priceIndex | Query Margin PriceIndex (MARKET_DATA)


## CrossMarginCollateralRatio

> CrossMarginCollateralRatioResponse CrossMarginCollateralRatio(ctx).Execute()

Cross margin collateral ratio (MARKET_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.CrossMarginCollateralRatio(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.CrossMarginCollateralRatio``: %v\n", err)
		return
	}

	// response from `CrossMarginCollateralRatio`: CrossMarginCollateralRatioResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

This endpoint does not need any parameter.

### Return type

[**CrossMarginCollateralRatioResponse**](CrossMarginCollateralRatioResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetAllCrossMarginPairs

> GetAllCrossMarginPairsResponse GetAllCrossMarginPairs(ctx).Symbol(symbol).Execute()

Get All Cross Margin Pairs (MARKET_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | isolated margin pair (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.GetAllCrossMarginPairs(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.GetAllCrossMarginPairs``: %v\n", err)
		return
	}

	// response from `GetAllCrossMarginPairs`: GetAllCrossMarginPairsResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | isolated margin pair | 

### Return type

[**GetAllCrossMarginPairsResponse**](GetAllCrossMarginPairsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetAllIsolatedMarginSymbol

> GetAllIsolatedMarginSymbolResponse GetAllIsolatedMarginSymbol(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Get All Isolated Margin Symbol(MARKET_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | isolated margin pair (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.GetAllIsolatedMarginSymbol(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.GetAllIsolatedMarginSymbol``: %v\n", err)
		return
	}

	// response from `GetAllIsolatedMarginSymbol`: GetAllIsolatedMarginSymbolResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | isolated margin pair | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**GetAllIsolatedMarginSymbolResponse**](GetAllIsolatedMarginSymbolResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetAllMarginAssets

> GetAllMarginAssetsResponse GetAllMarginAssets(ctx).Asset(asset).Execute()

Get All Margin Assets (MARKET_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	asset := "asset_example" // string |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.GetAllMarginAssets(context.Background()).Asset(asset).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.GetAllMarginAssets``: %v\n", err)
		return
	}

	// response from `GetAllMarginAssets`: GetAllMarginAssetsResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | 

### Return type

[**GetAllMarginAssetsResponse**](GetAllMarginAssetsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetDelistSchedule

> GetDelistScheduleResponse GetDelistSchedule(ctx).RecvWindow(recvWindow).Execute()

Get Delist Schedule (MARKET_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.GetDelistSchedule(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.GetDelistSchedule``: %v\n", err)
		return
	}

	// response from `GetDelistSchedule`: GetDelistScheduleResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**GetDelistScheduleResponse**](GetDelistScheduleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetLimitPricePairs

> GetLimitPricePairsResponse GetLimitPricePairs(ctx).Execute()

Get Limit Price Pairs(MARKET_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.GetLimitPricePairs(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.GetLimitPricePairs``: %v\n", err)
		return
	}

	// response from `GetLimitPricePairs`: GetLimitPricePairsResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

This endpoint does not need any parameter.

### Return type

[**GetLimitPricePairsResponse**](GetLimitPricePairsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetListSchedule

> GetListScheduleResponse GetListSchedule(ctx).RecvWindow(recvWindow).Execute()

Get list Schedule (MARKET_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.GetListSchedule(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.GetListSchedule``: %v\n", err)
		return
	}

	// response from `GetListSchedule`: GetListScheduleResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**GetListScheduleResponse**](GetListScheduleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryIsolatedMarginTierData

> QueryIsolatedMarginTierDataResponse QueryIsolatedMarginTierData(ctx).Symbol(symbol).Tier(tier).RecvWindow(recvWindow).Execute()

Query Isolated Margin Tier Data (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	tier := int64(789) // int64 | All margin tier data will be returned if tier is omitted (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.QueryIsolatedMarginTierData(context.Background()).Symbol(symbol).Tier(tier).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.QueryIsolatedMarginTierData``: %v\n", err)
		return
	}

	// response from `QueryIsolatedMarginTierData`: QueryIsolatedMarginTierDataResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | 
 **tier** | **int64** | All margin tier data will be returned if tier is omitted | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**QueryIsolatedMarginTierDataResponse**](QueryIsolatedMarginTierDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryLiabilityCoinLeverageBracketInCrossMarginProMode

> QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse QueryLiabilityCoinLeverageBracketInCrossMarginProMode(ctx).Execute()

Query Liability Coin Leverage Bracket in Cross Margin Pro Mode(MARKET_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.QueryLiabilityCoinLeverageBracketInCrossMarginProMode(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.QueryLiabilityCoinLeverageBracketInCrossMarginProMode``: %v\n", err)
		return
	}

	// response from `QueryLiabilityCoinLeverageBracketInCrossMarginProMode`: QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

This endpoint does not need any parameter.

### Return type

[**QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse**](QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryMarginAvailableInventory

> QueryMarginAvailableInventoryResponse QueryMarginAvailableInventory(ctx).Type(type_).Execute()

Query Margin Available Inventory(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	type_ := "type__example" // string | `MARGIN`,`ISOLATED`

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.QueryMarginAvailableInventory(context.Background()).Type(type_).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.QueryMarginAvailableInventory``: %v\n", err)
		return
	}

	// response from `QueryMarginAvailableInventory`: QueryMarginAvailableInventoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | &#x60;MARGIN&#x60;,&#x60;ISOLATED&#x60; | 

### Return type

[**QueryMarginAvailableInventoryResponse**](QueryMarginAvailableInventoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryMarginPriceindex

> QueryMarginPriceindexResponse QueryMarginPriceindex(ctx).Symbol(symbol).Execute()

Query Margin PriceIndex (MARKET_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.QueryMarginPriceindex(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.QueryMarginPriceindex``: %v\n", err)
		return
	}

	// response from `QueryMarginPriceindex`: QueryMarginPriceindexResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | 

### Return type

[**QueryMarginPriceindexResponse**](QueryMarginPriceindexResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

