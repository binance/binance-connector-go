/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the CmPositionAdlQuantileEstimationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CmPositionAdlQuantileEstimationResponse{}

// CmPositionAdlQuantileEstimationResponse struct for CmPositionAdlQuantileEstimationResponse
type CmPositionAdlQuantileEstimationResponse struct {
	Items []CmPositionAdlQuantileEstimationResponseInner
}

// NewCmPositionAdlQuantileEstimationResponse instantiates a new CmPositionAdlQuantileEstimationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCmPositionAdlQuantileEstimationResponse() *CmPositionAdlQuantileEstimationResponse {
	this := CmPositionAdlQuantileEstimationResponse{}
	return &this
}

// NewCmPositionAdlQuantileEstimationResponseWithDefaults instantiates a new CmPositionAdlQuantileEstimationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCmPositionAdlQuantileEstimationResponseWithDefaults() *CmPositionAdlQuantileEstimationResponse {
	this := CmPositionAdlQuantileEstimationResponse{}
	return &this
}

func (o CmPositionAdlQuantileEstimationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CmPositionAdlQuantileEstimationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *CmPositionAdlQuantileEstimationResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableCmPositionAdlQuantileEstimationResponse struct {
	value CmPositionAdlQuantileEstimationResponse
	isSet bool
}

func (v NullableCmPositionAdlQuantileEstimationResponse) Get() CmPositionAdlQuantileEstimationResponse {
	return v.value
}

func (v *NullableCmPositionAdlQuantileEstimationResponse) Set(val CmPositionAdlQuantileEstimationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCmPositionAdlQuantileEstimationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCmPositionAdlQuantileEstimationResponse) Unset() {
	v.value = CmPositionAdlQuantileEstimationResponse{}
	v.isSet = false
}

func NewNullableCmPositionAdlQuantileEstimationResponse(val CmPositionAdlQuantileEstimationResponse) *NullableCmPositionAdlQuantileEstimationResponse {
	return &NullableCmPositionAdlQuantileEstimationResponse{value: val, isSet: true}
}

func (v NullableCmPositionAdlQuantileEstimationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCmPositionAdlQuantileEstimationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
