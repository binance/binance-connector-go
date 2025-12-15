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

// checks if the QueryMarginAccountsOpenOcoResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryMarginAccountsOpenOcoResponse{}

// QueryMarginAccountsOpenOcoResponse struct for QueryMarginAccountsOpenOcoResponse
type QueryMarginAccountsOpenOcoResponse struct {
	Items []QueryMarginAccountsOpenOcoResponseInner
}

// NewQueryMarginAccountsOpenOcoResponse instantiates a new QueryMarginAccountsOpenOcoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryMarginAccountsOpenOcoResponse() *QueryMarginAccountsOpenOcoResponse {
	this := QueryMarginAccountsOpenOcoResponse{}
	return &this
}

// NewQueryMarginAccountsOpenOcoResponseWithDefaults instantiates a new QueryMarginAccountsOpenOcoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryMarginAccountsOpenOcoResponseWithDefaults() *QueryMarginAccountsOpenOcoResponse {
	this := QueryMarginAccountsOpenOcoResponse{}
	return &this
}

func (o QueryMarginAccountsOpenOcoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryMarginAccountsOpenOcoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryMarginAccountsOpenOcoResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryMarginAccountsOpenOcoResponse struct {
	value QueryMarginAccountsOpenOcoResponse
	isSet bool
}

func (v NullableQueryMarginAccountsOpenOcoResponse) Get() QueryMarginAccountsOpenOcoResponse {
	return v.value
}

func (v *NullableQueryMarginAccountsOpenOcoResponse) Set(val QueryMarginAccountsOpenOcoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryMarginAccountsOpenOcoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryMarginAccountsOpenOcoResponse) Unset() {
	v.value = QueryMarginAccountsOpenOcoResponse{}
	v.isSet = false
}

func NewNullableQueryMarginAccountsOpenOcoResponse(val QueryMarginAccountsOpenOcoResponse) *NullableQueryMarginAccountsOpenOcoResponse {
	return &NullableQueryMarginAccountsOpenOcoResponse{value: val, isSet: true}
}

func (v NullableQueryMarginAccountsOpenOcoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryMarginAccountsOpenOcoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
