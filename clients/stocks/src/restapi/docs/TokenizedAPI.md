# \TokenizedAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**BstocksCorporateActionRecords**](TokenizedAPI.md#BstocksCorporateActionRecords) | **Get** /sapi/v1/bstocks/corp-action | Bstocks Corporate Action Records (USER_DATA)
[**CorporateActionRecords**](TokenizedAPI.md#CorporateActionRecords) | **Get** /sapi/v1/equity/corp-action | Corporate Action Records (USER_DATA)
[**TokenizedConvertHistory**](TokenizedAPI.md#TokenizedConvertHistory) | **Get** /sapi/v1/equity/tokenized/history | Tokenized Convert History (USER_DATA)
[**TokenizedConvertStatus**](TokenizedAPI.md#TokenizedConvertStatus) | **Get** /sapi/v1/equity/tokenized/convert-status | Tokenized Convert Status (USER_DATA)
[**TokenizedMint**](TokenizedAPI.md#TokenizedMint) | **Post** /sapi/v1/equity/tokenized/mint | Tokenized Mint (TRADE)
[**TokenizedRedeem**](TokenizedAPI.md#TokenizedRedeem) | **Post** /sapi/v1/equity/tokenized/redeem | Tokenized Redeem (TRADE)


## BstocksCorporateActionRecords

> BstocksCorporateActionRecordsResponse BstocksCorporateActionRecords(ctx).Limit(limit).LastTime(lastTime).LastId(lastId).RecvWindow(recvWindow).Execute()

Bstocks Corporate Action Records (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/stocks"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	limit := int32(20) // int32 | Number of records to return. Default `20`, range `1` to `100`.
	lastTime := int64(1795593600000) // int64 | The `nextTime` returned by the previous page, in milliseconds. Must be sent together with `lastId`. (optional)
	lastId := int64(123456) // int64 | The `nextId` returned by the previous page. Must be sent together with `lastTime`. (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than `60000`. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStocksClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TokenizedAPI.BstocksCorporateActionRecords(context.Background()).Limit(limit).LastTime(lastTime).LastId(lastId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TokenizedAPI.BstocksCorporateActionRecords``: %v\n", err)
		return
	}

	// response from `BstocksCorporateActionRecords`: BstocksCorporateActionRecordsResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of records to return. Default &#x60;20&#x60;, range &#x60;1&#x60; to &#x60;100&#x60;. | 
 **lastTime** | **int64** | The &#x60;nextTime&#x60; returned by the previous page, in milliseconds. Must be sent together with &#x60;lastId&#x60;. | 
 **lastId** | **int64** | The &#x60;nextId&#x60; returned by the previous page. Must be sent together with &#x60;lastTime&#x60;. | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60;. | 

### Return type

[**BstocksCorporateActionRecordsResponse**](BstocksCorporateActionRecordsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CorporateActionRecords

> CorporateActionRecordsResponse CorporateActionRecords(ctx).Limit(limit).LastTime(lastTime).LastId(lastId).RecvWindow(recvWindow).Execute()

Corporate Action Records (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/stocks"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	limit := int32(20) // int32 | Number of records to return. Default `20`, range `1` to `100`.
	lastTime := int64(1795593600000) // int64 | The `nextTime` returned by the previous page, in milliseconds. Must be sent together with `lastId`. (optional)
	lastId := int64(123456) // int64 | The `nextId` returned by the previous page. Must be sent together with `lastTime`. (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than `60000`. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStocksClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TokenizedAPI.CorporateActionRecords(context.Background()).Limit(limit).LastTime(lastTime).LastId(lastId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TokenizedAPI.CorporateActionRecords``: %v\n", err)
		return
	}

	// response from `CorporateActionRecords`: CorporateActionRecordsResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of records to return. Default &#x60;20&#x60;, range &#x60;1&#x60; to &#x60;100&#x60;. | 
 **lastTime** | **int64** | The &#x60;nextTime&#x60; returned by the previous page, in milliseconds. Must be sent together with &#x60;lastId&#x60;. | 
 **lastId** | **int64** | The &#x60;nextId&#x60; returned by the previous page. Must be sent together with &#x60;lastTime&#x60;. | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60;. | 

### Return type

[**CorporateActionRecordsResponse**](CorporateActionRecordsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## TokenizedConvertHistory

> TokenizedConvertHistoryResponse TokenizedConvertHistory(ctx).StartTime(startTime).EndTime(endTime).LastId(lastId).Size(size).RecvWindow(recvWindow).Execute()

Tokenized Convert History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/stocks"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	startTime := int64(1735800000000) // int64 | Start time (ms epoch). (optional)
	endTime := int64(1735900000000) // int64 | End time (ms epoch). (optional)
	lastId := int64(10019) // int64 | Last record id from the previous page. Omit (or leave unset) to fetch the first page. (optional)
	size := int32(20) // int32 | Page size. Default `20`, max `100`. (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than `60000`. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStocksClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TokenizedAPI.TokenizedConvertHistory(context.Background()).StartTime(startTime).EndTime(endTime).LastId(lastId).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TokenizedAPI.TokenizedConvertHistory``: %v\n", err)
		return
	}

	// response from `TokenizedConvertHistory`: TokenizedConvertHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** | Start time (ms epoch). | 
 **endTime** | **int64** | End time (ms epoch). | 
 **lastId** | **int64** | Last record id from the previous page. Omit (or leave unset) to fetch the first page. | 
 **size** | **int32** | Page size. Default &#x60;20&#x60;, max &#x60;100&#x60;. | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60;. | 

### Return type

[**TokenizedConvertHistoryResponse**](TokenizedConvertHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## TokenizedConvertStatus

> TokenizedConvertStatusResponse TokenizedConvertStatus(ctx).IssuerRequestId(issuerRequestId).ConvertType(convertType).RecvWindow(recvWindow).Execute()

Tokenized Convert Status (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/stocks"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	issuerRequestId := "mint-20260505-8f3b9e1a2d3c4b5a" // string | Convert request id returned by `/tokenized/mint` or `/redeem`.
	convertType := models.TokenizedConvertStatusConvertTypeParameterMint // TokenizedConvertStatusConvertTypeParameter | `MINT` or `REDEEM`.
	recvWindow := int64(5000) // int64 | The value cannot be greater than `60000`. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStocksClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TokenizedAPI.TokenizedConvertStatus(context.Background()).IssuerRequestId(issuerRequestId).ConvertType(convertType).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TokenizedAPI.TokenizedConvertStatus``: %v\n", err)
		return
	}

	// response from `TokenizedConvertStatus`: TokenizedConvertStatusResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **issuerRequestId** | **string** | Convert request id returned by &#x60;/tokenized/mint&#x60; or &#x60;/redeem&#x60;. | 
 **convertType** | [**TokenizedConvertStatusConvertTypeParameter**](TokenizedConvertStatusConvertTypeParameter.md) | &#x60;MINT&#x60; or &#x60;REDEEM&#x60;. | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60;. | 

### Return type

[**TokenizedConvertStatusResponse**](TokenizedConvertStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## TokenizedMint

> TokenizedMintResponse TokenizedMint(ctx).UnderlyingAsset(underlyingAsset).UnderlyingAssetAmount(underlyingAssetAmount).ClientOrderId(clientOrderId).TokenizedAssetSettleWalletType(tokenizedAssetSettleWalletType).RecvWindow(recvWindow).Execute()

Tokenized Mint (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/stocks"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	underlyingAsset := "AAPL" // string | Underlying US-equity ticker, e.g. `AAPL`, `TSLA`. Resolved against the active-symbol list; unknown tickers return `-26004`. The target tokenized asset is looked up from this field via `/market/tokenized-assets`.
	underlyingAssetAmount := "1" // string | Quantity of the underlying asset to mint from. Must be > 0.
	clientOrderId := "mint-client-id-32chars-0000000001" // string | Client order id for idempotency. Format `^[a-zA-Z0-9-_]{32,36}$`. Auto-generated when omitted. (optional)
	tokenizedAssetSettleWalletType := models.PlaceEquityOrderTokenizedAssetSettleWalletTypeParameterMain // PlaceEquityOrderTokenizedAssetSettleWalletTypeParameter | Wallet that receives the minted tokenized asset. `MAIN` = Spot wallet, `CARD` = Funding wallet. Default `CARD`. If `MAIN` is sent while Spot wallet settlement is not enabled, the request is rejected as an illegal parameter. (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than `60000`. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStocksClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TokenizedAPI.TokenizedMint(context.Background()).UnderlyingAsset(underlyingAsset).UnderlyingAssetAmount(underlyingAssetAmount).ClientOrderId(clientOrderId).TokenizedAssetSettleWalletType(tokenizedAssetSettleWalletType).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TokenizedAPI.TokenizedMint``: %v\n", err)
		return
	}

	// response from `TokenizedMint`: TokenizedMintResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **underlyingAsset** | **string** | Underlying US-equity ticker, e.g. &#x60;AAPL&#x60;, &#x60;TSLA&#x60;. Resolved against the active-symbol list; unknown tickers return &#x60;-26004&#x60;. The target tokenized asset is looked up from this field via &#x60;/market/tokenized-assets&#x60;. | 
 **underlyingAssetAmount** | **string** | Quantity of the underlying asset to mint from. Must be &gt; 0. | 
 **clientOrderId** | **string** | Client order id for idempotency. Format &#x60;^[a-zA-Z0-9-_]{32,36}$&#x60;. Auto-generated when omitted. | 
 **tokenizedAssetSettleWalletType** | [**PlaceEquityOrderTokenizedAssetSettleWalletTypeParameter**](PlaceEquityOrderTokenizedAssetSettleWalletTypeParameter.md) | Wallet that receives the minted tokenized asset. &#x60;MAIN&#x60; &#x3D; Spot wallet, &#x60;CARD&#x60; &#x3D; Funding wallet. Default &#x60;CARD&#x60;. If &#x60;MAIN&#x60; is sent while Spot wallet settlement is not enabled, the request is rejected as an illegal parameter. | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60;. | 

### Return type

[**TokenizedMintResponse**](TokenizedMintResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## TokenizedRedeem

> TokenizedRedeemResponse TokenizedRedeem(ctx).TokenizedAsset(tokenizedAsset).TokenizedAssetAmount(tokenizedAssetAmount).ClientOrderId(clientOrderId).TokenizedAssetSettleWalletType(tokenizedAssetSettleWalletType).RecvWindow(recvWindow).Execute()

Tokenized Redeem (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/stocks"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	tokenizedAsset := "AAPLB" // string | Tokenized asset to redeem, e.g. `AAPLB`. Not a US-equity ticker — this is the on-chain tokenized asset identifier. Unknown asset returns `-1102` (the message currently says the parameter was empty/malformed, but it was in fact sent — it is simply unknown). The target underlying ticker is looked up from this field via `/market/tokenized-assets`.
	tokenizedAssetAmount := "1" // string | Quantity of the tokenized asset to redeem. Must be > 0.
	clientOrderId := "redeem-client-id-32chars-000000001" // string | Client order id for idempotency. Format `^[a-zA-Z0-9-_]{32,36}$`. Auto-generated when omitted. (optional)
	tokenizedAssetSettleWalletType := models.PlaceEquityOrderTokenizedAssetSettleWalletTypeParameterMain // PlaceEquityOrderTokenizedAssetSettleWalletTypeParameter | Wallet the tokenized asset is deducted from on redeem. `MAIN` = Spot wallet, `CARD` = Funding wallet. Default `CARD`. If `MAIN` is sent while Spot wallet settlement is not enabled, the request is rejected as an illegal parameter. (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than `60000`. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStocksClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TokenizedAPI.TokenizedRedeem(context.Background()).TokenizedAsset(tokenizedAsset).TokenizedAssetAmount(tokenizedAssetAmount).ClientOrderId(clientOrderId).TokenizedAssetSettleWalletType(tokenizedAssetSettleWalletType).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TokenizedAPI.TokenizedRedeem``: %v\n", err)
		return
	}

	// response from `TokenizedRedeem`: TokenizedRedeemResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **tokenizedAsset** | **string** | Tokenized asset to redeem, e.g. &#x60;AAPLB&#x60;. Not a US-equity ticker — this is the on-chain tokenized asset identifier. Unknown asset returns &#x60;-1102&#x60; (the message currently says the parameter was empty/malformed, but it was in fact sent — it is simply unknown). The target underlying ticker is looked up from this field via &#x60;/market/tokenized-assets&#x60;. | 
 **tokenizedAssetAmount** | **string** | Quantity of the tokenized asset to redeem. Must be &gt; 0. | 
 **clientOrderId** | **string** | Client order id for idempotency. Format &#x60;^[a-zA-Z0-9-_]{32,36}$&#x60;. Auto-generated when omitted. | 
 **tokenizedAssetSettleWalletType** | [**PlaceEquityOrderTokenizedAssetSettleWalletTypeParameter**](PlaceEquityOrderTokenizedAssetSettleWalletTypeParameter.md) | Wallet the tokenized asset is deducted from on redeem. &#x60;MAIN&#x60; &#x3D; Spot wallet, &#x60;CARD&#x60; &#x3D; Funding wallet. Default &#x60;CARD&#x60;. If &#x60;MAIN&#x60; is sent while Spot wallet settlement is not enabled, the request is rejected as an illegal parameter. | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60;. | 

### Return type

[**TokenizedRedeemResponse**](TokenizedRedeemResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

