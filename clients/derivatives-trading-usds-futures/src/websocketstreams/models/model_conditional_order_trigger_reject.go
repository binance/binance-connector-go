/*
Binance Derivatives Trading USDS Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket Market Streams

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the ConditionalOrderTriggerReject type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ConditionalOrderTriggerReject{}

// ConditionalOrderTriggerReject struct for ConditionalOrderTriggerReject
type ConditionalOrderTriggerReject struct {
	E                    *int64                           `json:"E,omitempty"`
	T                    *int64                           `json:"T,omitempty"`
	Smallor              *ConditionalOrderTriggerRejectOr `json:"or,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConditionalOrderTriggerReject ConditionalOrderTriggerReject

// NewConditionalOrderTriggerReject instantiates a new ConditionalOrderTriggerReject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionalOrderTriggerReject() *ConditionalOrderTriggerReject {
	this := ConditionalOrderTriggerReject{}
	return &this
}

// NewConditionalOrderTriggerRejectWithDefaults instantiates a new ConditionalOrderTriggerReject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionalOrderTriggerRejectWithDefaults() *ConditionalOrderTriggerReject {
	this := ConditionalOrderTriggerReject{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *ConditionalOrderTriggerReject) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalOrderTriggerReject) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *ConditionalOrderTriggerReject) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *ConditionalOrderTriggerReject) SetE(v int64) {
	o.E = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *ConditionalOrderTriggerReject) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalOrderTriggerReject) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *ConditionalOrderTriggerReject) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *ConditionalOrderTriggerReject) SetT(v int64) {
	o.T = &v
}

// GetOr returns the Or field value if set, zero value otherwise.
func (o *ConditionalOrderTriggerReject) GetSmallor() ConditionalOrderTriggerRejectOr {
	if o == nil || common.IsNil(o.Smallor) {
		var ret ConditionalOrderTriggerRejectOr
		return ret
	}
	return *o.Smallor
}

// GetOrOk returns a tuple with the Or field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalOrderTriggerReject) GetSmallorOk() (*ConditionalOrderTriggerRejectOr, bool) {
	if o == nil || common.IsNil(o.Smallor) {
		return nil, false
	}
	return o.Smallor, true
}

// HasOr returns a boolean if a field has been set.
func (o *ConditionalOrderTriggerReject) HasSmallor() bool {
	if o != nil && !common.IsNil(o.Smallor) {
		return true
	}

	return false
}

// SetOr gets a reference to the given ConditionalOrderTriggerRejectOr and assigns it to the Or field.
func (o *ConditionalOrderTriggerReject) SetSmallor(v ConditionalOrderTriggerRejectOr) {
	o.Smallor = &v
}

func (o ConditionalOrderTriggerReject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConditionalOrderTriggerReject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.Smallor) {
		toSerialize["or"] = o.Smallor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConditionalOrderTriggerReject) UnmarshalJSON(data []byte) (err error) {
	varConditionalOrderTriggerReject := _ConditionalOrderTriggerReject{}

	err = json.Unmarshal(data, &varConditionalOrderTriggerReject)

	if err != nil {
		return err
	}

	*o = ConditionalOrderTriggerReject(varConditionalOrderTriggerReject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "E")
		delete(additionalProperties, "T")
		delete(additionalProperties, "or")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConditionalOrderTriggerReject struct {
	value *ConditionalOrderTriggerReject
	isSet bool
}

func (v NullableConditionalOrderTriggerReject) Get() *ConditionalOrderTriggerReject {
	return v.value
}

func (v *NullableConditionalOrderTriggerReject) Set(val *ConditionalOrderTriggerReject) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionalOrderTriggerReject) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionalOrderTriggerReject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionalOrderTriggerReject(val *ConditionalOrderTriggerReject) *NullableConditionalOrderTriggerReject {
	return &NullableConditionalOrderTriggerReject{value: val, isSet: true}
}

func (v NullableConditionalOrderTriggerReject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionalOrderTriggerReject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
