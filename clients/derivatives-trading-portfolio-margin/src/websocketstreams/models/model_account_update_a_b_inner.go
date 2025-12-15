/*
Binance Derivatives Trading Portfolio Margin WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin WebSocket Market Streams

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AccountUpdateABInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountUpdateABInner{}

// AccountUpdateABInner struct for AccountUpdateABInner
type AccountUpdateABInner struct {
	Smalla               *string `json:"a,omitempty"`
	Smallwb              *string `json:"wb,omitempty"`
	Smallcw              *string `json:"cw,omitempty"`
	Smallbc              *string `json:"bc,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountUpdateABInner AccountUpdateABInner

// NewAccountUpdateABInner instantiates a new AccountUpdateABInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountUpdateABInner() *AccountUpdateABInner {
	this := AccountUpdateABInner{}
	return &this
}

// NewAccountUpdateABInnerWithDefaults instantiates a new AccountUpdateABInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountUpdateABInnerWithDefaults() *AccountUpdateABInner {
	this := AccountUpdateABInner{}
	return &this
}

// GetA returns the A field value if set, zero value otherwise.
func (o *AccountUpdateABInner) GetSmalla() string {
	if o == nil || common.IsNil(o.Smalla) {
		var ret string
		return ret
	}
	return *o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateABInner) GetSmallaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *AccountUpdateABInner) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *AccountUpdateABInner) SetSmalla(v string) {
	o.Smalla = &v
}

// GetWb returns the Wb field value if set, zero value otherwise.
func (o *AccountUpdateABInner) GetSmallwb() string {
	if o == nil || common.IsNil(o.Smallwb) {
		var ret string
		return ret
	}
	return *o.Smallwb
}

// GetWbOk returns a tuple with the Wb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateABInner) GetSmallwbOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallwb) {
		return nil, false
	}
	return o.Smallwb, true
}

// HasWb returns a boolean if a field has been set.
func (o *AccountUpdateABInner) HasSmallwb() bool {
	if o != nil && !common.IsNil(o.Smallwb) {
		return true
	}

	return false
}

// SetWb gets a reference to the given string and assigns it to the Wb field.
func (o *AccountUpdateABInner) SetSmallwb(v string) {
	o.Smallwb = &v
}

// GetCw returns the Cw field value if set, zero value otherwise.
func (o *AccountUpdateABInner) GetSmallcw() string {
	if o == nil || common.IsNil(o.Smallcw) {
		var ret string
		return ret
	}
	return *o.Smallcw
}

// GetCwOk returns a tuple with the Cw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateABInner) GetSmallcwOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallcw) {
		return nil, false
	}
	return o.Smallcw, true
}

// HasCw returns a boolean if a field has been set.
func (o *AccountUpdateABInner) HasSmallcw() bool {
	if o != nil && !common.IsNil(o.Smallcw) {
		return true
	}

	return false
}

// SetCw gets a reference to the given string and assigns it to the Cw field.
func (o *AccountUpdateABInner) SetSmallcw(v string) {
	o.Smallcw = &v
}

// GetBc returns the Bc field value if set, zero value otherwise.
func (o *AccountUpdateABInner) GetSmallbc() string {
	if o == nil || common.IsNil(o.Smallbc) {
		var ret string
		return ret
	}
	return *o.Smallbc
}

// GetBcOk returns a tuple with the Bc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateABInner) GetSmallbcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallbc) {
		return nil, false
	}
	return o.Smallbc, true
}

// HasBc returns a boolean if a field has been set.
func (o *AccountUpdateABInner) HasSmallbc() bool {
	if o != nil && !common.IsNil(o.Smallbc) {
		return true
	}

	return false
}

// SetBc gets a reference to the given string and assigns it to the Bc field.
func (o *AccountUpdateABInner) SetSmallbc(v string) {
	o.Smallbc = &v
}

func (o AccountUpdateABInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountUpdateABInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalla) {
		toSerialize["a"] = o.Smalla
	}
	if !common.IsNil(o.Smallwb) {
		toSerialize["wb"] = o.Smallwb
	}
	if !common.IsNil(o.Smallcw) {
		toSerialize["cw"] = o.Smallcw
	}
	if !common.IsNil(o.Smallbc) {
		toSerialize["bc"] = o.Smallbc
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountUpdateABInner) UnmarshalJSON(data []byte) (err error) {
	varAccountUpdateABInner := _AccountUpdateABInner{}

	err = json.Unmarshal(data, &varAccountUpdateABInner)

	if err != nil {
		return err
	}

	*o = AccountUpdateABInner(varAccountUpdateABInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "a")
		delete(additionalProperties, "wb")
		delete(additionalProperties, "cw")
		delete(additionalProperties, "bc")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountUpdateABInner struct {
	value *AccountUpdateABInner
	isSet bool
}

func (v NullableAccountUpdateABInner) Get() *AccountUpdateABInner {
	return v.value
}

func (v *NullableAccountUpdateABInner) Set(val *AccountUpdateABInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountUpdateABInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountUpdateABInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountUpdateABInner(val *AccountUpdateABInner) *NullableAccountUpdateABInner {
	return &NullableAccountUpdateABInner{value: val, isSet: true}
}

func (v NullableAccountUpdateABInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountUpdateABInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
