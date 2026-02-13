/*
Binance Derivatives Trading USDS Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GridUpdate type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GridUpdate{}

// GridUpdate struct for GridUpdate
type GridUpdate struct {
	T                    *int64        `json:"T,omitempty"`
	E                    *int64        `json:"E,omitempty"`
	Smallgu              *GridUpdateGu `json:"gu,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GridUpdate GridUpdate

// NewGridUpdate instantiates a new GridUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGridUpdate() *GridUpdate {
	this := GridUpdate{}
	return &this
}

// NewGridUpdateWithDefaults instantiates a new GridUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGridUpdateWithDefaults() *GridUpdate {
	this := GridUpdate{}
	return &this
}

// GetT returns the T field value if set, zero value otherwise.
func (o *GridUpdate) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridUpdate) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *GridUpdate) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *GridUpdate) SetT(v int64) {
	o.T = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *GridUpdate) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridUpdate) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *GridUpdate) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *GridUpdate) SetE(v int64) {
	o.E = &v
}

// GetGu returns the Gu field value if set, zero value otherwise.
func (o *GridUpdate) GetSmallgu() GridUpdateGu {
	if o == nil || common.IsNil(o.Smallgu) {
		var ret GridUpdateGu
		return ret
	}
	return *o.Smallgu
}

// GetGuOk returns a tuple with the Gu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridUpdate) GetSmallguOk() (*GridUpdateGu, bool) {
	if o == nil || common.IsNil(o.Smallgu) {
		return nil, false
	}
	return o.Smallgu, true
}

// HasGu returns a boolean if a field has been set.
func (o *GridUpdate) HasSmallgu() bool {
	if o != nil && !common.IsNil(o.Smallgu) {
		return true
	}

	return false
}

// SetGu gets a reference to the given GridUpdateGu and assigns it to the Gu field.
func (o *GridUpdate) SetSmallgu(v GridUpdateGu) {
	o.Smallgu = &v
}

func (o GridUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GridUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smallgu) {
		toSerialize["gu"] = o.Smallgu
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GridUpdate) UnmarshalJSON(data []byte) (err error) {
	varGridUpdate := _GridUpdate{}

	err = json.Unmarshal(data, &varGridUpdate)

	if err != nil {
		return err
	}

	*o = GridUpdate(varGridUpdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "T")
		delete(additionalProperties, "E")
		delete(additionalProperties, "gu")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGridUpdate struct {
	value *GridUpdate
	isSet bool
}

func (v NullableGridUpdate) Get() *GridUpdate {
	return v.value
}

func (v *NullableGridUpdate) Set(val *GridUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableGridUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableGridUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGridUpdate(val *GridUpdate) *NullableGridUpdate {
	return &NullableGridUpdate{value: val, isSet: true}
}

func (v NullableGridUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGridUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
