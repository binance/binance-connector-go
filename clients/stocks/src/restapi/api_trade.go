/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package binancestocksrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/stocks/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// TradeAPIService TradeAPI Service
type TradeAPIService Service

type ApiCancelAllEquityOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	recvWindow *int64
}

// The value cannot be greater than &#x60;60000&#x60;.
func (r ApiCancelAllEquityOrdersRequest) RecvWindow(recvWindow int64) ApiCancelAllEquityOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelAllEquityOrdersRequest) Execute() (*common.RestApiResponse[models.CancelAllEquityOrdersResponse], error) {
	return r.ApiService.CancelAllEquityOrdersExecute(r)
}

/*
CancelAllEquityOrders Cancel All Equity Orders (TRADE)
Post /sapi/v1/equity/order/cancel-all

https://developers.binance.com/en/docs/catalog/advanced-trading-stocks-trading/api/rest-api/trade#cancel-all-equity-orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -  The value cannot be greater than `60000`.
@return ApiCancelAllEquityOrdersRequest
*/
func (a *TradeAPIService) CancelAllEquityOrders(ctx context.Context) ApiCancelAllEquityOrdersRequest {
	return ApiCancelAllEquityOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CancelAllEquityOrdersResponse
func (a *TradeAPIService) CancelAllEquityOrdersExecute(r ApiCancelAllEquityOrdersRequest) (*common.RestApiResponse[models.CancelAllEquityOrdersResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/equity/order/cancel-all"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CancelAllEquityOrdersResponse](
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

type ApiCancelEquityOrderRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	orderId    *string
	recvWindow *int64
}

// Equity order id returned by &#x60;/order/place&#x60; or a query endpoint.
func (r ApiCancelEquityOrderRequest) OrderId(orderId string) ApiCancelEquityOrderRequest {
	r.orderId = &orderId
	return r
}

// The value cannot be greater than &#x60;60000&#x60;.
func (r ApiCancelEquityOrderRequest) RecvWindow(recvWindow int64) ApiCancelEquityOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelEquityOrderRequest) Execute() (*common.RestApiResponse[models.CancelEquityOrderResponse], error) {
	return r.ApiService.CancelEquityOrderExecute(r)
}

/*
CancelEquityOrder Cancel Equity Order (TRADE)
Post /sapi/v1/equity/order/cancel

https://developers.binance.com/en/docs/catalog/advanced-trading-stocks-trading/api/rest-api/trade#cancel-equity-order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param orderId -  Equity order id returned by `/order/place` or a query endpoint.
@param recvWindow -  The value cannot be greater than `60000`.
@return ApiCancelEquityOrderRequest
*/
func (a *TradeAPIService) CancelEquityOrder(ctx context.Context) ApiCancelEquityOrderRequest {
	return ApiCancelEquityOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CancelEquityOrderResponse
func (a *TradeAPIService) CancelEquityOrderExecute(r ApiCancelEquityOrderRequest) (*common.RestApiResponse[models.CancelEquityOrderResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/equity/order/cancel"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.orderId == nil {
		return nil, common.ReportError("orderId is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CancelEquityOrderResponse](
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

type ApiCurrentOpenOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	recvWindow *int64
}

// The value cannot be greater than &#x60;60000&#x60;.
func (r ApiCurrentOpenOrdersRequest) RecvWindow(recvWindow int64) ApiCurrentOpenOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCurrentOpenOrdersRequest) Execute() (*common.RestApiResponse[models.CurrentOpenOrdersResponse], error) {
	return r.ApiService.CurrentOpenOrdersExecute(r)
}

/*
CurrentOpenOrders Current Open Orders (USER_DATA)
Get /sapi/v1/equity/order/open-orders

https://developers.binance.com/en/docs/catalog/advanced-trading-stocks-trading/api/rest-api/trade#current-open-orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -  The value cannot be greater than `60000`.
@return ApiCurrentOpenOrdersRequest
*/
func (a *TradeAPIService) CurrentOpenOrders(ctx context.Context) ApiCurrentOpenOrdersRequest {
	return ApiCurrentOpenOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CurrentOpenOrdersResponse
func (a *TradeAPIService) CurrentOpenOrdersExecute(r ApiCurrentOpenOrdersRequest) (*common.RestApiResponse[models.CurrentOpenOrdersResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/equity/order/open-orders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CurrentOpenOrdersResponse](
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

type ApiEquityOrderDetailRequest struct {
	ctx           context.Context
	ApiService    *TradeAPIService
	orderId       *string
	clientOrderId *string
	recvWindow    *int64
}

// Equity order id. Either &#x60;orderId&#x60; or &#x60;clientOrderId&#x60; must be provided.
func (r ApiEquityOrderDetailRequest) OrderId(orderId string) ApiEquityOrderDetailRequest {
	r.orderId = &orderId
	return r
}

// Client-supplied order id. Either &#x60;orderId&#x60; or &#x60;clientOrderId&#x60; must be provided.
func (r ApiEquityOrderDetailRequest) ClientOrderId(clientOrderId string) ApiEquityOrderDetailRequest {
	r.clientOrderId = &clientOrderId
	return r
}

// The value cannot be greater than &#x60;60000&#x60;.
func (r ApiEquityOrderDetailRequest) RecvWindow(recvWindow int64) ApiEquityOrderDetailRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiEquityOrderDetailRequest) Execute() (*common.RestApiResponse[models.EquityOrderDetailResponse], error) {
	return r.ApiService.EquityOrderDetailExecute(r)
}

/*
EquityOrderDetail Equity Order Detail (USER_DATA)
Get /sapi/v1/equity/order/detail

https://developers.binance.com/en/docs/catalog/advanced-trading-stocks-trading/api/rest-api/trade#equity-order-detail

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param orderId -  Equity order id. Either `orderId` or `clientOrderId` must be provided.
@param clientOrderId -  Client-supplied order id. Either `orderId` or `clientOrderId` must be provided.
@param recvWindow -  The value cannot be greater than `60000`.
@return ApiEquityOrderDetailRequest
*/
func (a *TradeAPIService) EquityOrderDetail(ctx context.Context) ApiEquityOrderDetailRequest {
	return ApiEquityOrderDetailRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return EquityOrderDetailResponse
func (a *TradeAPIService) EquityOrderDetailExecute(r ApiEquityOrderDetailRequest) (*common.RestApiResponse[models.EquityOrderDetailResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/equity/order/detail"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.clientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "clientOrderId", r.clientOrderId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.EquityOrderDetailResponse](
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

type ApiEquityOrderHistoryRequest struct {
	ctx         context.Context
	ApiService  *TradeAPIService
	startTime   *int64
	endTime     *int64
	symbol      *string
	orderType   *models.PlaceEquityOrderOrderTypeParameter
	side        *models.PlaceEquityOrderSideParameter
	orderStatus *string
	current     *int32
	size        *int32
	recvWindow  *int64
}

// Start time (ms epoch).
func (r ApiEquityOrderHistoryRequest) StartTime(startTime int64) ApiEquityOrderHistoryRequest {
	r.startTime = &startTime
	return r
}

// End time (ms epoch).
func (r ApiEquityOrderHistoryRequest) EndTime(endTime int64) ApiEquityOrderHistoryRequest {
	r.endTime = &endTime
	return r
}

// US-equity ticker filter, e.g. &#x60;NVDA&#x60;.
func (r ApiEquityOrderHistoryRequest) Symbol(symbol string) ApiEquityOrderHistoryRequest {
	r.symbol = &symbol
	return r
}

// Order type filter: &#x60;MARKET&#x60; / &#x60;LIMIT&#x60;.
func (r ApiEquityOrderHistoryRequest) OrderType(orderType models.PlaceEquityOrderOrderTypeParameter) ApiEquityOrderHistoryRequest {
	r.orderType = &orderType
	return r
}

// Side filter: &#x60;BUY&#x60; / &#x60;SELL&#x60;.
func (r ApiEquityOrderHistoryRequest) Side(side models.PlaceEquityOrderSideParameter) ApiEquityOrderHistoryRequest {
	r.side = &side
	return r
}

// Comma-separated status filter. Allowed values: &#x60;FILLED&#x60;, &#x60;PARTIALLY_FILLED&#x60;, &#x60;CANCELED&#x60;, &#x60;EXPIRED&#x60;, &#x60;REJECTED&#x60;.
func (r ApiEquityOrderHistoryRequest) OrderStatus(orderStatus string) ApiEquityOrderHistoryRequest {
	r.orderStatus = &orderStatus
	return r
}

// Page number, 1-based. Default &#x60;1&#x60;.
func (r ApiEquityOrderHistoryRequest) Current(current int32) ApiEquityOrderHistoryRequest {
	r.current = &current
	return r
}

// Page size. Default &#x60;20&#x60;, max &#x60;100&#x60;.
func (r ApiEquityOrderHistoryRequest) Size(size int32) ApiEquityOrderHistoryRequest {
	r.size = &size
	return r
}

// The value cannot be greater than &#x60;60000&#x60;.
func (r ApiEquityOrderHistoryRequest) RecvWindow(recvWindow int64) ApiEquityOrderHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiEquityOrderHistoryRequest) Execute() (*common.RestApiResponse[models.EquityOrderHistoryResponse], error) {
	return r.ApiService.EquityOrderHistoryExecute(r)
}

/*
EquityOrderHistory Equity Order History (USER_DATA)
Get /sapi/v1/equity/order/history

https://developers.binance.com/en/docs/catalog/advanced-trading-stocks-trading/api/rest-api/trade#equity-order-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -  Start time (ms epoch).
@param endTime -  End time (ms epoch).
@param symbol -  US-equity ticker filter, e.g. `NVDA`.
@param orderType -  Order type filter: `MARKET` / `LIMIT`.
@param side -  Side filter: `BUY` / `SELL`.
@param orderStatus -  Comma-separated status filter. Allowed values: `FILLED`, `PARTIALLY_FILLED`, `CANCELED`, `EXPIRED`, `REJECTED`.
@param current -  Page number, 1-based. Default `1`.
@param size -  Page size. Default `20`, max `100`.
@param recvWindow -  The value cannot be greater than `60000`.
@return ApiEquityOrderHistoryRequest
*/
func (a *TradeAPIService) EquityOrderHistory(ctx context.Context) ApiEquityOrderHistoryRequest {
	return ApiEquityOrderHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return EquityOrderHistoryResponse
func (a *TradeAPIService) EquityOrderHistoryExecute(r ApiEquityOrderHistoryRequest) (*common.RestApiResponse[models.EquityOrderHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/equity/order/history"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.startTime == nil {
		return nil, common.ReportError("startTime is required and must be specified")
	}

	if r.endTime == nil {
		return nil, common.ReportError("endTime is required and must be specified")
	}

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.orderType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderType", r.orderType, "form", "")
	}
	if r.side != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "side", r.side, "form", "")
	}
	if r.orderStatus != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderStatus", r.orderStatus, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	if r.current != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "current", r.current, "form", "")
	}
	if r.size != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.EquityOrderHistoryResponse](
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

type ApiEquityTradeHistoryRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	startTime  *int64
	endTime    *int64
	symbol     *string
	side       *models.PlaceEquityOrderSideParameter
	orderId    *string
	current    *int32
	size       *int32
	recvWindow *int64
}

// Start time (ms epoch).
func (r ApiEquityTradeHistoryRequest) StartTime(startTime int64) ApiEquityTradeHistoryRequest {
	r.startTime = &startTime
	return r
}

// End time (ms epoch).
func (r ApiEquityTradeHistoryRequest) EndTime(endTime int64) ApiEquityTradeHistoryRequest {
	r.endTime = &endTime
	return r
}

// US-equity ticker filter, e.g. &#x60;NVDA&#x60;.
func (r ApiEquityTradeHistoryRequest) Symbol(symbol string) ApiEquityTradeHistoryRequest {
	r.symbol = &symbol
	return r
}

// Side filter: &#x60;BUY&#x60; / &#x60;SELL&#x60;.
func (r ApiEquityTradeHistoryRequest) Side(side models.PlaceEquityOrderSideParameter) ApiEquityTradeHistoryRequest {
	r.side = &side
	return r
}

// Narrow the result to executions of a single order.
func (r ApiEquityTradeHistoryRequest) OrderId(orderId string) ApiEquityTradeHistoryRequest {
	r.orderId = &orderId
	return r
}

// Page number, 1-based. Default &#x60;1&#x60;.
func (r ApiEquityTradeHistoryRequest) Current(current int32) ApiEquityTradeHistoryRequest {
	r.current = &current
	return r
}

// Page size. Default &#x60;20&#x60;, max &#x60;100&#x60;.
func (r ApiEquityTradeHistoryRequest) Size(size int32) ApiEquityTradeHistoryRequest {
	r.size = &size
	return r
}

// The value cannot be greater than &#x60;60000&#x60;.
func (r ApiEquityTradeHistoryRequest) RecvWindow(recvWindow int64) ApiEquityTradeHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiEquityTradeHistoryRequest) Execute() (*common.RestApiResponse[models.EquityTradeHistoryResponse], error) {
	return r.ApiService.EquityTradeHistoryExecute(r)
}

/*
EquityTradeHistory Equity Trade History (USER_DATA)
Get /sapi/v1/equity/trade/history

https://developers.binance.com/en/docs/catalog/advanced-trading-stocks-trading/api/rest-api/trade#equity-trade-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -  Start time (ms epoch).
@param endTime -  End time (ms epoch).
@param symbol -  US-equity ticker filter, e.g. `NVDA`.
@param side -  Side filter: `BUY` / `SELL`.
@param orderId -  Narrow the result to executions of a single order.
@param current -  Page number, 1-based. Default `1`.
@param size -  Page size. Default `20`, max `100`.
@param recvWindow -  The value cannot be greater than `60000`.
@return ApiEquityTradeHistoryRequest
*/
func (a *TradeAPIService) EquityTradeHistory(ctx context.Context) ApiEquityTradeHistoryRequest {
	return ApiEquityTradeHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return EquityTradeHistoryResponse
func (a *TradeAPIService) EquityTradeHistoryExecute(r ApiEquityTradeHistoryRequest) (*common.RestApiResponse[models.EquityTradeHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/equity/trade/history"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.startTime == nil {
		return nil, common.ReportError("startTime is required and must be specified")
	}

	if r.endTime == nil {
		return nil, common.ReportError("endTime is required and must be specified")
	}

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.side != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "side", r.side, "form", "")
	}
	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	if r.current != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "current", r.current, "form", "")
	}
	if r.size != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.EquityTradeHistoryResponse](
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

type ApiPlaceEquityOrderRequest struct {
	ctx            context.Context
	ApiService     *TradeAPIService
	symbol         *string
	side           *models.PlaceEquityOrderSideParameter
	orderType      *models.PlaceEquityOrderOrderTypeParameter
	quoteAsset     *string
	price          *string
	quantity       *string
	notional       *string
	timeInForce    *models.PlaceEquityOrderTimeInForceParameter
	tradingSession *models.PlaceEquityOrderTradingSessionParameter
	walletType     *models.PlaceEquityOrderWalletTypeParameter
	clientOrderId  *string
	tokenize       *bool
	recvWindow     *int64
}

// US stock ticker, e.g. &#x60;AAPL&#x60;, &#x60;TSLA&#x60;. Must be a symbol with tokenization enabled — check via &#x60;/market/tokenized-assets&#x60;.
func (r ApiPlaceEquityOrderRequest) Symbol(symbol string) ApiPlaceEquityOrderRequest {
	r.symbol = &symbol
	return r
}

// &#x60;BUY&#x60; / &#x60;SELL&#x60;.
func (r ApiPlaceEquityOrderRequest) Side(side models.PlaceEquityOrderSideParameter) ApiPlaceEquityOrderRequest {
	r.side = &side
	return r
}

// &#x60;MARKET&#x60; / &#x60;LIMIT&#x60;.
func (r ApiPlaceEquityOrderRequest) OrderType(orderType models.PlaceEquityOrderOrderTypeParameter) ApiPlaceEquityOrderRequest {
	r.orderType = &orderType
	return r
}

// Quote asset. Defaults to &#x60;USDC&#x60;; must be within the server&#39;s allowed set.
func (r ApiPlaceEquityOrderRequest) QuoteAsset(quoteAsset string) ApiPlaceEquityOrderRequest {
	r.quoteAsset = &quoteAsset
	return r
}

// **Required** for &#x60;LIMIT&#x60;; **forbidden** for &#x60;MARKET&#x60;. Maximum 2 decimal places.
func (r ApiPlaceEquityOrderRequest) Price(price string) ApiPlaceEquityOrderRequest {
	r.price = &price
	return r
}

// **Required** for &#x60;LIMIT&#x60; (both sides) and &#x60;SELL MARKET&#x60;; **forbidden** for &#x60;BUY MARKET&#x60;.
func (r ApiPlaceEquityOrderRequest) Quantity(quantity string) ApiPlaceEquityOrderRequest {
	r.quantity = &quantity
	return r
}

// **Required** for &#x60;BUY MARKET&#x60;; **forbidden** for &#x60;LIMIT&#x60; and &#x60;SELL MARKET&#x60;.
func (r ApiPlaceEquityOrderRequest) Notional(notional string) ApiPlaceEquityOrderRequest {
	r.notional = &notional
	return r
}

// &#x60;DAY&#x60; (default) / &#x60;GTC&#x60;. &#x60;GTC&#x60; is only supported for &#x60;LIMIT&#x60; orders; a fractional-share &#x60;GTC&#x60; order must be paired with &#x60;tradingSession &#x3D; EXTENDED&#x60; or &#x60;24H&#x60;.
func (r ApiPlaceEquityOrderRequest) TimeInForce(timeInForce models.PlaceEquityOrderTimeInForceParameter) ApiPlaceEquityOrderRequest {
	r.timeInForce = &timeInForce
	return r
}

// &#x60;RTH&#x60; / &#x60;EXTENDED&#x60; / &#x60;24H&#x60;. **Required** for &#x60;LIMIT&#x60;; **forbidden** for &#x60;MARKET&#x60;.
func (r ApiPlaceEquityOrderRequest) TradingSession(tradingSession models.PlaceEquityOrderTradingSessionParameter) ApiPlaceEquityOrderRequest {
	r.tradingSession = &tradingSession
	return r
}

// Payment wallet for &#x60;BUY&#x60; orders: &#x60;CARD&#x60; (default) / &#x60;MAIN&#x60;. &#x60;SELL&#x60; orders always settle to &#x60;CARD&#x60;.
func (r ApiPlaceEquityOrderRequest) WalletType(walletType models.PlaceEquityOrderWalletTypeParameter) ApiPlaceEquityOrderRequest {
	r.walletType = &walletType
	return r
}

// Client-supplied order id. Format &#x60;^[a-zA-Z0-9-_]{32,36}$&#x60;. Auto-generated when omitted.
func (r ApiPlaceEquityOrderRequest) ClientOrderId(clientOrderId string) ApiPlaceEquityOrderRequest {
	r.clientOrderId = &clientOrderId
	return r
}

// Whether to tokenize the purchased stock asset upon settlement. Default &#x60;true&#x60;. Set to &#x60;false&#x60; to receive the underlying equity directly instead of a tokenized asset.
func (r ApiPlaceEquityOrderRequest) Tokenize(tokenize bool) ApiPlaceEquityOrderRequest {
	r.tokenize = &tokenize
	return r
}

// The value cannot be greater than &#x60;60000&#x60;.
func (r ApiPlaceEquityOrderRequest) RecvWindow(recvWindow int64) ApiPlaceEquityOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiPlaceEquityOrderRequest) Execute() (*common.RestApiResponse[models.PlaceEquityOrderResponse], error) {
	return r.ApiService.PlaceEquityOrderExecute(r)
}

/*
PlaceEquityOrder Place Equity Order (TRADE)
Post /sapi/v1/equity/order/place

https://developers.binance.com/en/docs/catalog/advanced-trading-stocks-trading/api/rest-api/trade#place-equity-order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  US stock ticker, e.g. `AAPL`, `TSLA`. Must be a symbol with tokenization enabled — check via `/market/tokenized-assets`.
@param side -  `BUY` / `SELL`.
@param orderType -  `MARKET` / `LIMIT`.
@param quoteAsset -  Quote asset. Defaults to `USDC`; must be within the server's allowed set.
@param price -  **Required** for `LIMIT`; **forbidden** for `MARKET`. Maximum 2 decimal places.
@param quantity -  **Required** for `LIMIT` (both sides) and `SELL MARKET`; **forbidden** for `BUY MARKET`.
@param notional -  **Required** for `BUY MARKET`; **forbidden** for `LIMIT` and `SELL MARKET`.
@param timeInForce -  `DAY` (default) / `GTC`. `GTC` is only supported for `LIMIT` orders; a fractional-share `GTC` order must be paired with `tradingSession = EXTENDED` or `24H`.
@param tradingSession -  `RTH` / `EXTENDED` / `24H`. **Required** for `LIMIT`; **forbidden** for `MARKET`.
@param walletType -  Payment wallet for `BUY` orders: `CARD` (default) / `MAIN`. `SELL` orders always settle to `CARD`.
@param clientOrderId -  Client-supplied order id. Format `^[a-zA-Z0-9-_]{32,36}$`. Auto-generated when omitted.
@param tokenize -  Whether to tokenize the purchased stock asset upon settlement. Default `true`. Set to `false` to receive the underlying equity directly instead of a tokenized asset.
@param recvWindow -  The value cannot be greater than `60000`.
@return ApiPlaceEquityOrderRequest
*/
func (a *TradeAPIService) PlaceEquityOrder(ctx context.Context) ApiPlaceEquityOrderRequest {
	return ApiPlaceEquityOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PlaceEquityOrderResponse
func (a *TradeAPIService) PlaceEquityOrderExecute(r ApiPlaceEquityOrderRequest) (*common.RestApiResponse[models.PlaceEquityOrderResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/equity/order/place"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	if r.side == nil {
		return nil, common.ReportError("side is required and must be specified")
	}

	if r.orderType == nil {
		return nil, common.ReportError("orderType is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.quoteAsset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quoteAsset", r.quoteAsset, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "side", r.side, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderType", r.orderType, "form", "")
	if r.price != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "price", r.price, "form", "")
	}
	if r.quantity != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quantity", r.quantity, "form", "")
	}
	if r.notional != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "notional", r.notional, "form", "")
	}
	if r.timeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "timeInForce", r.timeInForce, "form", "")
	}
	if r.tradingSession != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "tradingSession", r.tradingSession, "form", "")
	}
	if r.walletType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "walletType", r.walletType, "form", "")
	}
	if r.clientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "clientOrderId", r.clientOrderId, "form", "")
	}
	if r.tokenize != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "tokenize", r.tokenize, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.PlaceEquityOrderResponse](
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
