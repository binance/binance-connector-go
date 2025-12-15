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

// checks if the AccountConfigUpdate type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountConfigUpdate{}

// AccountConfigUpdate struct for AccountConfigUpdate
type AccountConfigUpdate struct {
	E                    *int64                 `json:"E,omitempty"`
	T                    *int64                 `json:"T,omitempty"`
	Smallac              *AccountConfigUpdateAc `json:"ac,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountConfigUpdate AccountConfigUpdate

// NewAccountConfigUpdate instantiates a new AccountConfigUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountConfigUpdate() *AccountConfigUpdate {
	this := AccountConfigUpdate{}
	return &this
}

// NewAccountConfigUpdateWithDefaults instantiates a new AccountConfigUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountConfigUpdateWithDefaults() *AccountConfigUpdate {
	this := AccountConfigUpdate{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AccountConfigUpdate) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountConfigUpdate) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *AccountConfigUpdate) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *AccountConfigUpdate) SetE(v int64) {
	o.E = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *AccountConfigUpdate) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountConfigUpdate) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *AccountConfigUpdate) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *AccountConfigUpdate) SetT(v int64) {
	o.T = &v
}

// GetAc returns the Ac field value if set, zero value otherwise.
func (o *AccountConfigUpdate) GetSmallac() AccountConfigUpdateAc {
	if o == nil || common.IsNil(o.Smallac) {
		var ret AccountConfigUpdateAc
		return ret
	}
	return *o.Smallac
}

// GetAcOk returns a tuple with the Ac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountConfigUpdate) GetSmallacOk() (*AccountConfigUpdateAc, bool) {
	if o == nil || common.IsNil(o.Smallac) {
		return nil, false
	}
	return o.Smallac, true
}

// HasAc returns a boolean if a field has been set.
func (o *AccountConfigUpdate) HasSmallac() bool {
	if o != nil && !common.IsNil(o.Smallac) {
		return true
	}

	return false
}

// SetAc gets a reference to the given AccountConfigUpdateAc and assigns it to the Ac field.
func (o *AccountConfigUpdate) SetSmallac(v AccountConfigUpdateAc) {
	o.Smallac = &v
}

func (o AccountConfigUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountConfigUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.Smallac) {
		toSerialize["ac"] = o.Smallac
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountConfigUpdate) UnmarshalJSON(data []byte) (err error) {
	varAccountConfigUpdate := _AccountConfigUpdate{}

	err = json.Unmarshal(data, &varAccountConfigUpdate)

	if err != nil {
		return err
	}

	*o = AccountConfigUpdate(varAccountConfigUpdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "E")
		delete(additionalProperties, "T")
		delete(additionalProperties, "ac")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountConfigUpdate struct {
	value *AccountConfigUpdate
	isSet bool
}

func (v NullableAccountConfigUpdate) Get() *AccountConfigUpdate {
	return v.value
}

func (v *NullableAccountConfigUpdate) Set(val *AccountConfigUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountConfigUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountConfigUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountConfigUpdate(val *AccountConfigUpdate) *NullableAccountConfigUpdate {
	return &NullableAccountConfigUpdate{value: val, isSet: true}
}

func (v NullableAccountConfigUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountConfigUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
