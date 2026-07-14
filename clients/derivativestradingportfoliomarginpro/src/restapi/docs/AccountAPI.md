# \AccountAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**BnbTransfer**](AccountAPI.md#BnbTransfer) | **Post** /sapi/v1/portfolio/bnb-transfer | BNB transfer (USER_DATA)
[**ChangeAutoRepayFuturesStatus**](AccountAPI.md#ChangeAutoRepayFuturesStatus) | **Post** /sapi/v1/portfolio/repay-futures-switch | Change Auto-repay-futures Status (TRADE)
[**DeleteMarginCallLevel**](AccountAPI.md#DeleteMarginCallLevel) | **Delete** /sapi/v1/portfolio/margin-call-level | Delete Margin Call Level (USER_DATA)
[**FundAutoCollection**](AccountAPI.md#FundAutoCollection) | **Post** /sapi/v1/portfolio/auto-collection | Fund Auto-collection (USER_DATA)
[**FundCollectionByAsset**](AccountAPI.md#FundCollectionByAsset) | **Post** /sapi/v1/portfolio/asset-collection | Fund Collection by Asset (USER_DATA)
[**GetAutoRepayFuturesStatus**](AccountAPI.md#GetAutoRepayFuturesStatus) | **Get** /sapi/v1/portfolio/repay-futures-switch | Get Auto-repay-futures Status (USER_DATA)
[**GetDeltaModeStatus**](AccountAPI.md#GetDeltaModeStatus) | **Get** /sapi/v1/portfolio/delta-mode | Get Delta Mode Status (USER_DATA)
[**GetMarginCallLevel**](AccountAPI.md#GetMarginCallLevel) | **Get** /sapi/v1/portfolio/margin-call-level | Get Margin Call Level (USER_DATA)
[**GetPortfolioMarginProAccountBalance**](AccountAPI.md#GetPortfolioMarginProAccountBalance) | **Get** /sapi/v1/portfolio/balance | Get Portfolio Margin Pro Account Balance (USER_DATA)
[**GetPortfolioMarginProAccountInfo**](AccountAPI.md#GetPortfolioMarginProAccountInfo) | **Get** /sapi/v1/portfolio/account | Get Portfolio Margin Pro Account Info (USER_DATA)
[**GetPortfolioMarginProSpanAccountInfo**](AccountAPI.md#GetPortfolioMarginProSpanAccountInfo) | **Get** /sapi/v2/portfolio/account | Get Portfolio Margin Pro SPAN Account Info (USER_DATA)
[**GetTransferableEarnAssetBalanceForPortfolioMargin**](AccountAPI.md#GetTransferableEarnAssetBalanceForPortfolioMargin) | **Get** /sapi/v1/portfolio/earn-asset-balance | Get Transferable Earn Asset Balance for Portfolio Margin (USER_DATA)
[**PortfolioMarginProBankruptcyLoanRepay**](AccountAPI.md#PortfolioMarginProBankruptcyLoanRepay) | **Post** /sapi/v1/portfolio/repay | Portfolio Margin Pro Bankruptcy Loan Repay (TRADE)
[**QueryPortfolioMarginProBankruptcyLoanAmount**](AccountAPI.md#QueryPortfolioMarginProBankruptcyLoanAmount) | **Get** /sapi/v1/portfolio/pmLoan | Query Portfolio Margin Pro Bankruptcy Loan Amount (USER_DATA)
[**QueryPortfolioMarginProBankruptcyLoanRepayHistory**](AccountAPI.md#QueryPortfolioMarginProBankruptcyLoanRepayHistory) | **Get** /sapi/v1/portfolio/pmloan-history | Query Portfolio Margin Pro Bankruptcy Loan Repay History (USER_DATA)
[**QueryPortfolioMarginProNegativeBalanceInterestHistory**](AccountAPI.md#QueryPortfolioMarginProNegativeBalanceInterestHistory) | **Get** /sapi/v1/portfolio/interest-history | Query Portfolio Margin Pro Negative Balance Interest History (USER_DATA)
[**RepayFuturesNegativeBalance**](AccountAPI.md#RepayFuturesNegativeBalance) | **Post** /sapi/v1/portfolio/repay-futures-negative-balance | Repay futures Negative Balance (USER_DATA)
[**SetMarginCallLevel**](AccountAPI.md#SetMarginCallLevel) | **Post** /sapi/v1/portfolio/margin-call-level | Set Margin Call Level (USER_DATA)
[**SwitchDeltaMode**](AccountAPI.md#SwitchDeltaMode) | **Post** /sapi/v1/portfolio/delta-mode | Switch Delta Mode (TRADE)
[**TransferLdusdtRwusdForPortfolioMargin**](AccountAPI.md#TransferLdusdtRwusdForPortfolioMargin) | **Post** /sapi/v1/portfolio/earn-asset-transfer | Transfer LDUSDT/RWUSD for Portfolio Margin (TRADE)


## BnbTransfer

> BnbTransferResponse BnbTransfer(ctx).Amount(amount).TransferSide(transferSide).RecvWindow(recvWindow).Execute()

BNB transfer (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomarginpro"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	amount := float32(1.0) // float32 | 
	transferSide := models.BnbTransferTransferSideParameterToUm // BnbTransferTransferSideParameter | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginProClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.BnbTransfer(context.Background()).Amount(amount).TransferSide(transferSide).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.BnbTransfer``: %v\n", err)
		return
	}

	// response from `BnbTransfer`: BnbTransferResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **float32** |  | 
 **transferSide** | [**BnbTransferTransferSideParameter**](BnbTransferTransferSideParameter.md) |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**BnbTransferResponse**](BnbTransferResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## ChangeAutoRepayFuturesStatus

> ChangeAutoRepayFuturesStatusResponse ChangeAutoRepayFuturesStatus(ctx).AutoRepay(autoRepay).RecvWindow(recvWindow).Execute()

Change Auto-repay-futures Status (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomarginpro"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	autoRepay := models.ChangeAutoRepayFuturesStatusAutoRepayParameterTrue // ChangeAutoRepayFuturesStatusAutoRepayParameter | `false` for turn off the auto-repay futures negative balance function
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginProClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.ChangeAutoRepayFuturesStatus(context.Background()).AutoRepay(autoRepay).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.ChangeAutoRepayFuturesStatus``: %v\n", err)
		return
	}

	// response from `ChangeAutoRepayFuturesStatus`: ChangeAutoRepayFuturesStatusResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **autoRepay** | [**ChangeAutoRepayFuturesStatusAutoRepayParameter**](ChangeAutoRepayFuturesStatusAutoRepayParameter.md) | &#x60;false&#x60; for turn off the auto-repay futures negative balance function | 
 **recvWindow** | **int64** |  | 

### Return type

[**ChangeAutoRepayFuturesStatusResponse**](ChangeAutoRepayFuturesStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## DeleteMarginCallLevel

> DeleteMarginCallLevelResponse DeleteMarginCallLevel(ctx).RecvWindow(recvWindow).Execute()

Delete Margin Call Level (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomarginpro"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginProClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.DeleteMarginCallLevel(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.DeleteMarginCallLevel``: %v\n", err)
		return
	}

	// response from `DeleteMarginCallLevel`: DeleteMarginCallLevelResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**DeleteMarginCallLevelResponse**](DeleteMarginCallLevelResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## FundAutoCollection

> FundAutoCollectionResponse FundAutoCollection(ctx).RecvWindow(recvWindow).Execute()

Fund Auto-collection (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomarginpro"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginProClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.FundAutoCollection(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.FundAutoCollection``: %v\n", err)
		return
	}

	// response from `FundAutoCollection`: FundAutoCollectionResponse
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

[**FundAutoCollectionResponse**](FundAutoCollectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## FundCollectionByAsset

> FundCollectionByAssetResponse FundCollectionByAsset(ctx).Asset(asset).RecvWindow(recvWindow).Execute()

Fund Collection by Asset (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomarginpro"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	asset := "USDT" // string | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginProClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.FundCollectionByAsset(context.Background()).Asset(asset).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.FundCollectionByAsset``: %v\n", err)
		return
	}

	// response from `FundCollectionByAsset`: FundCollectionByAssetResponse
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
 **recvWindow** | **int64** |  | 

### Return type

[**FundCollectionByAssetResponse**](FundCollectionByAssetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetAutoRepayFuturesStatus

> GetAutoRepayFuturesStatusResponse GetAutoRepayFuturesStatus(ctx).RecvWindow(recvWindow).Execute()

Get Auto-repay-futures Status (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomarginpro"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginProClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetAutoRepayFuturesStatus(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetAutoRepayFuturesStatus``: %v\n", err)
		return
	}

	// response from `GetAutoRepayFuturesStatus`: GetAutoRepayFuturesStatusResponse
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

[**GetAutoRepayFuturesStatusResponse**](GetAutoRepayFuturesStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetDeltaModeStatus

> GetDeltaModeStatusResponse GetDeltaModeStatus(ctx).RecvWindow(recvWindow).Execute()

Get Delta Mode Status (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomarginpro"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginProClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetDeltaModeStatus(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetDeltaModeStatus``: %v\n", err)
		return
	}

	// response from `GetDeltaModeStatus`: GetDeltaModeStatusResponse
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

[**GetDeltaModeStatusResponse**](GetDeltaModeStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetMarginCallLevel

> GetMarginCallLevelResponse GetMarginCallLevel(ctx).RecvWindow(recvWindow).Execute()

Get Margin Call Level (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomarginpro"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginProClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetMarginCallLevel(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetMarginCallLevel``: %v\n", err)
		return
	}

	// response from `GetMarginCallLevel`: GetMarginCallLevelResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**GetMarginCallLevelResponse**](GetMarginCallLevelResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetPortfolioMarginProAccountBalance

> GetPortfolioMarginProAccountBalanceResponse GetPortfolioMarginProAccountBalance(ctx).Asset(asset).RecvWindow(recvWindow).Execute()

Get Portfolio Margin Pro Account Balance (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomarginpro"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	asset := "BTC" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginProClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetPortfolioMarginProAccountBalance(context.Background()).Asset(asset).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetPortfolioMarginProAccountBalance``: %v\n", err)
		return
	}

	// response from `GetPortfolioMarginProAccountBalance`: GetPortfolioMarginProAccountBalanceResponse
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
 **recvWindow** | **int64** |  | 

### Return type

[**GetPortfolioMarginProAccountBalanceResponse**](GetPortfolioMarginProAccountBalanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetPortfolioMarginProAccountInfo

> GetPortfolioMarginProAccountInfoResponse GetPortfolioMarginProAccountInfo(ctx).RecvWindow(recvWindow).Execute()

Get Portfolio Margin Pro Account Info (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomarginpro"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginProClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetPortfolioMarginProAccountInfo(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetPortfolioMarginProAccountInfo``: %v\n", err)
		return
	}

	// response from `GetPortfolioMarginProAccountInfo`: GetPortfolioMarginProAccountInfoResponse
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

[**GetPortfolioMarginProAccountInfoResponse**](GetPortfolioMarginProAccountInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetPortfolioMarginProSpanAccountInfo

> GetPortfolioMarginProSpanAccountInfoResponse GetPortfolioMarginProSpanAccountInfo(ctx).RecvWindow(recvWindow).Execute()

Get Portfolio Margin Pro SPAN Account Info (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomarginpro"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginProClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetPortfolioMarginProSpanAccountInfo(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetPortfolioMarginProSpanAccountInfo``: %v\n", err)
		return
	}

	// response from `GetPortfolioMarginProSpanAccountInfo`: GetPortfolioMarginProSpanAccountInfoResponse
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

[**GetPortfolioMarginProSpanAccountInfoResponse**](GetPortfolioMarginProSpanAccountInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetTransferableEarnAssetBalanceForPortfolioMargin

> GetTransferableEarnAssetBalanceForPortfolioMarginResponse GetTransferableEarnAssetBalanceForPortfolioMargin(ctx).Asset(asset).TransferType(transferType).RecvWindow(recvWindow).Execute()

Get Transferable Earn Asset Balance for Portfolio Margin (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomarginpro"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	asset := "LDUSDT" // string | `LDUSDT` only
	transferType := models.GetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameterEarnToFuture // GetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginProClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetTransferableEarnAssetBalanceForPortfolioMargin(context.Background()).Asset(asset).TransferType(transferType).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetTransferableEarnAssetBalanceForPortfolioMargin``: %v\n", err)
		return
	}

	// response from `GetTransferableEarnAssetBalanceForPortfolioMargin`: GetTransferableEarnAssetBalanceForPortfolioMarginResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** | &#x60;LDUSDT&#x60; only | 
 **transferType** | [**GetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter**](GetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter.md) |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetTransferableEarnAssetBalanceForPortfolioMarginResponse**](GetTransferableEarnAssetBalanceForPortfolioMarginResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## PortfolioMarginProBankruptcyLoanRepay

> PortfolioMarginProBankruptcyLoanRepayResponse PortfolioMarginProBankruptcyLoanRepay(ctx).From(from).RecvWindow(recvWindow).Execute()

Portfolio Margin Pro Bankruptcy Loan Repay (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomarginpro"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	from := models.PortfolioMarginProBankruptcyLoanRepayFromParameterSpot // PortfolioMarginProBankruptcyLoanRepayFromParameter |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginProClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.PortfolioMarginProBankruptcyLoanRepay(context.Background()).From(from).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.PortfolioMarginProBankruptcyLoanRepay``: %v\n", err)
		return
	}

	// response from `PortfolioMarginProBankruptcyLoanRepay`: PortfolioMarginProBankruptcyLoanRepayResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **from** | [**PortfolioMarginProBankruptcyLoanRepayFromParameter**](PortfolioMarginProBankruptcyLoanRepayFromParameter.md) |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PortfolioMarginProBankruptcyLoanRepayResponse**](PortfolioMarginProBankruptcyLoanRepayResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryPortfolioMarginProBankruptcyLoanAmount

> QueryPortfolioMarginProBankruptcyLoanAmountResponse QueryPortfolioMarginProBankruptcyLoanAmount(ctx).RecvWindow(recvWindow).Execute()

Query Portfolio Margin Pro Bankruptcy Loan Amount (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomarginpro"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginProClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.QueryPortfolioMarginProBankruptcyLoanAmount(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.QueryPortfolioMarginProBankruptcyLoanAmount``: %v\n", err)
		return
	}

	// response from `QueryPortfolioMarginProBankruptcyLoanAmount`: QueryPortfolioMarginProBankruptcyLoanAmountResponse
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

[**QueryPortfolioMarginProBankruptcyLoanAmountResponse**](QueryPortfolioMarginProBankruptcyLoanAmountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryPortfolioMarginProBankruptcyLoanRepayHistory

> QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse QueryPortfolioMarginProBankruptcyLoanRepayHistory(ctx).StartTime(startTime).EndTime(endTime).Size(size).Current(current).RecvWindow(recvWindow).Execute()

Query Portfolio Margin Pro Bankruptcy Loan Repay History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomarginpro"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	startTime := int64(1623319461670) // int64 | Start time (optional)
	endTime := int64(1641782889000) // int64 | End time (optional)
	size := int64(10) // int64 | Number of results returned. (optional)
	current := int64(1) // int64 | Currently querying page. Start from 1. (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginProClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.QueryPortfolioMarginProBankruptcyLoanRepayHistory(context.Background()).StartTime(startTime).EndTime(endTime).Size(size).Current(current).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.QueryPortfolioMarginProBankruptcyLoanRepayHistory``: %v\n", err)
		return
	}

	// response from `QueryPortfolioMarginProBankruptcyLoanRepayHistory`: QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** | Start time | 
 **endTime** | **int64** | End time | 
 **size** | **int64** | Number of results returned. | 
 **current** | **int64** | Currently querying page. Start from 1. | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse**](QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryPortfolioMarginProNegativeBalanceInterestHistory

> QueryPortfolioMarginProNegativeBalanceInterestHistoryResponse QueryPortfolioMarginProNegativeBalanceInterestHistory(ctx).Asset(asset).StartTime(startTime).EndTime(endTime).Size(size).RecvWindow(recvWindow).Execute()

Query Portfolio Margin Pro Negative Balance Interest History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomarginpro"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	asset := "USDT" // string |  (optional)
	startTime := int64(1623319461670) // int64 | Start time (optional)
	endTime := int64(1641782889000) // int64 | End time (optional)
	size := int64(10) // int64 | Number of results returned. (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginProClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.QueryPortfolioMarginProNegativeBalanceInterestHistory(context.Background()).Asset(asset).StartTime(startTime).EndTime(endTime).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.QueryPortfolioMarginProNegativeBalanceInterestHistory``: %v\n", err)
		return
	}

	// response from `QueryPortfolioMarginProNegativeBalanceInterestHistory`: QueryPortfolioMarginProNegativeBalanceInterestHistoryResponse
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
 **startTime** | **int64** | Start time | 
 **endTime** | **int64** | End time | 
 **size** | **int64** | Number of results returned. | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryPortfolioMarginProNegativeBalanceInterestHistoryResponse**](QueryPortfolioMarginProNegativeBalanceInterestHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## RepayFuturesNegativeBalance

> RepayFuturesNegativeBalanceResponse RepayFuturesNegativeBalance(ctx).From(from).RecvWindow(recvWindow).Execute()

Repay futures Negative Balance (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomarginpro"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	from := models.PortfolioMarginProBankruptcyLoanRepayFromParameterSpot // PortfolioMarginProBankruptcyLoanRepayFromParameter |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginProClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.RepayFuturesNegativeBalance(context.Background()).From(from).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.RepayFuturesNegativeBalance``: %v\n", err)
		return
	}

	// response from `RepayFuturesNegativeBalance`: RepayFuturesNegativeBalanceResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **from** | [**PortfolioMarginProBankruptcyLoanRepayFromParameter**](PortfolioMarginProBankruptcyLoanRepayFromParameter.md) |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**RepayFuturesNegativeBalanceResponse**](RepayFuturesNegativeBalanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SetMarginCallLevel

> SetMarginCallLevelResponse SetMarginCallLevel(ctx).MarginCallLevel(marginCallLevel).RecvWindow(recvWindow).Execute()

Set Margin Call Level (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomarginpro"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	marginCallLevel := float32(1.5) // float32 | The value must be within the range [1.1, 2.0].
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginProClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.SetMarginCallLevel(context.Background()).MarginCallLevel(marginCallLevel).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.SetMarginCallLevel``: %v\n", err)
		return
	}

	// response from `SetMarginCallLevel`: SetMarginCallLevelResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **marginCallLevel** | **float32** | The value must be within the range [1.1, 2.0]. | 
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**SetMarginCallLevelResponse**](SetMarginCallLevelResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SwitchDeltaMode

> SwitchDeltaModeResponse SwitchDeltaMode(ctx).DeltaEnabled(deltaEnabled).RecvWindow(recvWindow).Execute()

Switch Delta Mode (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomarginpro"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	deltaEnabled := models.ChangeAutoRepayFuturesStatusAutoRepayParameterTrue // ChangeAutoRepayFuturesStatusAutoRepayParameter | `true` to enable Delta mode; `false` to disable Delta mode
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginProClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.SwitchDeltaMode(context.Background()).DeltaEnabled(deltaEnabled).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.SwitchDeltaMode``: %v\n", err)
		return
	}

	// response from `SwitchDeltaMode`: SwitchDeltaModeResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **deltaEnabled** | [**ChangeAutoRepayFuturesStatusAutoRepayParameter**](ChangeAutoRepayFuturesStatusAutoRepayParameter.md) | &#x60;true&#x60; to enable Delta mode; &#x60;false&#x60; to disable Delta mode | 
 **recvWindow** | **int64** |  | 

### Return type

[**SwitchDeltaModeResponse**](SwitchDeltaModeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## TransferLdusdtRwusdForPortfolioMargin

> TransferLdusdtRwusdForPortfolioMarginResponse TransferLdusdtRwusdForPortfolioMargin(ctx).Asset(asset).TransferType(transferType).Amount(amount).RecvWindow(recvWindow).Execute()

Transfer LDUSDT/RWUSD for Portfolio Margin (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomarginpro"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	asset := models.TransferLdusdtRwusdForPortfolioMarginAssetParameterLdusdt // TransferLdusdtRwusdForPortfolioMarginAssetParameter | 
	transferType := models.GetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameterEarnToFuture // GetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter | 
	amount := float32(1) // float32 | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginProClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.TransferLdusdtRwusdForPortfolioMargin(context.Background()).Asset(asset).TransferType(transferType).Amount(amount).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.TransferLdusdtRwusdForPortfolioMargin``: %v\n", err)
		return
	}

	// response from `TransferLdusdtRwusdForPortfolioMargin`: TransferLdusdtRwusdForPortfolioMarginResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **asset** | [**TransferLdusdtRwusdForPortfolioMarginAssetParameter**](TransferLdusdtRwusdForPortfolioMarginAssetParameter.md) |  | 
 **transferType** | [**GetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter**](GetTransferableEarnAssetBalanceForPortfolioMarginTransferTypeParameter.md) |  | 
 **amount** | **float32** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**TransferLdusdtRwusdForPortfolioMarginResponse**](TransferLdusdtRwusdForPortfolioMarginResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

