/*
Binance Derivatives Trading USDS Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the PartialBookDepthStreamsResponseAItem type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PartialBookDepthStreamsResponseAItem{}

// PartialBookDepthStreamsResponseAItem struct for PartialBookDepthStreamsResponseAItem
type PartialBookDepthStreamsResponseAItem struct {
	Items []string
}

// NewPartialBookDepthStreamsResponseAItem instantiates a new PartialBookDepthStreamsResponseAItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartialBookDepthStreamsResponseAItem() *PartialBookDepthStreamsResponseAItem {
	this := PartialBookDepthStreamsResponseAItem{}
	return &this
}

// NewPartialBookDepthStreamsResponseAItemWithDefaults instantiates a new PartialBookDepthStreamsResponseAItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartialBookDepthStreamsResponseAItemWithDefaults() *PartialBookDepthStreamsResponseAItem {
	this := PartialBookDepthStreamsResponseAItem{}
	return &this
}

func (o PartialBookDepthStreamsResponseAItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PartialBookDepthStreamsResponseAItem) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *PartialBookDepthStreamsResponseAItem) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullablePartialBookDepthStreamsResponseAItem struct {
	value PartialBookDepthStreamsResponseAItem
	isSet bool
}

func (v NullablePartialBookDepthStreamsResponseAItem) Get() PartialBookDepthStreamsResponseAItem {
	return v.value
}

func (v *NullablePartialBookDepthStreamsResponseAItem) Set(val PartialBookDepthStreamsResponseAItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePartialBookDepthStreamsResponseAItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePartialBookDepthStreamsResponseAItem) Unset() {
	v.value = PartialBookDepthStreamsResponseAItem{}
	v.isSet = false
}

func NewNullablePartialBookDepthStreamsResponseAItem(val PartialBookDepthStreamsResponseAItem) *NullablePartialBookDepthStreamsResponseAItem {
	return &NullablePartialBookDepthStreamsResponseAItem{value: val, isSet: true}
}

func (v NullablePartialBookDepthStreamsResponseAItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartialBookDepthStreamsResponseAItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
