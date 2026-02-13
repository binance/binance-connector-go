/*
Binance Margin Trading WebSocket Market Streams

OpenAPI Specification for the Binance Margin Trading WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ListstatusOInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ListstatusOInner{}

// ListstatusOInner struct for ListstatusOInner
type ListstatusOInner struct {
	Smalls               *string `json:"s,omitempty"`
	Smalli               *int64  `json:"i,omitempty"`
	Smallc               *string `json:"c,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListstatusOInner ListstatusOInner

// NewListstatusOInner instantiates a new ListstatusOInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListstatusOInner() *ListstatusOInner {
	this := ListstatusOInner{}
	return &this
}

// NewListstatusOInnerWithDefaults instantiates a new ListstatusOInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListstatusOInnerWithDefaults() *ListstatusOInner {
	this := ListstatusOInner{}
	return &this
}

// GetS returns the S field value if set, zero value otherwise.
func (o *ListstatusOInner) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListstatusOInner) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *ListstatusOInner) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *ListstatusOInner) SetSmalls(v string) {
	o.Smalls = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *ListstatusOInner) GetSmalli() int64 {
	if o == nil || common.IsNil(o.Smalli) {
		var ret int64
		return ret
	}
	return *o.Smalli
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListstatusOInner) GetSmalliOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smalli) {
		return nil, false
	}
	return o.Smalli, true
}

// HasI returns a boolean if a field has been set.
func (o *ListstatusOInner) HasSmalli() bool {
	if o != nil && !common.IsNil(o.Smalli) {
		return true
	}

	return false
}

// SetI gets a reference to the given int64 and assigns it to the I field.
func (o *ListstatusOInner) SetSmalli(v int64) {
	o.Smalli = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *ListstatusOInner) GetSmallc() string {
	if o == nil || common.IsNil(o.Smallc) {
		var ret string
		return ret
	}
	return *o.Smallc
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListstatusOInner) GetSmallcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallc) {
		return nil, false
	}
	return o.Smallc, true
}

// HasC returns a boolean if a field has been set.
func (o *ListstatusOInner) HasSmallc() bool {
	if o != nil && !common.IsNil(o.Smallc) {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *ListstatusOInner) SetSmallc(v string) {
	o.Smallc = &v
}

func (o ListstatusOInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListstatusOInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smalli) {
		toSerialize["i"] = o.Smalli
	}
	if !common.IsNil(o.Smallc) {
		toSerialize["c"] = o.Smallc
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListstatusOInner) UnmarshalJSON(data []byte) (err error) {
	varListstatusOInner := _ListstatusOInner{}

	err = json.Unmarshal(data, &varListstatusOInner)

	if err != nil {
		return err
	}

	*o = ListstatusOInner(varListstatusOInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "s")
		delete(additionalProperties, "i")
		delete(additionalProperties, "c")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListstatusOInner struct {
	value *ListstatusOInner
	isSet bool
}

func (v NullableListstatusOInner) Get() *ListstatusOInner {
	return v.value
}

func (v *NullableListstatusOInner) Set(val *ListstatusOInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListstatusOInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListstatusOInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListstatusOInner(val *ListstatusOInner) *NullableListstatusOInner {
	return &NullableListstatusOInner{value: val, isSet: true}
}

func (v NullableListstatusOInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListstatusOInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
