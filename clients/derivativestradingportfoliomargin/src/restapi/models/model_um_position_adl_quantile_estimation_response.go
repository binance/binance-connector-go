/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the UmPositionAdlQuantileEstimationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UmPositionAdlQuantileEstimationResponse{}

// UmPositionAdlQuantileEstimationResponse struct for UmPositionAdlQuantileEstimationResponse
type UmPositionAdlQuantileEstimationResponse struct {
	Items []UmPositionAdlQuantileEstimationResponseInner
}

// NewUmPositionAdlQuantileEstimationResponse instantiates a new UmPositionAdlQuantileEstimationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUmPositionAdlQuantileEstimationResponse() *UmPositionAdlQuantileEstimationResponse {
	this := UmPositionAdlQuantileEstimationResponse{}
	return &this
}

// NewUmPositionAdlQuantileEstimationResponseWithDefaults instantiates a new UmPositionAdlQuantileEstimationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUmPositionAdlQuantileEstimationResponseWithDefaults() *UmPositionAdlQuantileEstimationResponse {
	this := UmPositionAdlQuantileEstimationResponse{}
	return &this
}

func (o UmPositionAdlQuantileEstimationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UmPositionAdlQuantileEstimationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *UmPositionAdlQuantileEstimationResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableUmPositionAdlQuantileEstimationResponse struct {
	value UmPositionAdlQuantileEstimationResponse
	isSet bool
}

func (v NullableUmPositionAdlQuantileEstimationResponse) Get() UmPositionAdlQuantileEstimationResponse {
	return v.value
}

func (v *NullableUmPositionAdlQuantileEstimationResponse) Set(val UmPositionAdlQuantileEstimationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUmPositionAdlQuantileEstimationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUmPositionAdlQuantileEstimationResponse) Unset() {
	v.value = UmPositionAdlQuantileEstimationResponse{}
	v.isSet = false
}

func NewNullableUmPositionAdlQuantileEstimationResponse(val UmPositionAdlQuantileEstimationResponse) *NullableUmPositionAdlQuantileEstimationResponse {
	return &NullableUmPositionAdlQuantileEstimationResponse{value: val, isSet: true}
}

func (v NullableUmPositionAdlQuantileEstimationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUmPositionAdlQuantileEstimationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
