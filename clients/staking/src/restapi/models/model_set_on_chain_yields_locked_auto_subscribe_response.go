/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the SetOnChainYieldsLockedAutoSubscribeResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SetOnChainYieldsLockedAutoSubscribeResponse{}

// SetOnChainYieldsLockedAutoSubscribeResponse struct for SetOnChainYieldsLockedAutoSubscribeResponse
type SetOnChainYieldsLockedAutoSubscribeResponse struct {
	Success              *bool `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SetOnChainYieldsLockedAutoSubscribeResponse SetOnChainYieldsLockedAutoSubscribeResponse

// NewSetOnChainYieldsLockedAutoSubscribeResponse instantiates a new SetOnChainYieldsLockedAutoSubscribeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetOnChainYieldsLockedAutoSubscribeResponse() *SetOnChainYieldsLockedAutoSubscribeResponse {
	this := SetOnChainYieldsLockedAutoSubscribeResponse{}
	return &this
}

// NewSetOnChainYieldsLockedAutoSubscribeResponseWithDefaults instantiates a new SetOnChainYieldsLockedAutoSubscribeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetOnChainYieldsLockedAutoSubscribeResponseWithDefaults() *SetOnChainYieldsLockedAutoSubscribeResponse {
	this := SetOnChainYieldsLockedAutoSubscribeResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SetOnChainYieldsLockedAutoSubscribeResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetOnChainYieldsLockedAutoSubscribeResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SetOnChainYieldsLockedAutoSubscribeResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *SetOnChainYieldsLockedAutoSubscribeResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o SetOnChainYieldsLockedAutoSubscribeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetOnChainYieldsLockedAutoSubscribeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SetOnChainYieldsLockedAutoSubscribeResponse) UnmarshalJSON(data []byte) (err error) {
	varSetOnChainYieldsLockedAutoSubscribeResponse := _SetOnChainYieldsLockedAutoSubscribeResponse{}

	err = json.Unmarshal(data, &varSetOnChainYieldsLockedAutoSubscribeResponse)

	if err != nil {
		return err
	}

	*o = SetOnChainYieldsLockedAutoSubscribeResponse(varSetOnChainYieldsLockedAutoSubscribeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSetOnChainYieldsLockedAutoSubscribeResponse struct {
	value *SetOnChainYieldsLockedAutoSubscribeResponse
	isSet bool
}

func (v NullableSetOnChainYieldsLockedAutoSubscribeResponse) Get() *SetOnChainYieldsLockedAutoSubscribeResponse {
	return v.value
}

func (v *NullableSetOnChainYieldsLockedAutoSubscribeResponse) Set(val *SetOnChainYieldsLockedAutoSubscribeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSetOnChainYieldsLockedAutoSubscribeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSetOnChainYieldsLockedAutoSubscribeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetOnChainYieldsLockedAutoSubscribeResponse(val *SetOnChainYieldsLockedAutoSubscribeResponse) *NullableSetOnChainYieldsLockedAutoSubscribeResponse {
	return &NullableSetOnChainYieldsLockedAutoSubscribeResponse{value: val, isSet: true}
}

func (v NullableSetOnChainYieldsLockedAutoSubscribeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetOnChainYieldsLockedAutoSubscribeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
