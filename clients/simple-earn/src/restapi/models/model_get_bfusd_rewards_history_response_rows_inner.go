/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetBfusdRewardsHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetBfusdRewardsHistoryResponseRowsInner{}

// GetBfusdRewardsHistoryResponseRowsInner struct for GetBfusdRewardsHistoryResponseRowsInner
type GetBfusdRewardsHistoryResponseRowsInner struct {
	Time                 *int64  `json:"time,omitempty"`
	RewardAsset          *string `json:"rewardAsset,omitempty"`
	RewardsAmount        *string `json:"rewardsAmount,omitempty"`
	BFUSDPosition        *string `json:"BFUSDPosition,omitempty"`
	AnnualPercentageRate *string `json:"annualPercentageRate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetBfusdRewardsHistoryResponseRowsInner GetBfusdRewardsHistoryResponseRowsInner

// NewGetBfusdRewardsHistoryResponseRowsInner instantiates a new GetBfusdRewardsHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBfusdRewardsHistoryResponseRowsInner() *GetBfusdRewardsHistoryResponseRowsInner {
	this := GetBfusdRewardsHistoryResponseRowsInner{}
	return &this
}

// NewGetBfusdRewardsHistoryResponseRowsInnerWithDefaults instantiates a new GetBfusdRewardsHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBfusdRewardsHistoryResponseRowsInnerWithDefaults() *GetBfusdRewardsHistoryResponseRowsInner {
	this := GetBfusdRewardsHistoryResponseRowsInner{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetBfusdRewardsHistoryResponseRowsInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdRewardsHistoryResponseRowsInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetBfusdRewardsHistoryResponseRowsInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetBfusdRewardsHistoryResponseRowsInner) SetTime(v int64) {
	o.Time = &v
}

// GetRewardAsset returns the RewardAsset field value if set, zero value otherwise.
func (o *GetBfusdRewardsHistoryResponseRowsInner) GetRewardAsset() string {
	if o == nil || common.IsNil(o.RewardAsset) {
		var ret string
		return ret
	}
	return *o.RewardAsset
}

// GetRewardAssetOk returns a tuple with the RewardAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdRewardsHistoryResponseRowsInner) GetRewardAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.RewardAsset) {
		return nil, false
	}
	return o.RewardAsset, true
}

// HasRewardAsset returns a boolean if a field has been set.
func (o *GetBfusdRewardsHistoryResponseRowsInner) HasRewardAsset() bool {
	if o != nil && !common.IsNil(o.RewardAsset) {
		return true
	}

	return false
}

// SetRewardAsset gets a reference to the given string and assigns it to the RewardAsset field.
func (o *GetBfusdRewardsHistoryResponseRowsInner) SetRewardAsset(v string) {
	o.RewardAsset = &v
}

// GetRewardsAmount returns the RewardsAmount field value if set, zero value otherwise.
func (o *GetBfusdRewardsHistoryResponseRowsInner) GetRewardsAmount() string {
	if o == nil || common.IsNil(o.RewardsAmount) {
		var ret string
		return ret
	}
	return *o.RewardsAmount
}

// GetRewardsAmountOk returns a tuple with the RewardsAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdRewardsHistoryResponseRowsInner) GetRewardsAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.RewardsAmount) {
		return nil, false
	}
	return o.RewardsAmount, true
}

// HasRewardsAmount returns a boolean if a field has been set.
func (o *GetBfusdRewardsHistoryResponseRowsInner) HasRewardsAmount() bool {
	if o != nil && !common.IsNil(o.RewardsAmount) {
		return true
	}

	return false
}

// SetRewardsAmount gets a reference to the given string and assigns it to the RewardsAmount field.
func (o *GetBfusdRewardsHistoryResponseRowsInner) SetRewardsAmount(v string) {
	o.RewardsAmount = &v
}

// GetBFUSDPosition returns the BFUSDPosition field value if set, zero value otherwise.
func (o *GetBfusdRewardsHistoryResponseRowsInner) GetBFUSDPosition() string {
	if o == nil || common.IsNil(o.BFUSDPosition) {
		var ret string
		return ret
	}
	return *o.BFUSDPosition
}

// GetBFUSDPositionOk returns a tuple with the BFUSDPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdRewardsHistoryResponseRowsInner) GetBFUSDPositionOk() (*string, bool) {
	if o == nil || common.IsNil(o.BFUSDPosition) {
		return nil, false
	}
	return o.BFUSDPosition, true
}

// HasBFUSDPosition returns a boolean if a field has been set.
func (o *GetBfusdRewardsHistoryResponseRowsInner) HasBFUSDPosition() bool {
	if o != nil && !common.IsNil(o.BFUSDPosition) {
		return true
	}

	return false
}

// SetBFUSDPosition gets a reference to the given string and assigns it to the BFUSDPosition field.
func (o *GetBfusdRewardsHistoryResponseRowsInner) SetBFUSDPosition(v string) {
	o.BFUSDPosition = &v
}

// GetAnnualPercentageRate returns the AnnualPercentageRate field value if set, zero value otherwise.
func (o *GetBfusdRewardsHistoryResponseRowsInner) GetAnnualPercentageRate() string {
	if o == nil || common.IsNil(o.AnnualPercentageRate) {
		var ret string
		return ret
	}
	return *o.AnnualPercentageRate
}

// GetAnnualPercentageRateOk returns a tuple with the AnnualPercentageRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBfusdRewardsHistoryResponseRowsInner) GetAnnualPercentageRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.AnnualPercentageRate) {
		return nil, false
	}
	return o.AnnualPercentageRate, true
}

// HasAnnualPercentageRate returns a boolean if a field has been set.
func (o *GetBfusdRewardsHistoryResponseRowsInner) HasAnnualPercentageRate() bool {
	if o != nil && !common.IsNil(o.AnnualPercentageRate) {
		return true
	}

	return false
}

// SetAnnualPercentageRate gets a reference to the given string and assigns it to the AnnualPercentageRate field.
func (o *GetBfusdRewardsHistoryResponseRowsInner) SetAnnualPercentageRate(v string) {
	o.AnnualPercentageRate = &v
}

func (o GetBfusdRewardsHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBfusdRewardsHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.RewardAsset) {
		toSerialize["rewardAsset"] = o.RewardAsset
	}
	if !common.IsNil(o.RewardsAmount) {
		toSerialize["rewardsAmount"] = o.RewardsAmount
	}
	if !common.IsNil(o.BFUSDPosition) {
		toSerialize["BFUSDPosition"] = o.BFUSDPosition
	}
	if !common.IsNil(o.AnnualPercentageRate) {
		toSerialize["annualPercentageRate"] = o.AnnualPercentageRate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetBfusdRewardsHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetBfusdRewardsHistoryResponseRowsInner := _GetBfusdRewardsHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varGetBfusdRewardsHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetBfusdRewardsHistoryResponseRowsInner(varGetBfusdRewardsHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "time")
		delete(additionalProperties, "rewardAsset")
		delete(additionalProperties, "rewardsAmount")
		delete(additionalProperties, "BFUSDPosition")
		delete(additionalProperties, "annualPercentageRate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetBfusdRewardsHistoryResponseRowsInner struct {
	value *GetBfusdRewardsHistoryResponseRowsInner
	isSet bool
}

func (v NullableGetBfusdRewardsHistoryResponseRowsInner) Get() *GetBfusdRewardsHistoryResponseRowsInner {
	return v.value
}

func (v *NullableGetBfusdRewardsHistoryResponseRowsInner) Set(val *GetBfusdRewardsHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBfusdRewardsHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBfusdRewardsHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBfusdRewardsHistoryResponseRowsInner(val *GetBfusdRewardsHistoryResponseRowsInner) *NullableGetBfusdRewardsHistoryResponseRowsInner {
	return &NullableGetBfusdRewardsHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetBfusdRewardsHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBfusdRewardsHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
