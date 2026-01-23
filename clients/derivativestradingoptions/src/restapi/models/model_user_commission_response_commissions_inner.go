/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the UserCommissionResponseCommissionsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UserCommissionResponseCommissionsInner{}

// UserCommissionResponseCommissionsInner struct for UserCommissionResponseCommissionsInner
type UserCommissionResponseCommissionsInner struct {
	Underlying           *string `json:"underlying,omitempty"`
	MakerFee             *string `json:"makerFee,omitempty"`
	TakerFee             *string `json:"takerFee,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserCommissionResponseCommissionsInner UserCommissionResponseCommissionsInner

// NewUserCommissionResponseCommissionsInner instantiates a new UserCommissionResponseCommissionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCommissionResponseCommissionsInner() *UserCommissionResponseCommissionsInner {
	this := UserCommissionResponseCommissionsInner{}
	return &this
}

// NewUserCommissionResponseCommissionsInnerWithDefaults instantiates a new UserCommissionResponseCommissionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCommissionResponseCommissionsInnerWithDefaults() *UserCommissionResponseCommissionsInner {
	this := UserCommissionResponseCommissionsInner{}
	return &this
}

// GetUnderlying returns the Underlying field value if set, zero value otherwise.
func (o *UserCommissionResponseCommissionsInner) GetUnderlying() string {
	if o == nil || common.IsNil(o.Underlying) {
		var ret string
		return ret
	}
	return *o.Underlying
}

// GetUnderlyingOk returns a tuple with the Underlying field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCommissionResponseCommissionsInner) GetUnderlyingOk() (*string, bool) {
	if o == nil || common.IsNil(o.Underlying) {
		return nil, false
	}
	return o.Underlying, true
}

// HasUnderlying returns a boolean if a field has been set.
func (o *UserCommissionResponseCommissionsInner) HasUnderlying() bool {
	if o != nil && !common.IsNil(o.Underlying) {
		return true
	}

	return false
}

// SetUnderlying gets a reference to the given string and assigns it to the Underlying field.
func (o *UserCommissionResponseCommissionsInner) SetUnderlying(v string) {
	o.Underlying = &v
}

// GetMakerFee returns the MakerFee field value if set, zero value otherwise.
func (o *UserCommissionResponseCommissionsInner) GetMakerFee() string {
	if o == nil || common.IsNil(o.MakerFee) {
		var ret string
		return ret
	}
	return *o.MakerFee
}

// GetMakerFeeOk returns a tuple with the MakerFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCommissionResponseCommissionsInner) GetMakerFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.MakerFee) {
		return nil, false
	}
	return o.MakerFee, true
}

// HasMakerFee returns a boolean if a field has been set.
func (o *UserCommissionResponseCommissionsInner) HasMakerFee() bool {
	if o != nil && !common.IsNil(o.MakerFee) {
		return true
	}

	return false
}

// SetMakerFee gets a reference to the given string and assigns it to the MakerFee field.
func (o *UserCommissionResponseCommissionsInner) SetMakerFee(v string) {
	o.MakerFee = &v
}

// GetTakerFee returns the TakerFee field value if set, zero value otherwise.
func (o *UserCommissionResponseCommissionsInner) GetTakerFee() string {
	if o == nil || common.IsNil(o.TakerFee) {
		var ret string
		return ret
	}
	return *o.TakerFee
}

// GetTakerFeeOk returns a tuple with the TakerFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCommissionResponseCommissionsInner) GetTakerFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.TakerFee) {
		return nil, false
	}
	return o.TakerFee, true
}

// HasTakerFee returns a boolean if a field has been set.
func (o *UserCommissionResponseCommissionsInner) HasTakerFee() bool {
	if o != nil && !common.IsNil(o.TakerFee) {
		return true
	}

	return false
}

// SetTakerFee gets a reference to the given string and assigns it to the TakerFee field.
func (o *UserCommissionResponseCommissionsInner) SetTakerFee(v string) {
	o.TakerFee = &v
}

func (o UserCommissionResponseCommissionsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserCommissionResponseCommissionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Underlying) {
		toSerialize["underlying"] = o.Underlying
	}
	if !common.IsNil(o.MakerFee) {
		toSerialize["makerFee"] = o.MakerFee
	}
	if !common.IsNil(o.TakerFee) {
		toSerialize["takerFee"] = o.TakerFee
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserCommissionResponseCommissionsInner) UnmarshalJSON(data []byte) (err error) {
	varUserCommissionResponseCommissionsInner := _UserCommissionResponseCommissionsInner{}

	err = json.Unmarshal(data, &varUserCommissionResponseCommissionsInner)

	if err != nil {
		return err
	}

	*o = UserCommissionResponseCommissionsInner(varUserCommissionResponseCommissionsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "underlying")
		delete(additionalProperties, "makerFee")
		delete(additionalProperties, "takerFee")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserCommissionResponseCommissionsInner struct {
	value *UserCommissionResponseCommissionsInner
	isSet bool
}

func (v NullableUserCommissionResponseCommissionsInner) Get() *UserCommissionResponseCommissionsInner {
	return v.value
}

func (v *NullableUserCommissionResponseCommissionsInner) Set(val *UserCommissionResponseCommissionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCommissionResponseCommissionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCommissionResponseCommissionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCommissionResponseCommissionsInner(val *UserCommissionResponseCommissionsInner) *NullableUserCommissionResponseCommissionsInner {
	return &NullableUserCommissionResponseCommissionsInner{value: val, isSet: true}
}

func (v NullableUserCommissionResponseCommissionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCommissionResponseCommissionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
