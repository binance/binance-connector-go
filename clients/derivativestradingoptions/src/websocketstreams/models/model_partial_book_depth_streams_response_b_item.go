/*
Binance Derivatives Trading Options WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Options WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the PartialBookDepthStreamsResponseBItem type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PartialBookDepthStreamsResponseBItem{}

// PartialBookDepthStreamsResponseBItem struct for PartialBookDepthStreamsResponseBItem
type PartialBookDepthStreamsResponseBItem struct {
	Items []string
}

// NewPartialBookDepthStreamsResponseBItem instantiates a new PartialBookDepthStreamsResponseBItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartialBookDepthStreamsResponseBItem() *PartialBookDepthStreamsResponseBItem {
	this := PartialBookDepthStreamsResponseBItem{}
	return &this
}

// NewPartialBookDepthStreamsResponseBItemWithDefaults instantiates a new PartialBookDepthStreamsResponseBItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartialBookDepthStreamsResponseBItemWithDefaults() *PartialBookDepthStreamsResponseBItem {
	this := PartialBookDepthStreamsResponseBItem{}
	return &this
}

func (o PartialBookDepthStreamsResponseBItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PartialBookDepthStreamsResponseBItem) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *PartialBookDepthStreamsResponseBItem) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullablePartialBookDepthStreamsResponseBItem struct {
	value PartialBookDepthStreamsResponseBItem
	isSet bool
}

func (v NullablePartialBookDepthStreamsResponseBItem) Get() PartialBookDepthStreamsResponseBItem {
	return v.value
}

func (v *NullablePartialBookDepthStreamsResponseBItem) Set(val PartialBookDepthStreamsResponseBItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePartialBookDepthStreamsResponseBItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePartialBookDepthStreamsResponseBItem) Unset() {
	v.value = PartialBookDepthStreamsResponseBItem{}
	v.isSet = false
}

func NewNullablePartialBookDepthStreamsResponseBItem(val PartialBookDepthStreamsResponseBItem) *NullablePartialBookDepthStreamsResponseBItem {
	return &NullablePartialBookDepthStreamsResponseBItem{value: val, isSet: true}
}

func (v NullablePartialBookDepthStreamsResponseBItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartialBookDepthStreamsResponseBItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
