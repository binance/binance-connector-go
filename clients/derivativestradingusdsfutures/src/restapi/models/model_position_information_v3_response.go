/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the PositionInformationV3Response type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PositionInformationV3Response{}

// PositionInformationV3Response struct for PositionInformationV3Response
type PositionInformationV3Response struct {
	Items []PositionInformationV3ResponseInner
}

// NewPositionInformationV3Response instantiates a new PositionInformationV3Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPositionInformationV3Response() *PositionInformationV3Response {
	this := PositionInformationV3Response{}
	return &this
}

// NewPositionInformationV3ResponseWithDefaults instantiates a new PositionInformationV3Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPositionInformationV3ResponseWithDefaults() *PositionInformationV3Response {
	this := PositionInformationV3Response{}
	return &this
}

func (o PositionInformationV3Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PositionInformationV3Response) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *PositionInformationV3Response) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullablePositionInformationV3Response struct {
	value PositionInformationV3Response
	isSet bool
}

func (v NullablePositionInformationV3Response) Get() PositionInformationV3Response {
	return v.value
}

func (v *NullablePositionInformationV3Response) Set(val PositionInformationV3Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePositionInformationV3Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePositionInformationV3Response) Unset() {
	v.value = PositionInformationV3Response{}
	v.isSet = false
}

func NewNullablePositionInformationV3Response(val PositionInformationV3Response) *NullablePositionInformationV3Response {
	return &NullablePositionInformationV3Response{value: val, isSet: true}
}

func (v NullablePositionInformationV3Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePositionInformationV3Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
