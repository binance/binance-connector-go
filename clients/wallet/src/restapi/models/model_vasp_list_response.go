/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the VaspListResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &VaspListResponse{}

// VaspListResponse struct for VaspListResponse
type VaspListResponse struct {
	Items []VaspListResponseInner
}

// NewVaspListResponse instantiates a new VaspListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVaspListResponse() *VaspListResponse {
	this := VaspListResponse{}
	return &this
}

// NewVaspListResponseWithDefaults instantiates a new VaspListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVaspListResponseWithDefaults() *VaspListResponse {
	this := VaspListResponse{}
	return &this
}

func (o VaspListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VaspListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *VaspListResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableVaspListResponse struct {
	value VaspListResponse
	isSet bool
}

func (v NullableVaspListResponse) Get() VaspListResponse {
	return v.value
}

func (v *NullableVaspListResponse) Set(val VaspListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVaspListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVaspListResponse) Unset() {
	v.value = VaspListResponse{}
	v.isSet = false
}

func NewNullableVaspListResponse(val VaspListResponse) *NullableVaspListResponse {
	return &NullableVaspListResponse{value: val, isSet: true}
}

func (v NullableVaspListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVaspListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
