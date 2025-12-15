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

// checks if the AlgoUpdate type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AlgoUpdate{}

// AlgoUpdate struct for AlgoUpdate
type AlgoUpdate struct {
	T                    *int64       `json:"T,omitempty"`
	E                    *int64       `json:"E,omitempty"`
	Smallo               *AlgoUpdateO `json:"o,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AlgoUpdate AlgoUpdate

// NewAlgoUpdate instantiates a new AlgoUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlgoUpdate() *AlgoUpdate {
	this := AlgoUpdate{}
	return &this
}

// NewAlgoUpdateWithDefaults instantiates a new AlgoUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlgoUpdateWithDefaults() *AlgoUpdate {
	this := AlgoUpdate{}
	return &this
}

// GetT returns the T field value if set, zero value otherwise.
func (o *AlgoUpdate) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoUpdate) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *AlgoUpdate) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *AlgoUpdate) SetT(v int64) {
	o.T = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AlgoUpdate) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoUpdate) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *AlgoUpdate) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *AlgoUpdate) SetE(v int64) {
	o.E = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *AlgoUpdate) GetSmallo() AlgoUpdateO {
	if o == nil || common.IsNil(o.Smallo) {
		var ret AlgoUpdateO
		return ret
	}
	return *o.Smallo
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoUpdate) GetSmalloOk() (*AlgoUpdateO, bool) {
	if o == nil || common.IsNil(o.Smallo) {
		return nil, false
	}
	return o.Smallo, true
}

// HasO returns a boolean if a field has been set.
func (o *AlgoUpdate) HasSmallo() bool {
	if o != nil && !common.IsNil(o.Smallo) {
		return true
	}

	return false
}

// SetO gets a reference to the given AlgoUpdateO and assigns it to the O field.
func (o *AlgoUpdate) SetSmallo(v AlgoUpdateO) {
	o.Smallo = &v
}

func (o AlgoUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlgoUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smallo) {
		toSerialize["o"] = o.Smallo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AlgoUpdate) UnmarshalJSON(data []byte) (err error) {
	varAlgoUpdate := _AlgoUpdate{}

	err = json.Unmarshal(data, &varAlgoUpdate)

	if err != nil {
		return err
	}

	*o = AlgoUpdate(varAlgoUpdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "T")
		delete(additionalProperties, "E")
		delete(additionalProperties, "o")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAlgoUpdate struct {
	value *AlgoUpdate
	isSet bool
}

func (v NullableAlgoUpdate) Get() *AlgoUpdate {
	return v.value
}

func (v *NullableAlgoUpdate) Set(val *AlgoUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAlgoUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAlgoUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlgoUpdate(val *AlgoUpdate) *NullableAlgoUpdate {
	return &NullableAlgoUpdate{value: val, isSet: true}
}

func (v NullableAlgoUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlgoUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
