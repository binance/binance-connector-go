/*
Binance Derivatives Trading COIN Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the DiffBookDepthStreamsResponseAItem type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DiffBookDepthStreamsResponseAItem{}

// DiffBookDepthStreamsResponseAItem struct for DiffBookDepthStreamsResponseAItem
type DiffBookDepthStreamsResponseAItem struct {
	Items []string
}

// NewDiffBookDepthStreamsResponseAItem instantiates a new DiffBookDepthStreamsResponseAItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiffBookDepthStreamsResponseAItem() *DiffBookDepthStreamsResponseAItem {
	this := DiffBookDepthStreamsResponseAItem{}
	return &this
}

// NewDiffBookDepthStreamsResponseAItemWithDefaults instantiates a new DiffBookDepthStreamsResponseAItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiffBookDepthStreamsResponseAItemWithDefaults() *DiffBookDepthStreamsResponseAItem {
	this := DiffBookDepthStreamsResponseAItem{}
	return &this
}

func (o DiffBookDepthStreamsResponseAItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiffBookDepthStreamsResponseAItem) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *DiffBookDepthStreamsResponseAItem) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableDiffBookDepthStreamsResponseAItem struct {
	value DiffBookDepthStreamsResponseAItem
	isSet bool
}

func (v NullableDiffBookDepthStreamsResponseAItem) Get() DiffBookDepthStreamsResponseAItem {
	return v.value
}

func (v *NullableDiffBookDepthStreamsResponseAItem) Set(val DiffBookDepthStreamsResponseAItem) {
	v.value = val
	v.isSet = true
}

func (v NullableDiffBookDepthStreamsResponseAItem) IsSet() bool {
	return v.isSet
}

func (v *NullableDiffBookDepthStreamsResponseAItem) Unset() {
	v.value = DiffBookDepthStreamsResponseAItem{}
	v.isSet = false
}

func NewNullableDiffBookDepthStreamsResponseAItem(val DiffBookDepthStreamsResponseAItem) *NullableDiffBookDepthStreamsResponseAItem {
	return &NullableDiffBookDepthStreamsResponseAItem{value: val, isSet: true}
}

func (v NullableDiffBookDepthStreamsResponseAItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiffBookDepthStreamsResponseAItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
