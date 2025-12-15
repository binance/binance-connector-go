/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryAllCurrentCmOpenConditionalOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryAllCurrentCmOpenConditionalOrdersResponse{}

// QueryAllCurrentCmOpenConditionalOrdersResponse struct for QueryAllCurrentCmOpenConditionalOrdersResponse
type QueryAllCurrentCmOpenConditionalOrdersResponse struct {
	Items []QueryAllCurrentCmOpenConditionalOrdersResponseInner
}

// NewQueryAllCurrentCmOpenConditionalOrdersResponse instantiates a new QueryAllCurrentCmOpenConditionalOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryAllCurrentCmOpenConditionalOrdersResponse() *QueryAllCurrentCmOpenConditionalOrdersResponse {
	this := QueryAllCurrentCmOpenConditionalOrdersResponse{}
	return &this
}

// NewQueryAllCurrentCmOpenConditionalOrdersResponseWithDefaults instantiates a new QueryAllCurrentCmOpenConditionalOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryAllCurrentCmOpenConditionalOrdersResponseWithDefaults() *QueryAllCurrentCmOpenConditionalOrdersResponse {
	this := QueryAllCurrentCmOpenConditionalOrdersResponse{}
	return &this
}

func (o QueryAllCurrentCmOpenConditionalOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryAllCurrentCmOpenConditionalOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryAllCurrentCmOpenConditionalOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryAllCurrentCmOpenConditionalOrdersResponse struct {
	value QueryAllCurrentCmOpenConditionalOrdersResponse
	isSet bool
}

func (v NullableQueryAllCurrentCmOpenConditionalOrdersResponse) Get() QueryAllCurrentCmOpenConditionalOrdersResponse {
	return v.value
}

func (v *NullableQueryAllCurrentCmOpenConditionalOrdersResponse) Set(val QueryAllCurrentCmOpenConditionalOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryAllCurrentCmOpenConditionalOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryAllCurrentCmOpenConditionalOrdersResponse) Unset() {
	v.value = QueryAllCurrentCmOpenConditionalOrdersResponse{}
	v.isSet = false
}

func NewNullableQueryAllCurrentCmOpenConditionalOrdersResponse(val QueryAllCurrentCmOpenConditionalOrdersResponse) *NullableQueryAllCurrentCmOpenConditionalOrdersResponse {
	return &NullableQueryAllCurrentCmOpenConditionalOrdersResponse{value: val, isSet: true}
}

func (v NullableQueryAllCurrentCmOpenConditionalOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryAllCurrentCmOpenConditionalOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
