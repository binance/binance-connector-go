/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the KlineCandlestickDataResponseItem type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &KlineCandlestickDataResponseItem{}

// KlineCandlestickDataResponseItem struct for KlineCandlestickDataResponseItem
type KlineCandlestickDataResponseItem struct {
	Items []KlineCandlestickDataResponseItemInner
}

// NewKlineCandlestickDataResponseItem instantiates a new KlineCandlestickDataResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKlineCandlestickDataResponseItem() *KlineCandlestickDataResponseItem {
	this := KlineCandlestickDataResponseItem{}
	return &this
}

// NewKlineCandlestickDataResponseItemWithDefaults instantiates a new KlineCandlestickDataResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKlineCandlestickDataResponseItemWithDefaults() *KlineCandlestickDataResponseItem {
	this := KlineCandlestickDataResponseItem{}
	return &this
}

func (o KlineCandlestickDataResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KlineCandlestickDataResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *KlineCandlestickDataResponseItem) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableKlineCandlestickDataResponseItem struct {
	value KlineCandlestickDataResponseItem
	isSet bool
}

func (v NullableKlineCandlestickDataResponseItem) Get() KlineCandlestickDataResponseItem {
	return v.value
}

func (v *NullableKlineCandlestickDataResponseItem) Set(val KlineCandlestickDataResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableKlineCandlestickDataResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableKlineCandlestickDataResponseItem) Unset() {
	v.value = KlineCandlestickDataResponseItem{}
	v.isSet = false
}

func NewNullableKlineCandlestickDataResponseItem(val KlineCandlestickDataResponseItem) *NullableKlineCandlestickDataResponseItem {
	return &NullableKlineCandlestickDataResponseItem{value: val, isSet: true}
}

func (v NullableKlineCandlestickDataResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlineCandlestickDataResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
