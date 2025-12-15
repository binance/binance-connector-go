/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MaxNumOrderListsFilter type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MaxNumOrderListsFilter{}

// MaxNumOrderListsFilter struct for MaxNumOrderListsFilter
type MaxNumOrderListsFilter struct {
	FilterType           *string `json:"filterType,omitempty"`
	MaxNumOrderLists     *int64  `json:"maxNumOrderLists,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MaxNumOrderListsFilter MaxNumOrderListsFilter

// NewMaxNumOrderListsFilter instantiates a new MaxNumOrderListsFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaxNumOrderListsFilter() *MaxNumOrderListsFilter {
	this := MaxNumOrderListsFilter{}
	return &this
}

// NewMaxNumOrderListsFilterWithDefaults instantiates a new MaxNumOrderListsFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaxNumOrderListsFilterWithDefaults() *MaxNumOrderListsFilter {
	this := MaxNumOrderListsFilter{}
	return &this
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *MaxNumOrderListsFilter) GetFilterType() string {
	if o == nil || common.IsNil(o.FilterType) {
		var ret string
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaxNumOrderListsFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *MaxNumOrderListsFilter) HasFilterType() bool {
	if o != nil && !common.IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given string and assigns it to the FilterType field.
func (o *MaxNumOrderListsFilter) SetFilterType(v string) {
	o.FilterType = &v
}

// GetMaxNumOrderLists returns the MaxNumOrderLists field value if set, zero value otherwise.
func (o *MaxNumOrderListsFilter) GetMaxNumOrderLists() int64 {
	if o == nil || common.IsNil(o.MaxNumOrderLists) {
		var ret int64
		return ret
	}
	return *o.MaxNumOrderLists
}

// GetMaxNumOrderListsOk returns a tuple with the MaxNumOrderLists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaxNumOrderListsFilter) GetMaxNumOrderListsOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MaxNumOrderLists) {
		return nil, false
	}
	return o.MaxNumOrderLists, true
}

// HasMaxNumOrderLists returns a boolean if a field has been set.
func (o *MaxNumOrderListsFilter) HasMaxNumOrderLists() bool {
	if o != nil && !common.IsNil(o.MaxNumOrderLists) {
		return true
	}

	return false
}

// SetMaxNumOrderLists gets a reference to the given int64 and assigns it to the MaxNumOrderLists field.
func (o *MaxNumOrderListsFilter) SetMaxNumOrderLists(v int64) {
	o.MaxNumOrderLists = &v
}

func (o MaxNumOrderListsFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaxNumOrderListsFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FilterType) {
		toSerialize["filterType"] = o.FilterType
	}
	if !common.IsNil(o.MaxNumOrderLists) {
		toSerialize["maxNumOrderLists"] = o.MaxNumOrderLists
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MaxNumOrderListsFilter) UnmarshalJSON(data []byte) (err error) {
	varMaxNumOrderListsFilter := _MaxNumOrderListsFilter{}

	err = json.Unmarshal(data, &varMaxNumOrderListsFilter)

	if err != nil {
		return err
	}

	*o = MaxNumOrderListsFilter(varMaxNumOrderListsFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filterType")
		delete(additionalProperties, "maxNumOrderLists")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMaxNumOrderListsFilter struct {
	value *MaxNumOrderListsFilter
	isSet bool
}

func (v NullableMaxNumOrderListsFilter) Get() *MaxNumOrderListsFilter {
	return v.value
}

func (v *NullableMaxNumOrderListsFilter) Set(val *MaxNumOrderListsFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableMaxNumOrderListsFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableMaxNumOrderListsFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaxNumOrderListsFilter(val *MaxNumOrderListsFilter) *NullableMaxNumOrderListsFilter {
	return &NullableMaxNumOrderListsFilter{value: val, isSet: true}
}

func (v NullableMaxNumOrderListsFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaxNumOrderListsFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
