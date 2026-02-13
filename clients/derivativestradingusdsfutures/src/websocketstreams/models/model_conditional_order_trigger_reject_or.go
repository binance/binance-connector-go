/*
Binance Derivatives Trading USDS Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ConditionalOrderTriggerRejectOr type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ConditionalOrderTriggerRejectOr{}

// ConditionalOrderTriggerRejectOr struct for ConditionalOrderTriggerRejectOr
type ConditionalOrderTriggerRejectOr struct {
	Smalls               *string `json:"s,omitempty"`
	Smalli               *int64  `json:"i,omitempty"`
	Smallr               *string `json:"r,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConditionalOrderTriggerRejectOr ConditionalOrderTriggerRejectOr

// NewConditionalOrderTriggerRejectOr instantiates a new ConditionalOrderTriggerRejectOr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionalOrderTriggerRejectOr() *ConditionalOrderTriggerRejectOr {
	this := ConditionalOrderTriggerRejectOr{}
	return &this
}

// NewConditionalOrderTriggerRejectOrWithDefaults instantiates a new ConditionalOrderTriggerRejectOr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionalOrderTriggerRejectOrWithDefaults() *ConditionalOrderTriggerRejectOr {
	this := ConditionalOrderTriggerRejectOr{}
	return &this
}

// GetS returns the S field value if set, zero value otherwise.
func (o *ConditionalOrderTriggerRejectOr) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalOrderTriggerRejectOr) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *ConditionalOrderTriggerRejectOr) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *ConditionalOrderTriggerRejectOr) SetSmalls(v string) {
	o.Smalls = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *ConditionalOrderTriggerRejectOr) GetSmalli() int64 {
	if o == nil || common.IsNil(o.Smalli) {
		var ret int64
		return ret
	}
	return *o.Smalli
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalOrderTriggerRejectOr) GetSmalliOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smalli) {
		return nil, false
	}
	return o.Smalli, true
}

// HasI returns a boolean if a field has been set.
func (o *ConditionalOrderTriggerRejectOr) HasSmalli() bool {
	if o != nil && !common.IsNil(o.Smalli) {
		return true
	}

	return false
}

// SetI gets a reference to the given int64 and assigns it to the I field.
func (o *ConditionalOrderTriggerRejectOr) SetSmalli(v int64) {
	o.Smalli = &v
}

// GetR returns the R field value if set, zero value otherwise.
func (o *ConditionalOrderTriggerRejectOr) GetSmallr() string {
	if o == nil || common.IsNil(o.Smallr) {
		var ret string
		return ret
	}
	return *o.Smallr
}

// GetROk returns a tuple with the R field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalOrderTriggerRejectOr) GetSmallrOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallr) {
		return nil, false
	}
	return o.Smallr, true
}

// HasR returns a boolean if a field has been set.
func (o *ConditionalOrderTriggerRejectOr) HasSmallr() bool {
	if o != nil && !common.IsNil(o.Smallr) {
		return true
	}

	return false
}

// SetR gets a reference to the given string and assigns it to the R field.
func (o *ConditionalOrderTriggerRejectOr) SetSmallr(v string) {
	o.Smallr = &v
}

func (o ConditionalOrderTriggerRejectOr) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConditionalOrderTriggerRejectOr) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smalli) {
		toSerialize["i"] = o.Smalli
	}
	if !common.IsNil(o.Smallr) {
		toSerialize["r"] = o.Smallr
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConditionalOrderTriggerRejectOr) UnmarshalJSON(data []byte) (err error) {
	varConditionalOrderTriggerRejectOr := _ConditionalOrderTriggerRejectOr{}

	err = json.Unmarshal(data, &varConditionalOrderTriggerRejectOr)

	if err != nil {
		return err
	}

	*o = ConditionalOrderTriggerRejectOr(varConditionalOrderTriggerRejectOr)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "s")
		delete(additionalProperties, "i")
		delete(additionalProperties, "r")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConditionalOrderTriggerRejectOr struct {
	value *ConditionalOrderTriggerRejectOr
	isSet bool
}

func (v NullableConditionalOrderTriggerRejectOr) Get() *ConditionalOrderTriggerRejectOr {
	return v.value
}

func (v *NullableConditionalOrderTriggerRejectOr) Set(val *ConditionalOrderTriggerRejectOr) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionalOrderTriggerRejectOr) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionalOrderTriggerRejectOr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionalOrderTriggerRejectOr(val *ConditionalOrderTriggerRejectOr) *NullableConditionalOrderTriggerRejectOr {
	return &NullableConditionalOrderTriggerRejectOr{value: val, isSet: true}
}

func (v NullableConditionalOrderTriggerRejectOr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionalOrderTriggerRejectOr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
