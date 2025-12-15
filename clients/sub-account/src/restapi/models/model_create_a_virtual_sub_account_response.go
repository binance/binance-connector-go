/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CreateAVirtualSubAccountResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CreateAVirtualSubAccountResponse{}

// CreateAVirtualSubAccountResponse struct for CreateAVirtualSubAccountResponse
type CreateAVirtualSubAccountResponse struct {
	Email                *string `json:"email,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateAVirtualSubAccountResponse CreateAVirtualSubAccountResponse

// NewCreateAVirtualSubAccountResponse instantiates a new CreateAVirtualSubAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAVirtualSubAccountResponse() *CreateAVirtualSubAccountResponse {
	this := CreateAVirtualSubAccountResponse{}
	return &this
}

// NewCreateAVirtualSubAccountResponseWithDefaults instantiates a new CreateAVirtualSubAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAVirtualSubAccountResponseWithDefaults() *CreateAVirtualSubAccountResponse {
	this := CreateAVirtualSubAccountResponse{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CreateAVirtualSubAccountResponse) GetEmail() string {
	if o == nil || common.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAVirtualSubAccountResponse) GetEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CreateAVirtualSubAccountResponse) HasEmail() bool {
	if o != nil && !common.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CreateAVirtualSubAccountResponse) SetEmail(v string) {
	o.Email = &v
}

func (o CreateAVirtualSubAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAVirtualSubAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateAVirtualSubAccountResponse) UnmarshalJSON(data []byte) (err error) {
	varCreateAVirtualSubAccountResponse := _CreateAVirtualSubAccountResponse{}

	err = json.Unmarshal(data, &varCreateAVirtualSubAccountResponse)

	if err != nil {
		return err
	}

	*o = CreateAVirtualSubAccountResponse(varCreateAVirtualSubAccountResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateAVirtualSubAccountResponse struct {
	value *CreateAVirtualSubAccountResponse
	isSet bool
}

func (v NullableCreateAVirtualSubAccountResponse) Get() *CreateAVirtualSubAccountResponse {
	return v.value
}

func (v *NullableCreateAVirtualSubAccountResponse) Set(val *CreateAVirtualSubAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAVirtualSubAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAVirtualSubAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAVirtualSubAccountResponse(val *CreateAVirtualSubAccountResponse) *NullableCreateAVirtualSubAccountResponse {
	return &NullableCreateAVirtualSubAccountResponse{value: val, isSet: true}
}

func (v NullableCreateAVirtualSubAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAVirtualSubAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
