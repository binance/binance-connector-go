# \AccountAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**BnbTransfer**](AccountAPI.md#BnbTransfer) | **Post** /sapi/v1/portfolio/bnb-transfer | BNB transfer(USER_DATA)
[**ChangeAutoRepayFuturesStatus**](AccountAPI.md#ChangeAutoRepayFuturesStatus) | **Post** /sapi/v1/portfolio/repay-futures-switch | Change Auto-repay-futures Status(TRADE)
[**FundAutoCollection**](AccountAPI.md#FundAutoCollection) | **Post** /sapi/v1/portfolio/auto-collection | Fund Auto-collection(USER_DATA)
[**FundCollectionByAsset**](AccountAPI.md#FundCollectionByAsset) | **Post** /sapi/v1/portfolio/asset-collection | Fund Collection by Asset(USER_DATA)
[**GetAutoRepayFuturesStatus**](AccountAPI.md#GetAutoRepayFuturesStatus) | **Get** /sapi/v1/portfolio/repay-futures-switch | Get Auto-repay-futures Status(USER_DATA)
[**GetDeltaModeStatus**](AccountAPI.md#GetDeltaModeStatus) | **Get** /sapi/v1/portfolio/delta-mode | Get Delta Mode Status(USER_DATA)
[**GetPortfolioMarginProAccountBalance**](AccountAPI.md#GetPortfolioMarginProAccountBalance) | **Get** /sapi/v1/portfolio/balance | Get Portfolio Margin Pro Account Balance(USER_DATA)
[**GetPortfolioMarginProAccountInfo**](AccountAPI.md#GetPortfolioMarginProAccountInfo) | **Get** /sapi/v1/portfolio/account | Get Portfolio Margin Pro Account Info(USER_DATA)
[**GetPortfolioMarginProSpanAccountInfo**](AccountAPI.md#GetPortfolioMarginProSpanAccountInfo) | **Get** /sapi/v2/portfolio/account | Get Portfolio Margin Pro SPAN Account Info(USER_DATA)
[**GetTransferableEarnAssetBalanceForPortfolioMargin**](AccountAPI.md#GetTransferableEarnAssetBalanceForPortfolioMargin) | **Get** /sapi/v1/portfolio/earn-asset-balance | Get Transferable Earn Asset Balance for Portfolio Margin (USER_DATA)
[**PortfolioMarginProBankruptcyLoanRepay**](AccountAPI.md#PortfolioMarginProBankruptcyLoanRepay) | **Post** /sapi/v1/portfolio/repay | Portfolio Margin Pro Bankruptcy Loan Repay
[**QueryPortfolioMarginProBankruptcyLoanAmount**](AccountAPI.md#QueryPortfolioMarginProBankruptcyLoanAmount) | **Get** /sapi/v1/portfolio/pmLoan | Query Portfolio Margin Pro Bankruptcy Loan Amount(USER_DATA)
[**QueryPortfolioMarginProBankruptcyLoanRepayHistory**](AccountAPI.md#QueryPortfolioMarginProBankruptcyLoanRepayHistory) | **Get** /sapi/v1/portfolio/pmloan-history | Query Portfolio Margin Pro Bankruptcy Loan Repay History(USER_DATA)
[**QueryPortfolioMarginProNegativeBalanceInterestHistory**](AccountAPI.md#QueryPortfolioMarginProNegativeBalanceInterestHistory) | **Get** /sapi/v1/portfolio/interest-history | Query Portfolio Margin Pro Negative Balance Interest History(USER_DATA)
[**RepayFuturesNegativeBalance**](AccountAPI.md#RepayFuturesNegativeBalance) | **Post** /sapi/v1/portfolio/repay-futures-negative-balance | Repay futures Negative Balance(USER_DATA)
[**SwitchDeltaMode**](AccountAPI.md#SwitchDeltaMode) | **Post** /sapi/v1/portfolio/delta-mode | Switch Delta Mode(TRADE)
[**TransferLdusdtRwusdForPortfolioMargin**](AccountAPI.md#TransferLdusdtRwusdForPortfolioMargin) | **Post** /sapi/v1/portfolio/earn-asset-transfer | Transfer LDUSDT/RWUSD for Portfolio Margin(TRADE)


## BnbTransfer

> BnbTransferResponse BnbTransfer(ctx).Amount(amount).TransferSide(transferSide).RecvWindow(recvWindow).Execute()

BNB transfer(USER_DATA)


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
	transferSide := "transferSide_example" // string | \"TO_UM\",\"FROM_UM\"
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
 **transferSide** | **string** | \&quot;TO_UM\&quot;,\&quot;FROM_UM\&quot; | 
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

Change Auto-repay-futures Status(TRADE)


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
	autoRepay := "true" // string | Default: `true`; `false` for turn off the auto-repay futures negative balance function
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
 **autoRepay** | **string** | Default: &#x60;true&#x60;; &#x60;false&#x60; for turn off the auto-repay futures negative balance function | 
 **recvWindow** | **int64** |  | 

### Return type

[**ChangeAutoRepayFuturesStatusResponse**](ChangeAutoRepayFuturesStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## FundAutoCollection

> FundAutoCollectionResponse FundAutoCollection(ctx).RecvWindow(recvWindow).Execute()

Fund Auto-collection(USER_DATA)


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

Fund Collection by Asset(USER_DATA)


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
	asset := "asset_example" // string | `LDUSDT` and `RWUSD`
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
 **asset** | **string** | &#x60;LDUSDT&#x60; and &#x60;RWUSD&#x60; | 
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

Get Auto-repay-futures Status(USER_DATA)


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

Get Delta Mode Status(USER_DATA)


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


## GetPortfolioMarginProAccountBalance

> GetPortfolioMarginProAccountBalanceResponse GetPortfolioMarginProAccountBalance(ctx).Asset(asset).RecvWindow(recvWindow).Execute()

Get Portfolio Margin Pro Account Balance(USER_DATA)


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
	asset := "asset_example" // string |  (optional)
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

Get Portfolio Margin Pro Account Info(USER_DATA)


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

Get Portfolio Margin Pro SPAN Account Info(USER_DATA)


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
	asset := "asset_example" // string | `LDUSDT` and `RWUSD`
	transferType := "transferType_example" // string | `EARN_TO_FUTURE` /`FUTURE_TO_EARN`
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
 **asset** | **string** | &#x60;LDUSDT&#x60; and &#x60;RWUSD&#x60; | 
 **transferType** | **string** | &#x60;EARN_TO_FUTURE&#x60; /&#x60;FUTURE_TO_EARN&#x60; | 
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

Portfolio Margin Pro Bankruptcy Loan Repay


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
	from := "SPOT" // string | SPOT or MARGIN，default SPOT (optional)
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
 **from** | **string** | SPOT or MARGIN，default SPOT | 
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

Query Portfolio Margin Pro Bankruptcy Loan Amount(USER_DATA)


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

> QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse QueryPortfolioMarginProBankruptcyLoanRepayHistory(ctx).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Query Portfolio Margin Pro Bankruptcy Loan Repay History(USER_DATA)


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
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	current := int64(1) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(10) // int64 | Default:10 Max:100 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginProClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.QueryPortfolioMarginProBankruptcyLoanRepayHistory(context.Background()).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
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
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10 Max:100 | 
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

Query Portfolio Margin Pro Negative Balance Interest History(USER_DATA)


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
	asset := "asset_example" // string |  (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	size := int64(10) // int64 | Default:10 Max:100 (optional)
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
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **size** | **int64** | Default:10 Max:100 | 
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

Repay futures Negative Balance(USER_DATA)


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
	from := "SPOT" // string | SPOT or MARGIN，default SPOT (optional)
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
 **from** | **string** | SPOT or MARGIN，default SPOT | 
 **recvWindow** | **int64** |  | 

### Return type

[**RepayFuturesNegativeBalanceResponse**](RepayFuturesNegativeBalanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SwitchDeltaMode

> SwitchDeltaModeResponse SwitchDeltaMode(ctx).DeltaEnabled(deltaEnabled).RecvWindow(recvWindow).Execute()

Switch Delta Mode(TRADE)


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
	deltaEnabled := "deltaEnabled_example" // string | `true` to enable Delta mode; `false` to disable Delta mode
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
 **deltaEnabled** | **string** | &#x60;true&#x60; to enable Delta mode; &#x60;false&#x60; to disable Delta mode | 
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

Transfer LDUSDT/RWUSD for Portfolio Margin(TRADE)


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
	asset := "asset_example" // string | `LDUSDT` and `RWUSD`
	transferType := "transferType_example" // string | `EARN_TO_FUTURE` /`FUTURE_TO_EARN`
	amount := float32(1.0) // float32 | 
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
 **asset** | **string** | &#x60;LDUSDT&#x60; and &#x60;RWUSD&#x60; | 
 **transferType** | **string** | &#x60;EARN_TO_FUTURE&#x60; /&#x60;FUTURE_TO_EARN&#x60; | 
 **amount** | **float32** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**TransferLdusdtRwusdForPortfolioMarginResponse**](TransferLdusdtRwusdForPortfolioMarginResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

