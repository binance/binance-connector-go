/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryAllCurrentUmOpenOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryAllCurrentUmOpenOrdersResponse{}

// QueryAllCurrentUmOpenOrdersResponse struct for QueryAllCurrentUmOpenOrdersResponse
type QueryAllCurrentUmOpenOrdersResponse struct {
	Items []QueryAllCurrentUmOpenOrdersResponseInner
}

// NewQueryAllCurrentUmOpenOrdersResponse instantiates a new QueryAllCurrentUmOpenOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryAllCurrentUmOpenOrdersResponse() *QueryAllCurrentUmOpenOrdersResponse {
	this := QueryAllCurrentUmOpenOrdersResponse{}
	return &this
}

// NewQueryAllCurrentUmOpenOrdersResponseWithDefaults instantiates a new QueryAllCurrentUmOpenOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryAllCurrentUmOpenOrdersResponseWithDefaults() *QueryAllCurrentUmOpenOrdersResponse {
	this := QueryAllCurrentUmOpenOrdersResponse{}
	return &this
}

func (o QueryAllCurrentUmOpenOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryAllCurrentUmOpenOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryAllCurrentUmOpenOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryAllCurrentUmOpenOrdersResponse struct {
	value QueryAllCurrentUmOpenOrdersResponse
	isSet bool
}

func (v NullableQueryAllCurrentUmOpenOrdersResponse) Get() QueryAllCurrentUmOpenOrdersResponse {
	return v.value
}

func (v *NullableQueryAllCurrentUmOpenOrdersResponse) Set(val QueryAllCurrentUmOpenOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryAllCurrentUmOpenOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryAllCurrentUmOpenOrdersResponse) Unset() {
	v.value = QueryAllCurrentUmOpenOrdersResponse{}
	v.isSet = false
}

func NewNullableQueryAllCurrentUmOpenOrdersResponse(val QueryAllCurrentUmOpenOrdersResponse) *NullableQueryAllCurrentUmOpenOrdersResponse {
	return &NullableQueryAllCurrentUmOpenOrdersResponse{value: val, isSet: true}
}

func (v NullableQueryAllCurrentUmOpenOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryAllCurrentUmOpenOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
