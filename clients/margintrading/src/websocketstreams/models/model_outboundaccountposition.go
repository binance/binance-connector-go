/*
Binance Margin Trading WebSocket Market Streams

OpenAPI Specification for the Binance Margin Trading WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the Outboundaccountposition type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Outboundaccountposition{}

// Outboundaccountposition struct for Outboundaccountposition
type Outboundaccountposition struct {
	E                    *int64                          `json:"E,omitempty"`
	Smallu               *int64                          `json:"u,omitempty"`
	B                    []OutboundaccountpositionBInner `json:"B,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Outboundaccountposition Outboundaccountposition

// NewOutboundaccountposition instantiates a new Outboundaccountposition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutboundaccountposition() *Outboundaccountposition {
	this := Outboundaccountposition{}
	return &this
}

// NewOutboundaccountpositionWithDefaults instantiates a new Outboundaccountposition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutboundaccountpositionWithDefaults() *Outboundaccountposition {
	this := Outboundaccountposition{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *Outboundaccountposition) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outboundaccountposition) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *Outboundaccountposition) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *Outboundaccountposition) SetE(v int64) {
	o.E = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *Outboundaccountposition) GetSmallu() int64 {
	if o == nil || common.IsNil(o.Smallu) {
		var ret int64
		return ret
	}
	return *o.Smallu
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outboundaccountposition) GetSmalluOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallu) {
		return nil, false
	}
	return o.Smallu, true
}

// HasU returns a boolean if a field has been set.
func (o *Outboundaccountposition) HasSmallu() bool {
	if o != nil && !common.IsNil(o.Smallu) {
		return true
	}

	return false
}

// SetU gets a reference to the given int64 and assigns it to the U field.
func (o *Outboundaccountposition) SetSmallu(v int64) {
	o.Smallu = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *Outboundaccountposition) GetB() []OutboundaccountpositionBInner {
	if o == nil || common.IsNil(o.B) {
		var ret []OutboundaccountpositionBInner
		return ret
	}
	return o.B
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outboundaccountposition) GetBOk() ([]OutboundaccountpositionBInner, bool) {
	if o == nil || common.IsNil(o.B) {
		return nil, false
	}
	return o.B, true
}

// HasB returns a boolean if a field has been set.
func (o *Outboundaccountposition) HasB() bool {
	if o != nil && !common.IsNil(o.B) {
		return true
	}

	return false
}

// SetB gets a reference to the given []OutboundaccountpositionBInner and assigns it to the B field.
func (o *Outboundaccountposition) SetB(v []OutboundaccountpositionBInner) {
	o.B = v
}

func (o Outboundaccountposition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Outboundaccountposition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smallu) {
		toSerialize["u"] = o.Smallu
	}
	if !common.IsNil(o.B) {
		toSerialize["B"] = o.B
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Outboundaccountposition) UnmarshalJSON(data []byte) (err error) {
	varOutboundaccountposition := _Outboundaccountposition{}

	err = json.Unmarshal(data, &varOutboundaccountposition)

	if err != nil {
		return err
	}

	*o = Outboundaccountposition(varOutboundaccountposition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "E")
		delete(additionalProperties, "u")
		delete(additionalProperties, "B")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOutboundaccountposition struct {
	value *Outboundaccountposition
	isSet bool
}

func (v NullableOutboundaccountposition) Get() *Outboundaccountposition {
	return v.value
}

func (v *NullableOutboundaccountposition) Set(val *Outboundaccountposition) {
	v.value = val
	v.isSet = true
}

func (v NullableOutboundaccountposition) IsSet() bool {
	return v.isSet
}

func (v *NullableOutboundaccountposition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutboundaccountposition(val *Outboundaccountposition) *NullableOutboundaccountposition {
	return &NullableOutboundaccountposition{value: val, isSet: true}
}

func (v NullableOutboundaccountposition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutboundaccountposition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
