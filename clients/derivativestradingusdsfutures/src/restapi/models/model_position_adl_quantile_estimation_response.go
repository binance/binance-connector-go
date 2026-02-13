/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the PositionAdlQuantileEstimationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PositionAdlQuantileEstimationResponse{}

// PositionAdlQuantileEstimationResponse struct for PositionAdlQuantileEstimationResponse
type PositionAdlQuantileEstimationResponse struct {
	Items []PositionAdlQuantileEstimationResponseInner
}

// NewPositionAdlQuantileEstimationResponse instantiates a new PositionAdlQuantileEstimationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPositionAdlQuantileEstimationResponse() *PositionAdlQuantileEstimationResponse {
	this := PositionAdlQuantileEstimationResponse{}
	return &this
}

// NewPositionAdlQuantileEstimationResponseWithDefaults instantiates a new PositionAdlQuantileEstimationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPositionAdlQuantileEstimationResponseWithDefaults() *PositionAdlQuantileEstimationResponse {
	this := PositionAdlQuantileEstimationResponse{}
	return &this
}

func (o PositionAdlQuantileEstimationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PositionAdlQuantileEstimationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *PositionAdlQuantileEstimationResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullablePositionAdlQuantileEstimationResponse struct {
	value PositionAdlQuantileEstimationResponse
	isSet bool
}

func (v NullablePositionAdlQuantileEstimationResponse) Get() PositionAdlQuantileEstimationResponse {
	return v.value
}

func (v *NullablePositionAdlQuantileEstimationResponse) Set(val PositionAdlQuantileEstimationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePositionAdlQuantileEstimationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePositionAdlQuantileEstimationResponse) Unset() {
	v.value = PositionAdlQuantileEstimationResponse{}
	v.isSet = false
}

func NewNullablePositionAdlQuantileEstimationResponse(val PositionAdlQuantileEstimationResponse) *NullablePositionAdlQuantileEstimationResponse {
	return &NullablePositionAdlQuantileEstimationResponse{value: val, isSet: true}
}

func (v NullablePositionAdlQuantileEstimationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePositionAdlQuantileEstimationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
