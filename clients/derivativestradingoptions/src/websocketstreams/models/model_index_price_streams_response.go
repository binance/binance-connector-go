/*
Binance Derivatives Trading Options WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Options WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the IndexPriceStreamsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &IndexPriceStreamsResponse{}

// IndexPriceStreamsResponse struct for IndexPriceStreamsResponse
type IndexPriceStreamsResponse struct {
	Items []IndexPriceStreamsResponseInner
}

// NewIndexPriceStreamsResponse instantiates a new IndexPriceStreamsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexPriceStreamsResponse() *IndexPriceStreamsResponse {
	this := IndexPriceStreamsResponse{}
	return &this
}

// NewIndexPriceStreamsResponseWithDefaults instantiates a new IndexPriceStreamsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexPriceStreamsResponseWithDefaults() *IndexPriceStreamsResponse {
	this := IndexPriceStreamsResponse{}
	return &this
}

func (o IndexPriceStreamsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndexPriceStreamsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *IndexPriceStreamsResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableIndexPriceStreamsResponse struct {
	value IndexPriceStreamsResponse
	isSet bool
}

func (v NullableIndexPriceStreamsResponse) Get() IndexPriceStreamsResponse {
	return v.value
}

func (v *NullableIndexPriceStreamsResponse) Set(val IndexPriceStreamsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexPriceStreamsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexPriceStreamsResponse) Unset() {
	v.value = IndexPriceStreamsResponse{}
	v.isSet = false
}

func NewNullableIndexPriceStreamsResponse(val IndexPriceStreamsResponse) *NullableIndexPriceStreamsResponse {
	return &NullableIndexPriceStreamsResponse{value: val, isSet: true}
}

func (v NullableIndexPriceStreamsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexPriceStreamsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
