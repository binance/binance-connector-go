# \AssetAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AssetDetail**](AssetAPI.md#AssetDetail) | **Get** /sapi/v1/asset/assetDetail | Asset Detail (USER_DATA)
[**AssetDividendRecord**](AssetAPI.md#AssetDividendRecord) | **Get** /sapi/v1/asset/assetDividend | Asset Dividend Record (USER_DATA)
[**DustConvert**](AssetAPI.md#DustConvert) | **Post** /sapi/v1/asset/dust-convert/convert | Dust Convert (USER_DATA)
[**DustConvertibleAssets**](AssetAPI.md#DustConvertibleAssets) | **Post** /sapi/v1/asset/dust-convert/query-convertible-assets | Dust Convertible Assets (USER_DATA)
[**DustTransfer**](AssetAPI.md#DustTransfer) | **Post** /sapi/v1/asset/dust | Dust Transfer (USER_DATA)
[**Dustlog**](AssetAPI.md#Dustlog) | **Get** /sapi/v1/asset/dribblet | DustLog(USER_DATA)
[**FundingWallet**](AssetAPI.md#FundingWallet) | **Post** /sapi/v1/asset/get-funding-asset | Funding Wallet (USER_DATA)
[**GetAssetsThatCanBeConvertedIntoBnb**](AssetAPI.md#GetAssetsThatCanBeConvertedIntoBnb) | **Post** /sapi/v1/asset/dust-btc | Get Assets That Can Be Converted Into BNB (USER_DATA)
[**GetCloudMiningPaymentAndRefundHistory**](AssetAPI.md#GetCloudMiningPaymentAndRefundHistory) | **Get** /sapi/v1/asset/ledger-transfer/cloud-mining/queryByPage | Get Cloud-Mining payment and refund history (USER_DATA)
[**GetOpenSymbolList**](AssetAPI.md#GetOpenSymbolList) | **Get** /sapi/v1/spot/open-symbol-list | Get Open Symbol List (MARKET_DATA)
[**QueryUserDelegationHistory**](AssetAPI.md#QueryUserDelegationHistory) | **Get** /sapi/v1/asset/custody/transfer-history | Query User Delegation History(For Master Account)(USER_DATA)
[**QueryUserUniversalTransferHistory**](AssetAPI.md#QueryUserUniversalTransferHistory) | **Get** /sapi/v1/asset/transfer | Query User Universal Transfer History(USER_DATA)
[**QueryUserWalletBalance**](AssetAPI.md#QueryUserWalletBalance) | **Get** /sapi/v1/asset/wallet/balance | Query User Wallet Balance (USER_DATA)
[**ToggleBnbBurnOnSpotTradeAndMarginInterest**](AssetAPI.md#ToggleBnbBurnOnSpotTradeAndMarginInterest) | **Post** /sapi/v1/bnbBurn | Toggle BNB Burn On Spot Trade And Margin Interest (USER_DATA)
[**TradeFee**](AssetAPI.md#TradeFee) | **Get** /sapi/v1/asset/tradeFee | Trade Fee (USER_DATA)
[**UserAsset**](AssetAPI.md#UserAsset) | **Post** /sapi/v3/asset/getUserAsset | User Asset (USER_DATA)
[**UserUniversalTransfer**](AssetAPI.md#UserUniversalTransfer) | **Post** /sapi/v1/asset/transfer | User Universal Transfer (USER_DATA)


## AssetDetail

> AssetDetailResponse AssetDetail(ctx).Asset(asset).RecvWindow(recvWindow).Execute()

Asset Detail (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	asset := "asset_example" // string | If asset is blank, then query all positive assets user have. (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetAPI.AssetDetail(context.Background()).Asset(asset).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetAPI.AssetDetail``: %v\n", err)
		return
	}

	// response from `AssetDetail`: AssetDetailResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** | If asset is blank, then query all positive assets user have. | 
 **recvWindow** | **int64** |  | 

### Return type

[**AssetDetailResponse**](AssetDetailResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## AssetDividendRecord

> AssetDividendRecordResponse AssetDividendRecord(ctx).Asset(asset).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Asset Dividend Record (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	asset := "asset_example" // string | If asset is blank, then query all positive assets user have. (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	limit := int64(7) // int64 | min 7, max 30, default 7 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetAPI.AssetDividendRecord(context.Background()).Asset(asset).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetAPI.AssetDividendRecord``: %v\n", err)
		return
	}

	// response from `AssetDividendRecord`: AssetDividendRecordResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** | If asset is blank, then query all positive assets user have. | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int64** | min 7, max 30, default 7 | 
 **recvWindow** | **int64** |  | 

### Return type

[**AssetDividendRecordResponse**](AssetDividendRecordResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## DustConvert

> DustConvertResponse DustConvert(ctx).Asset(asset).ClientId(clientId).TargetAsset(targetAsset).ThirdPartyClientId(thirdPartyClientId).DustQuotaAssetToTargetAssetPrice(dustQuotaAssetToTargetAssetPrice).Execute()

Dust Convert (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	asset := "asset_example" // string | 
	clientId := "1" // string | A unique id for the request (optional)
	targetAsset := "targetAsset_example" // string |  (optional)
	thirdPartyClientId := "1" // string |  (optional)
	dustQuotaAssetToTargetAssetPrice := float32(1.0) // float32 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetAPI.DustConvert(context.Background()).Asset(asset).ClientId(clientId).TargetAsset(targetAsset).ThirdPartyClientId(thirdPartyClientId).DustQuotaAssetToTargetAssetPrice(dustQuotaAssetToTargetAssetPrice).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetAPI.DustConvert``: %v\n", err)
		return
	}

	// response from `DustConvert`: DustConvertResponse
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
 **clientId** | **string** | A unique id for the request | 
 **targetAsset** | **string** |  | 
 **thirdPartyClientId** | **string** |  | 
 **dustQuotaAssetToTargetAssetPrice** | **float32** |  | 

### Return type

[**DustConvertResponse**](DustConvertResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## DustConvertibleAssets

> DustConvertibleAssetsResponse DustConvertibleAssets(ctx).TargetAsset(targetAsset).DustQuotaAssetToTargetAssetPrice(dustQuotaAssetToTargetAssetPrice).Execute()

Dust Convertible Assets (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	targetAsset := "targetAsset_example" // string | 
	dustQuotaAssetToTargetAssetPrice := float32(1.0) // float32 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetAPI.DustConvertibleAssets(context.Background()).TargetAsset(targetAsset).DustQuotaAssetToTargetAssetPrice(dustQuotaAssetToTargetAssetPrice).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetAPI.DustConvertibleAssets``: %v\n", err)
		return
	}

	// response from `DustConvertibleAssets`: DustConvertibleAssetsResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **targetAsset** | **string** |  | 
 **dustQuotaAssetToTargetAssetPrice** | **float32** |  | 

### Return type

[**DustConvertibleAssetsResponse**](DustConvertibleAssetsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## DustTransfer

> DustTransferResponse DustTransfer(ctx).Asset(asset).AccountType(accountType).RecvWindow(recvWindow).Execute()

Dust Transfer (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	asset := "asset_example" // string | 
	accountType := "SPOT" // string | `SPOT` or `MARGIN`,default `SPOT` (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetAPI.DustTransfer(context.Background()).Asset(asset).AccountType(accountType).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetAPI.DustTransfer``: %v\n", err)
		return
	}

	// response from `DustTransfer`: DustTransferResponse
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
 **accountType** | **string** | &#x60;SPOT&#x60; or &#x60;MARGIN&#x60;,default &#x60;SPOT&#x60; | 
 **recvWindow** | **int64** |  | 

### Return type

[**DustTransferResponse**](DustTransferResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## Dustlog

> DustlogResponse Dustlog(ctx).AccountType(accountType).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

DustLog(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	accountType := "SPOT" // string | `SPOT` or `MARGIN`,default `SPOT` (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetAPI.Dustlog(context.Background()).AccountType(accountType).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetAPI.Dustlog``: %v\n", err)
		return
	}

	// response from `Dustlog`: DustlogResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **accountType** | **string** | &#x60;SPOT&#x60; or &#x60;MARGIN&#x60;,default &#x60;SPOT&#x60; | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**DustlogResponse**](DustlogResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## FundingWallet

> FundingWalletResponse FundingWallet(ctx).Asset(asset).NeedBtcValuation(needBtcValuation).RecvWindow(recvWindow).Execute()

Funding Wallet (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	asset := "asset_example" // string | If asset is blank, then query all positive assets user have. (optional)
	needBtcValuation := "needBtcValuation_example" // string | true or false (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetAPI.FundingWallet(context.Background()).Asset(asset).NeedBtcValuation(needBtcValuation).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetAPI.FundingWallet``: %v\n", err)
		return
	}

	// response from `FundingWallet`: FundingWalletResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** | If asset is blank, then query all positive assets user have. | 
 **needBtcValuation** | **string** | true or false | 
 **recvWindow** | **int64** |  | 

### Return type

[**FundingWalletResponse**](FundingWalletResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetAssetsThatCanBeConvertedIntoBnb

> GetAssetsThatCanBeConvertedIntoBnbResponse GetAssetsThatCanBeConvertedIntoBnb(ctx).AccountType(accountType).RecvWindow(recvWindow).Execute()

Get Assets That Can Be Converted Into BNB (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	accountType := "SPOT" // string | `SPOT` or `MARGIN`,default `SPOT` (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetAPI.GetAssetsThatCanBeConvertedIntoBnb(context.Background()).AccountType(accountType).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetAPI.GetAssetsThatCanBeConvertedIntoBnb``: %v\n", err)
		return
	}

	// response from `GetAssetsThatCanBeConvertedIntoBnb`: GetAssetsThatCanBeConvertedIntoBnbResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **accountType** | **string** | &#x60;SPOT&#x60; or &#x60;MARGIN&#x60;,default &#x60;SPOT&#x60; | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetAssetsThatCanBeConvertedIntoBnbResponse**](GetAssetsThatCanBeConvertedIntoBnbResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetCloudMiningPaymentAndRefundHistory

> GetCloudMiningPaymentAndRefundHistoryResponse GetCloudMiningPaymentAndRefundHistory(ctx).StartTime(startTime).EndTime(endTime).TranId(tranId).ClientTranId(clientTranId).Asset(asset).Current(current).Size(size).Execute()

Get Cloud-Mining payment and refund history (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	startTime := int64(1623319461670) // int64 | 
	endTime := int64(1641782889000) // int64 | 
	tranId := int64(1) // int64 | The transaction id (optional)
	clientTranId := "1" // string | The unique flag (optional)
	asset := "asset_example" // string | If asset is blank, then query all positive assets user have. (optional)
	current := int64(1) // int64 | current page, default 1, the min value is 1 (optional)
	size := int64(10) // int64 | page size, default 10, the max value is 100 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetAPI.GetCloudMiningPaymentAndRefundHistory(context.Background()).StartTime(startTime).EndTime(endTime).TranId(tranId).ClientTranId(clientTranId).Asset(asset).Current(current).Size(size).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetAPI.GetCloudMiningPaymentAndRefundHistory``: %v\n", err)
		return
	}

	// response from `GetCloudMiningPaymentAndRefundHistory`: GetCloudMiningPaymentAndRefundHistoryResponse
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
 **tranId** | **int64** | The transaction id | 
 **clientTranId** | **string** | The unique flag | 
 **asset** | **string** | If asset is blank, then query all positive assets user have. | 
 **current** | **int64** | current page, default 1, the min value is 1 | 
 **size** | **int64** | page size, default 10, the max value is 100 | 

### Return type

[**GetCloudMiningPaymentAndRefundHistoryResponse**](GetCloudMiningPaymentAndRefundHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetOpenSymbolList

> GetOpenSymbolListResponse GetOpenSymbolList(ctx).Execute()

Get Open Symbol List (MARKET_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetAPI.GetOpenSymbolList(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetAPI.GetOpenSymbolList``: %v\n", err)
		return
	}

	// response from `GetOpenSymbolList`: GetOpenSymbolListResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

This endpoint does not need any parameter.

### Return type

[**GetOpenSymbolListResponse**](GetOpenSymbolListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryUserDelegationHistory

> QueryUserDelegationHistoryResponse QueryUserDelegationHistory(ctx).Email(email).StartTime(startTime).EndTime(endTime).Type(type_).Asset(asset).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Query User Delegation History(For Master Account)(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	email := "email_example" // string | 
	startTime := int64(1623319461670) // int64 | 
	endTime := int64(1641782889000) // int64 | 
	type_ := "type__example" // string | Delegate/Undelegate (optional)
	asset := "asset_example" // string | If asset is blank, then query all positive assets user have. (optional)
	current := int64(1) // int64 | current page, default 1, the min value is 1 (optional)
	size := int64(10) // int64 | page size, default 10, the max value is 100 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetAPI.QueryUserDelegationHistory(context.Background()).Email(email).StartTime(startTime).EndTime(endTime).Type(type_).Asset(asset).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetAPI.QueryUserDelegationHistory``: %v\n", err)
		return
	}

	// response from `QueryUserDelegationHistory`: QueryUserDelegationHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **type_** | **string** | Delegate/Undelegate | 
 **asset** | **string** | If asset is blank, then query all positive assets user have. | 
 **current** | **int64** | current page, default 1, the min value is 1 | 
 **size** | **int64** | page size, default 10, the max value is 100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryUserDelegationHistoryResponse**](QueryUserDelegationHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryUserUniversalTransferHistory

> QueryUserUniversalTransferHistoryResponse QueryUserUniversalTransferHistory(ctx).Type(type_).StartTime(startTime).EndTime(endTime).Current(current).Size(size).FromSymbol(fromSymbol).ToSymbol(toSymbol).RecvWindow(recvWindow).Execute()

Query User Universal Transfer History(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	type_ := "type__example" // string | 
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	current := int64(1) // int64 | current page, default 1, the min value is 1 (optional)
	size := int64(10) // int64 | page size, default 10, the max value is 100 (optional)
	fromSymbol := "fromSymbol_example" // string |  (optional)
	toSymbol := "toSymbol_example" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetAPI.QueryUserUniversalTransferHistory(context.Background()).Type(type_).StartTime(startTime).EndTime(endTime).Current(current).Size(size).FromSymbol(fromSymbol).ToSymbol(toSymbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetAPI.QueryUserUniversalTransferHistory``: %v\n", err)
		return
	}

	// response from `QueryUserUniversalTransferHistory`: QueryUserUniversalTransferHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | current page, default 1, the min value is 1 | 
 **size** | **int64** | page size, default 10, the max value is 100 | 
 **fromSymbol** | **string** |  | 
 **toSymbol** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryUserUniversalTransferHistoryResponse**](QueryUserUniversalTransferHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryUserWalletBalance

> QueryUserWalletBalanceResponse QueryUserWalletBalance(ctx).QuoteAsset(quoteAsset).RecvWindow(recvWindow).Execute()

Query User Wallet Balance (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	quoteAsset := "BTC" // string | `USDT`, `ETH`, `USDC`, `BNB`, etc. default `BTC` (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetAPI.QueryUserWalletBalance(context.Background()).QuoteAsset(quoteAsset).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetAPI.QueryUserWalletBalance``: %v\n", err)
		return
	}

	// response from `QueryUserWalletBalance`: QueryUserWalletBalanceResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **quoteAsset** | **string** | &#x60;USDT&#x60;, &#x60;ETH&#x60;, &#x60;USDC&#x60;, &#x60;BNB&#x60;, etc. default &#x60;BTC&#x60; | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryUserWalletBalanceResponse**](QueryUserWalletBalanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## ToggleBnbBurnOnSpotTradeAndMarginInterest

> ToggleBnbBurnOnSpotTradeAndMarginInterestResponse ToggleBnbBurnOnSpotTradeAndMarginInterest(ctx).SpotBNBBurn(spotBNBBurn).InterestBNBBurn(interestBNBBurn).RecvWindow(recvWindow).Execute()

Toggle BNB Burn On Spot Trade And Margin Interest (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	spotBNBBurn := "spotBNBBurn_example" // string | \"true\" or \"false\"; Determines whether to use BNB to pay for trading fees on SPOT (optional)
	interestBNBBurn := "interestBNBBurn_example" // string | \"true\" or \"false\"; Determines whether to use BNB to pay for margin loan's interest (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetAPI.ToggleBnbBurnOnSpotTradeAndMarginInterest(context.Background()).SpotBNBBurn(spotBNBBurn).InterestBNBBurn(interestBNBBurn).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetAPI.ToggleBnbBurnOnSpotTradeAndMarginInterest``: %v\n", err)
		return
	}

	// response from `ToggleBnbBurnOnSpotTradeAndMarginInterest`: ToggleBnbBurnOnSpotTradeAndMarginInterestResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **spotBNBBurn** | **string** | \&quot;true\&quot; or \&quot;false\&quot;; Determines whether to use BNB to pay for trading fees on SPOT | 
 **interestBNBBurn** | **string** | \&quot;true\&quot; or \&quot;false\&quot;; Determines whether to use BNB to pay for margin loan&#39;s interest | 
 **recvWindow** | **int64** |  | 

### Return type

[**ToggleBnbBurnOnSpotTradeAndMarginInterestResponse**](ToggleBnbBurnOnSpotTradeAndMarginInterestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## TradeFee

> TradeFeeResponse TradeFee(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Trade Fee (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetAPI.TradeFee(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetAPI.TradeFee``: %v\n", err)
		return
	}

	// response from `TradeFee`: TradeFeeResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**TradeFeeResponse**](TradeFeeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## UserAsset

> UserAssetResponse UserAsset(ctx).Asset(asset).NeedBtcValuation(needBtcValuation).RecvWindow(recvWindow).Execute()

User Asset (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	asset := "asset_example" // string | If asset is blank, then query all positive assets user have. (optional)
	needBtcValuation := true // bool | Whether need btc valuation or not. (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetAPI.UserAsset(context.Background()).Asset(asset).NeedBtcValuation(needBtcValuation).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetAPI.UserAsset``: %v\n", err)
		return
	}

	// response from `UserAsset`: UserAssetResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** | If asset is blank, then query all positive assets user have. | 
 **needBtcValuation** | **bool** | Whether need btc valuation or not. | 
 **recvWindow** | **int64** |  | 

### Return type

[**UserAssetResponse**](UserAssetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## UserUniversalTransfer

> UserUniversalTransferResponse UserUniversalTransfer(ctx).Type(type_).Asset(asset).Amount(amount).FromSymbol(fromSymbol).ToSymbol(toSymbol).RecvWindow(recvWindow).Execute()

User Universal Transfer (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	type_ := "type__example" // string | 
	asset := "asset_example" // string | 
	amount := float32(1.0) // float32 | 
	fromSymbol := "fromSymbol_example" // string |  (optional)
	toSymbol := "toSymbol_example" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetAPI.UserUniversalTransfer(context.Background()).Type(type_).Asset(asset).Amount(amount).FromSymbol(fromSymbol).ToSymbol(toSymbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetAPI.UserUniversalTransfer``: %v\n", err)
		return
	}

	// response from `UserUniversalTransfer`: UserUniversalTransferResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** |  | 
 **asset** | **string** |  | 
 **amount** | **float32** |  | 
 **fromSymbol** | **string** |  | 
 **toSymbol** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UserUniversalTransferResponse**](UserUniversalTransferResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

