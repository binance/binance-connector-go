# \GeneralAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**ExchangeInfo**](GeneralAPI.md#ExchangeInfo) | **Get** /api/v3/exchangeInfo | Exchange information
[**Ping**](GeneralAPI.md#Ping) | **Get** /api/v3/ping | Test connectivity
[**Time**](GeneralAPI.md#Time) | **Get** /api/v3/time | Check server time


## ExchangeInfo

> ExchangeInfoResponse ExchangeInfo(ctx).Symbol(symbol).Symbols(symbols).Permissions(permissions).ShowPermissionSets(showPermissionSets).SymbolStatus(symbolStatus).Execute()

Exchange information


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BNBUSDT" // string | Symbol to query (optional)
	symbols := []string{"Inner_example"} // []string | List of symbols to query (optional)
	permissions := []string{"Inner_example"} // []string | List of permissions to query (optional)
	showPermissionSets := true // bool | Controls whether the content of the `permissionSets` field is populated or not. Defaults to `true` (optional)
	symbolStatus := models.ExchangeInfoSymbolStatusParameterTrading // ExchangeInfoSymbolStatusParameter |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.GeneralAPI.ExchangeInfo(context.Background()).Symbol(symbol).Symbols(symbols).Permissions(permissions).ShowPermissionSets(showPermissionSets).SymbolStatus(symbolStatus).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `GeneralAPI.ExchangeInfo``: %v\n", err)
		return
	}

	// response from `ExchangeInfo`: ExchangeInfoResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol to query | 
 **symbols** | **[]string** | List of symbols to query | 
 **permissions** | **[]string** | List of permissions to query | 
 **showPermissionSets** | **bool** | Controls whether the content of the &#x60;permissionSets&#x60; field is populated or not. Defaults to &#x60;true&#x60; | 
 **symbolStatus** | [**ExchangeInfoSymbolStatusParameter**](ExchangeInfoSymbolStatusParameter.md) |  | 

### Return type

[**ExchangeInfoResponse**](ExchangeInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## Ping

> Ping(ctx).Execute()

Test connectivity


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	 err := apiClient.RestApi.GeneralAPI.Ping(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `GeneralAPI.Ping``: %v\n", err)
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


## Time

> TimeResponse Time(ctx).Execute()

Check server time


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.GeneralAPI.Time(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `GeneralAPI.Time``: %v\n", err)
		return
	}

	// response from `Time`: TimeResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

This endpoint does not need any parameter.

### Return type

[**TimeResponse**](TimeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

