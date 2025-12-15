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

// checks if the KlineCandlestickDataResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &KlineCandlestickDataResponse{}

// KlineCandlestickDataResponse struct for KlineCandlestickDataResponse
type KlineCandlestickDataResponse struct {
	Items []KlineCandlestickDataResponseInner
}

// NewKlineCandlestickDataResponse instantiates a new KlineCandlestickDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKlineCandlestickDataResponse() *KlineCandlestickDataResponse {
	this := KlineCandlestickDataResponse{}
	return &this
}

// NewKlineCandlestickDataResponseWithDefaults instantiates a new KlineCandlestickDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKlineCandlestickDataResponseWithDefaults() *KlineCandlestickDataResponse {
	this := KlineCandlestickDataResponse{}
	return &this
}

func (o KlineCandlestickDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KlineCandlestickDataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *KlineCandlestickDataResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableKlineCandlestickDataResponse struct {
	value KlineCandlestickDataResponse
	isSet bool
}

func (v NullableKlineCandlestickDataResponse) Get() KlineCandlestickDataResponse {
	return v.value
}

func (v *NullableKlineCandlestickDataResponse) Set(val KlineCandlestickDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKlineCandlestickDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKlineCandlestickDataResponse) Unset() {
	v.value = KlineCandlestickDataResponse{}
	v.isSet = false
}

func NewNullableKlineCandlestickDataResponse(val KlineCandlestickDataResponse) *NullableKlineCandlestickDataResponse {
	return &NullableKlineCandlestickDataResponse{value: val, isSet: true}
}

func (v NullableKlineCandlestickDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlineCandlestickDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
