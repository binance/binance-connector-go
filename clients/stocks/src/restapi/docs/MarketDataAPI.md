# \MarketDataAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**ExchangeInfo**](MarketDataAPI.md#ExchangeInfo) | **Get** /sapi/v1/equity/market/exchangeInfo | Exchange Info (MARKET_DATA)
[**LatestQuote**](MarketDataAPI.md#LatestQuote) | **Get** /sapi/v1/equity/market/quote | Latest Quote (MARKET_DATA)
[**TokenizedAssets**](MarketDataAPI.md#TokenizedAssets) | **Get** /sapi/v1/equity/market/tokenized-assets | Tokenized Assets (MARKET_DATA)


## ExchangeInfo

> ExchangeInfoResponse ExchangeInfo(ctx).Symbol(symbol).Execute()

Exchange Info (MARKET_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/stocks"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "AAPL" // string | Filter to a single US-equity ticker, e.g. `AAPL`. When omitted, returns all active symbols. An unknown ticker returns an empty `symbols` array (HTTP 200), not an error. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStocksClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.ExchangeInfo(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.ExchangeInfo``: %v\n", err)
		return
	}

	// response from `ExchangeInfo`: ExchangeInfoResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Filter to a single US-equity ticker, e.g. &#x60;AAPL&#x60;. When omitted, returns all active symbols. An unknown ticker returns an empty &#x60;symbols&#x60; array (HTTP 200), not an error. | 

### Return type

[**ExchangeInfoResponse**](ExchangeInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## LatestQuote

> LatestQuoteResponse LatestQuote(ctx).Symbol(symbol).Execute()

Latest Quote (MARKET_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/stocks"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "AAPL" // string | US-equity ticker, e.g. `AAPL`, `TSLA`. Case-insensitive; uppercased server-side.

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStocksClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.LatestQuote(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.LatestQuote``: %v\n", err)
		return
	}

	// response from `LatestQuote`: LatestQuoteResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | US-equity ticker, e.g. &#x60;AAPL&#x60;, &#x60;TSLA&#x60;. Case-insensitive; uppercased server-side. | 

### Return type

[**LatestQuoteResponse**](LatestQuoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## TokenizedAssets

> TokenizedAssetsResponse TokenizedAssets(ctx).Execute()

Tokenized Assets (MARKET_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/stocks"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStocksClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.TokenizedAssets(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.TokenizedAssets``: %v\n", err)
		return
	}

	// response from `TokenizedAssets`: TokenizedAssetsResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

This endpoint does not need any parameter.

### Return type

[**TokenizedAssetsResponse**](TokenizedAssetsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

