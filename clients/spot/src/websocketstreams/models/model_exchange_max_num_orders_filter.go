/*
Binance Spot WebSocket Streams

OpenAPI Specifications for the Binance Spot WebSocket Streams  API documents:   - [Github web-socket-streams documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-streams.md)   - [General API information for web-socket-streams on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the ExchangeMaxNumOrdersFilter type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExchangeMaxNumOrdersFilter{}

// ExchangeMaxNumOrdersFilter struct for ExchangeMaxNumOrdersFilter
type ExchangeMaxNumOrdersFilter struct {
	FilterType           *string `json:"filterType,omitempty"`
	MaxNumOrders         *int64  `json:"maxNumOrders,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExchangeMaxNumOrdersFilter ExchangeMaxNumOrdersFilter

// NewExchangeMaxNumOrdersFilter instantiates a new ExchangeMaxNumOrdersFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeMaxNumOrdersFilter() *ExchangeMaxNumOrdersFilter {
	this := ExchangeMaxNumOrdersFilter{}
	return &this
}

// NewExchangeMaxNumOrdersFilterWithDefaults instantiates a new ExchangeMaxNumOrdersFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeMaxNumOrdersFilterWithDefaults() *ExchangeMaxNumOrdersFilter {
	this := ExchangeMaxNumOrdersFilter{}
	return &this
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *ExchangeMaxNumOrdersFilter) GetFilterType() string {
	if o == nil || common.IsNil(o.FilterType) {
		var ret string
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeMaxNumOrdersFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *ExchangeMaxNumOrdersFilter) HasFilterType() bool {
	if o != nil && !common.IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given string and assigns it to the FilterType field.
func (o *ExchangeMaxNumOrdersFilter) SetFilterType(v string) {
	o.FilterType = &v
}

// GetMaxNumOrders returns the MaxNumOrders field value if set, zero value otherwise.
func (o *ExchangeMaxNumOrdersFilter) GetMaxNumOrders() int64 {
	if o == nil || common.IsNil(o.MaxNumOrders) {
		var ret int64
		return ret
	}
	return *o.MaxNumOrders
}

// GetMaxNumOrdersOk returns a tuple with the MaxNumOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeMaxNumOrdersFilter) GetMaxNumOrdersOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MaxNumOrders) {
		return nil, false
	}
	return o.MaxNumOrders, true
}

// HasMaxNumOrders returns a boolean if a field has been set.
func (o *ExchangeMaxNumOrdersFilter) HasMaxNumOrders() bool {
	if o != nil && !common.IsNil(o.MaxNumOrders) {
		return true
	}

	return false
}

// SetMaxNumOrders gets a reference to the given int64 and assigns it to the MaxNumOrders field.
func (o *ExchangeMaxNumOrdersFilter) SetMaxNumOrders(v int64) {
	o.MaxNumOrders = &v
}

func (o ExchangeMaxNumOrdersFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeMaxNumOrdersFilter) ToMap() (map[string]interface{}, error) {
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

func (o *ExchangeMaxNumOrdersFilter) UnmarshalJSON(data []byte) (err error) {
	varExchangeMaxNumOrdersFilter := _ExchangeMaxNumOrdersFilter{}

	err = json.Unmarshal(data, &varExchangeMaxNumOrdersFilter)

	if err != nil {
		return err
	}

	*o = ExchangeMaxNumOrdersFilter(varExchangeMaxNumOrdersFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filterType")
		delete(additionalProperties, "maxNumOrders")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExchangeMaxNumOrdersFilter struct {
	value *ExchangeMaxNumOrdersFilter
	isSet bool
}

func (v NullableExchangeMaxNumOrdersFilter) Get() *ExchangeMaxNumOrdersFilter {
	return v.value
}

func (v *NullableExchangeMaxNumOrdersFilter) Set(val *ExchangeMaxNumOrdersFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeMaxNumOrdersFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeMaxNumOrdersFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeMaxNumOrdersFilter(val *ExchangeMaxNumOrdersFilter) *NullableExchangeMaxNumOrdersFilter {
	return &NullableExchangeMaxNumOrdersFilter{value: val, isSet: true}
}

func (v NullableExchangeMaxNumOrdersFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeMaxNumOrdersFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
