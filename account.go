package binance_connector

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

// Binance Test New Order endpoint (POST /api/v3/order/test)
type TestNewOrder struct {
	c                   *Client
	symbol              string
	side                string
	orderType           string
	timeInForce         *string
	quantity            *float64
	quoteOrderQty       *float64
	price               *float64
	newClientOrderId    *string
	strategyId          *int
	strategyType        *int
	stopPrice           *float64
	trailingDelta       *int
	icebergQty          *float64
	newOrderRespType    *string
	selfTradePrevention *string
}

// Symbol set symbol
func (s *TestNewOrder) Symbol(symbol string) *TestNewOrder {
	s.symbol = symbol
	return s
}

// Side set side
func (s *TestNewOrder) Side(side string) *TestNewOrder {
	s.side = side
	return s
}

// OrderType set orderType
func (s *TestNewOrder) OrderType(orderType string) *TestNewOrder {
	s.orderType = orderType
	return s
}

// TimeInForce set timeInForce
func (s *TestNewOrder) TimeInForce(timeInForce string) *TestNewOrder {
	s.timeInForce = &timeInForce
	return s
}

// Quantity set quantity
func (s *TestNewOrder) Quantity(quantity float64) *TestNewOrder {
	s.quantity = &quantity
	return s
}

// QuoteOrderQty set quoteOrderQty
func (s *TestNewOrder) QuoteOrderQty(quoteOrderQty float64) *TestNewOrder {
	s.quoteOrderQty = &quoteOrderQty
	return s
}

// Price set price
func (s *TestNewOrder) Price(price float64) *TestNewOrder {
	s.price = &price
	return s
}

// NewClientOrderId set newClientOrderId
func (s *TestNewOrder) NewClientOrderId(newClientOrderId string) *TestNewOrder {
	s.newClientOrderId = &newClientOrderId
	return s
}

// StrategyId set strategyId
func (s *TestNewOrder) StrategyId(strategyId int) *TestNewOrder {
	s.strategyId = &strategyId
	return s
}

// StrategyType set strategyType
func (s *TestNewOrder) StrategyType(strategyType int) *TestNewOrder {
	s.strategyType = &strategyType
	return s
}

// StopPrice set stopPrice
func (s *TestNewOrder) StopPrice(stopPrice float64) *TestNewOrder {
	s.stopPrice = &stopPrice
	return s
}

// TrailingDelta set trailingDelta
func (s *TestNewOrder) TrailingDelta(trailingDelta int) *TestNewOrder {
	s.trailingDelta = &trailingDelta
	return s
}

// IcebergQty set icebergQty
func (s *TestNewOrder) IcebergQty(icebergQty float64) *TestNewOrder {
	s.icebergQty = &icebergQty
	return s
}

// NewOrderRespType set newOrderRespType
func (s *TestNewOrder) NewOrderRespType(newOrderRespType string) *TestNewOrder {
	s.newOrderRespType = &newOrderRespType
	return s
}

// SelfTradePrevention set selfTradePrevention
func (s *TestNewOrder) SelfTradePrevention(selfTradePrevention string) *TestNewOrder {
	s.selfTradePrevention = &selfTradePrevention
	return s
}

// Send the request
func (s *TestNewOrder) Do(ctx context.Context, opts ...RequestOption) (res *AccountOrderBookResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/api/v3/order/test",
		secType:  secTypeNone,
	}
	r.setParam("symbol", s.symbol)
	r.setParam("side", s.side)
	r.setParam("type", s.orderType)
	if s.timeInForce != nil {
		r.setParam("timeInForce", *s.timeInForce)
	}
	if s.quantity != nil {
		r.setParam("quantity", strconv.FormatFloat(*s.quantity, 'f', -1, 64))
	}
	if s.quoteOrderQty != nil {
		r.setParam("quoteOrderQty", *s.quoteOrderQty)
	}
	if s.price != nil {
		r.setParam("price", *s.price)
	}
	if s.newClientOrderId != nil {
		r.setParam("newClientOrderId", *s.newClientOrderId)
	}
	if s.strategyId != nil {
		r.setParam("strategyId", *s.strategyId)
	}
	if s.strategyType != nil {
		r.setParam("strategyType", *s.strategyType)
	}
	if s.stopPrice != nil {
		r.setParam("stopPrice", *s.stopPrice)
	}
	if s.trailingDelta != nil {
		r.setParam("trailingDelta", *s.trailingDelta)
	}
	if s.icebergQty != nil {
		r.setParam("icebergQty", *s.icebergQty)
	}
	if s.newOrderRespType != nil {
		r.setParam("newOrderRespType", *s.newOrderRespType)
	}
	if s.selfTradePrevention != nil {
		r.setParam("selfTradePreventionMode", *s.selfTradePrevention)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(AccountOrderBookResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Create AccountOrderBookResponse
type AccountOrderBookResponse struct {
}

// Binance New Order endpoint (POST /api/v3/order)
// CreateOrderService create order
type CreateOrderService struct {
	c                       *Client
	symbol                  string
	side                    string
	orderType               string
	timeInForce             *string
	quantity                *float64
	quoteOrderQty           *float64
	price                   *float64
	newClientOrderId        *string
	strategyId              *int
	strategyType            *int
	stopPrice               *float64
	trailingDelta           *int
	icebergQty              *float64
	newOrderRespType        *string
	selfTradePreventionMode *string
}

// Symbol set symbol
func (s *CreateOrderService) Symbol(symbol string) *CreateOrderService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *CreateOrderService) Side(side string) *CreateOrderService {
	s.side = side
	return s
}

// Type set type
func (s *CreateOrderService) Type(orderType string) *CreateOrderService {
	s.orderType = orderType
	return s
}

// TimeInForce set timeInForce
func (s *CreateOrderService) TimeInForce(timeInForce string) *CreateOrderService {
	s.timeInForce = &timeInForce
	return s
}

// Quantity set quantity
func (s *CreateOrderService) Quantity(quantity float64) *CreateOrderService {
	s.quantity = &quantity
	return s
}

// QuoteOrderQty set quoteOrderQty
func (s *CreateOrderService) QuoteOrderQty(quoteOrderQty float64) *CreateOrderService {
	s.quoteOrderQty = &quoteOrderQty
	return s
}

// Price set price
func (s *CreateOrderService) Price(price float64) *CreateOrderService {
	s.price = &price
	return s
}

// NewClientOrderId set newClientOrderId
func (s *CreateOrderService) NewClientOrderId(newClientOrderId string) *CreateOrderService {
	s.newClientOrderId = &newClientOrderId
	return s
}

// StrategyId set strategyId
func (s *CreateOrderService) StrategyId(strategyId int) *CreateOrderService {
	s.strategyId = &strategyId
	return s
}

// StrategyType set strategyType
func (s *CreateOrderService) StrategyType(strategyType int) *CreateOrderService {
	s.strategyType = &strategyType
	return s
}

// StopPrice set stopPrice
func (s *CreateOrderService) StopPrice(stopPrice float64) *CreateOrderService {
	s.stopPrice = &stopPrice
	return s
}

// TrailingDelta set trailingDelta
func (s *CreateOrderService) TrailingDelta(trailingDelta int) *CreateOrderService {
	s.trailingDelta = &trailingDelta
	return s
}

// IcebergQuantity set icebergQuantity
func (s *CreateOrderService) IcebergQuantity(icebergQty float64) *CreateOrderService {
	s.icebergQty = &icebergQty
	return s
}

// NewOrderRespType set newOrderRespType
func (s *CreateOrderService) NewOrderRespType(newOrderRespType string) *CreateOrderService {
	s.newOrderRespType = &newOrderRespType
	return s
}

// SelfTradePreventionMode set selfTradePreventionMode
func (s *CreateOrderService) SelfTradePreventionMode(selfTradePreventionMode string) *CreateOrderService {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

const (
	ACK    = 1
	RESULT = 2
	FULL   = 3
)

// Do send request
func (s *CreateOrderService) Do(ctx context.Context, opts ...RequestOption) (res interface{}, err error) {
	respType := ACK
	r := &request{
		method:   http.MethodPost,
		endpoint: "/api/v3/order",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	r.setParam("side", s.side)
	r.setParam("type", s.orderType)
	switch s.orderType {
	case "MARKET":
		respType = FULL
	case "LIMIT":
		respType = FULL
	}
	if s.timeInForce != nil {
		r.setParam("timeInForce", *s.timeInForce)
	}
	if s.quantity != nil {
		r.setParam("quantity", strconv.FormatFloat(*s.quantity, 'f', -1, 64))
	}
	if s.quoteOrderQty != nil {
		r.setParam("quoteOrderQty", *s.quoteOrderQty)
	}
	if s.price != nil {
		r.setParam("price", *s.price)
	}
	if s.newClientOrderId != nil {
		r.setParam("newClientOrderId", *s.newClientOrderId)
	}
	if s.strategyId != nil {
		r.setParam("strategyId", *s.strategyId)
	}
	if s.strategyType != nil {
		r.setParam("strategyType", *s.strategyType)
	}
	if s.stopPrice != nil {
		r.setParam("stopPrice", *s.stopPrice)
	}
	if s.trailingDelta != nil {
		r.setParam("trailingDelta", *s.trailingDelta)
	}
	if s.icebergQty != nil {
		r.setParam("icebergQty", *s.icebergQty)
	}
	if s.newOrderRespType != nil {
		r.setParam("newOrderRespType", *s.newOrderRespType)
		switch *s.newOrderRespType {
		case "ACK":
			respType = ACK
		case "RESULT":
			respType = RESULT
		case "FULL":
			respType = FULL
		}

	}
	if s.selfTradePreventionMode != nil {
		r.setParam("selfTradePreventionMode", *s.selfTradePreventionMode)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	switch respType {
	case ACK:
		res = new(CreateOrderResponseACK)
	case RESULT:
		res = new(CreateOrderResponseRESULT)
	case FULL:
		res = new(CreateOrderResponseFULL)
	}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Create CreateOrderResponseACK
type CreateOrderResponseACK struct {
	Symbol        string `json:"symbol"`
	OrderId       int64  `json:"orderId"`
	OrderListId   int64  `json:"orderListId"`
	ClientOrderId string `json:"clientOrderId"`
	TransactTime  uint64 `json:"transactTime"`
}

// Create CreateOrderResponseRESULT
type CreateOrderResponseRESULT struct {
	Symbol                  string `json:"symbol"`
	OrderId                 int64  `json:"orderId"`
	OrderListId             int64  `json:"orderListId"`
	ClientOrderId           string `json:"clientOrderId"`
	TransactTime            uint64 `json:"transactTime"`
	Price                   string `json:"price"`
	OrigQty                 string `json:"origQty"`
	ExecutedQty             string `json:"executedQty"`
	CumulativeQuoteQty      string `json:"cummulativeQuoteQty"`
	Status                  string `json:"status"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	Side                    string `json:"side"`
	WorkingTime             uint64 `json:"workingTime"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	IcebergQty              string `json:"icebergQty,omitempty"`
	PreventedMatchId        int64  `json:"preventedMatchId,omitempty"`
	PreventedQuantity       string `json:"preventedQuantity,omitempty"`
	StopPrice               string `json:"stopPrice,omitempty"`
	StrategyId              int64  `json:"strategyId,omitempty"`
	StrategyType            int64  `json:"strategyType,omitempty"`
	TrailingDelta           string `json:"trailingDelta,omitempty"`
	TrailingTime            int64  `json:"trailingTime,omitempty"`
}

// Create CreateOrderResponseFULL
type CreateOrderResponseFULL struct {
	Symbol                  string `json:"symbol"`
	OrderId                 int64  `json:"orderId"`
	OrderListId             int64  `json:"orderListId"`
	ClientOrderId           string `json:"clientOrderId"`
	TransactTime            uint64 `json:"transactTime"`
	Price                   string `json:"price"`
	OrigQty                 string `json:"origQty"`
	ExecutedQty             string `json:"executedQty"`
	CumulativeQuoteQty      string `json:"cummulativeQuoteQty"`
	Status                  string `json:"status"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	Side                    string `json:"side"`
	WorkingTime             uint64 `json:"workingTime"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	IcebergQty              string `json:"icebergQty,omitempty"`
	PreventedMatchId        int64  `json:"preventedMatchId,omitempty"`
	PreventedQuantity       string `json:"preventedQuantity,omitempty"`
	StopPrice               string `json:"stopPrice,omitempty"`
	StrategyId              int64  `json:"strategyId,omitempty"`
	StrategyType            int64  `json:"strategyType,omitempty"`
	TrailingDelta           string `json:"trailingDelta,omitempty"`
	TrailingTime            int64  `json:"trailingTime,omitempty"`
	Fills                   []struct {
		Price           string `json:"price"`
		Qty             string `json:"qty"`
		Commission      string `json:"commission"`
		CommissionAsset string `json:"commissionAsset"`
		TradeId         int64  `json:"tradeId"`
	} `json:"fills"`
}

// Binance Cancel Order endpoint (DELETE /api/v3/order)
// CancelOrderService cancel order
type CancelOrderService struct {
	c                  *Client
	symbol             string
	orderId            *int64
	origClientOrderId  *string
	newClientOrderId   *string
	cancelRestrictions *string
}

// Symbol set symbol
func (s *CancelOrderService) Symbol(symbol string) *CancelOrderService {
	s.symbol = symbol
	return s
}

// OrderId set orderId
func (s *CancelOrderService) OrderId(orderId int64) *CancelOrderService {
	s.orderId = &orderId
	return s
}

// OrigClientOrderId set origClientOrderId
func (s *CancelOrderService) OrigClientOrderId(origClientOrderId string) *CancelOrderService {
	s.origClientOrderId = &origClientOrderId
	return s
}

// NewClientOrderId set newClientOrderId
func (s *CancelOrderService) NewClientOrderId(newClientOrderId string) *CancelOrderService {
	s.newClientOrderId = &newClientOrderId
	return s
}

// CancelRestrictions set cancelRestrictions
func (s *CancelOrderService) CancelRestrictions(cancelRestrictions string) *CancelOrderService {
	s.cancelRestrictions = &cancelRestrictions
	return s
}

// Do send request
func (s *CancelOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CancelOrderResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/api/v3/order",
		secType:  secTypeSigned,
	}
	m := params{
		"symbol": s.symbol,
	}
	if s.orderId != nil {
		m["orderId"] = *s.orderId
	}
	if s.origClientOrderId != nil {
		m["origClientOrderId"] = *s.origClientOrderId
	}
	if s.newClientOrderId != nil {
		m["newClientOrderId"] = *s.newClientOrderId
	}
	if s.cancelRestrictions != nil {
		m["cancelRestrictions"] = *s.cancelRestrictions
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CancelOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Binance Cancel all open orders on a symbol (DELETE /api/v3/openOrders)
// CancelOpenOrdersService cancel open orders
type CancelOpenOrdersService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *CancelOpenOrdersService) Symbol(symbol string) *CancelOpenOrdersService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *CancelOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*CancelOrderResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/api/v3/openOrders",
		secType:  secTypeSigned,
	}
	m := params{
		"symbol": s.symbol,
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*CancelOrderResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Create CancelOrderResponse
type CancelOrderResponse struct {
	Symbol              string `json:"symbol"`
	OrigClientOrderId   string `json:"origClientOrderId"`
	OrderId             int64  `json:"orderId"`
	OrderListId         int64  `json:"orderListId"`
	ClientOrderId       string `json:"clientOrderId"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CumulativeQuoteQty  string `json:"cumulativeQuoteQty"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
	SelfTradePrevention string `json:"selfTradePrevention"`
	IcebergQty          string `json:"icebergQty,omitempty"`
	PreventedMatchId    int64  `json:"preventedMatchId,omitempty"`
	PreventedQuantity   string `json:"preventedQuantity,omitempty"`
	StopPrice           string `json:"stopPrice,omitempty"`
	StrategyId          int64  `json:"strategyId,omitempty"`
	StrategyType        int64  `json:"strategyType,omitempty"`
	TrailingDelta       string `json:"trailingDelta,omitempty"`
	TrailingTime        int64  `json:"trailingTime,omitempty"`
}

// Query Order (USER_DATA)
// Binance Query Order (USER_DATA) (GET /api/v3/order)
// GetOrderService get order
type GetOrderService struct {
	c                 *Client
	symbol            string
	orderId           *int64
	origClientOrderId *string
}

// Symbol set symbol
func (s *GetOrderService) Symbol(symbol string) *GetOrderService {
	s.symbol = symbol
	return s
}

// OrderId set orderId
func (s *GetOrderService) OrderId(orderId int64) *GetOrderService {
	s.orderId = &orderId
	return s
}

// OrigClientOrderId set origClientOrderId
func (s *GetOrderService) OrigClientOrderId(origClientOrderId string) *GetOrderService {
	s.origClientOrderId = &origClientOrderId
	return s
}

// Do send request
func (s *GetOrderService) Do(ctx context.Context, opts ...RequestOption) (res *GetOrderResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/order",
		secType:  secTypeSigned,
	}
	m := params{
		"symbol": s.symbol,
	}
	if s.orderId != nil {
		m["orderId"] = *s.orderId
	}
	if s.origClientOrderId != nil {
		m["origClientOrderId"] = *s.origClientOrderId
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetOrderResponse)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Create GetOrderResponse
type GetOrderResponse struct {
	Symbol                  string `json:"symbol"`
	OrderId                 int64  `json:"orderId"`
	OrderListId             int64  `json:"orderListId"`
	ClientOrderId           string `json:"clientOrderId"`
	Price                   string `json:"price"`
	OrigQty                 string `json:"origQty"`
	ExecutedQty             string `json:"executedQty"`
	CumulativeQuoteQty      string `json:"cumulativeQuoteQty"`
	Status                  string `json:"status"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	Side                    string `json:"side"`
	StopPrice               string `json:"stopPrice"`
	IcebergQty              string `json:"icebergQty,omitempty"`
	Time                    uint64 `json:"time"`
	UpdateTime              uint64 `json:"updateTime"`
	IsWorking               bool   `json:"isWorking"`
	WorkingTime             uint64 `json:"workingTime"`
	OrigQuoteOrderQty       string `json:"origQuoteOrderQty"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	PreventedMatchId        int64  `json:"preventedMatchId,omitempty"`
	PreventedQuantity       string `json:"preventedQuantity,omitempty"`
	StrategyId              int64  `json:"strategyId,omitempty"`
	StrategyType            int64  `json:"strategyType,omitempty"`
	TrailingDelta           string `json:"trailingDelta,omitempty"`
	TrailingTime            int64  `json:"trailingTime,omitempty"`
}

// Cancel an Existing Order and Send a New Order (TRADE)
type CancelReplaceService struct {
	c                       *Client
	symbol                  string
	side                    string
	orderType               string
	cancelReplaceMode       string
	timeInForce             *string
	quantity                *float64
	quoteOrderQty           *float64
	price                   *float64
	cancelNewClientOrderId  *string
	cancelOrigClientOrderId *string
	cancelOrderId           *int64
	newClientOrderId        *string
	strategyId              *int32
	strategyType            *int32
	stopPrice               *float64
	trailingDelta           *int64
	icebergQty              *float64
	newOrderRespType        *string
	selfTradePreventionMode *string
	cancelRestrictions      *string
}

// Symbol set symbol
func (s *CancelReplaceService) Symbol(symbol string) *CancelReplaceService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *CancelReplaceService) Side(side string) *CancelReplaceService {
	s.side = side
	return s
}

// OrderType set orderType
func (s *CancelReplaceService) OrderType(orderType string) *CancelReplaceService {
	s.orderType = orderType
	return s
}

// CancelReplaceMode set cancelReplaceMode
func (s *CancelReplaceService) CancelReplaceMode(cancelReplaceMode string) *CancelReplaceService {
	s.cancelReplaceMode = cancelReplaceMode
	return s
}

// TimeInForce set timeInForce
func (s *CancelReplaceService) TimeInForce(timeInForce string) *CancelReplaceService {
	s.timeInForce = &timeInForce
	return s
}

// Quantity set quantity
func (s *CancelReplaceService) Quantity(quantity float64) *CancelReplaceService {
	s.quantity = &quantity
	return s
}

// QuoteOrderQty set quoteOrderQty
func (s *CancelReplaceService) QuoteOrderQty(quoteOrderQty float64) *CancelReplaceService {
	s.quoteOrderQty = &quoteOrderQty
	return s
}

// Price set price
func (s *CancelReplaceService) Price(price float64) *CancelReplaceService {
	s.price = &price
	return s
}

// CancelNewClientOrderId set cancelNewClientOrderId
func (s *CancelReplaceService) CancelNewClientOrderId(cancelNewClientOrderId string) *CancelReplaceService {
	s.cancelNewClientOrderId = &cancelNewClientOrderId
	return s
}

// CancelOrigClientOrderId set cancelOrigClientOrderId
func (s *CancelReplaceService) CancelOrigClientOrderId(cancelOrigClientOrderId string) *CancelReplaceService {
	s.cancelOrigClientOrderId = &cancelOrigClientOrderId
	return s
}

// CancelOrderId set cancelOrderId
func (s *CancelReplaceService) CancelOrderId(cancelOrderId int64) *CancelReplaceService {
	s.cancelOrderId = &cancelOrderId
	return s
}

// NewClientOrderId set newClientOrderId
func (s *CancelReplaceService) NewClientOrderId(newClientOrderId string) *CancelReplaceService {
	s.newClientOrderId = &newClientOrderId
	return s
}

// StrategyId set strategyId
func (s *CancelReplaceService) StrategyId(strategyId int32) *CancelReplaceService {
	s.strategyId = &strategyId
	return s
}

// StrategyType set strategyType
func (s *CancelReplaceService) StrategyType(strategyType int32) *CancelReplaceService {
	s.strategyType = &strategyType
	return s
}

// StopPrice set stopPrice
func (s *CancelReplaceService) StopPrice(stopPrice float64) *CancelReplaceService {
	s.stopPrice = &stopPrice
	return s
}

// TrailingDelta set trailingDelta
func (s *CancelReplaceService) TrailingDelta(trailingDelta int64) *CancelReplaceService {
	s.trailingDelta = &trailingDelta
	return s
}

// IcebergQty set icebergQty
func (s *CancelReplaceService) IcebergQty(icebergQty float64) *CancelReplaceService {
	s.icebergQty = &icebergQty
	return s
}

// NewOrderRespType set newOrderRespType
func (s *CancelReplaceService) NewOrderRespType(newOrderRespType string) *CancelReplaceService {
	s.newOrderRespType = &newOrderRespType
	return s
}

// SelfTradePreventionMode set selfTradePreventionMode
func (s *CancelReplaceService) SelfTradePreventionMode(selfTradePreventionMode string) *CancelReplaceService {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

// CancelRestrictions set cancelRestrictions
func (s *CancelReplaceService) CancelRestrictions(cancelRestrictions string) *CancelReplaceService {
	s.cancelRestrictions = &cancelRestrictions
	return s
}

// Do send request
func (s *CancelReplaceService) Do(ctx context.Context, opts ...RequestOption) (res *CancelReplaceResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/api/v3/order/cancelReplace",
		secType:  secTypeSigned,
	}
	m := params{
		"symbol":            s.symbol,
		"side":              s.side,
		"type":              s.orderType,
		"cancelReplaceMode": s.cancelReplaceMode,
	}
	if s.timeInForce != nil {
		m["timeInForce"] = *s.timeInForce
	}
	if s.quantity != nil {
		m["quantity"] = strconv.FormatFloat(*s.quantity, 'f', -1, 64)
	}
	if s.quoteOrderQty != nil {
		m["quoteOrderQty"] = *s.quoteOrderQty
	}
	if s.price != nil {
		m["price"] = *s.price
	}
	if s.cancelNewClientOrderId != nil {
		m["cancelNewClientOrderId"] = *s.cancelNewClientOrderId
	}
	if s.cancelOrigClientOrderId != nil {
		m["cancelOrigClientOrderId"] = *s.cancelOrigClientOrderId
	}
	if s.cancelOrderId != nil {
		m["cancelOrderId"] = *s.cancelOrderId
	}
	if s.newClientOrderId != nil {
		m["newClientOrderId"] = *s.newClientOrderId
	}
	if s.strategyId != nil {
		m["strategyId"] = *s.strategyId
	}
	if s.strategyType != nil {
		m["strategyType"] = *s.strategyType
	}
	if s.stopPrice != nil {
		m["stopPrice"] = *s.stopPrice
	}
	if s.trailingDelta != nil {
		m["trailingDelta"] = *s.trailingDelta
	}
	if s.icebergQty != nil {
		m["icebergQty"] = *s.icebergQty
	}
	if s.newOrderRespType != nil {
		m["newOrderRespType"] = *s.newOrderRespType
	}
	if s.selfTradePreventionMode != nil {
		m["selfTradePreventionMode"] = *s.selfTradePreventionMode
	}
	if s.cancelRestrictions != nil {
		m["cancelRestrictions"] = *s.cancelRestrictions
	}
	r.setParams(m)
	data, _ := s.c.callAPI(ctx, r, opts...)
	res = new(CancelReplaceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, err
}

type CancelReplaceResponse struct {
	Code           int64  `json:"code,omitempty"`
	Msg            string `json:"msg,omitempty"`
	CancelResult   string `json:"cancelResult,omitempty"`
	NewOrderResult string `json:"newOrderResult,omitempty"`
	CancelResponse *struct {
		Code                    int    `json:"code,omitempty"`
		Msg                     string `json:"msg,omitempty"`
		Symbol                  string `json:"symbol,omitempty"`
		OrigClientOrderId       string `json:"origClientOrderId,omitempty"`
		OrderId                 int64  `json:"orderId,omitempty"`
		OrderListId             int64  `json:"orderListId,omitempty"`
		ClientOrderId           string `json:"clientOrderId,omitempty"`
		Price                   string `json:"price,omitempty"`
		OrigQty                 string `json:"origQty,omitempty"`
		ExecutedQty             string `json:"executedQty,omitempty"`
		CumulativeQuoteQty      string `json:"cumulativeQuoteQty,omitempty"`
		Status                  string `json:"status,omitempty"`
		TimeInForce             string `json:"timeInForce,omitempty"`
		Type                    string `json:"type,omitempty"`
		Side                    string `json:"side,omitempty"`
		SelfTradePreventionMode string `json:"selfTradePreventionMode,omitempty"`
	} `json:"cancelResponse,omitempty"`
	NewOrderResponse *struct {
		Code                    int64    `json:"code,omitempty"`
		Msg                     string   `json:"msg,omitempty"`
		Symbol                  string   `json:"symbol,omitempty"`
		OrderId                 int64    `json:"orderId,omitempty"`
		OrderListId             int64    `json:"orderListId,omitempty"`
		ClientOrderId           string   `json:"clientOrderId,omitempty"`
		TransactTime            uint64   `json:"transactTime,omitempty"`
		Price                   string   `json:"price,omitempty"`
		OrigQty                 string   `json:"origQty,omitempty"`
		ExecutedQty             string   `json:"executedQty,omitempty"`
		CumulativeQuoteQty      string   `json:"cumulativeQuoteQty,omitempty"`
		Status                  string   `json:"status,omitempty"`
		TimeInForce             string   `json:"timeInForce,omitempty"`
		Type                    string   `json:"type,omitempty"`
		Side                    string   `json:"side,omitempty"`
		Fills                   []string `json:"fills,omitempty"`
		SelfTradePreventionMode string   `json:"selfTradePreventionMode,omitempty"`
	} `json:"newOrderResponse,omitempty"`
	Data *struct {
		CancelResult   string `json:"cancelResult,omitempty"`
		NewOrderResult string `json:"newOrderResult,omitempty"`
		CancelResponse *struct {
			Code                    int64  `json:"code,omitempty"`
			Msg                     string `json:"msg,omitempty"`
			Symbol                  string `json:"symbol,omitempty"`
			OrigClientOrderId       string `json:"origClientOrderId,omitempty"`
			OrderId                 int64  `json:"orderId,omitempty"`
			OrderListId             int64  `json:"orderListId,omitempty"`
			ClientOrderId           string `json:"clientOrderId,omitempty"`
			Price                   string `json:"price,omitempty"`
			OrigQty                 string `json:"origQty,omitempty"`
			ExecutedQty             string `json:"executedQty,omitempty"`
			CumulativeQuoteQty      string `json:"cumulativeQuoteQty,omitempty"`
			Status                  string `json:"status,omitempty"`
			TimeInForce             string `json:"timeInForce,omitempty"`
			Type                    string `json:"type,omitempty"`
			Side                    string `json:"side,omitempty"`
			SelfTradePreventionMode string `json:"selfTradePreventionMode,omitempty"`
		} `json:"cancelResponse,omitempty"`
		NewOrderResponse struct {
			Code                    int64    `json:"code,omitempty"`
			Msg                     string   `json:"msg,omitempty"`
			Symbol                  string   `json:"symbol,omitempty"`
			OrderId                 int64    `json:"orderId,omitempty"`
			OrderListId             int64    `json:"orderListId,omitempty"`
			ClientOrderId           string   `json:"clientOrderId,omitempty"`
			TransactTime            uint64   `json:"transactTime,omitempty"`
			Price                   string   `json:"price,omitempty"`
			OrigQty                 string   `json:"origQty,omitempty"`
			ExecutedQty             string   `json:"executedQty,omitempty"`
			CumulativeQuoteQty      string   `json:"cumulativeQuoteQty,omitempty"`
			Status                  string   `json:"status,omitempty"`
			TimeInForce             string   `json:"timeInForce,omitempty"`
			Type                    string   `json:"type,omitempty"`
			Side                    string   `json:"side,omitempty"`
			Fills                   []string `json:"fills,omitempty"`
			SelfTradePreventionMode string   `json:"selfTradePreventionMode,omitempty"`
		} `json:"newOrderResponse"`
	} `json:"data,omitempty"`
}

// Binance Get current open orders (GET /api/v3/openOrders)
// GetOpenOrdersService get open orders
type GetOpenOrdersService struct {
	c      *Client
	symbol *string
}

// Symbol set symbol
func (s *GetOpenOrdersService) Symbol(symbol string) *GetOpenOrdersService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *GetOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*NewOpenOrdersResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/openOrders",
		secType:  secTypeSigned,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*NewOpenOrdersResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Create NewOpenOrdersResponse
type NewOpenOrdersResponse struct {
	Symbol                  string `json:"symbol"`
	OrderId                 int64  `json:"orderId"`
	OrderListId             int64  `json:"orderListId"`
	ClientOrderId           string `json:"clientOrderId"`
	Price                   string `json:"price"`
	OrigQty                 string `json:"origQty"`
	ExecutedQty             string `json:"executedQty"`
	CumulativeQuoteQty      string `json:"cumulativeQuoteQty"`
	Status                  string `json:"status"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	Side                    string `json:"side"`
	StopPrice               string `json:"stopPrice"`
	IcebergQty              string `json:"icebergQty,omitempty"`
	Time                    uint64 `json:"time"`
	UpdateTime              uint64 `json:"updateTime"`
	IsWorking               bool   `json:"isWorking"`
	WorkingTime             uint64 `json:"workingTime"`
	OrigQuoteOrderQty       string `json:"origQuoteOrderQty"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	PreventedMatchId        int64  `json:"preventedMatchId,omitempty"`
	PreventedQuantity       string `json:"preventedQuantity,omitempty"`
	StrategyId              int64  `json:"strategyId,omitempty"`
	StrategyType            int64  `json:"strategyType,omitempty"`
	TrailingDelta           string `json:"trailingDelta,omitempty"`
	TrailingTime            int64  `json:"trailingTime,omitempty"`
}

// Binance Get all account orders; active, canceled, or filled (GET /api/v3/allOrders)
// GetAllOrdersService get all orders
type GetAllOrdersService struct {
	c         *Client
	symbol    string
	orderId   *int64
	startTime *uint64
	endTime   *uint64
	limit     *int
}

// Symbol set symbol
func (s *GetAllOrdersService) Symbol(symbol string) *GetAllOrdersService {
	s.symbol = symbol
	return s
}

// OrderId set orderId
func (s *GetAllOrdersService) OrderId(orderId int64) *GetAllOrdersService {
	s.orderId = &orderId
	return s
}

// StartTime set startTime
func (s *GetAllOrdersService) StartTime(startTime uint64) *GetAllOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetAllOrdersService) EndTime(endTime uint64) *GetAllOrdersService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *GetAllOrdersService) Limit(limit int) *GetAllOrdersService {
	s.limit = &limit
	return s
}

// Do send request
func (s *GetAllOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*NewAllOrdersResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/allOrders",
		secType:  secTypeSigned,
	}
	m := params{
		"symbol": s.symbol,
	}
	if s.orderId != nil {
		m["orderId"] = *s.orderId
	}
	if s.startTime != nil {
		m["startTime"] = *s.startTime
	}
	if s.endTime != nil {
		m["endTime"] = *s.endTime
	}
	if s.limit != nil {
		m["limit"] = *s.limit
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*NewAllOrdersResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Create NewAllOrdersResponse
type NewAllOrdersResponse struct {
	Symbol                  string `json:"symbol"`
	ListClientOrderId       string `json:"listClientOrderId"`
	OrderId                 int64  `json:"orderId"`
	OrderListId             int64  `json:"orderListId"`
	ClientOrderId           string `json:"clientOrderId"`
	Price                   string `json:"price"`
	OrigQty                 string `json:"origQty"`
	ExecutedQty             string `json:"executedQty"`
	CumulativeQuoteQty      string `json:"cumulativeQuoteQty"`
	Status                  string `json:"status"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	Side                    string `json:"side"`
	StopPrice               string `json:"stopPrice"`
	IcebergQty              string `json:"icebergQty,omitempty"`
	Time                    uint64 `json:"time"`
	UpdateTime              uint64 `json:"updateTime"`
	IsWorking               bool   `json:"isWorking"`
	OrigQuoteOrderQty       string `json:"origQuoteOrderQty"`
	WorkingTime             uint64 `json:"workingTime"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	PreventedMatchId        int64  `json:"preventedMatchId,omitempty"`
	PreventedQuantity       string `json:"preventedQuantity,omitempty"`
	StrategyId              int64  `json:"strategyId,omitempty"`
	StrategyType            int64  `json:"strategyType,omitempty"`
	TrailingDelta           string `json:"trailingDelta,omitempty"`
	TrailingTime            int64  `json:"trailingTime,omitempty"`
}

// Binance New OCO (TRADE) (POST /api/v3/order/oco)
// NewOCOService create new OCO order
type NewOCOService struct {
	c                       *Client
	symbol                  string
	listClientOrderId       *string
	side                    string
	quantity                float64
	limitClientOrderId      *string
	limitStrategyId         *int
	limitStrategyType       *int
	price                   float64
	limitIcebergQty         *float64
	trailingDelta           *int
	stopClientOrderId       *string
	stopPrice               float64
	stopStrategyId          *int
	stopStrategyType        *int
	stopLimitPrice          *float64
	stopIcebergQty          *float64
	stopLimitTimeInForce    *string
	newOrderRespType        *string
	selfTradePreventionMode *string
}

// Symbol set symbol
func (s *NewOCOService) Symbol(symbol string) *NewOCOService {
	s.symbol = symbol
	return s
}

// ListClientOrderId set listClientOrderId
func (s *NewOCOService) ListClientOrderId(listClientOrderId string) *NewOCOService {
	s.listClientOrderId = &listClientOrderId
	return s
}

// Side set side
func (s *NewOCOService) Side(side string) *NewOCOService {
	s.side = side
	return s
}

// Quantity set quantity
func (s *NewOCOService) Quantity(quantity float64) *NewOCOService {
	s.quantity = quantity
	return s
}

// LimitClientOrderId set limitClientOrderId
func (s *NewOCOService) LimitClientOrderId(limitClientOrderId string) *NewOCOService {
	s.limitClientOrderId = &limitClientOrderId
	return s
}

// LimitStrategyId set limitStrategyId
func (s *NewOCOService) LimitStrategyId(limitStrategyId int) *NewOCOService {
	s.limitStrategyId = &limitStrategyId
	return s
}

// LimitStrategyType set limitStrategyType
func (s *NewOCOService) LimitStrategyType(limitStrategyType int) *NewOCOService {
	s.limitStrategyType = &limitStrategyType
	return s
}

// Price set price
func (s *NewOCOService) Price(price float64) *NewOCOService {
	s.price = price
	return s
}

// LimitIcebergQty set limitIcebergQty
func (s *NewOCOService) LimitIcebergQty(limitIcebergQty float64) *NewOCOService {
	s.limitIcebergQty = &limitIcebergQty
	return s
}

// TrailingDelta set trailingDelta
func (s *NewOCOService) TrailingDelta(trailingDelta int) *NewOCOService {
	s.trailingDelta = &trailingDelta
	return s
}

// StopClientOrderId set stopClientOrderId
func (s *NewOCOService) StopClientOrderId(stopClientOrderId string) *NewOCOService {
	s.stopClientOrderId = &stopClientOrderId
	return s
}

// StopPrice set stopPrice
func (s *NewOCOService) StopPrice(stopPrice float64) *NewOCOService {
	s.stopPrice = stopPrice
	return s
}

// StopStrategyId set stopStrategyId
func (s *NewOCOService) StopStrategyId(stopStrategyId int) *NewOCOService {
	s.stopStrategyId = &stopStrategyId
	return s
}

// StopStrategyType set stopStrategyType
func (s *NewOCOService) StopStrategyType(stopStrategyType int) *NewOCOService {
	s.stopStrategyType = &stopStrategyType
	return s
}

// StopLimitPrice set stopLimitPrice
func (s *NewOCOService) StopLimitPrice(stopLimitPrice float64) *NewOCOService {
	s.stopLimitPrice = &stopLimitPrice
	return s
}

// StopIcebergQty set stopIcebergQty
func (s *NewOCOService) StopIcebergQty(stopIcebergQty float64) *NewOCOService {
	s.stopIcebergQty = &stopIcebergQty
	return s
}

// StopLimitTimeInForce set stopLimitTimeInForce
func (s *NewOCOService) StopLimitTimeInForce(stopLimitTimeInForce string) *NewOCOService {
	s.stopLimitTimeInForce = &stopLimitTimeInForce
	return s
}

// NewOrderRespType set newOrderRespType
func (s *NewOCOService) NewOrderRespType(newOrderRespType string) *NewOCOService {
	s.newOrderRespType = &newOrderRespType
	return s
}

// selfTradePreventionMode set selfTradePreventionMode
func (s *NewOCOService) SelfTradePreventionMode(selfTradePreventionMode string) *NewOCOService {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

// Do send request
func (s *NewOCOService) Do(ctx context.Context, opts ...RequestOption) (res *OrderOCOResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/api/v3/order/oco",
		secType:  secTypeSigned,
	}
	m := params{
		"symbol":    s.symbol,
		"side":      s.side,
		"quantity":  strconv.FormatFloat(s.quantity, 'f', -1, 64),
		"price":     s.price,
		"stopPrice": s.stopPrice,
	}
	if s.listClientOrderId != nil {
		m["listClientOrderId"] = *s.listClientOrderId
	}
	if s.limitClientOrderId != nil {
		m["limitClientOrderId"] = *s.limitClientOrderId
	}
	if s.limitStrategyId != nil {
		m["limitStrategyId"] = *s.limitStrategyId
	}
	if s.limitStrategyType != nil {
		m["limitStrategyType"] = *s.limitStrategyType
	}
	if s.limitIcebergQty != nil {
		m["limitIcebergQty"] = *s.limitIcebergQty
	}
	if s.trailingDelta != nil {
		m["trailingDelta"] = *s.trailingDelta
	}
	if s.stopClientOrderId != nil {
		m["stopClientOrderId"] = *s.stopClientOrderId
	}
	if s.stopStrategyId != nil {
		m["stopStrategyId"] = *s.stopStrategyId
	}
	if s.stopStrategyType != nil {
		m["stopStrategyType"] = *s.stopStrategyType
	}
	if s.stopLimitPrice != nil {
		m["stopLimitPrice"] = *s.stopLimitPrice
	}
	if s.stopIcebergQty != nil {
		m["stopIcebergQty"] = *s.stopIcebergQty
	}
	if s.stopLimitTimeInForce != nil {
		m["stopLimitTimeInForce"] = *s.stopLimitTimeInForce
	}
	if s.newOrderRespType != nil {
		m["newOrderRespType"] = *s.newOrderRespType
	}
	if s.selfTradePreventionMode != nil {
		m["selfTradePreventionMode"] = *s.selfTradePreventionMode
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(OrderOCOResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type OrderOCOResponse struct {
	OrderListId       int64  `json:"orderListId"`
	ContingencyType   string `json:"contingencyType"`
	ListStatusType    string `json:"listStatusType"`
	ListOrderStatus   string `json:"listOrderStatus"`
	ListClientOrderId string `json:"listClientOrderId"`
	TransactionTime   uint64 `json:"transactionTime"`
	Symbol            string `json:"symbol"`
	Orders            []struct {
		Symbol        string `json:"symbol"`
		OrderId       int64  `json:"orderId"`
		ClientOrderId string `json:"clientOrderId"`
	} `json:"orders"`
	OrderReports []struct {
		Symbol                  string  `json:"symbol"`
		OrderId                 int64   `json:"orderId"`
		OrderListId             int64   `json:"orderListId"`
		ClientOrderId           string  `json:"clientOrderId"`
		TransactTime            uint64  `json:"transactTime"`
		Price                   float64 `json:"price"`
		OrigQty                 float64 `json:"origQty"`
		ExecutedQty             float64 `json:"executedQty"`
		CummulativeQuoteQty     float64 `json:"cummulativeQuoteQty"`
		Status                  string  `json:"status"`
		TimeInForce             string  `json:"timeInForce"`
		Type                    string  `json:"type"`
		Side                    string  `json:"side"`
		StopPrice               string  `json:"stopPrice"`
		WorkingTime             uint64  `json:"workingTime"`
		SelfTradePreventionMode string  `json:"selfTradePreventionMode"`
		IcebergQty              string  `json:"icebergQty,omitempty"`
		PreventedMatchId        int64   `json:"preventedMatchId,omitempty"`
		PreventedQuantity       string  `json:"preventedQuantity,omitempty"`
		StrategyId              int64   `json:"strategyId,omitempty"`
		StrategyType            int64   `json:"strategyType,omitempty"`
		TrailingDelta           string  `json:"trailingDelta,omitempty"`
		TrailingTime            int64   `json:"trailingTime,omitempty"`
	} `json:"orderReports"`
}

// Binance Cancel OCO (TRADE) (DELETE /api/v3/orderList)
// CancelOCOService cancel OCO order
type CancelOCOService struct {
	c                 *Client
	symbol            string
	orderListId       *int
	listClientOrderId *string
	newClientOrderId  *string
}

// Symbol set symbol
func (s *CancelOCOService) Symbol(symbol string) *CancelOCOService {
	s.symbol = symbol
	return s
}

// OrderListId set orderListId
func (s *CancelOCOService) OrderListId(orderListId int) *CancelOCOService {
	s.orderListId = &orderListId
	return s
}

// ListClientId set listClientId
func (s *CancelOCOService) ListClientOrderId(ListClientOrderId string) *CancelOCOService {
	s.listClientOrderId = &ListClientOrderId
	return s
}

// NewClientOrderId set newClientOrderId
func (s *CancelOCOService) NewClientOrderId(newClientOrderId string) *CancelOCOService {
	s.newClientOrderId = &newClientOrderId
	return s
}

// Do send request
func (s *CancelOCOService) Do(ctx context.Context, opts ...RequestOption) (res *OrderOCOResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/api/v3/orderList",
		secType:  secTypeSigned,
	}
	m := params{
		"symbol": s.symbol,
	}
	if s.orderListId != nil {
		m["orderListId"] = *s.orderListId
	}
	if s.listClientOrderId != nil {
		m["listClientOrderId"] = *s.listClientOrderId
	}
	if s.newClientOrderId != nil {
		m["newClientOrderId"] = *s.newClientOrderId
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(OrderOCOResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Binance Query OCO (USER_DATA) (GET /api/v3/orderList)
// QueryOCOService query OCO order
type QueryOCOService struct {
	c                 *Client
	orderListId       *int64
	origClientOrderId *string
}

// OrderListId set orderListId
func (s *QueryOCOService) OrderListId(orderListId int64) *QueryOCOService {
	s.orderListId = &orderListId
	return s
}

// OrigClientOrderId set origClientOrderId
func (s *QueryOCOService) OrigClientOrderId(origClientOrderId string) *QueryOCOService {
	s.origClientOrderId = &origClientOrderId
	return s
}

// Do send request
func (s *QueryOCOService) Do(ctx context.Context, opts ...RequestOption) (res *OCOResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/orderList",
		secType:  secTypeSigned,
	}
	if s.orderListId != nil {
		r.setParam("orderListId", *s.orderListId)
	}
	if s.origClientOrderId != nil {
		r.setParam("origClientOrderId", s.origClientOrderId)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(OCOResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Binance Query all OCO (USER_DATA) (GET /api/v3/allOrderList)
// QueryAllOCOService query all OCO order
type QueryAllOCOService struct {
	c         *Client
	fromId    *int64
	startTime *uint64
	endTime   *uint64
	limit     *int
}

// FromId set fromId
func (s *QueryAllOCOService) FromId(fromId int64) *QueryAllOCOService {
	s.fromId = &fromId
	return s
}

// StartTime set startTime
func (s *QueryAllOCOService) StartTime(startTime uint64) *QueryAllOCOService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *QueryAllOCOService) EndTime(endTime uint64) *QueryAllOCOService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *QueryAllOCOService) Limit(limit int) *QueryAllOCOService {
	s.limit = &limit
	return s
}

// Do send request
func (s *QueryAllOCOService) Do(ctx context.Context, opts ...RequestOption) (res []*OCOResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/allOrderList",
		secType:  secTypeSigned,
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
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*OCOResponse{}, err
	}
	res = make([]*OCOResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*OCOResponse{}, err
	}
	return res, nil
}

// Create OCOResponse
type OCOResponse struct {
	OrderListId       int64  `json:"orderListId"`
	ContingencyType   string `json:"contingencyType"`
	ListStatusType    string `json:"listStatusType"`
	ListOrderStatus   string `json:"listOrderStatus"`
	ListClientOrderId string `json:"listClientOrderId"`
	TransactionTime   uint64 `json:"transactionTime"`
	Symbol            string `json:"symbol"`
	Orders            []struct {
		Symbol        string `json:"symbol"`
		OrderId       int64  `json:"orderId"`
		ClientOrderId string `json:"clientOrderId"`
	} `json:"orders"`
}

// Binance Query open OCO (USER_DATA) (GET /api/v3/openOrderList)
// QueryOpenOCOService query open OCO order
type QueryOpenOCOService struct {
	c *Client
}

// Do send request
func (s *QueryOpenOCOService) Do(ctx context.Context, opts ...RequestOption) (res *OCOResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/openOrderList",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(OCOResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Binance Account Information (USER_DATA) (GET /api/v3/account)
// GetAccountService get account information
type GetAccountService struct {
	c *Client
}

// Do send request
func (s *GetAccountService) Do(ctx context.Context, opts ...RequestOption) (res *AccountResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/account",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(AccountResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Create AccountResponse
type AccountResponse struct {
	MakerCommission  int64     `json:"makerCommission"`
	TakerCommission  int64     `json:"takerCommission"`
	BuyerCommission  int64     `json:"buyerCommission"`
	SellerCommission int64     `json:"sellerCommission"`
	CanTrade         bool      `json:"canTrade"`
	CanWithdraw      bool      `json:"canWithdraw"`
	CanDeposit       bool      `json:"canDeposit"`
	UpdateTime       uint64    `json:"updateTime"`
	AccountType      string    `json:"accountType"`
	Balances         []Balance `json:"balances"`
	Permissions      []string  `json:"permissions"`
}

// Balance define user balance of your account
type Balance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

// Binance Get trades for a specific account and symbol (USER_DATA) (GET /api/v3/myTrades)
// GetMyTradesService get trades for a specific account and symbol
type GetMyTradesService struct {
	c         *Client
	symbol    string
	orderId   *int64
	startTime *uint64
	endTime   *uint64
	fromId    *int64
	limit     *int
}

// Symbol set symbol
func (s *GetMyTradesService) Symbol(symbol string) *GetMyTradesService {
	s.symbol = symbol
	return s
}

// OrderId set orderId
func (s *GetMyTradesService) OrderId(orderId int64) *GetMyTradesService {
	s.orderId = &orderId
	return s
}

// StartTime set startTime
func (s *GetMyTradesService) StartTime(startTime uint64) *GetMyTradesService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetMyTradesService) EndTime(endTime uint64) *GetMyTradesService {
	s.endTime = &endTime
	return s
}

// FromId set fromId
func (s *GetMyTradesService) FromId(fromId int64) *GetMyTradesService {
	s.fromId = &fromId
	return s
}

// Limit set limit
func (s *GetMyTradesService) Limit(limit int) *GetMyTradesService {
	s.limit = &limit
	return s
}

// Do send request
func (s *GetMyTradesService) Do(ctx context.Context, opts ...RequestOption) (res []*AccountTradeListResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/myTrades",
		secType:  secTypeSigned,
	}
	m := params{
		"symbol": s.symbol,
	}
	if s.orderId != nil {
		m["orderId"] = *s.orderId
	}
	if s.startTime != nil {
		m["startTime"] = *s.startTime
	}
	if s.endTime != nil {
		m["endTime"] = *s.endTime
	}
	if s.fromId != nil {
		m["fromId"] = *s.fromId
	}
	if s.limit != nil {
		m["limit"] = *s.limit
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*AccountTradeListResponse{}, err
	}
	res = make([]*AccountTradeListResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*AccountTradeListResponse{}, err
	}
	return res, nil
}

//type AccountTradeListResponse []AccountTrade

type AccountTradeListResponse struct {
	Id              int64  `json:"id"`
	Symbol          string `json:"symbol"`
	OrderId         int64  `json:"orderId"`
	OrderListId     int64  `json:"orderListId"`
	Price           string `json:"price"`
	Quantity        string `json:"qty"`
	QuoteQuantity   string `json:"quoteQty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Time            uint64 `json:"time"`
	IsBuyer         bool   `json:"isBuyer"`
	IsMaker         bool   `json:"isMaker"`
	IsBestMatch     bool   `json:"isBestMatch"`
}

// Query Current Order Count Usage (TRADE)
// GetQueryCurrentOrderCountUsageService query current order count usage
type GetQueryCurrentOrderCountUsageService struct {
	c *Client
}

// Do send request
func (s *GetQueryCurrentOrderCountUsageService) Do(ctx context.Context, opts ...RequestOption) (res []*QueryCurrentOrderCountUsageResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/rateLimit/order",
		secType:  secTypeSigned,
	}
	res = make([]*QueryCurrentOrderCountUsageResponse, 0)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

// Create QueryCurrentOrderCountUsageResponse
type QueryCurrentOrderCountUsageResponse struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
	Count         int    `json:"count"`
}

// Query Prevented Matches (USER_DATA)
// GetQueryPreventedMatchesService query prevented matches
type GetQueryPreventedMatchesService struct {
	c                    *Client
	symbol               string
	preventMatchId       *int64
	orderId              *int64
	fromPreventedMatchId *int64
	limit                *int
}

// Symbol set symbol
func (s *GetQueryPreventedMatchesService) Symbol(symbol string) *GetQueryPreventedMatchesService {
	s.symbol = symbol
	return s
}

// PreventMatchId set preventMatchId
func (s *GetQueryPreventedMatchesService) PreventMatchId(preventMatchId int64) *GetQueryPreventedMatchesService {
	s.preventMatchId = &preventMatchId
	return s
}

// OrderId set orderId
func (s *GetQueryPreventedMatchesService) OrderId(orderId int64) *GetQueryPreventedMatchesService {
	s.orderId = &orderId
	return s
}

// FromPreventedMatchId set fromPreventedMatchId
func (s *GetQueryPreventedMatchesService) FromPreventedMatchId(fromPreventedMatchId int64) *GetQueryPreventedMatchesService {
	s.fromPreventedMatchId = &fromPreventedMatchId
	return s
}

// Limit set limit
func (s *GetQueryPreventedMatchesService) Limit(limit int) *GetQueryPreventedMatchesService {
	s.limit = &limit
	return s
}

// Do send request
func (s *GetQueryPreventedMatchesService) Do(ctx context.Context, opts ...RequestOption) (res *QueryPreventedMatchesResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/myPreventedMatches",
		secType:  secTypeSigned,
	}
	m := params{
		"symbol": s.symbol,
	}
	if s.preventMatchId != nil {
		m["preventedMatchId"] = *s.preventMatchId
	}
	if s.orderId != nil {
		m["orderId"] = *s.orderId
	}
	if s.fromPreventedMatchId != nil {
		m["fromPreventedMatchId"] = *s.fromPreventedMatchId
	}
	if s.limit != nil {
		m["limit"] = *s.limit
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(QueryPreventedMatchesResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Create QueryPreventedMatchesResponse
type QueryPreventedMatchesResponse struct {
	PreventedMatches []struct {
		Symbol                  string `json:"symbol"`
		PreventedMatchId        int64  `json:"preventedMatchId"`
		TakerOrderId            int64  `json:"takerOrderId"`
		MakerOrderId            int64  `json:"makerOrderId"`
		TradeGroupId            int64  `json:"tradeGroupId"`
		SelfTradePreventionMode string `json:"selfTradePreventionMode"`
		Price                   string `json:"price"`
		MakerPreventedQuantity  string `json:"makerPreventedQuantity"`
		TransactTime            uint64 `json:"transactTime"`
	} `json:"preventedMatches"`
}
