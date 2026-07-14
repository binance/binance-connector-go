/*
Spot WebSocket API

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the OrderCancelReplaceResponseResultCancelResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderCancelReplaceResponseResultCancelResponse{}

// OrderCancelReplaceResponseResultCancelResponse struct for OrderCancelReplaceResponseResultCancelResponse
type OrderCancelReplaceResponseResultCancelResponse struct {
	Symbol                  *string `json:"symbol,omitempty"`
	OrigClientOrderId       *string `json:"origClientOrderId,omitempty"`
	OrderId                 *int64  `json:"orderId,omitempty"`
	OrderListId             *int64  `json:"orderListId,omitempty"`
	ClientOrderId           *string `json:"clientOrderId,omitempty"`
	TransactTime            *int64  `json:"transactTime,omitempty"`
	Price                   *string `json:"price,omitempty"`
	OrigQty                 *string `json:"origQty,omitempty"`
	ExecutedQty             *string `json:"executedQty,omitempty"`
	OrigQuoteOrderQty       *string `json:"origQuoteOrderQty,omitempty"`
	CummulativeQuoteQty     *string `json:"cummulativeQuoteQty,omitempty"`
	Status                  *string `json:"status,omitempty"`
	TimeInForce             *string `json:"timeInForce,omitempty"`
	Type                    *string `json:"type,omitempty"`
	Side                    *string `json:"side,omitempty"`
	SelfTradePreventionMode *string `json:"selfTradePreventionMode,omitempty"`
	// Appears only if the parameter icebergQty was sent in the request.
	IcebergQty *string `json:"icebergQty,omitempty"`
	// Appears only if the order expired due to STP.
	PreventedMatchId *int64 `json:"preventedMatchId,omitempty"`
	// Order quantity that expired due to STP.
	PreventedQuantity *string `json:"preventedQuantity,omitempty"`
	// Appears for STOP_LOSS, TAKE_PROFIT, STOP_LOSS_LIMIT, and TAKE_PROFIT_LIMIT orders.
	StopPrice *string `json:"stopPrice,omitempty"`
	// Appears only if the strategyId parameter was provided upon order placement.
	StrategyId *int64 `json:"strategyId,omitempty"`
	// Appears only if the strategyType parameter was provided upon order placement.
	StrategyType *int64 `json:"strategyType,omitempty"`
	// Delta price change required before order activation.
	TrailingDelta *int64 `json:"trailingDelta,omitempty"`
	// Time when the trailing order is now active and tracking price changes.
	TrailingTime *int64 `json:"trailingTime,omitempty"`
	// Field that determines whether order used SOR.
	UsedSor *bool `json:"usedSor,omitempty"`
	// Determines whether the order is being filled by the SOR or by the order book.
	WorkingFloor *string `json:"workingFloor,omitempty"`
	// Price peg type. Only for pegged orders.
	PegPriceType *string `json:"pegPriceType,omitempty"`
	// Price peg offset type. Only for pegged orders, if requested.
	PegOffsetType *string `json:"pegOffsetType,omitempty"`
	// Price peg offset value. Only for pegged orders, if requested.
	PegOffsetValue *int64 `json:"pegOffsetValue,omitempty"`
	// Current price order is pegged at. Only for pegged orders, once determined.
	PeggedPrice *string `json:"peggedPrice,omitempty"`
	// Cause of the order's expiration. Appears when an order has expired.
	ExpiryReason         *string `json:"expiryReason,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderCancelReplaceResponseResultCancelResponse OrderCancelReplaceResponseResultCancelResponse

// NewOrderCancelReplaceResponseResultCancelResponse instantiates a new OrderCancelReplaceResponseResultCancelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCancelReplaceResponseResultCancelResponse() *OrderCancelReplaceResponseResultCancelResponse {
	this := OrderCancelReplaceResponseResultCancelResponse{}
	return &this
}

// NewOrderCancelReplaceResponseResultCancelResponseWithDefaults instantiates a new OrderCancelReplaceResponseResultCancelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCancelReplaceResponseResultCancelResponseWithDefaults() *OrderCancelReplaceResponseResultCancelResponse {
	this := OrderCancelReplaceResponseResultCancelResponse{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *OrderCancelReplaceResponseResultCancelResponse) SetSymbol(v string) {
	o.Symbol = &v
}

// GetOrigClientOrderId returns the OrigClientOrderId field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetOrigClientOrderId() string {
	if o == nil || common.IsNil(o.OrigClientOrderId) {
		var ret string
		return ret
	}
	return *o.OrigClientOrderId
}

// GetOrigClientOrderIdOk returns a tuple with the OrigClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetOrigClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrigClientOrderId) {
		return nil, false
	}
	return o.OrigClientOrderId, true
}

// HasOrigClientOrderId returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) HasOrigClientOrderId() bool {
	if o != nil && !common.IsNil(o.OrigClientOrderId) {
		return true
	}

	return false
}

// SetOrigClientOrderId gets a reference to the given string and assigns it to the OrigClientOrderId field.
func (o *OrderCancelReplaceResponseResultCancelResponse) SetOrigClientOrderId(v string) {
	o.OrigClientOrderId = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *OrderCancelReplaceResponseResultCancelResponse) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetOrderListId returns the OrderListId field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetOrderListId() int64 {
	if o == nil || common.IsNil(o.OrderListId) {
		var ret int64
		return ret
	}
	return *o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetOrderListIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderListId) {
		return nil, false
	}
	return o.OrderListId, true
}

// HasOrderListId returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) HasOrderListId() bool {
	if o != nil && !common.IsNil(o.OrderListId) {
		return true
	}

	return false
}

// SetOrderListId gets a reference to the given int64 and assigns it to the OrderListId field.
func (o *OrderCancelReplaceResponseResultCancelResponse) SetOrderListId(v int64) {
	o.OrderListId = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *OrderCancelReplaceResponseResultCancelResponse) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetTransactTime returns the TransactTime field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetTransactTime() int64 {
	if o == nil || common.IsNil(o.TransactTime) {
		var ret int64
		return ret
	}
	return *o.TransactTime
}

// GetTransactTimeOk returns a tuple with the TransactTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetTransactTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TransactTime) {
		return nil, false
	}
	return o.TransactTime, true
}

// HasTransactTime returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) HasTransactTime() bool {
	if o != nil && !common.IsNil(o.TransactTime) {
		return true
	}

	return false
}

// SetTransactTime gets a reference to the given int64 and assigns it to the TransactTime field.
func (o *OrderCancelReplaceResponseResultCancelResponse) SetTransactTime(v int64) {
	o.TransactTime = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *OrderCancelReplaceResponseResultCancelResponse) SetPrice(v string) {
	o.Price = &v
}

// GetOrigQty returns the OrigQty field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetOrigQty() string {
	if o == nil || common.IsNil(o.OrigQty) {
		var ret string
		return ret
	}
	return *o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetOrigQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrigQty) {
		return nil, false
	}
	return o.OrigQty, true
}

// HasOrigQty returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) HasOrigQty() bool {
	if o != nil && !common.IsNil(o.OrigQty) {
		return true
	}

	return false
}

// SetOrigQty gets a reference to the given string and assigns it to the OrigQty field.
func (o *OrderCancelReplaceResponseResultCancelResponse) SetOrigQty(v string) {
	o.OrigQty = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetExecutedQty() string {
	if o == nil || common.IsNil(o.ExecutedQty) {
		var ret string
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetExecutedQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) HasExecutedQty() bool {
	if o != nil && !common.IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given string and assigns it to the ExecutedQty field.
func (o *OrderCancelReplaceResponseResultCancelResponse) SetExecutedQty(v string) {
	o.ExecutedQty = &v
}

// GetOrigQuoteOrderQty returns the OrigQuoteOrderQty field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetOrigQuoteOrderQty() string {
	if o == nil || common.IsNil(o.OrigQuoteOrderQty) {
		var ret string
		return ret
	}
	return *o.OrigQuoteOrderQty
}

// GetOrigQuoteOrderQtyOk returns a tuple with the OrigQuoteOrderQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetOrigQuoteOrderQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrigQuoteOrderQty) {
		return nil, false
	}
	return o.OrigQuoteOrderQty, true
}

// HasOrigQuoteOrderQty returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) HasOrigQuoteOrderQty() bool {
	if o != nil && !common.IsNil(o.OrigQuoteOrderQty) {
		return true
	}

	return false
}

// SetOrigQuoteOrderQty gets a reference to the given string and assigns it to the OrigQuoteOrderQty field.
func (o *OrderCancelReplaceResponseResultCancelResponse) SetOrigQuoteOrderQty(v string) {
	o.OrigQuoteOrderQty = &v
}

// GetCummulativeQuoteQty returns the CummulativeQuoteQty field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetCummulativeQuoteQty() string {
	if o == nil || common.IsNil(o.CummulativeQuoteQty) {
		var ret string
		return ret
	}
	return *o.CummulativeQuoteQty
}

// GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetCummulativeQuoteQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.CummulativeQuoteQty) {
		return nil, false
	}
	return o.CummulativeQuoteQty, true
}

// HasCummulativeQuoteQty returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) HasCummulativeQuoteQty() bool {
	if o != nil && !common.IsNil(o.CummulativeQuoteQty) {
		return true
	}

	return false
}

// SetCummulativeQuoteQty gets a reference to the given string and assigns it to the CummulativeQuoteQty field.
func (o *OrderCancelReplaceResponseResultCancelResponse) SetCummulativeQuoteQty(v string) {
	o.CummulativeQuoteQty = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OrderCancelReplaceResponseResultCancelResponse) SetStatus(v string) {
	o.Status = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetTimeInForce() string {
	if o == nil || common.IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetTimeInForceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) HasTimeInForce() bool {
	if o != nil && !common.IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *OrderCancelReplaceResponseResultCancelResponse) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OrderCancelReplaceResponseResultCancelResponse) SetType(v string) {
	o.Type = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *OrderCancelReplaceResponseResultCancelResponse) SetSide(v string) {
	o.Side = &v
}

// GetSelfTradePreventionMode returns the SelfTradePreventionMode field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetSelfTradePreventionMode() string {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		var ret string
		return ret
	}
	return *o.SelfTradePreventionMode
}

// GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetSelfTradePreventionModeOk() (*string, bool) {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		return nil, false
	}
	return o.SelfTradePreventionMode, true
}

// HasSelfTradePreventionMode returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) HasSelfTradePreventionMode() bool {
	if o != nil && !common.IsNil(o.SelfTradePreventionMode) {
		return true
	}

	return false
}

// SetSelfTradePreventionMode gets a reference to the given string and assigns it to the SelfTradePreventionMode field.
func (o *OrderCancelReplaceResponseResultCancelResponse) SetSelfTradePreventionMode(v string) {
	o.SelfTradePreventionMode = &v
}

// GetIcebergQty returns the IcebergQty field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetIcebergQty() string {
	if o == nil || common.IsNil(o.IcebergQty) {
		var ret string
		return ret
	}
	return *o.IcebergQty
}

// GetIcebergQtyOk returns a tuple with the IcebergQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetIcebergQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.IcebergQty) {
		return nil, false
	}
	return o.IcebergQty, true
}

// HasIcebergQty returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) HasIcebergQty() bool {
	if o != nil && !common.IsNil(o.IcebergQty) {
		return true
	}

	return false
}

// SetIcebergQty gets a reference to the given string and assigns it to the IcebergQty field.
func (o *OrderCancelReplaceResponseResultCancelResponse) SetIcebergQty(v string) {
	o.IcebergQty = &v
}

// GetPreventedMatchId returns the PreventedMatchId field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetPreventedMatchId() int64 {
	if o == nil || common.IsNil(o.PreventedMatchId) {
		var ret int64
		return ret
	}
	return *o.PreventedMatchId
}

// GetPreventedMatchIdOk returns a tuple with the PreventedMatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetPreventedMatchIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PreventedMatchId) {
		return nil, false
	}
	return o.PreventedMatchId, true
}

// HasPreventedMatchId returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) HasPreventedMatchId() bool {
	if o != nil && !common.IsNil(o.PreventedMatchId) {
		return true
	}

	return false
}

// SetPreventedMatchId gets a reference to the given int64 and assigns it to the PreventedMatchId field.
func (o *OrderCancelReplaceResponseResultCancelResponse) SetPreventedMatchId(v int64) {
	o.PreventedMatchId = &v
}

// GetPreventedQuantity returns the PreventedQuantity field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetPreventedQuantity() string {
	if o == nil || common.IsNil(o.PreventedQuantity) {
		var ret string
		return ret
	}
	return *o.PreventedQuantity
}

// GetPreventedQuantityOk returns a tuple with the PreventedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetPreventedQuantityOk() (*string, bool) {
	if o == nil || common.IsNil(o.PreventedQuantity) {
		return nil, false
	}
	return o.PreventedQuantity, true
}

// HasPreventedQuantity returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) HasPreventedQuantity() bool {
	if o != nil && !common.IsNil(o.PreventedQuantity) {
		return true
	}

	return false
}

// SetPreventedQuantity gets a reference to the given string and assigns it to the PreventedQuantity field.
func (o *OrderCancelReplaceResponseResultCancelResponse) SetPreventedQuantity(v string) {
	o.PreventedQuantity = &v
}

// GetStopPrice returns the StopPrice field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetStopPrice() string {
	if o == nil || common.IsNil(o.StopPrice) {
		var ret string
		return ret
	}
	return *o.StopPrice
}

// GetStopPriceOk returns a tuple with the StopPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetStopPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.StopPrice) {
		return nil, false
	}
	return o.StopPrice, true
}

// HasStopPrice returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) HasStopPrice() bool {
	if o != nil && !common.IsNil(o.StopPrice) {
		return true
	}

	return false
}

// SetStopPrice gets a reference to the given string and assigns it to the StopPrice field.
func (o *OrderCancelReplaceResponseResultCancelResponse) SetStopPrice(v string) {
	o.StopPrice = &v
}

// GetStrategyId returns the StrategyId field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetStrategyId() int64 {
	if o == nil || common.IsNil(o.StrategyId) {
		var ret int64
		return ret
	}
	return *o.StrategyId
}

// GetStrategyIdOk returns a tuple with the StrategyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetStrategyIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.StrategyId) {
		return nil, false
	}
	return o.StrategyId, true
}

// HasStrategyId returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) HasStrategyId() bool {
	if o != nil && !common.IsNil(o.StrategyId) {
		return true
	}

	return false
}

// SetStrategyId gets a reference to the given int64 and assigns it to the StrategyId field.
func (o *OrderCancelReplaceResponseResultCancelResponse) SetStrategyId(v int64) {
	o.StrategyId = &v
}

// GetStrategyType returns the StrategyType field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetStrategyType() int64 {
	if o == nil || common.IsNil(o.StrategyType) {
		var ret int64
		return ret
	}
	return *o.StrategyType
}

// GetStrategyTypeOk returns a tuple with the StrategyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetStrategyTypeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.StrategyType) {
		return nil, false
	}
	return o.StrategyType, true
}

// HasStrategyType returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) HasStrategyType() bool {
	if o != nil && !common.IsNil(o.StrategyType) {
		return true
	}

	return false
}

// SetStrategyType gets a reference to the given int64 and assigns it to the StrategyType field.
func (o *OrderCancelReplaceResponseResultCancelResponse) SetStrategyType(v int64) {
	o.StrategyType = &v
}

// GetTrailingDelta returns the TrailingDelta field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetTrailingDelta() int64 {
	if o == nil || common.IsNil(o.TrailingDelta) {
		var ret int64
		return ret
	}
	return *o.TrailingDelta
}

// GetTrailingDeltaOk returns a tuple with the TrailingDelta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetTrailingDeltaOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TrailingDelta) {
		return nil, false
	}
	return o.TrailingDelta, true
}

// HasTrailingDelta returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) HasTrailingDelta() bool {
	if o != nil && !common.IsNil(o.TrailingDelta) {
		return true
	}

	return false
}

// SetTrailingDelta gets a reference to the given int64 and assigns it to the TrailingDelta field.
func (o *OrderCancelReplaceResponseResultCancelResponse) SetTrailingDelta(v int64) {
	o.TrailingDelta = &v
}

// GetTrailingTime returns the TrailingTime field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetTrailingTime() int64 {
	if o == nil || common.IsNil(o.TrailingTime) {
		var ret int64
		return ret
	}
	return *o.TrailingTime
}

// GetTrailingTimeOk returns a tuple with the TrailingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetTrailingTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TrailingTime) {
		return nil, false
	}
	return o.TrailingTime, true
}

// HasTrailingTime returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) HasTrailingTime() bool {
	if o != nil && !common.IsNil(o.TrailingTime) {
		return true
	}

	return false
}

// SetTrailingTime gets a reference to the given int64 and assigns it to the TrailingTime field.
func (o *OrderCancelReplaceResponseResultCancelResponse) SetTrailingTime(v int64) {
	o.TrailingTime = &v
}

// GetUsedSor returns the UsedSor field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetUsedSor() bool {
	if o == nil || common.IsNil(o.UsedSor) {
		var ret bool
		return ret
	}
	return *o.UsedSor
}

// GetUsedSorOk returns a tuple with the UsedSor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetUsedSorOk() (*bool, bool) {
	if o == nil || common.IsNil(o.UsedSor) {
		return nil, false
	}
	return o.UsedSor, true
}

// HasUsedSor returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) HasUsedSor() bool {
	if o != nil && !common.IsNil(o.UsedSor) {
		return true
	}

	return false
}

// SetUsedSor gets a reference to the given bool and assigns it to the UsedSor field.
func (o *OrderCancelReplaceResponseResultCancelResponse) SetUsedSor(v bool) {
	o.UsedSor = &v
}

// GetWorkingFloor returns the WorkingFloor field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetWorkingFloor() string {
	if o == nil || common.IsNil(o.WorkingFloor) {
		var ret string
		return ret
	}
	return *o.WorkingFloor
}

// GetWorkingFloorOk returns a tuple with the WorkingFloor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetWorkingFloorOk() (*string, bool) {
	if o == nil || common.IsNil(o.WorkingFloor) {
		return nil, false
	}
	return o.WorkingFloor, true
}

// HasWorkingFloor returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) HasWorkingFloor() bool {
	if o != nil && !common.IsNil(o.WorkingFloor) {
		return true
	}

	return false
}

// SetWorkingFloor gets a reference to the given string and assigns it to the WorkingFloor field.
func (o *OrderCancelReplaceResponseResultCancelResponse) SetWorkingFloor(v string) {
	o.WorkingFloor = &v
}

// GetPegPriceType returns the PegPriceType field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetPegPriceType() string {
	if o == nil || common.IsNil(o.PegPriceType) {
		var ret string
		return ret
	}
	return *o.PegPriceType
}

// GetPegPriceTypeOk returns a tuple with the PegPriceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetPegPriceTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PegPriceType) {
		return nil, false
	}
	return o.PegPriceType, true
}

// HasPegPriceType returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) HasPegPriceType() bool {
	if o != nil && !common.IsNil(o.PegPriceType) {
		return true
	}

	return false
}

// SetPegPriceType gets a reference to the given string and assigns it to the PegPriceType field.
func (o *OrderCancelReplaceResponseResultCancelResponse) SetPegPriceType(v string) {
	o.PegPriceType = &v
}

// GetPegOffsetType returns the PegOffsetType field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetPegOffsetType() string {
	if o == nil || common.IsNil(o.PegOffsetType) {
		var ret string
		return ret
	}
	return *o.PegOffsetType
}

// GetPegOffsetTypeOk returns a tuple with the PegOffsetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetPegOffsetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PegOffsetType) {
		return nil, false
	}
	return o.PegOffsetType, true
}

// HasPegOffsetType returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) HasPegOffsetType() bool {
	if o != nil && !common.IsNil(o.PegOffsetType) {
		return true
	}

	return false
}

// SetPegOffsetType gets a reference to the given string and assigns it to the PegOffsetType field.
func (o *OrderCancelReplaceResponseResultCancelResponse) SetPegOffsetType(v string) {
	o.PegOffsetType = &v
}

// GetPegOffsetValue returns the PegOffsetValue field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetPegOffsetValue() int64 {
	if o == nil || common.IsNil(o.PegOffsetValue) {
		var ret int64
		return ret
	}
	return *o.PegOffsetValue
}

// GetPegOffsetValueOk returns a tuple with the PegOffsetValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetPegOffsetValueOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PegOffsetValue) {
		return nil, false
	}
	return o.PegOffsetValue, true
}

// HasPegOffsetValue returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) HasPegOffsetValue() bool {
	if o != nil && !common.IsNil(o.PegOffsetValue) {
		return true
	}

	return false
}

// SetPegOffsetValue gets a reference to the given int64 and assigns it to the PegOffsetValue field.
func (o *OrderCancelReplaceResponseResultCancelResponse) SetPegOffsetValue(v int64) {
	o.PegOffsetValue = &v
}

// GetPeggedPrice returns the PeggedPrice field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetPeggedPrice() string {
	if o == nil || common.IsNil(o.PeggedPrice) {
		var ret string
		return ret
	}
	return *o.PeggedPrice
}

// GetPeggedPriceOk returns a tuple with the PeggedPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetPeggedPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.PeggedPrice) {
		return nil, false
	}
	return o.PeggedPrice, true
}

// HasPeggedPrice returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) HasPeggedPrice() bool {
	if o != nil && !common.IsNil(o.PeggedPrice) {
		return true
	}

	return false
}

// SetPeggedPrice gets a reference to the given string and assigns it to the PeggedPrice field.
func (o *OrderCancelReplaceResponseResultCancelResponse) SetPeggedPrice(v string) {
	o.PeggedPrice = &v
}

// GetExpiryReason returns the ExpiryReason field value if set, zero value otherwise.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetExpiryReason() string {
	if o == nil || common.IsNil(o.ExpiryReason) {
		var ret string
		return ret
	}
	return *o.ExpiryReason
}

// GetExpiryReasonOk returns a tuple with the ExpiryReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) GetExpiryReasonOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExpiryReason) {
		return nil, false
	}
	return o.ExpiryReason, true
}

// HasExpiryReason returns a boolean if a field has been set.
func (o *OrderCancelReplaceResponseResultCancelResponse) HasExpiryReason() bool {
	if o != nil && !common.IsNil(o.ExpiryReason) {
		return true
	}

	return false
}

// SetExpiryReason gets a reference to the given string and assigns it to the ExpiryReason field.
func (o *OrderCancelReplaceResponseResultCancelResponse) SetExpiryReason(v string) {
	o.ExpiryReason = &v
}

func (o OrderCancelReplaceResponseResultCancelResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCancelReplaceResponseResultCancelResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.OrigClientOrderId) {
		toSerialize["origClientOrderId"] = o.OrigClientOrderId
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.OrderListId) {
		toSerialize["orderListId"] = o.OrderListId
	}
	if !common.IsNil(o.ClientOrderId) {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}
	if !common.IsNil(o.TransactTime) {
		toSerialize["transactTime"] = o.TransactTime
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.OrigQty) {
		toSerialize["origQty"] = o.OrigQty
	}
	if !common.IsNil(o.ExecutedQty) {
		toSerialize["executedQty"] = o.ExecutedQty
	}
	if !common.IsNil(o.OrigQuoteOrderQty) {
		toSerialize["origQuoteOrderQty"] = o.OrigQuoteOrderQty
	}
	if !common.IsNil(o.CummulativeQuoteQty) {
		toSerialize["cummulativeQuoteQty"] = o.CummulativeQuoteQty
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.TimeInForce) {
		toSerialize["timeInForce"] = o.TimeInForce
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !common.IsNil(o.SelfTradePreventionMode) {
		toSerialize["selfTradePreventionMode"] = o.SelfTradePreventionMode
	}
	if !common.IsNil(o.IcebergQty) {
		toSerialize["icebergQty"] = o.IcebergQty
	}
	if !common.IsNil(o.PreventedMatchId) {
		toSerialize["preventedMatchId"] = o.PreventedMatchId
	}
	if !common.IsNil(o.PreventedQuantity) {
		toSerialize["preventedQuantity"] = o.PreventedQuantity
	}
	if !common.IsNil(o.StopPrice) {
		toSerialize["stopPrice"] = o.StopPrice
	}
	if !common.IsNil(o.StrategyId) {
		toSerialize["strategyId"] = o.StrategyId
	}
	if !common.IsNil(o.StrategyType) {
		toSerialize["strategyType"] = o.StrategyType
	}
	if !common.IsNil(o.TrailingDelta) {
		toSerialize["trailingDelta"] = o.TrailingDelta
	}
	if !common.IsNil(o.TrailingTime) {
		toSerialize["trailingTime"] = o.TrailingTime
	}
	if !common.IsNil(o.UsedSor) {
		toSerialize["usedSor"] = o.UsedSor
	}
	if !common.IsNil(o.WorkingFloor) {
		toSerialize["workingFloor"] = o.WorkingFloor
	}
	if !common.IsNil(o.PegPriceType) {
		toSerialize["pegPriceType"] = o.PegPriceType
	}
	if !common.IsNil(o.PegOffsetType) {
		toSerialize["pegOffsetType"] = o.PegOffsetType
	}
	if !common.IsNil(o.PegOffsetValue) {
		toSerialize["pegOffsetValue"] = o.PegOffsetValue
	}
	if !common.IsNil(o.PeggedPrice) {
		toSerialize["peggedPrice"] = o.PeggedPrice
	}
	if !common.IsNil(o.ExpiryReason) {
		toSerialize["expiryReason"] = o.ExpiryReason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderCancelReplaceResponseResultCancelResponse) UnmarshalJSON(data []byte) (err error) {
	varOrderCancelReplaceResponseResultCancelResponse := _OrderCancelReplaceResponseResultCancelResponse{}

	err = json.Unmarshal(data, &varOrderCancelReplaceResponseResultCancelResponse)

	if err != nil {
		return err
	}

	*o = OrderCancelReplaceResponseResultCancelResponse(varOrderCancelReplaceResponseResultCancelResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "origClientOrderId")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "orderListId")
		delete(additionalProperties, "clientOrderId")
		delete(additionalProperties, "transactTime")
		delete(additionalProperties, "price")
		delete(additionalProperties, "origQty")
		delete(additionalProperties, "executedQty")
		delete(additionalProperties, "origQuoteOrderQty")
		delete(additionalProperties, "cummulativeQuoteQty")
		delete(additionalProperties, "status")
		delete(additionalProperties, "timeInForce")
		delete(additionalProperties, "type")
		delete(additionalProperties, "side")
		delete(additionalProperties, "selfTradePreventionMode")
		delete(additionalProperties, "icebergQty")
		delete(additionalProperties, "preventedMatchId")
		delete(additionalProperties, "preventedQuantity")
		delete(additionalProperties, "stopPrice")
		delete(additionalProperties, "strategyId")
		delete(additionalProperties, "strategyType")
		delete(additionalProperties, "trailingDelta")
		delete(additionalProperties, "trailingTime")
		delete(additionalProperties, "usedSor")
		delete(additionalProperties, "workingFloor")
		delete(additionalProperties, "pegPriceType")
		delete(additionalProperties, "pegOffsetType")
		delete(additionalProperties, "pegOffsetValue")
		delete(additionalProperties, "peggedPrice")
		delete(additionalProperties, "expiryReason")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderCancelReplaceResponseResultCancelResponse struct {
	value *OrderCancelReplaceResponseResultCancelResponse
	isSet bool
}

func (v NullableOrderCancelReplaceResponseResultCancelResponse) Get() *OrderCancelReplaceResponseResultCancelResponse {
	return v.value
}

func (v *NullableOrderCancelReplaceResponseResultCancelResponse) Set(val *OrderCancelReplaceResponseResultCancelResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCancelReplaceResponseResultCancelResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCancelReplaceResponseResultCancelResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCancelReplaceResponseResultCancelResponse(val *OrderCancelReplaceResponseResultCancelResponse) *NullableOrderCancelReplaceResponseResultCancelResponse {
	return &NullableOrderCancelReplaceResponseResultCancelResponse{value: val, isSet: true}
}

func (v NullableOrderCancelReplaceResponseResultCancelResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCancelReplaceResponseResultCancelResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
