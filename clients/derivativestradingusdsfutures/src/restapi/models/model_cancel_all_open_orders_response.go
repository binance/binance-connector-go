/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CancelAllOpenOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CancelAllOpenOrdersResponse{}

// CancelAllOpenOrdersResponse struct for CancelAllOpenOrdersResponse
type CancelAllOpenOrdersResponse struct {
	Code                 *int64  `json:"code,omitempty"`
	Msg                  *string `json:"msg,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CancelAllOpenOrdersResponse CancelAllOpenOrdersResponse

// NewCancelAllOpenOrdersResponse instantiates a new CancelAllOpenOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelAllOpenOrdersResponse() *CancelAllOpenOrdersResponse {
	this := CancelAllOpenOrdersResponse{}
	return &this
}

// NewCancelAllOpenOrdersResponseWithDefaults instantiates a new CancelAllOpenOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelAllOpenOrdersResponseWithDefaults() *CancelAllOpenOrdersResponse {
	this := CancelAllOpenOrdersResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CancelAllOpenOrdersResponse) GetCode() int64 {
	if o == nil || common.IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelAllOpenOrdersResponse) GetCodeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CancelAllOpenOrdersResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *CancelAllOpenOrdersResponse) SetCode(v int64) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *CancelAllOpenOrdersResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelAllOpenOrdersResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *CancelAllOpenOrdersResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *CancelAllOpenOrdersResponse) SetMsg(v string) {
	o.Msg = &v
}

func (o CancelAllOpenOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelAllOpenOrdersResponse) ToMap() (map[string]interface{}, error) {
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

func (o *CancelAllOpenOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	varCancelAllOpenOrdersResponse := _CancelAllOpenOrdersResponse{}

	err = json.Unmarshal(data, &varCancelAllOpenOrdersResponse)

	if err != nil {
		return err
	}

	*o = CancelAllOpenOrdersResponse(varCancelAllOpenOrdersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "msg")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCancelAllOpenOrdersResponse struct {
	value *CancelAllOpenOrdersResponse
	isSet bool
}

func (v NullableCancelAllOpenOrdersResponse) Get() *CancelAllOpenOrdersResponse {
	return v.value
}

func (v *NullableCancelAllOpenOrdersResponse) Set(val *CancelAllOpenOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelAllOpenOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelAllOpenOrdersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelAllOpenOrdersResponse(val *CancelAllOpenOrdersResponse) *NullableCancelAllOpenOrdersResponse {
	return &NullableCancelAllOpenOrdersResponse{value: val, isSet: true}
}

func (v NullableCancelAllOpenOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelAllOpenOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
