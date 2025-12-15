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

// checks if the CreateADualTokenGiftCardResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CreateADualTokenGiftCardResponse{}

// CreateADualTokenGiftCardResponse struct for CreateADualTokenGiftCardResponse
type CreateADualTokenGiftCardResponse struct {
	Code                 *string                               `json:"code,omitempty"`
	Message              *string                               `json:"message,omitempty"`
	Data                 *CreateADualTokenGiftCardResponseData `json:"data,omitempty"`
	Success              *bool                                 `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateADualTokenGiftCardResponse CreateADualTokenGiftCardResponse

// NewCreateADualTokenGiftCardResponse instantiates a new CreateADualTokenGiftCardResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateADualTokenGiftCardResponse() *CreateADualTokenGiftCardResponse {
	this := CreateADualTokenGiftCardResponse{}
	return &this
}

// NewCreateADualTokenGiftCardResponseWithDefaults instantiates a new CreateADualTokenGiftCardResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateADualTokenGiftCardResponseWithDefaults() *CreateADualTokenGiftCardResponse {
	this := CreateADualTokenGiftCardResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CreateADualTokenGiftCardResponse) GetCode() string {
	if o == nil || common.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateADualTokenGiftCardResponse) GetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CreateADualTokenGiftCardResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CreateADualTokenGiftCardResponse) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CreateADualTokenGiftCardResponse) GetMessage() string {
	if o == nil || common.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateADualTokenGiftCardResponse) GetMessageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CreateADualTokenGiftCardResponse) HasMessage() bool {
	if o != nil && !common.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CreateADualTokenGiftCardResponse) SetMessage(v string) {
	o.Message = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateADualTokenGiftCardResponse) GetData() CreateADualTokenGiftCardResponseData {
	if o == nil || common.IsNil(o.Data) {
		var ret CreateADualTokenGiftCardResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateADualTokenGiftCardResponse) GetDataOk() (*CreateADualTokenGiftCardResponseData, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateADualTokenGiftCardResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CreateADualTokenGiftCardResponseData and assigns it to the Data field.
func (o *CreateADualTokenGiftCardResponse) SetData(v CreateADualTokenGiftCardResponseData) {
	o.Data = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *CreateADualTokenGiftCardResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateADualTokenGiftCardResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *CreateADualTokenGiftCardResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *CreateADualTokenGiftCardResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o CreateADualTokenGiftCardResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateADualTokenGiftCardResponse) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateADualTokenGiftCardResponse) UnmarshalJSON(data []byte) (err error) {
	varCreateADualTokenGiftCardResponse := _CreateADualTokenGiftCardResponse{}

	err = json.Unmarshal(data, &varCreateADualTokenGiftCardResponse)

	if err != nil {
		return err
	}

	*o = CreateADualTokenGiftCardResponse(varCreateADualTokenGiftCardResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "message")
		delete(additionalProperties, "data")
		delete(additionalProperties, "success")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateADualTokenGiftCardResponse struct {
	value *CreateADualTokenGiftCardResponse
	isSet bool
}

func (v NullableCreateADualTokenGiftCardResponse) Get() *CreateADualTokenGiftCardResponse {
	return v.value
}

func (v *NullableCreateADualTokenGiftCardResponse) Set(val *CreateADualTokenGiftCardResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateADualTokenGiftCardResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateADualTokenGiftCardResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateADualTokenGiftCardResponse(val *CreateADualTokenGiftCardResponse) *NullableCreateADualTokenGiftCardResponse {
	return &NullableCreateADualTokenGiftCardResponse{value: val, isSet: true}
}

func (v NullableCreateADualTokenGiftCardResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateADualTokenGiftCardResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
