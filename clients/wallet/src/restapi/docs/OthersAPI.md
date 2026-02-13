# \OthersAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**GetSymbolsDelistScheduleForSpot**](OthersAPI.md#GetSymbolsDelistScheduleForSpot) | **Get** /sapi/v1/spot/delist-schedule | Get symbols delist schedule for spot (MARKET_DATA)
[**SystemStatus**](OthersAPI.md#SystemStatus) | **Get** /sapi/v1/system/status | System Status (System)


## GetSymbolsDelistScheduleForSpot

> GetSymbolsDelistScheduleForSpotResponse GetSymbolsDelistScheduleForSpot(ctx).RecvWindow(recvWindow).Execute()

Get symbols delist schedule for spot (MARKET_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.OthersAPI.GetSymbolsDelistScheduleForSpot(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `OthersAPI.GetSymbolsDelistScheduleForSpot``: %v\n", err)
		return
	}

	// response from `GetSymbolsDelistScheduleForSpot`: GetSymbolsDelistScheduleForSpotResponse
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

[**GetSymbolsDelistScheduleForSpotResponse**](GetSymbolsDelistScheduleForSpotResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SystemStatus

> SystemStatusResponse SystemStatus(ctx).Execute()

System Status (System)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.OthersAPI.SystemStatus(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `OthersAPI.SystemStatus``: %v\n", err)
		return
	}

	// response from `SystemStatus`: SystemStatusResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

This endpoint does not need any parameter.

### Return type

[**SystemStatusResponse**](SystemStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

