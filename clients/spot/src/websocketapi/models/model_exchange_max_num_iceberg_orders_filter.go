/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the ExchangeMaxNumIcebergOrdersFilter type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExchangeMaxNumIcebergOrdersFilter{}

// ExchangeMaxNumIcebergOrdersFilter struct for ExchangeMaxNumIcebergOrdersFilter
type ExchangeMaxNumIcebergOrdersFilter struct {
	FilterType           *string `json:"filterType,omitempty"`
	MaxNumIcebergOrders  *int64  `json:"maxNumIcebergOrders,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExchangeMaxNumIcebergOrdersFilter ExchangeMaxNumIcebergOrdersFilter

// NewExchangeMaxNumIcebergOrdersFilter instantiates a new ExchangeMaxNumIcebergOrdersFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeMaxNumIcebergOrdersFilter() *ExchangeMaxNumIcebergOrdersFilter {
	this := ExchangeMaxNumIcebergOrdersFilter{}
	return &this
}

// NewExchangeMaxNumIcebergOrdersFilterWithDefaults instantiates a new ExchangeMaxNumIcebergOrdersFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeMaxNumIcebergOrdersFilterWithDefaults() *ExchangeMaxNumIcebergOrdersFilter {
	this := ExchangeMaxNumIcebergOrdersFilter{}
	return &this
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *ExchangeMaxNumIcebergOrdersFilter) GetFilterType() string {
	if o == nil || common.IsNil(o.FilterType) {
		var ret string
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeMaxNumIcebergOrdersFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *ExchangeMaxNumIcebergOrdersFilter) HasFilterType() bool {
	if o != nil && !common.IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given string and assigns it to the FilterType field.
func (o *ExchangeMaxNumIcebergOrdersFilter) SetFilterType(v string) {
	o.FilterType = &v
}

// GetMaxNumIcebergOrders returns the MaxNumIcebergOrders field value if set, zero value otherwise.
func (o *ExchangeMaxNumIcebergOrdersFilter) GetMaxNumIcebergOrders() int64 {
	if o == nil || common.IsNil(o.MaxNumIcebergOrders) {
		var ret int64
		return ret
	}
	return *o.MaxNumIcebergOrders
}

// GetMaxNumIcebergOrdersOk returns a tuple with the MaxNumIcebergOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeMaxNumIcebergOrdersFilter) GetMaxNumIcebergOrdersOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MaxNumIcebergOrders) {
		return nil, false
	}
	return o.MaxNumIcebergOrders, true
}

// HasMaxNumIcebergOrders returns a boolean if a field has been set.
func (o *ExchangeMaxNumIcebergOrdersFilter) HasMaxNumIcebergOrders() bool {
	if o != nil && !common.IsNil(o.MaxNumIcebergOrders) {
		return true
	}

	return false
}

// SetMaxNumIcebergOrders gets a reference to the given int64 and assigns it to the MaxNumIcebergOrders field.
func (o *ExchangeMaxNumIcebergOrdersFilter) SetMaxNumIcebergOrders(v int64) {
	o.MaxNumIcebergOrders = &v
}

func (o ExchangeMaxNumIcebergOrdersFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeMaxNumIcebergOrdersFilter) ToMap() (map[string]interface{}, error) {
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

func (o *ExchangeMaxNumIcebergOrdersFilter) UnmarshalJSON(data []byte) (err error) {
	varExchangeMaxNumIcebergOrdersFilter := _ExchangeMaxNumIcebergOrdersFilter{}

	err = json.Unmarshal(data, &varExchangeMaxNumIcebergOrdersFilter)

	if err != nil {
		return err
	}

	*o = ExchangeMaxNumIcebergOrdersFilter(varExchangeMaxNumIcebergOrdersFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filterType")
		delete(additionalProperties, "maxNumIcebergOrders")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExchangeMaxNumIcebergOrdersFilter struct {
	value *ExchangeMaxNumIcebergOrdersFilter
	isSet bool
}

func (v NullableExchangeMaxNumIcebergOrdersFilter) Get() *ExchangeMaxNumIcebergOrdersFilter {
	return v.value
}

func (v *NullableExchangeMaxNumIcebergOrdersFilter) Set(val *ExchangeMaxNumIcebergOrdersFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeMaxNumIcebergOrdersFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeMaxNumIcebergOrdersFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeMaxNumIcebergOrdersFilter(val *ExchangeMaxNumIcebergOrdersFilter) *NullableExchangeMaxNumIcebergOrdersFilter {
	return &NullableExchangeMaxNumIcebergOrdersFilter{value: val, isSet: true}
}

func (v NullableExchangeMaxNumIcebergOrdersFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeMaxNumIcebergOrdersFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
