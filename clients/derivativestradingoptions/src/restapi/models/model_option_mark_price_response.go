/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OptionMarkPriceResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OptionMarkPriceResponse{}

// OptionMarkPriceResponse struct for OptionMarkPriceResponse
type OptionMarkPriceResponse struct {
	Items []OptionMarkPriceResponseInner
}

// NewOptionMarkPriceResponse instantiates a new OptionMarkPriceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionMarkPriceResponse() *OptionMarkPriceResponse {
	this := OptionMarkPriceResponse{}
	return &this
}

// NewOptionMarkPriceResponseWithDefaults instantiates a new OptionMarkPriceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionMarkPriceResponseWithDefaults() *OptionMarkPriceResponse {
	this := OptionMarkPriceResponse{}
	return &this
}

func (o OptionMarkPriceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionMarkPriceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *OptionMarkPriceResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableOptionMarkPriceResponse struct {
	value OptionMarkPriceResponse
	isSet bool
}

func (v NullableOptionMarkPriceResponse) Get() OptionMarkPriceResponse {
	return v.value
}

func (v *NullableOptionMarkPriceResponse) Set(val OptionMarkPriceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionMarkPriceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionMarkPriceResponse) Unset() {
	v.value = OptionMarkPriceResponse{}
	v.isSet = false
}

func NewNullableOptionMarkPriceResponse(val OptionMarkPriceResponse) *NullableOptionMarkPriceResponse {
	return &NullableOptionMarkPriceResponse{value: val, isSet: true}
}

func (v NullableOptionMarkPriceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionMarkPriceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
