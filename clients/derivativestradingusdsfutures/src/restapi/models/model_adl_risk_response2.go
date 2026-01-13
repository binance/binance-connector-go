/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AdlRiskResponse2 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AdlRiskResponse2{}

// AdlRiskResponse2 struct for AdlRiskResponse2
type AdlRiskResponse2 struct {
	Items []AdlRiskResponse2Inner
}

// NewAdlRiskResponse2 instantiates a new AdlRiskResponse2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdlRiskResponse2() *AdlRiskResponse2 {
	this := AdlRiskResponse2{}
	return &this
}

// NewAdlRiskResponse2WithDefaults instantiates a new AdlRiskResponse2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdlRiskResponse2WithDefaults() *AdlRiskResponse2 {
	this := AdlRiskResponse2{}
	return &this
}

func (o AdlRiskResponse2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdlRiskResponse2) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *AdlRiskResponse2) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableAdlRiskResponse2 struct {
	value AdlRiskResponse2
	isSet bool
}

func (v NullableAdlRiskResponse2) Get() AdlRiskResponse2 {
	return v.value
}

func (v *NullableAdlRiskResponse2) Set(val AdlRiskResponse2) {
	v.value = val
	v.isSet = true
}

func (v NullableAdlRiskResponse2) IsSet() bool {
	return v.isSet
}

func (v *NullableAdlRiskResponse2) Unset() {
	v.value = AdlRiskResponse2{}
	v.isSet = false
}

func NewNullableAdlRiskResponse2(val AdlRiskResponse2) *NullableAdlRiskResponse2 {
	return &NullableAdlRiskResponse2{value: val, isSet: true}
}

func (v NullableAdlRiskResponse2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdlRiskResponse2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
