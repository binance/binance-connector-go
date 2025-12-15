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

// checks if the QueryUsersCmForceOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryUsersCmForceOrdersResponse{}

// QueryUsersCmForceOrdersResponse struct for QueryUsersCmForceOrdersResponse
type QueryUsersCmForceOrdersResponse struct {
	Items []QueryUsersCmForceOrdersResponseInner
}

// NewQueryUsersCmForceOrdersResponse instantiates a new QueryUsersCmForceOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryUsersCmForceOrdersResponse() *QueryUsersCmForceOrdersResponse {
	this := QueryUsersCmForceOrdersResponse{}
	return &this
}

// NewQueryUsersCmForceOrdersResponseWithDefaults instantiates a new QueryUsersCmForceOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryUsersCmForceOrdersResponseWithDefaults() *QueryUsersCmForceOrdersResponse {
	this := QueryUsersCmForceOrdersResponse{}
	return &this
}

func (o QueryUsersCmForceOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryUsersCmForceOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryUsersCmForceOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryUsersCmForceOrdersResponse struct {
	value QueryUsersCmForceOrdersResponse
	isSet bool
}

func (v NullableQueryUsersCmForceOrdersResponse) Get() QueryUsersCmForceOrdersResponse {
	return v.value
}

func (v *NullableQueryUsersCmForceOrdersResponse) Set(val QueryUsersCmForceOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryUsersCmForceOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryUsersCmForceOrdersResponse) Unset() {
	v.value = QueryUsersCmForceOrdersResponse{}
	v.isSet = false
}

func NewNullableQueryUsersCmForceOrdersResponse(val QueryUsersCmForceOrdersResponse) *NullableQueryUsersCmForceOrdersResponse {
	return &NullableQueryUsersCmForceOrdersResponse{value: val, isSet: true}
}

func (v NullableQueryUsersCmForceOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryUsersCmForceOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
