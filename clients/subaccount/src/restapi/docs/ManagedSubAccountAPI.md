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
[**QueryManagedSubAccountTransferLogMasterAccountInvestor**](ManagedSubAccountAPI.md#QueryManagedSubAccountTransferLogMasterAccountInvestor) | **Get** /sapi/v1/managed-subaccount/queryTransLogForInvestor | Query Managed Sub Account Transfer Log (For Investor Master Account) (USER_DATA)
[**QueryManagedSubAccountTransferLogMasterAccountTrading**](ManagedSubAccountAPI.md#QueryManagedSubAccountTransferLogMasterAccountTrading) | **Get** /sapi/v1/managed-subaccount/queryTransLogForTradeParent | Query Managed Sub Account Transfer Log (For Trading Team Master Account) (USER_DATA)
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
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	toEmail := "toEmail_example" // string | 
	asset := "asset_example" // string | 
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
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	email := "sub-account-email@email.com" // string | [Sub-account email](#email-address)
	coin := "coin_example" // string | 
	network := "network_example" // string | networks can be found in `GET /sapi/v1/capital/deposit/address` (optional)
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
 **email** | **string** | [Sub-account email](#email-address) | 
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
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	email := "sub-account-email@email.com" // string | [Sub-account email](#email-address)
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
 **email** | **string** | [Sub-account email](#email-address) | 
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
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	email := "sub-account-email@email.com" // string | [Sub-account email](#email-address)
	accountType := "accountType_example" // string | No input or input \"MARGIN\" to get Cross Margin account details. Input \"ISOLATED_MARGIN\" to get Isolated Margin account details. (optional)

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
 **email** | **string** | [Sub-account email](#email-address) | 
 **accountType** | **string** | No input or input \&quot;MARGIN\&quot; to get Cross Margin account details. Input \&quot;ISOLATED_MARGIN\&quot; to get Isolated Margin account details. | 

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
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	email := "email_example" // string | Managed sub-account email (optional)
	page := int64(1) // int64 | Default value: 1 (optional)
	limit := int64(1) // int64 | Default value: 1, Max value: 200 (optional)
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
 **email** | **string** | Managed sub-account email | 
 **page** | **int64** | Default value: 1 | 
 **limit** | **int64** | Default value: 1, Max value: 200 | 
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
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	email := "sub-account-email@email.com" // string | [Sub-account email](#email-address)
	accountType := "accountType_example" // string | No input or input \"MARGIN\" to get Cross Margin account details. Input \"ISOLATED_MARGIN\" to get Isolated Margin account details. (optional)

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
 **email** | **string** | [Sub-account email](#email-address) | 
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
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	email := "sub-account-email@email.com" // string | [Sub-account email](#email-address)
	type_ := "type__example" // string | \"SPOT\", \"MARGIN\"（cross）, \"FUTURES\"（UM）
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	limit := int64(1) // int64 | Default value: 1, Max value: 200 (optional)
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
 **email** | **string** | [Sub-account email](#email-address) | 
 **type_** | **string** | \&quot;SPOT\&quot;, \&quot;MARGIN\&quot;（cross）, \&quot;FUTURES\&quot;（UM） | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int64** | Default value: 1, Max value: 200 | 
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

Query Managed Sub Account Transfer Log (For Investor Master Account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	email := "sub-account-email@email.com" // string | [Sub-account email](#email-address)
	startTime := int64(1623319461670) // int64 | Start Time
	endTime := int64(1641782889000) // int64 | End Time (The start time and end time interval cannot exceed half a year)
	page := int64(789) // int64 | Page
	limit := int64(789) // int64 | Limit (Max: 500)
	transfers := "transfers_example" // string | Transfer Direction (FROM/TO) (optional)
	transferFunctionAccountType := "transferFunctionAccountType_example" // string | Transfer function account type (SPOT/MARGIN/ISOLATED_MARGIN/USDT_FUTURE/COIN_FUTURE) (optional)

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
 **email** | **string** | [Sub-account email](#email-address) | 
 **startTime** | **int64** | Start Time | 
 **endTime** | **int64** | End Time (The start time and end time interval cannot exceed half a year) | 
 **page** | **int64** | Page | 
 **limit** | **int64** | Limit (Max: 500) | 
 **transfers** | **string** | Transfer Direction (FROM/TO) | 
 **transferFunctionAccountType** | **string** | Transfer function account type (SPOT/MARGIN/ISOLATED_MARGIN/USDT_FUTURE/COIN_FUTURE) | 

### Return type

[**QueryManagedSubAccountTransferLogMasterAccountInvestorResponse**](QueryManagedSubAccountTransferLogMasterAccountInvestorResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryManagedSubAccountTransferLogMasterAccountTrading

> QueryManagedSubAccountTransferLogMasterAccountTradingResponse QueryManagedSubAccountTransferLogMasterAccountTrading(ctx).Email(email).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).Transfers(transfers).TransferFunctionAccountType(transferFunctionAccountType).Execute()

Query Managed Sub Account Transfer Log (For Trading Team Master Account) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	email := "sub-account-email@email.com" // string | [Sub-account email](#email-address)
	startTime := int64(1623319461670) // int64 | Start Time
	endTime := int64(1641782889000) // int64 | End Time (The start time and end time interval cannot exceed half a year)
	page := int64(789) // int64 | Page
	limit := int64(789) // int64 | Limit (Max: 500)
	transfers := "transfers_example" // string | Transfer Direction (FROM/TO) (optional)
	transferFunctionAccountType := "transferFunctionAccountType_example" // string | Transfer function account type (SPOT/MARGIN/ISOLATED_MARGIN/USDT_FUTURE/COIN_FUTURE) (optional)

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
 **email** | **string** | [Sub-account email](#email-address) | 
 **startTime** | **int64** | Start Time | 
 **endTime** | **int64** | End Time (The start time and end time interval cannot exceed half a year) | 
 **page** | **int64** | Page | 
 **limit** | **int64** | Limit (Max: 500) | 
 **transfers** | **string** | Transfer Direction (FROM/TO) | 
 **transferFunctionAccountType** | **string** | Transfer function account type (SPOT/MARGIN/ISOLATED_MARGIN/USDT_FUTURE/COIN_FUTURE) | 

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
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	startTime := int64(1623319461670) // int64 | Start Time
	endTime := int64(1641782889000) // int64 | End Time (The start time and end time interval cannot exceed half a year)
	page := int64(789) // int64 | Page
	limit := int64(789) // int64 | Limit (Max: 500)
	transfers := "transfers_example" // string | Transfer Direction (FROM/TO) (optional)
	transferFunctionAccountType := "transferFunctionAccountType_example" // string | Transfer function account type (SPOT/MARGIN/ISOLATED_MARGIN/USDT_FUTURE/COIN_FUTURE) (optional)
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
 **page** | **int64** | Page | 
 **limit** | **int64** | Limit (Max: 500) | 
 **transfers** | **string** | Transfer Direction (FROM/TO) | 
 **transferFunctionAccountType** | **string** | Transfer function account type (SPOT/MARGIN/ISOLATED_MARGIN/USDT_FUTURE/COIN_FUTURE) | 
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
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	fromEmail := "fromEmail_example" // string | 
	asset := "asset_example" // string | 
	amount := float32(1.0) // float32 | 
	transferDate := int64(789) // int64 | Withdrawals is automatically occur on the transfer date(UTC0). If a date is not selected, the withdrawal occurs right now (optional)
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
 **transferDate** | **int64** | Withdrawals is automatically occur on the transfer date(UTC0). If a date is not selected, the withdrawal occurs right now | 
 **recvWindow** | **int64** |  | 

### Return type

[**WithdrawlAssetsFromTheManagedSubAccountResponse**](WithdrawlAssetsFromTheManagedSubAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

