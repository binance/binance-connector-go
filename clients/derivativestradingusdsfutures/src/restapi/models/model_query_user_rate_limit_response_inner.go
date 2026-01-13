/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryUserRateLimitResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryUserRateLimitResponseInner{}

// QueryUserRateLimitResponseInner struct for QueryUserRateLimitResponseInner
type QueryUserRateLimitResponseInner struct {
	RateLimitType        *string `json:"rateLimitType,omitempty"`
	Interval             *string `json:"interval,omitempty"`
	IntervalNum          *int64  `json:"intervalNum,omitempty"`
	Limit                *int64  `json:"limit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryUserRateLimitResponseInner QueryUserRateLimitResponseInner

// NewQueryUserRateLimitResponseInner instantiates a new QueryUserRateLimitResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryUserRateLimitResponseInner() *QueryUserRateLimitResponseInner {
	this := QueryUserRateLimitResponseInner{}
	return &this
}

// NewQueryUserRateLimitResponseInnerWithDefaults instantiates a new QueryUserRateLimitResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryUserRateLimitResponseInnerWithDefaults() *QueryUserRateLimitResponseInner {
	this := QueryUserRateLimitResponseInner{}
	return &this
}

// GetRateLimitType returns the RateLimitType field value if set, zero value otherwise.
func (o *QueryUserRateLimitResponseInner) GetRateLimitType() string {
	if o == nil || common.IsNil(o.RateLimitType) {
		var ret string
		return ret
	}
	return *o.RateLimitType
}

// GetRateLimitTypeOk returns a tuple with the RateLimitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserRateLimitResponseInner) GetRateLimitTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.RateLimitType) {
		return nil, false
	}
	return o.RateLimitType, true
}

// HasRateLimitType returns a boolean if a field has been set.
func (o *QueryUserRateLimitResponseInner) HasRateLimitType() bool {
	if o != nil && !common.IsNil(o.RateLimitType) {
		return true
	}

	return false
}

// SetRateLimitType gets a reference to the given string and assigns it to the RateLimitType field.
func (o *QueryUserRateLimitResponseInner) SetRateLimitType(v string) {
	o.RateLimitType = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *QueryUserRateLimitResponseInner) GetInterval() string {
	if o == nil || common.IsNil(o.Interval) {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserRateLimitResponseInner) GetIntervalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *QueryUserRateLimitResponseInner) HasInterval() bool {
	if o != nil && !common.IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *QueryUserRateLimitResponseInner) SetInterval(v string) {
	o.Interval = &v
}

// GetIntervalNum returns the IntervalNum field value if set, zero value otherwise.
func (o *QueryUserRateLimitResponseInner) GetIntervalNum() int64 {
	if o == nil || common.IsNil(o.IntervalNum) {
		var ret int64
		return ret
	}
	return *o.IntervalNum
}

// GetIntervalNumOk returns a tuple with the IntervalNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserRateLimitResponseInner) GetIntervalNumOk() (*int64, bool) {
	if o == nil || common.IsNil(o.IntervalNum) {
		return nil, false
	}
	return o.IntervalNum, true
}

// HasIntervalNum returns a boolean if a field has been set.
func (o *QueryUserRateLimitResponseInner) HasIntervalNum() bool {
	if o != nil && !common.IsNil(o.IntervalNum) {
		return true
	}

	return false
}

// SetIntervalNum gets a reference to the given int64 and assigns it to the IntervalNum field.
func (o *QueryUserRateLimitResponseInner) SetIntervalNum(v int64) {
	o.IntervalNum = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *QueryUserRateLimitResponseInner) GetLimit() int64 {
	if o == nil || common.IsNil(o.Limit) {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserRateLimitResponseInner) GetLimitOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *QueryUserRateLimitResponseInner) HasLimit() bool {
	if o != nil && !common.IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *QueryUserRateLimitResponseInner) SetLimit(v int64) {
	o.Limit = &v
}

func (o QueryUserRateLimitResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryUserRateLimitResponseInner) ToMap() (map[string]interface{}, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryUserRateLimitResponseInner) UnmarshalJSON(data []byte) (err error) {
	varQueryUserRateLimitResponseInner := _QueryUserRateLimitResponseInner{}

	err = json.Unmarshal(data, &varQueryUserRateLimitResponseInner)

	if err != nil {
		return err
	}

	*o = QueryUserRateLimitResponseInner(varQueryUserRateLimitResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rateLimitType")
		delete(additionalProperties, "interval")
		delete(additionalProperties, "intervalNum")
		delete(additionalProperties, "limit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryUserRateLimitResponseInner struct {
	value *QueryUserRateLimitResponseInner
	isSet bool
}

func (v NullableQueryUserRateLimitResponseInner) Get() *QueryUserRateLimitResponseInner {
	return v.value
}

func (v *NullableQueryUserRateLimitResponseInner) Set(val *QueryUserRateLimitResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryUserRateLimitResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryUserRateLimitResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryUserRateLimitResponseInner(val *QueryUserRateLimitResponseInner) *NullableQueryUserRateLimitResponseInner {
	return &NullableQueryUserRateLimitResponseInner{value: val, isSet: true}
}

func (v NullableQueryUserRateLimitResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryUserRateLimitResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
