/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)

API version: 1.0.0
*/

package binancespotwebsocketapi

import (
	"github.com/binance/binance-connector-go/clients/spot/src/websocketapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// GeneralAPIService GeneralAPI Service
type GeneralAPIService struct {
	Ws *common.WebsocketAPI
}

type ApiExchangeInfoRequest struct {
	ApiService         *GeneralAPIService
	id                 *string
	symbol             *string
	symbols            *[]string
	permissions        *[]string
	showPermissionSets *bool
	symbolStatus       *models.ExchangeInfoSymbolStatusParameter
}

// Unique WebSocket request ID.
func (r ApiExchangeInfoRequest) Id(id string) ApiExchangeInfoRequest {
	r.id = &id
	return r
}

// Describe a single symbol
func (r ApiExchangeInfoRequest) Symbol(symbol string) ApiExchangeInfoRequest {
	r.symbol = &symbol
	return r
}

// List of symbols to query
func (r ApiExchangeInfoRequest) Symbols(symbols []string) ApiExchangeInfoRequest {
	r.symbols = &symbols
	return r
}

func (r ApiExchangeInfoRequest) Permissions(permissions []string) ApiExchangeInfoRequest {
	r.permissions = &permissions
	return r
}

func (r ApiExchangeInfoRequest) ShowPermissionSets(showPermissionSets bool) ApiExchangeInfoRequest {
	r.showPermissionSets = &showPermissionSets
	return r
}

func (r ApiExchangeInfoRequest) SymbolStatus(symbolStatus models.ExchangeInfoSymbolStatusParameter) ApiExchangeInfoRequest {
	r.symbolStatus = &symbolStatus
	return r
}

func (r ApiExchangeInfoRequest) Execute() (*common.ResponseOrRaw[models.ExchangeInfoResponse], error) {
	respChan, errChan, err := r.ApiService.ExchangeInfoExecute(r)
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

func (r ApiExchangeInfoRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.ExchangeInfoResponse], chan error, error) {
	return r.ApiService.ExchangeInfoExecute(r)
}

/*
ExchangeInfo WebSocket Exchange information
/exchangeInfo

https://developers.binance.com/docs/binance-spot-api-docs/websocket-api/general-requests#exchange-information

@param id Unique WebSocket request ID.	@param symbol Describe a single symbol	@param symbols List of symbols to query	@param permissions	@param showPermissionSets	@param symbolStatus
@return ApiExchangeInfoRequest
*/
func (a *GeneralAPIService) ExchangeInfo() ApiExchangeInfoRequest {
	return ApiExchangeInfoRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return ExchangeInfoResponse
func (a *GeneralAPIService) ExchangeInfoExecute(r ApiExchangeInfoRequest) (chan *common.ResponseOrRaw[models.ExchangeInfoResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.symbol != nil {
		localVarQueryParams["symbol"] = *r.symbol
	}
	if r.symbols != nil {
		localVarQueryParams["symbols"] = *r.symbols
	}
	if r.permissions != nil {
		localVarQueryParams["permissions"] = *r.permissions
	}
	if r.showPermissionSets != nil {
		localVarQueryParams["showPermissionSets"] = *r.showPermissionSets
	}
	if r.symbolStatus != nil {
		localVarQueryParams["symbolStatus"] = *r.symbolStatus
	}

	localPayload := map[string]any{
		"method": "/exchangeInfo"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           false,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.ExchangeInfoResponse](a.Ws, localPayload, sendParams)
}

type ApiPingRequest struct {
	ApiService *GeneralAPIService
	id         *string
}

// Unique WebSocket request ID.
func (r ApiPingRequest) Id(id string) ApiPingRequest {
	r.id = &id
	return r
}

func (r ApiPingRequest) Execute() (*common.ResponseOrRaw[models.PingResponse], error) {
	respChan, errChan, err := r.ApiService.PingExecute(r)
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

func (r ApiPingRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.PingResponse], chan error, error) {
	return r.ApiService.PingExecute(r)
}

/*
Ping WebSocket Test connectivity
/ping

https://developers.binance.com/docs/binance-spot-api-docs/websocket-api/general-requests#test-connectivity

@param id Unique WebSocket request ID.
@return ApiPingRequest
*/
func (a *GeneralAPIService) Ping() ApiPingRequest {
	return ApiPingRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return PingResponse
func (a *GeneralAPIService) PingExecute(r ApiPingRequest) (chan *common.ResponseOrRaw[models.PingResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}

	localPayload := map[string]any{
		"method": "/ping"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           false,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.PingResponse](a.Ws, localPayload, sendParams)
}

type ApiTimeRequest struct {
	ApiService *GeneralAPIService
	id         *string
}

// Unique WebSocket request ID.
func (r ApiTimeRequest) Id(id string) ApiTimeRequest {
	r.id = &id
	return r
}

func (r ApiTimeRequest) Execute() (*common.ResponseOrRaw[models.TimeResponse], error) {
	respChan, errChan, err := r.ApiService.TimeExecute(r)
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

func (r ApiTimeRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.TimeResponse], chan error, error) {
	return r.ApiService.TimeExecute(r)
}

/*
Time WebSocket Check server time
/time

https://developers.binance.com/docs/binance-spot-api-docs/websocket-api/general-requests#check-server-time

@param id Unique WebSocket request ID.
@return ApiTimeRequest
*/
func (a *GeneralAPIService) Time() ApiTimeRequest {
	return ApiTimeRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return TimeResponse
func (a *GeneralAPIService) TimeExecute(r ApiTimeRequest) (chan *common.ResponseOrRaw[models.TimeResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}

	localPayload := map[string]any{
		"method": "/time"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           false,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.TimeResponse](a.Ws, localPayload, sendParams)
}
