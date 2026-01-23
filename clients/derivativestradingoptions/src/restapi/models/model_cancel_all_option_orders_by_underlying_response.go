/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CancelAllOptionOrdersByUnderlyingResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CancelAllOptionOrdersByUnderlyingResponse{}

// CancelAllOptionOrdersByUnderlyingResponse struct for CancelAllOptionOrdersByUnderlyingResponse
type CancelAllOptionOrdersByUnderlyingResponse struct {
	Code                 *int64  `json:"code,omitempty"`
	Msg                  *string `json:"msg,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CancelAllOptionOrdersByUnderlyingResponse CancelAllOptionOrdersByUnderlyingResponse

// NewCancelAllOptionOrdersByUnderlyingResponse instantiates a new CancelAllOptionOrdersByUnderlyingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelAllOptionOrdersByUnderlyingResponse() *CancelAllOptionOrdersByUnderlyingResponse {
	this := CancelAllOptionOrdersByUnderlyingResponse{}
	return &this
}

// NewCancelAllOptionOrdersByUnderlyingResponseWithDefaults instantiates a new CancelAllOptionOrdersByUnderlyingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelAllOptionOrdersByUnderlyingResponseWithDefaults() *CancelAllOptionOrdersByUnderlyingResponse {
	this := CancelAllOptionOrdersByUnderlyingResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CancelAllOptionOrdersByUnderlyingResponse) GetCode() int64 {
	if o == nil || common.IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelAllOptionOrdersByUnderlyingResponse) GetCodeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CancelAllOptionOrdersByUnderlyingResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *CancelAllOptionOrdersByUnderlyingResponse) SetCode(v int64) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *CancelAllOptionOrdersByUnderlyingResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelAllOptionOrdersByUnderlyingResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *CancelAllOptionOrdersByUnderlyingResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *CancelAllOptionOrdersByUnderlyingResponse) SetMsg(v string) {
	o.Msg = &v
}

func (o CancelAllOptionOrdersByUnderlyingResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelAllOptionOrdersByUnderlyingResponse) ToMap() (map[string]interface{}, error) {
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

func (o *CancelAllOptionOrdersByUnderlyingResponse) UnmarshalJSON(data []byte) (err error) {
	varCancelAllOptionOrdersByUnderlyingResponse := _CancelAllOptionOrdersByUnderlyingResponse{}

	err = json.Unmarshal(data, &varCancelAllOptionOrdersByUnderlyingResponse)

	if err != nil {
		return err
	}

	*o = CancelAllOptionOrdersByUnderlyingResponse(varCancelAllOptionOrdersByUnderlyingResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "msg")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCancelAllOptionOrdersByUnderlyingResponse struct {
	value *CancelAllOptionOrdersByUnderlyingResponse
	isSet bool
}

func (v NullableCancelAllOptionOrdersByUnderlyingResponse) Get() *CancelAllOptionOrdersByUnderlyingResponse {
	return v.value
}

func (v *NullableCancelAllOptionOrdersByUnderlyingResponse) Set(val *CancelAllOptionOrdersByUnderlyingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelAllOptionOrdersByUnderlyingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelAllOptionOrdersByUnderlyingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelAllOptionOrdersByUnderlyingResponse(val *CancelAllOptionOrdersByUnderlyingResponse) *NullableCancelAllOptionOrdersByUnderlyingResponse {
	return &NullableCancelAllOptionOrdersByUnderlyingResponse{value: val, isSet: true}
}

func (v NullableCancelAllOptionOrdersByUnderlyingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelAllOptionOrdersByUnderlyingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
