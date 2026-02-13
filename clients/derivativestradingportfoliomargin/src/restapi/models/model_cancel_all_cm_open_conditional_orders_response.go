/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the CancelAllCmOpenConditionalOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CancelAllCmOpenConditionalOrdersResponse{}

// CancelAllCmOpenConditionalOrdersResponse struct for CancelAllCmOpenConditionalOrdersResponse
type CancelAllCmOpenConditionalOrdersResponse struct {
	Code                 *string `json:"code,omitempty"`
	Msg                  *string `json:"msg,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CancelAllCmOpenConditionalOrdersResponse CancelAllCmOpenConditionalOrdersResponse

// NewCancelAllCmOpenConditionalOrdersResponse instantiates a new CancelAllCmOpenConditionalOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelAllCmOpenConditionalOrdersResponse() *CancelAllCmOpenConditionalOrdersResponse {
	this := CancelAllCmOpenConditionalOrdersResponse{}
	return &this
}

// NewCancelAllCmOpenConditionalOrdersResponseWithDefaults instantiates a new CancelAllCmOpenConditionalOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelAllCmOpenConditionalOrdersResponseWithDefaults() *CancelAllCmOpenConditionalOrdersResponse {
	this := CancelAllCmOpenConditionalOrdersResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CancelAllCmOpenConditionalOrdersResponse) GetCode() string {
	if o == nil || common.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelAllCmOpenConditionalOrdersResponse) GetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CancelAllCmOpenConditionalOrdersResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CancelAllCmOpenConditionalOrdersResponse) SetCode(v string) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *CancelAllCmOpenConditionalOrdersResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelAllCmOpenConditionalOrdersResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *CancelAllCmOpenConditionalOrdersResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *CancelAllCmOpenConditionalOrdersResponse) SetMsg(v string) {
	o.Msg = &v
}

func (o CancelAllCmOpenConditionalOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelAllCmOpenConditionalOrdersResponse) ToMap() (map[string]interface{}, error) {
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

func (o *CancelAllCmOpenConditionalOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	varCancelAllCmOpenConditionalOrdersResponse := _CancelAllCmOpenConditionalOrdersResponse{}

	err = json.Unmarshal(data, &varCancelAllCmOpenConditionalOrdersResponse)

	if err != nil {
		return err
	}

	*o = CancelAllCmOpenConditionalOrdersResponse(varCancelAllCmOpenConditionalOrdersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "msg")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCancelAllCmOpenConditionalOrdersResponse struct {
	value *CancelAllCmOpenConditionalOrdersResponse
	isSet bool
}

func (v NullableCancelAllCmOpenConditionalOrdersResponse) Get() *CancelAllCmOpenConditionalOrdersResponse {
	return v.value
}

func (v *NullableCancelAllCmOpenConditionalOrdersResponse) Set(val *CancelAllCmOpenConditionalOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelAllCmOpenConditionalOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelAllCmOpenConditionalOrdersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelAllCmOpenConditionalOrdersResponse(val *CancelAllCmOpenConditionalOrdersResponse) *NullableCancelAllCmOpenConditionalOrdersResponse {
	return &NullableCancelAllCmOpenConditionalOrdersResponse{value: val, isSet: true}
}

func (v NullableCancelAllCmOpenConditionalOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelAllCmOpenConditionalOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
