/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the IndexPriceAndMarkPriceResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &IndexPriceAndMarkPriceResponse{}

// IndexPriceAndMarkPriceResponse struct for IndexPriceAndMarkPriceResponse
type IndexPriceAndMarkPriceResponse struct {
	Items []IndexPriceAndMarkPriceResponseInner
}

// NewIndexPriceAndMarkPriceResponse instantiates a new IndexPriceAndMarkPriceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexPriceAndMarkPriceResponse() *IndexPriceAndMarkPriceResponse {
	this := IndexPriceAndMarkPriceResponse{}
	return &this
}

// NewIndexPriceAndMarkPriceResponseWithDefaults instantiates a new IndexPriceAndMarkPriceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexPriceAndMarkPriceResponseWithDefaults() *IndexPriceAndMarkPriceResponse {
	this := IndexPriceAndMarkPriceResponse{}
	return &this
}

func (o IndexPriceAndMarkPriceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndexPriceAndMarkPriceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *IndexPriceAndMarkPriceResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableIndexPriceAndMarkPriceResponse struct {
	value IndexPriceAndMarkPriceResponse
	isSet bool
}

func (v NullableIndexPriceAndMarkPriceResponse) Get() IndexPriceAndMarkPriceResponse {
	return v.value
}

func (v *NullableIndexPriceAndMarkPriceResponse) Set(val IndexPriceAndMarkPriceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexPriceAndMarkPriceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexPriceAndMarkPriceResponse) Unset() {
	v.value = IndexPriceAndMarkPriceResponse{}
	v.isSet = false
}

func NewNullableIndexPriceAndMarkPriceResponse(val IndexPriceAndMarkPriceResponse) *NullableIndexPriceAndMarkPriceResponse {
	return &NullableIndexPriceAndMarkPriceResponse{value: val, isSet: true}
}

func (v NullableIndexPriceAndMarkPriceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexPriceAndMarkPriceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
