/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetSymbolsDelistScheduleForSpotResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSymbolsDelistScheduleForSpotResponse{}

// GetSymbolsDelistScheduleForSpotResponse struct for GetSymbolsDelistScheduleForSpotResponse
type GetSymbolsDelistScheduleForSpotResponse struct {
	Items []GetSymbolsDelistScheduleForSpotResponseInner
}

// NewGetSymbolsDelistScheduleForSpotResponse instantiates a new GetSymbolsDelistScheduleForSpotResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSymbolsDelistScheduleForSpotResponse() *GetSymbolsDelistScheduleForSpotResponse {
	this := GetSymbolsDelistScheduleForSpotResponse{}
	return &this
}

// NewGetSymbolsDelistScheduleForSpotResponseWithDefaults instantiates a new GetSymbolsDelistScheduleForSpotResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSymbolsDelistScheduleForSpotResponseWithDefaults() *GetSymbolsDelistScheduleForSpotResponse {
	this := GetSymbolsDelistScheduleForSpotResponse{}
	return &this
}

func (o GetSymbolsDelistScheduleForSpotResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSymbolsDelistScheduleForSpotResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *GetSymbolsDelistScheduleForSpotResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableGetSymbolsDelistScheduleForSpotResponse struct {
	value GetSymbolsDelistScheduleForSpotResponse
	isSet bool
}

func (v NullableGetSymbolsDelistScheduleForSpotResponse) Get() GetSymbolsDelistScheduleForSpotResponse {
	return v.value
}

func (v *NullableGetSymbolsDelistScheduleForSpotResponse) Set(val GetSymbolsDelistScheduleForSpotResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSymbolsDelistScheduleForSpotResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSymbolsDelistScheduleForSpotResponse) Unset() {
	v.value = GetSymbolsDelistScheduleForSpotResponse{}
	v.isSet = false
}

func NewNullableGetSymbolsDelistScheduleForSpotResponse(val GetSymbolsDelistScheduleForSpotResponse) *NullableGetSymbolsDelistScheduleForSpotResponse {
	return &NullableGetSymbolsDelistScheduleForSpotResponse{value: val, isSet: true}
}

func (v NullableGetSymbolsDelistScheduleForSpotResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSymbolsDelistScheduleForSpotResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
