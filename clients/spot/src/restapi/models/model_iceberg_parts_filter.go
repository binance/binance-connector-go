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

// checks if the IcebergPartsFilter type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &IcebergPartsFilter{}

// IcebergPartsFilter struct for IcebergPartsFilter
type IcebergPartsFilter struct {
	FilterType           *string `json:"filterType,omitempty"`
	Limit                *int64  `json:"limit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IcebergPartsFilter IcebergPartsFilter

// NewIcebergPartsFilter instantiates a new IcebergPartsFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIcebergPartsFilter() *IcebergPartsFilter {
	this := IcebergPartsFilter{}
	return &this
}

// NewIcebergPartsFilterWithDefaults instantiates a new IcebergPartsFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIcebergPartsFilterWithDefaults() *IcebergPartsFilter {
	this := IcebergPartsFilter{}
	return &this
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *IcebergPartsFilter) GetFilterType() string {
	if o == nil || common.IsNil(o.FilterType) {
		var ret string
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IcebergPartsFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *IcebergPartsFilter) HasFilterType() bool {
	if o != nil && !common.IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given string and assigns it to the FilterType field.
func (o *IcebergPartsFilter) SetFilterType(v string) {
	o.FilterType = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *IcebergPartsFilter) GetLimit() int64 {
	if o == nil || common.IsNil(o.Limit) {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IcebergPartsFilter) GetLimitOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *IcebergPartsFilter) HasLimit() bool {
	if o != nil && !common.IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *IcebergPartsFilter) SetLimit(v int64) {
	o.Limit = &v
}

func (o IcebergPartsFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IcebergPartsFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FilterType) {
		toSerialize["filterType"] = o.FilterType
	}
	if !common.IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IcebergPartsFilter) UnmarshalJSON(data []byte) (err error) {
	varIcebergPartsFilter := _IcebergPartsFilter{}

	err = json.Unmarshal(data, &varIcebergPartsFilter)

	if err != nil {
		return err
	}

	*o = IcebergPartsFilter(varIcebergPartsFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filterType")
		delete(additionalProperties, "limit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIcebergPartsFilter struct {
	value *IcebergPartsFilter
	isSet bool
}

func (v NullableIcebergPartsFilter) Get() *IcebergPartsFilter {
	return v.value
}

func (v *NullableIcebergPartsFilter) Set(val *IcebergPartsFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableIcebergPartsFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableIcebergPartsFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIcebergPartsFilter(val *IcebergPartsFilter) *NullableIcebergPartsFilter {
	return &NullableIcebergPartsFilter{value: val, isSet: true}
}

func (v NullableIcebergPartsFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIcebergPartsFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
