# \ApiManagementAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AddIpRestrictionForSubAccountApiKey**](ApiManagementAPI.md#AddIpRestrictionForSubAccountApiKey) | **Post** /sapi/v2/sub-account/subAccountApi/ipRestriction | Add IP Restriction for Sub-Account API key (For Master Account) (USER_DATA)
[**DeleteIpListForASubAccountApiKey**](ApiManagementAPI.md#DeleteIpListForASubAccountApiKey) | **Delete** /sapi/v1/sub-account/subAccountApi/ipRestriction/ipList | Delete IP List For a Sub-account API Key (For Master Account) (USER_DATA)
[**GetIpRestrictionForASubAccountApiKey**](ApiManagementAPI.md#GetIpRestrictionForASubAccountApiKey) | **Get** /sapi/v1/sub-account/subAccountApi/ipRestriction | Get IP Restriction for a Sub-account API Key (For Master Account) (USER_DATA)


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
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	email := "sub-account-email@email.com" // string | [Sub-account email](#email-address)
	subAccountApiKey := "subAccountApiKey_example" // string | 
	status := "status_example" // string | IP Restriction status. 1 = IP Unrestricted. 2 = Restrict access to trusted IPs only.
	ipAddress := "ipAddress_example" // string | Insert static IP in batch, separated by commas. (optional)
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
 **email** | **string** | [Sub-account email](#email-address) | 
 **subAccountApiKey** | **string** |  | 
 **status** | **string** | IP Restriction status. 1 &#x3D; IP Unrestricted. 2 &#x3D; Restrict access to trusted IPs only. | 
 **ipAddress** | **string** | Insert static IP in batch, separated by commas. | 
 **recvWindow** | **int64** |  | 

### Return type

[**AddIpRestrictionForSubAccountApiKeyResponse**](AddIpRestrictionForSubAccountApiKeyResponse.md)

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
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	email := "sub-account-email@email.com" // string | [Sub-account email](#email-address)
	subAccountApiKey := "subAccountApiKey_example" // string | 
	ipAddress := "ipAddress_example" // string | IPs to be deleted. Can be added in batches, separated by commas
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
 **email** | **string** | [Sub-account email](#email-address) | 
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
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	email := "sub-account-email@email.com" // string | [Sub-account email](#email-address)
	subAccountApiKey := "subAccountApiKey_example" // string | 
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
 **email** | **string** | [Sub-account email](#email-address) | 
 **subAccountApiKey** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetIpRestrictionForASubAccountApiKeyResponse**](GetIpRestrictionForASubAccountApiKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

