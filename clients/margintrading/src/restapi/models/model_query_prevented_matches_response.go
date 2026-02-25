/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryPreventedMatchesResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryPreventedMatchesResponse{}

// QueryPreventedMatchesResponse struct for QueryPreventedMatchesResponse
type QueryPreventedMatchesResponse struct {
	Items []QueryPreventedMatchesResponseInner
}

// NewQueryPreventedMatchesResponse instantiates a new QueryPreventedMatchesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryPreventedMatchesResponse() *QueryPreventedMatchesResponse {
	this := QueryPreventedMatchesResponse{}
	return &this
}

// NewQueryPreventedMatchesResponseWithDefaults instantiates a new QueryPreventedMatchesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryPreventedMatchesResponseWithDefaults() *QueryPreventedMatchesResponse {
	this := QueryPreventedMatchesResponse{}
	return &this
}

func (o QueryPreventedMatchesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryPreventedMatchesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryPreventedMatchesResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryPreventedMatchesResponse struct {
	value QueryPreventedMatchesResponse
	isSet bool
}

func (v NullableQueryPreventedMatchesResponse) Get() QueryPreventedMatchesResponse {
	return v.value
}

func (v *NullableQueryPreventedMatchesResponse) Set(val QueryPreventedMatchesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryPreventedMatchesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryPreventedMatchesResponse) Unset() {
	v.value = QueryPreventedMatchesResponse{}
	v.isSet = false
}

func NewNullableQueryPreventedMatchesResponse(val QueryPreventedMatchesResponse) *NullableQueryPreventedMatchesResponse {
	return &NullableQueryPreventedMatchesResponse{value: val, isSet: true}
}

func (v NullableQueryPreventedMatchesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryPreventedMatchesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
