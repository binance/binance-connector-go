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
