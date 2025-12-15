/*
Binance Derivatives Trading Options WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Options WebSocket Market Streams

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AccountUpdatePInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountUpdatePInner{}

// AccountUpdatePInner struct for AccountUpdatePInner
type AccountUpdatePInner struct {
	Smalls               *string `json:"s,omitempty"`
	Smallc               *string `json:"c,omitempty"`
	Smallr               *string `json:"r,omitempty"`
	Smallp               *string `json:"p,omitempty"`
	Smalla               *string `json:"a,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountUpdatePInner AccountUpdatePInner

// NewAccountUpdatePInner instantiates a new AccountUpdatePInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountUpdatePInner() *AccountUpdatePInner {
	this := AccountUpdatePInner{}
	return &this
}

// NewAccountUpdatePInnerWithDefaults instantiates a new AccountUpdatePInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountUpdatePInnerWithDefaults() *AccountUpdatePInner {
	this := AccountUpdatePInner{}
	return &this
}

// GetS returns the S field value if set, zero value otherwise.
func (o *AccountUpdatePInner) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdatePInner) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *AccountUpdatePInner) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *AccountUpdatePInner) SetSmalls(v string) {
	o.Smalls = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *AccountUpdatePInner) GetSmallc() string {
	if o == nil || common.IsNil(o.Smallc) {
		var ret string
		return ret
	}
	return *o.Smallc
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdatePInner) GetSmallcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallc) {
		return nil, false
	}
	return o.Smallc, true
}

// HasC returns a boolean if a field has been set.
func (o *AccountUpdatePInner) HasSmallc() bool {
	if o != nil && !common.IsNil(o.Smallc) {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *AccountUpdatePInner) SetSmallc(v string) {
	o.Smallc = &v
}

// GetR returns the R field value if set, zero value otherwise.
func (o *AccountUpdatePInner) GetSmallr() string {
	if o == nil || common.IsNil(o.Smallr) {
		var ret string
		return ret
	}
	return *o.Smallr
}

// GetROk returns a tuple with the R field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdatePInner) GetSmallrOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallr) {
		return nil, false
	}
	return o.Smallr, true
}

// HasR returns a boolean if a field has been set.
func (o *AccountUpdatePInner) HasSmallr() bool {
	if o != nil && !common.IsNil(o.Smallr) {
		return true
	}

	return false
}

// SetR gets a reference to the given string and assigns it to the R field.
func (o *AccountUpdatePInner) SetSmallr(v string) {
	o.Smallr = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *AccountUpdatePInner) GetSmallp() string {
	if o == nil || common.IsNil(o.Smallp) {
		var ret string
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdatePInner) GetSmallpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *AccountUpdatePInner) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *AccountUpdatePInner) SetSmallp(v string) {
	o.Smallp = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *AccountUpdatePInner) GetSmalla() string {
	if o == nil || common.IsNil(o.Smalla) {
		var ret string
		return ret
	}
	return *o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdatePInner) GetSmallaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *AccountUpdatePInner) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *AccountUpdatePInner) SetSmalla(v string) {
	o.Smalla = &v
}

func (o AccountUpdatePInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountUpdatePInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smallc) {
		toSerialize["c"] = o.Smallc
	}
	if !common.IsNil(o.Smallr) {
		toSerialize["r"] = o.Smallr
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

func (o *AccountUpdatePInner) UnmarshalJSON(data []byte) (err error) {
	varAccountUpdatePInner := _AccountUpdatePInner{}

	err = json.Unmarshal(data, &varAccountUpdatePInner)

	if err != nil {
		return err
	}

	*o = AccountUpdatePInner(varAccountUpdatePInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "s")
		delete(additionalProperties, "c")
		delete(additionalProperties, "r")
		delete(additionalProperties, "p")
		delete(additionalProperties, "a")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountUpdatePInner struct {
	value *AccountUpdatePInner
	isSet bool
}

func (v NullableAccountUpdatePInner) Get() *AccountUpdatePInner {
	return v.value
}

func (v *NullableAccountUpdatePInner) Set(val *AccountUpdatePInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountUpdatePInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountUpdatePInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountUpdatePInner(val *AccountUpdatePInner) *NullableAccountUpdatePInner {
	return &NullableAccountUpdatePInner{value: val, isSet: true}
}

func (v NullableAccountUpdatePInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountUpdatePInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
