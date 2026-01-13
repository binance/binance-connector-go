/*
Binance Derivatives Trading USDS Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the RpiDiffBookDepthStreamsResponseAItem type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RpiDiffBookDepthStreamsResponseAItem{}

// RpiDiffBookDepthStreamsResponseAItem struct for RpiDiffBookDepthStreamsResponseAItem
type RpiDiffBookDepthStreamsResponseAItem struct {
	Items []string
}

// NewRpiDiffBookDepthStreamsResponseAItem instantiates a new RpiDiffBookDepthStreamsResponseAItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRpiDiffBookDepthStreamsResponseAItem() *RpiDiffBookDepthStreamsResponseAItem {
	this := RpiDiffBookDepthStreamsResponseAItem{}
	return &this
}

// NewRpiDiffBookDepthStreamsResponseAItemWithDefaults instantiates a new RpiDiffBookDepthStreamsResponseAItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRpiDiffBookDepthStreamsResponseAItemWithDefaults() *RpiDiffBookDepthStreamsResponseAItem {
	this := RpiDiffBookDepthStreamsResponseAItem{}
	return &this
}

func (o RpiDiffBookDepthStreamsResponseAItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RpiDiffBookDepthStreamsResponseAItem) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *RpiDiffBookDepthStreamsResponseAItem) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableRpiDiffBookDepthStreamsResponseAItem struct {
	value RpiDiffBookDepthStreamsResponseAItem
	isSet bool
}

func (v NullableRpiDiffBookDepthStreamsResponseAItem) Get() RpiDiffBookDepthStreamsResponseAItem {
	return v.value
}

func (v *NullableRpiDiffBookDepthStreamsResponseAItem) Set(val RpiDiffBookDepthStreamsResponseAItem) {
	v.value = val
	v.isSet = true
}

func (v NullableRpiDiffBookDepthStreamsResponseAItem) IsSet() bool {
	return v.isSet
}

func (v *NullableRpiDiffBookDepthStreamsResponseAItem) Unset() {
	v.value = RpiDiffBookDepthStreamsResponseAItem{}
	v.isSet = false
}

func NewNullableRpiDiffBookDepthStreamsResponseAItem(val RpiDiffBookDepthStreamsResponseAItem) *NullableRpiDiffBookDepthStreamsResponseAItem {
	return &NullableRpiDiffBookDepthStreamsResponseAItem{value: val, isSet: true}
}

func (v NullableRpiDiffBookDepthStreamsResponseAItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRpiDiffBookDepthStreamsResponseAItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
