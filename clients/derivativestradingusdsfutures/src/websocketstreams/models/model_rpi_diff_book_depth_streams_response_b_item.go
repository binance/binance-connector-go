/*
Binance Derivatives Trading USDS Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the RpiDiffBookDepthStreamsResponseBItem type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RpiDiffBookDepthStreamsResponseBItem{}

// RpiDiffBookDepthStreamsResponseBItem struct for RpiDiffBookDepthStreamsResponseBItem
type RpiDiffBookDepthStreamsResponseBItem struct {
	Items []string
}

// NewRpiDiffBookDepthStreamsResponseBItem instantiates a new RpiDiffBookDepthStreamsResponseBItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRpiDiffBookDepthStreamsResponseBItem() *RpiDiffBookDepthStreamsResponseBItem {
	this := RpiDiffBookDepthStreamsResponseBItem{}
	return &this
}

// NewRpiDiffBookDepthStreamsResponseBItemWithDefaults instantiates a new RpiDiffBookDepthStreamsResponseBItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRpiDiffBookDepthStreamsResponseBItemWithDefaults() *RpiDiffBookDepthStreamsResponseBItem {
	this := RpiDiffBookDepthStreamsResponseBItem{}
	return &this
}

func (o RpiDiffBookDepthStreamsResponseBItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RpiDiffBookDepthStreamsResponseBItem) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *RpiDiffBookDepthStreamsResponseBItem) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableRpiDiffBookDepthStreamsResponseBItem struct {
	value RpiDiffBookDepthStreamsResponseBItem
	isSet bool
}

func (v NullableRpiDiffBookDepthStreamsResponseBItem) Get() RpiDiffBookDepthStreamsResponseBItem {
	return v.value
}

func (v *NullableRpiDiffBookDepthStreamsResponseBItem) Set(val RpiDiffBookDepthStreamsResponseBItem) {
	v.value = val
	v.isSet = true
}

func (v NullableRpiDiffBookDepthStreamsResponseBItem) IsSet() bool {
	return v.isSet
}

func (v *NullableRpiDiffBookDepthStreamsResponseBItem) Unset() {
	v.value = RpiDiffBookDepthStreamsResponseBItem{}
	v.isSet = false
}

func NewNullableRpiDiffBookDepthStreamsResponseBItem(val RpiDiffBookDepthStreamsResponseBItem) *NullableRpiDiffBookDepthStreamsResponseBItem {
	return &NullableRpiDiffBookDepthStreamsResponseBItem{value: val, isSet: true}
}

func (v NullableRpiDiffBookDepthStreamsResponseBItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRpiDiffBookDepthStreamsResponseBItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
