/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the FuturesTradfiPerpsContractResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FuturesTradfiPerpsContractResponse{}

// FuturesTradfiPerpsContractResponse struct for FuturesTradfiPerpsContractResponse
type FuturesTradfiPerpsContractResponse struct {
	Code                 *int64  `json:"code,omitempty"`
	Msg                  *string `json:"msg,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FuturesTradfiPerpsContractResponse FuturesTradfiPerpsContractResponse

// NewFuturesTradfiPerpsContractResponse instantiates a new FuturesTradfiPerpsContractResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFuturesTradfiPerpsContractResponse() *FuturesTradfiPerpsContractResponse {
	this := FuturesTradfiPerpsContractResponse{}
	return &this
}

// NewFuturesTradfiPerpsContractResponseWithDefaults instantiates a new FuturesTradfiPerpsContractResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFuturesTradfiPerpsContractResponseWithDefaults() *FuturesTradfiPerpsContractResponse {
	this := FuturesTradfiPerpsContractResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *FuturesTradfiPerpsContractResponse) GetCode() int64 {
	if o == nil || common.IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesTradfiPerpsContractResponse) GetCodeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *FuturesTradfiPerpsContractResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *FuturesTradfiPerpsContractResponse) SetCode(v int64) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *FuturesTradfiPerpsContractResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesTradfiPerpsContractResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *FuturesTradfiPerpsContractResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *FuturesTradfiPerpsContractResponse) SetMsg(v string) {
	o.Msg = &v
}

func (o FuturesTradfiPerpsContractResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FuturesTradfiPerpsContractResponse) ToMap() (map[string]interface{}, error) {
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

func (o *FuturesTradfiPerpsContractResponse) UnmarshalJSON(data []byte) (err error) {
	varFuturesTradfiPerpsContractResponse := _FuturesTradfiPerpsContractResponse{}

	err = json.Unmarshal(data, &varFuturesTradfiPerpsContractResponse)

	if err != nil {
		return err
	}

	*o = FuturesTradfiPerpsContractResponse(varFuturesTradfiPerpsContractResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "msg")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFuturesTradfiPerpsContractResponse struct {
	value *FuturesTradfiPerpsContractResponse
	isSet bool
}

func (v NullableFuturesTradfiPerpsContractResponse) Get() *FuturesTradfiPerpsContractResponse {
	return v.value
}

func (v *NullableFuturesTradfiPerpsContractResponse) Set(val *FuturesTradfiPerpsContractResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFuturesTradfiPerpsContractResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFuturesTradfiPerpsContractResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFuturesTradfiPerpsContractResponse(val *FuturesTradfiPerpsContractResponse) *NullableFuturesTradfiPerpsContractResponse {
	return &NullableFuturesTradfiPerpsContractResponse{value: val, isSet: true}
}

func (v NullableFuturesTradfiPerpsContractResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFuturesTradfiPerpsContractResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
