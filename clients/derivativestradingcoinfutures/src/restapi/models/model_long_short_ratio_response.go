/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the LongShortRatioResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &LongShortRatioResponse{}

// LongShortRatioResponse struct for LongShortRatioResponse
type LongShortRatioResponse struct {
	Items []LongShortRatioResponseInner
}

// NewLongShortRatioResponse instantiates a new LongShortRatioResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLongShortRatioResponse() *LongShortRatioResponse {
	this := LongShortRatioResponse{}
	return &this
}

// NewLongShortRatioResponseWithDefaults instantiates a new LongShortRatioResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLongShortRatioResponseWithDefaults() *LongShortRatioResponse {
	this := LongShortRatioResponse{}
	return &this
}

func (o LongShortRatioResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LongShortRatioResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *LongShortRatioResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableLongShortRatioResponse struct {
	value LongShortRatioResponse
	isSet bool
}

func (v NullableLongShortRatioResponse) Get() LongShortRatioResponse {
	return v.value
}

func (v *NullableLongShortRatioResponse) Set(val LongShortRatioResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLongShortRatioResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLongShortRatioResponse) Unset() {
	v.value = LongShortRatioResponse{}
	v.isSet = false
}

func NewNullableLongShortRatioResponse(val LongShortRatioResponse) *NullableLongShortRatioResponse {
	return &NullableLongShortRatioResponse{value: val, isSet: true}
}

func (v NullableLongShortRatioResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLongShortRatioResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
