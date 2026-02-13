/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryUserRateLimitResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryUserRateLimitResponse{}

// QueryUserRateLimitResponse struct for QueryUserRateLimitResponse
type QueryUserRateLimitResponse struct {
	Items []QueryUserRateLimitResponseInner
}

// NewQueryUserRateLimitResponse instantiates a new QueryUserRateLimitResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryUserRateLimitResponse() *QueryUserRateLimitResponse {
	this := QueryUserRateLimitResponse{}
	return &this
}

// NewQueryUserRateLimitResponseWithDefaults instantiates a new QueryUserRateLimitResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryUserRateLimitResponseWithDefaults() *QueryUserRateLimitResponse {
	this := QueryUserRateLimitResponse{}
	return &this
}

func (o QueryUserRateLimitResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryUserRateLimitResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryUserRateLimitResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryUserRateLimitResponse struct {
	value QueryUserRateLimitResponse
	isSet bool
}

func (v NullableQueryUserRateLimitResponse) Get() QueryUserRateLimitResponse {
	return v.value
}

func (v *NullableQueryUserRateLimitResponse) Set(val QueryUserRateLimitResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryUserRateLimitResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryUserRateLimitResponse) Unset() {
	v.value = QueryUserRateLimitResponse{}
	v.isSet = false
}

func NewNullableQueryUserRateLimitResponse(val QueryUserRateLimitResponse) *NullableQueryUserRateLimitResponse {
	return &NullableQueryUserRateLimitResponse{value: val, isSet: true}
}

func (v NullableQueryUserRateLimitResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryUserRateLimitResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
