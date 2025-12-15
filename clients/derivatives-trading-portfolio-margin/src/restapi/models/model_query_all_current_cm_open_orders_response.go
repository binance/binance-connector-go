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

// checks if the QueryAllCurrentCmOpenOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryAllCurrentCmOpenOrdersResponse{}

// QueryAllCurrentCmOpenOrdersResponse struct for QueryAllCurrentCmOpenOrdersResponse
type QueryAllCurrentCmOpenOrdersResponse struct {
	Items []QueryAllCmOrdersResponseInner
}

// NewQueryAllCurrentCmOpenOrdersResponse instantiates a new QueryAllCurrentCmOpenOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryAllCurrentCmOpenOrdersResponse() *QueryAllCurrentCmOpenOrdersResponse {
	this := QueryAllCurrentCmOpenOrdersResponse{}
	return &this
}

// NewQueryAllCurrentCmOpenOrdersResponseWithDefaults instantiates a new QueryAllCurrentCmOpenOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryAllCurrentCmOpenOrdersResponseWithDefaults() *QueryAllCurrentCmOpenOrdersResponse {
	this := QueryAllCurrentCmOpenOrdersResponse{}
	return &this
}

func (o QueryAllCurrentCmOpenOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryAllCurrentCmOpenOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryAllCurrentCmOpenOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryAllCurrentCmOpenOrdersResponse struct {
	value QueryAllCurrentCmOpenOrdersResponse
	isSet bool
}

func (v NullableQueryAllCurrentCmOpenOrdersResponse) Get() QueryAllCurrentCmOpenOrdersResponse {
	return v.value
}

func (v *NullableQueryAllCurrentCmOpenOrdersResponse) Set(val QueryAllCurrentCmOpenOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryAllCurrentCmOpenOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryAllCurrentCmOpenOrdersResponse) Unset() {
	v.value = QueryAllCurrentCmOpenOrdersResponse{}
	v.isSet = false
}

func NewNullableQueryAllCurrentCmOpenOrdersResponse(val QueryAllCurrentCmOpenOrdersResponse) *NullableQueryAllCurrentCmOpenOrdersResponse {
	return &NullableQueryAllCurrentCmOpenOrdersResponse{value: val, isSet: true}
}

func (v NullableQueryAllCurrentCmOpenOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryAllCurrentCmOpenOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
