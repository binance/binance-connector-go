/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CreateSpecialKeyResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CreateSpecialKeyResponse{}

// CreateSpecialKeyResponse struct for CreateSpecialKeyResponse
type CreateSpecialKeyResponse struct {
	ApiKey               *string `json:"apiKey,omitempty"`
	SecretKey            *string `json:"secretKey,omitempty"`
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateSpecialKeyResponse CreateSpecialKeyResponse

// NewCreateSpecialKeyResponse instantiates a new CreateSpecialKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSpecialKeyResponse() *CreateSpecialKeyResponse {
	this := CreateSpecialKeyResponse{}
	return &this
}

// NewCreateSpecialKeyResponseWithDefaults instantiates a new CreateSpecialKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSpecialKeyResponseWithDefaults() *CreateSpecialKeyResponse {
	this := CreateSpecialKeyResponse{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *CreateSpecialKeyResponse) GetApiKey() string {
	if o == nil || common.IsNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSpecialKeyResponse) GetApiKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *CreateSpecialKeyResponse) HasApiKey() bool {
	if o != nil && !common.IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *CreateSpecialKeyResponse) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *CreateSpecialKeyResponse) GetSecretKey() string {
	if o == nil || common.IsNil(o.SecretKey) {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSpecialKeyResponse) GetSecretKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.SecretKey) {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *CreateSpecialKeyResponse) HasSecretKey() bool {
	if o != nil && !common.IsNil(o.SecretKey) {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *CreateSpecialKeyResponse) SetSecretKey(v string) {
	o.SecretKey = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateSpecialKeyResponse) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSpecialKeyResponse) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateSpecialKeyResponse) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateSpecialKeyResponse) SetType(v string) {
	o.Type = &v
}

func (o CreateSpecialKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSpecialKeyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ApiKey) {
		toSerialize["apiKey"] = o.ApiKey
	}
	if !common.IsNil(o.SecretKey) {
		toSerialize["secretKey"] = o.SecretKey
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateSpecialKeyResponse) UnmarshalJSON(data []byte) (err error) {
	varCreateSpecialKeyResponse := _CreateSpecialKeyResponse{}

	err = json.Unmarshal(data, &varCreateSpecialKeyResponse)

	if err != nil {
		return err
	}

	*o = CreateSpecialKeyResponse(varCreateSpecialKeyResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "apiKey")
		delete(additionalProperties, "secretKey")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateSpecialKeyResponse struct {
	value *CreateSpecialKeyResponse
	isSet bool
}

func (v NullableCreateSpecialKeyResponse) Get() *CreateSpecialKeyResponse {
	return v.value
}

func (v *NullableCreateSpecialKeyResponse) Set(val *CreateSpecialKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSpecialKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSpecialKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSpecialKeyResponse(val *CreateSpecialKeyResponse) *NullableCreateSpecialKeyResponse {
	return &NullableCreateSpecialKeyResponse{value: val, isSet: true}
}

func (v NullableCreateSpecialKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSpecialKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
