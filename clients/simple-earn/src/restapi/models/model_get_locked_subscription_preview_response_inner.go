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

// checks if the GetLockedSubscriptionPreviewResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetLockedSubscriptionPreviewResponseInner{}

// GetLockedSubscriptionPreviewResponseInner struct for GetLockedSubscriptionPreviewResponseInner
type GetLockedSubscriptionPreviewResponseInner struct {
	RewardAsset            *string `json:"rewardAsset,omitempty"`
	TotalRewardAmt         *string `json:"totalRewardAmt,omitempty"`
	ExtraRewardAsset       *string `json:"extraRewardAsset,omitempty"`
	EstTotalExtraRewardAmt *string `json:"estTotalExtraRewardAmt,omitempty"`
	BoostRewardAsset       *string `json:"boostRewardAsset,omitempty"`
	EstDailyRewardAmt      *string `json:"estDailyRewardAmt,omitempty"`
	NextPay                *string `json:"nextPay,omitempty"`
	NextPayDate            *string `json:"nextPayDate,omitempty"`
	ValueDate              *string `json:"valueDate,omitempty"`
	RewardsEndDate         *string `json:"rewardsEndDate,omitempty"`
	DeliverDate            *string `json:"deliverDate,omitempty"`
	NextSubscriptionDate   *string `json:"nextSubscriptionDate,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _GetLockedSubscriptionPreviewResponseInner GetLockedSubscriptionPreviewResponseInner

// NewGetLockedSubscriptionPreviewResponseInner instantiates a new GetLockedSubscriptionPreviewResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLockedSubscriptionPreviewResponseInner() *GetLockedSubscriptionPreviewResponseInner {
	this := GetLockedSubscriptionPreviewResponseInner{}
	return &this
}

// NewGetLockedSubscriptionPreviewResponseInnerWithDefaults instantiates a new GetLockedSubscriptionPreviewResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLockedSubscriptionPreviewResponseInnerWithDefaults() *GetLockedSubscriptionPreviewResponseInner {
	this := GetLockedSubscriptionPreviewResponseInner{}
	return &this
}

// GetRewardAsset returns the RewardAsset field value if set, zero value otherwise.
func (o *GetLockedSubscriptionPreviewResponseInner) GetRewardAsset() string {
	if o == nil || common.IsNil(o.RewardAsset) {
		var ret string
		return ret
	}
	return *o.RewardAsset
}

// GetRewardAssetOk returns a tuple with the RewardAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedSubscriptionPreviewResponseInner) GetRewardAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.RewardAsset) {
		return nil, false
	}
	return o.RewardAsset, true
}

// HasRewardAsset returns a boolean if a field has been set.
func (o *GetLockedSubscriptionPreviewResponseInner) HasRewardAsset() bool {
	if o != nil && !common.IsNil(o.RewardAsset) {
		return true
	}

	return false
}

// SetRewardAsset gets a reference to the given string and assigns it to the RewardAsset field.
func (o *GetLockedSubscriptionPreviewResponseInner) SetRewardAsset(v string) {
	o.RewardAsset = &v
}

// GetTotalRewardAmt returns the TotalRewardAmt field value if set, zero value otherwise.
func (o *GetLockedSubscriptionPreviewResponseInner) GetTotalRewardAmt() string {
	if o == nil || common.IsNil(o.TotalRewardAmt) {
		var ret string
		return ret
	}
	return *o.TotalRewardAmt
}

// GetTotalRewardAmtOk returns a tuple with the TotalRewardAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedSubscriptionPreviewResponseInner) GetTotalRewardAmtOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalRewardAmt) {
		return nil, false
	}
	return o.TotalRewardAmt, true
}

// HasTotalRewardAmt returns a boolean if a field has been set.
func (o *GetLockedSubscriptionPreviewResponseInner) HasTotalRewardAmt() bool {
	if o != nil && !common.IsNil(o.TotalRewardAmt) {
		return true
	}

	return false
}

// SetTotalRewardAmt gets a reference to the given string and assigns it to the TotalRewardAmt field.
func (o *GetLockedSubscriptionPreviewResponseInner) SetTotalRewardAmt(v string) {
	o.TotalRewardAmt = &v
}

// GetExtraRewardAsset returns the ExtraRewardAsset field value if set, zero value otherwise.
func (o *GetLockedSubscriptionPreviewResponseInner) GetExtraRewardAsset() string {
	if o == nil || common.IsNil(o.ExtraRewardAsset) {
		var ret string
		return ret
	}
	return *o.ExtraRewardAsset
}

// GetExtraRewardAssetOk returns a tuple with the ExtraRewardAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedSubscriptionPreviewResponseInner) GetExtraRewardAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExtraRewardAsset) {
		return nil, false
	}
	return o.ExtraRewardAsset, true
}

// HasExtraRewardAsset returns a boolean if a field has been set.
func (o *GetLockedSubscriptionPreviewResponseInner) HasExtraRewardAsset() bool {
	if o != nil && !common.IsNil(o.ExtraRewardAsset) {
		return true
	}

	return false
}

// SetExtraRewardAsset gets a reference to the given string and assigns it to the ExtraRewardAsset field.
func (o *GetLockedSubscriptionPreviewResponseInner) SetExtraRewardAsset(v string) {
	o.ExtraRewardAsset = &v
}

// GetEstTotalExtraRewardAmt returns the EstTotalExtraRewardAmt field value if set, zero value otherwise.
func (o *GetLockedSubscriptionPreviewResponseInner) GetEstTotalExtraRewardAmt() string {
	if o == nil || common.IsNil(o.EstTotalExtraRewardAmt) {
		var ret string
		return ret
	}
	return *o.EstTotalExtraRewardAmt
}

// GetEstTotalExtraRewardAmtOk returns a tuple with the EstTotalExtraRewardAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedSubscriptionPreviewResponseInner) GetEstTotalExtraRewardAmtOk() (*string, bool) {
	if o == nil || common.IsNil(o.EstTotalExtraRewardAmt) {
		return nil, false
	}
	return o.EstTotalExtraRewardAmt, true
}

// HasEstTotalExtraRewardAmt returns a boolean if a field has been set.
func (o *GetLockedSubscriptionPreviewResponseInner) HasEstTotalExtraRewardAmt() bool {
	if o != nil && !common.IsNil(o.EstTotalExtraRewardAmt) {
		return true
	}

	return false
}

// SetEstTotalExtraRewardAmt gets a reference to the given string and assigns it to the EstTotalExtraRewardAmt field.
func (o *GetLockedSubscriptionPreviewResponseInner) SetEstTotalExtraRewardAmt(v string) {
	o.EstTotalExtraRewardAmt = &v
}

// GetBoostRewardAsset returns the BoostRewardAsset field value if set, zero value otherwise.
func (o *GetLockedSubscriptionPreviewResponseInner) GetBoostRewardAsset() string {
	if o == nil || common.IsNil(o.BoostRewardAsset) {
		var ret string
		return ret
	}
	return *o.BoostRewardAsset
}

// GetBoostRewardAssetOk returns a tuple with the BoostRewardAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedSubscriptionPreviewResponseInner) GetBoostRewardAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.BoostRewardAsset) {
		return nil, false
	}
	return o.BoostRewardAsset, true
}

// HasBoostRewardAsset returns a boolean if a field has been set.
func (o *GetLockedSubscriptionPreviewResponseInner) HasBoostRewardAsset() bool {
	if o != nil && !common.IsNil(o.BoostRewardAsset) {
		return true
	}

	return false
}

// SetBoostRewardAsset gets a reference to the given string and assigns it to the BoostRewardAsset field.
func (o *GetLockedSubscriptionPreviewResponseInner) SetBoostRewardAsset(v string) {
	o.BoostRewardAsset = &v
}

// GetEstDailyRewardAmt returns the EstDailyRewardAmt field value if set, zero value otherwise.
func (o *GetLockedSubscriptionPreviewResponseInner) GetEstDailyRewardAmt() string {
	if o == nil || common.IsNil(o.EstDailyRewardAmt) {
		var ret string
		return ret
	}
	return *o.EstDailyRewardAmt
}

// GetEstDailyRewardAmtOk returns a tuple with the EstDailyRewardAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedSubscriptionPreviewResponseInner) GetEstDailyRewardAmtOk() (*string, bool) {
	if o == nil || common.IsNil(o.EstDailyRewardAmt) {
		return nil, false
	}
	return o.EstDailyRewardAmt, true
}

// HasEstDailyRewardAmt returns a boolean if a field has been set.
func (o *GetLockedSubscriptionPreviewResponseInner) HasEstDailyRewardAmt() bool {
	if o != nil && !common.IsNil(o.EstDailyRewardAmt) {
		return true
	}

	return false
}

// SetEstDailyRewardAmt gets a reference to the given string and assigns it to the EstDailyRewardAmt field.
func (o *GetLockedSubscriptionPreviewResponseInner) SetEstDailyRewardAmt(v string) {
	o.EstDailyRewardAmt = &v
}

// GetNextPay returns the NextPay field value if set, zero value otherwise.
func (o *GetLockedSubscriptionPreviewResponseInner) GetNextPay() string {
	if o == nil || common.IsNil(o.NextPay) {
		var ret string
		return ret
	}
	return *o.NextPay
}

// GetNextPayOk returns a tuple with the NextPay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedSubscriptionPreviewResponseInner) GetNextPayOk() (*string, bool) {
	if o == nil || common.IsNil(o.NextPay) {
		return nil, false
	}
	return o.NextPay, true
}

// HasNextPay returns a boolean if a field has been set.
func (o *GetLockedSubscriptionPreviewResponseInner) HasNextPay() bool {
	if o != nil && !common.IsNil(o.NextPay) {
		return true
	}

	return false
}

// SetNextPay gets a reference to the given string and assigns it to the NextPay field.
func (o *GetLockedSubscriptionPreviewResponseInner) SetNextPay(v string) {
	o.NextPay = &v
}

// GetNextPayDate returns the NextPayDate field value if set, zero value otherwise.
func (o *GetLockedSubscriptionPreviewResponseInner) GetNextPayDate() string {
	if o == nil || common.IsNil(o.NextPayDate) {
		var ret string
		return ret
	}
	return *o.NextPayDate
}

// GetNextPayDateOk returns a tuple with the NextPayDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedSubscriptionPreviewResponseInner) GetNextPayDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.NextPayDate) {
		return nil, false
	}
	return o.NextPayDate, true
}

// HasNextPayDate returns a boolean if a field has been set.
func (o *GetLockedSubscriptionPreviewResponseInner) HasNextPayDate() bool {
	if o != nil && !common.IsNil(o.NextPayDate) {
		return true
	}

	return false
}

// SetNextPayDate gets a reference to the given string and assigns it to the NextPayDate field.
func (o *GetLockedSubscriptionPreviewResponseInner) SetNextPayDate(v string) {
	o.NextPayDate = &v
}

// GetValueDate returns the ValueDate field value if set, zero value otherwise.
func (o *GetLockedSubscriptionPreviewResponseInner) GetValueDate() string {
	if o == nil || common.IsNil(o.ValueDate) {
		var ret string
		return ret
	}
	return *o.ValueDate
}

// GetValueDateOk returns a tuple with the ValueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedSubscriptionPreviewResponseInner) GetValueDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.ValueDate) {
		return nil, false
	}
	return o.ValueDate, true
}

// HasValueDate returns a boolean if a field has been set.
func (o *GetLockedSubscriptionPreviewResponseInner) HasValueDate() bool {
	if o != nil && !common.IsNil(o.ValueDate) {
		return true
	}

	return false
}

// SetValueDate gets a reference to the given string and assigns it to the ValueDate field.
func (o *GetLockedSubscriptionPreviewResponseInner) SetValueDate(v string) {
	o.ValueDate = &v
}

// GetRewardsEndDate returns the RewardsEndDate field value if set, zero value otherwise.
func (o *GetLockedSubscriptionPreviewResponseInner) GetRewardsEndDate() string {
	if o == nil || common.IsNil(o.RewardsEndDate) {
		var ret string
		return ret
	}
	return *o.RewardsEndDate
}

// GetRewardsEndDateOk returns a tuple with the RewardsEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedSubscriptionPreviewResponseInner) GetRewardsEndDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.RewardsEndDate) {
		return nil, false
	}
	return o.RewardsEndDate, true
}

// HasRewardsEndDate returns a boolean if a field has been set.
func (o *GetLockedSubscriptionPreviewResponseInner) HasRewardsEndDate() bool {
	if o != nil && !common.IsNil(o.RewardsEndDate) {
		return true
	}

	return false
}

// SetRewardsEndDate gets a reference to the given string and assigns it to the RewardsEndDate field.
func (o *GetLockedSubscriptionPreviewResponseInner) SetRewardsEndDate(v string) {
	o.RewardsEndDate = &v
}

// GetDeliverDate returns the DeliverDate field value if set, zero value otherwise.
func (o *GetLockedSubscriptionPreviewResponseInner) GetDeliverDate() string {
	if o == nil || common.IsNil(o.DeliverDate) {
		var ret string
		return ret
	}
	return *o.DeliverDate
}

// GetDeliverDateOk returns a tuple with the DeliverDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedSubscriptionPreviewResponseInner) GetDeliverDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.DeliverDate) {
		return nil, false
	}
	return o.DeliverDate, true
}

// HasDeliverDate returns a boolean if a field has been set.
func (o *GetLockedSubscriptionPreviewResponseInner) HasDeliverDate() bool {
	if o != nil && !common.IsNil(o.DeliverDate) {
		return true
	}

	return false
}

// SetDeliverDate gets a reference to the given string and assigns it to the DeliverDate field.
func (o *GetLockedSubscriptionPreviewResponseInner) SetDeliverDate(v string) {
	o.DeliverDate = &v
}

// GetNextSubscriptionDate returns the NextSubscriptionDate field value if set, zero value otherwise.
func (o *GetLockedSubscriptionPreviewResponseInner) GetNextSubscriptionDate() string {
	if o == nil || common.IsNil(o.NextSubscriptionDate) {
		var ret string
		return ret
	}
	return *o.NextSubscriptionDate
}

// GetNextSubscriptionDateOk returns a tuple with the NextSubscriptionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedSubscriptionPreviewResponseInner) GetNextSubscriptionDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.NextSubscriptionDate) {
		return nil, false
	}
	return o.NextSubscriptionDate, true
}

// HasNextSubscriptionDate returns a boolean if a field has been set.
func (o *GetLockedSubscriptionPreviewResponseInner) HasNextSubscriptionDate() bool {
	if o != nil && !common.IsNil(o.NextSubscriptionDate) {
		return true
	}

	return false
}

// SetNextSubscriptionDate gets a reference to the given string and assigns it to the NextSubscriptionDate field.
func (o *GetLockedSubscriptionPreviewResponseInner) SetNextSubscriptionDate(v string) {
	o.NextSubscriptionDate = &v
}

func (o GetLockedSubscriptionPreviewResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLockedSubscriptionPreviewResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.RewardAsset) {
		toSerialize["rewardAsset"] = o.RewardAsset
	}
	if !common.IsNil(o.TotalRewardAmt) {
		toSerialize["totalRewardAmt"] = o.TotalRewardAmt
	}
	if !common.IsNil(o.ExtraRewardAsset) {
		toSerialize["extraRewardAsset"] = o.ExtraRewardAsset
	}
	if !common.IsNil(o.EstTotalExtraRewardAmt) {
		toSerialize["estTotalExtraRewardAmt"] = o.EstTotalExtraRewardAmt
	}
	if !common.IsNil(o.BoostRewardAsset) {
		toSerialize["boostRewardAsset"] = o.BoostRewardAsset
	}
	if !common.IsNil(o.EstDailyRewardAmt) {
		toSerialize["estDailyRewardAmt"] = o.EstDailyRewardAmt
	}
	if !common.IsNil(o.NextPay) {
		toSerialize["nextPay"] = o.NextPay
	}
	if !common.IsNil(o.NextPayDate) {
		toSerialize["nextPayDate"] = o.NextPayDate
	}
	if !common.IsNil(o.ValueDate) {
		toSerialize["valueDate"] = o.ValueDate
	}
	if !common.IsNil(o.RewardsEndDate) {
		toSerialize["rewardsEndDate"] = o.RewardsEndDate
	}
	if !common.IsNil(o.DeliverDate) {
		toSerialize["deliverDate"] = o.DeliverDate
	}
	if !common.IsNil(o.NextSubscriptionDate) {
		toSerialize["nextSubscriptionDate"] = o.NextSubscriptionDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetLockedSubscriptionPreviewResponseInner) UnmarshalJSON(data []byte) (err error) {
	varGetLockedSubscriptionPreviewResponseInner := _GetLockedSubscriptionPreviewResponseInner{}

	err = json.Unmarshal(data, &varGetLockedSubscriptionPreviewResponseInner)

	if err != nil {
		return err
	}

	*o = GetLockedSubscriptionPreviewResponseInner(varGetLockedSubscriptionPreviewResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rewardAsset")
		delete(additionalProperties, "totalRewardAmt")
		delete(additionalProperties, "extraRewardAsset")
		delete(additionalProperties, "estTotalExtraRewardAmt")
		delete(additionalProperties, "boostRewardAsset")
		delete(additionalProperties, "estDailyRewardAmt")
		delete(additionalProperties, "nextPay")
		delete(additionalProperties, "nextPayDate")
		delete(additionalProperties, "valueDate")
		delete(additionalProperties, "rewardsEndDate")
		delete(additionalProperties, "deliverDate")
		delete(additionalProperties, "nextSubscriptionDate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetLockedSubscriptionPreviewResponseInner struct {
	value *GetLockedSubscriptionPreviewResponseInner
	isSet bool
}

func (v NullableGetLockedSubscriptionPreviewResponseInner) Get() *GetLockedSubscriptionPreviewResponseInner {
	return v.value
}

func (v *NullableGetLockedSubscriptionPreviewResponseInner) Set(val *GetLockedSubscriptionPreviewResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLockedSubscriptionPreviewResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLockedSubscriptionPreviewResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLockedSubscriptionPreviewResponseInner(val *GetLockedSubscriptionPreviewResponseInner) *NullableGetLockedSubscriptionPreviewResponseInner {
	return &NullableGetLockedSubscriptionPreviewResponseInner{value: val, isSet: true}
}

func (v NullableGetLockedSubscriptionPreviewResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLockedSubscriptionPreviewResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
