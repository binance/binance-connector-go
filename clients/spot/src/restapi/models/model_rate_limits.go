/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the RateLimits type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RateLimits{}

// RateLimits struct for RateLimits
type RateLimits struct {
	RateLimitType        *string `json:"rateLimitType,omitempty"`
	Interval             *string `json:"interval,omitempty"`
	IntervalNum          *int64  `json:"intervalNum,omitempty"`
	Limit                *int64  `json:"limit,omitempty"`
	Count                *int64  `json:"count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RateLimits RateLimits

// NewRateLimits instantiates a new RateLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateLimits() *RateLimits {
	this := RateLimits{}
	return &this
}

// NewRateLimitsWithDefaults instantiates a new RateLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateLimitsWithDefaults() *RateLimits {
	this := RateLimits{}
	return &this
}

// GetRateLimitType returns the RateLimitType field value if set, zero value otherwise.
func (o *RateLimits) GetRateLimitType() string {
	if o == nil || common.IsNil(o.RateLimitType) {
		var ret string
		return ret
	}
	return *o.RateLimitType
}

// GetRateLimitTypeOk returns a tuple with the RateLimitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimits) GetRateLimitTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.RateLimitType) {
		return nil, false
	}
	return o.RateLimitType, true
}

// HasRateLimitType returns a boolean if a field has been set.
func (o *RateLimits) HasRateLimitType() bool {
	if o != nil && !common.IsNil(o.RateLimitType) {
		return true
	}

	return false
}

// SetRateLimitType gets a reference to the given string and assigns it to the RateLimitType field.
func (o *RateLimits) SetRateLimitType(v string) {
	o.RateLimitType = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *RateLimits) GetInterval() string {
	if o == nil || common.IsNil(o.Interval) {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimits) GetIntervalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *RateLimits) HasInterval() bool {
	if o != nil && !common.IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *RateLimits) SetInterval(v string) {
	o.Interval = &v
}

// GetIntervalNum returns the IntervalNum field value if set, zero value otherwise.
func (o *RateLimits) GetIntervalNum() int64 {
	if o == nil || common.IsNil(o.IntervalNum) {
		var ret int64
		return ret
	}
	return *o.IntervalNum
}

// GetIntervalNumOk returns a tuple with the IntervalNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimits) GetIntervalNumOk() (*int64, bool) {
	if o == nil || common.IsNil(o.IntervalNum) {
		return nil, false
	}
	return o.IntervalNum, true
}

// HasIntervalNum returns a boolean if a field has been set.
func (o *RateLimits) HasIntervalNum() bool {
	if o != nil && !common.IsNil(o.IntervalNum) {
		return true
	}

	return false
}

// SetIntervalNum gets a reference to the given int64 and assigns it to the IntervalNum field.
func (o *RateLimits) SetIntervalNum(v int64) {
	o.IntervalNum = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *RateLimits) GetLimit() int64 {
	if o == nil || common.IsNil(o.Limit) {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimits) GetLimitOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *RateLimits) HasLimit() bool {
	if o != nil && !common.IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *RateLimits) SetLimit(v int64) {
	o.Limit = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *RateLimits) GetCount() int64 {
	if o == nil || common.IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimits) GetCountOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *RateLimits) HasCount() bool {
	if o != nil && !common.IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *RateLimits) SetCount(v int64) {
	o.Count = &v
}

func (o RateLimits) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RateLimits) ToMap() (map[string]interface{}, error) {
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

func (o *RateLimits) UnmarshalJSON(data []byte) (err error) {
	varRateLimits := _RateLimits{}

	err = json.Unmarshal(data, &varRateLimits)

	if err != nil {
		return err
	}

	*o = RateLimits(varRateLimits)

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

type NullableRateLimits struct {
	value *RateLimits
	isSet bool
}

func (v NullableRateLimits) Get() *RateLimits {
	return v.value
}

func (v *NullableRateLimits) Set(val *RateLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableRateLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableRateLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateLimits(val *RateLimits) *NullableRateLimits {
	return &NullableRateLimits{value: val, isSet: true}
}

func (v NullableRateLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
