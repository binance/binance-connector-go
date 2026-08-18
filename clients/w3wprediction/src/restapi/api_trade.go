/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package binancew3wpredictionrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/w3wprediction/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// TradeAPIService TradeAPI Service
type TradeAPIService Service

type ApiBatchCancelOrdersRequest struct {
	ctx            context.Context
	ApiService     *TradeAPIService
	walletAddress  *string
	walletId       *string
	cancelInfoList *[]models.BatchCancelOrdersCancelInfoListParameterInner
}

// User&#39;s prediction wallet address
func (r ApiBatchCancelOrdersRequest) WalletAddress(walletAddress string) ApiBatchCancelOrdersRequest {
	r.walletAddress = &walletAddress
	return r
}

// Wallet ID
func (r ApiBatchCancelOrdersRequest) WalletId(walletId string) ApiBatchCancelOrdersRequest {
	r.walletId = &walletId
	return r
}

// List of orders to cancel (index &#x60;i&#x60; starts from 0)
func (r ApiBatchCancelOrdersRequest) CancelInfoList(cancelInfoList []models.BatchCancelOrdersCancelInfoListParameterInner) ApiBatchCancelOrdersRequest {
	r.cancelInfoList = &cancelInfoList
	return r
}

func (r ApiBatchCancelOrdersRequest) Execute() (*common.RestApiResponse[models.BatchCancelOrdersResponse], error) {
	return r.ApiService.BatchCancelOrdersExecute(r)
}

/*
BatchCancelOrders Batch Cancel Orders (PREDICTION_TRADE)
Post /sapi/v1/w3w/wallet/prediction/trade/batch-cancel

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/trade#batch-cancel-orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param walletAddress -  User's prediction wallet address
@param walletId -  Wallet ID
@param cancelInfoList -  List of orders to cancel (index `i` starts from 0)
@return ApiBatchCancelOrdersRequest
*/
func (a *TradeAPIService) BatchCancelOrders(ctx context.Context) ApiBatchCancelOrdersRequest {
	return ApiBatchCancelOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BatchCancelOrdersResponse
func (a *TradeAPIService) BatchCancelOrdersExecute(r ApiBatchCancelOrdersRequest) (*common.RestApiResponse[models.BatchCancelOrdersResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/trade/batch-cancel"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.walletAddress == nil {
		return nil, common.ReportError("walletAddress is required and must be specified")
	}

	if r.walletId == nil {
		return nil, common.ReportError("walletId is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "walletAddress", r.walletAddress, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "walletId", r.walletId, "form", "")
	if r.cancelInfoList != nil {
		t := *r.cancelInfoList
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "cancelInfoList", t, "form", "multi")
	}

	resp, err := SendRequest[models.BatchCancelOrdersResponse](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		true,
	)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetQuoteRequest struct {
	ctx                context.Context
	ApiService         *TradeAPIService
	walletAddress      *string
	tokenId            *string
	side               *models.GetQuoteSideParameter
	amountIn           *string
	orderType          *models.GetQuoteOrderTypeParameter
	slippageBps        *int32
	priceLimit         *string
	chainId            *string
	feeRateBps         *int32
	fundingSource      *models.GetQuoteFundingSourceParameter
	fundTransferAmount *string
}

// User&#39;s prediction wallet address
func (r ApiGetQuoteRequest) WalletAddress(walletAddress string) ApiGetQuoteRequest {
	r.walletAddress = &walletAddress
	return r
}

// Prediction outcome token ID
func (r ApiGetQuoteRequest) TokenId(tokenId string) ApiGetQuoteRequest {
	r.tokenId = &tokenId
	return r
}

// Trade direction. Enum: &#x60;BUY&#x60;, &#x60;SELL&#x60;
func (r ApiGetQuoteRequest) Side(side models.GetQuoteSideParameter) ApiGetQuoteRequest {
	r.side = &side
	return r
}

// Input amount in wei (18 decimals). Must be &gt; 0. For &#x60;MARKET&#x60; orders, minimum is approximately 1.5 USDT (varies by market depth). Example: &#x60;1000000000000000000&#x60; &#x3D; 1 USDT
func (r ApiGetQuoteRequest) AmountIn(amountIn string) ApiGetQuoteRequest {
	r.amountIn = &amountIn
	return r
}

// Order type. Enum: &#x60;MARKET&#x60;, &#x60;LIMIT&#x60;
func (r ApiGetQuoteRequest) OrderType(orderType models.GetQuoteOrderTypeParameter) ApiGetQuoteRequest {
	r.orderType = &orderType
	return r
}

// Slippage tolerance in basis points. Range 1–10000
func (r ApiGetQuoteRequest) SlippageBps(slippageBps int32) ApiGetQuoteRequest {
	r.slippageBps = &slippageBps
	return r
}

// Limit price. Required when &#x60;orderType&#x3D;LIMIT&#x60;. Must be &gt; 0
func (r ApiGetQuoteRequest) PriceLimit(priceLimit string) ApiGetQuoteRequest {
	r.priceLimit = &priceLimit
	return r
}

// Chain ID. Default &#x60;56&#x60; (BSC)
func (r ApiGetQuoteRequest) ChainId(chainId string) ApiGetQuoteRequest {
	r.chainId = &chainId
	return r
}

// Fee rate in basis points. Default &#x60;200&#x60;, range 1–10000
func (r ApiGetQuoteRequest) FeeRateBps(feeRateBps int32) ApiGetQuoteRequest {
	r.feeRateBps = &feeRateBps
	return r
}

// Funding source. Enum: &#x60;MPC&#x60;, &#x60;CEX&#x60;. Default &#x60;MPC&#x60;
func (r ApiGetQuoteRequest) FundingSource(fundingSource models.GetQuoteFundingSourceParameter) ApiGetQuoteRequest {
	r.fundingSource = &fundingSource
	return r
}

// Auto-transfer amount before order (wei). Must be &gt; 0 if provided
func (r ApiGetQuoteRequest) FundTransferAmount(fundTransferAmount string) ApiGetQuoteRequest {
	r.fundTransferAmount = &fundTransferAmount
	return r
}

func (r ApiGetQuoteRequest) Execute() (*common.RestApiResponse[models.GetQuoteResponse], error) {
	return r.ApiService.GetQuoteExecute(r)
}

/*
GetQuote Get Quote (PREDICTION_TRADE)
Post /sapi/v1/w3w/wallet/prediction/trade/get-quote

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/trade#get-quote

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param walletAddress -  User's prediction wallet address
@param tokenId -  Prediction outcome token ID
@param side -  Trade direction. Enum: `BUY`, `SELL`
@param amountIn -  Input amount in wei (18 decimals). Must be > 0. For `MARKET` orders, minimum is approximately 1.5 USDT (varies by market depth). Example: `1000000000000000000` = 1 USDT
@param orderType -  Order type. Enum: `MARKET`, `LIMIT`
@param slippageBps -  Slippage tolerance in basis points. Range 1–10000
@param priceLimit -  Limit price. Required when `orderType=LIMIT`. Must be > 0
@param chainId -  Chain ID. Default `56` (BSC)
@param feeRateBps -  Fee rate in basis points. Default `200`, range 1–10000
@param fundingSource -  Funding source. Enum: `MPC`, `CEX`. Default `MPC`
@param fundTransferAmount -  Auto-transfer amount before order (wei). Must be > 0 if provided
@return ApiGetQuoteRequest
*/
func (a *TradeAPIService) GetQuote(ctx context.Context) ApiGetQuoteRequest {
	return ApiGetQuoteRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetQuoteResponse
func (a *TradeAPIService) GetQuoteExecute(r ApiGetQuoteRequest) (*common.RestApiResponse[models.GetQuoteResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/trade/get-quote"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.walletAddress == nil {
		return nil, common.ReportError("walletAddress is required and must be specified")
	}

	if r.tokenId == nil {
		return nil, common.ReportError("tokenId is required and must be specified")
	}

	if r.side == nil {
		return nil, common.ReportError("side is required and must be specified")
	}

	if r.amountIn == nil {
		return nil, common.ReportError("amountIn is required and must be specified")
	}

	if r.orderType == nil {
		return nil, common.ReportError("orderType is required and must be specified")
	}

	if r.slippageBps == nil {
		return nil, common.ReportError("slippageBps is required and must be specified")
	}
	if *r.slippageBps < 1 {
		return nil, common.ReportError("slippageBps must be greater than 1")
	}
	if *r.slippageBps > 10000 {
		return nil, common.ReportError("slippageBps must be less than 10000")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "walletAddress", r.walletAddress, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "tokenId", r.tokenId, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "side", r.side, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amountIn", r.amountIn, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderType", r.orderType, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "slippageBps", r.slippageBps, "form", "")
	if r.priceLimit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "priceLimit", r.priceLimit, "form", "")
	}
	if r.chainId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "chainId", r.chainId, "form", "")
	}
	if r.feeRateBps != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "feeRateBps", r.feeRateBps, "form", "")
	}
	if r.fundingSource != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fundingSource", r.fundingSource, "form", "")
	}
	if r.fundTransferAmount != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fundTransferAmount", r.fundTransferAmount, "form", "")
	}

	resp, err := SendRequest[models.GetQuoteResponse](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		true,
	)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiPlaceOrderRequest struct {
	ctx                context.Context
	ApiService         *TradeAPIService
	walletAddress      *string
	walletId           *string
	quoteId            *string
	timeInForce        *string
	accountType        *models.PlaceOrderAccountTypeParameter
	orderType          *models.GetQuoteOrderTypeParameter
	slippageBps        *int32
	priceLimit         *string
	fundingSource      *models.GetQuoteFundingSourceParameter
	fundTransferAmount *string
}

// User&#39;s prediction wallet address
func (r ApiPlaceOrderRequest) WalletAddress(walletAddress string) ApiPlaceOrderRequest {
	r.walletAddress = &walletAddress
	return r
}

// Wallet ID
func (r ApiPlaceOrderRequest) WalletId(walletId string) ApiPlaceOrderRequest {
	r.walletId = &walletId
	return r
}

// Quote ID obtained from &#x60;Get Quote&#x60;
func (r ApiPlaceOrderRequest) QuoteId(quoteId string) ApiPlaceOrderRequest {
	r.quoteId = &quoteId
	return r
}

// Must match &#x60;orderType&#x60;: &#x60;FOK&#x60; for &#x60;MARKET&#x60;, &#x60;GTC&#x60; for &#x60;LIMIT&#x60;
func (r ApiPlaceOrderRequest) TimeInForce(timeInForce string) ApiPlaceOrderRequest {
	r.timeInForce = &timeInForce
	return r
}

// Payment account type. Enum: &#x60;SPOT&#x60;, &#x60;FUNDING&#x60;
func (r ApiPlaceOrderRequest) AccountType(accountType models.PlaceOrderAccountTypeParameter) ApiPlaceOrderRequest {
	r.accountType = &accountType
	return r
}

// Order type. Enum: &#x60;MARKET&#x60;, &#x60;LIMIT&#x60;
func (r ApiPlaceOrderRequest) OrderType(orderType models.GetQuoteOrderTypeParameter) ApiPlaceOrderRequest {
	r.orderType = &orderType
	return r
}

// Slippage tolerance in basis points. Range 1–10000
func (r ApiPlaceOrderRequest) SlippageBps(slippageBps int32) ApiPlaceOrderRequest {
	r.slippageBps = &slippageBps
	return r
}

// Limit price. Required when &#x60;orderType&#x3D;LIMIT&#x60;. Must be &gt; 0
func (r ApiPlaceOrderRequest) PriceLimit(priceLimit string) ApiPlaceOrderRequest {
	r.priceLimit = &priceLimit
	return r
}

// Funding source. Enum: &#x60;MPC&#x60;, &#x60;CEX&#x60;. Default &#x60;MPC&#x60;
func (r ApiPlaceOrderRequest) FundingSource(fundingSource models.GetQuoteFundingSourceParameter) ApiPlaceOrderRequest {
	r.fundingSource = &fundingSource
	return r
}

// Auto-transfer amount before order (wei). Must be &gt; 0 if provided
func (r ApiPlaceOrderRequest) FundTransferAmount(fundTransferAmount string) ApiPlaceOrderRequest {
	r.fundTransferAmount = &fundTransferAmount
	return r
}

func (r ApiPlaceOrderRequest) Execute() (*common.RestApiResponse[models.PlaceOrderResponse], error) {
	return r.ApiService.PlaceOrderExecute(r)
}

/*
PlaceOrder Place Order (PREDICTION_TRADE)
Post /sapi/v1/w3w/wallet/prediction/trade/place-order-bundle

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/trade#place-order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param walletAddress -  User's prediction wallet address
@param walletId -  Wallet ID
@param quoteId -  Quote ID obtained from `Get Quote`
@param timeInForce -  Must match `orderType`: `FOK` for `MARKET`, `GTC` for `LIMIT`
@param accountType -  Payment account type. Enum: `SPOT`, `FUNDING`
@param orderType -  Order type. Enum: `MARKET`, `LIMIT`
@param slippageBps -  Slippage tolerance in basis points. Range 1–10000
@param priceLimit -  Limit price. Required when `orderType=LIMIT`. Must be > 0
@param fundingSource -  Funding source. Enum: `MPC`, `CEX`. Default `MPC`
@param fundTransferAmount -  Auto-transfer amount before order (wei). Must be > 0 if provided
@return ApiPlaceOrderRequest
*/
func (a *TradeAPIService) PlaceOrder(ctx context.Context) ApiPlaceOrderRequest {
	return ApiPlaceOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PlaceOrderResponse
func (a *TradeAPIService) PlaceOrderExecute(r ApiPlaceOrderRequest) (*common.RestApiResponse[models.PlaceOrderResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/trade/place-order-bundle"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.walletAddress == nil {
		return nil, common.ReportError("walletAddress is required and must be specified")
	}

	if r.walletId == nil {
		return nil, common.ReportError("walletId is required and must be specified")
	}

	if r.quoteId == nil {
		return nil, common.ReportError("quoteId is required and must be specified")
	}

	if r.timeInForce == nil {
		return nil, common.ReportError("timeInForce is required and must be specified")
	}

	if r.accountType == nil {
		return nil, common.ReportError("accountType is required and must be specified")
	}

	if r.orderType == nil {
		return nil, common.ReportError("orderType is required and must be specified")
	}

	if r.slippageBps == nil {
		return nil, common.ReportError("slippageBps is required and must be specified")
	}
	if *r.slippageBps < 1 {
		return nil, common.ReportError("slippageBps must be greater than 1")
	}
	if *r.slippageBps > 10000 {
		return nil, common.ReportError("slippageBps must be less than 10000")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "walletAddress", r.walletAddress, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "walletId", r.walletId, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quoteId", r.quoteId, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "timeInForce", r.timeInForce, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "accountType", r.accountType, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderType", r.orderType, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "slippageBps", r.slippageBps, "form", "")
	if r.priceLimit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "priceLimit", r.priceLimit, "form", "")
	}
	if r.fundingSource != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fundingSource", r.fundingSource, "form", "")
	}
	if r.fundTransferAmount != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fundTransferAmount", r.fundTransferAmount, "form", "")
	}

	resp, err := SendRequest[models.PlaceOrderResponse](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		true,
	)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryActiveOrdersRequest struct {
	ctx           context.Context
	ApiService    *TradeAPIService
	walletAddress *string
	tradeSide     *models.GetQuoteSideParameter
	l1Category    *string
	marketId      *int64
	offset        *int32
	limit         *int32
	recvWindow    *int64
}

// User&#39;s prediction wallet address
func (r ApiQueryActiveOrdersRequest) WalletAddress(walletAddress string) ApiQueryActiveOrdersRequest {
	r.walletAddress = &walletAddress
	return r
}

// Filter by trade side. Enum: &#x60;BUY&#x60;, &#x60;SELL&#x60;
func (r ApiQueryActiveOrdersRequest) TradeSide(tradeSide models.GetQuoteSideParameter) ApiQueryActiveOrdersRequest {
	r.tradeSide = &tradeSide
	return r
}

// Filter by level-1 category
func (r ApiQueryActiveOrdersRequest) L1Category(l1Category string) ApiQueryActiveOrdersRequest {
	r.l1Category = &l1Category
	return r
}

// Filter by market ID
func (r ApiQueryActiveOrdersRequest) MarketId(marketId int64) ApiQueryActiveOrdersRequest {
	r.marketId = &marketId
	return r
}

// Pagination offset. Default &#x60;0&#x60;
func (r ApiQueryActiveOrdersRequest) Offset(offset int32) ApiQueryActiveOrdersRequest {
	r.offset = &offset
	return r
}

// Page size. Default &#x60;20&#x60;, range 1–100
func (r ApiQueryActiveOrdersRequest) Limit(limit int32) ApiQueryActiveOrdersRequest {
	r.limit = &limit
	return r
}

// Request validity window in milliseconds
func (r ApiQueryActiveOrdersRequest) RecvWindow(recvWindow int64) ApiQueryActiveOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryActiveOrdersRequest) Execute() (*common.RestApiResponse[models.QueryActiveOrdersResponse], error) {
	return r.ApiService.QueryActiveOrdersExecute(r)
}

/*
QueryActiveOrders Query Active Orders (PREDICTION_TRADE)
Get /sapi/v1/w3w/wallet/prediction/order/list

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/trade#query-active-orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param walletAddress -  User's prediction wallet address
@param tradeSide -  Filter by trade side. Enum: `BUY`, `SELL`
@param l1Category -  Filter by level-1 category
@param marketId -  Filter by market ID
@param offset -  Pagination offset. Default `0`
@param limit -  Page size. Default `20`, range 1–100
@param recvWindow -  Request validity window in milliseconds
@return ApiQueryActiveOrdersRequest
*/
func (a *TradeAPIService) QueryActiveOrders(ctx context.Context) ApiQueryActiveOrdersRequest {
	return ApiQueryActiveOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryActiveOrdersResponse
func (a *TradeAPIService) QueryActiveOrdersExecute(r ApiQueryActiveOrdersRequest) (*common.RestApiResponse[models.QueryActiveOrdersResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/order/list"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.walletAddress == nil {
		return nil, common.ReportError("walletAddress is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "walletAddress", r.walletAddress, "form", "")
	if r.tradeSide != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "tradeSide", r.tradeSide, "form", "")
	}
	if r.l1Category != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "l1Category", r.l1Category, "form", "")
	}
	if r.marketId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "marketId", r.marketId, "form", "")
	}
	if r.offset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryActiveOrdersResponse](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		true,
	)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryOrderHistoryRequest struct {
	ctx           context.Context
	ApiService    *TradeAPIService
	walletAddress *string
	l1Category    *string
	orderType     *models.GetQuoteOrderTypeParameter
	status        *string
	startDate     *string
	endDate       *string
	offset        *int32
	limit         *int32
	recvWindow    *int64
}

// User&#39;s prediction wallet address
func (r ApiQueryOrderHistoryRequest) WalletAddress(walletAddress string) ApiQueryOrderHistoryRequest {
	r.walletAddress = &walletAddress
	return r
}

// Filter by level-1 category
func (r ApiQueryOrderHistoryRequest) L1Category(l1Category string) ApiQueryOrderHistoryRequest {
	r.l1Category = &l1Category
	return r
}

// Filter by order type. Enum: &#x60;MARKET&#x60;, &#x60;LIMIT&#x60;
func (r ApiQueryOrderHistoryRequest) OrderType(orderType models.GetQuoteOrderTypeParameter) ApiQueryOrderHistoryRequest {
	r.orderType = &orderType
	return r
}

// Filter by order status
func (r ApiQueryOrderHistoryRequest) Status(status string) ApiQueryOrderHistoryRequest {
	r.status = &status
	return r
}

// Start date. Format: &#x60;yyyy-MM-dd&#x60;. Must be ≤ &#x60;endDate&#x60;
func (r ApiQueryOrderHistoryRequest) StartDate(startDate string) ApiQueryOrderHistoryRequest {
	r.startDate = &startDate
	return r
}

// End date. Format: &#x60;yyyy-MM-dd&#x60;. Must be ≥ &#x60;startDate&#x60;
func (r ApiQueryOrderHistoryRequest) EndDate(endDate string) ApiQueryOrderHistoryRequest {
	r.endDate = &endDate
	return r
}

// Pagination offset. Default &#x60;0&#x60;
func (r ApiQueryOrderHistoryRequest) Offset(offset int32) ApiQueryOrderHistoryRequest {
	r.offset = &offset
	return r
}

// Page size. Default &#x60;20&#x60;, range 1–100
func (r ApiQueryOrderHistoryRequest) Limit(limit int32) ApiQueryOrderHistoryRequest {
	r.limit = &limit
	return r
}

// Request validity window in milliseconds
func (r ApiQueryOrderHistoryRequest) RecvWindow(recvWindow int64) ApiQueryOrderHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryOrderHistoryRequest) Execute() (*common.RestApiResponse[models.QueryOrderHistoryResponse], error) {
	return r.ApiService.QueryOrderHistoryExecute(r)
}

/*
QueryOrderHistory Query Order History (PREDICTION_TRADE)
Get /sapi/v1/w3w/wallet/prediction/order/history

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/trade#query-order-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param walletAddress -  User's prediction wallet address
@param l1Category -  Filter by level-1 category
@param orderType -  Filter by order type. Enum: `MARKET`, `LIMIT`
@param status -  Filter by order status
@param startDate -  Start date. Format: `yyyy-MM-dd`. Must be ≤ `endDate`
@param endDate -  End date. Format: `yyyy-MM-dd`. Must be ≥ `startDate`
@param offset -  Pagination offset. Default `0`
@param limit -  Page size. Default `20`, range 1–100
@param recvWindow -  Request validity window in milliseconds
@return ApiQueryOrderHistoryRequest
*/
func (a *TradeAPIService) QueryOrderHistory(ctx context.Context) ApiQueryOrderHistoryRequest {
	return ApiQueryOrderHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryOrderHistoryResponse
func (a *TradeAPIService) QueryOrderHistoryExecute(r ApiQueryOrderHistoryRequest) (*common.RestApiResponse[models.QueryOrderHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/order/history"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.walletAddress == nil {
		return nil, common.ReportError("walletAddress is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "walletAddress", r.walletAddress, "form", "")
	if r.l1Category != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "l1Category", r.l1Category, "form", "")
	}
	if r.orderType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderType", r.orderType, "form", "")
	}
	if r.status != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	if r.startDate != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startDate", r.startDate, "form", "")
	}
	if r.endDate != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endDate", r.endDate, "form", "")
	}
	if r.offset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryOrderHistoryResponse](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		true,
	)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
