# \DefaultAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**GetC2CTradeHistory**](DefaultAPI.md#GetC2CTradeHistory) | **Get** /sapi/v1/c2c/orderMatch/listUserOrderHistory | Get C2C Trade History (USER_DATA)


## GetC2CTradeHistory

> GetC2CTradeHistoryResponse GetC2CTradeHistory(ctx).TradeType(tradeType).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).Page(page).Rows(rows).Execute()

Get C2C Trade History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/c2c"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	tradeType := models.GetC2CTradeHistoryTradeTypeParameterBuy // GetC2CTradeHistoryTradeTypeParameter | Trade side filter (optional)
	startTimestamp := int64(1770736694138) // int64 |  (optional)
	endTimestamp := int64(1770736694138) // int64 |  (optional)
	page := int64(1) // int64 | Page number (optional)
	rows := int64(100) // int64 | Number of records per page (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceC2CClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.DefaultAPI.GetC2CTradeHistory(context.Background()).TradeType(tradeType).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).Page(page).Rows(rows).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.GetC2CTradeHistory``: %v\n", err)
		return
	}

	// response from `GetC2CTradeHistory`: GetC2CTradeHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **tradeType** | [**GetC2CTradeHistoryTradeTypeParameter**](GetC2CTradeHistoryTradeTypeParameter.md) | Trade side filter | 
 **startTimestamp** | **int64** |  | 
 **endTimestamp** | **int64** |  | 
 **page** | **int64** | Page number | 
 **rows** | **int64** | Number of records per page | 

### Return type

[**GetC2CTradeHistoryResponse**](GetC2CTradeHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

