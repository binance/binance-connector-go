/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryCurrentMarginOpenOrderResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryCurrentMarginOpenOrderResponse{}

// QueryCurrentMarginOpenOrderResponse struct for QueryCurrentMarginOpenOrderResponse
type QueryCurrentMarginOpenOrderResponse struct {
	Items []QueryCurrentMarginOpenOrderResponseInner
}

// NewQueryCurrentMarginOpenOrderResponse instantiates a new QueryCurrentMarginOpenOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryCurrentMarginOpenOrderResponse() *QueryCurrentMarginOpenOrderResponse {
	this := QueryCurrentMarginOpenOrderResponse{}
	return &this
}

// NewQueryCurrentMarginOpenOrderResponseWithDefaults instantiates a new QueryCurrentMarginOpenOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryCurrentMarginOpenOrderResponseWithDefaults() *QueryCurrentMarginOpenOrderResponse {
	this := QueryCurrentMarginOpenOrderResponse{}
	return &this
}

func (o QueryCurrentMarginOpenOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryCurrentMarginOpenOrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryCurrentMarginOpenOrderResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryCurrentMarginOpenOrderResponse struct {
	value QueryCurrentMarginOpenOrderResponse
	isSet bool
}

func (v NullableQueryCurrentMarginOpenOrderResponse) Get() QueryCurrentMarginOpenOrderResponse {
	return v.value
}

func (v *NullableQueryCurrentMarginOpenOrderResponse) Set(val QueryCurrentMarginOpenOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryCurrentMarginOpenOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryCurrentMarginOpenOrderResponse) Unset() {
	v.value = QueryCurrentMarginOpenOrderResponse{}
	v.isSet = false
}

func NewNullableQueryCurrentMarginOpenOrderResponse(val QueryCurrentMarginOpenOrderResponse) *NullableQueryCurrentMarginOpenOrderResponse {
	return &NullableQueryCurrentMarginOpenOrderResponse{value: val, isSet: true}
}

func (v NullableQueryCurrentMarginOpenOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryCurrentMarginOpenOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
