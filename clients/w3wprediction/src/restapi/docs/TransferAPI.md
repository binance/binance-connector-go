# \TransferAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**ApplyMmDeposit**](TransferAPI.md#ApplyMmDeposit) | **Post** /sapi/v1/w3w/wallet/prediction/deposit/apply | Apply MM Deposit (PREDICTION_TRADE)
[**ApplyMmWithdraw**](TransferAPI.md#ApplyMmWithdraw) | **Post** /sapi/v1/w3w/wallet/prediction/withdraw/apply | Apply MM Withdraw (PREDICTION_TRADE)
[**CreateInboundTransfer**](TransferAPI.md#CreateInboundTransfer) | **Post** /sapi/v1/w3w/wallet/prediction/transfer/inbound | Create Inbound Transfer (PREDICTION_TRADE)
[**CreateOutboundTransfer**](TransferAPI.md#CreateOutboundTransfer) | **Post** /sapi/v1/w3w/wallet/prediction/transfer/outbound | Create Outbound Transfer (PREDICTION_TRADE)
[**QueryTransferList**](TransferAPI.md#QueryTransferList) | **Get** /sapi/v1/w3w/wallet/prediction/transfer/list | Query Transfer List (PREDICTION_TRADE)
[**QueryTransferStatus**](TransferAPI.md#QueryTransferStatus) | **Get** /sapi/v1/w3w/wallet/prediction/transfer/status | Query Transfer Status (PREDICTION_TRADE)


## ApplyMmDeposit

> ApplyMmDepositResponse ApplyMmDeposit(ctx).FromToken(fromToken).FromTokenAmount(fromTokenAmount).ToToken(toToken).AccountType(accountType).ChainId(chainId).Execute()

Apply MM Deposit (PREDICTION_TRADE)


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
	fromToken := "USDT" // string | Source token symbol (e.g. `USDT`)
	fromTokenAmount := "1000000000000000000" // string | Source token amount in WEI (18 decimals). Example: `1000000000000000000` = 1 USDT
	toToken := "USDT" // string | Target token symbol (e.g. `USDT`)
	accountType := models.PlaceOrderAccountTypeParameterSpot // PlaceOrderAccountTypeParameter | Target CEX account type. Enum: `SPOT`, `FUNDING`
	chainId := "56" // string | Chain ID. Default `56` (BSC) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TransferAPI.ApplyMmDeposit(context.Background()).FromToken(fromToken).FromTokenAmount(fromTokenAmount).ToToken(toToken).AccountType(accountType).ChainId(chainId).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TransferAPI.ApplyMmDeposit``: %v\n", err)
		return
	}

	// response from `ApplyMmDeposit`: ApplyMmDepositResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **fromToken** | **string** | Source token symbol (e.g. &#x60;USDT&#x60;) | 
 **fromTokenAmount** | **string** | Source token amount in WEI (18 decimals). Example: &#x60;1000000000000000000&#x60; &#x3D; 1 USDT | 
 **toToken** | **string** | Target token symbol (e.g. &#x60;USDT&#x60;) | 
 **accountType** | [**PlaceOrderAccountTypeParameter**](PlaceOrderAccountTypeParameter.md) | Target CEX account type. Enum: &#x60;SPOT&#x60;, &#x60;FUNDING&#x60; | 
 **chainId** | **string** | Chain ID. Default &#x60;56&#x60; (BSC) | 

### Return type

[**ApplyMmDepositResponse**](ApplyMmDepositResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## ApplyMmWithdraw

> ApplyMmWithdrawResponse ApplyMmWithdraw(ctx).Coin(coin).Network(network).Amount(amount).WithdrawOrderId(withdrawOrderId).WalletType(walletType).Name(name).Execute()

Apply MM Withdraw (PREDICTION_TRADE)


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
	coin := "USDT" // string | Coin to withdraw (e.g. `USDT`)
	network := "BEP20" // string | Network (e.g. `BEP20`)
	amount := "100.00" // string | Amount to withdraw (must be > 0)
	withdrawOrderId := "wd_20260814_001" // string | Client withdraw order id (idempotency key) (optional)
	walletType := models.ApplyMmWithdrawWalletTypeParameterWalletType0 // ApplyMmWithdrawWalletTypeParameter | Source CEX account type. Enum: `0` (SPOT), `1` (FUNDING). Default `0`. Must be `0` or `1`; any other value is rejected (optional)
	name := "MM withdraw" // string | Remark for the withdraw (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TransferAPI.ApplyMmWithdraw(context.Background()).Coin(coin).Network(network).Amount(amount).WithdrawOrderId(withdrawOrderId).WalletType(walletType).Name(name).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TransferAPI.ApplyMmWithdraw``: %v\n", err)
		return
	}

	// response from `ApplyMmWithdraw`: ApplyMmWithdrawResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **coin** | **string** | Coin to withdraw (e.g. &#x60;USDT&#x60;) | 
 **network** | **string** | Network (e.g. &#x60;BEP20&#x60;) | 
 **amount** | **string** | Amount to withdraw (must be &gt; 0) | 
 **withdrawOrderId** | **string** | Client withdraw order id (idempotency key) | 
 **walletType** | [**ApplyMmWithdrawWalletTypeParameter**](ApplyMmWithdrawWalletTypeParameter.md) | Source CEX account type. Enum: &#x60;0&#x60; (SPOT), &#x60;1&#x60; (FUNDING). Default &#x60;0&#x60;. Must be &#x60;0&#x60; or &#x60;1&#x60;; any other value is rejected | 
 **name** | **string** | Remark for the withdraw | 

### Return type

[**ApplyMmWithdrawResponse**](ApplyMmWithdrawResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CreateInboundTransfer

> CreateInboundTransferResponse CreateInboundTransfer(ctx).WalletId(walletId).WalletAddress(walletAddress).FromTokenAmount(fromTokenAmount).AccountType(accountType).FromToken(fromToken).ToToken(toToken).ChainId(chainId).Execute()

Create Inbound Transfer (PREDICTION_TRADE)


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
	walletId := "5b5c1ec3be4e4416a5872b21c1ca5d20" // string | Wallet ID
	walletAddress := "0x12e32db8817e292508c34111cbc4b23340df542c" // string | User's prediction wallet address
	fromTokenAmount := "1000000000000000000" // string | Transfer amount in wei (18 decimals). Must be > 0. Example: `1000000000000000000` = 1 USDT
	accountType := models.PlaceOrderAccountTypeParameterSpot // PlaceOrderAccountTypeParameter | Destination CEX account. Enum: `SPOT`, `FUNDING`
	fromToken := "USDT" // string | Source token symbol. Default `USDT` (optional)
	toToken := "USDT" // string | Destination token symbol. Default `USDT` (optional)
	chainId := "56" // string | Chain ID. Default `56` (BSC) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TransferAPI.CreateInboundTransfer(context.Background()).WalletId(walletId).WalletAddress(walletAddress).FromTokenAmount(fromTokenAmount).AccountType(accountType).FromToken(fromToken).ToToken(toToken).ChainId(chainId).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TransferAPI.CreateInboundTransfer``: %v\n", err)
		return
	}

	// response from `CreateInboundTransfer`: CreateInboundTransferResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **walletId** | **string** | Wallet ID | 
 **walletAddress** | **string** | User&#39;s prediction wallet address | 
 **fromTokenAmount** | **string** | Transfer amount in wei (18 decimals). Must be &gt; 0. Example: &#x60;1000000000000000000&#x60; &#x3D; 1 USDT | 
 **accountType** | [**PlaceOrderAccountTypeParameter**](PlaceOrderAccountTypeParameter.md) | Destination CEX account. Enum: &#x60;SPOT&#x60;, &#x60;FUNDING&#x60; | 
 **fromToken** | **string** | Source token symbol. Default &#x60;USDT&#x60; | 
 **toToken** | **string** | Destination token symbol. Default &#x60;USDT&#x60; | 
 **chainId** | **string** | Chain ID. Default &#x60;56&#x60; (BSC) | 

### Return type

[**CreateInboundTransferResponse**](CreateInboundTransferResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CreateOutboundTransfer

> CreateOutboundTransferResponse CreateOutboundTransfer(ctx).WalletId(walletId).WalletAddress(walletAddress).FromTokenAmount(fromTokenAmount).AccountType(accountType).SourceBiz(sourceBiz).FromToken(fromToken).ToToken(toToken).ChainId(chainId).Execute()

Create Outbound Transfer (PREDICTION_TRADE)


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
	walletId := "5b5c1ec3be4e4416a5872b21c1ca5d20" // string | Wallet ID
	walletAddress := "0x12e32db8817e292508c34111cbc4b23340df542c" // string | User's prediction wallet address
	fromTokenAmount := "1000000000000000000" // string | Transfer amount in wei (18 decimals). Must be > 0. Example: `1000000000000000000` = 1 USDT
	accountType := models.PlaceOrderAccountTypeParameterSpot // PlaceOrderAccountTypeParameter | Source CEX account. Enum: `SPOT`, `FUNDING`
	sourceBiz := models.CreateOutboundTransferSourceBizParameterUserTransfer // CreateOutboundTransferSourceBizParameter | Business source. Enum: `USER_TRANSFER`, `PREDICTION_BUY`
	fromToken := "USDT" // string | Source token symbol. Default `USDT` (optional)
	toToken := "USDT" // string | Destination token symbol. Default `USDT` (optional)
	chainId := "56" // string | Chain ID. Default `56` (BSC) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TransferAPI.CreateOutboundTransfer(context.Background()).WalletId(walletId).WalletAddress(walletAddress).FromTokenAmount(fromTokenAmount).AccountType(accountType).SourceBiz(sourceBiz).FromToken(fromToken).ToToken(toToken).ChainId(chainId).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TransferAPI.CreateOutboundTransfer``: %v\n", err)
		return
	}

	// response from `CreateOutboundTransfer`: CreateOutboundTransferResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **walletId** | **string** | Wallet ID | 
 **walletAddress** | **string** | User&#39;s prediction wallet address | 
 **fromTokenAmount** | **string** | Transfer amount in wei (18 decimals). Must be &gt; 0. Example: &#x60;1000000000000000000&#x60; &#x3D; 1 USDT | 
 **accountType** | [**PlaceOrderAccountTypeParameter**](PlaceOrderAccountTypeParameter.md) | Source CEX account. Enum: &#x60;SPOT&#x60;, &#x60;FUNDING&#x60; | 
 **sourceBiz** | [**CreateOutboundTransferSourceBizParameter**](CreateOutboundTransferSourceBizParameter.md) | Business source. Enum: &#x60;USER_TRANSFER&#x60;, &#x60;PREDICTION_BUY&#x60; | 
 **fromToken** | **string** | Source token symbol. Default &#x60;USDT&#x60; | 
 **toToken** | **string** | Destination token symbol. Default &#x60;USDT&#x60; | 
 **chainId** | **string** | Chain ID. Default &#x60;56&#x60; (BSC) | 

### Return type

[**CreateOutboundTransferResponse**](CreateOutboundTransferResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryTransferList

> QueryTransferListResponse QueryTransferList(ctx).WalletAddress(walletAddress).StartDate(startDate).EndDate(endDate).TokenSymbol(tokenSymbol).Direction(direction).Offset(offset).Limit(limit).RecvWindow(recvWindow).Execute()

Query Transfer List (PREDICTION_TRADE)


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
	startDate := "2026-05-01" // string | Start date. Format: `yyyy-MM-dd`. Must be ≤ `endDate`
	endDate := "2026-05-25" // string | End date. Format: `yyyy-MM-dd`. Must be ≥ `startDate`
	tokenSymbol := "USDT" // string | Filter by token symbol (e.g. `USDT`) (optional)
	direction := models.QueryTransferListDirectionParameterInbound // QueryTransferListDirectionParameter | Filter by direction. Enum: `INBOUND`, `OUTBOUND` (optional)
	offset := int32(0) // int32 | Pagination offset. Default `0` (optional)
	limit := int32(20) // int32 | Page size. Default `20`, range 1–100 (optional)
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TransferAPI.QueryTransferList(context.Background()).WalletAddress(walletAddress).StartDate(startDate).EndDate(endDate).TokenSymbol(tokenSymbol).Direction(direction).Offset(offset).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TransferAPI.QueryTransferList``: %v\n", err)
		return
	}

	// response from `QueryTransferList`: QueryTransferListResponse
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
 **startDate** | **string** | Start date. Format: &#x60;yyyy-MM-dd&#x60;. Must be ≤ &#x60;endDate&#x60; | 
 **endDate** | **string** | End date. Format: &#x60;yyyy-MM-dd&#x60;. Must be ≥ &#x60;startDate&#x60; | 
 **tokenSymbol** | **string** | Filter by token symbol (e.g. &#x60;USDT&#x60;) | 
 **direction** | [**QueryTransferListDirectionParameter**](QueryTransferListDirectionParameter.md) | Filter by direction. Enum: &#x60;INBOUND&#x60;, &#x60;OUTBOUND&#x60; | 
 **offset** | **int32** | Pagination offset. Default &#x60;0&#x60; | 
 **limit** | **int32** | Page size. Default &#x60;20&#x60;, range 1–100 | 
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**QueryTransferListResponse**](QueryTransferListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryTransferStatus

> QueryTransferStatusResponse QueryTransferStatus(ctx).TransferId(transferId).RecvWindow(recvWindow).Execute()

Query Transfer Status (PREDICTION_TRADE)


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
	transferId := "tf_20260525_out_001" // string | Transfer ID returned from outbound/inbound transfer
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TransferAPI.QueryTransferStatus(context.Background()).TransferId(transferId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TransferAPI.QueryTransferStatus``: %v\n", err)
		return
	}

	// response from `QueryTransferStatus`: QueryTransferStatusResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **transferId** | **string** | Transfer ID returned from outbound/inbound transfer | 
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**QueryTransferStatusResponse**](QueryTransferStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

