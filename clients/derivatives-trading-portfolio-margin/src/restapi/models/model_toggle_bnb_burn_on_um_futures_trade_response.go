/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the ToggleBnbBurnOnUmFuturesTradeResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ToggleBnbBurnOnUmFuturesTradeResponse{}

// ToggleBnbBurnOnUmFuturesTradeResponse struct for ToggleBnbBurnOnUmFuturesTradeResponse
type ToggleBnbBurnOnUmFuturesTradeResponse struct {
	Code                 *int64  `json:"code,omitempty"`
	Msg                  *string `json:"msg,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ToggleBnbBurnOnUmFuturesTradeResponse ToggleBnbBurnOnUmFuturesTradeResponse

// NewToggleBnbBurnOnUmFuturesTradeResponse instantiates a new ToggleBnbBurnOnUmFuturesTradeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToggleBnbBurnOnUmFuturesTradeResponse() *ToggleBnbBurnOnUmFuturesTradeResponse {
	this := ToggleBnbBurnOnUmFuturesTradeResponse{}
	return &this
}

// NewToggleBnbBurnOnUmFuturesTradeResponseWithDefaults instantiates a new ToggleBnbBurnOnUmFuturesTradeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToggleBnbBurnOnUmFuturesTradeResponseWithDefaults() *ToggleBnbBurnOnUmFuturesTradeResponse {
	this := ToggleBnbBurnOnUmFuturesTradeResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ToggleBnbBurnOnUmFuturesTradeResponse) GetCode() int64 {
	if o == nil || common.IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToggleBnbBurnOnUmFuturesTradeResponse) GetCodeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ToggleBnbBurnOnUmFuturesTradeResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *ToggleBnbBurnOnUmFuturesTradeResponse) SetCode(v int64) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *ToggleBnbBurnOnUmFuturesTradeResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToggleBnbBurnOnUmFuturesTradeResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *ToggleBnbBurnOnUmFuturesTradeResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *ToggleBnbBurnOnUmFuturesTradeResponse) SetMsg(v string) {
	o.Msg = &v
}

func (o ToggleBnbBurnOnUmFuturesTradeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToggleBnbBurnOnUmFuturesTradeResponse) ToMap() (map[string]interface{}, error) {
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

func (o *ToggleBnbBurnOnUmFuturesTradeResponse) UnmarshalJSON(data []byte) (err error) {
	varToggleBnbBurnOnUmFuturesTradeResponse := _ToggleBnbBurnOnUmFuturesTradeResponse{}

	err = json.Unmarshal(data, &varToggleBnbBurnOnUmFuturesTradeResponse)

	if err != nil {
		return err
	}

	*o = ToggleBnbBurnOnUmFuturesTradeResponse(varToggleBnbBurnOnUmFuturesTradeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "msg")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableToggleBnbBurnOnUmFuturesTradeResponse struct {
	value *ToggleBnbBurnOnUmFuturesTradeResponse
	isSet bool
}

func (v NullableToggleBnbBurnOnUmFuturesTradeResponse) Get() *ToggleBnbBurnOnUmFuturesTradeResponse {
	return v.value
}

func (v *NullableToggleBnbBurnOnUmFuturesTradeResponse) Set(val *ToggleBnbBurnOnUmFuturesTradeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableToggleBnbBurnOnUmFuturesTradeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableToggleBnbBurnOnUmFuturesTradeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToggleBnbBurnOnUmFuturesTradeResponse(val *ToggleBnbBurnOnUmFuturesTradeResponse) *NullableToggleBnbBurnOnUmFuturesTradeResponse {
	return &NullableToggleBnbBurnOnUmFuturesTradeResponse{value: val, isSet: true}
}

func (v NullableToggleBnbBurnOnUmFuturesTradeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToggleBnbBurnOnUmFuturesTradeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
