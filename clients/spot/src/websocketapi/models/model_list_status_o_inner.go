/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the ListStatusOInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ListStatusOInner{}

// ListStatusOInner struct for ListStatusOInner
type ListStatusOInner struct {
	Smalls               *string `json:"s,omitempty"`
	Smalli               *int64  `json:"i,omitempty"`
	Smallc               *string `json:"c,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListStatusOInner ListStatusOInner

// NewListStatusOInner instantiates a new ListStatusOInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListStatusOInner() *ListStatusOInner {
	this := ListStatusOInner{}
	return &this
}

// NewListStatusOInnerWithDefaults instantiates a new ListStatusOInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListStatusOInnerWithDefaults() *ListStatusOInner {
	this := ListStatusOInner{}
	return &this
}

// GetS returns the S field value if set, zero value otherwise.
func (o *ListStatusOInner) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStatusOInner) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *ListStatusOInner) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *ListStatusOInner) SetSmalls(v string) {
	o.Smalls = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *ListStatusOInner) GetSmalli() int64 {
	if o == nil || common.IsNil(o.Smalli) {
		var ret int64
		return ret
	}
	return *o.Smalli
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStatusOInner) GetSmalliOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smalli) {
		return nil, false
	}
	return o.Smalli, true
}

// HasI returns a boolean if a field has been set.
func (o *ListStatusOInner) HasSmalli() bool {
	if o != nil && !common.IsNil(o.Smalli) {
		return true
	}

	return false
}

// SetI gets a reference to the given int64 and assigns it to the I field.
func (o *ListStatusOInner) SetSmalli(v int64) {
	o.Smalli = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *ListStatusOInner) GetSmallc() string {
	if o == nil || common.IsNil(o.Smallc) {
		var ret string
		return ret
	}
	return *o.Smallc
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStatusOInner) GetSmallcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallc) {
		return nil, false
	}
	return o.Smallc, true
}

// HasC returns a boolean if a field has been set.
func (o *ListStatusOInner) HasSmallc() bool {
	if o != nil && !common.IsNil(o.Smallc) {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *ListStatusOInner) SetSmallc(v string) {
	o.Smallc = &v
}

func (o ListStatusOInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListStatusOInner) ToMap() (map[string]interface{}, error) {
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

func (o *ListStatusOInner) UnmarshalJSON(data []byte) (err error) {
	varListStatusOInner := _ListStatusOInner{}

	err = json.Unmarshal(data, &varListStatusOInner)

	if err != nil {
		return err
	}

	*o = ListStatusOInner(varListStatusOInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "s")
		delete(additionalProperties, "i")
		delete(additionalProperties, "c")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListStatusOInner struct {
	value *ListStatusOInner
	isSet bool
}

func (v NullableListStatusOInner) Get() *ListStatusOInner {
	return v.value
}

func (v *NullableListStatusOInner) Set(val *ListStatusOInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListStatusOInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListStatusOInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListStatusOInner(val *ListStatusOInner) *NullableListStatusOInner {
	return &NullableListStatusOInner{value: val, isSet: true}
}

func (v NullableListStatusOInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListStatusOInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
