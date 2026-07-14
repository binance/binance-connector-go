/*
Spot WebSocket API

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the BlockTradesHistoricalResponseRateLimitsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BlockTradesHistoricalResponseRateLimitsInner{}

// BlockTradesHistoricalResponseRateLimitsInner struct for BlockTradesHistoricalResponseRateLimitsInner
type BlockTradesHistoricalResponseRateLimitsInner struct {
	RateLimitType        *string `json:"rateLimitType,omitempty"`
	Interval             *string `json:"interval,omitempty"`
	IntervalNum          *int32  `json:"intervalNum,omitempty"`
	Limit                *int32  `json:"limit,omitempty"`
	Count                *int32  `json:"count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BlockTradesHistoricalResponseRateLimitsInner BlockTradesHistoricalResponseRateLimitsInner

// NewBlockTradesHistoricalResponseRateLimitsInner instantiates a new BlockTradesHistoricalResponseRateLimitsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockTradesHistoricalResponseRateLimitsInner() *BlockTradesHistoricalResponseRateLimitsInner {
	this := BlockTradesHistoricalResponseRateLimitsInner{}
	return &this
}

// NewBlockTradesHistoricalResponseRateLimitsInnerWithDefaults instantiates a new BlockTradesHistoricalResponseRateLimitsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockTradesHistoricalResponseRateLimitsInnerWithDefaults() *BlockTradesHistoricalResponseRateLimitsInner {
	this := BlockTradesHistoricalResponseRateLimitsInner{}
	return &this
}

// GetRateLimitType returns the RateLimitType field value if set, zero value otherwise.
func (o *BlockTradesHistoricalResponseRateLimitsInner) GetRateLimitType() string {
	if o == nil || common.IsNil(o.RateLimitType) {
		var ret string
		return ret
	}
	return *o.RateLimitType
}

// GetRateLimitTypeOk returns a tuple with the RateLimitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTradesHistoricalResponseRateLimitsInner) GetRateLimitTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.RateLimitType) {
		return nil, false
	}
	return o.RateLimitType, true
}

// HasRateLimitType returns a boolean if a field has been set.
func (o *BlockTradesHistoricalResponseRateLimitsInner) HasRateLimitType() bool {
	if o != nil && !common.IsNil(o.RateLimitType) {
		return true
	}

	return false
}

// SetRateLimitType gets a reference to the given string and assigns it to the RateLimitType field.
func (o *BlockTradesHistoricalResponseRateLimitsInner) SetRateLimitType(v string) {
	o.RateLimitType = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *BlockTradesHistoricalResponseRateLimitsInner) GetInterval() string {
	if o == nil || common.IsNil(o.Interval) {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTradesHistoricalResponseRateLimitsInner) GetIntervalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *BlockTradesHistoricalResponseRateLimitsInner) HasInterval() bool {
	if o != nil && !common.IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *BlockTradesHistoricalResponseRateLimitsInner) SetInterval(v string) {
	o.Interval = &v
}

// GetIntervalNum returns the IntervalNum field value if set, zero value otherwise.
func (o *BlockTradesHistoricalResponseRateLimitsInner) GetIntervalNum() int32 {
	if o == nil || common.IsNil(o.IntervalNum) {
		var ret int32
		return ret
	}
	return *o.IntervalNum
}

// GetIntervalNumOk returns a tuple with the IntervalNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTradesHistoricalResponseRateLimitsInner) GetIntervalNumOk() (*int32, bool) {
	if o == nil || common.IsNil(o.IntervalNum) {
		return nil, false
	}
	return o.IntervalNum, true
}

// HasIntervalNum returns a boolean if a field has been set.
func (o *BlockTradesHistoricalResponseRateLimitsInner) HasIntervalNum() bool {
	if o != nil && !common.IsNil(o.IntervalNum) {
		return true
	}

	return false
}

// SetIntervalNum gets a reference to the given int32 and assigns it to the IntervalNum field.
func (o *BlockTradesHistoricalResponseRateLimitsInner) SetIntervalNum(v int32) {
	o.IntervalNum = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *BlockTradesHistoricalResponseRateLimitsInner) GetLimit() int32 {
	if o == nil || common.IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTradesHistoricalResponseRateLimitsInner) GetLimitOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *BlockTradesHistoricalResponseRateLimitsInner) HasLimit() bool {
	if o != nil && !common.IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *BlockTradesHistoricalResponseRateLimitsInner) SetLimit(v int32) {
	o.Limit = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *BlockTradesHistoricalResponseRateLimitsInner) GetCount() int32 {
	if o == nil || common.IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockTradesHistoricalResponseRateLimitsInner) GetCountOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *BlockTradesHistoricalResponseRateLimitsInner) HasCount() bool {
	if o != nil && !common.IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *BlockTradesHistoricalResponseRateLimitsInner) SetCount(v int32) {
	o.Count = &v
}

func (o BlockTradesHistoricalResponseRateLimitsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockTradesHistoricalResponseRateLimitsInner) ToMap() (map[string]interface{}, error) {
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

func (o *BlockTradesHistoricalResponseRateLimitsInner) UnmarshalJSON(data []byte) (err error) {
	varBlockTradesHistoricalResponseRateLimitsInner := _BlockTradesHistoricalResponseRateLimitsInner{}

	err = json.Unmarshal(data, &varBlockTradesHistoricalResponseRateLimitsInner)

	if err != nil {
		return err
	}

	*o = BlockTradesHistoricalResponseRateLimitsInner(varBlockTradesHistoricalResponseRateLimitsInner)

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

type NullableBlockTradesHistoricalResponseRateLimitsInner struct {
	value *BlockTradesHistoricalResponseRateLimitsInner
	isSet bool
}

func (v NullableBlockTradesHistoricalResponseRateLimitsInner) Get() *BlockTradesHistoricalResponseRateLimitsInner {
	return v.value
}

func (v *NullableBlockTradesHistoricalResponseRateLimitsInner) Set(val *BlockTradesHistoricalResponseRateLimitsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockTradesHistoricalResponseRateLimitsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockTradesHistoricalResponseRateLimitsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockTradesHistoricalResponseRateLimitsInner(val *BlockTradesHistoricalResponseRateLimitsInner) *NullableBlockTradesHistoricalResponseRateLimitsInner {
	return &NullableBlockTradesHistoricalResponseRateLimitsInner{value: val, isSet: true}
}

func (v NullableBlockTradesHistoricalResponseRateLimitsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockTradesHistoricalResponseRateLimitsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
