# \RedeemAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**BatchRedeem**](RedeemAPI.md#BatchRedeem) | **Post** /sapi/v1/w3w/wallet/prediction/batch-redeem | Batch Redeem (TRADE)
[**GetRedeemStatus**](RedeemAPI.md#GetRedeemStatus) | **Get** /sapi/v1/w3w/wallet/prediction/redeem/status | Get Redeem Status (USER_DATA)


## BatchRedeem

> BatchRedeemResponse BatchRedeem(ctx).WalletAddress(walletAddress).WalletId(walletId).TokenIds(tokenIds).ChainId(chainId).Execute()

Batch Redeem (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/w3wprediction"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	walletAddress := "0x12e32db8817e292508c34111cbc4b23340df542c" // string | User's prediction wallet address
	walletId := "5b5c1ec3be4e4416a5872b21c1ca5d20" // string | Wallet ID
	tokenIds := []string{"112233"} // []string | List of prediction token IDs to redeem. Not empty. Example: `tokenIds=112233&tokenIds=112234`
	chainId := "56" // string | Chain ID. Default `56` (BSC) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.RedeemAPI.BatchRedeem(context.Background()).WalletAddress(walletAddress).WalletId(walletId).TokenIds(tokenIds).ChainId(chainId).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `RedeemAPI.BatchRedeem``: %v\n", err)
		return
	}

	// response from `BatchRedeem`: BatchRedeemResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **walletAddress** | **string** | User&#39;s prediction wallet address | 
 **walletId** | **string** | Wallet ID | 
 **tokenIds** | **[]string** | List of prediction token IDs to redeem. Not empty. Example: &#x60;tokenIds&#x3D;112233&amp;tokenIds&#x3D;112234&#x60; | 
 **chainId** | **string** | Chain ID. Default &#x60;56&#x60; (BSC) | 

### Return type

[**BatchRedeemResponse**](BatchRedeemResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetRedeemStatus

> GetRedeemStatusResponse GetRedeemStatus(ctx).WalletAddress(walletAddress).TxHash(txHash).RecvWindow(recvWindow).Execute()

Get Redeem Status (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/w3wprediction"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	walletAddress := "0x12e32db8817e292508c34111cbc4b23340df542c" // string | User's prediction wallet address
	txHash := "0xabc123def456789abcdef123456789abcdef123456789abcdef123456789abcd" // string | Redeem transaction hash
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.RedeemAPI.GetRedeemStatus(context.Background()).WalletAddress(walletAddress).TxHash(txHash).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `RedeemAPI.GetRedeemStatus``: %v\n", err)
		return
	}

	// response from `GetRedeemStatus`: GetRedeemStatusResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **walletAddress** | **string** | User&#39;s prediction wallet address | 
 **txHash** | **string** | Redeem transaction hash | 
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**GetRedeemStatusResponse**](GetRedeemStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

