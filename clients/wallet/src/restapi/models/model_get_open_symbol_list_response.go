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

// checks if the GetOpenSymbolListResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetOpenSymbolListResponse{}

// GetOpenSymbolListResponse struct for GetOpenSymbolListResponse
type GetOpenSymbolListResponse struct {
	Items []GetOpenSymbolListResponseInner
}

// NewGetOpenSymbolListResponse instantiates a new GetOpenSymbolListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOpenSymbolListResponse() *GetOpenSymbolListResponse {
	this := GetOpenSymbolListResponse{}
	return &this
}

// NewGetOpenSymbolListResponseWithDefaults instantiates a new GetOpenSymbolListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOpenSymbolListResponseWithDefaults() *GetOpenSymbolListResponse {
	this := GetOpenSymbolListResponse{}
	return &this
}

func (o GetOpenSymbolListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOpenSymbolListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *GetOpenSymbolListResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableGetOpenSymbolListResponse struct {
	value GetOpenSymbolListResponse
	isSet bool
}

func (v NullableGetOpenSymbolListResponse) Get() GetOpenSymbolListResponse {
	return v.value
}

func (v *NullableGetOpenSymbolListResponse) Set(val GetOpenSymbolListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOpenSymbolListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOpenSymbolListResponse) Unset() {
	v.value = GetOpenSymbolListResponse{}
	v.isSet = false
}

func NewNullableGetOpenSymbolListResponse(val GetOpenSymbolListResponse) *NullableGetOpenSymbolListResponse {
	return &NullableGetOpenSymbolListResponse{value: val, isSet: true}
}

func (v NullableGetOpenSymbolListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOpenSymbolListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
