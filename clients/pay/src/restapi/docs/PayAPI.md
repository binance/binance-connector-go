# \PayAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**GetPayTradeHistory**](PayAPI.md#GetPayTradeHistory) | **Get** /sapi/v1/pay/transactions | Get Pay Trade History


## GetPayTradeHistory

> GetPayTradeHistoryResponse GetPayTradeHistory(ctx).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Get Pay Trade History


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/pay"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	limit := int64(100) // int64 | default 100, max 100 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinancePayClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.PayAPI.GetPayTradeHistory(context.Background()).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `PayAPI.GetPayTradeHistory``: %v\n", err)
		return
	}

	// response from `GetPayTradeHistory`: GetPayTradeHistoryResponse
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
 **limit** | **int64** | default 100, max 100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetPayTradeHistoryResponse**](GetPayTradeHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

