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

// checks if the MaxNumIcebergOrdersFilter type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MaxNumIcebergOrdersFilter{}

// MaxNumIcebergOrdersFilter struct for MaxNumIcebergOrdersFilter
type MaxNumIcebergOrdersFilter struct {
	FilterType           *string `json:"filterType,omitempty"`
	MaxNumIcebergOrders  *int64  `json:"maxNumIcebergOrders,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MaxNumIcebergOrdersFilter MaxNumIcebergOrdersFilter

// NewMaxNumIcebergOrdersFilter instantiates a new MaxNumIcebergOrdersFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaxNumIcebergOrdersFilter() *MaxNumIcebergOrdersFilter {
	this := MaxNumIcebergOrdersFilter{}
	return &this
}

// NewMaxNumIcebergOrdersFilterWithDefaults instantiates a new MaxNumIcebergOrdersFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaxNumIcebergOrdersFilterWithDefaults() *MaxNumIcebergOrdersFilter {
	this := MaxNumIcebergOrdersFilter{}
	return &this
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *MaxNumIcebergOrdersFilter) GetFilterType() string {
	if o == nil || common.IsNil(o.FilterType) {
		var ret string
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaxNumIcebergOrdersFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *MaxNumIcebergOrdersFilter) HasFilterType() bool {
	if o != nil && !common.IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given string and assigns it to the FilterType field.
func (o *MaxNumIcebergOrdersFilter) SetFilterType(v string) {
	o.FilterType = &v
}

// GetMaxNumIcebergOrders returns the MaxNumIcebergOrders field value if set, zero value otherwise.
func (o *MaxNumIcebergOrdersFilter) GetMaxNumIcebergOrders() int64 {
	if o == nil || common.IsNil(o.MaxNumIcebergOrders) {
		var ret int64
		return ret
	}
	return *o.MaxNumIcebergOrders
}

// GetMaxNumIcebergOrdersOk returns a tuple with the MaxNumIcebergOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaxNumIcebergOrdersFilter) GetMaxNumIcebergOrdersOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MaxNumIcebergOrders) {
		return nil, false
	}
	return o.MaxNumIcebergOrders, true
}

// HasMaxNumIcebergOrders returns a boolean if a field has been set.
func (o *MaxNumIcebergOrdersFilter) HasMaxNumIcebergOrders() bool {
	if o != nil && !common.IsNil(o.MaxNumIcebergOrders) {
		return true
	}

	return false
}

// SetMaxNumIcebergOrders gets a reference to the given int64 and assigns it to the MaxNumIcebergOrders field.
func (o *MaxNumIcebergOrdersFilter) SetMaxNumIcebergOrders(v int64) {
	o.MaxNumIcebergOrders = &v
}

func (o MaxNumIcebergOrdersFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaxNumIcebergOrdersFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FilterType) {
		toSerialize["filterType"] = o.FilterType
	}
	if !common.IsNil(o.MaxNumIcebergOrders) {
		toSerialize["maxNumIcebergOrders"] = o.MaxNumIcebergOrders
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MaxNumIcebergOrdersFilter) UnmarshalJSON(data []byte) (err error) {
	varMaxNumIcebergOrdersFilter := _MaxNumIcebergOrdersFilter{}

	err = json.Unmarshal(data, &varMaxNumIcebergOrdersFilter)

	if err != nil {
		return err
	}

	*o = MaxNumIcebergOrdersFilter(varMaxNumIcebergOrdersFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filterType")
		delete(additionalProperties, "maxNumIcebergOrders")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMaxNumIcebergOrdersFilter struct {
	value *MaxNumIcebergOrdersFilter
	isSet bool
}

func (v NullableMaxNumIcebergOrdersFilter) Get() *MaxNumIcebergOrdersFilter {
	return v.value
}

func (v *NullableMaxNumIcebergOrdersFilter) Set(val *MaxNumIcebergOrdersFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableMaxNumIcebergOrdersFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableMaxNumIcebergOrdersFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaxNumIcebergOrdersFilter(val *MaxNumIcebergOrdersFilter) *NullableMaxNumIcebergOrdersFilter {
	return &NullableMaxNumIcebergOrdersFilter{value: val, isSet: true}
}

func (v NullableMaxNumIcebergOrdersFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaxNumIcebergOrdersFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
