/*
Binance Alpha REST API

OpenAPI Specification for the Binance Alpha REST API

*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the KlinesResponseDataItem type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &KlinesResponseDataItem{}

// KlinesResponseDataItem struct for KlinesResponseDataItem
type KlinesResponseDataItem struct {
	Items []KlinesResponseDataItemInner
}

// NewKlinesResponseDataItem instantiates a new KlinesResponseDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKlinesResponseDataItem() *KlinesResponseDataItem {
	this := KlinesResponseDataItem{}
	return &this
}

// NewKlinesResponseDataItemWithDefaults instantiates a new KlinesResponseDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKlinesResponseDataItemWithDefaults() *KlinesResponseDataItem {
	this := KlinesResponseDataItem{}
	return &this
}

func (o KlinesResponseDataItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KlinesResponseDataItem) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *KlinesResponseDataItem) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableKlinesResponseDataItem struct {
	value KlinesResponseDataItem
	isSet bool
}

func (v NullableKlinesResponseDataItem) Get() KlinesResponseDataItem {
	return v.value
}

func (v *NullableKlinesResponseDataItem) Set(val KlinesResponseDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableKlinesResponseDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableKlinesResponseDataItem) Unset() {
	v.value = KlinesResponseDataItem{}
	v.isSet = false
}

func NewNullableKlinesResponseDataItem(val KlinesResponseDataItem) *NullableKlinesResponseDataItem {
	return &NullableKlinesResponseDataItem{value: val, isSet: true}
}

func (v NullableKlinesResponseDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlinesResponseDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
