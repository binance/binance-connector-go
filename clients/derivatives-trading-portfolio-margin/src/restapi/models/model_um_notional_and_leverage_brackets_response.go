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

// checks if the UmNotionalAndLeverageBracketsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UmNotionalAndLeverageBracketsResponse{}

// UmNotionalAndLeverageBracketsResponse struct for UmNotionalAndLeverageBracketsResponse
type UmNotionalAndLeverageBracketsResponse struct {
	Items []UmNotionalAndLeverageBracketsResponseInner
}

// NewUmNotionalAndLeverageBracketsResponse instantiates a new UmNotionalAndLeverageBracketsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUmNotionalAndLeverageBracketsResponse() *UmNotionalAndLeverageBracketsResponse {
	this := UmNotionalAndLeverageBracketsResponse{}
	return &this
}

// NewUmNotionalAndLeverageBracketsResponseWithDefaults instantiates a new UmNotionalAndLeverageBracketsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUmNotionalAndLeverageBracketsResponseWithDefaults() *UmNotionalAndLeverageBracketsResponse {
	this := UmNotionalAndLeverageBracketsResponse{}
	return &this
}

func (o UmNotionalAndLeverageBracketsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UmNotionalAndLeverageBracketsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *UmNotionalAndLeverageBracketsResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableUmNotionalAndLeverageBracketsResponse struct {
	value UmNotionalAndLeverageBracketsResponse
	isSet bool
}

func (v NullableUmNotionalAndLeverageBracketsResponse) Get() UmNotionalAndLeverageBracketsResponse {
	return v.value
}

func (v *NullableUmNotionalAndLeverageBracketsResponse) Set(val UmNotionalAndLeverageBracketsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUmNotionalAndLeverageBracketsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUmNotionalAndLeverageBracketsResponse) Unset() {
	v.value = UmNotionalAndLeverageBracketsResponse{}
	v.isSet = false
}

func NewNullableUmNotionalAndLeverageBracketsResponse(val UmNotionalAndLeverageBracketsResponse) *NullableUmNotionalAndLeverageBracketsResponse {
	return &NullableUmNotionalAndLeverageBracketsResponse{value: val, isSet: true}
}

func (v NullableUmNotionalAndLeverageBracketsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUmNotionalAndLeverageBracketsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
