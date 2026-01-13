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
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	optionType := "optionType_example" // string | Input CALL or PUT
	exercisedCoin := "exercisedCoin_example" // string | Target exercised asset, e.g.: if you subscribe to a high sell product (call option), you should input: `optionType`:CALL,`exercisedCoin`:USDT,`investCoin`:BNB; if you subscribe to a low buy product (put option), you should input: `optionType`:PUT,`exercisedCoin`:BNB,`investCoin`:USDT
	investCoin := "investCoin_example" // string | Asset used for subscribing, e.g.: if you subscribe to a high sell product (call option), you should input: `optionType`:CALL,`exercisedCoin`:USDT,`investCoin`:BNB; if you subscribe to a low buy product (put option), you should input: `optionType`:PUT,`exercisedCoin`:BNB,`investCoin`:USDT
	pageSize := int64(10) // int64 | Default: 10, Maximum: 100 (optional)
	pageIndex := int64(1) // int64 | Default: 1 (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

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
 **optionType** | **string** | Input CALL or PUT | 
 **exercisedCoin** | **string** | Target exercised asset, e.g.: if you subscribe to a high sell product (call option), you should input: &#x60;optionType&#x60;:CALL,&#x60;exercisedCoin&#x60;:USDT,&#x60;investCoin&#x60;:BNB; if you subscribe to a low buy product (put option), you should input: &#x60;optionType&#x60;:PUT,&#x60;exercisedCoin&#x60;:BNB,&#x60;investCoin&#x60;:USDT | 
 **investCoin** | **string** | Asset used for subscribing, e.g.: if you subscribe to a high sell product (call option), you should input: &#x60;optionType&#x60;:CALL,&#x60;exercisedCoin&#x60;:USDT,&#x60;investCoin&#x60;:BNB; if you subscribe to a low buy product (put option), you should input: &#x60;optionType&#x60;:PUT,&#x60;exercisedCoin&#x60;:BNB,&#x60;investCoin&#x60;:USDT | 
 **pageSize** | **int64** | Default: 10, Maximum: 100 | 
 **pageIndex** | **int64** | Default: 1 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**GetDualInvestmentProductListResponse**](GetDualInvestmentProductListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

