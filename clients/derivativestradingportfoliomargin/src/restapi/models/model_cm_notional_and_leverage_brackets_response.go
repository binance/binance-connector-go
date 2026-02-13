/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the CmNotionalAndLeverageBracketsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CmNotionalAndLeverageBracketsResponse{}

// CmNotionalAndLeverageBracketsResponse struct for CmNotionalAndLeverageBracketsResponse
type CmNotionalAndLeverageBracketsResponse struct {
	Items []CmNotionalAndLeverageBracketsResponseInner
}

// NewCmNotionalAndLeverageBracketsResponse instantiates a new CmNotionalAndLeverageBracketsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCmNotionalAndLeverageBracketsResponse() *CmNotionalAndLeverageBracketsResponse {
	this := CmNotionalAndLeverageBracketsResponse{}
	return &this
}

// NewCmNotionalAndLeverageBracketsResponseWithDefaults instantiates a new CmNotionalAndLeverageBracketsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCmNotionalAndLeverageBracketsResponseWithDefaults() *CmNotionalAndLeverageBracketsResponse {
	this := CmNotionalAndLeverageBracketsResponse{}
	return &this
}

func (o CmNotionalAndLeverageBracketsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CmNotionalAndLeverageBracketsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *CmNotionalAndLeverageBracketsResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableCmNotionalAndLeverageBracketsResponse struct {
	value CmNotionalAndLeverageBracketsResponse
	isSet bool
}

func (v NullableCmNotionalAndLeverageBracketsResponse) Get() CmNotionalAndLeverageBracketsResponse {
	return v.value
}

func (v *NullableCmNotionalAndLeverageBracketsResponse) Set(val CmNotionalAndLeverageBracketsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCmNotionalAndLeverageBracketsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCmNotionalAndLeverageBracketsResponse) Unset() {
	v.value = CmNotionalAndLeverageBracketsResponse{}
	v.isSet = false
}

func NewNullableCmNotionalAndLeverageBracketsResponse(val CmNotionalAndLeverageBracketsResponse) *NullableCmNotionalAndLeverageBracketsResponse {
	return &NullableCmNotionalAndLeverageBracketsResponse{value: val, isSet: true}
}

func (v NullableCmNotionalAndLeverageBracketsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCmNotionalAndLeverageBracketsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
