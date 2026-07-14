/*
Spot WebSocket API

Access market data, manage accounts, and trade on Binance Spot.
*/

package binancespotwebsocketapi

import (
	"github.com/binance/binance-connector-go/clients/spot/src/websocketapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// AccountAPIService AccountAPI Service
type AccountAPIService struct {
	Ws *common.WebsocketAPI
}

type ApiAccountCommissionRequest struct {
	ApiService *AccountAPIService
	symbol     *string
	id         *string
}

func (r ApiAccountCommissionRequest) Symbol(symbol string) ApiAccountCommissionRequest {
	r.symbol = &symbol
	return r
}

// Client-generated request identifier.
func (r ApiAccountCommissionRequest) Id(id string) ApiAccountCommissionRequest {
	r.id = &id
	return r
}

func (r ApiAccountCommissionRequest) Execute() (*common.ResponseOrRaw[models.AccountCommissionResponse], error) {
	respChan, errChan, err := r.ApiService.AccountCommissionExecute(r)
	if err != nil {
		return nil, err
	}

	select {
	case resp := <-respChan:
		return resp, nil
	case err := <-errChan:
		return nil, err
	}
}

func (r ApiAccountCommissionRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.AccountCommissionResponse], chan error, error) {
	return r.ApiService.AccountCommissionExecute(r)
}

/*
AccountCommission Account Commission Rates (USER_DATA)
/account.commission

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/account#account-commission

@param symbol	@param id Client-generated request identifier.
@return ApiAccountCommissionRequest
*/
func (a *AccountAPIService) AccountCommission() ApiAccountCommissionRequest {
	return ApiAccountCommissionRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AccountCommissionResponse
func (a *AccountAPIService) AccountCommissionExecute(r ApiAccountCommissionRequest) (chan *common.ResponseOrRaw[models.AccountCommissionResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}

	localPayload := map[string]any{
		"method": "/account.commission"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.AccountCommissionResponse](a.Ws, localPayload, sendParams)
}

type ApiAccountRateLimitsOrdersRequest struct {
	ApiService *AccountAPIService
	id         *string
	recvWindow *float32
}

// Client-generated request identifier.
func (r ApiAccountRateLimitsOrdersRequest) Id(id string) ApiAccountRateLimitsOrdersRequest {
	r.id = &id
	return r
}

// Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiAccountRateLimitsOrdersRequest) RecvWindow(recvWindow float32) ApiAccountRateLimitsOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiAccountRateLimitsOrdersRequest) Execute() (*common.ResponseOrRaw[models.AccountRateLimitsOrdersResponse], error) {
	respChan, errChan, err := r.ApiService.AccountRateLimitsOrdersExecute(r)
	if err != nil {
		return nil, err
	}

	select {
	case resp := <-respChan:
		return resp, nil
	case err := <-errChan:
		return nil, err
	}
}

func (r ApiAccountRateLimitsOrdersRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.AccountRateLimitsOrdersResponse], chan error, error) {
	return r.ApiService.AccountRateLimitsOrdersExecute(r)
}

/*
AccountRateLimitsOrders Unfilled Order Count (USER_DATA)
/account.rateLimits.orders

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/account#account-rate-limits-orders

@param id Client-generated request identifier.	@param recvWindow Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiAccountRateLimitsOrdersRequest
*/
func (a *AccountAPIService) AccountRateLimitsOrders() ApiAccountRateLimitsOrdersRequest {
	return ApiAccountRateLimitsOrdersRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AccountRateLimitsOrdersResponse
func (a *AccountAPIService) AccountRateLimitsOrdersExecute(r ApiAccountRateLimitsOrdersRequest) (chan *common.ResponseOrRaw[models.AccountRateLimitsOrdersResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/account.rateLimits.orders"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.AccountRateLimitsOrdersResponse](a.Ws, localPayload, sendParams)
}

type ApiAccountStatusRequest struct {
	ApiService       *AccountAPIService
	id               *string
	omitZeroBalances *bool
	recvWindow       *float32
}

// Client-generated request identifier.
func (r ApiAccountStatusRequest) Id(id string) ApiAccountStatusRequest {
	r.id = &id
	return r
}

// When set to &#x60;true&#x60;, emits only the non-zero balances of an account. Default value: &#x60;false&#x60;.
func (r ApiAccountStatusRequest) OmitZeroBalances(omitZeroBalances bool) ApiAccountStatusRequest {
	r.omitZeroBalances = &omitZeroBalances
	return r
}

// Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiAccountStatusRequest) RecvWindow(recvWindow float32) ApiAccountStatusRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiAccountStatusRequest) Execute() (*common.ResponseOrRaw[models.AccountStatusResponse], error) {
	respChan, errChan, err := r.ApiService.AccountStatusExecute(r)
	if err != nil {
		return nil, err
	}

	select {
	case resp := <-respChan:
		return resp, nil
	case err := <-errChan:
		return nil, err
	}
}

func (r ApiAccountStatusRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.AccountStatusResponse], chan error, error) {
	return r.ApiService.AccountStatusExecute(r)
}

/*
AccountStatus Account information (USER_DATA)
/account.status

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/account#account-status

@param id Client-generated request identifier.	@param omitZeroBalances When set to `true`, emits only the non-zero balances of an account. Default value: `false`.	@param recvWindow Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiAccountStatusRequest
*/
func (a *AccountAPIService) AccountStatus() ApiAccountStatusRequest {
	return ApiAccountStatusRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AccountStatusResponse
func (a *AccountAPIService) AccountStatusExecute(r ApiAccountStatusRequest) (chan *common.ResponseOrRaw[models.AccountStatusResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.omitZeroBalances != nil {
		localVarQueryParams["omitZeroBalances"] = *r.omitZeroBalances
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/account.status"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.AccountStatusResponse](a.Ws, localPayload, sendParams)
}

type ApiAllOrderListsRequest struct {
	ApiService *AccountAPIService
	id         *string
	fromId     *int32
	startTime  *int64
	endTime    *int64
	limit      *int32
	recvWindow *float32
}

// Client-generated request identifier.
func (r ApiAllOrderListsRequest) Id(id string) ApiAllOrderListsRequest {
	r.id = &id
	return r
}

// Order list ID to begin at
func (r ApiAllOrderListsRequest) FromId(fromId int32) ApiAllOrderListsRequest {
	r.fromId = &fromId
	return r
}

// Timestamp in ms
func (r ApiAllOrderListsRequest) StartTime(startTime int64) ApiAllOrderListsRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms
func (r ApiAllOrderListsRequest) EndTime(endTime int64) ApiAllOrderListsRequest {
	r.endTime = &endTime
	return r
}

// Default: 500; Maximum: 1000
func (r ApiAllOrderListsRequest) Limit(limit int32) ApiAllOrderListsRequest {
	r.limit = &limit
	return r
}

// Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiAllOrderListsRequest) RecvWindow(recvWindow float32) ApiAllOrderListsRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiAllOrderListsRequest) Execute() (*common.ResponseOrRaw[models.AllOrderListsResponse], error) {
	respChan, errChan, err := r.ApiService.AllOrderListsExecute(r)
	if err != nil {
		return nil, err
	}

	select {
	case resp := <-respChan:
		return resp, nil
	case err := <-errChan:
		return nil, err
	}
}

func (r ApiAllOrderListsRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.AllOrderListsResponse], chan error, error) {
	return r.ApiService.AllOrderListsExecute(r)
}

/*
AllOrderLists Account order list history (USER_DATA)
/allOrderLists

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/account#all-order-lists

@param id Client-generated request identifier.	@param fromId Order list ID to begin at	@param startTime Timestamp in ms	@param endTime Timestamp in ms	@param limit Default: 500; Maximum: 1000	@param recvWindow Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiAllOrderListsRequest
*/
func (a *AccountAPIService) AllOrderLists() ApiAllOrderListsRequest {
	return ApiAllOrderListsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AllOrderListsResponse
func (a *AccountAPIService) AllOrderListsExecute(r ApiAllOrderListsRequest) (chan *common.ResponseOrRaw[models.AllOrderListsResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.fromId != nil {
		localVarQueryParams["fromId"] = *r.fromId
	}
	if r.startTime != nil {
		localVarQueryParams["startTime"] = *r.startTime
	}
	if r.endTime != nil {
		localVarQueryParams["endTime"] = *r.endTime
	}
	if r.limit != nil {
		localVarQueryParams["limit"] = *r.limit
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/allOrderLists"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.AllOrderListsResponse](a.Ws, localPayload, sendParams)
}

type ApiAllOrdersRequest struct {
	ApiService *AccountAPIService
	symbol     *string
	id         *string
	orderId    *int64
	startTime  *int64
	endTime    *int64
	limit      *int32
	recvWindow *float32
}

func (r ApiAllOrdersRequest) Symbol(symbol string) ApiAllOrdersRequest {
	r.symbol = &symbol
	return r
}

// Client-generated request identifier.
func (r ApiAllOrdersRequest) Id(id string) ApiAllOrdersRequest {
	r.id = &id
	return r
}

// Order ID to begin at
func (r ApiAllOrdersRequest) OrderId(orderId int64) ApiAllOrdersRequest {
	r.orderId = &orderId
	return r
}

// Timestamp in ms
func (r ApiAllOrdersRequest) StartTime(startTime int64) ApiAllOrdersRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms
func (r ApiAllOrdersRequest) EndTime(endTime int64) ApiAllOrdersRequest {
	r.endTime = &endTime
	return r
}

// Default: 500; Maximum: 1000
func (r ApiAllOrdersRequest) Limit(limit int32) ApiAllOrdersRequest {
	r.limit = &limit
	return r
}

// Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiAllOrdersRequest) RecvWindow(recvWindow float32) ApiAllOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiAllOrdersRequest) Execute() (*common.ResponseOrRaw[models.AllOrdersResponse], error) {
	respChan, errChan, err := r.ApiService.AllOrdersExecute(r)
	if err != nil {
		return nil, err
	}

	select {
	case resp := <-respChan:
		return resp, nil
	case err := <-errChan:
		return nil, err
	}
}

func (r ApiAllOrdersRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.AllOrdersResponse], chan error, error) {
	return r.ApiService.AllOrdersExecute(r)
}

/*
AllOrders Account order history (USER_DATA)
/allOrders

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/account#all-orders

@param symbol	@param id Client-generated request identifier.	@param orderId Order ID to begin at	@param startTime Timestamp in ms	@param endTime Timestamp in ms	@param limit Default: 500; Maximum: 1000	@param recvWindow Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiAllOrdersRequest
*/
func (a *AccountAPIService) AllOrders() ApiAllOrdersRequest {
	return ApiAllOrdersRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AllOrdersResponse
func (a *AccountAPIService) AllOrdersExecute(r ApiAllOrdersRequest) (chan *common.ResponseOrRaw[models.AllOrdersResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.orderId != nil {
		localVarQueryParams["orderId"] = *r.orderId
	}
	if r.startTime != nil {
		localVarQueryParams["startTime"] = *r.startTime
	}
	if r.endTime != nil {
		localVarQueryParams["endTime"] = *r.endTime
	}
	if r.limit != nil {
		localVarQueryParams["limit"] = *r.limit
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/allOrders"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.AllOrdersResponse](a.Ws, localPayload, sendParams)
}

type ApiMyAllocationsRequest struct {
	ApiService       *AccountAPIService
	symbol           *string
	id               *string
	startTime        *int64
	endTime          *int64
	fromAllocationId *int32
	limit            *int32
	orderId          *int64
	recvWindow       *float32
}

func (r ApiMyAllocationsRequest) Symbol(symbol string) ApiMyAllocationsRequest {
	r.symbol = &symbol
	return r
}

// Client-generated request identifier.
func (r ApiMyAllocationsRequest) Id(id string) ApiMyAllocationsRequest {
	r.id = &id
	return r
}

// Timestamp in ms
func (r ApiMyAllocationsRequest) StartTime(startTime int64) ApiMyAllocationsRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms
func (r ApiMyAllocationsRequest) EndTime(endTime int64) ApiMyAllocationsRequest {
	r.endTime = &endTime
	return r
}

// Allocation ID to begin at
func (r ApiMyAllocationsRequest) FromAllocationId(fromAllocationId int32) ApiMyAllocationsRequest {
	r.fromAllocationId = &fromAllocationId
	return r
}

// Default: 500; Maximum: 1000
func (r ApiMyAllocationsRequest) Limit(limit int32) ApiMyAllocationsRequest {
	r.limit = &limit
	return r
}

// Order ID
func (r ApiMyAllocationsRequest) OrderId(orderId int64) ApiMyAllocationsRequest {
	r.orderId = &orderId
	return r
}

// Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiMyAllocationsRequest) RecvWindow(recvWindow float32) ApiMyAllocationsRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiMyAllocationsRequest) Execute() (*common.ResponseOrRaw[models.MyAllocationsResponse], error) {
	respChan, errChan, err := r.ApiService.MyAllocationsExecute(r)
	if err != nil {
		return nil, err
	}

	select {
	case resp := <-respChan:
		return resp, nil
	case err := <-errChan:
		return nil, err
	}
}

func (r ApiMyAllocationsRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.MyAllocationsResponse], chan error, error) {
	return r.ApiService.MyAllocationsExecute(r)
}

/*
MyAllocations Account allocations (USER_DATA)
/myAllocations

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/account#my-allocations

@param symbol	@param id Client-generated request identifier.	@param startTime Timestamp in ms	@param endTime Timestamp in ms	@param fromAllocationId Allocation ID to begin at	@param limit Default: 500; Maximum: 1000	@param orderId Order ID	@param recvWindow Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiMyAllocationsRequest
*/
func (a *AccountAPIService) MyAllocations() ApiMyAllocationsRequest {
	return ApiMyAllocationsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return MyAllocationsResponse
func (a *AccountAPIService) MyAllocationsExecute(r ApiMyAllocationsRequest) (chan *common.ResponseOrRaw[models.MyAllocationsResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.startTime != nil {
		localVarQueryParams["startTime"] = *r.startTime
	}
	if r.endTime != nil {
		localVarQueryParams["endTime"] = *r.endTime
	}
	if r.fromAllocationId != nil {
		localVarQueryParams["fromAllocationId"] = *r.fromAllocationId
	}
	if r.limit != nil {
		localVarQueryParams["limit"] = *r.limit
	}
	if r.orderId != nil {
		localVarQueryParams["orderId"] = *r.orderId
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/myAllocations"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.MyAllocationsResponse](a.Ws, localPayload, sendParams)
}

type ApiMyFiltersRequest struct {
	ApiService *AccountAPIService
	symbol     *string
	id         *string
	recvWindow *float32
}

func (r ApiMyFiltersRequest) Symbol(symbol string) ApiMyFiltersRequest {
	r.symbol = &symbol
	return r
}

// Client-generated request identifier.
func (r ApiMyFiltersRequest) Id(id string) ApiMyFiltersRequest {
	r.id = &id
	return r
}

// Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiMyFiltersRequest) RecvWindow(recvWindow float32) ApiMyFiltersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiMyFiltersRequest) Execute() (*common.ResponseOrRaw[models.MyFiltersResponse], error) {
	respChan, errChan, err := r.ApiService.MyFiltersExecute(r)
	if err != nil {
		return nil, err
	}

	select {
	case resp := <-respChan:
		return resp, nil
	case err := <-errChan:
		return nil, err
	}
}

func (r ApiMyFiltersRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.MyFiltersResponse], chan error, error) {
	return r.ApiService.MyFiltersExecute(r)
}

/*
MyFilters Query Relevant Filters (USER_DATA)
/myFilters

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/account#my-filters

@param symbol	@param id Client-generated request identifier.	@param recvWindow Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiMyFiltersRequest
*/
func (a *AccountAPIService) MyFilters() ApiMyFiltersRequest {
	return ApiMyFiltersRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return MyFiltersResponse
func (a *AccountAPIService) MyFiltersExecute(r ApiMyFiltersRequest) (chan *common.ResponseOrRaw[models.MyFiltersResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/myFilters"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.MyFiltersResponse](a.Ws, localPayload, sendParams)
}

type ApiMyPreventedMatchesRequest struct {
	ApiService           *AccountAPIService
	symbol               *string
	id                   *string
	preventedMatchId     *int64
	orderId              *int64
	fromPreventedMatchId *int64
	limit                *int32
	recvWindow           *float32
}

func (r ApiMyPreventedMatchesRequest) Symbol(symbol string) ApiMyPreventedMatchesRequest {
	r.symbol = &symbol
	return r
}

// Client-generated request identifier.
func (r ApiMyPreventedMatchesRequest) Id(id string) ApiMyPreventedMatchesRequest {
	r.id = &id
	return r
}

// Prevented match ID
func (r ApiMyPreventedMatchesRequest) PreventedMatchId(preventedMatchId int64) ApiMyPreventedMatchesRequest {
	r.preventedMatchId = &preventedMatchId
	return r
}

// Order ID
func (r ApiMyPreventedMatchesRequest) OrderId(orderId int64) ApiMyPreventedMatchesRequest {
	r.orderId = &orderId
	return r
}

// Prevented match ID to begin at
func (r ApiMyPreventedMatchesRequest) FromPreventedMatchId(fromPreventedMatchId int64) ApiMyPreventedMatchesRequest {
	r.fromPreventedMatchId = &fromPreventedMatchId
	return r
}

// Default: 500; Maximum: 1000
func (r ApiMyPreventedMatchesRequest) Limit(limit int32) ApiMyPreventedMatchesRequest {
	r.limit = &limit
	return r
}

// Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiMyPreventedMatchesRequest) RecvWindow(recvWindow float32) ApiMyPreventedMatchesRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiMyPreventedMatchesRequest) Execute() (*common.ResponseOrRaw[models.MyPreventedMatchesResponse], error) {
	respChan, errChan, err := r.ApiService.MyPreventedMatchesExecute(r)
	if err != nil {
		return nil, err
	}

	select {
	case resp := <-respChan:
		return resp, nil
	case err := <-errChan:
		return nil, err
	}
}

func (r ApiMyPreventedMatchesRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.MyPreventedMatchesResponse], chan error, error) {
	return r.ApiService.MyPreventedMatchesExecute(r)
}

/*
MyPreventedMatches Account prevented matches (USER_DATA)
/myPreventedMatches

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/account#my-prevented-matches

@param symbol	@param id Client-generated request identifier.	@param preventedMatchId Prevented match ID	@param orderId Order ID	@param fromPreventedMatchId Prevented match ID to begin at	@param limit Default: 500; Maximum: 1000	@param recvWindow Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiMyPreventedMatchesRequest
*/
func (a *AccountAPIService) MyPreventedMatches() ApiMyPreventedMatchesRequest {
	return ApiMyPreventedMatchesRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return MyPreventedMatchesResponse
func (a *AccountAPIService) MyPreventedMatchesExecute(r ApiMyPreventedMatchesRequest) (chan *common.ResponseOrRaw[models.MyPreventedMatchesResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.preventedMatchId != nil {
		localVarQueryParams["preventedMatchId"] = *r.preventedMatchId
	}
	if r.orderId != nil {
		localVarQueryParams["orderId"] = *r.orderId
	}
	if r.fromPreventedMatchId != nil {
		localVarQueryParams["fromPreventedMatchId"] = *r.fromPreventedMatchId
	}
	if r.limit != nil {
		localVarQueryParams["limit"] = *r.limit
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/myPreventedMatches"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.MyPreventedMatchesResponse](a.Ws, localPayload, sendParams)
}

type ApiMyTradesRequest struct {
	ApiService *AccountAPIService
	symbol     *string
	id         *string
	orderId    *int64
	startTime  *int64
	endTime    *int64
	fromId     *int32
	limit      *int32
	recvWindow *float32
}

func (r ApiMyTradesRequest) Symbol(symbol string) ApiMyTradesRequest {
	r.symbol = &symbol
	return r
}

// Client-generated request identifier.
func (r ApiMyTradesRequest) Id(id string) ApiMyTradesRequest {
	r.id = &id
	return r
}

// This can only be used in combination with &#x60;symbol&#x60;.
func (r ApiMyTradesRequest) OrderId(orderId int64) ApiMyTradesRequest {
	r.orderId = &orderId
	return r
}

// Timestamp in ms
func (r ApiMyTradesRequest) StartTime(startTime int64) ApiMyTradesRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms
func (r ApiMyTradesRequest) EndTime(endTime int64) ApiMyTradesRequest {
	r.endTime = &endTime
	return r
}

// First trade ID to query
func (r ApiMyTradesRequest) FromId(fromId int32) ApiMyTradesRequest {
	r.fromId = &fromId
	return r
}

// Default: 500; Maximum: 1000
func (r ApiMyTradesRequest) Limit(limit int32) ApiMyTradesRequest {
	r.limit = &limit
	return r
}

// Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiMyTradesRequest) RecvWindow(recvWindow float32) ApiMyTradesRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiMyTradesRequest) Execute() (*common.ResponseOrRaw[models.MyTradesResponse], error) {
	respChan, errChan, err := r.ApiService.MyTradesExecute(r)
	if err != nil {
		return nil, err
	}

	select {
	case resp := <-respChan:
		return resp, nil
	case err := <-errChan:
		return nil, err
	}
}

func (r ApiMyTradesRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.MyTradesResponse], chan error, error) {
	return r.ApiService.MyTradesExecute(r)
}

/*
MyTrades Account trade history (USER_DATA)
/myTrades

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/account#my-trades

@param symbol	@param id Client-generated request identifier.	@param orderId This can only be used in combination with `symbol`.	@param startTime Timestamp in ms	@param endTime Timestamp in ms	@param fromId First trade ID to query	@param limit Default: 500; Maximum: 1000	@param recvWindow Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiMyTradesRequest
*/
func (a *AccountAPIService) MyTrades() ApiMyTradesRequest {
	return ApiMyTradesRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return MyTradesResponse
func (a *AccountAPIService) MyTradesExecute(r ApiMyTradesRequest) (chan *common.ResponseOrRaw[models.MyTradesResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.orderId != nil {
		localVarQueryParams["orderId"] = *r.orderId
	}
	if r.startTime != nil {
		localVarQueryParams["startTime"] = *r.startTime
	}
	if r.endTime != nil {
		localVarQueryParams["endTime"] = *r.endTime
	}
	if r.fromId != nil {
		localVarQueryParams["fromId"] = *r.fromId
	}
	if r.limit != nil {
		localVarQueryParams["limit"] = *r.limit
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/myTrades"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.MyTradesResponse](a.Ws, localPayload, sendParams)
}

type ApiOpenOrderListsStatusRequest struct {
	ApiService *AccountAPIService
	id         *string
	recvWindow *float32
}

// Client-generated request identifier.
func (r ApiOpenOrderListsStatusRequest) Id(id string) ApiOpenOrderListsStatusRequest {
	r.id = &id
	return r
}

// Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOpenOrderListsStatusRequest) RecvWindow(recvWindow float32) ApiOpenOrderListsStatusRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOpenOrderListsStatusRequest) Execute() (*common.ResponseOrRaw[models.OpenOrderListsStatusResponse], error) {
	respChan, errChan, err := r.ApiService.OpenOrderListsStatusExecute(r)
	if err != nil {
		return nil, err
	}

	select {
	case resp := <-respChan:
		return resp, nil
	case err := <-errChan:
		return nil, err
	}
}

func (r ApiOpenOrderListsStatusRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.OpenOrderListsStatusResponse], chan error, error) {
	return r.ApiService.OpenOrderListsStatusExecute(r)
}

/*
OpenOrderListsStatus Current open Order lists (USER_DATA)
/openOrderLists.status

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/account#open-order-lists-status

@param id Client-generated request identifier.	@param recvWindow Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiOpenOrderListsStatusRequest
*/
func (a *AccountAPIService) OpenOrderListsStatus() ApiOpenOrderListsStatusRequest {
	return ApiOpenOrderListsStatusRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return OpenOrderListsStatusResponse
func (a *AccountAPIService) OpenOrderListsStatusExecute(r ApiOpenOrderListsStatusRequest) (chan *common.ResponseOrRaw[models.OpenOrderListsStatusResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/openOrderLists.status"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.OpenOrderListsStatusResponse](a.Ws, localPayload, sendParams)
}

type ApiOpenOrdersStatusRequest struct {
	ApiService *AccountAPIService
	id         *string
	symbol     *string
	recvWindow *float32
}

// Client-generated request identifier.
func (r ApiOpenOrdersStatusRequest) Id(id string) ApiOpenOrdersStatusRequest {
	r.id = &id
	return r
}

// If omitted, open orders for all symbols are returned
func (r ApiOpenOrdersStatusRequest) Symbol(symbol string) ApiOpenOrdersStatusRequest {
	r.symbol = &symbol
	return r
}

// Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOpenOrdersStatusRequest) RecvWindow(recvWindow float32) ApiOpenOrdersStatusRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOpenOrdersStatusRequest) Execute() (*common.ResponseOrRaw[models.OpenOrdersStatusResponse], error) {
	respChan, errChan, err := r.ApiService.OpenOrdersStatusExecute(r)
	if err != nil {
		return nil, err
	}

	select {
	case resp := <-respChan:
		return resp, nil
	case err := <-errChan:
		return nil, err
	}
}

func (r ApiOpenOrdersStatusRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.OpenOrdersStatusResponse], chan error, error) {
	return r.ApiService.OpenOrdersStatusExecute(r)
}

/*
OpenOrdersStatus Current open orders (USER_DATA)
/openOrders.status

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/account#open-orders-status

@param id Client-generated request identifier.	@param symbol If omitted, open orders for all symbols are returned	@param recvWindow Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiOpenOrdersStatusRequest
*/
func (a *AccountAPIService) OpenOrdersStatus() ApiOpenOrdersStatusRequest {
	return ApiOpenOrdersStatusRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return OpenOrdersStatusResponse
func (a *AccountAPIService) OpenOrdersStatusExecute(r ApiOpenOrdersStatusRequest) (chan *common.ResponseOrRaw[models.OpenOrdersStatusResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.symbol != nil {
		localVarQueryParams["symbol"] = *r.symbol
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/openOrders.status"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.OpenOrdersStatusResponse](a.Ws, localPayload, sendParams)
}

type ApiOrderAmendmentsRequest struct {
	ApiService      *AccountAPIService
	symbol          *string
	orderId         *int64
	id              *string
	fromExecutionId *int64
	limit           *int64
	recvWindow      *float32
}

func (r ApiOrderAmendmentsRequest) Symbol(symbol string) ApiOrderAmendmentsRequest {
	r.symbol = &symbol
	return r
}

// Order ID
func (r ApiOrderAmendmentsRequest) OrderId(orderId int64) ApiOrderAmendmentsRequest {
	r.orderId = &orderId
	return r
}

// Client-generated request identifier.
func (r ApiOrderAmendmentsRequest) Id(id string) ApiOrderAmendmentsRequest {
	r.id = &id
	return r
}

// Execution ID to begin at
func (r ApiOrderAmendmentsRequest) FromExecutionId(fromExecutionId int64) ApiOrderAmendmentsRequest {
	r.fromExecutionId = &fromExecutionId
	return r
}

// Default: 500; Maximum: 1000
func (r ApiOrderAmendmentsRequest) Limit(limit int64) ApiOrderAmendmentsRequest {
	r.limit = &limit
	return r
}

// Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderAmendmentsRequest) RecvWindow(recvWindow float32) ApiOrderAmendmentsRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOrderAmendmentsRequest) Execute() (*common.ResponseOrRaw[models.OrderAmendmentsResponse], error) {
	respChan, errChan, err := r.ApiService.OrderAmendmentsExecute(r)
	if err != nil {
		return nil, err
	}

	select {
	case resp := <-respChan:
		return resp, nil
	case err := <-errChan:
		return nil, err
	}
}

func (r ApiOrderAmendmentsRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.OrderAmendmentsResponse], chan error, error) {
	return r.ApiService.OrderAmendmentsExecute(r)
}

/*
OrderAmendments Query Order Amendments (USER_DATA)
/order.amendments

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/account#order-amendments

@param symbol	@param orderId Order ID	@param id Client-generated request identifier.	@param fromExecutionId Execution ID to begin at	@param limit Default: 500; Maximum: 1000	@param recvWindow Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiOrderAmendmentsRequest
*/
func (a *AccountAPIService) OrderAmendments() ApiOrderAmendmentsRequest {
	return ApiOrderAmendmentsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return OrderAmendmentsResponse
func (a *AccountAPIService) OrderAmendmentsExecute(r ApiOrderAmendmentsRequest) (chan *common.ResponseOrRaw[models.OrderAmendmentsResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.orderId == nil {
		return nil, nil, common.ReportError("orderId is required and must be specified")
	}
	localVarQueryParams["orderId"] = *r.orderId

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.fromExecutionId != nil {
		localVarQueryParams["fromExecutionId"] = *r.fromExecutionId
	}
	if r.limit != nil {
		localVarQueryParams["limit"] = *r.limit
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/order.amendments"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.OrderAmendmentsResponse](a.Ws, localPayload, sendParams)
}

type ApiOrderListStatusRequest struct {
	ApiService        *AccountAPIService
	id                *string
	origClientOrderId *string
	orderListId       *int32
	recvWindow        *float32
}

// Client-generated request identifier.
func (r ApiOrderListStatusRequest) Id(id string) ApiOrderListStatusRequest {
	r.id = &id
	return r
}

// Query order list by &#x60;listClientOrderId&#x60;. &#x60;orderListId&#x60; or &#x60;origClientOrderId&#x60; must be provided.
func (r ApiOrderListStatusRequest) OrigClientOrderId(origClientOrderId string) ApiOrderListStatusRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

// Query order list by &#x60;orderListId&#x60;. &#x60;orderListId&#x60; or &#x60;origClientOrderId&#x60; must be provided.
func (r ApiOrderListStatusRequest) OrderListId(orderListId int32) ApiOrderListStatusRequest {
	r.orderListId = &orderListId
	return r
}

// Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderListStatusRequest) RecvWindow(recvWindow float32) ApiOrderListStatusRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOrderListStatusRequest) Execute() (*common.ResponseOrRaw[models.OrderListStatusResponse], error) {
	respChan, errChan, err := r.ApiService.OrderListStatusExecute(r)
	if err != nil {
		return nil, err
	}

	select {
	case resp := <-respChan:
		return resp, nil
	case err := <-errChan:
		return nil, err
	}
}

func (r ApiOrderListStatusRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.OrderListStatusResponse], chan error, error) {
	return r.ApiService.OrderListStatusExecute(r)
}

/*
OrderListStatus Query Order list (USER_DATA)
/orderList.status

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/account#order-list-status

@param id Client-generated request identifier.	@param origClientOrderId Query order list by `listClientOrderId`. `orderListId` or `origClientOrderId` must be provided.	@param orderListId Query order list by `orderListId`. `orderListId` or `origClientOrderId` must be provided.	@param recvWindow Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiOrderListStatusRequest
*/
func (a *AccountAPIService) OrderListStatus() ApiOrderListStatusRequest {
	return ApiOrderListStatusRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return OrderListStatusResponse
func (a *AccountAPIService) OrderListStatusExecute(r ApiOrderListStatusRequest) (chan *common.ResponseOrRaw[models.OrderListStatusResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.origClientOrderId != nil {
		localVarQueryParams["origClientOrderId"] = *r.origClientOrderId
	}
	if r.orderListId != nil {
		localVarQueryParams["orderListId"] = *r.orderListId
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/orderList.status"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.OrderListStatusResponse](a.Ws, localPayload, sendParams)
}

type ApiOrderStatusRequest struct {
	ApiService        *AccountAPIService
	symbol            *string
	id                *string
	orderId           *int64
	origClientOrderId *string
	recvWindow        *float32
}

func (r ApiOrderStatusRequest) Symbol(symbol string) ApiOrderStatusRequest {
	r.symbol = &symbol
	return r
}

// Client-generated request identifier.
func (r ApiOrderStatusRequest) Id(id string) ApiOrderStatusRequest {
	r.id = &id
	return r
}

// Lookup order by &#x60;orderId&#x60;
func (r ApiOrderStatusRequest) OrderId(orderId int64) ApiOrderStatusRequest {
	r.orderId = &orderId
	return r
}

// Lookup order by &#x60;clientOrderId&#x60;
func (r ApiOrderStatusRequest) OrigClientOrderId(origClientOrderId string) ApiOrderStatusRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

// Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderStatusRequest) RecvWindow(recvWindow float32) ApiOrderStatusRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOrderStatusRequest) Execute() (*common.ResponseOrRaw[models.OrderStatusResponse], error) {
	respChan, errChan, err := r.ApiService.OrderStatusExecute(r)
	if err != nil {
		return nil, err
	}

	select {
	case resp := <-respChan:
		return resp, nil
	case err := <-errChan:
		return nil, err
	}
}

func (r ApiOrderStatusRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.OrderStatusResponse], chan error, error) {
	return r.ApiService.OrderStatusExecute(r)
}

/*
OrderStatus Query order (USER_DATA)
/order.status

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/account#order-status

@param symbol	@param id Client-generated request identifier.	@param orderId Lookup order by `orderId`	@param origClientOrderId Lookup order by `clientOrderId`	@param recvWindow Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiOrderStatusRequest
*/
func (a *AccountAPIService) OrderStatus() ApiOrderStatusRequest {
	return ApiOrderStatusRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return OrderStatusResponse
func (a *AccountAPIService) OrderStatusExecute(r ApiOrderStatusRequest) (chan *common.ResponseOrRaw[models.OrderStatusResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.orderId != nil {
		localVarQueryParams["orderId"] = *r.orderId
	}
	if r.origClientOrderId != nil {
		localVarQueryParams["origClientOrderId"] = *r.origClientOrderId
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/order.status"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.OrderStatusResponse](a.Ws, localPayload, sendParams)
}
