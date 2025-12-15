/*
Binance Spot WebSocket Streams

OpenAPI Specifications for the Binance Spot WebSocket Streams  API documents:   - [Github web-socket-streams documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-streams.md)   - [General API information for web-socket-streams on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams)

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the TPlusSellFilter type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TPlusSellFilter{}

// TPlusSellFilter struct for TPlusSellFilter
type TPlusSellFilter struct {
	FilterType           *string `json:"filterType,omitempty"`
	EndTime              *int64  `json:"endTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TPlusSellFilter TPlusSellFilter

// NewTPlusSellFilter instantiates a new TPlusSellFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTPlusSellFilter() *TPlusSellFilter {
	this := TPlusSellFilter{}
	return &this
}

// NewTPlusSellFilterWithDefaults instantiates a new TPlusSellFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTPlusSellFilterWithDefaults() *TPlusSellFilter {
	this := TPlusSellFilter{}
	return &this
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *TPlusSellFilter) GetFilterType() string {
	if o == nil || common.IsNil(o.FilterType) {
		var ret string
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TPlusSellFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *TPlusSellFilter) HasFilterType() bool {
	if o != nil && !common.IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given string and assigns it to the FilterType field.
func (o *TPlusSellFilter) SetFilterType(v string) {
	o.FilterType = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *TPlusSellFilter) GetEndTime() int64 {
	if o == nil || common.IsNil(o.EndTime) {
		var ret int64
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TPlusSellFilter) GetEndTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *TPlusSellFilter) HasEndTime() bool {
	if o != nil && !common.IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given int64 and assigns it to the EndTime field.
func (o *TPlusSellFilter) SetEndTime(v int64) {
	o.EndTime = &v
}

func (o TPlusSellFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TPlusSellFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FilterType) {
		toSerialize["filterType"] = o.FilterType
	}
	if !common.IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TPlusSellFilter) UnmarshalJSON(data []byte) (err error) {
	varTPlusSellFilter := _TPlusSellFilter{}

	err = json.Unmarshal(data, &varTPlusSellFilter)

	if err != nil {
		return err
	}

	*o = TPlusSellFilter(varTPlusSellFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filterType")
		delete(additionalProperties, "endTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTPlusSellFilter struct {
	value *TPlusSellFilter
	isSet bool
}

func (v NullableTPlusSellFilter) Get() *TPlusSellFilter {
	return v.value
}

func (v *NullableTPlusSellFilter) Set(val *TPlusSellFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableTPlusSellFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableTPlusSellFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTPlusSellFilter(val *TPlusSellFilter) *NullableTPlusSellFilter {
	return &NullableTPlusSellFilter{value: val, isSet: true}
}

func (v NullableTPlusSellFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTPlusSellFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
