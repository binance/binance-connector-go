/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the ChangePositionModeResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ChangePositionModeResponse{}

// ChangePositionModeResponse struct for ChangePositionModeResponse
type ChangePositionModeResponse struct {
	Code                 *int64  `json:"code,omitempty"`
	Msg                  *string `json:"msg,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChangePositionModeResponse ChangePositionModeResponse

// NewChangePositionModeResponse instantiates a new ChangePositionModeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangePositionModeResponse() *ChangePositionModeResponse {
	this := ChangePositionModeResponse{}
	return &this
}

// NewChangePositionModeResponseWithDefaults instantiates a new ChangePositionModeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangePositionModeResponseWithDefaults() *ChangePositionModeResponse {
	this := ChangePositionModeResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ChangePositionModeResponse) GetCode() int64 {
	if o == nil || common.IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangePositionModeResponse) GetCodeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ChangePositionModeResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *ChangePositionModeResponse) SetCode(v int64) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *ChangePositionModeResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangePositionModeResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *ChangePositionModeResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *ChangePositionModeResponse) SetMsg(v string) {
	o.Msg = &v
}

func (o ChangePositionModeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangePositionModeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

func (o *ChangePositionModeResponse) UnmarshalJSON(data []byte) (err error) {
	varChangePositionModeResponse := _ChangePositionModeResponse{}

	err = json.Unmarshal(data, &varChangePositionModeResponse)

	if err != nil {
		return err
	}

	*o = ChangePositionModeResponse(varChangePositionModeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "msg")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChangePositionModeResponse struct {
	value *ChangePositionModeResponse
	isSet bool
}

func (v NullableChangePositionModeResponse) Get() *ChangePositionModeResponse {
	return v.value
}

func (v *NullableChangePositionModeResponse) Set(val *ChangePositionModeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChangePositionModeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChangePositionModeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangePositionModeResponse(val *ChangePositionModeResponse) *NullableChangePositionModeResponse {
	return &NullableChangePositionModeResponse{value: val, isSet: true}
}

func (v NullableChangePositionModeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangePositionModeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
