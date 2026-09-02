/*
Convert REST API

Request quotes and execute cryptocurrency conversions via the Convert REST API.
*/

package binanceconvertrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/convert/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// TradeAPIService TradeAPI Service
type TradeAPIService Service

type ApiAcceptQuoteRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	quoteId    *string
	recvWindow *int64
}

func (r ApiAcceptQuoteRequest) QuoteId(quoteId string) ApiAcceptQuoteRequest {
	r.quoteId = &quoteId
	return r
}

// Request validity window in milliseconds
func (r ApiAcceptQuoteRequest) RecvWindow(recvWindow int64) ApiAcceptQuoteRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiAcceptQuoteRequest) Execute() (*common.RestApiResponse[models.AcceptQuoteResponse], error) {
	return r.ApiService.AcceptQuoteExecute(r)
}

/*
AcceptQuote Accept Quote (TRADE)
Post /sapi/v1/convert/acceptQuote

https://developers.binance.com/en/docs/catalog/core-trading-convert/api/rest-api/trade#accept-quote

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param quoteId -
@param recvWindow -  Request validity window in milliseconds
@return ApiAcceptQuoteRequest
*/
func (a *TradeAPIService) AcceptQuote(ctx context.Context) ApiAcceptQuoteRequest {
	return ApiAcceptQuoteRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AcceptQuoteResponse
func (a *TradeAPIService) AcceptQuoteExecute(r ApiAcceptQuoteRequest) (*common.RestApiResponse[models.AcceptQuoteResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/convert/acceptQuote"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.quoteId == nil {
		return nil, common.ReportError("quoteId is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quoteId", r.quoteId, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.AcceptQuoteResponse](
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

type ApiCancelLimitOrderRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	orderId    *int64
	recvWindow *int64
}

// The orderId from &#x60;placeOrder&#x60; api
func (r ApiCancelLimitOrderRequest) OrderId(orderId int64) ApiCancelLimitOrderRequest {
	r.orderId = &orderId
	return r
}

// Request validity window in milliseconds
func (r ApiCancelLimitOrderRequest) RecvWindow(recvWindow int64) ApiCancelLimitOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelLimitOrderRequest) Execute() (*common.RestApiResponse[models.CancelLimitOrderResponse], error) {
	return r.ApiService.CancelLimitOrderExecute(r)
}

/*
CancelLimitOrder Cancel limit order (TRADE)
Post /sapi/v1/convert/limit/cancelOrder

https://developers.binance.com/en/docs/catalog/core-trading-convert/api/rest-api/trade#cancel-limit-order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param orderId -  The orderId from `placeOrder` api
@param recvWindow -  Request validity window in milliseconds
@return ApiCancelLimitOrderRequest
*/
func (a *TradeAPIService) CancelLimitOrder(ctx context.Context) ApiCancelLimitOrderRequest {
	return ApiCancelLimitOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CancelLimitOrderResponse
func (a *TradeAPIService) CancelLimitOrderExecute(r ApiCancelLimitOrderRequest) (*common.RestApiResponse[models.CancelLimitOrderResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/convert/limit/cancelOrder"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.orderId == nil {
		return nil, common.ReportError("orderId is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CancelLimitOrderResponse](
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

type ApiGetConvertTradeHistoryRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	startTime  *int64
	endTime    *int64
	limit      *int64
	recvWindow *int64
}

func (r ApiGetConvertTradeHistoryRequest) StartTime(startTime int64) ApiGetConvertTradeHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetConvertTradeHistoryRequest) EndTime(endTime int64) ApiGetConvertTradeHistoryRequest {
	r.endTime = &endTime
	return r
}

// Number of records to return
func (r ApiGetConvertTradeHistoryRequest) Limit(limit int64) ApiGetConvertTradeHistoryRequest {
	r.limit = &limit
	return r
}

// Request validity window in milliseconds
func (r ApiGetConvertTradeHistoryRequest) RecvWindow(recvWindow int64) ApiGetConvertTradeHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetConvertTradeHistoryRequest) Execute() (*common.RestApiResponse[models.GetConvertTradeHistoryResponse], error) {
	return r.ApiService.GetConvertTradeHistoryExecute(r)
}

/*
GetConvertTradeHistory Get Convert Trade History (USER_DATA)
Get /sapi/v1/convert/tradeFlow

https://developers.binance.com/en/docs/catalog/core-trading-convert/api/rest-api/trade#get-convert-trade-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -
@param endTime -
@param limit -  Number of records to return
@param recvWindow -  Request validity window in milliseconds
@return ApiGetConvertTradeHistoryRequest
*/
func (a *TradeAPIService) GetConvertTradeHistory(ctx context.Context) ApiGetConvertTradeHistoryRequest {
	return ApiGetConvertTradeHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetConvertTradeHistoryResponse
func (a *TradeAPIService) GetConvertTradeHistoryExecute(r ApiGetConvertTradeHistoryRequest) (*common.RestApiResponse[models.GetConvertTradeHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/convert/tradeFlow"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.startTime == nil {
		return nil, common.ReportError("startTime is required and must be specified")
	}

	if r.endTime == nil {
		return nil, common.ReportError("endTime is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetConvertTradeHistoryResponse](
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

type ApiOrderStatusRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	orderId    *string
	quoteId    *string
}

// Either orderId or quoteId is required
func (r ApiOrderStatusRequest) OrderId(orderId string) ApiOrderStatusRequest {
	r.orderId = &orderId
	return r
}

// Either orderId or quoteId is required
func (r ApiOrderStatusRequest) QuoteId(quoteId string) ApiOrderStatusRequest {
	r.quoteId = &quoteId
	return r
}

func (r ApiOrderStatusRequest) Execute() (*common.RestApiResponse[models.OrderStatusResponse], error) {
	return r.ApiService.OrderStatusExecute(r)
}

/*
OrderStatus Order status (USER_DATA)
Get /sapi/v1/convert/orderStatus

https://developers.binance.com/en/docs/catalog/core-trading-convert/api/rest-api/trade#order-status

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param orderId -  Either orderId or quoteId is required
@param quoteId -  Either orderId or quoteId is required
@return ApiOrderStatusRequest
*/
func (a *TradeAPIService) OrderStatus(ctx context.Context) ApiOrderStatusRequest {
	return ApiOrderStatusRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OrderStatusResponse
func (a *TradeAPIService) OrderStatusExecute(r ApiOrderStatusRequest) (*common.RestApiResponse[models.OrderStatusResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/convert/orderStatus"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.quoteId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quoteId", r.quoteId, "form", "")
	}

	resp, err := SendRequest[models.OrderStatusResponse](
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

type ApiPlaceLimitOrderRequest struct {
	ctx         context.Context
	ApiService  *TradeAPIService
	baseAsset   *string
	quoteAsset  *string
	limitPrice  *float64
	side        *models.PlaceLimitOrderSideParameter
	expiredType *models.PlaceLimitOrderExpiredTypeParameter
	baseAmount  *float64
	quoteAmount *float64
	walletType  *models.PlaceLimitOrderWalletTypeParameter
	recvWindow  *int64
}

// base asset (use the response &#x60;fromIsBase&#x60; from &#x60;GET /sapi/v1/convert/exchangeInfo&#x60; api to check which one is baseAsset )
func (r ApiPlaceLimitOrderRequest) BaseAsset(baseAsset string) ApiPlaceLimitOrderRequest {
	r.baseAsset = &baseAsset
	return r
}

// quote asset
func (r ApiPlaceLimitOrderRequest) QuoteAsset(quoteAsset string) ApiPlaceLimitOrderRequest {
	r.quoteAsset = &quoteAsset
	return r
}

// Symbol limit price (from baseAsset to quoteAsset)
func (r ApiPlaceLimitOrderRequest) LimitPrice(limitPrice float64) ApiPlaceLimitOrderRequest {
	r.limitPrice = &limitPrice
	return r
}

// &#x60;BUY&#x60; or &#x60;SELL&#x60;
func (r ApiPlaceLimitOrderRequest) Side(side models.PlaceLimitOrderSideParameter) ApiPlaceLimitOrderRequest {
	r.side = &side
	return r
}

// Order expiry duration. 1_D, 3_D, 7_D, 30_D (D means day)
func (r ApiPlaceLimitOrderRequest) ExpiredType(expiredType models.PlaceLimitOrderExpiredTypeParameter) ApiPlaceLimitOrderRequest {
	r.expiredType = &expiredType
	return r
}

// Base asset amount. (One of &#x60;baseAmount&#x60; or &#x60;quoteAmount&#x60; is required)
func (r ApiPlaceLimitOrderRequest) BaseAmount(baseAmount float64) ApiPlaceLimitOrderRequest {
	r.baseAmount = &baseAmount
	return r
}

// Quote asset amount. (One of &#x60;baseAmount&#x60; or &#x60;quoteAmount&#x60; is required)
func (r ApiPlaceLimitOrderRequest) QuoteAmount(quoteAmount float64) ApiPlaceLimitOrderRequest {
	r.quoteAmount = &quoteAmount
	return r
}

// Wallet to use for payment. Supported values: &#x60;SPOT&#x60;, &#x60;FUNDING&#x60;, &#x60;EARN&#x60;. Combined wallets also supported: &#x60;SPOT_FUNDING&#x60;, &#x60;FUNDING_EARN&#x60;, &#x60;SPOT_FUNDING_EARN&#x60;, &#x60;SPOT_EARN&#x60;. Default is &#x60;SPOT&#x60;.
func (r ApiPlaceLimitOrderRequest) WalletType(walletType models.PlaceLimitOrderWalletTypeParameter) ApiPlaceLimitOrderRequest {
	r.walletType = &walletType
	return r
}

// Request validity window in milliseconds
func (r ApiPlaceLimitOrderRequest) RecvWindow(recvWindow int64) ApiPlaceLimitOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiPlaceLimitOrderRequest) Execute() (*common.RestApiResponse[models.PlaceLimitOrderResponse], error) {
	return r.ApiService.PlaceLimitOrderExecute(r)
}

/*
PlaceLimitOrder Place limit order (TRADE)
Post /sapi/v1/convert/limit/placeOrder

https://developers.binance.com/en/docs/catalog/core-trading-convert/api/rest-api/trade#place-limit-order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param baseAsset -  base asset (use the response `fromIsBase` from `GET /sapi/v1/convert/exchangeInfo` api to check which one is baseAsset )
@param quoteAsset -  quote asset
@param limitPrice -  Symbol limit price (from baseAsset to quoteAsset)
@param side -  `BUY` or `SELL`
@param expiredType -  Order expiry duration. 1_D, 3_D, 7_D, 30_D (D means day)
@param baseAmount -  Base asset amount. (One of `baseAmount` or `quoteAmount` is required)
@param quoteAmount -  Quote asset amount. (One of `baseAmount` or `quoteAmount` is required)
@param walletType -  Wallet to use for payment. Supported values: `SPOT`, `FUNDING`, `EARN`. Combined wallets also supported: `SPOT_FUNDING`, `FUNDING_EARN`, `SPOT_FUNDING_EARN`, `SPOT_EARN`. Default is `SPOT`.
@param recvWindow -  Request validity window in milliseconds
@return ApiPlaceLimitOrderRequest
*/
func (a *TradeAPIService) PlaceLimitOrder(ctx context.Context) ApiPlaceLimitOrderRequest {
	return ApiPlaceLimitOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PlaceLimitOrderResponse
func (a *TradeAPIService) PlaceLimitOrderExecute(r ApiPlaceLimitOrderRequest) (*common.RestApiResponse[models.PlaceLimitOrderResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/convert/limit/placeOrder"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.baseAsset == nil {
		return nil, common.ReportError("baseAsset is required and must be specified")
	}

	if r.quoteAsset == nil {
		return nil, common.ReportError("quoteAsset is required and must be specified")
	}

	if r.limitPrice == nil {
		return nil, common.ReportError("limitPrice is required and must be specified")
	}

	if r.side == nil {
		return nil, common.ReportError("side is required and must be specified")
	}

	if r.expiredType == nil {
		return nil, common.ReportError("expiredType is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "baseAsset", r.baseAsset, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quoteAsset", r.quoteAsset, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limitPrice", r.limitPrice, "form", "")
	if r.baseAmount != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "baseAmount", r.baseAmount, "form", "")
	}
	if r.quoteAmount != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quoteAmount", r.quoteAmount, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "side", r.side, "form", "")
	if r.walletType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "walletType", r.walletType, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "expiredType", r.expiredType, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.PlaceLimitOrderResponse](
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

type ApiQueryLimitOpenOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	recvWindow *int64
}

// Request validity window in milliseconds
func (r ApiQueryLimitOpenOrdersRequest) RecvWindow(recvWindow int64) ApiQueryLimitOpenOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryLimitOpenOrdersRequest) Execute() (*common.RestApiResponse[models.QueryLimitOpenOrdersResponse], error) {
	return r.ApiService.QueryLimitOpenOrdersExecute(r)
}

/*
QueryLimitOpenOrders Query limit open orders (USER_DATA)
Get /sapi/v1/convert/limit/queryOpenOrders

https://developers.binance.com/en/docs/catalog/core-trading-convert/api/rest-api/trade#query-limit-open-orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -  Request validity window in milliseconds
@return ApiQueryLimitOpenOrdersRequest
*/
func (a *TradeAPIService) QueryLimitOpenOrders(ctx context.Context) ApiQueryLimitOpenOrdersRequest {
	return ApiQueryLimitOpenOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryLimitOpenOrdersResponse
func (a *TradeAPIService) QueryLimitOpenOrdersExecute(r ApiQueryLimitOpenOrdersRequest) (*common.RestApiResponse[models.QueryLimitOpenOrdersResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/convert/limit/queryOpenOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryLimitOpenOrdersResponse](
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

type ApiSendQuoteRequestRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	fromAsset  *string
	toAsset    *string
	fromAmount *float64
	toAmount   *float64
	walletType *models.PlaceLimitOrderWalletTypeParameter
	validTime  *models.SendQuoteRequestValidTimeParameter
	recvWindow *int64
}

func (r ApiSendQuoteRequestRequest) FromAsset(fromAsset string) ApiSendQuoteRequestRequest {
	r.fromAsset = &fromAsset
	return r
}

func (r ApiSendQuoteRequestRequest) ToAsset(toAsset string) ApiSendQuoteRequestRequest {
	r.toAsset = &toAsset
	return r
}

// When specified, it is the amount you will be debited after the conversion
func (r ApiSendQuoteRequestRequest) FromAmount(fromAmount float64) ApiSendQuoteRequestRequest {
	r.fromAmount = &fromAmount
	return r
}

// When specified, it is the amount you will be credited after the conversion
func (r ApiSendQuoteRequestRequest) ToAmount(toAmount float64) ApiSendQuoteRequestRequest {
	r.toAmount = &toAmount
	return r
}

// Wallet to use for payment. Supported values: &#x60;SPOT&#x60;, &#x60;FUNDING&#x60;, &#x60;EARN&#x60;. Combined wallets also supported: &#x60;SPOT_FUNDING&#x60;, &#x60;FUNDING_EARN&#x60;, &#x60;SPOT_FUNDING_EARN&#x60;, &#x60;SPOT_EARN&#x60;. Default is &#x60;SPOT&#x60;.
func (r ApiSendQuoteRequestRequest) WalletType(walletType models.PlaceLimitOrderWalletTypeParameter) ApiSendQuoteRequestRequest {
	r.walletType = &walletType
	return r
}

// Quote valid duration. Supported values: 10s, 30s, 1m. Default is 10s.
func (r ApiSendQuoteRequestRequest) ValidTime(validTime models.SendQuoteRequestValidTimeParameter) ApiSendQuoteRequestRequest {
	r.validTime = &validTime
	return r
}

// Request validity window in milliseconds
func (r ApiSendQuoteRequestRequest) RecvWindow(recvWindow int64) ApiSendQuoteRequestRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSendQuoteRequestRequest) Execute() (*common.RestApiResponse[models.SendQuoteRequestResponse], error) {
	return r.ApiService.SendQuoteRequestExecute(r)
}

/*
SendQuoteRequest Send Quote Request (TRADE)
Post /sapi/v1/convert/getQuote

https://developers.binance.com/en/docs/catalog/core-trading-convert/api/rest-api/trade#send-quote-request

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param fromAsset -
@param toAsset -
@param fromAmount -  When specified, it is the amount you will be debited after the conversion
@param toAmount -  When specified, it is the amount you will be credited after the conversion
@param walletType -  Wallet to use for payment. Supported values: `SPOT`, `FUNDING`, `EARN`. Combined wallets also supported: `SPOT_FUNDING`, `FUNDING_EARN`, `SPOT_FUNDING_EARN`, `SPOT_EARN`. Default is `SPOT`.
@param validTime -  Quote valid duration. Supported values: 10s, 30s, 1m. Default is 10s.
@param recvWindow -  Request validity window in milliseconds
@return ApiSendQuoteRequestRequest
*/
func (a *TradeAPIService) SendQuoteRequest(ctx context.Context) ApiSendQuoteRequestRequest {
	return ApiSendQuoteRequestRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SendQuoteRequestResponse
func (a *TradeAPIService) SendQuoteRequestExecute(r ApiSendQuoteRequestRequest) (*common.RestApiResponse[models.SendQuoteRequestResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/convert/getQuote"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.fromAsset == nil {
		return nil, common.ReportError("fromAsset is required and must be specified")
	}

	if r.toAsset == nil {
		return nil, common.ReportError("toAsset is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromAsset", r.fromAsset, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "toAsset", r.toAsset, "form", "")
	if r.fromAmount != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromAmount", r.fromAmount, "form", "")
	}
	if r.toAmount != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "toAmount", r.toAmount, "form", "")
	}
	if r.walletType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "walletType", r.walletType, "form", "")
	}
	if r.validTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "validTime", r.validTime, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.SendQuoteRequestResponse](
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
