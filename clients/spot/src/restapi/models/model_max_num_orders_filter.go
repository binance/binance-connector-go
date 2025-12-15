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

// checks if the MaxNumOrdersFilter type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MaxNumOrdersFilter{}

// MaxNumOrdersFilter struct for MaxNumOrdersFilter
type MaxNumOrdersFilter struct {
	FilterType           *string `json:"filterType,omitempty"`
	MaxNumOrders         *int64  `json:"maxNumOrders,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MaxNumOrdersFilter MaxNumOrdersFilter

// NewMaxNumOrdersFilter instantiates a new MaxNumOrdersFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaxNumOrdersFilter() *MaxNumOrdersFilter {
	this := MaxNumOrdersFilter{}
	return &this
}

// NewMaxNumOrdersFilterWithDefaults instantiates a new MaxNumOrdersFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaxNumOrdersFilterWithDefaults() *MaxNumOrdersFilter {
	this := MaxNumOrdersFilter{}
	return &this
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *MaxNumOrdersFilter) GetFilterType() string {
	if o == nil || common.IsNil(o.FilterType) {
		var ret string
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaxNumOrdersFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *MaxNumOrdersFilter) HasFilterType() bool {
	if o != nil && !common.IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given string and assigns it to the FilterType field.
func (o *MaxNumOrdersFilter) SetFilterType(v string) {
	o.FilterType = &v
}

// GetMaxNumOrders returns the MaxNumOrders field value if set, zero value otherwise.
func (o *MaxNumOrdersFilter) GetMaxNumOrders() int64 {
	if o == nil || common.IsNil(o.MaxNumOrders) {
		var ret int64
		return ret
	}
	return *o.MaxNumOrders
}

// GetMaxNumOrdersOk returns a tuple with the MaxNumOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaxNumOrdersFilter) GetMaxNumOrdersOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MaxNumOrders) {
		return nil, false
	}
	return o.MaxNumOrders, true
}

// HasMaxNumOrders returns a boolean if a field has been set.
func (o *MaxNumOrdersFilter) HasMaxNumOrders() bool {
	if o != nil && !common.IsNil(o.MaxNumOrders) {
		return true
	}

	return false
}

// SetMaxNumOrders gets a reference to the given int64 and assigns it to the MaxNumOrders field.
func (o *MaxNumOrdersFilter) SetMaxNumOrders(v int64) {
	o.MaxNumOrders = &v
}

func (o MaxNumOrdersFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaxNumOrdersFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FilterType) {
		toSerialize["filterType"] = o.FilterType
	}
	if !common.IsNil(o.MaxNumOrders) {
		toSerialize["maxNumOrders"] = o.MaxNumOrders
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MaxNumOrdersFilter) UnmarshalJSON(data []byte) (err error) {
	varMaxNumOrdersFilter := _MaxNumOrdersFilter{}

	err = json.Unmarshal(data, &varMaxNumOrdersFilter)

	if err != nil {
		return err
	}

	*o = MaxNumOrdersFilter(varMaxNumOrdersFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filterType")
		delete(additionalProperties, "maxNumOrders")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMaxNumOrdersFilter struct {
	value *MaxNumOrdersFilter
	isSet bool
}

func (v NullableMaxNumOrdersFilter) Get() *MaxNumOrdersFilter {
	return v.value
}

func (v *NullableMaxNumOrdersFilter) Set(val *MaxNumOrdersFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableMaxNumOrdersFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableMaxNumOrdersFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaxNumOrdersFilter(val *MaxNumOrdersFilter) *NullableMaxNumOrdersFilter {
	return &NullableMaxNumOrdersFilter{value: val, isSet: true}
}

func (v NullableMaxNumOrdersFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaxNumOrdersFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
