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

// checks if the TimeWeightedAveragePriceFutureAlgoResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TimeWeightedAveragePriceFutureAlgoResponse{}

// TimeWeightedAveragePriceFutureAlgoResponse struct for TimeWeightedAveragePriceFutureAlgoResponse
type TimeWeightedAveragePriceFutureAlgoResponse struct {
	ClientAlgoId         *string `json:"clientAlgoId,omitempty"`
	Success              *bool   `json:"success,omitempty"`
	Code                 *int64  `json:"code,omitempty"`
	Msg                  *string `json:"msg,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TimeWeightedAveragePriceFutureAlgoResponse TimeWeightedAveragePriceFutureAlgoResponse

// NewTimeWeightedAveragePriceFutureAlgoResponse instantiates a new TimeWeightedAveragePriceFutureAlgoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeWeightedAveragePriceFutureAlgoResponse() *TimeWeightedAveragePriceFutureAlgoResponse {
	this := TimeWeightedAveragePriceFutureAlgoResponse{}
	return &this
}

// NewTimeWeightedAveragePriceFutureAlgoResponseWithDefaults instantiates a new TimeWeightedAveragePriceFutureAlgoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeWeightedAveragePriceFutureAlgoResponseWithDefaults() *TimeWeightedAveragePriceFutureAlgoResponse {
	this := TimeWeightedAveragePriceFutureAlgoResponse{}
	return &this
}

// GetClientAlgoId returns the ClientAlgoId field value if set, zero value otherwise.
func (o *TimeWeightedAveragePriceFutureAlgoResponse) GetClientAlgoId() string {
	if o == nil || common.IsNil(o.ClientAlgoId) {
		var ret string
		return ret
	}
	return *o.ClientAlgoId
}

// GetClientAlgoIdOk returns a tuple with the ClientAlgoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeWeightedAveragePriceFutureAlgoResponse) GetClientAlgoIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientAlgoId) {
		return nil, false
	}
	return o.ClientAlgoId, true
}

// HasClientAlgoId returns a boolean if a field has been set.
func (o *TimeWeightedAveragePriceFutureAlgoResponse) HasClientAlgoId() bool {
	if o != nil && !common.IsNil(o.ClientAlgoId) {
		return true
	}

	return false
}

// SetClientAlgoId gets a reference to the given string and assigns it to the ClientAlgoId field.
func (o *TimeWeightedAveragePriceFutureAlgoResponse) SetClientAlgoId(v string) {
	o.ClientAlgoId = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *TimeWeightedAveragePriceFutureAlgoResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeWeightedAveragePriceFutureAlgoResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *TimeWeightedAveragePriceFutureAlgoResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *TimeWeightedAveragePriceFutureAlgoResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *TimeWeightedAveragePriceFutureAlgoResponse) GetCode() int64 {
	if o == nil || common.IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeWeightedAveragePriceFutureAlgoResponse) GetCodeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *TimeWeightedAveragePriceFutureAlgoResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *TimeWeightedAveragePriceFutureAlgoResponse) SetCode(v int64) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *TimeWeightedAveragePriceFutureAlgoResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeWeightedAveragePriceFutureAlgoResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *TimeWeightedAveragePriceFutureAlgoResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *TimeWeightedAveragePriceFutureAlgoResponse) SetMsg(v string) {
	o.Msg = &v
}

func (o TimeWeightedAveragePriceFutureAlgoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeWeightedAveragePriceFutureAlgoResponse) ToMap() (map[string]interface{}, error) {
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

func (o *TimeWeightedAveragePriceFutureAlgoResponse) UnmarshalJSON(data []byte) (err error) {
	varTimeWeightedAveragePriceFutureAlgoResponse := _TimeWeightedAveragePriceFutureAlgoResponse{}

	err = json.Unmarshal(data, &varTimeWeightedAveragePriceFutureAlgoResponse)

	if err != nil {
		return err
	}

	*o = TimeWeightedAveragePriceFutureAlgoResponse(varTimeWeightedAveragePriceFutureAlgoResponse)

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

type NullableTimeWeightedAveragePriceFutureAlgoResponse struct {
	value *TimeWeightedAveragePriceFutureAlgoResponse
	isSet bool
}

func (v NullableTimeWeightedAveragePriceFutureAlgoResponse) Get() *TimeWeightedAveragePriceFutureAlgoResponse {
	return v.value
}

func (v *NullableTimeWeightedAveragePriceFutureAlgoResponse) Set(val *TimeWeightedAveragePriceFutureAlgoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeWeightedAveragePriceFutureAlgoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeWeightedAveragePriceFutureAlgoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeWeightedAveragePriceFutureAlgoResponse(val *TimeWeightedAveragePriceFutureAlgoResponse) *NullableTimeWeightedAveragePriceFutureAlgoResponse {
	return &NullableTimeWeightedAveragePriceFutureAlgoResponse{value: val, isSet: true}
}

func (v NullableTimeWeightedAveragePriceFutureAlgoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeWeightedAveragePriceFutureAlgoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
