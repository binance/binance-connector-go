package binance_connector

import (
	"context"
	"encoding/json"
	"strconv"
)

type AccountInformationService struct {
	websocketAPI *WebsocketAPIClient
	recvWindow   *int64
}

func (s *AccountInformationService) RecvWindow(recvWindow int64) *AccountInformationService {
	s.recvWindow = &recvWindow
	return s
}

func (s *AccountInformationService) Do(ctx context.Context) (*AccountInformationResponse, error) {
	parameters := map[string]string{}

	if s.recvWindow != nil {
		parameters["recvWindow"] = strconv.FormatInt(*s.recvWindow, 10)
	}

	signedParams, err := websocketAPISignature(s.websocketAPI.APIKey, s.websocketAPI.APISecret, parameters)
	if err != nil {
		panic(err)
	}

	id := getUUID()

	payload := map[string]interface{}{
		"id":     id,
		"method": "account.status",
		"params": signedParams,
	}

	messageCh := make(chan []byte)
	s.websocketAPI.ReqResponseMap[id] = messageCh

	err2 := s.websocketAPI.SendMessage(payload)
	if err2 != nil {
		return nil, err2
	}

	defer delete(s.websocketAPI.ReqResponseMap, id)

	select {
	case response := <-messageCh:
		var accInfoResponse AccountInformationResponse
		err = json.Unmarshal(response, &accInfoResponse)
		if err != nil {
			return nil, err
		}
		return &accInfoResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type AccountInformationResponse struct {
	ID         string              `json:"id"`
	Status     int                 `json:"status"`
	Error      *WsAPIErrorResponse `json:"error,omitempty"`
	Result     AccountInformation  `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit   `json:"rateLimits"`
}

type AccountInformation struct {
	MakerCommission            int               `json:"makerCommission"`
	TakerCommission            int               `json:"takerCommission"`
	BuyerCommission            int               `json:"buyerCommission"`
	SellerCommission           int               `json:"sellerCommission"`
	CanTrade                   bool              `json:"canTrade"`
	CanWithdraw                bool              `json:"canWithdraw"`
	CanDeposit                 bool              `json:"canDeposit"`
	CommissionRates            WsCommissionRates `json:"commissionRates"`
	Brokered                   bool              `json:"brokered"`
	RequireSelfTradePrevention bool              `json:"requireSelfTradePrevention"`
	UpdateTime                 int64             `json:"updateTime"`
	AccountType                string            `json:"accountType"`
	Balances                   []WsBalance       `json:"balances"`
	Permissions                []string          `json:"permissions"`
}

type WsCommissionRates struct {
	Maker  string `json:"maker"`
	Taker  string `json:"taker"`
	Buyer  string `json:"buyer"`
	Seller string `json:"seller"`
}

type WsBalance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

type AccountOrderRateLimitsService struct {
	websocketAPI *WebsocketAPIClient
	recvWindow   *int64
}

func (s *AccountOrderRateLimitsService) RecvWindow(recvWindow int64) *AccountOrderRateLimitsService {
	s.recvWindow = &recvWindow
	return s
}

func (s *AccountOrderRateLimitsService) Do(ctx context.Context) (*AccountOrderRateLimitsResponse, error) {
	parameters := map[string]string{}

	if s.recvWindow != nil {
		parameters["recvWindow"] = strconv.FormatInt(*s.recvWindow, 10)
	}

	signedParams, err := websocketAPISignature(s.websocketAPI.APIKey, s.websocketAPI.APISecret, parameters)
	if err != nil {
		panic(err)
	}

	id := getUUID()

	payload := map[string]interface{}{
		"id":     id,
		"method": "acc.rateLimits.orders",
		"params": signedParams,
	}

	messageCh := make(chan []byte)
	s.websocketAPI.ReqResponseMap[id] = messageCh

	err2 := s.websocketAPI.SendMessage(payload)
	if err2 != nil {
		return nil, err2
	}

	defer delete(s.websocketAPI.ReqResponseMap, id)

	select {
	case response := <-messageCh:
		var orderRateLimitsResponse AccountOrderRateLimitsResponse
		err = json.Unmarshal(response, &orderRateLimitsResponse)
		if err != nil {
			return nil, err
		}
		return &orderRateLimitsResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type AccountOrderRateLimitsResponse struct {
	ID         string                   `json:"id"`
	Status     int                      `json:"status"`
	Error      *WsAPIErrorResponse      `json:"error,omitempty"`
	Result     []*AccountOrderRateLimit `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit        `json:"rateLimits"`
}

type AccountOrderRateLimit struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
	Count         int    `json:"count"`
}

type AccountOrderHistoryService struct {
	websocketAPI *WebsocketAPIClient
	symbol       string
	orderId      *string
	startTime    *uint64
	endTime      *uint64
	limit        *int
	recvWindow   *int64
}

func (s *AccountOrderHistoryService) Symbol(symbol string) *AccountOrderHistoryService {
	s.symbol = symbol
	return s
}

func (s *AccountOrderHistoryService) OrderId(orderId string) *AccountOrderHistoryService {
	s.orderId = &orderId
	return s
}

func (s *AccountOrderHistoryService) StartTime(startTime uint64) *AccountOrderHistoryService {
	s.startTime = &startTime
	return s
}

func (s *AccountOrderHistoryService) EndTime(endTime uint64) *AccountOrderHistoryService {
	s.endTime = &endTime
	return s
}

func (s *AccountOrderHistoryService) Limit(limit int) *AccountOrderHistoryService {
	s.limit = &limit
	return s
}

func (s *AccountOrderHistoryService) RecvWindow(recvWindow int64) *AccountOrderHistoryService {
	s.recvWindow = &recvWindow
	return s
}

func (s *AccountOrderHistoryService) Do(ctx context.Context) (*AccountOrderHistoryResponse, error) {
	parameters := map[string]string{
		"symbol": s.symbol,
	}

	if s.orderId != nil {
		parameters["orderId"] = *s.orderId
	}

	if s.startTime != nil {
		parameters["startTime"] = strconv.FormatUint(*s.startTime, 10)
	}

	if s.endTime != nil {
		parameters["endTime"] = strconv.FormatUint(*s.endTime, 10)
	}

	if s.limit != nil {
		parameters["limit"] = strconv.Itoa(*s.limit)
	}

	if s.recvWindow != nil {
		parameters["recvWindow"] = strconv.FormatInt(*s.recvWindow, 10)
	}

	signedParams, err := websocketAPISignature(s.websocketAPI.APIKey, s.websocketAPI.APISecret, parameters)
	if err != nil {
		panic(err)
	}

	id := getUUID()

	payload := map[string]interface{}{
		"id":     id,
		"method": "allOrders",
		"params": signedParams,
	}

	messageCh := make(chan []byte)
	s.websocketAPI.ReqResponseMap[id] = messageCh

	err2 := s.websocketAPI.SendMessage(payload)
	if err2 != nil {
		return nil, err2
	}

	defer delete(s.websocketAPI.ReqResponseMap, id)

	select {
	case response := <-messageCh:
		var orderHistoryResponse AccountOrderHistoryResponse
		err = json.Unmarshal(response, &orderHistoryResponse)
		if err != nil {
			return nil, err
		}
		return &orderHistoryResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type AccountOrderHistoryResponse struct {
	ID         string              `json:"id"`
	Status     int                 `json:"status"`
	Error      *WsAPIErrorResponse `json:"error,omitempty"`
	Result     []*OrderHistoryItem `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit   `json:"rateLimits"`
}

type OrderHistoryItem struct {
	Symbol                  string `json:"symbol"`
	OrderID                 int64  `json:"orderId"`
	OrderListID             int64  `json:"orderListId"`
	ClientOrderID           string `json:"clientOrderId"`
	Price                   string `json:"price"`
	OrigQty                 string `json:"origQty"`
	ExecutedQty             string `json:"executedQty"`
	CumulativeQuoteQty      string `json:"cummulativeQuoteQty"`
	Status                  string `json:"status"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	Side                    string `json:"side"`
	StopPrice               string `json:"stopPrice"`
	IcebergQty              string `json:"icebergQty,omitempty"`
	Time                    int64  `json:"time"`
	UpdateTime              int64  `json:"updateTime"`
	IsWorking               bool   `json:"isWorking"`
	WorkingTime             int64  `json:"workingTime"`
	OrigQuoteOrderQty       string `json:"origQuoteOrderQty"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	PreventedMatchID        int64  `json:"preventedMatchId,omitempty"`
	PreventedQuantity       string `json:"preventedQuantity,omitempty"`
}

type AccountOCOHistoryService struct {
	websocketAPI *WebsocketAPIClient
	limit        *int
	fromID       *int64
	startTime    *uint64
	endTime      *uint64
	recvWindow   *int64
}

func (s *AccountOCOHistoryService) Limit(limit int) *AccountOCOHistoryService {
	s.limit = &limit
	return s
}

func (s *AccountOCOHistoryService) FromID(fromID int64) *AccountOCOHistoryService {
	s.fromID = &fromID
	return s
}

func (s *AccountOCOHistoryService) StartTime(startTime uint64) *AccountOCOHistoryService {
	s.startTime = &startTime
	return s
}

func (s *AccountOCOHistoryService) EndTime(endTime uint64) *AccountOCOHistoryService {
	s.endTime = &endTime
	return s
}

func (s *AccountOCOHistoryService) RecvWindow(recvWindow int64) *AccountOCOHistoryService {
	s.recvWindow = &recvWindow
	return s
}

func (s *AccountOCOHistoryService) Do(ctx context.Context) (*AccountOCOHistoryResponse, error) {
	parameters := map[string]string{}

	if s.limit != nil {
		parameters["limit"] = strconv.Itoa(*s.limit)
	}

	if s.fromID != nil {
		parameters["fromId"] = strconv.FormatInt(*s.fromID, 10)
	}

	if s.startTime != nil {
		parameters["startTime"] = strconv.FormatUint(*s.startTime, 10)
	}

	if s.endTime != nil {
		parameters["endTime"] = strconv.FormatUint(*s.endTime, 10)
	}

	if s.recvWindow != nil {
		parameters["recvWindow"] = strconv.FormatInt(*s.recvWindow, 10)
	}

	signedParams, err := websocketAPISignature(s.websocketAPI.APIKey, s.websocketAPI.APISecret, parameters)
	if err != nil {
		panic(err)
	}

	id := getUUID()

	payload := map[string]interface{}{
		"id":     id,
		"method": "allOrderLists",
		"params": signedParams,
	}

	messageCh := make(chan []byte)
	s.websocketAPI.ReqResponseMap[id] = messageCh

	err2 := s.websocketAPI.SendMessage(payload)
	if err2 != nil {
		return nil, err2
	}

	defer delete(s.websocketAPI.ReqResponseMap, id)

	select {
	case response := <-messageCh:
		var ocoHistoryResponse AccountOCOHistoryResponse
		err = json.Unmarshal(response, &ocoHistoryResponse)
		if err != nil {
			return nil, err
		}
		return &ocoHistoryResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type AccountOCOHistoryResponse struct {
	ID         string              `json:"id"`
	Status     int                 `json:"status"`
	Error      *WsAPIErrorResponse `json:"error,omitempty"`
	Result     []*OCOOrder         `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit   `json:"rateLimits"`
}

type OCOOrder struct {
	OrderListID     int64       `json:"orderListId"`
	ContingencyType string      `json:"contingencyType"`
	ListStatusType  string      `json:"listStatusType"`
	ListOrderStatus string      `json:"listOrderStatus"`
	TransactionTime int64       `json:"transactionTime"`
	Symbol          string      `json:"symbol"`
	Orders          []*OCOChild `json:"orders"`
}

type OCOChild struct {
	Symbol           string `json:"symbol"`
	OrderID          int64  `json:"orderId"`
	ClientOrderID    string `json:"clientOrderId"`
	Price            string `json:"price"`
	OrigQty          string `json:"origQty"`
	ExecutedQty      string `json:"executedQty"`
	CummulativeQuote string `json:"cummulativeQuoteQty"`
	Status           string `json:"status"`
	TimeInForce      string `json:"timeInForce"`
	Type             string `json:"type"`
	Side             string `json:"side"`
	StopPrice        string `json:"stopPrice"`
	IcebergQty       string `json:"icebergQty"`
	Time             int64  `json:"time"`
	Update           int64  `json:"updateTime"`
	IsWorking        bool   `json:"isWorking"`
}

type AccountTradeHistoryService struct {
	websocketAPI *WebsocketAPIClient
	symbol       string
	limit        *int
	fromID       *int64
	startTime    *int64
	endTime      *int64
	recvWindow   *int64
}

func (s *AccountTradeHistoryService) Symbol(symbol string) *AccountTradeHistoryService {
	s.symbol = symbol
	return s
}

func (s *AccountTradeHistoryService) Limit(limit int) *AccountTradeHistoryService {
	s.limit = &limit
	return s
}

func (s *AccountTradeHistoryService) FromID(fromID int64) *AccountTradeHistoryService {
	s.fromID = &fromID
	return s
}

func (s *AccountTradeHistoryService) StartTime(startTime int64) *AccountTradeHistoryService {
	s.startTime = &startTime
	return s
}

func (s *AccountTradeHistoryService) EndTime(endTime int64) *AccountTradeHistoryService {
	s.endTime = &endTime
	return s
}

func (s *AccountTradeHistoryService) RecvWindow(recvWindow int64) *AccountTradeHistoryService {
	s.recvWindow = &recvWindow
	return s
}

func (s *AccountTradeHistoryService) Do(ctx context.Context) (*AccountTradeHistoryResponse, error) {
	parameters := map[string]string{
		"symbol": s.symbol,
	}

	if s.limit != nil {
		parameters["limit"] = strconv.Itoa(*s.limit)
	}

	if s.fromID != nil {
		parameters["fromId"] = strconv.FormatInt(*s.fromID, 10)
	}

	if s.startTime != nil {
		parameters["startTime"] = strconv.FormatInt(*s.startTime, 10)
	}

	if s.endTime != nil {
		parameters["endTime"] = strconv.FormatInt(*s.endTime, 10)
	}

	if s.recvWindow != nil {
		parameters["recvWindow"] = strconv.FormatInt(*s.recvWindow, 10)
	}

	signedParams, err := websocketAPISignature(s.websocketAPI.APIKey, s.websocketAPI.APISecret, parameters)
	if err != nil {
		panic(err)
	}

	id := getUUID()

	payload := map[string]interface{}{
		"id":     id,
		"method": "myTrades",
		"params": signedParams,
	}

	messageCh := make(chan []byte)
	s.websocketAPI.ReqResponseMap[id] = messageCh

	err2 := s.websocketAPI.SendMessage(payload)
	if err2 != nil {
		return nil, err2
	}

	defer delete(s.websocketAPI.ReqResponseMap, id)

	select {
	case response := <-messageCh:
		var tradeHistoryResponse AccountTradeHistoryResponse
		err = json.Unmarshal(response, &tradeHistoryResponse)
		if err != nil {
			return nil, err
		}
		return &tradeHistoryResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type AccountTradeHistoryResponse struct {
	ID         string              `json:"id"`
	Status     int                 `json:"status"`
	Error      *WsAPIErrorResponse `json:"error,omitempty"`
	Result     []*TradeHistoryItem `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit   `json:"rateLimits"`
}

type TradeHistoryItem struct {
	Symbol          string `json:"symbol"`
	ID              int64  `json:"id"`
	OrderID         int64  `json:"orderId"`
	OrderListID     int64  `json:"orderListId"`
	Price           string `json:"price"`
	Quantity        string `json:"qty"`
	QuoteQuantity   string `json:"quoteQty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Time            int64  `json:"time"`
	IsBuyer         bool   `json:"isBuyer"`
	IsMaker         bool   `json:"isMaker"`
	IsBestMatch     bool   `json:"isBestMatch"`
}

type AccountPreventedMatchesService struct {
	websocketAPI *WebsocketAPIClient
	recvWindow   *int64
}

func (s *AccountPreventedMatchesService) RecvWindow(recvWindow int64) *AccountPreventedMatchesService {
	s.recvWindow = &recvWindow
	return s
}

func (s *AccountPreventedMatchesService) Do(ctx context.Context) (*AccountPreventedMatchesResponse, error) {
	parameters := map[string]string{}

	if s.recvWindow != nil {
		parameters["recvWindow"] = strconv.FormatInt(*s.recvWindow, 10)
	}

	signedParams, err := websocketAPISignature(s.websocketAPI.APIKey, s.websocketAPI.APISecret, parameters)
	if err != nil {
		panic(err)
	}

	id := getUUID()

	payload := map[string]interface{}{
		"id":     id,
		"method": "myPreventedMatches",
		"params": signedParams,
	}

	messageCh := make(chan []byte)
	s.websocketAPI.ReqResponseMap[id] = messageCh

	err2 := s.websocketAPI.SendMessage(payload)
	if err2 != nil {
		return nil, err2
	}

	defer delete(s.websocketAPI.ReqResponseMap, id)

	select {
	case response := <-messageCh:
		var preventedMatchesResponse AccountPreventedMatchesResponse
		err = json.Unmarshal(response, &preventedMatchesResponse)
		if err != nil {
			return nil, err
		}
		return &preventedMatchesResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type AccountPreventedMatchesResponse struct {
	ID         string              `json:"id"`
	Status     int                 `json:"status"`
	Error      *WsAPIErrorResponse `json:"error,omitempty"`
	Result     []*PreventedMatch   `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit   `json:"rateLimits"`
}

type PreventedMatch struct {
	Symbol                  string `json:"symbol"`
	PreventedMatchId        int64  `json:"preventedMatchId"`
	TakerOrderId            int64  `json:"takerOrderId"`
	MakerOrderId            int64  `json:"makerOrderId"`
	TradeGroupId            int64  `json:"tradeGroupId"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	Price                   string `json:"price"`
	MakerPreventedQuantity  string `json:"makerPreventedQuantity"`
	TransactTime            uint64 `json:"transactTime"`
}
