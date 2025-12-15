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

// checks if the StrategyUpdate type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &StrategyUpdate{}

// StrategyUpdate struct for StrategyUpdate
type StrategyUpdate struct {
	T                    *int64            `json:"T,omitempty"`
	E                    *int64            `json:"E,omitempty"`
	Smallsu              *StrategyUpdateSu `json:"su,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StrategyUpdate StrategyUpdate

// NewStrategyUpdate instantiates a new StrategyUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStrategyUpdate() *StrategyUpdate {
	this := StrategyUpdate{}
	return &this
}

// NewStrategyUpdateWithDefaults instantiates a new StrategyUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStrategyUpdateWithDefaults() *StrategyUpdate {
	this := StrategyUpdate{}
	return &this
}

// GetT returns the T field value if set, zero value otherwise.
func (o *StrategyUpdate) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StrategyUpdate) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *StrategyUpdate) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *StrategyUpdate) SetT(v int64) {
	o.T = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *StrategyUpdate) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StrategyUpdate) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *StrategyUpdate) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *StrategyUpdate) SetE(v int64) {
	o.E = &v
}

// GetSu returns the Su field value if set, zero value otherwise.
func (o *StrategyUpdate) GetSmallsu() StrategyUpdateSu {
	if o == nil || common.IsNil(o.Smallsu) {
		var ret StrategyUpdateSu
		return ret
	}
	return *o.Smallsu
}

// GetSuOk returns a tuple with the Su field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StrategyUpdate) GetSmallsuOk() (*StrategyUpdateSu, bool) {
	if o == nil || common.IsNil(o.Smallsu) {
		return nil, false
	}
	return o.Smallsu, true
}

// HasSu returns a boolean if a field has been set.
func (o *StrategyUpdate) HasSmallsu() bool {
	if o != nil && !common.IsNil(o.Smallsu) {
		return true
	}

	return false
}

// SetSu gets a reference to the given StrategyUpdateSu and assigns it to the Su field.
func (o *StrategyUpdate) SetSmallsu(v StrategyUpdateSu) {
	o.Smallsu = &v
}

func (o StrategyUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StrategyUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smallsu) {
		toSerialize["su"] = o.Smallsu
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StrategyUpdate) UnmarshalJSON(data []byte) (err error) {
	varStrategyUpdate := _StrategyUpdate{}

	err = json.Unmarshal(data, &varStrategyUpdate)

	if err != nil {
		return err
	}

	*o = StrategyUpdate(varStrategyUpdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "T")
		delete(additionalProperties, "E")
		delete(additionalProperties, "su")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStrategyUpdate struct {
	value *StrategyUpdate
	isSet bool
}

func (v NullableStrategyUpdate) Get() *StrategyUpdate {
	return v.value
}

func (v *NullableStrategyUpdate) Set(val *StrategyUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableStrategyUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableStrategyUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStrategyUpdate(val *StrategyUpdate) *NullableStrategyUpdate {
	return &NullableStrategyUpdate{value: val, isSet: true}
}

func (v NullableStrategyUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStrategyUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
