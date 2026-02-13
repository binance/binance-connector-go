/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryUsersUmForceOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryUsersUmForceOrdersResponse{}

// QueryUsersUmForceOrdersResponse struct for QueryUsersUmForceOrdersResponse
type QueryUsersUmForceOrdersResponse struct {
	Items []QueryUsersUmForceOrdersResponseInner
}

// NewQueryUsersUmForceOrdersResponse instantiates a new QueryUsersUmForceOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryUsersUmForceOrdersResponse() *QueryUsersUmForceOrdersResponse {
	this := QueryUsersUmForceOrdersResponse{}
	return &this
}

// NewQueryUsersUmForceOrdersResponseWithDefaults instantiates a new QueryUsersUmForceOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryUsersUmForceOrdersResponseWithDefaults() *QueryUsersUmForceOrdersResponse {
	this := QueryUsersUmForceOrdersResponse{}
	return &this
}

func (o QueryUsersUmForceOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryUsersUmForceOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryUsersUmForceOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryUsersUmForceOrdersResponse struct {
	value QueryUsersUmForceOrdersResponse
	isSet bool
}

func (v NullableQueryUsersUmForceOrdersResponse) Get() QueryUsersUmForceOrdersResponse {
	return v.value
}

func (v *NullableQueryUsersUmForceOrdersResponse) Set(val QueryUsersUmForceOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryUsersUmForceOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryUsersUmForceOrdersResponse) Unset() {
	v.value = QueryUsersUmForceOrdersResponse{}
	v.isSet = false
}

func NewNullableQueryUsersUmForceOrdersResponse(val QueryUsersUmForceOrdersResponse) *NullableQueryUsersUmForceOrdersResponse {
	return &NullableQueryUsersUmForceOrdersResponse{value: val, isSet: true}
}

func (v NullableQueryUsersUmForceOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryUsersUmForceOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
