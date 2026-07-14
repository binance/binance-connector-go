# \MarketDataAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**GetDualInvestmentProductList**](MarketDataAPI.md#GetDualInvestmentProductList) | **Get** /sapi/v1/dci/product/list | Get Dual Investment product list


## GetDualInvestmentProductList

> GetDualInvestmentProductListResponse GetDualInvestmentProductList(ctx).OptionType(optionType).ExercisedCoin(exercisedCoin).InvestCoin(investCoin).PageSize(pageSize).PageIndex(pageIndex).RecvWindow(recvWindow).Execute()

Get Dual Investment product list


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/dualinvestment"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	optionType := models.GetDualInvestmentProductListOptionTypeParameterCall // GetDualInvestmentProductListOptionTypeParameter | Input CALL or PUT
	exercisedCoin := "USDT" // string | Target exercised asset, e.g.: if you subscribe to a high sell product (call option), you should input: `optionType: CALL`, `exercisedCoin: USDT`, `investCoin: BNB`; if you subscribe to a low buy product (put option), you should input: `optionType: PUT`, `exercisedCoin: BNB`, `investCoin: USDT`
	investCoin := "BNB" // string | Asset used for subscribing, e.g.: if you subscribe to a high sell product (call option), you should input: `optionType: CALL`, `exercisedCoin: USDT`, `investCoin: BNB`; if you subscribe to a low buy product (put option), you should input: `optionType: PUT`, `exercisedCoin: BNB`, `investCoin: USDT`
	pageSize := int64(10) // int64 | Number of records per page (optional)
	pageIndex := int64(1) // int64 | Page index (optional)
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDualInvestmentClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.GetDualInvestmentProductList(context.Background()).OptionType(optionType).ExercisedCoin(exercisedCoin).InvestCoin(investCoin).PageSize(pageSize).PageIndex(pageIndex).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.GetDualInvestmentProductList``: %v\n", err)
		return
	}

	// response from `GetDualInvestmentProductList`: GetDualInvestmentProductListResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **optionType** | [**GetDualInvestmentProductListOptionTypeParameter**](GetDualInvestmentProductListOptionTypeParameter.md) | Input CALL or PUT | 
 **exercisedCoin** | **string** | Target exercised asset, e.g.: if you subscribe to a high sell product (call option), you should input: &#x60;optionType: CALL&#x60;, &#x60;exercisedCoin: USDT&#x60;, &#x60;investCoin: BNB&#x60;; if you subscribe to a low buy product (put option), you should input: &#x60;optionType: PUT&#x60;, &#x60;exercisedCoin: BNB&#x60;, &#x60;investCoin: USDT&#x60; | 
 **investCoin** | **string** | Asset used for subscribing, e.g.: if you subscribe to a high sell product (call option), you should input: &#x60;optionType: CALL&#x60;, &#x60;exercisedCoin: USDT&#x60;, &#x60;investCoin: BNB&#x60;; if you subscribe to a low buy product (put option), you should input: &#x60;optionType: PUT&#x60;, &#x60;exercisedCoin: BNB&#x60;, &#x60;investCoin: USDT&#x60; | 
 **pageSize** | **int64** | Number of records per page | 
 **pageIndex** | **int64** | Page index | 
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**GetDualInvestmentProductListResponse**](GetDualInvestmentProductListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

