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

// AccountAPIService AccountAPI Service
type AccountAPIService Service

type ApiSignUsEquityDisclaimerRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

// The value cannot be greater than &#x60;60000&#x60;.
func (r ApiSignUsEquityDisclaimerRequest) RecvWindow(recvWindow int64) ApiSignUsEquityDisclaimerRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSignUsEquityDisclaimerRequest) Execute() (*common.RestApiResponse[models.SignUsEquityDisclaimerResponse], error) {
	return r.ApiService.SignUsEquityDisclaimerExecute(r)
}

/*
SignUsEquityDisclaimer Sign US Equity Disclaimer (TRADE)
Post /sapi/v1/equity/account/disclaimer

https://developers.binance.com/en/docs/catalog/advanced-trading-stocks-trading/api/rest-api/account#sign-us-equity-disclaimer

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -  The value cannot be greater than `60000`.
@return ApiSignUsEquityDisclaimerRequest
*/
func (a *AccountAPIService) SignUsEquityDisclaimer(ctx context.Context) ApiSignUsEquityDisclaimerRequest {
	return ApiSignUsEquityDisclaimerRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SignUsEquityDisclaimerResponse
func (a *AccountAPIService) SignUsEquityDisclaimerExecute(r ApiSignUsEquityDisclaimerRequest) (*common.RestApiResponse[models.SignUsEquityDisclaimerResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/equity/account/disclaimer"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.SignUsEquityDisclaimerResponse](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		true,
	)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
