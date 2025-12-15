/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the SetFlexibleAutoSubscribeResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SetFlexibleAutoSubscribeResponse{}

// SetFlexibleAutoSubscribeResponse struct for SetFlexibleAutoSubscribeResponse
type SetFlexibleAutoSubscribeResponse struct {
	Success              *bool `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SetFlexibleAutoSubscribeResponse SetFlexibleAutoSubscribeResponse

// NewSetFlexibleAutoSubscribeResponse instantiates a new SetFlexibleAutoSubscribeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetFlexibleAutoSubscribeResponse() *SetFlexibleAutoSubscribeResponse {
	this := SetFlexibleAutoSubscribeResponse{}
	return &this
}

// NewSetFlexibleAutoSubscribeResponseWithDefaults instantiates a new SetFlexibleAutoSubscribeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetFlexibleAutoSubscribeResponseWithDefaults() *SetFlexibleAutoSubscribeResponse {
	this := SetFlexibleAutoSubscribeResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SetFlexibleAutoSubscribeResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetFlexibleAutoSubscribeResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SetFlexibleAutoSubscribeResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *SetFlexibleAutoSubscribeResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o SetFlexibleAutoSubscribeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetFlexibleAutoSubscribeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SetFlexibleAutoSubscribeResponse) UnmarshalJSON(data []byte) (err error) {
	varSetFlexibleAutoSubscribeResponse := _SetFlexibleAutoSubscribeResponse{}

	err = json.Unmarshal(data, &varSetFlexibleAutoSubscribeResponse)

	if err != nil {
		return err
	}

	*o = SetFlexibleAutoSubscribeResponse(varSetFlexibleAutoSubscribeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSetFlexibleAutoSubscribeResponse struct {
	value *SetFlexibleAutoSubscribeResponse
	isSet bool
}

func (v NullableSetFlexibleAutoSubscribeResponse) Get() *SetFlexibleAutoSubscribeResponse {
	return v.value
}

func (v *NullableSetFlexibleAutoSubscribeResponse) Set(val *SetFlexibleAutoSubscribeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSetFlexibleAutoSubscribeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSetFlexibleAutoSubscribeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetFlexibleAutoSubscribeResponse(val *SetFlexibleAutoSubscribeResponse) *NullableSetFlexibleAutoSubscribeResponse {
	return &NullableSetFlexibleAutoSubscribeResponse{value: val, isSet: true}
}

func (v NullableSetFlexibleAutoSubscribeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetFlexibleAutoSubscribeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
