/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryCmModifyOrderHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryCmModifyOrderHistoryResponse{}

// QueryCmModifyOrderHistoryResponse struct for QueryCmModifyOrderHistoryResponse
type QueryCmModifyOrderHistoryResponse struct {
	Items []QueryCmModifyOrderHistoryResponseInner
}

// NewQueryCmModifyOrderHistoryResponse instantiates a new QueryCmModifyOrderHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryCmModifyOrderHistoryResponse() *QueryCmModifyOrderHistoryResponse {
	this := QueryCmModifyOrderHistoryResponse{}
	return &this
}

// NewQueryCmModifyOrderHistoryResponseWithDefaults instantiates a new QueryCmModifyOrderHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryCmModifyOrderHistoryResponseWithDefaults() *QueryCmModifyOrderHistoryResponse {
	this := QueryCmModifyOrderHistoryResponse{}
	return &this
}

func (o QueryCmModifyOrderHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryCmModifyOrderHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryCmModifyOrderHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryCmModifyOrderHistoryResponse struct {
	value QueryCmModifyOrderHistoryResponse
	isSet bool
}

func (v NullableQueryCmModifyOrderHistoryResponse) Get() QueryCmModifyOrderHistoryResponse {
	return v.value
}

func (v *NullableQueryCmModifyOrderHistoryResponse) Set(val QueryCmModifyOrderHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryCmModifyOrderHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryCmModifyOrderHistoryResponse) Unset() {
	v.value = QueryCmModifyOrderHistoryResponse{}
	v.isSet = false
}

func NewNullableQueryCmModifyOrderHistoryResponse(val QueryCmModifyOrderHistoryResponse) *NullableQueryCmModifyOrderHistoryResponse {
	return &NullableQueryCmModifyOrderHistoryResponse{value: val, isSet: true}
}

func (v NullableQueryCmModifyOrderHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryCmModifyOrderHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
