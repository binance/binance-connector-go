# \TradeAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**VipLoanBorrow**](TradeAPI.md#VipLoanBorrow) | **Post** /sapi/v1/loan/vip/borrow | VIP Loan Borrow(TRADE)
[**VipLoanRenew**](TradeAPI.md#VipLoanRenew) | **Post** /sapi/v1/loan/vip/renew | VIP Loan Renew(TRADE)
[**VipLoanRepay**](TradeAPI.md#VipLoanRepay) | **Post** /sapi/v1/loan/vip/repay | VIP Loan Repay(TRADE)


## VipLoanBorrow

> VipLoanBorrowResponse VipLoanBorrow(ctx).LoanAccountId(loanAccountId).LoanCoin(loanCoin).LoanAmount(loanAmount).CollateralAccountId(collateralAccountId).CollateralCoin(collateralCoin).IsFlexibleRate(isFlexibleRate).LoanTerm(loanTerm).RecvWindow(recvWindow).Execute()

VIP Loan Borrow(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/viploan"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	loanAccountId := int64(1) // int64 | 
	loanCoin := "loanCoin_example" // string | 
	loanAmount := float32(1.0) // float32 | 
	collateralAccountId := "1" // string | Multiple split by `,`
	collateralCoin := "collateralCoin_example" // string | Multiple split by `,`
	isFlexibleRate := true // bool | Default: TRUE. TRUE : flexible rate; FALSE: fixed rate
	loanTerm := int64(789) // int64 | Mandatory for fixed rate. Optional for fixed interest rate. Eg: 30/60 days (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceVipLoanClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.VipLoanBorrow(context.Background()).LoanAccountId(loanAccountId).LoanCoin(loanCoin).LoanAmount(loanAmount).CollateralAccountId(collateralAccountId).CollateralCoin(collateralCoin).IsFlexibleRate(isFlexibleRate).LoanTerm(loanTerm).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.VipLoanBorrow``: %v\n", err)
		return
	}

	// response from `VipLoanBorrow`: VipLoanBorrowResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **loanAccountId** | **int64** |  | 
 **loanCoin** | **string** |  | 
 **loanAmount** | **float32** |  | 
 **collateralAccountId** | **string** | Multiple split by &#x60;,&#x60; | 
 **collateralCoin** | **string** | Multiple split by &#x60;,&#x60; | 
 **isFlexibleRate** | **bool** | Default: TRUE. TRUE : flexible rate; FALSE: fixed rate | 
 **loanTerm** | **int64** | Mandatory for fixed rate. Optional for fixed interest rate. Eg: 30/60 days | 
 **recvWindow** | **int64** |  | 

### Return type

[**VipLoanBorrowResponse**](VipLoanBorrowResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## VipLoanRenew

> VipLoanRenewResponse VipLoanRenew(ctx).OrderId(orderId).LoanTerm(loanTerm).RecvWindow(recvWindow).Execute()

VIP Loan Renew(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/viploan"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	orderId := int64(1) // int64 | 
	loanTerm := int64(789) // int64 | 30/60 days
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceVipLoanClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.VipLoanRenew(context.Background()).OrderId(orderId).LoanTerm(loanTerm).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.VipLoanRenew``: %v\n", err)
		return
	}

	// response from `VipLoanRenew`: VipLoanRenewResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **int64** |  | 
 **loanTerm** | **int64** | 30/60 days | 
 **recvWindow** | **int64** |  | 

### Return type

[**VipLoanRenewResponse**](VipLoanRenewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## VipLoanRepay

> VipLoanRepayResponse VipLoanRepay(ctx).OrderId(orderId).Amount(amount).RecvWindow(recvWindow).Execute()

VIP Loan Repay(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/viploan"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	orderId := int64(1) // int64 | 
	amount := float32(1.0) // float32 | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceVipLoanClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.VipLoanRepay(context.Background()).OrderId(orderId).Amount(amount).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.VipLoanRepay``: %v\n", err)
		return
	}

	// response from `VipLoanRepay`: VipLoanRepayResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **int64** |  | 
 **amount** | **float32** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**VipLoanRepayResponse**](VipLoanRepayResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

