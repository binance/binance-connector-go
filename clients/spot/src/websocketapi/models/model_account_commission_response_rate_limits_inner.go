/*
Spot WebSocket API

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AccountCommissionResponseRateLimitsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountCommissionResponseRateLimitsInner{}

// AccountCommissionResponseRateLimitsInner struct for AccountCommissionResponseRateLimitsInner
type AccountCommissionResponseRateLimitsInner struct {
	RateLimitType        *string `json:"rateLimitType,omitempty"`
	Interval             *string `json:"interval,omitempty"`
	IntervalNum          *int64  `json:"intervalNum,omitempty"`
	Limit                *int64  `json:"limit,omitempty"`
	Count                *int64  `json:"count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountCommissionResponseRateLimitsInner AccountCommissionResponseRateLimitsInner

// NewAccountCommissionResponseRateLimitsInner instantiates a new AccountCommissionResponseRateLimitsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountCommissionResponseRateLimitsInner() *AccountCommissionResponseRateLimitsInner {
	this := AccountCommissionResponseRateLimitsInner{}
	return &this
}

// NewAccountCommissionResponseRateLimitsInnerWithDefaults instantiates a new AccountCommissionResponseRateLimitsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCommissionResponseRateLimitsInnerWithDefaults() *AccountCommissionResponseRateLimitsInner {
	this := AccountCommissionResponseRateLimitsInner{}
	return &this
}

// GetRateLimitType returns the RateLimitType field value if set, zero value otherwise.
func (o *AccountCommissionResponseRateLimitsInner) GetRateLimitType() string {
	if o == nil || common.IsNil(o.RateLimitType) {
		var ret string
		return ret
	}
	return *o.RateLimitType
}

// GetRateLimitTypeOk returns a tuple with the RateLimitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseRateLimitsInner) GetRateLimitTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.RateLimitType) {
		return nil, false
	}
	return o.RateLimitType, true
}

// HasRateLimitType returns a boolean if a field has been set.
func (o *AccountCommissionResponseRateLimitsInner) HasRateLimitType() bool {
	if o != nil && !common.IsNil(o.RateLimitType) {
		return true
	}

	return false
}

// SetRateLimitType gets a reference to the given string and assigns it to the RateLimitType field.
func (o *AccountCommissionResponseRateLimitsInner) SetRateLimitType(v string) {
	o.RateLimitType = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *AccountCommissionResponseRateLimitsInner) GetInterval() string {
	if o == nil || common.IsNil(o.Interval) {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseRateLimitsInner) GetIntervalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *AccountCommissionResponseRateLimitsInner) HasInterval() bool {
	if o != nil && !common.IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *AccountCommissionResponseRateLimitsInner) SetInterval(v string) {
	o.Interval = &v
}

// GetIntervalNum returns the IntervalNum field value if set, zero value otherwise.
func (o *AccountCommissionResponseRateLimitsInner) GetIntervalNum() int64 {
	if o == nil || common.IsNil(o.IntervalNum) {
		var ret int64
		return ret
	}
	return *o.IntervalNum
}

// GetIntervalNumOk returns a tuple with the IntervalNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseRateLimitsInner) GetIntervalNumOk() (*int64, bool) {
	if o == nil || common.IsNil(o.IntervalNum) {
		return nil, false
	}
	return o.IntervalNum, true
}

// HasIntervalNum returns a boolean if a field has been set.
func (o *AccountCommissionResponseRateLimitsInner) HasIntervalNum() bool {
	if o != nil && !common.IsNil(o.IntervalNum) {
		return true
	}

	return false
}

// SetIntervalNum gets a reference to the given int64 and assigns it to the IntervalNum field.
func (o *AccountCommissionResponseRateLimitsInner) SetIntervalNum(v int64) {
	o.IntervalNum = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *AccountCommissionResponseRateLimitsInner) GetLimit() int64 {
	if o == nil || common.IsNil(o.Limit) {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseRateLimitsInner) GetLimitOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *AccountCommissionResponseRateLimitsInner) HasLimit() bool {
	if o != nil && !common.IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *AccountCommissionResponseRateLimitsInner) SetLimit(v int64) {
	o.Limit = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *AccountCommissionResponseRateLimitsInner) GetCount() int64 {
	if o == nil || common.IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCommissionResponseRateLimitsInner) GetCountOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *AccountCommissionResponseRateLimitsInner) HasCount() bool {
	if o != nil && !common.IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *AccountCommissionResponseRateLimitsInner) SetCount(v int64) {
	o.Count = &v
}

func (o AccountCommissionResponseRateLimitsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountCommissionResponseRateLimitsInner) ToMap() (map[string]interface{}, error) {
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

func (o *AccountCommissionResponseRateLimitsInner) UnmarshalJSON(data []byte) (err error) {
	varAccountCommissionResponseRateLimitsInner := _AccountCommissionResponseRateLimitsInner{}

	err = json.Unmarshal(data, &varAccountCommissionResponseRateLimitsInner)

	if err != nil {
		return err
	}

	*o = AccountCommissionResponseRateLimitsInner(varAccountCommissionResponseRateLimitsInner)

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

type NullableAccountCommissionResponseRateLimitsInner struct {
	value *AccountCommissionResponseRateLimitsInner
	isSet bool
}

func (v NullableAccountCommissionResponseRateLimitsInner) Get() *AccountCommissionResponseRateLimitsInner {
	return v.value
}

func (v *NullableAccountCommissionResponseRateLimitsInner) Set(val *AccountCommissionResponseRateLimitsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountCommissionResponseRateLimitsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountCommissionResponseRateLimitsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountCommissionResponseRateLimitsInner(val *AccountCommissionResponseRateLimitsInner) *NullableAccountCommissionResponseRateLimitsInner {
	return &NullableAccountCommissionResponseRateLimitsInner{value: val, isSet: true}
}

func (v NullableAccountCommissionResponseRateLimitsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountCommissionResponseRateLimitsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
