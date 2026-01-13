/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the ModifyIsolatedPositionMarginResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ModifyIsolatedPositionMarginResponse{}

// ModifyIsolatedPositionMarginResponse struct for ModifyIsolatedPositionMarginResponse
type ModifyIsolatedPositionMarginResponse struct {
	Amount               *float32 `json:"amount,omitempty"`
	Code                 *int64   `json:"code,omitempty"`
	Msg                  *string  `json:"msg,omitempty"`
	Type                 *int64   `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModifyIsolatedPositionMarginResponse ModifyIsolatedPositionMarginResponse

// NewModifyIsolatedPositionMarginResponse instantiates a new ModifyIsolatedPositionMarginResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyIsolatedPositionMarginResponse() *ModifyIsolatedPositionMarginResponse {
	this := ModifyIsolatedPositionMarginResponse{}
	return &this
}

// NewModifyIsolatedPositionMarginResponseWithDefaults instantiates a new ModifyIsolatedPositionMarginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyIsolatedPositionMarginResponseWithDefaults() *ModifyIsolatedPositionMarginResponse {
	this := ModifyIsolatedPositionMarginResponse{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ModifyIsolatedPositionMarginResponse) GetAmount() float32 {
	if o == nil || common.IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIsolatedPositionMarginResponse) GetAmountOk() (*float32, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ModifyIsolatedPositionMarginResponse) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *ModifyIsolatedPositionMarginResponse) SetAmount(v float32) {
	o.Amount = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ModifyIsolatedPositionMarginResponse) GetCode() int64 {
	if o == nil || common.IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIsolatedPositionMarginResponse) GetCodeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ModifyIsolatedPositionMarginResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *ModifyIsolatedPositionMarginResponse) SetCode(v int64) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *ModifyIsolatedPositionMarginResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIsolatedPositionMarginResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *ModifyIsolatedPositionMarginResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *ModifyIsolatedPositionMarginResponse) SetMsg(v string) {
	o.Msg = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ModifyIsolatedPositionMarginResponse) GetType() int64 {
	if o == nil || common.IsNil(o.Type) {
		var ret int64
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIsolatedPositionMarginResponse) GetTypeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ModifyIsolatedPositionMarginResponse) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int64 and assigns it to the Type field.
func (o *ModifyIsolatedPositionMarginResponse) SetType(v int64) {
	o.Type = &v
}

func (o ModifyIsolatedPositionMarginResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyIsolatedPositionMarginResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !common.IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModifyIsolatedPositionMarginResponse) UnmarshalJSON(data []byte) (err error) {
	varModifyIsolatedPositionMarginResponse := _ModifyIsolatedPositionMarginResponse{}

	err = json.Unmarshal(data, &varModifyIsolatedPositionMarginResponse)

	if err != nil {
		return err
	}

	*o = ModifyIsolatedPositionMarginResponse(varModifyIsolatedPositionMarginResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "code")
		delete(additionalProperties, "msg")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModifyIsolatedPositionMarginResponse struct {
	value *ModifyIsolatedPositionMarginResponse
	isSet bool
}

func (v NullableModifyIsolatedPositionMarginResponse) Get() *ModifyIsolatedPositionMarginResponse {
	return v.value
}

func (v *NullableModifyIsolatedPositionMarginResponse) Set(val *ModifyIsolatedPositionMarginResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyIsolatedPositionMarginResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyIsolatedPositionMarginResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyIsolatedPositionMarginResponse(val *ModifyIsolatedPositionMarginResponse) *NullableModifyIsolatedPositionMarginResponse {
	return &NullableModifyIsolatedPositionMarginResponse{value: val, isSet: true}
}

func (v NullableModifyIsolatedPositionMarginResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyIsolatedPositionMarginResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
