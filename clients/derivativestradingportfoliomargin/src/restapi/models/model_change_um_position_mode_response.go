/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ChangeUmPositionModeResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ChangeUmPositionModeResponse{}

// ChangeUmPositionModeResponse struct for ChangeUmPositionModeResponse
type ChangeUmPositionModeResponse struct {
	Code                 *int64  `json:"code,omitempty"`
	Msg                  *string `json:"msg,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChangeUmPositionModeResponse ChangeUmPositionModeResponse

// NewChangeUmPositionModeResponse instantiates a new ChangeUmPositionModeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeUmPositionModeResponse() *ChangeUmPositionModeResponse {
	this := ChangeUmPositionModeResponse{}
	return &this
}

// NewChangeUmPositionModeResponseWithDefaults instantiates a new ChangeUmPositionModeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeUmPositionModeResponseWithDefaults() *ChangeUmPositionModeResponse {
	this := ChangeUmPositionModeResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ChangeUmPositionModeResponse) GetCode() int64 {
	if o == nil || common.IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeUmPositionModeResponse) GetCodeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ChangeUmPositionModeResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *ChangeUmPositionModeResponse) SetCode(v int64) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *ChangeUmPositionModeResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeUmPositionModeResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *ChangeUmPositionModeResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *ChangeUmPositionModeResponse) SetMsg(v string) {
	o.Msg = &v
}

func (o ChangeUmPositionModeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeUmPositionModeResponse) ToMap() (map[string]interface{}, error) {
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

func (o *ChangeUmPositionModeResponse) UnmarshalJSON(data []byte) (err error) {
	varChangeUmPositionModeResponse := _ChangeUmPositionModeResponse{}

	err = json.Unmarshal(data, &varChangeUmPositionModeResponse)

	if err != nil {
		return err
	}

	*o = ChangeUmPositionModeResponse(varChangeUmPositionModeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "msg")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChangeUmPositionModeResponse struct {
	value *ChangeUmPositionModeResponse
	isSet bool
}

func (v NullableChangeUmPositionModeResponse) Get() *ChangeUmPositionModeResponse {
	return v.value
}

func (v *NullableChangeUmPositionModeResponse) Set(val *ChangeUmPositionModeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeUmPositionModeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeUmPositionModeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeUmPositionModeResponse(val *ChangeUmPositionModeResponse) *NullableChangeUmPositionModeResponse {
	return &NullableChangeUmPositionModeResponse{value: val, isSet: true}
}

func (v NullableChangeUmPositionModeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeUmPositionModeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
