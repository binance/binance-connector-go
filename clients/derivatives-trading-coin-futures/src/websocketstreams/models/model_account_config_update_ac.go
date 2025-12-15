/*
Binance Derivatives Trading COIN Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket Market Streams

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AccountConfigUpdateAc type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountConfigUpdateAc{}

// AccountConfigUpdateAc struct for AccountConfigUpdateAc
type AccountConfigUpdateAc struct {
	Smalls               *string `json:"s,omitempty"`
	Smalll               *int64  `json:"l,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountConfigUpdateAc AccountConfigUpdateAc

// NewAccountConfigUpdateAc instantiates a new AccountConfigUpdateAc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountConfigUpdateAc() *AccountConfigUpdateAc {
	this := AccountConfigUpdateAc{}
	return &this
}

// NewAccountConfigUpdateAcWithDefaults instantiates a new AccountConfigUpdateAc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountConfigUpdateAcWithDefaults() *AccountConfigUpdateAc {
	this := AccountConfigUpdateAc{}
	return &this
}

// GetS returns the S field value if set, zero value otherwise.
func (o *AccountConfigUpdateAc) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountConfigUpdateAc) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *AccountConfigUpdateAc) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *AccountConfigUpdateAc) SetSmalls(v string) {
	o.Smalls = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *AccountConfigUpdateAc) GetSmalll() int64 {
	if o == nil || common.IsNil(o.Smalll) {
		var ret int64
		return ret
	}
	return *o.Smalll
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountConfigUpdateAc) GetSmalllOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smalll) {
		return nil, false
	}
	return o.Smalll, true
}

// HasL returns a boolean if a field has been set.
func (o *AccountConfigUpdateAc) HasSmalll() bool {
	if o != nil && !common.IsNil(o.Smalll) {
		return true
	}

	return false
}

// SetL gets a reference to the given int64 and assigns it to the L field.
func (o *AccountConfigUpdateAc) SetSmalll(v int64) {
	o.Smalll = &v
}

func (o AccountConfigUpdateAc) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountConfigUpdateAc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smalll) {
		toSerialize["l"] = o.Smalll
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountConfigUpdateAc) UnmarshalJSON(data []byte) (err error) {
	varAccountConfigUpdateAc := _AccountConfigUpdateAc{}

	err = json.Unmarshal(data, &varAccountConfigUpdateAc)

	if err != nil {
		return err
	}

	*o = AccountConfigUpdateAc(varAccountConfigUpdateAc)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "s")
		delete(additionalProperties, "l")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountConfigUpdateAc struct {
	value *AccountConfigUpdateAc
	isSet bool
}

func (v NullableAccountConfigUpdateAc) Get() *AccountConfigUpdateAc {
	return v.value
}

func (v *NullableAccountConfigUpdateAc) Set(val *AccountConfigUpdateAc) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountConfigUpdateAc) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountConfigUpdateAc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountConfigUpdateAc(val *AccountConfigUpdateAc) *NullableAccountConfigUpdateAc {
	return &NullableAccountConfigUpdateAc{value: val, isSet: true}
}

func (v NullableAccountConfigUpdateAc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountConfigUpdateAc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
