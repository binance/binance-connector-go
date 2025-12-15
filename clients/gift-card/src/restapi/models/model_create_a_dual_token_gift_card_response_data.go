/*
Binance Gift Card REST API

OpenAPI Specification for the Binance Gift Card REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CreateADualTokenGiftCardResponseData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CreateADualTokenGiftCardResponseData{}

// CreateADualTokenGiftCardResponseData struct for CreateADualTokenGiftCardResponseData
type CreateADualTokenGiftCardResponseData struct {
	ReferenceNo          *string `json:"referenceNo,omitempty"`
	Code                 *string `json:"code,omitempty"`
	ExpiredTime          *int64  `json:"expiredTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateADualTokenGiftCardResponseData CreateADualTokenGiftCardResponseData

// NewCreateADualTokenGiftCardResponseData instantiates a new CreateADualTokenGiftCardResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateADualTokenGiftCardResponseData() *CreateADualTokenGiftCardResponseData {
	this := CreateADualTokenGiftCardResponseData{}
	return &this
}

// NewCreateADualTokenGiftCardResponseDataWithDefaults instantiates a new CreateADualTokenGiftCardResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateADualTokenGiftCardResponseDataWithDefaults() *CreateADualTokenGiftCardResponseData {
	this := CreateADualTokenGiftCardResponseData{}
	return &this
}

// GetReferenceNo returns the ReferenceNo field value if set, zero value otherwise.
func (o *CreateADualTokenGiftCardResponseData) GetReferenceNo() string {
	if o == nil || common.IsNil(o.ReferenceNo) {
		var ret string
		return ret
	}
	return *o.ReferenceNo
}

// GetReferenceNoOk returns a tuple with the ReferenceNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateADualTokenGiftCardResponseData) GetReferenceNoOk() (*string, bool) {
	if o == nil || common.IsNil(o.ReferenceNo) {
		return nil, false
	}
	return o.ReferenceNo, true
}

// HasReferenceNo returns a boolean if a field has been set.
func (o *CreateADualTokenGiftCardResponseData) HasReferenceNo() bool {
	if o != nil && !common.IsNil(o.ReferenceNo) {
		return true
	}

	return false
}

// SetReferenceNo gets a reference to the given string and assigns it to the ReferenceNo field.
func (o *CreateADualTokenGiftCardResponseData) SetReferenceNo(v string) {
	o.ReferenceNo = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CreateADualTokenGiftCardResponseData) GetCode() string {
	if o == nil || common.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateADualTokenGiftCardResponseData) GetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CreateADualTokenGiftCardResponseData) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CreateADualTokenGiftCardResponseData) SetCode(v string) {
	o.Code = &v
}

// GetExpiredTime returns the ExpiredTime field value if set, zero value otherwise.
func (o *CreateADualTokenGiftCardResponseData) GetExpiredTime() int64 {
	if o == nil || common.IsNil(o.ExpiredTime) {
		var ret int64
		return ret
	}
	return *o.ExpiredTime
}

// GetExpiredTimeOk returns a tuple with the ExpiredTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateADualTokenGiftCardResponseData) GetExpiredTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ExpiredTime) {
		return nil, false
	}
	return o.ExpiredTime, true
}

// HasExpiredTime returns a boolean if a field has been set.
func (o *CreateADualTokenGiftCardResponseData) HasExpiredTime() bool {
	if o != nil && !common.IsNil(o.ExpiredTime) {
		return true
	}

	return false
}

// SetExpiredTime gets a reference to the given int64 and assigns it to the ExpiredTime field.
func (o *CreateADualTokenGiftCardResponseData) SetExpiredTime(v int64) {
	o.ExpiredTime = &v
}

func (o CreateADualTokenGiftCardResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateADualTokenGiftCardResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ReferenceNo) {
		toSerialize["referenceNo"] = o.ReferenceNo
	}
	if !common.IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !common.IsNil(o.ExpiredTime) {
		toSerialize["expiredTime"] = o.ExpiredTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateADualTokenGiftCardResponseData) UnmarshalJSON(data []byte) (err error) {
	varCreateADualTokenGiftCardResponseData := _CreateADualTokenGiftCardResponseData{}

	err = json.Unmarshal(data, &varCreateADualTokenGiftCardResponseData)

	if err != nil {
		return err
	}

	*o = CreateADualTokenGiftCardResponseData(varCreateADualTokenGiftCardResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "referenceNo")
		delete(additionalProperties, "code")
		delete(additionalProperties, "expiredTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateADualTokenGiftCardResponseData struct {
	value *CreateADualTokenGiftCardResponseData
	isSet bool
}

func (v NullableCreateADualTokenGiftCardResponseData) Get() *CreateADualTokenGiftCardResponseData {
	return v.value
}

func (v *NullableCreateADualTokenGiftCardResponseData) Set(val *CreateADualTokenGiftCardResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateADualTokenGiftCardResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateADualTokenGiftCardResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateADualTokenGiftCardResponseData(val *CreateADualTokenGiftCardResponseData) *NullableCreateADualTokenGiftCardResponseData {
	return &NullableCreateADualTokenGiftCardResponseData{value: val, isSet: true}
}

func (v NullableCreateADualTokenGiftCardResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateADualTokenGiftCardResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
