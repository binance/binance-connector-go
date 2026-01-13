/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryAllMarginAccountOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryAllMarginAccountOrdersResponse{}

// QueryAllMarginAccountOrdersResponse struct for QueryAllMarginAccountOrdersResponse
type QueryAllMarginAccountOrdersResponse struct {
	Items []QueryAllMarginAccountOrdersResponseInner
}

// NewQueryAllMarginAccountOrdersResponse instantiates a new QueryAllMarginAccountOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryAllMarginAccountOrdersResponse() *QueryAllMarginAccountOrdersResponse {
	this := QueryAllMarginAccountOrdersResponse{}
	return &this
}

// NewQueryAllMarginAccountOrdersResponseWithDefaults instantiates a new QueryAllMarginAccountOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryAllMarginAccountOrdersResponseWithDefaults() *QueryAllMarginAccountOrdersResponse {
	this := QueryAllMarginAccountOrdersResponse{}
	return &this
}

func (o QueryAllMarginAccountOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryAllMarginAccountOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryAllMarginAccountOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryAllMarginAccountOrdersResponse struct {
	value QueryAllMarginAccountOrdersResponse
	isSet bool
}

func (v NullableQueryAllMarginAccountOrdersResponse) Get() QueryAllMarginAccountOrdersResponse {
	return v.value
}

func (v *NullableQueryAllMarginAccountOrdersResponse) Set(val QueryAllMarginAccountOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryAllMarginAccountOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryAllMarginAccountOrdersResponse) Unset() {
	v.value = QueryAllMarginAccountOrdersResponse{}
	v.isSet = false
}

func NewNullableQueryAllMarginAccountOrdersResponse(val QueryAllMarginAccountOrdersResponse) *NullableQueryAllMarginAccountOrdersResponse {
	return &NullableQueryAllMarginAccountOrdersResponse{value: val, isSet: true}
}

func (v NullableQueryAllMarginAccountOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryAllMarginAccountOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
