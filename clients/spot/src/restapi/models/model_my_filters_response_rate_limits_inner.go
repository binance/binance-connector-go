/*
Spot REST API

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the MyFiltersResponseRateLimitsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MyFiltersResponseRateLimitsInner{}

// MyFiltersResponseRateLimitsInner struct for MyFiltersResponseRateLimitsInner
type MyFiltersResponseRateLimitsInner struct {
	RateLimitType        *string `json:"rateLimitType,omitempty"`
	Interval             *string `json:"interval,omitempty"`
	IntervalNum          *int64  `json:"intervalNum,omitempty"`
	Limit                *int64  `json:"limit,omitempty"`
	Count                *int64  `json:"count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MyFiltersResponseRateLimitsInner MyFiltersResponseRateLimitsInner

// NewMyFiltersResponseRateLimitsInner instantiates a new MyFiltersResponseRateLimitsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMyFiltersResponseRateLimitsInner() *MyFiltersResponseRateLimitsInner {
	this := MyFiltersResponseRateLimitsInner{}
	return &this
}

// NewMyFiltersResponseRateLimitsInnerWithDefaults instantiates a new MyFiltersResponseRateLimitsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMyFiltersResponseRateLimitsInnerWithDefaults() *MyFiltersResponseRateLimitsInner {
	this := MyFiltersResponseRateLimitsInner{}
	return &this
}

// GetRateLimitType returns the RateLimitType field value if set, zero value otherwise.
func (o *MyFiltersResponseRateLimitsInner) GetRateLimitType() string {
	if o == nil || common.IsNil(o.RateLimitType) {
		var ret string
		return ret
	}
	return *o.RateLimitType
}

// GetRateLimitTypeOk returns a tuple with the RateLimitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyFiltersResponseRateLimitsInner) GetRateLimitTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.RateLimitType) {
		return nil, false
	}
	return o.RateLimitType, true
}

// HasRateLimitType returns a boolean if a field has been set.
func (o *MyFiltersResponseRateLimitsInner) HasRateLimitType() bool {
	if o != nil && !common.IsNil(o.RateLimitType) {
		return true
	}

	return false
}

// SetRateLimitType gets a reference to the given string and assigns it to the RateLimitType field.
func (o *MyFiltersResponseRateLimitsInner) SetRateLimitType(v string) {
	o.RateLimitType = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *MyFiltersResponseRateLimitsInner) GetInterval() string {
	if o == nil || common.IsNil(o.Interval) {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyFiltersResponseRateLimitsInner) GetIntervalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *MyFiltersResponseRateLimitsInner) HasInterval() bool {
	if o != nil && !common.IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *MyFiltersResponseRateLimitsInner) SetInterval(v string) {
	o.Interval = &v
}

// GetIntervalNum returns the IntervalNum field value if set, zero value otherwise.
func (o *MyFiltersResponseRateLimitsInner) GetIntervalNum() int64 {
	if o == nil || common.IsNil(o.IntervalNum) {
		var ret int64
		return ret
	}
	return *o.IntervalNum
}

// GetIntervalNumOk returns a tuple with the IntervalNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyFiltersResponseRateLimitsInner) GetIntervalNumOk() (*int64, bool) {
	if o == nil || common.IsNil(o.IntervalNum) {
		return nil, false
	}
	return o.IntervalNum, true
}

// HasIntervalNum returns a boolean if a field has been set.
func (o *MyFiltersResponseRateLimitsInner) HasIntervalNum() bool {
	if o != nil && !common.IsNil(o.IntervalNum) {
		return true
	}

	return false
}

// SetIntervalNum gets a reference to the given int64 and assigns it to the IntervalNum field.
func (o *MyFiltersResponseRateLimitsInner) SetIntervalNum(v int64) {
	o.IntervalNum = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *MyFiltersResponseRateLimitsInner) GetLimit() int64 {
	if o == nil || common.IsNil(o.Limit) {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyFiltersResponseRateLimitsInner) GetLimitOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *MyFiltersResponseRateLimitsInner) HasLimit() bool {
	if o != nil && !common.IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *MyFiltersResponseRateLimitsInner) SetLimit(v int64) {
	o.Limit = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *MyFiltersResponseRateLimitsInner) GetCount() int64 {
	if o == nil || common.IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyFiltersResponseRateLimitsInner) GetCountOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *MyFiltersResponseRateLimitsInner) HasCount() bool {
	if o != nil && !common.IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *MyFiltersResponseRateLimitsInner) SetCount(v int64) {
	o.Count = &v
}

func (o MyFiltersResponseRateLimitsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MyFiltersResponseRateLimitsInner) ToMap() (map[string]interface{}, error) {
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

func (o *MyFiltersResponseRateLimitsInner) UnmarshalJSON(data []byte) (err error) {
	varMyFiltersResponseRateLimitsInner := _MyFiltersResponseRateLimitsInner{}

	err = json.Unmarshal(data, &varMyFiltersResponseRateLimitsInner)

	if err != nil {
		return err
	}

	*o = MyFiltersResponseRateLimitsInner(varMyFiltersResponseRateLimitsInner)

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

type NullableMyFiltersResponseRateLimitsInner struct {
	value *MyFiltersResponseRateLimitsInner
	isSet bool
}

func (v NullableMyFiltersResponseRateLimitsInner) Get() *MyFiltersResponseRateLimitsInner {
	return v.value
}

func (v *NullableMyFiltersResponseRateLimitsInner) Set(val *MyFiltersResponseRateLimitsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMyFiltersResponseRateLimitsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMyFiltersResponseRateLimitsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMyFiltersResponseRateLimitsInner(val *MyFiltersResponseRateLimitsInner) *NullableMyFiltersResponseRateLimitsInner {
	return &NullableMyFiltersResponseRateLimitsInner{value: val, isSet: true}
}

func (v NullableMyFiltersResponseRateLimitsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMyFiltersResponseRateLimitsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
