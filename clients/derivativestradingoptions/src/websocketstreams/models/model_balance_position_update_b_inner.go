/*
Binance Derivatives Trading Options WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Options WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the BalancePositionUpdateBInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BalancePositionUpdateBInner{}

// BalancePositionUpdateBInner struct for BalancePositionUpdateBInner
type BalancePositionUpdateBInner struct {
	Smalla               *string `json:"a,omitempty"`
	Smallb               *string `json:"b,omitempty"`
	Smallbc              *string `json:"bc,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BalancePositionUpdateBInner BalancePositionUpdateBInner

// NewBalancePositionUpdateBInner instantiates a new BalancePositionUpdateBInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalancePositionUpdateBInner() *BalancePositionUpdateBInner {
	this := BalancePositionUpdateBInner{}
	return &this
}

// NewBalancePositionUpdateBInnerWithDefaults instantiates a new BalancePositionUpdateBInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalancePositionUpdateBInnerWithDefaults() *BalancePositionUpdateBInner {
	this := BalancePositionUpdateBInner{}
	return &this
}

// GetA returns the A field value if set, zero value otherwise.
func (o *BalancePositionUpdateBInner) GetSmalla() string {
	if o == nil || common.IsNil(o.Smalla) {
		var ret string
		return ret
	}
	return *o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalancePositionUpdateBInner) GetSmallaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *BalancePositionUpdateBInner) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *BalancePositionUpdateBInner) SetSmalla(v string) {
	o.Smalla = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *BalancePositionUpdateBInner) GetSmallb() string {
	if o == nil || common.IsNil(o.Smallb) {
		var ret string
		return ret
	}
	return *o.Smallb
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalancePositionUpdateBInner) GetSmallbOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallb) {
		return nil, false
	}
	return o.Smallb, true
}

// HasB returns a boolean if a field has been set.
func (o *BalancePositionUpdateBInner) HasSmallb() bool {
	if o != nil && !common.IsNil(o.Smallb) {
		return true
	}

	return false
}

// SetB gets a reference to the given string and assigns it to the B field.
func (o *BalancePositionUpdateBInner) SetSmallb(v string) {
	o.Smallb = &v
}

// GetBc returns the Bc field value if set, zero value otherwise.
func (o *BalancePositionUpdateBInner) GetSmallbc() string {
	if o == nil || common.IsNil(o.Smallbc) {
		var ret string
		return ret
	}
	return *o.Smallbc
}

// GetBcOk returns a tuple with the Bc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalancePositionUpdateBInner) GetSmallbcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallbc) {
		return nil, false
	}
	return o.Smallbc, true
}

// HasBc returns a boolean if a field has been set.
func (o *BalancePositionUpdateBInner) HasSmallbc() bool {
	if o != nil && !common.IsNil(o.Smallbc) {
		return true
	}

	return false
}

// SetBc gets a reference to the given string and assigns it to the Bc field.
func (o *BalancePositionUpdateBInner) SetSmallbc(v string) {
	o.Smallbc = &v
}

func (o BalancePositionUpdateBInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BalancePositionUpdateBInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalla) {
		toSerialize["a"] = o.Smalla
	}
	if !common.IsNil(o.Smallb) {
		toSerialize["b"] = o.Smallb
	}
	if !common.IsNil(o.Smallbc) {
		toSerialize["bc"] = o.Smallbc
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BalancePositionUpdateBInner) UnmarshalJSON(data []byte) (err error) {
	varBalancePositionUpdateBInner := _BalancePositionUpdateBInner{}

	err = json.Unmarshal(data, &varBalancePositionUpdateBInner)

	if err != nil {
		return err
	}

	*o = BalancePositionUpdateBInner(varBalancePositionUpdateBInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "a")
		delete(additionalProperties, "b")
		delete(additionalProperties, "bc")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBalancePositionUpdateBInner struct {
	value *BalancePositionUpdateBInner
	isSet bool
}

func (v NullableBalancePositionUpdateBInner) Get() *BalancePositionUpdateBInner {
	return v.value
}

func (v *NullableBalancePositionUpdateBInner) Set(val *BalancePositionUpdateBInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBalancePositionUpdateBInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBalancePositionUpdateBInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalancePositionUpdateBInner(val *BalancePositionUpdateBInner) *NullableBalancePositionUpdateBInner {
	return &NullableBalancePositionUpdateBInner{value: val, isSet: true}
}

func (v NullableBalancePositionUpdateBInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalancePositionUpdateBInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
