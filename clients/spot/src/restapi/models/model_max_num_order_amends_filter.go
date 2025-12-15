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

// checks if the MaxNumOrderAmendsFilter type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MaxNumOrderAmendsFilter{}

// MaxNumOrderAmendsFilter struct for MaxNumOrderAmendsFilter
type MaxNumOrderAmendsFilter struct {
	FilterType           *string `json:"filterType,omitempty"`
	MaxNumOrderAmends    *int64  `json:"maxNumOrderAmends,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MaxNumOrderAmendsFilter MaxNumOrderAmendsFilter

// NewMaxNumOrderAmendsFilter instantiates a new MaxNumOrderAmendsFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaxNumOrderAmendsFilter() *MaxNumOrderAmendsFilter {
	this := MaxNumOrderAmendsFilter{}
	return &this
}

// NewMaxNumOrderAmendsFilterWithDefaults instantiates a new MaxNumOrderAmendsFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaxNumOrderAmendsFilterWithDefaults() *MaxNumOrderAmendsFilter {
	this := MaxNumOrderAmendsFilter{}
	return &this
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *MaxNumOrderAmendsFilter) GetFilterType() string {
	if o == nil || common.IsNil(o.FilterType) {
		var ret string
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaxNumOrderAmendsFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *MaxNumOrderAmendsFilter) HasFilterType() bool {
	if o != nil && !common.IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given string and assigns it to the FilterType field.
func (o *MaxNumOrderAmendsFilter) SetFilterType(v string) {
	o.FilterType = &v
}

// GetMaxNumOrderAmends returns the MaxNumOrderAmends field value if set, zero value otherwise.
func (o *MaxNumOrderAmendsFilter) GetMaxNumOrderAmends() int64 {
	if o == nil || common.IsNil(o.MaxNumOrderAmends) {
		var ret int64
		return ret
	}
	return *o.MaxNumOrderAmends
}

// GetMaxNumOrderAmendsOk returns a tuple with the MaxNumOrderAmends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaxNumOrderAmendsFilter) GetMaxNumOrderAmendsOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MaxNumOrderAmends) {
		return nil, false
	}
	return o.MaxNumOrderAmends, true
}

// HasMaxNumOrderAmends returns a boolean if a field has been set.
func (o *MaxNumOrderAmendsFilter) HasMaxNumOrderAmends() bool {
	if o != nil && !common.IsNil(o.MaxNumOrderAmends) {
		return true
	}

	return false
}

// SetMaxNumOrderAmends gets a reference to the given int64 and assigns it to the MaxNumOrderAmends field.
func (o *MaxNumOrderAmendsFilter) SetMaxNumOrderAmends(v int64) {
	o.MaxNumOrderAmends = &v
}

func (o MaxNumOrderAmendsFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaxNumOrderAmendsFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FilterType) {
		toSerialize["filterType"] = o.FilterType
	}
	if !common.IsNil(o.MaxNumOrderAmends) {
		toSerialize["maxNumOrderAmends"] = o.MaxNumOrderAmends
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MaxNumOrderAmendsFilter) UnmarshalJSON(data []byte) (err error) {
	varMaxNumOrderAmendsFilter := _MaxNumOrderAmendsFilter{}

	err = json.Unmarshal(data, &varMaxNumOrderAmendsFilter)

	if err != nil {
		return err
	}

	*o = MaxNumOrderAmendsFilter(varMaxNumOrderAmendsFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filterType")
		delete(additionalProperties, "maxNumOrderAmends")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMaxNumOrderAmendsFilter struct {
	value *MaxNumOrderAmendsFilter
	isSet bool
}

func (v NullableMaxNumOrderAmendsFilter) Get() *MaxNumOrderAmendsFilter {
	return v.value
}

func (v *NullableMaxNumOrderAmendsFilter) Set(val *MaxNumOrderAmendsFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableMaxNumOrderAmendsFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableMaxNumOrderAmendsFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaxNumOrderAmendsFilter(val *MaxNumOrderAmendsFilter) *NullableMaxNumOrderAmendsFilter {
	return &NullableMaxNumOrderAmendsFilter{value: val, isSet: true}
}

func (v NullableMaxNumOrderAmendsFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaxNumOrderAmendsFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
