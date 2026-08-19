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

// UserDataStreamsAPIService UserDataStreamsAPI Service
type UserDataStreamsAPIService Service

type ApiCreateRenewListenKeyRequest struct {
	ctx        context.Context
	ApiService *UserDataStreamsAPIService
	recvWindow *int64
}

// The value cannot be greater than &#x60;60000&#x60;.
func (r ApiCreateRenewListenKeyRequest) RecvWindow(recvWindow int64) ApiCreateRenewListenKeyRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCreateRenewListenKeyRequest) Execute() (*common.RestApiResponse[models.CreateRenewListenKeyResponse], error) {
	return r.ApiService.CreateRenewListenKeyExecute(r)
}

/*
CreateRenewListenKey Create / Renew Listen Key (USER_STREAM)
Post /sapi/v1/equity/listenKey

https://developers.binance.com/en/docs/catalog/advanced-trading-stocks-trading/api/rest-api/user-data-streams#create-renew-listen-key

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -  The value cannot be greater than `60000`.
@return ApiCreateRenewListenKeyRequest
*/
func (a *UserDataStreamsAPIService) CreateRenewListenKey(ctx context.Context) ApiCreateRenewListenKeyRequest {
	return ApiCreateRenewListenKeyRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateRenewListenKeyResponse
func (a *UserDataStreamsAPIService) CreateRenewListenKeyExecute(r ApiCreateRenewListenKeyRequest) (*common.RestApiResponse[models.CreateRenewListenKeyResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/equity/listenKey"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CreateRenewListenKeyResponse](
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
