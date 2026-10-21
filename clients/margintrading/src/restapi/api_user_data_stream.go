/*
Margin REST API

Access account information, borrow and repay assets, and trade with Binance Margin.
*/

package binancemargintradingrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/margintrading/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// UserDataStreamAPIService UserDataStreamAPI Service
type UserDataStreamAPIService Service

type ApiCloseUserDataStreamRequest struct {
	ctx        context.Context
	ApiService *UserDataStreamAPIService
}

func (r ApiCloseUserDataStreamRequest) Execute() (struct{}, error) {
	return r.ApiService.CloseUserDataStreamExecute(r)
}

/*
CloseUserDataStream Close User Data Stream (USER_STREAM)
Delete /sapi/v1/margin/listen-key

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/user-data-stream#close-user-data-stream

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiCloseUserDataStreamRequest
*/
func (a *UserDataStreamAPIService) CloseUserDataStream(ctx context.Context) ApiCloseUserDataStreamRequest {
	return ApiCloseUserDataStreamRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *UserDataStreamAPIService) CloseUserDataStreamExecute(r ApiCloseUserDataStreamRequest) (struct{}, error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/listen-key"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	_, err := SendRequest[struct{}](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		false,
	)
	if err != nil {
		return struct{}{}, err
	}

	return struct{}{}, nil
}

type ApiCloseUserDataStreamV2Request struct {
	ctx        context.Context
	ApiService *UserDataStreamAPIService
}

func (r ApiCloseUserDataStreamV2Request) Execute() (struct{}, error) {
	return r.ApiService.CloseUserDataStreamV2Execute(r)
}

/*
CloseUserDataStreamV2 Close User Data Stream V2 (USER_STREAM)
Delete /sapi/v2/margin/listen-key

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/user-data-stream#close-user-data-stream-v2

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiCloseUserDataStreamV2Request
*/
func (a *UserDataStreamAPIService) CloseUserDataStreamV2(ctx context.Context) ApiCloseUserDataStreamV2Request {
	return ApiCloseUserDataStreamV2Request{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *UserDataStreamAPIService) CloseUserDataStreamV2Execute(r ApiCloseUserDataStreamV2Request) (struct{}, error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/sapi/v2/margin/listen-key"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	_, err := SendRequest[struct{}](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		false,
	)
	if err != nil {
		return struct{}{}, err
	}

	return struct{}{}, nil
}

type ApiKeepaliveUserDataStreamRequest struct {
	ctx        context.Context
	ApiService *UserDataStreamAPIService
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

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/user-data-stream#keepalive-user-data-stream

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param listenKey -
@return ApiKeepaliveUserDataStreamRequest
*/
func (a *UserDataStreamAPIService) KeepaliveUserDataStream(ctx context.Context) ApiKeepaliveUserDataStreamRequest {
	return ApiKeepaliveUserDataStreamRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *UserDataStreamAPIService) KeepaliveUserDataStreamExecute(r ApiKeepaliveUserDataStreamRequest) (struct{}, error) {
	localVarHTTPMethod := http.MethodPut
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/listen-key"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.listenKey == nil {
		return struct{}{}, common.ReportError("listenKey is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "listenKey", r.listenKey, "form", "")

	_, err := SendRequest[struct{}](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		false,
	)
	if err != nil {
		return struct{}{}, err
	}

	return struct{}{}, nil
}

type ApiKeepaliveUserDataStreamV2Request struct {
	ctx        context.Context
	ApiService *UserDataStreamAPIService
	listenKey  *string
}

func (r ApiKeepaliveUserDataStreamV2Request) ListenKey(listenKey string) ApiKeepaliveUserDataStreamV2Request {
	r.listenKey = &listenKey
	return r
}

func (r ApiKeepaliveUserDataStreamV2Request) Execute() (struct{}, error) {
	return r.ApiService.KeepaliveUserDataStreamV2Execute(r)
}

/*
KeepaliveUserDataStreamV2 Keepalive User Data Stream V2 (USER_STREAM)
Put /sapi/v2/margin/listen-key

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/user-data-stream#keepalive-user-data-stream-v2

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param listenKey -
@return ApiKeepaliveUserDataStreamV2Request
*/
func (a *UserDataStreamAPIService) KeepaliveUserDataStreamV2(ctx context.Context) ApiKeepaliveUserDataStreamV2Request {
	return ApiKeepaliveUserDataStreamV2Request{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *UserDataStreamAPIService) KeepaliveUserDataStreamV2Execute(r ApiKeepaliveUserDataStreamV2Request) (struct{}, error) {
	localVarHTTPMethod := http.MethodPut
	localVarPath := a.client.cfg.BasePath + "/sapi/v2/margin/listen-key"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.listenKey == nil {
		return struct{}{}, common.ReportError("listenKey is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "listenKey", r.listenKey, "form", "")

	_, err := SendRequest[struct{}](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		false,
	)
	if err != nil {
		return struct{}{}, err
	}

	return struct{}{}, nil
}

type ApiStartUserDataStreamRequest struct {
	ctx        context.Context
	ApiService *UserDataStreamAPIService
}

func (r ApiStartUserDataStreamRequest) Execute() (*common.RestApiResponse[models.StartUserDataStreamResponse], error) {
	return r.ApiService.StartUserDataStreamExecute(r)
}

/*
StartUserDataStream Start User Data Stream (USER_STREAM)
Post /sapi/v1/margin/listen-key

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/user-data-stream#start-user-data-stream

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiStartUserDataStreamRequest
*/
func (a *UserDataStreamAPIService) StartUserDataStream(ctx context.Context) ApiStartUserDataStreamRequest {
	return ApiStartUserDataStreamRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StartUserDataStreamResponse
func (a *UserDataStreamAPIService) StartUserDataStreamExecute(r ApiStartUserDataStreamRequest) (*common.RestApiResponse[models.StartUserDataStreamResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/listen-key"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	resp, err := SendRequest[models.StartUserDataStreamResponse](
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

type ApiStartUserDataStreamV2Request struct {
	ctx        context.Context
	ApiService *UserDataStreamAPIService
}

func (r ApiStartUserDataStreamV2Request) Execute() (*common.RestApiResponse[models.StartUserDataStreamV2Response], error) {
	return r.ApiService.StartUserDataStreamV2Execute(r)
}

/*
StartUserDataStreamV2 Start User Data Stream V2 (USER_STREAM)
Post /sapi/v2/margin/listen-key

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/user-data-stream#start-user-data-stream-v2

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiStartUserDataStreamV2Request
*/
func (a *UserDataStreamAPIService) StartUserDataStreamV2(ctx context.Context) ApiStartUserDataStreamV2Request {
	return ApiStartUserDataStreamV2Request{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StartUserDataStreamV2Response
func (a *UserDataStreamAPIService) StartUserDataStreamV2Execute(r ApiStartUserDataStreamV2Request) (*common.RestApiResponse[models.StartUserDataStreamV2Response], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v2/margin/listen-key"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	resp, err := SendRequest[models.StartUserDataStreamV2Response](
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
