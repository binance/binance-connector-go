/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryCurrentMarginOrderCountUsageResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryCurrentMarginOrderCountUsageResponseInner{}

// QueryCurrentMarginOrderCountUsageResponseInner struct for QueryCurrentMarginOrderCountUsageResponseInner
type QueryCurrentMarginOrderCountUsageResponseInner struct {
	RateLimitType        *string `json:"rateLimitType,omitempty"`
	Interval             *string `json:"interval,omitempty"`
	IntervalNum          *int64  `json:"intervalNum,omitempty"`
	Limit                *int64  `json:"limit,omitempty"`
	Count                *int64  `json:"count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryCurrentMarginOrderCountUsageResponseInner QueryCurrentMarginOrderCountUsageResponseInner

// NewQueryCurrentMarginOrderCountUsageResponseInner instantiates a new QueryCurrentMarginOrderCountUsageResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryCurrentMarginOrderCountUsageResponseInner() *QueryCurrentMarginOrderCountUsageResponseInner {
	this := QueryCurrentMarginOrderCountUsageResponseInner{}
	return &this
}

// NewQueryCurrentMarginOrderCountUsageResponseInnerWithDefaults instantiates a new QueryCurrentMarginOrderCountUsageResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryCurrentMarginOrderCountUsageResponseInnerWithDefaults() *QueryCurrentMarginOrderCountUsageResponseInner {
	this := QueryCurrentMarginOrderCountUsageResponseInner{}
	return &this
}

// GetRateLimitType returns the RateLimitType field value if set, zero value otherwise.
func (o *QueryCurrentMarginOrderCountUsageResponseInner) GetRateLimitType() string {
	if o == nil || common.IsNil(o.RateLimitType) {
		var ret string
		return ret
	}
	return *o.RateLimitType
}

// GetRateLimitTypeOk returns a tuple with the RateLimitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentMarginOrderCountUsageResponseInner) GetRateLimitTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.RateLimitType) {
		return nil, false
	}
	return o.RateLimitType, true
}

// HasRateLimitType returns a boolean if a field has been set.
func (o *QueryCurrentMarginOrderCountUsageResponseInner) HasRateLimitType() bool {
	if o != nil && !common.IsNil(o.RateLimitType) {
		return true
	}

	return false
}

// SetRateLimitType gets a reference to the given string and assigns it to the RateLimitType field.
func (o *QueryCurrentMarginOrderCountUsageResponseInner) SetRateLimitType(v string) {
	o.RateLimitType = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *QueryCurrentMarginOrderCountUsageResponseInner) GetInterval() string {
	if o == nil || common.IsNil(o.Interval) {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentMarginOrderCountUsageResponseInner) GetIntervalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *QueryCurrentMarginOrderCountUsageResponseInner) HasInterval() bool {
	if o != nil && !common.IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *QueryCurrentMarginOrderCountUsageResponseInner) SetInterval(v string) {
	o.Interval = &v
}

// GetIntervalNum returns the IntervalNum field value if set, zero value otherwise.
func (o *QueryCurrentMarginOrderCountUsageResponseInner) GetIntervalNum() int64 {
	if o == nil || common.IsNil(o.IntervalNum) {
		var ret int64
		return ret
	}
	return *o.IntervalNum
}

// GetIntervalNumOk returns a tuple with the IntervalNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentMarginOrderCountUsageResponseInner) GetIntervalNumOk() (*int64, bool) {
	if o == nil || common.IsNil(o.IntervalNum) {
		return nil, false
	}
	return o.IntervalNum, true
}

// HasIntervalNum returns a boolean if a field has been set.
func (o *QueryCurrentMarginOrderCountUsageResponseInner) HasIntervalNum() bool {
	if o != nil && !common.IsNil(o.IntervalNum) {
		return true
	}

	return false
}

// SetIntervalNum gets a reference to the given int64 and assigns it to the IntervalNum field.
func (o *QueryCurrentMarginOrderCountUsageResponseInner) SetIntervalNum(v int64) {
	o.IntervalNum = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *QueryCurrentMarginOrderCountUsageResponseInner) GetLimit() int64 {
	if o == nil || common.IsNil(o.Limit) {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentMarginOrderCountUsageResponseInner) GetLimitOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *QueryCurrentMarginOrderCountUsageResponseInner) HasLimit() bool {
	if o != nil && !common.IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *QueryCurrentMarginOrderCountUsageResponseInner) SetLimit(v int64) {
	o.Limit = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *QueryCurrentMarginOrderCountUsageResponseInner) GetCount() int64 {
	if o == nil || common.IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentMarginOrderCountUsageResponseInner) GetCountOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *QueryCurrentMarginOrderCountUsageResponseInner) HasCount() bool {
	if o != nil && !common.IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *QueryCurrentMarginOrderCountUsageResponseInner) SetCount(v int64) {
	o.Count = &v
}

func (o QueryCurrentMarginOrderCountUsageResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryCurrentMarginOrderCountUsageResponseInner) ToMap() (map[string]interface{}, error) {
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

func (o *QueryCurrentMarginOrderCountUsageResponseInner) UnmarshalJSON(data []byte) (err error) {
	varQueryCurrentMarginOrderCountUsageResponseInner := _QueryCurrentMarginOrderCountUsageResponseInner{}

	err = json.Unmarshal(data, &varQueryCurrentMarginOrderCountUsageResponseInner)

	if err != nil {
		return err
	}

	*o = QueryCurrentMarginOrderCountUsageResponseInner(varQueryCurrentMarginOrderCountUsageResponseInner)

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

type NullableQueryCurrentMarginOrderCountUsageResponseInner struct {
	value *QueryCurrentMarginOrderCountUsageResponseInner
	isSet bool
}

func (v NullableQueryCurrentMarginOrderCountUsageResponseInner) Get() *QueryCurrentMarginOrderCountUsageResponseInner {
	return v.value
}

func (v *NullableQueryCurrentMarginOrderCountUsageResponseInner) Set(val *QueryCurrentMarginOrderCountUsageResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryCurrentMarginOrderCountUsageResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryCurrentMarginOrderCountUsageResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryCurrentMarginOrderCountUsageResponseInner(val *QueryCurrentMarginOrderCountUsageResponseInner) *NullableQueryCurrentMarginOrderCountUsageResponseInner {
	return &NullableQueryCurrentMarginOrderCountUsageResponseInner{value: val, isSet: true}
}

func (v NullableQueryCurrentMarginOrderCountUsageResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryCurrentMarginOrderCountUsageResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
