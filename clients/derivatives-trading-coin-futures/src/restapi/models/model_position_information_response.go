/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the PositionInformationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PositionInformationResponse{}

// PositionInformationResponse struct for PositionInformationResponse
type PositionInformationResponse struct {
	Items []PositionInformationResponseInner
}

// NewPositionInformationResponse instantiates a new PositionInformationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPositionInformationResponse() *PositionInformationResponse {
	this := PositionInformationResponse{}
	return &this
}

// NewPositionInformationResponseWithDefaults instantiates a new PositionInformationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPositionInformationResponseWithDefaults() *PositionInformationResponse {
	this := PositionInformationResponse{}
	return &this
}

func (o PositionInformationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PositionInformationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *PositionInformationResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullablePositionInformationResponse struct {
	value PositionInformationResponse
	isSet bool
}

func (v NullablePositionInformationResponse) Get() PositionInformationResponse {
	return v.value
}

func (v *NullablePositionInformationResponse) Set(val PositionInformationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePositionInformationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePositionInformationResponse) Unset() {
	v.value = PositionInformationResponse{}
	v.isSet = false
}

func NewNullablePositionInformationResponse(val PositionInformationResponse) *NullablePositionInformationResponse {
	return &NullablePositionInformationResponse{value: val, isSet: true}
}

func (v NullablePositionInformationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePositionInformationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
