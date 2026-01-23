# \AccountAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AdjustCrossMarginMaxLeverage**](AccountAPI.md#AdjustCrossMarginMaxLeverage) | **Post** /sapi/v1/margin/max-leverage | Adjust cross margin max leverage (USER_DATA)
[**DisableIsolatedMarginAccount**](AccountAPI.md#DisableIsolatedMarginAccount) | **Delete** /sapi/v1/margin/isolated/account | Disable Isolated Margin Account (TRADE)
[**EnableIsolatedMarginAccount**](AccountAPI.md#EnableIsolatedMarginAccount) | **Post** /sapi/v1/margin/isolated/account | Enable Isolated Margin Account (TRADE)
[**GetBnbBurnStatus**](AccountAPI.md#GetBnbBurnStatus) | **Get** /sapi/v1/bnbBurn | Get BNB Burn Status (USER_DATA)
[**GetSummaryOfMarginAccount**](AccountAPI.md#GetSummaryOfMarginAccount) | **Get** /sapi/v1/margin/tradeCoeff | Get Summary of Margin account (USER_DATA)
[**QueryCrossIsolatedMarginCapitalFlow**](AccountAPI.md#QueryCrossIsolatedMarginCapitalFlow) | **Get** /sapi/v1/margin/capital-flow | Query Cross Isolated Margin Capital Flow (USER_DATA)
[**QueryCrossMarginAccountDetails**](AccountAPI.md#QueryCrossMarginAccountDetails) | **Get** /sapi/v1/margin/account | Query Cross Margin Account Details (USER_DATA)
[**QueryCrossMarginFeeData**](AccountAPI.md#QueryCrossMarginFeeData) | **Get** /sapi/v1/margin/crossMarginData | Query Cross Margin Fee Data (USER_DATA)
[**QueryEnabledIsolatedMarginAccountLimit**](AccountAPI.md#QueryEnabledIsolatedMarginAccountLimit) | **Get** /sapi/v1/margin/isolated/accountLimit | Query Enabled Isolated Margin Account Limit (USER_DATA)
[**QueryIsolatedMarginAccountInfo**](AccountAPI.md#QueryIsolatedMarginAccountInfo) | **Get** /sapi/v1/margin/isolated/account | Query Isolated Margin Account Info (USER_DATA)
[**QueryIsolatedMarginFeeData**](AccountAPI.md#QueryIsolatedMarginFeeData) | **Get** /sapi/v1/margin/isolatedMarginData | Query Isolated Margin Fee Data (USER_DATA)


## AdjustCrossMarginMaxLeverage

> AdjustCrossMarginMaxLeverageResponse AdjustCrossMarginMaxLeverage(ctx).MaxLeverage(maxLeverage).Execute()

Adjust cross margin max leverage (USER_DATA)


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
	maxLeverage := int64(789) // int64 | Can only adjust 3 , 5 or 10，Example: maxLeverage = 5 or 3 for Cross Margin Classic; maxLeverage=10 for Cross Margin Pro 10x leverage or 20x if compliance allows.

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.AdjustCrossMarginMaxLeverage(context.Background()).MaxLeverage(maxLeverage).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.AdjustCrossMarginMaxLeverage``: %v\n", err)
		return
	}

	// response from `AdjustCrossMarginMaxLeverage`: AdjustCrossMarginMaxLeverageResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **maxLeverage** | **int64** | Can only adjust 3 , 5 or 10，Example: maxLeverage &#x3D; 5 or 3 for Cross Margin Classic; maxLeverage&#x3D;10 for Cross Margin Pro 10x leverage or 20x if compliance allows. | 

### Return type

[**AdjustCrossMarginMaxLeverageResponse**](AdjustCrossMarginMaxLeverageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## DisableIsolatedMarginAccount

> DisableIsolatedMarginAccountResponse DisableIsolatedMarginAccount(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Disable Isolated Margin Account (TRADE)


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
	symbol := "symbol_example" // string | 
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.DisableIsolatedMarginAccount(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.DisableIsolatedMarginAccount``: %v\n", err)
		return
	}

	// response from `DisableIsolatedMarginAccount`: DisableIsolatedMarginAccountResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**DisableIsolatedMarginAccountResponse**](DisableIsolatedMarginAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## EnableIsolatedMarginAccount

> EnableIsolatedMarginAccountResponse EnableIsolatedMarginAccount(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Enable Isolated Margin Account (TRADE)


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
	symbol := "symbol_example" // string | 
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.EnableIsolatedMarginAccount(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.EnableIsolatedMarginAccount``: %v\n", err)
		return
	}

	// response from `EnableIsolatedMarginAccount`: EnableIsolatedMarginAccountResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**EnableIsolatedMarginAccountResponse**](EnableIsolatedMarginAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetBnbBurnStatus

> GetBnbBurnStatusResponse GetBnbBurnStatus(ctx).RecvWindow(recvWindow).Execute()

Get BNB Burn Status (USER_DATA)


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
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetBnbBurnStatus(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetBnbBurnStatus``: %v\n", err)
		return
	}

	// response from `GetBnbBurnStatus`: GetBnbBurnStatusResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**GetBnbBurnStatusResponse**](GetBnbBurnStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetSummaryOfMarginAccount

> GetSummaryOfMarginAccountResponse GetSummaryOfMarginAccount(ctx).RecvWindow(recvWindow).Execute()

Get Summary of Margin account (USER_DATA)


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
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetSummaryOfMarginAccount(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetSummaryOfMarginAccount``: %v\n", err)
		return
	}

	// response from `GetSummaryOfMarginAccount`: GetSummaryOfMarginAccountResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**GetSummaryOfMarginAccountResponse**](GetSummaryOfMarginAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryCrossIsolatedMarginCapitalFlow

> QueryCrossIsolatedMarginCapitalFlowResponse QueryCrossIsolatedMarginCapitalFlow(ctx).Asset(asset).Symbol(symbol).Type(type_).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()

Query Cross Isolated Margin Capital Flow (USER_DATA)


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
	asset := "asset_example" // string |  (optional)
	symbol := "symbol_example" // string | isolated margin pair (optional)
	type_ := "type__example" // string | Transfer Type: ROLL_IN, ROLL_OUT (optional)
	startTime := int64(1623319461670) // int64 | Only supports querying data from the past 90 days. (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	fromId := int64(1) // int64 | If `fromId` is set, data with `id` greater than `fromId` will be returned. Otherwise, the latest data will be returned. (optional)
	limit := int64(500) // int64 | Limit on the number of data records returned per request. Default: 500; Maximum: 1000. (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.QueryCrossIsolatedMarginCapitalFlow(context.Background()).Asset(asset).Symbol(symbol).Type(type_).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.QueryCrossIsolatedMarginCapitalFlow``: %v\n", err)
		return
	}

	// response from `QueryCrossIsolatedMarginCapitalFlow`: QueryCrossIsolatedMarginCapitalFlowResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | 
 **symbol** | **string** | isolated margin pair | 
 **type_** | **string** | Transfer Type: ROLL_IN, ROLL_OUT | 
 **startTime** | **int64** | Only supports querying data from the past 90 days. | 
 **endTime** | **int64** |  | 
 **fromId** | **int64** | If &#x60;fromId&#x60; is set, data with &#x60;id&#x60; greater than &#x60;fromId&#x60; will be returned. Otherwise, the latest data will be returned. | 
 **limit** | **int64** | Limit on the number of data records returned per request. Default: 500; Maximum: 1000. | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**QueryCrossIsolatedMarginCapitalFlowResponse**](QueryCrossIsolatedMarginCapitalFlowResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryCrossMarginAccountDetails

> QueryCrossMarginAccountDetailsResponse QueryCrossMarginAccountDetails(ctx).RecvWindow(recvWindow).Execute()

Query Cross Margin Account Details (USER_DATA)


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
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.QueryCrossMarginAccountDetails(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.QueryCrossMarginAccountDetails``: %v\n", err)
		return
	}

	// response from `QueryCrossMarginAccountDetails`: QueryCrossMarginAccountDetailsResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**QueryCrossMarginAccountDetailsResponse**](QueryCrossMarginAccountDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryCrossMarginFeeData

> QueryCrossMarginFeeDataResponse QueryCrossMarginFeeData(ctx).VipLevel(vipLevel).Coin(coin).RecvWindow(recvWindow).Execute()

Query Cross Margin Fee Data (USER_DATA)


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
	vipLevel := int64(1) // int64 | User's current specific margin data will be returned if vipLevel is omitted (optional)
	coin := "coin_example" // string |  (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.QueryCrossMarginFeeData(context.Background()).VipLevel(vipLevel).Coin(coin).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.QueryCrossMarginFeeData``: %v\n", err)
		return
	}

	// response from `QueryCrossMarginFeeData`: QueryCrossMarginFeeDataResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **vipLevel** | **int64** | User&#39;s current specific margin data will be returned if vipLevel is omitted | 
 **coin** | **string** |  | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**QueryCrossMarginFeeDataResponse**](QueryCrossMarginFeeDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryEnabledIsolatedMarginAccountLimit

> QueryEnabledIsolatedMarginAccountLimitResponse QueryEnabledIsolatedMarginAccountLimit(ctx).RecvWindow(recvWindow).Execute()

Query Enabled Isolated Margin Account Limit (USER_DATA)


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
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.QueryEnabledIsolatedMarginAccountLimit(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.QueryEnabledIsolatedMarginAccountLimit``: %v\n", err)
		return
	}

	// response from `QueryEnabledIsolatedMarginAccountLimit`: QueryEnabledIsolatedMarginAccountLimitResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**QueryEnabledIsolatedMarginAccountLimitResponse**](QueryEnabledIsolatedMarginAccountLimitResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryIsolatedMarginAccountInfo

> QueryIsolatedMarginAccountInfoResponse QueryIsolatedMarginAccountInfo(ctx).Symbols(symbols).RecvWindow(recvWindow).Execute()

Query Isolated Margin Account Info (USER_DATA)


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
	symbols := "symbols_example" // string | Max 5 symbols can be sent; separated by \",\". e.g. \"BTCUSDT,BNBUSDT,ADAUSDT\" (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.QueryIsolatedMarginAccountInfo(context.Background()).Symbols(symbols).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.QueryIsolatedMarginAccountInfo``: %v\n", err)
		return
	}

	// response from `QueryIsolatedMarginAccountInfo`: QueryIsolatedMarginAccountInfoResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbols** | **string** | Max 5 symbols can be sent; separated by \&quot;,\&quot;. e.g. \&quot;BTCUSDT,BNBUSDT,ADAUSDT\&quot; | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**QueryIsolatedMarginAccountInfoResponse**](QueryIsolatedMarginAccountInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryIsolatedMarginFeeData

> QueryIsolatedMarginFeeDataResponse QueryIsolatedMarginFeeData(ctx).VipLevel(vipLevel).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query Isolated Margin Fee Data (USER_DATA)


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
	vipLevel := int64(1) // int64 | User's current specific margin data will be returned if vipLevel is omitted (optional)
	symbol := "symbol_example" // string | isolated margin pair (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.QueryIsolatedMarginFeeData(context.Background()).VipLevel(vipLevel).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.QueryIsolatedMarginFeeData``: %v\n", err)
		return
	}

	// response from `QueryIsolatedMarginFeeData`: QueryIsolatedMarginFeeDataResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **vipLevel** | **int64** | User&#39;s current specific margin data will be returned if vipLevel is omitted | 
 **symbol** | **string** | isolated margin pair | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**QueryIsolatedMarginFeeDataResponse**](QueryIsolatedMarginFeeDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

