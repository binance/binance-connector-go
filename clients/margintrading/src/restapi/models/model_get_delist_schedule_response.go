/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetDelistScheduleResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetDelistScheduleResponse{}

// GetDelistScheduleResponse struct for GetDelistScheduleResponse
type GetDelistScheduleResponse struct {
	Items []GetDelistScheduleResponseInner
}

// NewGetDelistScheduleResponse instantiates a new GetDelistScheduleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDelistScheduleResponse() *GetDelistScheduleResponse {
	this := GetDelistScheduleResponse{}
	return &this
}

// NewGetDelistScheduleResponseWithDefaults instantiates a new GetDelistScheduleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDelistScheduleResponseWithDefaults() *GetDelistScheduleResponse {
	this := GetDelistScheduleResponse{}
	return &this
}

func (o GetDelistScheduleResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDelistScheduleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *GetDelistScheduleResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableGetDelistScheduleResponse struct {
	value GetDelistScheduleResponse
	isSet bool
}

func (v NullableGetDelistScheduleResponse) Get() GetDelistScheduleResponse {
	return v.value
}

func (v *NullableGetDelistScheduleResponse) Set(val GetDelistScheduleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDelistScheduleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDelistScheduleResponse) Unset() {
	v.value = GetDelistScheduleResponse{}
	v.isSet = false
}

func NewNullableGetDelistScheduleResponse(val GetDelistScheduleResponse) *NullableGetDelistScheduleResponse {
	return &NullableGetDelistScheduleResponse{value: val, isSet: true}
}

func (v NullableGetDelistScheduleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDelistScheduleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
