# \AccountAPI

All URIs are relative to *https://eapi.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AccountFundingFlow**](AccountAPI.md#AccountFundingFlow) | **Get** /eapi/v1/bill | Account Funding Flow (USER_DATA)
[**OptionMarginAccountInformation**](AccountAPI.md#OptionMarginAccountInformation) | **Get** /eapi/v1/marginAccount | Option Margin Account Information (USER_DATA)


## AccountFundingFlow

> AccountFundingFlowResponse AccountFundingFlow(ctx).Currency(currency).RecordId(recordId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Account Funding Flow (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	currency := "currency_example" // string | Asset type, only support USDT  as of now
	recordId := int64(1) // int64 | Return the recordId and subsequent data, the latest data is returned by default, e.g 100000 (optional)
	startTime := int64(1623319461670) // int64 | Start Time, e.g 1593511200000 (optional)
	endTime := int64(1641782889000) // int64 | End Time, e.g 1593512200000 (optional)
	limit := int64(100) // int64 | Number of result sets returned Default:100 Max:1000 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.AccountFundingFlow(context.Background()).Currency(currency).RecordId(recordId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.AccountFundingFlow``: %v\n", err)
		return
	}

	// response from `AccountFundingFlow`: AccountFundingFlowResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **currency** | **string** | Asset type, only support USDT  as of now | 
 **recordId** | **int64** | Return the recordId and subsequent data, the latest data is returned by default, e.g 100000 | 
 **startTime** | **int64** | Start Time, e.g 1593511200000 | 
 **endTime** | **int64** | End Time, e.g 1593512200000 | 
 **limit** | **int64** | Number of result sets returned Default:100 Max:1000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**AccountFundingFlowResponse**](AccountFundingFlowResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## OptionMarginAccountInformation

> OptionMarginAccountInformationResponse OptionMarginAccountInformation(ctx).RecvWindow(recvWindow).Execute()

Option Margin Account Information (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.OptionMarginAccountInformation(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.OptionMarginAccountInformation``: %v\n", err)
		return
	}

	// response from `OptionMarginAccountInformation`: OptionMarginAccountInformationResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** |  | 

### Return type

[**OptionMarginAccountInformationResponse**](OptionMarginAccountInformationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

