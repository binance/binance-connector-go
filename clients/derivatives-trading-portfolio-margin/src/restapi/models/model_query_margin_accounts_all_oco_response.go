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

// checks if the QueryMarginAccountsAllOcoResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryMarginAccountsAllOcoResponse{}

// QueryMarginAccountsAllOcoResponse struct for QueryMarginAccountsAllOcoResponse
type QueryMarginAccountsAllOcoResponse struct {
	Items []QueryMarginAccountsAllOcoResponseInner
}

// NewQueryMarginAccountsAllOcoResponse instantiates a new QueryMarginAccountsAllOcoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryMarginAccountsAllOcoResponse() *QueryMarginAccountsAllOcoResponse {
	this := QueryMarginAccountsAllOcoResponse{}
	return &this
}

// NewQueryMarginAccountsAllOcoResponseWithDefaults instantiates a new QueryMarginAccountsAllOcoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryMarginAccountsAllOcoResponseWithDefaults() *QueryMarginAccountsAllOcoResponse {
	this := QueryMarginAccountsAllOcoResponse{}
	return &this
}

func (o QueryMarginAccountsAllOcoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryMarginAccountsAllOcoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryMarginAccountsAllOcoResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryMarginAccountsAllOcoResponse struct {
	value QueryMarginAccountsAllOcoResponse
	isSet bool
}

func (v NullableQueryMarginAccountsAllOcoResponse) Get() QueryMarginAccountsAllOcoResponse {
	return v.value
}

func (v *NullableQueryMarginAccountsAllOcoResponse) Set(val QueryMarginAccountsAllOcoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryMarginAccountsAllOcoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryMarginAccountsAllOcoResponse) Unset() {
	v.value = QueryMarginAccountsAllOcoResponse{}
	v.isSet = false
}

func NewNullableQueryMarginAccountsAllOcoResponse(val QueryMarginAccountsAllOcoResponse) *NullableQueryMarginAccountsAllOcoResponse {
	return &NullableQueryMarginAccountsAllOcoResponse{value: val, isSet: true}
}

func (v NullableQueryMarginAccountsAllOcoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryMarginAccountsAllOcoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
