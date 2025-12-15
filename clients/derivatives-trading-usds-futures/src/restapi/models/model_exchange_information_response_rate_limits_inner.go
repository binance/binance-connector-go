/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the ExchangeInformationResponseRateLimitsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExchangeInformationResponseRateLimitsInner{}

// ExchangeInformationResponseRateLimitsInner struct for ExchangeInformationResponseRateLimitsInner
type ExchangeInformationResponseRateLimitsInner struct {
	Interval             *string `json:"interval,omitempty"`
	IntervalNum          *int64  `json:"intervalNum,omitempty"`
	Limit                *int64  `json:"limit,omitempty"`
	RateLimitType        *string `json:"rateLimitType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExchangeInformationResponseRateLimitsInner ExchangeInformationResponseRateLimitsInner

// NewExchangeInformationResponseRateLimitsInner instantiates a new ExchangeInformationResponseRateLimitsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeInformationResponseRateLimitsInner() *ExchangeInformationResponseRateLimitsInner {
	this := ExchangeInformationResponseRateLimitsInner{}
	return &this
}

// NewExchangeInformationResponseRateLimitsInnerWithDefaults instantiates a new ExchangeInformationResponseRateLimitsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeInformationResponseRateLimitsInnerWithDefaults() *ExchangeInformationResponseRateLimitsInner {
	this := ExchangeInformationResponseRateLimitsInner{}
	return &this
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *ExchangeInformationResponseRateLimitsInner) GetInterval() string {
	if o == nil || common.IsNil(o.Interval) {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseRateLimitsInner) GetIntervalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *ExchangeInformationResponseRateLimitsInner) HasInterval() bool {
	if o != nil && !common.IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *ExchangeInformationResponseRateLimitsInner) SetInterval(v string) {
	o.Interval = &v
}

// GetIntervalNum returns the IntervalNum field value if set, zero value otherwise.
func (o *ExchangeInformationResponseRateLimitsInner) GetIntervalNum() int64 {
	if o == nil || common.IsNil(o.IntervalNum) {
		var ret int64
		return ret
	}
	return *o.IntervalNum
}

// GetIntervalNumOk returns a tuple with the IntervalNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseRateLimitsInner) GetIntervalNumOk() (*int64, bool) {
	if o == nil || common.IsNil(o.IntervalNum) {
		return nil, false
	}
	return o.IntervalNum, true
}

// HasIntervalNum returns a boolean if a field has been set.
func (o *ExchangeInformationResponseRateLimitsInner) HasIntervalNum() bool {
	if o != nil && !common.IsNil(o.IntervalNum) {
		return true
	}

	return false
}

// SetIntervalNum gets a reference to the given int64 and assigns it to the IntervalNum field.
func (o *ExchangeInformationResponseRateLimitsInner) SetIntervalNum(v int64) {
	o.IntervalNum = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ExchangeInformationResponseRateLimitsInner) GetLimit() int64 {
	if o == nil || common.IsNil(o.Limit) {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseRateLimitsInner) GetLimitOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ExchangeInformationResponseRateLimitsInner) HasLimit() bool {
	if o != nil && !common.IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *ExchangeInformationResponseRateLimitsInner) SetLimit(v int64) {
	o.Limit = &v
}

// GetRateLimitType returns the RateLimitType field value if set, zero value otherwise.
func (o *ExchangeInformationResponseRateLimitsInner) GetRateLimitType() string {
	if o == nil || common.IsNil(o.RateLimitType) {
		var ret string
		return ret
	}
	return *o.RateLimitType
}

// GetRateLimitTypeOk returns a tuple with the RateLimitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInformationResponseRateLimitsInner) GetRateLimitTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.RateLimitType) {
		return nil, false
	}
	return o.RateLimitType, true
}

// HasRateLimitType returns a boolean if a field has been set.
func (o *ExchangeInformationResponseRateLimitsInner) HasRateLimitType() bool {
	if o != nil && !common.IsNil(o.RateLimitType) {
		return true
	}

	return false
}

// SetRateLimitType gets a reference to the given string and assigns it to the RateLimitType field.
func (o *ExchangeInformationResponseRateLimitsInner) SetRateLimitType(v string) {
	o.RateLimitType = &v
}

func (o ExchangeInformationResponseRateLimitsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeInformationResponseRateLimitsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !common.IsNil(o.IntervalNum) {
		toSerialize["intervalNum"] = o.IntervalNum
	}
	if !common.IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !common.IsNil(o.RateLimitType) {
		toSerialize["rateLimitType"] = o.RateLimitType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExchangeInformationResponseRateLimitsInner) UnmarshalJSON(data []byte) (err error) {
	varExchangeInformationResponseRateLimitsInner := _ExchangeInformationResponseRateLimitsInner{}

	err = json.Unmarshal(data, &varExchangeInformationResponseRateLimitsInner)

	if err != nil {
		return err
	}

	*o = ExchangeInformationResponseRateLimitsInner(varExchangeInformationResponseRateLimitsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "interval")
		delete(additionalProperties, "intervalNum")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "rateLimitType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExchangeInformationResponseRateLimitsInner struct {
	value *ExchangeInformationResponseRateLimitsInner
	isSet bool
}

func (v NullableExchangeInformationResponseRateLimitsInner) Get() *ExchangeInformationResponseRateLimitsInner {
	return v.value
}

func (v *NullableExchangeInformationResponseRateLimitsInner) Set(val *ExchangeInformationResponseRateLimitsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeInformationResponseRateLimitsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeInformationResponseRateLimitsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeInformationResponseRateLimitsInner(val *ExchangeInformationResponseRateLimitsInner) *NullableExchangeInformationResponseRateLimitsInner {
	return &NullableExchangeInformationResponseRateLimitsInner{value: val, isSet: true}
}

func (v NullableExchangeInformationResponseRateLimitsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeInformationResponseRateLimitsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
