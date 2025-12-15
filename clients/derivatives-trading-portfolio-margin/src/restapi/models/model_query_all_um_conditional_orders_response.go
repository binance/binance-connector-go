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

// checks if the QueryAllUmConditionalOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryAllUmConditionalOrdersResponse{}

// QueryAllUmConditionalOrdersResponse struct for QueryAllUmConditionalOrdersResponse
type QueryAllUmConditionalOrdersResponse struct {
	Items []QueryAllUmConditionalOrdersResponseInner
}

// NewQueryAllUmConditionalOrdersResponse instantiates a new QueryAllUmConditionalOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryAllUmConditionalOrdersResponse() *QueryAllUmConditionalOrdersResponse {
	this := QueryAllUmConditionalOrdersResponse{}
	return &this
}

// NewQueryAllUmConditionalOrdersResponseWithDefaults instantiates a new QueryAllUmConditionalOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryAllUmConditionalOrdersResponseWithDefaults() *QueryAllUmConditionalOrdersResponse {
	this := QueryAllUmConditionalOrdersResponse{}
	return &this
}

func (o QueryAllUmConditionalOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryAllUmConditionalOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryAllUmConditionalOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryAllUmConditionalOrdersResponse struct {
	value QueryAllUmConditionalOrdersResponse
	isSet bool
}

func (v NullableQueryAllUmConditionalOrdersResponse) Get() QueryAllUmConditionalOrdersResponse {
	return v.value
}

func (v *NullableQueryAllUmConditionalOrdersResponse) Set(val QueryAllUmConditionalOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryAllUmConditionalOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryAllUmConditionalOrdersResponse) Unset() {
	v.value = QueryAllUmConditionalOrdersResponse{}
	v.isSet = false
}

func NewNullableQueryAllUmConditionalOrdersResponse(val QueryAllUmConditionalOrdersResponse) *NullableQueryAllUmConditionalOrdersResponse {
	return &NullableQueryAllUmConditionalOrdersResponse{value: val, isSet: true}
}

func (v NullableQueryAllUmConditionalOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryAllUmConditionalOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
