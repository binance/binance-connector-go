# \UserInformationAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**CheckVIPLoanCollateralAccount**](UserInformationAPI.md#CheckVIPLoanCollateralAccount) | **Get** /sapi/v1/loan/vip/collateral/account | Check VIP Loan Collateral Account (USER_DATA)
[**GetVIPLoanOngoingOrders**](UserInformationAPI.md#GetVIPLoanOngoingOrders) | **Get** /sapi/v1/loan/vip/ongoing/orders | Get VIP Loan Ongoing Orders(USER_DATA)
[**QueryApplicationStatus**](UserInformationAPI.md#QueryApplicationStatus) | **Get** /sapi/v1/loan/vip/request/data | Query Application Status(USER_DATA)


## CheckVIPLoanCollateralAccount

> CheckVIPLoanCollateralAccountResponse CheckVIPLoanCollateralAccount(ctx).OrderId(orderId).CollateralAccountId(collateralAccountId).RecvWindow(recvWindow).Execute()

Check VIP Loan Collateral Account (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/viploan"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	orderId := int64(1) // int64 |  (optional)
	collateralAccountId := int64(1) // int64 |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceVipLoanClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.UserInformationAPI.CheckVIPLoanCollateralAccount(context.Background()).OrderId(orderId).CollateralAccountId(collateralAccountId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `UserInformationAPI.CheckVIPLoanCollateralAccount``: %v\n", err)
		return
	}

	// response from `CheckVIPLoanCollateralAccount`: CheckVIPLoanCollateralAccountResponse
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
 **collateralAccountId** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CheckVIPLoanCollateralAccountResponse**](CheckVIPLoanCollateralAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetVIPLoanOngoingOrders

> GetVIPLoanOngoingOrdersResponse GetVIPLoanOngoingOrders(ctx).OrderId(orderId).CollateralAccountId(collateralAccountId).LoanCoin(loanCoin).CollateralCoin(collateralCoin).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()

Get VIP Loan Ongoing Orders(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/viploan"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	orderId := int64(1) // int64 |  (optional)
	collateralAccountId := int64(1) // int64 |  (optional)
	loanCoin := "loanCoin_example" // string |  (optional)
	collateralCoin := "collateralCoin_example" // string |  (optional)
	current := int64(1) // int64 | Currently querying page. Start from 1, Default:1, Max: 1000. (optional)
	limit := int64(10) // int64 | Default: 10, Max: 100 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceVipLoanClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.UserInformationAPI.GetVIPLoanOngoingOrders(context.Background()).OrderId(orderId).CollateralAccountId(collateralAccountId).LoanCoin(loanCoin).CollateralCoin(collateralCoin).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `UserInformationAPI.GetVIPLoanOngoingOrders``: %v\n", err)
		return
	}

	// response from `GetVIPLoanOngoingOrders`: GetVIPLoanOngoingOrdersResponse
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
 **collateralAccountId** | **int64** |  | 
 **loanCoin** | **string** |  | 
 **collateralCoin** | **string** |  | 
 **current** | **int64** | Currently querying page. Start from 1, Default:1, Max: 1000. | 
 **limit** | **int64** | Default: 10, Max: 100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetVIPLoanOngoingOrdersResponse**](GetVIPLoanOngoingOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryApplicationStatus

> QueryApplicationStatusResponse QueryApplicationStatus(ctx).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()

Query Application Status(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/viploan"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	current := int64(1) // int64 | Currently querying page. Start from 1, Default:1, Max: 1000. (optional)
	limit := int64(10) // int64 | Default: 10, Max: 100 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceVipLoanClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.UserInformationAPI.QueryApplicationStatus(context.Background()).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `UserInformationAPI.QueryApplicationStatus``: %v\n", err)
		return
	}

	// response from `QueryApplicationStatus`: QueryApplicationStatusResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **current** | **int64** | Currently querying page. Start from 1, Default:1, Max: 1000. | 
 **limit** | **int64** | Default: 10, Max: 100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryApplicationStatusResponse**](QueryApplicationStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

