# \ManagedSubAccountAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**DepositAssetsIntoTheManagedSubAccount**](ManagedSubAccountAPI.md#DepositAssetsIntoTheManagedSubAccount) | **Post** /sapi/v1/managed-subaccount/deposit | Deposit Assets Into The Managed Sub-account (For Investor Master Account) (USER_DATA)
[**GetManagedSubAccountDepositAddress**](ManagedSubAccountAPI.md#GetManagedSubAccountDepositAddress) | **Get** /sapi/v1/managed-subaccount/deposit/address | Get Managed Sub-account Deposit Address (For Investor Master Account) (USER_DATA)
[**QueryManagedSubAccountAssetDetails**](ManagedSubAccountAPI.md#QueryManagedSubAccountAssetDetails) | **Get** /sapi/v1/managed-subaccount/asset | Query Managed Sub-account Asset Details (For Investor Master Account) (USER_DATA)
[**QueryManagedSubAccountFuturesAssetDetails**](ManagedSubAccountAPI.md#QueryManagedSubAccountFuturesAssetDetails) | **Get** /sapi/v1/managed-subaccount/fetch-future-asset | Query Managed Sub-account Futures Asset Details (For Investor Master Account) (USER_DATA)
[**QueryManagedSubAccountList**](ManagedSubAccountAPI.md#QueryManagedSubAccountList) | **Get** /sapi/v1/managed-subaccount/info | Query Managed Sub-account List (For Investor) (USER_DATA)
[**QueryManagedSubAccountMarginAssetDetails**](ManagedSubAccountAPI.md#QueryManagedSubAccountMarginAssetDetails) | **Get** /sapi/v1/managed-subaccount/marginAsset | Query Managed Sub-account Margin Asset Details (For Investor Master Account) (USER_DATA)
[**QueryManagedSubAccountSnapshot**](ManagedSubAccountAPI.md#QueryManagedSubAccountSnapshot) | **Get** /sapi/v1/managed-subaccount/accountSnapshot | Query Managed Sub-account Snapshot (For Investor Master Account) (USER_DATA)
[**QueryManagedSubAccountTransferLogMasterAccountInvestor**](ManagedSubAccountAPI.md#QueryManagedSubAccountTransferLogMasterAccountInvestor) | **Get** /sapi/v1/managed-subaccount/queryTransLogForInvestor | Query Managed Sub Account Transfer Log For Investor Master Account (USER_DATA)
[**QueryManagedSubAccountTransferLogMasterAccountTrading**](ManagedSubAccountAPI.md#QueryManagedSubAccountTransferLogMasterAccountTrading) | **Get** /sapi/v1/managed-subaccount/queryTransLogForTradeParent | Query Managed Sub Account Transfer Log For Trading Team Master Account (USER_DATA)
[**QueryManagedSubAccountTransferLogSubAccountTrading**](ManagedSubAccountAPI.md#QueryManagedSubAccountTransferLogSubAccountTrading) | **Get** /sapi/v1/managed-subaccount/query-trans-log | Query Managed Sub Account Transfer Log (For Trading Team Sub Account) (USER_DATA)
[**WithdrawlAssetsFromTheManagedSubAccount**](ManagedSubAccountAPI.md#WithdrawlAssetsFromTheManagedSubAccount) | **Post** /sapi/v1/managed-subaccount/withdraw | Withdrawl Assets From The Managed Sub-account (For Investor Master Account) (USER_DATA)


## DepositAssetsIntoTheManagedSubAccount

> DepositAssetsIntoTheManagedSubAccountResponse DepositAssetsIntoTheManagedSubAccount(ctx).ToEmail(toEmail).Asset(asset).Amount(amount).RecvWindow(recvWindow).Execute()

Deposit Assets Into The Managed Sub-account (For Investor Master Account) (USER_DATA)


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
	amount := float32(1.0) // float32 | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.ManagedSubAccountAPI.DepositAssetsIntoTheManagedSubAccount(context.Background()).ToEmail(toEmail).Asset(asset).Amount(amount).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `ManagedSubAccountAPI.DepositAssetsIntoTheManagedSubAccount``: %v\n", err)
		return
	}

	// response from `DepositAssetsIntoTheManagedSubAccount`: DepositAssetsIntoTheManagedSubAccountResponse
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
 **amount** | **float32** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**DepositAssetsIntoTheManagedSubAccountResponse**](DepositAssetsIntoTheManagedSubAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetManagedSubAccountDepositAddress

> GetManagedSubAccountDepositAddressResponse GetManagedSubAccountDepositAddress(ctx).Email(email).Coin(coin).Network(network).Amount(amount).RecvWindow(recvWindow).Execute()

Get Managed Sub-account Deposit Address (For Investor Master Account) (USER_DATA)


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
	email := "abc@test.com" // string | 
	coin := "USDT" // string | 
	network := "LIGHTNING" // string | networks can be found in `GET /sapi/v1/capital/deposit/address` (optional)
	amount := float32(1.0) // float32 |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.ManagedSubAccountAPI.GetManagedSubAccountDepositAddress(context.Background()).Email(email).Coin(coin).Network(network).Amount(amount).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `ManagedSubAccountAPI.GetManagedSubAccountDepositAddress``: %v\n", err)
		return
	}

	// response from `GetManagedSubAccountDepositAddress`: GetManagedSubAccountDepositAddressResponse
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
 **amount** | **float32** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetManagedSubAccountDepositAddressResponse**](GetManagedSubAccountDepositAddressResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryManagedSubAccountAssetDetails

> QueryManagedSubAccountAssetDetailsResponse QueryManagedSubAccountAssetDetails(ctx).Email(email).RecvWindow(recvWindow).Execute()

Query Managed Sub-account Asset Details (For Investor Master Account) (USER_DATA)


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
	email := "abc@test.com" // string | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountAssetDetails(context.Background()).Email(email).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `ManagedSubAccountAPI.QueryManagedSubAccountAssetDetails``: %v\n", err)
		return
	}

	// response from `QueryManagedSubAccountAssetDetails`: QueryManagedSubAccountAssetDetailsResponse
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

[**QueryManagedSubAccountAssetDetailsResponse**](QueryManagedSubAccountAssetDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryManagedSubAccountFuturesAssetDetails

> QueryManagedSubAccountFuturesAssetDetailsResponse QueryManagedSubAccountFuturesAssetDetails(ctx).Email(email).AccountType(accountType).Execute()

Query Managed Sub-account Futures Asset Details (For Investor Master Account) (USER_DATA)


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
	email := "abc@test.com" // string | 
	accountType := "MARGIN" // string | No input or input \"USDT_FUTURE\" to get UM Futures account details. Input \"COIN_FUTURE\" to get CM Futures account details. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountFuturesAssetDetails(context.Background()).Email(email).AccountType(accountType).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `ManagedSubAccountAPI.QueryManagedSubAccountFuturesAssetDetails``: %v\n", err)
		return
	}

	// response from `QueryManagedSubAccountFuturesAssetDetails`: QueryManagedSubAccountFuturesAssetDetailsResponse
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
 **accountType** | **string** | No input or input \&quot;USDT_FUTURE\&quot; to get UM Futures account details. Input \&quot;COIN_FUTURE\&quot; to get CM Futures account details. | 

### Return type

[**QueryManagedSubAccountFuturesAssetDetailsResponse**](QueryManagedSubAccountFuturesAssetDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryManagedSubAccountList

> QueryManagedSubAccountListResponse QueryManagedSubAccountList(ctx).Email(email).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Query Managed Sub-account List (For Investor) (USER_DATA)


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
	email := "abc@test.com" // string |  (optional)
	page := int64(1) // int64 |  (optional)
	limit := int64(10) // int64 |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountList(context.Background()).Email(email).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `ManagedSubAccountAPI.QueryManagedSubAccountList``: %v\n", err)
		return
	}

	// response from `QueryManagedSubAccountList`: QueryManagedSubAccountListResponse
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
 **page** | **int64** |  | 
 **limit** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryManagedSubAccountListResponse**](QueryManagedSubAccountListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryManagedSubAccountMarginAssetDetails

> QueryManagedSubAccountMarginAssetDetailsResponse QueryManagedSubAccountMarginAssetDetails(ctx).Email(email).AccountType(accountType).Execute()

Query Managed Sub-account Margin Asset Details (For Investor Master Account) (USER_DATA)


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
	email := "abc@test.com" // string | 
	accountType := "MARGIN" // string | No input or input \"MARGIN\" to get Cross Margin account details. Input \"ISOLATED_MARGIN\" to get Isolated Margin account details. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountMarginAssetDetails(context.Background()).Email(email).AccountType(accountType).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `ManagedSubAccountAPI.QueryManagedSubAccountMarginAssetDetails``: %v\n", err)
		return
	}

	// response from `QueryManagedSubAccountMarginAssetDetails`: QueryManagedSubAccountMarginAssetDetailsResponse
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
 **accountType** | **string** | No input or input \&quot;MARGIN\&quot; to get Cross Margin account details. Input \&quot;ISOLATED_MARGIN\&quot; to get Isolated Margin account details. | 

### Return type

[**QueryManagedSubAccountMarginAssetDetailsResponse**](QueryManagedSubAccountMarginAssetDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryManagedSubAccountSnapshot

> QueryManagedSubAccountSnapshotResponse QueryManagedSubAccountSnapshot(ctx).Email(email).Type(type_).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query Managed Sub-account Snapshot (For Investor Master Account) (USER_DATA)


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
	email := "abc@test.com" // string | 
	type_ := models.QueryManagedSubAccountSnapshotTypeParameterSpot // QueryManagedSubAccountSnapshotTypeParameter | 
	startTime := int64(1623319461670) // int64 | Query time range must be within 30 days and only supports data within the last month. (optional)
	endTime := int64(1641782889000) // int64 | If both startTime and endTime are omitted, records from the last 7 days are returned by default. (optional)
	limit := int64(10) // int64 |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountSnapshot(context.Background()).Email(email).Type(type_).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `ManagedSubAccountAPI.QueryManagedSubAccountSnapshot``: %v\n", err)
		return
	}

	// response from `QueryManagedSubAccountSnapshot`: QueryManagedSubAccountSnapshotResponse
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
 **type_** | [**QueryManagedSubAccountSnapshotTypeParameter**](QueryManagedSubAccountSnapshotTypeParameter.md) |  | 
 **startTime** | **int64** | Query time range must be within 30 days and only supports data within the last month. | 
 **endTime** | **int64** | If both startTime and endTime are omitted, records from the last 7 days are returned by default. | 
 **limit** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryManagedSubAccountSnapshotResponse**](QueryManagedSubAccountSnapshotResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryManagedSubAccountTransferLogMasterAccountInvestor

> QueryManagedSubAccountTransferLogMasterAccountInvestorResponse QueryManagedSubAccountTransferLogMasterAccountInvestor(ctx).Email(email).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).Transfers(transfers).TransferFunctionAccountType(transferFunctionAccountType).Execute()

Query Managed Sub Account Transfer Log For Investor Master Account (USER_DATA)


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
	email := "abc@test.com" // string | 
	startTime := int64(1623319461670) // int64 | Start Time
	endTime := int64(1641782889000) // int64 | End Time (The start time and end time interval cannot exceed half a year)
	page := int64(1) // int64 | Page
	limit := int64(1) // int64 | 
	transfers := "transfers_example" // string | Transfer Direction (FROM/TO) (optional)
	transferFunctionAccountType := models.QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameterSpot // QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountTransferLogMasterAccountInvestor(context.Background()).Email(email).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).Transfers(transfers).TransferFunctionAccountType(transferFunctionAccountType).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `ManagedSubAccountAPI.QueryManagedSubAccountTransferLogMasterAccountInvestor``: %v\n", err)
		return
	}

	// response from `QueryManagedSubAccountTransferLogMasterAccountInvestor`: QueryManagedSubAccountTransferLogMasterAccountInvestorResponse
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
 **startTime** | **int64** | Start Time | 
 **endTime** | **int64** | End Time (The start time and end time interval cannot exceed half a year) | 
 **page** | **int64** | Page | 
 **limit** | **int64** |  | 
 **transfers** | **string** | Transfer Direction (FROM/TO) | 
 **transferFunctionAccountType** | [**QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter**](QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter.md) |  | 

### Return type

[**QueryManagedSubAccountTransferLogMasterAccountInvestorResponse**](QueryManagedSubAccountTransferLogMasterAccountInvestorResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryManagedSubAccountTransferLogMasterAccountTrading

> QueryManagedSubAccountTransferLogMasterAccountTradingResponse QueryManagedSubAccountTransferLogMasterAccountTrading(ctx).Email(email).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).Transfers(transfers).TransferFunctionAccountType(transferFunctionAccountType).Execute()

Query Managed Sub Account Transfer Log For Trading Team Master Account (USER_DATA)


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
	email := "abc@test.com" // string | 
	startTime := int64(1623319461670) // int64 | Start Time
	endTime := int64(1641782889000) // int64 | End Time (The start time and end time interval cannot exceed half a year)
	page := int64(1) // int64 | 
	limit := int64(10) // int64 | 
	transfers := "transfers_example" // string | Transfer Direction (FROM/TO) (optional)
	transferFunctionAccountType := models.QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameterSpot // QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountTransferLogMasterAccountTrading(context.Background()).Email(email).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).Transfers(transfers).TransferFunctionAccountType(transferFunctionAccountType).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `ManagedSubAccountAPI.QueryManagedSubAccountTransferLogMasterAccountTrading``: %v\n", err)
		return
	}

	// response from `QueryManagedSubAccountTransferLogMasterAccountTrading`: QueryManagedSubAccountTransferLogMasterAccountTradingResponse
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
 **startTime** | **int64** | Start Time | 
 **endTime** | **int64** | End Time (The start time and end time interval cannot exceed half a year) | 
 **page** | **int64** |  | 
 **limit** | **int64** |  | 
 **transfers** | **string** | Transfer Direction (FROM/TO) | 
 **transferFunctionAccountType** | [**QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter**](QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter.md) |  | 

### Return type

[**QueryManagedSubAccountTransferLogMasterAccountTradingResponse**](QueryManagedSubAccountTransferLogMasterAccountTradingResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryManagedSubAccountTransferLogSubAccountTrading

> QueryManagedSubAccountTransferLogSubAccountTradingResponse QueryManagedSubAccountTransferLogSubAccountTrading(ctx).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).Transfers(transfers).TransferFunctionAccountType(transferFunctionAccountType).RecvWindow(recvWindow).Execute()

Query Managed Sub Account Transfer Log (For Trading Team Sub Account) (USER_DATA)


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
	startTime := int64(1623319461670) // int64 | Start Time
	endTime := int64(1641782889000) // int64 | End Time (The start time and end time interval cannot exceed half a year)
	page := int64(1) // int64 | 
	limit := int64(10) // int64 | 
	transfers := "transfers_example" // string | Transfer Direction (from/to) (optional)
	transferFunctionAccountType := models.QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameterSpot // QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountTransferLogSubAccountTrading(context.Background()).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).Transfers(transfers).TransferFunctionAccountType(transferFunctionAccountType).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `ManagedSubAccountAPI.QueryManagedSubAccountTransferLogSubAccountTrading``: %v\n", err)
		return
	}

	// response from `QueryManagedSubAccountTransferLogSubAccountTrading`: QueryManagedSubAccountTransferLogSubAccountTradingResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** | Start Time | 
 **endTime** | **int64** | End Time (The start time and end time interval cannot exceed half a year) | 
 **page** | **int64** |  | 
 **limit** | **int64** |  | 
 **transfers** | **string** | Transfer Direction (from/to) | 
 **transferFunctionAccountType** | [**QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter**](QueryManagedSubAccountTransferLogMasterAccountInvestorTransferFunctionAccountTypeParameter.md) |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryManagedSubAccountTransferLogSubAccountTradingResponse**](QueryManagedSubAccountTransferLogSubAccountTradingResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## WithdrawlAssetsFromTheManagedSubAccount

> WithdrawlAssetsFromTheManagedSubAccountResponse WithdrawlAssetsFromTheManagedSubAccount(ctx).FromEmail(fromEmail).Asset(asset).Amount(amount).TransferDate(transferDate).RecvWindow(recvWindow).Execute()

Withdrawl Assets From The Managed Sub-account (For Investor Master Account) (USER_DATA)


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
	fromEmail := "from@test.com" // string | 
	asset := "BTC" // string | 
	amount := float32(1.0) // float32 | 
	transferDate := int64(789) // int64 | Withdrawal will happen automatically on the selected date (UTC 0). If no date is selected, withdrawal takes effect immediately. (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.ManagedSubAccountAPI.WithdrawlAssetsFromTheManagedSubAccount(context.Background()).FromEmail(fromEmail).Asset(asset).Amount(amount).TransferDate(transferDate).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `ManagedSubAccountAPI.WithdrawlAssetsFromTheManagedSubAccount``: %v\n", err)
		return
	}

	// response from `WithdrawlAssetsFromTheManagedSubAccount`: WithdrawlAssetsFromTheManagedSubAccountResponse
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
 **asset** | **string** |  | 
 **amount** | **float32** |  | 
 **transferDate** | **int64** | Withdrawal will happen automatically on the selected date (UTC 0). If no date is selected, withdrawal takes effect immediately. | 
 **recvWindow** | **int64** |  | 

### Return type

[**WithdrawlAssetsFromTheManagedSubAccountResponse**](WithdrawlAssetsFromTheManagedSubAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

