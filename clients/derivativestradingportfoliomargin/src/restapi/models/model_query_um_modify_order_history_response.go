/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryUmModifyOrderHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryUmModifyOrderHistoryResponse{}

// QueryUmModifyOrderHistoryResponse struct for QueryUmModifyOrderHistoryResponse
type QueryUmModifyOrderHistoryResponse struct {
	Items []QueryUmModifyOrderHistoryResponseInner
}

// NewQueryUmModifyOrderHistoryResponse instantiates a new QueryUmModifyOrderHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryUmModifyOrderHistoryResponse() *QueryUmModifyOrderHistoryResponse {
	this := QueryUmModifyOrderHistoryResponse{}
	return &this
}

// NewQueryUmModifyOrderHistoryResponseWithDefaults instantiates a new QueryUmModifyOrderHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryUmModifyOrderHistoryResponseWithDefaults() *QueryUmModifyOrderHistoryResponse {
	this := QueryUmModifyOrderHistoryResponse{}
	return &this
}

func (o QueryUmModifyOrderHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryUmModifyOrderHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryUmModifyOrderHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryUmModifyOrderHistoryResponse struct {
	value QueryUmModifyOrderHistoryResponse
	isSet bool
}

func (v NullableQueryUmModifyOrderHistoryResponse) Get() QueryUmModifyOrderHistoryResponse {
	return v.value
}

func (v *NullableQueryUmModifyOrderHistoryResponse) Set(val QueryUmModifyOrderHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryUmModifyOrderHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryUmModifyOrderHistoryResponse) Unset() {
	v.value = QueryUmModifyOrderHistoryResponse{}
	v.isSet = false
}

func NewNullableQueryUmModifyOrderHistoryResponse(val QueryUmModifyOrderHistoryResponse) *NullableQueryUmModifyOrderHistoryResponse {
	return &NullableQueryUmModifyOrderHistoryResponse{value: val, isSet: true}
}

func (v NullableQueryUmModifyOrderHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryUmModifyOrderHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
