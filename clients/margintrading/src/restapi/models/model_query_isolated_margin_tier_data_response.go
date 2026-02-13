/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryIsolatedMarginTierDataResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryIsolatedMarginTierDataResponse{}

// QueryIsolatedMarginTierDataResponse struct for QueryIsolatedMarginTierDataResponse
type QueryIsolatedMarginTierDataResponse struct {
	Items []QueryIsolatedMarginTierDataResponseInner
}

// NewQueryIsolatedMarginTierDataResponse instantiates a new QueryIsolatedMarginTierDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryIsolatedMarginTierDataResponse() *QueryIsolatedMarginTierDataResponse {
	this := QueryIsolatedMarginTierDataResponse{}
	return &this
}

// NewQueryIsolatedMarginTierDataResponseWithDefaults instantiates a new QueryIsolatedMarginTierDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryIsolatedMarginTierDataResponseWithDefaults() *QueryIsolatedMarginTierDataResponse {
	this := QueryIsolatedMarginTierDataResponse{}
	return &this
}

func (o QueryIsolatedMarginTierDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryIsolatedMarginTierDataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryIsolatedMarginTierDataResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryIsolatedMarginTierDataResponse struct {
	value QueryIsolatedMarginTierDataResponse
	isSet bool
}

func (v NullableQueryIsolatedMarginTierDataResponse) Get() QueryIsolatedMarginTierDataResponse {
	return v.value
}

func (v *NullableQueryIsolatedMarginTierDataResponse) Set(val QueryIsolatedMarginTierDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryIsolatedMarginTierDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryIsolatedMarginTierDataResponse) Unset() {
	v.value = QueryIsolatedMarginTierDataResponse{}
	v.isSet = false
}

func NewNullableQueryIsolatedMarginTierDataResponse(val QueryIsolatedMarginTierDataResponse) *NullableQueryIsolatedMarginTierDataResponse {
	return &NullableQueryIsolatedMarginTierDataResponse{value: val, isSet: true}
}

func (v NullableQueryIsolatedMarginTierDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryIsolatedMarginTierDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
