/*
Binance Gift Card REST API

OpenAPI Specification for the Binance Gift Card REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CreateASingleTokenGiftCardResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CreateASingleTokenGiftCardResponse{}

// CreateASingleTokenGiftCardResponse struct for CreateASingleTokenGiftCardResponse
type CreateASingleTokenGiftCardResponse struct {
	Code                 *string                               `json:"code,omitempty"`
	Message              *string                               `json:"message,omitempty"`
	Data                 *CreateADualTokenGiftCardResponseData `json:"data,omitempty"`
	Success              *bool                                 `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateASingleTokenGiftCardResponse CreateASingleTokenGiftCardResponse

// NewCreateASingleTokenGiftCardResponse instantiates a new CreateASingleTokenGiftCardResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateASingleTokenGiftCardResponse() *CreateASingleTokenGiftCardResponse {
	this := CreateASingleTokenGiftCardResponse{}
	return &this
}

// NewCreateASingleTokenGiftCardResponseWithDefaults instantiates a new CreateASingleTokenGiftCardResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateASingleTokenGiftCardResponseWithDefaults() *CreateASingleTokenGiftCardResponse {
	this := CreateASingleTokenGiftCardResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CreateASingleTokenGiftCardResponse) GetCode() string {
	if o == nil || common.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateASingleTokenGiftCardResponse) GetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CreateASingleTokenGiftCardResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CreateASingleTokenGiftCardResponse) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CreateASingleTokenGiftCardResponse) GetMessage() string {
	if o == nil || common.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateASingleTokenGiftCardResponse) GetMessageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CreateASingleTokenGiftCardResponse) HasMessage() bool {
	if o != nil && !common.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CreateASingleTokenGiftCardResponse) SetMessage(v string) {
	o.Message = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateASingleTokenGiftCardResponse) GetData() CreateADualTokenGiftCardResponseData {
	if o == nil || common.IsNil(o.Data) {
		var ret CreateADualTokenGiftCardResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateASingleTokenGiftCardResponse) GetDataOk() (*CreateADualTokenGiftCardResponseData, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateASingleTokenGiftCardResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CreateADualTokenGiftCardResponseData and assigns it to the Data field.
func (o *CreateASingleTokenGiftCardResponse) SetData(v CreateADualTokenGiftCardResponseData) {
	o.Data = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *CreateASingleTokenGiftCardResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateASingleTokenGiftCardResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *CreateASingleTokenGiftCardResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *CreateASingleTokenGiftCardResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o CreateASingleTokenGiftCardResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateASingleTokenGiftCardResponse) ToMap() (map[string]interface{}, error) {
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

func (o *CreateASingleTokenGiftCardResponse) UnmarshalJSON(data []byte) (err error) {
	varCreateASingleTokenGiftCardResponse := _CreateASingleTokenGiftCardResponse{}

	err = json.Unmarshal(data, &varCreateASingleTokenGiftCardResponse)

	if err != nil {
		return err
	}

	*o = CreateASingleTokenGiftCardResponse(varCreateASingleTokenGiftCardResponse)

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

type NullableCreateASingleTokenGiftCardResponse struct {
	value *CreateASingleTokenGiftCardResponse
	isSet bool
}

func (v NullableCreateASingleTokenGiftCardResponse) Get() *CreateASingleTokenGiftCardResponse {
	return v.value
}

func (v *NullableCreateASingleTokenGiftCardResponse) Set(val *CreateASingleTokenGiftCardResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateASingleTokenGiftCardResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateASingleTokenGiftCardResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateASingleTokenGiftCardResponse(val *CreateASingleTokenGiftCardResponse) *NullableCreateASingleTokenGiftCardResponse {
	return &NullableCreateASingleTokenGiftCardResponse{value: val, isSet: true}
}

func (v NullableCreateASingleTokenGiftCardResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateASingleTokenGiftCardResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
