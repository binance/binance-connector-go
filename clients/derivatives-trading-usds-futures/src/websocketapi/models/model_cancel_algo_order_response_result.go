/*
Binance Derivatives Trading USDS Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CancelAlgoOrderResponseResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CancelAlgoOrderResponseResult{}

// CancelAlgoOrderResponseResult struct for CancelAlgoOrderResponseResult
type CancelAlgoOrderResponseResult struct {
	AlgoId               *int64  `json:"algoId,omitempty"`
	ClientAlgoId         *string `json:"clientAlgoId,omitempty"`
	Code                 *string `json:"code,omitempty"`
	Msg                  *string `json:"msg,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CancelAlgoOrderResponseResult CancelAlgoOrderResponseResult

// NewCancelAlgoOrderResponseResult instantiates a new CancelAlgoOrderResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelAlgoOrderResponseResult() *CancelAlgoOrderResponseResult {
	this := CancelAlgoOrderResponseResult{}
	return &this
}

// NewCancelAlgoOrderResponseResultWithDefaults instantiates a new CancelAlgoOrderResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelAlgoOrderResponseResultWithDefaults() *CancelAlgoOrderResponseResult {
	this := CancelAlgoOrderResponseResult{}
	return &this
}

// GetAlgoId returns the AlgoId field value if set, zero value otherwise.
func (o *CancelAlgoOrderResponseResult) GetAlgoId() int64 {
	if o == nil || common.IsNil(o.AlgoId) {
		var ret int64
		return ret
	}
	return *o.AlgoId
}

// GetAlgoIdOk returns a tuple with the AlgoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelAlgoOrderResponseResult) GetAlgoIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.AlgoId) {
		return nil, false
	}
	return o.AlgoId, true
}

// HasAlgoId returns a boolean if a field has been set.
func (o *CancelAlgoOrderResponseResult) HasAlgoId() bool {
	if o != nil && !common.IsNil(o.AlgoId) {
		return true
	}

	return false
}

// SetAlgoId gets a reference to the given int64 and assigns it to the AlgoId field.
func (o *CancelAlgoOrderResponseResult) SetAlgoId(v int64) {
	o.AlgoId = &v
}

// GetClientAlgoId returns the ClientAlgoId field value if set, zero value otherwise.
func (o *CancelAlgoOrderResponseResult) GetClientAlgoId() string {
	if o == nil || common.IsNil(o.ClientAlgoId) {
		var ret string
		return ret
	}
	return *o.ClientAlgoId
}

// GetClientAlgoIdOk returns a tuple with the ClientAlgoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelAlgoOrderResponseResult) GetClientAlgoIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientAlgoId) {
		return nil, false
	}
	return o.ClientAlgoId, true
}

// HasClientAlgoId returns a boolean if a field has been set.
func (o *CancelAlgoOrderResponseResult) HasClientAlgoId() bool {
	if o != nil && !common.IsNil(o.ClientAlgoId) {
		return true
	}

	return false
}

// SetClientAlgoId gets a reference to the given string and assigns it to the ClientAlgoId field.
func (o *CancelAlgoOrderResponseResult) SetClientAlgoId(v string) {
	o.ClientAlgoId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CancelAlgoOrderResponseResult) GetCode() string {
	if o == nil || common.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelAlgoOrderResponseResult) GetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CancelAlgoOrderResponseResult) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CancelAlgoOrderResponseResult) SetCode(v string) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *CancelAlgoOrderResponseResult) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelAlgoOrderResponseResult) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *CancelAlgoOrderResponseResult) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *CancelAlgoOrderResponseResult) SetMsg(v string) {
	o.Msg = &v
}

func (o CancelAlgoOrderResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelAlgoOrderResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AlgoId) {
		toSerialize["algoId"] = o.AlgoId
	}
	if !common.IsNil(o.ClientAlgoId) {
		toSerialize["clientAlgoId"] = o.ClientAlgoId
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

func (o *CancelAlgoOrderResponseResult) UnmarshalJSON(data []byte) (err error) {
	varCancelAlgoOrderResponseResult := _CancelAlgoOrderResponseResult{}

	err = json.Unmarshal(data, &varCancelAlgoOrderResponseResult)

	if err != nil {
		return err
	}

	*o = CancelAlgoOrderResponseResult(varCancelAlgoOrderResponseResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "algoId")
		delete(additionalProperties, "clientAlgoId")
		delete(additionalProperties, "code")
		delete(additionalProperties, "msg")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCancelAlgoOrderResponseResult struct {
	value *CancelAlgoOrderResponseResult
	isSet bool
}

func (v NullableCancelAlgoOrderResponseResult) Get() *CancelAlgoOrderResponseResult {
	return v.value
}

func (v *NullableCancelAlgoOrderResponseResult) Set(val *CancelAlgoOrderResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelAlgoOrderResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelAlgoOrderResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelAlgoOrderResponseResult(val *CancelAlgoOrderResponseResult) *NullableCancelAlgoOrderResponseResult {
	return &NullableCancelAlgoOrderResponseResult{value: val, isSet: true}
}

func (v NullableCancelAlgoOrderResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelAlgoOrderResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
