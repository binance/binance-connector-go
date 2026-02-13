/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetLockedRedemptionRecordResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetLockedRedemptionRecordResponseRowsInner{}

// GetLockedRedemptionRecordResponseRowsInner struct for GetLockedRedemptionRecordResponseRowsInner
type GetLockedRedemptionRecordResponseRowsInner struct {
	PositionId           *int64  `json:"positionId,omitempty"`
	RedeemId             *int64  `json:"redeemId,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	LockPeriod           *string `json:"lockPeriod,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	OriginalAmount       *string `json:"originalAmount,omitempty"`
	Type                 *string `json:"type,omitempty"`
	DeliverDate          *string `json:"deliverDate,omitempty"`
	LossAmount           *string `json:"lossAmount,omitempty"`
	IsComplete           *bool   `json:"isComplete,omitempty"`
	RewardAsset          *string `json:"rewardAsset,omitempty"`
	RewardAmt            *string `json:"rewardAmt,omitempty"`
	ExtraRewardAsset     *string `json:"extraRewardAsset,omitempty"`
	EstExtraRewardAmt    *string `json:"estExtraRewardAmt,omitempty"`
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetLockedRedemptionRecordResponseRowsInner GetLockedRedemptionRecordResponseRowsInner

// NewGetLockedRedemptionRecordResponseRowsInner instantiates a new GetLockedRedemptionRecordResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLockedRedemptionRecordResponseRowsInner() *GetLockedRedemptionRecordResponseRowsInner {
	this := GetLockedRedemptionRecordResponseRowsInner{}
	return &this
}

// NewGetLockedRedemptionRecordResponseRowsInnerWithDefaults instantiates a new GetLockedRedemptionRecordResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLockedRedemptionRecordResponseRowsInnerWithDefaults() *GetLockedRedemptionRecordResponseRowsInner {
	this := GetLockedRedemptionRecordResponseRowsInner{}
	return &this
}

// GetPositionId returns the PositionId field value if set, zero value otherwise.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetPositionId() int64 {
	if o == nil || common.IsNil(o.PositionId) {
		var ret int64
		return ret
	}
	return *o.PositionId
}

// GetPositionIdOk returns a tuple with the PositionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetPositionIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PositionId) {
		return nil, false
	}
	return o.PositionId, true
}

// HasPositionId returns a boolean if a field has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) HasPositionId() bool {
	if o != nil && !common.IsNil(o.PositionId) {
		return true
	}

	return false
}

// SetPositionId gets a reference to the given int64 and assigns it to the PositionId field.
func (o *GetLockedRedemptionRecordResponseRowsInner) SetPositionId(v int64) {
	o.PositionId = &v
}

// GetRedeemId returns the RedeemId field value if set, zero value otherwise.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetRedeemId() int64 {
	if o == nil || common.IsNil(o.RedeemId) {
		var ret int64
		return ret
	}
	return *o.RedeemId
}

// GetRedeemIdOk returns a tuple with the RedeemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetRedeemIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.RedeemId) {
		return nil, false
	}
	return o.RedeemId, true
}

// HasRedeemId returns a boolean if a field has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) HasRedeemId() bool {
	if o != nil && !common.IsNil(o.RedeemId) {
		return true
	}

	return false
}

// SetRedeemId gets a reference to the given int64 and assigns it to the RedeemId field.
func (o *GetLockedRedemptionRecordResponseRowsInner) SetRedeemId(v int64) {
	o.RedeemId = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetLockedRedemptionRecordResponseRowsInner) SetTime(v int64) {
	o.Time = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetLockedRedemptionRecordResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetLockPeriod returns the LockPeriod field value if set, zero value otherwise.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetLockPeriod() string {
	if o == nil || common.IsNil(o.LockPeriod) {
		var ret string
		return ret
	}
	return *o.LockPeriod
}

// GetLockPeriodOk returns a tuple with the LockPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetLockPeriodOk() (*string, bool) {
	if o == nil || common.IsNil(o.LockPeriod) {
		return nil, false
	}
	return o.LockPeriod, true
}

// HasLockPeriod returns a boolean if a field has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) HasLockPeriod() bool {
	if o != nil && !common.IsNil(o.LockPeriod) {
		return true
	}

	return false
}

// SetLockPeriod gets a reference to the given string and assigns it to the LockPeriod field.
func (o *GetLockedRedemptionRecordResponseRowsInner) SetLockPeriod(v string) {
	o.LockPeriod = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetLockedRedemptionRecordResponseRowsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetOriginalAmount returns the OriginalAmount field value if set, zero value otherwise.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetOriginalAmount() string {
	if o == nil || common.IsNil(o.OriginalAmount) {
		var ret string
		return ret
	}
	return *o.OriginalAmount
}

// GetOriginalAmountOk returns a tuple with the OriginalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetOriginalAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.OriginalAmount) {
		return nil, false
	}
	return o.OriginalAmount, true
}

// HasOriginalAmount returns a boolean if a field has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) HasOriginalAmount() bool {
	if o != nil && !common.IsNil(o.OriginalAmount) {
		return true
	}

	return false
}

// SetOriginalAmount gets a reference to the given string and assigns it to the OriginalAmount field.
func (o *GetLockedRedemptionRecordResponseRowsInner) SetOriginalAmount(v string) {
	o.OriginalAmount = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetLockedRedemptionRecordResponseRowsInner) SetType(v string) {
	o.Type = &v
}

// GetDeliverDate returns the DeliverDate field value if set, zero value otherwise.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetDeliverDate() string {
	if o == nil || common.IsNil(o.DeliverDate) {
		var ret string
		return ret
	}
	return *o.DeliverDate
}

// GetDeliverDateOk returns a tuple with the DeliverDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetDeliverDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.DeliverDate) {
		return nil, false
	}
	return o.DeliverDate, true
}

// HasDeliverDate returns a boolean if a field has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) HasDeliverDate() bool {
	if o != nil && !common.IsNil(o.DeliverDate) {
		return true
	}

	return false
}

// SetDeliverDate gets a reference to the given string and assigns it to the DeliverDate field.
func (o *GetLockedRedemptionRecordResponseRowsInner) SetDeliverDate(v string) {
	o.DeliverDate = &v
}

// GetLossAmount returns the LossAmount field value if set, zero value otherwise.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetLossAmount() string {
	if o == nil || common.IsNil(o.LossAmount) {
		var ret string
		return ret
	}
	return *o.LossAmount
}

// GetLossAmountOk returns a tuple with the LossAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetLossAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.LossAmount) {
		return nil, false
	}
	return o.LossAmount, true
}

// HasLossAmount returns a boolean if a field has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) HasLossAmount() bool {
	if o != nil && !common.IsNil(o.LossAmount) {
		return true
	}

	return false
}

// SetLossAmount gets a reference to the given string and assigns it to the LossAmount field.
func (o *GetLockedRedemptionRecordResponseRowsInner) SetLossAmount(v string) {
	o.LossAmount = &v
}

// GetIsComplete returns the IsComplete field value if set, zero value otherwise.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetIsComplete() bool {
	if o == nil || common.IsNil(o.IsComplete) {
		var ret bool
		return ret
	}
	return *o.IsComplete
}

// GetIsCompleteOk returns a tuple with the IsComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetIsCompleteOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsComplete) {
		return nil, false
	}
	return o.IsComplete, true
}

// HasIsComplete returns a boolean if a field has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) HasIsComplete() bool {
	if o != nil && !common.IsNil(o.IsComplete) {
		return true
	}

	return false
}

// SetIsComplete gets a reference to the given bool and assigns it to the IsComplete field.
func (o *GetLockedRedemptionRecordResponseRowsInner) SetIsComplete(v bool) {
	o.IsComplete = &v
}

// GetRewardAsset returns the RewardAsset field value if set, zero value otherwise.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetRewardAsset() string {
	if o == nil || common.IsNil(o.RewardAsset) {
		var ret string
		return ret
	}
	return *o.RewardAsset
}

// GetRewardAssetOk returns a tuple with the RewardAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetRewardAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.RewardAsset) {
		return nil, false
	}
	return o.RewardAsset, true
}

// HasRewardAsset returns a boolean if a field has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) HasRewardAsset() bool {
	if o != nil && !common.IsNil(o.RewardAsset) {
		return true
	}

	return false
}

// SetRewardAsset gets a reference to the given string and assigns it to the RewardAsset field.
func (o *GetLockedRedemptionRecordResponseRowsInner) SetRewardAsset(v string) {
	o.RewardAsset = &v
}

// GetRewardAmt returns the RewardAmt field value if set, zero value otherwise.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetRewardAmt() string {
	if o == nil || common.IsNil(o.RewardAmt) {
		var ret string
		return ret
	}
	return *o.RewardAmt
}

// GetRewardAmtOk returns a tuple with the RewardAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetRewardAmtOk() (*string, bool) {
	if o == nil || common.IsNil(o.RewardAmt) {
		return nil, false
	}
	return o.RewardAmt, true
}

// HasRewardAmt returns a boolean if a field has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) HasRewardAmt() bool {
	if o != nil && !common.IsNil(o.RewardAmt) {
		return true
	}

	return false
}

// SetRewardAmt gets a reference to the given string and assigns it to the RewardAmt field.
func (o *GetLockedRedemptionRecordResponseRowsInner) SetRewardAmt(v string) {
	o.RewardAmt = &v
}

// GetExtraRewardAsset returns the ExtraRewardAsset field value if set, zero value otherwise.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetExtraRewardAsset() string {
	if o == nil || common.IsNil(o.ExtraRewardAsset) {
		var ret string
		return ret
	}
	return *o.ExtraRewardAsset
}

// GetExtraRewardAssetOk returns a tuple with the ExtraRewardAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetExtraRewardAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExtraRewardAsset) {
		return nil, false
	}
	return o.ExtraRewardAsset, true
}

// HasExtraRewardAsset returns a boolean if a field has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) HasExtraRewardAsset() bool {
	if o != nil && !common.IsNil(o.ExtraRewardAsset) {
		return true
	}

	return false
}

// SetExtraRewardAsset gets a reference to the given string and assigns it to the ExtraRewardAsset field.
func (o *GetLockedRedemptionRecordResponseRowsInner) SetExtraRewardAsset(v string) {
	o.ExtraRewardAsset = &v
}

// GetEstExtraRewardAmt returns the EstExtraRewardAmt field value if set, zero value otherwise.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetEstExtraRewardAmt() string {
	if o == nil || common.IsNil(o.EstExtraRewardAmt) {
		var ret string
		return ret
	}
	return *o.EstExtraRewardAmt
}

// GetEstExtraRewardAmtOk returns a tuple with the EstExtraRewardAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetEstExtraRewardAmtOk() (*string, bool) {
	if o == nil || common.IsNil(o.EstExtraRewardAmt) {
		return nil, false
	}
	return o.EstExtraRewardAmt, true
}

// HasEstExtraRewardAmt returns a boolean if a field has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) HasEstExtraRewardAmt() bool {
	if o != nil && !common.IsNil(o.EstExtraRewardAmt) {
		return true
	}

	return false
}

// SetEstExtraRewardAmt gets a reference to the given string and assigns it to the EstExtraRewardAmt field.
func (o *GetLockedRedemptionRecordResponseRowsInner) SetEstExtraRewardAmt(v string) {
	o.EstExtraRewardAmt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetLockedRedemptionRecordResponseRowsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetLockedRedemptionRecordResponseRowsInner) SetStatus(v string) {
	o.Status = &v
}

func (o GetLockedRedemptionRecordResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLockedRedemptionRecordResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.PositionId) {
		toSerialize["positionId"] = o.PositionId
	}
	if !common.IsNil(o.RedeemId) {
		toSerialize["redeemId"] = o.RedeemId
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.LockPeriod) {
		toSerialize["lockPeriod"] = o.LockPeriod
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.OriginalAmount) {
		toSerialize["originalAmount"] = o.OriginalAmount
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.DeliverDate) {
		toSerialize["deliverDate"] = o.DeliverDate
	}
	if !common.IsNil(o.LossAmount) {
		toSerialize["lossAmount"] = o.LossAmount
	}
	if !common.IsNil(o.IsComplete) {
		toSerialize["isComplete"] = o.IsComplete
	}
	if !common.IsNil(o.RewardAsset) {
		toSerialize["rewardAsset"] = o.RewardAsset
	}
	if !common.IsNil(o.RewardAmt) {
		toSerialize["rewardAmt"] = o.RewardAmt
	}
	if !common.IsNil(o.ExtraRewardAsset) {
		toSerialize["extraRewardAsset"] = o.ExtraRewardAsset
	}
	if !common.IsNil(o.EstExtraRewardAmt) {
		toSerialize["estExtraRewardAmt"] = o.EstExtraRewardAmt
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetLockedRedemptionRecordResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetLockedRedemptionRecordResponseRowsInner := _GetLockedRedemptionRecordResponseRowsInner{}

	err = json.Unmarshal(data, &varGetLockedRedemptionRecordResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetLockedRedemptionRecordResponseRowsInner(varGetLockedRedemptionRecordResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "positionId")
		delete(additionalProperties, "redeemId")
		delete(additionalProperties, "time")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "lockPeriod")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "originalAmount")
		delete(additionalProperties, "type")
		delete(additionalProperties, "deliverDate")
		delete(additionalProperties, "lossAmount")
		delete(additionalProperties, "isComplete")
		delete(additionalProperties, "rewardAsset")
		delete(additionalProperties, "rewardAmt")
		delete(additionalProperties, "extraRewardAsset")
		delete(additionalProperties, "estExtraRewardAmt")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetLockedRedemptionRecordResponseRowsInner struct {
	value *GetLockedRedemptionRecordResponseRowsInner
	isSet bool
}

func (v NullableGetLockedRedemptionRecordResponseRowsInner) Get() *GetLockedRedemptionRecordResponseRowsInner {
	return v.value
}

func (v *NullableGetLockedRedemptionRecordResponseRowsInner) Set(val *GetLockedRedemptionRecordResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLockedRedemptionRecordResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLockedRedemptionRecordResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLockedRedemptionRecordResponseRowsInner(val *GetLockedRedemptionRecordResponseRowsInner) *NullableGetLockedRedemptionRecordResponseRowsInner {
	return &NullableGetLockedRedemptionRecordResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetLockedRedemptionRecordResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLockedRedemptionRecordResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
