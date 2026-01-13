/*
Binance Spot WebSocket Streams

OpenAPI Specifications for the Binance Spot WebSocket Streams  API documents:   - [Github web-socket-streams documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-streams.md)   - [General API information for web-socket-streams on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the TrailingDeltaFilter type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TrailingDeltaFilter{}

// TrailingDeltaFilter struct for TrailingDeltaFilter
type TrailingDeltaFilter struct {
	FilterType            *string `json:"filterType,omitempty"`
	MinTrailingAboveDelta *int64  `json:"minTrailingAboveDelta,omitempty"`
	MaxTrailingAboveDelta *int64  `json:"maxTrailingAboveDelta,omitempty"`
	MinTrailingBelowDelta *int64  `json:"minTrailingBelowDelta,omitempty"`
	MaxTrailingBelowDelta *int64  `json:"maxTrailingBelowDelta,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _TrailingDeltaFilter TrailingDeltaFilter

// NewTrailingDeltaFilter instantiates a new TrailingDeltaFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrailingDeltaFilter() *TrailingDeltaFilter {
	this := TrailingDeltaFilter{}
	return &this
}

// NewTrailingDeltaFilterWithDefaults instantiates a new TrailingDeltaFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrailingDeltaFilterWithDefaults() *TrailingDeltaFilter {
	this := TrailingDeltaFilter{}
	return &this
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *TrailingDeltaFilter) GetFilterType() string {
	if o == nil || common.IsNil(o.FilterType) {
		var ret string
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrailingDeltaFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *TrailingDeltaFilter) HasFilterType() bool {
	if o != nil && !common.IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given string and assigns it to the FilterType field.
func (o *TrailingDeltaFilter) SetFilterType(v string) {
	o.FilterType = &v
}

// GetMinTrailingAboveDelta returns the MinTrailingAboveDelta field value if set, zero value otherwise.
func (o *TrailingDeltaFilter) GetMinTrailingAboveDelta() int64 {
	if o == nil || common.IsNil(o.MinTrailingAboveDelta) {
		var ret int64
		return ret
	}
	return *o.MinTrailingAboveDelta
}

// GetMinTrailingAboveDeltaOk returns a tuple with the MinTrailingAboveDelta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrailingDeltaFilter) GetMinTrailingAboveDeltaOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MinTrailingAboveDelta) {
		return nil, false
	}
	return o.MinTrailingAboveDelta, true
}

// HasMinTrailingAboveDelta returns a boolean if a field has been set.
func (o *TrailingDeltaFilter) HasMinTrailingAboveDelta() bool {
	if o != nil && !common.IsNil(o.MinTrailingAboveDelta) {
		return true
	}

	return false
}

// SetMinTrailingAboveDelta gets a reference to the given int64 and assigns it to the MinTrailingAboveDelta field.
func (o *TrailingDeltaFilter) SetMinTrailingAboveDelta(v int64) {
	o.MinTrailingAboveDelta = &v
}

// GetMaxTrailingAboveDelta returns the MaxTrailingAboveDelta field value if set, zero value otherwise.
func (o *TrailingDeltaFilter) GetMaxTrailingAboveDelta() int64 {
	if o == nil || common.IsNil(o.MaxTrailingAboveDelta) {
		var ret int64
		return ret
	}
	return *o.MaxTrailingAboveDelta
}

// GetMaxTrailingAboveDeltaOk returns a tuple with the MaxTrailingAboveDelta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrailingDeltaFilter) GetMaxTrailingAboveDeltaOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MaxTrailingAboveDelta) {
		return nil, false
	}
	return o.MaxTrailingAboveDelta, true
}

// HasMaxTrailingAboveDelta returns a boolean if a field has been set.
func (o *TrailingDeltaFilter) HasMaxTrailingAboveDelta() bool {
	if o != nil && !common.IsNil(o.MaxTrailingAboveDelta) {
		return true
	}

	return false
}

// SetMaxTrailingAboveDelta gets a reference to the given int64 and assigns it to the MaxTrailingAboveDelta field.
func (o *TrailingDeltaFilter) SetMaxTrailingAboveDelta(v int64) {
	o.MaxTrailingAboveDelta = &v
}

// GetMinTrailingBelowDelta returns the MinTrailingBelowDelta field value if set, zero value otherwise.
func (o *TrailingDeltaFilter) GetMinTrailingBelowDelta() int64 {
	if o == nil || common.IsNil(o.MinTrailingBelowDelta) {
		var ret int64
		return ret
	}
	return *o.MinTrailingBelowDelta
}

// GetMinTrailingBelowDeltaOk returns a tuple with the MinTrailingBelowDelta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrailingDeltaFilter) GetMinTrailingBelowDeltaOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MinTrailingBelowDelta) {
		return nil, false
	}
	return o.MinTrailingBelowDelta, true
}

// HasMinTrailingBelowDelta returns a boolean if a field has been set.
func (o *TrailingDeltaFilter) HasMinTrailingBelowDelta() bool {
	if o != nil && !common.IsNil(o.MinTrailingBelowDelta) {
		return true
	}

	return false
}

// SetMinTrailingBelowDelta gets a reference to the given int64 and assigns it to the MinTrailingBelowDelta field.
func (o *TrailingDeltaFilter) SetMinTrailingBelowDelta(v int64) {
	o.MinTrailingBelowDelta = &v
}

// GetMaxTrailingBelowDelta returns the MaxTrailingBelowDelta field value if set, zero value otherwise.
func (o *TrailingDeltaFilter) GetMaxTrailingBelowDelta() int64 {
	if o == nil || common.IsNil(o.MaxTrailingBelowDelta) {
		var ret int64
		return ret
	}
	return *o.MaxTrailingBelowDelta
}

// GetMaxTrailingBelowDeltaOk returns a tuple with the MaxTrailingBelowDelta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrailingDeltaFilter) GetMaxTrailingBelowDeltaOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MaxTrailingBelowDelta) {
		return nil, false
	}
	return o.MaxTrailingBelowDelta, true
}

// HasMaxTrailingBelowDelta returns a boolean if a field has been set.
func (o *TrailingDeltaFilter) HasMaxTrailingBelowDelta() bool {
	if o != nil && !common.IsNil(o.MaxTrailingBelowDelta) {
		return true
	}

	return false
}

// SetMaxTrailingBelowDelta gets a reference to the given int64 and assigns it to the MaxTrailingBelowDelta field.
func (o *TrailingDeltaFilter) SetMaxTrailingBelowDelta(v int64) {
	o.MaxTrailingBelowDelta = &v
}

func (o TrailingDeltaFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrailingDeltaFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FilterType) {
		toSerialize["filterType"] = o.FilterType
	}
	if !common.IsNil(o.MinTrailingAboveDelta) {
		toSerialize["minTrailingAboveDelta"] = o.MinTrailingAboveDelta
	}
	if !common.IsNil(o.MaxTrailingAboveDelta) {
		toSerialize["maxTrailingAboveDelta"] = o.MaxTrailingAboveDelta
	}
	if !common.IsNil(o.MinTrailingBelowDelta) {
		toSerialize["minTrailingBelowDelta"] = o.MinTrailingBelowDelta
	}
	if !common.IsNil(o.MaxTrailingBelowDelta) {
		toSerialize["maxTrailingBelowDelta"] = o.MaxTrailingBelowDelta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TrailingDeltaFilter) UnmarshalJSON(data []byte) (err error) {
	varTrailingDeltaFilter := _TrailingDeltaFilter{}

	err = json.Unmarshal(data, &varTrailingDeltaFilter)

	if err != nil {
		return err
	}

	*o = TrailingDeltaFilter(varTrailingDeltaFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filterType")
		delete(additionalProperties, "minTrailingAboveDelta")
		delete(additionalProperties, "maxTrailingAboveDelta")
		delete(additionalProperties, "minTrailingBelowDelta")
		delete(additionalProperties, "maxTrailingBelowDelta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTrailingDeltaFilter struct {
	value *TrailingDeltaFilter
	isSet bool
}

func (v NullableTrailingDeltaFilter) Get() *TrailingDeltaFilter {
	return v.value
}

func (v *NullableTrailingDeltaFilter) Set(val *TrailingDeltaFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableTrailingDeltaFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableTrailingDeltaFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrailingDeltaFilter(val *TrailingDeltaFilter) *NullableTrailingDeltaFilter {
	return &NullableTrailingDeltaFilter{value: val, isSet: true}
}

func (v NullableTrailingDeltaFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrailingDeltaFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
