/*
Binance Algo REST API

OpenAPI Specification for the Binance Algo REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CancelAlgoOrderFutureAlgoResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CancelAlgoOrderFutureAlgoResponse{}

// CancelAlgoOrderFutureAlgoResponse struct for CancelAlgoOrderFutureAlgoResponse
type CancelAlgoOrderFutureAlgoResponse struct {
	AlgoId               *int64  `json:"algoId,omitempty"`
	Success              *bool   `json:"success,omitempty"`
	Code                 *int64  `json:"code,omitempty"`
	Msg                  *string `json:"msg,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CancelAlgoOrderFutureAlgoResponse CancelAlgoOrderFutureAlgoResponse

// NewCancelAlgoOrderFutureAlgoResponse instantiates a new CancelAlgoOrderFutureAlgoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelAlgoOrderFutureAlgoResponse() *CancelAlgoOrderFutureAlgoResponse {
	this := CancelAlgoOrderFutureAlgoResponse{}
	return &this
}

// NewCancelAlgoOrderFutureAlgoResponseWithDefaults instantiates a new CancelAlgoOrderFutureAlgoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelAlgoOrderFutureAlgoResponseWithDefaults() *CancelAlgoOrderFutureAlgoResponse {
	this := CancelAlgoOrderFutureAlgoResponse{}
	return &this
}

// GetAlgoId returns the AlgoId field value if set, zero value otherwise.
func (o *CancelAlgoOrderFutureAlgoResponse) GetAlgoId() int64 {
	if o == nil || common.IsNil(o.AlgoId) {
		var ret int64
		return ret
	}
	return *o.AlgoId
}

// GetAlgoIdOk returns a tuple with the AlgoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelAlgoOrderFutureAlgoResponse) GetAlgoIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.AlgoId) {
		return nil, false
	}
	return o.AlgoId, true
}

// HasAlgoId returns a boolean if a field has been set.
func (o *CancelAlgoOrderFutureAlgoResponse) HasAlgoId() bool {
	if o != nil && !common.IsNil(o.AlgoId) {
		return true
	}

	return false
}

// SetAlgoId gets a reference to the given int64 and assigns it to the AlgoId field.
func (o *CancelAlgoOrderFutureAlgoResponse) SetAlgoId(v int64) {
	o.AlgoId = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *CancelAlgoOrderFutureAlgoResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelAlgoOrderFutureAlgoResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *CancelAlgoOrderFutureAlgoResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *CancelAlgoOrderFutureAlgoResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CancelAlgoOrderFutureAlgoResponse) GetCode() int64 {
	if o == nil || common.IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelAlgoOrderFutureAlgoResponse) GetCodeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CancelAlgoOrderFutureAlgoResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *CancelAlgoOrderFutureAlgoResponse) SetCode(v int64) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *CancelAlgoOrderFutureAlgoResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelAlgoOrderFutureAlgoResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *CancelAlgoOrderFutureAlgoResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *CancelAlgoOrderFutureAlgoResponse) SetMsg(v string) {
	o.Msg = &v
}

func (o CancelAlgoOrderFutureAlgoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelAlgoOrderFutureAlgoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AlgoId) {
		toSerialize["algoId"] = o.AlgoId
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

func (o *CancelAlgoOrderFutureAlgoResponse) UnmarshalJSON(data []byte) (err error) {
	varCancelAlgoOrderFutureAlgoResponse := _CancelAlgoOrderFutureAlgoResponse{}

	err = json.Unmarshal(data, &varCancelAlgoOrderFutureAlgoResponse)

	if err != nil {
		return err
	}

	*o = CancelAlgoOrderFutureAlgoResponse(varCancelAlgoOrderFutureAlgoResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "algoId")
		delete(additionalProperties, "success")
		delete(additionalProperties, "code")
		delete(additionalProperties, "msg")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCancelAlgoOrderFutureAlgoResponse struct {
	value *CancelAlgoOrderFutureAlgoResponse
	isSet bool
}

func (v NullableCancelAlgoOrderFutureAlgoResponse) Get() *CancelAlgoOrderFutureAlgoResponse {
	return v.value
}

func (v *NullableCancelAlgoOrderFutureAlgoResponse) Set(val *CancelAlgoOrderFutureAlgoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelAlgoOrderFutureAlgoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelAlgoOrderFutureAlgoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelAlgoOrderFutureAlgoResponse(val *CancelAlgoOrderFutureAlgoResponse) *NullableCancelAlgoOrderFutureAlgoResponse {
	return &NullableCancelAlgoOrderFutureAlgoResponse{value: val, isSet: true}
}

func (v NullableCancelAlgoOrderFutureAlgoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelAlgoOrderFutureAlgoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
