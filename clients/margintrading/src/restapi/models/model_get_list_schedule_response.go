/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetListScheduleResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetListScheduleResponse{}

// GetListScheduleResponse struct for GetListScheduleResponse
type GetListScheduleResponse struct {
	Items []GetListScheduleResponseInner
}

// NewGetListScheduleResponse instantiates a new GetListScheduleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetListScheduleResponse() *GetListScheduleResponse {
	this := GetListScheduleResponse{}
	return &this
}

// NewGetListScheduleResponseWithDefaults instantiates a new GetListScheduleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetListScheduleResponseWithDefaults() *GetListScheduleResponse {
	this := GetListScheduleResponse{}
	return &this
}

func (o GetListScheduleResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetListScheduleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *GetListScheduleResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableGetListScheduleResponse struct {
	value GetListScheduleResponse
	isSet bool
}

func (v NullableGetListScheduleResponse) Get() GetListScheduleResponse {
	return v.value
}

func (v *NullableGetListScheduleResponse) Set(val GetListScheduleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetListScheduleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetListScheduleResponse) Unset() {
	v.value = GetListScheduleResponse{}
	v.isSet = false
}

func NewNullableGetListScheduleResponse(val GetListScheduleResponse) *NullableGetListScheduleResponse {
	return &NullableGetListScheduleResponse{value: val, isSet: true}
}

func (v NullableGetListScheduleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetListScheduleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
