/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MarkPriceResponse2 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarkPriceResponse2{}

// MarkPriceResponse2 struct for MarkPriceResponse2
type MarkPriceResponse2 struct {
	Items []MarkPriceResponse2Inner
}

// NewMarkPriceResponse2 instantiates a new MarkPriceResponse2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarkPriceResponse2() *MarkPriceResponse2 {
	this := MarkPriceResponse2{}
	return &this
}

// NewMarkPriceResponse2WithDefaults instantiates a new MarkPriceResponse2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarkPriceResponse2WithDefaults() *MarkPriceResponse2 {
	this := MarkPriceResponse2{}
	return &this
}

func (o MarkPriceResponse2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarkPriceResponse2) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *MarkPriceResponse2) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableMarkPriceResponse2 struct {
	value MarkPriceResponse2
	isSet bool
}

func (v NullableMarkPriceResponse2) Get() MarkPriceResponse2 {
	return v.value
}

func (v *NullableMarkPriceResponse2) Set(val MarkPriceResponse2) {
	v.value = val
	v.isSet = true
}

func (v NullableMarkPriceResponse2) IsSet() bool {
	return v.isSet
}

func (v *NullableMarkPriceResponse2) Unset() {
	v.value = MarkPriceResponse2{}
	v.isSet = false
}

func NewNullableMarkPriceResponse2(val MarkPriceResponse2) *NullableMarkPriceResponse2 {
	return &NullableMarkPriceResponse2{value: val, isSet: true}
}

func (v NullableMarkPriceResponse2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarkPriceResponse2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
