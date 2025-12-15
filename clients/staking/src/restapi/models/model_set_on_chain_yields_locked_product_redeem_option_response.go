/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the SetOnChainYieldsLockedProductRedeemOptionResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SetOnChainYieldsLockedProductRedeemOptionResponse{}

// SetOnChainYieldsLockedProductRedeemOptionResponse struct for SetOnChainYieldsLockedProductRedeemOptionResponse
type SetOnChainYieldsLockedProductRedeemOptionResponse struct {
	Success              *bool `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SetOnChainYieldsLockedProductRedeemOptionResponse SetOnChainYieldsLockedProductRedeemOptionResponse

// NewSetOnChainYieldsLockedProductRedeemOptionResponse instantiates a new SetOnChainYieldsLockedProductRedeemOptionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetOnChainYieldsLockedProductRedeemOptionResponse() *SetOnChainYieldsLockedProductRedeemOptionResponse {
	this := SetOnChainYieldsLockedProductRedeemOptionResponse{}
	return &this
}

// NewSetOnChainYieldsLockedProductRedeemOptionResponseWithDefaults instantiates a new SetOnChainYieldsLockedProductRedeemOptionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetOnChainYieldsLockedProductRedeemOptionResponseWithDefaults() *SetOnChainYieldsLockedProductRedeemOptionResponse {
	this := SetOnChainYieldsLockedProductRedeemOptionResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SetOnChainYieldsLockedProductRedeemOptionResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetOnChainYieldsLockedProductRedeemOptionResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SetOnChainYieldsLockedProductRedeemOptionResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *SetOnChainYieldsLockedProductRedeemOptionResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o SetOnChainYieldsLockedProductRedeemOptionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetOnChainYieldsLockedProductRedeemOptionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SetOnChainYieldsLockedProductRedeemOptionResponse) UnmarshalJSON(data []byte) (err error) {
	varSetOnChainYieldsLockedProductRedeemOptionResponse := _SetOnChainYieldsLockedProductRedeemOptionResponse{}

	err = json.Unmarshal(data, &varSetOnChainYieldsLockedProductRedeemOptionResponse)

	if err != nil {
		return err
	}

	*o = SetOnChainYieldsLockedProductRedeemOptionResponse(varSetOnChainYieldsLockedProductRedeemOptionResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSetOnChainYieldsLockedProductRedeemOptionResponse struct {
	value *SetOnChainYieldsLockedProductRedeemOptionResponse
	isSet bool
}

func (v NullableSetOnChainYieldsLockedProductRedeemOptionResponse) Get() *SetOnChainYieldsLockedProductRedeemOptionResponse {
	return v.value
}

func (v *NullableSetOnChainYieldsLockedProductRedeemOptionResponse) Set(val *SetOnChainYieldsLockedProductRedeemOptionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSetOnChainYieldsLockedProductRedeemOptionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSetOnChainYieldsLockedProductRedeemOptionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetOnChainYieldsLockedProductRedeemOptionResponse(val *SetOnChainYieldsLockedProductRedeemOptionResponse) *NullableSetOnChainYieldsLockedProductRedeemOptionResponse {
	return &NullableSetOnChainYieldsLockedProductRedeemOptionResponse{value: val, isSet: true}
}

func (v NullableSetOnChainYieldsLockedProductRedeemOptionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetOnChainYieldsLockedProductRedeemOptionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
