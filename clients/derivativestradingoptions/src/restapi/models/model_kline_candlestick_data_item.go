/*
Options REST API

Access market data, manage accounts, and trade Binance Options.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the KlineCandlestickDataItem type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &KlineCandlestickDataItem{}

// KlineCandlestickDataItem struct for KlineCandlestickDataItem
type KlineCandlestickDataItem struct {
	Items []KlineCandlestickDataItemInner
}

// NewKlineCandlestickDataItem instantiates a new KlineCandlestickDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKlineCandlestickDataItem() *KlineCandlestickDataItem {
	this := KlineCandlestickDataItem{}
	return &this
}

// NewKlineCandlestickDataItemWithDefaults instantiates a new KlineCandlestickDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKlineCandlestickDataItemWithDefaults() *KlineCandlestickDataItem {
	this := KlineCandlestickDataItem{}
	return &this
}

func (o KlineCandlestickDataItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KlineCandlestickDataItem) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *KlineCandlestickDataItem) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableKlineCandlestickDataItem struct {
	value KlineCandlestickDataItem
	isSet bool
}

func (v NullableKlineCandlestickDataItem) Get() KlineCandlestickDataItem {
	return v.value
}

func (v *NullableKlineCandlestickDataItem) Set(val KlineCandlestickDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableKlineCandlestickDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableKlineCandlestickDataItem) Unset() {
	v.value = KlineCandlestickDataItem{}
	v.isSet = false
}

func NewNullableKlineCandlestickDataItem(val KlineCandlestickDataItem) *NullableKlineCandlestickDataItem {
	return &NullableKlineCandlestickDataItem{value: val, isSet: true}
}

func (v NullableKlineCandlestickDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlineCandlestickDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
