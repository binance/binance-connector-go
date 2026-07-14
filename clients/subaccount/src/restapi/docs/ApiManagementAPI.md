# \ApiManagementAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AddIpRestrictionForSubAccountApiKey**](ApiManagementAPI.md#AddIpRestrictionForSubAccountApiKey) | **Post** /sapi/v2/sub-account/subAccountApi/ipRestriction | Add IP Restriction for Sub-Account API key (For Master Account) (USER_DATA)
[**CreateSubAccountApiKey**](ApiManagementAPI.md#CreateSubAccountApiKey) | **Post** /sapi/v1/sub-account/subAccountApi | Create Sub-account API Key (For Master Account) (USER_DATA)
[**DeleteIpListForASubAccountApiKey**](ApiManagementAPI.md#DeleteIpListForASubAccountApiKey) | **Delete** /sapi/v1/sub-account/subAccountApi/ipRestriction/ipList | Delete IP List For a Sub-account API Key (For Master Account) (USER_DATA)
[**DeleteSubAccountApiKey**](ApiManagementAPI.md#DeleteSubAccountApiKey) | **Delete** /sapi/v1/sub-account/subAccountApi | Delete Sub-account API Key (For Master Account) (USER_DATA)
[**GetIpRestrictionForASubAccountApiKey**](ApiManagementAPI.md#GetIpRestrictionForASubAccountApiKey) | **Get** /sapi/v1/sub-account/subAccountApi/ipRestriction | Get IP Restriction for a Sub-account API Key (For Master Account) (USER_DATA)
[**ModifySubAccountApiKeyPermission**](ApiManagementAPI.md#ModifySubAccountApiKeyPermission) | **Post** /sapi/v1/sub-account/subAccountApiPermission | Modify Sub-account API Key Permission (For Master Account) (USER_DATA)
[**QuerySubAccountApiKey**](ApiManagementAPI.md#QuerySubAccountApiKey) | **Get** /sapi/v1/sub-account/subAccountApi | Query Sub-account API Key (For Master Account) (USER_DATA)


## AddIpRestrictionForSubAccountApiKey

> AddIpRestrictionForSubAccountApiKeyResponse AddIpRestrictionForSubAccountApiKey(ctx).Email(email).SubAccountApiKey(subAccountApiKey).Status(status).IpAddress(ipAddress).RecvWindow(recvWindow).Execute()

Add IP Restriction for Sub-Account API key (For Master Account) (USER_DATA)


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
	subAccountApiKey := "k5V49ldtn4tszj6W3hystegdfvmGbqDzjmkCtpTvC0G74WhK7yd4rfCTo4lShf" // string | 
	status := int64(1) // int64 | IP Restriction status. 1 = IP Unrestricted. 2 = Restrict access to trusted IPs only.
	ipAddress := "69.210.67.14" // string | Insert static IP in batch, separated by commas. (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.ApiManagementAPI.AddIpRestrictionForSubAccountApiKey(context.Background()).Email(email).SubAccountApiKey(subAccountApiKey).Status(status).IpAddress(ipAddress).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `ApiManagementAPI.AddIpRestrictionForSubAccountApiKey``: %v\n", err)
		return
	}

	// response from `AddIpRestrictionForSubAccountApiKey`: AddIpRestrictionForSubAccountApiKeyResponse
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
 **subAccountApiKey** | **string** |  | 
 **status** | **int64** | IP Restriction status. 1 &#x3D; IP Unrestricted. 2 &#x3D; Restrict access to trusted IPs only. | 
 **ipAddress** | **string** | Insert static IP in batch, separated by commas. | 
 **recvWindow** | **int64** |  | 

### Return type

[**AddIpRestrictionForSubAccountApiKeyResponse**](AddIpRestrictionForSubAccountApiKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CreateSubAccountApiKey

> CreateSubAccountApiKeyResponse CreateSubAccountApiKey(ctx).Email(email).ApiName(apiName).Status(status).CanTrade(canTrade).CanMarginLoanRepay(canMarginLoanRepay).CanFuturesTrade(canFuturesTrade).CanUniversalTransfer(canUniversalTransfer).CanVanillaOptions(canVanillaOptions).IpAddress(ipAddress).ThirdPartyName(thirdPartyName).PublicKey(publicKey).RecvWindow(recvWindow).Execute()

Create Sub-account API Key (For Master Account) (USER_DATA)


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
	email := "123@test.com" // string | Sub-account email
	apiName := "myKey" // string | API Key name
	status := int64(2) // int64 | IP restriction status. 1 = unrestricted, 2 = restricted to trusted IPs, 3 = third-party IP restriction
	canTrade := true // bool | Spot & Margin trading permission, default false (optional)
	canMarginLoanRepay := false // bool | Margin borrow/repay permission, default false (optional)
	canFuturesTrade := false // bool | Futures trading permission, default false (optional)
	canUniversalTransfer := false // bool | Universal transfer permission, default false (optional)
	canVanillaOptions := false // bool | Vanilla options permission, default false (optional)
	ipAddress := "69.210.67.14" // string | Required when status=2. IP address list, max 500 chars (optional)
	thirdPartyName := "thirdParty" // string | Required when status=3. Third-party name (optional)
	publicKey := "publicKey_example" // string | Ed25519 public key (optional, for Ed25519 type API Key) (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.ApiManagementAPI.CreateSubAccountApiKey(context.Background()).Email(email).ApiName(apiName).Status(status).CanTrade(canTrade).CanMarginLoanRepay(canMarginLoanRepay).CanFuturesTrade(canFuturesTrade).CanUniversalTransfer(canUniversalTransfer).CanVanillaOptions(canVanillaOptions).IpAddress(ipAddress).ThirdPartyName(thirdPartyName).PublicKey(publicKey).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `ApiManagementAPI.CreateSubAccountApiKey``: %v\n", err)
		return
	}

	// response from `CreateSubAccountApiKey`: CreateSubAccountApiKeyResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub-account email | 
 **apiName** | **string** | API Key name | 
 **status** | **int64** | IP restriction status. 1 &#x3D; unrestricted, 2 &#x3D; restricted to trusted IPs, 3 &#x3D; third-party IP restriction | 
 **canTrade** | **bool** | Spot &amp; Margin trading permission, default false | 
 **canMarginLoanRepay** | **bool** | Margin borrow/repay permission, default false | 
 **canFuturesTrade** | **bool** | Futures trading permission, default false | 
 **canUniversalTransfer** | **bool** | Universal transfer permission, default false | 
 **canVanillaOptions** | **bool** | Vanilla options permission, default false | 
 **ipAddress** | **string** | Required when status&#x3D;2. IP address list, max 500 chars | 
 **thirdPartyName** | **string** | Required when status&#x3D;3. Third-party name | 
 **publicKey** | **string** | Ed25519 public key (optional, for Ed25519 type API Key) | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateSubAccountApiKeyResponse**](CreateSubAccountApiKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## DeleteIpListForASubAccountApiKey

> DeleteIpListForASubAccountApiKeyResponse DeleteIpListForASubAccountApiKey(ctx).Email(email).SubAccountApiKey(subAccountApiKey).IpAddress(ipAddress).RecvWindow(recvWindow).Execute()

Delete IP List For a Sub-account API Key (For Master Account) (USER_DATA)


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
	subAccountApiKey := "k5V49ldtn4tszj6W3hystegdfvmGbqDzjmkCtpTvC0G74WhK7yd4rfCTo4lShf" // string | 
	ipAddress := "69.210.67.14" // string | IPs to be deleted. Can be added in batches, separated by commas
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.ApiManagementAPI.DeleteIpListForASubAccountApiKey(context.Background()).Email(email).SubAccountApiKey(subAccountApiKey).IpAddress(ipAddress).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `ApiManagementAPI.DeleteIpListForASubAccountApiKey``: %v\n", err)
		return
	}

	// response from `DeleteIpListForASubAccountApiKey`: DeleteIpListForASubAccountApiKeyResponse
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
 **subAccountApiKey** | **string** |  | 
 **ipAddress** | **string** | IPs to be deleted. Can be added in batches, separated by commas | 
 **recvWindow** | **int64** |  | 

### Return type

[**DeleteIpListForASubAccountApiKeyResponse**](DeleteIpListForASubAccountApiKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## DeleteSubAccountApiKey

> map[string]interface{} DeleteSubAccountApiKey(ctx).Email(email).SubAccountApiKey(subAccountApiKey).RecvWindow(recvWindow).Execute()

Delete Sub-account API Key (For Master Account) (USER_DATA)


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
	email := "123@test.com" // string | Sub-account email
	subAccountApiKey := "k5V49ldtn4tszj6W3hystegdfvmGbqDzjmkCtpTvC0G74WhK7yd4rfCTo4lShf" // string | The sub-account API Key to be deleted
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.ApiManagementAPI.DeleteSubAccountApiKey(context.Background()).Email(email).SubAccountApiKey(subAccountApiKey).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `ApiManagementAPI.DeleteSubAccountApiKey``: %v\n", err)
		return
	}

	// response from `DeleteSubAccountApiKey`: map[string]interface{}
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub-account email | 
 **subAccountApiKey** | **string** | The sub-account API Key to be deleted | 
 **recvWindow** | **int64** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetIpRestrictionForASubAccountApiKey

> GetIpRestrictionForASubAccountApiKeyResponse GetIpRestrictionForASubAccountApiKey(ctx).Email(email).SubAccountApiKey(subAccountApiKey).RecvWindow(recvWindow).Execute()

Get IP Restriction for a Sub-account API Key (For Master Account) (USER_DATA)


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
	subAccountApiKey := "k5V49ldtn4tszj6W3hystegdfvmGbqDzjmkCtpTvC0G74WhK7yd4rfCTo4lShf" // string | 
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.ApiManagementAPI.GetIpRestrictionForASubAccountApiKey(context.Background()).Email(email).SubAccountApiKey(subAccountApiKey).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `ApiManagementAPI.GetIpRestrictionForASubAccountApiKey``: %v\n", err)
		return
	}

	// response from `GetIpRestrictionForASubAccountApiKey`: GetIpRestrictionForASubAccountApiKeyResponse
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
 **subAccountApiKey** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetIpRestrictionForASubAccountApiKeyResponse**](GetIpRestrictionForASubAccountApiKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## ModifySubAccountApiKeyPermission

> ModifySubAccountApiKeyPermissionResponse ModifySubAccountApiKeyPermission(ctx).Email(email).SubAccountApiKey(subAccountApiKey).CanTrade(canTrade).CanMarginLoanRepay(canMarginLoanRepay).CanFuturesTrade(canFuturesTrade).CanUniversalTransfer(canUniversalTransfer).CanVanillaOptions(canVanillaOptions).RecvWindow(recvWindow).Execute()

Modify Sub-account API Key Permission (For Master Account) (USER_DATA)


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
	email := "123@test.com" // string | Sub-account email
	subAccountApiKey := "k5V49ldtn4tszj6W3hystegdfvmGbqDzjmkCtpTvC0G74WhK7yd4rfCTo4lShf" // string | Sub-account API Key
	canTrade := true // bool | Spot & Margin trading permission (optional)
	canMarginLoanRepay := false // bool | Margin borrow/repay permission (optional)
	canFuturesTrade := true // bool | Futures trading permission (optional)
	canUniversalTransfer := false // bool | Universal transfer permission (optional)
	canVanillaOptions := false // bool | Vanilla options permission (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.ApiManagementAPI.ModifySubAccountApiKeyPermission(context.Background()).Email(email).SubAccountApiKey(subAccountApiKey).CanTrade(canTrade).CanMarginLoanRepay(canMarginLoanRepay).CanFuturesTrade(canFuturesTrade).CanUniversalTransfer(canUniversalTransfer).CanVanillaOptions(canVanillaOptions).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `ApiManagementAPI.ModifySubAccountApiKeyPermission``: %v\n", err)
		return
	}

	// response from `ModifySubAccountApiKeyPermission`: ModifySubAccountApiKeyPermissionResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub-account email | 
 **subAccountApiKey** | **string** | Sub-account API Key | 
 **canTrade** | **bool** | Spot &amp; Margin trading permission | 
 **canMarginLoanRepay** | **bool** | Margin borrow/repay permission | 
 **canFuturesTrade** | **bool** | Futures trading permission | 
 **canUniversalTransfer** | **bool** | Universal transfer permission | 
 **canVanillaOptions** | **bool** | Vanilla options permission | 
 **recvWindow** | **int64** |  | 

### Return type

[**ModifySubAccountApiKeyPermissionResponse**](ModifySubAccountApiKeyPermissionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QuerySubAccountApiKey

> QuerySubAccountApiKeyResponse QuerySubAccountApiKey(ctx).Email(email).SubAccountApiKey(subAccountApiKey).Page(page).Size(size).RecvWindow(recvWindow).Execute()

Query Sub-account API Key (For Master Account) (USER_DATA)


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
	email := "123@test.com" // string | Sub-account email
	subAccountApiKey := "k5V49ldtn4tszj6W3hystegdfvmGbqDzjmkCtpTvC0G74WhK7yd4rfCTo4lShf" // string | Specify an API Key for exact match (optional)
	page := int64(1) // int64 | Page number, default 1, minimum 1 (optional)
	size := int64(30) // int64 | Page size, default 30, maximum 100 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSubAccountClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.ApiManagementAPI.QuerySubAccountApiKey(context.Background()).Email(email).SubAccountApiKey(subAccountApiKey).Page(page).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `ApiManagementAPI.QuerySubAccountApiKey``: %v\n", err)
		return
	}

	// response from `QuerySubAccountApiKey`: QuerySubAccountApiKeyResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub-account email | 
 **subAccountApiKey** | **string** | Specify an API Key for exact match | 
 **page** | **int64** | Page number, default 1, minimum 1 | 
 **size** | **int64** | Page size, default 30, maximum 100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**QuerySubAccountApiKeyResponse**](QuerySubAccountApiKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

