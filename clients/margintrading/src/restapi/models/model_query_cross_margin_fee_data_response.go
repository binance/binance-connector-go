/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryCrossMarginFeeDataResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryCrossMarginFeeDataResponse{}

// QueryCrossMarginFeeDataResponse struct for QueryCrossMarginFeeDataResponse
type QueryCrossMarginFeeDataResponse struct {
	Items []QueryCrossMarginFeeDataResponseInner
}

// NewQueryCrossMarginFeeDataResponse instantiates a new QueryCrossMarginFeeDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryCrossMarginFeeDataResponse() *QueryCrossMarginFeeDataResponse {
	this := QueryCrossMarginFeeDataResponse{}
	return &this
}

// NewQueryCrossMarginFeeDataResponseWithDefaults instantiates a new QueryCrossMarginFeeDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryCrossMarginFeeDataResponseWithDefaults() *QueryCrossMarginFeeDataResponse {
	this := QueryCrossMarginFeeDataResponse{}
	return &this
}

func (o QueryCrossMarginFeeDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryCrossMarginFeeDataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryCrossMarginFeeDataResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryCrossMarginFeeDataResponse struct {
	value QueryCrossMarginFeeDataResponse
	isSet bool
}

func (v NullableQueryCrossMarginFeeDataResponse) Get() QueryCrossMarginFeeDataResponse {
	return v.value
}

func (v *NullableQueryCrossMarginFeeDataResponse) Set(val QueryCrossMarginFeeDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryCrossMarginFeeDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryCrossMarginFeeDataResponse) Unset() {
	v.value = QueryCrossMarginFeeDataResponse{}
	v.isSet = false
}

func NewNullableQueryCrossMarginFeeDataResponse(val QueryCrossMarginFeeDataResponse) *NullableQueryCrossMarginFeeDataResponse {
	return &NullableQueryCrossMarginFeeDataResponse{value: val, isSet: true}
}

func (v NullableQueryCrossMarginFeeDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryCrossMarginFeeDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
