/*
Binance Fiat REST API

OpenAPI Specification for the Binance Fiat REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the FiatWithdrawResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FiatWithdrawResponse{}

// FiatWithdrawResponse struct for FiatWithdrawResponse
type FiatWithdrawResponse struct {
	Code                 *string              `json:"code,omitempty"`
	Message              *string              `json:"message,omitempty"`
	Data                 *DepositResponseData `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FiatWithdrawResponse FiatWithdrawResponse

// NewFiatWithdrawResponse instantiates a new FiatWithdrawResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiatWithdrawResponse() *FiatWithdrawResponse {
	this := FiatWithdrawResponse{}
	return &this
}

// NewFiatWithdrawResponseWithDefaults instantiates a new FiatWithdrawResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiatWithdrawResponseWithDefaults() *FiatWithdrawResponse {
	this := FiatWithdrawResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *FiatWithdrawResponse) GetCode() string {
	if o == nil || common.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiatWithdrawResponse) GetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *FiatWithdrawResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *FiatWithdrawResponse) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *FiatWithdrawResponse) GetMessage() string {
	if o == nil || common.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiatWithdrawResponse) GetMessageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *FiatWithdrawResponse) HasMessage() bool {
	if o != nil && !common.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *FiatWithdrawResponse) SetMessage(v string) {
	o.Message = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *FiatWithdrawResponse) GetData() DepositResponseData {
	if o == nil || common.IsNil(o.Data) {
		var ret DepositResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiatWithdrawResponse) GetDataOk() (*DepositResponseData, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *FiatWithdrawResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DepositResponseData and assigns it to the Data field.
func (o *FiatWithdrawResponse) SetData(v DepositResponseData) {
	o.Data = &v
}

func (o FiatWithdrawResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FiatWithdrawResponse) ToMap() (map[string]interface{}, error) {
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

func (o *FiatWithdrawResponse) UnmarshalJSON(data []byte) (err error) {
	varFiatWithdrawResponse := _FiatWithdrawResponse{}

	err = json.Unmarshal(data, &varFiatWithdrawResponse)

	if err != nil {
		return err
	}

	*o = FiatWithdrawResponse(varFiatWithdrawResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "message")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFiatWithdrawResponse struct {
	value *FiatWithdrawResponse
	isSet bool
}

func (v NullableFiatWithdrawResponse) Get() *FiatWithdrawResponse {
	return v.value
}

func (v *NullableFiatWithdrawResponse) Set(val *FiatWithdrawResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFiatWithdrawResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFiatWithdrawResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiatWithdrawResponse(val *FiatWithdrawResponse) *NullableFiatWithdrawResponse {
	return &NullableFiatWithdrawResponse{value: val, isSet: true}
}

func (v NullableFiatWithdrawResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiatWithdrawResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
