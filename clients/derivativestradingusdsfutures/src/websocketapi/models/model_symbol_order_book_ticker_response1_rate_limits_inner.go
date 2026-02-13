/*
Binance Derivatives Trading USDS Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the SymbolOrderBookTickerResponse1RateLimitsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SymbolOrderBookTickerResponse1RateLimitsInner{}

// SymbolOrderBookTickerResponse1RateLimitsInner struct for SymbolOrderBookTickerResponse1RateLimitsInner
type SymbolOrderBookTickerResponse1RateLimitsInner struct {
	RateLimitType        *string `json:"rateLimitType,omitempty"`
	Interval             *string `json:"interval,omitempty"`
	IntervalNum          *int64  `json:"intervalNum,omitempty"`
	Limit                *int64  `json:"limit,omitempty"`
	Count                *int64  `json:"count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SymbolOrderBookTickerResponse1RateLimitsInner SymbolOrderBookTickerResponse1RateLimitsInner

// NewSymbolOrderBookTickerResponse1RateLimitsInner instantiates a new SymbolOrderBookTickerResponse1RateLimitsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSymbolOrderBookTickerResponse1RateLimitsInner() *SymbolOrderBookTickerResponse1RateLimitsInner {
	this := SymbolOrderBookTickerResponse1RateLimitsInner{}
	return &this
}

// NewSymbolOrderBookTickerResponse1RateLimitsInnerWithDefaults instantiates a new SymbolOrderBookTickerResponse1RateLimitsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSymbolOrderBookTickerResponse1RateLimitsInnerWithDefaults() *SymbolOrderBookTickerResponse1RateLimitsInner {
	this := SymbolOrderBookTickerResponse1RateLimitsInner{}
	return &this
}

// GetRateLimitType returns the RateLimitType field value if set, zero value otherwise.
func (o *SymbolOrderBookTickerResponse1RateLimitsInner) GetRateLimitType() string {
	if o == nil || common.IsNil(o.RateLimitType) {
		var ret string
		return ret
	}
	return *o.RateLimitType
}

// GetRateLimitTypeOk returns a tuple with the RateLimitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolOrderBookTickerResponse1RateLimitsInner) GetRateLimitTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.RateLimitType) {
		return nil, false
	}
	return o.RateLimitType, true
}

// HasRateLimitType returns a boolean if a field has been set.
func (o *SymbolOrderBookTickerResponse1RateLimitsInner) HasRateLimitType() bool {
	if o != nil && !common.IsNil(o.RateLimitType) {
		return true
	}

	return false
}

// SetRateLimitType gets a reference to the given string and assigns it to the RateLimitType field.
func (o *SymbolOrderBookTickerResponse1RateLimitsInner) SetRateLimitType(v string) {
	o.RateLimitType = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *SymbolOrderBookTickerResponse1RateLimitsInner) GetInterval() string {
	if o == nil || common.IsNil(o.Interval) {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolOrderBookTickerResponse1RateLimitsInner) GetIntervalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *SymbolOrderBookTickerResponse1RateLimitsInner) HasInterval() bool {
	if o != nil && !common.IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *SymbolOrderBookTickerResponse1RateLimitsInner) SetInterval(v string) {
	o.Interval = &v
}

// GetIntervalNum returns the IntervalNum field value if set, zero value otherwise.
func (o *SymbolOrderBookTickerResponse1RateLimitsInner) GetIntervalNum() int64 {
	if o == nil || common.IsNil(o.IntervalNum) {
		var ret int64
		return ret
	}
	return *o.IntervalNum
}

// GetIntervalNumOk returns a tuple with the IntervalNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolOrderBookTickerResponse1RateLimitsInner) GetIntervalNumOk() (*int64, bool) {
	if o == nil || common.IsNil(o.IntervalNum) {
		return nil, false
	}
	return o.IntervalNum, true
}

// HasIntervalNum returns a boolean if a field has been set.
func (o *SymbolOrderBookTickerResponse1RateLimitsInner) HasIntervalNum() bool {
	if o != nil && !common.IsNil(o.IntervalNum) {
		return true
	}

	return false
}

// SetIntervalNum gets a reference to the given int64 and assigns it to the IntervalNum field.
func (o *SymbolOrderBookTickerResponse1RateLimitsInner) SetIntervalNum(v int64) {
	o.IntervalNum = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *SymbolOrderBookTickerResponse1RateLimitsInner) GetLimit() int64 {
	if o == nil || common.IsNil(o.Limit) {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolOrderBookTickerResponse1RateLimitsInner) GetLimitOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *SymbolOrderBookTickerResponse1RateLimitsInner) HasLimit() bool {
	if o != nil && !common.IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *SymbolOrderBookTickerResponse1RateLimitsInner) SetLimit(v int64) {
	o.Limit = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *SymbolOrderBookTickerResponse1RateLimitsInner) GetCount() int64 {
	if o == nil || common.IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolOrderBookTickerResponse1RateLimitsInner) GetCountOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *SymbolOrderBookTickerResponse1RateLimitsInner) HasCount() bool {
	if o != nil && !common.IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *SymbolOrderBookTickerResponse1RateLimitsInner) SetCount(v int64) {
	o.Count = &v
}

func (o SymbolOrderBookTickerResponse1RateLimitsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SymbolOrderBookTickerResponse1RateLimitsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.RateLimitType) {
		toSerialize["rateLimitType"] = o.RateLimitType
	}
	if !common.IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !common.IsNil(o.IntervalNum) {
		toSerialize["intervalNum"] = o.IntervalNum
	}
	if !common.IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !common.IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SymbolOrderBookTickerResponse1RateLimitsInner) UnmarshalJSON(data []byte) (err error) {
	varSymbolOrderBookTickerResponse1RateLimitsInner := _SymbolOrderBookTickerResponse1RateLimitsInner{}

	err = json.Unmarshal(data, &varSymbolOrderBookTickerResponse1RateLimitsInner)

	if err != nil {
		return err
	}

	*o = SymbolOrderBookTickerResponse1RateLimitsInner(varSymbolOrderBookTickerResponse1RateLimitsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rateLimitType")
		delete(additionalProperties, "interval")
		delete(additionalProperties, "intervalNum")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSymbolOrderBookTickerResponse1RateLimitsInner struct {
	value *SymbolOrderBookTickerResponse1RateLimitsInner
	isSet bool
}

func (v NullableSymbolOrderBookTickerResponse1RateLimitsInner) Get() *SymbolOrderBookTickerResponse1RateLimitsInner {
	return v.value
}

func (v *NullableSymbolOrderBookTickerResponse1RateLimitsInner) Set(val *SymbolOrderBookTickerResponse1RateLimitsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSymbolOrderBookTickerResponse1RateLimitsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSymbolOrderBookTickerResponse1RateLimitsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSymbolOrderBookTickerResponse1RateLimitsInner(val *SymbolOrderBookTickerResponse1RateLimitsInner) *NullableSymbolOrderBookTickerResponse1RateLimitsInner {
	return &NullableSymbolOrderBookTickerResponse1RateLimitsInner{value: val, isSet: true}
}

func (v NullableSymbolOrderBookTickerResponse1RateLimitsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSymbolOrderBookTickerResponse1RateLimitsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
