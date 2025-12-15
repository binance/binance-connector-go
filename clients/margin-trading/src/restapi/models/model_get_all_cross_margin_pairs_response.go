/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetAllCrossMarginPairsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetAllCrossMarginPairsResponse{}

// GetAllCrossMarginPairsResponse struct for GetAllCrossMarginPairsResponse
type GetAllCrossMarginPairsResponse struct {
	Items []GetAllCrossMarginPairsResponseInner
}

// NewGetAllCrossMarginPairsResponse instantiates a new GetAllCrossMarginPairsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllCrossMarginPairsResponse() *GetAllCrossMarginPairsResponse {
	this := GetAllCrossMarginPairsResponse{}
	return &this
}

// NewGetAllCrossMarginPairsResponseWithDefaults instantiates a new GetAllCrossMarginPairsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllCrossMarginPairsResponseWithDefaults() *GetAllCrossMarginPairsResponse {
	this := GetAllCrossMarginPairsResponse{}
	return &this
}

func (o GetAllCrossMarginPairsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllCrossMarginPairsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *GetAllCrossMarginPairsResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableGetAllCrossMarginPairsResponse struct {
	value GetAllCrossMarginPairsResponse
	isSet bool
}

func (v NullableGetAllCrossMarginPairsResponse) Get() GetAllCrossMarginPairsResponse {
	return v.value
}

func (v *NullableGetAllCrossMarginPairsResponse) Set(val GetAllCrossMarginPairsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllCrossMarginPairsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllCrossMarginPairsResponse) Unset() {
	v.value = GetAllCrossMarginPairsResponse{}
	v.isSet = false
}

func NewNullableGetAllCrossMarginPairsResponse(val GetAllCrossMarginPairsResponse) *NullableGetAllCrossMarginPairsResponse {
	return &NullableGetAllCrossMarginPairsResponse{value: val, isSet: true}
}

func (v NullableGetAllCrossMarginPairsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllCrossMarginPairsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
