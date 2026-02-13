/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryMarginInterestRateHistoryResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryMarginInterestRateHistoryResponseInner{}

// QueryMarginInterestRateHistoryResponseInner struct for QueryMarginInterestRateHistoryResponseInner
type QueryMarginInterestRateHistoryResponseInner struct {
	Asset                *string `json:"asset,omitempty"`
	DailyInterestRate    *string `json:"dailyInterestRate,omitempty"`
	Timestamp            *int64  `json:"timestamp,omitempty"`
	VipLevel             *int64  `json:"vipLevel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryMarginInterestRateHistoryResponseInner QueryMarginInterestRateHistoryResponseInner

// NewQueryMarginInterestRateHistoryResponseInner instantiates a new QueryMarginInterestRateHistoryResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryMarginInterestRateHistoryResponseInner() *QueryMarginInterestRateHistoryResponseInner {
	this := QueryMarginInterestRateHistoryResponseInner{}
	return &this
}

// NewQueryMarginInterestRateHistoryResponseInnerWithDefaults instantiates a new QueryMarginInterestRateHistoryResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryMarginInterestRateHistoryResponseInnerWithDefaults() *QueryMarginInterestRateHistoryResponseInner {
	this := QueryMarginInterestRateHistoryResponseInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *QueryMarginInterestRateHistoryResponseInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginInterestRateHistoryResponseInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *QueryMarginInterestRateHistoryResponseInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *QueryMarginInterestRateHistoryResponseInner) SetAsset(v string) {
	o.Asset = &v
}

// GetDailyInterestRate returns the DailyInterestRate field value if set, zero value otherwise.
func (o *QueryMarginInterestRateHistoryResponseInner) GetDailyInterestRate() string {
	if o == nil || common.IsNil(o.DailyInterestRate) {
		var ret string
		return ret
	}
	return *o.DailyInterestRate
}

// GetDailyInterestRateOk returns a tuple with the DailyInterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginInterestRateHistoryResponseInner) GetDailyInterestRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.DailyInterestRate) {
		return nil, false
	}
	return o.DailyInterestRate, true
}

// HasDailyInterestRate returns a boolean if a field has been set.
func (o *QueryMarginInterestRateHistoryResponseInner) HasDailyInterestRate() bool {
	if o != nil && !common.IsNil(o.DailyInterestRate) {
		return true
	}

	return false
}

// SetDailyInterestRate gets a reference to the given string and assigns it to the DailyInterestRate field.
func (o *QueryMarginInterestRateHistoryResponseInner) SetDailyInterestRate(v string) {
	o.DailyInterestRate = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *QueryMarginInterestRateHistoryResponseInner) GetTimestamp() int64 {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginInterestRateHistoryResponseInner) GetTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *QueryMarginInterestRateHistoryResponseInner) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *QueryMarginInterestRateHistoryResponseInner) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetVipLevel returns the VipLevel field value if set, zero value otherwise.
func (o *QueryMarginInterestRateHistoryResponseInner) GetVipLevel() int64 {
	if o == nil || common.IsNil(o.VipLevel) {
		var ret int64
		return ret
	}
	return *o.VipLevel
}

// GetVipLevelOk returns a tuple with the VipLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryMarginInterestRateHistoryResponseInner) GetVipLevelOk() (*int64, bool) {
	if o == nil || common.IsNil(o.VipLevel) {
		return nil, false
	}
	return o.VipLevel, true
}

// HasVipLevel returns a boolean if a field has been set.
func (o *QueryMarginInterestRateHistoryResponseInner) HasVipLevel() bool {
	if o != nil && !common.IsNil(o.VipLevel) {
		return true
	}

	return false
}

// SetVipLevel gets a reference to the given int64 and assigns it to the VipLevel field.
func (o *QueryMarginInterestRateHistoryResponseInner) SetVipLevel(v int64) {
	o.VipLevel = &v
}

func (o QueryMarginInterestRateHistoryResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryMarginInterestRateHistoryResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.DailyInterestRate) {
		toSerialize["dailyInterestRate"] = o.DailyInterestRate
	}
	if !common.IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !common.IsNil(o.VipLevel) {
		toSerialize["vipLevel"] = o.VipLevel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryMarginInterestRateHistoryResponseInner) UnmarshalJSON(data []byte) (err error) {
	varQueryMarginInterestRateHistoryResponseInner := _QueryMarginInterestRateHistoryResponseInner{}

	err = json.Unmarshal(data, &varQueryMarginInterestRateHistoryResponseInner)

	if err != nil {
		return err
	}

	*o = QueryMarginInterestRateHistoryResponseInner(varQueryMarginInterestRateHistoryResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "dailyInterestRate")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "vipLevel")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryMarginInterestRateHistoryResponseInner struct {
	value *QueryMarginInterestRateHistoryResponseInner
	isSet bool
}

func (v NullableQueryMarginInterestRateHistoryResponseInner) Get() *QueryMarginInterestRateHistoryResponseInner {
	return v.value
}

func (v *NullableQueryMarginInterestRateHistoryResponseInner) Set(val *QueryMarginInterestRateHistoryResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryMarginInterestRateHistoryResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryMarginInterestRateHistoryResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryMarginInterestRateHistoryResponseInner(val *QueryMarginInterestRateHistoryResponseInner) *NullableQueryMarginInterestRateHistoryResponseInner {
	return &NullableQueryMarginInterestRateHistoryResponseInner{value: val, isSet: true}
}

func (v NullableQueryMarginInterestRateHistoryResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryMarginInterestRateHistoryResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
