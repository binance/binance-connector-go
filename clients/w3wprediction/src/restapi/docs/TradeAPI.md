# \TradeAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**BatchCancelOrders**](TradeAPI.md#BatchCancelOrders) | **Post** /sapi/v1/w3w/wallet/prediction/trade/batch-cancel | Batch Cancel Orders (TRADE)
[**GetQuote**](TradeAPI.md#GetQuote) | **Post** /sapi/v1/w3w/wallet/prediction/trade/get-quote | Get Quote (TRADE)
[**PlaceOrder**](TradeAPI.md#PlaceOrder) | **Post** /sapi/v1/w3w/wallet/prediction/trade/place-order-bundle | Place Order (TRADE)
[**QueryActiveOrders**](TradeAPI.md#QueryActiveOrders) | **Get** /sapi/v1/w3w/wallet/prediction/order/list | Query Active Orders (USER_DATA)
[**QueryOrderHistory**](TradeAPI.md#QueryOrderHistory) | **Get** /sapi/v1/w3w/wallet/prediction/order/history | Query Order History (USER_DATA)


## BatchCancelOrders

> BatchCancelOrdersResponse BatchCancelOrders(ctx).WalletAddress(walletAddress).WalletId(walletId).CancelInfoList(cancelInfoList).Execute()

Batch Cancel Orders (TRADE)


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
	cancelInfoList := []models.BatchCancelOrdersCancelInfoListParameterInner{*models.NewBatchCancelOrdersCancelInfoListParameterInner("54124")} // []BatchCancelOrdersCancelInfoListParameterInner | List of orders to cancel (index `i` starts from 0) (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.BatchCancelOrders(context.Background()).WalletAddress(walletAddress).WalletId(walletId).CancelInfoList(cancelInfoList).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.BatchCancelOrders``: %v\n", err)
		return
	}

	// response from `BatchCancelOrders`: BatchCancelOrdersResponse
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
 **cancelInfoList** | [**[]BatchCancelOrdersCancelInfoListParameterInner**](BatchCancelOrdersCancelInfoListParameterInner.md) | List of orders to cancel (index &#x60;i&#x60; starts from 0) | 

### Return type

[**BatchCancelOrdersResponse**](BatchCancelOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetQuote

> GetQuoteResponse GetQuote(ctx).WalletAddress(walletAddress).TokenId(tokenId).Side(side).AmountIn(amountIn).OrderType(orderType).SlippageBps(slippageBps).PriceLimit(priceLimit).ChainId(chainId).FeeRateBps(feeRateBps).FundingSource(fundingSource).FundTransferAmount(fundTransferAmount).Execute()

Get Quote (TRADE)


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
	tokenId := "112233" // string | Prediction outcome token ID
	side := models.GetQuoteSideParameterBuy // GetQuoteSideParameter | Trade direction. Enum: `BUY`, `SELL`
	amountIn := "1000000000000000000" // string | Input amount in wei (18 decimals). Must be > 0. For `MARKET` orders, minimum is approximately 1.5 USDT (varies by market depth). Example: `1000000000000000000` = 1 USDT
	orderType := models.GetQuoteOrderTypeParameterMarket // GetQuoteOrderTypeParameter | Order type. Enum: `MARKET`, `LIMIT`
	slippageBps := int32(1200) // int32 | Slippage tolerance in basis points. Range 1–10000
	priceLimit := "0.5" // string | Limit price. Required when `orderType=LIMIT`. Must be > 0 (optional)
	chainId := "56" // string | Chain ID. Default `56` (BSC) (optional)
	feeRateBps := int32(200) // int32 | Fee rate in basis points. Default `200`, range 1–10000 (optional)
	fundingSource := models.GetQuoteFundingSourceParameterMpc // GetQuoteFundingSourceParameter | Funding source. Enum: `MPC`, `CEX`. Default `MPC` (optional)
	fundTransferAmount := "1000000000000000000" // string | Auto-transfer amount before order (wei). Must be > 0 if provided (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.GetQuote(context.Background()).WalletAddress(walletAddress).TokenId(tokenId).Side(side).AmountIn(amountIn).OrderType(orderType).SlippageBps(slippageBps).PriceLimit(priceLimit).ChainId(chainId).FeeRateBps(feeRateBps).FundingSource(fundingSource).FundTransferAmount(fundTransferAmount).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.GetQuote``: %v\n", err)
		return
	}

	// response from `GetQuote`: GetQuoteResponse
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
 **tokenId** | **string** | Prediction outcome token ID | 
 **side** | [**GetQuoteSideParameter**](GetQuoteSideParameter.md) | Trade direction. Enum: &#x60;BUY&#x60;, &#x60;SELL&#x60; | 
 **amountIn** | **string** | Input amount in wei (18 decimals). Must be &gt; 0. For &#x60;MARKET&#x60; orders, minimum is approximately 1.5 USDT (varies by market depth). Example: &#x60;1000000000000000000&#x60; &#x3D; 1 USDT | 
 **orderType** | [**GetQuoteOrderTypeParameter**](GetQuoteOrderTypeParameter.md) | Order type. Enum: &#x60;MARKET&#x60;, &#x60;LIMIT&#x60; | 
 **slippageBps** | **int32** | Slippage tolerance in basis points. Range 1–10000 | 
 **priceLimit** | **string** | Limit price. Required when &#x60;orderType&#x3D;LIMIT&#x60;. Must be &gt; 0 | 
 **chainId** | **string** | Chain ID. Default &#x60;56&#x60; (BSC) | 
 **feeRateBps** | **int32** | Fee rate in basis points. Default &#x60;200&#x60;, range 1–10000 | 
 **fundingSource** | [**GetQuoteFundingSourceParameter**](GetQuoteFundingSourceParameter.md) | Funding source. Enum: &#x60;MPC&#x60;, &#x60;CEX&#x60;. Default &#x60;MPC&#x60; | 
 **fundTransferAmount** | **string** | Auto-transfer amount before order (wei). Must be &gt; 0 if provided | 

### Return type

[**GetQuoteResponse**](GetQuoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## PlaceOrder

> PlaceOrderResponse PlaceOrder(ctx).WalletAddress(walletAddress).WalletId(walletId).QuoteId(quoteId).TimeInForce(timeInForce).AccountType(accountType).OrderType(orderType).SlippageBps(slippageBps).PriceLimit(priceLimit).FundingSource(fundingSource).FundTransferAmount(fundTransferAmount).Execute()

Place Order (TRADE)


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
	quoteId := "q_20260525_abc123xyz" // string | Quote ID obtained from `Get Quote`
	timeInForce := "FOK" // string | Must match `orderType`: `FOK` for `MARKET`, `GTC` for `LIMIT`
	accountType := models.PlaceOrderAccountTypeParameterSpot // PlaceOrderAccountTypeParameter | Payment account type. Enum: `SPOT`, `FUNDING`
	orderType := models.GetQuoteOrderTypeParameterMarket // GetQuoteOrderTypeParameter | Order type. Enum: `MARKET`, `LIMIT`
	slippageBps := int32(1200) // int32 | Slippage tolerance in basis points. Range 1–10000
	priceLimit := "0.5" // string | Limit price. Required when `orderType=LIMIT`. Must be > 0 (optional)
	fundingSource := models.GetQuoteFundingSourceParameterMpc // GetQuoteFundingSourceParameter | Funding source. Enum: `MPC`, `CEX`. Default `MPC` (optional)
	fundTransferAmount := "1000000000000000000" // string | Auto-transfer amount before order (wei). Must be > 0 if provided (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.PlaceOrder(context.Background()).WalletAddress(walletAddress).WalletId(walletId).QuoteId(quoteId).TimeInForce(timeInForce).AccountType(accountType).OrderType(orderType).SlippageBps(slippageBps).PriceLimit(priceLimit).FundingSource(fundingSource).FundTransferAmount(fundTransferAmount).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.PlaceOrder``: %v\n", err)
		return
	}

	// response from `PlaceOrder`: PlaceOrderResponse
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
 **quoteId** | **string** | Quote ID obtained from &#x60;Get Quote&#x60; | 
 **timeInForce** | **string** | Must match &#x60;orderType&#x60;: &#x60;FOK&#x60; for &#x60;MARKET&#x60;, &#x60;GTC&#x60; for &#x60;LIMIT&#x60; | 
 **accountType** | [**PlaceOrderAccountTypeParameter**](PlaceOrderAccountTypeParameter.md) | Payment account type. Enum: &#x60;SPOT&#x60;, &#x60;FUNDING&#x60; | 
 **orderType** | [**GetQuoteOrderTypeParameter**](GetQuoteOrderTypeParameter.md) | Order type. Enum: &#x60;MARKET&#x60;, &#x60;LIMIT&#x60; | 
 **slippageBps** | **int32** | Slippage tolerance in basis points. Range 1–10000 | 
 **priceLimit** | **string** | Limit price. Required when &#x60;orderType&#x3D;LIMIT&#x60;. Must be &gt; 0 | 
 **fundingSource** | [**GetQuoteFundingSourceParameter**](GetQuoteFundingSourceParameter.md) | Funding source. Enum: &#x60;MPC&#x60;, &#x60;CEX&#x60;. Default &#x60;MPC&#x60; | 
 **fundTransferAmount** | **string** | Auto-transfer amount before order (wei). Must be &gt; 0 if provided | 

### Return type

[**PlaceOrderResponse**](PlaceOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryActiveOrders

> QueryActiveOrdersResponse QueryActiveOrders(ctx).WalletAddress(walletAddress).TradeSide(tradeSide).L1Category(l1Category).MarketId(marketId).Offset(offset).Limit(limit).RecvWindow(recvWindow).Execute()

Query Active Orders (USER_DATA)


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
	tradeSide := models.GetQuoteSideParameterBuy // GetQuoteSideParameter | Filter by trade side. Enum: `BUY`, `SELL` (optional)
	l1Category := "crypto" // string | Filter by level-1 category (optional)
	marketId := int64(5567895) // int64 | Filter by market ID (optional)
	offset := int32(0) // int32 | Pagination offset. Default `0` (optional)
	limit := int32(20) // int32 | Page size. Default `20`, range 1–100 (optional)
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryActiveOrders(context.Background()).WalletAddress(walletAddress).TradeSide(tradeSide).L1Category(l1Category).MarketId(marketId).Offset(offset).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryActiveOrders``: %v\n", err)
		return
	}

	// response from `QueryActiveOrders`: QueryActiveOrdersResponse
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
 **tradeSide** | [**GetQuoteSideParameter**](GetQuoteSideParameter.md) | Filter by trade side. Enum: &#x60;BUY&#x60;, &#x60;SELL&#x60; | 
 **l1Category** | **string** | Filter by level-1 category | 
 **marketId** | **int64** | Filter by market ID | 
 **offset** | **int32** | Pagination offset. Default &#x60;0&#x60; | 
 **limit** | **int32** | Page size. Default &#x60;20&#x60;, range 1–100 | 
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**QueryActiveOrdersResponse**](QueryActiveOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryOrderHistory

> QueryOrderHistoryResponse QueryOrderHistory(ctx).WalletAddress(walletAddress).L1Category(l1Category).OrderType(orderType).Status(status).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).RecvWindow(recvWindow).Execute()

Query Order History (USER_DATA)


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
	l1Category := "crypto" // string | Filter by level-1 category (optional)
	orderType := models.GetQuoteOrderTypeParameterMarket // GetQuoteOrderTypeParameter | Filter by order type. Enum: `MARKET`, `LIMIT` (optional)
	status := "CLOSED" // string | Filter by order status (optional)
	startDate := "2026-05-01" // string | Start date. Format: `yyyy-MM-dd`. Must be ≤ `endDate` (optional)
	endDate := "2026-05-25" // string | End date. Format: `yyyy-MM-dd`. Must be ≥ `startDate` (optional)
	offset := int32(0) // int32 | Pagination offset. Default `0` (optional)
	limit := int32(20) // int32 | Page size. Default `20`, range 1–100 (optional)
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceW3wPredictionClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryOrderHistory(context.Background()).WalletAddress(walletAddress).L1Category(l1Category).OrderType(orderType).Status(status).StartDate(startDate).EndDate(endDate).Offset(offset).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryOrderHistory``: %v\n", err)
		return
	}

	// response from `QueryOrderHistory`: QueryOrderHistoryResponse
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
 **l1Category** | **string** | Filter by level-1 category | 
 **orderType** | [**GetQuoteOrderTypeParameter**](GetQuoteOrderTypeParameter.md) | Filter by order type. Enum: &#x60;MARKET&#x60;, &#x60;LIMIT&#x60; | 
 **status** | **string** | Filter by order status | 
 **startDate** | **string** | Start date. Format: &#x60;yyyy-MM-dd&#x60;. Must be ≤ &#x60;endDate&#x60; | 
 **endDate** | **string** | End date. Format: &#x60;yyyy-MM-dd&#x60;. Must be ≥ &#x60;startDate&#x60; | 
 **offset** | **int32** | Pagination offset. Default &#x60;0&#x60; | 
 **limit** | **int32** | Page size. Default &#x60;20&#x60;, range 1–100 | 
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**QueryOrderHistoryResponse**](QueryOrderHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

