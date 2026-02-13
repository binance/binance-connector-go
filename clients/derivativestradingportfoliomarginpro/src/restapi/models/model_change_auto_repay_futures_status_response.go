/*
Binance Derivatives Trading Portfolio Margin Pro REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin Pro REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ChangeAutoRepayFuturesStatusResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ChangeAutoRepayFuturesStatusResponse{}

// ChangeAutoRepayFuturesStatusResponse struct for ChangeAutoRepayFuturesStatusResponse
type ChangeAutoRepayFuturesStatusResponse struct {
	Msg                  *string `json:"msg,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChangeAutoRepayFuturesStatusResponse ChangeAutoRepayFuturesStatusResponse

// NewChangeAutoRepayFuturesStatusResponse instantiates a new ChangeAutoRepayFuturesStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeAutoRepayFuturesStatusResponse() *ChangeAutoRepayFuturesStatusResponse {
	this := ChangeAutoRepayFuturesStatusResponse{}
	return &this
}

// NewChangeAutoRepayFuturesStatusResponseWithDefaults instantiates a new ChangeAutoRepayFuturesStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeAutoRepayFuturesStatusResponseWithDefaults() *ChangeAutoRepayFuturesStatusResponse {
	this := ChangeAutoRepayFuturesStatusResponse{}
	return &this
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *ChangeAutoRepayFuturesStatusResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeAutoRepayFuturesStatusResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *ChangeAutoRepayFuturesStatusResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *ChangeAutoRepayFuturesStatusResponse) SetMsg(v string) {
	o.Msg = &v
}

func (o ChangeAutoRepayFuturesStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeAutoRepayFuturesStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChangeAutoRepayFuturesStatusResponse) UnmarshalJSON(data []byte) (err error) {
	varChangeAutoRepayFuturesStatusResponse := _ChangeAutoRepayFuturesStatusResponse{}

	err = json.Unmarshal(data, &varChangeAutoRepayFuturesStatusResponse)

	if err != nil {
		return err
	}

	*o = ChangeAutoRepayFuturesStatusResponse(varChangeAutoRepayFuturesStatusResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "msg")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChangeAutoRepayFuturesStatusResponse struct {
	value *ChangeAutoRepayFuturesStatusResponse
	isSet bool
}

func (v NullableChangeAutoRepayFuturesStatusResponse) Get() *ChangeAutoRepayFuturesStatusResponse {
	return v.value
}

func (v *NullableChangeAutoRepayFuturesStatusResponse) Set(val *ChangeAutoRepayFuturesStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeAutoRepayFuturesStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeAutoRepayFuturesStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeAutoRepayFuturesStatusResponse(val *ChangeAutoRepayFuturesStatusResponse) *NullableChangeAutoRepayFuturesStatusResponse {
	return &NullableChangeAutoRepayFuturesStatusResponse{value: val, isSet: true}
}

func (v NullableChangeAutoRepayFuturesStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeAutoRepayFuturesStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
