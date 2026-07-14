/*
Options WebSocket Market Streams

Access market data, manage accounts, and trade Binance Options.
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
	updateSpeed *models.DiffBookDepthStreamsUpdateSpeedParameter
	id          *int32
}

// The symbol parameter
func (r ApiDiffBookDepthStreamsRequest) Symbol(symbol string) ApiDiffBookDepthStreamsRequest {
	r.symbol = &symbol
	return r
}

// WebSocket stream update speed
func (r ApiDiffBookDepthStreamsRequest) UpdateSpeed(updateSpeed models.DiffBookDepthStreamsUpdateSpeedParameter) ApiDiffBookDepthStreamsRequest {
	r.updateSpeed = &updateSpeed
	return r
}

// Unique WebSocket request ID.
func (r ApiDiffBookDepthStreamsRequest) Id(id int32) ApiDiffBookDepthStreamsRequest {
	r.id = &id
	return r
}

func (r ApiDiffBookDepthStreamsRequest) Execute() (*common.StreamHandler[models.DiffBookDepthStreamsResponse], error) {
	return r.ApiService.DiffBookDepthStreamsExecute(r)
}

/*
DiffBookDepthStreams Diff Book Depth Streams
/<symbol>@depth@<updateSpeed>

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-options/api/ws-streams/public#diff-book-depth-streams

@param symbol The symbol parameter	@param updateSpeed WebSocket stream update speed	@param id Unique WebSocket request ID.
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
	if r.updateSpeed == nil {
		return nil, common.ReportError("updateSpeed is required and must be specified")
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
			"updateSpeed": func() string {
				if r.updateSpeed == nil {
					return ""
				}
				return string(*r.updateSpeed)
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
	resp, err := common.CreateStreamHandler[models.DiffBookDepthStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, true)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiHour24TickerRequest struct {
	ApiService     *PublicAPIService
	symbol         *string
	id             *int32
	expirationDate *string
}

// The symbol parameter
func (r ApiHour24TickerRequest) Symbol(symbol string) ApiHour24TickerRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiHour24TickerRequest) Id(id int32) ApiHour24TickerRequest {
	r.id = &id
	return r
}

// The expiration date parameter
func (r ApiHour24TickerRequest) ExpirationDate(expirationDate string) ApiHour24TickerRequest {
	r.expirationDate = &expirationDate
	return r
}

func (r ApiHour24TickerRequest) Execute() (*common.StreamHandler[models.Hour24TickerResponse], error) {
	return r.ApiService.Hour24TickerExecute(r)
}

/*
Hour24Ticker 24-hour TICKER
/<symbol>@optionTicker<expirationDate>

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-options/api/ws-streams/public#hour24-ticker

@param symbol The symbol parameter	@param id Unique WebSocket request ID.	@param expirationDate The expiration date parameter
@return ApiHour24TickerRequest
*/
func (a *PublicAPIService) Hour24Ticker() ApiHour24TickerRequest {
	return ApiHour24TickerRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return Hour24TickerResponse
func (a *PublicAPIService) Hour24TickerExecute(r ApiHour24TickerRequest) (*common.StreamHandler[models.Hour24TickerResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@optionTicker<expirationDate>"[1:],
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
			"expirationDate": func() string {
				if r.expirationDate == nil {
					return ""
				}
				return *r.expirationDate
			}(),
		},
	)
	ws := a.client.WsPublic

	id := []any{common.GenerateIntUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.Hour24TickerResponse](&common.StreamHandlerWrapper{
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-options/api/ws-streams/public#individual-symbol-book-ticker-streams

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
	level       *models.PartialBookDepthStreamsLevelParameter
	updateSpeed *models.DiffBookDepthStreamsUpdateSpeedParameter
	id          *int32
}

// The symbol parameter
func (r ApiPartialBookDepthStreamsRequest) Symbol(symbol string) ApiPartialBookDepthStreamsRequest {
	r.symbol = &symbol
	return r
}

// The level parameter
func (r ApiPartialBookDepthStreamsRequest) Level(level models.PartialBookDepthStreamsLevelParameter) ApiPartialBookDepthStreamsRequest {
	r.level = &level
	return r
}

// WebSocket stream update speed
func (r ApiPartialBookDepthStreamsRequest) UpdateSpeed(updateSpeed models.DiffBookDepthStreamsUpdateSpeedParameter) ApiPartialBookDepthStreamsRequest {
	r.updateSpeed = &updateSpeed
	return r
}

// Unique WebSocket request ID.
func (r ApiPartialBookDepthStreamsRequest) Id(id int32) ApiPartialBookDepthStreamsRequest {
	r.id = &id
	return r
}

func (r ApiPartialBookDepthStreamsRequest) Execute() (*common.StreamHandler[models.PartialBookDepthStreamsResponse], error) {
	return r.ApiService.PartialBookDepthStreamsExecute(r)
}

/*
PartialBookDepthStreams Partial Book Depth Streams
/<symbol>@depth<level>@<updateSpeed>

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-options/api/ws-streams/public#partial-book-depth-streams

@param symbol The symbol parameter	@param level The level parameter	@param updateSpeed WebSocket stream update speed	@param id Unique WebSocket request ID.
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
	if r.updateSpeed == nil {
		return nil, common.ReportError("updateSpeed is required and must be specified")
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
				return string(*r.level)
			}(),
			"updateSpeed": func() string {
				if r.updateSpeed == nil {
					return ""
				}
				return string(*r.updateSpeed)
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
	resp, err := common.CreateStreamHandler[models.PartialBookDepthStreamsResponse](&common.StreamHandlerWrapper{
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-options/api/ws-streams/public#trade-streams

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
