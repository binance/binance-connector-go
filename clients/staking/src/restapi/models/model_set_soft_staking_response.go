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

// checks if the SetSoftStakingResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SetSoftStakingResponse{}

// SetSoftStakingResponse struct for SetSoftStakingResponse
type SetSoftStakingResponse struct {
	Success              *bool `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SetSoftStakingResponse SetSoftStakingResponse

// NewSetSoftStakingResponse instantiates a new SetSoftStakingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetSoftStakingResponse() *SetSoftStakingResponse {
	this := SetSoftStakingResponse{}
	return &this
}

// NewSetSoftStakingResponseWithDefaults instantiates a new SetSoftStakingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetSoftStakingResponseWithDefaults() *SetSoftStakingResponse {
	this := SetSoftStakingResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SetSoftStakingResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetSoftStakingResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SetSoftStakingResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *SetSoftStakingResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o SetSoftStakingResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetSoftStakingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SetSoftStakingResponse) UnmarshalJSON(data []byte) (err error) {
	varSetSoftStakingResponse := _SetSoftStakingResponse{}

	err = json.Unmarshal(data, &varSetSoftStakingResponse)

	if err != nil {
		return err
	}

	*o = SetSoftStakingResponse(varSetSoftStakingResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSetSoftStakingResponse struct {
	value *SetSoftStakingResponse
	isSet bool
}

func (v NullableSetSoftStakingResponse) Get() *SetSoftStakingResponse {
	return v.value
}

func (v *NullableSetSoftStakingResponse) Set(val *SetSoftStakingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSetSoftStakingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSetSoftStakingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetSoftStakingResponse(val *SetSoftStakingResponse) *NullableSetSoftStakingResponse {
	return &NullableSetSoftStakingResponse{value: val, isSet: true}
}

func (v NullableSetSoftStakingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetSoftStakingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
