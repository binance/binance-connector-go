/*
Binance Margin Trading WebSocket Market Streams

OpenAPI Specification for the Binance Margin Trading WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OutboundaccountpositionBInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OutboundaccountpositionBInner{}

// OutboundaccountpositionBInner struct for OutboundaccountpositionBInner
type OutboundaccountpositionBInner struct {
	Smalla               *string `json:"a,omitempty"`
	Smallf               *string `json:"f,omitempty"`
	Smalll               *string `json:"l,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OutboundaccountpositionBInner OutboundaccountpositionBInner

// NewOutboundaccountpositionBInner instantiates a new OutboundaccountpositionBInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutboundaccountpositionBInner() *OutboundaccountpositionBInner {
	this := OutboundaccountpositionBInner{}
	return &this
}

// NewOutboundaccountpositionBInnerWithDefaults instantiates a new OutboundaccountpositionBInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutboundaccountpositionBInnerWithDefaults() *OutboundaccountpositionBInner {
	this := OutboundaccountpositionBInner{}
	return &this
}

// GetA returns the A field value if set, zero value otherwise.
func (o *OutboundaccountpositionBInner) GetSmalla() string {
	if o == nil || common.IsNil(o.Smalla) {
		var ret string
		return ret
	}
	return *o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutboundaccountpositionBInner) GetSmallaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *OutboundaccountpositionBInner) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *OutboundaccountpositionBInner) SetSmalla(v string) {
	o.Smalla = &v
}

// GetF returns the F field value if set, zero value otherwise.
func (o *OutboundaccountpositionBInner) GetSmallf() string {
	if o == nil || common.IsNil(o.Smallf) {
		var ret string
		return ret
	}
	return *o.Smallf
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutboundaccountpositionBInner) GetSmallfOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallf) {
		return nil, false
	}
	return o.Smallf, true
}

// HasF returns a boolean if a field has been set.
func (o *OutboundaccountpositionBInner) HasSmallf() bool {
	if o != nil && !common.IsNil(o.Smallf) {
		return true
	}

	return false
}

// SetF gets a reference to the given string and assigns it to the F field.
func (o *OutboundaccountpositionBInner) SetSmallf(v string) {
	o.Smallf = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *OutboundaccountpositionBInner) GetSmalll() string {
	if o == nil || common.IsNil(o.Smalll) {
		var ret string
		return ret
	}
	return *o.Smalll
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutboundaccountpositionBInner) GetSmalllOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalll) {
		return nil, false
	}
	return o.Smalll, true
}

// HasL returns a boolean if a field has been set.
func (o *OutboundaccountpositionBInner) HasSmalll() bool {
	if o != nil && !common.IsNil(o.Smalll) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *OutboundaccountpositionBInner) SetSmalll(v string) {
	o.Smalll = &v
}

func (o OutboundaccountpositionBInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutboundaccountpositionBInner) ToMap() (map[string]interface{}, error) {
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

func (o *OutboundaccountpositionBInner) UnmarshalJSON(data []byte) (err error) {
	varOutboundaccountpositionBInner := _OutboundaccountpositionBInner{}

	err = json.Unmarshal(data, &varOutboundaccountpositionBInner)

	if err != nil {
		return err
	}

	*o = OutboundaccountpositionBInner(varOutboundaccountpositionBInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "a")
		delete(additionalProperties, "f")
		delete(additionalProperties, "l")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOutboundaccountpositionBInner struct {
	value *OutboundaccountpositionBInner
	isSet bool
}

func (v NullableOutboundaccountpositionBInner) Get() *OutboundaccountpositionBInner {
	return v.value
}

func (v *NullableOutboundaccountpositionBInner) Set(val *OutboundaccountpositionBInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOutboundaccountpositionBInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOutboundaccountpositionBInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutboundaccountpositionBInner(val *OutboundaccountpositionBInner) *NullableOutboundaccountpositionBInner {
	return &NullableOutboundaccountpositionBInner{value: val, isSet: true}
}

func (v NullableOutboundaccountpositionBInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutboundaccountpositionBInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
