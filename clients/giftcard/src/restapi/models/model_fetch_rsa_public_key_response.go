/*
Binance Gift Card REST API

OpenAPI Specification for the Binance Gift Card REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the FetchRsaPublicKeyResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FetchRsaPublicKeyResponse{}

// FetchRsaPublicKeyResponse struct for FetchRsaPublicKeyResponse
type FetchRsaPublicKeyResponse struct {
	Code                 *string `json:"code,omitempty"`
	Message              *string `json:"message,omitempty"`
	Data                 *string `json:"data,omitempty"`
	Success              *bool   `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FetchRsaPublicKeyResponse FetchRsaPublicKeyResponse

// NewFetchRsaPublicKeyResponse instantiates a new FetchRsaPublicKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchRsaPublicKeyResponse() *FetchRsaPublicKeyResponse {
	this := FetchRsaPublicKeyResponse{}
	return &this
}

// NewFetchRsaPublicKeyResponseWithDefaults instantiates a new FetchRsaPublicKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchRsaPublicKeyResponseWithDefaults() *FetchRsaPublicKeyResponse {
	this := FetchRsaPublicKeyResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *FetchRsaPublicKeyResponse) GetCode() string {
	if o == nil || common.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchRsaPublicKeyResponse) GetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *FetchRsaPublicKeyResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *FetchRsaPublicKeyResponse) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *FetchRsaPublicKeyResponse) GetMessage() string {
	if o == nil || common.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchRsaPublicKeyResponse) GetMessageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *FetchRsaPublicKeyResponse) HasMessage() bool {
	if o != nil && !common.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *FetchRsaPublicKeyResponse) SetMessage(v string) {
	o.Message = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *FetchRsaPublicKeyResponse) GetData() string {
	if o == nil || common.IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchRsaPublicKeyResponse) GetDataOk() (*string, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *FetchRsaPublicKeyResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *FetchRsaPublicKeyResponse) SetData(v string) {
	o.Data = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *FetchRsaPublicKeyResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchRsaPublicKeyResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *FetchRsaPublicKeyResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *FetchRsaPublicKeyResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o FetchRsaPublicKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchRsaPublicKeyResponse) ToMap() (map[string]interface{}, error) {
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

func (o *FetchRsaPublicKeyResponse) UnmarshalJSON(data []byte) (err error) {
	varFetchRsaPublicKeyResponse := _FetchRsaPublicKeyResponse{}

	err = json.Unmarshal(data, &varFetchRsaPublicKeyResponse)

	if err != nil {
		return err
	}

	*o = FetchRsaPublicKeyResponse(varFetchRsaPublicKeyResponse)

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

type NullableFetchRsaPublicKeyResponse struct {
	value *FetchRsaPublicKeyResponse
	isSet bool
}

func (v NullableFetchRsaPublicKeyResponse) Get() *FetchRsaPublicKeyResponse {
	return v.value
}

func (v *NullableFetchRsaPublicKeyResponse) Set(val *FetchRsaPublicKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchRsaPublicKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchRsaPublicKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchRsaPublicKeyResponse(val *FetchRsaPublicKeyResponse) *NullableFetchRsaPublicKeyResponse {
	return &NullableFetchRsaPublicKeyResponse{value: val, isSet: true}
}

func (v NullableFetchRsaPublicKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchRsaPublicKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
