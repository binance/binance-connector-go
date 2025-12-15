/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CancelMarginAccountOcoOrdersResponseOrderReportsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CancelMarginAccountOcoOrdersResponseOrderReportsInner{}

// CancelMarginAccountOcoOrdersResponseOrderReportsInner struct for CancelMarginAccountOcoOrdersResponseOrderReportsInner
type CancelMarginAccountOcoOrdersResponseOrderReportsInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	OrigClientOrderId    *string `json:"origClientOrderId,omitempty"`
	OrderId              *int64  `json:"orderId,omitempty"`
	OrderListId          *int64  `json:"orderListId,omitempty"`
	ClientOrderId        *string `json:"clientOrderId,omitempty"`
	Price                *string `json:"price,omitempty"`
	OrigQty              *string `json:"origQty,omitempty"`
	ExecutedQty          *string `json:"executedQty,omitempty"`
	CummulativeQuoteQty  *string `json:"cummulativeQuoteQty,omitempty"`
	Status               *string `json:"status,omitempty"`
	TimeInForce          *string `json:"timeInForce,omitempty"`
	Type                 *string `json:"type,omitempty"`
	Side                 *string `json:"side,omitempty"`
	StopPrice            *string `json:"stopPrice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CancelMarginAccountOcoOrdersResponseOrderReportsInner CancelMarginAccountOcoOrdersResponseOrderReportsInner

// NewCancelMarginAccountOcoOrdersResponseOrderReportsInner instantiates a new CancelMarginAccountOcoOrdersResponseOrderReportsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelMarginAccountOcoOrdersResponseOrderReportsInner() *CancelMarginAccountOcoOrdersResponseOrderReportsInner {
	this := CancelMarginAccountOcoOrdersResponseOrderReportsInner{}
	return &this
}

// NewCancelMarginAccountOcoOrdersResponseOrderReportsInnerWithDefaults instantiates a new CancelMarginAccountOcoOrdersResponseOrderReportsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelMarginAccountOcoOrdersResponseOrderReportsInnerWithDefaults() *CancelMarginAccountOcoOrdersResponseOrderReportsInner {
	this := CancelMarginAccountOcoOrdersResponseOrderReportsInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetOrigClientOrderId returns the OrigClientOrderId field value if set, zero value otherwise.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) GetOrigClientOrderId() string {
	if o == nil || common.IsNil(o.OrigClientOrderId) {
		var ret string
		return ret
	}
	return *o.OrigClientOrderId
}

// GetOrigClientOrderIdOk returns a tuple with the OrigClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) GetOrigClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrigClientOrderId) {
		return nil, false
	}
	return o.OrigClientOrderId, true
}

// HasOrigClientOrderId returns a boolean if a field has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) HasOrigClientOrderId() bool {
	if o != nil && !common.IsNil(o.OrigClientOrderId) {
		return true
	}

	return false
}

// SetOrigClientOrderId gets a reference to the given string and assigns it to the OrigClientOrderId field.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) SetOrigClientOrderId(v string) {
	o.OrigClientOrderId = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetOrderListId returns the OrderListId field value if set, zero value otherwise.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) GetOrderListId() int64 {
	if o == nil || common.IsNil(o.OrderListId) {
		var ret int64
		return ret
	}
	return *o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) GetOrderListIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderListId) {
		return nil, false
	}
	return o.OrderListId, true
}

// HasOrderListId returns a boolean if a field has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) HasOrderListId() bool {
	if o != nil && !common.IsNil(o.OrderListId) {
		return true
	}

	return false
}

// SetOrderListId gets a reference to the given int64 and assigns it to the OrderListId field.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) SetOrderListId(v int64) {
	o.OrderListId = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) SetPrice(v string) {
	o.Price = &v
}

// GetOrigQty returns the OrigQty field value if set, zero value otherwise.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) GetOrigQty() string {
	if o == nil || common.IsNil(o.OrigQty) {
		var ret string
		return ret
	}
	return *o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) GetOrigQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrigQty) {
		return nil, false
	}
	return o.OrigQty, true
}

// HasOrigQty returns a boolean if a field has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) HasOrigQty() bool {
	if o != nil && !common.IsNil(o.OrigQty) {
		return true
	}

	return false
}

// SetOrigQty gets a reference to the given string and assigns it to the OrigQty field.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) SetOrigQty(v string) {
	o.OrigQty = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) GetExecutedQty() string {
	if o == nil || common.IsNil(o.ExecutedQty) {
		var ret string
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) GetExecutedQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) HasExecutedQty() bool {
	if o != nil && !common.IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given string and assigns it to the ExecutedQty field.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) SetExecutedQty(v string) {
	o.ExecutedQty = &v
}

// GetCummulativeQuoteQty returns the CummulativeQuoteQty field value if set, zero value otherwise.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) GetCummulativeQuoteQty() string {
	if o == nil || common.IsNil(o.CummulativeQuoteQty) {
		var ret string
		return ret
	}
	return *o.CummulativeQuoteQty
}

// GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) GetCummulativeQuoteQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.CummulativeQuoteQty) {
		return nil, false
	}
	return o.CummulativeQuoteQty, true
}

// HasCummulativeQuoteQty returns a boolean if a field has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) HasCummulativeQuoteQty() bool {
	if o != nil && !common.IsNil(o.CummulativeQuoteQty) {
		return true
	}

	return false
}

// SetCummulativeQuoteQty gets a reference to the given string and assigns it to the CummulativeQuoteQty field.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) SetCummulativeQuoteQty(v string) {
	o.CummulativeQuoteQty = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) SetStatus(v string) {
	o.Status = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) GetTimeInForce() string {
	if o == nil || common.IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) GetTimeInForceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) HasTimeInForce() bool {
	if o != nil && !common.IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) SetType(v string) {
	o.Type = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) SetSide(v string) {
	o.Side = &v
}

// GetStopPrice returns the StopPrice field value if set, zero value otherwise.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) GetStopPrice() string {
	if o == nil || common.IsNil(o.StopPrice) {
		var ret string
		return ret
	}
	return *o.StopPrice
}

// GetStopPriceOk returns a tuple with the StopPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) GetStopPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.StopPrice) {
		return nil, false
	}
	return o.StopPrice, true
}

// HasStopPrice returns a boolean if a field has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) HasStopPrice() bool {
	if o != nil && !common.IsNil(o.StopPrice) {
		return true
	}

	return false
}

// SetStopPrice gets a reference to the given string and assigns it to the StopPrice field.
func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) SetStopPrice(v string) {
	o.StopPrice = &v
}

func (o CancelMarginAccountOcoOrdersResponseOrderReportsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelMarginAccountOcoOrdersResponseOrderReportsInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.OrigQty) {
		toSerialize["origQty"] = o.OrigQty
	}
	if !common.IsNil(o.ExecutedQty) {
		toSerialize["executedQty"] = o.ExecutedQty
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
	if !common.IsNil(o.StopPrice) {
		toSerialize["stopPrice"] = o.StopPrice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CancelMarginAccountOcoOrdersResponseOrderReportsInner) UnmarshalJSON(data []byte) (err error) {
	varCancelMarginAccountOcoOrdersResponseOrderReportsInner := _CancelMarginAccountOcoOrdersResponseOrderReportsInner{}

	err = json.Unmarshal(data, &varCancelMarginAccountOcoOrdersResponseOrderReportsInner)

	if err != nil {
		return err
	}

	*o = CancelMarginAccountOcoOrdersResponseOrderReportsInner(varCancelMarginAccountOcoOrdersResponseOrderReportsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "origClientOrderId")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "orderListId")
		delete(additionalProperties, "clientOrderId")
		delete(additionalProperties, "price")
		delete(additionalProperties, "origQty")
		delete(additionalProperties, "executedQty")
		delete(additionalProperties, "cummulativeQuoteQty")
		delete(additionalProperties, "status")
		delete(additionalProperties, "timeInForce")
		delete(additionalProperties, "type")
		delete(additionalProperties, "side")
		delete(additionalProperties, "stopPrice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCancelMarginAccountOcoOrdersResponseOrderReportsInner struct {
	value *CancelMarginAccountOcoOrdersResponseOrderReportsInner
	isSet bool
}

func (v NullableCancelMarginAccountOcoOrdersResponseOrderReportsInner) Get() *CancelMarginAccountOcoOrdersResponseOrderReportsInner {
	return v.value
}

func (v *NullableCancelMarginAccountOcoOrdersResponseOrderReportsInner) Set(val *CancelMarginAccountOcoOrdersResponseOrderReportsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelMarginAccountOcoOrdersResponseOrderReportsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelMarginAccountOcoOrdersResponseOrderReportsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelMarginAccountOcoOrdersResponseOrderReportsInner(val *CancelMarginAccountOcoOrdersResponseOrderReportsInner) *NullableCancelMarginAccountOcoOrdersResponseOrderReportsInner {
	return &NullableCancelMarginAccountOcoOrdersResponseOrderReportsInner{value: val, isSet: true}
}

func (v NullableCancelMarginAccountOcoOrdersResponseOrderReportsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelMarginAccountOcoOrdersResponseOrderReportsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
