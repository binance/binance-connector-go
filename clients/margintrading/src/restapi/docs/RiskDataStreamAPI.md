# \RiskDataStreamAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**CloseUserDataStream**](RiskDataStreamAPI.md#CloseUserDataStream) | **Delete** /sapi/v1/margin/listen-key | Close User Data Stream (USER_STREAM)
[**KeepaliveUserDataStream**](RiskDataStreamAPI.md#KeepaliveUserDataStream) | **Put** /sapi/v1/margin/listen-key | Keepalive User Data Stream (USER_STREAM)
[**StartUserDataStream**](RiskDataStreamAPI.md#StartUserDataStream) | **Post** /sapi/v1/margin/listen-key | Start User Data Stream (USER_STREAM)


## CloseUserDataStream

> CloseUserDataStream(ctx).Execute()

Close User Data Stream (USER_STREAM)


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

	 err := apiClient.RestApi.RiskDataStreamAPI.CloseUserDataStream(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `RiskDataStreamAPI.CloseUserDataStream``: %v\n", err)
		return
	}

}
```

### Path Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: Not defined

[[Back to README]](../../../README.md)


## KeepaliveUserDataStream

> KeepaliveUserDataStream(ctx).ListenKey(listenKey).Execute()

Keepalive User Data Stream (USER_STREAM)


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
	listenKey := "listenKey_example" // string | 

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	 err := apiClient.RestApi.RiskDataStreamAPI.KeepaliveUserDataStream(context.Background()).ListenKey(listenKey).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `RiskDataStreamAPI.KeepaliveUserDataStream``: %v\n", err)
		return
	}

}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **listenKey** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: Not defined

[[Back to README]](../../../README.md)


## StartUserDataStream

> StartUserDataStreamResponse StartUserDataStream(ctx).Execute()

Start User Data Stream (USER_STREAM)


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

	resp, err := apiClient.RestApi.RiskDataStreamAPI.StartUserDataStream(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `RiskDataStreamAPI.StartUserDataStream``: %v\n", err)
		return
	}

	// response from `StartUserDataStream`: StartUserDataStreamResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

This endpoint does not need any parameter.

### Return type

[**StartUserDataStreamResponse**](StartUserDataStreamResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

