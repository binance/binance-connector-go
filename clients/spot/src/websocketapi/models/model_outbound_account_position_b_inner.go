/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OutboundAccountPositionBInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OutboundAccountPositionBInner{}

// OutboundAccountPositionBInner struct for OutboundAccountPositionBInner
type OutboundAccountPositionBInner struct {
	A                    *string `json:"a,omitempty"`
	F                    *string `json:"f,omitempty"`
	L                    *string `json:"l,omitempty"`
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
func (o *OutboundAccountPositionBInner) GetA() string {
	if o == nil || common.IsNil(o.A) {
		var ret string
		return ret
	}
	return *o.A
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutboundAccountPositionBInner) GetAOk() (*string, bool) {
	if o == nil || common.IsNil(o.A) {
		return nil, false
	}
	return o.A, true
}

// HasA returns a boolean if a field has been set.
func (o *OutboundAccountPositionBInner) HasA() bool {
	if o != nil && !common.IsNil(o.A) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *OutboundAccountPositionBInner) SetA(v string) {
	o.A = &v
}

// GetF returns the F field value if set, zero value otherwise.
func (o *OutboundAccountPositionBInner) GetF() string {
	if o == nil || common.IsNil(o.F) {
		var ret string
		return ret
	}
	return *o.F
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutboundAccountPositionBInner) GetFOk() (*string, bool) {
	if o == nil || common.IsNil(o.F) {
		return nil, false
	}
	return o.F, true
}

// HasF returns a boolean if a field has been set.
func (o *OutboundAccountPositionBInner) HasF() bool {
	if o != nil && !common.IsNil(o.F) {
		return true
	}

	return false
}

// SetF gets a reference to the given string and assigns it to the F field.
func (o *OutboundAccountPositionBInner) SetF(v string) {
	o.F = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *OutboundAccountPositionBInner) GetL() string {
	if o == nil || common.IsNil(o.L) {
		var ret string
		return ret
	}
	return *o.L
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutboundAccountPositionBInner) GetLOk() (*string, bool) {
	if o == nil || common.IsNil(o.L) {
		return nil, false
	}
	return o.L, true
}

// HasL returns a boolean if a field has been set.
func (o *OutboundAccountPositionBInner) HasL() bool {
	if o != nil && !common.IsNil(o.L) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *OutboundAccountPositionBInner) SetL(v string) {
	o.L = &v
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
	if !common.IsNil(o.A) {
		toSerialize["a"] = o.A
	}
	if !common.IsNil(o.F) {
		toSerialize["f"] = o.F
	}
	if !common.IsNil(o.L) {
		toSerialize["l"] = o.L
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
