/*
Binance Spot WebSocket Streams

OpenAPI Specifications for the Binance Spot WebSocket Streams  API documents:   - [Github web-socket-streams documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-streams.md)   - [General API information for web-socket-streams on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ExchangeMaxNumAlgoOrdersFilter type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExchangeMaxNumAlgoOrdersFilter{}

// ExchangeMaxNumAlgoOrdersFilter struct for ExchangeMaxNumAlgoOrdersFilter
type ExchangeMaxNumAlgoOrdersFilter struct {
	FilterType           *string `json:"filterType,omitempty"`
	MaxNumAlgoOrders     *int64  `json:"maxNumAlgoOrders,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExchangeMaxNumAlgoOrdersFilter ExchangeMaxNumAlgoOrdersFilter

// NewExchangeMaxNumAlgoOrdersFilter instantiates a new ExchangeMaxNumAlgoOrdersFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeMaxNumAlgoOrdersFilter() *ExchangeMaxNumAlgoOrdersFilter {
	this := ExchangeMaxNumAlgoOrdersFilter{}
	return &this
}

// NewExchangeMaxNumAlgoOrdersFilterWithDefaults instantiates a new ExchangeMaxNumAlgoOrdersFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeMaxNumAlgoOrdersFilterWithDefaults() *ExchangeMaxNumAlgoOrdersFilter {
	this := ExchangeMaxNumAlgoOrdersFilter{}
	return &this
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *ExchangeMaxNumAlgoOrdersFilter) GetFilterType() string {
	if o == nil || common.IsNil(o.FilterType) {
		var ret string
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeMaxNumAlgoOrdersFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *ExchangeMaxNumAlgoOrdersFilter) HasFilterType() bool {
	if o != nil && !common.IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given string and assigns it to the FilterType field.
func (o *ExchangeMaxNumAlgoOrdersFilter) SetFilterType(v string) {
	o.FilterType = &v
}

// GetMaxNumAlgoOrders returns the MaxNumAlgoOrders field value if set, zero value otherwise.
func (o *ExchangeMaxNumAlgoOrdersFilter) GetMaxNumAlgoOrders() int64 {
	if o == nil || common.IsNil(o.MaxNumAlgoOrders) {
		var ret int64
		return ret
	}
	return *o.MaxNumAlgoOrders
}

// GetMaxNumAlgoOrdersOk returns a tuple with the MaxNumAlgoOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeMaxNumAlgoOrdersFilter) GetMaxNumAlgoOrdersOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MaxNumAlgoOrders) {
		return nil, false
	}
	return o.MaxNumAlgoOrders, true
}

// HasMaxNumAlgoOrders returns a boolean if a field has been set.
func (o *ExchangeMaxNumAlgoOrdersFilter) HasMaxNumAlgoOrders() bool {
	if o != nil && !common.IsNil(o.MaxNumAlgoOrders) {
		return true
	}

	return false
}

// SetMaxNumAlgoOrders gets a reference to the given int64 and assigns it to the MaxNumAlgoOrders field.
func (o *ExchangeMaxNumAlgoOrdersFilter) SetMaxNumAlgoOrders(v int64) {
	o.MaxNumAlgoOrders = &v
}

func (o ExchangeMaxNumAlgoOrdersFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeMaxNumAlgoOrdersFilter) ToMap() (map[string]interface{}, error) {
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

func (o *ExchangeMaxNumAlgoOrdersFilter) UnmarshalJSON(data []byte) (err error) {
	varExchangeMaxNumAlgoOrdersFilter := _ExchangeMaxNumAlgoOrdersFilter{}

	err = json.Unmarshal(data, &varExchangeMaxNumAlgoOrdersFilter)

	if err != nil {
		return err
	}

	*o = ExchangeMaxNumAlgoOrdersFilter(varExchangeMaxNumAlgoOrdersFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filterType")
		delete(additionalProperties, "maxNumAlgoOrders")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExchangeMaxNumAlgoOrdersFilter struct {
	value *ExchangeMaxNumAlgoOrdersFilter
	isSet bool
}

func (v NullableExchangeMaxNumAlgoOrdersFilter) Get() *ExchangeMaxNumAlgoOrdersFilter {
	return v.value
}

func (v *NullableExchangeMaxNumAlgoOrdersFilter) Set(val *ExchangeMaxNumAlgoOrdersFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeMaxNumAlgoOrdersFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeMaxNumAlgoOrdersFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeMaxNumAlgoOrdersFilter(val *ExchangeMaxNumAlgoOrdersFilter) *NullableExchangeMaxNumAlgoOrdersFilter {
	return &NullableExchangeMaxNumAlgoOrdersFilter{value: val, isSet: true}
}

func (v NullableExchangeMaxNumAlgoOrdersFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeMaxNumAlgoOrdersFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
