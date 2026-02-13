/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the CancelAllAlgoOpenOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CancelAllAlgoOpenOrdersResponse{}

// CancelAllAlgoOpenOrdersResponse struct for CancelAllAlgoOpenOrdersResponse
type CancelAllAlgoOpenOrdersResponse struct {
	Code                 *int64  `json:"code,omitempty"`
	Msg                  *string `json:"msg,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CancelAllAlgoOpenOrdersResponse CancelAllAlgoOpenOrdersResponse

// NewCancelAllAlgoOpenOrdersResponse instantiates a new CancelAllAlgoOpenOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelAllAlgoOpenOrdersResponse() *CancelAllAlgoOpenOrdersResponse {
	this := CancelAllAlgoOpenOrdersResponse{}
	return &this
}

// NewCancelAllAlgoOpenOrdersResponseWithDefaults instantiates a new CancelAllAlgoOpenOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelAllAlgoOpenOrdersResponseWithDefaults() *CancelAllAlgoOpenOrdersResponse {
	this := CancelAllAlgoOpenOrdersResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CancelAllAlgoOpenOrdersResponse) GetCode() int64 {
	if o == nil || common.IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelAllAlgoOpenOrdersResponse) GetCodeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CancelAllAlgoOpenOrdersResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *CancelAllAlgoOpenOrdersResponse) SetCode(v int64) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *CancelAllAlgoOpenOrdersResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelAllAlgoOpenOrdersResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *CancelAllAlgoOpenOrdersResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *CancelAllAlgoOpenOrdersResponse) SetMsg(v string) {
	o.Msg = &v
}

func (o CancelAllAlgoOpenOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelAllAlgoOpenOrdersResponse) ToMap() (map[string]interface{}, error) {
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

func (o *CancelAllAlgoOpenOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	varCancelAllAlgoOpenOrdersResponse := _CancelAllAlgoOpenOrdersResponse{}

	err = json.Unmarshal(data, &varCancelAllAlgoOpenOrdersResponse)

	if err != nil {
		return err
	}

	*o = CancelAllAlgoOpenOrdersResponse(varCancelAllAlgoOpenOrdersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "msg")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCancelAllAlgoOpenOrdersResponse struct {
	value *CancelAllAlgoOpenOrdersResponse
	isSet bool
}

func (v NullableCancelAllAlgoOpenOrdersResponse) Get() *CancelAllAlgoOpenOrdersResponse {
	return v.value
}

func (v *NullableCancelAllAlgoOpenOrdersResponse) Set(val *CancelAllAlgoOpenOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelAllAlgoOpenOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelAllAlgoOpenOrdersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelAllAlgoOpenOrdersResponse(val *CancelAllAlgoOpenOrdersResponse) *NullableCancelAllAlgoOpenOrdersResponse {
	return &NullableCancelAllAlgoOpenOrdersResponse{value: val, isSet: true}
}

func (v NullableCancelAllAlgoOpenOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelAllAlgoOpenOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
