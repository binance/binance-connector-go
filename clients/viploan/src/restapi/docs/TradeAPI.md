# \TradeAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**VipLoanBorrow**](TradeAPI.md#VipLoanBorrow) | **Post** /sapi/v1/loan/vip/borrow | VIP Loan Borrow (TRADE)
[**VipLoanFixedRateBorrow**](TradeAPI.md#VipLoanFixedRateBorrow) | **Post** /sapi/v1/loan/vip/fixed/borrow | VIP Loan Fixed Rate Borrow (TRADE)
[**VipLoanRenew**](TradeAPI.md#VipLoanRenew) | **Post** /sapi/v1/loan/vip/renew | VIP Loan Renew (TRADE)
[**VipLoanRepay**](TradeAPI.md#VipLoanRepay) | **Post** /sapi/v1/loan/vip/repay | VIP Loan Repay (TRADE)


## VipLoanBorrow

> VipLoanBorrowResponse VipLoanBorrow(ctx).LoanAccountId(loanAccountId).LoanCoin(loanCoin).LoanAmount(loanAmount).CollateralAccountId(collateralAccountId).CollateralCoin(collateralCoin).IsFlexibleRate(isFlexibleRate).LoanTerm(loanTerm).RecvWindow(recvWindow).Execute()

VIP Loan Borrow (TRADE)


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
	loanCoin := "BTC" // string | 
	loanAmount := float64(1.0) // float64 | 
	collateralAccountId := "12345678,12345678,12345678" // string | Collateral account ID(s). Multiple split by `,`
	collateralCoin := "BUSD,USDT,ETH" // string | 
	isFlexibleRate := true // bool | TRUE: flexible rate; FALSE: fixed rate
	loanTerm := int64(30) // int64 | Mandatory for fixed rate. Optional for flexible rate. e.g. 30/60 days (optional)
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
 **loanAmount** | **float64** |  | 
 **collateralAccountId** | **string** | Collateral account ID(s). Multiple split by &#x60;,&#x60; | 
 **collateralCoin** | **string** |  | 
 **isFlexibleRate** | **bool** | TRUE: flexible rate; FALSE: fixed rate | 
 **loanTerm** | **int64** | Mandatory for fixed rate. Optional for flexible rate. e.g. 30/60 days | 
 **recvWindow** | **int64** |  | 

### Return type

[**VipLoanBorrowResponse**](VipLoanBorrowResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## VipLoanFixedRateBorrow

> VipLoanFixedRateBorrowResponse VipLoanFixedRateBorrow(ctx).SupplyRequest(supplyRequest).BorrowCoin(borrowCoin).LoanTerm(loanTerm).BorrowUid(borrowUid).CollateralCoin(collateralCoin).CollateralAccountId(collateralAccountId).AutoRepay(autoRepay).RecvWindow(recvWindow).Execute()

VIP Loan Fixed Rate Borrow (TRADE)


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
	supplyRequest := "1212:0.12:100;3434:0.13:50" // string | Supply request string, positional encoding (no key). Multiple entries separated by `;`, fields separated by `:`, order: `<requestId>:<interestRate>:<amount>`. Example: `1212:0.12:100;3434:0.13:50`
	borrowCoin := "BUSD" // string | Borrow coin
	loanTerm := int64(30) // int64 | Loan term in days
	borrowUid := int64(12345678) // int64 | Borrow receiving account UID
	collateralCoin := "BNB,ETH,BTC" // string | Collateral coin(s), multiple separated by `,`. Only coin names, no amount (VIP loan collateral amount = entire spot account balance)
	collateralAccountId := "12345,67890,13579" // string | Collateral account ID(s), multiple separated by `,`
	autoRepay := true // bool | Default: `true`. `true`: auto repay at expiration; `false`: auto-convert to flexible (floating rate) at expiration (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than `60000` (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceVipLoanClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.VipLoanFixedRateBorrow(context.Background()).SupplyRequest(supplyRequest).BorrowCoin(borrowCoin).LoanTerm(loanTerm).BorrowUid(borrowUid).CollateralCoin(collateralCoin).CollateralAccountId(collateralAccountId).AutoRepay(autoRepay).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.VipLoanFixedRateBorrow``: %v\n", err)
		return
	}

	// response from `VipLoanFixedRateBorrow`: VipLoanFixedRateBorrowResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **supplyRequest** | **string** | Supply request string, positional encoding (no key). Multiple entries separated by &#x60;;&#x60;, fields separated by &#x60;:&#x60;, order: &#x60;&lt;requestId&gt;:&lt;interestRate&gt;:&lt;amount&gt;&#x60;. Example: &#x60;1212:0.12:100;3434:0.13:50&#x60; | 
 **borrowCoin** | **string** | Borrow coin | 
 **loanTerm** | **int64** | Loan term in days | 
 **borrowUid** | **int64** | Borrow receiving account UID | 
 **collateralCoin** | **string** | Collateral coin(s), multiple separated by &#x60;,&#x60;. Only coin names, no amount (VIP loan collateral amount &#x3D; entire spot account balance) | 
 **collateralAccountId** | **string** | Collateral account ID(s), multiple separated by &#x60;,&#x60; | 
 **autoRepay** | **bool** | Default: &#x60;true&#x60;. &#x60;true&#x60;: auto repay at expiration; &#x60;false&#x60;: auto-convert to flexible (floating rate) at expiration | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**VipLoanFixedRateBorrowResponse**](VipLoanFixedRateBorrowResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## VipLoanRenew

> VipLoanRenewResponse VipLoanRenew(ctx).OrderId(orderId).LoanTerm(loanTerm).RecvWindow(recvWindow).Execute()

VIP Loan Renew (TRADE)


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
	loanTerm := int64(30) // int64 | 30/60 days
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

VIP Loan Repay (TRADE)


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
	amount := float64(1.0) // float64 | 
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
 **amount** | **float64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**VipLoanRepayResponse**](VipLoanRepayResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

