# \YieldArenaAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**GetYieldArenaActivities**](YieldArenaAPI.md#GetYieldArenaActivities) | **Get** /sapi/v1/earn/arena/activities | Get Yield Arena Activities (USER_DATA)


## GetYieldArenaActivities

> GetYieldArenaActivitiesResponse GetYieldArenaActivities(ctx).Lang(lang).RecvWindow(recvWindow).Execute()

Get Yield Arena Activities (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/simpleearn"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	lang := "en" // string | Locale tag for `title` and `description` (e.g. `en`, `zh-CN`, `pt-BR`). Default: `en`. If the value is missing, malformed, or has no translation configured, content is returned in `en`. (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSimpleEarnClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.YieldArenaAPI.GetYieldArenaActivities(context.Background()).Lang(lang).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `YieldArenaAPI.GetYieldArenaActivities``: %v\n", err)
		return
	}

	// response from `GetYieldArenaActivities`: GetYieldArenaActivitiesResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **lang** | **string** | Locale tag for &#x60;title&#x60; and &#x60;description&#x60; (e.g. &#x60;en&#x60;, &#x60;zh-CN&#x60;, &#x60;pt-BR&#x60;). Default: &#x60;en&#x60;. If the value is missing, malformed, or has no translation configured, content is returned in &#x60;en&#x60;. | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetYieldArenaActivitiesResponse**](GetYieldArenaActivitiesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

