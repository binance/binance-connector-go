/*
Options REST API

Access market data, manage accounts, and trade Binance Options.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the TradfiOptionsContractResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TradfiOptionsContractResponse{}

// TradfiOptionsContractResponse struct for TradfiOptionsContractResponse
type TradfiOptionsContractResponse struct {
	Code                 *int64  `json:"code,omitempty"`
	Msg                  *string `json:"msg,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TradfiOptionsContractResponse TradfiOptionsContractResponse

// NewTradfiOptionsContractResponse instantiates a new TradfiOptionsContractResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradfiOptionsContractResponse() *TradfiOptionsContractResponse {
	this := TradfiOptionsContractResponse{}
	return &this
}

// NewTradfiOptionsContractResponseWithDefaults instantiates a new TradfiOptionsContractResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradfiOptionsContractResponseWithDefaults() *TradfiOptionsContractResponse {
	this := TradfiOptionsContractResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *TradfiOptionsContractResponse) GetCode() int64 {
	if o == nil || common.IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradfiOptionsContractResponse) GetCodeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *TradfiOptionsContractResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *TradfiOptionsContractResponse) SetCode(v int64) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *TradfiOptionsContractResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradfiOptionsContractResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *TradfiOptionsContractResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *TradfiOptionsContractResponse) SetMsg(v string) {
	o.Msg = &v
}

func (o TradfiOptionsContractResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TradfiOptionsContractResponse) ToMap() (map[string]interface{}, error) {
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

func (o *TradfiOptionsContractResponse) UnmarshalJSON(data []byte) (err error) {
	varTradfiOptionsContractResponse := _TradfiOptionsContractResponse{}

	err = json.Unmarshal(data, &varTradfiOptionsContractResponse)

	if err != nil {
		return err
	}

	*o = TradfiOptionsContractResponse(varTradfiOptionsContractResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "msg")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTradfiOptionsContractResponse struct {
	value *TradfiOptionsContractResponse
	isSet bool
}

func (v NullableTradfiOptionsContractResponse) Get() *TradfiOptionsContractResponse {
	return v.value
}

func (v *NullableTradfiOptionsContractResponse) Set(val *TradfiOptionsContractResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTradfiOptionsContractResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTradfiOptionsContractResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradfiOptionsContractResponse(val *TradfiOptionsContractResponse) *NullableTradfiOptionsContractResponse {
	return &NullableTradfiOptionsContractResponse{value: val, isSet: true}
}

func (v NullableTradfiOptionsContractResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradfiOptionsContractResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
