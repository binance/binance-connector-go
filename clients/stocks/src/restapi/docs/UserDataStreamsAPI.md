# \UserDataStreamsAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**CreateRenewListenKey**](UserDataStreamsAPI.md#CreateRenewListenKey) | **Post** /sapi/v1/equity/listenKey | Create / Renew Listen Key (USER_STREAM)


## CreateRenewListenKey

> CreateRenewListenKeyResponse CreateRenewListenKey(ctx).RecvWindow(recvWindow).Execute()

Create / Renew Listen Key (USER_STREAM)


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
	recvWindow := int64(5000) // int64 | The value cannot be greater than `60000`. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStocksClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.UserDataStreamsAPI.CreateRenewListenKey(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `UserDataStreamsAPI.CreateRenewListenKey``: %v\n", err)
		return
	}

	// response from `CreateRenewListenKey`: CreateRenewListenKeyResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60;. | 

### Return type

[**CreateRenewListenKeyResponse**](CreateRenewListenKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

