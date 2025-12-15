/*
Binance Derivatives Trading Options WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Options WebSocket Market Streams

API version: 1.0.0
*/

package binancederivativestradingoptionswebsocketstreams

import (
	"strconv"

	"github.com/binance/binance-connector-go/clients/derivativestradingoptions/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/common"
)

// WebsocketMarketStreamsAPIService WebsocketMarketStreamsAPI Service
type WebsocketMarketStreamsAPIService Service

type ApiIndexPriceStreamsRequest struct {
	ApiService *WebsocketMarketStreamsAPIService
	symbol     *string
	id         *string
}

// The symbol parameter
func (r ApiIndexPriceStreamsRequest) Symbol(symbol string) ApiIndexPriceStreamsRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiIndexPriceStreamsRequest) Id(id string) ApiIndexPriceStreamsRequest {
	r.id = &id
	return r
}

func (r ApiIndexPriceStreamsRequest) Execute() (*common.StreamHandler[models.IndexPriceStreamsResponse], error) {
	return r.ApiService.IndexPriceStreamsExecute(r)
}

/*
IndexPriceStreams Index Price Streams
/<symbol>@index

https://developers.binance.com/docs/derivatives/option/websocket-market-streams/Index-Price-Streams

@param symbol The symbol parameter	@param id Unique WebSocket request ID.
@return ApiIndexPriceStreamsRequest
*/
func (a *WebsocketMarketStreamsAPIService) IndexPriceStreams() ApiIndexPriceStreamsRequest {
	return ApiIndexPriceStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return IndexPriceStreamsResponse
func (a *WebsocketMarketStreamsAPIService) IndexPriceStreamsExecute(r ApiIndexPriceStreamsRequest) (*common.StreamHandler[models.IndexPriceStreamsResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@index"[1:],
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
	ws := a.client.Ws

	id := []string{common.GenerateUUID()}
	if r.id != nil {
		id = []string{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.IndexPriceStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiKlineCandlestickStreamsRequest struct {
	ApiService *WebsocketMarketStreamsAPIService
	symbol     *string
	interval   *string
	id         *string
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
func (r ApiKlineCandlestickStreamsRequest) Id(id string) ApiKlineCandlestickStreamsRequest {
	r.id = &id
	return r
}

func (r ApiKlineCandlestickStreamsRequest) Execute() (*common.StreamHandler[models.KlineCandlestickStreamsResponse], error) {
	return r.ApiService.KlineCandlestickStreamsExecute(r)
}

/*
KlineCandlestickStreams Kline/Candlestick Streams
/<symbol>@kline_<interval>

https://developers.binance.com/docs/derivatives/option/websocket-market-streams/Kline-Candlestick-Streams

@param symbol The symbol parameter	@param interval The interval parameter	@param id Unique WebSocket request ID.
@return ApiKlineCandlestickStreamsRequest
*/
func (a *WebsocketMarketStreamsAPIService) KlineCandlestickStreams() ApiKlineCandlestickStreamsRequest {
	return ApiKlineCandlestickStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return KlineCandlestickStreamsResponse
func (a *WebsocketMarketStreamsAPIService) KlineCandlestickStreamsExecute(r ApiKlineCandlestickStreamsRequest) (*common.StreamHandler[models.KlineCandlestickStreamsResponse], error) {
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
				return *r.id
			}(),
		},
	)
	ws := a.client.Ws

	id := []string{common.GenerateUUID()}
	if r.id != nil {
		id = []string{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.KlineCandlestickStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiMarkPriceRequest struct {
	ApiService      *WebsocketMarketStreamsAPIService
	underlyingAsset *string
	id              *string
}

// The underlyingAsset parameter
func (r ApiMarkPriceRequest) UnderlyingAsset(underlyingAsset string) ApiMarkPriceRequest {
	r.underlyingAsset = &underlyingAsset
	return r
}

// Unique WebSocket request ID.
func (r ApiMarkPriceRequest) Id(id string) ApiMarkPriceRequest {
	r.id = &id
	return r
}

func (r ApiMarkPriceRequest) Execute() (*common.StreamHandler[models.MarkPriceResponse], error) {
	return r.ApiService.MarkPriceExecute(r)
}

/*
MarkPrice Mark Price
/<underlyingAsset>@markPrice

https://developers.binance.com/docs/derivatives/option/websocket-market-streams/Mark-Price

@param underlyingAsset The underlyingAsset parameter	@param id Unique WebSocket request ID.
@return ApiMarkPriceRequest
*/
func (a *WebsocketMarketStreamsAPIService) MarkPrice() ApiMarkPriceRequest {
	return ApiMarkPriceRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return MarkPriceResponse
func (a *WebsocketMarketStreamsAPIService) MarkPriceExecute(r ApiMarkPriceRequest) (*common.StreamHandler[models.MarkPriceResponse], error) {
	if r.underlyingAsset == nil {
		return nil, common.ReportError("underlyingAsset is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<underlyingAsset>@markPrice"[1:],
		map[string]string{
			"underlyingAsset": func() string {
				if r.underlyingAsset == nil {
					return ""
				}
				return *r.underlyingAsset
			}(),
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return *r.id
			}(),
		},
	)
	ws := a.client.Ws

	id := []string{common.GenerateUUID()}
	if r.id != nil {
		id = []string{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.MarkPriceResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiNewSymbolInfoRequest struct {
	ApiService *WebsocketMarketStreamsAPIService
	id         *string
}

// Unique WebSocket request ID.
func (r ApiNewSymbolInfoRequest) Id(id string) ApiNewSymbolInfoRequest {
	r.id = &id
	return r
}

func (r ApiNewSymbolInfoRequest) Execute() (*common.StreamHandler[models.NewSymbolInfoResponse], error) {
	return r.ApiService.NewSymbolInfoExecute(r)
}

/*
NewSymbolInfo New Symbol Info
/option_pair

https://developers.binance.com/docs/derivatives/option/websocket-market-streams/New-Symbol-Info

@param id Unique WebSocket request ID.
@return ApiNewSymbolInfoRequest
*/
func (a *WebsocketMarketStreamsAPIService) NewSymbolInfo() ApiNewSymbolInfoRequest {
	return ApiNewSymbolInfoRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return NewSymbolInfoResponse
func (a *WebsocketMarketStreamsAPIService) NewSymbolInfoExecute(r ApiNewSymbolInfoRequest) (*common.StreamHandler[models.NewSymbolInfoResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/option_pair"[1:],
		map[string]string{
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return *r.id
			}(),
		},
	)
	ws := a.client.Ws

	id := []string{common.GenerateUUID()}
	if r.id != nil {
		id = []string{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.NewSymbolInfoResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiOpenInterestRequest struct {
	ApiService      *WebsocketMarketStreamsAPIService
	underlyingAsset *string
	expirationDate  *string
	id              *string
}

// The underlyingAsset parameter
func (r ApiOpenInterestRequest) UnderlyingAsset(underlyingAsset string) ApiOpenInterestRequest {
	r.underlyingAsset = &underlyingAsset
	return r
}

// The expirationDate parameter
func (r ApiOpenInterestRequest) ExpirationDate(expirationDate string) ApiOpenInterestRequest {
	r.expirationDate = &expirationDate
	return r
}

// Unique WebSocket request ID.
func (r ApiOpenInterestRequest) Id(id string) ApiOpenInterestRequest {
	r.id = &id
	return r
}

func (r ApiOpenInterestRequest) Execute() (*common.StreamHandler[models.OpenInterestResponse], error) {
	return r.ApiService.OpenInterestExecute(r)
}

/*
OpenInterest Open Interest
/<underlyingAsset>@openInterest@<expirationDate>

https://developers.binance.com/docs/derivatives/option/websocket-market-streams/Open-Interest

@param underlyingAsset The underlyingAsset parameter	@param expirationDate The expirationDate parameter	@param id Unique WebSocket request ID.
@return ApiOpenInterestRequest
*/
func (a *WebsocketMarketStreamsAPIService) OpenInterest() ApiOpenInterestRequest {
	return ApiOpenInterestRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return OpenInterestResponse
func (a *WebsocketMarketStreamsAPIService) OpenInterestExecute(r ApiOpenInterestRequest) (*common.StreamHandler[models.OpenInterestResponse], error) {
	if r.underlyingAsset == nil {
		return nil, common.ReportError("underlyingAsset is required and must be specified")
	}
	if r.expirationDate == nil {
		return nil, common.ReportError("expirationDate is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<underlyingAsset>@openInterest@<expirationDate>"[1:],
		map[string]string{
			"underlyingAsset": func() string {
				if r.underlyingAsset == nil {
					return ""
				}
				return *r.underlyingAsset
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
				return *r.id
			}(),
		},
	)
	ws := a.client.Ws

	id := []string{common.GenerateUUID()}
	if r.id != nil {
		id = []string{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.OpenInterestResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiPartialBookDepthStreamsRequest struct {
	ApiService  *WebsocketMarketStreamsAPIService
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

https://developers.binance.com/docs/derivatives/option/websocket-market-streams/Partial-Book-Depth-Streams

@param symbol The symbol parameter	@param levels The levels parameter	@param id Unique WebSocket request ID.	@param updateSpeed WebSocket stream update speed
@return ApiPartialBookDepthStreamsRequest
*/
func (a *WebsocketMarketStreamsAPIService) PartialBookDepthStreams() ApiPartialBookDepthStreamsRequest {
	return ApiPartialBookDepthStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return PartialBookDepthStreamsResponse
func (a *WebsocketMarketStreamsAPIService) PartialBookDepthStreamsExecute(r ApiPartialBookDepthStreamsRequest) (*common.StreamHandler[models.PartialBookDepthStreamsResponse], error) {
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
	ws := a.client.Ws

	id := []string{common.GenerateUUID()}
	if r.id != nil {
		id = []string{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.PartialBookDepthStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiTicker24HourRequest struct {
	ApiService *WebsocketMarketStreamsAPIService
	symbol     *string
	id         *string
}

// The symbol parameter
func (r ApiTicker24HourRequest) Symbol(symbol string) ApiTicker24HourRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiTicker24HourRequest) Id(id string) ApiTicker24HourRequest {
	r.id = &id
	return r
}

func (r ApiTicker24HourRequest) Execute() (*common.StreamHandler[models.Ticker24HourResponse], error) {
	return r.ApiService.Ticker24HourExecute(r)
}

/*
Ticker24Hour 24-hour TICKER
/<symbol>@ticker

https://developers.binance.com/docs/derivatives/option/websocket-market-streams/24-hour-TICKER

@param symbol The symbol parameter	@param id Unique WebSocket request ID.
@return ApiTicker24HourRequest
*/
func (a *WebsocketMarketStreamsAPIService) Ticker24Hour() ApiTicker24HourRequest {
	return ApiTicker24HourRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return Ticker24HourResponse
func (a *WebsocketMarketStreamsAPIService) Ticker24HourExecute(r ApiTicker24HourRequest) (*common.StreamHandler[models.Ticker24HourResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@ticker"[1:],
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
	ws := a.client.Ws

	id := []string{common.GenerateUUID()}
	if r.id != nil {
		id = []string{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.Ticker24HourResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiTicker24HourByUnderlyingAssetAndExpirationDataRequest struct {
	ApiService      *WebsocketMarketStreamsAPIService
	underlyingAsset *string
	expirationDate  *string
	id              *string
}

// The underlyingAsset parameter
func (r ApiTicker24HourByUnderlyingAssetAndExpirationDataRequest) UnderlyingAsset(underlyingAsset string) ApiTicker24HourByUnderlyingAssetAndExpirationDataRequest {
	r.underlyingAsset = &underlyingAsset
	return r
}

// The expirationDate parameter
func (r ApiTicker24HourByUnderlyingAssetAndExpirationDataRequest) ExpirationDate(expirationDate string) ApiTicker24HourByUnderlyingAssetAndExpirationDataRequest {
	r.expirationDate = &expirationDate
	return r
}

// Unique WebSocket request ID.
func (r ApiTicker24HourByUnderlyingAssetAndExpirationDataRequest) Id(id string) ApiTicker24HourByUnderlyingAssetAndExpirationDataRequest {
	r.id = &id
	return r
}

func (r ApiTicker24HourByUnderlyingAssetAndExpirationDataRequest) Execute() (*common.StreamHandler[models.Ticker24HourByUnderlyingAssetAndExpirationDataResponse], error) {
	return r.ApiService.Ticker24HourByUnderlyingAssetAndExpirationDataExecute(r)
}

/*
Ticker24HourByUnderlyingAssetAndExpirationData 24-hour TICKER by underlying asset and expiration data
/<underlyingAsset>@ticker@<expirationDate>

https://developers.binance.com/docs/derivatives/option/websocket-market-streams/24-hour-TICKER-by-underlying-asset-and-expiration-data

@param underlyingAsset The underlyingAsset parameter	@param expirationDate The expirationDate parameter	@param id Unique WebSocket request ID.
@return ApiTicker24HourByUnderlyingAssetAndExpirationDataRequest
*/
func (a *WebsocketMarketStreamsAPIService) Ticker24HourByUnderlyingAssetAndExpirationData() ApiTicker24HourByUnderlyingAssetAndExpirationDataRequest {
	return ApiTicker24HourByUnderlyingAssetAndExpirationDataRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return Ticker24HourByUnderlyingAssetAndExpirationDataResponse
func (a *WebsocketMarketStreamsAPIService) Ticker24HourByUnderlyingAssetAndExpirationDataExecute(r ApiTicker24HourByUnderlyingAssetAndExpirationDataRequest) (*common.StreamHandler[models.Ticker24HourByUnderlyingAssetAndExpirationDataResponse], error) {
	if r.underlyingAsset == nil {
		return nil, common.ReportError("underlyingAsset is required and must be specified")
	}
	if r.expirationDate == nil {
		return nil, common.ReportError("expirationDate is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<underlyingAsset>@ticker@<expirationDate>"[1:],
		map[string]string{
			"underlyingAsset": func() string {
				if r.underlyingAsset == nil {
					return ""
				}
				return *r.underlyingAsset
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
				return *r.id
			}(),
		},
	)
	ws := a.client.Ws

	id := []string{common.GenerateUUID()}
	if r.id != nil {
		id = []string{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.Ticker24HourByUnderlyingAssetAndExpirationDataResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiTradeStreamsRequest struct {
	ApiService *WebsocketMarketStreamsAPIService
	symbol     *string
	id         *string
}

// The symbol parameter
func (r ApiTradeStreamsRequest) Symbol(symbol string) ApiTradeStreamsRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiTradeStreamsRequest) Id(id string) ApiTradeStreamsRequest {
	r.id = &id
	return r
}

func (r ApiTradeStreamsRequest) Execute() (*common.StreamHandler[models.TradeStreamsResponse], error) {
	return r.ApiService.TradeStreamsExecute(r)
}

/*
TradeStreams Trade Streams
/<symbol>@trade

https://developers.binance.com/docs/derivatives/option/websocket-market-streams/Trade-Streams

@param symbol The symbol parameter	@param id Unique WebSocket request ID.
@return ApiTradeStreamsRequest
*/
func (a *WebsocketMarketStreamsAPIService) TradeStreams() ApiTradeStreamsRequest {
	return ApiTradeStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return TradeStreamsResponse
func (a *WebsocketMarketStreamsAPIService) TradeStreamsExecute(r ApiTradeStreamsRequest) (*common.StreamHandler[models.TradeStreamsResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@trade"[1:],
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
	ws := a.client.Ws

	id := []string{common.GenerateUUID()}
	if r.id != nil {
		id = []string{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.TradeStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
