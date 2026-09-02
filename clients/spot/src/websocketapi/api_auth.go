/*
Spot WebSocket API

Access market data, manage accounts, and trade on Binance Spot.
*/

package binancespotwebsocketapi

import (
	"github.com/binance/binance-connector-go/clients/spot/src/websocketapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// AuthAPIService AuthAPI Service
type AuthAPIService struct {
	Ws *common.WebsocketAPI
}

type ApiSessionLogonRequest struct {
	ApiService *AuthAPIService
	id         *string
	recvWindow *float64
}

// Client-generated request identifier.
func (r ApiSessionLogonRequest) Id(id string) ApiSessionLogonRequest {
	r.id = &id
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiSessionLogonRequest) RecvWindow(recvWindow float64) ApiSessionLogonRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSessionLogonRequest) Execute() (*common.ResponseOrRaw[models.SessionLogonResponse], error) {
	respChan, errChan, err := r.ApiService.SessionLogonExecute(r)
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

func (r ApiSessionLogonRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.SessionLogonResponse], chan error, error) {
	return r.ApiService.SessionLogonExecute(r)
}

/*
SessionLogon Log in with API key (USER_DATA)
/session.logon

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/auth#session-logon

@param id Client-generated request identifier.	@param recvWindow The value cannot be greater than `60000`. Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiSessionLogonRequest
*/
func (a *AuthAPIService) SessionLogon() ApiSessionLogonRequest {
	return ApiSessionLogonRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return SessionLogonResponse
func (a *AuthAPIService) SessionLogonExecute(r ApiSessionLogonRequest) (chan *common.ResponseOrRaw[models.SessionLogonResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/session.logon"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: true,
	}

	return SendMessage[models.SessionLogonResponse](a.Ws, localPayload, sendParams)
}

type ApiSessionLogoutRequest struct {
	ApiService *AuthAPIService
	id         *string
}

// Client-generated request identifier.
func (r ApiSessionLogoutRequest) Id(id string) ApiSessionLogoutRequest {
	r.id = &id
	return r
}

func (r ApiSessionLogoutRequest) Execute() (*common.ResponseOrRaw[models.SessionLogoutResponse], error) {
	respChan, errChan, err := r.ApiService.SessionLogoutExecute(r)
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

func (r ApiSessionLogoutRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.SessionLogoutResponse], chan error, error) {
	return r.ApiService.SessionLogoutExecute(r)
}

/*
SessionLogout Log out of the session
/session.logout

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/auth#session-logout

@param id Client-generated request identifier.
@return ApiSessionLogoutRequest
*/
func (a *AuthAPIService) SessionLogout() ApiSessionLogoutRequest {
	return ApiSessionLogoutRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return SessionLogoutResponse
func (a *AuthAPIService) SessionLogoutExecute(r ApiSessionLogoutRequest) (chan *common.ResponseOrRaw[models.SessionLogoutResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}

	localPayload := map[string]any{
		"method": "/session.logout"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           false,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.SessionLogoutResponse](a.Ws, localPayload, sendParams)
}

type ApiSessionStatusRequest struct {
	ApiService *AuthAPIService
	id         *string
}

// Client-generated request identifier.
func (r ApiSessionStatusRequest) Id(id string) ApiSessionStatusRequest {
	r.id = &id
	return r
}

func (r ApiSessionStatusRequest) Execute() (*common.ResponseOrRaw[models.SessionStatusResponse], error) {
	respChan, errChan, err := r.ApiService.SessionStatusExecute(r)
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

func (r ApiSessionStatusRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.SessionStatusResponse], chan error, error) {
	return r.ApiService.SessionStatusExecute(r)
}

/*
SessionStatus Query session status
/session.status

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/auth#session-status

@param id Client-generated request identifier.
@return ApiSessionStatusRequest
*/
func (a *AuthAPIService) SessionStatus() ApiSessionStatusRequest {
	return ApiSessionStatusRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return SessionStatusResponse
func (a *AuthAPIService) SessionStatusExecute(r ApiSessionStatusRequest) (chan *common.ResponseOrRaw[models.SessionStatusResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}

	localPayload := map[string]any{
		"method": "/session.status"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           false,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.SessionStatusResponse](a.Ws, localPayload, sendParams)
}
