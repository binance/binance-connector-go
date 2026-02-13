/*
Binance Derivatives Trading Portfolio Margin WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AccountUpdateA type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountUpdateA{}

// AccountUpdateA struct for AccountUpdateA
type AccountUpdateA struct {
	Smallm               *string                `json:"m,omitempty"`
	B                    []AccountUpdateABInner `json:"B,omitempty"`
	P                    []AccountUpdateAPInner `json:"P,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountUpdateA AccountUpdateA

// NewAccountUpdateA instantiates a new AccountUpdateA object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountUpdateA() *AccountUpdateA {
	this := AccountUpdateA{}
	return &this
}

// NewAccountUpdateAWithDefaults instantiates a new AccountUpdateA object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountUpdateAWithDefaults() *AccountUpdateA {
	this := AccountUpdateA{}
	return &this
}

// GetM returns the M field value if set, zero value otherwise.
func (o *AccountUpdateA) GetSmallm() string {
	if o == nil || common.IsNil(o.Smallm) {
		var ret string
		return ret
	}
	return *o.Smallm
}

// GetMOk returns a tuple with the M field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateA) GetSmallmOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallm) {
		return nil, false
	}
	return o.Smallm, true
}

// HasM returns a boolean if a field has been set.
func (o *AccountUpdateA) HasSmallm() bool {
	if o != nil && !common.IsNil(o.Smallm) {
		return true
	}

	return false
}

// SetM gets a reference to the given string and assigns it to the M field.
func (o *AccountUpdateA) SetSmallm(v string) {
	o.Smallm = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *AccountUpdateA) GetB() []AccountUpdateABInner {
	if o == nil || common.IsNil(o.B) {
		var ret []AccountUpdateABInner
		return ret
	}
	return o.B
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateA) GetBOk() ([]AccountUpdateABInner, bool) {
	if o == nil || common.IsNil(o.B) {
		return nil, false
	}
	return o.B, true
}

// HasB returns a boolean if a field has been set.
func (o *AccountUpdateA) HasB() bool {
	if o != nil && !common.IsNil(o.B) {
		return true
	}

	return false
}

// SetB gets a reference to the given []AccountUpdateABInner and assigns it to the B field.
func (o *AccountUpdateA) SetB(v []AccountUpdateABInner) {
	o.B = v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *AccountUpdateA) GetP() []AccountUpdateAPInner {
	if o == nil || common.IsNil(o.P) {
		var ret []AccountUpdateAPInner
		return ret
	}
	return o.P
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateA) GetPOk() ([]AccountUpdateAPInner, bool) {
	if o == nil || common.IsNil(o.P) {
		return nil, false
	}
	return o.P, true
}

// HasP returns a boolean if a field has been set.
func (o *AccountUpdateA) HasP() bool {
	if o != nil && !common.IsNil(o.P) {
		return true
	}

	return false
}

// SetP gets a reference to the given []AccountUpdateAPInner and assigns it to the P field.
func (o *AccountUpdateA) SetP(v []AccountUpdateAPInner) {
	o.P = v
}

func (o AccountUpdateA) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountUpdateA) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smallm) {
		toSerialize["m"] = o.Smallm
	}
	if !common.IsNil(o.B) {
		toSerialize["B"] = o.B
	}
	if !common.IsNil(o.P) {
		toSerialize["P"] = o.P
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountUpdateA) UnmarshalJSON(data []byte) (err error) {
	varAccountUpdateA := _AccountUpdateA{}

	err = json.Unmarshal(data, &varAccountUpdateA)

	if err != nil {
		return err
	}

	*o = AccountUpdateA(varAccountUpdateA)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "m")
		delete(additionalProperties, "B")
		delete(additionalProperties, "P")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountUpdateA struct {
	value *AccountUpdateA
	isSet bool
}

func (v NullableAccountUpdateA) Get() *AccountUpdateA {
	return v.value
}

func (v *NullableAccountUpdateA) Set(val *AccountUpdateA) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountUpdateA) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountUpdateA) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountUpdateA(val *AccountUpdateA) *NullableAccountUpdateA {
	return &NullableAccountUpdateA{value: val, isSet: true}
}

func (v NullableAccountUpdateA) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountUpdateA) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
