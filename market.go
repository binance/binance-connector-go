package binance_connector

import (
	"context"
	"encoding/json"
	"math/big"
	"net/http"
)

// Binance Test Connectivity endpoint (GET /api/v3/ping)
type Ping struct {
	c *Client
}

// Send the request
func (s *Ping) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/ping",
		secType:  secTypeNone,
	}
	_, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	return nil
}

// Binance Check Server Time endpoint (GET /api/v3/time)
type ServerTime struct {
	c *Client
}

// Send the request
func (s *ServerTime) Do(ctx context.Context, opts ...RequestOption) (res *ServerTimeResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/time",
		secType:  secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ServerTimeResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ServerTimeResponse define server time response
type ServerTimeResponse struct {
	ServerTime uint64 `json:"serverTime"`
}

// Binance Exchange Information endpoint (GET /api/v3/exchangeInfo)
type ExchangeInfo struct {
	c *Client
}

// Send the request
func (s *ExchangeInfo) Do(ctx context.Context, opts ...RequestOption) (res *ExchangeInfoResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/exchangeInfo",
		secType:  secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ExchangeInfoResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ExchangeInfoResponse define exchange info response
type ExchangeInfoResponse struct {
	Timezone        string            `json:"timezone"`
	ServerTime      uint64            `json:"serverTime"`
	RateLimits      []*RateLimit      `json:"rateLimits"`
	ExchangeFilters []*ExchangeFilter `json:"exchangeFilters"`
	Symbols         []*SymbolInfo     `json:"symbols"`
}

// RateLimit define rate limit
type RateLimit struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	Limit         int    `json:"limit"`
}

// ExchangeFilter define exchange filter
type ExchangeFilter struct {
	FilterType string `json:"filterType"`
	MaxNumAlgo int64  `json:"maxNumAlgoOrders"`
}

// Symbol define symbol
type SymbolInfo struct {
	Symbol                     string          `json:"symbol"`
	Status                     string          `json:"status"`
	BaseAsset                  string          `json:"baseAsset"`
	BaseAssetPrecision         int64           `json:"baseAssetPrecision"`
	QuoteAsset                 string          `json:"quoteAsset"`
	QuotePrecision             int64           `json:"quotePrecision"`
	OrderTypes                 []string        `json:"orderTypes"`
	IcebergAllowed             bool            `json:"icebergAllowed"`
	OcoAllowed                 bool            `json:"ocoAllowed"`
	QuoteOrderQtyMarketAllowed bool            `json:"quoteOrderQtyMarketAllowed"`
	IsSpotTradingAllowed       bool            `json:"isSpotTradingAllowed"`
	IsMarginTradingAllowed     bool            `json:"isMarginTradingAllowed"`
	Filters                    []*SymbolFilter `json:"filters"`
	Permissions                []string        `json:"permissions"`
}

// SymbolFilter define symbol filter
type SymbolFilter struct {
	FilterType       string `json:"filterType"`
	MinPrice         string `json:"minPrice"`
	MaxPrice         string `json:"maxPrice"`
	TickSize         string `json:"tickSize"`
	MinQty           string `json:"minQty"`
	MaxQty           string `json:"maxQty"`
	StepSize         string `json:"stepSize"`
	MinNotional      string `json:"minNotional"`
	Limit            uint   `json:"limit"`
	MaxNumAlgoOrders int64  `json:"maxNumAlgoOrders"`
}

// Binance Order Book endpoint (GET /api/v3/depth)
type OrderBook struct {
	c      *Client
	symbol string
	limit  *int
}

// Symbol set symbol
func (s *OrderBook) Symbol(symbol string) *OrderBook {
	s.symbol = symbol
	return s
}

// Limit set limit
func (s *OrderBook) Limit(limit int) *OrderBook {
	s.limit = &limit
	return s
}

// Send the request
func (s *OrderBook) Do(ctx context.Context, opts ...RequestOption) (res *OrderBookResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/depth",
		secType:  secTypeNone,
	}
	r.setParam("symbol", s.symbol)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(OrderBookResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// OrderBookResponse define order book response
type OrderBookResponse struct {
	LastUpdateId uint64         `json:"lastUpdateId"`
	Bids         [][]*big.Float `json:"bids"`
	Asks         [][]*big.Float `json:"asks"`
}

// Binance Recent Trades List endpoint (GET /api/v3/trades)
type RecentTradesList struct {
	c      *Client
	symbol string
	limit  *int
}

// Symbol set symbol
func (s *RecentTradesList) Symbol(symbol string) *RecentTradesList {
	s.symbol = symbol
	return s
}

// Limit set limit
func (s *RecentTradesList) Limit(limit int) *RecentTradesList {
	s.limit = &limit
	return s
}

// Send the request
func (s *RecentTradesList) Do(ctx context.Context, opts ...RequestOption) (res []*RecentTradesListResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/trades",
		secType:  secTypeNone,
	}
	r.setParam("symbol", s.symbol)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// RecentTradesListResponse define recent trades list response
type RecentTradesListResponse struct {
	Id           uint64 `json:"id"`
	Price        string `json:"price"`
	Qty          string `json:"qty"`
	Time         uint64 `json:"time"`
	QuoteQty     string `json:"quoteQty"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
	IsBest       bool   `json:"isBestMatch"`
}

// Binance Old Trade Lookup endpoint (GET /api/v3/historicalTrades)
type HistoricalTradeLookup struct {
	c      *Client
	symbol string
	limit  *uint
	fromId *int64
}

// Symbol set symbol
func (s *HistoricalTradeLookup) Symbol(symbol string) *HistoricalTradeLookup {
	s.symbol = symbol
	return s
}

// Limit set limit
func (s *HistoricalTradeLookup) Limit(limit uint) *HistoricalTradeLookup {
	s.limit = &limit
	return s
}

// FromId set fromId
func (s *HistoricalTradeLookup) FromId(fromId int64) *HistoricalTradeLookup {
	s.fromId = &fromId
	return s
}

// Send the request
func (s *HistoricalTradeLookup) Do(ctx context.Context, opts ...RequestOption) (res []*RecentTradesListResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/historicalTrades",
		secType:  secTypeAPIKey,
	}
	r.setParam("symbol", s.symbol)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.fromId != nil {
		r.setParam("fromId", *s.fromId)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Binance Compressed/Aggregate Trades List endpoint (GET /api/v3/aggTrades)
type AggTradesList struct {
	c         *Client
	symbol    string
	limit     *int
	fromId    *int
	startTime *uint64
	endTime   *uint64
}

// Symbol set symbol
func (s *AggTradesList) Symbol(symbol string) *AggTradesList {
	s.symbol = symbol
	return s
}

// Limit set limit
func (s *AggTradesList) Limit(limit int) *AggTradesList {
	s.limit = &limit
	return s
}

// FromId set fromId
func (s *AggTradesList) FromId(fromId int) *AggTradesList {
	s.fromId = &fromId
	return s
}

// StartTime set startTime
func (s *AggTradesList) StartTime(startTime uint64) *AggTradesList {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *AggTradesList) EndTime(endTime uint64) *AggTradesList {
	s.endTime = &endTime
	return s
}

// Send the request
func (s *AggTradesList) Do(ctx context.Context, opts ...RequestOption) (res []*AggTradesListResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/aggTrades",
		secType:  secTypeNone,
	}
	r.setParam("symbol", s.symbol)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.fromId != nil {
		r.setParam("fromId", *s.fromId)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AggTradesListResponse define compressed trades list response
type AggTradesListResponse struct {
	AggTradeId   uint64 `json:"a"`
	Price        string `json:"p"`
	Qty          string `json:"q"`
	FirstTradeId uint64 `json:"f"`
	LastTradeId  uint64 `json:"l"`
	Time         uint64 `json:"T"`
	IsBuyer      bool   `json:"m"`
	IsBest       bool   `json:"M"`
}

// Binance Kline/Candlestick Data endpoint (GET /api/v3/klines)
type Klines struct {
	c         *Client
	symbol    string
	interval  string
	limit     *int
	startTime *uint64
	endTime   *uint64
}

// Symbol set symbol
func (s *Klines) Symbol(symbol string) *Klines {
	s.symbol = symbol
	return s
}

// Interval set interval
func (s *Klines) Interval(interval string) *Klines {
	s.interval = interval
	return s
}

// Limit set limit
func (s *Klines) Limit(limit int) *Klines {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *Klines) StartTime(startTime uint64) *Klines {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *Klines) EndTime(endTime uint64) *Klines {
	s.endTime = &endTime
	return s
}

func (s *Klines) Do(ctx context.Context, opts ...RequestOption) (res []*KlinesResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/klines",
		secType:  secTypeNone,
	}
	r.setParam("symbol", s.symbol)
	r.setParam("interval", s.interval)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var klinesResponseArray KlinesResponseArray
	err = json.Unmarshal(data, &klinesResponseArray)
	if err != nil {
		return nil, err
	}
	var klines []*KlinesResponse
	for _, kline := range klinesResponseArray {
		openTime := kline[0].(float64)
		open := kline[1].(string)
		high := kline[2].(string)
		low := kline[3].(string)
		close := kline[4].(string)
		volume := kline[5].(string)
		closeTime := kline[6].(float64)
		quoteAssetVolume := kline[7].(string)
		numberOfTrades := kline[8].(float64)
		takerBuyBaseAssetVolume := kline[9].(string)
		takerBuyQuoteAssetVolume := kline[10].(string)

		// create a KlinesResponse struct using the parsed fields
		klinesResponse := &KlinesResponse{
			OpenTime:                 uint64(openTime),
			Open:                     open,
			High:                     high,
			Low:                      low,
			Close:                    close,
			Volume:                   volume,
			CloseTime:                uint64(closeTime),
			QuoteAssetVolume:         quoteAssetVolume,
			NumberOfTrades:           uint64(numberOfTrades),
			TakerBuyBaseAssetVolume:  takerBuyBaseAssetVolume,
			TakerBuyQuoteAssetVolume: takerBuyQuoteAssetVolume,
		}
		klines = append(klines, klinesResponse)
	}
	return klines, nil
}

type KlinesResponseArray [][]interface{}

// Define Klines response data
type KlinesResponse struct {
	OpenTime                 uint64 `json:"openTime"`
	Open                     string `json:"open"`
	High                     string `json:"high"`
	Low                      string `json:"low"`
	Close                    string `json:"close"`
	Volume                   string `json:"volume"`
	CloseTime                uint64 `json:"closeTime"`
	QuoteAssetVolume         string `json:"quoteAssetVolume"`
	NumberOfTrades           uint64 `json:"numberOfTrades"`
	TakerBuyBaseAssetVolume  string `json:"takerBuyBaseAssetVolume"`
	TakerBuyQuoteAssetVolume string `json:"takerBuyQuoteAssetVolume"`
}

// Binance UI Klines GET /api/v3/uiKlines
type UiKlines struct {
	c         *Client
	symbol    string
	interval  string
	limit     *int
	startTime *uint64
	endTime   *uint64
}

// Symbol set symbol
func (s *UiKlines) Symbol(symbol string) *UiKlines {
	s.symbol = symbol
	return s
}

// Interval set interval
func (s *UiKlines) Interval(interval string) *UiKlines {
	s.interval = interval
	return s
}

// Limit set limit
func (s *UiKlines) Limit(limit int) *UiKlines {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *UiKlines) StartTime(startTime uint64) *UiKlines {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *UiKlines) EndTime(endTime uint64) *UiKlines {
	s.endTime = &endTime
	return s
}

// Send the request
func (s *UiKlines) Do(ctx context.Context, opts ...RequestOption) (res []*UiKlinesResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/uiKlines",
		secType:  secTypeNone,
	}
	r.setParam("symbol", s.symbol)
	r.setParam("interval", s.interval)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var uiklinesResponseArray UiKlinesResponseArray
	err = json.Unmarshal(data, &uiklinesResponseArray)
	if err != nil {
		return nil, err
	}
	var uiklines []*UiKlinesResponse
	for _, uikline := range uiklinesResponseArray {
		openTime := uikline[0].(float64)
		open := uikline[1].(string)
		high := uikline[2].(string)
		low := uikline[3].(string)
		close := uikline[4].(string)
		volume := uikline[5].(string)
		closeTime := uikline[6].(float64)
		quoteAssetVolume := uikline[7].(string)
		numberOfTrades := uikline[8].(float64)
		takerBuyBaseAssetVolume := uikline[9].(string)
		takerBuyQuoteAssetVolume := uikline[10].(string)

		// create a KlinesResponse struct using the parsed fields
		uiklinesResponse := &UiKlinesResponse{
			OpenTime:                 uint64(openTime),
			Open:                     open,
			High:                     high,
			Low:                      low,
			Close:                    close,
			Volume:                   volume,
			CloseTime:                uint64(closeTime),
			QuoteAssetVolume:         quoteAssetVolume,
			NumberOfTrades:           uint64(numberOfTrades),
			TakerBuyBaseAssetVolume:  takerBuyBaseAssetVolume,
			TakerBuyQuoteAssetVolume: takerBuyQuoteAssetVolume,
		}
		uiklines = append(uiklines, uiklinesResponse)
	}
	return uiklines, nil
}

type UiKlinesResponseArray [][]interface{}

// Define UiKlines response data
type UiKlinesResponse struct {
	OpenTime                 uint64 `json:"openTime"`
	Open                     string `json:"open"`
	High                     string `json:"high"`
	Low                      string `json:"low"`
	Close                    string `json:"close"`
	Volume                   string `json:"volume"`
	CloseTime                uint64 `json:"closeTime"`
	QuoteAssetVolume         string `json:"quoteAssetVolume"`
	NumberOfTrades           uint64 `json:"numberOfTrades"`
	TakerBuyBaseAssetVolume  string `json:"takerBuyBaseAssetVolume"`
	TakerBuyQuoteAssetVolume string `json:"takerBuyQuoteAssetVolume"`
}

// Binance Current Average Price (GET /api/v3/avgPrice)
type AvgPrice struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *AvgPrice) Symbol(symbol string) *AvgPrice {
	s.symbol = symbol
	return s
}

// Send the request
func (s *AvgPrice) Do(ctx context.Context, opts ...RequestOption) (res *AvgPriceResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/avgPrice",
		secType:  secTypeNone,
	}
	r.setParam("symbol", s.symbol)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(AvgPriceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Define AvgPrice response data
type AvgPriceResponse struct {
	Mins  uint64 `json:"mins"`
	Price string `json:"price"`
}

// Binance 24hr Ticker Price Change Statistics (GET /api/v3/ticker/24hr)
type Ticker24hr struct {
	c       *Client
	symbol  *string
	symbols *[]string
}

// Symbol set symbol
func (s *Ticker24hr) Symbol(symbol string) *Ticker24hr {
	s.symbol = &symbol
	return s
}

// Symbols set symbols
func (s *Ticker24hr) Symbols(symbols []string) *Ticker24hr {
	s.symbols = &symbols
	return s
}

// Send the request
func (s *Ticker24hr) Do(ctx context.Context, opts ...RequestOption) (res *Ticker24hrResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/ticker/24hr",
		secType:  secTypeNone,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.symbols != nil {
		r.setParam("symbols", *s.symbols)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(Ticker24hrResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Define Ticker24hr response data
type Ticker24hrResponse struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	AskPrice           string `json:"askPrice"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           uint64 `json:"openTime"`
	CloseTime          uint64 `json:"closeTime"`
	FirstId            uint64 `json:"firstId"`
	LastId             uint64 `json:"lastId"`
	Count              uint64 `json:"count"`
}

// Binance Symbol Price Ticker (GET /api/v3/ticker/price)
type TickerPrice struct {
	c       *Client
	symbol  *string
	symbols *[]string
}

// Symbol set symbol
func (s *TickerPrice) Symbol(symbol string) *TickerPrice {
	s.symbol = &symbol
	return s
}

// Symbols set symbols
func (s *TickerPrice) Symbols(symbols []string) *TickerPrice {
	s.symbols = &symbols
	return s
}

// Send the request
func (s *TickerPrice) Do(ctx context.Context, opts ...RequestOption) (res *TickerPriceResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/ticker/price",
		secType:  secTypeNone,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.symbols != nil {
		r.setParam("symbols", *s.symbols)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(TickerPriceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Define TickerPrice response data
type TickerPriceResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

// Binance Symbol Order Book Ticker (GET /api/v3/ticker/bookTicker)
type TickerBookTicker struct {
	c       *Client
	symbol  *string
	symbols *[]string
}

// Symbol set symbol
func (s *TickerBookTicker) Symbol(symbol string) *TickerBookTicker {
	s.symbol = &symbol
	return s
}

// Symbols set symbols
func (s *TickerBookTicker) Symbols(symbols []string) *TickerBookTicker {
	s.symbols = &symbols
	return s
}

func (s *TickerBookTicker) Do(ctx context.Context, opts ...RequestOption) (res []*TickerBookTickerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/ticker/bookTicker",
		secType:  secTypeNone,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.symbols != nil {
		s, _ := json.Marshal(s.symbols)
		r.setParam("symbols", string(s))
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*TickerBookTickerResponse{}, err
	}
	var raw json.RawMessage
	err = json.Unmarshal(data, &raw)
	if err != nil {
		return []*TickerBookTickerResponse{}, err
	}

	if raw[0] == '[' {
		// The response is an array, unmarshal it as before
		res = make([]*TickerBookTickerResponse, 0)
		err = json.Unmarshal(data, &res)
		if err != nil {
			return []*TickerBookTickerResponse{}, err
		}
	} else {
		// The response is a single object, not an array, make sure to add it to the slice
		singleRes := new(TickerBookTickerResponse)
		err = json.Unmarshal(data, &singleRes)
		if err != nil {
			return []*TickerBookTickerResponse{}, err
		}
		res = append(res, singleRes)
	}

	return res, nil
}

// Define TickerBookTicker response data
type TickerBookTickerResponse struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	BidQty   string `json:"bidQty"`
	AskPrice string `json:"askPrice"`
	AskQty   string `json:"askQty"`
}

// Binance Rolling window price change statistics (GET /api/v3/ticker)
type Ticker struct {
	c          *Client
	symbol     string
	windowSize *string
	tickerType *string
}

// Symbol set symbol
func (s *Ticker) Symbol(symbol string) *Ticker {
	s.symbol = symbol
	return s
}

// WindowSize set windowSize
func (s *Ticker) WindowSize(windowSize string) *Ticker {
	s.windowSize = &windowSize
	return s
}

// Type set type
func (s *Ticker) Type(tickerType string) *Ticker {
	s.tickerType = &tickerType
	return s
}

// Send the request
func (s *Ticker) Do(ctx context.Context, opts ...RequestOption) (res *TickerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/ticker",
		secType:  secTypeNone,
	}
	r.setParam("symbol", s.symbol)
	if s.windowSize != nil {
		r.setParam("windowSize", *s.windowSize)
	}
	if s.tickerType != nil {
		r.setParam("type", *s.tickerType)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(TickerResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Define Ticker response data
type TickerResponse struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	AskPrice           string `json:"askPrice"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           uint64 `json:"openTime"`
	CloseTime          uint64 `json:"closeTime"`
	FirstId            uint64 `json:"firstId"`
	LastId             uint64 `json:"lastId"`
	Count              uint64 `json:"count"`
}
