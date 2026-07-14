/*
Spot WebSocket Market Streams

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the OutboundAccountPositionBInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OutboundAccountPositionBInner{}

// OutboundAccountPositionBInner struct for OutboundAccountPositionBInner
type OutboundAccountPositionBInner struct {
	// Asset
	Smalla *string `json:"a,omitempty"`
	// Free
	Smallf *string `json:"f,omitempty"`
	// Locked
	Smalll               *string `json:"l,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OutboundAccountPositionBInner OutboundAccountPositionBInner

// NewOutboundAccountPositionBInner instantiates a new OutboundAccountPositionBInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutboundAccountPositionBInner() *OutboundAccountPositionBInner {
	this := OutboundAccountPositionBInner{}
	return &this
}

// NewOutboundAccountPositionBInnerWithDefaults instantiates a new OutboundAccountPositionBInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutboundAccountPositionBInnerWithDefaults() *OutboundAccountPositionBInner {
	this := OutboundAccountPositionBInner{}
	return &this
}

// GetA returns the A field value if set, zero value otherwise.
func (o *OutboundAccountPositionBInner) GetSmalla() string {
	if o == nil || common.IsNil(o.Smalla) {
		var ret string
		return ret
	}
	return *o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutboundAccountPositionBInner) GetSmallaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *OutboundAccountPositionBInner) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *OutboundAccountPositionBInner) SetSmalla(v string) {
	o.Smalla = &v
}

// GetF returns the F field value if set, zero value otherwise.
func (o *OutboundAccountPositionBInner) GetSmallf() string {
	if o == nil || common.IsNil(o.Smallf) {
		var ret string
		return ret
	}
	return *o.Smallf
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutboundAccountPositionBInner) GetSmallfOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallf) {
		return nil, false
	}
	return o.Smallf, true
}

// HasF returns a boolean if a field has been set.
func (o *OutboundAccountPositionBInner) HasSmallf() bool {
	if o != nil && !common.IsNil(o.Smallf) {
		return true
	}

	return false
}

// SetF gets a reference to the given string and assigns it to the F field.
func (o *OutboundAccountPositionBInner) SetSmallf(v string) {
	o.Smallf = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *OutboundAccountPositionBInner) GetSmalll() string {
	if o == nil || common.IsNil(o.Smalll) {
		var ret string
		return ret
	}
	return *o.Smalll
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutboundAccountPositionBInner) GetSmalllOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalll) {
		return nil, false
	}
	return o.Smalll, true
}

// HasL returns a boolean if a field has been set.
func (o *OutboundAccountPositionBInner) HasSmalll() bool {
	if o != nil && !common.IsNil(o.Smalll) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *OutboundAccountPositionBInner) SetSmalll(v string) {
	o.Smalll = &v
}

func (o OutboundAccountPositionBInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutboundAccountPositionBInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalla) {
		toSerialize["a"] = o.Smalla
	}
	if !common.IsNil(o.Smallf) {
		toSerialize["f"] = o.Smallf
	}
	if !common.IsNil(o.Smalll) {
		toSerialize["l"] = o.Smalll
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OutboundAccountPositionBInner) UnmarshalJSON(data []byte) (err error) {
	varOutboundAccountPositionBInner := _OutboundAccountPositionBInner{}

	err = json.Unmarshal(data, &varOutboundAccountPositionBInner)

	if err != nil {
		return err
	}

	*o = OutboundAccountPositionBInner(varOutboundAccountPositionBInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "a")
		delete(additionalProperties, "f")
		delete(additionalProperties, "l")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOutboundAccountPositionBInner struct {
	value *OutboundAccountPositionBInner
	isSet bool
}

func (v NullableOutboundAccountPositionBInner) Get() *OutboundAccountPositionBInner {
	return v.value
}

func (v *NullableOutboundAccountPositionBInner) Set(val *OutboundAccountPositionBInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOutboundAccountPositionBInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOutboundAccountPositionBInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutboundAccountPositionBInner(val *OutboundAccountPositionBInner) *NullableOutboundAccountPositionBInner {
	return &NullableOutboundAccountPositionBInner{value: val, isSet: true}
}

func (v NullableOutboundAccountPositionBInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutboundAccountPositionBInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
