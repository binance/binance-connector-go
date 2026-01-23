/*
Binance Derivatives Trading Portfolio Margin Pro REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin Pro REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the SwitchDeltaModeResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SwitchDeltaModeResponse{}

// SwitchDeltaModeResponse struct for SwitchDeltaModeResponse
type SwitchDeltaModeResponse struct {
	Msg                  *string `json:"msg,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SwitchDeltaModeResponse SwitchDeltaModeResponse

// NewSwitchDeltaModeResponse instantiates a new SwitchDeltaModeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSwitchDeltaModeResponse() *SwitchDeltaModeResponse {
	this := SwitchDeltaModeResponse{}
	return &this
}

// NewSwitchDeltaModeResponseWithDefaults instantiates a new SwitchDeltaModeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSwitchDeltaModeResponseWithDefaults() *SwitchDeltaModeResponse {
	this := SwitchDeltaModeResponse{}
	return &this
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *SwitchDeltaModeResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchDeltaModeResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *SwitchDeltaModeResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *SwitchDeltaModeResponse) SetMsg(v string) {
	o.Msg = &v
}

func (o SwitchDeltaModeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SwitchDeltaModeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SwitchDeltaModeResponse) UnmarshalJSON(data []byte) (err error) {
	varSwitchDeltaModeResponse := _SwitchDeltaModeResponse{}

	err = json.Unmarshal(data, &varSwitchDeltaModeResponse)

	if err != nil {
		return err
	}

	*o = SwitchDeltaModeResponse(varSwitchDeltaModeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "msg")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSwitchDeltaModeResponse struct {
	value *SwitchDeltaModeResponse
	isSet bool
}

func (v NullableSwitchDeltaModeResponse) Get() *SwitchDeltaModeResponse {
	return v.value
}

func (v *NullableSwitchDeltaModeResponse) Set(val *SwitchDeltaModeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSwitchDeltaModeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSwitchDeltaModeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSwitchDeltaModeResponse(val *SwitchDeltaModeResponse) *NullableSwitchDeltaModeResponse {
	return &NullableSwitchDeltaModeResponse{value: val, isSet: true}
}

func (v NullableSwitchDeltaModeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSwitchDeltaModeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
