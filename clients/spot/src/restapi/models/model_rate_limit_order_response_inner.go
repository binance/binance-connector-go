/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the RateLimitOrderResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RateLimitOrderResponseInner{}

// RateLimitOrderResponseInner struct for RateLimitOrderResponseInner
type RateLimitOrderResponseInner struct {
	RateLimitType        *string `json:"rateLimitType,omitempty"`
	Interval             *string `json:"interval,omitempty"`
	IntervalNum          *int64  `json:"intervalNum,omitempty"`
	Limit                *int64  `json:"limit,omitempty"`
	Count                *int64  `json:"count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RateLimitOrderResponseInner RateLimitOrderResponseInner

// NewRateLimitOrderResponseInner instantiates a new RateLimitOrderResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateLimitOrderResponseInner() *RateLimitOrderResponseInner {
	this := RateLimitOrderResponseInner{}
	return &this
}

// NewRateLimitOrderResponseInnerWithDefaults instantiates a new RateLimitOrderResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateLimitOrderResponseInnerWithDefaults() *RateLimitOrderResponseInner {
	this := RateLimitOrderResponseInner{}
	return &this
}

// GetRateLimitType returns the RateLimitType field value if set, zero value otherwise.
func (o *RateLimitOrderResponseInner) GetRateLimitType() string {
	if o == nil || common.IsNil(o.RateLimitType) {
		var ret string
		return ret
	}
	return *o.RateLimitType
}

// GetRateLimitTypeOk returns a tuple with the RateLimitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimitOrderResponseInner) GetRateLimitTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.RateLimitType) {
		return nil, false
	}
	return o.RateLimitType, true
}

// HasRateLimitType returns a boolean if a field has been set.
func (o *RateLimitOrderResponseInner) HasRateLimitType() bool {
	if o != nil && !common.IsNil(o.RateLimitType) {
		return true
	}

	return false
}

// SetRateLimitType gets a reference to the given string and assigns it to the RateLimitType field.
func (o *RateLimitOrderResponseInner) SetRateLimitType(v string) {
	o.RateLimitType = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *RateLimitOrderResponseInner) GetInterval() string {
	if o == nil || common.IsNil(o.Interval) {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimitOrderResponseInner) GetIntervalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *RateLimitOrderResponseInner) HasInterval() bool {
	if o != nil && !common.IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *RateLimitOrderResponseInner) SetInterval(v string) {
	o.Interval = &v
}

// GetIntervalNum returns the IntervalNum field value if set, zero value otherwise.
func (o *RateLimitOrderResponseInner) GetIntervalNum() int64 {
	if o == nil || common.IsNil(o.IntervalNum) {
		var ret int64
		return ret
	}
	return *o.IntervalNum
}

// GetIntervalNumOk returns a tuple with the IntervalNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimitOrderResponseInner) GetIntervalNumOk() (*int64, bool) {
	if o == nil || common.IsNil(o.IntervalNum) {
		return nil, false
	}
	return o.IntervalNum, true
}

// HasIntervalNum returns a boolean if a field has been set.
func (o *RateLimitOrderResponseInner) HasIntervalNum() bool {
	if o != nil && !common.IsNil(o.IntervalNum) {
		return true
	}

	return false
}

// SetIntervalNum gets a reference to the given int64 and assigns it to the IntervalNum field.
func (o *RateLimitOrderResponseInner) SetIntervalNum(v int64) {
	o.IntervalNum = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *RateLimitOrderResponseInner) GetLimit() int64 {
	if o == nil || common.IsNil(o.Limit) {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimitOrderResponseInner) GetLimitOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *RateLimitOrderResponseInner) HasLimit() bool {
	if o != nil && !common.IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *RateLimitOrderResponseInner) SetLimit(v int64) {
	o.Limit = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *RateLimitOrderResponseInner) GetCount() int64 {
	if o == nil || common.IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimitOrderResponseInner) GetCountOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *RateLimitOrderResponseInner) HasCount() bool {
	if o != nil && !common.IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *RateLimitOrderResponseInner) SetCount(v int64) {
	o.Count = &v
}

func (o RateLimitOrderResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RateLimitOrderResponseInner) ToMap() (map[string]interface{}, error) {
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

func (o *RateLimitOrderResponseInner) UnmarshalJSON(data []byte) (err error) {
	varRateLimitOrderResponseInner := _RateLimitOrderResponseInner{}

	err = json.Unmarshal(data, &varRateLimitOrderResponseInner)

	if err != nil {
		return err
	}

	*o = RateLimitOrderResponseInner(varRateLimitOrderResponseInner)

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

type NullableRateLimitOrderResponseInner struct {
	value *RateLimitOrderResponseInner
	isSet bool
}

func (v NullableRateLimitOrderResponseInner) Get() *RateLimitOrderResponseInner {
	return v.value
}

func (v *NullableRateLimitOrderResponseInner) Set(val *RateLimitOrderResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRateLimitOrderResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRateLimitOrderResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateLimitOrderResponseInner(val *RateLimitOrderResponseInner) *NullableRateLimitOrderResponseInner {
	return &NullableRateLimitOrderResponseInner{value: val, isSet: true}
}

func (v NullableRateLimitOrderResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateLimitOrderResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
