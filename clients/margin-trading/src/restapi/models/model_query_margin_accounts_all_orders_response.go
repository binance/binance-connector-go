/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryMarginAccountsAllOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryMarginAccountsAllOrdersResponse{}

// QueryMarginAccountsAllOrdersResponse struct for QueryMarginAccountsAllOrdersResponse
type QueryMarginAccountsAllOrdersResponse struct {
	Items []QueryMarginAccountsAllOrdersResponseInner
}

// NewQueryMarginAccountsAllOrdersResponse instantiates a new QueryMarginAccountsAllOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryMarginAccountsAllOrdersResponse() *QueryMarginAccountsAllOrdersResponse {
	this := QueryMarginAccountsAllOrdersResponse{}
	return &this
}

// NewQueryMarginAccountsAllOrdersResponseWithDefaults instantiates a new QueryMarginAccountsAllOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryMarginAccountsAllOrdersResponseWithDefaults() *QueryMarginAccountsAllOrdersResponse {
	this := QueryMarginAccountsAllOrdersResponse{}
	return &this
}

func (o QueryMarginAccountsAllOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryMarginAccountsAllOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryMarginAccountsAllOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryMarginAccountsAllOrdersResponse struct {
	value QueryMarginAccountsAllOrdersResponse
	isSet bool
}

func (v NullableQueryMarginAccountsAllOrdersResponse) Get() QueryMarginAccountsAllOrdersResponse {
	return v.value
}

func (v *NullableQueryMarginAccountsAllOrdersResponse) Set(val QueryMarginAccountsAllOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryMarginAccountsAllOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryMarginAccountsAllOrdersResponse) Unset() {
	v.value = QueryMarginAccountsAllOrdersResponse{}
	v.isSet = false
}

func NewNullableQueryMarginAccountsAllOrdersResponse(val QueryMarginAccountsAllOrdersResponse) *NullableQueryMarginAccountsAllOrdersResponse {
	return &NullableQueryMarginAccountsAllOrdersResponse{value: val, isSet: true}
}

func (v NullableQueryMarginAccountsAllOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryMarginAccountsAllOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
