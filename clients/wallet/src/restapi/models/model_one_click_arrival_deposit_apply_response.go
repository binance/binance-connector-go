/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OneClickArrivalDepositApplyResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OneClickArrivalDepositApplyResponse{}

// OneClickArrivalDepositApplyResponse struct for OneClickArrivalDepositApplyResponse
type OneClickArrivalDepositApplyResponse struct {
	Code                 *string `json:"code,omitempty"`
	Message              *string `json:"message,omitempty"`
	Data                 *bool   `json:"data,omitempty"`
	Success              *bool   `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OneClickArrivalDepositApplyResponse OneClickArrivalDepositApplyResponse

// NewOneClickArrivalDepositApplyResponse instantiates a new OneClickArrivalDepositApplyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOneClickArrivalDepositApplyResponse() *OneClickArrivalDepositApplyResponse {
	this := OneClickArrivalDepositApplyResponse{}
	return &this
}

// NewOneClickArrivalDepositApplyResponseWithDefaults instantiates a new OneClickArrivalDepositApplyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOneClickArrivalDepositApplyResponseWithDefaults() *OneClickArrivalDepositApplyResponse {
	this := OneClickArrivalDepositApplyResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OneClickArrivalDepositApplyResponse) GetCode() string {
	if o == nil || common.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneClickArrivalDepositApplyResponse) GetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *OneClickArrivalDepositApplyResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *OneClickArrivalDepositApplyResponse) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *OneClickArrivalDepositApplyResponse) GetMessage() string {
	if o == nil || common.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneClickArrivalDepositApplyResponse) GetMessageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *OneClickArrivalDepositApplyResponse) HasMessage() bool {
	if o != nil && !common.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *OneClickArrivalDepositApplyResponse) SetMessage(v string) {
	o.Message = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OneClickArrivalDepositApplyResponse) GetData() bool {
	if o == nil || common.IsNil(o.Data) {
		var ret bool
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneClickArrivalDepositApplyResponse) GetDataOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OneClickArrivalDepositApplyResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given bool and assigns it to the Data field.
func (o *OneClickArrivalDepositApplyResponse) SetData(v bool) {
	o.Data = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *OneClickArrivalDepositApplyResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneClickArrivalDepositApplyResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *OneClickArrivalDepositApplyResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *OneClickArrivalDepositApplyResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o OneClickArrivalDepositApplyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OneClickArrivalDepositApplyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !common.IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OneClickArrivalDepositApplyResponse) UnmarshalJSON(data []byte) (err error) {
	varOneClickArrivalDepositApplyResponse := _OneClickArrivalDepositApplyResponse{}

	err = json.Unmarshal(data, &varOneClickArrivalDepositApplyResponse)

	if err != nil {
		return err
	}

	*o = OneClickArrivalDepositApplyResponse(varOneClickArrivalDepositApplyResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "message")
		delete(additionalProperties, "data")
		delete(additionalProperties, "success")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOneClickArrivalDepositApplyResponse struct {
	value *OneClickArrivalDepositApplyResponse
	isSet bool
}

func (v NullableOneClickArrivalDepositApplyResponse) Get() *OneClickArrivalDepositApplyResponse {
	return v.value
}

func (v *NullableOneClickArrivalDepositApplyResponse) Set(val *OneClickArrivalDepositApplyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOneClickArrivalDepositApplyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOneClickArrivalDepositApplyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOneClickArrivalDepositApplyResponse(val *OneClickArrivalDepositApplyResponse) *NullableOneClickArrivalDepositApplyResponse {
	return &NullableOneClickArrivalDepositApplyResponse{value: val, isSet: true}
}

func (v NullableOneClickArrivalDepositApplyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOneClickArrivalDepositApplyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
