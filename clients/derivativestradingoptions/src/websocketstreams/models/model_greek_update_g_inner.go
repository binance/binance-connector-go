/*
Binance Derivatives Trading Options WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Options WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GreekUpdateGInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GreekUpdateGInner{}

// GreekUpdateGInner struct for GreekUpdateGInner
type GreekUpdateGInner struct {
	Smallu               *string `json:"u,omitempty"`
	Smalld               *string `json:"d,omitempty"`
	Smallg               *string `json:"g,omitempty"`
	Smallt               *string `json:"t,omitempty"`
	Smallv               *string `json:"v,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GreekUpdateGInner GreekUpdateGInner

// NewGreekUpdateGInner instantiates a new GreekUpdateGInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGreekUpdateGInner() *GreekUpdateGInner {
	this := GreekUpdateGInner{}
	return &this
}

// NewGreekUpdateGInnerWithDefaults instantiates a new GreekUpdateGInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGreekUpdateGInnerWithDefaults() *GreekUpdateGInner {
	this := GreekUpdateGInner{}
	return &this
}

// GetU returns the U field value if set, zero value otherwise.
func (o *GreekUpdateGInner) GetSmallu() string {
	if o == nil || common.IsNil(o.Smallu) {
		var ret string
		return ret
	}
	return *o.Smallu
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GreekUpdateGInner) GetSmalluOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallu) {
		return nil, false
	}
	return o.Smallu, true
}

// HasU returns a boolean if a field has been set.
func (o *GreekUpdateGInner) HasSmallu() bool {
	if o != nil && !common.IsNil(o.Smallu) {
		return true
	}

	return false
}

// SetU gets a reference to the given string and assigns it to the U field.
func (o *GreekUpdateGInner) SetSmallu(v string) {
	o.Smallu = &v
}

// GetD returns the D field value if set, zero value otherwise.
func (o *GreekUpdateGInner) GetSmalld() string {
	if o == nil || common.IsNil(o.Smalld) {
		var ret string
		return ret
	}
	return *o.Smalld
}

// GetDOk returns a tuple with the D field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GreekUpdateGInner) GetSmalldOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalld) {
		return nil, false
	}
	return o.Smalld, true
}

// HasD returns a boolean if a field has been set.
func (o *GreekUpdateGInner) HasSmalld() bool {
	if o != nil && !common.IsNil(o.Smalld) {
		return true
	}

	return false
}

// SetD gets a reference to the given string and assigns it to the D field.
func (o *GreekUpdateGInner) SetSmalld(v string) {
	o.Smalld = &v
}

// GetG returns the G field value if set, zero value otherwise.
func (o *GreekUpdateGInner) GetSmallg() string {
	if o == nil || common.IsNil(o.Smallg) {
		var ret string
		return ret
	}
	return *o.Smallg
}

// GetGOk returns a tuple with the G field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GreekUpdateGInner) GetSmallgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallg) {
		return nil, false
	}
	return o.Smallg, true
}

// HasG returns a boolean if a field has been set.
func (o *GreekUpdateGInner) HasSmallg() bool {
	if o != nil && !common.IsNil(o.Smallg) {
		return true
	}

	return false
}

// SetG gets a reference to the given string and assigns it to the G field.
func (o *GreekUpdateGInner) SetSmallg(v string) {
	o.Smallg = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *GreekUpdateGInner) GetSmallt() string {
	if o == nil || common.IsNil(o.Smallt) {
		var ret string
		return ret
	}
	return *o.Smallt
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GreekUpdateGInner) GetSmalltOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallt) {
		return nil, false
	}
	return o.Smallt, true
}

// HasT returns a boolean if a field has been set.
func (o *GreekUpdateGInner) HasSmallt() bool {
	if o != nil && !common.IsNil(o.Smallt) {
		return true
	}

	return false
}

// SetT gets a reference to the given string and assigns it to the T field.
func (o *GreekUpdateGInner) SetSmallt(v string) {
	o.Smallt = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *GreekUpdateGInner) GetSmallv() string {
	if o == nil || common.IsNil(o.Smallv) {
		var ret string
		return ret
	}
	return *o.Smallv
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GreekUpdateGInner) GetSmallvOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallv) {
		return nil, false
	}
	return o.Smallv, true
}

// HasV returns a boolean if a field has been set.
func (o *GreekUpdateGInner) HasSmallv() bool {
	if o != nil && !common.IsNil(o.Smallv) {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *GreekUpdateGInner) SetSmallv(v string) {
	o.Smallv = &v
}

func (o GreekUpdateGInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GreekUpdateGInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smallu) {
		toSerialize["u"] = o.Smallu
	}
	if !common.IsNil(o.Smalld) {
		toSerialize["d"] = o.Smalld
	}
	if !common.IsNil(o.Smallg) {
		toSerialize["g"] = o.Smallg
	}
	if !common.IsNil(o.Smallt) {
		toSerialize["t"] = o.Smallt
	}
	if !common.IsNil(o.Smallv) {
		toSerialize["v"] = o.Smallv
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GreekUpdateGInner) UnmarshalJSON(data []byte) (err error) {
	varGreekUpdateGInner := _GreekUpdateGInner{}

	err = json.Unmarshal(data, &varGreekUpdateGInner)

	if err != nil {
		return err
	}

	*o = GreekUpdateGInner(varGreekUpdateGInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "u")
		delete(additionalProperties, "d")
		delete(additionalProperties, "g")
		delete(additionalProperties, "t")
		delete(additionalProperties, "v")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGreekUpdateGInner struct {
	value *GreekUpdateGInner
	isSet bool
}

func (v NullableGreekUpdateGInner) Get() *GreekUpdateGInner {
	return v.value
}

func (v *NullableGreekUpdateGInner) Set(val *GreekUpdateGInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGreekUpdateGInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGreekUpdateGInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGreekUpdateGInner(val *GreekUpdateGInner) *NullableGreekUpdateGInner {
	return &NullableGreekUpdateGInner{value: val, isSet: true}
}

func (v NullableGreekUpdateGInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGreekUpdateGInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
