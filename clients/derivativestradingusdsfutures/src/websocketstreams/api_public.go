/*
Binance Derivatives Trading USDS Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket Market Streams
*/

package binancederivativestradingusdsfutureswebsocketstreams

import (
	"strconv"

	"github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// PublicAPIService PublicAPI Service
type PublicAPIService Service

type ApiAllBookTickersStreamRequest struct {
	ApiService *PublicAPIService
	id         *string
}

// Unique WebSocket request ID.
func (r ApiAllBookTickersStreamRequest) Id(id string) ApiAllBookTickersStreamRequest {
	r.id = &id
	return r
}

func (r ApiAllBookTickersStreamRequest) Execute() (*common.StreamHandler[models.AllBookTickersStreamResponse], error) {
	return r.ApiService.AllBookTickersStreamExecute(r)
}

/*
AllBookTickersStream All Book Tickers Stream
/!bookTicker

https://developers.binance.com/docs/derivatives/usds-margined-futures/websocket-market-streams/All-Book-Tickers-Stream

@param id Unique WebSocket request ID.
@return ApiAllBookTickersStreamRequest
*/
func (a *PublicAPIService) AllBookTickersStream() ApiAllBookTickersStreamRequest {
	return ApiAllBookTickersStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AllBookTickersStreamResponse
func (a *PublicAPIService) AllBookTickersStreamExecute(r ApiAllBookTickersStreamRequest) (*common.StreamHandler[models.AllBookTickersStreamResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/!bookTicker"[1:],
		map[string]string{
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return *r.id
			}(),
		},
	)
	ws := a.client.WsPublic

	id := []any{common.GenerateUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.AllBookTickersStreamResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiDiffBookDepthStreamsRequest struct {
	ApiService  *PublicAPIService
	symbol      *string
	id          *string
	updateSpeed *string
}

// The symbol parameter
func (r ApiDiffBookDepthStreamsRequest) Symbol(symbol string) ApiDiffBookDepthStreamsRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiDiffBookDepthStreamsRequest) Id(id string) ApiDiffBookDepthStreamsRequest {
	r.id = &id
	return r
}

// WebSocket stream update speed
func (r ApiDiffBookDepthStreamsRequest) UpdateSpeed(updateSpeed string) ApiDiffBookDepthStreamsRequest {
	r.updateSpeed = &updateSpeed
	return r
}

func (r ApiDiffBookDepthStreamsRequest) Execute() (*common.StreamHandler[models.DiffBookDepthStreamsResponse], error) {
	return r.ApiService.DiffBookDepthStreamsExecute(r)
}

/*
DiffBookDepthStreams Diff. Book Depth Streams
/<symbol>@depth@<updateSpeed>

https://developers.binance.com/docs/derivatives/usds-margined-futures/websocket-market-streams/Diff-Book-Depth-Streams

@param symbol The symbol parameter	@param id Unique WebSocket request ID.	@param updateSpeed WebSocket stream update speed
@return ApiDiffBookDepthStreamsRequest
*/
func (a *PublicAPIService) DiffBookDepthStreams() ApiDiffBookDepthStreamsRequest {
	return ApiDiffBookDepthStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return DiffBookDepthStreamsResponse
func (a *PublicAPIService) DiffBookDepthStreamsExecute(r ApiDiffBookDepthStreamsRequest) (*common.StreamHandler[models.DiffBookDepthStreamsResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@depth@<updateSpeed>"[1:],
		map[string]string{
			"symbol": func() string {
				if r.symbol == nil {
					return ""
				}
				return *r.symbol
			}(),
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return *r.id
			}(),
			"updateSpeed": func() string {
				if r.updateSpeed == nil {
					return ""
				}
				return *r.updateSpeed
			}(),
		},
	)
	ws := a.client.WsPublic

	id := []any{common.GenerateUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.DiffBookDepthStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiIndividualSymbolBookTickerStreamsRequest struct {
	ApiService *PublicAPIService
	symbol     *string
	id         *string
}

// The symbol parameter
func (r ApiIndividualSymbolBookTickerStreamsRequest) Symbol(symbol string) ApiIndividualSymbolBookTickerStreamsRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiIndividualSymbolBookTickerStreamsRequest) Id(id string) ApiIndividualSymbolBookTickerStreamsRequest {
	r.id = &id
	return r
}

func (r ApiIndividualSymbolBookTickerStreamsRequest) Execute() (*common.StreamHandler[models.IndividualSymbolBookTickerStreamsResponse], error) {
	return r.ApiService.IndividualSymbolBookTickerStreamsExecute(r)
}

/*
IndividualSymbolBookTickerStreams Individual Symbol Book Ticker Streams
/<symbol>@bookTicker

https://developers.binance.com/docs/derivatives/usds-margined-futures/websocket-market-streams/Individual-Symbol-Book-Ticker-Streams

@param symbol The symbol parameter	@param id Unique WebSocket request ID.
@return ApiIndividualSymbolBookTickerStreamsRequest
*/
func (a *PublicAPIService) IndividualSymbolBookTickerStreams() ApiIndividualSymbolBookTickerStreamsRequest {
	return ApiIndividualSymbolBookTickerStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return IndividualSymbolBookTickerStreamsResponse
func (a *PublicAPIService) IndividualSymbolBookTickerStreamsExecute(r ApiIndividualSymbolBookTickerStreamsRequest) (*common.StreamHandler[models.IndividualSymbolBookTickerStreamsResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@bookTicker"[1:],
		map[string]string{
			"symbol": func() string {
				if r.symbol == nil {
					return ""
				}
				return *r.symbol
			}(),
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return *r.id
			}(),
		},
	)
	ws := a.client.WsPublic

	id := []any{common.GenerateUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.IndividualSymbolBookTickerStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiPartialBookDepthStreamsRequest struct {
	ApiService  *PublicAPIService
	symbol      *string
	levels      *int64
	id          *string
	updateSpeed *string
}

// The symbol parameter
func (r ApiPartialBookDepthStreamsRequest) Symbol(symbol string) ApiPartialBookDepthStreamsRequest {
	r.symbol = &symbol
	return r
}

// The levels parameter
func (r ApiPartialBookDepthStreamsRequest) Levels(levels int64) ApiPartialBookDepthStreamsRequest {
	r.levels = &levels
	return r
}

// Unique WebSocket request ID.
func (r ApiPartialBookDepthStreamsRequest) Id(id string) ApiPartialBookDepthStreamsRequest {
	r.id = &id
	return r
}

// WebSocket stream update speed
func (r ApiPartialBookDepthStreamsRequest) UpdateSpeed(updateSpeed string) ApiPartialBookDepthStreamsRequest {
	r.updateSpeed = &updateSpeed
	return r
}

func (r ApiPartialBookDepthStreamsRequest) Execute() (*common.StreamHandler[models.PartialBookDepthStreamsResponse], error) {
	return r.ApiService.PartialBookDepthStreamsExecute(r)
}

/*
PartialBookDepthStreams Partial Book Depth Streams
/<symbol>@depth<levels>@<updateSpeed>

https://developers.binance.com/docs/derivatives/usds-margined-futures/websocket-market-streams/Partial-Book-Depth-Streams

@param symbol The symbol parameter	@param levels The levels parameter	@param id Unique WebSocket request ID.	@param updateSpeed WebSocket stream update speed
@return ApiPartialBookDepthStreamsRequest
*/
func (a *PublicAPIService) PartialBookDepthStreams() ApiPartialBookDepthStreamsRequest {
	return ApiPartialBookDepthStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return PartialBookDepthStreamsResponse
func (a *PublicAPIService) PartialBookDepthStreamsExecute(r ApiPartialBookDepthStreamsRequest) (*common.StreamHandler[models.PartialBookDepthStreamsResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.levels == nil {
		return nil, common.ReportError("levels is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@depth<levels>@<updateSpeed>"[1:],
		map[string]string{
			"symbol": func() string {
				if r.symbol == nil {
					return ""
				}
				return *r.symbol
			}(),
			"levels": func() string {
				if r.levels == nil {
					return ""
				}
				return strconv.FormatInt(*r.levels, 10)
			}(),
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return *r.id
			}(),
			"updateSpeed": func() string {
				if r.updateSpeed == nil {
					return ""
				}
				return *r.updateSpeed
			}(),
		},
	)
	ws := a.client.WsPublic

	id := []any{common.GenerateUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.PartialBookDepthStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiRpiDiffBookDepthStreamsRequest struct {
	ApiService *PublicAPIService
	symbol     *string
	id         *string
}

// The symbol parameter
func (r ApiRpiDiffBookDepthStreamsRequest) Symbol(symbol string) ApiRpiDiffBookDepthStreamsRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiRpiDiffBookDepthStreamsRequest) Id(id string) ApiRpiDiffBookDepthStreamsRequest {
	r.id = &id
	return r
}

func (r ApiRpiDiffBookDepthStreamsRequest) Execute() (*common.StreamHandler[models.RpiDiffBookDepthStreamsResponse], error) {
	return r.ApiService.RpiDiffBookDepthStreamsExecute(r)
}

/*
RpiDiffBookDepthStreams RPI Diff. Book Depth Streams
/<symbol>@rpiDepth@500ms

https://developers.binance.com/docs/derivatives/usds-margined-futures/websocket-market-streams/Diff-Book-Depth-Streams-RPI

@param symbol The symbol parameter	@param id Unique WebSocket request ID.
@return ApiRpiDiffBookDepthStreamsRequest
*/
func (a *PublicAPIService) RpiDiffBookDepthStreams() ApiRpiDiffBookDepthStreamsRequest {
	return ApiRpiDiffBookDepthStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return RpiDiffBookDepthStreamsResponse
func (a *PublicAPIService) RpiDiffBookDepthStreamsExecute(r ApiRpiDiffBookDepthStreamsRequest) (*common.StreamHandler[models.RpiDiffBookDepthStreamsResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@rpiDepth@500ms"[1:],
		map[string]string{
			"symbol": func() string {
				if r.symbol == nil {
					return ""
				}
				return *r.symbol
			}(),
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return *r.id
			}(),
		},
	)
	ws := a.client.WsPublic

	id := []any{common.GenerateUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.RpiDiffBookDepthStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
