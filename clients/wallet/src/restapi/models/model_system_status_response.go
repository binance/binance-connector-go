/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the SystemStatusResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SystemStatusResponse{}

// SystemStatusResponse struct for SystemStatusResponse
type SystemStatusResponse struct {
	Status               *int64  `json:"status,omitempty"`
	Msg                  *string `json:"msg,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SystemStatusResponse SystemStatusResponse

// NewSystemStatusResponse instantiates a new SystemStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemStatusResponse() *SystemStatusResponse {
	this := SystemStatusResponse{}
	return &this
}

// NewSystemStatusResponseWithDefaults instantiates a new SystemStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemStatusResponseWithDefaults() *SystemStatusResponse {
	this := SystemStatusResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SystemStatusResponse) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemStatusResponse) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SystemStatusResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *SystemStatusResponse) SetStatus(v int64) {
	o.Status = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *SystemStatusResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemStatusResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *SystemStatusResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *SystemStatusResponse) SetMsg(v string) {
	o.Msg = &v
}

func (o SystemStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SystemStatusResponse) UnmarshalJSON(data []byte) (err error) {
	varSystemStatusResponse := _SystemStatusResponse{}

	err = json.Unmarshal(data, &varSystemStatusResponse)

	if err != nil {
		return err
	}

	*o = SystemStatusResponse(varSystemStatusResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "msg")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSystemStatusResponse struct {
	value *SystemStatusResponse
	isSet bool
}

func (v NullableSystemStatusResponse) Get() *SystemStatusResponse {
	return v.value
}

func (v *NullableSystemStatusResponse) Set(val *SystemStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemStatusResponse(val *SystemStatusResponse) *NullableSystemStatusResponse {
	return &NullableSystemStatusResponse{value: val, isSet: true}
}

func (v NullableSystemStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
