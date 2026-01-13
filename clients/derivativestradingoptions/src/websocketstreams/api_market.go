/*
Binance Derivatives Trading Options WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Options WebSocket Market Streams
*/

package binancederivativestradingoptionswebsocketstreams

import (
	"github.com/binance/binance-connector-go/clients/derivativestradingoptions/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/common"
)

// MarketAPIService MarketAPI Service
type MarketAPIService Service

type ApiIndexPriceStreamsRequest struct {
	ApiService *MarketAPIService
	id         *int32
}

// Unique WebSocket request ID.
func (r ApiIndexPriceStreamsRequest) Id(id int32) ApiIndexPriceStreamsRequest {
	r.id = &id
	return r
}

func (r ApiIndexPriceStreamsRequest) Execute() (*common.StreamHandler[models.IndexPriceStreamsResponse], error) {
	return r.ApiService.IndexPriceStreamsExecute(r)
}

/*
IndexPriceStreams Index Price Streams
/!index@arr

https://developers.binance.com/docs/derivatives/options-trading/websocket-market-streams/Index-Price-Streams

@param id Unique WebSocket request ID.
@return ApiIndexPriceStreamsRequest
*/
func (a *MarketAPIService) IndexPriceStreams() ApiIndexPriceStreamsRequest {
	return ApiIndexPriceStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return IndexPriceStreamsResponse
func (a *MarketAPIService) IndexPriceStreamsExecute(r ApiIndexPriceStreamsRequest) (*common.StreamHandler[models.IndexPriceStreamsResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/!index@arr"[1:],
		map[string]string{
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return string(*r.id)
			}(),
		},
	)
	ws := a.client.WsMarket

	id := []any{common.GenerateIntUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.IndexPriceStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, true)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiKlineCandlestickStreamsRequest struct {
	ApiService *MarketAPIService
	symbol     *string
	interval   *string
	id         *int32
}

// The symbol parameter
func (r ApiKlineCandlestickStreamsRequest) Symbol(symbol string) ApiKlineCandlestickStreamsRequest {
	r.symbol = &symbol
	return r
}

// The interval parameter
func (r ApiKlineCandlestickStreamsRequest) Interval(interval string) ApiKlineCandlestickStreamsRequest {
	r.interval = &interval
	return r
}

// Unique WebSocket request ID.
func (r ApiKlineCandlestickStreamsRequest) Id(id int32) ApiKlineCandlestickStreamsRequest {
	r.id = &id
	return r
}

func (r ApiKlineCandlestickStreamsRequest) Execute() (*common.StreamHandler[models.KlineCandlestickStreamsResponse], error) {
	return r.ApiService.KlineCandlestickStreamsExecute(r)
}

/*
KlineCandlestickStreams Kline/Candlestick Streams
/<symbol>@kline_<interval>

https://developers.binance.com/docs/derivatives/options-trading/websocket-market-streams/Kline-Candlestick-Streams

@param symbol The symbol parameter	@param interval The interval parameter	@param id Unique WebSocket request ID.
@return ApiKlineCandlestickStreamsRequest
*/
func (a *MarketAPIService) KlineCandlestickStreams() ApiKlineCandlestickStreamsRequest {
	return ApiKlineCandlestickStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return KlineCandlestickStreamsResponse
func (a *MarketAPIService) KlineCandlestickStreamsExecute(r ApiKlineCandlestickStreamsRequest) (*common.StreamHandler[models.KlineCandlestickStreamsResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.interval == nil {
		return nil, common.ReportError("interval is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@kline_<interval>"[1:],
		map[string]string{
			"symbol": func() string {
				if r.symbol == nil {
					return ""
				}
				return *r.symbol
			}(),
			"interval": func() string {
				if r.interval == nil {
					return ""
				}
				return *r.interval
			}(),
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return string(*r.id)
			}(),
		},
	)
	ws := a.client.WsMarket

	id := []any{common.GenerateIntUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.KlineCandlestickStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, true)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiMarkPriceRequest struct {
	ApiService *MarketAPIService
	underlying *string
	id         *int32
}

// The underlying parameter
func (r ApiMarkPriceRequest) Underlying(underlying string) ApiMarkPriceRequest {
	r.underlying = &underlying
	return r
}

// Unique WebSocket request ID.
func (r ApiMarkPriceRequest) Id(id int32) ApiMarkPriceRequest {
	r.id = &id
	return r
}

func (r ApiMarkPriceRequest) Execute() (*common.StreamHandler[models.MarkPriceResponse], error) {
	return r.ApiService.MarkPriceExecute(r)
}

/*
MarkPrice Mark Price
/<underlying>@optionMarkPrice

https://developers.binance.com/docs/derivatives/options-trading/websocket-market-streams/Mark-Price

@param underlying The underlying parameter	@param id Unique WebSocket request ID.
@return ApiMarkPriceRequest
*/
func (a *MarketAPIService) MarkPrice() ApiMarkPriceRequest {
	return ApiMarkPriceRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return MarkPriceResponse
func (a *MarketAPIService) MarkPriceExecute(r ApiMarkPriceRequest) (*common.StreamHandler[models.MarkPriceResponse], error) {
	if r.underlying == nil {
		return nil, common.ReportError("underlying is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<underlying>@optionMarkPrice"[1:],
		map[string]string{
			"underlying": func() string {
				if r.underlying == nil {
					return ""
				}
				return *r.underlying
			}(),
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return string(*r.id)
			}(),
		},
	)
	ws := a.client.WsMarket

	id := []any{common.GenerateIntUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.MarkPriceResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, true)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiNewSymbolInfoRequest struct {
	ApiService *MarketAPIService
	id         *int32
}

// Unique WebSocket request ID.
func (r ApiNewSymbolInfoRequest) Id(id int32) ApiNewSymbolInfoRequest {
	r.id = &id
	return r
}

func (r ApiNewSymbolInfoRequest) Execute() (*common.StreamHandler[models.NewSymbolInfoResponse], error) {
	return r.ApiService.NewSymbolInfoExecute(r)
}

/*
NewSymbolInfo New Symbol Info
/!optionSymbol

https://developers.binance.com/docs/derivatives/options-trading/websocket-market-streams/New-Symbol-Info

@param id Unique WebSocket request ID.
@return ApiNewSymbolInfoRequest
*/
func (a *MarketAPIService) NewSymbolInfo() ApiNewSymbolInfoRequest {
	return ApiNewSymbolInfoRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return NewSymbolInfoResponse
func (a *MarketAPIService) NewSymbolInfoExecute(r ApiNewSymbolInfoRequest) (*common.StreamHandler[models.NewSymbolInfoResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/!optionSymbol"[1:],
		map[string]string{
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return string(*r.id)
			}(),
		},
	)
	ws := a.client.WsMarket

	id := []any{common.GenerateIntUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.NewSymbolInfoResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, true)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiOpenInterestRequest struct {
	ApiService     *MarketAPIService
	expirationDate *string
	id             *int32
}

// The expirationDate parameter
func (r ApiOpenInterestRequest) ExpirationDate(expirationDate string) ApiOpenInterestRequest {
	r.expirationDate = &expirationDate
	return r
}

// Unique WebSocket request ID.
func (r ApiOpenInterestRequest) Id(id int32) ApiOpenInterestRequest {
	r.id = &id
	return r
}

func (r ApiOpenInterestRequest) Execute() (*common.StreamHandler[models.OpenInterestResponse], error) {
	return r.ApiService.OpenInterestExecute(r)
}

/*
OpenInterest Open Interest
/underlying@optionOpenInterest@<expirationDate>

https://developers.binance.com/docs/derivatives/options-trading/websocket-market-streams/Open-Interest

@param expirationDate The expirationDate parameter	@param id Unique WebSocket request ID.
@return ApiOpenInterestRequest
*/
func (a *MarketAPIService) OpenInterest() ApiOpenInterestRequest {
	return ApiOpenInterestRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return OpenInterestResponse
func (a *MarketAPIService) OpenInterestExecute(r ApiOpenInterestRequest) (*common.StreamHandler[models.OpenInterestResponse], error) {
	if r.expirationDate == nil {
		return nil, common.ReportError("expirationDate is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/underlying@optionOpenInterest@<expirationDate>"[1:],
		map[string]string{
			"expirationDate": func() string {
				if r.expirationDate == nil {
					return ""
				}
				return *r.expirationDate
			}(),
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return string(*r.id)
			}(),
		},
	)
	ws := a.client.WsMarket

	id := []any{common.GenerateIntUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.OpenInterestResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, true)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
