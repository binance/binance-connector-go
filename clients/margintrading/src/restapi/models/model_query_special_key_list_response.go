/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QuerySpecialKeyListResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QuerySpecialKeyListResponse{}

// QuerySpecialKeyListResponse struct for QuerySpecialKeyListResponse
type QuerySpecialKeyListResponse struct {
	Items []QuerySpecialKeyListResponseInner
}

// NewQuerySpecialKeyListResponse instantiates a new QuerySpecialKeyListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuerySpecialKeyListResponse() *QuerySpecialKeyListResponse {
	this := QuerySpecialKeyListResponse{}
	return &this
}

// NewQuerySpecialKeyListResponseWithDefaults instantiates a new QuerySpecialKeyListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuerySpecialKeyListResponseWithDefaults() *QuerySpecialKeyListResponse {
	this := QuerySpecialKeyListResponse{}
	return &this
}

func (o QuerySpecialKeyListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuerySpecialKeyListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QuerySpecialKeyListResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQuerySpecialKeyListResponse struct {
	value QuerySpecialKeyListResponse
	isSet bool
}

func (v NullableQuerySpecialKeyListResponse) Get() QuerySpecialKeyListResponse {
	return v.value
}

func (v *NullableQuerySpecialKeyListResponse) Set(val QuerySpecialKeyListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQuerySpecialKeyListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQuerySpecialKeyListResponse) Unset() {
	v.value = QuerySpecialKeyListResponse{}
	v.isSet = false
}

func NewNullableQuerySpecialKeyListResponse(val QuerySpecialKeyListResponse) *NullableQuerySpecialKeyListResponse {
	return &NullableQuerySpecialKeyListResponse{value: val, isSet: true}
}

func (v NullableQuerySpecialKeyListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuerySpecialKeyListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
