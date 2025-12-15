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

// checks if the NotionalBracketForPairResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &NotionalBracketForPairResponseInner{}

// NotionalBracketForPairResponseInner struct for NotionalBracketForPairResponseInner
type NotionalBracketForPairResponseInner struct {
	Pair                 *string                                            `json:"pair,omitempty"`
	Brackets             []NotionalBracketForPairResponseInnerBracketsInner `json:"brackets,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NotionalBracketForPairResponseInner NotionalBracketForPairResponseInner

// NewNotionalBracketForPairResponseInner instantiates a new NotionalBracketForPairResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotionalBracketForPairResponseInner() *NotionalBracketForPairResponseInner {
	this := NotionalBracketForPairResponseInner{}
	return &this
}

// NewNotionalBracketForPairResponseInnerWithDefaults instantiates a new NotionalBracketForPairResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotionalBracketForPairResponseInnerWithDefaults() *NotionalBracketForPairResponseInner {
	this := NotionalBracketForPairResponseInner{}
	return &this
}

// GetPair returns the Pair field value if set, zero value otherwise.
func (o *NotionalBracketForPairResponseInner) GetPair() string {
	if o == nil || common.IsNil(o.Pair) {
		var ret string
		return ret
	}
	return *o.Pair
}

// GetPairOk returns a tuple with the Pair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotionalBracketForPairResponseInner) GetPairOk() (*string, bool) {
	if o == nil || common.IsNil(o.Pair) {
		return nil, false
	}
	return o.Pair, true
}

// HasPair returns a boolean if a field has been set.
func (o *NotionalBracketForPairResponseInner) HasPair() bool {
	if o != nil && !common.IsNil(o.Pair) {
		return true
	}

	return false
}

// SetPair gets a reference to the given string and assigns it to the Pair field.
func (o *NotionalBracketForPairResponseInner) SetPair(v string) {
	o.Pair = &v
}

// GetBrackets returns the Brackets field value if set, zero value otherwise.
func (o *NotionalBracketForPairResponseInner) GetBrackets() []NotionalBracketForPairResponseInnerBracketsInner {
	if o == nil || common.IsNil(o.Brackets) {
		var ret []NotionalBracketForPairResponseInnerBracketsInner
		return ret
	}
	return o.Brackets
}

// GetBracketsOk returns a tuple with the Brackets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotionalBracketForPairResponseInner) GetBracketsOk() ([]NotionalBracketForPairResponseInnerBracketsInner, bool) {
	if o == nil || common.IsNil(o.Brackets) {
		return nil, false
	}
	return o.Brackets, true
}

// HasBrackets returns a boolean if a field has been set.
func (o *NotionalBracketForPairResponseInner) HasBrackets() bool {
	if o != nil && !common.IsNil(o.Brackets) {
		return true
	}

	return false
}

// SetBrackets gets a reference to the given []NotionalBracketForPairResponseInnerBracketsInner and assigns it to the Brackets field.
func (o *NotionalBracketForPairResponseInner) SetBrackets(v []NotionalBracketForPairResponseInnerBracketsInner) {
	o.Brackets = v
}

func (o NotionalBracketForPairResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotionalBracketForPairResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Pair) {
		toSerialize["pair"] = o.Pair
	}
	if !common.IsNil(o.Brackets) {
		toSerialize["brackets"] = o.Brackets
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NotionalBracketForPairResponseInner) UnmarshalJSON(data []byte) (err error) {
	varNotionalBracketForPairResponseInner := _NotionalBracketForPairResponseInner{}

	err = json.Unmarshal(data, &varNotionalBracketForPairResponseInner)

	if err != nil {
		return err
	}

	*o = NotionalBracketForPairResponseInner(varNotionalBracketForPairResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pair")
		delete(additionalProperties, "brackets")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNotionalBracketForPairResponseInner struct {
	value *NotionalBracketForPairResponseInner
	isSet bool
}

func (v NullableNotionalBracketForPairResponseInner) Get() *NotionalBracketForPairResponseInner {
	return v.value
}

func (v *NullableNotionalBracketForPairResponseInner) Set(val *NotionalBracketForPairResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableNotionalBracketForPairResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableNotionalBracketForPairResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotionalBracketForPairResponseInner(val *NotionalBracketForPairResponseInner) *NullableNotionalBracketForPairResponseInner {
	return &NullableNotionalBracketForPairResponseInner{value: val, isSet: true}
}

func (v NullableNotionalBracketForPairResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotionalBracketForPairResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
