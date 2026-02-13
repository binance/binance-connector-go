/*
Binance Derivatives Trading COIN Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket API
*/

package binancederivativestradingcoinfutureswebsocketapi

import (
	"github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/websocketapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// TradeAPIService TradeAPI Service
type TradeAPIService struct {
	Ws *common.WebsocketAPI
}

type ApiCancelOrderRequest struct {
	ApiService        *TradeAPIService
	symbol            *string
	id                *string
	orderId           *int64
	origClientOrderId *string
	recvWindow        *int64
}

func (r ApiCancelOrderRequest) Symbol(symbol string) ApiCancelOrderRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiCancelOrderRequest) Id(id string) ApiCancelOrderRequest {
	r.id = &id
	return r
}

func (r ApiCancelOrderRequest) OrderId(orderId int64) ApiCancelOrderRequest {
	r.orderId = &orderId
	return r
}

func (r ApiCancelOrderRequest) OrigClientOrderId(origClientOrderId string) ApiCancelOrderRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

func (r ApiCancelOrderRequest) RecvWindow(recvWindow int64) ApiCancelOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelOrderRequest) Execute() (*common.ResponseOrRaw[models.CancelOrderResponse], error) {
	respChan, errChan, err := r.ApiService.CancelOrderExecute(r)
	if err != nil {
		return nil, err
	}

	select {
	case resp := <-respChan:
		return resp, nil
	case err := <-errChan:
		return nil, err
	}
}

func (r ApiCancelOrderRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.CancelOrderResponse], chan error, error) {
	return r.ApiService.CancelOrderExecute(r)
}

/*
CancelOrder Cancel Order (TRADE)
/order.cancel

https://developers.binance.com/docs/derivatives/coin-margined-futures/trade/websocket-api/Cancel-Order

@param symbol	@param id Unique WebSocket request ID.	@param orderId	@param origClientOrderId	@param recvWindow
@return ApiCancelOrderRequest
*/
func (a *TradeAPIService) CancelOrder() ApiCancelOrderRequest {
	return ApiCancelOrderRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return CancelOrderResponse
func (a *TradeAPIService) CancelOrderExecute(r ApiCancelOrderRequest) (chan *common.ResponseOrRaw[models.CancelOrderResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.orderId != nil {
		localVarQueryParams["orderId"] = *r.orderId
	}
	if r.origClientOrderId != nil {
		localVarQueryParams["origClientOrderId"] = *r.origClientOrderId
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/order.cancel"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.CancelOrderResponse](a.Ws, localPayload, sendParams)
}

type ApiModifyOrderRequest struct {
	ApiService        *TradeAPIService
	symbol            *string
	side              *models.ModifyOrderSideParameter
	quantity          *float32
	price             *float32
	id                *string
	orderId           *int64
	origClientOrderId *string
	priceMatch        *models.ModifyOrderPriceMatchParameter
	recvWindow        *int64
}

func (r ApiModifyOrderRequest) Symbol(symbol string) ApiModifyOrderRequest {
	r.symbol = &symbol
	return r
}

// &#x60;SELL&#x60;, &#x60;BUY&#x60;
func (r ApiModifyOrderRequest) Side(side models.ModifyOrderSideParameter) ApiModifyOrderRequest {
	r.side = &side
	return r
}

// Order quantity, cannot be sent with &#x60;closePosition&#x3D;true&#x60;
func (r ApiModifyOrderRequest) Quantity(quantity float32) ApiModifyOrderRequest {
	r.quantity = &quantity
	return r
}

func (r ApiModifyOrderRequest) Price(price float32) ApiModifyOrderRequest {
	r.price = &price
	return r
}

// Unique WebSocket request ID.
func (r ApiModifyOrderRequest) Id(id string) ApiModifyOrderRequest {
	r.id = &id
	return r
}

func (r ApiModifyOrderRequest) OrderId(orderId int64) ApiModifyOrderRequest {
	r.orderId = &orderId
	return r
}

func (r ApiModifyOrderRequest) OrigClientOrderId(origClientOrderId string) ApiModifyOrderRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

// only available for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; can be set to &#x60;OPPONENT&#x60;/ &#x60;OPPONENT_5&#x60;/ &#x60;OPPONENT_10&#x60;/ &#x60;OPPONENT_20&#x60;: /&#x60;QUEUE&#x60;/ &#x60;QUEUE_5&#x60;/ &#x60;QUEUE_10&#x60;/ &#x60;QUEUE_20&#x60;; Can&#39;t be passed together with &#x60;price&#x60;
func (r ApiModifyOrderRequest) PriceMatch(priceMatch models.ModifyOrderPriceMatchParameter) ApiModifyOrderRequest {
	r.priceMatch = &priceMatch
	return r
}

func (r ApiModifyOrderRequest) RecvWindow(recvWindow int64) ApiModifyOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiModifyOrderRequest) Execute() (*common.ResponseOrRaw[models.ModifyOrderResponse], error) {
	respChan, errChan, err := r.ApiService.ModifyOrderExecute(r)
	if err != nil {
		return nil, err
	}

	select {
	case resp := <-respChan:
		return resp, nil
	case err := <-errChan:
		return nil, err
	}
}

func (r ApiModifyOrderRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.ModifyOrderResponse], chan error, error) {
	return r.ApiService.ModifyOrderExecute(r)
}

/*
ModifyOrder Modify Order (TRADE)
/order.modify

https://developers.binance.com/docs/derivatives/coin-margined-futures/trade/websocket-api/Modify-Order

@param symbol	@param side `SELL`, `BUY`	@param quantity Order quantity, cannot be sent with `closePosition=true`	@param price	@param id Unique WebSocket request ID.	@param orderId	@param origClientOrderId	@param priceMatch only available for `LIMIT`/`STOP`/`TAKE_PROFIT` order; can be set to `OPPONENT`/ `OPPONENT_5`/ `OPPONENT_10`/ `OPPONENT_20`: /`QUEUE`/ `QUEUE_5`/ `QUEUE_10`/ `QUEUE_20`; Can't be passed together with `price`	@param recvWindow
@return ApiModifyOrderRequest
*/
func (a *TradeAPIService) ModifyOrder() ApiModifyOrderRequest {
	return ApiModifyOrderRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return ModifyOrderResponse
func (a *TradeAPIService) ModifyOrderExecute(r ApiModifyOrderRequest) (chan *common.ResponseOrRaw[models.ModifyOrderResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.side == nil {
		return nil, nil, common.ReportError("side is required and must be specified")
	}
	localVarQueryParams["side"] = *r.side

	if r.quantity == nil {
		return nil, nil, common.ReportError("quantity is required and must be specified")
	}
	localVarQueryParams["quantity"] = *r.quantity

	if r.price == nil {
		return nil, nil, common.ReportError("price is required and must be specified")
	}
	localVarQueryParams["price"] = *r.price

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.orderId != nil {
		localVarQueryParams["orderId"] = *r.orderId
	}
	if r.origClientOrderId != nil {
		localVarQueryParams["origClientOrderId"] = *r.origClientOrderId
	}
	if r.priceMatch != nil {
		localVarQueryParams["priceMatch"] = *r.priceMatch
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/order.modify"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.ModifyOrderResponse](a.Ws, localPayload, sendParams)
}

type ApiNewOrderRequest struct {
	ApiService              *TradeAPIService
	symbol                  *string
	side                    *models.ModifyOrderSideParameter
	type_                   *models.NewOrderTypeParameter
	id                      *string
	positionSide            *models.NewOrderPositionSideParameter
	timeInForce             *models.NewOrderTimeInForceParameter
	quantity                *float32
	reduceOnly              *string
	price                   *float32
	newClientOrderId        *string
	stopPrice               *float32
	closePosition           *string
	activationPrice         *float32
	callbackRate            *float32
	workingType             *models.NewOrderWorkingTypeParameter
	priceProtect            *string
	newOrderRespType        *models.NewOrderNewOrderRespTypeParameter
	priceMatch              *models.ModifyOrderPriceMatchParameter
	selfTradePreventionMode *models.NewOrderSelfTradePreventionModeParameter
	recvWindow              *int64
}

func (r ApiNewOrderRequest) Symbol(symbol string) ApiNewOrderRequest {
	r.symbol = &symbol
	return r
}

// &#x60;SELL&#x60;, &#x60;BUY&#x60;
func (r ApiNewOrderRequest) Side(side models.ModifyOrderSideParameter) ApiNewOrderRequest {
	r.side = &side
	return r
}

// &#x60;LIMIT&#x60;, &#x60;MARKET&#x60;, &#x60;STOP&#x60;, &#x60;STOP_MARKET&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_MARKET&#x60;, &#x60;TRAILING_STOP_MARKET&#x60;
func (r ApiNewOrderRequest) Type(type_ models.NewOrderTypeParameter) ApiNewOrderRequest {
	r.type_ = &type_
	return r
}

// Unique WebSocket request ID.
func (r ApiNewOrderRequest) Id(id string) ApiNewOrderRequest {
	r.id = &id
	return r
}

// Default &#x60;BOTH&#x60; for One-way Mode; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode.  It must be sent in Hedge Mode.
func (r ApiNewOrderRequest) PositionSide(positionSide models.NewOrderPositionSideParameter) ApiNewOrderRequest {
	r.positionSide = &positionSide
	return r
}

func (r ApiNewOrderRequest) TimeInForce(timeInForce models.NewOrderTimeInForceParameter) ApiNewOrderRequest {
	r.timeInForce = &timeInForce
	return r
}

// Quantity measured by contract number, Cannot be sent with &#x60;closePosition&#x60;&#x3D;&#x60;true&#x60;
func (r ApiNewOrderRequest) Quantity(quantity float32) ApiNewOrderRequest {
	r.quantity = &quantity
	return r
}

// &#x60;true&#x60; or &#x60;false&#x60;. default &#x60;false&#x60;. Cannot be sent in Hedge Mode; cannot be sent with &#x60;closePosition&#x60;&#x3D;&#x60;true&#x60; (Close-All)
func (r ApiNewOrderRequest) ReduceOnly(reduceOnly string) ApiNewOrderRequest {
	r.reduceOnly = &reduceOnly
	return r
}

func (r ApiNewOrderRequest) Price(price float32) ApiNewOrderRequest {
	r.price = &price
	return r
}

// A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: &#x60;^[\\.A-Z\\:/a-z0-9_-]{1,36}$&#x60;
func (r ApiNewOrderRequest) NewClientOrderId(newClientOrderId string) ApiNewOrderRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

// Used with &#x60;STOP/STOP_MARKET&#x60; or &#x60;TAKE_PROFIT/TAKE_PROFIT_MARKET&#x60; orders.
func (r ApiNewOrderRequest) StopPrice(stopPrice float32) ApiNewOrderRequest {
	r.stopPrice = &stopPrice
	return r
}

// &#x60;true&#x60;, &#x60;false&#x60;；Close-All，used with &#x60;STOP_MARKET&#x60; or &#x60;TAKE_PROFIT_MARKET&#x60;.
func (r ApiNewOrderRequest) ClosePosition(closePosition string) ApiNewOrderRequest {
	r.closePosition = &closePosition
	return r
}

// Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, default as the latest price(supporting different workingType)
func (r ApiNewOrderRequest) ActivationPrice(activationPrice float32) ApiNewOrderRequest {
	r.activationPrice = &activationPrice
	return r
}

// Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, min 0.1, max 10 where 1 for 1%
func (r ApiNewOrderRequest) CallbackRate(callbackRate float32) ApiNewOrderRequest {
	r.callbackRate = &callbackRate
	return r
}

// stopPrice triggered by: \&quot;MARK_PRICE\&quot;, \&quot;CONTRACT_PRICE\&quot;. Default \&quot;CONTRACT_PRICE\&quot;
func (r ApiNewOrderRequest) WorkingType(workingType models.NewOrderWorkingTypeParameter) ApiNewOrderRequest {
	r.workingType = &workingType
	return r
}

// \&quot;TRUE\&quot; or \&quot;FALSE\&quot;, default \&quot;FALSE\&quot;. Used with &#x60;STOP/STOP_MARKET&#x60; or &#x60;TAKE_PROFIT/TAKE_PROFIT_MARKET&#x60; orders.
func (r ApiNewOrderRequest) PriceProtect(priceProtect string) ApiNewOrderRequest {
	r.priceProtect = &priceProtect
	return r
}

// &#x60;ACK&#x60;,&#x60;RESULT&#x60;, default &#x60;ACK&#x60;
func (r ApiNewOrderRequest) NewOrderRespType(newOrderRespType models.NewOrderNewOrderRespTypeParameter) ApiNewOrderRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

// only available for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; can be set to &#x60;OPPONENT&#x60;/ &#x60;OPPONENT_5&#x60;/ &#x60;OPPONENT_10&#x60;/ &#x60;OPPONENT_20&#x60;: /&#x60;QUEUE&#x60;/ &#x60;QUEUE_5&#x60;/ &#x60;QUEUE_10&#x60;/ &#x60;QUEUE_20&#x60;; Can&#39;t be passed together with &#x60;price&#x60;
func (r ApiNewOrderRequest) PriceMatch(priceMatch models.ModifyOrderPriceMatchParameter) ApiNewOrderRequest {
	r.priceMatch = &priceMatch
	return r
}

// &#x60;NONE&#x60;: No STP / &#x60;EXPIRE_TAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_MAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_BOTH&#x60;:expire both orders when STP triggers; default &#x60;NONE&#x60;
func (r ApiNewOrderRequest) SelfTradePreventionMode(selfTradePreventionMode models.NewOrderSelfTradePreventionModeParameter) ApiNewOrderRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

func (r ApiNewOrderRequest) RecvWindow(recvWindow int64) ApiNewOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiNewOrderRequest) Execute() (*common.ResponseOrRaw[models.NewOrderResponse], error) {
	respChan, errChan, err := r.ApiService.NewOrderExecute(r)
	if err != nil {
		return nil, err
	}

	select {
	case resp := <-respChan:
		return resp, nil
	case err := <-errChan:
		return nil, err
	}
}

func (r ApiNewOrderRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.NewOrderResponse], chan error, error) {
	return r.ApiService.NewOrderExecute(r)
}

/*
NewOrder New Order(TRADE)
/order.place

https://developers.binance.com/docs/derivatives/coin-margined-futures/trade/websocket-api/New-Order

@param symbol	@param side `SELL`, `BUY`	@param type_ `LIMIT`, `MARKET`, `STOP`, `STOP_MARKET`, `TAKE_PROFIT`, `TAKE_PROFIT_MARKET`, `TRAILING_STOP_MARKET`	@param id Unique WebSocket request ID.	@param positionSide Default `BOTH` for One-way Mode; `LONG` or `SHORT` for Hedge Mode.  It must be sent in Hedge Mode.	@param timeInForce	@param quantity Quantity measured by contract number, Cannot be sent with `closePosition`=`true`	@param reduceOnly `true` or `false`. default `false`. Cannot be sent in Hedge Mode; cannot be sent with `closePosition`=`true` (Close-All)	@param price	@param newClientOrderId A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: `^[\\.A-Z\\:/a-z0-9_-]{1,36}$`	@param stopPrice Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders.	@param closePosition `true`, `false`；Close-All，used with `STOP_MARKET` or `TAKE_PROFIT_MARKET`.	@param activationPrice Used with `TRAILING_STOP_MARKET` orders, default as the latest price(supporting different workingType)	@param callbackRate Used with `TRAILING_STOP_MARKET` orders, min 0.1, max 10 where 1 for 1%	@param workingType stopPrice triggered by: \"MARK_PRICE\", \"CONTRACT_PRICE\". Default \"CONTRACT_PRICE\"	@param priceProtect \"TRUE\" or \"FALSE\", default \"FALSE\". Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders.	@param newOrderRespType `ACK`,`RESULT`, default `ACK`	@param priceMatch only available for `LIMIT`/`STOP`/`TAKE_PROFIT` order; can be set to `OPPONENT`/ `OPPONENT_5`/ `OPPONENT_10`/ `OPPONENT_20`: /`QUEUE`/ `QUEUE_5`/ `QUEUE_10`/ `QUEUE_20`; Can't be passed together with `price`	@param selfTradePreventionMode `NONE`: No STP / `EXPIRE_TAKER`:expire taker order when STP triggers/ `EXPIRE_MAKER`:expire taker order when STP triggers/ `EXPIRE_BOTH`:expire both orders when STP triggers; default `NONE`	@param recvWindow
@return ApiNewOrderRequest
*/
func (a *TradeAPIService) NewOrder() ApiNewOrderRequest {
	return ApiNewOrderRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return NewOrderResponse
func (a *TradeAPIService) NewOrderExecute(r ApiNewOrderRequest) (chan *common.ResponseOrRaw[models.NewOrderResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.side == nil {
		return nil, nil, common.ReportError("side is required and must be specified")
	}
	localVarQueryParams["side"] = *r.side

	if r.type_ == nil {
		return nil, nil, common.ReportError("type_ is required and must be specified")
	}
	localVarQueryParams["type_"] = *r.type_

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.positionSide != nil {
		localVarQueryParams["positionSide"] = *r.positionSide
	}
	if r.timeInForce != nil {
		localVarQueryParams["timeInForce"] = *r.timeInForce
	}
	if r.quantity != nil {
		localVarQueryParams["quantity"] = *r.quantity
	}
	if r.reduceOnly != nil {
		localVarQueryParams["reduceOnly"] = *r.reduceOnly
	}
	if r.price != nil {
		localVarQueryParams["price"] = *r.price
	}
	if r.newClientOrderId != nil {
		localVarQueryParams["newClientOrderId"] = *r.newClientOrderId
	}
	if r.stopPrice != nil {
		localVarQueryParams["stopPrice"] = *r.stopPrice
	}
	if r.closePosition != nil {
		localVarQueryParams["closePosition"] = *r.closePosition
	}
	if r.activationPrice != nil {
		localVarQueryParams["activationPrice"] = *r.activationPrice
	}
	if r.callbackRate != nil {
		localVarQueryParams["callbackRate"] = *r.callbackRate
	}
	if r.workingType != nil {
		localVarQueryParams["workingType"] = *r.workingType
	}
	if r.priceProtect != nil {
		localVarQueryParams["priceProtect"] = *r.priceProtect
	}
	if r.newOrderRespType != nil {
		localVarQueryParams["newOrderRespType"] = *r.newOrderRespType
	}
	if r.priceMatch != nil {
		localVarQueryParams["priceMatch"] = *r.priceMatch
	}
	if r.selfTradePreventionMode != nil {
		localVarQueryParams["selfTradePreventionMode"] = *r.selfTradePreventionMode
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/order.place"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.NewOrderResponse](a.Ws, localPayload, sendParams)
}

type ApiPositionInformationRequest struct {
	ApiService  *TradeAPIService
	id          *string
	marginAsset *string
	pair        *string
	recvWindow  *int64
}

// Unique WebSocket request ID.
func (r ApiPositionInformationRequest) Id(id string) ApiPositionInformationRequest {
	r.id = &id
	return r
}

func (r ApiPositionInformationRequest) MarginAsset(marginAsset string) ApiPositionInformationRequest {
	r.marginAsset = &marginAsset
	return r
}

func (r ApiPositionInformationRequest) Pair(pair string) ApiPositionInformationRequest {
	r.pair = &pair
	return r
}

func (r ApiPositionInformationRequest) RecvWindow(recvWindow int64) ApiPositionInformationRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiPositionInformationRequest) Execute() (*common.ResponseOrRaw[models.PositionInformationResponse], error) {
	respChan, errChan, err := r.ApiService.PositionInformationExecute(r)
	if err != nil {
		return nil, err
	}

	select {
	case resp := <-respChan:
		return resp, nil
	case err := <-errChan:
		return nil, err
	}
}

func (r ApiPositionInformationRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.PositionInformationResponse], chan error, error) {
	return r.ApiService.PositionInformationExecute(r)
}

/*
PositionInformation Position Information(USER_DATA)
/account.position

https://developers.binance.com/docs/derivatives/coin-margined-futures/trade/websocket-api/Position-Information

@param id Unique WebSocket request ID.	@param marginAsset	@param pair	@param recvWindow
@return ApiPositionInformationRequest
*/
func (a *TradeAPIService) PositionInformation() ApiPositionInformationRequest {
	return ApiPositionInformationRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return PositionInformationResponse
func (a *TradeAPIService) PositionInformationExecute(r ApiPositionInformationRequest) (chan *common.ResponseOrRaw[models.PositionInformationResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.marginAsset != nil {
		localVarQueryParams["marginAsset"] = *r.marginAsset
	}
	if r.pair != nil {
		localVarQueryParams["pair"] = *r.pair
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/account.position"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.PositionInformationResponse](a.Ws, localPayload, sendParams)
}

type ApiQueryOrderRequest struct {
	ApiService        *TradeAPIService
	symbol            *string
	id                *string
	orderId           *int64
	origClientOrderId *string
	recvWindow        *int64
}

func (r ApiQueryOrderRequest) Symbol(symbol string) ApiQueryOrderRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiQueryOrderRequest) Id(id string) ApiQueryOrderRequest {
	r.id = &id
	return r
}

func (r ApiQueryOrderRequest) OrderId(orderId int64) ApiQueryOrderRequest {
	r.orderId = &orderId
	return r
}

func (r ApiQueryOrderRequest) OrigClientOrderId(origClientOrderId string) ApiQueryOrderRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

func (r ApiQueryOrderRequest) RecvWindow(recvWindow int64) ApiQueryOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryOrderRequest) Execute() (*common.ResponseOrRaw[models.QueryOrderResponse], error) {
	respChan, errChan, err := r.ApiService.QueryOrderExecute(r)
	if err != nil {
		return nil, err
	}

	select {
	case resp := <-respChan:
		return resp, nil
	case err := <-errChan:
		return nil, err
	}
}

func (r ApiQueryOrderRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.QueryOrderResponse], chan error, error) {
	return r.ApiService.QueryOrderExecute(r)
}

/*
QueryOrder Query Order (USER_DATA)
/order.status

https://developers.binance.com/docs/derivatives/coin-margined-futures/trade/websocket-api/Query-Order

@param symbol	@param id Unique WebSocket request ID.	@param orderId	@param origClientOrderId	@param recvWindow
@return ApiQueryOrderRequest
*/
func (a *TradeAPIService) QueryOrder() ApiQueryOrderRequest {
	return ApiQueryOrderRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return QueryOrderResponse
func (a *TradeAPIService) QueryOrderExecute(r ApiQueryOrderRequest) (chan *common.ResponseOrRaw[models.QueryOrderResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.orderId != nil {
		localVarQueryParams["orderId"] = *r.orderId
	}
	if r.origClientOrderId != nil {
		localVarQueryParams["origClientOrderId"] = *r.origClientOrderId
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/order.status"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.QueryOrderResponse](a.Ws, localPayload, sendParams)
}
