package binance_connector

import (
	"context"
	"encoding/json"
	"strconv"
)

type DepthService struct {
	websocketAPI *WebsocketAPIClient
	symbol       string
	limit        *int
}

func (s *DepthService) Symbol(symbol string) *DepthService {
	s.symbol = symbol
	return s
}

func (s *DepthService) Limit(limit int) *DepthService {
	s.limit = &limit
	return s
}

func (s *DepthService) Do(ctx context.Context) (*DepthResponse, error) {
	parameters := map[string]interface{}{
		"symbol": s.symbol,
	}

	if s.limit != nil {
		parameters["limit"] = s.limit
	}

	id := getUUID()

	payload := map[string]interface{}{
		"id":     id,
		"method": "depth",
		"params": parameters,
	}

	messageCh := make(chan []byte)
	s.websocketAPI.ReqResponseMap[id] = messageCh

	err := s.websocketAPI.SendMessage(payload)
	if err != nil {
		return nil, err
	}

	defer delete(s.websocketAPI.ReqResponseMap, id)

	select {
	case response := <-messageCh:
		var depthResponse DepthResponse
		err = json.Unmarshal(response, &depthResponse)
		if err != nil {
			return nil, err
		}
		return &depthResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type DepthResponse struct {
	ID         string              `json:"id"`
	Status     int                 `json:"status"`
	Error      *WsAPIErrorResponse `json:"error,omitempty"`
	Result     *DepthResult        `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit   `json:"rateLimits"`
}

type DepthResult struct {
	LastUpdateID int64      `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}

type RecentTradesService struct {
	websocketAPI *WebsocketAPIClient
	symbol       string
	limit        *int
}

func (s *RecentTradesService) Symbol(symbol string) *RecentTradesService {
	s.symbol = symbol
	return s
}

func (s *RecentTradesService) Limit(limit int) *RecentTradesService {
	s.limit = &limit
	return s
}

func (s *RecentTradesService) Do(ctx context.Context) (*RecentTradesResponse, error) {
	parameters := map[string]string{
		"symbol": s.symbol,
	}

	if s.limit != nil {
		parameters["limit"] = strconv.Itoa(*s.limit)
	}

	id := getUUID()

	payload := map[string]interface{}{
		"id":     id,
		"method": "trades.recent",
		"params": parameters,
	}

	messageCh := make(chan []byte)
	s.websocketAPI.ReqResponseMap[id] = messageCh

	err := s.websocketAPI.SendMessage(payload)
	if err != nil {
		return nil, err
	}

	defer delete(s.websocketAPI.ReqResponseMap, id)

	select {
	case response := <-messageCh:
		var recentTradesResponse RecentTradesResponse
		err = json.Unmarshal(response, &recentTradesResponse)
		if err != nil {
			return nil, err
		}
		return &recentTradesResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type RecentTradesResponse struct {
	ID         string               `json:"id"`
	Status     int                  `json:"status"`
	Result     []*RecentTradeResult `json:"result"`
	Error      *WsAPIErrorResponse  `json:"error,omitempty"`
	RateLimits []*WsAPIRateLimit    `json:"rateLimits"`
}

type RecentTradeResult struct {
	ID           int64  `json:"id"`
	Price        string `json:"price"`
	Quantity     string `json:"qty"`
	Time         int64  `json:"time"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
	IsBestMatch  bool   `json:"isBestMatch"`
}

type HistoricalTradesService struct {
	websocketAPI *WebsocketAPIClient
	symbol       string
	limit        *int
	fromId       *int64
}

func (s *HistoricalTradesService) Symbol(symbol string) *HistoricalTradesService {
	s.symbol = symbol
	return s
}

func (s *HistoricalTradesService) Limit(limit int) *HistoricalTradesService {
	s.limit = &limit
	return s
}

func (s *HistoricalTradesService) FromId(fromId int64) *HistoricalTradesService {
	s.fromId = &fromId
	return s
}

func (s *HistoricalTradesService) Do(ctx context.Context) (*HistoricalTradesResponse, error) {
	parameters := map[string]interface{}{
		"symbol": s.symbol,
		"apiKey": s.websocketAPI.APIKey,
	}

	if s.limit != nil {
		parameters["limit"] = s.limit
	}

	if s.fromId != nil {
		parameters["fromId"] = s.fromId
	}

	id := getUUID()

	payload := map[string]interface{}{
		"id":     id,
		"method": "trades.historical",
		"params": parameters,
	}

	messageCh := make(chan []byte)
	s.websocketAPI.ReqResponseMap[id] = messageCh

	err := s.websocketAPI.SendMessage(payload)
	if err != nil {
		return nil, err
	}

	defer delete(s.websocketAPI.ReqResponseMap, id)

	select {
	case response := <-messageCh:
		var historicalTradesResponse HistoricalTradesResponse
		err = json.Unmarshal(response, &historicalTradesResponse)
		if err != nil {
			return nil, err
		}
		return &historicalTradesResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type HistoricalTradesResponse struct {
	ID         string                 `json:"id"`
	Status     int                    `json:"status"`
	Error      *WsAPIErrorResponse    `json:"error,omitempty"`
	Result     []*HistoricalTradeData `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit      `json:"rateLimits"`
}

type HistoricalTradeData struct {
	Id           int64  `json:"id"`
	Price        string `json:"price"`
	Quantity     string `json:"qty"`
	QuoteQty     string `json:"quoteQty"`
	Time         uint64 `json:"time"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
	IsBestMatch  bool   `json:"isBestMatch"`
}

type AggregateTradesService struct {
	websocketAPI *WebsocketAPIClient
	symbol       string
	fromId       *int64
	startTime    *uint64
	endTime      *uint64
	limit        *int
}

func (s *AggregateTradesService) Symbol(symbol string) *AggregateTradesService {
	s.symbol = symbol
	return s
}

func (s *AggregateTradesService) FromId(fromId int64) *AggregateTradesService {
	s.fromId = &fromId
	return s
}

func (s *AggregateTradesService) StartTime(startTime uint64) *AggregateTradesService {
	s.startTime = &startTime
	return s
}

func (s *AggregateTradesService) EndTime(endTime uint64) *AggregateTradesService {
	s.endTime = &endTime
	return s
}

func (s *AggregateTradesService) Limit(limit int) *AggregateTradesService {
	s.limit = &limit
	return s
}

func (s *AggregateTradesService) Do(ctx context.Context) (*AggregateTradesResponse, error) {
	parameters := map[string]interface{}{
		"symbol": s.symbol,
	}

	if s.fromId != nil {
		parameters["fromId"] = s.fromId
	}

	if s.startTime != nil {
		parameters["startTime"] = s.startTime
	}

	if s.endTime != nil {
		parameters["endTime"] = s.endTime
	}

	if s.limit != nil {
		parameters["limit"] = s.limit
	}

	id := getUUID()

	payload := map[string]interface{}{
		"id":     id,
		"method": "trades.aggregate",
		"params": parameters,
	}

	messageCh := make(chan []byte)
	s.websocketAPI.ReqResponseMap[id] = messageCh

	err := s.websocketAPI.SendMessage(payload)
	if err != nil {
		return nil, err
	}

	defer delete(s.websocketAPI.ReqResponseMap, id)

	select {
	case response := <-messageCh:
		var aggTradesResponse AggregateTradesResponse
		err = json.Unmarshal(response, &aggTradesResponse)
		if err != nil {
			return nil, err
		}
		return &aggTradesResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type AggregateTradesResponse struct {
	ID         string                `json:"id"`
	Status     int                   `json:"status"`
	Error      *WsAPIErrorResponse   `json:"error,omitempty"`
	Result     []*AggregateTradeData `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit     `json:"rateLimits"`
}

type AggregateTradeData struct {
	AggregateTradeID int64  `json:"a"`
	Price            string `json:"p"`
	Quantity         string `json:"q"`
	FirstTradeID     int64  `json:"f"`
	LastTradeID      int64  `json:"l"`
	Timestamp        int64  `json:"T"`
	IsBuyerMaker     bool   `json:"m"`
	IsBestMatch      bool   `json:"M"`
}

type KlinesService struct {
	websocketAPI *WebsocketAPIClient
	symbol       string
	interval     string
	startTime    *int64
	endTime      *int64
	limit        *int
}

func (s *KlinesService) Symbol(symbol string) *KlinesService {
	s.symbol = symbol
	return s
}

func (s *KlinesService) Interval(interval string) *KlinesService {
	s.interval = interval
	return s
}

func (s *KlinesService) StartTime(startTime int64) *KlinesService {
	s.startTime = &startTime
	return s
}

func (s *KlinesService) EndTime(endTime int64) *KlinesService {
	s.endTime = &endTime
	return s
}

func (s *KlinesService) Limit(limit int) *KlinesService {
	s.limit = &limit
	return s
}

func (s *KlinesService) Do(ctx context.Context) (*WsAPIKlinesResponse, error) {
	parameters := map[string]interface{}{
		"symbol":   s.symbol,
		"interval": s.interval,
	}

	if s.startTime != nil {
		parameters["startTime"] = *s.startTime
	}

	if s.endTime != nil {
		parameters["endTime"] = *s.endTime
	}

	if s.limit != nil {
		parameters["limit"] = *s.limit
	}

	id := getUUID()

	payload := map[string]interface{}{
		"id":     id,
		"method": "klines",
		"params": parameters,
	}

	messageCh := make(chan []byte)
	s.websocketAPI.ReqResponseMap[id] = messageCh

	err := s.websocketAPI.SendMessage(payload)
	if err != nil {
		return nil, err
	}

	defer delete(s.websocketAPI.ReqResponseMap, id)

	select {
	case response := <-messageCh:
		var klinesResponse WsAPIKlinesResponse
		err = json.Unmarshal(response, &klinesResponse)
		if err != nil {
			return nil, err
		}
		return &klinesResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type WsAPIKlinesResponse struct {
	ID         string              `json:"id"`
	Status     int                 `json:"status"`
	Error      *WsAPIErrorResponse `json:"error,omitempty"`
	Result     [][]interface{}     `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit   `json:"rateLimits"`
}

type AvgPriceService struct {
	websocketAPI *WebsocketAPIClient
	symbol       string
}

func (s *AvgPriceService) Symbol(symbol string) *AvgPriceService {
	s.symbol = symbol
	return s
}

func (s *AvgPriceService) Do(ctx context.Context) (*WsAPIAvgPriceResponse, error) {
	parameters := map[string]interface{}{
		"symbol": s.symbol,
	}

	id := getUUID()

	payload := map[string]interface{}{
		"id":     id,
		"method": "avgPrice",
		"params": parameters,
	}

	messageCh := make(chan []byte)
	s.websocketAPI.ReqResponseMap[id] = messageCh

	err := s.websocketAPI.SendMessage(payload)
	if err != nil {
		return nil, err
	}

	defer delete(s.websocketAPI.ReqResponseMap, id)

	select {
	case response := <-messageCh:
		var avgPriceResponse WsAPIAvgPriceResponse
		err = json.Unmarshal(response, &avgPriceResponse)
		if err != nil {
			return nil, err
		}
		return &avgPriceResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type WsAPIAvgPriceResponse struct {
	Id         string              `json:"id"`
	Status     int                 `json:"status"`
	Error      *WsAPIErrorResponse `json:"error,omitempty"`
	Result     *AvgPriceResult     `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit   `json:"rateLimits"`
}

type AvgPriceResult struct {
	Mins  int64  `json:"mins"`
	Price string `json:"price"`
}

type Ticker24hrService struct {
	websocketAPI *WebsocketAPIClient
	symbol       string
}

func (s *Ticker24hrService) Symbol(symbol string) *Ticker24hrService {
	s.symbol = symbol
	return s
}

func (s *Ticker24hrService) Do(ctx context.Context) (*WsAPITicker24hrResponse, error) {
	parameters := map[string]interface{}{
		"symbol": s.symbol,
	}

	id := getUUID()

	payload := map[string]interface{}{
		"id":     id,
		"method": "ticker.24hr",
		"params": parameters,
	}

	messageCh := make(chan []byte)
	s.websocketAPI.ReqResponseMap[id] = messageCh

	err := s.websocketAPI.SendMessage(payload)
	if err != nil {
		return nil, err
	}

	defer delete(s.websocketAPI.ReqResponseMap, id)

	select {
	case response := <-messageCh:
		var twentyFourHrResponse WsAPITicker24hrResponse
		err = json.Unmarshal(response, &twentyFourHrResponse)
		if err != nil {
			return nil, err
		}
		return &twentyFourHrResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type WsAPITicker24hrResponse struct {
	Id         string              `json:"id"`
	Status     int                 `json:"status"`
	Error      *WsAPIErrorResponse `json:"error,omitempty"`
	Result     *Ticker24hrData     `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit   `json:"rateLimits"`
}

type Ticker24hrData struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	CloseQty           string `json:"closeQty"`
	BidPrice           string `json:"bidPrice"`
	BidQty             string `json:"bidQty"`
	AskPrice           string `json:"askPrice"`
	AskQty             string `json:"askQty"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstTradeId       int64  `json:"firstTradeId"`
	LastTradeId        int64  `json:"lastTradeId"`
	TotalTrades        int64  `json:"totalTrades"`
}

type TickerService struct {
	websocketAPI *WebsocketAPIClient
	symbol       string
}

func (s *TickerService) Symbol(symbol string) *TickerService {
	s.symbol = symbol
	return s
}

func (s *TickerService) Do(ctx context.Context) (*WsAPITickerResponse, error) {
	parameters := map[string]interface{}{
		"symbol": s.symbol,
	}

	id := getUUID()

	payload := map[string]interface{}{
		"id":     id,
		"method": "ticker",
		"params": parameters,
	}

	messageCh := make(chan []byte)
	s.websocketAPI.ReqResponseMap[id] = messageCh

	err := s.websocketAPI.SendMessage(payload)
	if err != nil {
		return nil, err
	}

	defer delete(s.websocketAPI.ReqResponseMap, id)

	select {
	case response := <-messageCh:
		var tickerResponse WsAPITickerResponse
		err = json.Unmarshal(response, &tickerResponse)
		if err != nil {
			return nil, err
		}
		return &tickerResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type WsAPITickerResponse struct {
	Id         string              `json:"id"`
	Status     int                 `json:"status"`
	Error      *WsAPIErrorResponse `json:"error,omitempty"`
	Result     *TickerData         `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit   `json:"rateLimits"`
}

type TickerData struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	LastPrice          string `json:"lastPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           uint64 `json:"openTime"`
	CloseTime          uint64 `json:"closeTime"`
	FirstId            int64  `json:"firstId"`
	LastId             int64  `json:"lastId"`
	Count              int64  `json:"count"`
}

type TickerPriceService struct {
	websocketAPI *WebsocketAPIClient
	symbol       string
}

func (s *TickerPriceService) Symbol(symbol string) *TickerPriceService {
	s.symbol = symbol
	return s
}

func (s *TickerPriceService) Do(ctx context.Context) (*WsAPIPriceTickerResponse, error) {
	parameters := map[string]interface{}{
		"symbol": s.symbol,
	}

	id := getUUID()

	payload := map[string]interface{}{
		"id":     id,
		"method": "ticker.price",
		"params": parameters,
	}

	messageCh := make(chan []byte)
	s.websocketAPI.ReqResponseMap[id] = messageCh

	err := s.websocketAPI.SendMessage(payload)
	if err != nil {
		return nil, err
	}

	defer delete(s.websocketAPI.ReqResponseMap, id)

	select {
	case response := <-messageCh:
		var priceTickerResponse WsAPIPriceTickerResponse
		err = json.Unmarshal(response, &priceTickerResponse)
		if err != nil {
			return nil, err
		}
		return &priceTickerResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type WsAPIPriceTickerResponse struct {
	Id         string              `json:"id"`
	Status     int                 `json:"status"`
	Error      *WsAPIErrorResponse `json:"error,omitempty"`
	Result     *tickerPriceData    `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit   `json:"rateLimits"`
}

type tickerPriceData struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type TickerBookService struct {
	websocketAPI *WebsocketAPIClient
	symbol       string
}

func (s *TickerBookService) Symbol(symbol string) *TickerBookService {
	s.symbol = symbol
	return s
}

func (s *TickerBookService) Do(ctx context.Context) (*WsAPIBookTickerResponse, error) {
	parameters := map[string]interface{}{
		"symbol": s.symbol,
	}

	id := getUUID()

	payload := map[string]interface{}{
		"id":     id,
		"method": "ticker.book",
		"params": parameters,
	}

	messageCh := make(chan []byte)
	s.websocketAPI.ReqResponseMap[id] = messageCh

	err := s.websocketAPI.SendMessage(payload)
	if err != nil {
		return nil, err
	}

	defer delete(s.websocketAPI.ReqResponseMap, id)

	select {
	case response := <-messageCh:
		var bookTickerResponse WsAPIBookTickerResponse
		err = json.Unmarshal(response, &bookTickerResponse)
		if err != nil {
			return nil, err
		}
		return &bookTickerResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type WsAPIBookTickerResponse struct {
	Id         string              `json:"id"`
	Status     int                 `json:"status"`
	Error      *WsAPIErrorResponse `json:"error,omitempty"`
	Result     *BookTickerData     `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit   `json:"rateLimits"`
}

type BookTickerData struct {
	BidPrice string `json:"bidPrice"`
	BidQty   string `json:"bidQty"`
	AskPrice string `json:"askPrice"`
	AskQty   string `json:"askQty"`
}

type UIKlinesService struct {
	websocketAPI *WebsocketAPIClient
	symbol       string
	interval     string
	startTime    *int64
	endTime      *int64
	limit        *int
}

func (s *UIKlinesService) Symbol(symbol string) *UIKlinesService {
	s.symbol = symbol
	return s
}

func (s *UIKlinesService) Interval(interval string) *UIKlinesService {
	s.interval = interval
	return s
}

func (s *UIKlinesService) StartTime(startTime int64) *UIKlinesService {
	s.startTime = &startTime
	return s
}

func (s *UIKlinesService) EndTime(endTime int64) *UIKlinesService {
	s.endTime = &endTime
	return s
}

func (s *UIKlinesService) Limit(limit int) *UIKlinesService {
	s.limit = &limit
	return s
}

func (s *UIKlinesService) Do(ctx context.Context) (*UIKlinesResponse, error) {
	parameters := map[string]interface{}{
		"symbol":   s.symbol,
		"interval": s.interval,
	}

	if s.startTime != nil {
		parameters["startTime"] = *s.startTime
	}

	if s.endTime != nil {
		parameters["endTime"] = *s.endTime
	}

	if s.limit != nil {
		parameters["limit"] = *s.limit
	}

	id := getUUID()

	payload := map[string]interface{}{
		"id":     id,
		"method": "uiKlines",
		"params": parameters,
	}

	messageCh := make(chan []byte)
	s.websocketAPI.ReqResponseMap[id] = messageCh

	err := s.websocketAPI.SendMessage(payload)
	if err != nil {
		return nil, err
	}

	defer delete(s.websocketAPI.ReqResponseMap, id)

	select {
	case response := <-messageCh:
		var uiKlinesResponse UIKlinesResponse
		err = json.Unmarshal(response, &uiKlinesResponse)
		if err != nil {
			return nil, err
		}
		return &uiKlinesResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type UIKlinesResponse struct {
	ID         string              `json:"id"`
	Status     int                 `json:"status"`
	Error      *WsAPIErrorResponse `json:"error,omitempty"`
	Result     [][]interface{}     `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit   `json:"rateLimits"`
}
