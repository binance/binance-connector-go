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

// checks if the OptionPositionInformationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OptionPositionInformationResponse{}

// OptionPositionInformationResponse struct for OptionPositionInformationResponse
type OptionPositionInformationResponse struct {
	Items []OptionPositionInformationResponseInner
}

// NewOptionPositionInformationResponse instantiates a new OptionPositionInformationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionPositionInformationResponse() *OptionPositionInformationResponse {
	this := OptionPositionInformationResponse{}
	return &this
}

// NewOptionPositionInformationResponseWithDefaults instantiates a new OptionPositionInformationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionPositionInformationResponseWithDefaults() *OptionPositionInformationResponse {
	this := OptionPositionInformationResponse{}
	return &this
}

func (o OptionPositionInformationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionPositionInformationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *OptionPositionInformationResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableOptionPositionInformationResponse struct {
	value OptionPositionInformationResponse
	isSet bool
}

func (v NullableOptionPositionInformationResponse) Get() OptionPositionInformationResponse {
	return v.value
}

func (v *NullableOptionPositionInformationResponse) Set(val OptionPositionInformationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionPositionInformationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionPositionInformationResponse) Unset() {
	v.value = OptionPositionInformationResponse{}
	v.isSet = false
}

func NewNullableOptionPositionInformationResponse(val OptionPositionInformationResponse) *NullableOptionPositionInformationResponse {
	return &NullableOptionPositionInformationResponse{value: val, isSet: true}
}

func (v NullableOptionPositionInformationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionPositionInformationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
