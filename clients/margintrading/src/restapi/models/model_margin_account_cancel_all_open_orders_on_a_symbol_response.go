/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the MarginAccountCancelAllOpenOrdersOnASymbolResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarginAccountCancelAllOpenOrdersOnASymbolResponse{}

// MarginAccountCancelAllOpenOrdersOnASymbolResponse struct for MarginAccountCancelAllOpenOrdersOnASymbolResponse
type MarginAccountCancelAllOpenOrdersOnASymbolResponse struct {
	Items []MarginAccountCancelAllOpenOrdersOnASymbolResponseInner
}

// NewMarginAccountCancelAllOpenOrdersOnASymbolResponse instantiates a new MarginAccountCancelAllOpenOrdersOnASymbolResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginAccountCancelAllOpenOrdersOnASymbolResponse() *MarginAccountCancelAllOpenOrdersOnASymbolResponse {
	this := MarginAccountCancelAllOpenOrdersOnASymbolResponse{}
	return &this
}

// NewMarginAccountCancelAllOpenOrdersOnASymbolResponseWithDefaults instantiates a new MarginAccountCancelAllOpenOrdersOnASymbolResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginAccountCancelAllOpenOrdersOnASymbolResponseWithDefaults() *MarginAccountCancelAllOpenOrdersOnASymbolResponse {
	this := MarginAccountCancelAllOpenOrdersOnASymbolResponse{}
	return &this
}

func (o MarginAccountCancelAllOpenOrdersOnASymbolResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginAccountCancelAllOpenOrdersOnASymbolResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *MarginAccountCancelAllOpenOrdersOnASymbolResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableMarginAccountCancelAllOpenOrdersOnASymbolResponse struct {
	value MarginAccountCancelAllOpenOrdersOnASymbolResponse
	isSet bool
}

func (v NullableMarginAccountCancelAllOpenOrdersOnASymbolResponse) Get() MarginAccountCancelAllOpenOrdersOnASymbolResponse {
	return v.value
}

func (v *NullableMarginAccountCancelAllOpenOrdersOnASymbolResponse) Set(val MarginAccountCancelAllOpenOrdersOnASymbolResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginAccountCancelAllOpenOrdersOnASymbolResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginAccountCancelAllOpenOrdersOnASymbolResponse) Unset() {
	v.value = MarginAccountCancelAllOpenOrdersOnASymbolResponse{}
	v.isSet = false
}

func NewNullableMarginAccountCancelAllOpenOrdersOnASymbolResponse(val MarginAccountCancelAllOpenOrdersOnASymbolResponse) *NullableMarginAccountCancelAllOpenOrdersOnASymbolResponse {
	return &NullableMarginAccountCancelAllOpenOrdersOnASymbolResponse{value: val, isSet: true}
}

func (v NullableMarginAccountCancelAllOpenOrdersOnASymbolResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginAccountCancelAllOpenOrdersOnASymbolResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
