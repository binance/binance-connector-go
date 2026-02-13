/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the OrderAmendKeepPriorityResponseAmendedOrder type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderAmendKeepPriorityResponseAmendedOrder{}

// OrderAmendKeepPriorityResponseAmendedOrder struct for OrderAmendKeepPriorityResponseAmendedOrder
type OrderAmendKeepPriorityResponseAmendedOrder struct {
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

type _OrderAmendKeepPriorityResponseAmendedOrder OrderAmendKeepPriorityResponseAmendedOrder

// NewOrderAmendKeepPriorityResponseAmendedOrder instantiates a new OrderAmendKeepPriorityResponseAmendedOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderAmendKeepPriorityResponseAmendedOrder() *OrderAmendKeepPriorityResponseAmendedOrder {
	this := OrderAmendKeepPriorityResponseAmendedOrder{}
	return &this
}

// NewOrderAmendKeepPriorityResponseAmendedOrderWithDefaults instantiates a new OrderAmendKeepPriorityResponseAmendedOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderAmendKeepPriorityResponseAmendedOrderWithDefaults() *OrderAmendKeepPriorityResponseAmendedOrder {
	this := OrderAmendKeepPriorityResponseAmendedOrder{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetSymbol(v string) {
	o.Symbol = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetOrderListId returns the OrderListId field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetOrderListId() int64 {
	if o == nil || common.IsNil(o.OrderListId) {
		var ret int64
		return ret
	}
	return *o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetOrderListIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderListId) {
		return nil, false
	}
	return o.OrderListId, true
}

// HasOrderListId returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasOrderListId() bool {
	if o != nil && !common.IsNil(o.OrderListId) {
		return true
	}

	return false
}

// SetOrderListId gets a reference to the given int64 and assigns it to the OrderListId field.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetOrderListId(v int64) {
	o.OrderListId = &v
}

// GetOrigClientOrderId returns the OrigClientOrderId field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetOrigClientOrderId() string {
	if o == nil || common.IsNil(o.OrigClientOrderId) {
		var ret string
		return ret
	}
	return *o.OrigClientOrderId
}

// GetOrigClientOrderIdOk returns a tuple with the OrigClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetOrigClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrigClientOrderId) {
		return nil, false
	}
	return o.OrigClientOrderId, true
}

// HasOrigClientOrderId returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasOrigClientOrderId() bool {
	if o != nil && !common.IsNil(o.OrigClientOrderId) {
		return true
	}

	return false
}

// SetOrigClientOrderId gets a reference to the given string and assigns it to the OrigClientOrderId field.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetOrigClientOrderId(v string) {
	o.OrigClientOrderId = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetPrice(v string) {
	o.Price = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetQty() string {
	if o == nil || common.IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasQty() bool {
	if o != nil && !common.IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetQty(v string) {
	o.Qty = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetExecutedQty() string {
	if o == nil || common.IsNil(o.ExecutedQty) {
		var ret string
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetExecutedQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasExecutedQty() bool {
	if o != nil && !common.IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given string and assigns it to the ExecutedQty field.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetExecutedQty(v string) {
	o.ExecutedQty = &v
}

// GetPreventedQty returns the PreventedQty field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetPreventedQty() string {
	if o == nil || common.IsNil(o.PreventedQty) {
		var ret string
		return ret
	}
	return *o.PreventedQty
}

// GetPreventedQtyOk returns a tuple with the PreventedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetPreventedQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.PreventedQty) {
		return nil, false
	}
	return o.PreventedQty, true
}

// HasPreventedQty returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasPreventedQty() bool {
	if o != nil && !common.IsNil(o.PreventedQty) {
		return true
	}

	return false
}

// SetPreventedQty gets a reference to the given string and assigns it to the PreventedQty field.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetPreventedQty(v string) {
	o.PreventedQty = &v
}

// GetQuoteOrderQty returns the QuoteOrderQty field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetQuoteOrderQty() string {
	if o == nil || common.IsNil(o.QuoteOrderQty) {
		var ret string
		return ret
	}
	return *o.QuoteOrderQty
}

// GetQuoteOrderQtyOk returns a tuple with the QuoteOrderQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetQuoteOrderQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteOrderQty) {
		return nil, false
	}
	return o.QuoteOrderQty, true
}

// HasQuoteOrderQty returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasQuoteOrderQty() bool {
	if o != nil && !common.IsNil(o.QuoteOrderQty) {
		return true
	}

	return false
}

// SetQuoteOrderQty gets a reference to the given string and assigns it to the QuoteOrderQty field.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetQuoteOrderQty(v string) {
	o.QuoteOrderQty = &v
}

// GetCumulativeQuoteQty returns the CumulativeQuoteQty field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetCumulativeQuoteQty() string {
	if o == nil || common.IsNil(o.CumulativeQuoteQty) {
		var ret string
		return ret
	}
	return *o.CumulativeQuoteQty
}

// GetCumulativeQuoteQtyOk returns a tuple with the CumulativeQuoteQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetCumulativeQuoteQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.CumulativeQuoteQty) {
		return nil, false
	}
	return o.CumulativeQuoteQty, true
}

// HasCumulativeQuoteQty returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasCumulativeQuoteQty() bool {
	if o != nil && !common.IsNil(o.CumulativeQuoteQty) {
		return true
	}

	return false
}

// SetCumulativeQuoteQty gets a reference to the given string and assigns it to the CumulativeQuoteQty field.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetCumulativeQuoteQty(v string) {
	o.CumulativeQuoteQty = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetStatus(v string) {
	o.Status = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetTimeInForce() string {
	if o == nil || common.IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetTimeInForceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasTimeInForce() bool {
	if o != nil && !common.IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetType(v string) {
	o.Type = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetSide(v string) {
	o.Side = &v
}

// GetWorkingTime returns the WorkingTime field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetWorkingTime() int64 {
	if o == nil || common.IsNil(o.WorkingTime) {
		var ret int64
		return ret
	}
	return *o.WorkingTime
}

// GetWorkingTimeOk returns a tuple with the WorkingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetWorkingTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.WorkingTime) {
		return nil, false
	}
	return o.WorkingTime, true
}

// HasWorkingTime returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasWorkingTime() bool {
	if o != nil && !common.IsNil(o.WorkingTime) {
		return true
	}

	return false
}

// SetWorkingTime gets a reference to the given int64 and assigns it to the WorkingTime field.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetWorkingTime(v int64) {
	o.WorkingTime = &v
}

// GetSelfTradePreventionMode returns the SelfTradePreventionMode field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetSelfTradePreventionMode() string {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		var ret string
		return ret
	}
	return *o.SelfTradePreventionMode
}

// GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetSelfTradePreventionModeOk() (*string, bool) {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		return nil, false
	}
	return o.SelfTradePreventionMode, true
}

// HasSelfTradePreventionMode returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasSelfTradePreventionMode() bool {
	if o != nil && !common.IsNil(o.SelfTradePreventionMode) {
		return true
	}

	return false
}

// SetSelfTradePreventionMode gets a reference to the given string and assigns it to the SelfTradePreventionMode field.
func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetSelfTradePreventionMode(v string) {
	o.SelfTradePreventionMode = &v
}

func (o OrderAmendKeepPriorityResponseAmendedOrder) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderAmendKeepPriorityResponseAmendedOrder) ToMap() (map[string]interface{}, error) {
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

func (o *OrderAmendKeepPriorityResponseAmendedOrder) UnmarshalJSON(data []byte) (err error) {
	varOrderAmendKeepPriorityResponseAmendedOrder := _OrderAmendKeepPriorityResponseAmendedOrder{}

	err = json.Unmarshal(data, &varOrderAmendKeepPriorityResponseAmendedOrder)

	if err != nil {
		return err
	}

	*o = OrderAmendKeepPriorityResponseAmendedOrder(varOrderAmendKeepPriorityResponseAmendedOrder)

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

type NullableOrderAmendKeepPriorityResponseAmendedOrder struct {
	value *OrderAmendKeepPriorityResponseAmendedOrder
	isSet bool
}

func (v NullableOrderAmendKeepPriorityResponseAmendedOrder) Get() *OrderAmendKeepPriorityResponseAmendedOrder {
	return v.value
}

func (v *NullableOrderAmendKeepPriorityResponseAmendedOrder) Set(val *OrderAmendKeepPriorityResponseAmendedOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderAmendKeepPriorityResponseAmendedOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderAmendKeepPriorityResponseAmendedOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderAmendKeepPriorityResponseAmendedOrder(val *OrderAmendKeepPriorityResponseAmendedOrder) *NullableOrderAmendKeepPriorityResponseAmendedOrder {
	return &NullableOrderAmendKeepPriorityResponseAmendedOrder{value: val, isSet: true}
}

func (v NullableOrderAmendKeepPriorityResponseAmendedOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderAmendKeepPriorityResponseAmendedOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
