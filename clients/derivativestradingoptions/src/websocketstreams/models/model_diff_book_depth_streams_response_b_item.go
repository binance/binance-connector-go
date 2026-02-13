/*
Binance Derivatives Trading Options WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Options WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the DiffBookDepthStreamsResponseBItem type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DiffBookDepthStreamsResponseBItem{}

// DiffBookDepthStreamsResponseBItem struct for DiffBookDepthStreamsResponseBItem
type DiffBookDepthStreamsResponseBItem struct {
	Items []string
}

// NewDiffBookDepthStreamsResponseBItem instantiates a new DiffBookDepthStreamsResponseBItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiffBookDepthStreamsResponseBItem() *DiffBookDepthStreamsResponseBItem {
	this := DiffBookDepthStreamsResponseBItem{}
	return &this
}

// NewDiffBookDepthStreamsResponseBItemWithDefaults instantiates a new DiffBookDepthStreamsResponseBItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiffBookDepthStreamsResponseBItemWithDefaults() *DiffBookDepthStreamsResponseBItem {
	this := DiffBookDepthStreamsResponseBItem{}
	return &this
}

func (o DiffBookDepthStreamsResponseBItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiffBookDepthStreamsResponseBItem) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *DiffBookDepthStreamsResponseBItem) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableDiffBookDepthStreamsResponseBItem struct {
	value DiffBookDepthStreamsResponseBItem
	isSet bool
}

func (v NullableDiffBookDepthStreamsResponseBItem) Get() DiffBookDepthStreamsResponseBItem {
	return v.value
}

func (v *NullableDiffBookDepthStreamsResponseBItem) Set(val DiffBookDepthStreamsResponseBItem) {
	v.value = val
	v.isSet = true
}

func (v NullableDiffBookDepthStreamsResponseBItem) IsSet() bool {
	return v.isSet
}

func (v *NullableDiffBookDepthStreamsResponseBItem) Unset() {
	v.value = DiffBookDepthStreamsResponseBItem{}
	v.isSet = false
}

func NewNullableDiffBookDepthStreamsResponseBItem(val DiffBookDepthStreamsResponseBItem) *NullableDiffBookDepthStreamsResponseBItem {
	return &NullableDiffBookDepthStreamsResponseBItem{value: val, isSet: true}
}

func (v NullableDiffBookDepthStreamsResponseBItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiffBookDepthStreamsResponseBItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
