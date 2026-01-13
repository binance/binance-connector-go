/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OrderAmendKeepPriorityResponseResultAmendedOrder type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderAmendKeepPriorityResponseResultAmendedOrder{}

// OrderAmendKeepPriorityResponseResultAmendedOrder struct for OrderAmendKeepPriorityResponseResultAmendedOrder
type OrderAmendKeepPriorityResponseResultAmendedOrder struct {
	Symbol                  *string `json:"symbol,omitempty"`
	OrderId                 *int64  `json:"orderId,omitempty"`
	OrderListId             *int64  `json:"orderListId,omitempty"`
	OrigClientOrderId       *string `json:"origClientOrderId,omitempty"`
	ClientOrderId           *string `json:"clientOrderId,omitempty"`
	Price                   *string `json:"price,omitempty"`
	Qty                     *string `json:"qty,omitempty"`
	ExecutedQty             *string `json:"executedQty,omitempty"`
	PreventedQty            *string `json:"preventedQty,omitempty"`
	QuoteOrderQty           *string `json:"quoteOrderQty,omitempty"`
	CumulativeQuoteQty      *string `json:"cumulativeQuoteQty,omitempty"`
	Status                  *string `json:"status,omitempty"`
	TimeInForce             *string `json:"timeInForce,omitempty"`
	Type                    *string `json:"type,omitempty"`
	Side                    *string `json:"side,omitempty"`
	WorkingTime             *int64  `json:"workingTime,omitempty"`
	SelfTradePreventionMode *string `json:"selfTradePreventionMode,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _OrderAmendKeepPriorityResponseResultAmendedOrder OrderAmendKeepPriorityResponseResultAmendedOrder

// NewOrderAmendKeepPriorityResponseResultAmendedOrder instantiates a new OrderAmendKeepPriorityResponseResultAmendedOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderAmendKeepPriorityResponseResultAmendedOrder() *OrderAmendKeepPriorityResponseResultAmendedOrder {
	this := OrderAmendKeepPriorityResponseResultAmendedOrder{}
	return &this
}

// NewOrderAmendKeepPriorityResponseResultAmendedOrderWithDefaults instantiates a new OrderAmendKeepPriorityResponseResultAmendedOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderAmendKeepPriorityResponseResultAmendedOrderWithDefaults() *OrderAmendKeepPriorityResponseResultAmendedOrder {
	this := OrderAmendKeepPriorityResponseResultAmendedOrder{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetSymbol(v string) {
	o.Symbol = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetOrderListId returns the OrderListId field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetOrderListId() int64 {
	if o == nil || common.IsNil(o.OrderListId) {
		var ret int64
		return ret
	}
	return *o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetOrderListIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderListId) {
		return nil, false
	}
	return o.OrderListId, true
}

// HasOrderListId returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasOrderListId() bool {
	if o != nil && !common.IsNil(o.OrderListId) {
		return true
	}

	return false
}

// SetOrderListId gets a reference to the given int64 and assigns it to the OrderListId field.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetOrderListId(v int64) {
	o.OrderListId = &v
}

// GetOrigClientOrderId returns the OrigClientOrderId field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetOrigClientOrderId() string {
	if o == nil || common.IsNil(o.OrigClientOrderId) {
		var ret string
		return ret
	}
	return *o.OrigClientOrderId
}

// GetOrigClientOrderIdOk returns a tuple with the OrigClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetOrigClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrigClientOrderId) {
		return nil, false
	}
	return o.OrigClientOrderId, true
}

// HasOrigClientOrderId returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasOrigClientOrderId() bool {
	if o != nil && !common.IsNil(o.OrigClientOrderId) {
		return true
	}

	return false
}

// SetOrigClientOrderId gets a reference to the given string and assigns it to the OrigClientOrderId field.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetOrigClientOrderId(v string) {
	o.OrigClientOrderId = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetPrice(v string) {
	o.Price = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetQty() string {
	if o == nil || common.IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasQty() bool {
	if o != nil && !common.IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetQty(v string) {
	o.Qty = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetExecutedQty() string {
	if o == nil || common.IsNil(o.ExecutedQty) {
		var ret string
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetExecutedQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasExecutedQty() bool {
	if o != nil && !common.IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given string and assigns it to the ExecutedQty field.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetExecutedQty(v string) {
	o.ExecutedQty = &v
}

// GetPreventedQty returns the PreventedQty field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetPreventedQty() string {
	if o == nil || common.IsNil(o.PreventedQty) {
		var ret string
		return ret
	}
	return *o.PreventedQty
}

// GetPreventedQtyOk returns a tuple with the PreventedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetPreventedQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.PreventedQty) {
		return nil, false
	}
	return o.PreventedQty, true
}

// HasPreventedQty returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasPreventedQty() bool {
	if o != nil && !common.IsNil(o.PreventedQty) {
		return true
	}

	return false
}

// SetPreventedQty gets a reference to the given string and assigns it to the PreventedQty field.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetPreventedQty(v string) {
	o.PreventedQty = &v
}

// GetQuoteOrderQty returns the QuoteOrderQty field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetQuoteOrderQty() string {
	if o == nil || common.IsNil(o.QuoteOrderQty) {
		var ret string
		return ret
	}
	return *o.QuoteOrderQty
}

// GetQuoteOrderQtyOk returns a tuple with the QuoteOrderQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetQuoteOrderQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteOrderQty) {
		return nil, false
	}
	return o.QuoteOrderQty, true
}

// HasQuoteOrderQty returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasQuoteOrderQty() bool {
	if o != nil && !common.IsNil(o.QuoteOrderQty) {
		return true
	}

	return false
}

// SetQuoteOrderQty gets a reference to the given string and assigns it to the QuoteOrderQty field.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetQuoteOrderQty(v string) {
	o.QuoteOrderQty = &v
}

// GetCumulativeQuoteQty returns the CumulativeQuoteQty field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetCumulativeQuoteQty() string {
	if o == nil || common.IsNil(o.CumulativeQuoteQty) {
		var ret string
		return ret
	}
	return *o.CumulativeQuoteQty
}

// GetCumulativeQuoteQtyOk returns a tuple with the CumulativeQuoteQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetCumulativeQuoteQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.CumulativeQuoteQty) {
		return nil, false
	}
	return o.CumulativeQuoteQty, true
}

// HasCumulativeQuoteQty returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasCumulativeQuoteQty() bool {
	if o != nil && !common.IsNil(o.CumulativeQuoteQty) {
		return true
	}

	return false
}

// SetCumulativeQuoteQty gets a reference to the given string and assigns it to the CumulativeQuoteQty field.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetCumulativeQuoteQty(v string) {
	o.CumulativeQuoteQty = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetStatus(v string) {
	o.Status = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetTimeInForce() string {
	if o == nil || common.IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetTimeInForceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasTimeInForce() bool {
	if o != nil && !common.IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetType(v string) {
	o.Type = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetSide(v string) {
	o.Side = &v
}

// GetWorkingTime returns the WorkingTime field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetWorkingTime() int64 {
	if o == nil || common.IsNil(o.WorkingTime) {
		var ret int64
		return ret
	}
	return *o.WorkingTime
}

// GetWorkingTimeOk returns a tuple with the WorkingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetWorkingTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.WorkingTime) {
		return nil, false
	}
	return o.WorkingTime, true
}

// HasWorkingTime returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasWorkingTime() bool {
	if o != nil && !common.IsNil(o.WorkingTime) {
		return true
	}

	return false
}

// SetWorkingTime gets a reference to the given int64 and assigns it to the WorkingTime field.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetWorkingTime(v int64) {
	o.WorkingTime = &v
}

// GetSelfTradePreventionMode returns the SelfTradePreventionMode field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetSelfTradePreventionMode() string {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		var ret string
		return ret
	}
	return *o.SelfTradePreventionMode
}

// GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetSelfTradePreventionModeOk() (*string, bool) {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		return nil, false
	}
	return o.SelfTradePreventionMode, true
}

// HasSelfTradePreventionMode returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasSelfTradePreventionMode() bool {
	if o != nil && !common.IsNil(o.SelfTradePreventionMode) {
		return true
	}

	return false
}

// SetSelfTradePreventionMode gets a reference to the given string and assigns it to the SelfTradePreventionMode field.
func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetSelfTradePreventionMode(v string) {
	o.SelfTradePreventionMode = &v
}

func (o OrderAmendKeepPriorityResponseResultAmendedOrder) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderAmendKeepPriorityResponseResultAmendedOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.OrderListId) {
		toSerialize["orderListId"] = o.OrderListId
	}
	if !common.IsNil(o.OrigClientOrderId) {
		toSerialize["origClientOrderId"] = o.OrigClientOrderId
	}
	if !common.IsNil(o.ClientOrderId) {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.Qty) {
		toSerialize["qty"] = o.Qty
	}
	if !common.IsNil(o.ExecutedQty) {
		toSerialize["executedQty"] = o.ExecutedQty
	}
	if !common.IsNil(o.PreventedQty) {
		toSerialize["preventedQty"] = o.PreventedQty
	}
	if !common.IsNil(o.QuoteOrderQty) {
		toSerialize["quoteOrderQty"] = o.QuoteOrderQty
	}
	if !common.IsNil(o.CumulativeQuoteQty) {
		toSerialize["cumulativeQuoteQty"] = o.CumulativeQuoteQty
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
	if !common.IsNil(o.WorkingTime) {
		toSerialize["workingTime"] = o.WorkingTime
	}
	if !common.IsNil(o.SelfTradePreventionMode) {
		toSerialize["selfTradePreventionMode"] = o.SelfTradePreventionMode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) UnmarshalJSON(data []byte) (err error) {
	varOrderAmendKeepPriorityResponseResultAmendedOrder := _OrderAmendKeepPriorityResponseResultAmendedOrder{}

	err = json.Unmarshal(data, &varOrderAmendKeepPriorityResponseResultAmendedOrder)

	if err != nil {
		return err
	}

	*o = OrderAmendKeepPriorityResponseResultAmendedOrder(varOrderAmendKeepPriorityResponseResultAmendedOrder)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "orderListId")
		delete(additionalProperties, "origClientOrderId")
		delete(additionalProperties, "clientOrderId")
		delete(additionalProperties, "price")
		delete(additionalProperties, "qty")
		delete(additionalProperties, "executedQty")
		delete(additionalProperties, "preventedQty")
		delete(additionalProperties, "quoteOrderQty")
		delete(additionalProperties, "cumulativeQuoteQty")
		delete(additionalProperties, "status")
		delete(additionalProperties, "timeInForce")
		delete(additionalProperties, "type")
		delete(additionalProperties, "side")
		delete(additionalProperties, "workingTime")
		delete(additionalProperties, "selfTradePreventionMode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderAmendKeepPriorityResponseResultAmendedOrder struct {
	value *OrderAmendKeepPriorityResponseResultAmendedOrder
	isSet bool
}

func (v NullableOrderAmendKeepPriorityResponseResultAmendedOrder) Get() *OrderAmendKeepPriorityResponseResultAmendedOrder {
	return v.value
}

func (v *NullableOrderAmendKeepPriorityResponseResultAmendedOrder) Set(val *OrderAmendKeepPriorityResponseResultAmendedOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderAmendKeepPriorityResponseResultAmendedOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderAmendKeepPriorityResponseResultAmendedOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderAmendKeepPriorityResponseResultAmendedOrder(val *OrderAmendKeepPriorityResponseResultAmendedOrder) *NullableOrderAmendKeepPriorityResponseResultAmendedOrder {
	return &NullableOrderAmendKeepPriorityResponseResultAmendedOrder{value: val, isSet: true}
}

func (v NullableOrderAmendKeepPriorityResponseResultAmendedOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderAmendKeepPriorityResponseResultAmendedOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
