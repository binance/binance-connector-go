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

// checks if the UmNotionalAndLeverageBracketsResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UmNotionalAndLeverageBracketsResponseInner{}

// UmNotionalAndLeverageBracketsResponseInner struct for UmNotionalAndLeverageBracketsResponseInner
type UmNotionalAndLeverageBracketsResponseInner struct {
	Symbol               *string                                                   `json:"symbol,omitempty"`
	NotionalCoef         *string                                                   `json:"notionalCoef,omitempty"`
	Brackets             []UmNotionalAndLeverageBracketsResponseInnerBracketsInner `json:"brackets,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UmNotionalAndLeverageBracketsResponseInner UmNotionalAndLeverageBracketsResponseInner

// NewUmNotionalAndLeverageBracketsResponseInner instantiates a new UmNotionalAndLeverageBracketsResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUmNotionalAndLeverageBracketsResponseInner() *UmNotionalAndLeverageBracketsResponseInner {
	this := UmNotionalAndLeverageBracketsResponseInner{}
	return &this
}

// NewUmNotionalAndLeverageBracketsResponseInnerWithDefaults instantiates a new UmNotionalAndLeverageBracketsResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUmNotionalAndLeverageBracketsResponseInnerWithDefaults() *UmNotionalAndLeverageBracketsResponseInner {
	this := UmNotionalAndLeverageBracketsResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *UmNotionalAndLeverageBracketsResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmNotionalAndLeverageBracketsResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *UmNotionalAndLeverageBracketsResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *UmNotionalAndLeverageBracketsResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetNotionalCoef returns the NotionalCoef field value if set, zero value otherwise.
func (o *UmNotionalAndLeverageBracketsResponseInner) GetNotionalCoef() string {
	if o == nil || common.IsNil(o.NotionalCoef) {
		var ret string
		return ret
	}
	return *o.NotionalCoef
}

// GetNotionalCoefOk returns a tuple with the NotionalCoef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmNotionalAndLeverageBracketsResponseInner) GetNotionalCoefOk() (*string, bool) {
	if o == nil || common.IsNil(o.NotionalCoef) {
		return nil, false
	}
	return o.NotionalCoef, true
}

// HasNotionalCoef returns a boolean if a field has been set.
func (o *UmNotionalAndLeverageBracketsResponseInner) HasNotionalCoef() bool {
	if o != nil && !common.IsNil(o.NotionalCoef) {
		return true
	}

	return false
}

// SetNotionalCoef gets a reference to the given string and assigns it to the NotionalCoef field.
func (o *UmNotionalAndLeverageBracketsResponseInner) SetNotionalCoef(v string) {
	o.NotionalCoef = &v
}

// GetBrackets returns the Brackets field value if set, zero value otherwise.
func (o *UmNotionalAndLeverageBracketsResponseInner) GetBrackets() []UmNotionalAndLeverageBracketsResponseInnerBracketsInner {
	if o == nil || common.IsNil(o.Brackets) {
		var ret []UmNotionalAndLeverageBracketsResponseInnerBracketsInner
		return ret
	}
	return o.Brackets
}

// GetBracketsOk returns a tuple with the Brackets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmNotionalAndLeverageBracketsResponseInner) GetBracketsOk() ([]UmNotionalAndLeverageBracketsResponseInnerBracketsInner, bool) {
	if o == nil || common.IsNil(o.Brackets) {
		return nil, false
	}
	return o.Brackets, true
}

// HasBrackets returns a boolean if a field has been set.
func (o *UmNotionalAndLeverageBracketsResponseInner) HasBrackets() bool {
	if o != nil && !common.IsNil(o.Brackets) {
		return true
	}

	return false
}

// SetBrackets gets a reference to the given []UmNotionalAndLeverageBracketsResponseInnerBracketsInner and assigns it to the Brackets field.
func (o *UmNotionalAndLeverageBracketsResponseInner) SetBrackets(v []UmNotionalAndLeverageBracketsResponseInnerBracketsInner) {
	o.Brackets = v
}

func (o UmNotionalAndLeverageBracketsResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UmNotionalAndLeverageBracketsResponseInner) ToMap() (map[string]interface{}, error) {
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

func (o *UmNotionalAndLeverageBracketsResponseInner) UnmarshalJSON(data []byte) (err error) {
	varUmNotionalAndLeverageBracketsResponseInner := _UmNotionalAndLeverageBracketsResponseInner{}

	err = json.Unmarshal(data, &varUmNotionalAndLeverageBracketsResponseInner)

	if err != nil {
		return err
	}

	*o = UmNotionalAndLeverageBracketsResponseInner(varUmNotionalAndLeverageBracketsResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "notionalCoef")
		delete(additionalProperties, "brackets")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUmNotionalAndLeverageBracketsResponseInner struct {
	value *UmNotionalAndLeverageBracketsResponseInner
	isSet bool
}

func (v NullableUmNotionalAndLeverageBracketsResponseInner) Get() *UmNotionalAndLeverageBracketsResponseInner {
	return v.value
}

func (v *NullableUmNotionalAndLeverageBracketsResponseInner) Set(val *UmNotionalAndLeverageBracketsResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUmNotionalAndLeverageBracketsResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUmNotionalAndLeverageBracketsResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUmNotionalAndLeverageBracketsResponseInner(val *UmNotionalAndLeverageBracketsResponseInner) *NullableUmNotionalAndLeverageBracketsResponseInner {
	return &NullableUmNotionalAndLeverageBracketsResponseInner{value: val, isSet: true}
}

func (v NullableUmNotionalAndLeverageBracketsResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUmNotionalAndLeverageBracketsResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
