/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the NotionalAndLeverageBracketsResponse1Inner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &NotionalAndLeverageBracketsResponse1Inner{}

// NotionalAndLeverageBracketsResponse1Inner struct for NotionalAndLeverageBracketsResponse1Inner
type NotionalAndLeverageBracketsResponse1Inner struct {
	Symbol               *string                                                  `json:"symbol,omitempty"`
	NotionalCoef         *float32                                                 `json:"notionalCoef,omitempty"`
	Brackets             []NotionalAndLeverageBracketsResponse1InnerBracketsInner `json:"brackets,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NotionalAndLeverageBracketsResponse1Inner NotionalAndLeverageBracketsResponse1Inner

// NewNotionalAndLeverageBracketsResponse1Inner instantiates a new NotionalAndLeverageBracketsResponse1Inner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotionalAndLeverageBracketsResponse1Inner() *NotionalAndLeverageBracketsResponse1Inner {
	this := NotionalAndLeverageBracketsResponse1Inner{}
	return &this
}

// NewNotionalAndLeverageBracketsResponse1InnerWithDefaults instantiates a new NotionalAndLeverageBracketsResponse1Inner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotionalAndLeverageBracketsResponse1InnerWithDefaults() *NotionalAndLeverageBracketsResponse1Inner {
	this := NotionalAndLeverageBracketsResponse1Inner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *NotionalAndLeverageBracketsResponse1Inner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotionalAndLeverageBracketsResponse1Inner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *NotionalAndLeverageBracketsResponse1Inner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *NotionalAndLeverageBracketsResponse1Inner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetNotionalCoef returns the NotionalCoef field value if set, zero value otherwise.
func (o *NotionalAndLeverageBracketsResponse1Inner) GetNotionalCoef() float32 {
	if o == nil || common.IsNil(o.NotionalCoef) {
		var ret float32
		return ret
	}
	return *o.NotionalCoef
}

// GetNotionalCoefOk returns a tuple with the NotionalCoef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotionalAndLeverageBracketsResponse1Inner) GetNotionalCoefOk() (*float32, bool) {
	if o == nil || common.IsNil(o.NotionalCoef) {
		return nil, false
	}
	return o.NotionalCoef, true
}

// HasNotionalCoef returns a boolean if a field has been set.
func (o *NotionalAndLeverageBracketsResponse1Inner) HasNotionalCoef() bool {
	if o != nil && !common.IsNil(o.NotionalCoef) {
		return true
	}

	return false
}

// SetNotionalCoef gets a reference to the given float32 and assigns it to the NotionalCoef field.
func (o *NotionalAndLeverageBracketsResponse1Inner) SetNotionalCoef(v float32) {
	o.NotionalCoef = &v
}

// GetBrackets returns the Brackets field value if set, zero value otherwise.
func (o *NotionalAndLeverageBracketsResponse1Inner) GetBrackets() []NotionalAndLeverageBracketsResponse1InnerBracketsInner {
	if o == nil || common.IsNil(o.Brackets) {
		var ret []NotionalAndLeverageBracketsResponse1InnerBracketsInner
		return ret
	}
	return o.Brackets
}

// GetBracketsOk returns a tuple with the Brackets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotionalAndLeverageBracketsResponse1Inner) GetBracketsOk() ([]NotionalAndLeverageBracketsResponse1InnerBracketsInner, bool) {
	if o == nil || common.IsNil(o.Brackets) {
		return nil, false
	}
	return o.Brackets, true
}

// HasBrackets returns a boolean if a field has been set.
func (o *NotionalAndLeverageBracketsResponse1Inner) HasBrackets() bool {
	if o != nil && !common.IsNil(o.Brackets) {
		return true
	}

	return false
}

// SetBrackets gets a reference to the given []NotionalAndLeverageBracketsResponse1InnerBracketsInner and assigns it to the Brackets field.
func (o *NotionalAndLeverageBracketsResponse1Inner) SetBrackets(v []NotionalAndLeverageBracketsResponse1InnerBracketsInner) {
	o.Brackets = v
}

func (o NotionalAndLeverageBracketsResponse1Inner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotionalAndLeverageBracketsResponse1Inner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.NotionalCoef) {
		toSerialize["notionalCoef"] = o.NotionalCoef
	}
	if !common.IsNil(o.Brackets) {
		toSerialize["brackets"] = o.Brackets
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NotionalAndLeverageBracketsResponse1Inner) UnmarshalJSON(data []byte) (err error) {
	varNotionalAndLeverageBracketsResponse1Inner := _NotionalAndLeverageBracketsResponse1Inner{}

	err = json.Unmarshal(data, &varNotionalAndLeverageBracketsResponse1Inner)

	if err != nil {
		return err
	}

	*o = NotionalAndLeverageBracketsResponse1Inner(varNotionalAndLeverageBracketsResponse1Inner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "notionalCoef")
		delete(additionalProperties, "brackets")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNotionalAndLeverageBracketsResponse1Inner struct {
	value *NotionalAndLeverageBracketsResponse1Inner
	isSet bool
}

func (v NullableNotionalAndLeverageBracketsResponse1Inner) Get() *NotionalAndLeverageBracketsResponse1Inner {
	return v.value
}

func (v *NullableNotionalAndLeverageBracketsResponse1Inner) Set(val *NotionalAndLeverageBracketsResponse1Inner) {
	v.value = val
	v.isSet = true
}

func (v NullableNotionalAndLeverageBracketsResponse1Inner) IsSet() bool {
	return v.isSet
}

func (v *NullableNotionalAndLeverageBracketsResponse1Inner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotionalAndLeverageBracketsResponse1Inner(val *NotionalAndLeverageBracketsResponse1Inner) *NullableNotionalAndLeverageBracketsResponse1Inner {
	return &NullableNotionalAndLeverageBracketsResponse1Inner{value: val, isSet: true}
}

func (v NullableNotionalAndLeverageBracketsResponse1Inner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotionalAndLeverageBracketsResponse1Inner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
