/*
Binance Derivatives Trading Options WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Options WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GreekUpdate type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GreekUpdate{}

// GreekUpdate struct for GreekUpdate
type GreekUpdate struct {
	E                    *int64              `json:"E,omitempty"`
	T                    *int64              `json:"T,omitempty"`
	G                    []GreekUpdateGInner `json:"G,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GreekUpdate GreekUpdate

// NewGreekUpdate instantiates a new GreekUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGreekUpdate() *GreekUpdate {
	this := GreekUpdate{}
	return &this
}

// NewGreekUpdateWithDefaults instantiates a new GreekUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGreekUpdateWithDefaults() *GreekUpdate {
	this := GreekUpdate{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *GreekUpdate) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GreekUpdate) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *GreekUpdate) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *GreekUpdate) SetE(v int64) {
	o.E = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *GreekUpdate) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GreekUpdate) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *GreekUpdate) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *GreekUpdate) SetT(v int64) {
	o.T = &v
}

// GetG returns the G field value if set, zero value otherwise.
func (o *GreekUpdate) GetG() []GreekUpdateGInner {
	if o == nil || common.IsNil(o.G) {
		var ret []GreekUpdateGInner
		return ret
	}
	return o.G
}

// GetGOk returns a tuple with the G field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GreekUpdate) GetGOk() ([]GreekUpdateGInner, bool) {
	if o == nil || common.IsNil(o.G) {
		return nil, false
	}
	return o.G, true
}

// HasG returns a boolean if a field has been set.
func (o *GreekUpdate) HasG() bool {
	if o != nil && !common.IsNil(o.G) {
		return true
	}

	return false
}

// SetG gets a reference to the given []GreekUpdateGInner and assigns it to the G field.
func (o *GreekUpdate) SetG(v []GreekUpdateGInner) {
	o.G = v
}

func (o GreekUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GreekUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.G) {
		toSerialize["G"] = o.G
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GreekUpdate) UnmarshalJSON(data []byte) (err error) {
	varGreekUpdate := _GreekUpdate{}

	err = json.Unmarshal(data, &varGreekUpdate)

	if err != nil {
		return err
	}

	*o = GreekUpdate(varGreekUpdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "E")
		delete(additionalProperties, "T")
		delete(additionalProperties, "G")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGreekUpdate struct {
	value *GreekUpdate
	isSet bool
}

func (v NullableGreekUpdate) Get() *GreekUpdate {
	return v.value
}

func (v *NullableGreekUpdate) Set(val *GreekUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableGreekUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableGreekUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGreekUpdate(val *GreekUpdate) *NullableGreekUpdate {
	return &NullableGreekUpdate{value: val, isSet: true}
}

func (v NullableGreekUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGreekUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
