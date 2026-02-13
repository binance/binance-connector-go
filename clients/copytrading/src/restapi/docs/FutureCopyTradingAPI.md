# \FutureCopyTradingAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**GetFuturesLeadTraderStatus**](FutureCopyTradingAPI.md#GetFuturesLeadTraderStatus) | **Get** /sapi/v1/copyTrading/futures/userStatus | Get Futures Lead Trader Status(TRADE)
[**GetFuturesLeadTradingSymbolWhitelist**](FutureCopyTradingAPI.md#GetFuturesLeadTradingSymbolWhitelist) | **Get** /sapi/v1/copyTrading/futures/leadSymbol | Get Futures Lead Trading Symbol Whitelist(USER_DATA)


## GetFuturesLeadTraderStatus

> GetFuturesLeadTraderStatusResponse GetFuturesLeadTraderStatus(ctx).RecvWindow(recvWindow).Execute()

Get Futures Lead Trader Status(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/copytrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceCopyTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FutureCopyTradingAPI.GetFuturesLeadTraderStatus(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FutureCopyTradingAPI.GetFuturesLeadTraderStatus``: %v\n", err)
		return
	}

	// response from `GetFuturesLeadTraderStatus`: GetFuturesLeadTraderStatusResponse
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

[**GetFuturesLeadTraderStatusResponse**](GetFuturesLeadTraderStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetFuturesLeadTradingSymbolWhitelist

> GetFuturesLeadTradingSymbolWhitelistResponse GetFuturesLeadTradingSymbolWhitelist(ctx).RecvWindow(recvWindow).Execute()

Get Futures Lead Trading Symbol Whitelist(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/copytrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceCopyTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.FutureCopyTradingAPI.GetFuturesLeadTradingSymbolWhitelist(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `FutureCopyTradingAPI.GetFuturesLeadTradingSymbolWhitelist``: %v\n", err)
		return
	}

	// response from `GetFuturesLeadTradingSymbolWhitelist`: GetFuturesLeadTradingSymbolWhitelistResponse
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

[**GetFuturesLeadTradingSymbolWhitelistResponse**](GetFuturesLeadTradingSymbolWhitelistResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

