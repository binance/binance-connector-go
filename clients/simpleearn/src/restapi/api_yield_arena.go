/*
Simple Earn REST API

Earn rewards by subscribing to flexible or locked Simple Earn products.
*/

package binancesimpleearnrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/simpleearn/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// YieldArenaAPIService YieldArenaAPI Service
type YieldArenaAPIService Service

type ApiGetYieldArenaActivitiesRequest struct {
	ctx        context.Context
	ApiService *YieldArenaAPIService
	lang       *string
	recvWindow *int64
}

// Locale tag for &#x60;title&#x60; and &#x60;description&#x60; (e.g. &#x60;en&#x60;, &#x60;zh-CN&#x60;, &#x60;pt-BR&#x60;). Default: &#x60;en&#x60;. If the value is missing, malformed, or has no translation configured, content is returned in &#x60;en&#x60;.
func (r ApiGetYieldArenaActivitiesRequest) Lang(lang string) ApiGetYieldArenaActivitiesRequest {
	r.lang = &lang
	return r
}

func (r ApiGetYieldArenaActivitiesRequest) RecvWindow(recvWindow int64) ApiGetYieldArenaActivitiesRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetYieldArenaActivitiesRequest) Execute() (*common.RestApiResponse[models.GetYieldArenaActivitiesResponse], error) {
	return r.ApiService.GetYieldArenaActivitiesExecute(r)
}

/*
GetYieldArenaActivities Get Yield Arena Activities (USER_DATA)
Get /sapi/v1/earn/arena/activities

https://developers.binance.com/en/docs/catalog/investment-and-services-simple-earn/api/rest-api/yield-arena#get-yield-arena-activities

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param lang -  Locale tag for `title` and `description` (e.g. `en`, `zh-CN`, `pt-BR`). Default: `en`. If the value is missing, malformed, or has no translation configured, content is returned in `en`.
@param recvWindow -
@return ApiGetYieldArenaActivitiesRequest
*/
func (a *YieldArenaAPIService) GetYieldArenaActivities(ctx context.Context) ApiGetYieldArenaActivitiesRequest {
	return ApiGetYieldArenaActivitiesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetYieldArenaActivitiesResponse
func (a *YieldArenaAPIService) GetYieldArenaActivitiesExecute(r ApiGetYieldArenaActivitiesRequest) (*common.RestApiResponse[models.GetYieldArenaActivitiesResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/earn/arena/activities"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.lang != nil {
		localVarBodyParameters["lang"] = *r.lang
	}

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetYieldArenaActivitiesResponse](
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
