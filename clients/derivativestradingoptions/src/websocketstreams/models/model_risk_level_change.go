/*
Binance Derivatives Trading Options WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Options WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the RiskLevelChange type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RiskLevelChange{}

// RiskLevelChange struct for RiskLevelChange
type RiskLevelChange struct {
	E                    *int64  `json:"E,omitempty"`
	Smalls               *string `json:"s,omitempty"`
	Smallmb              *string `json:"mb,omitempty"`
	Smallmm              *string `json:"mm,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RiskLevelChange RiskLevelChange

// NewRiskLevelChange instantiates a new RiskLevelChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskLevelChange() *RiskLevelChange {
	this := RiskLevelChange{}
	return &this
}

// NewRiskLevelChangeWithDefaults instantiates a new RiskLevelChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskLevelChangeWithDefaults() *RiskLevelChange {
	this := RiskLevelChange{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *RiskLevelChange) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskLevelChange) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *RiskLevelChange) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *RiskLevelChange) SetE(v int64) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *RiskLevelChange) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskLevelChange) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *RiskLevelChange) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *RiskLevelChange) SetSmalls(v string) {
	o.Smalls = &v
}

// GetMb returns the Mb field value if set, zero value otherwise.
func (o *RiskLevelChange) GetSmallmb() string {
	if o == nil || common.IsNil(o.Smallmb) {
		var ret string
		return ret
	}
	return *o.Smallmb
}

// GetMbOk returns a tuple with the Mb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskLevelChange) GetSmallmbOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallmb) {
		return nil, false
	}
	return o.Smallmb, true
}

// HasMb returns a boolean if a field has been set.
func (o *RiskLevelChange) HasSmallmb() bool {
	if o != nil && !common.IsNil(o.Smallmb) {
		return true
	}

	return false
}

// SetMb gets a reference to the given string and assigns it to the Mb field.
func (o *RiskLevelChange) SetSmallmb(v string) {
	o.Smallmb = &v
}

// GetMm returns the Mm field value if set, zero value otherwise.
func (o *RiskLevelChange) GetSmallmm() string {
	if o == nil || common.IsNil(o.Smallmm) {
		var ret string
		return ret
	}
	return *o.Smallmm
}

// GetMmOk returns a tuple with the Mm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskLevelChange) GetSmallmmOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallmm) {
		return nil, false
	}
	return o.Smallmm, true
}

// HasMm returns a boolean if a field has been set.
func (o *RiskLevelChange) HasSmallmm() bool {
	if o != nil && !common.IsNil(o.Smallmm) {
		return true
	}

	return false
}

// SetMm gets a reference to the given string and assigns it to the Mm field.
func (o *RiskLevelChange) SetSmallmm(v string) {
	o.Smallmm = &v
}

func (o RiskLevelChange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskLevelChange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smallmb) {
		toSerialize["mb"] = o.Smallmb
	}
	if !common.IsNil(o.Smallmm) {
		toSerialize["mm"] = o.Smallmm
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RiskLevelChange) UnmarshalJSON(data []byte) (err error) {
	varRiskLevelChange := _RiskLevelChange{}

	err = json.Unmarshal(data, &varRiskLevelChange)

	if err != nil {
		return err
	}

	*o = RiskLevelChange(varRiskLevelChange)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "E")
		delete(additionalProperties, "s")
		delete(additionalProperties, "mb")
		delete(additionalProperties, "mm")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRiskLevelChange struct {
	value *RiskLevelChange
	isSet bool
}

func (v NullableRiskLevelChange) Get() *RiskLevelChange {
	return v.value
}

func (v *NullableRiskLevelChange) Set(val *RiskLevelChange) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskLevelChange) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskLevelChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskLevelChange(val *RiskLevelChange) *NullableRiskLevelChange {
	return &NullableRiskLevelChange{value: val, isSet: true}
}

func (v NullableRiskLevelChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskLevelChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
