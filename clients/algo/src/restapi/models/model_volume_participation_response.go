/*
Algo Trading REST API

Programmatic access to Binance’s execution algorithms for creating and managing Spot and Futures algo orders.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the VolumeParticipationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &VolumeParticipationResponse{}

// VolumeParticipationResponse struct for VolumeParticipationResponse
type VolumeParticipationResponse struct {
	ClientAlgoId         *string `json:"clientAlgoId,omitempty"`
	Success              *bool   `json:"success,omitempty"`
	Code                 *int64  `json:"code,omitempty"`
	Msg                  *string `json:"msg,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VolumeParticipationResponse VolumeParticipationResponse

// NewVolumeParticipationResponse instantiates a new VolumeParticipationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeParticipationResponse() *VolumeParticipationResponse {
	this := VolumeParticipationResponse{}
	return &this
}

// NewVolumeParticipationResponseWithDefaults instantiates a new VolumeParticipationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeParticipationResponseWithDefaults() *VolumeParticipationResponse {
	this := VolumeParticipationResponse{}
	return &this
}

// GetClientAlgoId returns the ClientAlgoId field value if set, zero value otherwise.
func (o *VolumeParticipationResponse) GetClientAlgoId() string {
	if o == nil || common.IsNil(o.ClientAlgoId) {
		var ret string
		return ret
	}
	return *o.ClientAlgoId
}

// GetClientAlgoIdOk returns a tuple with the ClientAlgoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeParticipationResponse) GetClientAlgoIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientAlgoId) {
		return nil, false
	}
	return o.ClientAlgoId, true
}

// HasClientAlgoId returns a boolean if a field has been set.
func (o *VolumeParticipationResponse) HasClientAlgoId() bool {
	if o != nil && !common.IsNil(o.ClientAlgoId) {
		return true
	}

	return false
}

// SetClientAlgoId gets a reference to the given string and assigns it to the ClientAlgoId field.
func (o *VolumeParticipationResponse) SetClientAlgoId(v string) {
	o.ClientAlgoId = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *VolumeParticipationResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeParticipationResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *VolumeParticipationResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *VolumeParticipationResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *VolumeParticipationResponse) GetCode() int64 {
	if o == nil || common.IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeParticipationResponse) GetCodeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *VolumeParticipationResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *VolumeParticipationResponse) SetCode(v int64) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *VolumeParticipationResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeParticipationResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *VolumeParticipationResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *VolumeParticipationResponse) SetMsg(v string) {
	o.Msg = &v
}

func (o VolumeParticipationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumeParticipationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ClientAlgoId) {
		toSerialize["clientAlgoId"] = o.ClientAlgoId
	}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !common.IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !common.IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VolumeParticipationResponse) UnmarshalJSON(data []byte) (err error) {
	varVolumeParticipationResponse := _VolumeParticipationResponse{}

	err = json.Unmarshal(data, &varVolumeParticipationResponse)

	if err != nil {
		return err
	}

	*o = VolumeParticipationResponse(varVolumeParticipationResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "clientAlgoId")
		delete(additionalProperties, "success")
		delete(additionalProperties, "code")
		delete(additionalProperties, "msg")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVolumeParticipationResponse struct {
	value *VolumeParticipationResponse
	isSet bool
}

func (v NullableVolumeParticipationResponse) Get() *VolumeParticipationResponse {
	return v.value
}

func (v *NullableVolumeParticipationResponse) Set(val *VolumeParticipationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeParticipationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeParticipationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeParticipationResponse(val *VolumeParticipationResponse) *NullableVolumeParticipationResponse {
	return &NullableVolumeParticipationResponse{value: val, isSet: true}
}

func (v NullableVolumeParticipationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeParticipationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
