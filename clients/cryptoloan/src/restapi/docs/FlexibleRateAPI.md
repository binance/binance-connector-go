# \FlexibleRateAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**CheckCollateralRepayRate**](FlexibleRateAPI.md#CheckCollateralRepayRate) | **Get** /sapi/v2/loan/flexible/repay/rate | Check Collateral Repay Rate (USER_DATA)
[**FlexibleLoanAdjustLtv**](FlexibleRateAPI.md#FlexibleLoanAdjustLtv) | **Post** /sapi/v2/loan/flexible/adjust/ltv | Flexible Loan Adjust LTV(TRADE)
[**FlexibleLoanBorrow**](FlexibleRateAPI.md#FlexibleLoanBorrow) | **Post** /sapi/v2/loan/flexible/borrow | Flexible Loan Borrow(TRADE)
[**FlexibleLoanRepay**](FlexibleRateAPI.md#FlexibleLoanRepay) | **Post** /sapi/v2/loan/flexible/repay | Flexible Loan Repay(TRADE)
[**GetFlexibleLoanAssetsData**](FlexibleRateAPI.md#GetFlexibleLoanAssetsData) | **Get** /sapi/v2/loan/flexible/loanable/data | Get Flexible Loan Assets Data(USER_DATA)
[**GetFlexibleLoanBorrowHistory**](FlexibleRateAPI.md#GetFlexibleLoanBorrowHistory) | **Get** /sapi/v2/loan/flexible/borrow/history | Get Flexible Loan Borrow History(USER_DATA)
[**GetFlexibleLoanCollateralAssetsData**](FlexibleRateAPI.md#GetFlexibleLoanCollateralAssetsData) | **Get** /sapi/v2/loan/flexible/collateral/data | Get Flexible Loan Collateral Assets Data(USER_DATA)
[**GetFlexibleLoanInterestRateHistory**](FlexibleRateAPI.md#GetFlexibleLoanInterestRateHistory) | **Get** /sapi/v2/loan/interestRateHistory | Get Flexible Loan Interest Rate History (USER_DATA)
[**GetFlexibleLoanLiquidationHistory**](FlexibleRateAPI.md#GetFlexibleLoanLiquidationHistory) | **Get** /sapi/v2/loan/flexible/liquidation/history | Get Flexible Loan Liquidation History (USER_DATA)
[**GetFlexibleLoanLtvAdjustmentHistory**](FlexibleRateAPI.md#GetFlexibleLoanLtvAdjustmentHistory) | **Get** /sapi/v2/loan/flexible/ltv/adjustment/history | Get Flexible Loan LTV Adjustment History(USER_DATA)
[**GetFlexibleLoanOngoingOrders**](FlexibleRateAPI.md#GetFlexibleLoanOngoingOrders) | **Get** /sapi/v2/loan/flexible/ongoing/orders | Get Flexible Loan Ongoing Orders(USER_DATA)
[**GetFlexibleLoanRepaymentHistory**](FlexibleRateAPI.md#GetFlexibleLoanRepaymentHistory) | **Get** /sapi/v2/loan/flexible/repay/history | Get Flexible Loan Repayment History(USER_DATA)


## CheckCollateralRepayRate

> CheckCollateralRepayRateResponse CheckCollateralRepayRate(ctx).LoanCoin(loanCoin).CollateralCoin(collateralCoin).RecvWindow(recvWindow).Execute()

Check Collateral Repay Rate (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/cryptoloan"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	loanCoin := "loanCoin_example" // string | 
	collateralCoin := "collateralCoin_example" // string | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceCryptoLoanClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleRateAPI.CheckCollateralRepayRate(context.Background()).LoanCoin(loanCoin).CollateralCoin(collateralCoin).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleRateAPI.CheckCollateralRepayRate``: %v\n", err)
		return
	}

	// response from `CheckCollateralRepayRate`: CheckCollateralRepayRateResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **loanCoin** | **string** |  | 
 **collateralCoin** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CheckCollateralRepayRateResponse**](CheckCollateralRepayRateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## FlexibleLoanAdjustLtv

> FlexibleLoanAdjustLtvResponse FlexibleLoanAdjustLtv(ctx).LoanCoin(loanCoin).CollateralCoin(collateralCoin).AdjustmentAmount(adjustmentAmount).Direction(direction).RecvWindow(recvWindow).Execute()

Flexible Loan Adjust LTV(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/cryptoloan"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	loanCoin := "loanCoin_example" // string | 
	collateralCoin := "collateralCoin_example" // string | 
	adjustmentAmount := float32(1.0) // float32 | 
	direction := "direction_example" // string | \"ADDITIONAL\", \"REDUCED\"
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceCryptoLoanClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleRateAPI.FlexibleLoanAdjustLtv(context.Background()).LoanCoin(loanCoin).CollateralCoin(collateralCoin).AdjustmentAmount(adjustmentAmount).Direction(direction).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleRateAPI.FlexibleLoanAdjustLtv``: %v\n", err)
		return
	}

	// response from `FlexibleLoanAdjustLtv`: FlexibleLoanAdjustLtvResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **loanCoin** | **string** |  | 
 **collateralCoin** | **string** |  | 
 **adjustmentAmount** | **float32** |  | 
 **direction** | **string** | \&quot;ADDITIONAL\&quot;, \&quot;REDUCED\&quot; | 
 **recvWindow** | **int64** |  | 

### Return type

[**FlexibleLoanAdjustLtvResponse**](FlexibleLoanAdjustLtvResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## FlexibleLoanBorrow

> FlexibleLoanBorrowResponse FlexibleLoanBorrow(ctx).LoanCoin(loanCoin).CollateralCoin(collateralCoin).LoanAmount(loanAmount).CollateralAmount(collateralAmount).RecvWindow(recvWindow).Execute()

Flexible Loan Borrow(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/cryptoloan"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	loanCoin := "loanCoin_example" // string | 
	collateralCoin := "collateralCoin_example" // string | 
	loanAmount := float32(1.0) // float32 | Mandatory when collateralAmount is empty (optional)
	collateralAmount := float32(1.0) // float32 | Mandatory when loanAmount is empty (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceCryptoLoanClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleRateAPI.FlexibleLoanBorrow(context.Background()).LoanCoin(loanCoin).CollateralCoin(collateralCoin).LoanAmount(loanAmount).CollateralAmount(collateralAmount).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleRateAPI.FlexibleLoanBorrow``: %v\n", err)
		return
	}

	// response from `FlexibleLoanBorrow`: FlexibleLoanBorrowResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **loanCoin** | **string** |  | 
 **collateralCoin** | **string** |  | 
 **loanAmount** | **float32** | Mandatory when collateralAmount is empty | 
 **collateralAmount** | **float32** | Mandatory when loanAmount is empty | 
 **recvWindow** | **int64** |  | 

### Return type

[**FlexibleLoanBorrowResponse**](FlexibleLoanBorrowResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## FlexibleLoanRepay

> FlexibleLoanRepayResponse FlexibleLoanRepay(ctx).LoanCoin(loanCoin).CollateralCoin(collateralCoin).RepayAmount(repayAmount).CollateralReturn(collateralReturn).FullRepayment(fullRepayment).RepaymentType(repaymentType).RecvWindow(recvWindow).Execute()

Flexible Loan Repay(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/cryptoloan"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	loanCoin := "loanCoin_example" // string | 
	collateralCoin := "collateralCoin_example" // string | 
	repayAmount := float32(1.0) // float32 | repay amount of loanCoin
	collateralReturn := true // bool | Default: TRUE. TRUE: Return extra collateral to spot account; FALSE: Keep extra collateral in the order, and lower LTV. (optional)
	fullRepayment := false // bool | Default: FALSE. TRUE: Full repayment; FALSE: Partial repayment, based on loanAmount (optional)
	repaymentType := int64(1) // int64 | Default: 1. 1: Repayment with loan asset; 2: Repayment with collateral (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceCryptoLoanClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleRateAPI.FlexibleLoanRepay(context.Background()).LoanCoin(loanCoin).CollateralCoin(collateralCoin).RepayAmount(repayAmount).CollateralReturn(collateralReturn).FullRepayment(fullRepayment).RepaymentType(repaymentType).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleRateAPI.FlexibleLoanRepay``: %v\n", err)
		return
	}

	// response from `FlexibleLoanRepay`: FlexibleLoanRepayResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **loanCoin** | **string** |  | 
 **collateralCoin** | **string** |  | 
 **repayAmount** | **float32** | repay amount of loanCoin | 
 **collateralReturn** | **bool** | Default: TRUE. TRUE: Return extra collateral to spot account; FALSE: Keep extra collateral in the order, and lower LTV. | 
 **fullRepayment** | **bool** | Default: FALSE. TRUE: Full repayment; FALSE: Partial repayment, based on loanAmount | 
 **repaymentType** | **int64** | Default: 1. 1: Repayment with loan asset; 2: Repayment with collateral | 
 **recvWindow** | **int64** |  | 

### Return type

[**FlexibleLoanRepayResponse**](FlexibleLoanRepayResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetFlexibleLoanAssetsData

> GetFlexibleLoanAssetsDataResponse GetFlexibleLoanAssetsData(ctx).LoanCoin(loanCoin).RecvWindow(recvWindow).Execute()

Get Flexible Loan Assets Data(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/cryptoloan"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	loanCoin := "loanCoin_example" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceCryptoLoanClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleRateAPI.GetFlexibleLoanAssetsData(context.Background()).LoanCoin(loanCoin).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleRateAPI.GetFlexibleLoanAssetsData``: %v\n", err)
		return
	}

	// response from `GetFlexibleLoanAssetsData`: GetFlexibleLoanAssetsDataResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **loanCoin** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetFlexibleLoanAssetsDataResponse**](GetFlexibleLoanAssetsDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetFlexibleLoanBorrowHistory

> GetFlexibleLoanBorrowHistoryResponse GetFlexibleLoanBorrowHistory(ctx).LoanCoin(loanCoin).CollateralCoin(collateralCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()

Get Flexible Loan Borrow History(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/cryptoloan"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	loanCoin := "loanCoin_example" // string |  (optional)
	collateralCoin := "collateralCoin_example" // string |  (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	current := int64(1) // int64 | Current querying page. Start from 1; default: 1; max: 1000 (optional)
	limit := int64(10) // int64 | Default: 10; max: 100 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceCryptoLoanClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleRateAPI.GetFlexibleLoanBorrowHistory(context.Background()).LoanCoin(loanCoin).CollateralCoin(collateralCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleRateAPI.GetFlexibleLoanBorrowHistory``: %v\n", err)
		return
	}

	// response from `GetFlexibleLoanBorrowHistory`: GetFlexibleLoanBorrowHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **loanCoin** | **string** |  | 
 **collateralCoin** | **string** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Current querying page. Start from 1; default: 1; max: 1000 | 
 **limit** | **int64** | Default: 10; max: 100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetFlexibleLoanBorrowHistoryResponse**](GetFlexibleLoanBorrowHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetFlexibleLoanCollateralAssetsData

> GetFlexibleLoanCollateralAssetsDataResponse GetFlexibleLoanCollateralAssetsData(ctx).CollateralCoin(collateralCoin).RecvWindow(recvWindow).Execute()

Get Flexible Loan Collateral Assets Data(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/cryptoloan"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	collateralCoin := "collateralCoin_example" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceCryptoLoanClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleRateAPI.GetFlexibleLoanCollateralAssetsData(context.Background()).CollateralCoin(collateralCoin).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleRateAPI.GetFlexibleLoanCollateralAssetsData``: %v\n", err)
		return
	}

	// response from `GetFlexibleLoanCollateralAssetsData`: GetFlexibleLoanCollateralAssetsDataResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **collateralCoin** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetFlexibleLoanCollateralAssetsDataResponse**](GetFlexibleLoanCollateralAssetsDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetFlexibleLoanInterestRateHistory

> GetFlexibleLoanInterestRateHistoryResponse GetFlexibleLoanInterestRateHistory(ctx).Coin(coin).RecvWindow(recvWindow).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).Execute()

Get Flexible Loan Interest Rate History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/cryptoloan"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	coin := "coin_example" // string | 
	recvWindow := int64(5000) // int64 | 
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	current := int64(1) // int64 | Current querying page. Start from 1; default: 1; max: 1000 (optional)
	limit := int64(10) // int64 | Default: 10; max: 100 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceCryptoLoanClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleRateAPI.GetFlexibleLoanInterestRateHistory(context.Background()).Coin(coin).RecvWindow(recvWindow).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleRateAPI.GetFlexibleLoanInterestRateHistory``: %v\n", err)
		return
	}

	// response from `GetFlexibleLoanInterestRateHistory`: GetFlexibleLoanInterestRateHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **coin** | **string** |  | 
 **recvWindow** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Current querying page. Start from 1; default: 1; max: 1000 | 
 **limit** | **int64** | Default: 10; max: 100 | 

### Return type

[**GetFlexibleLoanInterestRateHistoryResponse**](GetFlexibleLoanInterestRateHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetFlexibleLoanLiquidationHistory

> GetFlexibleLoanLiquidationHistoryResponse GetFlexibleLoanLiquidationHistory(ctx).LoanCoin(loanCoin).CollateralCoin(collateralCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()

Get Flexible Loan Liquidation History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/cryptoloan"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	loanCoin := "loanCoin_example" // string |  (optional)
	collateralCoin := "collateralCoin_example" // string |  (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	current := int64(1) // int64 | Current querying page. Start from 1; default: 1; max: 1000 (optional)
	limit := int64(10) // int64 | Default: 10; max: 100 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceCryptoLoanClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleRateAPI.GetFlexibleLoanLiquidationHistory(context.Background()).LoanCoin(loanCoin).CollateralCoin(collateralCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleRateAPI.GetFlexibleLoanLiquidationHistory``: %v\n", err)
		return
	}

	// response from `GetFlexibleLoanLiquidationHistory`: GetFlexibleLoanLiquidationHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **loanCoin** | **string** |  | 
 **collateralCoin** | **string** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Current querying page. Start from 1; default: 1; max: 1000 | 
 **limit** | **int64** | Default: 10; max: 100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetFlexibleLoanLiquidationHistoryResponse**](GetFlexibleLoanLiquidationHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetFlexibleLoanLtvAdjustmentHistory

> GetFlexibleLoanLtvAdjustmentHistoryResponse GetFlexibleLoanLtvAdjustmentHistory(ctx).LoanCoin(loanCoin).CollateralCoin(collateralCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()

Get Flexible Loan LTV Adjustment History(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/cryptoloan"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	loanCoin := "loanCoin_example" // string |  (optional)
	collateralCoin := "collateralCoin_example" // string |  (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	current := int64(1) // int64 | Current querying page. Start from 1; default: 1; max: 1000 (optional)
	limit := int64(10) // int64 | Default: 10; max: 100 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceCryptoLoanClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleRateAPI.GetFlexibleLoanLtvAdjustmentHistory(context.Background()).LoanCoin(loanCoin).CollateralCoin(collateralCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleRateAPI.GetFlexibleLoanLtvAdjustmentHistory``: %v\n", err)
		return
	}

	// response from `GetFlexibleLoanLtvAdjustmentHistory`: GetFlexibleLoanLtvAdjustmentHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **loanCoin** | **string** |  | 
 **collateralCoin** | **string** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Current querying page. Start from 1; default: 1; max: 1000 | 
 **limit** | **int64** | Default: 10; max: 100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetFlexibleLoanLtvAdjustmentHistoryResponse**](GetFlexibleLoanLtvAdjustmentHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetFlexibleLoanOngoingOrders

> GetFlexibleLoanOngoingOrdersResponse GetFlexibleLoanOngoingOrders(ctx).LoanCoin(loanCoin).CollateralCoin(collateralCoin).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()

Get Flexible Loan Ongoing Orders(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/cryptoloan"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	loanCoin := "loanCoin_example" // string |  (optional)
	collateralCoin := "collateralCoin_example" // string |  (optional)
	current := int64(1) // int64 | Current querying page. Start from 1; default: 1; max: 1000 (optional)
	limit := int64(10) // int64 | Default: 10; max: 100 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceCryptoLoanClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleRateAPI.GetFlexibleLoanOngoingOrders(context.Background()).LoanCoin(loanCoin).CollateralCoin(collateralCoin).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleRateAPI.GetFlexibleLoanOngoingOrders``: %v\n", err)
		return
	}

	// response from `GetFlexibleLoanOngoingOrders`: GetFlexibleLoanOngoingOrdersResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **loanCoin** | **string** |  | 
 **collateralCoin** | **string** |  | 
 **current** | **int64** | Current querying page. Start from 1; default: 1; max: 1000 | 
 **limit** | **int64** | Default: 10; max: 100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetFlexibleLoanOngoingOrdersResponse**](GetFlexibleLoanOngoingOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetFlexibleLoanRepaymentHistory

> GetFlexibleLoanRepaymentHistoryResponse GetFlexibleLoanRepaymentHistory(ctx).LoanCoin(loanCoin).CollateralCoin(collateralCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()

Get Flexible Loan Repayment History(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/cryptoloan"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	loanCoin := "loanCoin_example" // string |  (optional)
	collateralCoin := "collateralCoin_example" // string |  (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	current := int64(1) // int64 | Current querying page. Start from 1; default: 1; max: 1000 (optional)
	limit := int64(10) // int64 | Default: 10; max: 100 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceCryptoLoanClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FlexibleRateAPI.GetFlexibleLoanRepaymentHistory(context.Background()).LoanCoin(loanCoin).CollateralCoin(collateralCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FlexibleRateAPI.GetFlexibleLoanRepaymentHistory``: %v\n", err)
		return
	}

	// response from `GetFlexibleLoanRepaymentHistory`: GetFlexibleLoanRepaymentHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **loanCoin** | **string** |  | 
 **collateralCoin** | **string** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Current querying page. Start from 1; default: 1; max: 1000 | 
 **limit** | **int64** | Default: 10; max: 100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetFlexibleLoanRepaymentHistoryResponse**](GetFlexibleLoanRepaymentHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

