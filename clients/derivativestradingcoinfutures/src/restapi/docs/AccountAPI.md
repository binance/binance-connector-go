# \AccountAPI

All URIs are relative to *https://dapi.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AccountInformation**](AccountAPI.md#AccountInformation) | **Get** /dapi/v1/account | Account Information (USER_DATA)
[**FuturesAccountBalance**](AccountAPI.md#FuturesAccountBalance) | **Get** /dapi/v1/balance | Futures Account Balance (USER_DATA)
[**GetCurrentPositionMode**](AccountAPI.md#GetCurrentPositionMode) | **Get** /dapi/v1/positionSide/dual | Get Current Position Mode(USER_DATA)
[**GetDownloadIdForFuturesOrderHistory**](AccountAPI.md#GetDownloadIdForFuturesOrderHistory) | **Get** /dapi/v1/order/asyn | Get Download Id For Futures Order History (USER_DATA)
[**GetDownloadIdForFuturesTradeHistory**](AccountAPI.md#GetDownloadIdForFuturesTradeHistory) | **Get** /dapi/v1/trade/asyn | Get Download Id For Futures Trade History (USER_DATA)
[**GetDownloadIdForFuturesTransactionHistory**](AccountAPI.md#GetDownloadIdForFuturesTransactionHistory) | **Get** /dapi/v1/income/asyn | Get Download Id For Futures Transaction History(USER_DATA)
[**GetFuturesOrderHistoryDownloadLinkById**](AccountAPI.md#GetFuturesOrderHistoryDownloadLinkById) | **Get** /dapi/v1/order/asyn/id | Get Futures Order History Download Link by Id (USER_DATA)
[**GetFuturesTradeDownloadLinkById**](AccountAPI.md#GetFuturesTradeDownloadLinkById) | **Get** /dapi/v1/trade/asyn/id | Get Futures Trade Download Link by Id(USER_DATA)
[**GetFuturesTransactionHistoryDownloadLinkById**](AccountAPI.md#GetFuturesTransactionHistoryDownloadLinkById) | **Get** /dapi/v1/income/asyn/id | Get Futures Transaction History Download Link by Id (USER_DATA)
[**GetIncomeHistory**](AccountAPI.md#GetIncomeHistory) | **Get** /dapi/v1/income | Get Income History(USER_DATA)
[**NotionalBracketForPair**](AccountAPI.md#NotionalBracketForPair) | **Get** /dapi/v1/leverageBracket | Notional Bracket for Pair(USER_DATA)
[**NotionalBracketForSymbol**](AccountAPI.md#NotionalBracketForSymbol) | **Get** /dapi/v2/leverageBracket | Notional Bracket for Symbol(USER_DATA)
[**UserCommissionRate**](AccountAPI.md#UserCommissionRate) | **Get** /dapi/v1/commissionRate | User Commission Rate (USER_DATA)


## AccountInformation

> AccountInformationResponse AccountInformation(ctx).RecvWindow(recvWindow).Execute()

Account Information (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.AccountInformation(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.AccountInformation``: %v\n", err)
		return
	}

	// response from `AccountInformation`: AccountInformationResponse
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

[**AccountInformationResponse**](AccountInformationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## FuturesAccountBalance

> FuturesAccountBalanceResponse FuturesAccountBalance(ctx).RecvWindow(recvWindow).Execute()

Futures Account Balance (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.FuturesAccountBalance(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.FuturesAccountBalance``: %v\n", err)
		return
	}

	// response from `FuturesAccountBalance`: FuturesAccountBalanceResponse
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

[**FuturesAccountBalanceResponse**](FuturesAccountBalanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetCurrentPositionMode

> GetCurrentPositionModeResponse GetCurrentPositionMode(ctx).RecvWindow(recvWindow).Execute()

Get Current Position Mode(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetCurrentPositionMode(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetCurrentPositionMode``: %v\n", err)
		return
	}

	// response from `GetCurrentPositionMode`: GetCurrentPositionModeResponse
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

[**GetCurrentPositionModeResponse**](GetCurrentPositionModeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetDownloadIdForFuturesOrderHistory

> GetDownloadIdForFuturesOrderHistoryResponse GetDownloadIdForFuturesOrderHistory(ctx).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

Get Download Id For Futures Order History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	startTime := int64(1623319461670) // int64 | Timestamp in ms
	endTime := int64(1641782889000) // int64 | Timestamp in ms
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetDownloadIdForFuturesOrderHistory(context.Background()).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetDownloadIdForFuturesOrderHistory``: %v\n", err)
		return
	}

	// response from `GetDownloadIdForFuturesOrderHistory`: GetDownloadIdForFuturesOrderHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** | Timestamp in ms | 
 **endTime** | **int64** | Timestamp in ms | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetDownloadIdForFuturesOrderHistoryResponse**](GetDownloadIdForFuturesOrderHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetDownloadIdForFuturesTradeHistory

> GetDownloadIdForFuturesTradeHistoryResponse GetDownloadIdForFuturesTradeHistory(ctx).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

Get Download Id For Futures Trade History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	startTime := int64(1623319461670) // int64 | Timestamp in ms
	endTime := int64(1641782889000) // int64 | Timestamp in ms
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetDownloadIdForFuturesTradeHistory(context.Background()).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetDownloadIdForFuturesTradeHistory``: %v\n", err)
		return
	}

	// response from `GetDownloadIdForFuturesTradeHistory`: GetDownloadIdForFuturesTradeHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** | Timestamp in ms | 
 **endTime** | **int64** | Timestamp in ms | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetDownloadIdForFuturesTradeHistoryResponse**](GetDownloadIdForFuturesTradeHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetDownloadIdForFuturesTransactionHistory

> GetDownloadIdForFuturesTransactionHistoryResponse GetDownloadIdForFuturesTransactionHistory(ctx).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

Get Download Id For Futures Transaction History(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	startTime := int64(1623319461670) // int64 | Timestamp in ms
	endTime := int64(1641782889000) // int64 | Timestamp in ms
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetDownloadIdForFuturesTransactionHistory(context.Background()).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetDownloadIdForFuturesTransactionHistory``: %v\n", err)
		return
	}

	// response from `GetDownloadIdForFuturesTransactionHistory`: GetDownloadIdForFuturesTransactionHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** | Timestamp in ms | 
 **endTime** | **int64** | Timestamp in ms | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetDownloadIdForFuturesTransactionHistoryResponse**](GetDownloadIdForFuturesTransactionHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetFuturesOrderHistoryDownloadLinkById

> GetFuturesOrderHistoryDownloadLinkByIdResponse GetFuturesOrderHistoryDownloadLinkById(ctx).DownloadId(downloadId).RecvWindow(recvWindow).Execute()

Get Futures Order History Download Link by Id (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	downloadId := "1" // string | get by download id api
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetFuturesOrderHistoryDownloadLinkById(context.Background()).DownloadId(downloadId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetFuturesOrderHistoryDownloadLinkById``: %v\n", err)
		return
	}

	// response from `GetFuturesOrderHistoryDownloadLinkById`: GetFuturesOrderHistoryDownloadLinkByIdResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **downloadId** | **string** | get by download id api | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetFuturesOrderHistoryDownloadLinkByIdResponse**](GetFuturesOrderHistoryDownloadLinkByIdResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetFuturesTradeDownloadLinkById

> GetFuturesTradeDownloadLinkByIdResponse GetFuturesTradeDownloadLinkById(ctx).DownloadId(downloadId).RecvWindow(recvWindow).Execute()

Get Futures Trade Download Link by Id(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	downloadId := "1" // string | get by download id api
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetFuturesTradeDownloadLinkById(context.Background()).DownloadId(downloadId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetFuturesTradeDownloadLinkById``: %v\n", err)
		return
	}

	// response from `GetFuturesTradeDownloadLinkById`: GetFuturesTradeDownloadLinkByIdResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **downloadId** | **string** | get by download id api | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetFuturesTradeDownloadLinkByIdResponse**](GetFuturesTradeDownloadLinkByIdResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetFuturesTransactionHistoryDownloadLinkById

> GetFuturesTransactionHistoryDownloadLinkByIdResponse GetFuturesTransactionHistoryDownloadLinkById(ctx).DownloadId(downloadId).RecvWindow(recvWindow).Execute()

Get Futures Transaction History Download Link by Id (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	downloadId := "1" // string | get by download id api
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetFuturesTransactionHistoryDownloadLinkById(context.Background()).DownloadId(downloadId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetFuturesTransactionHistoryDownloadLinkById``: %v\n", err)
		return
	}

	// response from `GetFuturesTransactionHistoryDownloadLinkById`: GetFuturesTransactionHistoryDownloadLinkByIdResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **downloadId** | **string** | get by download id api | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetFuturesTransactionHistoryDownloadLinkByIdResponse**](GetFuturesTransactionHistoryDownloadLinkByIdResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetIncomeHistory

> GetIncomeHistoryResponse GetIncomeHistory(ctx).Symbol(symbol).IncomeType(incomeType).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Get Income History(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string |  (optional)
	incomeType := "incomeType_example" // string | \"TRANSFER\",\"WELCOME_BONUS\", \"FUNDING_FEE\", \"REALIZED_PNL\", \"COMMISSION\", \"INSURANCE_CLEAR\", and \"DELIVERED_SETTELMENT\" (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	page := int64(789) // int64 |  (optional)
	limit := int64(100) // int64 | Default 100; max 1000 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetIncomeHistory(context.Background()).Symbol(symbol).IncomeType(incomeType).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetIncomeHistory``: %v\n", err)
		return
	}

	// response from `GetIncomeHistory`: GetIncomeHistoryResponse
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
 **incomeType** | **string** | \&quot;TRANSFER\&quot;,\&quot;WELCOME_BONUS\&quot;, \&quot;FUNDING_FEE\&quot;, \&quot;REALIZED_PNL\&quot;, \&quot;COMMISSION\&quot;, \&quot;INSURANCE_CLEAR\&quot;, and \&quot;DELIVERED_SETTELMENT\&quot; | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **page** | **int64** |  | 
 **limit** | **int64** | Default 100; max 1000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetIncomeHistoryResponse**](GetIncomeHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## NotionalBracketForPair

> NotionalBracketForPairResponse NotionalBracketForPair(ctx).Pair(pair).RecvWindow(recvWindow).Execute()

Notional Bracket for Pair(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	pair := "pair_example" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.NotionalBracketForPair(context.Background()).Pair(pair).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.NotionalBracketForPair``: %v\n", err)
		return
	}

	// response from `NotionalBracketForPair`: NotionalBracketForPairResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **pair** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**NotionalBracketForPairResponse**](NotionalBracketForPairResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## NotionalBracketForSymbol

> NotionalBracketForSymbolResponse NotionalBracketForSymbol(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Notional Bracket for Symbol(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.NotionalBracketForSymbol(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.NotionalBracketForSymbol``: %v\n", err)
		return
	}

	// response from `NotionalBracketForSymbol`: NotionalBracketForSymbolResponse
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
 **recvWindow** | **int64** |  | 

### Return type

[**NotionalBracketForSymbolResponse**](NotionalBracketForSymbolResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## UserCommissionRate

> UserCommissionRateResponse UserCommissionRate(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

User Commission Rate (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.UserCommissionRate(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.UserCommissionRate``: %v\n", err)
		return
	}

	// response from `UserCommissionRate`: UserCommissionRateResponse
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
 **recvWindow** | **int64** |  | 

### Return type

[**UserCommissionRateResponse**](UserCommissionRateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

