/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the ChangeMarginTypeResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ChangeMarginTypeResponse{}

// ChangeMarginTypeResponse struct for ChangeMarginTypeResponse
type ChangeMarginTypeResponse struct {
	Code                 *int64  `json:"code,omitempty"`
	Msg                  *string `json:"msg,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChangeMarginTypeResponse ChangeMarginTypeResponse

// NewChangeMarginTypeResponse instantiates a new ChangeMarginTypeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeMarginTypeResponse() *ChangeMarginTypeResponse {
	this := ChangeMarginTypeResponse{}
	return &this
}

// NewChangeMarginTypeResponseWithDefaults instantiates a new ChangeMarginTypeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeMarginTypeResponseWithDefaults() *ChangeMarginTypeResponse {
	this := ChangeMarginTypeResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ChangeMarginTypeResponse) GetCode() int64 {
	if o == nil || common.IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeMarginTypeResponse) GetCodeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ChangeMarginTypeResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *ChangeMarginTypeResponse) SetCode(v int64) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *ChangeMarginTypeResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeMarginTypeResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *ChangeMarginTypeResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *ChangeMarginTypeResponse) SetMsg(v string) {
	o.Msg = &v
}

func (o ChangeMarginTypeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeMarginTypeResponse) ToMap() (map[string]interface{}, error) {
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

func (o *ChangeMarginTypeResponse) UnmarshalJSON(data []byte) (err error) {
	varChangeMarginTypeResponse := _ChangeMarginTypeResponse{}

	err = json.Unmarshal(data, &varChangeMarginTypeResponse)

	if err != nil {
		return err
	}

	*o = ChangeMarginTypeResponse(varChangeMarginTypeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "msg")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChangeMarginTypeResponse struct {
	value *ChangeMarginTypeResponse
	isSet bool
}

func (v NullableChangeMarginTypeResponse) Get() *ChangeMarginTypeResponse {
	return v.value
}

func (v *NullableChangeMarginTypeResponse) Set(val *ChangeMarginTypeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeMarginTypeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeMarginTypeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeMarginTypeResponse(val *ChangeMarginTypeResponse) *NullableChangeMarginTypeResponse {
	return &NullableChangeMarginTypeResponse{value: val, isSet: true}
}

func (v NullableChangeMarginTypeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeMarginTypeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
