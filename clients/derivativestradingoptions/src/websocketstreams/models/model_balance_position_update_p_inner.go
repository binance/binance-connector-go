/*
Binance Derivatives Trading Options WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Options WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the BalancePositionUpdatePInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BalancePositionUpdatePInner{}

// BalancePositionUpdatePInner struct for BalancePositionUpdatePInner
type BalancePositionUpdatePInner struct {
	Smalls               *string `json:"s,omitempty"`
	Smallc               *string `json:"c,omitempty"`
	Smallp               *string `json:"p,omitempty"`
	Smalla               *string `json:"a,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BalancePositionUpdatePInner BalancePositionUpdatePInner

// NewBalancePositionUpdatePInner instantiates a new BalancePositionUpdatePInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalancePositionUpdatePInner() *BalancePositionUpdatePInner {
	this := BalancePositionUpdatePInner{}
	return &this
}

// NewBalancePositionUpdatePInnerWithDefaults instantiates a new BalancePositionUpdatePInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalancePositionUpdatePInnerWithDefaults() *BalancePositionUpdatePInner {
	this := BalancePositionUpdatePInner{}
	return &this
}

// GetS returns the S field value if set, zero value otherwise.
func (o *BalancePositionUpdatePInner) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalancePositionUpdatePInner) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *BalancePositionUpdatePInner) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *BalancePositionUpdatePInner) SetSmalls(v string) {
	o.Smalls = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *BalancePositionUpdatePInner) GetSmallc() string {
	if o == nil || common.IsNil(o.Smallc) {
		var ret string
		return ret
	}
	return *o.Smallc
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalancePositionUpdatePInner) GetSmallcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallc) {
		return nil, false
	}
	return o.Smallc, true
}

// HasC returns a boolean if a field has been set.
func (o *BalancePositionUpdatePInner) HasSmallc() bool {
	if o != nil && !common.IsNil(o.Smallc) {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *BalancePositionUpdatePInner) SetSmallc(v string) {
	o.Smallc = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *BalancePositionUpdatePInner) GetSmallp() string {
	if o == nil || common.IsNil(o.Smallp) {
		var ret string
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalancePositionUpdatePInner) GetSmallpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *BalancePositionUpdatePInner) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *BalancePositionUpdatePInner) SetSmallp(v string) {
	o.Smallp = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *BalancePositionUpdatePInner) GetSmalla() string {
	if o == nil || common.IsNil(o.Smalla) {
		var ret string
		return ret
	}
	return *o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalancePositionUpdatePInner) GetSmallaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *BalancePositionUpdatePInner) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *BalancePositionUpdatePInner) SetSmalla(v string) {
	o.Smalla = &v
}

func (o BalancePositionUpdatePInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BalancePositionUpdatePInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smallc) {
		toSerialize["c"] = o.Smallc
	}
	if !common.IsNil(o.Smallp) {
		toSerialize["p"] = o.Smallp
	}
	if !common.IsNil(o.Smalla) {
		toSerialize["a"] = o.Smalla
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BalancePositionUpdatePInner) UnmarshalJSON(data []byte) (err error) {
	varBalancePositionUpdatePInner := _BalancePositionUpdatePInner{}

	err = json.Unmarshal(data, &varBalancePositionUpdatePInner)

	if err != nil {
		return err
	}

	*o = BalancePositionUpdatePInner(varBalancePositionUpdatePInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "s")
		delete(additionalProperties, "c")
		delete(additionalProperties, "p")
		delete(additionalProperties, "a")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBalancePositionUpdatePInner struct {
	value *BalancePositionUpdatePInner
	isSet bool
}

func (v NullableBalancePositionUpdatePInner) Get() *BalancePositionUpdatePInner {
	return v.value
}

func (v *NullableBalancePositionUpdatePInner) Set(val *BalancePositionUpdatePInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBalancePositionUpdatePInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBalancePositionUpdatePInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalancePositionUpdatePInner(val *BalancePositionUpdatePInner) *NullableBalancePositionUpdatePInner {
	return &NullableBalancePositionUpdatePInner{value: val, isSet: true}
}

func (v NullableBalancePositionUpdatePInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalancePositionUpdatePInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
