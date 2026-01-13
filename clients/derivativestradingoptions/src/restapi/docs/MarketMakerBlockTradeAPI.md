# \MarketMakerBlockTradeAPI

All URIs are relative to *https://eapi.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AcceptBlockTradeOrder**](MarketMakerBlockTradeAPI.md#AcceptBlockTradeOrder) | **Post** /eapi/v1/block/order/execute | Accept Block Trade Order (TRADE)
[**AccountBlockTradeList**](MarketMakerBlockTradeAPI.md#AccountBlockTradeList) | **Get** /eapi/v1/block/user-trades | Account Block Trade List (USER_DATA)
[**CancelBlockTradeOrder**](MarketMakerBlockTradeAPI.md#CancelBlockTradeOrder) | **Delete** /eapi/v1/block/order/create | Cancel Block Trade Order (TRADE)
[**ExtendBlockTradeOrder**](MarketMakerBlockTradeAPI.md#ExtendBlockTradeOrder) | **Put** /eapi/v1/block/order/create | Extend Block Trade Order (TRADE)
[**NewBlockTradeOrder**](MarketMakerBlockTradeAPI.md#NewBlockTradeOrder) | **Post** /eapi/v1/block/order/create | New Block Trade Order (TRADE)
[**QueryBlockTradeDetails**](MarketMakerBlockTradeAPI.md#QueryBlockTradeDetails) | **Get** /eapi/v1/block/order/execute | Query Block Trade Details (USER_DATA)
[**QueryBlockTradeOrder**](MarketMakerBlockTradeAPI.md#QueryBlockTradeOrder) | **Get** /eapi/v1/block/order/orders | Query Block Trade Order (TRADE)


## AcceptBlockTradeOrder

> AcceptBlockTradeOrderResponse AcceptBlockTradeOrder(ctx).BlockOrderMatchingKey(blockOrderMatchingKey).RecvWindow(recvWindow).Execute()

Accept Block Trade Order (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	blockOrderMatchingKey := "blockOrderMatchingKey_example" // string | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketMakerBlockTradeAPI.AcceptBlockTradeOrder(context.Background()).BlockOrderMatchingKey(blockOrderMatchingKey).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketMakerBlockTradeAPI.AcceptBlockTradeOrder``: %v\n", err)
		return
	}

	// response from `AcceptBlockTradeOrder`: AcceptBlockTradeOrderResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **blockOrderMatchingKey** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**AcceptBlockTradeOrderResponse**](AcceptBlockTradeOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## AccountBlockTradeList

> AccountBlockTradeListResponse AccountBlockTradeList(ctx).EndTime(endTime).StartTime(startTime).Underlying(underlying).RecvWindow(recvWindow).Execute()

Account Block Trade List (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	endTime := int64(1641782889000) // int64 | End Time, e.g 1593512200000 (optional)
	startTime := int64(1623319461670) // int64 | Start Time, e.g 1593511200000 (optional)
	underlying := "underlying_example" // string | underlying, e.g BTCUSDT (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketMakerBlockTradeAPI.AccountBlockTradeList(context.Background()).EndTime(endTime).StartTime(startTime).Underlying(underlying).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketMakerBlockTradeAPI.AccountBlockTradeList``: %v\n", err)
		return
	}

	// response from `AccountBlockTradeList`: AccountBlockTradeListResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **endTime** | **int64** | End Time, e.g 1593512200000 | 
 **startTime** | **int64** | Start Time, e.g 1593511200000 | 
 **underlying** | **string** | underlying, e.g BTCUSDT | 
 **recvWindow** | **int64** |  | 

### Return type

[**AccountBlockTradeListResponse**](AccountBlockTradeListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CancelBlockTradeOrder

> CancelBlockTradeOrder(ctx).BlockOrderMatchingKey(blockOrderMatchingKey).RecvWindow(recvWindow).Execute()

Cancel Block Trade Order (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	blockOrderMatchingKey := "blockOrderMatchingKey_example" // string | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	 err := apiClient.RestApi.MarketMakerBlockTradeAPI.CancelBlockTradeOrder(context.Background()).BlockOrderMatchingKey(blockOrderMatchingKey).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketMakerBlockTradeAPI.CancelBlockTradeOrder``: %v\n", err)
		return
	}

}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **blockOrderMatchingKey** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: Not defined

[[Back to README]](../../../README.md)


## ExtendBlockTradeOrder

> ExtendBlockTradeOrderResponse ExtendBlockTradeOrder(ctx).BlockOrderMatchingKey(blockOrderMatchingKey).RecvWindow(recvWindow).Execute()

Extend Block Trade Order (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	blockOrderMatchingKey := "blockOrderMatchingKey_example" // string | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketMakerBlockTradeAPI.ExtendBlockTradeOrder(context.Background()).BlockOrderMatchingKey(blockOrderMatchingKey).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketMakerBlockTradeAPI.ExtendBlockTradeOrder``: %v\n", err)
		return
	}

	// response from `ExtendBlockTradeOrder`: ExtendBlockTradeOrderResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **blockOrderMatchingKey** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**ExtendBlockTradeOrderResponse**](ExtendBlockTradeOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## NewBlockTradeOrder

> NewBlockTradeOrderResponse NewBlockTradeOrder(ctx).Liquidity(liquidity).Legs(legs).RecvWindow(recvWindow).Execute()

New Block Trade Order (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	liquidity := "liquidity_example" // string | Taker or Maker
	legs := []map[string]interface{}{map[string]interface{}([{"symbol":"BTC-210115-35000-C","price":"100","quantity":"0.0002","side":"BUY","type":"LIMIT"}])} // []map[string]interface{} | Max 1 (only single leg supported), list of legs parameters in JSON; example: eapi/v1/block/order/create?orders=[{\"symbol\":\"BTC-210115-35000-C\", \"price\":\"100\",\"quantity\":\"0.0002\",\"side\":\"BUY\",\"type\":\"LIMIT\"}]
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketMakerBlockTradeAPI.NewBlockTradeOrder(context.Background()).Liquidity(liquidity).Legs(legs).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketMakerBlockTradeAPI.NewBlockTradeOrder``: %v\n", err)
		return
	}

	// response from `NewBlockTradeOrder`: NewBlockTradeOrderResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **liquidity** | **string** | Taker or Maker | 
 **legs** | **[]map[string]interface{}** | Max 1 (only single leg supported), list of legs parameters in JSON; example: eapi/v1/block/order/create?orders&#x3D;[{\&quot;symbol\&quot;:\&quot;BTC-210115-35000-C\&quot;, \&quot;price\&quot;:\&quot;100\&quot;,\&quot;quantity\&quot;:\&quot;0.0002\&quot;,\&quot;side\&quot;:\&quot;BUY\&quot;,\&quot;type\&quot;:\&quot;LIMIT\&quot;}] | 
 **recvWindow** | **int64** |  | 

### Return type

[**NewBlockTradeOrderResponse**](NewBlockTradeOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryBlockTradeDetails

> QueryBlockTradeDetailsResponse QueryBlockTradeDetails(ctx).BlockOrderMatchingKey(blockOrderMatchingKey).RecvWindow(recvWindow).Execute()

Query Block Trade Details (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	blockOrderMatchingKey := "blockOrderMatchingKey_example" // string | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketMakerBlockTradeAPI.QueryBlockTradeDetails(context.Background()).BlockOrderMatchingKey(blockOrderMatchingKey).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketMakerBlockTradeAPI.QueryBlockTradeDetails``: %v\n", err)
		return
	}

	// response from `QueryBlockTradeDetails`: QueryBlockTradeDetailsResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **blockOrderMatchingKey** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryBlockTradeDetailsResponse**](QueryBlockTradeDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryBlockTradeOrder

> QueryBlockTradeOrderResponse QueryBlockTradeOrder(ctx).BlockOrderMatchingKey(blockOrderMatchingKey).EndTime(endTime).StartTime(startTime).Underlying(underlying).RecvWindow(recvWindow).Execute()

Query Block Trade Order (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	blockOrderMatchingKey := "blockOrderMatchingKey_example" // string | If specified, returns the specific block trade associated with the blockOrderMatchingKey (optional)
	endTime := int64(1641782889000) // int64 | End Time, e.g 1593512200000 (optional)
	startTime := int64(1623319461670) // int64 | Start Time, e.g 1593511200000 (optional)
	underlying := "underlying_example" // string | underlying, e.g BTCUSDT (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketMakerBlockTradeAPI.QueryBlockTradeOrder(context.Background()).BlockOrderMatchingKey(blockOrderMatchingKey).EndTime(endTime).StartTime(startTime).Underlying(underlying).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketMakerBlockTradeAPI.QueryBlockTradeOrder``: %v\n", err)
		return
	}

	// response from `QueryBlockTradeOrder`: QueryBlockTradeOrderResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **blockOrderMatchingKey** | **string** | If specified, returns the specific block trade associated with the blockOrderMatchingKey | 
 **endTime** | **int64** | End Time, e.g 1593512200000 | 
 **startTime** | **int64** | Start Time, e.g 1593511200000 | 
 **underlying** | **string** | underlying, e.g BTCUSDT | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryBlockTradeOrderResponse**](QueryBlockTradeOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

