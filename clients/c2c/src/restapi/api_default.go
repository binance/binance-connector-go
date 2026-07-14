/*
C2C REST API

Query fiat transaction history via the C2C REST API.
*/

package binancec2crestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/c2c/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// DefaultAPIService DefaultAPI Service
type DefaultAPIService Service

type ApiGetC2CTradeHistoryRequest struct {
	ctx            context.Context
	ApiService     *DefaultAPIService
	tradeType      *models.GetC2CTradeHistoryTradeTypeParameter
	startTimestamp *int64
	endTimestamp   *int64
	page           *int64
	rows           *int64
}

// Trade side filter
func (r ApiGetC2CTradeHistoryRequest) TradeType(tradeType models.GetC2CTradeHistoryTradeTypeParameter) ApiGetC2CTradeHistoryRequest {
	r.tradeType = &tradeType
	return r
}

func (r ApiGetC2CTradeHistoryRequest) StartTimestamp(startTimestamp int64) ApiGetC2CTradeHistoryRequest {
	r.startTimestamp = &startTimestamp
	return r
}

func (r ApiGetC2CTradeHistoryRequest) EndTimestamp(endTimestamp int64) ApiGetC2CTradeHistoryRequest {
	r.endTimestamp = &endTimestamp
	return r
}

// Page number
func (r ApiGetC2CTradeHistoryRequest) Page(page int64) ApiGetC2CTradeHistoryRequest {
	r.page = &page
	return r
}

// Number of records per page
func (r ApiGetC2CTradeHistoryRequest) Rows(rows int64) ApiGetC2CTradeHistoryRequest {
	r.rows = &rows
	return r
}

func (r ApiGetC2CTradeHistoryRequest) Execute() (*common.RestApiResponse[models.GetC2CTradeHistoryResponse], error) {
	return r.ApiService.GetC2CTradeHistoryExecute(r)
}

/*
GetC2CTradeHistory Get C2C Trade History (USER_DATA)
Get /sapi/v1/c2c/orderMatch/listUserOrderHistory

https://developers.binance.com/en/docs/catalog/investment-and-services-c2-c/api/rest-api/~#get-c2-ctrade-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param tradeType -  Trade side filter
@param startTimestamp -
@param endTimestamp -
@param page -  Page number
@param rows -  Number of records per page
@return ApiGetC2CTradeHistoryRequest
*/
func (a *DefaultAPIService) GetC2CTradeHistory(ctx context.Context) ApiGetC2CTradeHistoryRequest {
	return ApiGetC2CTradeHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetC2CTradeHistoryResponse
func (a *DefaultAPIService) GetC2CTradeHistoryExecute(r ApiGetC2CTradeHistoryRequest) (*common.RestApiResponse[models.GetC2CTradeHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/c2c/orderMatch/listUserOrderHistory"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.tradeType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "tradeType", r.tradeType, "form", "")
	}
	if r.startTimestamp != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTimestamp", r.startTimestamp, "form", "")
	}
	if r.endTimestamp != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTimestamp", r.endTimestamp, "form", "")
	}
	if r.page != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.rows != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "rows", r.rows, "form", "")
	}

	resp, err := SendRequest[models.GetC2CTradeHistoryResponse](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		false,
	)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
