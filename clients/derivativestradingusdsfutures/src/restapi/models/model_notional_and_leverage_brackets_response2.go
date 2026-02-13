/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the NotionalAndLeverageBracketsResponse2 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &NotionalAndLeverageBracketsResponse2{}

// NotionalAndLeverageBracketsResponse2 struct for NotionalAndLeverageBracketsResponse2
type NotionalAndLeverageBracketsResponse2 struct {
	Symbol               *string                                             `json:"symbol,omitempty"`
	NotionalCoef         *float32                                            `json:"notionalCoef,omitempty"`
	Brackets             []NotionalAndLeverageBracketsResponse2BracketsInner `json:"brackets,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NotionalAndLeverageBracketsResponse2 NotionalAndLeverageBracketsResponse2

// NewNotionalAndLeverageBracketsResponse2 instantiates a new NotionalAndLeverageBracketsResponse2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotionalAndLeverageBracketsResponse2() *NotionalAndLeverageBracketsResponse2 {
	this := NotionalAndLeverageBracketsResponse2{}
	return &this
}

// NewNotionalAndLeverageBracketsResponse2WithDefaults instantiates a new NotionalAndLeverageBracketsResponse2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotionalAndLeverageBracketsResponse2WithDefaults() *NotionalAndLeverageBracketsResponse2 {
	this := NotionalAndLeverageBracketsResponse2{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *NotionalAndLeverageBracketsResponse2) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotionalAndLeverageBracketsResponse2) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *NotionalAndLeverageBracketsResponse2) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *NotionalAndLeverageBracketsResponse2) SetSymbol(v string) {
	o.Symbol = &v
}

// GetNotionalCoef returns the NotionalCoef field value if set, zero value otherwise.
func (o *NotionalAndLeverageBracketsResponse2) GetNotionalCoef() float32 {
	if o == nil || common.IsNil(o.NotionalCoef) {
		var ret float32
		return ret
	}
	return *o.NotionalCoef
}

// GetNotionalCoefOk returns a tuple with the NotionalCoef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotionalAndLeverageBracketsResponse2) GetNotionalCoefOk() (*float32, bool) {
	if o == nil || common.IsNil(o.NotionalCoef) {
		return nil, false
	}
	return o.NotionalCoef, true
}

// HasNotionalCoef returns a boolean if a field has been set.
func (o *NotionalAndLeverageBracketsResponse2) HasNotionalCoef() bool {
	if o != nil && !common.IsNil(o.NotionalCoef) {
		return true
	}

	return false
}

// SetNotionalCoef gets a reference to the given float32 and assigns it to the NotionalCoef field.
func (o *NotionalAndLeverageBracketsResponse2) SetNotionalCoef(v float32) {
	o.NotionalCoef = &v
}

// GetBrackets returns the Brackets field value if set, zero value otherwise.
func (o *NotionalAndLeverageBracketsResponse2) GetBrackets() []NotionalAndLeverageBracketsResponse2BracketsInner {
	if o == nil || common.IsNil(o.Brackets) {
		var ret []NotionalAndLeverageBracketsResponse2BracketsInner
		return ret
	}
	return o.Brackets
}

// GetBracketsOk returns a tuple with the Brackets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotionalAndLeverageBracketsResponse2) GetBracketsOk() ([]NotionalAndLeverageBracketsResponse2BracketsInner, bool) {
	if o == nil || common.IsNil(o.Brackets) {
		return nil, false
	}
	return o.Brackets, true
}

// HasBrackets returns a boolean if a field has been set.
func (o *NotionalAndLeverageBracketsResponse2) HasBrackets() bool {
	if o != nil && !common.IsNil(o.Brackets) {
		return true
	}

	return false
}

// SetBrackets gets a reference to the given []NotionalAndLeverageBracketsResponse2BracketsInner and assigns it to the Brackets field.
func (o *NotionalAndLeverageBracketsResponse2) SetBrackets(v []NotionalAndLeverageBracketsResponse2BracketsInner) {
	o.Brackets = v
}

func (o NotionalAndLeverageBracketsResponse2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotionalAndLeverageBracketsResponse2) ToMap() (map[string]interface{}, error) {
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

func (o *NotionalAndLeverageBracketsResponse2) UnmarshalJSON(data []byte) (err error) {
	varNotionalAndLeverageBracketsResponse2 := _NotionalAndLeverageBracketsResponse2{}

	err = json.Unmarshal(data, &varNotionalAndLeverageBracketsResponse2)

	if err != nil {
		return err
	}

	*o = NotionalAndLeverageBracketsResponse2(varNotionalAndLeverageBracketsResponse2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "notionalCoef")
		delete(additionalProperties, "brackets")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNotionalAndLeverageBracketsResponse2 struct {
	value *NotionalAndLeverageBracketsResponse2
	isSet bool
}

func (v NullableNotionalAndLeverageBracketsResponse2) Get() *NotionalAndLeverageBracketsResponse2 {
	return v.value
}

func (v *NullableNotionalAndLeverageBracketsResponse2) Set(val *NotionalAndLeverageBracketsResponse2) {
	v.value = val
	v.isSet = true
}

func (v NullableNotionalAndLeverageBracketsResponse2) IsSet() bool {
	return v.isSet
}

func (v *NullableNotionalAndLeverageBracketsResponse2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotionalAndLeverageBracketsResponse2(val *NotionalAndLeverageBracketsResponse2) *NullableNotionalAndLeverageBracketsResponse2 {
	return &NullableNotionalAndLeverageBracketsResponse2{value: val, isSet: true}
}

func (v NullableNotionalAndLeverageBracketsResponse2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotionalAndLeverageBracketsResponse2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
