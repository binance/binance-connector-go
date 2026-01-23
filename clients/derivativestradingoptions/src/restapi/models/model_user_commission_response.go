/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the UserCommissionResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UserCommissionResponse{}

// UserCommissionResponse struct for UserCommissionResponse
type UserCommissionResponse struct {
	Commissions          []UserCommissionResponseCommissionsInner `json:"commissions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserCommissionResponse UserCommissionResponse

// NewUserCommissionResponse instantiates a new UserCommissionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCommissionResponse() *UserCommissionResponse {
	this := UserCommissionResponse{}
	return &this
}

// NewUserCommissionResponseWithDefaults instantiates a new UserCommissionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCommissionResponseWithDefaults() *UserCommissionResponse {
	this := UserCommissionResponse{}
	return &this
}

// GetCommissions returns the Commissions field value if set, zero value otherwise.
func (o *UserCommissionResponse) GetCommissions() []UserCommissionResponseCommissionsInner {
	if o == nil || common.IsNil(o.Commissions) {
		var ret []UserCommissionResponseCommissionsInner
		return ret
	}
	return o.Commissions
}

// GetCommissionsOk returns a tuple with the Commissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCommissionResponse) GetCommissionsOk() ([]UserCommissionResponseCommissionsInner, bool) {
	if o == nil || common.IsNil(o.Commissions) {
		return nil, false
	}
	return o.Commissions, true
}

// HasCommissions returns a boolean if a field has been set.
func (o *UserCommissionResponse) HasCommissions() bool {
	if o != nil && !common.IsNil(o.Commissions) {
		return true
	}

	return false
}

// SetCommissions gets a reference to the given []UserCommissionResponseCommissionsInner and assigns it to the Commissions field.
func (o *UserCommissionResponse) SetCommissions(v []UserCommissionResponseCommissionsInner) {
	o.Commissions = v
}

func (o UserCommissionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserCommissionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Commissions) {
		toSerialize["commissions"] = o.Commissions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserCommissionResponse) UnmarshalJSON(data []byte) (err error) {
	varUserCommissionResponse := _UserCommissionResponse{}

	err = json.Unmarshal(data, &varUserCommissionResponse)

	if err != nil {
		return err
	}

	*o = UserCommissionResponse(varUserCommissionResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "commissions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserCommissionResponse struct {
	value *UserCommissionResponse
	isSet bool
}

func (v NullableUserCommissionResponse) Get() *UserCommissionResponse {
	return v.value
}

func (v *NullableUserCommissionResponse) Set(val *UserCommissionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCommissionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCommissionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCommissionResponse(val *UserCommissionResponse) *NullableUserCommissionResponse {
	return &NullableUserCommissionResponse{value: val, isSet: true}
}

func (v NullableUserCommissionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCommissionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
