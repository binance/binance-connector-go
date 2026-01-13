/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CancelMarginAccountOcoOrdersResponseOrdersInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CancelMarginAccountOcoOrdersResponseOrdersInner{}

// CancelMarginAccountOcoOrdersResponseOrdersInner struct for CancelMarginAccountOcoOrdersResponseOrdersInner
type CancelMarginAccountOcoOrdersResponseOrdersInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	OrderId              *int64  `json:"orderId,omitempty"`
	ClientOrderId        *string `json:"clientOrderId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CancelMarginAccountOcoOrdersResponseOrdersInner CancelMarginAccountOcoOrdersResponseOrdersInner

// NewCancelMarginAccountOcoOrdersResponseOrdersInner instantiates a new CancelMarginAccountOcoOrdersResponseOrdersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelMarginAccountOcoOrdersResponseOrdersInner() *CancelMarginAccountOcoOrdersResponseOrdersInner {
	this := CancelMarginAccountOcoOrdersResponseOrdersInner{}
	return &this
}

// NewCancelMarginAccountOcoOrdersResponseOrdersInnerWithDefaults instantiates a new CancelMarginAccountOcoOrdersResponseOrdersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelMarginAccountOcoOrdersResponseOrdersInnerWithDefaults() *CancelMarginAccountOcoOrdersResponseOrdersInner {
	this := CancelMarginAccountOcoOrdersResponseOrdersInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *CancelMarginAccountOcoOrdersResponseOrdersInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrdersInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrdersInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *CancelMarginAccountOcoOrdersResponseOrdersInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *CancelMarginAccountOcoOrdersResponseOrdersInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrdersInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrdersInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *CancelMarginAccountOcoOrdersResponseOrdersInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *CancelMarginAccountOcoOrdersResponseOrdersInner) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrdersInner) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *CancelMarginAccountOcoOrdersResponseOrdersInner) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *CancelMarginAccountOcoOrdersResponseOrdersInner) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

func (o CancelMarginAccountOcoOrdersResponseOrdersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelMarginAccountOcoOrdersResponseOrdersInner) ToMap() (map[string]interface{}, error) {
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

func (o *CancelMarginAccountOcoOrdersResponseOrdersInner) UnmarshalJSON(data []byte) (err error) {
	varCancelMarginAccountOcoOrdersResponseOrdersInner := _CancelMarginAccountOcoOrdersResponseOrdersInner{}

	err = json.Unmarshal(data, &varCancelMarginAccountOcoOrdersResponseOrdersInner)

	if err != nil {
		return err
	}

	*o = CancelMarginAccountOcoOrdersResponseOrdersInner(varCancelMarginAccountOcoOrdersResponseOrdersInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "clientOrderId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCancelMarginAccountOcoOrdersResponseOrdersInner struct {
	value *CancelMarginAccountOcoOrdersResponseOrdersInner
	isSet bool
}

func (v NullableCancelMarginAccountOcoOrdersResponseOrdersInner) Get() *CancelMarginAccountOcoOrdersResponseOrdersInner {
	return v.value
}

func (v *NullableCancelMarginAccountOcoOrdersResponseOrdersInner) Set(val *CancelMarginAccountOcoOrdersResponseOrdersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelMarginAccountOcoOrdersResponseOrdersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelMarginAccountOcoOrdersResponseOrdersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelMarginAccountOcoOrdersResponseOrdersInner(val *CancelMarginAccountOcoOrdersResponseOrdersInner) *NullableCancelMarginAccountOcoOrdersResponseOrdersInner {
	return &NullableCancelMarginAccountOcoOrdersResponseOrdersInner{value: val, isSet: true}
}

func (v NullableCancelMarginAccountOcoOrdersResponseOrdersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelMarginAccountOcoOrdersResponseOrdersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
