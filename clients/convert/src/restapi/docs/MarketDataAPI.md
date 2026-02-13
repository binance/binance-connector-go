# \MarketDataAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**ListAllConvertPairs**](MarketDataAPI.md#ListAllConvertPairs) | **Get** /sapi/v1/convert/exchangeInfo | List All Convert Pairs
[**QueryOrderQuantityPrecisionPerAsset**](MarketDataAPI.md#QueryOrderQuantityPrecisionPerAsset) | **Get** /sapi/v1/convert/assetInfo | Query order quantity precision per asset(USER_DATA)


## ListAllConvertPairs

> ListAllConvertPairsResponse ListAllConvertPairs(ctx).FromAsset(fromAsset).ToAsset(toAsset).Execute()

List All Convert Pairs


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/convert"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	fromAsset := "fromAsset_example" // string | User spends coin (optional)
	toAsset := "toAsset_example" // string | User receives coin (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceConvertClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.ListAllConvertPairs(context.Background()).FromAsset(fromAsset).ToAsset(toAsset).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.ListAllConvertPairs``: %v\n", err)
		return
	}

	// response from `ListAllConvertPairs`: ListAllConvertPairsResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **fromAsset** | **string** | User spends coin | 
 **toAsset** | **string** | User receives coin | 

### Return type

[**ListAllConvertPairsResponse**](ListAllConvertPairsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryOrderQuantityPrecisionPerAsset

> QueryOrderQuantityPrecisionPerAssetResponse QueryOrderQuantityPrecisionPerAsset(ctx).RecvWindow(recvWindow).Execute()

Query order quantity precision per asset(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/convert"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceConvertClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.QueryOrderQuantityPrecisionPerAsset(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.QueryOrderQuantityPrecisionPerAsset``: %v\n", err)
		return
	}

	// response from `QueryOrderQuantityPrecisionPerAsset`: QueryOrderQuantityPrecisionPerAssetResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**QueryOrderQuantityPrecisionPerAssetResponse**](QueryOrderQuantityPrecisionPerAssetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

