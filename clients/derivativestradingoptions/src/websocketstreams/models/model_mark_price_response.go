/*
Binance Derivatives Trading Options WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Options WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the MarkPriceResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarkPriceResponse{}

// MarkPriceResponse struct for MarkPriceResponse
type MarkPriceResponse struct {
	Items []MarkPriceResponseInner
}

// NewMarkPriceResponse instantiates a new MarkPriceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarkPriceResponse() *MarkPriceResponse {
	this := MarkPriceResponse{}
	return &this
}

// NewMarkPriceResponseWithDefaults instantiates a new MarkPriceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarkPriceResponseWithDefaults() *MarkPriceResponse {
	this := MarkPriceResponse{}
	return &this
}

func (o MarkPriceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarkPriceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *MarkPriceResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableMarkPriceResponse struct {
	value MarkPriceResponse
	isSet bool
}

func (v NullableMarkPriceResponse) Get() MarkPriceResponse {
	return v.value
}

func (v *NullableMarkPriceResponse) Set(val MarkPriceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMarkPriceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMarkPriceResponse) Unset() {
	v.value = MarkPriceResponse{}
	v.isSet = false
}

func NewNullableMarkPriceResponse(val MarkPriceResponse) *NullableMarkPriceResponse {
	return &NullableMarkPriceResponse{value: val, isSet: true}
}

func (v NullableMarkPriceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarkPriceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
