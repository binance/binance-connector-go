# \TransferAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**GetCrossMarginTransferHistory**](TransferAPI.md#GetCrossMarginTransferHistory) | **Get** /sapi/v1/margin/transfer | Get Cross Margin Transfer History (USER_DATA)
[**QueryMaxTransferOutAmount**](TransferAPI.md#QueryMaxTransferOutAmount) | **Get** /sapi/v1/margin/maxTransferable | Query Max Transfer-Out Amount (USER_DATA)


## GetCrossMarginTransferHistory

> GetCrossMarginTransferHistoryResponse GetCrossMarginTransferHistory(ctx).Asset(asset).Type(type_).StartTime(startTime).EndTime(endTime).Current(current).Size(size).IsolatedSymbol(isolatedSymbol).RecvWindow(recvWindow).Execute()

Get Cross Margin Transfer History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	asset := "asset_example" // string |  (optional)
	type_ := "type__example" // string | Transfer Type: ROLL_IN, ROLL_OUT (optional)
	startTime := int64(1623319461670) // int64 | 只支持查询最近90天的数据 (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	current := int64(1) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(10) // int64 | Default:10 Max:100 (optional)
	isolatedSymbol := "isolatedSymbol_example" // string | isolated symbol (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TransferAPI.GetCrossMarginTransferHistory(context.Background()).Asset(asset).Type(type_).StartTime(startTime).EndTime(endTime).Current(current).Size(size).IsolatedSymbol(isolatedSymbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TransferAPI.GetCrossMarginTransferHistory``: %v\n", err)
		return
	}

	// response from `GetCrossMarginTransferHistory`: GetCrossMarginTransferHistoryResponse
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
 **type_** | **string** | Transfer Type: ROLL_IN, ROLL_OUT | 
 **startTime** | **int64** | 只支持查询最近90天的数据 | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10 Max:100 | 
 **isolatedSymbol** | **string** | isolated symbol | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**GetCrossMarginTransferHistoryResponse**](GetCrossMarginTransferHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryMaxTransferOutAmount

> QueryMaxTransferOutAmountResponse QueryMaxTransferOutAmount(ctx).Asset(asset).IsolatedSymbol(isolatedSymbol).RecvWindow(recvWindow).Execute()

Query Max Transfer-Out Amount (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	asset := "asset_example" // string | 
	isolatedSymbol := "isolatedSymbol_example" // string | isolated symbol (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TransferAPI.QueryMaxTransferOutAmount(context.Background()).Asset(asset).IsolatedSymbol(isolatedSymbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TransferAPI.QueryMaxTransferOutAmount``: %v\n", err)
		return
	}

	// response from `QueryMaxTransferOutAmount`: QueryMaxTransferOutAmountResponse
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
 **isolatedSymbol** | **string** | isolated symbol | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**QueryMaxTransferOutAmountResponse**](QueryMaxTransferOutAmountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

