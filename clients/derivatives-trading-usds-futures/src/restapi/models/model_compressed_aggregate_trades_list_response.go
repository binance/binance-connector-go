/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CompressedAggregateTradesListResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CompressedAggregateTradesListResponse{}

// CompressedAggregateTradesListResponse struct for CompressedAggregateTradesListResponse
type CompressedAggregateTradesListResponse struct {
	Items []CompressedAggregateTradesListResponseInner
}

// NewCompressedAggregateTradesListResponse instantiates a new CompressedAggregateTradesListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompressedAggregateTradesListResponse() *CompressedAggregateTradesListResponse {
	this := CompressedAggregateTradesListResponse{}
	return &this
}

// NewCompressedAggregateTradesListResponseWithDefaults instantiates a new CompressedAggregateTradesListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompressedAggregateTradesListResponseWithDefaults() *CompressedAggregateTradesListResponse {
	this := CompressedAggregateTradesListResponse{}
	return &this
}

func (o CompressedAggregateTradesListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompressedAggregateTradesListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *CompressedAggregateTradesListResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableCompressedAggregateTradesListResponse struct {
	value CompressedAggregateTradesListResponse
	isSet bool
}

func (v NullableCompressedAggregateTradesListResponse) Get() CompressedAggregateTradesListResponse {
	return v.value
}

func (v *NullableCompressedAggregateTradesListResponse) Set(val CompressedAggregateTradesListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCompressedAggregateTradesListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCompressedAggregateTradesListResponse) Unset() {
	v.value = CompressedAggregateTradesListResponse{}
	v.isSet = false
}

func NewNullableCompressedAggregateTradesListResponse(val CompressedAggregateTradesListResponse) *NullableCompressedAggregateTradesListResponse {
	return &NullableCompressedAggregateTradesListResponse{value: val, isSet: true}
}

func (v NullableCompressedAggregateTradesListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompressedAggregateTradesListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
