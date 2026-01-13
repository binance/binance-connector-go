/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the NotionalAndLeverageBracketsResponse1 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &NotionalAndLeverageBracketsResponse1{}

// NotionalAndLeverageBracketsResponse1 struct for NotionalAndLeverageBracketsResponse1
type NotionalAndLeverageBracketsResponse1 struct {
	Items []NotionalAndLeverageBracketsResponse1Inner
}

// NewNotionalAndLeverageBracketsResponse1 instantiates a new NotionalAndLeverageBracketsResponse1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotionalAndLeverageBracketsResponse1() *NotionalAndLeverageBracketsResponse1 {
	this := NotionalAndLeverageBracketsResponse1{}
	return &this
}

// NewNotionalAndLeverageBracketsResponse1WithDefaults instantiates a new NotionalAndLeverageBracketsResponse1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotionalAndLeverageBracketsResponse1WithDefaults() *NotionalAndLeverageBracketsResponse1 {
	this := NotionalAndLeverageBracketsResponse1{}
	return &this
}

func (o NotionalAndLeverageBracketsResponse1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotionalAndLeverageBracketsResponse1) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *NotionalAndLeverageBracketsResponse1) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableNotionalAndLeverageBracketsResponse1 struct {
	value NotionalAndLeverageBracketsResponse1
	isSet bool
}

func (v NullableNotionalAndLeverageBracketsResponse1) Get() NotionalAndLeverageBracketsResponse1 {
	return v.value
}

func (v *NullableNotionalAndLeverageBracketsResponse1) Set(val NotionalAndLeverageBracketsResponse1) {
	v.value = val
	v.isSet = true
}

func (v NullableNotionalAndLeverageBracketsResponse1) IsSet() bool {
	return v.isSet
}

func (v *NullableNotionalAndLeverageBracketsResponse1) Unset() {
	v.value = NotionalAndLeverageBracketsResponse1{}
	v.isSet = false
}

func NewNullableNotionalAndLeverageBracketsResponse1(val NotionalAndLeverageBracketsResponse1) *NullableNotionalAndLeverageBracketsResponse1 {
	return &NullableNotionalAndLeverageBracketsResponse1{value: val, isSet: true}
}

func (v NullableNotionalAndLeverageBracketsResponse1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotionalAndLeverageBracketsResponse1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
