/*
Binance Derivatives Trading Options WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Options WebSocket Market Streams
*/

package binancederivativestradingoptionswebsocketstreams

import (
	"github.com/binance/binance-connector-go/clients/derivativestradingoptions/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// PublicAPIService PublicAPI Service
type PublicAPIService Service

type ApiDiffBookDepthStreamsRequest struct {
	ApiService  *PublicAPIService
	symbol      *string
	id          *int32
	updateSpeed *string
}

// The symbol parameter
func (r ApiDiffBookDepthStreamsRequest) Symbol(symbol string) ApiDiffBookDepthStreamsRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiDiffBookDepthStreamsRequest) Id(id int32) ApiDiffBookDepthStreamsRequest {
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
DiffBookDepthStreams Diff Book Depth Streams
/<symbol>@depth@<updateSpeed>

https://developers.binance.com/docs/derivatives/options-trading/websocket-market-streams/Diff-Book-Depth-Streams

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
				return string(*r.id)
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

	id := []any{common.GenerateIntUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.DiffBookDepthStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, true)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiIndividualSymbolBookTickerStreamsRequest struct {
	ApiService *PublicAPIService
	symbol     *string
	id         *int32
}

// The symbol parameter
func (r ApiIndividualSymbolBookTickerStreamsRequest) Symbol(symbol string) ApiIndividualSymbolBookTickerStreamsRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiIndividualSymbolBookTickerStreamsRequest) Id(id int32) ApiIndividualSymbolBookTickerStreamsRequest {
	r.id = &id
	return r
}

func (r ApiIndividualSymbolBookTickerStreamsRequest) Execute() (*common.StreamHandler[models.IndividualSymbolBookTickerStreamsResponse], error) {
	return r.ApiService.IndividualSymbolBookTickerStreamsExecute(r)
}

/*
IndividualSymbolBookTickerStreams Individual Symbol Book Ticker Streams
/<symbol>@bookTicker

https://developers.binance.com/docs/derivatives/options-trading/websocket-market-streams/Individual-Symbol-Book-Ticker-Streams

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
				return string(*r.id)
			}(),
		},
	)
	ws := a.client.WsPublic

	id := []any{common.GenerateIntUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.IndividualSymbolBookTickerStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, true)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiPartialBookDepthStreamsRequest struct {
	ApiService  *PublicAPIService
	symbol      *string
	level       *string
	id          *int32
	updateSpeed *string
}

// The symbol parameter
func (r ApiPartialBookDepthStreamsRequest) Symbol(symbol string) ApiPartialBookDepthStreamsRequest {
	r.symbol = &symbol
	return r
}

// The level parameter
func (r ApiPartialBookDepthStreamsRequest) Level(level string) ApiPartialBookDepthStreamsRequest {
	r.level = &level
	return r
}

// Unique WebSocket request ID.
func (r ApiPartialBookDepthStreamsRequest) Id(id int32) ApiPartialBookDepthStreamsRequest {
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
/<symbol>@depth<level>@<updateSpeed>

https://developers.binance.com/docs/derivatives/options-trading/websocket-market-streams/Partial-Book-Depth-Streams

@param symbol The symbol parameter	@param level The level parameter	@param id Unique WebSocket request ID.	@param updateSpeed WebSocket stream update speed
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
	if r.level == nil {
		return nil, common.ReportError("level is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@depth<level>@<updateSpeed>"[1:],
		map[string]string{
			"symbol": func() string {
				if r.symbol == nil {
					return ""
				}
				return *r.symbol
			}(),
			"level": func() string {
				if r.level == nil {
					return ""
				}
				return *r.level
			}(),
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return string(*r.id)
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

	id := []any{common.GenerateIntUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.PartialBookDepthStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, true)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiTicker24HourRequest struct {
	ApiService *PublicAPIService
	symbol     *string
	id         *int32
}

// The symbol parameter
func (r ApiTicker24HourRequest) Symbol(symbol string) ApiTicker24HourRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiTicker24HourRequest) Id(id int32) ApiTicker24HourRequest {
	r.id = &id
	return r
}

func (r ApiTicker24HourRequest) Execute() (*common.StreamHandler[models.Ticker24HourResponse], error) {
	return r.ApiService.Ticker24HourExecute(r)
}

/*
Ticker24Hour 24-hour TICKER
/<symbol>@optionTicker

https://developers.binance.com/docs/derivatives/options-trading/websocket-market-streams/24-hour-TICKER

@param symbol The symbol parameter	@param id Unique WebSocket request ID.
@return ApiTicker24HourRequest
*/
func (a *PublicAPIService) Ticker24Hour() ApiTicker24HourRequest {
	return ApiTicker24HourRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return Ticker24HourResponse
func (a *PublicAPIService) Ticker24HourExecute(r ApiTicker24HourRequest) (*common.StreamHandler[models.Ticker24HourResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@optionTicker"[1:],
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
				return string(*r.id)
			}(),
		},
	)
	ws := a.client.WsPublic

	id := []any{common.GenerateIntUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.Ticker24HourResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, true)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiTradeStreamsRequest struct {
	ApiService *PublicAPIService
	symbol     *string
	id         *int32
}

// The symbol parameter
func (r ApiTradeStreamsRequest) Symbol(symbol string) ApiTradeStreamsRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiTradeStreamsRequest) Id(id int32) ApiTradeStreamsRequest {
	r.id = &id
	return r
}

func (r ApiTradeStreamsRequest) Execute() (*common.StreamHandler[models.TradeStreamsResponse], error) {
	return r.ApiService.TradeStreamsExecute(r)
}

/*
TradeStreams Trade Streams
/<symbol>@optionTrade

https://developers.binance.com/docs/derivatives/options-trading/websocket-market-streams/Trade-Streams

@param symbol The symbol parameter	@param id Unique WebSocket request ID.
@return ApiTradeStreamsRequest
*/
func (a *PublicAPIService) TradeStreams() ApiTradeStreamsRequest {
	return ApiTradeStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return TradeStreamsResponse
func (a *PublicAPIService) TradeStreamsExecute(r ApiTradeStreamsRequest) (*common.StreamHandler[models.TradeStreamsResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@optionTrade"[1:],
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
				return string(*r.id)
			}(),
		},
	)
	ws := a.client.WsPublic

	id := []any{common.GenerateIntUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.TradeStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, true)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
