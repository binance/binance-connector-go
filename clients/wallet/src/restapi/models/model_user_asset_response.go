/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the UserAssetResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UserAssetResponse{}

// UserAssetResponse struct for UserAssetResponse
type UserAssetResponse struct {
	Items []UserAssetResponseInner
}

// NewUserAssetResponse instantiates a new UserAssetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAssetResponse() *UserAssetResponse {
	this := UserAssetResponse{}
	return &this
}

// NewUserAssetResponseWithDefaults instantiates a new UserAssetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAssetResponseWithDefaults() *UserAssetResponse {
	this := UserAssetResponse{}
	return &this
}

func (o UserAssetResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserAssetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *UserAssetResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableUserAssetResponse struct {
	value UserAssetResponse
	isSet bool
}

func (v NullableUserAssetResponse) Get() UserAssetResponse {
	return v.value
}

func (v *NullableUserAssetResponse) Set(val UserAssetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAssetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAssetResponse) Unset() {
	v.value = UserAssetResponse{}
	v.isSet = false
}

func NewNullableUserAssetResponse(val UserAssetResponse) *NullableUserAssetResponse {
	return &NullableUserAssetResponse{value: val, isSet: true}
}

func (v NullableUserAssetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAssetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
