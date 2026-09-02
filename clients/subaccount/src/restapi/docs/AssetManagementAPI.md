# \AssetManagementAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**FuturesTransferForSubAccount**](AssetManagementAPI.md#FuturesTransferForSubAccount) | **Post** /sapi/v1/sub-account/futures/transfer | Futures Transfer for Sub-account (For Master Account) (USER_DATA)
[**GetDetailOnSubAccountsFuturesAccount**](AssetManagementAPI.md#GetDetailOnSubAccountsFuturesAccount) | **Get** /sapi/v1/sub-account/futures/account | Get Detail on Sub-account&#39;s Futures Account (For Master Account) (USER_DATA)
[**GetDetailOnSubAccountsFuturesAccountV2**](AssetManagementAPI.md#GetDetailOnSubAccountsFuturesAccountV2) | **Get** /sapi/v2/sub-account/futures/account | Get Detail on Sub-account&#39;s Futures Account V2 (For Master Account) (USER_DATA)
[**GetDetailOnSubAccountsMarginAccount**](AssetManagementAPI.md#GetDetailOnSubAccountsMarginAccount) | **Get** /sapi/v1/sub-account/margin/account | Get Detail on Sub-account&#39;s Margin Account (For Master Account) (USER_DATA)
[**GetMovePositionHistoryForSubAccount**](AssetManagementAPI.md#GetMovePositionHistoryForSubAccount) | **Get** /sapi/v1/sub-account/futures/move-position | Get Move Position History for Sub-account (For Master Account) (USER_DATA)
[**GetSubAccountDepositAddress**](AssetManagementAPI.md#GetSubAccountDepositAddress) | **Get** /sapi/v1/capital/deposit/subAddress | Get Sub-account Deposit Address (For Master Account) (USER_DATA)
[**GetSubAccountDepositHistory**](AssetManagementAPI.md#GetSubAccountDepositHistory) | **Get** /sapi/v1/capital/deposit/subHisrec | Get Sub-account Deposit History (For Master Account) (USER_DATA)
[**GetSummaryOfSubAccountsFuturesAccount**](AssetManagementAPI.md#GetSummaryOfSubAccountsFuturesAccount) | **Get** /sapi/v1/sub-account/futures/accountSummary | Get Summary of Sub-account&#39;s Futures Account (For Master Account) (USER_DATA)
[**GetSummaryOfSubAccountsFuturesAccountV2**](AssetManagementAPI.md#GetSummaryOfSubAccountsFuturesAccountV2) | **Get** /sapi/v2/sub-account/futures/accountSummary | Get Summary of Sub-account&#39;s Futures Account V2 (For Master Account) (USER_DATA)
[**GetSummaryOfSubAccountsMarginAccount**](AssetManagementAPI.md#GetSummaryOfSubAccountsMarginAccount) | **Get** /sapi/v1/sub-account/margin/accountSummary | Get Summary of Sub-account&#39;s Margin Account (For Master Account) (USER_DATA)
[**MarginTransferForSubAccount**](AssetManagementAPI.md#MarginTransferForSubAccount) | **Post** /sapi/v1/sub-account/margin/transfer | Margin Transfer for Sub-account (For Master Account) (USER_DATA)
[**MovePositionForSubAccount**](AssetManagementAPI.md#MovePositionForSubAccount) | **Post** /sapi/v1/sub-account/futures/move-position | Move Position for Sub-account (For Master Account) (USER_DATA)
[**QuerySubAccountAssets**](AssetManagementAPI.md#QuerySubAccountAssets) | **Get** /sapi/v3/sub-account/assets | Query Sub-account Assets (For Master Account) (USER_DATA)
[**QuerySubAccountAssetsAssetManagement**](AssetManagementAPI.md#QuerySubAccountAssetsAssetManagement) | **Get** /sapi/v4/sub-account/assets | Query Sub-account Assets V4 (For Master Account) (USER_DATA)
[**QuerySubAccountFuturesAssetTransferHistory**](AssetManagementAPI.md#QuerySubAccountFuturesAssetTransferHistory) | **Get** /sapi/v1/sub-account/futures/internalTransfer | Query Sub-account Futures Asset Transfer History (For Master Account) (USER_DATA)
[**QuerySubAccountSpotAssetTransferHistory**](AssetManagementAPI.md#QuerySubAccountSpotAssetTransferHistory) | **Get** /sapi/v1/sub-account/sub/transfer/history | Query Sub-account Spot Asset Transfer History (For Master Account) (USER_DATA)
[**QuerySubAccountSpotAssetsSummary**](AssetManagementAPI.md#QuerySubAccountSpotAssetsSummary) | **Get** /sapi/v1/sub-account/spotSummary | Query Sub-account Spot Assets Summary (For Master Account) (USER_DATA)
[**QueryUniversalTransferHistory**](AssetManagementAPI.md#QueryUniversalTransferHistory) | **Get** /sapi/v1/sub-account/universalTransfer | Query Universal Transfer History (For Master Account) (USER_DATA)
[**SubAccountFuturesAssetTransfer**](AssetManagementAPI.md#SubAccountFuturesAssetTransfer) | **Post** /sapi/v1/sub-account/futures/internalTransfer | Sub-account Futures Asset Transfer (For Master Account) (USER_DATA)
[**SubAccountTransferHistory**](AssetManagementAPI.md#SubAccountTransferHistory) | **Get** /sapi/v1/sub-account/transfer/subUserHistory | Sub-account Transfer History (For Sub-account) (USER_DATA)
[**TransferToMaster**](AssetManagementAPI.md#TransferToMaster) | **Post** /sapi/v1/sub-account/transfer/subToMaster | Transfer to Master (For Sub-account) (USER_DATA)
[**TransferToSubAccountOfSameMaster**](AssetManagementAPI.md#TransferToSubAccountOfSameMaster) | **Post** /sapi/v1/sub-account/transfer/subToSub | Transfer to Sub-account of Same Master (For Sub-account) (USER_DATA)
[**UniversalTransfer**](AssetManagementAPI.md#UniversalTransfer) | **Post** /sapi/v1/sub-account/universalTransfer | Universal Transfer (For Master Account) (USER_DATA)


## FuturesTransferForSubAccount

> FuturesTransferForSubAccountResponse FuturesTransferForSubAccount(ctx).Email(email).Asset(asset).Amount(amount).Type(type_).RecvWindow(recvWindow).Execute()

Futures Transfer for Sub-account (For Master Account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	email := "123@test.com" // string | 
	asset := "USDT" // string | The asset being transferred
	amount := float64(1.0) // float64 | The amount to be transferred
	type_ := int64(1) // int64 | 1: transfer from subaccount's spot account to its USDT-margined futures account 2: transfer from subaccount's USDT-margined futures account to its spot account 3: transfer from subaccount's spot account to its COIN-margined futures account 4:transfer from subaccount's COIN-margined futures account to its spot account
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetManagementAPI.FuturesTransferForSubAccount(context.Background()).Email(email).Asset(asset).Amount(amount).Type(type_).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetManagementAPI.FuturesTransferForSubAccount``: %v\n", err)
		return
	}

	// response from `FuturesTransferForSubAccount`: FuturesTransferForSubAccountResponse
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
 **asset** | **string** | The asset being transferred | 
 **amount** | **float64** | The amount to be transferred | 
 **type_** | **int64** | 1: transfer from subaccount&#39;s spot account to its USDT-margined futures account 2: transfer from subaccount&#39;s USDT-margined futures account to its spot account 3: transfer from subaccount&#39;s spot account to its COIN-margined futures account 4:transfer from subaccount&#39;s COIN-margined futures account to its spot account | 
 **recvWindow** | **int64** |  | 

### Return type

[**FuturesTransferForSubAccountResponse**](FuturesTransferForSubAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetDetailOnSubAccountsFuturesAccount

> GetDetailOnSubAccountsFuturesAccountResponse GetDetailOnSubAccountsFuturesAccount(ctx).Email(email).RecvWindow(recvWindow).Execute()

Get Detail on Sub-account's Futures Account (For Master Account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	email := "123@test.com" // string | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetManagementAPI.GetDetailOnSubAccountsFuturesAccount(context.Background()).Email(email).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetManagementAPI.GetDetailOnSubAccountsFuturesAccount``: %v\n", err)
		return
	}

	// response from `GetDetailOnSubAccountsFuturesAccount`: GetDetailOnSubAccountsFuturesAccountResponse
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
 **recvWindow** | **int64** |  | 

### Return type

[**GetDetailOnSubAccountsFuturesAccountResponse**](GetDetailOnSubAccountsFuturesAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetDetailOnSubAccountsFuturesAccountV2

> GetDetailOnSubAccountsFuturesAccountV2Response GetDetailOnSubAccountsFuturesAccountV2(ctx).Email(email).FuturesType(futuresType).RecvWindow(recvWindow).Execute()

Get Detail on Sub-account's Futures Account V2 (For Master Account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	email := "123@test.com" // string | 
	futuresType := int64(1) // int64 | 1:USDT-margined Futures，2: Coin-margined Futures
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetManagementAPI.GetDetailOnSubAccountsFuturesAccountV2(context.Background()).Email(email).FuturesType(futuresType).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetManagementAPI.GetDetailOnSubAccountsFuturesAccountV2``: %v\n", err)
		return
	}

	// response from `GetDetailOnSubAccountsFuturesAccountV2`: GetDetailOnSubAccountsFuturesAccountV2Response
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
 **futuresType** | **int64** | 1:USDT-margined Futures，2: Coin-margined Futures | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetDetailOnSubAccountsFuturesAccountV2Response**](GetDetailOnSubAccountsFuturesAccountV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetDetailOnSubAccountsMarginAccount

> GetDetailOnSubAccountsMarginAccountResponse GetDetailOnSubAccountsMarginAccount(ctx).Email(email).RecvWindow(recvWindow).Execute()

Get Detail on Sub-account's Margin Account (For Master Account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	email := "123@test.com" // string | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetManagementAPI.GetDetailOnSubAccountsMarginAccount(context.Background()).Email(email).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetManagementAPI.GetDetailOnSubAccountsMarginAccount``: %v\n", err)
		return
	}

	// response from `GetDetailOnSubAccountsMarginAccount`: GetDetailOnSubAccountsMarginAccountResponse
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
 **recvWindow** | **int64** |  | 

### Return type

[**GetDetailOnSubAccountsMarginAccountResponse**](GetDetailOnSubAccountsMarginAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetMovePositionHistoryForSubAccount

> GetMovePositionHistoryForSubAccountResponse GetMovePositionHistoryForSubAccount(ctx).Symbol(symbol).Page(page).Rows(rows).ProductType(productType).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

Get Move Position History for Sub-account (For Master Account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSDT" // string | 
	page := int64(1) // int64 | 
	rows := int64(1) // int64 | Max 100.
	productType := models.GetMovePositionHistoryForSubAccountProductTypeParameterUm // GetMovePositionHistoryForSubAccountProductTypeParameter | Default UM. (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetManagementAPI.GetMovePositionHistoryForSubAccount(context.Background()).Symbol(symbol).Page(page).Rows(rows).ProductType(productType).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetManagementAPI.GetMovePositionHistoryForSubAccount``: %v\n", err)
		return
	}

	// response from `GetMovePositionHistoryForSubAccount`: GetMovePositionHistoryForSubAccountResponse
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
 **page** | **int64** |  | 
 **rows** | **int64** | Max 100. | 
 **productType** | [**GetMovePositionHistoryForSubAccountProductTypeParameter**](GetMovePositionHistoryForSubAccountProductTypeParameter.md) | Default UM. | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetMovePositionHistoryForSubAccountResponse**](GetMovePositionHistoryForSubAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetSubAccountDepositAddress

> GetSubAccountDepositAddressResponse GetSubAccountDepositAddress(ctx).Email(email).Coin(coin).Network(network).Amount(amount).RecvWindow(recvWindow).Execute()

Get Sub-account Deposit Address (For Master Account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	email := "123@test.com" // string | 
	coin := "BTC" // string | 
	network := "network_example" // string | networks can be found in `GET /sapi/v1/capital/deposit/address` (optional)
	amount := float64(1.0) // float64 |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetManagementAPI.GetSubAccountDepositAddress(context.Background()).Email(email).Coin(coin).Network(network).Amount(amount).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetManagementAPI.GetSubAccountDepositAddress``: %v\n", err)
		return
	}

	// response from `GetSubAccountDepositAddress`: GetSubAccountDepositAddressResponse
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
 **coin** | **string** |  | 
 **network** | **string** | networks can be found in &#x60;GET /sapi/v1/capital/deposit/address&#x60; | 
 **amount** | **float64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSubAccountDepositAddressResponse**](GetSubAccountDepositAddressResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetSubAccountDepositHistory

> GetSubAccountDepositHistoryResponse GetSubAccountDepositHistory(ctx).Email(email).IncludeSource(includeSource).Coin(coin).Status(status).StartTime(startTime).EndTime(endTime).Limit(limit).Offset(offset).RecvWindow(recvWindow).TxId(txId).Execute()

Get Sub-account Deposit History (For Master Account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	email := "123@test.com" // string | 
	includeSource := false // bool | Default `false`, return `sourceAddress` field when set to `true` (optional)
	coin := "BTC" // string |  (optional)
	status := int64(0) // int64 | Deposit status: 0=pending, 6=credited but cannot withdraw, 7=wrong deposit, 8=waiting user confirmation, 1=success. (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	limit := int64(1) // int64 |  (optional)
	offset := int64(0) // int64 |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)
	txId := "1" // string |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetManagementAPI.GetSubAccountDepositHistory(context.Background()).Email(email).IncludeSource(includeSource).Coin(coin).Status(status).StartTime(startTime).EndTime(endTime).Limit(limit).Offset(offset).RecvWindow(recvWindow).TxId(txId).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetManagementAPI.GetSubAccountDepositHistory``: %v\n", err)
		return
	}

	// response from `GetSubAccountDepositHistory`: GetSubAccountDepositHistoryResponse
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
 **includeSource** | **bool** | Default &#x60;false&#x60;, return &#x60;sourceAddress&#x60; field when set to &#x60;true&#x60; | 
 **coin** | **string** |  | 
 **status** | **int64** | Deposit status: 0&#x3D;pending, 6&#x3D;credited but cannot withdraw, 7&#x3D;wrong deposit, 8&#x3D;waiting user confirmation, 1&#x3D;success. | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int64** |  | 
 **offset** | **int64** |  | 
 **recvWindow** | **int64** |  | 
 **txId** | **string** |  | 

### Return type

[**GetSubAccountDepositHistoryResponse**](GetSubAccountDepositHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetSummaryOfSubAccountsFuturesAccount

> GetSummaryOfSubAccountsFuturesAccountResponse GetSummaryOfSubAccountsFuturesAccount(ctx).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Get Summary of Sub-account's Futures Account (For Master Account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	page := int64(1) // int64 | 
	limit := int64(1) // int64 | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetManagementAPI.GetSummaryOfSubAccountsFuturesAccount(context.Background()).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetManagementAPI.GetSummaryOfSubAccountsFuturesAccount``: %v\n", err)
		return
	}

	// response from `GetSummaryOfSubAccountsFuturesAccount`: GetSummaryOfSubAccountsFuturesAccountResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | 
 **limit** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSummaryOfSubAccountsFuturesAccountResponse**](GetSummaryOfSubAccountsFuturesAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetSummaryOfSubAccountsFuturesAccountV2

> GetSummaryOfSubAccountsFuturesAccountV2Response GetSummaryOfSubAccountsFuturesAccountV2(ctx).FuturesType(futuresType).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Get Summary of Sub-account's Futures Account V2 (For Master Account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	futuresType := int64(1) // int64 | 1:USDT-margined Futures，2: Coin-margined Futures
	page := int64(1) // int64 |  (optional)
	limit := int64(10) // int64 |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetManagementAPI.GetSummaryOfSubAccountsFuturesAccountV2(context.Background()).FuturesType(futuresType).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetManagementAPI.GetSummaryOfSubAccountsFuturesAccountV2``: %v\n", err)
		return
	}

	// response from `GetSummaryOfSubAccountsFuturesAccountV2`: GetSummaryOfSubAccountsFuturesAccountV2Response
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **futuresType** | **int64** | 1:USDT-margined Futures，2: Coin-margined Futures | 
 **page** | **int64** |  | 
 **limit** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSummaryOfSubAccountsFuturesAccountV2Response**](GetSummaryOfSubAccountsFuturesAccountV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetSummaryOfSubAccountsMarginAccount

> GetSummaryOfSubAccountsMarginAccountResponse GetSummaryOfSubAccountsMarginAccount(ctx).RecvWindow(recvWindow).Execute()

Get Summary of Sub-account's Margin Account (For Master Account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetManagementAPI.GetSummaryOfSubAccountsMarginAccount(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetManagementAPI.GetSummaryOfSubAccountsMarginAccount``: %v\n", err)
		return
	}

	// response from `GetSummaryOfSubAccountsMarginAccount`: GetSummaryOfSubAccountsMarginAccountResponse
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

[**GetSummaryOfSubAccountsMarginAccountResponse**](GetSummaryOfSubAccountsMarginAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MarginTransferForSubAccount

> MarginTransferForSubAccountResponse MarginTransferForSubAccount(ctx).Email(email).Asset(asset).Amount(amount).Type(type_).RecvWindow(recvWindow).Execute()

Margin Transfer for Sub-account (For Master Account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	email := "123@test.com" // string | 
	asset := "BTC" // string | The asset being transferred
	amount := float64(1.0) // float64 | The amount to be transferred
	type_ := int64(1) // int64 | 1: transfer from subaccount's spot account to margin account 2: transfer from subaccount's margin account to its spot account
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetManagementAPI.MarginTransferForSubAccount(context.Background()).Email(email).Asset(asset).Amount(amount).Type(type_).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetManagementAPI.MarginTransferForSubAccount``: %v\n", err)
		return
	}

	// response from `MarginTransferForSubAccount`: MarginTransferForSubAccountResponse
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
 **asset** | **string** | The asset being transferred | 
 **amount** | **float64** | The amount to be transferred | 
 **type_** | **int64** | 1: transfer from subaccount&#39;s spot account to margin account 2: transfer from subaccount&#39;s margin account to its spot account | 
 **recvWindow** | **int64** |  | 

### Return type

[**MarginTransferForSubAccountResponse**](MarginTransferForSubAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MovePositionForSubAccount

> MovePositionForSubAccountResponse MovePositionForSubAccount(ctx).FromUserEmail(fromUserEmail).ToUserEmail(toUserEmail).ProductType(productType).OrderArgs(orderArgs).RecvWindow(recvWindow).Execute()

Move Position for Sub-account (For Master Account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	fromUserEmail := "testFrom@google.com" // string | 
	toUserEmail := "testTo@google.com" // string | 
	productType := models.GetMovePositionHistoryForSubAccountProductTypeParameterUm // GetMovePositionHistoryForSubAccountProductTypeParameter | A single request cannot mix UM and OPTION positions.
	orderArgs := []models.MovePositionForSubAccountOrderArgsParameterInner{*models.NewMovePositionForSubAccountOrderArgsParameterInner()} // []MovePositionForSubAccountOrderArgsParameterInner | Max 10 positions supported. When input request parameter,orderArgs.symbol should be STRING, orderArgs.quantity should be BIGDECIMAL, and orderArgs.positionSide should be STRING, positionSide support BOTH,LONG and SHORT. Each entry should be like orderArgs[0].symbol=BTCUSDT,orderArgs[0].quantity=0.001,orderArgs[0].positionSide=BOTH. Example of the request parameter array: orderArgs[0].symbol=BTCUSDT orderArgs[0].quantity=0.001 orderArgs[0].positionSide=BOTH orderArgs[1].symbol=ETHUSDT orderArgs[1].quantity=0.01 orderArgs[1].positionSide=BOTH
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetManagementAPI.MovePositionForSubAccount(context.Background()).FromUserEmail(fromUserEmail).ToUserEmail(toUserEmail).ProductType(productType).OrderArgs(orderArgs).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetManagementAPI.MovePositionForSubAccount``: %v\n", err)
		return
	}

	// response from `MovePositionForSubAccount`: MovePositionForSubAccountResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **fromUserEmail** | **string** |  | 
 **toUserEmail** | **string** |  | 
 **productType** | [**GetMovePositionHistoryForSubAccountProductTypeParameter**](GetMovePositionHistoryForSubAccountProductTypeParameter.md) | A single request cannot mix UM and OPTION positions. | 
 **orderArgs** | [**[]MovePositionForSubAccountOrderArgsParameterInner**](MovePositionForSubAccountOrderArgsParameterInner.md) | Max 10 positions supported. When input request parameter,orderArgs.symbol should be STRING, orderArgs.quantity should be BIGDECIMAL, and orderArgs.positionSide should be STRING, positionSide support BOTH,LONG and SHORT. Each entry should be like orderArgs[0].symbol&#x3D;BTCUSDT,orderArgs[0].quantity&#x3D;0.001,orderArgs[0].positionSide&#x3D;BOTH. Example of the request parameter array: orderArgs[0].symbol&#x3D;BTCUSDT orderArgs[0].quantity&#x3D;0.001 orderArgs[0].positionSide&#x3D;BOTH orderArgs[1].symbol&#x3D;ETHUSDT orderArgs[1].quantity&#x3D;0.01 orderArgs[1].positionSide&#x3D;BOTH | 
 **recvWindow** | **int64** |  | 

### Return type

[**MovePositionForSubAccountResponse**](MovePositionForSubAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QuerySubAccountAssets

> QuerySubAccountAssetsResponse QuerySubAccountAssets(ctx).Email(email).RecvWindow(recvWindow).Execute()

Query Sub-account Assets (For Master Account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	email := "123@test.com" // string | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetManagementAPI.QuerySubAccountAssets(context.Background()).Email(email).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetManagementAPI.QuerySubAccountAssets``: %v\n", err)
		return
	}

	// response from `QuerySubAccountAssets`: QuerySubAccountAssetsResponse
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
 **recvWindow** | **int64** |  | 

### Return type

[**QuerySubAccountAssetsResponse**](QuerySubAccountAssetsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QuerySubAccountAssetsAssetManagement

> QuerySubAccountAssetsAssetManagementResponse QuerySubAccountAssetsAssetManagement(ctx).Email(email).RecvWindow(recvWindow).Execute()

Query Sub-account Assets V4 (For Master Account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	email := "123@test.com" // string | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetManagementAPI.QuerySubAccountAssetsAssetManagement(context.Background()).Email(email).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetManagementAPI.QuerySubAccountAssetsAssetManagement``: %v\n", err)
		return
	}

	// response from `QuerySubAccountAssetsAssetManagement`: QuerySubAccountAssetsAssetManagementResponse
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
 **recvWindow** | **int64** |  | 

### Return type

[**QuerySubAccountAssetsAssetManagementResponse**](QuerySubAccountAssetsAssetManagementResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QuerySubAccountFuturesAssetTransferHistory

> QuerySubAccountFuturesAssetTransferHistoryResponse QuerySubAccountFuturesAssetTransferHistory(ctx).Email(email).FuturesType(futuresType).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Query Sub-account Futures Asset Transfer History (For Master Account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	email := "123@test.com" // string | 
	futuresType := int64(1) // int64 | 1:USDT-margined Futures，2: Coin-margined Futures
	startTime := int64(1623319461670) // int64 | Cannot be earlier than 1 month ago (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	page := int64(1) // int64 |  (optional)
	limit := int64(10) // int64 |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetManagementAPI.QuerySubAccountFuturesAssetTransferHistory(context.Background()).Email(email).FuturesType(futuresType).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetManagementAPI.QuerySubAccountFuturesAssetTransferHistory``: %v\n", err)
		return
	}

	// response from `QuerySubAccountFuturesAssetTransferHistory`: QuerySubAccountFuturesAssetTransferHistoryResponse
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
 **futuresType** | **int64** | 1:USDT-margined Futures，2: Coin-margined Futures | 
 **startTime** | **int64** | Cannot be earlier than 1 month ago | 
 **endTime** | **int64** |  | 
 **page** | **int64** |  | 
 **limit** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**QuerySubAccountFuturesAssetTransferHistoryResponse**](QuerySubAccountFuturesAssetTransferHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QuerySubAccountSpotAssetTransferHistory

> QuerySubAccountSpotAssetTransferHistoryResponse QuerySubAccountSpotAssetTransferHistory(ctx).FromEmail(fromEmail).ToEmail(toEmail).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Query Sub-account Spot Asset Transfer History (For Master Account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	fromEmail := "aaa@test.com" // string |  (optional)
	toEmail := "bbb@test.com" // string |  (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	page := int64(1) // int64 |  (optional)
	limit := int64(10) // int64 |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetManagementAPI.QuerySubAccountSpotAssetTransferHistory(context.Background()).FromEmail(fromEmail).ToEmail(toEmail).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetManagementAPI.QuerySubAccountSpotAssetTransferHistory``: %v\n", err)
		return
	}

	// response from `QuerySubAccountSpotAssetTransferHistory`: QuerySubAccountSpotAssetTransferHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **fromEmail** | **string** |  | 
 **toEmail** | **string** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **page** | **int64** |  | 
 **limit** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**QuerySubAccountSpotAssetTransferHistoryResponse**](QuerySubAccountSpotAssetTransferHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QuerySubAccountSpotAssetsSummary

> QuerySubAccountSpotAssetsSummaryResponse QuerySubAccountSpotAssetsSummary(ctx).Email(email).Page(page).Size(size).RecvWindow(recvWindow).Execute()

Query Sub-account Spot Assets Summary (For Master Account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	email := "123@test.com" // string | Managed sub-account email (optional)
	page := int64(1) // int64 |  (optional)
	size := int64(10) // int64 |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetManagementAPI.QuerySubAccountSpotAssetsSummary(context.Background()).Email(email).Page(page).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetManagementAPI.QuerySubAccountSpotAssetsSummary``: %v\n", err)
		return
	}

	// response from `QuerySubAccountSpotAssetsSummary`: QuerySubAccountSpotAssetsSummaryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Managed sub-account email | 
 **page** | **int64** |  | 
 **size** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**QuerySubAccountSpotAssetsSummaryResponse**](QuerySubAccountSpotAssetsSummaryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryUniversalTransferHistory

> QueryUniversalTransferHistoryResponse QueryUniversalTransferHistory(ctx).FromEmail(fromEmail).ToEmail(toEmail).ClientTranId(clientTranId).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Query Universal Transfer History (For Master Account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	fromEmail := "abctest@gmail.com" // string |  (optional)
	toEmail := "deftest@gmail.com" // string |  (optional)
	clientTranId := "1" // string |  (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	page := int64(1) // int64 |  (optional)
	limit := int64(10) // int64 |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetManagementAPI.QueryUniversalTransferHistory(context.Background()).FromEmail(fromEmail).ToEmail(toEmail).ClientTranId(clientTranId).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetManagementAPI.QueryUniversalTransferHistory``: %v\n", err)
		return
	}

	// response from `QueryUniversalTransferHistory`: QueryUniversalTransferHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **fromEmail** | **string** |  | 
 **toEmail** | **string** |  | 
 **clientTranId** | **string** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **page** | **int64** |  | 
 **limit** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryUniversalTransferHistoryResponse**](QueryUniversalTransferHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SubAccountFuturesAssetTransfer

> SubAccountFuturesAssetTransferResponse SubAccountFuturesAssetTransfer(ctx).FromEmail(fromEmail).ToEmail(toEmail).FuturesType(futuresType).Asset(asset).Amount(amount).RecvWindow(recvWindow).Execute()

Sub-account Futures Asset Transfer (For Master Account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	fromEmail := "abc@test.com" // string | Sender email
	toEmail := "def@test.com" // string | Recipient email
	futuresType := int64(1) // int64 | 1:USDT-margined Futures，2: Coin-margined Futures
	asset := "BTC" // string | 
	amount := float64(1.0) // float64 | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetManagementAPI.SubAccountFuturesAssetTransfer(context.Background()).FromEmail(fromEmail).ToEmail(toEmail).FuturesType(futuresType).Asset(asset).Amount(amount).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetManagementAPI.SubAccountFuturesAssetTransfer``: %v\n", err)
		return
	}

	// response from `SubAccountFuturesAssetTransfer`: SubAccountFuturesAssetTransferResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **fromEmail** | **string** | Sender email | 
 **toEmail** | **string** | Recipient email | 
 **futuresType** | **int64** | 1:USDT-margined Futures，2: Coin-margined Futures | 
 **asset** | **string** |  | 
 **amount** | **float64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubAccountFuturesAssetTransferResponse**](SubAccountFuturesAssetTransferResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SubAccountTransferHistory

> SubAccountTransferHistoryResponse SubAccountTransferHistory(ctx).Asset(asset).Type(type_).StartTime(startTime).EndTime(endTime).Limit(limit).ReturnFailHistory(returnFailHistory).RecvWindow(recvWindow).Execute()

Sub-account Transfer History (For Sub-account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	asset := "BTC" // string | If not sent, result of all assets will be returned (optional)
	type_ := int64(1) // int64 | 1: transfer in, 2: transfer out (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	limit := int64(10) // int64 |  (optional)
	returnFailHistory := false // bool | Default `False`, return PROCESS and SUCCESS status history; If `True`,return PROCESS and SUCCESS and FAILURE status history (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetManagementAPI.SubAccountTransferHistory(context.Background()).Asset(asset).Type(type_).StartTime(startTime).EndTime(endTime).Limit(limit).ReturnFailHistory(returnFailHistory).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetManagementAPI.SubAccountTransferHistory``: %v\n", err)
		return
	}

	// response from `SubAccountTransferHistory`: SubAccountTransferHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** | If not sent, result of all assets will be returned | 
 **type_** | **int64** | 1: transfer in, 2: transfer out | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int64** |  | 
 **returnFailHistory** | **bool** | Default &#x60;False&#x60;, return PROCESS and SUCCESS status history; If &#x60;True&#x60;,return PROCESS and SUCCESS and FAILURE status history | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubAccountTransferHistoryResponse**](SubAccountTransferHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## TransferToMaster

> TransferToMasterResponse TransferToMaster(ctx).Asset(asset).Amount(amount).RecvWindow(recvWindow).Execute()

Transfer to Master (For Sub-account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	asset := "BTC" // string | 
	amount := float64(1.0) // float64 | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetManagementAPI.TransferToMaster(context.Background()).Asset(asset).Amount(amount).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetManagementAPI.TransferToMaster``: %v\n", err)
		return
	}

	// response from `TransferToMaster`: TransferToMasterResponse
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
 **amount** | **float64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**TransferToMasterResponse**](TransferToMasterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## TransferToSubAccountOfSameMaster

> TransferToSubAccountOfSameMasterResponse TransferToSubAccountOfSameMaster(ctx).ToEmail(toEmail).Asset(asset).Amount(amount).RecvWindow(recvWindow).Execute()

Transfer to Sub-account of Same Master (For Sub-account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	toEmail := "abc@test.com" // string | 
	asset := "BTC" // string | 
	amount := float64(1.0) // float64 | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetManagementAPI.TransferToSubAccountOfSameMaster(context.Background()).ToEmail(toEmail).Asset(asset).Amount(amount).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetManagementAPI.TransferToSubAccountOfSameMaster``: %v\n", err)
		return
	}

	// response from `TransferToSubAccountOfSameMaster`: TransferToSubAccountOfSameMasterResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **toEmail** | **string** |  | 
 **asset** | **string** |  | 
 **amount** | **float64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**TransferToSubAccountOfSameMasterResponse**](TransferToSubAccountOfSameMasterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## UniversalTransfer

> UniversalTransferResponse UniversalTransfer(ctx).FromAccountType(fromAccountType).ToAccountType(toAccountType).Asset(asset).Amount(amount).FromEmail(fromEmail).ToEmail(toEmail).ClientTranId(clientTranId).Symbol(symbol).RecvWindow(recvWindow).Execute()

Universal Transfer (For Master Account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	fromAccountType := models.UniversalTransferFromAccountTypeParameterSpot // UniversalTransferFromAccountTypeParameter | 
	toAccountType := models.UniversalTransferFromAccountTypeParameterSpot // UniversalTransferFromAccountTypeParameter | 
	asset := "BTC" // string | 
	amount := float64(1.0) // float64 | 
	fromEmail := "abc@test.com" // string |  (optional)
	toEmail := "def@test.com" // string |  (optional)
	clientTranId := "1" // string | Must be unique (optional)
	symbol := "BTCUSDT" // string | Only supported under ISOLATED_MARGIN type (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AssetManagementAPI.UniversalTransfer(context.Background()).FromAccountType(fromAccountType).ToAccountType(toAccountType).Asset(asset).Amount(amount).FromEmail(fromEmail).ToEmail(toEmail).ClientTranId(clientTranId).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AssetManagementAPI.UniversalTransfer``: %v\n", err)
		return
	}

	// response from `UniversalTransfer`: UniversalTransferResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **fromAccountType** | [**UniversalTransferFromAccountTypeParameter**](UniversalTransferFromAccountTypeParameter.md) |  | 
 **toAccountType** | [**UniversalTransferFromAccountTypeParameter**](UniversalTransferFromAccountTypeParameter.md) |  | 
 **asset** | **string** |  | 
 **amount** | **float64** |  | 
 **fromEmail** | **string** |  | 
 **toEmail** | **string** |  | 
 **clientTranId** | **string** | Must be unique | 
 **symbol** | **string** | Only supported under ISOLATED_MARGIN type | 
 **recvWindow** | **int64** |  | 

### Return type

[**UniversalTransferResponse**](UniversalTransferResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

