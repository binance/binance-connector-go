/*
Binance Convert REST API

OpenAPI Specification for the Binance Convert REST API
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

// The value cannot be greater than 60000
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

https://developers.binance.com/docs/convert/trade/Accept-Quote

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param quoteId -
@param recvWindow -  The value cannot be greater than 60000
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

	resp, err := SendRequest[models.AcceptQuoteResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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

// The value cannot be greater than 60000
func (r ApiCancelLimitOrderRequest) RecvWindow(recvWindow int64) ApiCancelLimitOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelLimitOrderRequest) Execute() (*common.RestApiResponse[models.CancelLimitOrderResponse], error) {
	return r.ApiService.CancelLimitOrderExecute(r)
}

/*
CancelLimitOrder Cancel limit order (USER_DATA)
Post /sapi/v1/convert/limit/cancelOrder

https://developers.binance.com/docs/convert/trade/Cancel-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param orderId -  The orderId from `placeOrder` api
@param recvWindow -  The value cannot be greater than 60000
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

	resp, err := SendRequest[models.CancelLimitOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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

// Default 100, Max 1000
func (r ApiGetConvertTradeHistoryRequest) Limit(limit int64) ApiGetConvertTradeHistoryRequest {
	r.limit = &limit
	return r
}

// The value cannot be greater than 60000
func (r ApiGetConvertTradeHistoryRequest) RecvWindow(recvWindow int64) ApiGetConvertTradeHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetConvertTradeHistoryRequest) Execute() (*common.RestApiResponse[models.GetConvertTradeHistoryResponse], error) {
	return r.ApiService.GetConvertTradeHistoryExecute(r)
}

/*
GetConvertTradeHistory Get Convert Trade History(USER_DATA)
Get /sapi/v1/convert/tradeFlow

https://developers.binance.com/docs/convert/trade/Get-Convert-Trade-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -
@param endTime -
@param limit -  Default 100, Max 1000
@param recvWindow -  The value cannot be greater than 60000
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

	resp, err := SendRequest[models.GetConvertTradeHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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
OrderStatus Order status(USER_DATA)
Get /sapi/v1/convert/orderStatus

https://developers.binance.com/docs/convert/trade/Order-Status

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

	resp, err := SendRequest[models.OrderStatusResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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
	limitPrice  *float32
	side        *string
	expiredType *string
	baseAmount  *float32
	quoteAmount *float32
	walletType  *string
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
func (r ApiPlaceLimitOrderRequest) LimitPrice(limitPrice float32) ApiPlaceLimitOrderRequest {
	r.limitPrice = &limitPrice
	return r
}

// &#x60;BUY&#x60; or &#x60;SELL&#x60;
func (r ApiPlaceLimitOrderRequest) Side(side string) ApiPlaceLimitOrderRequest {
	r.side = &side
	return r
}

// 1_D, 3_D, 7_D, 30_D  (D means day)
func (r ApiPlaceLimitOrderRequest) ExpiredType(expiredType string) ApiPlaceLimitOrderRequest {
	r.expiredType = &expiredType
	return r
}

// Base asset amount.  (One of &#x60;baseAmount&#x60; or &#x60;quoteAmount&#x60; is required)
func (r ApiPlaceLimitOrderRequest) BaseAmount(baseAmount float32) ApiPlaceLimitOrderRequest {
	r.baseAmount = &baseAmount
	return r
}

// Quote asset amount.  (One of &#x60;baseAmount&#x60; or &#x60;quoteAmount&#x60; is required)
func (r ApiPlaceLimitOrderRequest) QuoteAmount(quoteAmount float32) ApiPlaceLimitOrderRequest {
	r.quoteAmount = &quoteAmount
	return r
}

// It is to choose which wallet of assets. The wallet selection is &#x60;SPOT&#x60;, &#x60;FUNDING&#x60; and &#x60;EARN&#x60;. Combination of wallet is supported i.e. &#x60;SPOT_FUNDING&#x60;, &#x60;FUNDING_EARN&#x60;, &#x60;SPOT_FUNDING_EARN&#x60; or &#x60;SPOT_EARN&#x60;  Default is &#x60;SPOT&#x60;.
func (r ApiPlaceLimitOrderRequest) WalletType(walletType string) ApiPlaceLimitOrderRequest {
	r.walletType = &walletType
	return r
}

// The value cannot be greater than 60000
func (r ApiPlaceLimitOrderRequest) RecvWindow(recvWindow int64) ApiPlaceLimitOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiPlaceLimitOrderRequest) Execute() (*common.RestApiResponse[models.PlaceLimitOrderResponse], error) {
	return r.ApiService.PlaceLimitOrderExecute(r)
}

/*
PlaceLimitOrder Place limit order (USER_DATA)
Post /sapi/v1/convert/limit/placeOrder

https://developers.binance.com/docs/convert/trade/Place-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param baseAsset -  base asset (use the response `fromIsBase` from `GET /sapi/v1/convert/exchangeInfo` api to check which one is baseAsset )
@param quoteAsset -  quote asset
@param limitPrice -  Symbol limit price (from baseAsset to quoteAsset)
@param side -  `BUY` or `SELL`
@param expiredType -  1_D, 3_D, 7_D, 30_D  (D means day)
@param baseAmount -  Base asset amount.  (One of `baseAmount` or `quoteAmount` is required)
@param quoteAmount -  Quote asset amount.  (One of `baseAmount` or `quoteAmount` is required)
@param walletType -  It is to choose which wallet of assets. The wallet selection is `SPOT`, `FUNDING` and `EARN`. Combination of wallet is supported i.e. `SPOT_FUNDING`, `FUNDING_EARN`, `SPOT_FUNDING_EARN` or `SPOT_EARN`  Default is `SPOT`.
@param recvWindow -  The value cannot be greater than 60000
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

	resp, err := SendRequest[models.PlaceLimitOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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

// The value cannot be greater than 60000
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

https://developers.binance.com/docs/convert/trade/Query-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -  The value cannot be greater than 60000
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

	resp, err := SendRequest[models.QueryLimitOpenOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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
	fromAmount *float32
	toAmount   *float32
	walletType *string
	validTime  *string
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
func (r ApiSendQuoteRequestRequest) FromAmount(fromAmount float32) ApiSendQuoteRequestRequest {
	r.fromAmount = &fromAmount
	return r
}

// When specified, it is the amount you will be credited after the conversion
func (r ApiSendQuoteRequestRequest) ToAmount(toAmount float32) ApiSendQuoteRequestRequest {
	r.toAmount = &toAmount
	return r
}

// It is to choose which wallet of assets. The wallet selection is &#x60;SPOT&#x60;, &#x60;FUNDING&#x60; and &#x60;EARN&#x60;. Combination of wallet is supported i.e. &#x60;SPOT_FUNDING&#x60;, &#x60;FUNDING_EARN&#x60;, &#x60;SPOT_FUNDING_EARN&#x60; or &#x60;SPOT_EARN&#x60;  Default is &#x60;SPOT&#x60;.
func (r ApiSendQuoteRequestRequest) WalletType(walletType string) ApiSendQuoteRequestRequest {
	r.walletType = &walletType
	return r
}

// 10s, 30s, 1m, default 10s
func (r ApiSendQuoteRequestRequest) ValidTime(validTime string) ApiSendQuoteRequestRequest {
	r.validTime = &validTime
	return r
}

// The value cannot be greater than 60000
func (r ApiSendQuoteRequestRequest) RecvWindow(recvWindow int64) ApiSendQuoteRequestRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSendQuoteRequestRequest) Execute() (*common.RestApiResponse[models.SendQuoteRequestResponse], error) {
	return r.ApiService.SendQuoteRequestExecute(r)
}

/*
SendQuoteRequest Send Quote Request(USER_DATA)
Post /sapi/v1/convert/getQuote

https://developers.binance.com/docs/convert/trade/Send-quote-request

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param fromAsset -
@param toAsset -
@param fromAmount -  When specified, it is the amount you will be debited after the conversion
@param toAmount -  When specified, it is the amount you will be credited after the conversion
@param walletType -  It is to choose which wallet of assets. The wallet selection is `SPOT`, `FUNDING` and `EARN`. Combination of wallet is supported i.e. `SPOT_FUNDING`, `FUNDING_EARN`, `SPOT_FUNDING_EARN` or `SPOT_EARN`  Default is `SPOT`.
@param validTime -  10s, 30s, 1m, default 10s
@param recvWindow -  The value cannot be greater than 60000
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

	resp, err := SendRequest[models.SendQuoteRequestResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
