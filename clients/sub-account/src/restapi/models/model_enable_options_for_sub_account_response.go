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

// checks if the EnableOptionsForSubAccountResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &EnableOptionsForSubAccountResponse{}

// EnableOptionsForSubAccountResponse struct for EnableOptionsForSubAccountResponse
type EnableOptionsForSubAccountResponse struct {
	Email                *string `json:"email,omitempty"`
	IsEOptionsEnabled    *bool   `json:"isEOptionsEnabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EnableOptionsForSubAccountResponse EnableOptionsForSubAccountResponse

// NewEnableOptionsForSubAccountResponse instantiates a new EnableOptionsForSubAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnableOptionsForSubAccountResponse() *EnableOptionsForSubAccountResponse {
	this := EnableOptionsForSubAccountResponse{}
	return &this
}

// NewEnableOptionsForSubAccountResponseWithDefaults instantiates a new EnableOptionsForSubAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnableOptionsForSubAccountResponseWithDefaults() *EnableOptionsForSubAccountResponse {
	this := EnableOptionsForSubAccountResponse{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *EnableOptionsForSubAccountResponse) GetEmail() string {
	if o == nil || common.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableOptionsForSubAccountResponse) GetEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *EnableOptionsForSubAccountResponse) HasEmail() bool {
	if o != nil && !common.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *EnableOptionsForSubAccountResponse) SetEmail(v string) {
	o.Email = &v
}

// GetIsEOptionsEnabled returns the IsEOptionsEnabled field value if set, zero value otherwise.
func (o *EnableOptionsForSubAccountResponse) GetIsEOptionsEnabled() bool {
	if o == nil || common.IsNil(o.IsEOptionsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEOptionsEnabled
}

// GetIsEOptionsEnabledOk returns a tuple with the IsEOptionsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableOptionsForSubAccountResponse) GetIsEOptionsEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsEOptionsEnabled) {
		return nil, false
	}
	return o.IsEOptionsEnabled, true
}

// HasIsEOptionsEnabled returns a boolean if a field has been set.
func (o *EnableOptionsForSubAccountResponse) HasIsEOptionsEnabled() bool {
	if o != nil && !common.IsNil(o.IsEOptionsEnabled) {
		return true
	}

	return false
}

// SetIsEOptionsEnabled gets a reference to the given bool and assigns it to the IsEOptionsEnabled field.
func (o *EnableOptionsForSubAccountResponse) SetIsEOptionsEnabled(v bool) {
	o.IsEOptionsEnabled = &v
}

func (o EnableOptionsForSubAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnableOptionsForSubAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !common.IsNil(o.IsEOptionsEnabled) {
		toSerialize["isEOptionsEnabled"] = o.IsEOptionsEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EnableOptionsForSubAccountResponse) UnmarshalJSON(data []byte) (err error) {
	varEnableOptionsForSubAccountResponse := _EnableOptionsForSubAccountResponse{}

	err = json.Unmarshal(data, &varEnableOptionsForSubAccountResponse)

	if err != nil {
		return err
	}

	*o = EnableOptionsForSubAccountResponse(varEnableOptionsForSubAccountResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "isEOptionsEnabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEnableOptionsForSubAccountResponse struct {
	value *EnableOptionsForSubAccountResponse
	isSet bool
}

func (v NullableEnableOptionsForSubAccountResponse) Get() *EnableOptionsForSubAccountResponse {
	return v.value
}

func (v *NullableEnableOptionsForSubAccountResponse) Set(val *EnableOptionsForSubAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEnableOptionsForSubAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEnableOptionsForSubAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnableOptionsForSubAccountResponse(val *EnableOptionsForSubAccountResponse) *NullableEnableOptionsForSubAccountResponse {
	return &NullableEnableOptionsForSubAccountResponse{value: val, isSet: true}
}

func (v NullableEnableOptionsForSubAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnableOptionsForSubAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
