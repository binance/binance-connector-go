package binance_connector

import (
	"context"
	"net/http"

	json "github.com/json-iterator/go"
)

// Get all margin assets API Endpoint
const (
	getAllMarginAssetsEndpoint = "/sapi/v1/margin/allAssets"
)

// GetAllMarginAssetsService get all margin assets
type GetAllMarginAssetsService struct {
	c *Client
}

// Do send request
func (s *GetAllMarginAssetsService) Do(ctx context.Context, opts ...RequestOption) (res []*GetAllMarginAssetsResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: getAllMarginAssetsEndpoint,
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*GetAllMarginAssetsResponse{}, err
	}
	res = make([]*GetAllMarginAssetsResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*GetAllMarginAssetsResponse{}, err
	}
	return res, nil
}

// GetAllMarginAssetsResponse define get all margin assets response
type GetAllMarginAssetsResponse struct {
	AssetFullName  string `json:"assetFullName"`
	AssetName      string `json:"assetName"`
	IsBorrowable   bool   `json:"isBorrowable"`
	IsMortgageable bool   `json:"isMortgageable"`
	MinLoanAmt     string `json:"minLoanAmt"`
	MaxLoanAmt     string `json:"maxLoanAmt"`
	MinMortgageAmt string `json:"minMortgageAmt"`
	MaxMortgageAmt string `json:"maxMortgageAmt"`
	Asset          string `json:"asset"`
}

// Get all margin pairs API Endpoint
const (
	getAllMarginPairsEndpoint = "/sapi/v1/margin/allPairs"
)

// GetAllMarginPairsService get all margin pairs
type GetAllMarginPairsService struct {
	c *Client
}

// Do send request
func (s *GetAllMarginPairsService) Do(ctx context.Context, opts ...RequestOption) (res []*GetAllMarginPairsResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: getAllMarginPairsEndpoint,
		secType:  secTypeAPIKey,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*GetAllMarginPairsResponse{}, err
	}
	res = make([]*GetAllMarginPairsResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*GetAllMarginPairsResponse{}, err
	}
	return res, nil
}

// GetAllMarginPairsResponse define get all margin pairs response
type GetAllMarginPairsResponse struct {
	Base          string `json:"base"`
	Id            int    `json:"id"`
	IsBuyAllowed  bool   `json:"isBuyAllowed"`
	IsMarginTrade bool   `json:"isMarginTrade"`
	IsSellAllowed bool   `json:"isSellAllowed"`
	Quote         string `json:"quote"`
	Symbol        string `json:"symbol"`
}

// Query Margin Price Index API Endpoint
const (
	queryMarginPriceIndexEndpoint = "/sapi/v1/margin/priceIndex"
)

// QueryMarginPriceIndexService query margin price index
type QueryMarginPriceIndexService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *QueryMarginPriceIndexService) Symbol(symbol string) *QueryMarginPriceIndexService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *QueryMarginPriceIndexService) Do(ctx context.Context, opts ...RequestOption) (res *QueryMarginPriceIndexResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: queryMarginPriceIndexEndpoint,
		secType:  secTypeAPIKey,
	}
	r.setParam("symbol", s.symbol)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &QueryMarginPriceIndexResponse{}, err
	}
	res = new(QueryMarginPriceIndexResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return &QueryMarginPriceIndexResponse{}, err
	}
	return res, nil
}

// QueryMarginPriceIndexResponse define query margin price index response
type QueryMarginPriceIndexResponse struct {
	CalcTime int64  `json:"calcTime"`
	Price    string `json:"price"`
	Symbol   string `json:"symbol"`
}

// Query Margin Available Inventory (USER_DATA)
const (
	queryMarginAvailableInventoryEndpoint = "/sapi/v1/margin/available-inventory"
)

// QueryMarginAvailableInventoryService query margin available inventory
type QueryMarginAvailableInventoryService struct {
	c          *Client
	marginType string
}

// MarginType set marginType
func (s *QueryMarginAvailableInventoryService) MarginType(marginType string) *QueryMarginAvailableInventoryService {
	s.marginType = marginType
	return s
}

// Do send request
func (s *QueryMarginAvailableInventoryService) Do(ctx context.Context, opts ...RequestOption) (res *QueryMarginAvailableInventoryResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: queryMarginAvailableInventoryEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("type", s.marginType)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &QueryMarginAvailableInventoryResponse{}, err
	}
	res = new(QueryMarginAvailableInventoryResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return &QueryMarginAvailableInventoryResponse{}, err
	}
	return res, nil
}

// QueryMarginAvailableInventoryResponse define query margin available inventory response
type QueryMarginAvailableInventoryResponse struct {
	Assets     map[string]string `json:"assets"`
	UpdateTime int64             `json:"updateTime"`
}

// Query Liability Coin Leverage Bracket in Cross Margin Pro Mode(MARKET_DATA)
const (
	queryLiabilityCoinLeverageBracketEndpoint = "/sapi/v1/margin/leverageBracket"
)

// QueryLiabilityCoinLeverageBracketService query liability coin leverage bracket
type QueryLiabilityCoinLeverageBracketService struct {
	c *Client
}

// Do send request
func (s *QueryLiabilityCoinLeverageBracketService) Do(ctx context.Context, opts ...RequestOption) (res []*QueryLiabilityCoinLeverageBracketResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: queryLiabilityCoinLeverageBracketEndpoint,
		secType:  secTypeAPIKey,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*QueryLiabilityCoinLeverageBracketResponse{}, err
	}
	res = make([]*QueryLiabilityCoinLeverageBracketResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*QueryLiabilityCoinLeverageBracketResponse{}, err
	}
	return res, nil
}

// QueryLiabilityCoinLeverageBracketResponse define query liability coin leverage bracket response
type QueryLiabilityCoinLeverageBracketResponse struct {
	AssetNames []string `json:"assetNames"`
	Rank       int      `json:"rank"`
	Brackets   []struct {
		Leverage              int     `json:"leverage"`
		MaxDebt               float64 `json:"maxDebt"`
		MaintenanceMarginRate float64 `json:"maintenanceMarginRate"`
		InitialMarginRate     float64 `json:"initialMarginRate"`
		FastNum               float64 `json:"fastNum"`
	} `json:"brackets"`
}

// Margin Accouunt New Order (TRADE) API Endpoint
const (
	marginAccountNewOrderEndpoint = "/sapi/v1/margin/order"
)

// MarginAccountNewOrderService margin account new order
type MarginAccountNewOrderService struct {
	c                *Client
	symbol           string
	isIsolated       *string
	side             string
	orderType        string
	quantity         *float64
	quoteOrderQty    *float64
	price            *float64
	stopPrice        *float64
	newClientOrderId *string
	icebergQty       *float64
	newOrderRespType *string
	sideEffectType   *string
	timeInForce      *string
}

// Symbol set symbol
func (s *MarginAccountNewOrderService) Symbol(symbol string) *MarginAccountNewOrderService {
	s.symbol = symbol
	return s
}

// IsIsolated set isIsolated
func (s *MarginAccountNewOrderService) IsIsolated(isIsolated string) *MarginAccountNewOrderService {
	s.isIsolated = &isIsolated
	return s
}

// Side set side
func (s *MarginAccountNewOrderService) Side(side string) *MarginAccountNewOrderService {
	s.side = side
	return s
}

// OrderType set orderType
func (s *MarginAccountNewOrderService) OrderType(orderType string) *MarginAccountNewOrderService {
	s.orderType = orderType
	return s
}

// Quantity set quantity
func (s *MarginAccountNewOrderService) Quantity(quantity float64) *MarginAccountNewOrderService {
	s.quantity = &quantity
	return s
}

// QuoteOrderQty set quoteOrderQty
func (s *MarginAccountNewOrderService) QuoteOrderQty(quoteOrderQty float64) *MarginAccountNewOrderService {
	s.quoteOrderQty = &quoteOrderQty
	return s
}

// Price set price
func (s *MarginAccountNewOrderService) Price(price float64) *MarginAccountNewOrderService {
	s.price = &price
	return s
}

// StopPrice set stopPrice
func (s *MarginAccountNewOrderService) StopPrice(stopPrice float64) *MarginAccountNewOrderService {
	s.stopPrice = &stopPrice
	return s
}

// NewClientOrderId set newClientOrderId
func (s *MarginAccountNewOrderService) NewClientOrderId(newClientOrderId string) *MarginAccountNewOrderService {
	s.newClientOrderId = &newClientOrderId
	return s
}

// IcebergQty set icebergQty
func (s *MarginAccountNewOrderService) IcebergQty(icebergQty float64) *MarginAccountNewOrderService {
	s.icebergQty = &icebergQty
	return s
}

// NewOrderRespType set newOrderRespType
func (s *MarginAccountNewOrderService) NewOrderRespType(newOrderRespType string) *MarginAccountNewOrderService {
	s.newOrderRespType = &newOrderRespType
	return s
}

// SideEffectType set sideEffectType
func (s *MarginAccountNewOrderService) SideEffectType(sideEffectType string) *MarginAccountNewOrderService {
	s.sideEffectType = &sideEffectType
	return s
}

// TimeInForce set timeInForce
func (s *MarginAccountNewOrderService) TimeInForce(timeInForce string) *MarginAccountNewOrderService {
	s.timeInForce = &timeInForce
	return s
}

// Do send request
func (s *MarginAccountNewOrderService) Do(ctx context.Context, opts ...RequestOption) (res interface{}, err error) {
	respType := ACK
	r := &request{
		method:   http.MethodPost,
		endpoint: marginAccountNewOrderEndpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol": s.symbol,
		"side":   s.side,
		"type":   s.orderType,
	}
	switch s.orderType {
	case "MARKET":
		respType = FULL
	case "LIMIT":
		respType = FULL
	}
	if s.isIsolated != nil {
		m["isIsolated"] = *s.isIsolated
	}
	if s.quantity != nil {
		m["quantity"] = *s.quantity
	}
	if s.quoteOrderQty != nil {
		m["quoteOrderQty"] = *s.quoteOrderQty
	}
	if s.price != nil {
		m["price"] = *s.price
	}
	if s.stopPrice != nil {
		m["stopPrice"] = *s.stopPrice
	}
	if s.newClientOrderId != nil {
		m["newClientOrderId"] = *s.newClientOrderId
	}
	if s.icebergQty != nil {
		m["icebergQty"] = *s.icebergQty
	}
	if s.newOrderRespType != nil {
		m["newOrderRespType"] = *s.newOrderRespType
		switch *s.newOrderRespType {
		case "ACK":
			respType = ACK
		case "RESULT":
			respType = RESULT
		case "FULL":
			respType = FULL
		}
	}
	if s.sideEffectType != nil {
		m["sideEffectType"] = *s.sideEffectType
	}
	if s.timeInForce != nil {
		m["timeInForce"] = *s.timeInForce
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	switch respType {
	case ACK:
		res = new(MarginAccountNewOrderResponseACK)
	case RESULT:
		res = new(MarginAccountNewOrderResponseRESULT)
	case FULL:
		res = new(MarginAccountNewOrderResponseFULL)
	}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Create MarginAccountNewOrderResponseACK
type MarginAccountNewOrderResponseACK struct {
	Symbol        string `json:"symbol"`
	OrderId       int64  `json:"orderId"`
	ClientOrderId int64  `json:"clientOrderId"`
	IsIsolated    bool   `json:"isIsolated"`
	TransactTime  uint64 `json:"transactTime"`
}

// Create MarginAccountNewOrderResponseRESULT
type MarginAccountNewOrderResponseRESULT struct {
	Symbol              string `json:"symbol"`
	OrderId             int64  `json:"orderId"`
	ClientOrderId       string `json:"clientOrderId"`
	TransactTime        uint64 `json:"transactTime"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	IsIsolated          bool   `json:"isIsolated"`
	Side                string `json:"side"`
}

// Create MarginAccountNewOrderResponseFULL
type MarginAccountNewOrderResponseFULL struct {
	Symbol                string  `json:"symbol"`
	OrderId               int64   `json:"orderId"`
	ClientOrderId         string  `json:"clientOrderId"`
	TransactTime          uint64  `json:"transactTime"`
	Price                 string  `json:"price"`
	OrigQty               string  `json:"origQty"`
	ExecutedQty           string  `json:"executedQty"`
	CummulativeQuoteQty   string  `json:"cummulativeQuoteQty"`
	Status                string  `json:"status"`
	TimeInForce           string  `json:"timeInForce"`
	Type                  string  `json:"type"`
	Side                  string  `json:"side"`
	MarginBuyBorrowAmount float64 `json:"marginBuyBorrowAmount"`
	MarginBuyBorrowAsset  string  `json:"marginBuyBorrowAsset"`
	IsIsolated            bool    `json:"isIsolated"`
	Fills                 []struct {
		Price           string `json:"price"`
		Qty             string `json:"qty"`
		Commission      string `json:"commission"`
		CommissionAsset string `json:"commissionAsset"`
	} `json:"fills"`
}

// Margin Account Cancel Order (TRADE) API Endpoint
const (
	marginAccountCancelOrderEndpoint = "/sapi/v1/margin/order"
)

// MarginAccountCancelOrderService margin account cancel order
type MarginAccountCancelOrderService struct {
	c                 *Client
	symbol            string
	isIsolated        *string
	orderId           *int
	origClientOrderId *string
	newClientOrderId  *string
}

// Symbol set symbol
func (s *MarginAccountCancelOrderService) Symbol(symbol string) *MarginAccountCancelOrderService {
	s.symbol = symbol
	return s
}

// IsIsolated set isIsolated
func (s *MarginAccountCancelOrderService) IsIsolated(isIsolated string) *MarginAccountCancelOrderService {
	s.isIsolated = &isIsolated
	return s
}

// OrderId set orderId
func (s *MarginAccountCancelOrderService) OrderId(orderId int) *MarginAccountCancelOrderService {
	s.orderId = &orderId
	return s
}

// OrigClientOrderId set origClientOrderId
func (s *MarginAccountCancelOrderService) OrigClientOrderId(origClientOrderId string) *MarginAccountCancelOrderService {
	s.origClientOrderId = &origClientOrderId
	return s
}

// NewClientOrderId set newClientOrderId
func (s *MarginAccountCancelOrderService) NewClientOrderId(newClientOrderId string) *MarginAccountCancelOrderService {
	s.newClientOrderId = &newClientOrderId
	return s
}

// Do send request
func (s *MarginAccountCancelOrderService) Do(ctx context.Context, opts ...RequestOption) (res *MarginAccountCancelOrderResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: marginAccountCancelOrderEndpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol": s.symbol,
	}
	if s.isIsolated != nil {
		m["isIsolated"] = *s.isIsolated
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
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &MarginAccountCancelOrderResponse{}, err
	}
	res = new(MarginAccountCancelOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return &MarginAccountCancelOrderResponse{}, err
	}
	return res, nil
}

// MarginAccountCancelOrderResponse define margin account cancel order response
type MarginAccountCancelOrderResponse struct {
	Symbol              string `json:"symbol"`
	IsIsolated          bool   `json:"isIsolated"`
	OrderId             int    `json:"orderId"`
	OrigClientOrderId   string `json:"origClientOrderId"`
	ClientOrderId       string `json:"clientOrderId"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
}

// Margin Account Cancel All Orders (TRADE) API Endpoint
const (
	marginAccountCancelAllOrdersEndpoint = "/sapi/v1/margin/openOrders"
)

// MarginAccountCancelAllOrdersService margin account cancel all orders
type MarginAccountCancelAllOrdersService struct {
	c          *Client
	symbol     string
	isIsolated *string
}

// Symbol set symbol
func (s *MarginAccountCancelAllOrdersService) Symbol(symbol string) *MarginAccountCancelAllOrdersService {
	s.symbol = symbol
	return s
}

// IsIsolated set isIsolated
func (s *MarginAccountCancelAllOrdersService) IsIsolated(isIsolated string) *MarginAccountCancelAllOrdersService {
	s.isIsolated = &isIsolated
	return s
}

// Do send request
func (s *MarginAccountCancelAllOrdersService) Do(ctx context.Context, opts ...RequestOption) (res *MarginAccountCancelAllOrdersResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: marginAccountCancelAllOrdersEndpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol": s.symbol,
	}
	if s.isIsolated != nil {
		m["isIsolated"] = *s.isIsolated
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &MarginAccountCancelAllOrdersResponse{}, err
	}
	res = new(MarginAccountCancelAllOrdersResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return &MarginAccountCancelAllOrdersResponse{}, err
	}
	return res, nil
}

// MarginAccountCancelAllOrdersResponse define margin account cancel all orders response
type MarginAccountCancelAllOrdersResponse struct {
	Symbol              string `json:"symbol"`
	IsIsolated          bool   `json:"isIsolated"`
	OrigClientOrderId   string `json:"origClientOrderId"`
	OrderId             int    `json:"orderId"`
	OrderListId         int    `json:"orderListId"`
	ClientOrderId       string `json:"clientOrderId"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
}

// Get Cross Margin Transfer History (USER_DATA) API Endpoint
const (
	crossMarginTransferHistoryEndpoint = "/sapi/v1/margin/transfer"
)

// CrossMarginTransferHistoryService get cross margin transfer history
type CrossMarginTransferHistoryService struct {
	c         *Client
	asset     *string
	orderType *string
	startTime *uint64
	endTime   *uint64
	current   *int
	size      *int
	archived  *string
}

// Asset set asset
func (s *CrossMarginTransferHistoryService) Asset(asset string) *CrossMarginTransferHistoryService {
	s.asset = &asset
	return s
}

// OrderType set orderType
func (s *CrossMarginTransferHistoryService) OrderType(orderType string) *CrossMarginTransferHistoryService {
	s.orderType = &orderType
	return s
}

// StartTime set startTime
func (s *CrossMarginTransferHistoryService) StartTime(startTime uint64) *CrossMarginTransferHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *CrossMarginTransferHistoryService) EndTime(endTime uint64) *CrossMarginTransferHistoryService {
	s.endTime = &endTime
	return s
}

// Current set current
func (s *CrossMarginTransferHistoryService) Current(current int) *CrossMarginTransferHistoryService {
	s.current = &current
	return s
}

// Size set size
func (s *CrossMarginTransferHistoryService) Size(size int) *CrossMarginTransferHistoryService {
	s.size = &size
	return s
}

// Archived set archived
func (s *CrossMarginTransferHistoryService) Archived(archived string) *CrossMarginTransferHistoryService {
	s.archived = &archived
	return s
}

// Do send request
func (s *CrossMarginTransferHistoryService) Do(ctx context.Context, opts ...RequestOption) (res *CrossMarginTransferHistoryResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: crossMarginTransferHistoryEndpoint,
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.orderType != nil {
		r.setParam("type", *s.orderType)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	if s.archived != nil {
		r.setParam("archived", *s.archived)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &CrossMarginTransferHistoryResponse{}, err
	}
	res = new(CrossMarginTransferHistoryResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return &CrossMarginTransferHistoryResponse{}, err
	}
	return res, nil
}

// CrossMarginTransferHistoryResponse define cross margin transfer history response
type CrossMarginTransferHistoryResponse struct {
	Rows []struct {
		Amount    string `json:"amount"`
		Asset     string `json:"asset"`
		Status    string `json:"status"`
		Timestamp uint64 `json:"timestamp"`
		TxId      int64  `json:"txId"`
		Type      string `json:"type"`
	} `json:"rows"`
	Total int `json:"total"`
}

// Query Interest History (USER_DATA) API Endpoint
const (
	interestHistoryEndpoint = "/sapi/v1/margin/interestHistory"
)

// InterestHistoryService query interest history
type InterestHistoryService struct {
	c              *Client
	asset          *string
	isolatedSymbol *string
	startTime      *uint64
	endTime        *uint64
	current        *int
	size           *int
	archived       *string
}

// Asset set asset
func (s *InterestHistoryService) Asset(asset string) *InterestHistoryService {
	s.asset = &asset
	return s
}

// IsolatedSymbol set isolatedSymbol
func (s *InterestHistoryService) IsolatedSymbol(isolatedSymbol string) *InterestHistoryService {
	s.isolatedSymbol = &isolatedSymbol
	return s
}

// StartTime set startTime
func (s *InterestHistoryService) StartTime(startTime uint64) *InterestHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *InterestHistoryService) EndTime(endTime uint64) *InterestHistoryService {
	s.endTime = &endTime
	return s
}

// Current set current
func (s *InterestHistoryService) Current(current int) *InterestHistoryService {
	s.current = &current
	return s
}

// Size set size
func (s *InterestHistoryService) Size(size int) *InterestHistoryService {
	s.size = &size
	return s
}

// Archived set archived
func (s *InterestHistoryService) Archived(archived string) *InterestHistoryService {
	s.archived = &archived
	return s
}

// Do send request
func (s *InterestHistoryService) Do(ctx context.Context, opts ...RequestOption) (res *InterestHistoryResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: interestHistoryEndpoint,
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.isolatedSymbol != nil {
		r.setParam("isolatedSymbol", *s.isolatedSymbol)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	if s.archived != nil {
		r.setParam("archived", *s.archived)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &InterestHistoryResponse{}, err
	}
	res = new(InterestHistoryResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return &InterestHistoryResponse{}, err
	}
	return res, nil
}

// InterestHistoryResponse define interest history response
type InterestHistoryResponse struct {
	Rows []struct {
		TxId                int64  `json:"txId"`
		InterestAccuredTime uint64 `json:"interestAccuredTime"`
		Asset               string `json:"asset"`
		RawAsset            string `json:"rawAsset"`
		Principal           string `json:"principal"`
		Interest            string `json:"interest"`
		InterestRate        string `json:"interestRate"`
		Type                string `json:"type"`
		IsolatedSymbol      string `json:"isolatedSymbol"`
	} `json:"rows"`
	Total int `json:"total"`
}

// Query Force Liquidation Record (USER_DATA) API Endpoint
const (
	forceLiquidationRecordEndpoint = "/sapi/v1/margin/forceLiquidationRec"
)

// ForceLiquidationRecordService query force liquidation record
type ForceLiquidationRecordService struct {
	c              *Client
	startTime      *uint64
	endTime        *uint64
	isolatedSymbol *string
	current        *int
	size           *int
}

// IsolatedSymbol set isolatedSymbol
func (s *ForceLiquidationRecordService) IsolatedSymbol(isolatedSymbol string) *ForceLiquidationRecordService {
	s.isolatedSymbol = &isolatedSymbol
	return s
}

// StartTime set startTime
func (s *ForceLiquidationRecordService) StartTime(startTime uint64) *ForceLiquidationRecordService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *ForceLiquidationRecordService) EndTime(endTime uint64) *ForceLiquidationRecordService {
	s.endTime = &endTime
	return s
}

// Current set current
func (s *ForceLiquidationRecordService) Current(current int) *ForceLiquidationRecordService {
	s.current = &current
	return s
}

// Size set size
func (s *ForceLiquidationRecordService) Size(size int) *ForceLiquidationRecordService {
	s.size = &size
	return s
}

// Do send request
func (s *ForceLiquidationRecordService) Do(ctx context.Context, opts ...RequestOption) (res *ForceLiquidationRecordResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: forceLiquidationRecordEndpoint,
		secType:  secTypeSigned,
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.isolatedSymbol != nil {
		r.setParam("isolatedSymbol", *s.isolatedSymbol)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &ForceLiquidationRecordResponse{}, err
	}
	res = new(ForceLiquidationRecordResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return &ForceLiquidationRecordResponse{}, err
	}
	return res, nil
}

// ForceLiquidationRecordResponse define force liquidation record response
type ForceLiquidationRecordResponse struct {
	Rows []struct {
		AvgPrice    string `json:"avgPrice"`
		ExecutedQty string `json:"executedQty"`
		OrderId     int    `json:"orderId"`
		Price       string `json:"price"`
		Qty         string `json:"qty"`
		Side        string `json:"side"`
		Symbol      string `json:"symbol"`
		TimeInForce string `json:"timeInForce"`
		IsIsolated  bool   `json:"isIsolated"`
		UpdatedTime uint64 `json:"updatedTime"`
	} `json:"rows"`
	Total int `json:"total"`
}

// Query Query Cross Margin Account Details (USER_DATA) API Endpoint
const (
	crossMarginAccountDetailEndpoint = "/sapi/v1/margin/account"
)

// CrossMarginAccountDetailService query cross margin account details
type CrossMarginAccountDetailService struct {
	c *Client
}

// Do send request
func (s *CrossMarginAccountDetailService) Do(ctx context.Context, opts ...RequestOption) (res *CrossMarginAccountDetailResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: crossMarginAccountDetailEndpoint,
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &CrossMarginAccountDetailResponse{}, err
	}
	res = new(CrossMarginAccountDetailResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return &CrossMarginAccountDetailResponse{}, err
	}
	return res, nil
}

// CrossMarginAccountDetailResponse define cross margin account detail response
type CrossMarginAccountDetailResponse struct {
	Created                    bool   `json:"created"`
	BorrowEnabled              bool   `json:"borrowEnabled"`
	MarginLevel                string `json:"marginLevel"`
	CollateralMarginLevel      string `json:"collateralMarginLevel"`
	TotalAssetOfBtc            string `json:"totalAssetOfBtc"`
	TotalLiabilityOfBtc        string `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBtc         string `json:"totalNetAssetOfBtc"`
	TotalCollateralValueInUSDT string `json:"totalCollateralValueInUSDT"`
	TradeEnabled               bool   `json:"tradeEnabled"`
	TransferInEnabled          bool   `json:"transferInEnabled"`
	TransferOutEnabled         bool   `json:"transferOutEnabled"`
	AccountType                string `json:"accountType"`
	UserAssets                 []struct {
		Asset    string `json:"asset"`
		Borrowed string `json:"borrowed"`
		Free     string `json:"free"`
		Interest string `json:"interest"`
		Locked   string `json:"locked"`
		NetAsset string `json:"netAsset"`
	} `json:"userAssets"`
}

// Query Margin Account's Order (USER_DATA) API Endpoint
const (
	marginAccountOrderEndpoint = "/sapi/v1/margin/order"
)

// MarginAccountOrderService query margin account's order
type MarginAccountOrderService struct {
	c                 *Client
	symbol            string
	isIsolated        *string
	orderId           *int
	origClientOrderId *string
}

// Symbol set symbol
func (s *MarginAccountOrderService) Symbol(symbol string) *MarginAccountOrderService {
	s.symbol = symbol
	return s
}

// IsIsolated set isIsolated
func (s *MarginAccountOrderService) IsIsolated(isIsolated string) *MarginAccountOrderService {
	s.isIsolated = &isIsolated
	return s
}

// OrderId set orderId
func (s *MarginAccountOrderService) OrderId(orderId int) *MarginAccountOrderService {
	s.orderId = &orderId
	return s
}

// OrigClientOrderId set origClientOrderId
func (s *MarginAccountOrderService) OrigClientOrderId(origClientOrderId string) *MarginAccountOrderService {
	s.origClientOrderId = &origClientOrderId
	return s
}

// Do send request
func (s *MarginAccountOrderService) Do(ctx context.Context, opts ...RequestOption) (res *MarginAccountOrderResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: marginAccountOrderEndpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol": s.symbol,
	}
	if s.isIsolated != nil {
		m["isIsolated"] = *s.isIsolated
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
		return &MarginAccountOrderResponse{}, err
	}
	res = new(MarginAccountOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return &MarginAccountOrderResponse{}, err
	}
	return res, nil
}

// MarginAccountOrderResponse define margin account order response
type MarginAccountOrderResponse struct {
	ClientOrderId       string `json:"clientOrderId"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	ExecutedQty         string `json:"executedQty"`
	IcebergQty          string `json:"icebergQty"`
	IsWorking           bool   `json:"isWorking"`
	OrderId             int    `json:"orderId"`
	OrigQty             string `json:"origQty"`
	Price               string `json:"price"`
	Side                string `json:"side"`
	Status              string `json:"status"`
	StopPrice           string `json:"stopPrice"`
	Symbol              string `json:"symbol"`
	IsIsolated          bool   `json:"isIsolated"`
	Time                uint64 `json:"time"`
	TimeInForce         string `json:"timeInForce"`
	OrderType           string `json:"type"`
	UpdateTime          uint64 `json:"updateTime"`
}

// Query Margin Account's Open Order (USER_DATA) API Endpoint
const (
	marginAccountOpenOrderEndpoint = "/sapi/v1/margin/openOrders"
)

// MarginAccountOpenOrderService query margin account's open order
type MarginAccountOpenOrderService struct {
	c          *Client
	symbol     *string
	isIsolated *string
}

// Symbol set symbol
func (s *MarginAccountOpenOrderService) Symbol(symbol string) *MarginAccountOpenOrderService {
	s.symbol = &symbol
	return s
}

// IsIsolated set isIsolated
func (s *MarginAccountOpenOrderService) IsIsolated(isIsolated string) *MarginAccountOpenOrderService {
	s.isIsolated = &isIsolated
	return s
}

// Do send request
func (s *MarginAccountOpenOrderService) Do(ctx context.Context, opts ...RequestOption) (res []*MarginAccountOpenOrderResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: marginAccountOpenOrderEndpoint,
		secType:  secTypeSigned,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.isIsolated != nil {
		r.setParam("isIsolated", *s.isIsolated)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*MarginAccountOpenOrderResponse{}, err
	}
	res = make([]*MarginAccountOpenOrderResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*MarginAccountOpenOrderResponse{}, err
	}
	return res, nil
}

// MarginAccountOpenOrderResponse define margin account open order response
type MarginAccountOpenOrderResponse struct {
	ClientOrderId       string `json:"clientOrderId"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	ExecutedQty         string `json:"executedQty"`
	IcebergQty          string `json:"icebergQty"`
	IsWorking           bool   `json:"isWorking"`
	OrderId             int    `json:"orderId"`
	OrigQty             string `json:"origQty"`
	Price               string `json:"price"`
	Side                string `json:"side"`
	Status              string `json:"status"`
	StopPrice           string `json:"stopPrice"`
	Symbol              string `json:"symbol"`
	IsIsolated          bool   `json:"isIsolated"`
	Time                uint64 `json:"time"`
	TimeInForce         string `json:"timeInForce"`
	OrderType           string `json:"type"`
	UpdateTime          uint64 `json:"updateTime"`
}

// Query Margin Account's All Orders (USER_DATA) API Endpoint
const (
	marginAccountAllOrderEndpoint = "/sapi/v1/margin/allOrders"
)

// MarginAccountAllOrderService query margin account's all order
type MarginAccountAllOrderService struct {
	c          *Client
	symbol     string
	isIsolated *string
	orderId    *int
	startTime  *uint64
	endTime    *uint64
	limit      *int
}

// Symbol set symbol
func (s *MarginAccountAllOrderService) Symbol(symbol string) *MarginAccountAllOrderService {
	s.symbol = symbol
	return s
}

// IsIsolated set isIsolated
func (s *MarginAccountAllOrderService) IsIsolated(isIsolated string) *MarginAccountAllOrderService {
	s.isIsolated = &isIsolated
	return s
}

// OrderId set orderId
func (s *MarginAccountAllOrderService) OrderId(orderId int) *MarginAccountAllOrderService {
	s.orderId = &orderId
	return s
}

// StartTime set startTime
func (s *MarginAccountAllOrderService) StartTime(startTime uint64) *MarginAccountAllOrderService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *MarginAccountAllOrderService) EndTime(endTime uint64) *MarginAccountAllOrderService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *MarginAccountAllOrderService) Limit(limit int) *MarginAccountAllOrderService {
	s.limit = &limit
	return s
}

// Do send request
func (s *MarginAccountAllOrderService) Do(ctx context.Context, opts ...RequestOption) (res []*MarginAccountAllOrderResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: marginAccountAllOrderEndpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol": s.symbol,
	}
	if s.isIsolated != nil {
		m["isIsolated"] = *s.isIsolated
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
		return []*MarginAccountAllOrderResponse{}, err
	}
	res = make([]*MarginAccountAllOrderResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*MarginAccountAllOrderResponse{}, err
	}
	return res, nil
}

// MarginAccountAllOrderResponse define margin account all order response
type MarginAccountAllOrderResponse struct {
	ClientOrderId       string `json:"clientOrderId"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	ExecutedQty         string `json:"executedQty"`
	IcebergQty          string `json:"icebergQty"`
	IsWorking           bool   `json:"isWorking"`
	OrderId             int    `json:"orderId"`
	OrigQty             string `json:"origQty"`
	Price               string `json:"price"`
	Side                string `json:"side"`
	Status              string `json:"status"`
	StopPrice           string `json:"stopPrice"`
	Symbol              string `json:"symbol"`
	IsIsolated          bool   `json:"isIsolated"`
	Time                uint64 `json:"time"`
	TimeInForce         string `json:"timeInForce"`
	OrderType           string `json:"type"`
	UpdateTime          uint64 `json:"updateTime"`
}

// Margin Account New OCO (TRADE) API Endpoint
const (
	marginAccountNewOCOEndpoint = "/sapi/v1/margin/order/oco"
)

// MarginAccountNewOCOService create new oco order
type MarginAccountNewOCOService struct {
	c                    *Client
	symbol               string
	isIsolated           *string
	listClientOrderId    *string
	side                 string
	quantity             float64
	limitClientOrderId   *string
	price                float64
	limitIcebergQty      *float64
	stopClientOrderId    *string
	stopPrice            float64
	stopLimitPrice       *float64
	stopIcebergQty       *float64
	stopLimitTimeInForce *string
	newOrderRespType     *string
	sideEffectType       *string
}

// Symbol set symbol
func (s *MarginAccountNewOCOService) Symbol(symbol string) *MarginAccountNewOCOService {
	s.symbol = symbol
	return s
}

// IsIsolated set isIsolated
func (s *MarginAccountNewOCOService) IsIsolated(isIsolated string) *MarginAccountNewOCOService {
	s.isIsolated = &isIsolated
	return s
}

// ListClientOrderId set listClientOrderId
func (s *MarginAccountNewOCOService) ListClientOrderId(listClientOrderId string) *MarginAccountNewOCOService {
	s.listClientOrderId = &listClientOrderId
	return s
}

// Side set side
func (s *MarginAccountNewOCOService) Side(side string) *MarginAccountNewOCOService {
	s.side = side
	return s
}

// Quantity set quantity
func (s *MarginAccountNewOCOService) Quantity(quantity float64) *MarginAccountNewOCOService {
	s.quantity = quantity
	return s
}

// LimitClientOrderId set limitClientOrderId
func (s *MarginAccountNewOCOService) LimitClientOrderId(limitClientOrderId string) *MarginAccountNewOCOService {
	s.limitClientOrderId = &limitClientOrderId
	return s
}

// Price set price
func (s *MarginAccountNewOCOService) Price(price float64) *MarginAccountNewOCOService {
	s.price = price
	return s
}

// LimitIcebergQty set limitIcebergQty
func (s *MarginAccountNewOCOService) LimitIcebergQty(limitIcebergQty float64) *MarginAccountNewOCOService {
	s.limitIcebergQty = &limitIcebergQty
	return s
}

// StopClientOrderId set stopClientOrderId
func (s *MarginAccountNewOCOService) StopClientOrderId(stopClientOrderId string) *MarginAccountNewOCOService {
	s.stopClientOrderId = &stopClientOrderId
	return s
}

// StopPrice set stopPrice
func (s *MarginAccountNewOCOService) StopPrice(stopPrice float64) *MarginAccountNewOCOService {
	s.stopPrice = stopPrice
	return s
}

// StopLimitPrice set stopLimitPrice
func (s *MarginAccountNewOCOService) StopLimitPrice(stopLimitPrice float64) *MarginAccountNewOCOService {
	s.stopLimitPrice = &stopLimitPrice
	return s
}

// StopIcebergQty set stopIcebergQty
func (s *MarginAccountNewOCOService) StopIcebergQty(stopIcebergQty float64) *MarginAccountNewOCOService {
	s.stopIcebergQty = &stopIcebergQty
	return s
}

// StopLimitTimeInForce set stopLimitTimeInForce
func (s *MarginAccountNewOCOService) StopLimitTimeInForce(stopLimitTimeInForce string) *MarginAccountNewOCOService {
	s.stopLimitTimeInForce = &stopLimitTimeInForce
	return s
}

// NewOrderRespType set newOrderRespType
func (s *MarginAccountNewOCOService) NewOrderRespType(newOrderRespType string) *MarginAccountNewOCOService {
	s.newOrderRespType = &newOrderRespType
	return s
}

// SideEffectType set sideEffectType
func (s *MarginAccountNewOCOService) SideEffectType(sideEffectType string) *MarginAccountNewOCOService {
	s.sideEffectType = &sideEffectType
	return s
}

// Do send request
func (s *MarginAccountNewOCOService) Do(ctx context.Context, opts ...RequestOption) (res *MarginAccountNewOCOResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: marginAccountNewOCOEndpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol":    s.symbol,
		"side":      s.side,
		"quantity":  s.quantity,
		"price":     s.price,
		"stopPrice": s.stopPrice,
	}
	if s.isIsolated != nil {
		m["isIsolated"] = *s.isIsolated
	}
	if s.listClientOrderId != nil {
		m["listClientOrderId"] = *s.listClientOrderId
	}
	if s.limitClientOrderId != nil {
		m["limitClientOrderId"] = *s.limitClientOrderId
	}
	if s.limitIcebergQty != nil {
		m["limitIcebergQty"] = *s.limitIcebergQty
	}
	if s.stopClientOrderId != nil {
		m["stopClientOrderId"] = *s.stopClientOrderId
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
	if s.sideEffectType != nil {
		m["sideEffectType"] = *s.sideEffectType
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &MarginAccountNewOCOResponse{}, err
	}
	res = new(MarginAccountNewOCOResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return &MarginAccountNewOCOResponse{}, err
	}
	return res, nil
}

// MarginAccountNewOCOService response
type MarginAccountNewOCOResponse struct {
	OrderListId           int    `json:"orderListId"`
	ContingencyType       string `json:"contingencyType"`
	ListStatusType        string `json:"listStatusType"`
	ListOrderStatus       string `json:"listOrderStatus"`
	ListClientOrderId     string `json:"listClientOrderId"`
	TransactionTime       uint64 `json:"transactionTime"`
	Symbol                string `json:"symbol"`
	MarginBuyBorrowAmount string `json:"marginBuyBorrowAmount"`
	MarginBuyBorrowAsset  string `json:"marginBuyBorrowAsset"`
	Orders                []struct {
		Symbol        string `json:"symbol"`
		OrderId       int    `json:"orderId"`
		ClientOrderId string `json:"clientOrderId"`
	} `json:"orders"`
	OrderReports []struct {
		Symbol              string `json:"symbol"`
		OrderId             int    `json:"orderId"`
		OrderListId         int    `json:"orderListId"`
		ClientOrderId       string `json:"clientOrderId"`
		TransactTime        uint64 `json:"transactTime"`
		Price               string `json:"price"`
		OrigQty             string `json:"origQty"`
		ExecutedQty         string `json:"executedQty"`
		CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
		Status              string `json:"status"`
		TimeInForce         string `json:"timeInForce"`
		OrderType           string `json:"type"`
		Side                string `json:"side"`
		StopPrice           string `json:"stopPrice"`
	} `json:"orderReports"`
}

// Margin Account Cancel OCO (TRADE)
const (
	marginAccountCancelOCOEndpoint = "/sapi/v1/margin/orderList"
)

type MarginAccountCancelOCOService struct {
	c                 *Client
	symbol            string
	isIsolated        *string
	orderListId       *int
	listClientOrderId *string
	newClientOrderId  *string
}

// Symbol set symbol
func (s *MarginAccountCancelOCOService) Symbol(symbol string) *MarginAccountCancelOCOService {
	s.symbol = symbol
	return s
}

// IsIsolated set isIsolated
func (s *MarginAccountCancelOCOService) IsIsolated(isIsolated string) *MarginAccountCancelOCOService {
	s.isIsolated = &isIsolated
	return s
}

// OrderListId set orderListId
func (s *MarginAccountCancelOCOService) OrderListId(orderListId int) *MarginAccountCancelOCOService {
	s.orderListId = &orderListId
	return s
}

// ListClientOrderId set listClientOrderId
func (s *MarginAccountCancelOCOService) ListClientOrderId(listClientOrderId string) *MarginAccountCancelOCOService {
	s.listClientOrderId = &listClientOrderId
	return s
}

// NewClientOrderId set newClientOrderId
func (s *MarginAccountCancelOCOService) NewClientOrderId(newClientOrderId string) *MarginAccountCancelOCOService {
	s.newClientOrderId = &newClientOrderId
	return s
}

// Do send request
func (s *MarginAccountCancelOCOService) Do(ctx context.Context, opts ...RequestOption) (res *MarginAccountCancelOCOResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: marginAccountCancelOCOEndpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol": s.symbol,
	}
	if s.isIsolated != nil {
		m["isIsolated"] = *s.isIsolated
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
		return &MarginAccountCancelOCOResponse{}, err
	}
	res = new(MarginAccountCancelOCOResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return &MarginAccountCancelOCOResponse{}, err
	}
	return res, nil
}

// MarginAccountCancelOCOService response
type MarginAccountCancelOCOResponse struct {
	OrderListId       int    `json:"orderListId"`
	ContingencyType   string `json:"contingencyType"`
	ListStatusType    string `json:"listStatusType"`
	ListOrderStatus   string `json:"listOrderStatus"`
	ListClientOrderId string `json:"listClientOrderId"`
	TransactionTime   uint64 `json:"transactionTime"`
	Symbol            string `json:"symbol"`
	IsIsolated        bool   `json:"isIsolated"`
	Orders            []struct {
		Symbol        string `json:"symbol"`
		OrderId       int    `json:"orderId"`
		ClientOrderId string `json:"clientOrderId"`
	} `json:"orders"`
	OrderReports []struct {
		Symbol              string `json:"symbol"`
		OrigClientOrderId   string `json:"origClientOrderId"`
		OrderId             int    `json:"orderId"`
		OrderListId         int    `json:"orderListId"`
		ClientOrderId       string `json:"clientOrderId"`
		Price               string `json:"price"`
		OrigQty             string `json:"origQty"`
		ExecutedQty         string `json:"executedQty"`
		CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
		Status              string `json:"status"`
		TimeInForce         string `json:"timeInForce"`
		OrderType           string `json:"type"`
		Side                string `json:"side"`
		StopPrice           string `json:"stopPrice"`
	} `json:"orderReports"`
}

// Query Margin Account's OCO (USER_DATA) (HMAC SHA256)
const (
	marginAccountQueryOCOEndpoint = "/sapi/v1/margin/orderList"
)

type MarginAccountQueryOCOService struct {
	c                 *Client
	isIsolated        *string
	symbol            *string
	orderListId       *int
	origClientOrderId *string
}

// IsIsolated set isIsolated
func (s *MarginAccountQueryOCOService) IsIsolated(isIsolated string) *MarginAccountQueryOCOService {
	s.isIsolated = &isIsolated
	return s
}

// Symbol set symbol
func (s *MarginAccountQueryOCOService) Symbol(symbol string) *MarginAccountQueryOCOService {
	s.symbol = &symbol
	return s
}

// OrderListId set orderListId
func (s *MarginAccountQueryOCOService) OrderListId(orderListId int) *MarginAccountQueryOCOService {
	s.orderListId = &orderListId
	return s
}

// OrigClientOrderId set origClientOrderId
func (s *MarginAccountQueryOCOService) OrigClientOrderId(origClientOrderId string) *MarginAccountQueryOCOService {
	s.origClientOrderId = &origClientOrderId
	return s
}

// Do send request
func (s *MarginAccountQueryOCOService) Do(ctx context.Context, opts ...RequestOption) (res *MarginAccountQueryOCOResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: marginAccountQueryOCOEndpoint,
		secType:  secTypeSigned,
	}
	if s.isIsolated != nil {
		r.setParam("isIsolated", *s.isIsolated)
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.orderListId != nil {
		r.setParam("orderListId", *s.orderListId)
	}
	if s.origClientOrderId != nil {
		r.setParam("origClientOrderId", *s.origClientOrderId)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &MarginAccountQueryOCOResponse{}, err
	}
	res = new(MarginAccountQueryOCOResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return &MarginAccountQueryOCOResponse{}, err
	}
	return res, nil
}

// MarginAccountQueryOCOService response
type MarginAccountQueryOCOResponse struct {
	OrderListId       int    `json:"orderListId"`
	ContingencyType   string `json:"contingencyType"`
	ListStatusType    string `json:"listStatusType"`
	ListOrderStatus   string `json:"listOrderStatus"`
	ListClientOrderId string `json:"listClientOrderId"`
	TransactionTime   uint64 `json:"transactionTime"`
	Symbol            string `json:"symbol"`
	IsIsolated        bool   `json:"isIsolated"`
	Orders            []struct {
		Symbol        string `json:"symbol"`
		OrderId       int    `json:"orderId"`
		ClientOrderId string `json:"clientOrderId"`
	} `json:"orders"`
}

// Query Margin Account's all OCO (USER_DATA)
const (
	marginAccountQueryAllOCOEndpoint = "/sapi/v1/margin/allOrderList"
)

type MarginAccountQueryAllOCOService struct {
	c          *Client
	isIsolated *string
	symbol     *string
	fromId     *int
	startTime  *uint64
	endTime    *uint64
	limit      *int
}

// IsIsolated set isIsolated
func (s *MarginAccountQueryAllOCOService) IsIsolated(isIsolated string) *MarginAccountQueryAllOCOService {
	s.isIsolated = &isIsolated
	return s
}

// Symbol set symbol
func (s *MarginAccountQueryAllOCOService) Symbol(symbol string) *MarginAccountQueryAllOCOService {
	s.symbol = &symbol
	return s
}

// FromId set fromId
func (s *MarginAccountQueryAllOCOService) FromId(fromId int) *MarginAccountQueryAllOCOService {
	s.fromId = &fromId
	return s
}

// StartTime set startTime
func (s *MarginAccountQueryAllOCOService) StartTime(startTime uint64) *MarginAccountQueryAllOCOService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *MarginAccountQueryAllOCOService) EndTime(endTime uint64) *MarginAccountQueryAllOCOService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *MarginAccountQueryAllOCOService) Limit(limit int) *MarginAccountQueryAllOCOService {
	s.limit = &limit
	return s
}

// Do send request
func (s *MarginAccountQueryAllOCOService) Do(ctx context.Context, opts ...RequestOption) (res []*MarginAccountQueryAllOCOResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: marginAccountQueryAllOCOEndpoint,
		secType:  secTypeSigned,
	}
	if s.isIsolated != nil {
		r.setParam("isIsolated", *s.isIsolated)
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
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
		return []*MarginAccountQueryAllOCOResponse{}, err
	}
	res = make([]*MarginAccountQueryAllOCOResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*MarginAccountQueryAllOCOResponse{}, err
	}
	return res, nil
}

// MarginAccountQueryAllOCOService response
type MarginAccountQueryAllOCOResponse struct {
	OrderListId       int    `json:"orderListId"`
	ContingencyType   string `json:"contingencyType"`
	ListStatusType    string `json:"listStatusType"`
	ListOrderStatus   string `json:"listOrderStatus"`
	ListClientOrderId string `json:"listClientOrderId"`
	TransactionTime   uint64 `json:"transactionTime"`
	Symbol            string `json:"symbol"`
	IsIsolated        bool   `json:"isIsolated"`
	Orders            []struct {
		Symbol        string `json:"symbol"`
		OrderId       int    `json:"orderId"`
		ClientOrderId string `json:"clientOrderId"`
	} `json:"orders"`
}

// Query Margin Account's Open OCO (USER_DATA)
const (
	marginAccountQueryOpenOCOEndpoint = "/sapi/v1/margin/openOrderList"
)

type MarginAccountQueryOpenOCOService struct {
	c          *Client
	isIsolated *string
	symbol     *string
}

// IsIsolated set isIsolated
func (s *MarginAccountQueryOpenOCOService) IsIsolated(isIsolated string) *MarginAccountQueryOpenOCOService {
	s.isIsolated = &isIsolated
	return s
}

// Symbol set symbol
func (s *MarginAccountQueryOpenOCOService) Symbol(symbol string) *MarginAccountQueryOpenOCOService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *MarginAccountQueryOpenOCOService) Do(ctx context.Context, opts ...RequestOption) (res []*MarginAccountQueryOpenOCOResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: marginAccountQueryOpenOCOEndpoint,
		secType:  secTypeSigned,
	}
	if s.isIsolated != nil {
		r.setParam("isIsolated", *s.isIsolated)
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*MarginAccountQueryOpenOCOResponse{}, err
	}
	res = make([]*MarginAccountQueryOpenOCOResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*MarginAccountQueryOpenOCOResponse{}, err
	}
	return res, nil
}

// MarginAccountQueryOpenOCOService response
type MarginAccountQueryOpenOCOResponse struct {
	OrderListId       int    `json:"orderListId"`
	ContingencyType   string `json:"contingencyType"`
	ListStatusType    string `json:"listStatusType"`
	ListOrderStatus   string `json:"listOrderStatus"`
	ListClientOrderId string `json:"listClientOrderId"`
	TransactionTime   uint64 `json:"transactionTime"`
	Symbol            string `json:"symbol"`
	IsIsolated        bool   `json:"isIsolated"`
	Orders            []struct {
		Symbol        string `json:"symbol"`
		OrderId       int    `json:"orderId"`
		ClientOrderId string `json:"clientOrderId"`
	} `json:"orders"`
}

// Query Margin Account's Trade List (USER_DATA)
const (
	marginAccountQueryTradeListEndpoint = "/sapi/v1/margin/myTrades"
)

type MarginAccountQueryTradeListService struct {
	c          *Client
	symbol     string
	isIsolated *string
	orderId    *int
	startTime  *uint64
	endTime    *uint64
	fromId     *int
	limit      *int
}

// Symbol set symbol
func (s *MarginAccountQueryTradeListService) Symbol(symbol string) *MarginAccountQueryTradeListService {
	s.symbol = symbol
	return s
}

// IsIsolated set isIsolated
func (s *MarginAccountQueryTradeListService) IsIsolated(isIsolated string) *MarginAccountQueryTradeListService {
	s.isIsolated = &isIsolated
	return s
}

// OrderId set orderId
func (s *MarginAccountQueryTradeListService) OrderId(orderId int) *MarginAccountQueryTradeListService {
	s.orderId = &orderId
	return s
}

// StartTime set startTime
func (s *MarginAccountQueryTradeListService) StartTime(startTime uint64) *MarginAccountQueryTradeListService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *MarginAccountQueryTradeListService) EndTime(endTime uint64) *MarginAccountQueryTradeListService {
	s.endTime = &endTime
	return s
}

// FromId set fromId
func (s *MarginAccountQueryTradeListService) FromId(fromId int) *MarginAccountQueryTradeListService {
	s.fromId = &fromId
	return s
}

// Limit set limit
func (s *MarginAccountQueryTradeListService) Limit(limit int) *MarginAccountQueryTradeListService {
	s.limit = &limit
	return s
}

// Do send request
func (s *MarginAccountQueryTradeListService) Do(ctx context.Context, opts ...RequestOption) (res []*MarginAccountQueryTradeListResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: marginAccountQueryTradeListEndpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol": s.symbol,
	}
	if s.isIsolated != nil {
		m["isIsolated"] = *s.isIsolated
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
		return []*MarginAccountQueryTradeListResponse{}, err
	}
	res = make([]*MarginAccountQueryTradeListResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*MarginAccountQueryTradeListResponse{}, err
	}
	return res, nil
}

// MarginAccountQueryTradeListService response
type MarginAccountQueryTradeListResponse struct {
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Id              int    `json:"id"`
	IsBestMatch     bool   `json:"isBestMatch"`
	IsBuyer         bool   `json:"isBuyer"`
	IsMaker         bool   `json:"isMaker"`
	OrderId         int    `json:"orderId"`
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	Symbol          string `json:"symbol"`
	IsIsolated      bool   `json:"isIsolated"`
	Time            uint64 `json:"time"`
}

// Query Margin Account's Max Borrow (USER_DATA)
const (
	marginAccountQueryMaxBorrowEndpoint = "/sapi/v1/margin/maxBorrowable"
)

type MarginAccountQueryMaxBorrowService struct {
	c              *Client
	asset          string
	isolatedSymbol *string
}

// Asset set asset
func (s *MarginAccountQueryMaxBorrowService) Asset(asset string) *MarginAccountQueryMaxBorrowService {
	s.asset = asset
	return s
}

// IsolatedSymbol set isolatedSymbol
func (s *MarginAccountQueryMaxBorrowService) IsolatedSymbol(isolatedSymbol string) *MarginAccountQueryMaxBorrowService {
	s.isolatedSymbol = &isolatedSymbol
	return s
}

// Do send request
func (s *MarginAccountQueryMaxBorrowService) Do(ctx context.Context, opts ...RequestOption) (res *MarginAccountQueryMaxBorrowResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: marginAccountQueryMaxBorrowEndpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"asset": s.asset,
	}
	if s.isolatedSymbol != nil {
		m["isolatedSymbol"] = *s.isolatedSymbol
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &MarginAccountQueryMaxBorrowResponse{}, err
	}
	res = new(MarginAccountQueryMaxBorrowResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return &MarginAccountQueryMaxBorrowResponse{}, err
	}
	return res, nil
}

// MarginAccountQueryMaxBorrowService response
type MarginAccountQueryMaxBorrowResponse struct {
	Amount      string `json:"amount"`
	BorrowLimit string `json:"borrowLimit"`
}

// Query Margin Account's Max Transfer-Out Amount (USER_DATA)
const (
	marginAccountQueryMaxTransferOutAmountEndpoint = "/sapi/v1/margin/maxTransferable"
)

type MarginAccountQueryMaxTransferOutAmountService struct {
	c              *Client
	asset          string
	isolatedSymbol *string
}

// Asset set asset
func (s *MarginAccountQueryMaxTransferOutAmountService) Asset(asset string) *MarginAccountQueryMaxTransferOutAmountService {
	s.asset = asset
	return s
}

// IsolatedSymbol set isolatedSymbol
func (s *MarginAccountQueryMaxTransferOutAmountService) IsolatedSymbol(isolatedSymbol string) *MarginAccountQueryMaxTransferOutAmountService {
	s.isolatedSymbol = &isolatedSymbol
	return s
}

// Do send request
func (s *MarginAccountQueryMaxTransferOutAmountService) Do(ctx context.Context, opts ...RequestOption) (res *MarginAccountQueryMaxTransferOutAmountResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: marginAccountQueryMaxTransferOutAmountEndpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"asset": s.asset,
	}
	if s.isolatedSymbol != nil {
		m["isolatedSymbol"] = *s.isolatedSymbol
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &MarginAccountQueryMaxTransferOutAmountResponse{}, err
	}
	res = new(MarginAccountQueryMaxTransferOutAmountResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return &MarginAccountQueryMaxTransferOutAmountResponse{}, err
	}
	return res, nil
}

// MarginAccountQueryMaxTransferOutAmountService response
type MarginAccountQueryMaxTransferOutAmountResponse struct {
	Amount string `json:"amount"`
}

// Adjust cross margin max leverage (USER_DATA)
const (
	marginAccountAdjustCrossMaxLeverageEndpoint = "/sapi/v1/margin/max-leverage"
)

type MarginAccountAdjustCrossMaxLeverageService struct {
	c           *Client
	maxLeverage int32
}

// MaxLeverage set maxLeverage
func (s *MarginAccountAdjustCrossMaxLeverageService) MaxLeverage(maxLeverage int32) *MarginAccountAdjustCrossMaxLeverageService {
	s.maxLeverage = maxLeverage
	return s
}

// Do send request
func (s *MarginAccountAdjustCrossMaxLeverageService) Do(ctx context.Context, opts ...RequestOption) (res *MarginAccountAdjustCrossMaxLeverageResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: marginAccountAdjustCrossMaxLeverageEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("maxLeverage", s.maxLeverage)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &MarginAccountAdjustCrossMaxLeverageResponse{}, err
	}
	res = new(MarginAccountAdjustCrossMaxLeverageResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return &MarginAccountAdjustCrossMaxLeverageResponse{}, err
	}
	return res, nil
}

// MarginAccountAdjustCrossMaxLeverageService response
type MarginAccountAdjustCrossMaxLeverageResponse struct {
	Success bool `json:"success"`
}

// Get Summary of Margin account (USER_DATA) - GET /sapi/v1/margin/tradeCoeff (HMAC SHA256)
const (
	marginAccountSummaryEndpoint = "/sapi/v1/margin/tradeCoeff"
)

type MarginAccountSummaryService struct {
	c *Client
}

// Do send request
func (s *MarginAccountSummaryService) Do(ctx context.Context, opts ...RequestOption) (res *MarginAccountSummaryResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: marginAccountSummaryEndpoint,
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &MarginAccountSummaryResponse{}, err
	}
	res = new(MarginAccountSummaryResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return &MarginAccountSummaryResponse{}, err
	}
	return res, nil
}

// MarginAccountSummaryService response
type MarginAccountSummaryResponse struct {
	NormalBar           string `json:"normalBar"`
	MarginCallBar       string `json:"marginCallBar"`
	ForceLiquidationBar string `json:"forceLiquidationBar"`
}

// Query Isolated Margin Account Info (USER_DATA)
const (
	marginIsolatedAccountInfoEndpoint = "/sapi/v1/margin/isolated/account"
)

type MarginIsolatedAccountInfoService struct {
	c       *Client
	symbols *string
}

// Symbols set symbols
func (s *MarginIsolatedAccountInfoService) Symbols(symbols string) *MarginIsolatedAccountInfoService {
	s.symbols = &symbols
	return s
}

// Do send request
func (s *MarginIsolatedAccountInfoService) Do(ctx context.Context, opts ...RequestOption) (res interface{}, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: marginIsolatedAccountInfoEndpoint,
		secType:  secTypeSigned,
	}
	if s.symbols != nil {
		r.addParam("symbols", s.symbols)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		if s.symbols != nil {
			return &MarginIsolatedAccountInfoResponseSymbols{}, err
		} else {
			return &MarginIsolatedAccountInfoResponse{}, err
		}
	}
	if s.symbols != nil {
		res = new(MarginIsolatedAccountInfoResponseSymbols)
	} else {
		res = new(MarginIsolatedAccountInfoResponse)
	}
	err = json.Unmarshal(data, res)
	if err != nil {
		if s.symbols != nil {
			return &MarginIsolatedAccountInfoResponseSymbols{}, err
		} else {
			return &MarginIsolatedAccountInfoResponse{}, err
		}
	}
	return res, nil
}

type MarginIsolatedAccountInfoAssets struct {
	BaseAsset struct {
		Asset         string `json:"asset"`
		BorrowEnabled bool   `json:"borrowEnabled"`
		Free          string `json:"free"`
		Interest      string `json:"interest"`
		Locked        string `json:"locked"`
		NetAsset      string `json:"netAsset"`
		NetAssetOfBtc string `json:"netAssetOfBtc"`
		RepayEnabled  bool   `json:"repayEnabled"`
		TotalAsset    string `json:"totalAsset"`
	} `json:"baseAsset"`
	QuoteAsset struct {
		Asset         string `json:"asset"`
		BorrowEnabled bool   `json:"borrowEnabled"`
		Free          string `json:"free"`
		Interest      string `json:"interest"`
		Locked        string `json:"locked"`
		NetAsset      string `json:"netAsset"`
		NetAssetOfBtc string `json:"netAssetOfBtc"`
		RepayEnabled  bool   `json:"repayEnabled"`
		TotalAsset    string `json:"totalAsset"`
	} `json:"quoteAsset"`
	Symbol            string `json:"symbol"`
	IsolatedCreated   bool   `json:"isolatedCreated"`
	Enabled           bool   `json:"enabled"`
	MarginLevel       string `json:"marginLevel"`
	MarginLevelStatus string `json:"marginLevelStatus"`
	MarginRatio       string `json:"marginRatio"`
	IndexPrice        string `json:"indexPrice"`
	LiquidatePrice    string `json:"liquidatePrice"`
	LiquidateRate     string `json:"liquidateRate"`
	TradeEnabled      bool   `json:"tradeEnabled"`
}

// MarginIsolatedAccountInfoService response if symbols parameter is not sent
type MarginIsolatedAccountInfoResponseSymbols struct {
	Assets []*MarginIsolatedAccountInfoAssets `json:"assets"`
}

// MarginIsolatedAccountInfoService response if symbols parameter is sent
type MarginIsolatedAccountInfoResponse struct {
	Assets              []*MarginIsolatedAccountInfoAssets `json:"assets"`
	TotalAssetOfBtc     string                             `json:"totalAssetOfBtc"`
	TotalLiabilityOfBtc string                             `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBtc  string                             `json:"totalNetAssetOfBtc"`
}

// Disable Isolated Margin Account (TRADE)
const (
	marginIsolatedAccountDisableEndpoint = "/sapi/v1/margin/isolated/account"
)

type MarginIsolatedAccountDisableService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *MarginIsolatedAccountDisableService) Symbol(symbol string) *MarginIsolatedAccountDisableService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *MarginIsolatedAccountDisableService) Do(ctx context.Context, opts ...RequestOption) (res *MarginIsolatedAccountDisableResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: marginIsolatedAccountDisableEndpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol": s.symbol,
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &MarginIsolatedAccountDisableResponse{}, err
	}
	res = new(MarginIsolatedAccountDisableResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return &MarginIsolatedAccountDisableResponse{}, err
	}
	return res, nil
}

// MarginIsolatedAccountDisableService response
type MarginIsolatedAccountDisableResponse struct {
	Success bool `json:"success"`
}

// Enable Isolated Margin Account (TRADE)
const (
	marginIsolatedAccountEnableEndpoint = "/sapi/v1/margin/isolated/account"
)

type MarginIsolatedAccountEnableService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *MarginIsolatedAccountEnableService) Symbol(symbol string) *MarginIsolatedAccountEnableService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *MarginIsolatedAccountEnableService) Do(ctx context.Context, opts ...RequestOption) (res *MarginIsolatedAccountEnableResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: marginIsolatedAccountEnableEndpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol": s.symbol,
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &MarginIsolatedAccountEnableResponse{}, err
	}
	res = new(MarginIsolatedAccountEnableResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return &MarginIsolatedAccountEnableResponse{}, err
	}
	return res, nil
}

// MarginIsolatedAccountEnableService response
type MarginIsolatedAccountEnableResponse struct {
	Success bool `json:"success"`
}

// Query Enabled Isolated Margin Account Limit (USER_DATA)
const (
	marginIsolatedAccountLimitEndpoint = "/sapi/v1/margin/isolated/accountLimit"
)

type MarginIsolatedAccountLimitService struct {
	c *Client
}

// Do send request
func (s *MarginIsolatedAccountLimitService) Do(ctx context.Context, opts ...RequestOption) (res *MarginIsolatedAccountLimitResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: marginIsolatedAccountLimitEndpoint,
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &MarginIsolatedAccountLimitResponse{}, err
	}
	res = new(MarginIsolatedAccountLimitResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return &MarginIsolatedAccountLimitResponse{}, err
	}
	return res, nil
}

// MarginIsolatedAccountLimitService response
type MarginIsolatedAccountLimitResponse struct {
	EnabledAccount int `json:"enabledAccount"`
	MaxAccount     int `json:"maxAccount"`
}

// Get All Isolated Margin Symbol(USER_DATA)
const (
	marginIsolatedSymbolAllEndpoint = "/sapi/v1/margin/isolated/allPairs"
)

// AllIsolatedMarginSymbolService returns all isolated margin symbols
type AllIsolatedMarginSymbolService struct {
	c *Client
}

// Do send request
func (s *AllIsolatedMarginSymbolService) Do(ctx context.Context, opts ...RequestOption) (res []*MarginIsolatedSymbolResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: marginIsolatedSymbolAllEndpoint,
		secType:  secTypeAPIKey,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*MarginIsolatedSymbolResponse{}, err
	}
	res = make([]*MarginIsolatedSymbolResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*MarginIsolatedSymbolResponse{}, err
	}
	return res, nil
}

// MarginIsolatedSymbolAllService returns all isolated margin symbols
type MarginIsolatedSymbolResponse struct {
	Symbol        string `json:"symbol"`
	Base          string `json:"base"`
	Quote         string `json:"quote"`
	IsMarginTrade bool   `json:"isMarginTrade"`
	IsBuyAllowed  bool   `json:"isBuyAllowed"`
	IsSellAllowed bool   `json:"isSellAllowed"`
}

// Toggle BNB Burn On Spot Trade And Margin Interest (USER_DATA)
const (
	marginToggleBnbBurnEndpoint = "/sapi/v1/bnbBurn"
)

type MarginToggleBnbBurnService struct {
	c               *Client
	spotBNBBurn     *string
	interestBNBBurn *string
}

// SpotBNBBurn set spotBNBBurn
func (s *MarginToggleBnbBurnService) SpotBNBBurn(spotBNBBurn string) *MarginToggleBnbBurnService {
	s.spotBNBBurn = &spotBNBBurn
	return s
}

// InterestBNBBurn set interestBNBBurn
func (s *MarginToggleBnbBurnService) InterestBNBBurn(interestBNBBurn string) *MarginToggleBnbBurnService {
	s.interestBNBBurn = &interestBNBBurn
	return s
}

// Do send request
func (s *MarginToggleBnbBurnService) Do(ctx context.Context, opts ...RequestOption) (res *MarginToggleBnbBurnResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: marginToggleBnbBurnEndpoint,
		secType:  secTypeSigned,
	}
	if s.spotBNBBurn != nil {
		r.addParam("spotBNBBurn", s.spotBNBBurn)
	}
	if s.interestBNBBurn != nil {
		r.addParam("interestBNBBurn", s.interestBNBBurn)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &MarginToggleBnbBurnResponse{}, err
	}
	res = new(MarginToggleBnbBurnResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return &MarginToggleBnbBurnResponse{}, err
	}
	return res, nil
}

// MarginToggleBnbBurnService response
type MarginToggleBnbBurnResponse struct {
	SpotBNBBurn     bool `json:"spotBNBBurn"`
	InterestBNBBurn bool `json:"interestBNBBurn"`
}

// Query Cross Isolated Margin Capital Flow (USER_DATA)
const (
	marginIsolatedCapitalFlowEndpoint = "/sapi/v1/margin/capital-flow"
)

type MarginIsolatedCapitalFlowService struct {
	c            *Client
	asset        *string
	symbol       *string
	isolatedType *string
	startTime    *uint64
	endTime      *uint64
	fromId       *int
	limit        *int
}

// Asset set asset
func (s *MarginIsolatedCapitalFlowService) Asset(asset string) *MarginIsolatedCapitalFlowService {
	s.asset = &asset
	return s
}

// Symbol set symbol
func (s *MarginIsolatedCapitalFlowService) Symbol(symbol string) *MarginIsolatedCapitalFlowService {
	s.symbol = &symbol
	return s
}

// IsolatedType set isolatedType
func (s *MarginIsolatedCapitalFlowService) IsolatedType(isolatedType string) *MarginIsolatedCapitalFlowService {
	s.isolatedType = &isolatedType
	return s
}

// StartTime set startTime
func (s *MarginIsolatedCapitalFlowService) StartTime(startTime uint64) *MarginIsolatedCapitalFlowService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *MarginIsolatedCapitalFlowService) EndTime(endTime uint64) *MarginIsolatedCapitalFlowService {
	s.endTime = &endTime
	return s
}

// FromId set fromId
func (s *MarginIsolatedCapitalFlowService) FromId(fromId int) *MarginIsolatedCapitalFlowService {
	s.fromId = &fromId
	return s
}

// Limit set limit
func (s *MarginIsolatedCapitalFlowService) Limit(limit int) *MarginIsolatedCapitalFlowService {
	s.limit = &limit
	return s
}

// Do send request
func (s *MarginIsolatedCapitalFlowService) Do(ctx context.Context, opts ...RequestOption) (res []*MarginIsolatedCapitalFlowResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: marginIsolatedCapitalFlowEndpoint,
		secType:  secTypeSigned,
	}
	m := params{}
	if s.asset != nil {
		m["asset"] = *s.asset
	}
	if s.symbol != nil {
		m["symbol"] = *s.symbol
	}
	if s.isolatedType != nil {
		m["isolatedType"] = *s.isolatedType
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
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*MarginIsolatedCapitalFlowResponse{}, err
	}
	res = make([]*MarginIsolatedCapitalFlowResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*MarginIsolatedCapitalFlowResponse{}, err
	}
	return res, nil
}

// MarginIsolatedCapitalFlowService response
type MarginIsolatedCapitalFlowResponse struct {
	Id        int    `json:"id"`
	TranId    int    `json:"tranId"`
	Timestamp uint64 `json:"timestamp"`
	Asset     string `json:"asset"`
	Symbol    string `json:"symbol"`
	Type      string `json:"type"`
	Amount    string `json:"amount"`
}

// Get BNB Burn Status (USER_DATA)
const (
	marginBnbBurnStatusEndpoint = "/sapi/v1/bnbBurn"
)

type MarginBnbBurnStatusService struct {
	c *Client
}

// Do send request
func (s *MarginBnbBurnStatusService) Do(ctx context.Context, opts ...RequestOption) (res *MarginBnbBurnStatusResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: marginBnbBurnStatusEndpoint,
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &MarginBnbBurnStatusResponse{}, err
	}
	res = new(MarginBnbBurnStatusResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return &MarginBnbBurnStatusResponse{}, err
	}
	return res, nil
}

// MarginBnbBurnStatusService response
type MarginBnbBurnStatusResponse struct {
	SpotBNBBurn     bool `json:"spotBNBBurn"`
	InterestBNBBurn bool `json:"interestBNBBurn"`
}

// Query Margin Interest Rate History (USER_DATA)
const (
	marginInterestRateHistoryEndpoint = "/sapi/v1/margin/interestRateHistory"
)

type MarginInterestRateHistoryService struct {
	c         *Client
	asset     string
	vipLevel  *int
	startTime *uint64
	endTime   *uint64
}

// Asset set asset
func (s *MarginInterestRateHistoryService) Asset(asset string) *MarginInterestRateHistoryService {
	s.asset = asset
	return s
}

// VipLevel set vipLevel
func (s *MarginInterestRateHistoryService) VipLevel(vipLevel int) *MarginInterestRateHistoryService {
	s.vipLevel = &vipLevel
	return s
}

// StartTime set startTime
func (s *MarginInterestRateHistoryService) StartTime(startTime uint64) *MarginInterestRateHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *MarginInterestRateHistoryService) EndTime(endTime uint64) *MarginInterestRateHistoryService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *MarginInterestRateHistoryService) Do(ctx context.Context, opts ...RequestOption) (res []*MarginInterestRateHistoryResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: marginInterestRateHistoryEndpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"asset": s.asset,
	}
	if s.vipLevel != nil {
		m["vipLevel"] = *s.vipLevel
	}
	if s.startTime != nil {
		m["startTime"] = *s.startTime
	}
	if s.endTime != nil {
		m["endTime"] = *s.endTime
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*MarginInterestRateHistoryResponse{}, err
	}
	res = make([]*MarginInterestRateHistoryResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*MarginInterestRateHistoryResponse{}, err
	}
	return res, nil
}

// MarginInterestRateHistoryService response
type MarginInterestRateHistoryResponse struct {
	Asset             string `json:"asset"`
	DailyInterestRate string `json:"dailyInterestRate"`
	Timestamp         uint64 `json:"timestamp"`
	VIPLevel          int    `json:"vipLevel"`
}

// Query Cross Margin Fee Data (USER_DATA)
const (
	marginCrossMarginFeeEndpoint = "/sapi/v1/margin/crossMarginData"
)

type MarginCrossMarginFeeService struct {
	c        *Client
	vipLevel *int
	coin     *string
}

// VipLevel set vipLevel
func (s *MarginCrossMarginFeeService) VipLevel(vipLevel int) *MarginCrossMarginFeeService {
	s.vipLevel = &vipLevel
	return s
}

// Coin set coin
func (s *MarginCrossMarginFeeService) Coin(coin string) *MarginCrossMarginFeeService {
	s.coin = &coin
	return s
}

// Do send request
func (s *MarginCrossMarginFeeService) Do(ctx context.Context, opts ...RequestOption) (res []*MarginCrossMarginFeeResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: marginCrossMarginFeeEndpoint,
		secType:  secTypeSigned,
	}
	if s.vipLevel != nil {
		r.setParam("vipLevel", *s.vipLevel)
	}
	if s.coin != nil {
		r.setParam("coin", *s.coin)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*MarginCrossMarginFeeResponse{}, err
	}
	res = make([]*MarginCrossMarginFeeResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*MarginCrossMarginFeeResponse{}, err
	}
	return res, nil
}

// MarginCrossMarginFeeService response
type MarginCrossMarginFeeResponse struct {
	VIPLevel        int    `json:"vipLevel"`
	Coin            string `json:"coin"`
	TransferIn      bool   `json:"transferIn"`
	Borrowable      bool   `json:"transferOut"`
	DailyInterest   string `json:"dailyInterest"`
	YearlyInterest  string `json:"yearlyInterest"`
	BorrowLimit     string `json:"borrowLimit"`
	MarginablePairs struct {
		Pair string `json:"pair"`
	} `json:"marginablePairs"`
}

// Query Isolated Margin Fee Data (USER_DATA)
const (
	marginIsolatedMarginFeeEndpoint = "/sapi/v1/margin/isolatedMarginData"
)

type MarginIsolatedMarginFeeService struct {
	c        *Client
	vipLevel *int
	symbol   *string
}

// VipLevel set vipLevel
func (s *MarginIsolatedMarginFeeService) VipLevel(vipLevel int) *MarginIsolatedMarginFeeService {
	s.vipLevel = &vipLevel
	return s
}

// Symbol set symbol
func (s *MarginIsolatedMarginFeeService) Symbol(symbol string) *MarginIsolatedMarginFeeService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *MarginIsolatedMarginFeeService) Do(ctx context.Context, opts ...RequestOption) (res []*MarginIsolatedMarginFeeResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: marginIsolatedMarginFeeEndpoint,
		secType:  secTypeSigned,
	}
	if s.vipLevel != nil {
		r.setParam("vipLevel", *s.vipLevel)
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*MarginIsolatedMarginFeeResponse{}, err
	}
	res = make([]*MarginIsolatedMarginFeeResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*MarginIsolatedMarginFeeResponse{}, err
	}
	return res, nil
}

// MarginIsolatedMarginFeeService response
type MarginIsolatedMarginFeeResponse struct {
	VIPLevel int    `json:"vipLevel"`
	Symbol   string `json:"symbol"`
	Leverage string `json:"leverage"`
	Data     struct {
		Coin          string `json:"coin"`
		DailyInterest string `json:"dailyInterest"`
		BorrowLimit   string `json:"borrowLimit"`
	} `json:"data"`
}

// Query Isolated Margin Tier Data (USER_DATA)
const (
	marginIsolatedMarginTierEndpoint = "/sapi/v1/margin/isolatedMarginTier"
)

type MarginIsolatedMarginTierService struct {
	c      *Client
	symbol string
	tier   *int
}

// Symbol set symbol
func (s *MarginIsolatedMarginTierService) Symbol(symbol string) *MarginIsolatedMarginTierService {
	s.symbol = symbol
	return s
}

// Tier set tier
func (s *MarginIsolatedMarginTierService) Tier(tier int) *MarginIsolatedMarginTierService {
	s.tier = &tier
	return s
}

// Do send request
func (s *MarginIsolatedMarginTierService) Do(ctx context.Context, opts ...RequestOption) (res []*MarginIsolatedMarginTierResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: marginIsolatedMarginTierEndpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol": s.symbol,
	}
	if s.tier != nil {
		m["tier"] = *s.tier
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*MarginIsolatedMarginTierResponse{}, err
	}
	res = make([]*MarginIsolatedMarginTierResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*MarginIsolatedMarginTierResponse{}, err
	}
	return res, nil
}

// MarginIsolatedMarginTierService response
type MarginIsolatedMarginTierResponse struct {
	Symbol                  string `json:"symbol"`
	Tier                    int    `json:"tier"`
	EffectiveMultiple       string `json:"effectiveMultiple"`
	InitialRiskRatio        string `json:"initialRiskRatio"`
	LiquidationRiskRatio    string `json:"liquidationRiskRatio"`
	BaseAssetMaxBorrowable  string `json:"baseAssetMaxBorrowable"`
	QuoteAssetMaxBorrowable string `json:"quoteAssetMaxBorrowable"`
}

// Query Current Margin Order Count Usage (TRADE)
const (
	marginCurrentOrderCountEndpoint = "/sapi/v1/margin/rateLimit/order"
)

type MarginCurrentOrderCountService struct {
	c          *Client
	isIsolated *string
	symbol     *string
}

// IsIsolated set isIsolated
func (s *MarginCurrentOrderCountService) IsIsolated(isIsolated string) *MarginCurrentOrderCountService {
	s.isIsolated = &isIsolated
	return s
}

// Symbol set symbol
func (s *MarginCurrentOrderCountService) Symbol(symbol string) *MarginCurrentOrderCountService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *MarginCurrentOrderCountService) Do(ctx context.Context, opts ...RequestOption) (res []*MarginCurrentOrderCountResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: marginCurrentOrderCountEndpoint,
		secType:  secTypeSigned,
	}
	if s.isIsolated != nil {
		r.setParam("isIsolated", *s.isIsolated)
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*MarginCurrentOrderCountResponse{}, err
	}
	res = make([]*MarginCurrentOrderCountResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*MarginCurrentOrderCountResponse{}, err
	}
	return res, nil
}

// MarginCurrentOrderCountService response
type MarginCurrentOrderCountResponse struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
	Count         int    `json:"count"`
}

// UserAssetDribblet represents one dust log row
type UserAssetDribblet struct {
	OperateTime              int64                     `json:"operateTime"`
	TotalTransferedAmount    string                    `json:"totalTransferedAmount"`    //Total transfered BNB amount for this exchange.
	TotalServiceChargeAmount string                    `json:"totalServiceChargeAmount"` //Total service charge amount for this exchange.
	TransId                  int64                     `json:"transId"`
	UserAssetDribbletDetails []UserAssetDribbletDetail `json:"userAssetDribbletDetails"` //Details of this exchange.
}

// DustLog represents one dust log informations
type UserAssetDribbletDetail struct {
	TransId             int    `json:"transId"`
	ServiceChargeAmount string `json:"serviceChargeAmount"`
	Amount              string `json:"amount"`
	OperateTime         int64  `json:"operateTime"` //The time of this exchange.
	TransferedAmount    string `json:"transferedAmount"`
	FromAsset           string `json:"fromAsset"`
}

// Cross margin collateral ratio (MARKET_DATA)
const (
	marginCrossCollateralRatioEndpoint = "/sapi/v1/margin/crossMarginCollateralRatio"
)

type MarginCrossCollateralRatioService struct {
	c *Client
}

// Do send request
func (s *MarginCrossCollateralRatioService) Do(ctx context.Context, opts ...RequestOption) (res []*MarginCrossCollateralRatioResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: marginCrossCollateralRatioEndpoint,
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*MarginCrossCollateralRatioResponse{}, err
	}
	res = make([]*MarginCrossCollateralRatioResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*MarginCrossCollateralRatioResponse{}, err
	}
	return res, nil
}

// MarginCrossCollateralRatioService response
type MarginCrossCollateralRatioResponse struct {
	Collaterals []*struct {
		MinUsdValue  string `json:"minUsdValue"`
		MaxUsdValue  string `json:"maxUsdValue"`
		DiscountRate string `json:"discountRate"`
	} `json:"collaterals"`
	AssetNames []*struct {
		Asset string `json:"asset"`
	} `json:"assetNames"`
}

// Get Small Liability Exchange Coin List (USER_DATA)
const (
	marginSmallLiabilityExchangeCoinListEndpoint = "/sapi/v1/margin/exchange-small-liability"
)

type MarginSmallLiabilityExchangeCoinListService struct {
	c *Client
}

// Do send request
func (s *MarginSmallLiabilityExchangeCoinListService) Do(ctx context.Context, opts ...RequestOption) (res []*MarginSmallLiabilityExchangeCoinListResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: marginSmallLiabilityExchangeCoinListEndpoint,
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*MarginSmallLiabilityExchangeCoinListResponse{}, err
	}
	res = make([]*MarginSmallLiabilityExchangeCoinListResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*MarginSmallLiabilityExchangeCoinListResponse{}, err
	}
	return res, nil
}

// MarginSmallLiabilityExchangeCoinListService response
type MarginSmallLiabilityExchangeCoinListResponse struct {
	Asset           string `json:"asset"`
	Interest        string `json:"interest"`
	Principal       string `json:"principal"`
	LiabilityOfBUSD string `json:"liabilityOfBUSD"`
}

// Margin Manual Liquidation
const (
	marginManualLiquidationEndpoint = "/sapi/v1/margin/manual-liquidation"
)

type MarginManualLiquidationService struct {
	c          *Client
	marginType string
	symbol     *string
}

// MarginType set marginType
func (s *MarginManualLiquidationService) MarginType(marginType string) *MarginManualLiquidationService {
	s.marginType = marginType
	return s
}

// Symbol set symbol
func (s *MarginManualLiquidationService) Symbol(symbol string) *MarginManualLiquidationService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *MarginManualLiquidationService) Do(ctx context.Context, opts ...RequestOption) (res []*MarginManualLiquidationResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: marginManualLiquidationEndpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"marginType": s.marginType,
	}
	if s.symbol != nil {
		m["symbol"] = *s.symbol
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*MarginManualLiquidationResponse{}, err
	}
	res = make([]*MarginManualLiquidationResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*MarginManualLiquidationResponse{}, err
	}
	return res, nil
}

// MarginManualLiquidationService response
type MarginManualLiquidationResponse struct {
	Asset          string  `json:"asset"`
	Interest       string  `json:"interest"`
	Principal      string  `json:"principal"`
	LiabilityAsset string  `json:"liabilityAsset"`
	LiabilityQty   float64 `json:"liabilityQty"`
}

// Margin Account New OTO (TRADE)
const (
	marginAccountNewOTOEndpoint = "/sapi/v1/margin/order/oto"
)

type MarginAccountNewOTOService struct {
	c                       *Client
	symbol                  string
	workingType             string
	workingSide             string
	workingPrice            float64
	workingQuantity         float64
	pendingType             string
	pendingSide             string
	pendingQuantity         float64
	isIsolated              *string
	listClientOrderId       *string
	newOrderRespType        *string
	sideEffectType          *string
	selfTradePreventionMode *string
	autoRepayAtCancel       *bool
	workingClientOrderId    *string
	workingIcebergQty       *float64
	workingTimeInForce      *string
	pendingClientOrderId    *string
	pendingPrice            *float64
	pendingStopPrice        *float64
	pendingTrailingDelta    *float64
	pendingIcebergQty       *float64
	pendingTimeInForce      *string
}

// Symbol set symbol
func (s *MarginAccountNewOTOService) Symbol(symbol string) *MarginAccountNewOTOService {
	s.symbol = symbol
	return s
}

// WorkingType set workingType
func (s *MarginAccountNewOTOService) WorkingType(workingType string) *MarginAccountNewOTOService {
	s.workingType = workingType
	return s
}

// WorkingSide set workingSide
func (s *MarginAccountNewOTOService) WorkingSide(workingSide string) *MarginAccountNewOTOService {
	s.workingSide = workingSide
	return s
}

// WorkingPrice set workingPrice
func (s *MarginAccountNewOTOService) WorkingPrice(workingPrice float64) *MarginAccountNewOTOService {
	s.workingPrice = workingPrice
	return s
}

// WorkingQuantity set workingQuantity
func (s *MarginAccountNewOTOService) WorkingQuantity(workingQuantity float64) *MarginAccountNewOTOService {
	s.workingQuantity = workingQuantity
	return s
}

// PendingType set pendingType
func (s *MarginAccountNewOTOService) PendingType(pendingType string) *MarginAccountNewOTOService {
	s.pendingType = pendingType
	return s
}

// PendingSide set pendingSide
func (s *MarginAccountNewOTOService) PendingSide(pendingSide string) *MarginAccountNewOTOService {
	s.pendingSide = pendingSide
	return s
}

// PendingQuantity set pendingQuantity
func (s *MarginAccountNewOTOService) PendingQuantity(pendingQuantity float64) *MarginAccountNewOTOService {
	s.pendingQuantity = pendingQuantity
	return s
}

// IsIsolated set isIsolated
func (s *MarginAccountNewOTOService) IsIsolated(isIsolated string) *MarginAccountNewOTOService {
	s.isIsolated = &isIsolated
	return s
}

// ListClientOrderId set listClientOrderId
func (s *MarginAccountNewOTOService) ListClientOrderId(listClientOrderId string) *MarginAccountNewOTOService {
	s.listClientOrderId = &listClientOrderId
	return s
}

// NewOrderRespType set newOrderRespType
func (s *MarginAccountNewOTOService) NewOrderRespType(newOrderRespType string) *MarginAccountNewOTOService {
	s.newOrderRespType = &newOrderRespType
	return s
}

// SideEffectType set sideEffectType
func (s *MarginAccountNewOTOService) SideEffectType(sideEffectType string) *MarginAccountNewOTOService {
	s.sideEffectType = &sideEffectType
	return s
}

// SelfTradePreventionMode set selfTradePreventionMode
func (s *MarginAccountNewOTOService) SelfTradePreventionMode(selfTradePreventionMode string) *MarginAccountNewOTOService {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

// AutoRepayAtCancel set autoRepayAtCancel
func (s *MarginAccountNewOTOService) AutoRepayAtCancel(autoRepayAtCancel bool) *MarginAccountNewOTOService {
	s.autoRepayAtCancel = &autoRepayAtCancel
	return s
}

// WorkingClientOrderId set workingClientOrderId
func (s *MarginAccountNewOTOService) WorkingClientOrderId(workingClientOrderId string) *MarginAccountNewOTOService {
	s.workingClientOrderId = &workingClientOrderId
	return s
}

// WorkingIcebergQty set workingIcebergQty
func (s *MarginAccountNewOTOService) WorkingIcebergQty(workingIcebergQty float64) *MarginAccountNewOTOService {
	s.workingIcebergQty = &workingIcebergQty
	return s
}

// WorkingTimeInForce set workingTimeInForce
func (s *MarginAccountNewOTOService) WorkingTimeInForce(workingTimeInForce string) *MarginAccountNewOTOService {
	s.workingTimeInForce = &workingTimeInForce
	return s
}

// PendingClientOrderId set pendingClientOrderId
func (s *MarginAccountNewOTOService) PendingClientOrderId(pendingClientOrderId string) *MarginAccountNewOTOService {
	s.pendingClientOrderId = &pendingClientOrderId
	return s
}

// PendingPrice set pendingPrice
func (s *MarginAccountNewOTOService) PendingPrice(pendingPrice float64) *MarginAccountNewOTOService {
	s.pendingPrice = &pendingPrice
	return s
}

// PendingStopPrice set pendingStopPrice
func (s *MarginAccountNewOTOService) PendingStopPrice(pendingStopPrice float64) *MarginAccountNewOTOService {
	s.pendingStopPrice = &pendingStopPrice
	return s
}

// PendingTrailingDelta set pendingTrailingDelta
func (s *MarginAccountNewOTOService) PendingTrailingDelta(pendingTrailingDelta float64) *MarginAccountNewOTOService {
	s.pendingTrailingDelta = &pendingTrailingDelta
	return s
}

// PendingIcebergQty set pendingIcebergQty
func (s *MarginAccountNewOTOService) PendingIcebergQty(pendingIcebergQty float64) *MarginAccountNewOTOService {
	s.pendingIcebergQty = &pendingIcebergQty
	return s
}

// PendingTimeInForce set pendingTimeInForce
func (s *MarginAccountNewOTOService) PendingTimeInForce(pendingTimeInForce string) *MarginAccountNewOTOService {
	s.pendingTimeInForce = &pendingTimeInForce
	return s
}

// Do send request
func (s *MarginAccountNewOTOService) Do(ctx context.Context, opts ...RequestOption) (res *MarginAccountNewOTOResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: marginAccountNewOTOEndpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol":          s.symbol,
		"workingType":     s.workingType,
		"workingSide":     s.workingSide,
		"workingPrice":    s.workingPrice,
		"workingQuantity": s.workingQuantity,
		"pendingType":     s.pendingType,
		"pendingSide":     s.pendingSide,
		"pendingQuantity": s.pendingQuantity,
	}
	if s.isIsolated != nil {
		m["isIsolated"] = *s.isIsolated
	}
	if s.listClientOrderId != nil {
		m["listClientOrderId"] = *s.listClientOrderId
	}
	if s.newOrderRespType != nil {
		m["newOrderRespType"] = *s.newOrderRespType
	}
	if s.sideEffectType != nil {
		m["sideEffectType"] = *s.sideEffectType
	}
	if s.selfTradePreventionMode != nil {
		m["selfTradePreventionMode"] = *s.selfTradePreventionMode
	}
	if s.autoRepayAtCancel != nil {
		m["autoRepayAtCancel"] = *s.autoRepayAtCancel
	}
	if s.workingClientOrderId != nil {
		m["workingClientOrderId"] = *s.workingClientOrderId
	}
	if s.workingIcebergQty != nil {
		m["workingIcebergQty"] = *s.workingIcebergQty
	}
	if s.workingTimeInForce != nil {
		m["workingTimeInForce"] = *s.workingTimeInForce
	}
	if s.pendingClientOrderId != nil {
		m["pendingClientOrderId"] = *s.pendingClientOrderId
	}
	if s.pendingPrice != nil {
		m["pendingPrice"] = *s.pendingPrice
	}
	if s.pendingStopPrice != nil {
		m["pendingStopPrice"] = *s.pendingStopPrice
	}
	if s.pendingTrailingDelta != nil {
		m["pendingTrailingDelta"] = *s.pendingTrailingDelta
	}
	if s.pendingIcebergQty != nil {
		m["pendingIcebergQty"] = *s.pendingIcebergQty
	}
	if s.pendingTimeInForce != nil {
		m["pendingTimeInForce"] = *s.pendingTimeInForce
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &MarginAccountNewOTOResponse{}, err
	}
	res = new(MarginAccountNewOTOResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return &MarginAccountNewOTOResponse{}, err
	}
	return res, nil
}

// MarginAccountNewOTOResponse represents the response from the MarginAccountNewOTOService
type MarginAccountNewOTOResponse struct {
	OrderListId       int64  `json:"orderListId"`
	ContingencyType   string `json:"contingencyType"`
	ListStatusType    string `json:"listStatusType"`
	ListOrderStatus   string `json:"listOrderStatus"`
	ListClientOrderId string `json:"listClientOrderId"`
	TransactionTime   int64  `json:"transactionTime"`
	Symbol            string `json:"symbol"`
	IsIsolated        bool   `json:"isIsolated"`
	Orders            []struct {
		Symbol        string `json:"symbol"`
		OrderId       int64  `json:"orderId"`
		ClientOrderId string `json:"clientOrderId"`
	}
	OrderReports []struct {
		Symbol                  string `json:"symbol"`
		OrderId                 int64  `json:"orderId"`
		OrderListId             int64  `json:"orderListId"`
		ClientOrderId           string `json:"clientOrderId"`
		TransactTime            int64  `json:"transactTime"`
		Price                   string `json:"price"`
		OrigQty                 string `json:"origQty"`
		ExecutedQty             string `json:"executedQty"`
		CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`
		Status                  string `json:"status"`
		TimeInForce             string `json:"timeInForce"`
		Type                    string `json:"type"`
		Side                    string `json:"side"`
		SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	}
}

// Margin Account New OTOCO (TRADE)
const (
	marginAccountNewOTOCOEndpoint = "/sapi/v1/margin/order/otoco"
)

type MarginAccountNewOTOCOService struct {
	c                         *Client
	symbol                    string
	workingType               string
	workingSide               string
	workingPrice              float64
	workingQuantity           float64
	pendingSide               string
	pendingQuantity           float64
	pendingAboveType          string
	isIsolated                *string
	sideEffectType            *string
	autoRepayAtCancel         *bool
	listClientOrderId         *string
	newOrderRespType          *string
	selfTradePreventionMode   *string
	workingClientOrderId      *string
	workingIcebergQty         *float64
	workingTimeInForce        *string
	pendingAboveClientOrderId *string
	pendingAbovePrice         *float64
	pendingAboveStopPrice     *float64
	pendingAboveTrailingDelta *float64
	pendingAboveIcebergQty    *float64
	pendingAboveTimeInForce   *string
	pendingBelowType          *string
	pendingBelowClientOrderId *string
	pendingBelowPrice         *float64
	pendingBelowStopPrice     *float64
	pendingBelowTrailingDelta *float64
	pendingBelowIcebergQty    *float64
	pendingBelowTimeInForce   *string
}

// Symbol set symbol
func (s *MarginAccountNewOTOCOService) Symbol(symbol string) *MarginAccountNewOTOCOService {
	s.symbol = symbol
	return s
}

// WorkingType set workingType
func (s *MarginAccountNewOTOCOService) WorkingType(workingType string) *MarginAccountNewOTOCOService {
	s.workingType = workingType
	return s
}

// WorkingSide set workingSide
func (s *MarginAccountNewOTOCOService) WorkingSide(workingSide string) *MarginAccountNewOTOCOService {
	s.workingSide = workingSide
	return s
}

// WorkingPrice set workingPrice
func (s *MarginAccountNewOTOCOService) WorkingPrice(workingPrice float64) *MarginAccountNewOTOCOService {
	s.workingPrice = workingPrice
	return s
}

// WorkingQuantity set workingQuantity
func (s *MarginAccountNewOTOCOService) WorkingQuantity(workingQuantity float64) *MarginAccountNewOTOCOService {
	s.workingQuantity = workingQuantity
	return s
}

// PendingSide set pendingSide
func (s *MarginAccountNewOTOCOService) PendingSide(pendingSide string) *MarginAccountNewOTOCOService {
	s.pendingSide = pendingSide
	return s
}

// PendingQuantity set pendingQuantity
func (s *MarginAccountNewOTOCOService) PendingQuantity(pendingQuantity float64) *MarginAccountNewOTOCOService {
	s.pendingQuantity = pendingQuantity
	return s
}

// PendingAboveType set pendingAboveType
func (s *MarginAccountNewOTOCOService) PendingAboveType(pendingAboveType string) *MarginAccountNewOTOCOService {
	s.pendingAboveType = pendingAboveType
	return s
}

// IsIsolated set isIsolated
func (s *MarginAccountNewOTOCOService) IsIsolated(isIsolated string) *MarginAccountNewOTOCOService {
	s.isIsolated = &isIsolated
	return s
}

// SideEffectType set sideEffectType
func (s *MarginAccountNewOTOCOService) SideEffectType(sideEffectType string) *MarginAccountNewOTOCOService {
	s.sideEffectType = &sideEffectType
	return s
}

// AutoRepayAtCancel set autoRepayAtCancel
func (s *MarginAccountNewOTOCOService) AutoRepayAtCancel(autoRepayAtCancel bool) *MarginAccountNewOTOCOService {
	s.autoRepayAtCancel = &autoRepayAtCancel
	return s
}

// ListClientOrderId set listClientOrderId
func (s *MarginAccountNewOTOCOService) ListClientOrderId(listClientOrderId string) *MarginAccountNewOTOCOService {
	s.listClientOrderId = &listClientOrderId
	return s
}

// NewOrderRespType set newOrderRespType
func (s *MarginAccountNewOTOCOService) NewOrderRespType(newOrderRespType string) *MarginAccountNewOTOCOService {
	s.newOrderRespType = &newOrderRespType
	return s
}

// SelfTradePreventionMode set selfTradePreventionMode
func (s *MarginAccountNewOTOCOService) SelfTradePreventionMode(selfTradePreventionMode string) *MarginAccountNewOTOCOService {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

// WorkingClientOrderId set workingClientOrderId
func (s *MarginAccountNewOTOCOService) WorkingClientOrderId(workingClientOrderId string) *MarginAccountNewOTOCOService {
	s.workingClientOrderId = &workingClientOrderId
	return s
}

// WorkingIcebergQty set workingIcebergQty
func (s *MarginAccountNewOTOCOService) WorkingIcebergQty(workingIcebergQty float64) *MarginAccountNewOTOCOService {
	s.workingIcebergQty = &workingIcebergQty
	return s
}

// WorkingTimeInForce set workingTimeInForce
func (s *MarginAccountNewOTOCOService) WorkingTimeInForce(workingTimeInForce string) *MarginAccountNewOTOCOService {
	s.workingTimeInForce = &workingTimeInForce
	return s
}

// PendingAboveClientOrderId set pendingAboveClientOrderId
func (s *MarginAccountNewOTOCOService) PendingAboveClientOrderId(pendingAboveClientOrderId string) *MarginAccountNewOTOCOService {
	s.pendingAboveClientOrderId = &pendingAboveClientOrderId
	return s
}

// PendingAbovePrice set pendingAbovePrice
func (s *MarginAccountNewOTOCOService) PendingAbovePrice(pendingAbovePrice float64) *MarginAccountNewOTOCOService {
	s.pendingAbovePrice = &pendingAbovePrice
	return s
}

// PendingAboveStopPrice set pendingAboveStopPrice
func (s *MarginAccountNewOTOCOService) PendingAboveStopPrice(pendingAboveStopPrice float64) *MarginAccountNewOTOCOService {
	s.pendingAboveStopPrice = &pendingAboveStopPrice
	return s
}

// PendingAboveTrailingDelta set pendingAboveTrailingDelta
func (s *MarginAccountNewOTOCOService) PendingAboveTrailingDelta(pendingAboveTrailingDelta float64) *MarginAccountNewOTOCOService {
	s.pendingAboveTrailingDelta = &pendingAboveTrailingDelta
	return s
}

// PendingAboveIcebergQty set pendingAboveIcebergQty
func (s *MarginAccountNewOTOCOService) PendingAboveIcebergQty(pendingAboveIcebergQty float64) *MarginAccountNewOTOCOService {
	s.pendingAboveIcebergQty = &pendingAboveIcebergQty
	return s
}

// PendingAboveTimeInForce set pendingAboveTimeInForce
func (s *MarginAccountNewOTOCOService) PendingAboveTimeInForce(pendingAboveTimeInForce string) *MarginAccountNewOTOCOService {
	s.pendingAboveTimeInForce = &pendingAboveTimeInForce
	return s
}

// PendingBelowType set pendingBelowType
func (s *MarginAccountNewOTOCOService) PendingBelowType(pendingBelowType string) *MarginAccountNewOTOCOService {
	s.pendingBelowType = &pendingBelowType
	return s
}

// PendingBelowClientOrderId set pendingBelowClientOrderId
func (s *MarginAccountNewOTOCOService) PendingBelowClientOrderId(pendingBelowClientOrderId string) *MarginAccountNewOTOCOService {
	s.pendingBelowClientOrderId = &pendingBelowClientOrderId
	return s
}

// PendingBelowPrice set pendingBelowPrice
func (s *MarginAccountNewOTOCOService) PendingBelowPrice(pendingBelowPrice float64) *MarginAccountNewOTOCOService {
	s.pendingBelowPrice = &pendingBelowPrice
	return s
}

// PendingBelowStopPrice set pendingBelowStopPrice
func (s *MarginAccountNewOTOCOService) PendingBelowStopPrice(pendingBelowStopPrice float64) *MarginAccountNewOTOCOService {
	s.pendingBelowStopPrice = &pendingBelowStopPrice
	return s
}

// PendingBelowTrailingDelta set pendingBelowTrailingDelta
func (s *MarginAccountNewOTOCOService) PendingBelowTrailingDelta(pendingBelowTrailingDelta float64) *MarginAccountNewOTOCOService {
	s.pendingBelowTrailingDelta = &pendingBelowTrailingDelta
	return s
}

// PendingBelowIcebergQty set pendingBelowIcebergQty
func (s *MarginAccountNewOTOCOService) PendingBelowIcebergQty(pendingBelowIcebergQty float64) *MarginAccountNewOTOCOService {
	s.pendingBelowIcebergQty = &pendingBelowIcebergQty
	return s
}

// PendingBelowTimeInForce set pendingBelowTimeInForce
func (s *MarginAccountNewOTOCOService) PendingBelowTimeInForce(pendingBelowTimeInForce string) *MarginAccountNewOTOCOService {
	s.pendingBelowTimeInForce = &pendingBelowTimeInForce
	return s
}

// Do send request
func (s *MarginAccountNewOTOCOService) Do(ctx context.Context, opts ...RequestOption) (res *MarginAccountNewOTOCOResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: marginAccountNewOTOCOEndpoint,
		secType:  secTypeSigned,
	}

	m := params{
		"symbol":           s.symbol,
		"workingType":      s.workingType,
		"workingSide":      s.workingSide,
		"workingPrice":     s.workingPrice,
		"workingQuantity":  s.workingQuantity,
		"pendingSide":      s.pendingSide,
		"pendingQuantity":  s.pendingQuantity,
		"pendingAboveType": s.pendingAboveType,
	}

	if s.isIsolated != nil {
		m["isIsolated"] = *s.isIsolated
	}
	if s.sideEffectType != nil {
		m["sideEffectType"] = *s.sideEffectType
	}
	if s.autoRepayAtCancel != nil {
		m["autoRepayAtCancel"] = *s.autoRepayAtCancel
	}
	if s.listClientOrderId != nil {
		m["listClientOrderId"] = *s.listClientOrderId
	}
	if s.newOrderRespType != nil {
		m["newOrderRespType"] = *s.newOrderRespType
	}
	if s.selfTradePreventionMode != nil {
		m["selfTradePreventionMode"] = *s.selfTradePreventionMode
	}
	if s.workingClientOrderId != nil {
		m["workingClientOrderId"] = *s.workingClientOrderId
	}
	if s.workingIcebergQty != nil {
		m["workingIcebergQty"] = *s.workingIcebergQty
	}
	if s.workingTimeInForce != nil {
		m["workingTimeInForce"] = *s.workingTimeInForce
	}
	if s.pendingAboveClientOrderId != nil {
		m["pendingAboveClientOrderId"] = *s.pendingAboveClientOrderId
	}
	if s.pendingAbovePrice != nil {
		m["pendingAbovePrice"] = *s.pendingAbovePrice
	}
	if s.pendingAboveStopPrice != nil {
		m["pendingAboveStopPrice"] = *s.pendingAboveStopPrice
	}
	if s.pendingAboveTrailingDelta != nil {
		m["pendingAboveTrailingDelta"] = *s.pendingAboveTrailingDelta
	}
	if s.pendingAboveIcebergQty != nil {
		m["pendingAboveIcebergQty"] = *s.pendingAboveIcebergQty
	}
	if s.pendingAboveTimeInForce != nil {
		m["pendingAboveTimeInForce"] = *s.pendingAboveTimeInForce
	}
	if s.pendingBelowType != nil {
		m["pendingBelowType"] = *s.pendingBelowType
	}
	if s.pendingBelowClientOrderId != nil {
		m["pendingBelowClientOrderId"] = *s.pendingBelowClientOrderId
	}
	if s.pendingBelowPrice != nil {
		m["pendingBelowPrice"] = *s.pendingBelowPrice
	}
	if s.pendingBelowStopPrice != nil {
		m["pendingBelowStopPrice"] = *s.pendingBelowStopPrice
	}
	if s.pendingBelowTrailingDelta != nil {
		m["pendingBelowTrailingDelta"] = *s.pendingBelowTrailingDelta
	}
	if s.pendingBelowIcebergQty != nil {
		m["pendingBelowIcebergQty"] = *s.pendingBelowIcebergQty
	}
	if s.pendingBelowTimeInForce != nil {
		m["pendingBelowTimeInForce"] = *s.pendingBelowTimeInForce
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &MarginAccountNewOTOCOResponse{}, err
	}
	res = new(MarginAccountNewOTOCOResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return &MarginAccountNewOTOCOResponse{}, err
	}
	return res, nil
}

// MarginAccountNewOTOCOResponse represents the response from the MarginAccountNewOTOCOService
type MarginAccountNewOTOCOResponse struct {
	OrderListId       int64  `json:"orderListId"`
	ContingencyType   string `json:"contingencyType"`
	ListStatusType    string `json:"listStatusType"`
	ListOrderStatus   string `json:"listOrderStatus"`
	ListClientOrderId string `json:"listClientOrderId"`
	TransactionTime   int64  `json:"transactionTime"`
	Symbol            string `json:"symbol"`
	IsIsolated        bool   `json:"isIsolated"`
	Orders            []struct {
		Symbol        string `json:"symbol"`
		OrderId       int64  `json:"orderId"`
		ClientOrderId string `json:"clientOrderId"`
	}
	OrderReports []struct {
		Symbol                  string `json:"symbol"`
		OrderId                 int64  `json:"orderId"`
		OrderListId             int64  `json:"orderListId"`
		ClientOrderId           string `json:"clientOrderId"`
		TransactTime            int64  `json:"transactTime"`
		Price                   string `json:"price"`
		OrigQty                 string `json:"origQty"`
		ExecutedQty             string `json:"executedQty"`
		CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`
		Status                  string `json:"status"`
		TimeInForce             string `json:"timeInForce"`
		Type                    string `json:"type"`
		Side                    string `json:"side"`
		StopPrice               string `json:"stopPrice"`
		SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	}
}

// Small Liability Exchange (MARGIN)
const (
	marginSmallLiabilityExchangeEndpoint = "/sapi/v1/margin/exchange-small-liability"
)

type MarginSmallLiabilityExchangeService struct {
	c          *Client
	assetNames string
}

// AssetNames set assetNames
func (s *MarginSmallLiabilityExchangeService) AssetNames(assetNames string) *MarginSmallLiabilityExchangeService {
	s.assetNames = assetNames
	return s
}

// Do send request
func (s *MarginSmallLiabilityExchangeService) Do(ctx context.Context, opts ...RequestOption) (res []*MarginSmallLiabilityExchangeResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: marginSmallLiabilityExchangeEndpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"assetNames": s.assetNames,
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*MarginSmallLiabilityExchangeResponse{}, err
	}
	res = make([]*MarginSmallLiabilityExchangeResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*MarginSmallLiabilityExchangeResponse{}, err
	}
	return res, nil
}

// MarginSmallLiabilityExchangeService response
type MarginSmallLiabilityExchangeResponse struct {
}

// Get Small Liability Exchange History (USER_DATA)
const (
	marginSmallLiabilityExchangeHistoryEndpoint = "/sapi/v1/margin/exchange-small-liability-history"
)

type MarginSmallLiabilityExchangeHistoryService struct {
	c         *Client
	current   int
	size      int
	startTime *uint64
	endTime   *uint64
}

// Current set current
func (s *MarginSmallLiabilityExchangeHistoryService) Current(current int) *MarginSmallLiabilityExchangeHistoryService {
	s.current = current
	return s
}

// Size set size
func (s *MarginSmallLiabilityExchangeHistoryService) Size(size int) *MarginSmallLiabilityExchangeHistoryService {
	s.size = size
	return s
}

// StartTime set startTime
func (s *MarginSmallLiabilityExchangeHistoryService) StartTime(startTime uint64) *MarginSmallLiabilityExchangeHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *MarginSmallLiabilityExchangeHistoryService) EndTime(endTime uint64) *MarginSmallLiabilityExchangeHistoryService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *MarginSmallLiabilityExchangeHistoryService) Do(ctx context.Context, opts ...RequestOption) (res []*MarginSmallLiabilityExchangeHistoryResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: marginSmallLiabilityExchangeHistoryEndpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"current": s.current,
		"size":    s.size,
	}
	if s.startTime != nil {
		m["startTime"] = *s.startTime
	}
	if s.endTime != nil {
		m["endTime"] = *s.endTime
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*MarginSmallLiabilityExchangeHistoryResponse{}, err
	}
	res = make([]*MarginSmallLiabilityExchangeHistoryResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*MarginSmallLiabilityExchangeHistoryResponse{}, err
	}
	return res, nil
}

// MarginSmallLiabilityExchangeHistoryService response
type MarginSmallLiabilityExchangeHistoryResponse struct {
	Total int `json:"total"`
	Rows  []*struct {
		Asset        string `json:"asset"`
		Amount       string `json:"amount"`
		TargetAsset  string `json:"targetAsset"`
		TargetAmount string `json:"targetAmount"`
		BizType      string `json:"bizType"`
		Timestamp    uint64 `json:"timestamp"`
	} `json:"rows"`
}
