/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryOptionOrderHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryOptionOrderHistoryResponse{}

// QueryOptionOrderHistoryResponse struct for QueryOptionOrderHistoryResponse
type QueryOptionOrderHistoryResponse struct {
	Items []QueryOptionOrderHistoryResponseInner
}

// NewQueryOptionOrderHistoryResponse instantiates a new QueryOptionOrderHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryOptionOrderHistoryResponse() *QueryOptionOrderHistoryResponse {
	this := QueryOptionOrderHistoryResponse{}
	return &this
}

// NewQueryOptionOrderHistoryResponseWithDefaults instantiates a new QueryOptionOrderHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryOptionOrderHistoryResponseWithDefaults() *QueryOptionOrderHistoryResponse {
	this := QueryOptionOrderHistoryResponse{}
	return &this
}

func (o QueryOptionOrderHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryOptionOrderHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryOptionOrderHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryOptionOrderHistoryResponse struct {
	value QueryOptionOrderHistoryResponse
	isSet bool
}

func (v NullableQueryOptionOrderHistoryResponse) Get() QueryOptionOrderHistoryResponse {
	return v.value
}

func (v *NullableQueryOptionOrderHistoryResponse) Set(val QueryOptionOrderHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryOptionOrderHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryOptionOrderHistoryResponse) Unset() {
	v.value = QueryOptionOrderHistoryResponse{}
	v.isSet = false
}

func NewNullableQueryOptionOrderHistoryResponse(val QueryOptionOrderHistoryResponse) *NullableQueryOptionOrderHistoryResponse {
	return &NullableQueryOptionOrderHistoryResponse{value: val, isSet: true}
}

func (v NullableQueryOptionOrderHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryOptionOrderHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
