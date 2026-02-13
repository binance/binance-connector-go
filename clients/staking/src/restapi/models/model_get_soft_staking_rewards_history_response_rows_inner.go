/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetSoftStakingRewardsHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSoftStakingRewardsHistoryResponseRowsInner{}

// GetSoftStakingRewardsHistoryResponseRowsInner struct for GetSoftStakingRewardsHistoryResponseRowsInner
type GetSoftStakingRewardsHistoryResponseRowsInner struct {
	Asset                *string `json:"asset,omitempty"`
	Rewards              *string `json:"rewards,omitempty"`
	RewardAsset          *string `json:"rewardAsset,omitempty"`
	AvgAmount            *string `json:"avgAmount,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSoftStakingRewardsHistoryResponseRowsInner GetSoftStakingRewardsHistoryResponseRowsInner

// NewGetSoftStakingRewardsHistoryResponseRowsInner instantiates a new GetSoftStakingRewardsHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSoftStakingRewardsHistoryResponseRowsInner() *GetSoftStakingRewardsHistoryResponseRowsInner {
	this := GetSoftStakingRewardsHistoryResponseRowsInner{}
	return &this
}

// NewGetSoftStakingRewardsHistoryResponseRowsInnerWithDefaults instantiates a new GetSoftStakingRewardsHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSoftStakingRewardsHistoryResponseRowsInnerWithDefaults() *GetSoftStakingRewardsHistoryResponseRowsInner {
	this := GetSoftStakingRewardsHistoryResponseRowsInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetSoftStakingRewardsHistoryResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSoftStakingRewardsHistoryResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetSoftStakingRewardsHistoryResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetSoftStakingRewardsHistoryResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetRewards returns the Rewards field value if set, zero value otherwise.
func (o *GetSoftStakingRewardsHistoryResponseRowsInner) GetRewards() string {
	if o == nil || common.IsNil(o.Rewards) {
		var ret string
		return ret
	}
	return *o.Rewards
}

// GetRewardsOk returns a tuple with the Rewards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSoftStakingRewardsHistoryResponseRowsInner) GetRewardsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Rewards) {
		return nil, false
	}
	return o.Rewards, true
}

// HasRewards returns a boolean if a field has been set.
func (o *GetSoftStakingRewardsHistoryResponseRowsInner) HasRewards() bool {
	if o != nil && !common.IsNil(o.Rewards) {
		return true
	}

	return false
}

// SetRewards gets a reference to the given string and assigns it to the Rewards field.
func (o *GetSoftStakingRewardsHistoryResponseRowsInner) SetRewards(v string) {
	o.Rewards = &v
}

// GetRewardAsset returns the RewardAsset field value if set, zero value otherwise.
func (o *GetSoftStakingRewardsHistoryResponseRowsInner) GetRewardAsset() string {
	if o == nil || common.IsNil(o.RewardAsset) {
		var ret string
		return ret
	}
	return *o.RewardAsset
}

// GetRewardAssetOk returns a tuple with the RewardAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSoftStakingRewardsHistoryResponseRowsInner) GetRewardAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.RewardAsset) {
		return nil, false
	}
	return o.RewardAsset, true
}

// HasRewardAsset returns a boolean if a field has been set.
func (o *GetSoftStakingRewardsHistoryResponseRowsInner) HasRewardAsset() bool {
	if o != nil && !common.IsNil(o.RewardAsset) {
		return true
	}

	return false
}

// SetRewardAsset gets a reference to the given string and assigns it to the RewardAsset field.
func (o *GetSoftStakingRewardsHistoryResponseRowsInner) SetRewardAsset(v string) {
	o.RewardAsset = &v
}

// GetAvgAmount returns the AvgAmount field value if set, zero value otherwise.
func (o *GetSoftStakingRewardsHistoryResponseRowsInner) GetAvgAmount() string {
	if o == nil || common.IsNil(o.AvgAmount) {
		var ret string
		return ret
	}
	return *o.AvgAmount
}

// GetAvgAmountOk returns a tuple with the AvgAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSoftStakingRewardsHistoryResponseRowsInner) GetAvgAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.AvgAmount) {
		return nil, false
	}
	return o.AvgAmount, true
}

// HasAvgAmount returns a boolean if a field has been set.
func (o *GetSoftStakingRewardsHistoryResponseRowsInner) HasAvgAmount() bool {
	if o != nil && !common.IsNil(o.AvgAmount) {
		return true
	}

	return false
}

// SetAvgAmount gets a reference to the given string and assigns it to the AvgAmount field.
func (o *GetSoftStakingRewardsHistoryResponseRowsInner) SetAvgAmount(v string) {
	o.AvgAmount = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetSoftStakingRewardsHistoryResponseRowsInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSoftStakingRewardsHistoryResponseRowsInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetSoftStakingRewardsHistoryResponseRowsInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetSoftStakingRewardsHistoryResponseRowsInner) SetTime(v int64) {
	o.Time = &v
}

func (o GetSoftStakingRewardsHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSoftStakingRewardsHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Rewards) {
		toSerialize["rewards"] = o.Rewards
	}
	if !common.IsNil(o.RewardAsset) {
		toSerialize["rewardAsset"] = o.RewardAsset
	}
	if !common.IsNil(o.AvgAmount) {
		toSerialize["avgAmount"] = o.AvgAmount
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetSoftStakingRewardsHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetSoftStakingRewardsHistoryResponseRowsInner := _GetSoftStakingRewardsHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetSoftStakingRewardsHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetSoftStakingRewardsHistoryResponseRowsInner(varGetSoftStakingRewardsHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "rewards")
		delete(additionalProperties, "rewardAsset")
		delete(additionalProperties, "avgAmount")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSoftStakingRewardsHistoryResponseRowsInner struct {
	value *GetSoftStakingRewardsHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetSoftStakingRewardsHistoryResponseRowsInner) Get() *GetSoftStakingRewardsHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetSoftStakingRewardsHistoryResponseRowsInner) Set(val *GetSoftStakingRewardsHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSoftStakingRewardsHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSoftStakingRewardsHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSoftStakingRewardsHistoryResponseRowsInner(val *GetSoftStakingRewardsHistoryResponseRowsInner) *NullableGetSoftStakingRewardsHistoryResponseRowsInner {
	return &NullableGetSoftStakingRewardsHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetSoftStakingRewardsHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSoftStakingRewardsHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
