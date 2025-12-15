# \AccountAPI

All URIs are relative to *https://papi.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AccountBalance**](AccountAPI.md#AccountBalance) | **Get** /papi/v1/balance | Account Balance(USER_DATA)
[**AccountInformation**](AccountAPI.md#AccountInformation) | **Get** /papi/v1/account | Account Information(USER_DATA)
[**BnbTransfer**](AccountAPI.md#BnbTransfer) | **Post** /papi/v1/bnb-transfer | BNB transfer (TRADE)
[**ChangeAutoRepayFuturesStatus**](AccountAPI.md#ChangeAutoRepayFuturesStatus) | **Post** /papi/v1/repay-futures-switch | Change Auto-repay-futures Status(TRADE)
[**ChangeCmInitialLeverage**](AccountAPI.md#ChangeCmInitialLeverage) | **Post** /papi/v1/cm/leverage | Change CM Initial Leverage (TRADE)
[**ChangeCmPositionMode**](AccountAPI.md#ChangeCmPositionMode) | **Post** /papi/v1/cm/positionSide/dual | Change CM Position Mode(TRADE)
[**ChangeUmInitialLeverage**](AccountAPI.md#ChangeUmInitialLeverage) | **Post** /papi/v1/um/leverage | Change UM Initial Leverage(TRADE)
[**ChangeUmPositionMode**](AccountAPI.md#ChangeUmPositionMode) | **Post** /papi/v1/um/positionSide/dual | Change UM Position Mode(TRADE)
[**CmNotionalAndLeverageBrackets**](AccountAPI.md#CmNotionalAndLeverageBrackets) | **Get** /papi/v1/cm/leverageBracket | CM Notional and Leverage Brackets(USER_DATA)
[**FundAutoCollection**](AccountAPI.md#FundAutoCollection) | **Post** /papi/v1/auto-collection | Fund Auto-collection(TRADE)
[**FundCollectionByAsset**](AccountAPI.md#FundCollectionByAsset) | **Post** /papi/v1/asset-collection | Fund Collection by Asset(TRADE)
[**GetAutoRepayFuturesStatus**](AccountAPI.md#GetAutoRepayFuturesStatus) | **Get** /papi/v1/repay-futures-switch | Get Auto-repay-futures Status(USER_DATA)
[**GetCmAccountDetail**](AccountAPI.md#GetCmAccountDetail) | **Get** /papi/v1/cm/account | Get CM Account Detail(USER_DATA)
[**GetCmCurrentPositionMode**](AccountAPI.md#GetCmCurrentPositionMode) | **Get** /papi/v1/cm/positionSide/dual | Get CM Current Position Mode(USER_DATA)
[**GetCmIncomeHistory**](AccountAPI.md#GetCmIncomeHistory) | **Get** /papi/v1/cm/income | Get CM Income History(USER_DATA)
[**GetDownloadIdForUmFuturesOrderHistory**](AccountAPI.md#GetDownloadIdForUmFuturesOrderHistory) | **Get** /papi/v1/um/order/asyn | Get Download Id For UM Futures Order History (USER_DATA)
[**GetDownloadIdForUmFuturesTradeHistory**](AccountAPI.md#GetDownloadIdForUmFuturesTradeHistory) | **Get** /papi/v1/um/trade/asyn | Get Download Id For UM Futures Trade History (USER_DATA)
[**GetDownloadIdForUmFuturesTransactionHistory**](AccountAPI.md#GetDownloadIdForUmFuturesTransactionHistory) | **Get** /papi/v1/um/income/asyn | Get Download Id For UM Futures Transaction History (USER_DATA)
[**GetMarginBorrowLoanInterestHistory**](AccountAPI.md#GetMarginBorrowLoanInterestHistory) | **Get** /papi/v1/margin/marginInterestHistory | Get Margin Borrow/Loan Interest History(USER_DATA)
[**GetUmAccountDetail**](AccountAPI.md#GetUmAccountDetail) | **Get** /papi/v1/um/account | Get UM Account Detail(USER_DATA)
[**GetUmAccountDetailV2**](AccountAPI.md#GetUmAccountDetailV2) | **Get** /papi/v2/um/account | Get UM Account Detail V2(USER_DATA)
[**GetUmCurrentPositionMode**](AccountAPI.md#GetUmCurrentPositionMode) | **Get** /papi/v1/um/positionSide/dual | Get UM Current Position Mode(USER_DATA)
[**GetUmFuturesOrderDownloadLinkById**](AccountAPI.md#GetUmFuturesOrderDownloadLinkById) | **Get** /papi/v1/um/order/asyn/id | Get UM Futures Order Download Link by Id(USER_DATA)
[**GetUmFuturesTradeDownloadLinkById**](AccountAPI.md#GetUmFuturesTradeDownloadLinkById) | **Get** /papi/v1/um/trade/asyn/id | Get UM Futures Trade Download Link by Id(USER_DATA)
[**GetUmFuturesTransactionDownloadLinkById**](AccountAPI.md#GetUmFuturesTransactionDownloadLinkById) | **Get** /papi/v1/um/income/asyn/id | Get UM Futures Transaction Download Link by Id(USER_DATA)
[**GetUmIncomeHistory**](AccountAPI.md#GetUmIncomeHistory) | **Get** /papi/v1/um/income | Get UM Income History(USER_DATA)
[**GetUserCommissionRateForCm**](AccountAPI.md#GetUserCommissionRateForCm) | **Get** /papi/v1/cm/commissionRate | Get User Commission Rate for CM(USER_DATA)
[**GetUserCommissionRateForUm**](AccountAPI.md#GetUserCommissionRateForUm) | **Get** /papi/v1/um/commissionRate | Get User Commission Rate for UM(USER_DATA)
[**MarginMaxBorrow**](AccountAPI.md#MarginMaxBorrow) | **Get** /papi/v1/margin/maxBorrowable | Margin Max Borrow(USER_DATA)
[**PortfolioMarginUmTradingQuantitativeRulesIndicators**](AccountAPI.md#PortfolioMarginUmTradingQuantitativeRulesIndicators) | **Get** /papi/v1/um/apiTradingStatus | Portfolio Margin UM Trading Quantitative Rules Indicators(USER_DATA)
[**QueryCmPositionInformation**](AccountAPI.md#QueryCmPositionInformation) | **Get** /papi/v1/cm/positionRisk | Query CM Position Information(USER_DATA)
[**QueryMarginLoanRecord**](AccountAPI.md#QueryMarginLoanRecord) | **Get** /papi/v1/margin/marginLoan | Query Margin Loan Record(USER_DATA)
[**QueryMarginMaxWithdraw**](AccountAPI.md#QueryMarginMaxWithdraw) | **Get** /papi/v1/margin/maxWithdraw | Query Margin Max Withdraw(USER_DATA)
[**QueryMarginRepayRecord**](AccountAPI.md#QueryMarginRepayRecord) | **Get** /papi/v1/margin/repayLoan | Query Margin repay Record(USER_DATA)
[**QueryPortfolioMarginNegativeBalanceInterestHistory**](AccountAPI.md#QueryPortfolioMarginNegativeBalanceInterestHistory) | **Get** /papi/v1/portfolio/interest-history | Query Portfolio Margin Negative Balance Interest History(USER_DATA)
[**QueryUmPositionInformation**](AccountAPI.md#QueryUmPositionInformation) | **Get** /papi/v1/um/positionRisk | Query UM Position Information(USER_DATA)
[**QueryUserNegativeBalanceAutoExchangeRecord**](AccountAPI.md#QueryUserNegativeBalanceAutoExchangeRecord) | **Get** /papi/v1/portfolio/negative-balance-exchange-record | Query User Negative Balance Auto Exchange Record (USER_DATA)
[**QueryUserRateLimit**](AccountAPI.md#QueryUserRateLimit) | **Get** /papi/v1/rateLimit/order | Query User Rate Limit (USER_DATA)
[**RepayFuturesNegativeBalance**](AccountAPI.md#RepayFuturesNegativeBalance) | **Post** /papi/v1/repay-futures-negative-balance | Repay futures Negative Balance(USER_DATA)
[**UmFuturesAccountConfiguration**](AccountAPI.md#UmFuturesAccountConfiguration) | **Get** /papi/v1/um/accountConfig | UM Futures Account Configuration(USER_DATA)
[**UmFuturesSymbolConfiguration**](AccountAPI.md#UmFuturesSymbolConfiguration) | **Get** /papi/v1/um/symbolConfig | UM Futures Symbol Configuration(USER_DATA)
[**UmNotionalAndLeverageBrackets**](AccountAPI.md#UmNotionalAndLeverageBrackets) | **Get** /papi/v1/um/leverageBracket | UM Notional and Leverage Brackets (USER_DATA)


## AccountBalance

> AccountBalanceResponse AccountBalance(ctx).Asset(asset).RecvWindow(recvWindow).Execute()

Account Balance(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	asset := "asset_example" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.AccountBalance(context.Background()).Asset(asset).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.AccountBalance``: %v\n", err)
		return
	}

	// response from `AccountBalance`: AccountBalanceResponse
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

[**AccountBalanceResponse**](AccountBalanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## AccountInformation

> AccountInformationResponse AccountInformation(ctx).RecvWindow(recvWindow).Execute()

Account Information(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

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


## BnbTransfer

> BnbTransferResponse BnbTransfer(ctx).Amount(amount).TransferSide(transferSide).RecvWindow(recvWindow).Execute()

BNB transfer (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
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
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	autoRepay := "true" // string | Default: `true`; `false` for turn off the auto-repay futures negative balance function
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

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


## ChangeCmInitialLeverage

> ChangeCmInitialLeverageResponse ChangeCmInitialLeverage(ctx).Symbol(symbol).Leverage(leverage).RecvWindow(recvWindow).Execute()

Change CM Initial Leverage (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	leverage := int64(789) // int64 | target initial leverage: int from 1 to 125
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.ChangeCmInitialLeverage(context.Background()).Symbol(symbol).Leverage(leverage).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.ChangeCmInitialLeverage``: %v\n", err)
		return
	}

	// response from `ChangeCmInitialLeverage`: ChangeCmInitialLeverageResponse
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
 **leverage** | **int64** | target initial leverage: int from 1 to 125 | 
 **recvWindow** | **int64** |  | 

### Return type

[**ChangeCmInitialLeverageResponse**](ChangeCmInitialLeverageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## ChangeCmPositionMode

> ChangeCmPositionModeResponse ChangeCmPositionMode(ctx).DualSidePosition(dualSidePosition).RecvWindow(recvWindow).Execute()

Change CM Position Mode(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	dualSidePosition := "dualSidePosition_example" // string | \"true\": Hedge Mode; \"false\": One-way Mode
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.ChangeCmPositionMode(context.Background()).DualSidePosition(dualSidePosition).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.ChangeCmPositionMode``: %v\n", err)
		return
	}

	// response from `ChangeCmPositionMode`: ChangeCmPositionModeResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **dualSidePosition** | **string** | \&quot;true\&quot;: Hedge Mode; \&quot;false\&quot;: One-way Mode | 
 **recvWindow** | **int64** |  | 

### Return type

[**ChangeCmPositionModeResponse**](ChangeCmPositionModeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## ChangeUmInitialLeverage

> ChangeUmInitialLeverageResponse ChangeUmInitialLeverage(ctx).Symbol(symbol).Leverage(leverage).RecvWindow(recvWindow).Execute()

Change UM Initial Leverage(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	leverage := int64(789) // int64 | target initial leverage: int from 1 to 125
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.ChangeUmInitialLeverage(context.Background()).Symbol(symbol).Leverage(leverage).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.ChangeUmInitialLeverage``: %v\n", err)
		return
	}

	// response from `ChangeUmInitialLeverage`: ChangeUmInitialLeverageResponse
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
 **leverage** | **int64** | target initial leverage: int from 1 to 125 | 
 **recvWindow** | **int64** |  | 

### Return type

[**ChangeUmInitialLeverageResponse**](ChangeUmInitialLeverageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## ChangeUmPositionMode

> ChangeUmPositionModeResponse ChangeUmPositionMode(ctx).DualSidePosition(dualSidePosition).RecvWindow(recvWindow).Execute()

Change UM Position Mode(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	dualSidePosition := "dualSidePosition_example" // string | \"true\": Hedge Mode; \"false\": One-way Mode
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.ChangeUmPositionMode(context.Background()).DualSidePosition(dualSidePosition).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.ChangeUmPositionMode``: %v\n", err)
		return
	}

	// response from `ChangeUmPositionMode`: ChangeUmPositionModeResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **dualSidePosition** | **string** | \&quot;true\&quot;: Hedge Mode; \&quot;false\&quot;: One-way Mode | 
 **recvWindow** | **int64** |  | 

### Return type

[**ChangeUmPositionModeResponse**](ChangeUmPositionModeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CmNotionalAndLeverageBrackets

> CmNotionalAndLeverageBracketsResponse CmNotionalAndLeverageBrackets(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

CM Notional and Leverage Brackets(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
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
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.CmNotionalAndLeverageBrackets(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.CmNotionalAndLeverageBrackets``: %v\n", err)
		return
	}

	// response from `CmNotionalAndLeverageBrackets`: CmNotionalAndLeverageBracketsResponse
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

[**CmNotionalAndLeverageBracketsResponse**](CmNotionalAndLeverageBracketsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## FundAutoCollection

> FundAutoCollectionResponse FundAutoCollection(ctx).RecvWindow(recvWindow).Execute()

Fund Auto-collection(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

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

Fund Collection by Asset(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	asset := "asset_example" // string | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

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

Get Auto-repay-futures Status(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

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


## GetCmAccountDetail

> GetCmAccountDetailResponse GetCmAccountDetail(ctx).RecvWindow(recvWindow).Execute()

Get CM Account Detail(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetCmAccountDetail(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetCmAccountDetail``: %v\n", err)
		return
	}

	// response from `GetCmAccountDetail`: GetCmAccountDetailResponse
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

[**GetCmAccountDetailResponse**](GetCmAccountDetailResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetCmCurrentPositionMode

> GetCmCurrentPositionModeResponse GetCmCurrentPositionMode(ctx).RecvWindow(recvWindow).Execute()

Get CM Current Position Mode(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetCmCurrentPositionMode(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetCmCurrentPositionMode``: %v\n", err)
		return
	}

	// response from `GetCmCurrentPositionMode`: GetCmCurrentPositionModeResponse
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

[**GetCmCurrentPositionModeResponse**](GetCmCurrentPositionModeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetCmIncomeHistory

> GetCmIncomeHistoryResponse GetCmIncomeHistory(ctx).Symbol(symbol).IncomeType(incomeType).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Get CM Income History(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string |  (optional)
	incomeType := "incomeType_example" // string | TRANSFER, WELCOME_BONUS, REALIZED_PNL, FUNDING_FEE, COMMISSION, INSURANCE_CLEAR, REFERRAL_KICKBACK, COMMISSION_REBATE, API_REBATE, CONTEST_REWARD, CROSS_COLLATERAL_TRANSFER, OPTIONS_PREMIUM_FEE, OPTIONS_SETTLE_PROFIT, INTERNAL_TRANSFER, AUTO_EXCHANGE, DELIVERED_SETTELMENT, COIN_SWAP_DEPOSIT, COIN_SWAP_WITHDRAW, POSITION_LIMIT_INCREASE_FEE (optional)
	startTime := int64(1623319461670) // int64 | Timestamp in ms to get funding from INCLUSIVE. (optional)
	endTime := int64(1641782889000) // int64 | Timestamp in ms to get funding until INCLUSIVE. (optional)
	page := int64(789) // int64 |  (optional)
	limit := int64(100) // int64 | Default 100; max 1000 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetCmIncomeHistory(context.Background()).Symbol(symbol).IncomeType(incomeType).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetCmIncomeHistory``: %v\n", err)
		return
	}

	// response from `GetCmIncomeHistory`: GetCmIncomeHistoryResponse
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
 **incomeType** | **string** | TRANSFER, WELCOME_BONUS, REALIZED_PNL, FUNDING_FEE, COMMISSION, INSURANCE_CLEAR, REFERRAL_KICKBACK, COMMISSION_REBATE, API_REBATE, CONTEST_REWARD, CROSS_COLLATERAL_TRANSFER, OPTIONS_PREMIUM_FEE, OPTIONS_SETTLE_PROFIT, INTERNAL_TRANSFER, AUTO_EXCHANGE, DELIVERED_SETTELMENT, COIN_SWAP_DEPOSIT, COIN_SWAP_WITHDRAW, POSITION_LIMIT_INCREASE_FEE | 
 **startTime** | **int64** | Timestamp in ms to get funding from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding until INCLUSIVE. | 
 **page** | **int64** |  | 
 **limit** | **int64** | Default 100; max 1000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetCmIncomeHistoryResponse**](GetCmIncomeHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetDownloadIdForUmFuturesOrderHistory

> GetDownloadIdForUmFuturesOrderHistoryResponse GetDownloadIdForUmFuturesOrderHistory(ctx).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

Get Download Id For UM Futures Order History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	startTime := int64(1623319461670) // int64 | 
	endTime := int64(1641782889000) // int64 | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetDownloadIdForUmFuturesOrderHistory(context.Background()).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetDownloadIdForUmFuturesOrderHistory``: %v\n", err)
		return
	}

	// response from `GetDownloadIdForUmFuturesOrderHistory`: GetDownloadIdForUmFuturesOrderHistoryResponse
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
 **recvWindow** | **int64** |  | 

### Return type

[**GetDownloadIdForUmFuturesOrderHistoryResponse**](GetDownloadIdForUmFuturesOrderHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetDownloadIdForUmFuturesTradeHistory

> GetDownloadIdForUmFuturesTradeHistoryResponse GetDownloadIdForUmFuturesTradeHistory(ctx).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

Get Download Id For UM Futures Trade History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	startTime := int64(1623319461670) // int64 | 
	endTime := int64(1641782889000) // int64 | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetDownloadIdForUmFuturesTradeHistory(context.Background()).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetDownloadIdForUmFuturesTradeHistory``: %v\n", err)
		return
	}

	// response from `GetDownloadIdForUmFuturesTradeHistory`: GetDownloadIdForUmFuturesTradeHistoryResponse
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
 **recvWindow** | **int64** |  | 

### Return type

[**GetDownloadIdForUmFuturesTradeHistoryResponse**](GetDownloadIdForUmFuturesTradeHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetDownloadIdForUmFuturesTransactionHistory

> GetDownloadIdForUmFuturesTransactionHistoryResponse GetDownloadIdForUmFuturesTransactionHistory(ctx).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

Get Download Id For UM Futures Transaction History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	startTime := int64(1623319461670) // int64 | 
	endTime := int64(1641782889000) // int64 | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetDownloadIdForUmFuturesTransactionHistory(context.Background()).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetDownloadIdForUmFuturesTransactionHistory``: %v\n", err)
		return
	}

	// response from `GetDownloadIdForUmFuturesTransactionHistory`: GetDownloadIdForUmFuturesTransactionHistoryResponse
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
 **recvWindow** | **int64** |  | 

### Return type

[**GetDownloadIdForUmFuturesTransactionHistoryResponse**](GetDownloadIdForUmFuturesTransactionHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetMarginBorrowLoanInterestHistory

> GetMarginBorrowLoanInterestHistoryResponse GetMarginBorrowLoanInterestHistory(ctx).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Archived(archived).RecvWindow(recvWindow).Execute()

Get Margin Borrow/Loan Interest History(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	asset := "asset_example" // string |  (optional)
	startTime := int64(1623319461670) // int64 | Timestamp in ms to get funding from INCLUSIVE. (optional)
	endTime := int64(1641782889000) // int64 | Timestamp in ms to get funding until INCLUSIVE. (optional)
	current := int64(1) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(10) // int64 | Default:10 Max:100 (optional)
	archived := "archived_example" // string | Default: `false`. Set to `true` for archived data from 6 months ago (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetMarginBorrowLoanInterestHistory(context.Background()).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Archived(archived).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetMarginBorrowLoanInterestHistory``: %v\n", err)
		return
	}

	// response from `GetMarginBorrowLoanInterestHistory`: GetMarginBorrowLoanInterestHistoryResponse
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
 **startTime** | **int64** | Timestamp in ms to get funding from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding until INCLUSIVE. | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10 Max:100 | 
 **archived** | **string** | Default: &#x60;false&#x60;. Set to &#x60;true&#x60; for archived data from 6 months ago | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetMarginBorrowLoanInterestHistoryResponse**](GetMarginBorrowLoanInterestHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetUmAccountDetail

> GetUmAccountDetailResponse GetUmAccountDetail(ctx).RecvWindow(recvWindow).Execute()

Get UM Account Detail(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetUmAccountDetail(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetUmAccountDetail``: %v\n", err)
		return
	}

	// response from `GetUmAccountDetail`: GetUmAccountDetailResponse
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

[**GetUmAccountDetailResponse**](GetUmAccountDetailResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetUmAccountDetailV2

> GetUmAccountDetailV2Response GetUmAccountDetailV2(ctx).RecvWindow(recvWindow).Execute()

Get UM Account Detail V2(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetUmAccountDetailV2(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetUmAccountDetailV2``: %v\n", err)
		return
	}

	// response from `GetUmAccountDetailV2`: GetUmAccountDetailV2Response
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

[**GetUmAccountDetailV2Response**](GetUmAccountDetailV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetUmCurrentPositionMode

> GetUmCurrentPositionModeResponse GetUmCurrentPositionMode(ctx).RecvWindow(recvWindow).Execute()

Get UM Current Position Mode(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetUmCurrentPositionMode(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetUmCurrentPositionMode``: %v\n", err)
		return
	}

	// response from `GetUmCurrentPositionMode`: GetUmCurrentPositionModeResponse
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

[**GetUmCurrentPositionModeResponse**](GetUmCurrentPositionModeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetUmFuturesOrderDownloadLinkById

> GetUmFuturesOrderDownloadLinkByIdResponse GetUmFuturesOrderDownloadLinkById(ctx).DownloadId(downloadId).RecvWindow(recvWindow).Execute()

Get UM Futures Order Download Link by Id(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
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
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetUmFuturesOrderDownloadLinkById(context.Background()).DownloadId(downloadId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetUmFuturesOrderDownloadLinkById``: %v\n", err)
		return
	}

	// response from `GetUmFuturesOrderDownloadLinkById`: GetUmFuturesOrderDownloadLinkByIdResponse
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

[**GetUmFuturesOrderDownloadLinkByIdResponse**](GetUmFuturesOrderDownloadLinkByIdResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetUmFuturesTradeDownloadLinkById

> GetUmFuturesTradeDownloadLinkByIdResponse GetUmFuturesTradeDownloadLinkById(ctx).DownloadId(downloadId).RecvWindow(recvWindow).Execute()

Get UM Futures Trade Download Link by Id(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
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
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetUmFuturesTradeDownloadLinkById(context.Background()).DownloadId(downloadId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetUmFuturesTradeDownloadLinkById``: %v\n", err)
		return
	}

	// response from `GetUmFuturesTradeDownloadLinkById`: GetUmFuturesTradeDownloadLinkByIdResponse
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

[**GetUmFuturesTradeDownloadLinkByIdResponse**](GetUmFuturesTradeDownloadLinkByIdResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetUmFuturesTransactionDownloadLinkById

> GetUmFuturesTransactionDownloadLinkByIdResponse GetUmFuturesTransactionDownloadLinkById(ctx).DownloadId(downloadId).RecvWindow(recvWindow).Execute()

Get UM Futures Transaction Download Link by Id(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
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
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetUmFuturesTransactionDownloadLinkById(context.Background()).DownloadId(downloadId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetUmFuturesTransactionDownloadLinkById``: %v\n", err)
		return
	}

	// response from `GetUmFuturesTransactionDownloadLinkById`: GetUmFuturesTransactionDownloadLinkByIdResponse
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

[**GetUmFuturesTransactionDownloadLinkByIdResponse**](GetUmFuturesTransactionDownloadLinkByIdResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetUmIncomeHistory

> GetUmIncomeHistoryResponse GetUmIncomeHistory(ctx).Symbol(symbol).IncomeType(incomeType).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Get UM Income History(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string |  (optional)
	incomeType := "incomeType_example" // string | TRANSFER, WELCOME_BONUS, REALIZED_PNL, FUNDING_FEE, COMMISSION, INSURANCE_CLEAR, REFERRAL_KICKBACK, COMMISSION_REBATE, API_REBATE, CONTEST_REWARD, CROSS_COLLATERAL_TRANSFER, OPTIONS_PREMIUM_FEE, OPTIONS_SETTLE_PROFIT, INTERNAL_TRANSFER, AUTO_EXCHANGE, DELIVERED_SETTELMENT, COIN_SWAP_DEPOSIT, COIN_SWAP_WITHDRAW, POSITION_LIMIT_INCREASE_FEE (optional)
	startTime := int64(1623319461670) // int64 | Timestamp in ms to get funding from INCLUSIVE. (optional)
	endTime := int64(1641782889000) // int64 | Timestamp in ms to get funding until INCLUSIVE. (optional)
	page := int64(789) // int64 |  (optional)
	limit := int64(100) // int64 | Default 100; max 1000 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetUmIncomeHistory(context.Background()).Symbol(symbol).IncomeType(incomeType).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetUmIncomeHistory``: %v\n", err)
		return
	}

	// response from `GetUmIncomeHistory`: GetUmIncomeHistoryResponse
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
 **incomeType** | **string** | TRANSFER, WELCOME_BONUS, REALIZED_PNL, FUNDING_FEE, COMMISSION, INSURANCE_CLEAR, REFERRAL_KICKBACK, COMMISSION_REBATE, API_REBATE, CONTEST_REWARD, CROSS_COLLATERAL_TRANSFER, OPTIONS_PREMIUM_FEE, OPTIONS_SETTLE_PROFIT, INTERNAL_TRANSFER, AUTO_EXCHANGE, DELIVERED_SETTELMENT, COIN_SWAP_DEPOSIT, COIN_SWAP_WITHDRAW, POSITION_LIMIT_INCREASE_FEE | 
 **startTime** | **int64** | Timestamp in ms to get funding from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding until INCLUSIVE. | 
 **page** | **int64** |  | 
 **limit** | **int64** | Default 100; max 1000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetUmIncomeHistoryResponse**](GetUmIncomeHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetUserCommissionRateForCm

> GetUserCommissionRateForCmResponse GetUserCommissionRateForCm(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Get User Commission Rate for CM(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
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
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetUserCommissionRateForCm(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetUserCommissionRateForCm``: %v\n", err)
		return
	}

	// response from `GetUserCommissionRateForCm`: GetUserCommissionRateForCmResponse
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

[**GetUserCommissionRateForCmResponse**](GetUserCommissionRateForCmResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetUserCommissionRateForUm

> GetUserCommissionRateForUmResponse GetUserCommissionRateForUm(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Get User Commission Rate for UM(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
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
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetUserCommissionRateForUm(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetUserCommissionRateForUm``: %v\n", err)
		return
	}

	// response from `GetUserCommissionRateForUm`: GetUserCommissionRateForUmResponse
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

[**GetUserCommissionRateForUmResponse**](GetUserCommissionRateForUmResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MarginMaxBorrow

> MarginMaxBorrowResponse MarginMaxBorrow(ctx).Asset(asset).RecvWindow(recvWindow).Execute()

Margin Max Borrow(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	asset := "asset_example" // string | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.MarginMaxBorrow(context.Background()).Asset(asset).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.MarginMaxBorrow``: %v\n", err)
		return
	}

	// response from `MarginMaxBorrow`: MarginMaxBorrowResponse
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

[**MarginMaxBorrowResponse**](MarginMaxBorrowResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## PortfolioMarginUmTradingQuantitativeRulesIndicators

> PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse PortfolioMarginUmTradingQuantitativeRulesIndicators(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Portfolio Margin UM Trading Quantitative Rules Indicators(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
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
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.PortfolioMarginUmTradingQuantitativeRulesIndicators(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.PortfolioMarginUmTradingQuantitativeRulesIndicators``: %v\n", err)
		return
	}

	// response from `PortfolioMarginUmTradingQuantitativeRulesIndicators`: PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse
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

[**PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse**](PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryCmPositionInformation

> QueryCmPositionInformationResponse QueryCmPositionInformation(ctx).MarginAsset(marginAsset).Pair(pair).RecvWindow(recvWindow).Execute()

Query CM Position Information(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	marginAsset := "marginAsset_example" // string |  (optional)
	pair := "pair_example" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.QueryCmPositionInformation(context.Background()).MarginAsset(marginAsset).Pair(pair).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.QueryCmPositionInformation``: %v\n", err)
		return
	}

	// response from `QueryCmPositionInformation`: QueryCmPositionInformationResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **marginAsset** | **string** |  | 
 **pair** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryCmPositionInformationResponse**](QueryCmPositionInformationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryMarginLoanRecord

> QueryMarginLoanRecordResponse QueryMarginLoanRecord(ctx).Asset(asset).TxId(txId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Archived(archived).RecvWindow(recvWindow).Execute()

Query Margin Loan Record(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	asset := "asset_example" // string | 
	txId := int64(1) // int64 | the `tranId` in `POST/papi/v1/marginLoan` (optional)
	startTime := int64(1623319461670) // int64 | Timestamp in ms to get funding from INCLUSIVE. (optional)
	endTime := int64(1641782889000) // int64 | Timestamp in ms to get funding until INCLUSIVE. (optional)
	current := int64(1) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(10) // int64 | Default:10 Max:100 (optional)
	archived := "archived_example" // string | Default: `false`. Set to `true` for archived data from 6 months ago (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.QueryMarginLoanRecord(context.Background()).Asset(asset).TxId(txId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Archived(archived).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.QueryMarginLoanRecord``: %v\n", err)
		return
	}

	// response from `QueryMarginLoanRecord`: QueryMarginLoanRecordResponse
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
 **txId** | **int64** | the &#x60;tranId&#x60; in &#x60;POST/papi/v1/marginLoan&#x60; | 
 **startTime** | **int64** | Timestamp in ms to get funding from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding until INCLUSIVE. | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10 Max:100 | 
 **archived** | **string** | Default: &#x60;false&#x60;. Set to &#x60;true&#x60; for archived data from 6 months ago | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryMarginLoanRecordResponse**](QueryMarginLoanRecordResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryMarginMaxWithdraw

> QueryMarginMaxWithdrawResponse QueryMarginMaxWithdraw(ctx).Asset(asset).RecvWindow(recvWindow).Execute()

Query Margin Max Withdraw(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	asset := "asset_example" // string | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.QueryMarginMaxWithdraw(context.Background()).Asset(asset).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.QueryMarginMaxWithdraw``: %v\n", err)
		return
	}

	// response from `QueryMarginMaxWithdraw`: QueryMarginMaxWithdrawResponse
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

[**QueryMarginMaxWithdrawResponse**](QueryMarginMaxWithdrawResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryMarginRepayRecord

> QueryMarginRepayRecordResponse QueryMarginRepayRecord(ctx).Asset(asset).TxId(txId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Archived(archived).RecvWindow(recvWindow).Execute()

Query Margin repay Record(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	asset := "asset_example" // string | 
	txId := int64(1) // int64 | the `tranId` in `POST/papi/v1/marginLoan` (optional)
	startTime := int64(1623319461670) // int64 | Timestamp in ms to get funding from INCLUSIVE. (optional)
	endTime := int64(1641782889000) // int64 | Timestamp in ms to get funding until INCLUSIVE. (optional)
	current := int64(1) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(10) // int64 | Default:10 Max:100 (optional)
	archived := "archived_example" // string | Default: `false`. Set to `true` for archived data from 6 months ago (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.QueryMarginRepayRecord(context.Background()).Asset(asset).TxId(txId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Archived(archived).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.QueryMarginRepayRecord``: %v\n", err)
		return
	}

	// response from `QueryMarginRepayRecord`: QueryMarginRepayRecordResponse
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
 **txId** | **int64** | the &#x60;tranId&#x60; in &#x60;POST/papi/v1/marginLoan&#x60; | 
 **startTime** | **int64** | Timestamp in ms to get funding from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding until INCLUSIVE. | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10 Max:100 | 
 **archived** | **string** | Default: &#x60;false&#x60;. Set to &#x60;true&#x60; for archived data from 6 months ago | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryMarginRepayRecordResponse**](QueryMarginRepayRecordResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryPortfolioMarginNegativeBalanceInterestHistory

> QueryPortfolioMarginNegativeBalanceInterestHistoryResponse QueryPortfolioMarginNegativeBalanceInterestHistory(ctx).Asset(asset).StartTime(startTime).EndTime(endTime).Size(size).RecvWindow(recvWindow).Execute()

Query Portfolio Margin Negative Balance Interest History(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	asset := "asset_example" // string |  (optional)
	startTime := int64(1623319461670) // int64 | Timestamp in ms to get funding from INCLUSIVE. (optional)
	endTime := int64(1641782889000) // int64 | Timestamp in ms to get funding until INCLUSIVE. (optional)
	size := int64(10) // int64 | Default:10 Max:100 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.QueryPortfolioMarginNegativeBalanceInterestHistory(context.Background()).Asset(asset).StartTime(startTime).EndTime(endTime).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.QueryPortfolioMarginNegativeBalanceInterestHistory``: %v\n", err)
		return
	}

	// response from `QueryPortfolioMarginNegativeBalanceInterestHistory`: QueryPortfolioMarginNegativeBalanceInterestHistoryResponse
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
 **startTime** | **int64** | Timestamp in ms to get funding from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding until INCLUSIVE. | 
 **size** | **int64** | Default:10 Max:100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryPortfolioMarginNegativeBalanceInterestHistoryResponse**](QueryPortfolioMarginNegativeBalanceInterestHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryUmPositionInformation

> QueryUmPositionInformationResponse QueryUmPositionInformation(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query UM Position Information(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
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
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.QueryUmPositionInformation(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.QueryUmPositionInformation``: %v\n", err)
		return
	}

	// response from `QueryUmPositionInformation`: QueryUmPositionInformationResponse
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

[**QueryUmPositionInformationResponse**](QueryUmPositionInformationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryUserNegativeBalanceAutoExchangeRecord

> QueryUserNegativeBalanceAutoExchangeRecordResponse QueryUserNegativeBalanceAutoExchangeRecord(ctx).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

Query User Negative Balance Auto Exchange Record (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	startTime := int64(1623319461670) // int64 | 
	endTime := int64(1641782889000) // int64 | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.QueryUserNegativeBalanceAutoExchangeRecord(context.Background()).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.QueryUserNegativeBalanceAutoExchangeRecord``: %v\n", err)
		return
	}

	// response from `QueryUserNegativeBalanceAutoExchangeRecord`: QueryUserNegativeBalanceAutoExchangeRecordResponse
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
 **recvWindow** | **int64** |  | 

### Return type

[**QueryUserNegativeBalanceAutoExchangeRecordResponse**](QueryUserNegativeBalanceAutoExchangeRecordResponse.md)

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

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


## RepayFuturesNegativeBalance

> RepayFuturesNegativeBalanceResponse RepayFuturesNegativeBalance(ctx).RecvWindow(recvWindow).Execute()

Repay futures Negative Balance(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.RepayFuturesNegativeBalance(context.Background()).RecvWindow(recvWindow).Execute()
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
 **recvWindow** | **int64** |  | 

### Return type

[**RepayFuturesNegativeBalanceResponse**](RepayFuturesNegativeBalanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## UmFuturesAccountConfiguration

> UmFuturesAccountConfigurationResponse UmFuturesAccountConfiguration(ctx).RecvWindow(recvWindow).Execute()

UM Futures Account Configuration(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.UmFuturesAccountConfiguration(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.UmFuturesAccountConfiguration``: %v\n", err)
		return
	}

	// response from `UmFuturesAccountConfiguration`: UmFuturesAccountConfigurationResponse
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

[**UmFuturesAccountConfigurationResponse**](UmFuturesAccountConfigurationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## UmFuturesSymbolConfiguration

> UmFuturesSymbolConfigurationResponse UmFuturesSymbolConfiguration(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

UM Futures Symbol Configuration(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
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
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.UmFuturesSymbolConfiguration(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.UmFuturesSymbolConfiguration``: %v\n", err)
		return
	}

	// response from `UmFuturesSymbolConfiguration`: UmFuturesSymbolConfigurationResponse
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

[**UmFuturesSymbolConfigurationResponse**](UmFuturesSymbolConfigurationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## UmNotionalAndLeverageBrackets

> UmNotionalAndLeverageBracketsResponse UmNotionalAndLeverageBrackets(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

UM Notional and Leverage Brackets (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
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
	apiClient := models.NewBinanceDerivativesTradingPortfolioMarginClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.UmNotionalAndLeverageBrackets(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.UmNotionalAndLeverageBrackets``: %v\n", err)
		return
	}

	// response from `UmNotionalAndLeverageBrackets`: UmNotionalAndLeverageBracketsResponse
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

[**UmNotionalAndLeverageBracketsResponse**](UmNotionalAndLeverageBracketsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

