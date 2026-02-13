# \AccountAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AccountApiTradingStatus**](AccountAPI.md#AccountApiTradingStatus) | **Get** /sapi/v1/account/apiTradingStatus | Account API Trading Status (USER_DATA)
[**AccountInfo**](AccountAPI.md#AccountInfo) | **Get** /sapi/v1/account/info | Account info (USER_DATA)
[**AccountStatus**](AccountAPI.md#AccountStatus) | **Get** /sapi/v1/account/status | Account Status (USER_DATA)
[**DailyAccountSnapshot**](AccountAPI.md#DailyAccountSnapshot) | **Get** /sapi/v1/accountSnapshot | Daily Account Snapshot (USER_DATA)
[**DisableFastWithdrawSwitch**](AccountAPI.md#DisableFastWithdrawSwitch) | **Post** /sapi/v1/account/disableFastWithdrawSwitch | Disable Fast Withdraw Switch (USER_DATA)
[**EnableFastWithdrawSwitch**](AccountAPI.md#EnableFastWithdrawSwitch) | **Post** /sapi/v1/account/enableFastWithdrawSwitch | Enable Fast Withdraw Switch (USER_DATA)
[**GetApiKeyPermission**](AccountAPI.md#GetApiKeyPermission) | **Get** /sapi/v1/account/apiRestrictions | Get API Key Permission (USER_DATA)


## AccountApiTradingStatus

> AccountApiTradingStatusResponse AccountApiTradingStatus(ctx).RecvWindow(recvWindow).Execute()

Account API Trading Status (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.AccountApiTradingStatus(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.AccountApiTradingStatus``: %v\n", err)
		return
	}

	// response from `AccountApiTradingStatus`: AccountApiTradingStatusResponse
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

[**AccountApiTradingStatusResponse**](AccountApiTradingStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## AccountInfo

> AccountInfoResponse AccountInfo(ctx).RecvWindow(recvWindow).Execute()

Account info (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.AccountInfo(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.AccountInfo``: %v\n", err)
		return
	}

	// response from `AccountInfo`: AccountInfoResponse
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

[**AccountInfoResponse**](AccountInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## AccountStatus

> AccountStatusResponse AccountStatus(ctx).RecvWindow(recvWindow).Execute()

Account Status (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.AccountStatus(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.AccountStatus``: %v\n", err)
		return
	}

	// response from `AccountStatus`: AccountStatusResponse
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

[**AccountStatusResponse**](AccountStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## DailyAccountSnapshot

> DailyAccountSnapshotResponse DailyAccountSnapshot(ctx).Type(type_).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Daily Account Snapshot (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	type_ := "type__example" // string | 
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

	resp, err := apiClient.RestApi.AccountAPI.DailyAccountSnapshot(context.Background()).Type(type_).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.DailyAccountSnapshot``: %v\n", err)
		return
	}

	// response from `DailyAccountSnapshot`: DailyAccountSnapshotResponse
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
 **limit** | **int64** | min 7, max 30, default 7 | 
 **recvWindow** | **int64** |  | 

### Return type

[**DailyAccountSnapshotResponse**](DailyAccountSnapshotResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## DisableFastWithdrawSwitch

> DisableFastWithdrawSwitch(ctx).RecvWindow(recvWindow).Execute()

Disable Fast Withdraw Switch (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	 err := apiClient.RestApi.AccountAPI.DisableFastWithdrawSwitch(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.DisableFastWithdrawSwitch``: %v\n", err)
		return
	}

}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: Not defined

[[Back to README]](../../../README.md)


## EnableFastWithdrawSwitch

> EnableFastWithdrawSwitch(ctx).RecvWindow(recvWindow).Execute()

Enable Fast Withdraw Switch (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	 err := apiClient.RestApi.AccountAPI.EnableFastWithdrawSwitch(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.EnableFastWithdrawSwitch``: %v\n", err)
		return
	}

}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: Not defined

[[Back to README]](../../../README.md)


## GetApiKeyPermission

> GetApiKeyPermissionResponse GetApiKeyPermission(ctx).RecvWindow(recvWindow).Execute()

Get API Key Permission (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetApiKeyPermission(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetApiKeyPermission``: %v\n", err)
		return
	}

	// response from `GetApiKeyPermission`: GetApiKeyPermissionResponse
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

[**GetApiKeyPermissionResponse**](GetApiKeyPermissionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

