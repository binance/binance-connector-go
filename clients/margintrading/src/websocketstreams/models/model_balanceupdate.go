/*
Binance Margin Trading WebSocket Market Streams

OpenAPI Specification for the Binance Margin Trading WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the Balanceupdate type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Balanceupdate{}

// Balanceupdate struct for Balanceupdate
type Balanceupdate struct {
	E                    *int64  `json:"E,omitempty"`
	Smalla               *string `json:"a,omitempty"`
	Smalld               *string `json:"d,omitempty"`
	T                    *int64  `json:"T,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Balanceupdate Balanceupdate

// NewBalanceupdate instantiates a new Balanceupdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalanceupdate() *Balanceupdate {
	this := Balanceupdate{}
	return &this
}

// NewBalanceupdateWithDefaults instantiates a new Balanceupdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceupdateWithDefaults() *Balanceupdate {
	this := Balanceupdate{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *Balanceupdate) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Balanceupdate) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *Balanceupdate) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *Balanceupdate) SetE(v int64) {
	o.E = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *Balanceupdate) GetSmalla() string {
	if o == nil || common.IsNil(o.Smalla) {
		var ret string
		return ret
	}
	return *o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Balanceupdate) GetSmallaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *Balanceupdate) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *Balanceupdate) SetSmalla(v string) {
	o.Smalla = &v
}

// GetD returns the D field value if set, zero value otherwise.
func (o *Balanceupdate) GetSmalld() string {
	if o == nil || common.IsNil(o.Smalld) {
		var ret string
		return ret
	}
	return *o.Smalld
}

// GetDOk returns a tuple with the D field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Balanceupdate) GetSmalldOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalld) {
		return nil, false
	}
	return o.Smalld, true
}

// HasD returns a boolean if a field has been set.
func (o *Balanceupdate) HasSmalld() bool {
	if o != nil && !common.IsNil(o.Smalld) {
		return true
	}

	return false
}

// SetD gets a reference to the given string and assigns it to the D field.
func (o *Balanceupdate) SetSmalld(v string) {
	o.Smalld = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *Balanceupdate) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Balanceupdate) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *Balanceupdate) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *Balanceupdate) SetT(v int64) {
	o.T = &v
}

func (o Balanceupdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Balanceupdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smalla) {
		toSerialize["a"] = o.Smalla
	}
	if !common.IsNil(o.Smalld) {
		toSerialize["d"] = o.Smalld
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Balanceupdate) UnmarshalJSON(data []byte) (err error) {
	varBalanceupdate := _Balanceupdate{}

	err = json.Unmarshal(data, &varBalanceupdate)

	if err != nil {
		return err
	}

	*o = Balanceupdate(varBalanceupdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "E")
		delete(additionalProperties, "a")
		delete(additionalProperties, "d")
		delete(additionalProperties, "T")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBalanceupdate struct {
	value *Balanceupdate
	isSet bool
}

func (v NullableBalanceupdate) Get() *Balanceupdate {
	return v.value
}

func (v *NullableBalanceupdate) Set(val *Balanceupdate) {
	v.value = val
	v.isSet = true
}

func (v NullableBalanceupdate) IsSet() bool {
	return v.isSet
}

func (v *NullableBalanceupdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalanceupdate(val *Balanceupdate) *NullableBalanceupdate {
	return &NullableBalanceupdate{value: val, isSet: true}
}

func (v NullableBalanceupdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalanceupdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
