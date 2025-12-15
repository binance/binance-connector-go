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

// checks if the AccountUpdateBInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountUpdateBInner{}

// AccountUpdateBInner struct for AccountUpdateBInner
type AccountUpdateBInner struct {
	Smallb               *string `json:"b,omitempty"`
	Smallm               *string `json:"m,omitempty"`
	Smallu               *string `json:"u,omitempty"`
	U                    *int64  `json:"U,omitempty"`
	M                    *string `json:"M,omitempty"`
	Smalli               *string `json:"i,omitempty"`
	Smalla               *string `json:"a,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountUpdateBInner AccountUpdateBInner

// NewAccountUpdateBInner instantiates a new AccountUpdateBInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountUpdateBInner() *AccountUpdateBInner {
	this := AccountUpdateBInner{}
	return &this
}

// NewAccountUpdateBInnerWithDefaults instantiates a new AccountUpdateBInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountUpdateBInnerWithDefaults() *AccountUpdateBInner {
	this := AccountUpdateBInner{}
	return &this
}

// GetB returns the B field value if set, zero value otherwise.
func (o *AccountUpdateBInner) GetSmallb() string {
	if o == nil || common.IsNil(o.Smallb) {
		var ret string
		return ret
	}
	return *o.Smallb
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateBInner) GetSmallbOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallb) {
		return nil, false
	}
	return o.Smallb, true
}

// HasB returns a boolean if a field has been set.
func (o *AccountUpdateBInner) HasSmallb() bool {
	if o != nil && !common.IsNil(o.Smallb) {
		return true
	}

	return false
}

// SetB gets a reference to the given string and assigns it to the B field.
func (o *AccountUpdateBInner) SetSmallb(v string) {
	o.Smallb = &v
}

// GetM returns the M field value if set, zero value otherwise.
func (o *AccountUpdateBInner) GetSmallm() string {
	if o == nil || common.IsNil(o.Smallm) {
		var ret string
		return ret
	}
	return *o.Smallm
}

// GetMOk returns a tuple with the M field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateBInner) GetSmallmOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallm) {
		return nil, false
	}
	return o.Smallm, true
}

// HasM returns a boolean if a field has been set.
func (o *AccountUpdateBInner) HasSmallm() bool {
	if o != nil && !common.IsNil(o.Smallm) {
		return true
	}

	return false
}

// SetM gets a reference to the given string and assigns it to the M field.
func (o *AccountUpdateBInner) SetSmallm(v string) {
	o.Smallm = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *AccountUpdateBInner) GetSmallu() string {
	if o == nil || common.IsNil(o.Smallu) {
		var ret string
		return ret
	}
	return *o.Smallu
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateBInner) GetSmalluOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallu) {
		return nil, false
	}
	return o.Smallu, true
}

// HasU returns a boolean if a field has been set.
func (o *AccountUpdateBInner) HasSmallu() bool {
	if o != nil && !common.IsNil(o.Smallu) {
		return true
	}

	return false
}

// SetU gets a reference to the given string and assigns it to the U field.
func (o *AccountUpdateBInner) SetSmallu(v string) {
	o.Smallu = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *AccountUpdateBInner) GetU() int64 {
	if o == nil || common.IsNil(o.U) {
		var ret int64
		return ret
	}
	return *o.U
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateBInner) GetUOk() (*int64, bool) {
	if o == nil || common.IsNil(o.U) {
		return nil, false
	}
	return o.U, true
}

// HasU returns a boolean if a field has been set.
func (o *AccountUpdateBInner) HasU() bool {
	if o != nil && !common.IsNil(o.U) {
		return true
	}

	return false
}

// SetU gets a reference to the given int64 and assigns it to the U field.
func (o *AccountUpdateBInner) SetU(v int64) {
	o.U = &v
}

// GetM returns the M field value if set, zero value otherwise.
func (o *AccountUpdateBInner) GetM() string {
	if o == nil || common.IsNil(o.M) {
		var ret string
		return ret
	}
	return *o.M
}

// GetMOk returns a tuple with the M field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateBInner) GetMOk() (*string, bool) {
	if o == nil || common.IsNil(o.M) {
		return nil, false
	}
	return o.M, true
}

// HasM returns a boolean if a field has been set.
func (o *AccountUpdateBInner) HasM() bool {
	if o != nil && !common.IsNil(o.M) {
		return true
	}

	return false
}

// SetM gets a reference to the given string and assigns it to the M field.
func (o *AccountUpdateBInner) SetM(v string) {
	o.M = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *AccountUpdateBInner) GetSmalli() string {
	if o == nil || common.IsNil(o.Smalli) {
		var ret string
		return ret
	}
	return *o.Smalli
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateBInner) GetSmalliOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalli) {
		return nil, false
	}
	return o.Smalli, true
}

// HasI returns a boolean if a field has been set.
func (o *AccountUpdateBInner) HasSmalli() bool {
	if o != nil && !common.IsNil(o.Smalli) {
		return true
	}

	return false
}

// SetI gets a reference to the given string and assigns it to the I field.
func (o *AccountUpdateBInner) SetSmalli(v string) {
	o.Smalli = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *AccountUpdateBInner) GetSmalla() string {
	if o == nil || common.IsNil(o.Smalla) {
		var ret string
		return ret
	}
	return *o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateBInner) GetSmallaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *AccountUpdateBInner) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *AccountUpdateBInner) SetSmalla(v string) {
	o.Smalla = &v
}

func (o AccountUpdateBInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountUpdateBInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smallb) {
		toSerialize["b"] = o.Smallb
	}
	if !common.IsNil(o.Smallm) {
		toSerialize["m"] = o.Smallm
	}
	if !common.IsNil(o.Smallu) {
		toSerialize["u"] = o.Smallu
	}
	if !common.IsNil(o.U) {
		toSerialize["U"] = o.U
	}
	if !common.IsNil(o.M) {
		toSerialize["M"] = o.M
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

func (o *AccountUpdateBInner) UnmarshalJSON(data []byte) (err error) {
	varAccountUpdateBInner := _AccountUpdateBInner{}

	err = json.Unmarshal(data, &varAccountUpdateBInner)

	if err != nil {
		return err
	}

	*o = AccountUpdateBInner(varAccountUpdateBInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "b")
		delete(additionalProperties, "m")
		delete(additionalProperties, "u")
		delete(additionalProperties, "U")
		delete(additionalProperties, "M")
		delete(additionalProperties, "i")
		delete(additionalProperties, "a")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountUpdateBInner struct {
	value *AccountUpdateBInner
	isSet bool
}

func (v NullableAccountUpdateBInner) Get() *AccountUpdateBInner {
	return v.value
}

func (v *NullableAccountUpdateBInner) Set(val *AccountUpdateBInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountUpdateBInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountUpdateBInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountUpdateBInner(val *AccountUpdateBInner) *NullableAccountUpdateBInner {
	return &NullableAccountUpdateBInner{value: val, isSet: true}
}

func (v NullableAccountUpdateBInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountUpdateBInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
