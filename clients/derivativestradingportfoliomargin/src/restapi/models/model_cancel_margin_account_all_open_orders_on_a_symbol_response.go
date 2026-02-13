/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the CancelMarginAccountAllOpenOrdersOnASymbolResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CancelMarginAccountAllOpenOrdersOnASymbolResponse{}

// CancelMarginAccountAllOpenOrdersOnASymbolResponse struct for CancelMarginAccountAllOpenOrdersOnASymbolResponse
type CancelMarginAccountAllOpenOrdersOnASymbolResponse struct {
	Items []CancelMarginAccountAllOpenOrdersOnASymbolResponseInner
}

// NewCancelMarginAccountAllOpenOrdersOnASymbolResponse instantiates a new CancelMarginAccountAllOpenOrdersOnASymbolResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelMarginAccountAllOpenOrdersOnASymbolResponse() *CancelMarginAccountAllOpenOrdersOnASymbolResponse {
	this := CancelMarginAccountAllOpenOrdersOnASymbolResponse{}
	return &this
}

// NewCancelMarginAccountAllOpenOrdersOnASymbolResponseWithDefaults instantiates a new CancelMarginAccountAllOpenOrdersOnASymbolResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelMarginAccountAllOpenOrdersOnASymbolResponseWithDefaults() *CancelMarginAccountAllOpenOrdersOnASymbolResponse {
	this := CancelMarginAccountAllOpenOrdersOnASymbolResponse{}
	return &this
}

func (o CancelMarginAccountAllOpenOrdersOnASymbolResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelMarginAccountAllOpenOrdersOnASymbolResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *CancelMarginAccountAllOpenOrdersOnASymbolResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableCancelMarginAccountAllOpenOrdersOnASymbolResponse struct {
	value CancelMarginAccountAllOpenOrdersOnASymbolResponse
	isSet bool
}

func (v NullableCancelMarginAccountAllOpenOrdersOnASymbolResponse) Get() CancelMarginAccountAllOpenOrdersOnASymbolResponse {
	return v.value
}

func (v *NullableCancelMarginAccountAllOpenOrdersOnASymbolResponse) Set(val CancelMarginAccountAllOpenOrdersOnASymbolResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelMarginAccountAllOpenOrdersOnASymbolResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelMarginAccountAllOpenOrdersOnASymbolResponse) Unset() {
	v.value = CancelMarginAccountAllOpenOrdersOnASymbolResponse{}
	v.isSet = false
}

func NewNullableCancelMarginAccountAllOpenOrdersOnASymbolResponse(val CancelMarginAccountAllOpenOrdersOnASymbolResponse) *NullableCancelMarginAccountAllOpenOrdersOnASymbolResponse {
	return &NullableCancelMarginAccountAllOpenOrdersOnASymbolResponse{value: val, isSet: true}
}

func (v NullableCancelMarginAccountAllOpenOrdersOnASymbolResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelMarginAccountAllOpenOrdersOnASymbolResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
