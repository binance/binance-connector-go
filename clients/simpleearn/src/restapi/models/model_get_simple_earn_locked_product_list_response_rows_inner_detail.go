/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetSimpleEarnLockedProductListResponseRowsInnerDetail type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSimpleEarnLockedProductListResponseRowsInnerDetail{}

// GetSimpleEarnLockedProductListResponseRowsInnerDetail struct for GetSimpleEarnLockedProductListResponseRowsInnerDetail
type GetSimpleEarnLockedProductListResponseRowsInnerDetail struct {
	Asset                 *string `json:"asset,omitempty"`
	RewardAsset           *string `json:"rewardAsset,omitempty"`
	Duration              *int64  `json:"duration,omitempty"`
	Renewable             *bool   `json:"renewable,omitempty"`
	IsSoldOut             *bool   `json:"isSoldOut,omitempty"`
	Apr                   *string `json:"apr,omitempty"`
	Status                *string `json:"status,omitempty"`
	SubscriptionStartTime *int64  `json:"subscriptionStartTime,omitempty"`
	ExtraRewardAsset      *string `json:"extraRewardAsset,omitempty"`
	ExtraRewardAPR        *string `json:"extraRewardAPR,omitempty"`
	BoostRewardAsset      *string `json:"boostRewardAsset,omitempty"`
	BoostApr              *string `json:"boostApr,omitempty"`
	BoostEndTime          *int64  `json:"boostEndTime,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _GetSimpleEarnLockedProductListResponseRowsInnerDetail GetSimpleEarnLockedProductListResponseRowsInnerDetail

// NewGetSimpleEarnLockedProductListResponseRowsInnerDetail instantiates a new GetSimpleEarnLockedProductListResponseRowsInnerDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSimpleEarnLockedProductListResponseRowsInnerDetail() *GetSimpleEarnLockedProductListResponseRowsInnerDetail {
	this := GetSimpleEarnLockedProductListResponseRowsInnerDetail{}
	return &this
}

// NewGetSimpleEarnLockedProductListResponseRowsInnerDetailWithDefaults instantiates a new GetSimpleEarnLockedProductListResponseRowsInnerDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSimpleEarnLockedProductListResponseRowsInnerDetailWithDefaults() *GetSimpleEarnLockedProductListResponseRowsInnerDetail {
	this := GetSimpleEarnLockedProductListResponseRowsInnerDetail{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) SetAsset(v string) {
	o.Asset = &v
}

// GetRewardAsset returns the RewardAsset field value if set, zero value otherwise.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) GetRewardAsset() string {
	if o == nil || common.IsNil(o.RewardAsset) {
		var ret string
		return ret
	}
	return *o.RewardAsset
}

// GetRewardAssetOk returns a tuple with the RewardAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) GetRewardAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.RewardAsset) {
		return nil, false
	}
	return o.RewardAsset, true
}

// HasRewardAsset returns a boolean if a field has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) HasRewardAsset() bool {
	if o != nil && !common.IsNil(o.RewardAsset) {
		return true
	}

	return false
}

// SetRewardAsset gets a reference to the given string and assigns it to the RewardAsset field.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) SetRewardAsset(v string) {
	o.RewardAsset = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) GetDuration() int64 {
	if o == nil || common.IsNil(o.Duration) {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) GetDurationOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) HasDuration() bool {
	if o != nil && !common.IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) SetDuration(v int64) {
	o.Duration = &v
}

// GetRenewable returns the Renewable field value if set, zero value otherwise.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) GetRenewable() bool {
	if o == nil || common.IsNil(o.Renewable) {
		var ret bool
		return ret
	}
	return *o.Renewable
}

// GetRenewableOk returns a tuple with the Renewable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) GetRenewableOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Renewable) {
		return nil, false
	}
	return o.Renewable, true
}

// HasRenewable returns a boolean if a field has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) HasRenewable() bool {
	if o != nil && !common.IsNil(o.Renewable) {
		return true
	}

	return false
}

// SetRenewable gets a reference to the given bool and assigns it to the Renewable field.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) SetRenewable(v bool) {
	o.Renewable = &v
}

// GetIsSoldOut returns the IsSoldOut field value if set, zero value otherwise.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) GetIsSoldOut() bool {
	if o == nil || common.IsNil(o.IsSoldOut) {
		var ret bool
		return ret
	}
	return *o.IsSoldOut
}

// GetIsSoldOutOk returns a tuple with the IsSoldOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) GetIsSoldOutOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsSoldOut) {
		return nil, false
	}
	return o.IsSoldOut, true
}

// HasIsSoldOut returns a boolean if a field has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) HasIsSoldOut() bool {
	if o != nil && !common.IsNil(o.IsSoldOut) {
		return true
	}

	return false
}

// SetIsSoldOut gets a reference to the given bool and assigns it to the IsSoldOut field.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) SetIsSoldOut(v bool) {
	o.IsSoldOut = &v
}

// GetApr returns the Apr field value if set, zero value otherwise.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) GetApr() string {
	if o == nil || common.IsNil(o.Apr) {
		var ret string
		return ret
	}
	return *o.Apr
}

// GetAprOk returns a tuple with the Apr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) GetAprOk() (*string, bool) {
	if o == nil || common.IsNil(o.Apr) {
		return nil, false
	}
	return o.Apr, true
}

// HasApr returns a boolean if a field has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) HasApr() bool {
	if o != nil && !common.IsNil(o.Apr) {
		return true
	}

	return false
}

// SetApr gets a reference to the given string and assigns it to the Apr field.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) SetApr(v string) {
	o.Apr = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) SetStatus(v string) {
	o.Status = &v
}

// GetSubscriptionStartTime returns the SubscriptionStartTime field value if set, zero value otherwise.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) GetSubscriptionStartTime() int64 {
	if o == nil || common.IsNil(o.SubscriptionStartTime) {
		var ret int64
		return ret
	}
	return *o.SubscriptionStartTime
}

// GetSubscriptionStartTimeOk returns a tuple with the SubscriptionStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) GetSubscriptionStartTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.SubscriptionStartTime) {
		return nil, false
	}
	return o.SubscriptionStartTime, true
}

// HasSubscriptionStartTime returns a boolean if a field has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) HasSubscriptionStartTime() bool {
	if o != nil && !common.IsNil(o.SubscriptionStartTime) {
		return true
	}

	return false
}

// SetSubscriptionStartTime gets a reference to the given int64 and assigns it to the SubscriptionStartTime field.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) SetSubscriptionStartTime(v int64) {
	o.SubscriptionStartTime = &v
}

// GetExtraRewardAsset returns the ExtraRewardAsset field value if set, zero value otherwise.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) GetExtraRewardAsset() string {
	if o == nil || common.IsNil(o.ExtraRewardAsset) {
		var ret string
		return ret
	}
	return *o.ExtraRewardAsset
}

// GetExtraRewardAssetOk returns a tuple with the ExtraRewardAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) GetExtraRewardAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExtraRewardAsset) {
		return nil, false
	}
	return o.ExtraRewardAsset, true
}

// HasExtraRewardAsset returns a boolean if a field has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) HasExtraRewardAsset() bool {
	if o != nil && !common.IsNil(o.ExtraRewardAsset) {
		return true
	}

	return false
}

// SetExtraRewardAsset gets a reference to the given string and assigns it to the ExtraRewardAsset field.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) SetExtraRewardAsset(v string) {
	o.ExtraRewardAsset = &v
}

// GetExtraRewardAPR returns the ExtraRewardAPR field value if set, zero value otherwise.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) GetExtraRewardAPR() string {
	if o == nil || common.IsNil(o.ExtraRewardAPR) {
		var ret string
		return ret
	}
	return *o.ExtraRewardAPR
}

// GetExtraRewardAPROk returns a tuple with the ExtraRewardAPR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) GetExtraRewardAPROk() (*string, bool) {
	if o == nil || common.IsNil(o.ExtraRewardAPR) {
		return nil, false
	}
	return o.ExtraRewardAPR, true
}

// HasExtraRewardAPR returns a boolean if a field has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) HasExtraRewardAPR() bool {
	if o != nil && !common.IsNil(o.ExtraRewardAPR) {
		return true
	}

	return false
}

// SetExtraRewardAPR gets a reference to the given string and assigns it to the ExtraRewardAPR field.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) SetExtraRewardAPR(v string) {
	o.ExtraRewardAPR = &v
}

// GetBoostRewardAsset returns the BoostRewardAsset field value if set, zero value otherwise.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) GetBoostRewardAsset() string {
	if o == nil || common.IsNil(o.BoostRewardAsset) {
		var ret string
		return ret
	}
	return *o.BoostRewardAsset
}

// GetBoostRewardAssetOk returns a tuple with the BoostRewardAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) GetBoostRewardAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.BoostRewardAsset) {
		return nil, false
	}
	return o.BoostRewardAsset, true
}

// HasBoostRewardAsset returns a boolean if a field has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) HasBoostRewardAsset() bool {
	if o != nil && !common.IsNil(o.BoostRewardAsset) {
		return true
	}

	return false
}

// SetBoostRewardAsset gets a reference to the given string and assigns it to the BoostRewardAsset field.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) SetBoostRewardAsset(v string) {
	o.BoostRewardAsset = &v
}

// GetBoostApr returns the BoostApr field value if set, zero value otherwise.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) GetBoostApr() string {
	if o == nil || common.IsNil(o.BoostApr) {
		var ret string
		return ret
	}
	return *o.BoostApr
}

// GetBoostAprOk returns a tuple with the BoostApr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) GetBoostAprOk() (*string, bool) {
	if o == nil || common.IsNil(o.BoostApr) {
		return nil, false
	}
	return o.BoostApr, true
}

// HasBoostApr returns a boolean if a field has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) HasBoostApr() bool {
	if o != nil && !common.IsNil(o.BoostApr) {
		return true
	}

	return false
}

// SetBoostApr gets a reference to the given string and assigns it to the BoostApr field.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) SetBoostApr(v string) {
	o.BoostApr = &v
}

// GetBoostEndTime returns the BoostEndTime field value if set, zero value otherwise.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) GetBoostEndTime() int64 {
	if o == nil || common.IsNil(o.BoostEndTime) {
		var ret int64
		return ret
	}
	return *o.BoostEndTime
}

// GetBoostEndTimeOk returns a tuple with the BoostEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) GetBoostEndTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.BoostEndTime) {
		return nil, false
	}
	return o.BoostEndTime, true
}

// HasBoostEndTime returns a boolean if a field has been set.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) HasBoostEndTime() bool {
	if o != nil && !common.IsNil(o.BoostEndTime) {
		return true
	}

	return false
}

// SetBoostEndTime gets a reference to the given int64 and assigns it to the BoostEndTime field.
func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) SetBoostEndTime(v int64) {
	o.BoostEndTime = &v
}

func (o GetSimpleEarnLockedProductListResponseRowsInnerDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSimpleEarnLockedProductListResponseRowsInnerDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.RewardAsset) {
		toSerialize["rewardAsset"] = o.RewardAsset
	}
	if !common.IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !common.IsNil(o.Renewable) {
		toSerialize["renewable"] = o.Renewable
	}
	if !common.IsNil(o.IsSoldOut) {
		toSerialize["isSoldOut"] = o.IsSoldOut
	}
	if !common.IsNil(o.Apr) {
		toSerialize["apr"] = o.Apr
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.SubscriptionStartTime) {
		toSerialize["subscriptionStartTime"] = o.SubscriptionStartTime
	}
	if !common.IsNil(o.ExtraRewardAsset) {
		toSerialize["extraRewardAsset"] = o.ExtraRewardAsset
	}
	if !common.IsNil(o.ExtraRewardAPR) {
		toSerialize["extraRewardAPR"] = o.ExtraRewardAPR
	}
	if !common.IsNil(o.BoostRewardAsset) {
		toSerialize["boostRewardAsset"] = o.BoostRewardAsset
	}
	if !common.IsNil(o.BoostApr) {
		toSerialize["boostApr"] = o.BoostApr
	}
	if !common.IsNil(o.BoostEndTime) {
		toSerialize["boostEndTime"] = o.BoostEndTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetSimpleEarnLockedProductListResponseRowsInnerDetail) UnmarshalJSON(data []byte) (err error) {
	varGetSimpleEarnLockedProductListResponseRowsInnerDetail := _GetSimpleEarnLockedProductListResponseRowsInnerDetail{}

	err = json.Unmarshal(data, &varGetSimpleEarnLockedProductListResponseRowsInnerDetail)

	if err != nil {
		return err
	}

	*o = GetSimpleEarnLockedProductListResponseRowsInnerDetail(varGetSimpleEarnLockedProductListResponseRowsInnerDetail)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "rewardAsset")
		delete(additionalProperties, "duration")
		delete(additionalProperties, "renewable")
		delete(additionalProperties, "isSoldOut")
		delete(additionalProperties, "apr")
		delete(additionalProperties, "status")
		delete(additionalProperties, "subscriptionStartTime")
		delete(additionalProperties, "extraRewardAsset")
		delete(additionalProperties, "extraRewardAPR")
		delete(additionalProperties, "boostRewardAsset")
		delete(additionalProperties, "boostApr")
		delete(additionalProperties, "boostEndTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSimpleEarnLockedProductListResponseRowsInnerDetail struct {
	value *GetSimpleEarnLockedProductListResponseRowsInnerDetail
	isSet bool
}

func (v NullableGetSimpleEarnLockedProductListResponseRowsInnerDetail) Get() *GetSimpleEarnLockedProductListResponseRowsInnerDetail {
	return v.value
}

func (v *NullableGetSimpleEarnLockedProductListResponseRowsInnerDetail) Set(val *GetSimpleEarnLockedProductListResponseRowsInnerDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSimpleEarnLockedProductListResponseRowsInnerDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSimpleEarnLockedProductListResponseRowsInnerDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSimpleEarnLockedProductListResponseRowsInnerDetail(val *GetSimpleEarnLockedProductListResponseRowsInnerDetail) *NullableGetSimpleEarnLockedProductListResponseRowsInnerDetail {
	return &NullableGetSimpleEarnLockedProductListResponseRowsInnerDetail{value: val, isSet: true}
}

func (v NullableGetSimpleEarnLockedProductListResponseRowsInnerDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSimpleEarnLockedProductListResponseRowsInnerDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
