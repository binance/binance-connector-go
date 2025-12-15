/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API

API version: 1.0.0
*/

package binancemargintradingrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/margintrading/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// TransferAPIService TransferAPI Service
type TransferAPIService Service

type ApiGetCrossMarginTransferHistoryRequest struct {
	ctx            context.Context
	ApiService     *TransferAPIService
	asset          *string
	type_          *string
	startTime      *int64
	endTime        *int64
	current        *int64
	size           *int64
	isolatedSymbol *string
	recvWindow     *int64
}

func (r ApiGetCrossMarginTransferHistoryRequest) Asset(asset string) ApiGetCrossMarginTransferHistoryRequest {
	r.asset = &asset
	return r
}

// Transfer Type: ROLL_IN, ROLL_OUT
func (r ApiGetCrossMarginTransferHistoryRequest) Type(type_ string) ApiGetCrossMarginTransferHistoryRequest {
	r.type_ = &type_
	return r
}

// 只支持查询最近90天的数据
func (r ApiGetCrossMarginTransferHistoryRequest) StartTime(startTime int64) ApiGetCrossMarginTransferHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetCrossMarginTransferHistoryRequest) EndTime(endTime int64) ApiGetCrossMarginTransferHistoryRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page. Start from 1. Default:1
func (r ApiGetCrossMarginTransferHistoryRequest) Current(current int64) ApiGetCrossMarginTransferHistoryRequest {
	r.current = &current
	return r
}

// Default:10 Max:100
func (r ApiGetCrossMarginTransferHistoryRequest) Size(size int64) ApiGetCrossMarginTransferHistoryRequest {
	r.size = &size
	return r
}

// isolated symbol
func (r ApiGetCrossMarginTransferHistoryRequest) IsolatedSymbol(isolatedSymbol string) ApiGetCrossMarginTransferHistoryRequest {
	r.isolatedSymbol = &isolatedSymbol
	return r
}

// No more than 60000
func (r ApiGetCrossMarginTransferHistoryRequest) RecvWindow(recvWindow int64) ApiGetCrossMarginTransferHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetCrossMarginTransferHistoryRequest) Execute() (*common.RestApiResponse[models.GetCrossMarginTransferHistoryResponse], error) {
	return r.ApiService.GetCrossMarginTransferHistoryExecute(r)
}

/*
GetCrossMarginTransferHistory Get Cross Margin Transfer History (USER_DATA)
Get /sapi/v1/margin/transfer

https://developers.binance.com/docs/margin_trading/transfer/Get-Cross-Margin-Transfer-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -
@param type_ -  Transfer Type: ROLL_IN, ROLL_OUT
@param startTime -  只支持查询最近90天的数据
@param endTime -
@param current -  Currently querying page. Start from 1. Default:1
@param size -  Default:10 Max:100
@param isolatedSymbol -  isolated symbol
@param recvWindow -  No more than 60000
@return ApiGetCrossMarginTransferHistoryRequest
*/
func (a *TransferAPIService) GetCrossMarginTransferHistory(ctx context.Context) ApiGetCrossMarginTransferHistoryRequest {
	return ApiGetCrossMarginTransferHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetCrossMarginTransferHistoryResponse
func (a *TransferAPIService) GetCrossMarginTransferHistoryExecute(r ApiGetCrossMarginTransferHistoryRequest) (*common.RestApiResponse[models.GetCrossMarginTransferHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/transfer"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	}
	if r.type_ != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	}
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
	if r.isolatedSymbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "isolatedSymbol", r.isolatedSymbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetCrossMarginTransferHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryMaxTransferOutAmountRequest struct {
	ctx            context.Context
	ApiService     *TransferAPIService
	asset          *string
	isolatedSymbol *string
	recvWindow     *int64
}

func (r ApiQueryMaxTransferOutAmountRequest) Asset(asset string) ApiQueryMaxTransferOutAmountRequest {
	r.asset = &asset
	return r
}

// isolated symbol
func (r ApiQueryMaxTransferOutAmountRequest) IsolatedSymbol(isolatedSymbol string) ApiQueryMaxTransferOutAmountRequest {
	r.isolatedSymbol = &isolatedSymbol
	return r
}

// No more than 60000
func (r ApiQueryMaxTransferOutAmountRequest) RecvWindow(recvWindow int64) ApiQueryMaxTransferOutAmountRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryMaxTransferOutAmountRequest) Execute() (*common.RestApiResponse[models.QueryMaxTransferOutAmountResponse], error) {
	return r.ApiService.QueryMaxTransferOutAmountExecute(r)
}

/*
QueryMaxTransferOutAmount Query Max Transfer-Out Amount (USER_DATA)
Get /sapi/v1/margin/maxTransferable

https://developers.binance.com/docs/margin_trading/transfer/Query-Max-Transfer-Out-Amount

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -
@param isolatedSymbol -  isolated symbol
@param recvWindow -  No more than 60000
@return ApiQueryMaxTransferOutAmountRequest
*/
func (a *TransferAPIService) QueryMaxTransferOutAmount(ctx context.Context) ApiQueryMaxTransferOutAmountRequest {
	return ApiQueryMaxTransferOutAmountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryMaxTransferOutAmountResponse
func (a *TransferAPIService) QueryMaxTransferOutAmountExecute(r ApiQueryMaxTransferOutAmountRequest) (*common.RestApiResponse[models.QueryMaxTransferOutAmountResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/maxTransferable"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset == nil {
		return nil, common.ReportError("asset is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	if r.isolatedSymbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "isolatedSymbol", r.isolatedSymbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryMaxTransferOutAmountResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
