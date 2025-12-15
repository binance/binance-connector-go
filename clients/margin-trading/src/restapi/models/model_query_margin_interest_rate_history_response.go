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

// checks if the QueryMarginInterestRateHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryMarginInterestRateHistoryResponse{}

// QueryMarginInterestRateHistoryResponse struct for QueryMarginInterestRateHistoryResponse
type QueryMarginInterestRateHistoryResponse struct {
	Items []QueryMarginInterestRateHistoryResponseInner
}

// NewQueryMarginInterestRateHistoryResponse instantiates a new QueryMarginInterestRateHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryMarginInterestRateHistoryResponse() *QueryMarginInterestRateHistoryResponse {
	this := QueryMarginInterestRateHistoryResponse{}
	return &this
}

// NewQueryMarginInterestRateHistoryResponseWithDefaults instantiates a new QueryMarginInterestRateHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryMarginInterestRateHistoryResponseWithDefaults() *QueryMarginInterestRateHistoryResponse {
	this := QueryMarginInterestRateHistoryResponse{}
	return &this
}

func (o QueryMarginInterestRateHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryMarginInterestRateHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryMarginInterestRateHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryMarginInterestRateHistoryResponse struct {
	value QueryMarginInterestRateHistoryResponse
	isSet bool
}

func (v NullableQueryMarginInterestRateHistoryResponse) Get() QueryMarginInterestRateHistoryResponse {
	return v.value
}

func (v *NullableQueryMarginInterestRateHistoryResponse) Set(val QueryMarginInterestRateHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryMarginInterestRateHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryMarginInterestRateHistoryResponse) Unset() {
	v.value = QueryMarginInterestRateHistoryResponse{}
	v.isSet = false
}

func NewNullableQueryMarginInterestRateHistoryResponse(val QueryMarginInterestRateHistoryResponse) *NullableQueryMarginInterestRateHistoryResponse {
	return &NullableQueryMarginInterestRateHistoryResponse{value: val, isSet: true}
}

func (v NullableQueryMarginInterestRateHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryMarginInterestRateHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
