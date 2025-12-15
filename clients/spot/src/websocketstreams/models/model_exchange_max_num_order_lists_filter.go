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

// checks if the ExchangeMaxNumOrderListsFilter type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExchangeMaxNumOrderListsFilter{}

// ExchangeMaxNumOrderListsFilter struct for ExchangeMaxNumOrderListsFilter
type ExchangeMaxNumOrderListsFilter struct {
	FilterType           *string `json:"filterType,omitempty"`
	MaxNumOrderLists     *int64  `json:"maxNumOrderLists,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExchangeMaxNumOrderListsFilter ExchangeMaxNumOrderListsFilter

// NewExchangeMaxNumOrderListsFilter instantiates a new ExchangeMaxNumOrderListsFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeMaxNumOrderListsFilter() *ExchangeMaxNumOrderListsFilter {
	this := ExchangeMaxNumOrderListsFilter{}
	return &this
}

// NewExchangeMaxNumOrderListsFilterWithDefaults instantiates a new ExchangeMaxNumOrderListsFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeMaxNumOrderListsFilterWithDefaults() *ExchangeMaxNumOrderListsFilter {
	this := ExchangeMaxNumOrderListsFilter{}
	return &this
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *ExchangeMaxNumOrderListsFilter) GetFilterType() string {
	if o == nil || common.IsNil(o.FilterType) {
		var ret string
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeMaxNumOrderListsFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *ExchangeMaxNumOrderListsFilter) HasFilterType() bool {
	if o != nil && !common.IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given string and assigns it to the FilterType field.
func (o *ExchangeMaxNumOrderListsFilter) SetFilterType(v string) {
	o.FilterType = &v
}

// GetMaxNumOrderLists returns the MaxNumOrderLists field value if set, zero value otherwise.
func (o *ExchangeMaxNumOrderListsFilter) GetMaxNumOrderLists() int64 {
	if o == nil || common.IsNil(o.MaxNumOrderLists) {
		var ret int64
		return ret
	}
	return *o.MaxNumOrderLists
}

// GetMaxNumOrderListsOk returns a tuple with the MaxNumOrderLists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeMaxNumOrderListsFilter) GetMaxNumOrderListsOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MaxNumOrderLists) {
		return nil, false
	}
	return o.MaxNumOrderLists, true
}

// HasMaxNumOrderLists returns a boolean if a field has been set.
func (o *ExchangeMaxNumOrderListsFilter) HasMaxNumOrderLists() bool {
	if o != nil && !common.IsNil(o.MaxNumOrderLists) {
		return true
	}

	return false
}

// SetMaxNumOrderLists gets a reference to the given int64 and assigns it to the MaxNumOrderLists field.
func (o *ExchangeMaxNumOrderListsFilter) SetMaxNumOrderLists(v int64) {
	o.MaxNumOrderLists = &v
}

func (o ExchangeMaxNumOrderListsFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeMaxNumOrderListsFilter) ToMap() (map[string]interface{}, error) {
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

func (o *ExchangeMaxNumOrderListsFilter) UnmarshalJSON(data []byte) (err error) {
	varExchangeMaxNumOrderListsFilter := _ExchangeMaxNumOrderListsFilter{}

	err = json.Unmarshal(data, &varExchangeMaxNumOrderListsFilter)

	if err != nil {
		return err
	}

	*o = ExchangeMaxNumOrderListsFilter(varExchangeMaxNumOrderListsFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filterType")
		delete(additionalProperties, "maxNumOrderLists")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExchangeMaxNumOrderListsFilter struct {
	value *ExchangeMaxNumOrderListsFilter
	isSet bool
}

func (v NullableExchangeMaxNumOrderListsFilter) Get() *ExchangeMaxNumOrderListsFilter {
	return v.value
}

func (v *NullableExchangeMaxNumOrderListsFilter) Set(val *ExchangeMaxNumOrderListsFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeMaxNumOrderListsFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeMaxNumOrderListsFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeMaxNumOrderListsFilter(val *ExchangeMaxNumOrderListsFilter) *NullableExchangeMaxNumOrderListsFilter {
	return &NullableExchangeMaxNumOrderListsFilter{value: val, isSet: true}
}

func (v NullableExchangeMaxNumOrderListsFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeMaxNumOrderListsFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
