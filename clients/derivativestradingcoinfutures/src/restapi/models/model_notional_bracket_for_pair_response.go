/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the NotionalBracketForPairResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &NotionalBracketForPairResponse{}

// NotionalBracketForPairResponse struct for NotionalBracketForPairResponse
type NotionalBracketForPairResponse struct {
	Items []NotionalBracketForPairResponseInner
}

// NewNotionalBracketForPairResponse instantiates a new NotionalBracketForPairResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotionalBracketForPairResponse() *NotionalBracketForPairResponse {
	this := NotionalBracketForPairResponse{}
	return &this
}

// NewNotionalBracketForPairResponseWithDefaults instantiates a new NotionalBracketForPairResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotionalBracketForPairResponseWithDefaults() *NotionalBracketForPairResponse {
	this := NotionalBracketForPairResponse{}
	return &this
}

func (o NotionalBracketForPairResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotionalBracketForPairResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *NotionalBracketForPairResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableNotionalBracketForPairResponse struct {
	value NotionalBracketForPairResponse
	isSet bool
}

func (v NullableNotionalBracketForPairResponse) Get() NotionalBracketForPairResponse {
	return v.value
}

func (v *NullableNotionalBracketForPairResponse) Set(val NotionalBracketForPairResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNotionalBracketForPairResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNotionalBracketForPairResponse) Unset() {
	v.value = NotionalBracketForPairResponse{}
	v.isSet = false
}

func NewNullableNotionalBracketForPairResponse(val NotionalBracketForPairResponse) *NullableNotionalBracketForPairResponse {
	return &NullableNotionalBracketForPairResponse{value: val, isSet: true}
}

func (v NullableNotionalBracketForPairResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotionalBracketForPairResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
