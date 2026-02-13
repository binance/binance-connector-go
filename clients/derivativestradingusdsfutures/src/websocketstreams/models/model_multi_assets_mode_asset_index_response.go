/*
Binance Derivatives Trading USDS Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the MultiAssetsModeAssetIndexResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MultiAssetsModeAssetIndexResponse{}

// MultiAssetsModeAssetIndexResponse struct for MultiAssetsModeAssetIndexResponse
type MultiAssetsModeAssetIndexResponse struct {
	Items []MultiAssetsModeAssetIndexResponseInner
}

// NewMultiAssetsModeAssetIndexResponse instantiates a new MultiAssetsModeAssetIndexResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiAssetsModeAssetIndexResponse() *MultiAssetsModeAssetIndexResponse {
	this := MultiAssetsModeAssetIndexResponse{}
	return &this
}

// NewMultiAssetsModeAssetIndexResponseWithDefaults instantiates a new MultiAssetsModeAssetIndexResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiAssetsModeAssetIndexResponseWithDefaults() *MultiAssetsModeAssetIndexResponse {
	this := MultiAssetsModeAssetIndexResponse{}
	return &this
}

func (o MultiAssetsModeAssetIndexResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultiAssetsModeAssetIndexResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *MultiAssetsModeAssetIndexResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableMultiAssetsModeAssetIndexResponse struct {
	value MultiAssetsModeAssetIndexResponse
	isSet bool
}

func (v NullableMultiAssetsModeAssetIndexResponse) Get() MultiAssetsModeAssetIndexResponse {
	return v.value
}

func (v *NullableMultiAssetsModeAssetIndexResponse) Set(val MultiAssetsModeAssetIndexResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiAssetsModeAssetIndexResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiAssetsModeAssetIndexResponse) Unset() {
	v.value = MultiAssetsModeAssetIndexResponse{}
	v.isSet = false
}

func NewNullableMultiAssetsModeAssetIndexResponse(val MultiAssetsModeAssetIndexResponse) *NullableMultiAssetsModeAssetIndexResponse {
	return &NullableMultiAssetsModeAssetIndexResponse{value: val, isSet: true}
}

func (v NullableMultiAssetsModeAssetIndexResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiAssetsModeAssetIndexResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
