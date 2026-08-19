/*
Stocks Trading WebSocket Streams

WebSocket stream definitions for Binance Stocks Trading. Base URL: wss://nbstream.binance.com/equity
*/

package binancestockswebsocketstreams

import (
	"github.com/binance/binance-connector-go/clients/stocks/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// MarketStreamsAPIService MarketStreamsAPI Service
type MarketStreamsAPIService Service

type ApiCalendarStreamRequest struct {
	ApiService *MarketStreamsAPIService
}

func (r ApiCalendarStreamRequest) Execute() (*common.StreamHandler[models.CalendarStreamResponse], error) {
	return r.ApiService.CalendarStreamExecute(r)
}

/*
CalendarStream Calendar Stream
/calendar

https://developers.binance.com/en/docs/catalog/advanced-trading-stocks-trading/api/ws-streams/market-streams#calendar-stream

@return ApiCalendarStreamRequest
*/
func (a *MarketStreamsAPIService) CalendarStream() ApiCalendarStreamRequest {
	return ApiCalendarStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return CalendarStreamResponse
func (a *MarketStreamsAPIService) CalendarStreamExecute(r ApiCalendarStreamRequest) (*common.StreamHandler[models.CalendarStreamResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/calendar"[1:],
		map[string]string{},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	resp, err := common.CreateStreamHandler[models.CalendarStreamResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiKlineStreamRequest struct {
	ApiService *MarketStreamsAPIService
	symbol     *string
	interval   *string
}

// US-equity ticker (UPPERCASE), e.g. &#x60;AAPL&#x60;.
func (r ApiKlineStreamRequest) Symbol(symbol string) ApiKlineStreamRequest {
	r.symbol = &symbol
	return r
}

// Kline interval — &#x60;5m&#x60; / &#x60;1h&#x60; / &#x60;1d&#x60; / &#x60;1w&#x60; / &#x60;1M&#x60;.
func (r ApiKlineStreamRequest) Interval(interval string) ApiKlineStreamRequest {
	r.interval = &interval
	return r
}

func (r ApiKlineStreamRequest) Execute() (*common.StreamHandler[models.KlineStreamResponse], error) {
	return r.ApiService.KlineStreamExecute(r)
}

/*
KlineStream Kline Stream
/<symbol>@kline_<interval>

https://developers.binance.com/en/docs/catalog/advanced-trading-stocks-trading/api/ws-streams/market-streams#kline-stream

@param symbol US-equity ticker (UPPERCASE), e.g. `AAPL`.	@param interval Kline interval — `5m` / `1h` / `1d` / `1w` / `1M`.
@return ApiKlineStreamRequest
*/
func (a *MarketStreamsAPIService) KlineStream() ApiKlineStreamRequest {
	return ApiKlineStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return KlineStreamResponse
func (a *MarketStreamsAPIService) KlineStreamExecute(r ApiKlineStreamRequest) (*common.StreamHandler[models.KlineStreamResponse], error) {
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
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	resp, err := common.CreateStreamHandler[models.KlineStreamResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiPriceStreamRequest struct {
	ApiService *MarketStreamsAPIService
}

func (r ApiPriceStreamRequest) Execute() (*common.StreamHandler[models.PriceStreamResponse], error) {
	return r.ApiService.PriceStreamExecute(r)
}

/*
PriceStream Price Stream
/price

https://developers.binance.com/en/docs/catalog/advanced-trading-stocks-trading/api/ws-streams/market-streams#price-stream

@return ApiPriceStreamRequest
*/
func (a *MarketStreamsAPIService) PriceStream() ApiPriceStreamRequest {
	return ApiPriceStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return PriceStreamResponse
func (a *MarketStreamsAPIService) PriceStreamExecute(r ApiPriceStreamRequest) (*common.StreamHandler[models.PriceStreamResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/price"[1:],
		map[string]string{},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	resp, err := common.CreateStreamHandler[models.PriceStreamResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiQuoteStreamRequest struct {
	ApiService *MarketStreamsAPIService
	symbol     *string
}

// US-equity ticker (UPPERCASE), e.g. &#x60;AAPL&#x60;.
func (r ApiQuoteStreamRequest) Symbol(symbol string) ApiQuoteStreamRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQuoteStreamRequest) Execute() (*common.StreamHandler[models.QuoteStreamResponse], error) {
	return r.ApiService.QuoteStreamExecute(r)
}

/*
QuoteStream Quote Stream
/<symbol>@quote

https://developers.binance.com/en/docs/catalog/advanced-trading-stocks-trading/api/ws-streams/market-streams#quote-stream

@param symbol US-equity ticker (UPPERCASE), e.g. `AAPL`.
@return ApiQuoteStreamRequest
*/
func (a *MarketStreamsAPIService) QuoteStream() ApiQuoteStreamRequest {
	return ApiQuoteStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return QuoteStreamResponse
func (a *MarketStreamsAPIService) QuoteStreamExecute(r ApiQuoteStreamRequest) (*common.StreamHandler[models.QuoteStreamResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@quote"[1:],
		map[string]string{
			"symbol": func() string {
				if r.symbol == nil {
					return ""
				}
				return *r.symbol
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	resp, err := common.CreateStreamHandler[models.QuoteStreamResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiTradabilityStreamRequest struct {
	ApiService *MarketStreamsAPIService
	symbol     *string
}

// US-equity ticker (UPPERCASE), e.g. &#x60;AAPL&#x60;.
func (r ApiTradabilityStreamRequest) Symbol(symbol string) ApiTradabilityStreamRequest {
	r.symbol = &symbol
	return r
}

func (r ApiTradabilityStreamRequest) Execute() (*common.StreamHandler[models.TradabilityStreamResponse], error) {
	return r.ApiService.TradabilityStreamExecute(r)
}

/*
TradabilityStream Tradability Stream
/<symbol>@tradability

https://developers.binance.com/en/docs/catalog/advanced-trading-stocks-trading/api/ws-streams/market-streams#tradability-stream

@param symbol US-equity ticker (UPPERCASE), e.g. `AAPL`.
@return ApiTradabilityStreamRequest
*/
func (a *MarketStreamsAPIService) TradabilityStream() ApiTradabilityStreamRequest {
	return ApiTradabilityStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return TradabilityStreamResponse
func (a *MarketStreamsAPIService) TradabilityStreamExecute(r ApiTradabilityStreamRequest) (*common.StreamHandler[models.TradabilityStreamResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@tradability"[1:],
		map[string]string{
			"symbol": func() string {
				if r.symbol == nil {
					return ""
				}
				return *r.symbol
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	resp, err := common.CreateStreamHandler[models.TradabilityStreamResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiTradingStatusStreamRequest struct {
	ApiService *MarketStreamsAPIService
	symbol     *string
}

// US-equity ticker (UPPERCASE), e.g. &#x60;AAPL&#x60;.
func (r ApiTradingStatusStreamRequest) Symbol(symbol string) ApiTradingStatusStreamRequest {
	r.symbol = &symbol
	return r
}

func (r ApiTradingStatusStreamRequest) Execute() (*common.StreamHandler[models.TradingStatusStreamResponse], error) {
	return r.ApiService.TradingStatusStreamExecute(r)
}

/*
TradingStatusStream Trading Status Stream
/<symbol>@tradingStatus

https://developers.binance.com/en/docs/catalog/advanced-trading-stocks-trading/api/ws-streams/market-streams#trading-status-stream

@param symbol US-equity ticker (UPPERCASE), e.g. `AAPL`.
@return ApiTradingStatusStreamRequest
*/
func (a *MarketStreamsAPIService) TradingStatusStream() ApiTradingStatusStreamRequest {
	return ApiTradingStatusStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return TradingStatusStreamResponse
func (a *MarketStreamsAPIService) TradingStatusStreamExecute(r ApiTradingStatusStreamRequest) (*common.StreamHandler[models.TradingStatusStreamResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@tradingStatus"[1:],
		map[string]string{
			"symbol": func() string {
				if r.symbol == nil {
					return ""
				}
				return *r.symbol
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	resp, err := common.CreateStreamHandler[models.TradingStatusStreamResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
