/*
Binance Margin Trading WebSocket Market Streams

OpenAPI Specification for the Binance Margin Trading WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the MarginLevelStatusChange type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarginLevelStatusChange{}

// MarginLevelStatusChange struct for MarginLevelStatusChange
type MarginLevelStatusChange struct {
	E                    *int64  `json:"E,omitempty"`
	Smalll               *string `json:"l,omitempty"`
	Smalls               *string `json:"s,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarginLevelStatusChange MarginLevelStatusChange

// NewMarginLevelStatusChange instantiates a new MarginLevelStatusChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginLevelStatusChange() *MarginLevelStatusChange {
	this := MarginLevelStatusChange{}
	return &this
}

// NewMarginLevelStatusChangeWithDefaults instantiates a new MarginLevelStatusChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginLevelStatusChangeWithDefaults() *MarginLevelStatusChange {
	this := MarginLevelStatusChange{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *MarginLevelStatusChange) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginLevelStatusChange) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *MarginLevelStatusChange) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *MarginLevelStatusChange) SetE(v int64) {
	o.E = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *MarginLevelStatusChange) GetSmalll() string {
	if o == nil || common.IsNil(o.Smalll) {
		var ret string
		return ret
	}
	return *o.Smalll
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginLevelStatusChange) GetSmalllOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalll) {
		return nil, false
	}
	return o.Smalll, true
}

// HasL returns a boolean if a field has been set.
func (o *MarginLevelStatusChange) HasSmalll() bool {
	if o != nil && !common.IsNil(o.Smalll) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *MarginLevelStatusChange) SetSmalll(v string) {
	o.Smalll = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *MarginLevelStatusChange) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginLevelStatusChange) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *MarginLevelStatusChange) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *MarginLevelStatusChange) SetSmalls(v string) {
	o.Smalls = &v
}

func (o MarginLevelStatusChange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginLevelStatusChange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smalll) {
		toSerialize["l"] = o.Smalll
	}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MarginLevelStatusChange) UnmarshalJSON(data []byte) (err error) {
	varMarginLevelStatusChange := _MarginLevelStatusChange{}

	err = json.Unmarshal(data, &varMarginLevelStatusChange)

	if err != nil {
		return err
	}

	*o = MarginLevelStatusChange(varMarginLevelStatusChange)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "E")
		delete(additionalProperties, "l")
		delete(additionalProperties, "s")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginLevelStatusChange struct {
	value *MarginLevelStatusChange
	isSet bool
}

func (v NullableMarginLevelStatusChange) Get() *MarginLevelStatusChange {
	return v.value
}

func (v *NullableMarginLevelStatusChange) Set(val *MarginLevelStatusChange) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginLevelStatusChange) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginLevelStatusChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginLevelStatusChange(val *MarginLevelStatusChange) *NullableMarginLevelStatusChange {
	return &NullableMarginLevelStatusChange{value: val, isSet: true}
}

func (v NullableMarginLevelStatusChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginLevelStatusChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
