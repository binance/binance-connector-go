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

// checks if the AccountUpdateGInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountUpdateGInner{}

// AccountUpdateGInner struct for AccountUpdateGInner
type AccountUpdateGInner struct {
	Smallui              *string  `json:"ui,omitempty"`
	D                    *float32 `json:"d,omitempty"`
	T                    *float32 `json:"t,omitempty"`
	G                    *float32 `json:"g,omitempty"`
	V                    *float32 `json:"v,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountUpdateGInner AccountUpdateGInner

// NewAccountUpdateGInner instantiates a new AccountUpdateGInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountUpdateGInner() *AccountUpdateGInner {
	this := AccountUpdateGInner{}
	return &this
}

// NewAccountUpdateGInnerWithDefaults instantiates a new AccountUpdateGInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountUpdateGInnerWithDefaults() *AccountUpdateGInner {
	this := AccountUpdateGInner{}
	return &this
}

// GetUi returns the Ui field value if set, zero value otherwise.
func (o *AccountUpdateGInner) GetSmallui() string {
	if o == nil || common.IsNil(o.Smallui) {
		var ret string
		return ret
	}
	return *o.Smallui
}

// GetUiOk returns a tuple with the Ui field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateGInner) GetSmalluiOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallui) {
		return nil, false
	}
	return o.Smallui, true
}

// HasUi returns a boolean if a field has been set.
func (o *AccountUpdateGInner) HasSmallui() bool {
	if o != nil && !common.IsNil(o.Smallui) {
		return true
	}

	return false
}

// SetUi gets a reference to the given string and assigns it to the Ui field.
func (o *AccountUpdateGInner) SetSmallui(v string) {
	o.Smallui = &v
}

// GetD returns the D field value if set, zero value otherwise.
func (o *AccountUpdateGInner) GetD() float32 {
	if o == nil || common.IsNil(o.D) {
		var ret float32
		return ret
	}
	return *o.D
}

// GetDOk returns a tuple with the D field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateGInner) GetDOk() (*float32, bool) {
	if o == nil || common.IsNil(o.D) {
		return nil, false
	}
	return o.D, true
}

// HasD returns a boolean if a field has been set.
func (o *AccountUpdateGInner) HasD() bool {
	if o != nil && !common.IsNil(o.D) {
		return true
	}

	return false
}

// SetD gets a reference to the given float32 and assigns it to the D field.
func (o *AccountUpdateGInner) SetD(v float32) {
	o.D = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *AccountUpdateGInner) GetT() float32 {
	if o == nil || common.IsNil(o.T) {
		var ret float32
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateGInner) GetTOk() (*float32, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *AccountUpdateGInner) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given float32 and assigns it to the T field.
func (o *AccountUpdateGInner) SetT(v float32) {
	o.T = &v
}

// GetG returns the G field value if set, zero value otherwise.
func (o *AccountUpdateGInner) GetG() float32 {
	if o == nil || common.IsNil(o.G) {
		var ret float32
		return ret
	}
	return *o.G
}

// GetGOk returns a tuple with the G field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateGInner) GetGOk() (*float32, bool) {
	if o == nil || common.IsNil(o.G) {
		return nil, false
	}
	return o.G, true
}

// HasG returns a boolean if a field has been set.
func (o *AccountUpdateGInner) HasG() bool {
	if o != nil && !common.IsNil(o.G) {
		return true
	}

	return false
}

// SetG gets a reference to the given float32 and assigns it to the G field.
func (o *AccountUpdateGInner) SetG(v float32) {
	o.G = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *AccountUpdateGInner) GetV() float32 {
	if o == nil || common.IsNil(o.V) {
		var ret float32
		return ret
	}
	return *o.V
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateGInner) GetVOk() (*float32, bool) {
	if o == nil || common.IsNil(o.V) {
		return nil, false
	}
	return o.V, true
}

// HasV returns a boolean if a field has been set.
func (o *AccountUpdateGInner) HasV() bool {
	if o != nil && !common.IsNil(o.V) {
		return true
	}

	return false
}

// SetV gets a reference to the given float32 and assigns it to the V field.
func (o *AccountUpdateGInner) SetV(v float32) {
	o.V = &v
}

func (o AccountUpdateGInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountUpdateGInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smallui) {
		toSerialize["ui"] = o.Smallui
	}
	if !common.IsNil(o.D) {
		toSerialize["d"] = o.D
	}
	if !common.IsNil(o.T) {
		toSerialize["t"] = o.T
	}
	if !common.IsNil(o.G) {
		toSerialize["g"] = o.G
	}
	if !common.IsNil(o.V) {
		toSerialize["v"] = o.V
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountUpdateGInner) UnmarshalJSON(data []byte) (err error) {
	varAccountUpdateGInner := _AccountUpdateGInner{}

	err = json.Unmarshal(data, &varAccountUpdateGInner)

	if err != nil {
		return err
	}

	*o = AccountUpdateGInner(varAccountUpdateGInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ui")
		delete(additionalProperties, "d")
		delete(additionalProperties, "t")
		delete(additionalProperties, "g")
		delete(additionalProperties, "v")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountUpdateGInner struct {
	value *AccountUpdateGInner
	isSet bool
}

func (v NullableAccountUpdateGInner) Get() *AccountUpdateGInner {
	return v.value
}

func (v *NullableAccountUpdateGInner) Set(val *AccountUpdateGInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountUpdateGInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountUpdateGInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountUpdateGInner(val *AccountUpdateGInner) *NullableAccountUpdateGInner {
	return &NullableAccountUpdateGInner{value: val, isSet: true}
}

func (v NullableAccountUpdateGInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountUpdateGInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
