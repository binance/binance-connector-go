/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package binancederivativestradingportfoliomarginrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// TradeAPIService TradeAPI Service
type TradeAPIService Service

type ApiCancelAllCmOpenConditionalOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiCancelAllCmOpenConditionalOrdersRequest) Symbol(symbol string) ApiCancelAllCmOpenConditionalOrdersRequest {
	r.symbol = &symbol
	return r
}

func (r ApiCancelAllCmOpenConditionalOrdersRequest) RecvWindow(recvWindow int64) ApiCancelAllCmOpenConditionalOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelAllCmOpenConditionalOrdersRequest) Execute() (*common.RestApiResponse[models.CancelAllCmOpenConditionalOrdersResponse], error) {
	return r.ApiService.CancelAllCmOpenConditionalOrdersExecute(r)
}

/*
CancelAllCmOpenConditionalOrders Cancel All CM Open Conditional Orders(TRADE)
Delete /papi/v1/cm/conditional/allOpenOrders

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Cancel-All-CM-Open-Conditional-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -
@return ApiCancelAllCmOpenConditionalOrdersRequest
*/
func (a *TradeAPIService) CancelAllCmOpenConditionalOrders(ctx context.Context) ApiCancelAllCmOpenConditionalOrdersRequest {
	return ApiCancelAllCmOpenConditionalOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CancelAllCmOpenConditionalOrdersResponse
func (a *TradeAPIService) CancelAllCmOpenConditionalOrdersExecute(r ApiCancelAllCmOpenConditionalOrdersRequest) (*common.RestApiResponse[models.CancelAllCmOpenConditionalOrdersResponse], error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/papi/v1/cm/conditional/allOpenOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CancelAllCmOpenConditionalOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiCancelAllCmOpenOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiCancelAllCmOpenOrdersRequest) Symbol(symbol string) ApiCancelAllCmOpenOrdersRequest {
	r.symbol = &symbol
	return r
}

func (r ApiCancelAllCmOpenOrdersRequest) RecvWindow(recvWindow int64) ApiCancelAllCmOpenOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelAllCmOpenOrdersRequest) Execute() (*common.RestApiResponse[models.CancelAllCmOpenOrdersResponse], error) {
	return r.ApiService.CancelAllCmOpenOrdersExecute(r)
}

/*
CancelAllCmOpenOrders Cancel All CM Open Orders(TRADE)
Delete /papi/v1/cm/allOpenOrders

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Cancel-All-CM-Open-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -
@return ApiCancelAllCmOpenOrdersRequest
*/
func (a *TradeAPIService) CancelAllCmOpenOrders(ctx context.Context) ApiCancelAllCmOpenOrdersRequest {
	return ApiCancelAllCmOpenOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CancelAllCmOpenOrdersResponse
func (a *TradeAPIService) CancelAllCmOpenOrdersExecute(r ApiCancelAllCmOpenOrdersRequest) (*common.RestApiResponse[models.CancelAllCmOpenOrdersResponse], error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/papi/v1/cm/allOpenOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CancelAllCmOpenOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiCancelAllUmOpenConditionalOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiCancelAllUmOpenConditionalOrdersRequest) Symbol(symbol string) ApiCancelAllUmOpenConditionalOrdersRequest {
	r.symbol = &symbol
	return r
}

func (r ApiCancelAllUmOpenConditionalOrdersRequest) RecvWindow(recvWindow int64) ApiCancelAllUmOpenConditionalOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelAllUmOpenConditionalOrdersRequest) Execute() (*common.RestApiResponse[models.CancelAllUmOpenConditionalOrdersResponse], error) {
	return r.ApiService.CancelAllUmOpenConditionalOrdersExecute(r)
}

/*
CancelAllUmOpenConditionalOrders Cancel All UM Open Conditional Orders (TRADE)
Delete /papi/v1/um/conditional/allOpenOrders

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Cancel-All-UM-Open-Conditional-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -
@return ApiCancelAllUmOpenConditionalOrdersRequest
*/
func (a *TradeAPIService) CancelAllUmOpenConditionalOrders(ctx context.Context) ApiCancelAllUmOpenConditionalOrdersRequest {
	return ApiCancelAllUmOpenConditionalOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CancelAllUmOpenConditionalOrdersResponse
func (a *TradeAPIService) CancelAllUmOpenConditionalOrdersExecute(r ApiCancelAllUmOpenConditionalOrdersRequest) (*common.RestApiResponse[models.CancelAllUmOpenConditionalOrdersResponse], error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/conditional/allOpenOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CancelAllUmOpenConditionalOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiCancelAllUmOpenOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiCancelAllUmOpenOrdersRequest) Symbol(symbol string) ApiCancelAllUmOpenOrdersRequest {
	r.symbol = &symbol
	return r
}

func (r ApiCancelAllUmOpenOrdersRequest) RecvWindow(recvWindow int64) ApiCancelAllUmOpenOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelAllUmOpenOrdersRequest) Execute() (*common.RestApiResponse[models.CancelAllUmOpenOrdersResponse], error) {
	return r.ApiService.CancelAllUmOpenOrdersExecute(r)
}

/*
CancelAllUmOpenOrders Cancel All UM Open Orders(TRADE)
Delete /papi/v1/um/allOpenOrders

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Cancel-All-UM-Open-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -
@return ApiCancelAllUmOpenOrdersRequest
*/
func (a *TradeAPIService) CancelAllUmOpenOrders(ctx context.Context) ApiCancelAllUmOpenOrdersRequest {
	return ApiCancelAllUmOpenOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CancelAllUmOpenOrdersResponse
func (a *TradeAPIService) CancelAllUmOpenOrdersExecute(r ApiCancelAllUmOpenOrdersRequest) (*common.RestApiResponse[models.CancelAllUmOpenOrdersResponse], error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/allOpenOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CancelAllUmOpenOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiCancelCmConditionalOrderRequest struct {
	ctx                 context.Context
	ApiService          *TradeAPIService
	symbol              *string
	strategyId          *int64
	newClientStrategyId *string
	recvWindow          *int64
}

func (r ApiCancelCmConditionalOrderRequest) Symbol(symbol string) ApiCancelCmConditionalOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiCancelCmConditionalOrderRequest) StrategyId(strategyId int64) ApiCancelCmConditionalOrderRequest {
	r.strategyId = &strategyId
	return r
}

func (r ApiCancelCmConditionalOrderRequest) NewClientStrategyId(newClientStrategyId string) ApiCancelCmConditionalOrderRequest {
	r.newClientStrategyId = &newClientStrategyId
	return r
}

func (r ApiCancelCmConditionalOrderRequest) RecvWindow(recvWindow int64) ApiCancelCmConditionalOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelCmConditionalOrderRequest) Execute() (*common.RestApiResponse[models.CancelCmConditionalOrderResponse], error) {
	return r.ApiService.CancelCmConditionalOrderExecute(r)
}

/*
CancelCmConditionalOrder Cancel CM Conditional Order(TRADE)
Delete /papi/v1/cm/conditional/order

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Cancel-CM-Conditional-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param strategyId -
@param newClientStrategyId -
@param recvWindow -
@return ApiCancelCmConditionalOrderRequest
*/
func (a *TradeAPIService) CancelCmConditionalOrder(ctx context.Context) ApiCancelCmConditionalOrderRequest {
	return ApiCancelCmConditionalOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CancelCmConditionalOrderResponse
func (a *TradeAPIService) CancelCmConditionalOrderExecute(r ApiCancelCmConditionalOrderRequest) (*common.RestApiResponse[models.CancelCmConditionalOrderResponse], error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/papi/v1/cm/conditional/order"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.strategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "strategyId", r.strategyId, "form", "")
	}
	if r.newClientStrategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newClientStrategyId", r.newClientStrategyId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CancelCmConditionalOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiCancelCmOrderRequest struct {
	ctx               context.Context
	ApiService        *TradeAPIService
	symbol            *string
	orderId           *int64
	origClientOrderId *string
	recvWindow        *int64
}

func (r ApiCancelCmOrderRequest) Symbol(symbol string) ApiCancelCmOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiCancelCmOrderRequest) OrderId(orderId int64) ApiCancelCmOrderRequest {
	r.orderId = &orderId
	return r
}

func (r ApiCancelCmOrderRequest) OrigClientOrderId(origClientOrderId string) ApiCancelCmOrderRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

func (r ApiCancelCmOrderRequest) RecvWindow(recvWindow int64) ApiCancelCmOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelCmOrderRequest) Execute() (*common.RestApiResponse[models.CancelCmOrderResponse], error) {
	return r.ApiService.CancelCmOrderExecute(r)
}

/*
CancelCmOrder Cancel CM Order(TRADE)
Delete /papi/v1/cm/order

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Cancel-CM-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param orderId -
@param origClientOrderId -
@param recvWindow -
@return ApiCancelCmOrderRequest
*/
func (a *TradeAPIService) CancelCmOrder(ctx context.Context) ApiCancelCmOrderRequest {
	return ApiCancelCmOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CancelCmOrderResponse
func (a *TradeAPIService) CancelCmOrderExecute(r ApiCancelCmOrderRequest) (*common.RestApiResponse[models.CancelCmOrderResponse], error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/papi/v1/cm/order"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.origClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "origClientOrderId", r.origClientOrderId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CancelCmOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiCancelMarginAccountAllOpenOrdersOnASymbolRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiCancelMarginAccountAllOpenOrdersOnASymbolRequest) Symbol(symbol string) ApiCancelMarginAccountAllOpenOrdersOnASymbolRequest {
	r.symbol = &symbol
	return r
}

func (r ApiCancelMarginAccountAllOpenOrdersOnASymbolRequest) RecvWindow(recvWindow int64) ApiCancelMarginAccountAllOpenOrdersOnASymbolRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelMarginAccountAllOpenOrdersOnASymbolRequest) Execute() (*common.RestApiResponse[models.CancelMarginAccountAllOpenOrdersOnASymbolResponse], error) {
	return r.ApiService.CancelMarginAccountAllOpenOrdersOnASymbolExecute(r)
}

/*
CancelMarginAccountAllOpenOrdersOnASymbol Cancel Margin Account All Open Orders on a Symbol(TRADE)
Delete /papi/v1/margin/allOpenOrders

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Cancel-Margin-Account-All-Open-Orders-on-a-Symbol

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -
@return ApiCancelMarginAccountAllOpenOrdersOnASymbolRequest
*/
func (a *TradeAPIService) CancelMarginAccountAllOpenOrdersOnASymbol(ctx context.Context) ApiCancelMarginAccountAllOpenOrdersOnASymbolRequest {
	return ApiCancelMarginAccountAllOpenOrdersOnASymbolRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CancelMarginAccountAllOpenOrdersOnASymbolResponse
func (a *TradeAPIService) CancelMarginAccountAllOpenOrdersOnASymbolExecute(r ApiCancelMarginAccountAllOpenOrdersOnASymbolRequest) (*common.RestApiResponse[models.CancelMarginAccountAllOpenOrdersOnASymbolResponse], error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/papi/v1/margin/allOpenOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CancelMarginAccountAllOpenOrdersOnASymbolResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiCancelMarginAccountOcoOrdersRequest struct {
	ctx               context.Context
	ApiService        *TradeAPIService
	symbol            *string
	orderListId       *int64
	listClientOrderId *string
	newClientOrderId  *string
	recvWindow        *int64
}

func (r ApiCancelMarginAccountOcoOrdersRequest) Symbol(symbol string) ApiCancelMarginAccountOcoOrdersRequest {
	r.symbol = &symbol
	return r
}

// Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided
func (r ApiCancelMarginAccountOcoOrdersRequest) OrderListId(orderListId int64) ApiCancelMarginAccountOcoOrdersRequest {
	r.orderListId = &orderListId
	return r
}

// Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided
func (r ApiCancelMarginAccountOcoOrdersRequest) ListClientOrderId(listClientOrderId string) ApiCancelMarginAccountOcoOrdersRequest {
	r.listClientOrderId = &listClientOrderId
	return r
}

// Used to uniquely identify this cancel. Automatically generated by default
func (r ApiCancelMarginAccountOcoOrdersRequest) NewClientOrderId(newClientOrderId string) ApiCancelMarginAccountOcoOrdersRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

func (r ApiCancelMarginAccountOcoOrdersRequest) RecvWindow(recvWindow int64) ApiCancelMarginAccountOcoOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelMarginAccountOcoOrdersRequest) Execute() (*common.RestApiResponse[models.CancelMarginAccountOcoOrdersResponse], error) {
	return r.ApiService.CancelMarginAccountOcoOrdersExecute(r)
}

/*
CancelMarginAccountOcoOrders Cancel Margin Account OCO Orders(TRADE)
Delete /papi/v1/margin/orderList

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Cancel-Margin-Account-OCO-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param orderListId -  Either `orderListId` or `listClientOrderId` must be provided
@param listClientOrderId -  Either `orderListId` or `listClientOrderId` must be provided
@param newClientOrderId -  Used to uniquely identify this cancel. Automatically generated by default
@param recvWindow -
@return ApiCancelMarginAccountOcoOrdersRequest
*/
func (a *TradeAPIService) CancelMarginAccountOcoOrders(ctx context.Context) ApiCancelMarginAccountOcoOrdersRequest {
	return ApiCancelMarginAccountOcoOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CancelMarginAccountOcoOrdersResponse
func (a *TradeAPIService) CancelMarginAccountOcoOrdersExecute(r ApiCancelMarginAccountOcoOrdersRequest) (*common.RestApiResponse[models.CancelMarginAccountOcoOrdersResponse], error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/papi/v1/margin/orderList"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.orderListId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderListId", r.orderListId, "form", "")
	}
	if r.listClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "listClientOrderId", r.listClientOrderId, "form", "")
	}
	if r.newClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newClientOrderId", r.newClientOrderId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CancelMarginAccountOcoOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiCancelMarginAccountOrderRequest struct {
	ctx               context.Context
	ApiService        *TradeAPIService
	symbol            *string
	orderId           *int64
	origClientOrderId *string
	newClientOrderId  *string
	recvWindow        *int64
}

func (r ApiCancelMarginAccountOrderRequest) Symbol(symbol string) ApiCancelMarginAccountOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiCancelMarginAccountOrderRequest) OrderId(orderId int64) ApiCancelMarginAccountOrderRequest {
	r.orderId = &orderId
	return r
}

func (r ApiCancelMarginAccountOrderRequest) OrigClientOrderId(origClientOrderId string) ApiCancelMarginAccountOrderRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

// Used to uniquely identify this cancel. Automatically generated by default
func (r ApiCancelMarginAccountOrderRequest) NewClientOrderId(newClientOrderId string) ApiCancelMarginAccountOrderRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

func (r ApiCancelMarginAccountOrderRequest) RecvWindow(recvWindow int64) ApiCancelMarginAccountOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelMarginAccountOrderRequest) Execute() (*common.RestApiResponse[models.CancelMarginAccountOrderResponse], error) {
	return r.ApiService.CancelMarginAccountOrderExecute(r)
}

/*
CancelMarginAccountOrder Cancel Margin Account Order(TRADE)
Delete /papi/v1/margin/order

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Cancel-Margin-Account-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param orderId -
@param origClientOrderId -
@param newClientOrderId -  Used to uniquely identify this cancel. Automatically generated by default
@param recvWindow -
@return ApiCancelMarginAccountOrderRequest
*/
func (a *TradeAPIService) CancelMarginAccountOrder(ctx context.Context) ApiCancelMarginAccountOrderRequest {
	return ApiCancelMarginAccountOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CancelMarginAccountOrderResponse
func (a *TradeAPIService) CancelMarginAccountOrderExecute(r ApiCancelMarginAccountOrderRequest) (*common.RestApiResponse[models.CancelMarginAccountOrderResponse], error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/papi/v1/margin/order"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.origClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "origClientOrderId", r.origClientOrderId, "form", "")
	}
	if r.newClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newClientOrderId", r.newClientOrderId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CancelMarginAccountOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiCancelUmConditionalOrderRequest struct {
	ctx                 context.Context
	ApiService          *TradeAPIService
	symbol              *string
	strategyId          *int64
	newClientStrategyId *string
	recvWindow          *int64
}

func (r ApiCancelUmConditionalOrderRequest) Symbol(symbol string) ApiCancelUmConditionalOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiCancelUmConditionalOrderRequest) StrategyId(strategyId int64) ApiCancelUmConditionalOrderRequest {
	r.strategyId = &strategyId
	return r
}

func (r ApiCancelUmConditionalOrderRequest) NewClientStrategyId(newClientStrategyId string) ApiCancelUmConditionalOrderRequest {
	r.newClientStrategyId = &newClientStrategyId
	return r
}

func (r ApiCancelUmConditionalOrderRequest) RecvWindow(recvWindow int64) ApiCancelUmConditionalOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelUmConditionalOrderRequest) Execute() (*common.RestApiResponse[models.CancelUmConditionalOrderResponse], error) {
	return r.ApiService.CancelUmConditionalOrderExecute(r)
}

/*
CancelUmConditionalOrder Cancel UM Conditional Order(TRADE)
Delete /papi/v1/um/conditional/order

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Cancel-UM-Conditional-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param strategyId -
@param newClientStrategyId -
@param recvWindow -
@return ApiCancelUmConditionalOrderRequest
*/
func (a *TradeAPIService) CancelUmConditionalOrder(ctx context.Context) ApiCancelUmConditionalOrderRequest {
	return ApiCancelUmConditionalOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CancelUmConditionalOrderResponse
func (a *TradeAPIService) CancelUmConditionalOrderExecute(r ApiCancelUmConditionalOrderRequest) (*common.RestApiResponse[models.CancelUmConditionalOrderResponse], error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/conditional/order"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.strategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "strategyId", r.strategyId, "form", "")
	}
	if r.newClientStrategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newClientStrategyId", r.newClientStrategyId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CancelUmConditionalOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiCancelUmOrderRequest struct {
	ctx               context.Context
	ApiService        *TradeAPIService
	symbol            *string
	orderId           *int64
	origClientOrderId *string
	recvWindow        *int64
}

func (r ApiCancelUmOrderRequest) Symbol(symbol string) ApiCancelUmOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiCancelUmOrderRequest) OrderId(orderId int64) ApiCancelUmOrderRequest {
	r.orderId = &orderId
	return r
}

func (r ApiCancelUmOrderRequest) OrigClientOrderId(origClientOrderId string) ApiCancelUmOrderRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

func (r ApiCancelUmOrderRequest) RecvWindow(recvWindow int64) ApiCancelUmOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelUmOrderRequest) Execute() (*common.RestApiResponse[models.CancelUmOrderResponse], error) {
	return r.ApiService.CancelUmOrderExecute(r)
}

/*
CancelUmOrder Cancel UM Order(TRADE)
Delete /papi/v1/um/order

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Cancel-UM-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param orderId -
@param origClientOrderId -
@param recvWindow -
@return ApiCancelUmOrderRequest
*/
func (a *TradeAPIService) CancelUmOrder(ctx context.Context) ApiCancelUmOrderRequest {
	return ApiCancelUmOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CancelUmOrderResponse
func (a *TradeAPIService) CancelUmOrderExecute(r ApiCancelUmOrderRequest) (*common.RestApiResponse[models.CancelUmOrderResponse], error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/order"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.origClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "origClientOrderId", r.origClientOrderId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CancelUmOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiCmAccountTradeListRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	pair       *string
	startTime  *int64
	endTime    *int64
	fromId     *int64
	limit      *int64
	recvWindow *int64
}

func (r ApiCmAccountTradeListRequest) Symbol(symbol string) ApiCmAccountTradeListRequest {
	r.symbol = &symbol
	return r
}

func (r ApiCmAccountTradeListRequest) Pair(pair string) ApiCmAccountTradeListRequest {
	r.pair = &pair
	return r
}

// Timestamp in ms to get funding from INCLUSIVE.
func (r ApiCmAccountTradeListRequest) StartTime(startTime int64) ApiCmAccountTradeListRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms to get funding until INCLUSIVE.
func (r ApiCmAccountTradeListRequest) EndTime(endTime int64) ApiCmAccountTradeListRequest {
	r.endTime = &endTime
	return r
}

// Trade id to fetch from. Default gets most recent trades.
func (r ApiCmAccountTradeListRequest) FromId(fromId int64) ApiCmAccountTradeListRequest {
	r.fromId = &fromId
	return r
}

// Default 100; max 1000
func (r ApiCmAccountTradeListRequest) Limit(limit int64) ApiCmAccountTradeListRequest {
	r.limit = &limit
	return r
}

func (r ApiCmAccountTradeListRequest) RecvWindow(recvWindow int64) ApiCmAccountTradeListRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCmAccountTradeListRequest) Execute() (*common.RestApiResponse[models.CmAccountTradeListResponse], error) {
	return r.ApiService.CmAccountTradeListExecute(r)
}

/*
CmAccountTradeList CM Account Trade List(USER_DATA)
Get /papi/v1/cm/userTrades

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/CM-Account-Trade-List

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param pair -
@param startTime -  Timestamp in ms to get funding from INCLUSIVE.
@param endTime -  Timestamp in ms to get funding until INCLUSIVE.
@param fromId -  Trade id to fetch from. Default gets most recent trades.
@param limit -  Default 100; max 1000
@param recvWindow -
@return ApiCmAccountTradeListRequest
*/
func (a *TradeAPIService) CmAccountTradeList(ctx context.Context) ApiCmAccountTradeListRequest {
	return ApiCmAccountTradeListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CmAccountTradeListResponse
func (a *TradeAPIService) CmAccountTradeListExecute(r ApiCmAccountTradeListRequest) (*common.RestApiResponse[models.CmAccountTradeListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/cm/userTrades"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.pair != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pair", r.pair, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.fromId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromId", r.fromId, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CmAccountTradeListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiCmPositionAdlQuantileEstimationRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiCmPositionAdlQuantileEstimationRequest) Symbol(symbol string) ApiCmPositionAdlQuantileEstimationRequest {
	r.symbol = &symbol
	return r
}

func (r ApiCmPositionAdlQuantileEstimationRequest) RecvWindow(recvWindow int64) ApiCmPositionAdlQuantileEstimationRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCmPositionAdlQuantileEstimationRequest) Execute() (*common.RestApiResponse[models.CmPositionAdlQuantileEstimationResponse], error) {
	return r.ApiService.CmPositionAdlQuantileEstimationExecute(r)
}

/*
CmPositionAdlQuantileEstimation CM Position ADL Quantile Estimation(USER_DATA)
Get /papi/v1/cm/adlQuantile

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/CM-Position-ADL-Quantile-Estimation

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -
@return ApiCmPositionAdlQuantileEstimationRequest
*/
func (a *TradeAPIService) CmPositionAdlQuantileEstimation(ctx context.Context) ApiCmPositionAdlQuantileEstimationRequest {
	return ApiCmPositionAdlQuantileEstimationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CmPositionAdlQuantileEstimationResponse
func (a *TradeAPIService) CmPositionAdlQuantileEstimationExecute(r ApiCmPositionAdlQuantileEstimationRequest) (*common.RestApiResponse[models.CmPositionAdlQuantileEstimationResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/cm/adlQuantile"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CmPositionAdlQuantileEstimationResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetUmFuturesBnbBurnStatusRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	recvWindow *int64
}

func (r ApiGetUmFuturesBnbBurnStatusRequest) RecvWindow(recvWindow int64) ApiGetUmFuturesBnbBurnStatusRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetUmFuturesBnbBurnStatusRequest) Execute() (*common.RestApiResponse[models.GetUmFuturesBnbBurnStatusResponse], error) {
	return r.ApiService.GetUmFuturesBnbBurnStatusExecute(r)
}

/*
GetUmFuturesBnbBurnStatus Get UM Futures BNB Burn Status (USER_DATA)
Get /papi/v1/um/feeBurn

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Get-UM-Futures-BNB-Burn-Status

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiGetUmFuturesBnbBurnStatusRequest
*/
func (a *TradeAPIService) GetUmFuturesBnbBurnStatus(ctx context.Context) ApiGetUmFuturesBnbBurnStatusRequest {
	return ApiGetUmFuturesBnbBurnStatusRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetUmFuturesBnbBurnStatusResponse
func (a *TradeAPIService) GetUmFuturesBnbBurnStatusExecute(r ApiGetUmFuturesBnbBurnStatusRequest) (*common.RestApiResponse[models.GetUmFuturesBnbBurnStatusResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/feeBurn"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetUmFuturesBnbBurnStatusResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiMarginAccountBorrowRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	asset      *string
	amount     *float32
	recvWindow *int64
}

func (r ApiMarginAccountBorrowRequest) Asset(asset string) ApiMarginAccountBorrowRequest {
	r.asset = &asset
	return r
}

func (r ApiMarginAccountBorrowRequest) Amount(amount float32) ApiMarginAccountBorrowRequest {
	r.amount = &amount
	return r
}

func (r ApiMarginAccountBorrowRequest) RecvWindow(recvWindow int64) ApiMarginAccountBorrowRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiMarginAccountBorrowRequest) Execute() (*common.RestApiResponse[models.MarginAccountBorrowResponse], error) {
	return r.ApiService.MarginAccountBorrowExecute(r)
}

/*
MarginAccountBorrow Margin Account Borrow(MARGIN)
Post /papi/v1/marginLoan

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Margin-Account-Borrow

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -
@param amount -
@param recvWindow -
@return ApiMarginAccountBorrowRequest
*/
func (a *TradeAPIService) MarginAccountBorrow(ctx context.Context) ApiMarginAccountBorrowRequest {
	return ApiMarginAccountBorrowRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return MarginAccountBorrowResponse
func (a *TradeAPIService) MarginAccountBorrowExecute(r ApiMarginAccountBorrowRequest) (*common.RestApiResponse[models.MarginAccountBorrowResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/papi/v1/marginLoan"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset == nil {
		return nil, common.ReportError("asset is required and must be specified")
	}
	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.MarginAccountBorrowResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiMarginAccountNewOcoRequest struct {
	ctx                  context.Context
	ApiService           *TradeAPIService
	symbol               *string
	side                 *models.NewCmConditionalOrderSideParameter
	quantity             *float32
	price                *float32
	stopPrice            *float32
	listClientOrderId    *string
	limitClientOrderId   *string
	limitIcebergQty      *float32
	stopClientOrderId    *string
	stopLimitPrice       *float32
	stopIcebergQty       *float32
	stopLimitTimeInForce *models.MarginAccountNewOcoStopLimitTimeInForceParameter
	newOrderRespType     *models.NewCmOrderNewOrderRespTypeParameter
	sideEffectType       *models.NewMarginOrderSideEffectTypeParameter
	recvWindow           *int64
}

func (r ApiMarginAccountNewOcoRequest) Symbol(symbol string) ApiMarginAccountNewOcoRequest {
	r.symbol = &symbol
	return r
}

func (r ApiMarginAccountNewOcoRequest) Side(side models.NewCmConditionalOrderSideParameter) ApiMarginAccountNewOcoRequest {
	r.side = &side
	return r
}

// Order quantity
func (r ApiMarginAccountNewOcoRequest) Quantity(quantity float32) ApiMarginAccountNewOcoRequest {
	r.quantity = &quantity
	return r
}

func (r ApiMarginAccountNewOcoRequest) Price(price float32) ApiMarginAccountNewOcoRequest {
	r.price = &price
	return r
}

func (r ApiMarginAccountNewOcoRequest) StopPrice(stopPrice float32) ApiMarginAccountNewOcoRequest {
	r.stopPrice = &stopPrice
	return r
}

// Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided
func (r ApiMarginAccountNewOcoRequest) ListClientOrderId(listClientOrderId string) ApiMarginAccountNewOcoRequest {
	r.listClientOrderId = &listClientOrderId
	return r
}

// A unique Id for the limit order
func (r ApiMarginAccountNewOcoRequest) LimitClientOrderId(limitClientOrderId string) ApiMarginAccountNewOcoRequest {
	r.limitClientOrderId = &limitClientOrderId
	return r
}

func (r ApiMarginAccountNewOcoRequest) LimitIcebergQty(limitIcebergQty float32) ApiMarginAccountNewOcoRequest {
	r.limitIcebergQty = &limitIcebergQty
	return r
}

// A unique Id for the stop loss/stop loss limit leg
func (r ApiMarginAccountNewOcoRequest) StopClientOrderId(stopClientOrderId string) ApiMarginAccountNewOcoRequest {
	r.stopClientOrderId = &stopClientOrderId
	return r
}

// If provided, stopLimitTimeInForce is required.
func (r ApiMarginAccountNewOcoRequest) StopLimitPrice(stopLimitPrice float32) ApiMarginAccountNewOcoRequest {
	r.stopLimitPrice = &stopLimitPrice
	return r
}

func (r ApiMarginAccountNewOcoRequest) StopIcebergQty(stopIcebergQty float32) ApiMarginAccountNewOcoRequest {
	r.stopIcebergQty = &stopIcebergQty
	return r
}

// Valid values are &#x60;GTC/FOK/IOC&#x60;
func (r ApiMarginAccountNewOcoRequest) StopLimitTimeInForce(stopLimitTimeInForce models.MarginAccountNewOcoStopLimitTimeInForceParameter) ApiMarginAccountNewOcoRequest {
	r.stopLimitTimeInForce = &stopLimitTimeInForce
	return r
}

// \&quot;ACK\&quot;, \&quot;RESULT\&quot;, default \&quot;ACK\&quot;
func (r ApiMarginAccountNewOcoRequest) NewOrderRespType(newOrderRespType models.NewCmOrderNewOrderRespTypeParameter) ApiMarginAccountNewOcoRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

// NO_SIDE_EFFECT, MARGIN_BUY, AUTO_REPAY; default NO_SIDE_EFFECT.
func (r ApiMarginAccountNewOcoRequest) SideEffectType(sideEffectType models.NewMarginOrderSideEffectTypeParameter) ApiMarginAccountNewOcoRequest {
	r.sideEffectType = &sideEffectType
	return r
}

func (r ApiMarginAccountNewOcoRequest) RecvWindow(recvWindow int64) ApiMarginAccountNewOcoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiMarginAccountNewOcoRequest) Execute() (*common.RestApiResponse[models.MarginAccountNewOcoResponse], error) {
	return r.ApiService.MarginAccountNewOcoExecute(r)
}

/*
MarginAccountNewOco Margin Account New OCO(TRADE)
Post /papi/v1/margin/order/oco

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Margin-Account-New-OCO

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param side -
@param quantity -  Order quantity
@param price -
@param stopPrice -
@param listClientOrderId -  Either `orderListId` or `listClientOrderId` must be provided
@param limitClientOrderId -  A unique Id for the limit order
@param limitIcebergQty -
@param stopClientOrderId -  A unique Id for the stop loss/stop loss limit leg
@param stopLimitPrice -  If provided, stopLimitTimeInForce is required.
@param stopIcebergQty -
@param stopLimitTimeInForce -  Valid values are `GTC/FOK/IOC`
@param newOrderRespType -  \"ACK\", \"RESULT\", default \"ACK\"
@param sideEffectType -  NO_SIDE_EFFECT, MARGIN_BUY, AUTO_REPAY; default NO_SIDE_EFFECT.
@param recvWindow -
@return ApiMarginAccountNewOcoRequest
*/
func (a *TradeAPIService) MarginAccountNewOco(ctx context.Context) ApiMarginAccountNewOcoRequest {
	return ApiMarginAccountNewOcoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return MarginAccountNewOcoResponse
func (a *TradeAPIService) MarginAccountNewOcoExecute(r ApiMarginAccountNewOcoRequest) (*common.RestApiResponse[models.MarginAccountNewOcoResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/papi/v1/margin/order/oco"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.side == nil {
		return nil, common.ReportError("side is required and must be specified")
	}
	if r.quantity == nil {
		return nil, common.ReportError("quantity is required and must be specified")
	}
	if r.price == nil {
		return nil, common.ReportError("price is required and must be specified")
	}
	if r.stopPrice == nil {
		return nil, common.ReportError("stopPrice is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.listClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "listClientOrderId", r.listClientOrderId, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "side", r.side, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quantity", r.quantity, "form", "")
	if r.limitClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limitClientOrderId", r.limitClientOrderId, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "price", r.price, "form", "")
	if r.limitIcebergQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limitIcebergQty", r.limitIcebergQty, "form", "")
	}
	if r.stopClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "stopClientOrderId", r.stopClientOrderId, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "stopPrice", r.stopPrice, "form", "")
	if r.stopLimitPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "stopLimitPrice", r.stopLimitPrice, "form", "")
	}
	if r.stopIcebergQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "stopIcebergQty", r.stopIcebergQty, "form", "")
	}
	if r.stopLimitTimeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "stopLimitTimeInForce", r.stopLimitTimeInForce, "form", "")
	}
	if r.newOrderRespType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newOrderRespType", r.newOrderRespType, "form", "")
	}
	if r.sideEffectType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "sideEffectType", r.sideEffectType, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.MarginAccountNewOcoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiMarginAccountRepayRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	asset      *string
	amount     *float32
	recvWindow *int64
}

func (r ApiMarginAccountRepayRequest) Asset(asset string) ApiMarginAccountRepayRequest {
	r.asset = &asset
	return r
}

func (r ApiMarginAccountRepayRequest) Amount(amount float32) ApiMarginAccountRepayRequest {
	r.amount = &amount
	return r
}

func (r ApiMarginAccountRepayRequest) RecvWindow(recvWindow int64) ApiMarginAccountRepayRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiMarginAccountRepayRequest) Execute() (*common.RestApiResponse[models.MarginAccountRepayResponse], error) {
	return r.ApiService.MarginAccountRepayExecute(r)
}

/*
MarginAccountRepay Margin Account Repay(MARGIN)
Post /papi/v1/repayLoan

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Margin-Account-Repay

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -
@param amount -
@param recvWindow -
@return ApiMarginAccountRepayRequest
*/
func (a *TradeAPIService) MarginAccountRepay(ctx context.Context) ApiMarginAccountRepayRequest {
	return ApiMarginAccountRepayRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return MarginAccountRepayResponse
func (a *TradeAPIService) MarginAccountRepayExecute(r ApiMarginAccountRepayRequest) (*common.RestApiResponse[models.MarginAccountRepayResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/papi/v1/repayLoan"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset == nil {
		return nil, common.ReportError("asset is required and must be specified")
	}
	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.MarginAccountRepayResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiMarginAccountRepayDebtRequest struct {
	ctx                context.Context
	ApiService         *TradeAPIService
	asset              *string
	amount             *string
	specifyRepayAssets *string
	recvWindow         *int64
}

func (r ApiMarginAccountRepayDebtRequest) Asset(asset string) ApiMarginAccountRepayDebtRequest {
	r.asset = &asset
	return r
}

func (r ApiMarginAccountRepayDebtRequest) Amount(amount string) ApiMarginAccountRepayDebtRequest {
	r.amount = &amount
	return r
}

// Specific asset list to repay debt; Can be added in batch, separated by commas
func (r ApiMarginAccountRepayDebtRequest) SpecifyRepayAssets(specifyRepayAssets string) ApiMarginAccountRepayDebtRequest {
	r.specifyRepayAssets = &specifyRepayAssets
	return r
}

func (r ApiMarginAccountRepayDebtRequest) RecvWindow(recvWindow int64) ApiMarginAccountRepayDebtRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiMarginAccountRepayDebtRequest) Execute() (*common.RestApiResponse[models.MarginAccountRepayDebtResponse], error) {
	return r.ApiService.MarginAccountRepayDebtExecute(r)
}

/*
MarginAccountRepayDebt Margin Account Repay Debt(TRADE)
Post /papi/v1/margin/repay-debt

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Margin-Account-Repay-Debt

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -
@param amount -
@param specifyRepayAssets -  Specific asset list to repay debt; Can be added in batch, separated by commas
@param recvWindow -
@return ApiMarginAccountRepayDebtRequest
*/
func (a *TradeAPIService) MarginAccountRepayDebt(ctx context.Context) ApiMarginAccountRepayDebtRequest {
	return ApiMarginAccountRepayDebtRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return MarginAccountRepayDebtResponse
func (a *TradeAPIService) MarginAccountRepayDebtExecute(r ApiMarginAccountRepayDebtRequest) (*common.RestApiResponse[models.MarginAccountRepayDebtResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/papi/v1/margin/repay-debt"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset == nil {
		return nil, common.ReportError("asset is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	if r.amount != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	}
	if r.specifyRepayAssets != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "specifyRepayAssets", r.specifyRepayAssets, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.MarginAccountRepayDebtResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiMarginAccountTradeListRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	orderId    *int64
	startTime  *int64
	endTime    *int64
	fromId     *int64
	limit      *int64
	recvWindow *int64
}

func (r ApiMarginAccountTradeListRequest) Symbol(symbol string) ApiMarginAccountTradeListRequest {
	r.symbol = &symbol
	return r
}

func (r ApiMarginAccountTradeListRequest) OrderId(orderId int64) ApiMarginAccountTradeListRequest {
	r.orderId = &orderId
	return r
}

// Timestamp in ms to get funding from INCLUSIVE.
func (r ApiMarginAccountTradeListRequest) StartTime(startTime int64) ApiMarginAccountTradeListRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms to get funding until INCLUSIVE.
func (r ApiMarginAccountTradeListRequest) EndTime(endTime int64) ApiMarginAccountTradeListRequest {
	r.endTime = &endTime
	return r
}

// Trade id to fetch from. Default gets most recent trades.
func (r ApiMarginAccountTradeListRequest) FromId(fromId int64) ApiMarginAccountTradeListRequest {
	r.fromId = &fromId
	return r
}

// Default 100; max 1000
func (r ApiMarginAccountTradeListRequest) Limit(limit int64) ApiMarginAccountTradeListRequest {
	r.limit = &limit
	return r
}

func (r ApiMarginAccountTradeListRequest) RecvWindow(recvWindow int64) ApiMarginAccountTradeListRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiMarginAccountTradeListRequest) Execute() (*common.RestApiResponse[models.MarginAccountTradeListResponse], error) {
	return r.ApiService.MarginAccountTradeListExecute(r)
}

/*
MarginAccountTradeList Margin Account Trade List (USER_DATA)
Get /papi/v1/margin/myTrades

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Margin-Account-Trade-List

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param orderId -
@param startTime -  Timestamp in ms to get funding from INCLUSIVE.
@param endTime -  Timestamp in ms to get funding until INCLUSIVE.
@param fromId -  Trade id to fetch from. Default gets most recent trades.
@param limit -  Default 100; max 1000
@param recvWindow -
@return ApiMarginAccountTradeListRequest
*/
func (a *TradeAPIService) MarginAccountTradeList(ctx context.Context) ApiMarginAccountTradeListRequest {
	return ApiMarginAccountTradeListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return MarginAccountTradeListResponse
func (a *TradeAPIService) MarginAccountTradeListExecute(r ApiMarginAccountTradeListRequest) (*common.RestApiResponse[models.MarginAccountTradeListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/margin/myTrades"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.fromId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromId", r.fromId, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.MarginAccountTradeListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiModifyCmOrderRequest struct {
	ctx               context.Context
	ApiService        *TradeAPIService
	symbol            *string
	side              *models.NewCmConditionalOrderSideParameter
	quantity          *float32
	price             *float32
	orderId           *int64
	origClientOrderId *string
	priceMatch        *models.ModifyCmOrderPriceMatchParameter
	recvWindow        *int64
}

func (r ApiModifyCmOrderRequest) Symbol(symbol string) ApiModifyCmOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiModifyCmOrderRequest) Side(side models.NewCmConditionalOrderSideParameter) ApiModifyCmOrderRequest {
	r.side = &side
	return r
}

// Order quantity
func (r ApiModifyCmOrderRequest) Quantity(quantity float32) ApiModifyCmOrderRequest {
	r.quantity = &quantity
	return r
}

func (r ApiModifyCmOrderRequest) Price(price float32) ApiModifyCmOrderRequest {
	r.price = &price
	return r
}

func (r ApiModifyCmOrderRequest) OrderId(orderId int64) ApiModifyCmOrderRequest {
	r.orderId = &orderId
	return r
}

func (r ApiModifyCmOrderRequest) OrigClientOrderId(origClientOrderId string) ApiModifyCmOrderRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

// only avaliable for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; can be set to &#x60;OPPONENT&#x60;/ &#x60;OPPONENT_5&#x60;/ &#x60;OPPONENT_10&#x60;/ &#x60;OPPONENT_20&#x60;: /&#x60;QUEUE&#x60;/ &#x60;QUEUE_5&#x60;/ &#x60;QUEUE_10&#x60;/ &#x60;QUEUE_20&#x60;; Can&#39;t be passed together with &#x60;price&#x60;
func (r ApiModifyCmOrderRequest) PriceMatch(priceMatch models.ModifyCmOrderPriceMatchParameter) ApiModifyCmOrderRequest {
	r.priceMatch = &priceMatch
	return r
}

func (r ApiModifyCmOrderRequest) RecvWindow(recvWindow int64) ApiModifyCmOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiModifyCmOrderRequest) Execute() (*common.RestApiResponse[models.ModifyCmOrderResponse], error) {
	return r.ApiService.ModifyCmOrderExecute(r)
}

/*
ModifyCmOrder Modify CM Order(TRADE)
Put /papi/v1/cm/order

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Modify-CM-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param side -
@param quantity -  Order quantity
@param price -
@param orderId -
@param origClientOrderId -
@param priceMatch -  only avaliable for `LIMIT`/`STOP`/`TAKE_PROFIT` order; can be set to `OPPONENT`/ `OPPONENT_5`/ `OPPONENT_10`/ `OPPONENT_20`: /`QUEUE`/ `QUEUE_5`/ `QUEUE_10`/ `QUEUE_20`; Can't be passed together with `price`
@param recvWindow -
@return ApiModifyCmOrderRequest
*/
func (a *TradeAPIService) ModifyCmOrder(ctx context.Context) ApiModifyCmOrderRequest {
	return ApiModifyCmOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ModifyCmOrderResponse
func (a *TradeAPIService) ModifyCmOrderExecute(r ApiModifyCmOrderRequest) (*common.RestApiResponse[models.ModifyCmOrderResponse], error) {
	localVarHTTPMethod := http.MethodPut
	localVarPath := a.client.cfg.BasePath + "/papi/v1/cm/order"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.side == nil {
		return nil, common.ReportError("side is required and must be specified")
	}
	if r.quantity == nil {
		return nil, common.ReportError("quantity is required and must be specified")
	}
	if r.price == nil {
		return nil, common.ReportError("price is required and must be specified")
	}

	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.origClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "origClientOrderId", r.origClientOrderId, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "side", r.side, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quantity", r.quantity, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "price", r.price, "form", "")
	if r.priceMatch != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "priceMatch", r.priceMatch, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.ModifyCmOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiModifyUmOrderRequest struct {
	ctx               context.Context
	ApiService        *TradeAPIService
	symbol            *string
	side              *models.NewCmConditionalOrderSideParameter
	quantity          *float32
	price             *float32
	orderId           *int64
	origClientOrderId *string
	priceMatch        *models.ModifyCmOrderPriceMatchParameter
	recvWindow        *int64
}

func (r ApiModifyUmOrderRequest) Symbol(symbol string) ApiModifyUmOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiModifyUmOrderRequest) Side(side models.NewCmConditionalOrderSideParameter) ApiModifyUmOrderRequest {
	r.side = &side
	return r
}

// Order quantity
func (r ApiModifyUmOrderRequest) Quantity(quantity float32) ApiModifyUmOrderRequest {
	r.quantity = &quantity
	return r
}

func (r ApiModifyUmOrderRequest) Price(price float32) ApiModifyUmOrderRequest {
	r.price = &price
	return r
}

func (r ApiModifyUmOrderRequest) OrderId(orderId int64) ApiModifyUmOrderRequest {
	r.orderId = &orderId
	return r
}

func (r ApiModifyUmOrderRequest) OrigClientOrderId(origClientOrderId string) ApiModifyUmOrderRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

// only avaliable for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; can be set to &#x60;OPPONENT&#x60;/ &#x60;OPPONENT_5&#x60;/ &#x60;OPPONENT_10&#x60;/ &#x60;OPPONENT_20&#x60;: /&#x60;QUEUE&#x60;/ &#x60;QUEUE_5&#x60;/ &#x60;QUEUE_10&#x60;/ &#x60;QUEUE_20&#x60;; Can&#39;t be passed together with &#x60;price&#x60;
func (r ApiModifyUmOrderRequest) PriceMatch(priceMatch models.ModifyCmOrderPriceMatchParameter) ApiModifyUmOrderRequest {
	r.priceMatch = &priceMatch
	return r
}

func (r ApiModifyUmOrderRequest) RecvWindow(recvWindow int64) ApiModifyUmOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiModifyUmOrderRequest) Execute() (*common.RestApiResponse[models.ModifyUmOrderResponse], error) {
	return r.ApiService.ModifyUmOrderExecute(r)
}

/*
ModifyUmOrder Modify UM Order(TRADE)
Put /papi/v1/um/order

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Modify-UM-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param side -
@param quantity -  Order quantity
@param price -
@param orderId -
@param origClientOrderId -
@param priceMatch -  only avaliable for `LIMIT`/`STOP`/`TAKE_PROFIT` order; can be set to `OPPONENT`/ `OPPONENT_5`/ `OPPONENT_10`/ `OPPONENT_20`: /`QUEUE`/ `QUEUE_5`/ `QUEUE_10`/ `QUEUE_20`; Can't be passed together with `price`
@param recvWindow -
@return ApiModifyUmOrderRequest
*/
func (a *TradeAPIService) ModifyUmOrder(ctx context.Context) ApiModifyUmOrderRequest {
	return ApiModifyUmOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ModifyUmOrderResponse
func (a *TradeAPIService) ModifyUmOrderExecute(r ApiModifyUmOrderRequest) (*common.RestApiResponse[models.ModifyUmOrderResponse], error) {
	localVarHTTPMethod := http.MethodPut
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/order"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.side == nil {
		return nil, common.ReportError("side is required and must be specified")
	}
	if r.quantity == nil {
		return nil, common.ReportError("quantity is required and must be specified")
	}
	if r.price == nil {
		return nil, common.ReportError("price is required and must be specified")
	}

	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.origClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "origClientOrderId", r.origClientOrderId, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "side", r.side, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quantity", r.quantity, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "price", r.price, "form", "")
	if r.priceMatch != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "priceMatch", r.priceMatch, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.ModifyUmOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiNewCmConditionalOrderRequest struct {
	ctx                 context.Context
	ApiService          *TradeAPIService
	symbol              *string
	side                *models.NewCmConditionalOrderSideParameter
	strategyType        *models.NewCmConditionalOrderStrategyTypeParameter
	positionSide        *models.NewCmConditionalOrderPositionSideParameter
	timeInForce         *models.NewCmConditionalOrderTimeInForceParameter
	quantity            *float32
	reduceOnly          *string
	price               *float32
	workingType         *models.NewCmConditionalOrderWorkingTypeParameter
	priceProtect        *string
	newClientStrategyId *string
	stopPrice           *float32
	activationPrice     *float32
	callbackRate        *float32
	recvWindow          *int64
}

func (r ApiNewCmConditionalOrderRequest) Symbol(symbol string) ApiNewCmConditionalOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiNewCmConditionalOrderRequest) Side(side models.NewCmConditionalOrderSideParameter) ApiNewCmConditionalOrderRequest {
	r.side = &side
	return r
}

// \&quot;STOP\&quot;, \&quot;STOP_MARKET\&quot;, \&quot;TAKE_PROFIT\&quot;, \&quot;TAKE_PROFIT_MARKET\&quot;, and \&quot;TRAILING_STOP_MARKET\&quot;
func (r ApiNewCmConditionalOrderRequest) StrategyType(strategyType models.NewCmConditionalOrderStrategyTypeParameter) ApiNewCmConditionalOrderRequest {
	r.strategyType = &strategyType
	return r
}

// Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent in Hedge Mode.
func (r ApiNewCmConditionalOrderRequest) PositionSide(positionSide models.NewCmConditionalOrderPositionSideParameter) ApiNewCmConditionalOrderRequest {
	r.positionSide = &positionSide
	return r
}

func (r ApiNewCmConditionalOrderRequest) TimeInForce(timeInForce models.NewCmConditionalOrderTimeInForceParameter) ApiNewCmConditionalOrderRequest {
	r.timeInForce = &timeInForce
	return r
}

func (r ApiNewCmConditionalOrderRequest) Quantity(quantity float32) ApiNewCmConditionalOrderRequest {
	r.quantity = &quantity
	return r
}

// \&quot;true\&quot; or \&quot;false\&quot;. default \&quot;false\&quot;. Cannot be sent in Hedge Mode .
func (r ApiNewCmConditionalOrderRequest) ReduceOnly(reduceOnly string) ApiNewCmConditionalOrderRequest {
	r.reduceOnly = &reduceOnly
	return r
}

func (r ApiNewCmConditionalOrderRequest) Price(price float32) ApiNewCmConditionalOrderRequest {
	r.price = &price
	return r
}

// stopPrice triggered by: \&quot;MARK_PRICE\&quot;, \&quot;CONTRACT_PRICE\&quot;. Default \&quot;CONTRACT_PRICE\&quot;
func (r ApiNewCmConditionalOrderRequest) WorkingType(workingType models.NewCmConditionalOrderWorkingTypeParameter) ApiNewCmConditionalOrderRequest {
	r.workingType = &workingType
	return r
}

// \&quot;TRUE\&quot; or \&quot;FALSE\&quot;, default \&quot;FALSE\&quot;. Used with &#x60;STOP/STOP_MARKET&#x60; or &#x60;TAKE_PROFIT/TAKE_PROFIT_MARKET&#x60; orders
func (r ApiNewCmConditionalOrderRequest) PriceProtect(priceProtect string) ApiNewCmConditionalOrderRequest {
	r.priceProtect = &priceProtect
	return r
}

func (r ApiNewCmConditionalOrderRequest) NewClientStrategyId(newClientStrategyId string) ApiNewCmConditionalOrderRequest {
	r.newClientStrategyId = &newClientStrategyId
	return r
}

// Used with &#x60;STOP/STOP_MARKET&#x60; or &#x60;TAKE_PROFIT/TAKE_PROFIT_MARKET&#x60; orders.
func (r ApiNewCmConditionalOrderRequest) StopPrice(stopPrice float32) ApiNewCmConditionalOrderRequest {
	r.stopPrice = &stopPrice
	return r
}

// Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, default as the mark price
func (r ApiNewCmConditionalOrderRequest) ActivationPrice(activationPrice float32) ApiNewCmConditionalOrderRequest {
	r.activationPrice = &activationPrice
	return r
}

// Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, min 0.1, max 5 where 1 for 1%
func (r ApiNewCmConditionalOrderRequest) CallbackRate(callbackRate float32) ApiNewCmConditionalOrderRequest {
	r.callbackRate = &callbackRate
	return r
}

func (r ApiNewCmConditionalOrderRequest) RecvWindow(recvWindow int64) ApiNewCmConditionalOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiNewCmConditionalOrderRequest) Execute() (*common.RestApiResponse[models.NewCmConditionalOrderResponse], error) {
	return r.ApiService.NewCmConditionalOrderExecute(r)
}

/*
NewCmConditionalOrder New CM Conditional Order(TRADE)
Post /papi/v1/cm/conditional/order

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/New-CM-Conditional-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param side -
@param strategyType -  \"STOP\", \"STOP_MARKET\", \"TAKE_PROFIT\", \"TAKE_PROFIT_MARKET\", and \"TRAILING_STOP_MARKET\"
@param positionSide -  Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent in Hedge Mode.
@param timeInForce -
@param quantity -
@param reduceOnly -  \"true\" or \"false\". default \"false\". Cannot be sent in Hedge Mode .
@param price -
@param workingType -  stopPrice triggered by: \"MARK_PRICE\", \"CONTRACT_PRICE\". Default \"CONTRACT_PRICE\"
@param priceProtect -  \"TRUE\" or \"FALSE\", default \"FALSE\". Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders
@param newClientStrategyId -
@param stopPrice -  Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders.
@param activationPrice -  Used with `TRAILING_STOP_MARKET` orders, default as the mark price
@param callbackRate -  Used with `TRAILING_STOP_MARKET` orders, min 0.1, max 5 where 1 for 1%
@param recvWindow -
@return ApiNewCmConditionalOrderRequest
*/
func (a *TradeAPIService) NewCmConditionalOrder(ctx context.Context) ApiNewCmConditionalOrderRequest {
	return ApiNewCmConditionalOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return NewCmConditionalOrderResponse
func (a *TradeAPIService) NewCmConditionalOrderExecute(r ApiNewCmConditionalOrderRequest) (*common.RestApiResponse[models.NewCmConditionalOrderResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/papi/v1/cm/conditional/order"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.side == nil {
		return nil, common.ReportError("side is required and must be specified")
	}
	if r.strategyType == nil {
		return nil, common.ReportError("strategyType is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "side", r.side, "form", "")
	if r.positionSide != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "positionSide", r.positionSide, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "strategyType", r.strategyType, "form", "")
	if r.timeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "timeInForce", r.timeInForce, "form", "")
	}
	if r.quantity != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quantity", r.quantity, "form", "")
	}
	if r.reduceOnly != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "reduceOnly", r.reduceOnly, "form", "")
	}
	if r.price != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "price", r.price, "form", "")
	}
	if r.workingType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingType", r.workingType, "form", "")
	}
	if r.priceProtect != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "priceProtect", r.priceProtect, "form", "")
	}
	if r.newClientStrategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newClientStrategyId", r.newClientStrategyId, "form", "")
	}
	if r.stopPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "stopPrice", r.stopPrice, "form", "")
	}
	if r.activationPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "activationPrice", r.activationPrice, "form", "")
	}
	if r.callbackRate != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "callbackRate", r.callbackRate, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.NewCmConditionalOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiNewCmOrderRequest struct {
	ctx              context.Context
	ApiService       *TradeAPIService
	symbol           *string
	side             *models.NewCmConditionalOrderSideParameter
	type_            *models.NewCmOrderTypeParameter
	positionSide     *models.NewCmConditionalOrderPositionSideParameter
	timeInForce      *models.NewCmConditionalOrderTimeInForceParameter
	quantity         *float32
	reduceOnly       *string
	price            *float32
	priceMatch       *models.ModifyCmOrderPriceMatchParameter
	newClientOrderId *string
	newOrderRespType *models.NewCmOrderNewOrderRespTypeParameter
	recvWindow       *int64
}

func (r ApiNewCmOrderRequest) Symbol(symbol string) ApiNewCmOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiNewCmOrderRequest) Side(side models.NewCmConditionalOrderSideParameter) ApiNewCmOrderRequest {
	r.side = &side
	return r
}

// &#x60;LIMIT&#x60;, &#x60;MARKET&#x60;
func (r ApiNewCmOrderRequest) Type(type_ models.NewCmOrderTypeParameter) ApiNewCmOrderRequest {
	r.type_ = &type_
	return r
}

// Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent in Hedge Mode.
func (r ApiNewCmOrderRequest) PositionSide(positionSide models.NewCmConditionalOrderPositionSideParameter) ApiNewCmOrderRequest {
	r.positionSide = &positionSide
	return r
}

func (r ApiNewCmOrderRequest) TimeInForce(timeInForce models.NewCmConditionalOrderTimeInForceParameter) ApiNewCmOrderRequest {
	r.timeInForce = &timeInForce
	return r
}

func (r ApiNewCmOrderRequest) Quantity(quantity float32) ApiNewCmOrderRequest {
	r.quantity = &quantity
	return r
}

// \&quot;true\&quot; or \&quot;false\&quot;. default \&quot;false\&quot;. Cannot be sent in Hedge Mode .
func (r ApiNewCmOrderRequest) ReduceOnly(reduceOnly string) ApiNewCmOrderRequest {
	r.reduceOnly = &reduceOnly
	return r
}

func (r ApiNewCmOrderRequest) Price(price float32) ApiNewCmOrderRequest {
	r.price = &price
	return r
}

// only avaliable for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; can be set to &#x60;OPPONENT&#x60;/ &#x60;OPPONENT_5&#x60;/ &#x60;OPPONENT_10&#x60;/ &#x60;OPPONENT_20&#x60;: /&#x60;QUEUE&#x60;/ &#x60;QUEUE_5&#x60;/ &#x60;QUEUE_10&#x60;/ &#x60;QUEUE_20&#x60;; Can&#39;t be passed together with &#x60;price&#x60;
func (r ApiNewCmOrderRequest) PriceMatch(priceMatch models.ModifyCmOrderPriceMatchParameter) ApiNewCmOrderRequest {
	r.priceMatch = &priceMatch
	return r
}

// Used to uniquely identify this cancel. Automatically generated by default
func (r ApiNewCmOrderRequest) NewClientOrderId(newClientOrderId string) ApiNewCmOrderRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

// \&quot;ACK\&quot;, \&quot;RESULT\&quot;, default \&quot;ACK\&quot;
func (r ApiNewCmOrderRequest) NewOrderRespType(newOrderRespType models.NewCmOrderNewOrderRespTypeParameter) ApiNewCmOrderRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

func (r ApiNewCmOrderRequest) RecvWindow(recvWindow int64) ApiNewCmOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiNewCmOrderRequest) Execute() (*common.RestApiResponse[models.NewCmOrderResponse], error) {
	return r.ApiService.NewCmOrderExecute(r)
}

/*
NewCmOrder New CM Order(TRADE)
Post /papi/v1/cm/order

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/New-CM-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param side -
@param type_ -  `LIMIT`, `MARKET`
@param positionSide -  Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent in Hedge Mode.
@param timeInForce -
@param quantity -
@param reduceOnly -  \"true\" or \"false\". default \"false\". Cannot be sent in Hedge Mode .
@param price -
@param priceMatch -  only avaliable for `LIMIT`/`STOP`/`TAKE_PROFIT` order; can be set to `OPPONENT`/ `OPPONENT_5`/ `OPPONENT_10`/ `OPPONENT_20`: /`QUEUE`/ `QUEUE_5`/ `QUEUE_10`/ `QUEUE_20`; Can't be passed together with `price`
@param newClientOrderId -  Used to uniquely identify this cancel. Automatically generated by default
@param newOrderRespType -  \"ACK\", \"RESULT\", default \"ACK\"
@param recvWindow -
@return ApiNewCmOrderRequest
*/
func (a *TradeAPIService) NewCmOrder(ctx context.Context) ApiNewCmOrderRequest {
	return ApiNewCmOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return NewCmOrderResponse
func (a *TradeAPIService) NewCmOrderExecute(r ApiNewCmOrderRequest) (*common.RestApiResponse[models.NewCmOrderResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/papi/v1/cm/order"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.side == nil {
		return nil, common.ReportError("side is required and must be specified")
	}
	if r.type_ == nil {
		return nil, common.ReportError("type_ is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "side", r.side, "form", "")
	if r.positionSide != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "positionSide", r.positionSide, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	if r.timeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "timeInForce", r.timeInForce, "form", "")
	}
	if r.quantity != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quantity", r.quantity, "form", "")
	}
	if r.reduceOnly != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "reduceOnly", r.reduceOnly, "form", "")
	}
	if r.price != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "price", r.price, "form", "")
	}
	if r.priceMatch != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "priceMatch", r.priceMatch, "form", "")
	}
	if r.newClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newClientOrderId", r.newClientOrderId, "form", "")
	}
	if r.newOrderRespType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newOrderRespType", r.newOrderRespType, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.NewCmOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiNewMarginOrderRequest struct {
	ctx                     context.Context
	ApiService              *TradeAPIService
	symbol                  *string
	side                    *models.NewCmConditionalOrderSideParameter
	type_                   *models.NewCmOrderTypeParameter
	quantity                *float32
	quoteOrderQty           *float32
	price                   *float32
	stopPrice               *float32
	newClientOrderId        *string
	newOrderRespType        *models.NewCmOrderNewOrderRespTypeParameter
	icebergQty              *float32
	sideEffectType          *models.NewMarginOrderSideEffectTypeParameter
	timeInForce             *models.NewCmConditionalOrderTimeInForceParameter
	selfTradePreventionMode *models.NewMarginOrderSelfTradePreventionModeParameter
	autoRepayAtCancel       *bool
	recvWindow              *int64
}

func (r ApiNewMarginOrderRequest) Symbol(symbol string) ApiNewMarginOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiNewMarginOrderRequest) Side(side models.NewCmConditionalOrderSideParameter) ApiNewMarginOrderRequest {
	r.side = &side
	return r
}

// &#x60;LIMIT&#x60;, &#x60;MARKET&#x60;
func (r ApiNewMarginOrderRequest) Type(type_ models.NewCmOrderTypeParameter) ApiNewMarginOrderRequest {
	r.type_ = &type_
	return r
}

func (r ApiNewMarginOrderRequest) Quantity(quantity float32) ApiNewMarginOrderRequest {
	r.quantity = &quantity
	return r
}

func (r ApiNewMarginOrderRequest) QuoteOrderQty(quoteOrderQty float32) ApiNewMarginOrderRequest {
	r.quoteOrderQty = &quoteOrderQty
	return r
}

func (r ApiNewMarginOrderRequest) Price(price float32) ApiNewMarginOrderRequest {
	r.price = &price
	return r
}

// Used with &#x60;STOP/STOP_MARKET&#x60; or &#x60;TAKE_PROFIT/TAKE_PROFIT_MARKET&#x60; orders.
func (r ApiNewMarginOrderRequest) StopPrice(stopPrice float32) ApiNewMarginOrderRequest {
	r.stopPrice = &stopPrice
	return r
}

// Used to uniquely identify this cancel. Automatically generated by default
func (r ApiNewMarginOrderRequest) NewClientOrderId(newClientOrderId string) ApiNewMarginOrderRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

// \&quot;ACK\&quot;, \&quot;RESULT\&quot;, default \&quot;ACK\&quot;
func (r ApiNewMarginOrderRequest) NewOrderRespType(newOrderRespType models.NewCmOrderNewOrderRespTypeParameter) ApiNewMarginOrderRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

// Used with &#x60;LIMIT&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; to create an iceberg order
func (r ApiNewMarginOrderRequest) IcebergQty(icebergQty float32) ApiNewMarginOrderRequest {
	r.icebergQty = &icebergQty
	return r
}

// NO_SIDE_EFFECT, MARGIN_BUY, AUTO_REPAY; default NO_SIDE_EFFECT.
func (r ApiNewMarginOrderRequest) SideEffectType(sideEffectType models.NewMarginOrderSideEffectTypeParameter) ApiNewMarginOrderRequest {
	r.sideEffectType = &sideEffectType
	return r
}

func (r ApiNewMarginOrderRequest) TimeInForce(timeInForce models.NewCmConditionalOrderTimeInForceParameter) ApiNewMarginOrderRequest {
	r.timeInForce = &timeInForce
	return r
}

// &#x60;NONE&#x60;:No STP / &#x60;EXPIRE_TAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_MAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_BOTH&#x60;:expire both orders when STP triggers
func (r ApiNewMarginOrderRequest) SelfTradePreventionMode(selfTradePreventionMode models.NewMarginOrderSelfTradePreventionModeParameter) ApiNewMarginOrderRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// 只有在自动借款单或者自动借还单生效，true表示的是撤单后需要把订单产生的借款归还，默认为true
func (r ApiNewMarginOrderRequest) AutoRepayAtCancel(autoRepayAtCancel bool) ApiNewMarginOrderRequest {
	r.autoRepayAtCancel = &autoRepayAtCancel
	return r
}

func (r ApiNewMarginOrderRequest) RecvWindow(recvWindow int64) ApiNewMarginOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiNewMarginOrderRequest) Execute() (*common.RestApiResponse[models.NewMarginOrderResponse], error) {
	return r.ApiService.NewMarginOrderExecute(r)
}

/*
NewMarginOrder New Margin Order(TRADE)
Post /papi/v1/margin/order

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/New-Margin-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param side -
@param type_ -  `LIMIT`, `MARKET`
@param quantity -
@param quoteOrderQty -
@param price -
@param stopPrice -  Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders.
@param newClientOrderId -  Used to uniquely identify this cancel. Automatically generated by default
@param newOrderRespType -  \"ACK\", \"RESULT\", default \"ACK\"
@param icebergQty -  Used with `LIMIT`, `STOP_LOSS_LIMIT`, and `TAKE_PROFIT_LIMIT` to create an iceberg order
@param sideEffectType -  NO_SIDE_EFFECT, MARGIN_BUY, AUTO_REPAY; default NO_SIDE_EFFECT.
@param timeInForce -
@param selfTradePreventionMode -  `NONE`:No STP / `EXPIRE_TAKER`:expire taker order when STP triggers/ `EXPIRE_MAKER`:expire taker order when STP triggers/ `EXPIRE_BOTH`:expire both orders when STP triggers
@param autoRepayAtCancel -  只有在自动借款单或者自动借还单生效，true表示的是撤单后需要把订单产生的借款归还，默认为true
@param recvWindow -
@return ApiNewMarginOrderRequest
*/
func (a *TradeAPIService) NewMarginOrder(ctx context.Context) ApiNewMarginOrderRequest {
	return ApiNewMarginOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return NewMarginOrderResponse
func (a *TradeAPIService) NewMarginOrderExecute(r ApiNewMarginOrderRequest) (*common.RestApiResponse[models.NewMarginOrderResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/papi/v1/margin/order"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.side == nil {
		return nil, common.ReportError("side is required and must be specified")
	}
	if r.type_ == nil {
		return nil, common.ReportError("type_ is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "side", r.side, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	if r.quantity != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quantity", r.quantity, "form", "")
	}
	if r.quoteOrderQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quoteOrderQty", r.quoteOrderQty, "form", "")
	}
	if r.price != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "price", r.price, "form", "")
	}
	if r.stopPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "stopPrice", r.stopPrice, "form", "")
	}
	if r.newClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newClientOrderId", r.newClientOrderId, "form", "")
	}
	if r.newOrderRespType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newOrderRespType", r.newOrderRespType, "form", "")
	}
	if r.icebergQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "icebergQty", r.icebergQty, "form", "")
	}
	if r.sideEffectType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "sideEffectType", r.sideEffectType, "form", "")
	}
	if r.timeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "timeInForce", r.timeInForce, "form", "")
	}
	if r.selfTradePreventionMode != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "selfTradePreventionMode", r.selfTradePreventionMode, "form", "")
	}
	if r.autoRepayAtCancel != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "autoRepayAtCancel", r.autoRepayAtCancel, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.NewMarginOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiNewUmConditionalOrderRequest struct {
	ctx                     context.Context
	ApiService              *TradeAPIService
	symbol                  *string
	side                    *models.NewCmConditionalOrderSideParameter
	strategyType            *models.NewCmConditionalOrderStrategyTypeParameter
	positionSide            *models.NewCmConditionalOrderPositionSideParameter
	timeInForce             *models.NewCmConditionalOrderTimeInForceParameter
	quantity                *float32
	reduceOnly              *string
	price                   *float32
	workingType             *models.NewCmConditionalOrderWorkingTypeParameter
	priceProtect            *string
	newClientStrategyId     *string
	stopPrice               *float32
	activationPrice         *float32
	callbackRate            *float32
	priceMatch              *models.ModifyCmOrderPriceMatchParameter
	selfTradePreventionMode *models.NewMarginOrderSelfTradePreventionModeParameter
	goodTillDate            *int64
	recvWindow              *int64
}

func (r ApiNewUmConditionalOrderRequest) Symbol(symbol string) ApiNewUmConditionalOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiNewUmConditionalOrderRequest) Side(side models.NewCmConditionalOrderSideParameter) ApiNewUmConditionalOrderRequest {
	r.side = &side
	return r
}

// \&quot;STOP\&quot;, \&quot;STOP_MARKET\&quot;, \&quot;TAKE_PROFIT\&quot;, \&quot;TAKE_PROFIT_MARKET\&quot;, and \&quot;TRAILING_STOP_MARKET\&quot;
func (r ApiNewUmConditionalOrderRequest) StrategyType(strategyType models.NewCmConditionalOrderStrategyTypeParameter) ApiNewUmConditionalOrderRequest {
	r.strategyType = &strategyType
	return r
}

// Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent in Hedge Mode.
func (r ApiNewUmConditionalOrderRequest) PositionSide(positionSide models.NewCmConditionalOrderPositionSideParameter) ApiNewUmConditionalOrderRequest {
	r.positionSide = &positionSide
	return r
}

func (r ApiNewUmConditionalOrderRequest) TimeInForce(timeInForce models.NewCmConditionalOrderTimeInForceParameter) ApiNewUmConditionalOrderRequest {
	r.timeInForce = &timeInForce
	return r
}

func (r ApiNewUmConditionalOrderRequest) Quantity(quantity float32) ApiNewUmConditionalOrderRequest {
	r.quantity = &quantity
	return r
}

// \&quot;true\&quot; or \&quot;false\&quot;. default \&quot;false\&quot;. Cannot be sent in Hedge Mode .
func (r ApiNewUmConditionalOrderRequest) ReduceOnly(reduceOnly string) ApiNewUmConditionalOrderRequest {
	r.reduceOnly = &reduceOnly
	return r
}

func (r ApiNewUmConditionalOrderRequest) Price(price float32) ApiNewUmConditionalOrderRequest {
	r.price = &price
	return r
}

// stopPrice triggered by: \&quot;MARK_PRICE\&quot;, \&quot;CONTRACT_PRICE\&quot;. Default \&quot;CONTRACT_PRICE\&quot;
func (r ApiNewUmConditionalOrderRequest) WorkingType(workingType models.NewCmConditionalOrderWorkingTypeParameter) ApiNewUmConditionalOrderRequest {
	r.workingType = &workingType
	return r
}

// \&quot;TRUE\&quot; or \&quot;FALSE\&quot;, default \&quot;FALSE\&quot;. Used with &#x60;STOP/STOP_MARKET&#x60; or &#x60;TAKE_PROFIT/TAKE_PROFIT_MARKET&#x60; orders
func (r ApiNewUmConditionalOrderRequest) PriceProtect(priceProtect string) ApiNewUmConditionalOrderRequest {
	r.priceProtect = &priceProtect
	return r
}

func (r ApiNewUmConditionalOrderRequest) NewClientStrategyId(newClientStrategyId string) ApiNewUmConditionalOrderRequest {
	r.newClientStrategyId = &newClientStrategyId
	return r
}

// Used with &#x60;STOP/STOP_MARKET&#x60; or &#x60;TAKE_PROFIT/TAKE_PROFIT_MARKET&#x60; orders.
func (r ApiNewUmConditionalOrderRequest) StopPrice(stopPrice float32) ApiNewUmConditionalOrderRequest {
	r.stopPrice = &stopPrice
	return r
}

// Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, default as the mark price
func (r ApiNewUmConditionalOrderRequest) ActivationPrice(activationPrice float32) ApiNewUmConditionalOrderRequest {
	r.activationPrice = &activationPrice
	return r
}

// Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, min 0.1, max 5 where 1 for 1%
func (r ApiNewUmConditionalOrderRequest) CallbackRate(callbackRate float32) ApiNewUmConditionalOrderRequest {
	r.callbackRate = &callbackRate
	return r
}

// only avaliable for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; can be set to &#x60;OPPONENT&#x60;/ &#x60;OPPONENT_5&#x60;/ &#x60;OPPONENT_10&#x60;/ &#x60;OPPONENT_20&#x60;: /&#x60;QUEUE&#x60;/ &#x60;QUEUE_5&#x60;/ &#x60;QUEUE_10&#x60;/ &#x60;QUEUE_20&#x60;; Can&#39;t be passed together with &#x60;price&#x60;
func (r ApiNewUmConditionalOrderRequest) PriceMatch(priceMatch models.ModifyCmOrderPriceMatchParameter) ApiNewUmConditionalOrderRequest {
	r.priceMatch = &priceMatch
	return r
}

// &#x60;NONE&#x60;:No STP / &#x60;EXPIRE_TAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_MAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_BOTH&#x60;:expire both orders when STP triggers
func (r ApiNewUmConditionalOrderRequest) SelfTradePreventionMode(selfTradePreventionMode models.NewMarginOrderSelfTradePreventionModeParameter) ApiNewUmConditionalOrderRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// order cancel time for timeInForce &#x60;GTD&#x60;, mandatory when &#x60;timeInforce&#x60; set to &#x60;GTD&#x60;; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000Mode. It must be sent in Hedge Mode.
func (r ApiNewUmConditionalOrderRequest) GoodTillDate(goodTillDate int64) ApiNewUmConditionalOrderRequest {
	r.goodTillDate = &goodTillDate
	return r
}

func (r ApiNewUmConditionalOrderRequest) RecvWindow(recvWindow int64) ApiNewUmConditionalOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiNewUmConditionalOrderRequest) Execute() (*common.RestApiResponse[models.NewUmConditionalOrderResponse], error) {
	return r.ApiService.NewUmConditionalOrderExecute(r)
}

/*
NewUmConditionalOrder New UM Conditional Order (TRADE)
Post /papi/v1/um/conditional/order

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/New-UM-Conditional-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param side -
@param strategyType -  \"STOP\", \"STOP_MARKET\", \"TAKE_PROFIT\", \"TAKE_PROFIT_MARKET\", and \"TRAILING_STOP_MARKET\"
@param positionSide -  Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent in Hedge Mode.
@param timeInForce -
@param quantity -
@param reduceOnly -  \"true\" or \"false\". default \"false\". Cannot be sent in Hedge Mode .
@param price -
@param workingType -  stopPrice triggered by: \"MARK_PRICE\", \"CONTRACT_PRICE\". Default \"CONTRACT_PRICE\"
@param priceProtect -  \"TRUE\" or \"FALSE\", default \"FALSE\". Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders
@param newClientStrategyId -
@param stopPrice -  Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders.
@param activationPrice -  Used with `TRAILING_STOP_MARKET` orders, default as the mark price
@param callbackRate -  Used with `TRAILING_STOP_MARKET` orders, min 0.1, max 5 where 1 for 1%
@param priceMatch -  only avaliable for `LIMIT`/`STOP`/`TAKE_PROFIT` order; can be set to `OPPONENT`/ `OPPONENT_5`/ `OPPONENT_10`/ `OPPONENT_20`: /`QUEUE`/ `QUEUE_5`/ `QUEUE_10`/ `QUEUE_20`; Can't be passed together with `price`
@param selfTradePreventionMode -  `NONE`:No STP / `EXPIRE_TAKER`:expire taker order when STP triggers/ `EXPIRE_MAKER`:expire taker order when STP triggers/ `EXPIRE_BOTH`:expire both orders when STP triggers
@param goodTillDate -  order cancel time for timeInForce `GTD`, mandatory when `timeInforce` set to `GTD`; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000Mode. It must be sent in Hedge Mode.
@param recvWindow -
@return ApiNewUmConditionalOrderRequest
*/
func (a *TradeAPIService) NewUmConditionalOrder(ctx context.Context) ApiNewUmConditionalOrderRequest {
	return ApiNewUmConditionalOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return NewUmConditionalOrderResponse
func (a *TradeAPIService) NewUmConditionalOrderExecute(r ApiNewUmConditionalOrderRequest) (*common.RestApiResponse[models.NewUmConditionalOrderResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/conditional/order"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.side == nil {
		return nil, common.ReportError("side is required and must be specified")
	}
	if r.strategyType == nil {
		return nil, common.ReportError("strategyType is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "side", r.side, "form", "")
	if r.positionSide != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "positionSide", r.positionSide, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "strategyType", r.strategyType, "form", "")
	if r.timeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "timeInForce", r.timeInForce, "form", "")
	}
	if r.quantity != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quantity", r.quantity, "form", "")
	}
	if r.reduceOnly != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "reduceOnly", r.reduceOnly, "form", "")
	}
	if r.price != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "price", r.price, "form", "")
	}
	if r.workingType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingType", r.workingType, "form", "")
	}
	if r.priceProtect != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "priceProtect", r.priceProtect, "form", "")
	}
	if r.newClientStrategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newClientStrategyId", r.newClientStrategyId, "form", "")
	}
	if r.stopPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "stopPrice", r.stopPrice, "form", "")
	}
	if r.activationPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "activationPrice", r.activationPrice, "form", "")
	}
	if r.callbackRate != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "callbackRate", r.callbackRate, "form", "")
	}
	if r.priceMatch != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "priceMatch", r.priceMatch, "form", "")
	}
	if r.selfTradePreventionMode != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "selfTradePreventionMode", r.selfTradePreventionMode, "form", "")
	}
	if r.goodTillDate != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "goodTillDate", r.goodTillDate, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.NewUmConditionalOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiNewUmOrderRequest struct {
	ctx                     context.Context
	ApiService              *TradeAPIService
	symbol                  *string
	side                    *models.NewCmConditionalOrderSideParameter
	type_                   *models.NewCmOrderTypeParameter
	positionSide            *models.NewCmConditionalOrderPositionSideParameter
	timeInForce             *models.NewCmConditionalOrderTimeInForceParameter
	quantity                *float32
	reduceOnly              *string
	price                   *float32
	newClientOrderId        *string
	newOrderRespType        *models.NewCmOrderNewOrderRespTypeParameter
	priceMatch              *models.ModifyCmOrderPriceMatchParameter
	selfTradePreventionMode *models.NewMarginOrderSelfTradePreventionModeParameter
	goodTillDate            *int64
	recvWindow              *int64
}

func (r ApiNewUmOrderRequest) Symbol(symbol string) ApiNewUmOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiNewUmOrderRequest) Side(side models.NewCmConditionalOrderSideParameter) ApiNewUmOrderRequest {
	r.side = &side
	return r
}

// &#x60;LIMIT&#x60;, &#x60;MARKET&#x60;
func (r ApiNewUmOrderRequest) Type(type_ models.NewCmOrderTypeParameter) ApiNewUmOrderRequest {
	r.type_ = &type_
	return r
}

// Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent in Hedge Mode.
func (r ApiNewUmOrderRequest) PositionSide(positionSide models.NewCmConditionalOrderPositionSideParameter) ApiNewUmOrderRequest {
	r.positionSide = &positionSide
	return r
}

func (r ApiNewUmOrderRequest) TimeInForce(timeInForce models.NewCmConditionalOrderTimeInForceParameter) ApiNewUmOrderRequest {
	r.timeInForce = &timeInForce
	return r
}

func (r ApiNewUmOrderRequest) Quantity(quantity float32) ApiNewUmOrderRequest {
	r.quantity = &quantity
	return r
}

// \&quot;true\&quot; or \&quot;false\&quot;. default \&quot;false\&quot;. Cannot be sent in Hedge Mode .
func (r ApiNewUmOrderRequest) ReduceOnly(reduceOnly string) ApiNewUmOrderRequest {
	r.reduceOnly = &reduceOnly
	return r
}

func (r ApiNewUmOrderRequest) Price(price float32) ApiNewUmOrderRequest {
	r.price = &price
	return r
}

// Used to uniquely identify this cancel. Automatically generated by default
func (r ApiNewUmOrderRequest) NewClientOrderId(newClientOrderId string) ApiNewUmOrderRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

// \&quot;ACK\&quot;, \&quot;RESULT\&quot;, default \&quot;ACK\&quot;
func (r ApiNewUmOrderRequest) NewOrderRespType(newOrderRespType models.NewCmOrderNewOrderRespTypeParameter) ApiNewUmOrderRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

// only avaliable for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; can be set to &#x60;OPPONENT&#x60;/ &#x60;OPPONENT_5&#x60;/ &#x60;OPPONENT_10&#x60;/ &#x60;OPPONENT_20&#x60;: /&#x60;QUEUE&#x60;/ &#x60;QUEUE_5&#x60;/ &#x60;QUEUE_10&#x60;/ &#x60;QUEUE_20&#x60;; Can&#39;t be passed together with &#x60;price&#x60;
func (r ApiNewUmOrderRequest) PriceMatch(priceMatch models.ModifyCmOrderPriceMatchParameter) ApiNewUmOrderRequest {
	r.priceMatch = &priceMatch
	return r
}

// &#x60;NONE&#x60;:No STP / &#x60;EXPIRE_TAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_MAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_BOTH&#x60;:expire both orders when STP triggers
func (r ApiNewUmOrderRequest) SelfTradePreventionMode(selfTradePreventionMode models.NewMarginOrderSelfTradePreventionModeParameter) ApiNewUmOrderRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// order cancel time for timeInForce &#x60;GTD&#x60;, mandatory when &#x60;timeInforce&#x60; set to &#x60;GTD&#x60;; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000Mode. It must be sent in Hedge Mode.
func (r ApiNewUmOrderRequest) GoodTillDate(goodTillDate int64) ApiNewUmOrderRequest {
	r.goodTillDate = &goodTillDate
	return r
}

func (r ApiNewUmOrderRequest) RecvWindow(recvWindow int64) ApiNewUmOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiNewUmOrderRequest) Execute() (*common.RestApiResponse[models.NewUmOrderResponse], error) {
	return r.ApiService.NewUmOrderExecute(r)
}

/*
NewUmOrder New UM Order (TRADE)
Post /papi/v1/um/order

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/New-UM-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param side -
@param type_ -  `LIMIT`, `MARKET`
@param positionSide -  Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent in Hedge Mode.
@param timeInForce -
@param quantity -
@param reduceOnly -  \"true\" or \"false\". default \"false\". Cannot be sent in Hedge Mode .
@param price -
@param newClientOrderId -  Used to uniquely identify this cancel. Automatically generated by default
@param newOrderRespType -  \"ACK\", \"RESULT\", default \"ACK\"
@param priceMatch -  only avaliable for `LIMIT`/`STOP`/`TAKE_PROFIT` order; can be set to `OPPONENT`/ `OPPONENT_5`/ `OPPONENT_10`/ `OPPONENT_20`: /`QUEUE`/ `QUEUE_5`/ `QUEUE_10`/ `QUEUE_20`; Can't be passed together with `price`
@param selfTradePreventionMode -  `NONE`:No STP / `EXPIRE_TAKER`:expire taker order when STP triggers/ `EXPIRE_MAKER`:expire taker order when STP triggers/ `EXPIRE_BOTH`:expire both orders when STP triggers
@param goodTillDate -  order cancel time for timeInForce `GTD`, mandatory when `timeInforce` set to `GTD`; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000Mode. It must be sent in Hedge Mode.
@param recvWindow -
@return ApiNewUmOrderRequest
*/
func (a *TradeAPIService) NewUmOrder(ctx context.Context) ApiNewUmOrderRequest {
	return ApiNewUmOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return NewUmOrderResponse
func (a *TradeAPIService) NewUmOrderExecute(r ApiNewUmOrderRequest) (*common.RestApiResponse[models.NewUmOrderResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/order"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.side == nil {
		return nil, common.ReportError("side is required and must be specified")
	}
	if r.type_ == nil {
		return nil, common.ReportError("type_ is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "side", r.side, "form", "")
	if r.positionSide != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "positionSide", r.positionSide, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	if r.timeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "timeInForce", r.timeInForce, "form", "")
	}
	if r.quantity != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quantity", r.quantity, "form", "")
	}
	if r.reduceOnly != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "reduceOnly", r.reduceOnly, "form", "")
	}
	if r.price != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "price", r.price, "form", "")
	}
	if r.newClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newClientOrderId", r.newClientOrderId, "form", "")
	}
	if r.newOrderRespType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newOrderRespType", r.newOrderRespType, "form", "")
	}
	if r.priceMatch != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "priceMatch", r.priceMatch, "form", "")
	}
	if r.selfTradePreventionMode != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "selfTradePreventionMode", r.selfTradePreventionMode, "form", "")
	}
	if r.goodTillDate != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "goodTillDate", r.goodTillDate, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.NewUmOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryAllCmConditionalOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	strategyId *int64
	startTime  *int64
	endTime    *int64
	limit      *int64
	recvWindow *int64
}

func (r ApiQueryAllCmConditionalOrdersRequest) Symbol(symbol string) ApiQueryAllCmConditionalOrdersRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryAllCmConditionalOrdersRequest) StrategyId(strategyId int64) ApiQueryAllCmConditionalOrdersRequest {
	r.strategyId = &strategyId
	return r
}

// Timestamp in ms to get funding from INCLUSIVE.
func (r ApiQueryAllCmConditionalOrdersRequest) StartTime(startTime int64) ApiQueryAllCmConditionalOrdersRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms to get funding until INCLUSIVE.
func (r ApiQueryAllCmConditionalOrdersRequest) EndTime(endTime int64) ApiQueryAllCmConditionalOrdersRequest {
	r.endTime = &endTime
	return r
}

// Default 100; max 1000
func (r ApiQueryAllCmConditionalOrdersRequest) Limit(limit int64) ApiQueryAllCmConditionalOrdersRequest {
	r.limit = &limit
	return r
}

func (r ApiQueryAllCmConditionalOrdersRequest) RecvWindow(recvWindow int64) ApiQueryAllCmConditionalOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryAllCmConditionalOrdersRequest) Execute() (*common.RestApiResponse[models.QueryAllCmConditionalOrdersResponse], error) {
	return r.ApiService.QueryAllCmConditionalOrdersExecute(r)
}

/*
QueryAllCmConditionalOrders Query All CM Conditional Orders(USER_DATA)
Get /papi/v1/cm/conditional/allOrders

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Query-All-CM-Conditional-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param strategyId -
@param startTime -  Timestamp in ms to get funding from INCLUSIVE.
@param endTime -  Timestamp in ms to get funding until INCLUSIVE.
@param limit -  Default 100; max 1000
@param recvWindow -
@return ApiQueryAllCmConditionalOrdersRequest
*/
func (a *TradeAPIService) QueryAllCmConditionalOrders(ctx context.Context) ApiQueryAllCmConditionalOrdersRequest {
	return ApiQueryAllCmConditionalOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryAllCmConditionalOrdersResponse
func (a *TradeAPIService) QueryAllCmConditionalOrdersExecute(r ApiQueryAllCmConditionalOrdersRequest) (*common.RestApiResponse[models.QueryAllCmConditionalOrdersResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/cm/conditional/allOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.strategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "strategyId", r.strategyId, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryAllCmConditionalOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryAllCmOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	pair       *string
	orderId    *int64
	startTime  *int64
	endTime    *int64
	limit      *int64
	recvWindow *int64
}

func (r ApiQueryAllCmOrdersRequest) Symbol(symbol string) ApiQueryAllCmOrdersRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryAllCmOrdersRequest) Pair(pair string) ApiQueryAllCmOrdersRequest {
	r.pair = &pair
	return r
}

func (r ApiQueryAllCmOrdersRequest) OrderId(orderId int64) ApiQueryAllCmOrdersRequest {
	r.orderId = &orderId
	return r
}

// Timestamp in ms to get funding from INCLUSIVE.
func (r ApiQueryAllCmOrdersRequest) StartTime(startTime int64) ApiQueryAllCmOrdersRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms to get funding until INCLUSIVE.
func (r ApiQueryAllCmOrdersRequest) EndTime(endTime int64) ApiQueryAllCmOrdersRequest {
	r.endTime = &endTime
	return r
}

// Default 100; max 1000
func (r ApiQueryAllCmOrdersRequest) Limit(limit int64) ApiQueryAllCmOrdersRequest {
	r.limit = &limit
	return r
}

func (r ApiQueryAllCmOrdersRequest) RecvWindow(recvWindow int64) ApiQueryAllCmOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryAllCmOrdersRequest) Execute() (*common.RestApiResponse[models.QueryAllCmOrdersResponse], error) {
	return r.ApiService.QueryAllCmOrdersExecute(r)
}

/*
QueryAllCmOrders Query All CM Orders (USER_DATA)
Get /papi/v1/cm/allOrders

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Query-All-CM-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param pair -
@param orderId -
@param startTime -  Timestamp in ms to get funding from INCLUSIVE.
@param endTime -  Timestamp in ms to get funding until INCLUSIVE.
@param limit -  Default 100; max 1000
@param recvWindow -
@return ApiQueryAllCmOrdersRequest
*/
func (a *TradeAPIService) QueryAllCmOrders(ctx context.Context) ApiQueryAllCmOrdersRequest {
	return ApiQueryAllCmOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryAllCmOrdersResponse
func (a *TradeAPIService) QueryAllCmOrdersExecute(r ApiQueryAllCmOrdersRequest) (*common.RestApiResponse[models.QueryAllCmOrdersResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/cm/allOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.pair != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pair", r.pair, "form", "")
	}
	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryAllCmOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryAllCurrentCmOpenConditionalOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiQueryAllCurrentCmOpenConditionalOrdersRequest) Symbol(symbol string) ApiQueryAllCurrentCmOpenConditionalOrdersRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryAllCurrentCmOpenConditionalOrdersRequest) RecvWindow(recvWindow int64) ApiQueryAllCurrentCmOpenConditionalOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryAllCurrentCmOpenConditionalOrdersRequest) Execute() (*common.RestApiResponse[models.QueryAllCurrentCmOpenConditionalOrdersResponse], error) {
	return r.ApiService.QueryAllCurrentCmOpenConditionalOrdersExecute(r)
}

/*
QueryAllCurrentCmOpenConditionalOrders Query All Current CM Open Conditional Orders (USER_DATA)
Get /papi/v1/cm/conditional/openOrders

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Query-All-Current-CM-Open-Conditional-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -
@return ApiQueryAllCurrentCmOpenConditionalOrdersRequest
*/
func (a *TradeAPIService) QueryAllCurrentCmOpenConditionalOrders(ctx context.Context) ApiQueryAllCurrentCmOpenConditionalOrdersRequest {
	return ApiQueryAllCurrentCmOpenConditionalOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryAllCurrentCmOpenConditionalOrdersResponse
func (a *TradeAPIService) QueryAllCurrentCmOpenConditionalOrdersExecute(r ApiQueryAllCurrentCmOpenConditionalOrdersRequest) (*common.RestApiResponse[models.QueryAllCurrentCmOpenConditionalOrdersResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/cm/conditional/openOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryAllCurrentCmOpenConditionalOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryAllCurrentCmOpenOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	pair       *string
	recvWindow *int64
}

func (r ApiQueryAllCurrentCmOpenOrdersRequest) Symbol(symbol string) ApiQueryAllCurrentCmOpenOrdersRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryAllCurrentCmOpenOrdersRequest) Pair(pair string) ApiQueryAllCurrentCmOpenOrdersRequest {
	r.pair = &pair
	return r
}

func (r ApiQueryAllCurrentCmOpenOrdersRequest) RecvWindow(recvWindow int64) ApiQueryAllCurrentCmOpenOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryAllCurrentCmOpenOrdersRequest) Execute() (*common.RestApiResponse[models.QueryAllCurrentCmOpenOrdersResponse], error) {
	return r.ApiService.QueryAllCurrentCmOpenOrdersExecute(r)
}

/*
QueryAllCurrentCmOpenOrders Query All Current CM Open Orders(USER_DATA)
Get /papi/v1/cm/openOrders

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Query-All-Current-CM-Open-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param pair -
@param recvWindow -
@return ApiQueryAllCurrentCmOpenOrdersRequest
*/
func (a *TradeAPIService) QueryAllCurrentCmOpenOrders(ctx context.Context) ApiQueryAllCurrentCmOpenOrdersRequest {
	return ApiQueryAllCurrentCmOpenOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryAllCurrentCmOpenOrdersResponse
func (a *TradeAPIService) QueryAllCurrentCmOpenOrdersExecute(r ApiQueryAllCurrentCmOpenOrdersRequest) (*common.RestApiResponse[models.QueryAllCurrentCmOpenOrdersResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/cm/openOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.pair != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pair", r.pair, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryAllCurrentCmOpenOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryAllCurrentUmOpenConditionalOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiQueryAllCurrentUmOpenConditionalOrdersRequest) Symbol(symbol string) ApiQueryAllCurrentUmOpenConditionalOrdersRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryAllCurrentUmOpenConditionalOrdersRequest) RecvWindow(recvWindow int64) ApiQueryAllCurrentUmOpenConditionalOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryAllCurrentUmOpenConditionalOrdersRequest) Execute() (*common.RestApiResponse[models.QueryAllCurrentUmOpenConditionalOrdersResponse], error) {
	return r.ApiService.QueryAllCurrentUmOpenConditionalOrdersExecute(r)
}

/*
QueryAllCurrentUmOpenConditionalOrders Query All Current UM Open Conditional Orders(USER_DATA)
Get /papi/v1/um/conditional/openOrders

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Query-All-Current-UM-Open-Conditional-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -
@return ApiQueryAllCurrentUmOpenConditionalOrdersRequest
*/
func (a *TradeAPIService) QueryAllCurrentUmOpenConditionalOrders(ctx context.Context) ApiQueryAllCurrentUmOpenConditionalOrdersRequest {
	return ApiQueryAllCurrentUmOpenConditionalOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryAllCurrentUmOpenConditionalOrdersResponse
func (a *TradeAPIService) QueryAllCurrentUmOpenConditionalOrdersExecute(r ApiQueryAllCurrentUmOpenConditionalOrdersRequest) (*common.RestApiResponse[models.QueryAllCurrentUmOpenConditionalOrdersResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/conditional/openOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryAllCurrentUmOpenConditionalOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryAllCurrentUmOpenOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiQueryAllCurrentUmOpenOrdersRequest) Symbol(symbol string) ApiQueryAllCurrentUmOpenOrdersRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryAllCurrentUmOpenOrdersRequest) RecvWindow(recvWindow int64) ApiQueryAllCurrentUmOpenOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryAllCurrentUmOpenOrdersRequest) Execute() (*common.RestApiResponse[models.QueryAllCurrentUmOpenOrdersResponse], error) {
	return r.ApiService.QueryAllCurrentUmOpenOrdersExecute(r)
}

/*
QueryAllCurrentUmOpenOrders Query All Current UM Open Orders(USER_DATA)
Get /papi/v1/um/openOrders

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Query-All-Current-UM-Open-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -
@return ApiQueryAllCurrentUmOpenOrdersRequest
*/
func (a *TradeAPIService) QueryAllCurrentUmOpenOrders(ctx context.Context) ApiQueryAllCurrentUmOpenOrdersRequest {
	return ApiQueryAllCurrentUmOpenOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryAllCurrentUmOpenOrdersResponse
func (a *TradeAPIService) QueryAllCurrentUmOpenOrdersExecute(r ApiQueryAllCurrentUmOpenOrdersRequest) (*common.RestApiResponse[models.QueryAllCurrentUmOpenOrdersResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/openOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryAllCurrentUmOpenOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryAllMarginAccountOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	orderId    *int64
	startTime  *int64
	endTime    *int64
	limit      *int64
	recvWindow *int64
}

func (r ApiQueryAllMarginAccountOrdersRequest) Symbol(symbol string) ApiQueryAllMarginAccountOrdersRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryAllMarginAccountOrdersRequest) OrderId(orderId int64) ApiQueryAllMarginAccountOrdersRequest {
	r.orderId = &orderId
	return r
}

// Timestamp in ms to get funding from INCLUSIVE.
func (r ApiQueryAllMarginAccountOrdersRequest) StartTime(startTime int64) ApiQueryAllMarginAccountOrdersRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms to get funding until INCLUSIVE.
func (r ApiQueryAllMarginAccountOrdersRequest) EndTime(endTime int64) ApiQueryAllMarginAccountOrdersRequest {
	r.endTime = &endTime
	return r
}

// Default 100; max 1000
func (r ApiQueryAllMarginAccountOrdersRequest) Limit(limit int64) ApiQueryAllMarginAccountOrdersRequest {
	r.limit = &limit
	return r
}

func (r ApiQueryAllMarginAccountOrdersRequest) RecvWindow(recvWindow int64) ApiQueryAllMarginAccountOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryAllMarginAccountOrdersRequest) Execute() (*common.RestApiResponse[models.QueryAllMarginAccountOrdersResponse], error) {
	return r.ApiService.QueryAllMarginAccountOrdersExecute(r)
}

/*
QueryAllMarginAccountOrders Query All Margin Account Orders (USER_DATA)
Get /papi/v1/margin/allOrders

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Query-All-Margin-Account-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param orderId -
@param startTime -  Timestamp in ms to get funding from INCLUSIVE.
@param endTime -  Timestamp in ms to get funding until INCLUSIVE.
@param limit -  Default 100; max 1000
@param recvWindow -
@return ApiQueryAllMarginAccountOrdersRequest
*/
func (a *TradeAPIService) QueryAllMarginAccountOrders(ctx context.Context) ApiQueryAllMarginAccountOrdersRequest {
	return ApiQueryAllMarginAccountOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryAllMarginAccountOrdersResponse
func (a *TradeAPIService) QueryAllMarginAccountOrdersExecute(r ApiQueryAllMarginAccountOrdersRequest) (*common.RestApiResponse[models.QueryAllMarginAccountOrdersResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/margin/allOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryAllMarginAccountOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryAllUmConditionalOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	strategyId *int64
	startTime  *int64
	endTime    *int64
	limit      *int64
	recvWindow *int64
}

func (r ApiQueryAllUmConditionalOrdersRequest) Symbol(symbol string) ApiQueryAllUmConditionalOrdersRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryAllUmConditionalOrdersRequest) StrategyId(strategyId int64) ApiQueryAllUmConditionalOrdersRequest {
	r.strategyId = &strategyId
	return r
}

// Timestamp in ms to get funding from INCLUSIVE.
func (r ApiQueryAllUmConditionalOrdersRequest) StartTime(startTime int64) ApiQueryAllUmConditionalOrdersRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms to get funding until INCLUSIVE.
func (r ApiQueryAllUmConditionalOrdersRequest) EndTime(endTime int64) ApiQueryAllUmConditionalOrdersRequest {
	r.endTime = &endTime
	return r
}

// Default 100; max 1000
func (r ApiQueryAllUmConditionalOrdersRequest) Limit(limit int64) ApiQueryAllUmConditionalOrdersRequest {
	r.limit = &limit
	return r
}

func (r ApiQueryAllUmConditionalOrdersRequest) RecvWindow(recvWindow int64) ApiQueryAllUmConditionalOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryAllUmConditionalOrdersRequest) Execute() (*common.RestApiResponse[models.QueryAllUmConditionalOrdersResponse], error) {
	return r.ApiService.QueryAllUmConditionalOrdersExecute(r)
}

/*
QueryAllUmConditionalOrders Query All UM Conditional Orders(USER_DATA)
Get /papi/v1/um/conditional/allOrders

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Query-All-UM-Conditional-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param strategyId -
@param startTime -  Timestamp in ms to get funding from INCLUSIVE.
@param endTime -  Timestamp in ms to get funding until INCLUSIVE.
@param limit -  Default 100; max 1000
@param recvWindow -
@return ApiQueryAllUmConditionalOrdersRequest
*/
func (a *TradeAPIService) QueryAllUmConditionalOrders(ctx context.Context) ApiQueryAllUmConditionalOrdersRequest {
	return ApiQueryAllUmConditionalOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryAllUmConditionalOrdersResponse
func (a *TradeAPIService) QueryAllUmConditionalOrdersExecute(r ApiQueryAllUmConditionalOrdersRequest) (*common.RestApiResponse[models.QueryAllUmConditionalOrdersResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/conditional/allOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.strategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "strategyId", r.strategyId, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryAllUmConditionalOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryAllUmOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	orderId    *int64
	startTime  *int64
	endTime    *int64
	limit      *int64
	recvWindow *int64
}

func (r ApiQueryAllUmOrdersRequest) Symbol(symbol string) ApiQueryAllUmOrdersRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryAllUmOrdersRequest) OrderId(orderId int64) ApiQueryAllUmOrdersRequest {
	r.orderId = &orderId
	return r
}

// Timestamp in ms to get funding from INCLUSIVE.
func (r ApiQueryAllUmOrdersRequest) StartTime(startTime int64) ApiQueryAllUmOrdersRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms to get funding until INCLUSIVE.
func (r ApiQueryAllUmOrdersRequest) EndTime(endTime int64) ApiQueryAllUmOrdersRequest {
	r.endTime = &endTime
	return r
}

// Default 100; max 1000
func (r ApiQueryAllUmOrdersRequest) Limit(limit int64) ApiQueryAllUmOrdersRequest {
	r.limit = &limit
	return r
}

func (r ApiQueryAllUmOrdersRequest) RecvWindow(recvWindow int64) ApiQueryAllUmOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryAllUmOrdersRequest) Execute() (*common.RestApiResponse[models.QueryAllUmOrdersResponse], error) {
	return r.ApiService.QueryAllUmOrdersExecute(r)
}

/*
QueryAllUmOrders Query All UM Orders(USER_DATA)
Get /papi/v1/um/allOrders

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Query-All-UM-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param orderId -
@param startTime -  Timestamp in ms to get funding from INCLUSIVE.
@param endTime -  Timestamp in ms to get funding until INCLUSIVE.
@param limit -  Default 100; max 1000
@param recvWindow -
@return ApiQueryAllUmOrdersRequest
*/
func (a *TradeAPIService) QueryAllUmOrders(ctx context.Context) ApiQueryAllUmOrdersRequest {
	return ApiQueryAllUmOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryAllUmOrdersResponse
func (a *TradeAPIService) QueryAllUmOrdersExecute(r ApiQueryAllUmOrdersRequest) (*common.RestApiResponse[models.QueryAllUmOrdersResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/allOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryAllUmOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryCmConditionalOrderHistoryRequest struct {
	ctx                 context.Context
	ApiService          *TradeAPIService
	symbol              *string
	strategyId          *int64
	newClientStrategyId *string
	recvWindow          *int64
}

func (r ApiQueryCmConditionalOrderHistoryRequest) Symbol(symbol string) ApiQueryCmConditionalOrderHistoryRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryCmConditionalOrderHistoryRequest) StrategyId(strategyId int64) ApiQueryCmConditionalOrderHistoryRequest {
	r.strategyId = &strategyId
	return r
}

func (r ApiQueryCmConditionalOrderHistoryRequest) NewClientStrategyId(newClientStrategyId string) ApiQueryCmConditionalOrderHistoryRequest {
	r.newClientStrategyId = &newClientStrategyId
	return r
}

func (r ApiQueryCmConditionalOrderHistoryRequest) RecvWindow(recvWindow int64) ApiQueryCmConditionalOrderHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryCmConditionalOrderHistoryRequest) Execute() (*common.RestApiResponse[models.QueryCmConditionalOrderHistoryResponse], error) {
	return r.ApiService.QueryCmConditionalOrderHistoryExecute(r)
}

/*
QueryCmConditionalOrderHistory Query CM Conditional Order History(USER_DATA)
Get /papi/v1/cm/conditional/orderHistory

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Query-CM-Conditional-Order-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param strategyId -
@param newClientStrategyId -
@param recvWindow -
@return ApiQueryCmConditionalOrderHistoryRequest
*/
func (a *TradeAPIService) QueryCmConditionalOrderHistory(ctx context.Context) ApiQueryCmConditionalOrderHistoryRequest {
	return ApiQueryCmConditionalOrderHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryCmConditionalOrderHistoryResponse
func (a *TradeAPIService) QueryCmConditionalOrderHistoryExecute(r ApiQueryCmConditionalOrderHistoryRequest) (*common.RestApiResponse[models.QueryCmConditionalOrderHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/cm/conditional/orderHistory"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.strategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "strategyId", r.strategyId, "form", "")
	}
	if r.newClientStrategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newClientStrategyId", r.newClientStrategyId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryCmConditionalOrderHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryCmModifyOrderHistoryRequest struct {
	ctx               context.Context
	ApiService        *TradeAPIService
	symbol            *string
	orderId           *int64
	origClientOrderId *string
	startTime         *int64
	endTime           *int64
	limit             *int64
	recvWindow        *int64
}

func (r ApiQueryCmModifyOrderHistoryRequest) Symbol(symbol string) ApiQueryCmModifyOrderHistoryRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryCmModifyOrderHistoryRequest) OrderId(orderId int64) ApiQueryCmModifyOrderHistoryRequest {
	r.orderId = &orderId
	return r
}

func (r ApiQueryCmModifyOrderHistoryRequest) OrigClientOrderId(origClientOrderId string) ApiQueryCmModifyOrderHistoryRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

// Timestamp in ms to get funding from INCLUSIVE.
func (r ApiQueryCmModifyOrderHistoryRequest) StartTime(startTime int64) ApiQueryCmModifyOrderHistoryRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms to get funding until INCLUSIVE.
func (r ApiQueryCmModifyOrderHistoryRequest) EndTime(endTime int64) ApiQueryCmModifyOrderHistoryRequest {
	r.endTime = &endTime
	return r
}

// Default 100; max 1000
func (r ApiQueryCmModifyOrderHistoryRequest) Limit(limit int64) ApiQueryCmModifyOrderHistoryRequest {
	r.limit = &limit
	return r
}

func (r ApiQueryCmModifyOrderHistoryRequest) RecvWindow(recvWindow int64) ApiQueryCmModifyOrderHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryCmModifyOrderHistoryRequest) Execute() (*common.RestApiResponse[models.QueryCmModifyOrderHistoryResponse], error) {
	return r.ApiService.QueryCmModifyOrderHistoryExecute(r)
}

/*
QueryCmModifyOrderHistory Query CM Modify Order History(TRADE)
Get /papi/v1/cm/orderAmendment

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Query-CM-Modify-Order-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param orderId -
@param origClientOrderId -
@param startTime -  Timestamp in ms to get funding from INCLUSIVE.
@param endTime -  Timestamp in ms to get funding until INCLUSIVE.
@param limit -  Default 100; max 1000
@param recvWindow -
@return ApiQueryCmModifyOrderHistoryRequest
*/
func (a *TradeAPIService) QueryCmModifyOrderHistory(ctx context.Context) ApiQueryCmModifyOrderHistoryRequest {
	return ApiQueryCmModifyOrderHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryCmModifyOrderHistoryResponse
func (a *TradeAPIService) QueryCmModifyOrderHistoryExecute(r ApiQueryCmModifyOrderHistoryRequest) (*common.RestApiResponse[models.QueryCmModifyOrderHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/cm/orderAmendment"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.origClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "origClientOrderId", r.origClientOrderId, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryCmModifyOrderHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryCmOrderRequest struct {
	ctx               context.Context
	ApiService        *TradeAPIService
	symbol            *string
	orderId           *int64
	origClientOrderId *string
	recvWindow        *int64
}

func (r ApiQueryCmOrderRequest) Symbol(symbol string) ApiQueryCmOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryCmOrderRequest) OrderId(orderId int64) ApiQueryCmOrderRequest {
	r.orderId = &orderId
	return r
}

func (r ApiQueryCmOrderRequest) OrigClientOrderId(origClientOrderId string) ApiQueryCmOrderRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

func (r ApiQueryCmOrderRequest) RecvWindow(recvWindow int64) ApiQueryCmOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryCmOrderRequest) Execute() (*common.RestApiResponse[models.QueryCmOrderResponse], error) {
	return r.ApiService.QueryCmOrderExecute(r)
}

/*
QueryCmOrder Query CM Order(USER_DATA)
Get /papi/v1/cm/order

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Query-CM-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param orderId -
@param origClientOrderId -
@param recvWindow -
@return ApiQueryCmOrderRequest
*/
func (a *TradeAPIService) QueryCmOrder(ctx context.Context) ApiQueryCmOrderRequest {
	return ApiQueryCmOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryCmOrderResponse
func (a *TradeAPIService) QueryCmOrderExecute(r ApiQueryCmOrderRequest) (*common.RestApiResponse[models.QueryCmOrderResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/cm/order"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.origClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "origClientOrderId", r.origClientOrderId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryCmOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryCurrentCmOpenConditionalOrderRequest struct {
	ctx                 context.Context
	ApiService          *TradeAPIService
	symbol              *string
	strategyId          *int64
	newClientStrategyId *string
	recvWindow          *int64
}

func (r ApiQueryCurrentCmOpenConditionalOrderRequest) Symbol(symbol string) ApiQueryCurrentCmOpenConditionalOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryCurrentCmOpenConditionalOrderRequest) StrategyId(strategyId int64) ApiQueryCurrentCmOpenConditionalOrderRequest {
	r.strategyId = &strategyId
	return r
}

func (r ApiQueryCurrentCmOpenConditionalOrderRequest) NewClientStrategyId(newClientStrategyId string) ApiQueryCurrentCmOpenConditionalOrderRequest {
	r.newClientStrategyId = &newClientStrategyId
	return r
}

func (r ApiQueryCurrentCmOpenConditionalOrderRequest) RecvWindow(recvWindow int64) ApiQueryCurrentCmOpenConditionalOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryCurrentCmOpenConditionalOrderRequest) Execute() (*common.RestApiResponse[models.QueryCurrentCmOpenConditionalOrderResponse], error) {
	return r.ApiService.QueryCurrentCmOpenConditionalOrderExecute(r)
}

/*
QueryCurrentCmOpenConditionalOrder Query Current CM Open Conditional Order(USER_DATA)
Get /papi/v1/cm/conditional/openOrder

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Query-Current-CM-Open-Conditional-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param strategyId -
@param newClientStrategyId -
@param recvWindow -
@return ApiQueryCurrentCmOpenConditionalOrderRequest
*/
func (a *TradeAPIService) QueryCurrentCmOpenConditionalOrder(ctx context.Context) ApiQueryCurrentCmOpenConditionalOrderRequest {
	return ApiQueryCurrentCmOpenConditionalOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryCurrentCmOpenConditionalOrderResponse
func (a *TradeAPIService) QueryCurrentCmOpenConditionalOrderExecute(r ApiQueryCurrentCmOpenConditionalOrderRequest) (*common.RestApiResponse[models.QueryCurrentCmOpenConditionalOrderResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/cm/conditional/openOrder"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.strategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "strategyId", r.strategyId, "form", "")
	}
	if r.newClientStrategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newClientStrategyId", r.newClientStrategyId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryCurrentCmOpenConditionalOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryCurrentCmOpenOrderRequest struct {
	ctx               context.Context
	ApiService        *TradeAPIService
	symbol            *string
	orderId           *int64
	origClientOrderId *string
	recvWindow        *int64
}

func (r ApiQueryCurrentCmOpenOrderRequest) Symbol(symbol string) ApiQueryCurrentCmOpenOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryCurrentCmOpenOrderRequest) OrderId(orderId int64) ApiQueryCurrentCmOpenOrderRequest {
	r.orderId = &orderId
	return r
}

func (r ApiQueryCurrentCmOpenOrderRequest) OrigClientOrderId(origClientOrderId string) ApiQueryCurrentCmOpenOrderRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

func (r ApiQueryCurrentCmOpenOrderRequest) RecvWindow(recvWindow int64) ApiQueryCurrentCmOpenOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryCurrentCmOpenOrderRequest) Execute() (*common.RestApiResponse[models.QueryCurrentCmOpenOrderResponse], error) {
	return r.ApiService.QueryCurrentCmOpenOrderExecute(r)
}

/*
QueryCurrentCmOpenOrder Query Current CM Open Order (USER_DATA)
Get /papi/v1/cm/openOrder

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Query-Current-CM-Open-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param orderId -
@param origClientOrderId -
@param recvWindow -
@return ApiQueryCurrentCmOpenOrderRequest
*/
func (a *TradeAPIService) QueryCurrentCmOpenOrder(ctx context.Context) ApiQueryCurrentCmOpenOrderRequest {
	return ApiQueryCurrentCmOpenOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryCurrentCmOpenOrderResponse
func (a *TradeAPIService) QueryCurrentCmOpenOrderExecute(r ApiQueryCurrentCmOpenOrderRequest) (*common.RestApiResponse[models.QueryCurrentCmOpenOrderResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/cm/openOrder"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.origClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "origClientOrderId", r.origClientOrderId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryCurrentCmOpenOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryCurrentMarginOpenOrderRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiQueryCurrentMarginOpenOrderRequest) Symbol(symbol string) ApiQueryCurrentMarginOpenOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryCurrentMarginOpenOrderRequest) RecvWindow(recvWindow int64) ApiQueryCurrentMarginOpenOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryCurrentMarginOpenOrderRequest) Execute() (*common.RestApiResponse[models.QueryCurrentMarginOpenOrderResponse], error) {
	return r.ApiService.QueryCurrentMarginOpenOrderExecute(r)
}

/*
QueryCurrentMarginOpenOrder Query Current Margin Open Order (USER_DATA)
Get /papi/v1/margin/openOrders

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Query-Current-Margin-Open-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -
@return ApiQueryCurrentMarginOpenOrderRequest
*/
func (a *TradeAPIService) QueryCurrentMarginOpenOrder(ctx context.Context) ApiQueryCurrentMarginOpenOrderRequest {
	return ApiQueryCurrentMarginOpenOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryCurrentMarginOpenOrderResponse
func (a *TradeAPIService) QueryCurrentMarginOpenOrderExecute(r ApiQueryCurrentMarginOpenOrderRequest) (*common.RestApiResponse[models.QueryCurrentMarginOpenOrderResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/margin/openOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryCurrentMarginOpenOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryCurrentUmOpenConditionalOrderRequest struct {
	ctx                 context.Context
	ApiService          *TradeAPIService
	symbol              *string
	strategyId          *int64
	newClientStrategyId *string
	recvWindow          *int64
}

func (r ApiQueryCurrentUmOpenConditionalOrderRequest) Symbol(symbol string) ApiQueryCurrentUmOpenConditionalOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryCurrentUmOpenConditionalOrderRequest) StrategyId(strategyId int64) ApiQueryCurrentUmOpenConditionalOrderRequest {
	r.strategyId = &strategyId
	return r
}

func (r ApiQueryCurrentUmOpenConditionalOrderRequest) NewClientStrategyId(newClientStrategyId string) ApiQueryCurrentUmOpenConditionalOrderRequest {
	r.newClientStrategyId = &newClientStrategyId
	return r
}

func (r ApiQueryCurrentUmOpenConditionalOrderRequest) RecvWindow(recvWindow int64) ApiQueryCurrentUmOpenConditionalOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryCurrentUmOpenConditionalOrderRequest) Execute() (*common.RestApiResponse[models.QueryCurrentUmOpenConditionalOrderResponse], error) {
	return r.ApiService.QueryCurrentUmOpenConditionalOrderExecute(r)
}

/*
QueryCurrentUmOpenConditionalOrder Query Current UM Open Conditional Order(USER_DATA)
Get /papi/v1/um/conditional/openOrder

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Query-Current-UM-Open-Conditional-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param strategyId -
@param newClientStrategyId -
@param recvWindow -
@return ApiQueryCurrentUmOpenConditionalOrderRequest
*/
func (a *TradeAPIService) QueryCurrentUmOpenConditionalOrder(ctx context.Context) ApiQueryCurrentUmOpenConditionalOrderRequest {
	return ApiQueryCurrentUmOpenConditionalOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryCurrentUmOpenConditionalOrderResponse
func (a *TradeAPIService) QueryCurrentUmOpenConditionalOrderExecute(r ApiQueryCurrentUmOpenConditionalOrderRequest) (*common.RestApiResponse[models.QueryCurrentUmOpenConditionalOrderResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/conditional/openOrder"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.strategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "strategyId", r.strategyId, "form", "")
	}
	if r.newClientStrategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newClientStrategyId", r.newClientStrategyId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryCurrentUmOpenConditionalOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryCurrentUmOpenOrderRequest struct {
	ctx               context.Context
	ApiService        *TradeAPIService
	symbol            *string
	orderId           *int64
	origClientOrderId *string
	recvWindow        *int64
}

func (r ApiQueryCurrentUmOpenOrderRequest) Symbol(symbol string) ApiQueryCurrentUmOpenOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryCurrentUmOpenOrderRequest) OrderId(orderId int64) ApiQueryCurrentUmOpenOrderRequest {
	r.orderId = &orderId
	return r
}

func (r ApiQueryCurrentUmOpenOrderRequest) OrigClientOrderId(origClientOrderId string) ApiQueryCurrentUmOpenOrderRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

func (r ApiQueryCurrentUmOpenOrderRequest) RecvWindow(recvWindow int64) ApiQueryCurrentUmOpenOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryCurrentUmOpenOrderRequest) Execute() (*common.RestApiResponse[models.QueryCurrentUmOpenOrderResponse], error) {
	return r.ApiService.QueryCurrentUmOpenOrderExecute(r)
}

/*
QueryCurrentUmOpenOrder Query Current UM Open Order(USER_DATA)
Get /papi/v1/um/openOrder

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Query-Current-UM-Open-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param orderId -
@param origClientOrderId -
@param recvWindow -
@return ApiQueryCurrentUmOpenOrderRequest
*/
func (a *TradeAPIService) QueryCurrentUmOpenOrder(ctx context.Context) ApiQueryCurrentUmOpenOrderRequest {
	return ApiQueryCurrentUmOpenOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryCurrentUmOpenOrderResponse
func (a *TradeAPIService) QueryCurrentUmOpenOrderExecute(r ApiQueryCurrentUmOpenOrderRequest) (*common.RestApiResponse[models.QueryCurrentUmOpenOrderResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/openOrder"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.origClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "origClientOrderId", r.origClientOrderId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryCurrentUmOpenOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryMarginAccountOrderRequest struct {
	ctx               context.Context
	ApiService        *TradeAPIService
	symbol            *string
	orderId           *int64
	origClientOrderId *string
	recvWindow        *int64
}

func (r ApiQueryMarginAccountOrderRequest) Symbol(symbol string) ApiQueryMarginAccountOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryMarginAccountOrderRequest) OrderId(orderId int64) ApiQueryMarginAccountOrderRequest {
	r.orderId = &orderId
	return r
}

func (r ApiQueryMarginAccountOrderRequest) OrigClientOrderId(origClientOrderId string) ApiQueryMarginAccountOrderRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

func (r ApiQueryMarginAccountOrderRequest) RecvWindow(recvWindow int64) ApiQueryMarginAccountOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryMarginAccountOrderRequest) Execute() (*common.RestApiResponse[models.QueryMarginAccountOrderResponse], error) {
	return r.ApiService.QueryMarginAccountOrderExecute(r)
}

/*
QueryMarginAccountOrder Query Margin Account Order (USER_DATA)
Get /papi/v1/margin/order

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Query-Margin-Account-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param orderId -
@param origClientOrderId -
@param recvWindow -
@return ApiQueryMarginAccountOrderRequest
*/
func (a *TradeAPIService) QueryMarginAccountOrder(ctx context.Context) ApiQueryMarginAccountOrderRequest {
	return ApiQueryMarginAccountOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryMarginAccountOrderResponse
func (a *TradeAPIService) QueryMarginAccountOrderExecute(r ApiQueryMarginAccountOrderRequest) (*common.RestApiResponse[models.QueryMarginAccountOrderResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/margin/order"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.origClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "origClientOrderId", r.origClientOrderId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryMarginAccountOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryMarginAccountsAllOcoRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	fromId     *int64
	startTime  *int64
	endTime    *int64
	limit      *int64
	recvWindow *int64
}

// Trade id to fetch from. Default gets most recent trades.
func (r ApiQueryMarginAccountsAllOcoRequest) FromId(fromId int64) ApiQueryMarginAccountsAllOcoRequest {
	r.fromId = &fromId
	return r
}

// Timestamp in ms to get funding from INCLUSIVE.
func (r ApiQueryMarginAccountsAllOcoRequest) StartTime(startTime int64) ApiQueryMarginAccountsAllOcoRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms to get funding until INCLUSIVE.
func (r ApiQueryMarginAccountsAllOcoRequest) EndTime(endTime int64) ApiQueryMarginAccountsAllOcoRequest {
	r.endTime = &endTime
	return r
}

// Default 100; max 1000
func (r ApiQueryMarginAccountsAllOcoRequest) Limit(limit int64) ApiQueryMarginAccountsAllOcoRequest {
	r.limit = &limit
	return r
}

func (r ApiQueryMarginAccountsAllOcoRequest) RecvWindow(recvWindow int64) ApiQueryMarginAccountsAllOcoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryMarginAccountsAllOcoRequest) Execute() (*common.RestApiResponse[models.QueryMarginAccountsAllOcoResponse], error) {
	return r.ApiService.QueryMarginAccountsAllOcoExecute(r)
}

/*
QueryMarginAccountsAllOco Query Margin Account's all OCO (USER_DATA)
Get /papi/v1/margin/allOrderList

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Query-Margin-Account-all-OCO

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param fromId -  Trade id to fetch from. Default gets most recent trades.
@param startTime -  Timestamp in ms to get funding from INCLUSIVE.
@param endTime -  Timestamp in ms to get funding until INCLUSIVE.
@param limit -  Default 100; max 1000
@param recvWindow -
@return ApiQueryMarginAccountsAllOcoRequest
*/
func (a *TradeAPIService) QueryMarginAccountsAllOco(ctx context.Context) ApiQueryMarginAccountsAllOcoRequest {
	return ApiQueryMarginAccountsAllOcoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryMarginAccountsAllOcoResponse
func (a *TradeAPIService) QueryMarginAccountsAllOcoExecute(r ApiQueryMarginAccountsAllOcoRequest) (*common.RestApiResponse[models.QueryMarginAccountsAllOcoResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/margin/allOrderList"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.fromId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromId", r.fromId, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryMarginAccountsAllOcoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryMarginAccountsOcoRequest struct {
	ctx               context.Context
	ApiService        *TradeAPIService
	orderListId       *int64
	origClientOrderId *string
	recvWindow        *int64
}

// Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided
func (r ApiQueryMarginAccountsOcoRequest) OrderListId(orderListId int64) ApiQueryMarginAccountsOcoRequest {
	r.orderListId = &orderListId
	return r
}

func (r ApiQueryMarginAccountsOcoRequest) OrigClientOrderId(origClientOrderId string) ApiQueryMarginAccountsOcoRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

func (r ApiQueryMarginAccountsOcoRequest) RecvWindow(recvWindow int64) ApiQueryMarginAccountsOcoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryMarginAccountsOcoRequest) Execute() (*common.RestApiResponse[models.QueryMarginAccountsOcoResponse], error) {
	return r.ApiService.QueryMarginAccountsOcoExecute(r)
}

/*
QueryMarginAccountsOco Query Margin Account's OCO (USER_DATA)
Get /papi/v1/margin/orderList

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Query-Margin-Account-OCO

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param orderListId -  Either `orderListId` or `listClientOrderId` must be provided
@param origClientOrderId -
@param recvWindow -
@return ApiQueryMarginAccountsOcoRequest
*/
func (a *TradeAPIService) QueryMarginAccountsOco(ctx context.Context) ApiQueryMarginAccountsOcoRequest {
	return ApiQueryMarginAccountsOcoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryMarginAccountsOcoResponse
func (a *TradeAPIService) QueryMarginAccountsOcoExecute(r ApiQueryMarginAccountsOcoRequest) (*common.RestApiResponse[models.QueryMarginAccountsOcoResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/margin/orderList"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.orderListId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderListId", r.orderListId, "form", "")
	}
	if r.origClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "origClientOrderId", r.origClientOrderId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryMarginAccountsOcoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryMarginAccountsOpenOcoRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	recvWindow *int64
}

func (r ApiQueryMarginAccountsOpenOcoRequest) RecvWindow(recvWindow int64) ApiQueryMarginAccountsOpenOcoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryMarginAccountsOpenOcoRequest) Execute() (*common.RestApiResponse[models.QueryMarginAccountsOpenOcoResponse], error) {
	return r.ApiService.QueryMarginAccountsOpenOcoExecute(r)
}

/*
QueryMarginAccountsOpenOco Query Margin Account's Open OCO (USER_DATA)
Get /papi/v1/margin/openOrderList

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Query-Margin-Account-Open-OCO

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiQueryMarginAccountsOpenOcoRequest
*/
func (a *TradeAPIService) QueryMarginAccountsOpenOco(ctx context.Context) ApiQueryMarginAccountsOpenOcoRequest {
	return ApiQueryMarginAccountsOpenOcoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryMarginAccountsOpenOcoResponse
func (a *TradeAPIService) QueryMarginAccountsOpenOcoExecute(r ApiQueryMarginAccountsOpenOcoRequest) (*common.RestApiResponse[models.QueryMarginAccountsOpenOcoResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/margin/openOrderList"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryMarginAccountsOpenOcoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryUmConditionalOrderHistoryRequest struct {
	ctx                 context.Context
	ApiService          *TradeAPIService
	symbol              *string
	strategyId          *int64
	newClientStrategyId *string
	recvWindow          *int64
}

func (r ApiQueryUmConditionalOrderHistoryRequest) Symbol(symbol string) ApiQueryUmConditionalOrderHistoryRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryUmConditionalOrderHistoryRequest) StrategyId(strategyId int64) ApiQueryUmConditionalOrderHistoryRequest {
	r.strategyId = &strategyId
	return r
}

func (r ApiQueryUmConditionalOrderHistoryRequest) NewClientStrategyId(newClientStrategyId string) ApiQueryUmConditionalOrderHistoryRequest {
	r.newClientStrategyId = &newClientStrategyId
	return r
}

func (r ApiQueryUmConditionalOrderHistoryRequest) RecvWindow(recvWindow int64) ApiQueryUmConditionalOrderHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryUmConditionalOrderHistoryRequest) Execute() (*common.RestApiResponse[models.QueryUmConditionalOrderHistoryResponse], error) {
	return r.ApiService.QueryUmConditionalOrderHistoryExecute(r)
}

/*
QueryUmConditionalOrderHistory Query UM Conditional Order History(USER_DATA)
Get /papi/v1/um/conditional/orderHistory

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Query-UM-Conditional-Order-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param strategyId -
@param newClientStrategyId -
@param recvWindow -
@return ApiQueryUmConditionalOrderHistoryRequest
*/
func (a *TradeAPIService) QueryUmConditionalOrderHistory(ctx context.Context) ApiQueryUmConditionalOrderHistoryRequest {
	return ApiQueryUmConditionalOrderHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryUmConditionalOrderHistoryResponse
func (a *TradeAPIService) QueryUmConditionalOrderHistoryExecute(r ApiQueryUmConditionalOrderHistoryRequest) (*common.RestApiResponse[models.QueryUmConditionalOrderHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/conditional/orderHistory"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.strategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "strategyId", r.strategyId, "form", "")
	}
	if r.newClientStrategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newClientStrategyId", r.newClientStrategyId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryUmConditionalOrderHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryUmModifyOrderHistoryRequest struct {
	ctx               context.Context
	ApiService        *TradeAPIService
	symbol            *string
	orderId           *int64
	origClientOrderId *string
	startTime         *int64
	endTime           *int64
	limit             *int64
	recvWindow        *int64
}

func (r ApiQueryUmModifyOrderHistoryRequest) Symbol(symbol string) ApiQueryUmModifyOrderHistoryRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryUmModifyOrderHistoryRequest) OrderId(orderId int64) ApiQueryUmModifyOrderHistoryRequest {
	r.orderId = &orderId
	return r
}

func (r ApiQueryUmModifyOrderHistoryRequest) OrigClientOrderId(origClientOrderId string) ApiQueryUmModifyOrderHistoryRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

// Timestamp in ms to get funding from INCLUSIVE.
func (r ApiQueryUmModifyOrderHistoryRequest) StartTime(startTime int64) ApiQueryUmModifyOrderHistoryRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms to get funding until INCLUSIVE.
func (r ApiQueryUmModifyOrderHistoryRequest) EndTime(endTime int64) ApiQueryUmModifyOrderHistoryRequest {
	r.endTime = &endTime
	return r
}

// Default 100; max 1000
func (r ApiQueryUmModifyOrderHistoryRequest) Limit(limit int64) ApiQueryUmModifyOrderHistoryRequest {
	r.limit = &limit
	return r
}

func (r ApiQueryUmModifyOrderHistoryRequest) RecvWindow(recvWindow int64) ApiQueryUmModifyOrderHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryUmModifyOrderHistoryRequest) Execute() (*common.RestApiResponse[models.QueryUmModifyOrderHistoryResponse], error) {
	return r.ApiService.QueryUmModifyOrderHistoryExecute(r)
}

/*
QueryUmModifyOrderHistory Query UM Modify Order History(TRADE)
Get /papi/v1/um/orderAmendment

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Query-UM-Modify-Order-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param orderId -
@param origClientOrderId -
@param startTime -  Timestamp in ms to get funding from INCLUSIVE.
@param endTime -  Timestamp in ms to get funding until INCLUSIVE.
@param limit -  Default 100; max 1000
@param recvWindow -
@return ApiQueryUmModifyOrderHistoryRequest
*/
func (a *TradeAPIService) QueryUmModifyOrderHistory(ctx context.Context) ApiQueryUmModifyOrderHistoryRequest {
	return ApiQueryUmModifyOrderHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryUmModifyOrderHistoryResponse
func (a *TradeAPIService) QueryUmModifyOrderHistoryExecute(r ApiQueryUmModifyOrderHistoryRequest) (*common.RestApiResponse[models.QueryUmModifyOrderHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/orderAmendment"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.origClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "origClientOrderId", r.origClientOrderId, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryUmModifyOrderHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryUmOrderRequest struct {
	ctx               context.Context
	ApiService        *TradeAPIService
	symbol            *string
	orderId           *int64
	origClientOrderId *string
	recvWindow        *int64
}

func (r ApiQueryUmOrderRequest) Symbol(symbol string) ApiQueryUmOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryUmOrderRequest) OrderId(orderId int64) ApiQueryUmOrderRequest {
	r.orderId = &orderId
	return r
}

func (r ApiQueryUmOrderRequest) OrigClientOrderId(origClientOrderId string) ApiQueryUmOrderRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

func (r ApiQueryUmOrderRequest) RecvWindow(recvWindow int64) ApiQueryUmOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryUmOrderRequest) Execute() (*common.RestApiResponse[models.QueryUmOrderResponse], error) {
	return r.ApiService.QueryUmOrderExecute(r)
}

/*
QueryUmOrder Query UM Order (USER_DATA)
Get /papi/v1/um/order

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Query-UM-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param orderId -
@param origClientOrderId -
@param recvWindow -
@return ApiQueryUmOrderRequest
*/
func (a *TradeAPIService) QueryUmOrder(ctx context.Context) ApiQueryUmOrderRequest {
	return ApiQueryUmOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryUmOrderResponse
func (a *TradeAPIService) QueryUmOrderExecute(r ApiQueryUmOrderRequest) (*common.RestApiResponse[models.QueryUmOrderResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/order"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.origClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "origClientOrderId", r.origClientOrderId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryUmOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryUsersCmForceOrdersRequest struct {
	ctx           context.Context
	ApiService    *TradeAPIService
	symbol        *string
	autoCloseType *models.QueryUsersCmForceOrdersAutoCloseTypeParameter
	startTime     *int64
	endTime       *int64
	limit         *int64
	recvWindow    *int64
}

func (r ApiQueryUsersCmForceOrdersRequest) Symbol(symbol string) ApiQueryUsersCmForceOrdersRequest {
	r.symbol = &symbol
	return r
}

// &#x60;LIQUIDATION&#x60; for liquidation orders, &#x60;ADL&#x60; for ADL orders.
func (r ApiQueryUsersCmForceOrdersRequest) AutoCloseType(autoCloseType models.QueryUsersCmForceOrdersAutoCloseTypeParameter) ApiQueryUsersCmForceOrdersRequest {
	r.autoCloseType = &autoCloseType
	return r
}

// Timestamp in ms to get funding from INCLUSIVE.
func (r ApiQueryUsersCmForceOrdersRequest) StartTime(startTime int64) ApiQueryUsersCmForceOrdersRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms to get funding until INCLUSIVE.
func (r ApiQueryUsersCmForceOrdersRequest) EndTime(endTime int64) ApiQueryUsersCmForceOrdersRequest {
	r.endTime = &endTime
	return r
}

// Default 100; max 1000
func (r ApiQueryUsersCmForceOrdersRequest) Limit(limit int64) ApiQueryUsersCmForceOrdersRequest {
	r.limit = &limit
	return r
}

func (r ApiQueryUsersCmForceOrdersRequest) RecvWindow(recvWindow int64) ApiQueryUsersCmForceOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryUsersCmForceOrdersRequest) Execute() (*common.RestApiResponse[models.QueryUsersCmForceOrdersResponse], error) {
	return r.ApiService.QueryUsersCmForceOrdersExecute(r)
}

/*
QueryUsersCmForceOrders Query User's CM Force Orders(USER_DATA)
Get /papi/v1/cm/forceOrders

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Query-Users-CM-Force-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param autoCloseType -  `LIQUIDATION` for liquidation orders, `ADL` for ADL orders.
@param startTime -  Timestamp in ms to get funding from INCLUSIVE.
@param endTime -  Timestamp in ms to get funding until INCLUSIVE.
@param limit -  Default 100; max 1000
@param recvWindow -
@return ApiQueryUsersCmForceOrdersRequest
*/
func (a *TradeAPIService) QueryUsersCmForceOrders(ctx context.Context) ApiQueryUsersCmForceOrdersRequest {
	return ApiQueryUsersCmForceOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryUsersCmForceOrdersResponse
func (a *TradeAPIService) QueryUsersCmForceOrdersExecute(r ApiQueryUsersCmForceOrdersRequest) (*common.RestApiResponse[models.QueryUsersCmForceOrdersResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/cm/forceOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.autoCloseType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "autoCloseType", r.autoCloseType, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryUsersCmForceOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryUsersMarginForceOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

// Timestamp in ms to get funding from INCLUSIVE.
func (r ApiQueryUsersMarginForceOrdersRequest) StartTime(startTime int64) ApiQueryUsersMarginForceOrdersRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms to get funding until INCLUSIVE.
func (r ApiQueryUsersMarginForceOrdersRequest) EndTime(endTime int64) ApiQueryUsersMarginForceOrdersRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page. Start from 1. Default:1
func (r ApiQueryUsersMarginForceOrdersRequest) Current(current int64) ApiQueryUsersMarginForceOrdersRequest {
	r.current = &current
	return r
}

// Default:10 Max:100
func (r ApiQueryUsersMarginForceOrdersRequest) Size(size int64) ApiQueryUsersMarginForceOrdersRequest {
	r.size = &size
	return r
}

func (r ApiQueryUsersMarginForceOrdersRequest) RecvWindow(recvWindow int64) ApiQueryUsersMarginForceOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryUsersMarginForceOrdersRequest) Execute() (*common.RestApiResponse[models.QueryUsersMarginForceOrdersResponse], error) {
	return r.ApiService.QueryUsersMarginForceOrdersExecute(r)
}

/*
QueryUsersMarginForceOrders Query User's Margin Force Orders(USER_DATA)
Get /papi/v1/margin/forceOrders

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Query-Users-Margin-Force-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -  Timestamp in ms to get funding from INCLUSIVE.
@param endTime -  Timestamp in ms to get funding until INCLUSIVE.
@param current -  Currently querying page. Start from 1. Default:1
@param size -  Default:10 Max:100
@param recvWindow -
@return ApiQueryUsersMarginForceOrdersRequest
*/
func (a *TradeAPIService) QueryUsersMarginForceOrders(ctx context.Context) ApiQueryUsersMarginForceOrdersRequest {
	return ApiQueryUsersMarginForceOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryUsersMarginForceOrdersResponse
func (a *TradeAPIService) QueryUsersMarginForceOrdersExecute(r ApiQueryUsersMarginForceOrdersRequest) (*common.RestApiResponse[models.QueryUsersMarginForceOrdersResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/margin/forceOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.current != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "current", r.current, "form", "")
	}
	if r.size != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryUsersMarginForceOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryUsersUmForceOrdersRequest struct {
	ctx           context.Context
	ApiService    *TradeAPIService
	symbol        *string
	autoCloseType *models.QueryUsersCmForceOrdersAutoCloseTypeParameter
	startTime     *int64
	endTime       *int64
	limit         *int64
	recvWindow    *int64
}

func (r ApiQueryUsersUmForceOrdersRequest) Symbol(symbol string) ApiQueryUsersUmForceOrdersRequest {
	r.symbol = &symbol
	return r
}

// &#x60;LIQUIDATION&#x60; for liquidation orders, &#x60;ADL&#x60; for ADL orders.
func (r ApiQueryUsersUmForceOrdersRequest) AutoCloseType(autoCloseType models.QueryUsersCmForceOrdersAutoCloseTypeParameter) ApiQueryUsersUmForceOrdersRequest {
	r.autoCloseType = &autoCloseType
	return r
}

// Timestamp in ms to get funding from INCLUSIVE.
func (r ApiQueryUsersUmForceOrdersRequest) StartTime(startTime int64) ApiQueryUsersUmForceOrdersRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms to get funding until INCLUSIVE.
func (r ApiQueryUsersUmForceOrdersRequest) EndTime(endTime int64) ApiQueryUsersUmForceOrdersRequest {
	r.endTime = &endTime
	return r
}

// Default 100; max 1000
func (r ApiQueryUsersUmForceOrdersRequest) Limit(limit int64) ApiQueryUsersUmForceOrdersRequest {
	r.limit = &limit
	return r
}

func (r ApiQueryUsersUmForceOrdersRequest) RecvWindow(recvWindow int64) ApiQueryUsersUmForceOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryUsersUmForceOrdersRequest) Execute() (*common.RestApiResponse[models.QueryUsersUmForceOrdersResponse], error) {
	return r.ApiService.QueryUsersUmForceOrdersExecute(r)
}

/*
QueryUsersUmForceOrders Query User's UM Force Orders (USER_DATA)
Get /papi/v1/um/forceOrders

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Query-Users-UM-Force-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param autoCloseType -  `LIQUIDATION` for liquidation orders, `ADL` for ADL orders.
@param startTime -  Timestamp in ms to get funding from INCLUSIVE.
@param endTime -  Timestamp in ms to get funding until INCLUSIVE.
@param limit -  Default 100; max 1000
@param recvWindow -
@return ApiQueryUsersUmForceOrdersRequest
*/
func (a *TradeAPIService) QueryUsersUmForceOrders(ctx context.Context) ApiQueryUsersUmForceOrdersRequest {
	return ApiQueryUsersUmForceOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryUsersUmForceOrdersResponse
func (a *TradeAPIService) QueryUsersUmForceOrdersExecute(r ApiQueryUsersUmForceOrdersRequest) (*common.RestApiResponse[models.QueryUsersUmForceOrdersResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/forceOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.autoCloseType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "autoCloseType", r.autoCloseType, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryUsersUmForceOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiToggleBnbBurnOnUmFuturesTradeRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	feeBurn    *string
	recvWindow *int64
}

// \&quot;true\&quot;: Fee Discount On; \&quot;false\&quot;: Fee Discount Off
func (r ApiToggleBnbBurnOnUmFuturesTradeRequest) FeeBurn(feeBurn string) ApiToggleBnbBurnOnUmFuturesTradeRequest {
	r.feeBurn = &feeBurn
	return r
}

func (r ApiToggleBnbBurnOnUmFuturesTradeRequest) RecvWindow(recvWindow int64) ApiToggleBnbBurnOnUmFuturesTradeRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiToggleBnbBurnOnUmFuturesTradeRequest) Execute() (*common.RestApiResponse[models.ToggleBnbBurnOnUmFuturesTradeResponse], error) {
	return r.ApiService.ToggleBnbBurnOnUmFuturesTradeExecute(r)
}

/*
ToggleBnbBurnOnUmFuturesTrade Toggle BNB Burn On UM Futures Trade (TRADE)
Post /papi/v1/um/feeBurn

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Toggle-BNB-Burn-On-UM-Futures-Trade

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param feeBurn -  \"true\": Fee Discount On; \"false\": Fee Discount Off
@param recvWindow -
@return ApiToggleBnbBurnOnUmFuturesTradeRequest
*/
func (a *TradeAPIService) ToggleBnbBurnOnUmFuturesTrade(ctx context.Context) ApiToggleBnbBurnOnUmFuturesTradeRequest {
	return ApiToggleBnbBurnOnUmFuturesTradeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToggleBnbBurnOnUmFuturesTradeResponse
func (a *TradeAPIService) ToggleBnbBurnOnUmFuturesTradeExecute(r ApiToggleBnbBurnOnUmFuturesTradeRequest) (*common.RestApiResponse[models.ToggleBnbBurnOnUmFuturesTradeResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/feeBurn"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.feeBurn == nil {
		return nil, common.ReportError("feeBurn is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "feeBurn", r.feeBurn, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.ToggleBnbBurnOnUmFuturesTradeResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiUmAccountTradeListRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	startTime  *int64
	endTime    *int64
	fromId     *int64
	limit      *int64
	recvWindow *int64
}

func (r ApiUmAccountTradeListRequest) Symbol(symbol string) ApiUmAccountTradeListRequest {
	r.symbol = &symbol
	return r
}

// Timestamp in ms to get funding from INCLUSIVE.
func (r ApiUmAccountTradeListRequest) StartTime(startTime int64) ApiUmAccountTradeListRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms to get funding until INCLUSIVE.
func (r ApiUmAccountTradeListRequest) EndTime(endTime int64) ApiUmAccountTradeListRequest {
	r.endTime = &endTime
	return r
}

// Trade id to fetch from. Default gets most recent trades.
func (r ApiUmAccountTradeListRequest) FromId(fromId int64) ApiUmAccountTradeListRequest {
	r.fromId = &fromId
	return r
}

// Default 100; max 1000
func (r ApiUmAccountTradeListRequest) Limit(limit int64) ApiUmAccountTradeListRequest {
	r.limit = &limit
	return r
}

func (r ApiUmAccountTradeListRequest) RecvWindow(recvWindow int64) ApiUmAccountTradeListRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiUmAccountTradeListRequest) Execute() (*common.RestApiResponse[models.UmAccountTradeListResponse], error) {
	return r.ApiService.UmAccountTradeListExecute(r)
}

/*
UmAccountTradeList UM Account Trade List(USER_DATA)
Get /papi/v1/um/userTrades

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/UM-Account-Trade-List

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param startTime -  Timestamp in ms to get funding from INCLUSIVE.
@param endTime -  Timestamp in ms to get funding until INCLUSIVE.
@param fromId -  Trade id to fetch from. Default gets most recent trades.
@param limit -  Default 100; max 1000
@param recvWindow -
@return ApiUmAccountTradeListRequest
*/
func (a *TradeAPIService) UmAccountTradeList(ctx context.Context) ApiUmAccountTradeListRequest {
	return ApiUmAccountTradeListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return UmAccountTradeListResponse
func (a *TradeAPIService) UmAccountTradeListExecute(r ApiUmAccountTradeListRequest) (*common.RestApiResponse[models.UmAccountTradeListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/userTrades"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.fromId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromId", r.fromId, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.UmAccountTradeListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiUmPositionAdlQuantileEstimationRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiUmPositionAdlQuantileEstimationRequest) Symbol(symbol string) ApiUmPositionAdlQuantileEstimationRequest {
	r.symbol = &symbol
	return r
}

func (r ApiUmPositionAdlQuantileEstimationRequest) RecvWindow(recvWindow int64) ApiUmPositionAdlQuantileEstimationRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiUmPositionAdlQuantileEstimationRequest) Execute() (*common.RestApiResponse[models.UmPositionAdlQuantileEstimationResponse], error) {
	return r.ApiService.UmPositionAdlQuantileEstimationExecute(r)
}

/*
UmPositionAdlQuantileEstimation UM Position ADL Quantile Estimation(USER_DATA)
Get /papi/v1/um/adlQuantile

https://developers.binance.com/docs/derivatives/portfolio-margin/trade/UM-Position-ADL-Quantile-Estimation

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -
@return ApiUmPositionAdlQuantileEstimationRequest
*/
func (a *TradeAPIService) UmPositionAdlQuantileEstimation(ctx context.Context) ApiUmPositionAdlQuantileEstimationRequest {
	return ApiUmPositionAdlQuantileEstimationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return UmPositionAdlQuantileEstimationResponse
func (a *TradeAPIService) UmPositionAdlQuantileEstimationExecute(r ApiUmPositionAdlQuantileEstimationRequest) (*common.RestApiResponse[models.UmPositionAdlQuantileEstimationResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/adlQuantile"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.UmPositionAdlQuantileEstimationResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
