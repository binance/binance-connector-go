# \StableRateAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**CheckCollateralRepayRateStableRate**](StableRateAPI.md#CheckCollateralRepayRateStableRate) | **Get** /sapi/v1/loan/repay/collateral/rate | Check Collateral Repay Rate(USER_DATA)
[**GetCryptoLoansIncomeHistory**](StableRateAPI.md#GetCryptoLoansIncomeHistory) | **Get** /sapi/v1/loan/income | Get Crypto Loans Income History(USER_DATA)
[**GetLoanBorrowHistory**](StableRateAPI.md#GetLoanBorrowHistory) | **Get** /sapi/v1/loan/borrow/history | Get Loan Borrow History(USER_DATA)
[**GetLoanLtvAdjustmentHistory**](StableRateAPI.md#GetLoanLtvAdjustmentHistory) | **Get** /sapi/v1/loan/ltv/adjustment/history | Get Loan LTV Adjustment History(USER_DATA)
[**GetLoanRepaymentHistory**](StableRateAPI.md#GetLoanRepaymentHistory) | **Get** /sapi/v1/loan/repay/history | Get Loan Repayment History(USER_DATA)


## CheckCollateralRepayRateStableRate

> CheckCollateralRepayRateStableRateResponse CheckCollateralRepayRateStableRate(ctx).LoanCoin(loanCoin).CollateralCoin(collateralCoin).RepayAmount(repayAmount).RecvWindow(recvWindow).Execute()

Check Collateral Repay Rate(USER_DATA)


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
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceCryptoLoanClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.StableRateAPI.CheckCollateralRepayRateStableRate(context.Background()).LoanCoin(loanCoin).CollateralCoin(collateralCoin).RepayAmount(repayAmount).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `StableRateAPI.CheckCollateralRepayRateStableRate``: %v\n", err)
		return
	}

	// response from `CheckCollateralRepayRateStableRate`: CheckCollateralRepayRateStableRateResponse
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
 **recvWindow** | **int64** |  | 

### Return type

[**CheckCollateralRepayRateStableRateResponse**](CheckCollateralRepayRateStableRateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetCryptoLoansIncomeHistory

> GetCryptoLoansIncomeHistoryResponse GetCryptoLoansIncomeHistory(ctx).Asset(asset).Type(type_).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Get Crypto Loans Income History(USER_DATA)


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
	asset := "asset_example" // string |  (optional)
	type_ := "0" // string | All types will be returned by default. Enum：`borrowIn` ,`collateralSpent`, `repayAmount`, `collateralReturn`(Collateral return after repayment), `addCollateral`, `removeCollateral`, `collateralReturnAfterLiquidation` (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	limit := int64(10) // int64 | Default: 10; max: 100 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceCryptoLoanClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.StableRateAPI.GetCryptoLoansIncomeHistory(context.Background()).Asset(asset).Type(type_).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `StableRateAPI.GetCryptoLoansIncomeHistory``: %v\n", err)
		return
	}

	// response from `GetCryptoLoansIncomeHistory`: GetCryptoLoansIncomeHistoryResponse
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
 **type_** | **string** | All types will be returned by default. Enum：&#x60;borrowIn&#x60; ,&#x60;collateralSpent&#x60;, &#x60;repayAmount&#x60;, &#x60;collateralReturn&#x60;(Collateral return after repayment), &#x60;addCollateral&#x60;, &#x60;removeCollateral&#x60;, &#x60;collateralReturnAfterLiquidation&#x60; | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int64** | Default: 10; max: 100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetCryptoLoansIncomeHistoryResponse**](GetCryptoLoansIncomeHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetLoanBorrowHistory

> GetLoanBorrowHistoryResponse GetLoanBorrowHistory(ctx).OrderId(orderId).LoanCoin(loanCoin).CollateralCoin(collateralCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()

Get Loan Borrow History(USER_DATA)


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
	orderId := int64(1) // int64 | orderId in `POST /sapi/v1/loan/borrow` (optional)
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

	resp, err := apiClient.RestApi.StableRateAPI.GetLoanBorrowHistory(context.Background()).OrderId(orderId).LoanCoin(loanCoin).CollateralCoin(collateralCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `StableRateAPI.GetLoanBorrowHistory``: %v\n", err)
		return
	}

	// response from `GetLoanBorrowHistory`: GetLoanBorrowHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **int64** | orderId in &#x60;POST /sapi/v1/loan/borrow&#x60; | 
 **loanCoin** | **string** |  | 
 **collateralCoin** | **string** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Current querying page. Start from 1; default: 1; max: 1000 | 
 **limit** | **int64** | Default: 10; max: 100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetLoanBorrowHistoryResponse**](GetLoanBorrowHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetLoanLtvAdjustmentHistory

> GetLoanLtvAdjustmentHistoryResponse GetLoanLtvAdjustmentHistory(ctx).OrderId(orderId).LoanCoin(loanCoin).CollateralCoin(collateralCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()

Get Loan LTV Adjustment History(USER_DATA)


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
	orderId := int64(1) // int64 | orderId in `POST /sapi/v1/loan/borrow` (optional)
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

	resp, err := apiClient.RestApi.StableRateAPI.GetLoanLtvAdjustmentHistory(context.Background()).OrderId(orderId).LoanCoin(loanCoin).CollateralCoin(collateralCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `StableRateAPI.GetLoanLtvAdjustmentHistory``: %v\n", err)
		return
	}

	// response from `GetLoanLtvAdjustmentHistory`: GetLoanLtvAdjustmentHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **int64** | orderId in &#x60;POST /sapi/v1/loan/borrow&#x60; | 
 **loanCoin** | **string** |  | 
 **collateralCoin** | **string** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Current querying page. Start from 1; default: 1; max: 1000 | 
 **limit** | **int64** | Default: 10; max: 100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetLoanLtvAdjustmentHistoryResponse**](GetLoanLtvAdjustmentHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetLoanRepaymentHistory

> GetLoanRepaymentHistoryResponse GetLoanRepaymentHistory(ctx).OrderId(orderId).LoanCoin(loanCoin).CollateralCoin(collateralCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()

Get Loan Repayment History(USER_DATA)


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
	orderId := int64(1) // int64 | orderId in `POST /sapi/v1/loan/borrow` (optional)
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

	resp, err := apiClient.RestApi.StableRateAPI.GetLoanRepaymentHistory(context.Background()).OrderId(orderId).LoanCoin(loanCoin).CollateralCoin(collateralCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `StableRateAPI.GetLoanRepaymentHistory``: %v\n", err)
		return
	}

	// response from `GetLoanRepaymentHistory`: GetLoanRepaymentHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **int64** | orderId in &#x60;POST /sapi/v1/loan/borrow&#x60; | 
 **loanCoin** | **string** |  | 
 **collateralCoin** | **string** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Current querying page. Start from 1; default: 1; max: 1000 | 
 **limit** | **int64** | Default: 10; max: 100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetLoanRepaymentHistoryResponse**](GetLoanRepaymentHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

