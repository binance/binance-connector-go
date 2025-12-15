/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the ModifyMultipleOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ModifyMultipleOrdersResponse{}

// ModifyMultipleOrdersResponse struct for ModifyMultipleOrdersResponse
type ModifyMultipleOrdersResponse struct {
	Items []ModifyMultipleOrdersResponseInner
}

// NewModifyMultipleOrdersResponse instantiates a new ModifyMultipleOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyMultipleOrdersResponse() *ModifyMultipleOrdersResponse {
	this := ModifyMultipleOrdersResponse{}
	return &this
}

// NewModifyMultipleOrdersResponseWithDefaults instantiates a new ModifyMultipleOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyMultipleOrdersResponseWithDefaults() *ModifyMultipleOrdersResponse {
	this := ModifyMultipleOrdersResponse{}
	return &this
}

func (o ModifyMultipleOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyMultipleOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *ModifyMultipleOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableModifyMultipleOrdersResponse struct {
	value ModifyMultipleOrdersResponse
	isSet bool
}

func (v NullableModifyMultipleOrdersResponse) Get() ModifyMultipleOrdersResponse {
	return v.value
}

func (v *NullableModifyMultipleOrdersResponse) Set(val ModifyMultipleOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyMultipleOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyMultipleOrdersResponse) Unset() {
	v.value = ModifyMultipleOrdersResponse{}
	v.isSet = false
}

func NewNullableModifyMultipleOrdersResponse(val ModifyMultipleOrdersResponse) *NullableModifyMultipleOrdersResponse {
	return &NullableModifyMultipleOrdersResponse{value: val, isSet: true}
}

func (v NullableModifyMultipleOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyMultipleOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
