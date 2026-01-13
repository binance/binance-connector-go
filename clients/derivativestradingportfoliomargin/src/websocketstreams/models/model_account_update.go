/*
Binance Derivatives Trading Portfolio Margin WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AccountUpdate type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountUpdate{}

// AccountUpdate struct for AccountUpdate
type AccountUpdate struct {
	Smallfs              *string         `json:"fs,omitempty"`
	E                    *int64          `json:"E,omitempty"`
	T                    *int64          `json:"T,omitempty"`
	Smalli               *string         `json:"i,omitempty"`
	Smalla               *AccountUpdateA `json:"a,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountUpdate AccountUpdate

// NewAccountUpdate instantiates a new AccountUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountUpdate() *AccountUpdate {
	this := AccountUpdate{}
	return &this
}

// NewAccountUpdateWithDefaults instantiates a new AccountUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountUpdateWithDefaults() *AccountUpdate {
	this := AccountUpdate{}
	return &this
}

// GetFs returns the Fs field value if set, zero value otherwise.
func (o *AccountUpdate) GetSmallfs() string {
	if o == nil || common.IsNil(o.Smallfs) {
		var ret string
		return ret
	}
	return *o.Smallfs
}

// GetFsOk returns a tuple with the Fs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdate) GetSmallfsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallfs) {
		return nil, false
	}
	return o.Smallfs, true
}

// HasFs returns a boolean if a field has been set.
func (o *AccountUpdate) HasSmallfs() bool {
	if o != nil && !common.IsNil(o.Smallfs) {
		return true
	}

	return false
}

// SetFs gets a reference to the given string and assigns it to the Fs field.
func (o *AccountUpdate) SetSmallfs(v string) {
	o.Smallfs = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AccountUpdate) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdate) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *AccountUpdate) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *AccountUpdate) SetE(v int64) {
	o.E = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *AccountUpdate) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdate) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *AccountUpdate) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *AccountUpdate) SetT(v int64) {
	o.T = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *AccountUpdate) GetSmalli() string {
	if o == nil || common.IsNil(o.Smalli) {
		var ret string
		return ret
	}
	return *o.Smalli
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdate) GetSmalliOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalli) {
		return nil, false
	}
	return o.Smalli, true
}

// HasI returns a boolean if a field has been set.
func (o *AccountUpdate) HasSmalli() bool {
	if o != nil && !common.IsNil(o.Smalli) {
		return true
	}

	return false
}

// SetI gets a reference to the given string and assigns it to the I field.
func (o *AccountUpdate) SetSmalli(v string) {
	o.Smalli = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *AccountUpdate) GetSmalla() AccountUpdateA {
	if o == nil || common.IsNil(o.Smalla) {
		var ret AccountUpdateA
		return ret
	}
	return *o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdate) GetSmallaOk() (*AccountUpdateA, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *AccountUpdate) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given AccountUpdateA and assigns it to the A field.
func (o *AccountUpdate) SetSmalla(v AccountUpdateA) {
	o.Smalla = &v
}

func (o AccountUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smallfs) {
		toSerialize["fs"] = o.Smallfs
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.Smalli) {
		toSerialize["i"] = o.Smalli
	}
	if !common.IsNil(o.Smalla) {
		toSerialize["a"] = o.Smalla
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountUpdate) UnmarshalJSON(data []byte) (err error) {
	varAccountUpdate := _AccountUpdate{}

	err = json.Unmarshal(data, &varAccountUpdate)

	if err != nil {
		return err
	}

	*o = AccountUpdate(varAccountUpdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fs")
		delete(additionalProperties, "E")
		delete(additionalProperties, "T")
		delete(additionalProperties, "i")
		delete(additionalProperties, "a")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountUpdate struct {
	value *AccountUpdate
	isSet bool
}

func (v NullableAccountUpdate) Get() *AccountUpdate {
	return v.value
}

func (v *NullableAccountUpdate) Set(val *AccountUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountUpdate(val *AccountUpdate) *NullableAccountUpdate {
	return &NullableAccountUpdate{value: val, isSet: true}
}

func (v NullableAccountUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
