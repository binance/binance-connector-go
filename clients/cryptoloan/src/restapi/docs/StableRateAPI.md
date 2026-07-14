# \StableRateAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**GetCryptoLoansIncomeHistory**](StableRateAPI.md#GetCryptoLoansIncomeHistory) | **Get** /sapi/v1/loan/income | Get Crypto Loans Income History (USER_DATA)
[**GetLoanBorrowHistory**](StableRateAPI.md#GetLoanBorrowHistory) | **Get** /sapi/v1/loan/borrow/history | Get Loan Borrow History (USER_DATA)
[**GetLoanLtvAdjustmentHistory**](StableRateAPI.md#GetLoanLtvAdjustmentHistory) | **Get** /sapi/v1/loan/ltv/adjustment/history | Get Loan LTV Adjustment History (USER_DATA)
[**GetLoanRepaymentHistory**](StableRateAPI.md#GetLoanRepaymentHistory) | **Get** /sapi/v1/loan/repay/history | Get Loan Repayment History (USER_DATA)


## GetCryptoLoansIncomeHistory

> GetCryptoLoansIncomeHistoryResponse GetCryptoLoansIncomeHistory(ctx).Asset(asset).Type(type_).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Get Crypto Loans Income History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/cryptoloan"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	asset := "BUSD" // string |  (optional)
	type_ := models.GetCryptoLoansIncomeHistoryTypeParameterBorrowin // GetCryptoLoansIncomeHistoryTypeParameter | All types will be returned by default. (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	limit := int64(10) // int64 | Number of records to return (optional)
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

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
 **type_** | [**GetCryptoLoansIncomeHistoryTypeParameter**](GetCryptoLoansIncomeHistoryTypeParameter.md) | All types will be returned by default. | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int64** | Number of records to return | 
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**GetCryptoLoansIncomeHistoryResponse**](GetCryptoLoansIncomeHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetLoanBorrowHistory

> GetLoanBorrowHistoryResponse GetLoanBorrowHistory(ctx).OrderId(orderId).LoanCoin(loanCoin).CollateralCoin(collateralCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()

Get Loan Borrow History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/cryptoloan"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	orderId := int64(1) // int64 | orderId in `POST /sapi/v1/loan/borrow` (optional)
	loanCoin := "BUSD" // string |  (optional)
	collateralCoin := "BNB" // string |  (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	current := int64(1) // int64 | Current querying page (optional)
	limit := int64(10) // int64 | Number of records to return (optional)
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

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
 **current** | **int64** | Current querying page | 
 **limit** | **int64** | Number of records to return | 
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**GetLoanBorrowHistoryResponse**](GetLoanBorrowHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetLoanLtvAdjustmentHistory

> GetLoanLtvAdjustmentHistoryResponse GetLoanLtvAdjustmentHistory(ctx).OrderId(orderId).LoanCoin(loanCoin).CollateralCoin(collateralCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()

Get Loan LTV Adjustment History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/cryptoloan"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	orderId := int64(1) // int64 | orderId in `POST /sapi/v1/loan/borrow` (optional)
	loanCoin := "BUSD" // string |  (optional)
	collateralCoin := "BNB" // string |  (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	current := int64(1) // int64 | Current querying page (optional)
	limit := int64(10) // int64 | Number of records to return (optional)
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

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
 **current** | **int64** | Current querying page | 
 **limit** | **int64** | Number of records to return | 
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**GetLoanLtvAdjustmentHistoryResponse**](GetLoanLtvAdjustmentHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetLoanRepaymentHistory

> GetLoanRepaymentHistoryResponse GetLoanRepaymentHistory(ctx).OrderId(orderId).LoanCoin(loanCoin).CollateralCoin(collateralCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()

Get Loan Repayment History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/cryptoloan"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	orderId := int64(1) // int64 | orderId in `POST /sapi/v1/loan/borrow` (optional)
	loanCoin := "BUSD" // string |  (optional)
	collateralCoin := "BNB" // string |  (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	current := int64(1) // int64 | Current querying page (optional)
	limit := int64(10) // int64 | Number of records to return (optional)
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

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
 **current** | **int64** | Current querying page | 
 **limit** | **int64** | Number of records to return | 
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**GetLoanRepaymentHistoryResponse**](GetLoanRepaymentHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

