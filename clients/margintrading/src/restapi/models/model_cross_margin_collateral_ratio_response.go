/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CrossMarginCollateralRatioResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CrossMarginCollateralRatioResponse{}

// CrossMarginCollateralRatioResponse struct for CrossMarginCollateralRatioResponse
type CrossMarginCollateralRatioResponse struct {
	Items []CrossMarginCollateralRatioResponseInner
}

// NewCrossMarginCollateralRatioResponse instantiates a new CrossMarginCollateralRatioResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCrossMarginCollateralRatioResponse() *CrossMarginCollateralRatioResponse {
	this := CrossMarginCollateralRatioResponse{}
	return &this
}

// NewCrossMarginCollateralRatioResponseWithDefaults instantiates a new CrossMarginCollateralRatioResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCrossMarginCollateralRatioResponseWithDefaults() *CrossMarginCollateralRatioResponse {
	this := CrossMarginCollateralRatioResponse{}
	return &this
}

func (o CrossMarginCollateralRatioResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CrossMarginCollateralRatioResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *CrossMarginCollateralRatioResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableCrossMarginCollateralRatioResponse struct {
	value CrossMarginCollateralRatioResponse
	isSet bool
}

func (v NullableCrossMarginCollateralRatioResponse) Get() CrossMarginCollateralRatioResponse {
	return v.value
}

func (v *NullableCrossMarginCollateralRatioResponse) Set(val CrossMarginCollateralRatioResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCrossMarginCollateralRatioResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCrossMarginCollateralRatioResponse) Unset() {
	v.value = CrossMarginCollateralRatioResponse{}
	v.isSet = false
}

func NewNullableCrossMarginCollateralRatioResponse(val CrossMarginCollateralRatioResponse) *NullableCrossMarginCollateralRatioResponse {
	return &NullableCrossMarginCollateralRatioResponse{value: val, isSet: true}
}

func (v NullableCrossMarginCollateralRatioResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCrossMarginCollateralRatioResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
