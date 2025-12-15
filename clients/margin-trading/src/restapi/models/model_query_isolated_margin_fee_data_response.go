/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryIsolatedMarginFeeDataResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryIsolatedMarginFeeDataResponse{}

// QueryIsolatedMarginFeeDataResponse struct for QueryIsolatedMarginFeeDataResponse
type QueryIsolatedMarginFeeDataResponse struct {
	Items []QueryIsolatedMarginFeeDataResponseInner
}

// NewQueryIsolatedMarginFeeDataResponse instantiates a new QueryIsolatedMarginFeeDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryIsolatedMarginFeeDataResponse() *QueryIsolatedMarginFeeDataResponse {
	this := QueryIsolatedMarginFeeDataResponse{}
	return &this
}

// NewQueryIsolatedMarginFeeDataResponseWithDefaults instantiates a new QueryIsolatedMarginFeeDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryIsolatedMarginFeeDataResponseWithDefaults() *QueryIsolatedMarginFeeDataResponse {
	this := QueryIsolatedMarginFeeDataResponse{}
	return &this
}

func (o QueryIsolatedMarginFeeDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryIsolatedMarginFeeDataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryIsolatedMarginFeeDataResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryIsolatedMarginFeeDataResponse struct {
	value QueryIsolatedMarginFeeDataResponse
	isSet bool
}

func (v NullableQueryIsolatedMarginFeeDataResponse) Get() QueryIsolatedMarginFeeDataResponse {
	return v.value
}

func (v *NullableQueryIsolatedMarginFeeDataResponse) Set(val QueryIsolatedMarginFeeDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryIsolatedMarginFeeDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryIsolatedMarginFeeDataResponse) Unset() {
	v.value = QueryIsolatedMarginFeeDataResponse{}
	v.isSet = false
}

func NewNullableQueryIsolatedMarginFeeDataResponse(val QueryIsolatedMarginFeeDataResponse) *NullableQueryIsolatedMarginFeeDataResponse {
	return &NullableQueryIsolatedMarginFeeDataResponse{value: val, isSet: true}
}

func (v NullableQueryIsolatedMarginFeeDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryIsolatedMarginFeeDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
