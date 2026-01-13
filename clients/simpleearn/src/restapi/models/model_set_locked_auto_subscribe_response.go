/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the SetLockedAutoSubscribeResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SetLockedAutoSubscribeResponse{}

// SetLockedAutoSubscribeResponse struct for SetLockedAutoSubscribeResponse
type SetLockedAutoSubscribeResponse struct {
	Success              *bool `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SetLockedAutoSubscribeResponse SetLockedAutoSubscribeResponse

// NewSetLockedAutoSubscribeResponse instantiates a new SetLockedAutoSubscribeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetLockedAutoSubscribeResponse() *SetLockedAutoSubscribeResponse {
	this := SetLockedAutoSubscribeResponse{}
	return &this
}

// NewSetLockedAutoSubscribeResponseWithDefaults instantiates a new SetLockedAutoSubscribeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetLockedAutoSubscribeResponseWithDefaults() *SetLockedAutoSubscribeResponse {
	this := SetLockedAutoSubscribeResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SetLockedAutoSubscribeResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetLockedAutoSubscribeResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SetLockedAutoSubscribeResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *SetLockedAutoSubscribeResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o SetLockedAutoSubscribeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetLockedAutoSubscribeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SetLockedAutoSubscribeResponse) UnmarshalJSON(data []byte) (err error) {
	varSetLockedAutoSubscribeResponse := _SetLockedAutoSubscribeResponse{}

	err = json.Unmarshal(data, &varSetLockedAutoSubscribeResponse)

	if err != nil {
		return err
	}

	*o = SetLockedAutoSubscribeResponse(varSetLockedAutoSubscribeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSetLockedAutoSubscribeResponse struct {
	value *SetLockedAutoSubscribeResponse
	isSet bool
}

func (v NullableSetLockedAutoSubscribeResponse) Get() *SetLockedAutoSubscribeResponse {
	return v.value
}

func (v *NullableSetLockedAutoSubscribeResponse) Set(val *SetLockedAutoSubscribeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSetLockedAutoSubscribeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSetLockedAutoSubscribeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetLockedAutoSubscribeResponse(val *SetLockedAutoSubscribeResponse) *NullableSetLockedAutoSubscribeResponse {
	return &NullableSetLockedAutoSubscribeResponse{value: val, isSet: true}
}

func (v NullableSetLockedAutoSubscribeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetLockedAutoSubscribeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
