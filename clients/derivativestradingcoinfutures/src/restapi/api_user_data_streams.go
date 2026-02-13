/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package binancederivativestradingcoinfuturesrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// UserDataStreamsAPIService UserDataStreamsAPI Service
type UserDataStreamsAPIService Service

type ApiCloseUserDataStreamRequest struct {
	ctx        context.Context
	ApiService *UserDataStreamsAPIService
}

func (r ApiCloseUserDataStreamRequest) Execute() (struct{}, error) {
	return r.ApiService.CloseUserDataStreamExecute(r)
}

/*
CloseUserDataStream Close User Data Stream(USER_STREAM)
Delete /dapi/v1/listenKey

https://developers.binance.com/docs/derivatives/coin-margined-futures/user-data-streams/Close-User-Data-Stream

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiCloseUserDataStreamRequest
*/
func (a *UserDataStreamsAPIService) CloseUserDataStream(ctx context.Context) ApiCloseUserDataStreamRequest {
	return ApiCloseUserDataStreamRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *UserDataStreamsAPIService) CloseUserDataStreamExecute(r ApiCloseUserDataStreamRequest) (struct{}, error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/listenKey"

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
	ApiService *UserDataStreamsAPIService
}

func (r ApiKeepaliveUserDataStreamRequest) Execute() (*common.RestApiResponse[models.KeepaliveUserDataStreamResponse], error) {
	return r.ApiService.KeepaliveUserDataStreamExecute(r)
}

/*
KeepaliveUserDataStream Keepalive User Data Stream (USER_STREAM)
Put /dapi/v1/listenKey

https://developers.binance.com/docs/derivatives/coin-margined-futures/user-data-streams/Keepalive-User-Data-Stream

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiKeepaliveUserDataStreamRequest
*/
func (a *UserDataStreamsAPIService) KeepaliveUserDataStream(ctx context.Context) ApiKeepaliveUserDataStreamRequest {
	return ApiKeepaliveUserDataStreamRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return KeepaliveUserDataStreamResponse
func (a *UserDataStreamsAPIService) KeepaliveUserDataStreamExecute(r ApiKeepaliveUserDataStreamRequest) (*common.RestApiResponse[models.KeepaliveUserDataStreamResponse], error) {
	localVarHTTPMethod := http.MethodPut
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/listenKey"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	resp, err := SendRequest[models.KeepaliveUserDataStreamResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiStartUserDataStreamRequest struct {
	ctx        context.Context
	ApiService *UserDataStreamsAPIService
}

func (r ApiStartUserDataStreamRequest) Execute() (*common.RestApiResponse[models.StartUserDataStreamResponse], error) {
	return r.ApiService.StartUserDataStreamExecute(r)
}

/*
StartUserDataStream Start User Data Stream (USER_STREAM)
Post /dapi/v1/listenKey

https://developers.binance.com/docs/derivatives/coin-margined-futures/user-data-streams/Start-User-Data-Stream

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiStartUserDataStreamRequest
*/
func (a *UserDataStreamsAPIService) StartUserDataStream(ctx context.Context) ApiStartUserDataStreamRequest {
	return ApiStartUserDataStreamRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StartUserDataStreamResponse
func (a *UserDataStreamsAPIService) StartUserDataStreamExecute(r ApiStartUserDataStreamRequest) (*common.RestApiResponse[models.StartUserDataStreamResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/listenKey"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	resp, err := SendRequest[models.StartUserDataStreamResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
