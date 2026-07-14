/*
Spot WebSocket Market Streams

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the OutboundAccountPosition type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OutboundAccountPosition{}

// OutboundAccountPosition struct for OutboundAccountPosition
type OutboundAccountPosition struct {
	// Event Time
	E *int64 `json:"E,omitempty"`
	// Time of last account update
	Smallu *int64 `json:"u,omitempty"`
	// Balances Array
	B                    []OutboundAccountPositionBInner `json:"B,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OutboundAccountPosition OutboundAccountPosition

// NewOutboundAccountPosition instantiates a new OutboundAccountPosition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutboundAccountPosition() *OutboundAccountPosition {
	this := OutboundAccountPosition{}
	return &this
}

// NewOutboundAccountPositionWithDefaults instantiates a new OutboundAccountPosition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutboundAccountPositionWithDefaults() *OutboundAccountPosition {
	this := OutboundAccountPosition{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *OutboundAccountPosition) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutboundAccountPosition) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *OutboundAccountPosition) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *OutboundAccountPosition) SetE(v int64) {
	o.E = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *OutboundAccountPosition) GetSmallu() int64 {
	if o == nil || common.IsNil(o.Smallu) {
		var ret int64
		return ret
	}
	return *o.Smallu
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutboundAccountPosition) GetSmalluOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallu) {
		return nil, false
	}
	return o.Smallu, true
}

// HasU returns a boolean if a field has been set.
func (o *OutboundAccountPosition) HasSmallu() bool {
	if o != nil && !common.IsNil(o.Smallu) {
		return true
	}

	return false
}

// SetU gets a reference to the given int64 and assigns it to the U field.
func (o *OutboundAccountPosition) SetSmallu(v int64) {
	o.Smallu = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *OutboundAccountPosition) GetB() []OutboundAccountPositionBInner {
	if o == nil || common.IsNil(o.B) {
		var ret []OutboundAccountPositionBInner
		return ret
	}
	return o.B
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutboundAccountPosition) GetBOk() ([]OutboundAccountPositionBInner, bool) {
	if o == nil || common.IsNil(o.B) {
		return nil, false
	}
	return o.B, true
}

// HasB returns a boolean if a field has been set.
func (o *OutboundAccountPosition) HasB() bool {
	if o != nil && !common.IsNil(o.B) {
		return true
	}

	return false
}

// SetB gets a reference to the given []OutboundAccountPositionBInner and assigns it to the B field.
func (o *OutboundAccountPosition) SetB(v []OutboundAccountPositionBInner) {
	o.B = v
}

func (o OutboundAccountPosition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutboundAccountPosition) ToMap() (map[string]interface{}, error) {
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

func (o *OutboundAccountPosition) UnmarshalJSON(data []byte) (err error) {
	varOutboundAccountPosition := _OutboundAccountPosition{}

	err = json.Unmarshal(data, &varOutboundAccountPosition)

	if err != nil {
		return err
	}

	*o = OutboundAccountPosition(varOutboundAccountPosition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "E")
		delete(additionalProperties, "u")
		delete(additionalProperties, "B")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOutboundAccountPosition struct {
	value *OutboundAccountPosition
	isSet bool
}

func (v NullableOutboundAccountPosition) Get() *OutboundAccountPosition {
	return v.value
}

func (v *NullableOutboundAccountPosition) Set(val *OutboundAccountPosition) {
	v.value = val
	v.isSet = true
}

func (v NullableOutboundAccountPosition) IsSet() bool {
	return v.isSet
}

func (v *NullableOutboundAccountPosition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutboundAccountPosition(val *OutboundAccountPosition) *NullableOutboundAccountPosition {
	return &NullableOutboundAccountPosition{value: val, isSet: true}
}

func (v NullableOutboundAccountPosition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutboundAccountPosition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
