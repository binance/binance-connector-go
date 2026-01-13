/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CancelAllUmOpenConditionalOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CancelAllUmOpenConditionalOrdersResponse{}

// CancelAllUmOpenConditionalOrdersResponse struct for CancelAllUmOpenConditionalOrdersResponse
type CancelAllUmOpenConditionalOrdersResponse struct {
	Code                 *string `json:"code,omitempty"`
	Msg                  *string `json:"msg,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CancelAllUmOpenConditionalOrdersResponse CancelAllUmOpenConditionalOrdersResponse

// NewCancelAllUmOpenConditionalOrdersResponse instantiates a new CancelAllUmOpenConditionalOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelAllUmOpenConditionalOrdersResponse() *CancelAllUmOpenConditionalOrdersResponse {
	this := CancelAllUmOpenConditionalOrdersResponse{}
	return &this
}

// NewCancelAllUmOpenConditionalOrdersResponseWithDefaults instantiates a new CancelAllUmOpenConditionalOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelAllUmOpenConditionalOrdersResponseWithDefaults() *CancelAllUmOpenConditionalOrdersResponse {
	this := CancelAllUmOpenConditionalOrdersResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CancelAllUmOpenConditionalOrdersResponse) GetCode() string {
	if o == nil || common.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelAllUmOpenConditionalOrdersResponse) GetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CancelAllUmOpenConditionalOrdersResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CancelAllUmOpenConditionalOrdersResponse) SetCode(v string) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *CancelAllUmOpenConditionalOrdersResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelAllUmOpenConditionalOrdersResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *CancelAllUmOpenConditionalOrdersResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *CancelAllUmOpenConditionalOrdersResponse) SetMsg(v string) {
	o.Msg = &v
}

func (o CancelAllUmOpenConditionalOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelAllUmOpenConditionalOrdersResponse) ToMap() (map[string]interface{}, error) {
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

func (o *CancelAllUmOpenConditionalOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	varCancelAllUmOpenConditionalOrdersResponse := _CancelAllUmOpenConditionalOrdersResponse{}

	err = json.Unmarshal(data, &varCancelAllUmOpenConditionalOrdersResponse)

	if err != nil {
		return err
	}

	*o = CancelAllUmOpenConditionalOrdersResponse(varCancelAllUmOpenConditionalOrdersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "msg")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCancelAllUmOpenConditionalOrdersResponse struct {
	value *CancelAllUmOpenConditionalOrdersResponse
	isSet bool
}

func (v NullableCancelAllUmOpenConditionalOrdersResponse) Get() *CancelAllUmOpenConditionalOrdersResponse {
	return v.value
}

func (v *NullableCancelAllUmOpenConditionalOrdersResponse) Set(val *CancelAllUmOpenConditionalOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelAllUmOpenConditionalOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelAllUmOpenConditionalOrdersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelAllUmOpenConditionalOrdersResponse(val *CancelAllUmOpenConditionalOrdersResponse) *NullableCancelAllUmOpenConditionalOrdersResponse {
	return &NullableCancelAllUmOpenConditionalOrdersResponse{value: val, isSet: true}
}

func (v NullableCancelAllUmOpenConditionalOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelAllUmOpenConditionalOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
