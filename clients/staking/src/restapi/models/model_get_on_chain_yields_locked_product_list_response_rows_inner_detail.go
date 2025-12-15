/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetOnChainYieldsLockedProductListResponseRowsInnerDetail type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetOnChainYieldsLockedProductListResponseRowsInnerDetail{}

// GetOnChainYieldsLockedProductListResponseRowsInnerDetail struct for GetOnChainYieldsLockedProductListResponseRowsInnerDetail
type GetOnChainYieldsLockedProductListResponseRowsInnerDetail struct {
	Asset                 *string `json:"asset,omitempty"`
	RewardAsset           *string `json:"rewardAsset,omitempty"`
	Duration              *int64  `json:"duration,omitempty"`
	Renewable             *bool   `json:"renewable,omitempty"`
	IsSoldOut             *bool   `json:"isSoldOut,omitempty"`
	Apr                   *string `json:"apr,omitempty"`
	Status                *string `json:"status,omitempty"`
	SubscriptionStartTime *int64  `json:"subscriptionStartTime,omitempty"`
	CanRedeemToFlex       *bool   `json:"canRedeemToFlex,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _GetOnChainYieldsLockedProductListResponseRowsInnerDetail GetOnChainYieldsLockedProductListResponseRowsInnerDetail

// NewGetOnChainYieldsLockedProductListResponseRowsInnerDetail instantiates a new GetOnChainYieldsLockedProductListResponseRowsInnerDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOnChainYieldsLockedProductListResponseRowsInnerDetail() *GetOnChainYieldsLockedProductListResponseRowsInnerDetail {
	this := GetOnChainYieldsLockedProductListResponseRowsInnerDetail{}
	return &this
}

// NewGetOnChainYieldsLockedProductListResponseRowsInnerDetailWithDefaults instantiates a new GetOnChainYieldsLockedProductListResponseRowsInnerDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOnChainYieldsLockedProductListResponseRowsInnerDetailWithDefaults() *GetOnChainYieldsLockedProductListResponseRowsInnerDetail {
	this := GetOnChainYieldsLockedProductListResponseRowsInnerDetail{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) SetAsset(v string) {
	o.Asset = &v
}

// GetRewardAsset returns the RewardAsset field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) GetRewardAsset() string {
	if o == nil || common.IsNil(o.RewardAsset) {
		var ret string
		return ret
	}
	return *o.RewardAsset
}

// GetRewardAssetOk returns a tuple with the RewardAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) GetRewardAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.RewardAsset) {
		return nil, false
	}
	return o.RewardAsset, true
}

// HasRewardAsset returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) HasRewardAsset() bool {
	if o != nil && !common.IsNil(o.RewardAsset) {
		return true
	}

	return false
}

// SetRewardAsset gets a reference to the given string and assigns it to the RewardAsset field.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) SetRewardAsset(v string) {
	o.RewardAsset = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) GetDuration() int64 {
	if o == nil || common.IsNil(o.Duration) {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) GetDurationOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) HasDuration() bool {
	if o != nil && !common.IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) SetDuration(v int64) {
	o.Duration = &v
}

// GetRenewable returns the Renewable field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) GetRenewable() bool {
	if o == nil || common.IsNil(o.Renewable) {
		var ret bool
		return ret
	}
	return *o.Renewable
}

// GetRenewableOk returns a tuple with the Renewable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) GetRenewableOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Renewable) {
		return nil, false
	}
	return o.Renewable, true
}

// HasRenewable returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) HasRenewable() bool {
	if o != nil && !common.IsNil(o.Renewable) {
		return true
	}

	return false
}

// SetRenewable gets a reference to the given bool and assigns it to the Renewable field.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) SetRenewable(v bool) {
	o.Renewable = &v
}

// GetIsSoldOut returns the IsSoldOut field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) GetIsSoldOut() bool {
	if o == nil || common.IsNil(o.IsSoldOut) {
		var ret bool
		return ret
	}
	return *o.IsSoldOut
}

// GetIsSoldOutOk returns a tuple with the IsSoldOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) GetIsSoldOutOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsSoldOut) {
		return nil, false
	}
	return o.IsSoldOut, true
}

// HasIsSoldOut returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) HasIsSoldOut() bool {
	if o != nil && !common.IsNil(o.IsSoldOut) {
		return true
	}

	return false
}

// SetIsSoldOut gets a reference to the given bool and assigns it to the IsSoldOut field.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) SetIsSoldOut(v bool) {
	o.IsSoldOut = &v
}

// GetApr returns the Apr field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) GetApr() string {
	if o == nil || common.IsNil(o.Apr) {
		var ret string
		return ret
	}
	return *o.Apr
}

// GetAprOk returns a tuple with the Apr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) GetAprOk() (*string, bool) {
	if o == nil || common.IsNil(o.Apr) {
		return nil, false
	}
	return o.Apr, true
}

// HasApr returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) HasApr() bool {
	if o != nil && !common.IsNil(o.Apr) {
		return true
	}

	return false
}

// SetApr gets a reference to the given string and assigns it to the Apr field.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) SetApr(v string) {
	o.Apr = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) SetStatus(v string) {
	o.Status = &v
}

// GetSubscriptionStartTime returns the SubscriptionStartTime field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) GetSubscriptionStartTime() int64 {
	if o == nil || common.IsNil(o.SubscriptionStartTime) {
		var ret int64
		return ret
	}
	return *o.SubscriptionStartTime
}

// GetSubscriptionStartTimeOk returns a tuple with the SubscriptionStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) GetSubscriptionStartTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.SubscriptionStartTime) {
		return nil, false
	}
	return o.SubscriptionStartTime, true
}

// HasSubscriptionStartTime returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) HasSubscriptionStartTime() bool {
	if o != nil && !common.IsNil(o.SubscriptionStartTime) {
		return true
	}

	return false
}

// SetSubscriptionStartTime gets a reference to the given int64 and assigns it to the SubscriptionStartTime field.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) SetSubscriptionStartTime(v int64) {
	o.SubscriptionStartTime = &v
}

// GetCanRedeemToFlex returns the CanRedeemToFlex field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) GetCanRedeemToFlex() bool {
	if o == nil || common.IsNil(o.CanRedeemToFlex) {
		var ret bool
		return ret
	}
	return *o.CanRedeemToFlex
}

// GetCanRedeemToFlexOk returns a tuple with the CanRedeemToFlex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) GetCanRedeemToFlexOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanRedeemToFlex) {
		return nil, false
	}
	return o.CanRedeemToFlex, true
}

// HasCanRedeemToFlex returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) HasCanRedeemToFlex() bool {
	if o != nil && !common.IsNil(o.CanRedeemToFlex) {
		return true
	}

	return false
}

// SetCanRedeemToFlex gets a reference to the given bool and assigns it to the CanRedeemToFlex field.
func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) SetCanRedeemToFlex(v bool) {
	o.CanRedeemToFlex = &v
}

func (o GetOnChainYieldsLockedProductListResponseRowsInnerDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOnChainYieldsLockedProductListResponseRowsInnerDetail) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.CanRedeemToFlex) {
		toSerialize["canRedeemToFlex"] = o.CanRedeemToFlex
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) UnmarshalJSON(data []byte) (err error) {
	varGetOnChainYieldsLockedProductListResponseRowsInnerDetail := _GetOnChainYieldsLockedProductListResponseRowsInnerDetail{}

	err = json.Unmarshal(data, &varGetOnChainYieldsLockedProductListResponseRowsInnerDetail)

	if err != nil {
		return err
	}

	*o = GetOnChainYieldsLockedProductListResponseRowsInnerDetail(varGetOnChainYieldsLockedProductListResponseRowsInnerDetail)

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
		delete(additionalProperties, "canRedeemToFlex")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetOnChainYieldsLockedProductListResponseRowsInnerDetail struct {
	value *GetOnChainYieldsLockedProductListResponseRowsInnerDetail
	isSet bool
}

func (v NullableGetOnChainYieldsLockedProductListResponseRowsInnerDetail) Get() *GetOnChainYieldsLockedProductListResponseRowsInnerDetail {
	return v.value
}

func (v *NullableGetOnChainYieldsLockedProductListResponseRowsInnerDetail) Set(val *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOnChainYieldsLockedProductListResponseRowsInnerDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOnChainYieldsLockedProductListResponseRowsInnerDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOnChainYieldsLockedProductListResponseRowsInnerDetail(val *GetOnChainYieldsLockedProductListResponseRowsInnerDetail) *NullableGetOnChainYieldsLockedProductListResponseRowsInnerDetail {
	return &NullableGetOnChainYieldsLockedProductListResponseRowsInnerDetail{value: val, isSet: true}
}

func (v NullableGetOnChainYieldsLockedProductListResponseRowsInnerDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOnChainYieldsLockedProductListResponseRowsInnerDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
