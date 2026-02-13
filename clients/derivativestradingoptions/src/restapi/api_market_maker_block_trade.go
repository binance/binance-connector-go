/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package binancederivativestradingoptionsrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/derivativestradingoptions/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// MarketMakerBlockTradeAPIService MarketMakerBlockTradeAPI Service
type MarketMakerBlockTradeAPIService Service

type ApiAcceptBlockTradeOrderRequest struct {
	ctx                   context.Context
	ApiService            *MarketMakerBlockTradeAPIService
	blockOrderMatchingKey *string
	recvWindow            *int64
}

func (r ApiAcceptBlockTradeOrderRequest) BlockOrderMatchingKey(blockOrderMatchingKey string) ApiAcceptBlockTradeOrderRequest {
	r.blockOrderMatchingKey = &blockOrderMatchingKey
	return r
}

func (r ApiAcceptBlockTradeOrderRequest) RecvWindow(recvWindow int64) ApiAcceptBlockTradeOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiAcceptBlockTradeOrderRequest) Execute() (*common.RestApiResponse[models.AcceptBlockTradeOrderResponse], error) {
	return r.ApiService.AcceptBlockTradeOrderExecute(r)
}

/*
AcceptBlockTradeOrder Accept Block Trade Order (TRADE)
Post /eapi/v1/block/order/execute

https://developers.binance.com/docs/derivatives/options-trading/market-maker-block-trade/Accept-Block-Trade-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param blockOrderMatchingKey -
@param recvWindow -
@return ApiAcceptBlockTradeOrderRequest
*/
func (a *MarketMakerBlockTradeAPIService) AcceptBlockTradeOrder(ctx context.Context) ApiAcceptBlockTradeOrderRequest {
	return ApiAcceptBlockTradeOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AcceptBlockTradeOrderResponse
func (a *MarketMakerBlockTradeAPIService) AcceptBlockTradeOrderExecute(r ApiAcceptBlockTradeOrderRequest) (*common.RestApiResponse[models.AcceptBlockTradeOrderResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/block/order/execute"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.blockOrderMatchingKey == nil {
		return nil, common.ReportError("blockOrderMatchingKey is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "blockOrderMatchingKey", r.blockOrderMatchingKey, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.AcceptBlockTradeOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiAccountBlockTradeListRequest struct {
	ctx        context.Context
	ApiService *MarketMakerBlockTradeAPIService
	endTime    *int64
	startTime  *int64
	underlying *string
	recvWindow *int64
}

// End Time, e.g 1593512200000
func (r ApiAccountBlockTradeListRequest) EndTime(endTime int64) ApiAccountBlockTradeListRequest {
	r.endTime = &endTime
	return r
}

// Start Time, e.g 1593511200000
func (r ApiAccountBlockTradeListRequest) StartTime(startTime int64) ApiAccountBlockTradeListRequest {
	r.startTime = &startTime
	return r
}

// underlying, e.g BTCUSDT
func (r ApiAccountBlockTradeListRequest) Underlying(underlying string) ApiAccountBlockTradeListRequest {
	r.underlying = &underlying
	return r
}

func (r ApiAccountBlockTradeListRequest) RecvWindow(recvWindow int64) ApiAccountBlockTradeListRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiAccountBlockTradeListRequest) Execute() (*common.RestApiResponse[models.AccountBlockTradeListResponse], error) {
	return r.ApiService.AccountBlockTradeListExecute(r)
}

/*
AccountBlockTradeList Account Block Trade List (USER_DATA)
Get /eapi/v1/block/user-trades

https://developers.binance.com/docs/derivatives/options-trading/market-maker-block-trade/Account-Block-Trade-List

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param endTime -  End Time, e.g 1593512200000
@param startTime -  Start Time, e.g 1593511200000
@param underlying -  underlying, e.g BTCUSDT
@param recvWindow -
@return ApiAccountBlockTradeListRequest
*/
func (a *MarketMakerBlockTradeAPIService) AccountBlockTradeList(ctx context.Context) ApiAccountBlockTradeListRequest {
	return ApiAccountBlockTradeListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AccountBlockTradeListResponse
func (a *MarketMakerBlockTradeAPIService) AccountBlockTradeListExecute(r ApiAccountBlockTradeListRequest) (*common.RestApiResponse[models.AccountBlockTradeListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/block/user-trades"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.underlying != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "underlying", r.underlying, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.AccountBlockTradeListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiCancelBlockTradeOrderRequest struct {
	ctx                   context.Context
	ApiService            *MarketMakerBlockTradeAPIService
	blockOrderMatchingKey *string
	recvWindow            *int64
}

func (r ApiCancelBlockTradeOrderRequest) BlockOrderMatchingKey(blockOrderMatchingKey string) ApiCancelBlockTradeOrderRequest {
	r.blockOrderMatchingKey = &blockOrderMatchingKey
	return r
}

func (r ApiCancelBlockTradeOrderRequest) RecvWindow(recvWindow int64) ApiCancelBlockTradeOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelBlockTradeOrderRequest) Execute() (struct{}, error) {
	return r.ApiService.CancelBlockTradeOrderExecute(r)
}

/*
CancelBlockTradeOrder Cancel Block Trade Order (TRADE)
Delete /eapi/v1/block/order/create

https://developers.binance.com/docs/derivatives/options-trading/market-maker-block-trade/Cancel-Block-Trade-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param blockOrderMatchingKey -
@param recvWindow -
@return ApiCancelBlockTradeOrderRequest
*/
func (a *MarketMakerBlockTradeAPIService) CancelBlockTradeOrder(ctx context.Context) ApiCancelBlockTradeOrderRequest {
	return ApiCancelBlockTradeOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *MarketMakerBlockTradeAPIService) CancelBlockTradeOrderExecute(r ApiCancelBlockTradeOrderRequest) (struct{}, error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/block/order/create"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.blockOrderMatchingKey == nil {
		return struct{}{}, common.ReportError("blockOrderMatchingKey is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "blockOrderMatchingKey", r.blockOrderMatchingKey, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	_, err := SendRequest[struct{}](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil {
		return struct{}{}, err
	}

	return struct{}{}, nil
}

type ApiExtendBlockTradeOrderRequest struct {
	ctx                   context.Context
	ApiService            *MarketMakerBlockTradeAPIService
	blockOrderMatchingKey *string
	recvWindow            *int64
}

func (r ApiExtendBlockTradeOrderRequest) BlockOrderMatchingKey(blockOrderMatchingKey string) ApiExtendBlockTradeOrderRequest {
	r.blockOrderMatchingKey = &blockOrderMatchingKey
	return r
}

func (r ApiExtendBlockTradeOrderRequest) RecvWindow(recvWindow int64) ApiExtendBlockTradeOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiExtendBlockTradeOrderRequest) Execute() (*common.RestApiResponse[models.ExtendBlockTradeOrderResponse], error) {
	return r.ApiService.ExtendBlockTradeOrderExecute(r)
}

/*
ExtendBlockTradeOrder Extend Block Trade Order (TRADE)
Put /eapi/v1/block/order/create

https://developers.binance.com/docs/derivatives/options-trading/market-maker-block-trade/Extend-Block-Trade-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param blockOrderMatchingKey -
@param recvWindow -
@return ApiExtendBlockTradeOrderRequest
*/
func (a *MarketMakerBlockTradeAPIService) ExtendBlockTradeOrder(ctx context.Context) ApiExtendBlockTradeOrderRequest {
	return ApiExtendBlockTradeOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ExtendBlockTradeOrderResponse
func (a *MarketMakerBlockTradeAPIService) ExtendBlockTradeOrderExecute(r ApiExtendBlockTradeOrderRequest) (*common.RestApiResponse[models.ExtendBlockTradeOrderResponse], error) {
	localVarHTTPMethod := http.MethodPut
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/block/order/create"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.blockOrderMatchingKey == nil {
		return nil, common.ReportError("blockOrderMatchingKey is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "blockOrderMatchingKey", r.blockOrderMatchingKey, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.ExtendBlockTradeOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiNewBlockTradeOrderRequest struct {
	ctx        context.Context
	ApiService *MarketMakerBlockTradeAPIService
	liquidity  *string
	legs       *[]map[string]interface{}
	recvWindow *int64
}

// Taker or Maker
func (r ApiNewBlockTradeOrderRequest) Liquidity(liquidity string) ApiNewBlockTradeOrderRequest {
	r.liquidity = &liquidity
	return r
}

// Max 1 (only single leg supported), list of legs parameters in JSON; example: eapi/v1/block/order/create?orders&#x3D;[{\&quot;symbol\&quot;:\&quot;BTC-210115-35000-C\&quot;, \&quot;price\&quot;:\&quot;100\&quot;,\&quot;quantity\&quot;:\&quot;0.0002\&quot;,\&quot;side\&quot;:\&quot;BUY\&quot;,\&quot;type\&quot;:\&quot;LIMIT\&quot;}]
func (r ApiNewBlockTradeOrderRequest) Legs(legs []map[string]interface{}) ApiNewBlockTradeOrderRequest {
	r.legs = &legs
	return r
}

func (r ApiNewBlockTradeOrderRequest) RecvWindow(recvWindow int64) ApiNewBlockTradeOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiNewBlockTradeOrderRequest) Execute() (*common.RestApiResponse[models.NewBlockTradeOrderResponse], error) {
	return r.ApiService.NewBlockTradeOrderExecute(r)
}

/*
NewBlockTradeOrder New Block Trade Order (TRADE)
Post /eapi/v1/block/order/create

https://developers.binance.com/docs/derivatives/options-trading/market-maker-block-trade/New-Block-Trade-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param liquidity -  Taker or Maker
@param legs -  Max 1 (only single leg supported), list of legs parameters in JSON; example: eapi/v1/block/order/create?orders=[{\"symbol\":\"BTC-210115-35000-C\", \"price\":\"100\",\"quantity\":\"0.0002\",\"side\":\"BUY\",\"type\":\"LIMIT\"}]
@param recvWindow -
@return ApiNewBlockTradeOrderRequest
*/
func (a *MarketMakerBlockTradeAPIService) NewBlockTradeOrder(ctx context.Context) ApiNewBlockTradeOrderRequest {
	return ApiNewBlockTradeOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return NewBlockTradeOrderResponse
func (a *MarketMakerBlockTradeAPIService) NewBlockTradeOrderExecute(r ApiNewBlockTradeOrderRequest) (*common.RestApiResponse[models.NewBlockTradeOrderResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/block/order/create"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.liquidity == nil {
		return nil, common.ReportError("liquidity is required and must be specified")
	}
	if r.legs == nil {
		return nil, common.ReportError("legs is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "liquidity", r.liquidity, "form", "")
	{
		t := *r.legs
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "legs", t, "form", "multi")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.NewBlockTradeOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryBlockTradeDetailsRequest struct {
	ctx                   context.Context
	ApiService            *MarketMakerBlockTradeAPIService
	blockOrderMatchingKey *string
	recvWindow            *int64
}

func (r ApiQueryBlockTradeDetailsRequest) BlockOrderMatchingKey(blockOrderMatchingKey string) ApiQueryBlockTradeDetailsRequest {
	r.blockOrderMatchingKey = &blockOrderMatchingKey
	return r
}

func (r ApiQueryBlockTradeDetailsRequest) RecvWindow(recvWindow int64) ApiQueryBlockTradeDetailsRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryBlockTradeDetailsRequest) Execute() (*common.RestApiResponse[models.QueryBlockTradeDetailsResponse], error) {
	return r.ApiService.QueryBlockTradeDetailsExecute(r)
}

/*
QueryBlockTradeDetails Query Block Trade Details (USER_DATA)
Get /eapi/v1/block/order/execute

https://developers.binance.com/docs/derivatives/options-trading/market-maker-block-trade/Query-Block-Trade-Detail

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param blockOrderMatchingKey -
@param recvWindow -
@return ApiQueryBlockTradeDetailsRequest
*/
func (a *MarketMakerBlockTradeAPIService) QueryBlockTradeDetails(ctx context.Context) ApiQueryBlockTradeDetailsRequest {
	return ApiQueryBlockTradeDetailsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryBlockTradeDetailsResponse
func (a *MarketMakerBlockTradeAPIService) QueryBlockTradeDetailsExecute(r ApiQueryBlockTradeDetailsRequest) (*common.RestApiResponse[models.QueryBlockTradeDetailsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/block/order/execute"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.blockOrderMatchingKey == nil {
		return nil, common.ReportError("blockOrderMatchingKey is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "blockOrderMatchingKey", r.blockOrderMatchingKey, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryBlockTradeDetailsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryBlockTradeOrderRequest struct {
	ctx                   context.Context
	ApiService            *MarketMakerBlockTradeAPIService
	blockOrderMatchingKey *string
	endTime               *int64
	startTime             *int64
	underlying            *string
	recvWindow            *int64
}

// If specified, returns the specific block trade associated with the blockOrderMatchingKey
func (r ApiQueryBlockTradeOrderRequest) BlockOrderMatchingKey(blockOrderMatchingKey string) ApiQueryBlockTradeOrderRequest {
	r.blockOrderMatchingKey = &blockOrderMatchingKey
	return r
}

// End Time, e.g 1593512200000
func (r ApiQueryBlockTradeOrderRequest) EndTime(endTime int64) ApiQueryBlockTradeOrderRequest {
	r.endTime = &endTime
	return r
}

// Start Time, e.g 1593511200000
func (r ApiQueryBlockTradeOrderRequest) StartTime(startTime int64) ApiQueryBlockTradeOrderRequest {
	r.startTime = &startTime
	return r
}

// underlying, e.g BTCUSDT
func (r ApiQueryBlockTradeOrderRequest) Underlying(underlying string) ApiQueryBlockTradeOrderRequest {
	r.underlying = &underlying
	return r
}

func (r ApiQueryBlockTradeOrderRequest) RecvWindow(recvWindow int64) ApiQueryBlockTradeOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryBlockTradeOrderRequest) Execute() (*common.RestApiResponse[models.QueryBlockTradeOrderResponse], error) {
	return r.ApiService.QueryBlockTradeOrderExecute(r)
}

/*
QueryBlockTradeOrder Query Block Trade Order (TRADE)
Get /eapi/v1/block/order/orders

https://developers.binance.com/docs/derivatives/options-trading/market-maker-block-trade/Query-Block-Trade-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param blockOrderMatchingKey -  If specified, returns the specific block trade associated with the blockOrderMatchingKey
@param endTime -  End Time, e.g 1593512200000
@param startTime -  Start Time, e.g 1593511200000
@param underlying -  underlying, e.g BTCUSDT
@param recvWindow -
@return ApiQueryBlockTradeOrderRequest
*/
func (a *MarketMakerBlockTradeAPIService) QueryBlockTradeOrder(ctx context.Context) ApiQueryBlockTradeOrderRequest {
	return ApiQueryBlockTradeOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryBlockTradeOrderResponse
func (a *MarketMakerBlockTradeAPIService) QueryBlockTradeOrderExecute(r ApiQueryBlockTradeOrderRequest) (*common.RestApiResponse[models.QueryBlockTradeOrderResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/block/order/orders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.blockOrderMatchingKey != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "blockOrderMatchingKey", r.blockOrderMatchingKey, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.underlying != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "underlying", r.underlying, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryBlockTradeOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
