/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the PositionInformationV2Response type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PositionInformationV2Response{}

// PositionInformationV2Response struct for PositionInformationV2Response
type PositionInformationV2Response struct {
	Items []PositionInformationV2ResponseInner
}

// NewPositionInformationV2Response instantiates a new PositionInformationV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPositionInformationV2Response() *PositionInformationV2Response {
	this := PositionInformationV2Response{}
	return &this
}

// NewPositionInformationV2ResponseWithDefaults instantiates a new PositionInformationV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPositionInformationV2ResponseWithDefaults() *PositionInformationV2Response {
	this := PositionInformationV2Response{}
	return &this
}

func (o PositionInformationV2Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PositionInformationV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *PositionInformationV2Response) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullablePositionInformationV2Response struct {
	value PositionInformationV2Response
	isSet bool
}

func (v NullablePositionInformationV2Response) Get() PositionInformationV2Response {
	return v.value
}

func (v *NullablePositionInformationV2Response) Set(val PositionInformationV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePositionInformationV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePositionInformationV2Response) Unset() {
	v.value = PositionInformationV2Response{}
	v.isSet = false
}

func NewNullablePositionInformationV2Response(val PositionInformationV2Response) *NullablePositionInformationV2Response {
	return &NullablePositionInformationV2Response{value: val, isSet: true}
}

func (v NullablePositionInformationV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePositionInformationV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
