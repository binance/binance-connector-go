/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryUmAlgoOrderHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryUmAlgoOrderHistoryResponse{}

// QueryUmAlgoOrderHistoryResponse struct for QueryUmAlgoOrderHistoryResponse
type QueryUmAlgoOrderHistoryResponse struct {
	Items []QueryUmAlgoOrderHistoryResponseInner
}

// NewQueryUmAlgoOrderHistoryResponse instantiates a new QueryUmAlgoOrderHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryUmAlgoOrderHistoryResponse() *QueryUmAlgoOrderHistoryResponse {
	this := QueryUmAlgoOrderHistoryResponse{}
	return &this
}

// NewQueryUmAlgoOrderHistoryResponseWithDefaults instantiates a new QueryUmAlgoOrderHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryUmAlgoOrderHistoryResponseWithDefaults() *QueryUmAlgoOrderHistoryResponse {
	this := QueryUmAlgoOrderHistoryResponse{}
	return &this
}

func (o QueryUmAlgoOrderHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryUmAlgoOrderHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryUmAlgoOrderHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryUmAlgoOrderHistoryResponse struct {
	value QueryUmAlgoOrderHistoryResponse
	isSet bool
}

func (v NullableQueryUmAlgoOrderHistoryResponse) Get() QueryUmAlgoOrderHistoryResponse {
	return v.value
}

func (v *NullableQueryUmAlgoOrderHistoryResponse) Set(val QueryUmAlgoOrderHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryUmAlgoOrderHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryUmAlgoOrderHistoryResponse) Unset() {
	v.value = QueryUmAlgoOrderHistoryResponse{}
	v.isSet = false
}

func NewNullableQueryUmAlgoOrderHistoryResponse(val QueryUmAlgoOrderHistoryResponse) *NullableQueryUmAlgoOrderHistoryResponse {
	return &NullableQueryUmAlgoOrderHistoryResponse{value: val, isSet: true}
}

func (v NullableQueryUmAlgoOrderHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryUmAlgoOrderHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
