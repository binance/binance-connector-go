# \ConvertAPI

All URIs are relative to *https://fapi.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AcceptTheOfferedQuote**](ConvertAPI.md#AcceptTheOfferedQuote) | **Post** /fapi/v1/convert/acceptQuote | Accept the offered quote (USER_DATA)
[**ListAllConvertPairs**](ConvertAPI.md#ListAllConvertPairs) | **Get** /fapi/v1/convert/exchangeInfo | List All Convert Pairs
[**OrderStatus**](ConvertAPI.md#OrderStatus) | **Get** /fapi/v1/convert/orderStatus | Order status(USER_DATA)
[**SendQuoteRequest**](ConvertAPI.md#SendQuoteRequest) | **Post** /fapi/v1/convert/getQuote | Send Quote Request(USER_DATA)


## AcceptTheOfferedQuote

> AcceptTheOfferedQuoteResponse AcceptTheOfferedQuote(ctx).QuoteId(quoteId).RecvWindow(recvWindow).Execute()

Accept the offered quote (USER_DATA)


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
	quoteId := "1" // string | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.ConvertAPI.AcceptTheOfferedQuote(context.Background()).QuoteId(quoteId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `ConvertAPI.AcceptTheOfferedQuote``: %v\n", err)
		return
	}

	// response from `AcceptTheOfferedQuote`: AcceptTheOfferedQuoteResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **quoteId** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**AcceptTheOfferedQuoteResponse**](AcceptTheOfferedQuoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


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

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
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
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.ConvertAPI.ListAllConvertPairs(context.Background()).FromAsset(fromAsset).ToAsset(toAsset).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `ConvertAPI.ListAllConvertPairs``: %v\n", err)
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


## OrderStatus

> OrderStatusResponse OrderStatus(ctx).OrderId(orderId).QuoteId(quoteId).Execute()

Order status(USER_DATA)


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
	orderId := "1" // string | Either orderId or quoteId is required (optional)
	quoteId := "1" // string | Either orderId or quoteId is required (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.ConvertAPI.OrderStatus(context.Background()).OrderId(orderId).QuoteId(quoteId).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `ConvertAPI.OrderStatus``: %v\n", err)
		return
	}

	// response from `OrderStatus`: OrderStatusResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **string** | Either orderId or quoteId is required | 
 **quoteId** | **string** | Either orderId or quoteId is required | 

### Return type

[**OrderStatusResponse**](OrderStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SendQuoteRequest

> SendQuoteRequestResponse SendQuoteRequest(ctx).FromAsset(fromAsset).ToAsset(toAsset).FromAmount(fromAmount).ToAmount(toAmount).ValidTime(validTime).RecvWindow(recvWindow).Execute()

Send Quote Request(USER_DATA)


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
	fromAsset := "fromAsset_example" // string | 
	toAsset := "toAsset_example" // string | 
	fromAmount := float32(1.0) // float32 | When specified, it is the amount you will be debited after the conversion (optional)
	toAmount := float32(1.0) // float32 | When specified, it is the amount you will be credited after the conversion (optional)
	validTime := "10s" // string | 10s, default 10s (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.ConvertAPI.SendQuoteRequest(context.Background()).FromAsset(fromAsset).ToAsset(toAsset).FromAmount(fromAmount).ToAmount(toAmount).ValidTime(validTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `ConvertAPI.SendQuoteRequest``: %v\n", err)
		return
	}

	// response from `SendQuoteRequest`: SendQuoteRequestResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **fromAsset** | **string** |  | 
 **toAsset** | **string** |  | 
 **fromAmount** | **float32** | When specified, it is the amount you will be debited after the conversion | 
 **toAmount** | **float32** | When specified, it is the amount you will be credited after the conversion | 
 **validTime** | **string** | 10s, default 10s | 
 **recvWindow** | **int64** |  | 

### Return type

[**SendQuoteRequestResponse**](SendQuoteRequestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

