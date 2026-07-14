/*
Options WebSocket Market Streams

Access market data, manage accounts, and trade Binance Options.
*/

package binancederivativestradingoptionswebsocketstreams

import (
	"github.com/binance/binance-connector-go/clients/derivativestradingoptions/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-options/api/ws-streams/market#index-price-streams

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
	interval   *models.KlineCandlestickStreamsIntervalParameter
	id         *int32
}

// The symbol parameter
func (r ApiKlineCandlestickStreamsRequest) Symbol(symbol string) ApiKlineCandlestickStreamsRequest {
	r.symbol = &symbol
	return r
}

// The interval parameter
func (r ApiKlineCandlestickStreamsRequest) Interval(interval models.KlineCandlestickStreamsIntervalParameter) ApiKlineCandlestickStreamsRequest {
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-options/api/ws-streams/market#kline-candlestick-streams

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
				return string(*r.interval)
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-options/api/ws-streams/market#new-symbol-info

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
	underlying     *string
	expirationDate *string
	id             *int32
}

// The underlying parameter
func (r ApiOpenInterestRequest) Underlying(underlying string) ApiOpenInterestRequest {
	r.underlying = &underlying
	return r
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
/<underlying>@openInterest@<expirationDate>

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-options/api/ws-streams/market#open-interest

@param underlying The underlying parameter	@param expirationDate The expirationDate parameter	@param id Unique WebSocket request ID.
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
	if r.underlying == nil {
		return nil, common.ReportError("underlying is required and must be specified")
	}
	if r.expirationDate == nil {
		return nil, common.ReportError("expirationDate is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<underlying>@openInterest@<expirationDate>"[1:],
		map[string]string{
			"underlying": func() string {
				if r.underlying == nil {
					return ""
				}
				return *r.underlying
			}(),
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

type ApiOptionMarkPriceRequest struct {
	ApiService *MarketAPIService
	underlying *string
	id         *int32
}

// The underlying parameter
func (r ApiOptionMarkPriceRequest) Underlying(underlying string) ApiOptionMarkPriceRequest {
	r.underlying = &underlying
	return r
}

// Unique WebSocket request ID.
func (r ApiOptionMarkPriceRequest) Id(id int32) ApiOptionMarkPriceRequest {
	r.id = &id
	return r
}

func (r ApiOptionMarkPriceRequest) Execute() (*common.StreamHandler[models.OptionMarkPriceResponse], error) {
	return r.ApiService.OptionMarkPriceExecute(r)
}

/*
OptionMarkPrice Option Mark Price
/<underlying>@optionMarkPrice

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-options/api/ws-streams/market#option-mark-price

@param underlying The underlying parameter	@param id Unique WebSocket request ID.
@return ApiOptionMarkPriceRequest
*/
func (a *MarketAPIService) OptionMarkPrice() ApiOptionMarkPriceRequest {
	return ApiOptionMarkPriceRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return OptionMarkPriceResponse
func (a *MarketAPIService) OptionMarkPriceExecute(r ApiOptionMarkPriceRequest) (*common.StreamHandler[models.OptionMarkPriceResponse], error) {
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
	resp, err := common.CreateStreamHandler[models.OptionMarkPriceResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, true)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
