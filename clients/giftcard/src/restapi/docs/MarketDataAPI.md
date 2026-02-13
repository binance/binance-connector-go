# \MarketDataAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**CreateADualTokenGiftCard**](MarketDataAPI.md#CreateADualTokenGiftCard) | **Post** /sapi/v1/giftcard/buyCode | Create a dual-token gift card(fixed value, discount feature)(TRADE)
[**CreateASingleTokenGiftCard**](MarketDataAPI.md#CreateASingleTokenGiftCard) | **Post** /sapi/v1/giftcard/createCode | Create a single-token gift card (USER_DATA)
[**FetchRsaPublicKey**](MarketDataAPI.md#FetchRsaPublicKey) | **Get** /sapi/v1/giftcard/cryptography/rsa-public-key | Fetch RSA Public Key(USER_DATA)
[**FetchTokenLimit**](MarketDataAPI.md#FetchTokenLimit) | **Get** /sapi/v1/giftcard/buyCode/token-limit | Fetch Token Limit(USER_DATA)
[**RedeemABinanceGiftCard**](MarketDataAPI.md#RedeemABinanceGiftCard) | **Post** /sapi/v1/giftcard/redeemCode | Redeem a Binance Gift Card(USER_DATA)
[**VerifyBinanceGiftCardByGiftCardNumber**](MarketDataAPI.md#VerifyBinanceGiftCardByGiftCardNumber) | **Get** /sapi/v1/giftcard/verify | Verify Binance Gift Card by Gift Card Number(USER_DATA)


## CreateADualTokenGiftCard

> CreateADualTokenGiftCardResponse CreateADualTokenGiftCard(ctx).BaseToken(baseToken).FaceToken(faceToken).BaseTokenAmount(baseTokenAmount).RecvWindow(recvWindow).Execute()

Create a dual-token gift card(fixed value, discount feature)(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/giftcard"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	baseToken := "baseToken_example" // string | The token you want to pay, example: BUSD
	faceToken := "faceToken_example" // string | The token you want to buy, example: BNB. If faceToken = baseToken, it's the same as createCode endpoint.
	baseTokenAmount := float32(1.0) // float32 | The base token asset quantity, example : 1.002
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceGiftCardClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.CreateADualTokenGiftCard(context.Background()).BaseToken(baseToken).FaceToken(faceToken).BaseTokenAmount(baseTokenAmount).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.CreateADualTokenGiftCard``: %v\n", err)
		return
	}

	// response from `CreateADualTokenGiftCard`: CreateADualTokenGiftCardResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **baseToken** | **string** | The token you want to pay, example: BUSD | 
 **faceToken** | **string** | The token you want to buy, example: BNB. If faceToken &#x3D; baseToken, it&#39;s the same as createCode endpoint. | 
 **baseTokenAmount** | **float32** | The base token asset quantity, example : 1.002 | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateADualTokenGiftCardResponse**](CreateADualTokenGiftCardResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CreateASingleTokenGiftCard

> CreateASingleTokenGiftCardResponse CreateASingleTokenGiftCard(ctx).Token(token).Amount(amount).RecvWindow(recvWindow).Execute()

Create a single-token gift card (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/giftcard"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	token := "token_example" // string | The token type contained in the Binance Gift Card
	amount := float32(1.0) // float32 | The amount of the token contained in the Binance Gift Card
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceGiftCardClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.CreateASingleTokenGiftCard(context.Background()).Token(token).Amount(amount).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.CreateASingleTokenGiftCard``: %v\n", err)
		return
	}

	// response from `CreateASingleTokenGiftCard`: CreateASingleTokenGiftCardResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | The token type contained in the Binance Gift Card | 
 **amount** | **float32** | The amount of the token contained in the Binance Gift Card | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateASingleTokenGiftCardResponse**](CreateASingleTokenGiftCardResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## FetchRsaPublicKey

> FetchRsaPublicKeyResponse FetchRsaPublicKey(ctx).RecvWindow(recvWindow).Execute()

Fetch RSA Public Key(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/giftcard"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceGiftCardClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.FetchRsaPublicKey(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.FetchRsaPublicKey``: %v\n", err)
		return
	}

	// response from `FetchRsaPublicKey`: FetchRsaPublicKeyResponse
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

[**FetchRsaPublicKeyResponse**](FetchRsaPublicKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## FetchTokenLimit

> FetchTokenLimitResponse FetchTokenLimit(ctx).BaseToken(baseToken).RecvWindow(recvWindow).Execute()

Fetch Token Limit(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/giftcard"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	baseToken := "baseToken_example" // string | The token you want to pay, example: BUSD
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceGiftCardClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.FetchTokenLimit(context.Background()).BaseToken(baseToken).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.FetchTokenLimit``: %v\n", err)
		return
	}

	// response from `FetchTokenLimit`: FetchTokenLimitResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **baseToken** | **string** | The token you want to pay, example: BUSD | 
 **recvWindow** | **int64** |  | 

### Return type

[**FetchTokenLimitResponse**](FetchTokenLimitResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## RedeemABinanceGiftCard

> RedeemABinanceGiftCardResponse RedeemABinanceGiftCard(ctx).Code(code).ExternalUid(externalUid).RecvWindow(recvWindow).Execute()

Redeem a Binance Gift Card(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/giftcard"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	code := "code_example" // string | Redemption code of Binance Gift Card to be redeemed, supports both Plaintext & Encrypted code.
	externalUid := "externalUid_example" // string | Each external unique ID represents a unique user on the partner platform. The function helps you to identify the redemption behavior of different users, such as redemption frequency and amount. It also helps risk and limit control of a single account, such as daily limit on redemption volume, frequency, and incorrect number of entries. This will also prevent a single user account reach the partner's daily redemption limits. We strongly recommend you to use this feature and transfer us the User ID of your users if you have different users redeeming Binance Gift Cards on your platform. To protect user data privacy, you may choose to transfer the user id in any desired format (max. 400 characters). (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceGiftCardClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.RedeemABinanceGiftCard(context.Background()).Code(code).ExternalUid(externalUid).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.RedeemABinanceGiftCard``: %v\n", err)
		return
	}

	// response from `RedeemABinanceGiftCard`: RedeemABinanceGiftCardResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** | Redemption code of Binance Gift Card to be redeemed, supports both Plaintext &amp; Encrypted code. | 
 **externalUid** | **string** | Each external unique ID represents a unique user on the partner platform. The function helps you to identify the redemption behavior of different users, such as redemption frequency and amount. It also helps risk and limit control of a single account, such as daily limit on redemption volume, frequency, and incorrect number of entries. This will also prevent a single user account reach the partner&#39;s daily redemption limits. We strongly recommend you to use this feature and transfer us the User ID of your users if you have different users redeeming Binance Gift Cards on your platform. To protect user data privacy, you may choose to transfer the user id in any desired format (max. 400 characters). | 
 **recvWindow** | **int64** |  | 

### Return type

[**RedeemABinanceGiftCardResponse**](RedeemABinanceGiftCardResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## VerifyBinanceGiftCardByGiftCardNumber

> VerifyBinanceGiftCardByGiftCardNumberResponse VerifyBinanceGiftCardByGiftCardNumber(ctx).ReferenceNo(referenceNo).RecvWindow(recvWindow).Execute()

Verify Binance Gift Card by Gift Card Number(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/giftcard"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	referenceNo := "referenceNo_example" // string | Enter the Gift Card Number
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceGiftCardClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.MarketDataAPI.VerifyBinanceGiftCardByGiftCardNumber(context.Background()).ReferenceNo(referenceNo).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketDataAPI.VerifyBinanceGiftCardByGiftCardNumber``: %v\n", err)
		return
	}

	// response from `VerifyBinanceGiftCardByGiftCardNumber`: VerifyBinanceGiftCardByGiftCardNumberResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **referenceNo** | **string** | Enter the Gift Card Number | 
 **recvWindow** | **int64** |  | 

### Return type

[**VerifyBinanceGiftCardByGiftCardNumberResponse**](VerifyBinanceGiftCardByGiftCardNumberResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

