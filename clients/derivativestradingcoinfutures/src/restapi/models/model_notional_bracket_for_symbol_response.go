/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the NotionalBracketForSymbolResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &NotionalBracketForSymbolResponse{}

// NotionalBracketForSymbolResponse struct for NotionalBracketForSymbolResponse
type NotionalBracketForSymbolResponse struct {
	Items []NotionalBracketForSymbolResponseInner
}

// NewNotionalBracketForSymbolResponse instantiates a new NotionalBracketForSymbolResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotionalBracketForSymbolResponse() *NotionalBracketForSymbolResponse {
	this := NotionalBracketForSymbolResponse{}
	return &this
}

// NewNotionalBracketForSymbolResponseWithDefaults instantiates a new NotionalBracketForSymbolResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotionalBracketForSymbolResponseWithDefaults() *NotionalBracketForSymbolResponse {
	this := NotionalBracketForSymbolResponse{}
	return &this
}

func (o NotionalBracketForSymbolResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotionalBracketForSymbolResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *NotionalBracketForSymbolResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableNotionalBracketForSymbolResponse struct {
	value NotionalBracketForSymbolResponse
	isSet bool
}

func (v NullableNotionalBracketForSymbolResponse) Get() NotionalBracketForSymbolResponse {
	return v.value
}

func (v *NullableNotionalBracketForSymbolResponse) Set(val NotionalBracketForSymbolResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNotionalBracketForSymbolResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNotionalBracketForSymbolResponse) Unset() {
	v.value = NotionalBracketForSymbolResponse{}
	v.isSet = false
}

func NewNullableNotionalBracketForSymbolResponse(val NotionalBracketForSymbolResponse) *NullableNotionalBracketForSymbolResponse {
	return &NullableNotionalBracketForSymbolResponse{value: val, isSet: true}
}

func (v NullableNotionalBracketForSymbolResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotionalBracketForSymbolResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
