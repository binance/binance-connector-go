/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryCurrentCmOpenOrderResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryCurrentCmOpenOrderResponse{}

// QueryCurrentCmOpenOrderResponse struct for QueryCurrentCmOpenOrderResponse
type QueryCurrentCmOpenOrderResponse struct {
	Items []QueryAllCmOrdersResponseInner
}

// NewQueryCurrentCmOpenOrderResponse instantiates a new QueryCurrentCmOpenOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryCurrentCmOpenOrderResponse() *QueryCurrentCmOpenOrderResponse {
	this := QueryCurrentCmOpenOrderResponse{}
	return &this
}

// NewQueryCurrentCmOpenOrderResponseWithDefaults instantiates a new QueryCurrentCmOpenOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryCurrentCmOpenOrderResponseWithDefaults() *QueryCurrentCmOpenOrderResponse {
	this := QueryCurrentCmOpenOrderResponse{}
	return &this
}

func (o QueryCurrentCmOpenOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryCurrentCmOpenOrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryCurrentCmOpenOrderResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryCurrentCmOpenOrderResponse struct {
	value QueryCurrentCmOpenOrderResponse
	isSet bool
}

func (v NullableQueryCurrentCmOpenOrderResponse) Get() QueryCurrentCmOpenOrderResponse {
	return v.value
}

func (v *NullableQueryCurrentCmOpenOrderResponse) Set(val QueryCurrentCmOpenOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryCurrentCmOpenOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryCurrentCmOpenOrderResponse) Unset() {
	v.value = QueryCurrentCmOpenOrderResponse{}
	v.isSet = false
}

func NewNullableQueryCurrentCmOpenOrderResponse(val QueryCurrentCmOpenOrderResponse) *NullableQueryCurrentCmOpenOrderResponse {
	return &NullableQueryCurrentCmOpenOrderResponse{value: val, isSet: true}
}

func (v NullableQueryCurrentCmOpenOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryCurrentCmOpenOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
