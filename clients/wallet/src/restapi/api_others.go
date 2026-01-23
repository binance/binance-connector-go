/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package binancewalletrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/wallet/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// OthersAPIService OthersAPI Service
type OthersAPIService Service

type ApiGetSymbolsDelistScheduleForSpotRequest struct {
	ctx        context.Context
	ApiService *OthersAPIService
	recvWindow *int64
}

func (r ApiGetSymbolsDelistScheduleForSpotRequest) RecvWindow(recvWindow int64) ApiGetSymbolsDelistScheduleForSpotRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetSymbolsDelistScheduleForSpotRequest) Execute() (*common.RestApiResponse[models.GetSymbolsDelistScheduleForSpotResponse], error) {
	return r.ApiService.GetSymbolsDelistScheduleForSpotExecute(r)
}

/*
GetSymbolsDelistScheduleForSpot Get symbols delist schedule for spot (MARKET_DATA)
Get /sapi/v1/spot/delist-schedule

https://developers.binance.com/docs/wallet/others/delist-schedule

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiGetSymbolsDelistScheduleForSpotRequest
*/
func (a *OthersAPIService) GetSymbolsDelistScheduleForSpot(ctx context.Context) ApiGetSymbolsDelistScheduleForSpotRequest {
	return ApiGetSymbolsDelistScheduleForSpotRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetSymbolsDelistScheduleForSpotResponse
func (a *OthersAPIService) GetSymbolsDelistScheduleForSpotExecute(r ApiGetSymbolsDelistScheduleForSpotRequest) (*common.RestApiResponse[models.GetSymbolsDelistScheduleForSpotResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/spot/delist-schedule"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetSymbolsDelistScheduleForSpotResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiSystemStatusRequest struct {
	ctx        context.Context
	ApiService *OthersAPIService
}

func (r ApiSystemStatusRequest) Execute() (*common.RestApiResponse[models.SystemStatusResponse], error) {
	return r.ApiService.SystemStatusExecute(r)
}

/*
SystemStatus System Status (System)
Get /sapi/v1/system/status

https://developers.binance.com/docs/wallet/others/System-Status

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiSystemStatusRequest
*/
func (a *OthersAPIService) SystemStatus(ctx context.Context) ApiSystemStatusRequest {
	return ApiSystemStatusRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SystemStatusResponse
func (a *OthersAPIService) SystemStatusExecute(r ApiSystemStatusRequest) (*common.RestApiResponse[models.SystemStatusResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/system/status"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	resp, err := SendRequest[models.SystemStatusResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
