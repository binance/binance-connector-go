/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package binancemargintradingrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/margintrading/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// RiskDataStreamAPIService RiskDataStreamAPI Service
type RiskDataStreamAPIService Service

type ApiCloseUserDataStreamRequest struct {
	ctx        context.Context
	ApiService *RiskDataStreamAPIService
}

func (r ApiCloseUserDataStreamRequest) Execute() (struct{}, error) {
	return r.ApiService.CloseUserDataStreamExecute(r)
}

/*
CloseUserDataStream Close User Data Stream (USER_STREAM)
Delete /sapi/v1/margin/listen-key

https://developers.binance.com/docs/margin_trading/risk-data-stream/Close-User-Data-Stream

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiCloseUserDataStreamRequest
*/
func (a *RiskDataStreamAPIService) CloseUserDataStream(ctx context.Context) ApiCloseUserDataStreamRequest {
	return ApiCloseUserDataStreamRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *RiskDataStreamAPIService) CloseUserDataStreamExecute(r ApiCloseUserDataStreamRequest) (struct{}, error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/listen-key"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	_, err := SendRequest[struct{}](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil {
		return struct{}{}, err
	}

	return struct{}{}, nil
}

type ApiKeepaliveUserDataStreamRequest struct {
	ctx        context.Context
	ApiService *RiskDataStreamAPIService
	listenKey  *string
}

func (r ApiKeepaliveUserDataStreamRequest) ListenKey(listenKey string) ApiKeepaliveUserDataStreamRequest {
	r.listenKey = &listenKey
	return r
}

func (r ApiKeepaliveUserDataStreamRequest) Execute() (struct{}, error) {
	return r.ApiService.KeepaliveUserDataStreamExecute(r)
}

/*
KeepaliveUserDataStream Keepalive User Data Stream (USER_STREAM)
Put /sapi/v1/margin/listen-key

https://developers.binance.com/docs/margin_trading/risk-data-stream/Keepalive-User-Data-Stream

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param listenKey -
@return ApiKeepaliveUserDataStreamRequest
*/
func (a *RiskDataStreamAPIService) KeepaliveUserDataStream(ctx context.Context) ApiKeepaliveUserDataStreamRequest {
	return ApiKeepaliveUserDataStreamRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *RiskDataStreamAPIService) KeepaliveUserDataStreamExecute(r ApiKeepaliveUserDataStreamRequest) (struct{}, error) {
	localVarHTTPMethod := http.MethodPut
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/listen-key"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.listenKey == nil {
		return struct{}{}, common.ReportError("listenKey is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "listenKey", r.listenKey, "form", "")

	_, err := SendRequest[struct{}](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil {
		return struct{}{}, err
	}

	return struct{}{}, nil
}

type ApiStartUserDataStreamRequest struct {
	ctx        context.Context
	ApiService *RiskDataStreamAPIService
}

func (r ApiStartUserDataStreamRequest) Execute() (*common.RestApiResponse[models.StartUserDataStreamResponse], error) {
	return r.ApiService.StartUserDataStreamExecute(r)
}

/*
StartUserDataStream Start User Data Stream (USER_STREAM)
Post /sapi/v1/margin/listen-key

https://developers.binance.com/docs/margin_trading/risk-data-stream/Start-User-Data-Stream

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiStartUserDataStreamRequest
*/
func (a *RiskDataStreamAPIService) StartUserDataStream(ctx context.Context) ApiStartUserDataStreamRequest {
	return ApiStartUserDataStreamRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StartUserDataStreamResponse
func (a *RiskDataStreamAPIService) StartUserDataStreamExecute(r ApiStartUserDataStreamRequest) (*common.RestApiResponse[models.StartUserDataStreamResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/listen-key"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	resp, err := SendRequest[models.StartUserDataStreamResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
