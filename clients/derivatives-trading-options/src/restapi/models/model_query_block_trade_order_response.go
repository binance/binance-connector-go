/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryBlockTradeOrderResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryBlockTradeOrderResponse{}

// QueryBlockTradeOrderResponse struct for QueryBlockTradeOrderResponse
type QueryBlockTradeOrderResponse struct {
	Items []QueryBlockTradeOrderResponseInner
}

// NewQueryBlockTradeOrderResponse instantiates a new QueryBlockTradeOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryBlockTradeOrderResponse() *QueryBlockTradeOrderResponse {
	this := QueryBlockTradeOrderResponse{}
	return &this
}

// NewQueryBlockTradeOrderResponseWithDefaults instantiates a new QueryBlockTradeOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryBlockTradeOrderResponseWithDefaults() *QueryBlockTradeOrderResponse {
	this := QueryBlockTradeOrderResponse{}
	return &this
}

func (o QueryBlockTradeOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryBlockTradeOrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryBlockTradeOrderResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryBlockTradeOrderResponse struct {
	value QueryBlockTradeOrderResponse
	isSet bool
}

func (v NullableQueryBlockTradeOrderResponse) Get() QueryBlockTradeOrderResponse {
	return v.value
}

func (v *NullableQueryBlockTradeOrderResponse) Set(val QueryBlockTradeOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryBlockTradeOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryBlockTradeOrderResponse) Unset() {
	v.value = QueryBlockTradeOrderResponse{}
	v.isSet = false
}

func NewNullableQueryBlockTradeOrderResponse(val QueryBlockTradeOrderResponse) *NullableQueryBlockTradeOrderResponse {
	return &NullableQueryBlockTradeOrderResponse{value: val, isSet: true}
}

func (v NullableQueryBlockTradeOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryBlockTradeOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
