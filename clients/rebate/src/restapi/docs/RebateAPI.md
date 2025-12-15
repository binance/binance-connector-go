# \RebateAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**GetSpotRebateHistoryRecords**](RebateAPI.md#GetSpotRebateHistoryRecords) | **Get** /sapi/v1/rebate/taxQuery | Get Spot Rebate History Records (USER_DATA)


## GetSpotRebateHistoryRecords

> GetSpotRebateHistoryRecordsResponse GetSpotRebateHistoryRecords(ctx).StartTime(startTime).EndTime(endTime).Page(page).RecvWindow(recvWindow).Execute()

Get Spot Rebate History Records (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/rebate"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	page := int64(1) // int64 | Default 1 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceRebateClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.RebateAPI.GetSpotRebateHistoryRecords(context.Background()).StartTime(startTime).EndTime(endTime).Page(page).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `RebateAPI.GetSpotRebateHistoryRecords``: %v\n", err)
		return
	}

	// response from `GetSpotRebateHistoryRecords`: GetSpotRebateHistoryRecordsResponse
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
 **page** | **int64** | Default 1 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSpotRebateHistoryRecordsResponse**](GetSpotRebateHistoryRecordsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

