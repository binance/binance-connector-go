/*
Spot WebSocket API

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the OpenOrderListsStatusResponseRateLimitsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OpenOrderListsStatusResponseRateLimitsInner{}

// OpenOrderListsStatusResponseRateLimitsInner struct for OpenOrderListsStatusResponseRateLimitsInner
type OpenOrderListsStatusResponseRateLimitsInner struct {
	RateLimitType        *string `json:"rateLimitType,omitempty"`
	Interval             *string `json:"interval,omitempty"`
	IntervalNum          *int64  `json:"intervalNum,omitempty"`
	Limit                *int64  `json:"limit,omitempty"`
	Count                *int64  `json:"count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OpenOrderListsStatusResponseRateLimitsInner OpenOrderListsStatusResponseRateLimitsInner

// NewOpenOrderListsStatusResponseRateLimitsInner instantiates a new OpenOrderListsStatusResponseRateLimitsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenOrderListsStatusResponseRateLimitsInner() *OpenOrderListsStatusResponseRateLimitsInner {
	this := OpenOrderListsStatusResponseRateLimitsInner{}
	return &this
}

// NewOpenOrderListsStatusResponseRateLimitsInnerWithDefaults instantiates a new OpenOrderListsStatusResponseRateLimitsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenOrderListsStatusResponseRateLimitsInnerWithDefaults() *OpenOrderListsStatusResponseRateLimitsInner {
	this := OpenOrderListsStatusResponseRateLimitsInner{}
	return &this
}

// GetRateLimitType returns the RateLimitType field value if set, zero value otherwise.
func (o *OpenOrderListsStatusResponseRateLimitsInner) GetRateLimitType() string {
	if o == nil || common.IsNil(o.RateLimitType) {
		var ret string
		return ret
	}
	return *o.RateLimitType
}

// GetRateLimitTypeOk returns a tuple with the RateLimitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenOrderListsStatusResponseRateLimitsInner) GetRateLimitTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.RateLimitType) {
		return nil, false
	}
	return o.RateLimitType, true
}

// HasRateLimitType returns a boolean if a field has been set.
func (o *OpenOrderListsStatusResponseRateLimitsInner) HasRateLimitType() bool {
	if o != nil && !common.IsNil(o.RateLimitType) {
		return true
	}

	return false
}

// SetRateLimitType gets a reference to the given string and assigns it to the RateLimitType field.
func (o *OpenOrderListsStatusResponseRateLimitsInner) SetRateLimitType(v string) {
	o.RateLimitType = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *OpenOrderListsStatusResponseRateLimitsInner) GetInterval() string {
	if o == nil || common.IsNil(o.Interval) {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenOrderListsStatusResponseRateLimitsInner) GetIntervalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *OpenOrderListsStatusResponseRateLimitsInner) HasInterval() bool {
	if o != nil && !common.IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *OpenOrderListsStatusResponseRateLimitsInner) SetInterval(v string) {
	o.Interval = &v
}

// GetIntervalNum returns the IntervalNum field value if set, zero value otherwise.
func (o *OpenOrderListsStatusResponseRateLimitsInner) GetIntervalNum() int64 {
	if o == nil || common.IsNil(o.IntervalNum) {
		var ret int64
		return ret
	}
	return *o.IntervalNum
}

// GetIntervalNumOk returns a tuple with the IntervalNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenOrderListsStatusResponseRateLimitsInner) GetIntervalNumOk() (*int64, bool) {
	if o == nil || common.IsNil(o.IntervalNum) {
		return nil, false
	}
	return o.IntervalNum, true
}

// HasIntervalNum returns a boolean if a field has been set.
func (o *OpenOrderListsStatusResponseRateLimitsInner) HasIntervalNum() bool {
	if o != nil && !common.IsNil(o.IntervalNum) {
		return true
	}

	return false
}

// SetIntervalNum gets a reference to the given int64 and assigns it to the IntervalNum field.
func (o *OpenOrderListsStatusResponseRateLimitsInner) SetIntervalNum(v int64) {
	o.IntervalNum = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *OpenOrderListsStatusResponseRateLimitsInner) GetLimit() int64 {
	if o == nil || common.IsNil(o.Limit) {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenOrderListsStatusResponseRateLimitsInner) GetLimitOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *OpenOrderListsStatusResponseRateLimitsInner) HasLimit() bool {
	if o != nil && !common.IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *OpenOrderListsStatusResponseRateLimitsInner) SetLimit(v int64) {
	o.Limit = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *OpenOrderListsStatusResponseRateLimitsInner) GetCount() int64 {
	if o == nil || common.IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenOrderListsStatusResponseRateLimitsInner) GetCountOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *OpenOrderListsStatusResponseRateLimitsInner) HasCount() bool {
	if o != nil && !common.IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *OpenOrderListsStatusResponseRateLimitsInner) SetCount(v int64) {
	o.Count = &v
}

func (o OpenOrderListsStatusResponseRateLimitsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenOrderListsStatusResponseRateLimitsInner) ToMap() (map[string]interface{}, error) {
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

func (o *OpenOrderListsStatusResponseRateLimitsInner) UnmarshalJSON(data []byte) (err error) {
	varOpenOrderListsStatusResponseRateLimitsInner := _OpenOrderListsStatusResponseRateLimitsInner{}

	err = json.Unmarshal(data, &varOpenOrderListsStatusResponseRateLimitsInner)

	if err != nil {
		return err
	}

	*o = OpenOrderListsStatusResponseRateLimitsInner(varOpenOrderListsStatusResponseRateLimitsInner)

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

type NullableOpenOrderListsStatusResponseRateLimitsInner struct {
	value *OpenOrderListsStatusResponseRateLimitsInner
	isSet bool
}

func (v NullableOpenOrderListsStatusResponseRateLimitsInner) Get() *OpenOrderListsStatusResponseRateLimitsInner {
	return v.value
}

func (v *NullableOpenOrderListsStatusResponseRateLimitsInner) Set(val *OpenOrderListsStatusResponseRateLimitsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenOrderListsStatusResponseRateLimitsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenOrderListsStatusResponseRateLimitsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenOrderListsStatusResponseRateLimitsInner(val *OpenOrderListsStatusResponseRateLimitsInner) *NullableOpenOrderListsStatusResponseRateLimitsInner {
	return &NullableOpenOrderListsStatusResponseRateLimitsInner{value: val, isSet: true}
}

func (v NullableOpenOrderListsStatusResponseRateLimitsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenOrderListsStatusResponseRateLimitsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
