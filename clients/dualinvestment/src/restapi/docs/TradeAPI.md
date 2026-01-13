# \TradeAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**ChangeAutoCompoundStatus**](TradeAPI.md#ChangeAutoCompoundStatus) | **Post** /sapi/v1/dci/product/auto_compound/edit-status | Change Auto-Compound status(USER_DATA)
[**CheckDualInvestmentAccounts**](TradeAPI.md#CheckDualInvestmentAccounts) | **Get** /sapi/v1/dci/product/accounts | Check Dual Investment accounts(USER_DATA)
[**GetDualInvestmentPositions**](TradeAPI.md#GetDualInvestmentPositions) | **Get** /sapi/v1/dci/product/positions | Get Dual Investment positions(USER_DATA)
[**SubscribeDualInvestmentProducts**](TradeAPI.md#SubscribeDualInvestmentProducts) | **Post** /sapi/v1/dci/product/subscribe | Subscribe Dual Investment products(USER_DATA)


## ChangeAutoCompoundStatus

> ChangeAutoCompoundStatusResponse ChangeAutoCompoundStatus(ctx).PositionId(positionId).AutoCompoundPlan(autoCompoundPlan).RecvWindow(recvWindow).Execute()

Change Auto-Compound status(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/dualinvestment"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	positionId := "1" // string | Get positionId from `/sapi/v1/dci/product/positions`
	autoCompoundPlan := "autoCompoundPlan_example" // string |  (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDualInvestmentClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.ChangeAutoCompoundStatus(context.Background()).PositionId(positionId).AutoCompoundPlan(autoCompoundPlan).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.ChangeAutoCompoundStatus``: %v\n", err)
		return
	}

	// response from `ChangeAutoCompoundStatus`: ChangeAutoCompoundStatusResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **positionId** | **string** | Get positionId from &#x60;/sapi/v1/dci/product/positions&#x60; | 
 **autoCompoundPlan** | **string** |  | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**ChangeAutoCompoundStatusResponse**](ChangeAutoCompoundStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CheckDualInvestmentAccounts

> CheckDualInvestmentAccountsResponse CheckDualInvestmentAccounts(ctx).RecvWindow(recvWindow).Execute()

Check Dual Investment accounts(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/dualinvestment"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDualInvestmentClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CheckDualInvestmentAccounts(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CheckDualInvestmentAccounts``: %v\n", err)
		return
	}

	// response from `CheckDualInvestmentAccounts`: CheckDualInvestmentAccountsResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**CheckDualInvestmentAccountsResponse**](CheckDualInvestmentAccountsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetDualInvestmentPositions

> GetDualInvestmentPositionsResponse GetDualInvestmentPositions(ctx).Status(status).PageSize(pageSize).PageIndex(pageIndex).RecvWindow(recvWindow).Execute()

Get Dual Investment positions(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/dualinvestment"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	status := "status_example" // string | `PENDING`:Products are purchasing, will give results later;`PURCHASE_SUCCESS`:purchase successfully;`SETTLED`: Products are finish settling;`PURCHASE_FAIL`:fail to purchase;`REFUNDING`:refund ongoing;`REFUND_SUCCESS`:refund to spot account successfully; `SETTLING`:Products are settling. If don't fill this field, will response all the position status. (optional)
	pageSize := int64(10) // int64 | Default: 10, Maximum: 100 (optional)
	pageIndex := int64(1) // int64 | Default: 1 (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDualInvestmentClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.GetDualInvestmentPositions(context.Background()).Status(status).PageSize(pageSize).PageIndex(pageIndex).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.GetDualInvestmentPositions``: %v\n", err)
		return
	}

	// response from `GetDualInvestmentPositions`: GetDualInvestmentPositionsResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | &#x60;PENDING&#x60;:Products are purchasing, will give results later;&#x60;PURCHASE_SUCCESS&#x60;:purchase successfully;&#x60;SETTLED&#x60;: Products are finish settling;&#x60;PURCHASE_FAIL&#x60;:fail to purchase;&#x60;REFUNDING&#x60;:refund ongoing;&#x60;REFUND_SUCCESS&#x60;:refund to spot account successfully; &#x60;SETTLING&#x60;:Products are settling. If don&#39;t fill this field, will response all the position status. | 
 **pageSize** | **int64** | Default: 10, Maximum: 100 | 
 **pageIndex** | **int64** | Default: 1 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**GetDualInvestmentPositionsResponse**](GetDualInvestmentPositionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SubscribeDualInvestmentProducts

> SubscribeDualInvestmentProductsResponse SubscribeDualInvestmentProducts(ctx).Id(id).OrderId(orderId).DepositAmount(depositAmount).AutoCompoundPlan(autoCompoundPlan).RecvWindow(recvWindow).Execute()

Subscribe Dual Investment products(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/dualinvestment"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	id := "id_example" // string | get id from `/sapi/v1/dci/product/list`
	orderId := "1" // string | get orderId from `/sapi/v1/dci/product/list`
	depositAmount := float32(1.0) // float32 | the amount for subscribing
	autoCompoundPlan := "NONE" // string | `NONE`: switch off the plan, `STANDARD`:standard plan,`ADVANCED`:advanced plan
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDualInvestmentClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.SubscribeDualInvestmentProducts(context.Background()).Id(id).OrderId(orderId).DepositAmount(depositAmount).AutoCompoundPlan(autoCompoundPlan).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.SubscribeDualInvestmentProducts``: %v\n", err)
		return
	}

	// response from `SubscribeDualInvestmentProducts`: SubscribeDualInvestmentProductsResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | get id from &#x60;/sapi/v1/dci/product/list&#x60; | 
 **orderId** | **string** | get orderId from &#x60;/sapi/v1/dci/product/list&#x60; | 
 **depositAmount** | **float32** | the amount for subscribing | 
 **autoCompoundPlan** | **string** | &#x60;NONE&#x60;: switch off the plan, &#x60;STANDARD&#x60;:standard plan,&#x60;ADVANCED&#x60;:advanced plan | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**SubscribeDualInvestmentProductsResponse**](SubscribeDualInvestmentProductsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

