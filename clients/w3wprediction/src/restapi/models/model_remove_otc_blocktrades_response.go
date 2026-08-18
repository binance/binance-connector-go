/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the RemoveOtcBlocktradesResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RemoveOtcBlocktradesResponse{}

// RemoveOtcBlocktradesResponse struct for RemoveOtcBlocktradesResponse
type RemoveOtcBlocktradesResponse struct {
	Removed              []string `json:"removed,omitempty"`
	Noop                 []string `json:"noop,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RemoveOtcBlocktradesResponse RemoveOtcBlocktradesResponse

// NewRemoveOtcBlocktradesResponse instantiates a new RemoveOtcBlocktradesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveOtcBlocktradesResponse() *RemoveOtcBlocktradesResponse {
	this := RemoveOtcBlocktradesResponse{}
	return &this
}

// NewRemoveOtcBlocktradesResponseWithDefaults instantiates a new RemoveOtcBlocktradesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveOtcBlocktradesResponseWithDefaults() *RemoveOtcBlocktradesResponse {
	this := RemoveOtcBlocktradesResponse{}
	return &this
}

// GetRemoved returns the Removed field value if set, zero value otherwise.
func (o *RemoveOtcBlocktradesResponse) GetRemoved() []string {
	if o == nil || common.IsNil(o.Removed) {
		var ret []string
		return ret
	}
	return o.Removed
}

// GetRemovedOk returns a tuple with the Removed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoveOtcBlocktradesResponse) GetRemovedOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Removed) {
		return nil, false
	}
	return o.Removed, true
}

// HasRemoved returns a boolean if a field has been set.
func (o *RemoveOtcBlocktradesResponse) HasRemoved() bool {
	if o != nil && !common.IsNil(o.Removed) {
		return true
	}

	return false
}

// SetRemoved gets a reference to the given []string and assigns it to the Removed field.
func (o *RemoveOtcBlocktradesResponse) SetRemoved(v []string) {
	o.Removed = v
}

// GetNoop returns the Noop field value if set, zero value otherwise.
func (o *RemoveOtcBlocktradesResponse) GetNoop() []string {
	if o == nil || common.IsNil(o.Noop) {
		var ret []string
		return ret
	}
	return o.Noop
}

// GetNoopOk returns a tuple with the Noop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoveOtcBlocktradesResponse) GetNoopOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Noop) {
		return nil, false
	}
	return o.Noop, true
}

// HasNoop returns a boolean if a field has been set.
func (o *RemoveOtcBlocktradesResponse) HasNoop() bool {
	if o != nil && !common.IsNil(o.Noop) {
		return true
	}

	return false
}

// SetNoop gets a reference to the given []string and assigns it to the Noop field.
func (o *RemoveOtcBlocktradesResponse) SetNoop(v []string) {
	o.Noop = v
}

func (o RemoveOtcBlocktradesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveOtcBlocktradesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Removed) {
		toSerialize["removed"] = o.Removed
	}
	if !common.IsNil(o.Noop) {
		toSerialize["noop"] = o.Noop
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RemoveOtcBlocktradesResponse) UnmarshalJSON(data []byte) (err error) {
	varRemoveOtcBlocktradesResponse := _RemoveOtcBlocktradesResponse{}

	err = json.Unmarshal(data, &varRemoveOtcBlocktradesResponse)

	if err != nil {
		return err
	}

	*o = RemoveOtcBlocktradesResponse(varRemoveOtcBlocktradesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "removed")
		delete(additionalProperties, "noop")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRemoveOtcBlocktradesResponse struct {
	value *RemoveOtcBlocktradesResponse
	isSet bool
}

func (v NullableRemoveOtcBlocktradesResponse) Get() *RemoveOtcBlocktradesResponse {
	return v.value
}

func (v *NullableRemoveOtcBlocktradesResponse) Set(val *RemoveOtcBlocktradesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveOtcBlocktradesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveOtcBlocktradesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveOtcBlocktradesResponse(val *RemoveOtcBlocktradesResponse) *NullableRemoveOtcBlocktradesResponse {
	return &NullableRemoveOtcBlocktradesResponse{value: val, isSet: true}
}

func (v NullableRemoveOtcBlocktradesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveOtcBlocktradesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
