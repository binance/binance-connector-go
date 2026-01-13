/*
Binance Derivatives Trading COIN Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket API
*/

package binancederivativestradingcoinfutureswebsocketapi

import (
	"github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/websocketapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// UserDataStreamsAPIService UserDataStreamsAPI Service
type UserDataStreamsAPIService struct {
	Ws *common.WebsocketAPI
}

type ApiCloseUserDataStreamRequest struct {
	ApiService *UserDataStreamsAPIService
	id         *string
}

// Unique WebSocket request ID.
func (r ApiCloseUserDataStreamRequest) Id(id string) ApiCloseUserDataStreamRequest {
	r.id = &id
	return r
}

func (r ApiCloseUserDataStreamRequest) Execute() (*common.ResponseOrRaw[models.CloseUserDataStreamResponse], error) {
	respChan, errChan, err := r.ApiService.CloseUserDataStreamExecute(r)
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

func (r ApiCloseUserDataStreamRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.CloseUserDataStreamResponse], chan error, error) {
	return r.ApiService.CloseUserDataStreamExecute(r)
}

/*
CloseUserDataStream Close User Data Stream (USER_STREAM)
/userDataStream.stop

https://developers.binance.com/docs/derivatives/coin-margined-futures/user-data-streams/Close-User-Data-Stream-Wsp

@param id Unique WebSocket request ID.
@return ApiCloseUserDataStreamRequest
*/
func (a *UserDataStreamsAPIService) CloseUserDataStream() ApiCloseUserDataStreamRequest {
	return ApiCloseUserDataStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return CloseUserDataStreamResponse
func (a *UserDataStreamsAPIService) CloseUserDataStreamExecute(r ApiCloseUserDataStreamRequest) (chan *common.ResponseOrRaw[models.CloseUserDataStreamResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}

	localPayload := map[string]any{
		"method": "/userDataStream.stop"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           false,
		WithAPIKey:       true,
		WithSessionLogon: false,
	}

	return SendMessage[models.CloseUserDataStreamResponse](a.Ws, localPayload, sendParams)
}

type ApiKeepaliveUserDataStreamRequest struct {
	ApiService *UserDataStreamsAPIService
	id         *string
}

// Unique WebSocket request ID.
func (r ApiKeepaliveUserDataStreamRequest) Id(id string) ApiKeepaliveUserDataStreamRequest {
	r.id = &id
	return r
}

func (r ApiKeepaliveUserDataStreamRequest) Execute() (*common.ResponseOrRaw[models.KeepaliveUserDataStreamResponse], error) {
	respChan, errChan, err := r.ApiService.KeepaliveUserDataStreamExecute(r)
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

func (r ApiKeepaliveUserDataStreamRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.KeepaliveUserDataStreamResponse], chan error, error) {
	return r.ApiService.KeepaliveUserDataStreamExecute(r)
}

/*
KeepaliveUserDataStream Keepalive User Data Stream (USER_STREAM)
/userDataStream.ping

https://developers.binance.com/docs/derivatives/coin-margined-futures/user-data-streams/Keepalive-User-Data-Stream-Wsp

@param id Unique WebSocket request ID.
@return ApiKeepaliveUserDataStreamRequest
*/
func (a *UserDataStreamsAPIService) KeepaliveUserDataStream() ApiKeepaliveUserDataStreamRequest {
	return ApiKeepaliveUserDataStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return KeepaliveUserDataStreamResponse
func (a *UserDataStreamsAPIService) KeepaliveUserDataStreamExecute(r ApiKeepaliveUserDataStreamRequest) (chan *common.ResponseOrRaw[models.KeepaliveUserDataStreamResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}

	localPayload := map[string]any{
		"method": "/userDataStream.ping"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           false,
		WithAPIKey:       true,
		WithSessionLogon: false,
	}

	return SendMessage[models.KeepaliveUserDataStreamResponse](a.Ws, localPayload, sendParams)
}

type ApiStartUserDataStreamRequest struct {
	ApiService *UserDataStreamsAPIService
	id         *string
}

// Unique WebSocket request ID.
func (r ApiStartUserDataStreamRequest) Id(id string) ApiStartUserDataStreamRequest {
	r.id = &id
	return r
}

func (r ApiStartUserDataStreamRequest) Execute() (*common.ResponseOrRaw[models.StartUserDataStreamResponse], error) {
	respChan, errChan, err := r.ApiService.StartUserDataStreamExecute(r)
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

func (r ApiStartUserDataStreamRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.StartUserDataStreamResponse], chan error, error) {
	return r.ApiService.StartUserDataStreamExecute(r)
}

/*
StartUserDataStream Start User Data Stream (USER_STREAM)
/userDataStream.start

https://developers.binance.com/docs/derivatives/coin-margined-futures/user-data-streams/Start-User-Data-Stream-Wsp

@param id Unique WebSocket request ID.
@return ApiStartUserDataStreamRequest
*/
func (a *UserDataStreamsAPIService) StartUserDataStream() ApiStartUserDataStreamRequest {
	return ApiStartUserDataStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return StartUserDataStreamResponse
func (a *UserDataStreamsAPIService) StartUserDataStreamExecute(r ApiStartUserDataStreamRequest) (chan *common.ResponseOrRaw[models.StartUserDataStreamResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}

	localPayload := map[string]any{
		"method": "/userDataStream.start"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           false,
		WithAPIKey:       true,
		WithSessionLogon: false,
	}

	return SendMessage[models.StartUserDataStreamResponse](a.Ws, localPayload, sendParams)
}
