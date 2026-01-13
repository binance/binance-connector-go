/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MarginAccountCancelOcoResponseOrdersInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarginAccountCancelOcoResponseOrdersInner{}

// MarginAccountCancelOcoResponseOrdersInner struct for MarginAccountCancelOcoResponseOrdersInner
type MarginAccountCancelOcoResponseOrdersInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	OrderId              *int64  `json:"orderId,omitempty"`
	ClientOrderId        *string `json:"clientOrderId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarginAccountCancelOcoResponseOrdersInner MarginAccountCancelOcoResponseOrdersInner

// NewMarginAccountCancelOcoResponseOrdersInner instantiates a new MarginAccountCancelOcoResponseOrdersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginAccountCancelOcoResponseOrdersInner() *MarginAccountCancelOcoResponseOrdersInner {
	this := MarginAccountCancelOcoResponseOrdersInner{}
	return &this
}

// NewMarginAccountCancelOcoResponseOrdersInnerWithDefaults instantiates a new MarginAccountCancelOcoResponseOrdersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginAccountCancelOcoResponseOrdersInnerWithDefaults() *MarginAccountCancelOcoResponseOrdersInner {
	this := MarginAccountCancelOcoResponseOrdersInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *MarginAccountCancelOcoResponseOrdersInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelOcoResponseOrdersInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *MarginAccountCancelOcoResponseOrdersInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *MarginAccountCancelOcoResponseOrdersInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *MarginAccountCancelOcoResponseOrdersInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelOcoResponseOrdersInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *MarginAccountCancelOcoResponseOrdersInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *MarginAccountCancelOcoResponseOrdersInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *MarginAccountCancelOcoResponseOrdersInner) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginAccountCancelOcoResponseOrdersInner) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *MarginAccountCancelOcoResponseOrdersInner) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *MarginAccountCancelOcoResponseOrdersInner) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

func (o MarginAccountCancelOcoResponseOrdersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginAccountCancelOcoResponseOrdersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.ClientOrderId) {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MarginAccountCancelOcoResponseOrdersInner) UnmarshalJSON(data []byte) (err error) {
	varMarginAccountCancelOcoResponseOrdersInner := _MarginAccountCancelOcoResponseOrdersInner{}

	err = json.Unmarshal(data, &varMarginAccountCancelOcoResponseOrdersInner)

	if err != nil {
		return err
	}

	*o = MarginAccountCancelOcoResponseOrdersInner(varMarginAccountCancelOcoResponseOrdersInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "clientOrderId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginAccountCancelOcoResponseOrdersInner struct {
	value *MarginAccountCancelOcoResponseOrdersInner
	isSet bool
}

func (v NullableMarginAccountCancelOcoResponseOrdersInner) Get() *MarginAccountCancelOcoResponseOrdersInner {
	return v.value
}

func (v *NullableMarginAccountCancelOcoResponseOrdersInner) Set(val *MarginAccountCancelOcoResponseOrdersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginAccountCancelOcoResponseOrdersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginAccountCancelOcoResponseOrdersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginAccountCancelOcoResponseOrdersInner(val *MarginAccountCancelOcoResponseOrdersInner) *NullableMarginAccountCancelOcoResponseOrdersInner {
	return &NullableMarginAccountCancelOcoResponseOrdersInner{value: val, isSet: true}
}

func (v NullableMarginAccountCancelOcoResponseOrdersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginAccountCancelOcoResponseOrdersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
