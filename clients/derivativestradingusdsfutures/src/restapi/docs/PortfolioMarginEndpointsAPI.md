# \PortfolioMarginEndpointsAPI

All URIs are relative to *https://fapi.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**ClassicPortfolioMarginAccountInformation**](PortfolioMarginEndpointsAPI.md#ClassicPortfolioMarginAccountInformation) | **Get** /fapi/v1/pmAccountInfo | Classic Portfolio Margin Account Information (USER_DATA)


## ClassicPortfolioMarginAccountInformation

> ClassicPortfolioMarginAccountInformationResponse ClassicPortfolioMarginAccountInformation(ctx).Asset(asset).RecvWindow(recvWindow).Execute()

Classic Portfolio Margin Account Information (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	asset := "asset_example" // string | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.PortfolioMarginEndpointsAPI.ClassicPortfolioMarginAccountInformation(context.Background()).Asset(asset).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `PortfolioMarginEndpointsAPI.ClassicPortfolioMarginAccountInformation``: %v\n", err)
		return
	}

	// response from `ClassicPortfolioMarginAccountInformation`: ClassicPortfolioMarginAccountInformationResponse
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
 **recvWindow** | **int64** |  | 

### Return type

[**ClassicPortfolioMarginAccountInformationResponse**](ClassicPortfolioMarginAccountInformationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

