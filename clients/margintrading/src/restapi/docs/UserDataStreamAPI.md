# \UserDataStreamAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**CloseUserDataStream**](UserDataStreamAPI.md#CloseUserDataStream) | **Delete** /sapi/v1/margin/listen-key | Close User Data Stream (USER_STREAM)
[**CloseUserDataStreamV2**](UserDataStreamAPI.md#CloseUserDataStreamV2) | **Delete** /sapi/v2/margin/listen-key | Close User Data Stream V2 (USER_STREAM)
[**KeepaliveUserDataStream**](UserDataStreamAPI.md#KeepaliveUserDataStream) | **Put** /sapi/v1/margin/listen-key | Keepalive User Data Stream (USER_STREAM)
[**KeepaliveUserDataStreamV2**](UserDataStreamAPI.md#KeepaliveUserDataStreamV2) | **Put** /sapi/v2/margin/listen-key | Keepalive User Data Stream V2 (USER_STREAM)
[**StartUserDataStream**](UserDataStreamAPI.md#StartUserDataStream) | **Post** /sapi/v1/margin/listen-key | Start User Data Stream (USER_STREAM)
[**StartUserDataStreamV2**](UserDataStreamAPI.md#StartUserDataStreamV2) | **Post** /sapi/v2/margin/listen-key | Start User Data Stream V2 (USER_STREAM)


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
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	 err := apiClient.RestApi.UserDataStreamAPI.CloseUserDataStream(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `UserDataStreamAPI.CloseUserDataStream``: %v\n", err)
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


## CloseUserDataStreamV2

> CloseUserDataStreamV2(ctx).Execute()

Close User Data Stream V2 (USER_STREAM)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	 err := apiClient.RestApi.UserDataStreamAPI.CloseUserDataStreamV2(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `UserDataStreamAPI.CloseUserDataStreamV2``: %v\n", err)
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
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	listenKey := "listenKey_example" // string | 

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	 err := apiClient.RestApi.UserDataStreamAPI.KeepaliveUserDataStream(context.Background()).ListenKey(listenKey).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `UserDataStreamAPI.KeepaliveUserDataStream``: %v\n", err)
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


## KeepaliveUserDataStreamV2

> KeepaliveUserDataStreamV2(ctx).ListenKey(listenKey).Execute()

Keepalive User Data Stream V2 (USER_STREAM)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	listenKey := "listenKey_example" // string | 

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	 err := apiClient.RestApi.UserDataStreamAPI.KeepaliveUserDataStreamV2(context.Background()).ListenKey(listenKey).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `UserDataStreamAPI.KeepaliveUserDataStreamV2``: %v\n", err)
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
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.UserDataStreamAPI.StartUserDataStream(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `UserDataStreamAPI.StartUserDataStream``: %v\n", err)
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


## StartUserDataStreamV2

> StartUserDataStreamV2Response StartUserDataStreamV2(ctx).Execute()

Start User Data Stream V2 (USER_STREAM)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.UserDataStreamAPI.StartUserDataStreamV2(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `UserDataStreamAPI.StartUserDataStreamV2``: %v\n", err)
		return
	}

	// response from `StartUserDataStreamV2`: StartUserDataStreamV2Response
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

This endpoint does not need any parameter.

### Return type

[**StartUserDataStreamV2Response**](StartUserDataStreamV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

