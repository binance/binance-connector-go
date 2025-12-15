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

// checks if the QueryAllUmOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryAllUmOrdersResponse{}

// QueryAllUmOrdersResponse struct for QueryAllUmOrdersResponse
type QueryAllUmOrdersResponse struct {
	Items []QueryAllCurrentUmOpenOrdersResponseInner
}

// NewQueryAllUmOrdersResponse instantiates a new QueryAllUmOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryAllUmOrdersResponse() *QueryAllUmOrdersResponse {
	this := QueryAllUmOrdersResponse{}
	return &this
}

// NewQueryAllUmOrdersResponseWithDefaults instantiates a new QueryAllUmOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryAllUmOrdersResponseWithDefaults() *QueryAllUmOrdersResponse {
	this := QueryAllUmOrdersResponse{}
	return &this
}

func (o QueryAllUmOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryAllUmOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryAllUmOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryAllUmOrdersResponse struct {
	value QueryAllUmOrdersResponse
	isSet bool
}

func (v NullableQueryAllUmOrdersResponse) Get() QueryAllUmOrdersResponse {
	return v.value
}

func (v *NullableQueryAllUmOrdersResponse) Set(val QueryAllUmOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryAllUmOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryAllUmOrdersResponse) Unset() {
	v.value = QueryAllUmOrdersResponse{}
	v.isSet = false
}

func NewNullableQueryAllUmOrdersResponse(val QueryAllUmOrdersResponse) *NullableQueryAllUmOrdersResponse {
	return &NullableQueryAllUmOrdersResponse{value: val, isSet: true}
}

func (v NullableQueryAllUmOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryAllUmOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
