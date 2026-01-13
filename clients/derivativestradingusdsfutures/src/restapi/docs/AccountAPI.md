# \AccountAPI

All URIs are relative to *https://fapi.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AccountInformationV2**](AccountAPI.md#AccountInformationV2) | **Get** /fapi/v2/account | Account Information V2(USER_DATA)
[**AccountInformationV3**](AccountAPI.md#AccountInformationV3) | **Get** /fapi/v3/account | Account Information V3(USER_DATA)
[**FuturesAccountBalanceV2**](AccountAPI.md#FuturesAccountBalanceV2) | **Get** /fapi/v2/balance | Futures Account Balance V2 (USER_DATA)
[**FuturesAccountBalanceV3**](AccountAPI.md#FuturesAccountBalanceV3) | **Get** /fapi/v3/balance | Futures Account Balance V3 (USER_DATA)
[**FuturesAccountConfiguration**](AccountAPI.md#FuturesAccountConfiguration) | **Get** /fapi/v1/accountConfig | Futures Account Configuration(USER_DATA)
[**FuturesTradingQuantitativeRulesIndicators**](AccountAPI.md#FuturesTradingQuantitativeRulesIndicators) | **Get** /fapi/v1/apiTradingStatus | Futures Trading Quantitative Rules Indicators (USER_DATA)
[**GetBnbBurnStatus**](AccountAPI.md#GetBnbBurnStatus) | **Get** /fapi/v1/feeBurn | Get BNB Burn Status (USER_DATA)
[**GetCurrentMultiAssetsMode**](AccountAPI.md#GetCurrentMultiAssetsMode) | **Get** /fapi/v1/multiAssetsMargin | Get Current Multi-Assets Mode (USER_DATA)
[**GetCurrentPositionMode**](AccountAPI.md#GetCurrentPositionMode) | **Get** /fapi/v1/positionSide/dual | Get Current Position Mode(USER_DATA)
[**GetDownloadIdForFuturesOrderHistory**](AccountAPI.md#GetDownloadIdForFuturesOrderHistory) | **Get** /fapi/v1/order/asyn | Get Download Id For Futures Order History (USER_DATA)
[**GetDownloadIdForFuturesTradeHistory**](AccountAPI.md#GetDownloadIdForFuturesTradeHistory) | **Get** /fapi/v1/trade/asyn | Get Download Id For Futures Trade History (USER_DATA)
[**GetDownloadIdForFuturesTransactionHistory**](AccountAPI.md#GetDownloadIdForFuturesTransactionHistory) | **Get** /fapi/v1/income/asyn | Get Download Id For Futures Transaction History(USER_DATA)
[**GetFuturesOrderHistoryDownloadLinkById**](AccountAPI.md#GetFuturesOrderHistoryDownloadLinkById) | **Get** /fapi/v1/order/asyn/id | Get Futures Order History Download Link by Id (USER_DATA)
[**GetFuturesTradeDownloadLinkById**](AccountAPI.md#GetFuturesTradeDownloadLinkById) | **Get** /fapi/v1/trade/asyn/id | Get Futures Trade Download Link by Id(USER_DATA)
[**GetFuturesTransactionHistoryDownloadLinkById**](AccountAPI.md#GetFuturesTransactionHistoryDownloadLinkById) | **Get** /fapi/v1/income/asyn/id | Get Futures Transaction History Download Link by Id (USER_DATA)
[**GetIncomeHistory**](AccountAPI.md#GetIncomeHistory) | **Get** /fapi/v1/income | Get Income History (USER_DATA)
[**NotionalAndLeverageBrackets**](AccountAPI.md#NotionalAndLeverageBrackets) | **Get** /fapi/v1/leverageBracket | Notional and Leverage Brackets (USER_DATA)
[**QueryUserRateLimit**](AccountAPI.md#QueryUserRateLimit) | **Get** /fapi/v1/rateLimit/order | Query User Rate Limit (USER_DATA)
[**SymbolConfiguration**](AccountAPI.md#SymbolConfiguration) | **Get** /fapi/v1/symbolConfig | Symbol Configuration(USER_DATA)
[**ToggleBnbBurnOnFuturesTrade**](AccountAPI.md#ToggleBnbBurnOnFuturesTrade) | **Post** /fapi/v1/feeBurn | Toggle BNB Burn On Futures Trade (TRADE)
[**UserCommissionRate**](AccountAPI.md#UserCommissionRate) | **Get** /fapi/v1/commissionRate | User Commission Rate (USER_DATA)


## AccountInformationV2

> AccountInformationV2Response AccountInformationV2(ctx).RecvWindow(recvWindow).Execute()

Account Information V2(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.AccountInformationV2(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.AccountInformationV2``: %v\n", err)
		return
	}

	// response from `AccountInformationV2`: AccountInformationV2Response
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

[**AccountInformationV2Response**](AccountInformationV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## AccountInformationV3

> AccountInformationV3Response AccountInformationV3(ctx).RecvWindow(recvWindow).Execute()

Account Information V3(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.AccountInformationV3(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.AccountInformationV3``: %v\n", err)
		return
	}

	// response from `AccountInformationV3`: AccountInformationV3Response
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

[**AccountInformationV3Response**](AccountInformationV3Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## FuturesAccountBalanceV2

> FuturesAccountBalanceV2Response FuturesAccountBalanceV2(ctx).RecvWindow(recvWindow).Execute()

Futures Account Balance V2 (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.FuturesAccountBalanceV2(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.FuturesAccountBalanceV2``: %v\n", err)
		return
	}

	// response from `FuturesAccountBalanceV2`: FuturesAccountBalanceV2Response
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

[**FuturesAccountBalanceV2Response**](FuturesAccountBalanceV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## FuturesAccountBalanceV3

> FuturesAccountBalanceV3Response FuturesAccountBalanceV3(ctx).RecvWindow(recvWindow).Execute()

Futures Account Balance V3 (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.FuturesAccountBalanceV3(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.FuturesAccountBalanceV3``: %v\n", err)
		return
	}

	// response from `FuturesAccountBalanceV3`: FuturesAccountBalanceV3Response
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

[**FuturesAccountBalanceV3Response**](FuturesAccountBalanceV3Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## FuturesAccountConfiguration

> FuturesAccountConfigurationResponse FuturesAccountConfiguration(ctx).RecvWindow(recvWindow).Execute()

Futures Account Configuration(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.FuturesAccountConfiguration(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.FuturesAccountConfiguration``: %v\n", err)
		return
	}

	// response from `FuturesAccountConfiguration`: FuturesAccountConfigurationResponse
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

[**FuturesAccountConfigurationResponse**](FuturesAccountConfigurationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## FuturesTradingQuantitativeRulesIndicators

> FuturesTradingQuantitativeRulesIndicatorsResponse FuturesTradingQuantitativeRulesIndicators(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Futures Trading Quantitative Rules Indicators (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.FuturesTradingQuantitativeRulesIndicators(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.FuturesTradingQuantitativeRulesIndicators``: %v\n", err)
		return
	}

	// response from `FuturesTradingQuantitativeRulesIndicators`: FuturesTradingQuantitativeRulesIndicatorsResponse
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

[**FuturesTradingQuantitativeRulesIndicatorsResponse**](FuturesTradingQuantitativeRulesIndicatorsResponse.md)

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

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
 **recvWindow** | **int64** |  | 

### Return type

[**GetBnbBurnStatusResponse**](GetBnbBurnStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetCurrentMultiAssetsMode

> GetCurrentMultiAssetsModeResponse GetCurrentMultiAssetsMode(ctx).RecvWindow(recvWindow).Execute()

Get Current Multi-Assets Mode (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetCurrentMultiAssetsMode(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetCurrentMultiAssetsMode``: %v\n", err)
		return
	}

	// response from `GetCurrentMultiAssetsMode`: GetCurrentMultiAssetsModeResponse
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

[**GetCurrentMultiAssetsModeResponse**](GetCurrentMultiAssetsModeResponse.md)

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/common"
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
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/common"
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
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/common"
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
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	downloadId := "1" // string | get by download id api
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	downloadId := "1" // string | get by download id api
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	downloadId := "1" // string | get by download id api
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

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

Get Income History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string |  (optional)
	incomeType := "incomeType_example" // string | TRANSFER, WELCOME_BONUS, REALIZED_PNL, FUNDING_FEE, COMMISSION, INSURANCE_CLEAR, REFERRAL_KICKBACK, COMMISSION_REBATE, API_REBATE, CONTEST_REWARD, CROSS_COLLATERAL_TRANSFER, OPTIONS_PREMIUM_FEE, OPTIONS_SETTLE_PROFIT, INTERNAL_TRANSFER, AUTO_EXCHANGE, DELIVERED_SETTELMENT, COIN_SWAP_DEPOSIT, COIN_SWAP_WITHDRAW, POSITION_LIMIT_INCREASE_FEE, STRATEGY_UMFUTURES_TRANSFER，FEE_RETURN，BFUSD_REWARD (optional)
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
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

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
 **incomeType** | **string** | TRANSFER, WELCOME_BONUS, REALIZED_PNL, FUNDING_FEE, COMMISSION, INSURANCE_CLEAR, REFERRAL_KICKBACK, COMMISSION_REBATE, API_REBATE, CONTEST_REWARD, CROSS_COLLATERAL_TRANSFER, OPTIONS_PREMIUM_FEE, OPTIONS_SETTLE_PROFIT, INTERNAL_TRANSFER, AUTO_EXCHANGE, DELIVERED_SETTELMENT, COIN_SWAP_DEPOSIT, COIN_SWAP_WITHDRAW, POSITION_LIMIT_INCREASE_FEE, STRATEGY_UMFUTURES_TRANSFER，FEE_RETURN，BFUSD_REWARD | 
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


## NotionalAndLeverageBrackets

> NotionalAndLeverageBracketsResponse NotionalAndLeverageBrackets(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Notional and Leverage Brackets (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.NotionalAndLeverageBrackets(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.NotionalAndLeverageBrackets``: %v\n", err)
		return
	}

	// response from `NotionalAndLeverageBrackets`: NotionalAndLeverageBracketsResponse
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

[**NotionalAndLeverageBracketsResponse**](NotionalAndLeverageBracketsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryUserRateLimit

> QueryUserRateLimitResponse QueryUserRateLimit(ctx).RecvWindow(recvWindow).Execute()

Query User Rate Limit (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.QueryUserRateLimit(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.QueryUserRateLimit``: %v\n", err)
		return
	}

	// response from `QueryUserRateLimit`: QueryUserRateLimitResponse
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

[**QueryUserRateLimitResponse**](QueryUserRateLimitResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SymbolConfiguration

> SymbolConfigurationResponse SymbolConfiguration(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Symbol Configuration(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.SymbolConfiguration(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.SymbolConfiguration``: %v\n", err)
		return
	}

	// response from `SymbolConfiguration`: SymbolConfigurationResponse
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

[**SymbolConfigurationResponse**](SymbolConfigurationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## ToggleBnbBurnOnFuturesTrade

> ToggleBnbBurnOnFuturesTradeResponse ToggleBnbBurnOnFuturesTrade(ctx).FeeBurn(feeBurn).RecvWindow(recvWindow).Execute()

Toggle BNB Burn On Futures Trade (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	feeBurn := "feeBurn_example" // string | \"true\": Fee Discount On; \"false\": Fee Discount Off
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.ToggleBnbBurnOnFuturesTrade(context.Background()).FeeBurn(feeBurn).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.ToggleBnbBurnOnFuturesTrade``: %v\n", err)
		return
	}

	// response from `ToggleBnbBurnOnFuturesTrade`: ToggleBnbBurnOnFuturesTradeResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **feeBurn** | **string** | \&quot;true\&quot;: Fee Discount On; \&quot;false\&quot;: Fee Discount Off | 
 **recvWindow** | **int64** |  | 

### Return type

[**ToggleBnbBurnOnFuturesTradeResponse**](ToggleBnbBurnOnFuturesTradeResponse.md)

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithRestAPI(configuration))

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

