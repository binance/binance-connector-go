/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the EnableFuturesForSubAccountResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &EnableFuturesForSubAccountResponse{}

// EnableFuturesForSubAccountResponse struct for EnableFuturesForSubAccountResponse
type EnableFuturesForSubAccountResponse struct {
	Email                *string `json:"email,omitempty"`
	IsFuturesEnabled     *bool   `json:"isFuturesEnabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EnableFuturesForSubAccountResponse EnableFuturesForSubAccountResponse

// NewEnableFuturesForSubAccountResponse instantiates a new EnableFuturesForSubAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnableFuturesForSubAccountResponse() *EnableFuturesForSubAccountResponse {
	this := EnableFuturesForSubAccountResponse{}
	return &this
}

// NewEnableFuturesForSubAccountResponseWithDefaults instantiates a new EnableFuturesForSubAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnableFuturesForSubAccountResponseWithDefaults() *EnableFuturesForSubAccountResponse {
	this := EnableFuturesForSubAccountResponse{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *EnableFuturesForSubAccountResponse) GetEmail() string {
	if o == nil || common.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableFuturesForSubAccountResponse) GetEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *EnableFuturesForSubAccountResponse) HasEmail() bool {
	if o != nil && !common.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *EnableFuturesForSubAccountResponse) SetEmail(v string) {
	o.Email = &v
}

// GetIsFuturesEnabled returns the IsFuturesEnabled field value if set, zero value otherwise.
func (o *EnableFuturesForSubAccountResponse) GetIsFuturesEnabled() bool {
	if o == nil || common.IsNil(o.IsFuturesEnabled) {
		var ret bool
		return ret
	}
	return *o.IsFuturesEnabled
}

// GetIsFuturesEnabledOk returns a tuple with the IsFuturesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableFuturesForSubAccountResponse) GetIsFuturesEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsFuturesEnabled) {
		return nil, false
	}
	return o.IsFuturesEnabled, true
}

// HasIsFuturesEnabled returns a boolean if a field has been set.
func (o *EnableFuturesForSubAccountResponse) HasIsFuturesEnabled() bool {
	if o != nil && !common.IsNil(o.IsFuturesEnabled) {
		return true
	}

	return false
}

// SetIsFuturesEnabled gets a reference to the given bool and assigns it to the IsFuturesEnabled field.
func (o *EnableFuturesForSubAccountResponse) SetIsFuturesEnabled(v bool) {
	o.IsFuturesEnabled = &v
}

func (o EnableFuturesForSubAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnableFuturesForSubAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !common.IsNil(o.IsFuturesEnabled) {
		toSerialize["isFuturesEnabled"] = o.IsFuturesEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EnableFuturesForSubAccountResponse) UnmarshalJSON(data []byte) (err error) {
	varEnableFuturesForSubAccountResponse := _EnableFuturesForSubAccountResponse{}

	err = json.Unmarshal(data, &varEnableFuturesForSubAccountResponse)

	if err != nil {
		return err
	}

	*o = EnableFuturesForSubAccountResponse(varEnableFuturesForSubAccountResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "isFuturesEnabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEnableFuturesForSubAccountResponse struct {
	value *EnableFuturesForSubAccountResponse
	isSet bool
}

func (v NullableEnableFuturesForSubAccountResponse) Get() *EnableFuturesForSubAccountResponse {
	return v.value
}

func (v *NullableEnableFuturesForSubAccountResponse) Set(val *EnableFuturesForSubAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEnableFuturesForSubAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEnableFuturesForSubAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnableFuturesForSubAccountResponse(val *EnableFuturesForSubAccountResponse) *NullableEnableFuturesForSubAccountResponse {
	return &NullableEnableFuturesForSubAccountResponse{value: val, isSet: true}
}

func (v NullableEnableFuturesForSubAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnableFuturesForSubAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
