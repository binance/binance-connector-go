# \UserDataStreamsAPI

All URIs are relative to *https://papi.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**CloseUserDataStream**](UserDataStreamsAPI.md#CloseUserDataStream) | **Delete** /papi/v1/listenKey | Close User Data Stream(USER_STREAM)
[**KeepaliveUserDataStream**](UserDataStreamsAPI.md#KeepaliveUserDataStream) | **Put** /papi/v1/listenKey | Keepalive User Data Stream (USER_STREAM)
[**StartUserDataStream**](UserDataStreamsAPI.md#StartUserDataStream) | **Post** /papi/v1/listenKey | Start User Data Stream(USER_STREAM)


## CloseUserDataStream

> CloseUserDataStream(ctx).Execute()

Close User Data Stream(USER_STREAM)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	 err := apiClient.RestApi.UserDataStreamsAPI.CloseUserDataStream(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `UserDataStreamsAPI.CloseUserDataStream``: %v\n", err)
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

> KeepaliveUserDataStream(ctx).Execute()

Keepalive User Data Stream (USER_STREAM)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	 err := apiClient.RestApi.UserDataStreamsAPI.KeepaliveUserDataStream(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `UserDataStreamsAPI.KeepaliveUserDataStream``: %v\n", err)
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


## StartUserDataStream

> StartUserDataStreamResponse StartUserDataStream(ctx).Execute()

Start User Data Stream(USER_STREAM)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.UserDataStreamsAPI.StartUserDataStream(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `UserDataStreamsAPI.StartUserDataStream``: %v\n", err)
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

