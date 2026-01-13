/*
Binance Derivatives Trading COIN Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the StartUserDataStreamResponseRateLimitsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &StartUserDataStreamResponseRateLimitsInner{}

// StartUserDataStreamResponseRateLimitsInner struct for StartUserDataStreamResponseRateLimitsInner
type StartUserDataStreamResponseRateLimitsInner struct {
	RateLimitType        *string `json:"rateLimitType,omitempty"`
	Interval             *string `json:"interval,omitempty"`
	IntervalNum          *int64  `json:"intervalNum,omitempty"`
	Limit                *int64  `json:"limit,omitempty"`
	Count                *int64  `json:"count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StartUserDataStreamResponseRateLimitsInner StartUserDataStreamResponseRateLimitsInner

// NewStartUserDataStreamResponseRateLimitsInner instantiates a new StartUserDataStreamResponseRateLimitsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartUserDataStreamResponseRateLimitsInner() *StartUserDataStreamResponseRateLimitsInner {
	this := StartUserDataStreamResponseRateLimitsInner{}
	return &this
}

// NewStartUserDataStreamResponseRateLimitsInnerWithDefaults instantiates a new StartUserDataStreamResponseRateLimitsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartUserDataStreamResponseRateLimitsInnerWithDefaults() *StartUserDataStreamResponseRateLimitsInner {
	this := StartUserDataStreamResponseRateLimitsInner{}
	return &this
}

// GetRateLimitType returns the RateLimitType field value if set, zero value otherwise.
func (o *StartUserDataStreamResponseRateLimitsInner) GetRateLimitType() string {
	if o == nil || common.IsNil(o.RateLimitType) {
		var ret string
		return ret
	}
	return *o.RateLimitType
}

// GetRateLimitTypeOk returns a tuple with the RateLimitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartUserDataStreamResponseRateLimitsInner) GetRateLimitTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.RateLimitType) {
		return nil, false
	}
	return o.RateLimitType, true
}

// HasRateLimitType returns a boolean if a field has been set.
func (o *StartUserDataStreamResponseRateLimitsInner) HasRateLimitType() bool {
	if o != nil && !common.IsNil(o.RateLimitType) {
		return true
	}

	return false
}

// SetRateLimitType gets a reference to the given string and assigns it to the RateLimitType field.
func (o *StartUserDataStreamResponseRateLimitsInner) SetRateLimitType(v string) {
	o.RateLimitType = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *StartUserDataStreamResponseRateLimitsInner) GetInterval() string {
	if o == nil || common.IsNil(o.Interval) {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartUserDataStreamResponseRateLimitsInner) GetIntervalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *StartUserDataStreamResponseRateLimitsInner) HasInterval() bool {
	if o != nil && !common.IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *StartUserDataStreamResponseRateLimitsInner) SetInterval(v string) {
	o.Interval = &v
}

// GetIntervalNum returns the IntervalNum field value if set, zero value otherwise.
func (o *StartUserDataStreamResponseRateLimitsInner) GetIntervalNum() int64 {
	if o == nil || common.IsNil(o.IntervalNum) {
		var ret int64
		return ret
	}
	return *o.IntervalNum
}

// GetIntervalNumOk returns a tuple with the IntervalNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartUserDataStreamResponseRateLimitsInner) GetIntervalNumOk() (*int64, bool) {
	if o == nil || common.IsNil(o.IntervalNum) {
		return nil, false
	}
	return o.IntervalNum, true
}

// HasIntervalNum returns a boolean if a field has been set.
func (o *StartUserDataStreamResponseRateLimitsInner) HasIntervalNum() bool {
	if o != nil && !common.IsNil(o.IntervalNum) {
		return true
	}

	return false
}

// SetIntervalNum gets a reference to the given int64 and assigns it to the IntervalNum field.
func (o *StartUserDataStreamResponseRateLimitsInner) SetIntervalNum(v int64) {
	o.IntervalNum = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *StartUserDataStreamResponseRateLimitsInner) GetLimit() int64 {
	if o == nil || common.IsNil(o.Limit) {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartUserDataStreamResponseRateLimitsInner) GetLimitOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *StartUserDataStreamResponseRateLimitsInner) HasLimit() bool {
	if o != nil && !common.IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *StartUserDataStreamResponseRateLimitsInner) SetLimit(v int64) {
	o.Limit = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *StartUserDataStreamResponseRateLimitsInner) GetCount() int64 {
	if o == nil || common.IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartUserDataStreamResponseRateLimitsInner) GetCountOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *StartUserDataStreamResponseRateLimitsInner) HasCount() bool {
	if o != nil && !common.IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *StartUserDataStreamResponseRateLimitsInner) SetCount(v int64) {
	o.Count = &v
}

func (o StartUserDataStreamResponseRateLimitsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartUserDataStreamResponseRateLimitsInner) ToMap() (map[string]interface{}, error) {
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

func (o *StartUserDataStreamResponseRateLimitsInner) UnmarshalJSON(data []byte) (err error) {
	varStartUserDataStreamResponseRateLimitsInner := _StartUserDataStreamResponseRateLimitsInner{}

	err = json.Unmarshal(data, &varStartUserDataStreamResponseRateLimitsInner)

	if err != nil {
		return err
	}

	*o = StartUserDataStreamResponseRateLimitsInner(varStartUserDataStreamResponseRateLimitsInner)

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

type NullableStartUserDataStreamResponseRateLimitsInner struct {
	value *StartUserDataStreamResponseRateLimitsInner
	isSet bool
}

func (v NullableStartUserDataStreamResponseRateLimitsInner) Get() *StartUserDataStreamResponseRateLimitsInner {
	return v.value
}

func (v *NullableStartUserDataStreamResponseRateLimitsInner) Set(val *StartUserDataStreamResponseRateLimitsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableStartUserDataStreamResponseRateLimitsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableStartUserDataStreamResponseRateLimitsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartUserDataStreamResponseRateLimitsInner(val *StartUserDataStreamResponseRateLimitsInner) *NullableStartUserDataStreamResponseRateLimitsInner {
	return &NullableStartUserDataStreamResponseRateLimitsInner{value: val, isSet: true}
}

func (v NullableStartUserDataStreamResponseRateLimitsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartUserDataStreamResponseRateLimitsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
