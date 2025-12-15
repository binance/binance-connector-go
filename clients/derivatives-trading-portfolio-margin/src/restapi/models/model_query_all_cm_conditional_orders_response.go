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

// checks if the QueryAllCmConditionalOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryAllCmConditionalOrdersResponse{}

// QueryAllCmConditionalOrdersResponse struct for QueryAllCmConditionalOrdersResponse
type QueryAllCmConditionalOrdersResponse struct {
	Items []QueryAllCmConditionalOrdersResponseInner
}

// NewQueryAllCmConditionalOrdersResponse instantiates a new QueryAllCmConditionalOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryAllCmConditionalOrdersResponse() *QueryAllCmConditionalOrdersResponse {
	this := QueryAllCmConditionalOrdersResponse{}
	return &this
}

// NewQueryAllCmConditionalOrdersResponseWithDefaults instantiates a new QueryAllCmConditionalOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryAllCmConditionalOrdersResponseWithDefaults() *QueryAllCmConditionalOrdersResponse {
	this := QueryAllCmConditionalOrdersResponse{}
	return &this
}

func (o QueryAllCmConditionalOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryAllCmConditionalOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryAllCmConditionalOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryAllCmConditionalOrdersResponse struct {
	value QueryAllCmConditionalOrdersResponse
	isSet bool
}

func (v NullableQueryAllCmConditionalOrdersResponse) Get() QueryAllCmConditionalOrdersResponse {
	return v.value
}

func (v *NullableQueryAllCmConditionalOrdersResponse) Set(val QueryAllCmConditionalOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryAllCmConditionalOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryAllCmConditionalOrdersResponse) Unset() {
	v.value = QueryAllCmConditionalOrdersResponse{}
	v.isSet = false
}

func NewNullableQueryAllCmConditionalOrdersResponse(val QueryAllCmConditionalOrdersResponse) *NullableQueryAllCmConditionalOrdersResponse {
	return &NullableQueryAllCmConditionalOrdersResponse{value: val, isSet: true}
}

func (v NullableQueryAllCmConditionalOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryAllCmConditionalOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
