# \GeneralAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**ExchangeInfo**](GeneralAPI.md#ExchangeInfo) | **Get** /api/v3/exchangeInfo | Exchange information
[**ExecutionRules**](GeneralAPI.md#ExecutionRules) | **Get** /api/v3/executionRules | Query Execution Rules
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
	symbol := "ETHBTC" // string | Example: curl -X GET \"https://api.binance.com/api/v3/exchangeInfo?symbol=BNBBTC\" (optional)
	symbols := []string{"BTCUSDT"} // []string | Examples: curl -X GET \"https://api.binance.com/api/v3/exchangeInfo?symbols=%5B%22BNBBTC%22,%22BTCUSDT%22%5D\" or curl -g -X GET 'https://api.binance.com/api/v3/exchangeInfo?symbols=[\"BTCUSDT\",\"BNBBTC\"]' (optional)
	permissions := []models.ExchangeInfoPermissionsParameterInner{models.exchangeInfo_permissions_parameter_inner("SPOT")} // []ExchangeInfoPermissionsParameterInner | Examples: curl -X GET \"https://api.binance.com/api/v3/exchangeInfo?permissions=SPOT\"  curl -X GET \"https://api.binance.com/api/v3/exchangeInfo?permissions=%5B%22MARGIN%22%2C%22LEVERAGED%22%5D\" or curl -g -X GET 'https://api.binance.com/api/v3/exchangeInfo?permissions=[\"MARGIN\",\"LEVERAGED\"]' (optional)
	showPermissionSets := false // bool | Controls whether the content of the `permissionSets` field is populated or not. (optional)
	symbolStatus := models.ExchangeInfoSymbolStatusParameterTrading // ExchangeInfoSymbolStatusParameter | Filters for symbols that have this `tradingStatus`. Cannot be used in combination with `symbols` or `symbol`. (optional)

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
 **symbol** | **string** | Example: curl -X GET \&quot;https://api.binance.com/api/v3/exchangeInfo?symbol&#x3D;BNBBTC\&quot; | 
 **symbols** | **[]string** | Examples: curl -X GET \&quot;https://api.binance.com/api/v3/exchangeInfo?symbols&#x3D;%5B%22BNBBTC%22,%22BTCUSDT%22%5D\&quot; or curl -g -X GET &#39;https://api.binance.com/api/v3/exchangeInfo?symbols&#x3D;[\&quot;BTCUSDT\&quot;,\&quot;BNBBTC\&quot;]&#39; | 
 **permissions** | [**[]ExchangeInfoPermissionsParameterInner**](ExchangeInfoPermissionsParameterInner.md) | Examples: curl -X GET \&quot;https://api.binance.com/api/v3/exchangeInfo?permissions&#x3D;SPOT\&quot;  curl -X GET \&quot;https://api.binance.com/api/v3/exchangeInfo?permissions&#x3D;%5B%22MARGIN%22%2C%22LEVERAGED%22%5D\&quot; or curl -g -X GET &#39;https://api.binance.com/api/v3/exchangeInfo?permissions&#x3D;[\&quot;MARGIN\&quot;,\&quot;LEVERAGED\&quot;]&#39; | 
 **showPermissionSets** | **bool** | Controls whether the content of the &#x60;permissionSets&#x60; field is populated or not. | 
 **symbolStatus** | [**ExchangeInfoSymbolStatusParameter**](ExchangeInfoSymbolStatusParameter.md) | Filters for symbols that have this &#x60;tradingStatus&#x60;. Cannot be used in combination with &#x60;symbols&#x60; or &#x60;symbol&#x60;. | 

### Return type

[**ExchangeInfoResponse**](ExchangeInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## ExecutionRules

> ExecutionRulesResponse ExecutionRules(ctx).Symbol(symbol).Symbols(symbols).SymbolStatus(symbolStatus).Execute()

Query Execution Rules


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
	symbol := "BAZUSD" // string | Query for specified symbol. (optional)
	symbols := []string{"BAZUSD"} // []string | Query for multiple symbols. (optional)
	symbolStatus := models.ExchangeInfoSymbolStatusParameterTrading // ExchangeInfoSymbolStatusParameter | Query for all symbols with the specified status. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.GeneralAPI.ExecutionRules(context.Background()).Symbol(symbol).Symbols(symbols).SymbolStatus(symbolStatus).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `GeneralAPI.ExecutionRules``: %v\n", err)
		return
	}

	// response from `ExecutionRules`: ExecutionRulesResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Query for specified symbol. | 
 **symbols** | **[]string** | Query for multiple symbols. | 
 **symbolStatus** | [**ExchangeInfoSymbolStatusParameter**](ExchangeInfoSymbolStatusParameter.md) | Query for all symbols with the specified status. | 

### Return type

[**ExecutionRulesResponse**](ExecutionRulesResponse.md)

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

