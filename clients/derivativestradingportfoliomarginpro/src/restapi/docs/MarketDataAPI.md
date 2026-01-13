# \MarketDataAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**GetPortfolioMarginAssetLeverage**](MarketDataAPI.md#GetPortfolioMarginAssetLeverage) | **Get** /sapi/v1/portfolio/margin-asset-leverage | Get Portfolio Margin Asset Leverage(USER_DATA)
[**PortfolioMarginCollateralRate**](MarketDataAPI.md#PortfolioMarginCollateralRate) | **Get** /sapi/v1/portfolio/collateralRate | Portfolio Margin Collateral Rate(MARKET_DATA)
[**PortfolioMarginProTieredCollateralRate**](MarketDataAPI.md#PortfolioMarginProTieredCollateralRate) | **Get** /sapi/v2/portfolio/collateralRate | Portfolio Margin Pro Tiered Collateral Rate(USER_DATA)
[**QueryPortfolioMarginAssetIndexPrice**](MarketDataAPI.md#QueryPortfolioMarginAssetIndexPrice) | **Get** /sapi/v1/portfolio/asset-index-price | Query Portfolio Margin Asset Index Price (MARKET_DATA)


## GetPortfolioMarginAssetLeverage

> GetPortfolioMarginAssetLeverageResponse GetPortfolioMarginAssetLeverage(ctx).Execute()

Get Portfolio Margin Asset Leverage(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomarginpro"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginProClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.GetPortfolioMarginAssetLeverage(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.GetPortfolioMarginAssetLeverage``: %v\n", err)
		return
	}

	// response from `GetPortfolioMarginAssetLeverage`: GetPortfolioMarginAssetLeverageResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

This endpoint does not need any parameter.

### Return type

[**GetPortfolioMarginAssetLeverageResponse**](GetPortfolioMarginAssetLeverageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## PortfolioMarginCollateralRate

> PortfolioMarginCollateralRateResponse PortfolioMarginCollateralRate(ctx).Execute()

Portfolio Margin Collateral Rate(MARKET_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomarginpro"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginProClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.PortfolioMarginCollateralRate(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.PortfolioMarginCollateralRate``: %v\n", err)
		return
	}

	// response from `PortfolioMarginCollateralRate`: PortfolioMarginCollateralRateResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

This endpoint does not need any parameter.

### Return type

[**PortfolioMarginCollateralRateResponse**](PortfolioMarginCollateralRateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## PortfolioMarginProTieredCollateralRate

> PortfolioMarginProTieredCollateralRateResponse PortfolioMarginProTieredCollateralRate(ctx).RecvWindow(recvWindow).Execute()

Portfolio Margin Pro Tiered Collateral Rate(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomarginpro"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginProClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.PortfolioMarginProTieredCollateralRate(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.PortfolioMarginProTieredCollateralRate``: %v\n", err)
		return
	}

	// response from `PortfolioMarginProTieredCollateralRate`: PortfolioMarginProTieredCollateralRateResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** |  | 

### Return type

[**PortfolioMarginProTieredCollateralRateResponse**](PortfolioMarginProTieredCollateralRateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryPortfolioMarginAssetIndexPrice

> QueryPortfolioMarginAssetIndexPriceResponse QueryPortfolioMarginAssetIndexPrice(ctx).Asset(asset).Execute()

Query Portfolio Margin Asset Index Price (MARKET_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomarginpro"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	asset := "asset_example" // string |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginProClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.QueryPortfolioMarginAssetIndexPrice(context.Background()).Asset(asset).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.QueryPortfolioMarginAssetIndexPrice``: %v\n", err)
		return
	}

	// response from `QueryPortfolioMarginAssetIndexPrice`: QueryPortfolioMarginAssetIndexPriceResponse
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

[**QueryPortfolioMarginAssetIndexPriceResponse**](QueryPortfolioMarginAssetIndexPriceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

