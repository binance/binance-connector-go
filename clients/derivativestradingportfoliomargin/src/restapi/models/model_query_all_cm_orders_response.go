/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryAllCmOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryAllCmOrdersResponse{}

// QueryAllCmOrdersResponse struct for QueryAllCmOrdersResponse
type QueryAllCmOrdersResponse struct {
	Items []QueryAllCmOrdersResponseInner
}

// NewQueryAllCmOrdersResponse instantiates a new QueryAllCmOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryAllCmOrdersResponse() *QueryAllCmOrdersResponse {
	this := QueryAllCmOrdersResponse{}
	return &this
}

// NewQueryAllCmOrdersResponseWithDefaults instantiates a new QueryAllCmOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryAllCmOrdersResponseWithDefaults() *QueryAllCmOrdersResponse {
	this := QueryAllCmOrdersResponse{}
	return &this
}

func (o QueryAllCmOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryAllCmOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryAllCmOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryAllCmOrdersResponse struct {
	value QueryAllCmOrdersResponse
	isSet bool
}

func (v NullableQueryAllCmOrdersResponse) Get() QueryAllCmOrdersResponse {
	return v.value
}

func (v *NullableQueryAllCmOrdersResponse) Set(val QueryAllCmOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryAllCmOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryAllCmOrdersResponse) Unset() {
	v.value = QueryAllCmOrdersResponse{}
	v.isSet = false
}

func NewNullableQueryAllCmOrdersResponse(val QueryAllCmOrdersResponse) *NullableQueryAllCmOrdersResponse {
	return &NullableQueryAllCmOrdersResponse{value: val, isSet: true}
}

func (v NullableQueryAllCmOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryAllCmOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
