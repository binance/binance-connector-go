/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetAllMarginAssetsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetAllMarginAssetsResponse{}

// GetAllMarginAssetsResponse struct for GetAllMarginAssetsResponse
type GetAllMarginAssetsResponse struct {
	Items []GetAllMarginAssetsResponseInner
}

// NewGetAllMarginAssetsResponse instantiates a new GetAllMarginAssetsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllMarginAssetsResponse() *GetAllMarginAssetsResponse {
	this := GetAllMarginAssetsResponse{}
	return &this
}

// NewGetAllMarginAssetsResponseWithDefaults instantiates a new GetAllMarginAssetsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllMarginAssetsResponseWithDefaults() *GetAllMarginAssetsResponse {
	this := GetAllMarginAssetsResponse{}
	return &this
}

func (o GetAllMarginAssetsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllMarginAssetsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *GetAllMarginAssetsResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableGetAllMarginAssetsResponse struct {
	value GetAllMarginAssetsResponse
	isSet bool
}

func (v NullableGetAllMarginAssetsResponse) Get() GetAllMarginAssetsResponse {
	return v.value
}

func (v *NullableGetAllMarginAssetsResponse) Set(val GetAllMarginAssetsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllMarginAssetsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllMarginAssetsResponse) Unset() {
	v.value = GetAllMarginAssetsResponse{}
	v.isSet = false
}

func NewNullableGetAllMarginAssetsResponse(val GetAllMarginAssetsResponse) *NullableGetAllMarginAssetsResponse {
	return &NullableGetAllMarginAssetsResponse{value: val, isSet: true}
}

func (v NullableGetAllMarginAssetsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllMarginAssetsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
