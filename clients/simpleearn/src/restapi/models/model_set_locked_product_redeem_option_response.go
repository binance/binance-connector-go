/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the SetLockedProductRedeemOptionResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SetLockedProductRedeemOptionResponse{}

// SetLockedProductRedeemOptionResponse struct for SetLockedProductRedeemOptionResponse
type SetLockedProductRedeemOptionResponse struct {
	Success              *bool `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SetLockedProductRedeemOptionResponse SetLockedProductRedeemOptionResponse

// NewSetLockedProductRedeemOptionResponse instantiates a new SetLockedProductRedeemOptionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetLockedProductRedeemOptionResponse() *SetLockedProductRedeemOptionResponse {
	this := SetLockedProductRedeemOptionResponse{}
	return &this
}

// NewSetLockedProductRedeemOptionResponseWithDefaults instantiates a new SetLockedProductRedeemOptionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetLockedProductRedeemOptionResponseWithDefaults() *SetLockedProductRedeemOptionResponse {
	this := SetLockedProductRedeemOptionResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SetLockedProductRedeemOptionResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetLockedProductRedeemOptionResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SetLockedProductRedeemOptionResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *SetLockedProductRedeemOptionResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o SetLockedProductRedeemOptionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetLockedProductRedeemOptionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SetLockedProductRedeemOptionResponse) UnmarshalJSON(data []byte) (err error) {
	varSetLockedProductRedeemOptionResponse := _SetLockedProductRedeemOptionResponse{}

	err = json.Unmarshal(data, &varSetLockedProductRedeemOptionResponse)

	if err != nil {
		return err
	}

	*o = SetLockedProductRedeemOptionResponse(varSetLockedProductRedeemOptionResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSetLockedProductRedeemOptionResponse struct {
	value *SetLockedProductRedeemOptionResponse
	isSet bool
}

func (v NullableSetLockedProductRedeemOptionResponse) Get() *SetLockedProductRedeemOptionResponse {
	return v.value
}

func (v *NullableSetLockedProductRedeemOptionResponse) Set(val *SetLockedProductRedeemOptionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSetLockedProductRedeemOptionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSetLockedProductRedeemOptionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetLockedProductRedeemOptionResponse(val *SetLockedProductRedeemOptionResponse) *NullableSetLockedProductRedeemOptionResponse {
	return &NullableSetLockedProductRedeemOptionResponse{value: val, isSet: true}
}

func (v NullableSetLockedProductRedeemOptionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetLockedProductRedeemOptionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
