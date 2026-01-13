/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MultiAssetsModeAssetIndexResponse2 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MultiAssetsModeAssetIndexResponse2{}

// MultiAssetsModeAssetIndexResponse2 struct for MultiAssetsModeAssetIndexResponse2
type MultiAssetsModeAssetIndexResponse2 struct {
	Items []MultiAssetsModeAssetIndexResponse2Inner
}

// NewMultiAssetsModeAssetIndexResponse2 instantiates a new MultiAssetsModeAssetIndexResponse2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiAssetsModeAssetIndexResponse2() *MultiAssetsModeAssetIndexResponse2 {
	this := MultiAssetsModeAssetIndexResponse2{}
	return &this
}

// NewMultiAssetsModeAssetIndexResponse2WithDefaults instantiates a new MultiAssetsModeAssetIndexResponse2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiAssetsModeAssetIndexResponse2WithDefaults() *MultiAssetsModeAssetIndexResponse2 {
	this := MultiAssetsModeAssetIndexResponse2{}
	return &this
}

func (o MultiAssetsModeAssetIndexResponse2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultiAssetsModeAssetIndexResponse2) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *MultiAssetsModeAssetIndexResponse2) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableMultiAssetsModeAssetIndexResponse2 struct {
	value MultiAssetsModeAssetIndexResponse2
	isSet bool
}

func (v NullableMultiAssetsModeAssetIndexResponse2) Get() MultiAssetsModeAssetIndexResponse2 {
	return v.value
}

func (v *NullableMultiAssetsModeAssetIndexResponse2) Set(val MultiAssetsModeAssetIndexResponse2) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiAssetsModeAssetIndexResponse2) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiAssetsModeAssetIndexResponse2) Unset() {
	v.value = MultiAssetsModeAssetIndexResponse2{}
	v.isSet = false
}

func NewNullableMultiAssetsModeAssetIndexResponse2(val MultiAssetsModeAssetIndexResponse2) *NullableMultiAssetsModeAssetIndexResponse2 {
	return &NullableMultiAssetsModeAssetIndexResponse2{value: val, isSet: true}
}

func (v NullableMultiAssetsModeAssetIndexResponse2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiAssetsModeAssetIndexResponse2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
