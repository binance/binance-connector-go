/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner{}

// CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner struct for CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner
type CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	OrderId              *int64  `json:"orderId,omitempty"`
	ClientOrderId        *string `json:"clientOrderId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner

// NewCancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner instantiates a new CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner() *CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner {
	this := CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner{}
	return &this
}

// NewCancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInnerWithDefaults instantiates a new CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInnerWithDefaults() *CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner {
	this := CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

func (o CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner) ToMap() (map[string]interface{}, error) {
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

func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner) UnmarshalJSON(data []byte) (err error) {
	varCancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner := _CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner{}

	err = json.Unmarshal(data, &varCancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner)

	if err != nil {
		return err
	}

	*o = CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner(varCancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "clientOrderId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner struct {
	value *CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner
	isSet bool
}

func (v NullableCancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner) Get() *CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner {
	return v.value
}

func (v *NullableCancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner) Set(val *CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner(val *CancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner) *NullableCancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner {
	return &NullableCancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner{value: val, isSet: true}
}

func (v NullableCancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelMarginAccountAllOpenOrdersOnASymbolResponseInnerOrdersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
