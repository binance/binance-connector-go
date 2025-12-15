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

// checks if the QueryAllCurrentUmOpenConditionalOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryAllCurrentUmOpenConditionalOrdersResponse{}

// QueryAllCurrentUmOpenConditionalOrdersResponse struct for QueryAllCurrentUmOpenConditionalOrdersResponse
type QueryAllCurrentUmOpenConditionalOrdersResponse struct {
	Items []QueryAllCurrentUmOpenConditionalOrdersResponseInner
}

// NewQueryAllCurrentUmOpenConditionalOrdersResponse instantiates a new QueryAllCurrentUmOpenConditionalOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryAllCurrentUmOpenConditionalOrdersResponse() *QueryAllCurrentUmOpenConditionalOrdersResponse {
	this := QueryAllCurrentUmOpenConditionalOrdersResponse{}
	return &this
}

// NewQueryAllCurrentUmOpenConditionalOrdersResponseWithDefaults instantiates a new QueryAllCurrentUmOpenConditionalOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryAllCurrentUmOpenConditionalOrdersResponseWithDefaults() *QueryAllCurrentUmOpenConditionalOrdersResponse {
	this := QueryAllCurrentUmOpenConditionalOrdersResponse{}
	return &this
}

func (o QueryAllCurrentUmOpenConditionalOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryAllCurrentUmOpenConditionalOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryAllCurrentUmOpenConditionalOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryAllCurrentUmOpenConditionalOrdersResponse struct {
	value QueryAllCurrentUmOpenConditionalOrdersResponse
	isSet bool
}

func (v NullableQueryAllCurrentUmOpenConditionalOrdersResponse) Get() QueryAllCurrentUmOpenConditionalOrdersResponse {
	return v.value
}

func (v *NullableQueryAllCurrentUmOpenConditionalOrdersResponse) Set(val QueryAllCurrentUmOpenConditionalOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryAllCurrentUmOpenConditionalOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryAllCurrentUmOpenConditionalOrdersResponse) Unset() {
	v.value = QueryAllCurrentUmOpenConditionalOrdersResponse{}
	v.isSet = false
}

func NewNullableQueryAllCurrentUmOpenConditionalOrdersResponse(val QueryAllCurrentUmOpenConditionalOrdersResponse) *NullableQueryAllCurrentUmOpenConditionalOrdersResponse {
	return &NullableQueryAllCurrentUmOpenConditionalOrdersResponse{value: val, isSet: true}
}

func (v NullableQueryAllCurrentUmOpenConditionalOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryAllCurrentUmOpenConditionalOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
