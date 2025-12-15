/*
Binance Fiat REST API

OpenAPI Specification for the Binance Fiat REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the DepositResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DepositResponse{}

// DepositResponse struct for DepositResponse
type DepositResponse struct {
	Code                 *string              `json:"code,omitempty"`
	Message              *string              `json:"message,omitempty"`
	Data                 *DepositResponseData `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DepositResponse DepositResponse

// NewDepositResponse instantiates a new DepositResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositResponse() *DepositResponse {
	this := DepositResponse{}
	return &this
}

// NewDepositResponseWithDefaults instantiates a new DepositResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositResponseWithDefaults() *DepositResponse {
	this := DepositResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *DepositResponse) GetCode() string {
	if o == nil || common.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositResponse) GetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *DepositResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *DepositResponse) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DepositResponse) GetMessage() string {
	if o == nil || common.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositResponse) GetMessageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DepositResponse) HasMessage() bool {
	if o != nil && !common.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DepositResponse) SetMessage(v string) {
	o.Message = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DepositResponse) GetData() DepositResponseData {
	if o == nil || common.IsNil(o.Data) {
		var ret DepositResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositResponse) GetDataOk() (*DepositResponseData, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DepositResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DepositResponseData and assigns it to the Data field.
func (o *DepositResponse) SetData(v DepositResponseData) {
	o.Data = &v
}

func (o DepositResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepositResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !common.IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DepositResponse) UnmarshalJSON(data []byte) (err error) {
	varDepositResponse := _DepositResponse{}

	err = json.Unmarshal(data, &varDepositResponse)

	if err != nil {
		return err
	}

	*o = DepositResponse(varDepositResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "message")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDepositResponse struct {
	value *DepositResponse
	isSet bool
}

func (v NullableDepositResponse) Get() *DepositResponse {
	return v.value
}

func (v *NullableDepositResponse) Set(val *DepositResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositResponse(val *DepositResponse) *NullableDepositResponse {
	return &NullableDepositResponse{value: val, isSet: true}
}

func (v NullableDepositResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
