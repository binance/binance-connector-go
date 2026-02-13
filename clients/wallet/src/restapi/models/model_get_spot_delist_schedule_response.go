/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetSpotDelistScheduleResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSpotDelistScheduleResponse{}

// GetSpotDelistScheduleResponse struct for GetSpotDelistScheduleResponse
type GetSpotDelistScheduleResponse struct {
	Items []GetSpotDelistScheduleResponseInner
}

// NewGetSpotDelistScheduleResponse instantiates a new GetSpotDelistScheduleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSpotDelistScheduleResponse() *GetSpotDelistScheduleResponse {
	this := GetSpotDelistScheduleResponse{}
	return &this
}

// NewGetSpotDelistScheduleResponseWithDefaults instantiates a new GetSpotDelistScheduleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSpotDelistScheduleResponseWithDefaults() *GetSpotDelistScheduleResponse {
	this := GetSpotDelistScheduleResponse{}
	return &this
}

func (o GetSpotDelistScheduleResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSpotDelistScheduleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *GetSpotDelistScheduleResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableGetSpotDelistScheduleResponse struct {
	value GetSpotDelistScheduleResponse
	isSet bool
}

func (v NullableGetSpotDelistScheduleResponse) Get() GetSpotDelistScheduleResponse {
	return v.value
}

func (v *NullableGetSpotDelistScheduleResponse) Set(val GetSpotDelistScheduleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSpotDelistScheduleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSpotDelistScheduleResponse) Unset() {
	v.value = GetSpotDelistScheduleResponse{}
	v.isSet = false
}

func NewNullableGetSpotDelistScheduleResponse(val GetSpotDelistScheduleResponse) *NullableGetSpotDelistScheduleResponse {
	return &NullableGetSpotDelistScheduleResponse{value: val, isSet: true}
}

func (v NullableGetSpotDelistScheduleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSpotDelistScheduleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
