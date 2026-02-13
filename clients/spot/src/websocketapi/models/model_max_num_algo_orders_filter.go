/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the MaxNumAlgoOrdersFilter type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MaxNumAlgoOrdersFilter{}

// MaxNumAlgoOrdersFilter struct for MaxNumAlgoOrdersFilter
type MaxNumAlgoOrdersFilter struct {
	FilterType           *string `json:"filterType,omitempty"`
	MaxNumAlgoOrders     *int64  `json:"maxNumAlgoOrders,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MaxNumAlgoOrdersFilter MaxNumAlgoOrdersFilter

// NewMaxNumAlgoOrdersFilter instantiates a new MaxNumAlgoOrdersFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaxNumAlgoOrdersFilter() *MaxNumAlgoOrdersFilter {
	this := MaxNumAlgoOrdersFilter{}
	return &this
}

// NewMaxNumAlgoOrdersFilterWithDefaults instantiates a new MaxNumAlgoOrdersFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaxNumAlgoOrdersFilterWithDefaults() *MaxNumAlgoOrdersFilter {
	this := MaxNumAlgoOrdersFilter{}
	return &this
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *MaxNumAlgoOrdersFilter) GetFilterType() string {
	if o == nil || common.IsNil(o.FilterType) {
		var ret string
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaxNumAlgoOrdersFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *MaxNumAlgoOrdersFilter) HasFilterType() bool {
	if o != nil && !common.IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given string and assigns it to the FilterType field.
func (o *MaxNumAlgoOrdersFilter) SetFilterType(v string) {
	o.FilterType = &v
}

// GetMaxNumAlgoOrders returns the MaxNumAlgoOrders field value if set, zero value otherwise.
func (o *MaxNumAlgoOrdersFilter) GetMaxNumAlgoOrders() int64 {
	if o == nil || common.IsNil(o.MaxNumAlgoOrders) {
		var ret int64
		return ret
	}
	return *o.MaxNumAlgoOrders
}

// GetMaxNumAlgoOrdersOk returns a tuple with the MaxNumAlgoOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaxNumAlgoOrdersFilter) GetMaxNumAlgoOrdersOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MaxNumAlgoOrders) {
		return nil, false
	}
	return o.MaxNumAlgoOrders, true
}

// HasMaxNumAlgoOrders returns a boolean if a field has been set.
func (o *MaxNumAlgoOrdersFilter) HasMaxNumAlgoOrders() bool {
	if o != nil && !common.IsNil(o.MaxNumAlgoOrders) {
		return true
	}

	return false
}

// SetMaxNumAlgoOrders gets a reference to the given int64 and assigns it to the MaxNumAlgoOrders field.
func (o *MaxNumAlgoOrdersFilter) SetMaxNumAlgoOrders(v int64) {
	o.MaxNumAlgoOrders = &v
}

func (o MaxNumAlgoOrdersFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaxNumAlgoOrdersFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FilterType) {
		toSerialize["filterType"] = o.FilterType
	}
	if !common.IsNil(o.MaxNumAlgoOrders) {
		toSerialize["maxNumAlgoOrders"] = o.MaxNumAlgoOrders
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MaxNumAlgoOrdersFilter) UnmarshalJSON(data []byte) (err error) {
	varMaxNumAlgoOrdersFilter := _MaxNumAlgoOrdersFilter{}

	err = json.Unmarshal(data, &varMaxNumAlgoOrdersFilter)

	if err != nil {
		return err
	}

	*o = MaxNumAlgoOrdersFilter(varMaxNumAlgoOrdersFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filterType")
		delete(additionalProperties, "maxNumAlgoOrders")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMaxNumAlgoOrdersFilter struct {
	value *MaxNumAlgoOrdersFilter
	isSet bool
}

func (v NullableMaxNumAlgoOrdersFilter) Get() *MaxNumAlgoOrdersFilter {
	return v.value
}

func (v *NullableMaxNumAlgoOrdersFilter) Set(val *MaxNumAlgoOrdersFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableMaxNumAlgoOrdersFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableMaxNumAlgoOrdersFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaxNumAlgoOrdersFilter(val *MaxNumAlgoOrdersFilter) *NullableMaxNumAlgoOrdersFilter {
	return &NullableMaxNumAlgoOrdersFilter{value: val, isSet: true}
}

func (v NullableMaxNumAlgoOrdersFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaxNumAlgoOrdersFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
