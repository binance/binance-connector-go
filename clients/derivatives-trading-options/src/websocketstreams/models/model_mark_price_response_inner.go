/*
Binance Derivatives Trading Options WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Options WebSocket Market Streams

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MarkPriceResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarkPriceResponseInner{}

// MarkPriceResponseInner struct for MarkPriceResponseInner
type MarkPriceResponseInner struct {
	Smalle               *string `json:"e,omitempty"`
	E                    *int64  `json:"E,omitempty"`
	Smalls               *string `json:"s,omitempty"`
	Smallmp              *string `json:"mp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarkPriceResponseInner MarkPriceResponseInner

// NewMarkPriceResponseInner instantiates a new MarkPriceResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarkPriceResponseInner() *MarkPriceResponseInner {
	this := MarkPriceResponseInner{}
	return &this
}

// NewMarkPriceResponseInnerWithDefaults instantiates a new MarkPriceResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarkPriceResponseInnerWithDefaults() *MarkPriceResponseInner {
	this := MarkPriceResponseInner{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *MarkPriceResponseInner) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponseInner) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *MarkPriceResponseInner) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *MarkPriceResponseInner) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *MarkPriceResponseInner) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponseInner) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *MarkPriceResponseInner) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *MarkPriceResponseInner) SetE(v int64) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *MarkPriceResponseInner) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponseInner) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *MarkPriceResponseInner) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *MarkPriceResponseInner) SetSmalls(v string) {
	o.Smalls = &v
}

// GetMp returns the Mp field value if set, zero value otherwise.
func (o *MarkPriceResponseInner) GetSmallmp() string {
	if o == nil || common.IsNil(o.Smallmp) {
		var ret string
		return ret
	}
	return *o.Smallmp
}

// GetMpOk returns a tuple with the Mp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponseInner) GetSmallmpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallmp) {
		return nil, false
	}
	return o.Smallmp, true
}

// HasMp returns a boolean if a field has been set.
func (o *MarkPriceResponseInner) HasSmallmp() bool {
	if o != nil && !common.IsNil(o.Smallmp) {
		return true
	}

	return false
}

// SetMp gets a reference to the given string and assigns it to the Mp field.
func (o *MarkPriceResponseInner) SetSmallmp(v string) {
	o.Smallmp = &v
}

func (o MarkPriceResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarkPriceResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalle) {
		toSerialize["e"] = o.Smalle
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smallmp) {
		toSerialize["mp"] = o.Smallmp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MarkPriceResponseInner) UnmarshalJSON(data []byte) (err error) {
	varMarkPriceResponseInner := _MarkPriceResponseInner{}

	err = json.Unmarshal(data, &varMarkPriceResponseInner)

	if err != nil {
		return err
	}

	*o = MarkPriceResponseInner(varMarkPriceResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "s")
		delete(additionalProperties, "mp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarkPriceResponseInner struct {
	value *MarkPriceResponseInner
	isSet bool
}

func (v NullableMarkPriceResponseInner) Get() *MarkPriceResponseInner {
	return v.value
}

func (v *NullableMarkPriceResponseInner) Set(val *MarkPriceResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMarkPriceResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMarkPriceResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarkPriceResponseInner(val *MarkPriceResponseInner) *NullableMarkPriceResponseInner {
	return &NullableMarkPriceResponseInner{value: val, isSet: true}
}

func (v NullableMarkPriceResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarkPriceResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
