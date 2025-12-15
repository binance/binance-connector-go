# \CapitalAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AllCoinsInformation**](CapitalAPI.md#AllCoinsInformation) | **Get** /sapi/v1/capital/config/getall | All Coins&#39; Information (USER_DATA)
[**DepositAddress**](CapitalAPI.md#DepositAddress) | **Get** /sapi/v1/capital/deposit/address | Deposit Address(supporting network) (USER_DATA)
[**DepositHistory**](CapitalAPI.md#DepositHistory) | **Get** /sapi/v1/capital/deposit/hisrec | Deposit History (supporting network) (USER_DATA)
[**FetchDepositAddressListWithNetwork**](CapitalAPI.md#FetchDepositAddressListWithNetwork) | **Get** /sapi/v1/capital/deposit/address/list | Fetch deposit address list with network(USER_DATA)
[**FetchWithdrawAddressList**](CapitalAPI.md#FetchWithdrawAddressList) | **Get** /sapi/v1/capital/withdraw/address/list | Fetch withdraw address list (USER_DATA)
[**FetchWithdrawQuota**](CapitalAPI.md#FetchWithdrawQuota) | **Get** /sapi/v1/capital/withdraw/quota | Fetch withdraw quota (USER_DATA)
[**OneClickArrivalDepositApply**](CapitalAPI.md#OneClickArrivalDepositApply) | **Post** /sapi/v1/capital/deposit/credit-apply | One click arrival deposit apply (for expired address deposit) (USER_DATA)
[**Withdraw**](CapitalAPI.md#Withdraw) | **Post** /sapi/v1/capital/withdraw/apply | Withdraw(USER_DATA)
[**WithdrawHistory**](CapitalAPI.md#WithdrawHistory) | **Get** /sapi/v1/capital/withdraw/history | Withdraw History (supporting network) (USER_DATA)


## AllCoinsInformation

> AllCoinsInformationResponse AllCoinsInformation(ctx).RecvWindow(recvWindow).Execute()

All Coins' Information (USER_DATA)


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
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.CapitalAPI.AllCoinsInformation(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `CapitalAPI.AllCoinsInformation``: %v\n", err)
		return
	}

	// response from `AllCoinsInformation`: AllCoinsInformationResponse
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

[**AllCoinsInformationResponse**](AllCoinsInformationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## DepositAddress

> DepositAddressResponse DepositAddress(ctx).Coin(coin).Network(network).Amount(amount).RecvWindow(recvWindow).Execute()

Deposit Address(supporting network) (USER_DATA)


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
	coin := "coin_example" // string | 
	network := "network_example" // string |  (optional)
	amount := float32(1.0) // float32 |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.CapitalAPI.DepositAddress(context.Background()).Coin(coin).Network(network).Amount(amount).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `CapitalAPI.DepositAddress``: %v\n", err)
		return
	}

	// response from `DepositAddress`: DepositAddressResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **coin** | **string** |  | 
 **network** | **string** |  | 
 **amount** | **float32** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**DepositAddressResponse**](DepositAddressResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## DepositHistory

> DepositHistoryResponse DepositHistory(ctx).IncludeSource(includeSource).Coin(coin).Status(status).StartTime(startTime).EndTime(endTime).Offset(offset).Limit(limit).RecvWindow(recvWindow).TxId(txId).Execute()

Deposit History (supporting network) (USER_DATA)


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
	includeSource := false // bool | Default: `false`, return `sourceAddress`field when set to `true` (optional)
	coin := "coin_example" // string |  (optional)
	status := int64(789) // int64 | 0(0:Email Sent, 2:Awaiting Approval 3:Rejected 4:Processing 6:Completed) (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	offset := int64(0) // int64 | Default: 0 (optional)
	limit := int64(7) // int64 | min 7, max 30, default 7 (optional)
	recvWindow := int64(5000) // int64 |  (optional)
	txId := "1" // string |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.CapitalAPI.DepositHistory(context.Background()).IncludeSource(includeSource).Coin(coin).Status(status).StartTime(startTime).EndTime(endTime).Offset(offset).Limit(limit).RecvWindow(recvWindow).TxId(txId).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `CapitalAPI.DepositHistory``: %v\n", err)
		return
	}

	// response from `DepositHistory`: DepositHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **includeSource** | **bool** | Default: &#x60;false&#x60;, return &#x60;sourceAddress&#x60;field when set to &#x60;true&#x60; | 
 **coin** | **string** |  | 
 **status** | **int64** | 0(0:Email Sent, 2:Awaiting Approval 3:Rejected 4:Processing 6:Completed) | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **offset** | **int64** | Default: 0 | 
 **limit** | **int64** | min 7, max 30, default 7 | 
 **recvWindow** | **int64** |  | 
 **txId** | **string** |  | 

### Return type

[**DepositHistoryResponse**](DepositHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## FetchDepositAddressListWithNetwork

> FetchDepositAddressListWithNetworkResponse FetchDepositAddressListWithNetwork(ctx).Coin(coin).Network(network).Execute()

Fetch deposit address list with network(USER_DATA)


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
	coin := "coin_example" // string | 
	network := "network_example" // string |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.CapitalAPI.FetchDepositAddressListWithNetwork(context.Background()).Coin(coin).Network(network).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `CapitalAPI.FetchDepositAddressListWithNetwork``: %v\n", err)
		return
	}

	// response from `FetchDepositAddressListWithNetwork`: FetchDepositAddressListWithNetworkResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **coin** | **string** |  | 
 **network** | **string** |  | 

### Return type

[**FetchDepositAddressListWithNetworkResponse**](FetchDepositAddressListWithNetworkResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## FetchWithdrawAddressList

> FetchWithdrawAddressListResponse FetchWithdrawAddressList(ctx).Execute()

Fetch withdraw address list (USER_DATA)


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

	resp, err := apiClient.RestApi.CapitalAPI.FetchWithdrawAddressList(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `CapitalAPI.FetchWithdrawAddressList``: %v\n", err)
		return
	}

	// response from `FetchWithdrawAddressList`: FetchWithdrawAddressListResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

This endpoint does not need any parameter.

### Return type

[**FetchWithdrawAddressListResponse**](FetchWithdrawAddressListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## FetchWithdrawQuota

> FetchWithdrawQuotaResponse FetchWithdrawQuota(ctx).Execute()

Fetch withdraw quota (USER_DATA)


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

	resp, err := apiClient.RestApi.CapitalAPI.FetchWithdrawQuota(context.Background()).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `CapitalAPI.FetchWithdrawQuota``: %v\n", err)
		return
	}

	// response from `FetchWithdrawQuota`: FetchWithdrawQuotaResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

This endpoint does not need any parameter.

### Return type

[**FetchWithdrawQuotaResponse**](FetchWithdrawQuotaResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## OneClickArrivalDepositApply

> OneClickArrivalDepositApplyResponse OneClickArrivalDepositApply(ctx).DepositId(depositId).TxId(txId).SubAccountId(subAccountId).SubUserId(subUserId).Execute()

One click arrival deposit apply (for expired address deposit) (USER_DATA)


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
	depositId := int64(1) // int64 | Deposit record Id, priority use (optional)
	txId := "1" // string |  (optional)
	subAccountId := int64(1) // int64 | Sub-accountId of Cloud user (optional)
	subUserId := int64(1) // int64 | Sub-userId of parent user (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.CapitalAPI.OneClickArrivalDepositApply(context.Background()).DepositId(depositId).TxId(txId).SubAccountId(subAccountId).SubUserId(subUserId).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `CapitalAPI.OneClickArrivalDepositApply``: %v\n", err)
		return
	}

	// response from `OneClickArrivalDepositApply`: OneClickArrivalDepositApplyResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **depositId** | **int64** | Deposit record Id, priority use | 
 **txId** | **string** |  | 
 **subAccountId** | **int64** | Sub-accountId of Cloud user | 
 **subUserId** | **int64** | Sub-userId of parent user | 

### Return type

[**OneClickArrivalDepositApplyResponse**](OneClickArrivalDepositApplyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## Withdraw

> WithdrawResponse Withdraw(ctx).Coin(coin).Address(address).Amount(amount).WithdrawOrderId(withdrawOrderId).Network(network).AddressTag(addressTag).TransactionFeeFlag(transactionFeeFlag).Name(name).WalletType(walletType).RecvWindow(recvWindow).Execute()

Withdraw(USER_DATA)


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
	coin := "coin_example" // string | 
	address := "address_example" // string | 
	amount := float32(1.0) // float32 | 
	withdrawOrderId := "1" // string | client side id for withdrawal, if provided in POST `/sapi/v1/capital/withdraw/apply`, can be used here for query. (optional)
	network := "network_example" // string |  (optional)
	addressTag := "addressTag_example" // string | Secondary address identifier for coins like XRP,XMR etc. (optional)
	transactionFeeFlag := false // bool | When making internal transfer, `true` for returning the fee to the destination account; `false` for returning the fee back to the departure account. Default `false`. (optional)
	name := "name_example" // string | Description of the address. Address book cap is 200, space in name should be encoded into `%20` (optional)
	walletType := int64(0) // int64 | The wallet type for withdraw，0-spot wallet ，1-funding wallet. Default walletType is the current \"selected wallet\" under wallet->Fiat and Spot/Funding->Deposit (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.CapitalAPI.Withdraw(context.Background()).Coin(coin).Address(address).Amount(amount).WithdrawOrderId(withdrawOrderId).Network(network).AddressTag(addressTag).TransactionFeeFlag(transactionFeeFlag).Name(name).WalletType(walletType).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `CapitalAPI.Withdraw``: %v\n", err)
		return
	}

	// response from `Withdraw`: WithdrawResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **coin** | **string** |  | 
 **address** | **string** |  | 
 **amount** | **float32** |  | 
 **withdrawOrderId** | **string** | client side id for withdrawal, if provided in POST &#x60;/sapi/v1/capital/withdraw/apply&#x60;, can be used here for query. | 
 **network** | **string** |  | 
 **addressTag** | **string** | Secondary address identifier for coins like XRP,XMR etc. | 
 **transactionFeeFlag** | **bool** | When making internal transfer, &#x60;true&#x60; for returning the fee to the destination account; &#x60;false&#x60; for returning the fee back to the departure account. Default &#x60;false&#x60;. | 
 **name** | **string** | Description of the address. Address book cap is 200, space in name should be encoded into &#x60;%20&#x60; | 
 **walletType** | **int64** | The wallet type for withdraw，0-spot wallet ，1-funding wallet. Default walletType is the current \&quot;selected wallet\&quot; under wallet-&gt;Fiat and Spot/Funding-&gt;Deposit | 
 **recvWindow** | **int64** |  | 

### Return type

[**WithdrawResponse**](WithdrawResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## WithdrawHistory

> WithdrawHistoryResponse WithdrawHistory(ctx).Coin(coin).WithdrawOrderId(withdrawOrderId).Status(status).Offset(offset).Limit(limit).IdList(idList).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

Withdraw History (supporting network) (USER_DATA)


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
	coin := "coin_example" // string |  (optional)
	withdrawOrderId := "1" // string | client side id for withdrawal, if provided in POST `/sapi/v1/capital/withdraw/apply`, can be used here for query. (optional)
	status := int64(789) // int64 | 0(0:Email Sent, 2:Awaiting Approval 3:Rejected 4:Processing 6:Completed) (optional)
	offset := int64(0) // int64 | Default: 0 (optional)
	limit := int64(7) // int64 | min 7, max 30, default 7 (optional)
	idList := "idList_example" // string | id list returned in the response of POST `/sapi/v1/capital/withdraw/apply`, separated by `,` (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.CapitalAPI.WithdrawHistory(context.Background()).Coin(coin).WithdrawOrderId(withdrawOrderId).Status(status).Offset(offset).Limit(limit).IdList(idList).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `CapitalAPI.WithdrawHistory``: %v\n", err)
		return
	}

	// response from `WithdrawHistory`: WithdrawHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **coin** | **string** |  | 
 **withdrawOrderId** | **string** | client side id for withdrawal, if provided in POST &#x60;/sapi/v1/capital/withdraw/apply&#x60;, can be used here for query. | 
 **status** | **int64** | 0(0:Email Sent, 2:Awaiting Approval 3:Rejected 4:Processing 6:Completed) | 
 **offset** | **int64** | Default: 0 | 
 **limit** | **int64** | min 7, max 30, default 7 | 
 **idList** | **string** | id list returned in the response of POST &#x60;/sapi/v1/capital/withdraw/apply&#x60;, separated by &#x60;,&#x60; | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**WithdrawHistoryResponse**](WithdrawHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

