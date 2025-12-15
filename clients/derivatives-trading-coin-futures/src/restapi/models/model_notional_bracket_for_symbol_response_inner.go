/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the NotionalBracketForSymbolResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &NotionalBracketForSymbolResponseInner{}

// NotionalBracketForSymbolResponseInner struct for NotionalBracketForSymbolResponseInner
type NotionalBracketForSymbolResponseInner struct {
	Symbol               *string                                            `json:"symbol,omitempty"`
	NotionalCoef         *float32                                           `json:"notionalCoef,omitempty"`
	Brackets             []NotionalBracketForPairResponseInnerBracketsInner `json:"brackets,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NotionalBracketForSymbolResponseInner NotionalBracketForSymbolResponseInner

// NewNotionalBracketForSymbolResponseInner instantiates a new NotionalBracketForSymbolResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotionalBracketForSymbolResponseInner() *NotionalBracketForSymbolResponseInner {
	this := NotionalBracketForSymbolResponseInner{}
	return &this
}

// NewNotionalBracketForSymbolResponseInnerWithDefaults instantiates a new NotionalBracketForSymbolResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotionalBracketForSymbolResponseInnerWithDefaults() *NotionalBracketForSymbolResponseInner {
	this := NotionalBracketForSymbolResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *NotionalBracketForSymbolResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotionalBracketForSymbolResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *NotionalBracketForSymbolResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *NotionalBracketForSymbolResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetNotionalCoef returns the NotionalCoef field value if set, zero value otherwise.
func (o *NotionalBracketForSymbolResponseInner) GetNotionalCoef() float32 {
	if o == nil || common.IsNil(o.NotionalCoef) {
		var ret float32
		return ret
	}
	return *o.NotionalCoef
}

// GetNotionalCoefOk returns a tuple with the NotionalCoef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotionalBracketForSymbolResponseInner) GetNotionalCoefOk() (*float32, bool) {
	if o == nil || common.IsNil(o.NotionalCoef) {
		return nil, false
	}
	return o.NotionalCoef, true
}

// HasNotionalCoef returns a boolean if a field has been set.
func (o *NotionalBracketForSymbolResponseInner) HasNotionalCoef() bool {
	if o != nil && !common.IsNil(o.NotionalCoef) {
		return true
	}

	return false
}

// SetNotionalCoef gets a reference to the given float32 and assigns it to the NotionalCoef field.
func (o *NotionalBracketForSymbolResponseInner) SetNotionalCoef(v float32) {
	o.NotionalCoef = &v
}

// GetBrackets returns the Brackets field value if set, zero value otherwise.
func (o *NotionalBracketForSymbolResponseInner) GetBrackets() []NotionalBracketForPairResponseInnerBracketsInner {
	if o == nil || common.IsNil(o.Brackets) {
		var ret []NotionalBracketForPairResponseInnerBracketsInner
		return ret
	}
	return o.Brackets
}

// GetBracketsOk returns a tuple with the Brackets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotionalBracketForSymbolResponseInner) GetBracketsOk() ([]NotionalBracketForPairResponseInnerBracketsInner, bool) {
	if o == nil || common.IsNil(o.Brackets) {
		return nil, false
	}
	return o.Brackets, true
}

// HasBrackets returns a boolean if a field has been set.
func (o *NotionalBracketForSymbolResponseInner) HasBrackets() bool {
	if o != nil && !common.IsNil(o.Brackets) {
		return true
	}

	return false
}

// SetBrackets gets a reference to the given []NotionalBracketForPairResponseInnerBracketsInner and assigns it to the Brackets field.
func (o *NotionalBracketForSymbolResponseInner) SetBrackets(v []NotionalBracketForPairResponseInnerBracketsInner) {
	o.Brackets = v
}

func (o NotionalBracketForSymbolResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotionalBracketForSymbolResponseInner) ToMap() (map[string]interface{}, error) {
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

func (o *NotionalBracketForSymbolResponseInner) UnmarshalJSON(data []byte) (err error) {
	varNotionalBracketForSymbolResponseInner := _NotionalBracketForSymbolResponseInner{}

	err = json.Unmarshal(data, &varNotionalBracketForSymbolResponseInner)

	if err != nil {
		return err
	}

	*o = NotionalBracketForSymbolResponseInner(varNotionalBracketForSymbolResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "notionalCoef")
		delete(additionalProperties, "brackets")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNotionalBracketForSymbolResponseInner struct {
	value *NotionalBracketForSymbolResponseInner
	isSet bool
}

func (v NullableNotionalBracketForSymbolResponseInner) Get() *NotionalBracketForSymbolResponseInner {
	return v.value
}

func (v *NullableNotionalBracketForSymbolResponseInner) Set(val *NotionalBracketForSymbolResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableNotionalBracketForSymbolResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableNotionalBracketForSymbolResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotionalBracketForSymbolResponseInner(val *NotionalBracketForSymbolResponseInner) *NullableNotionalBracketForSymbolResponseInner {
	return &NullableNotionalBracketForSymbolResponseInner{value: val, isSet: true}
}

func (v NullableNotionalBracketForSymbolResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotionalBracketForSymbolResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
