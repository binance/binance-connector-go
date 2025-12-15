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

// checks if the GetOnChainYieldsLockedRedemptionRecordResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetOnChainYieldsLockedRedemptionRecordResponseRowsInner{}

// GetOnChainYieldsLockedRedemptionRecordResponseRowsInner struct for GetOnChainYieldsLockedRedemptionRecordResponseRowsInner
type GetOnChainYieldsLockedRedemptionRecordResponseRowsInner struct {
	PositionId           *string `json:"positionId,omitempty"`
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
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetOnChainYieldsLockedRedemptionRecordResponseRowsInner GetOnChainYieldsLockedRedemptionRecordResponseRowsInner

// NewGetOnChainYieldsLockedRedemptionRecordResponseRowsInner instantiates a new GetOnChainYieldsLockedRedemptionRecordResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOnChainYieldsLockedRedemptionRecordResponseRowsInner() *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner {
	this := GetOnChainYieldsLockedRedemptionRecordResponseRowsInner{}
	return &this
}

// NewGetOnChainYieldsLockedRedemptionRecordResponseRowsInnerWithDefaults instantiates a new GetOnChainYieldsLockedRedemptionRecordResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOnChainYieldsLockedRedemptionRecordResponseRowsInnerWithDefaults() *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner {
	this := GetOnChainYieldsLockedRedemptionRecordResponseRowsInner{}
	return &this
}

// GetPositionId returns the PositionId field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) GetPositionId() string {
	if o == nil || common.IsNil(o.PositionId) {
		var ret string
		return ret
	}
	return *o.PositionId
}

// GetPositionIdOk returns a tuple with the PositionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) GetPositionIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionId) {
		return nil, false
	}
	return o.PositionId, true
}

// HasPositionId returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) HasPositionId() bool {
	if o != nil && !common.IsNil(o.PositionId) {
		return true
	}

	return false
}

// SetPositionId gets a reference to the given string and assigns it to the PositionId field.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) SetPositionId(v string) {
	o.PositionId = &v
}

// GetRedeemId returns the RedeemId field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) GetRedeemId() int64 {
	if o == nil || common.IsNil(o.RedeemId) {
		var ret int64
		return ret
	}
	return *o.RedeemId
}

// GetRedeemIdOk returns a tuple with the RedeemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) GetRedeemIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.RedeemId) {
		return nil, false
	}
	return o.RedeemId, true
}

// HasRedeemId returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) HasRedeemId() bool {
	if o != nil && !common.IsNil(o.RedeemId) {
		return true
	}

	return false
}

// SetRedeemId gets a reference to the given int64 and assigns it to the RedeemId field.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) SetRedeemId(v int64) {
	o.RedeemId = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) SetTime(v int64) {
	o.Time = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetLockPeriod returns the LockPeriod field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) GetLockPeriod() string {
	if o == nil || common.IsNil(o.LockPeriod) {
		var ret string
		return ret
	}
	return *o.LockPeriod
}

// GetLockPeriodOk returns a tuple with the LockPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) GetLockPeriodOk() (*string, bool) {
	if o == nil || common.IsNil(o.LockPeriod) {
		return nil, false
	}
	return o.LockPeriod, true
}

// HasLockPeriod returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) HasLockPeriod() bool {
	if o != nil && !common.IsNil(o.LockPeriod) {
		return true
	}

	return false
}

// SetLockPeriod gets a reference to the given string and assigns it to the LockPeriod field.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) SetLockPeriod(v string) {
	o.LockPeriod = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetOriginalAmount returns the OriginalAmount field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) GetOriginalAmount() string {
	if o == nil || common.IsNil(o.OriginalAmount) {
		var ret string
		return ret
	}
	return *o.OriginalAmount
}

// GetOriginalAmountOk returns a tuple with the OriginalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) GetOriginalAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.OriginalAmount) {
		return nil, false
	}
	return o.OriginalAmount, true
}

// HasOriginalAmount returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) HasOriginalAmount() bool {
	if o != nil && !common.IsNil(o.OriginalAmount) {
		return true
	}

	return false
}

// SetOriginalAmount gets a reference to the given string and assigns it to the OriginalAmount field.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) SetOriginalAmount(v string) {
	o.OriginalAmount = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) SetType(v string) {
	o.Type = &v
}

// GetDeliverDate returns the DeliverDate field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) GetDeliverDate() string {
	if o == nil || common.IsNil(o.DeliverDate) {
		var ret string
		return ret
	}
	return *o.DeliverDate
}

// GetDeliverDateOk returns a tuple with the DeliverDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) GetDeliverDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.DeliverDate) {
		return nil, false
	}
	return o.DeliverDate, true
}

// HasDeliverDate returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) HasDeliverDate() bool {
	if o != nil && !common.IsNil(o.DeliverDate) {
		return true
	}

	return false
}

// SetDeliverDate gets a reference to the given string and assigns it to the DeliverDate field.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) SetDeliverDate(v string) {
	o.DeliverDate = &v
}

// GetLossAmount returns the LossAmount field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) GetLossAmount() string {
	if o == nil || common.IsNil(o.LossAmount) {
		var ret string
		return ret
	}
	return *o.LossAmount
}

// GetLossAmountOk returns a tuple with the LossAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) GetLossAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.LossAmount) {
		return nil, false
	}
	return o.LossAmount, true
}

// HasLossAmount returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) HasLossAmount() bool {
	if o != nil && !common.IsNil(o.LossAmount) {
		return true
	}

	return false
}

// SetLossAmount gets a reference to the given string and assigns it to the LossAmount field.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) SetLossAmount(v string) {
	o.LossAmount = &v
}

// GetIsComplete returns the IsComplete field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) GetIsComplete() bool {
	if o == nil || common.IsNil(o.IsComplete) {
		var ret bool
		return ret
	}
	return *o.IsComplete
}

// GetIsCompleteOk returns a tuple with the IsComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) GetIsCompleteOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsComplete) {
		return nil, false
	}
	return o.IsComplete, true
}

// HasIsComplete returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) HasIsComplete() bool {
	if o != nil && !common.IsNil(o.IsComplete) {
		return true
	}

	return false
}

// SetIsComplete gets a reference to the given bool and assigns it to the IsComplete field.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) SetIsComplete(v bool) {
	o.IsComplete = &v
}

// GetRewardAsset returns the RewardAsset field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) GetRewardAsset() string {
	if o == nil || common.IsNil(o.RewardAsset) {
		var ret string
		return ret
	}
	return *o.RewardAsset
}

// GetRewardAssetOk returns a tuple with the RewardAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) GetRewardAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.RewardAsset) {
		return nil, false
	}
	return o.RewardAsset, true
}

// HasRewardAsset returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) HasRewardAsset() bool {
	if o != nil && !common.IsNil(o.RewardAsset) {
		return true
	}

	return false
}

// SetRewardAsset gets a reference to the given string and assigns it to the RewardAsset field.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) SetRewardAsset(v string) {
	o.RewardAsset = &v
}

// GetRewardAmt returns the RewardAmt field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) GetRewardAmt() string {
	if o == nil || common.IsNil(o.RewardAmt) {
		var ret string
		return ret
	}
	return *o.RewardAmt
}

// GetRewardAmtOk returns a tuple with the RewardAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) GetRewardAmtOk() (*string, bool) {
	if o == nil || common.IsNil(o.RewardAmt) {
		return nil, false
	}
	return o.RewardAmt, true
}

// HasRewardAmt returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) HasRewardAmt() bool {
	if o != nil && !common.IsNil(o.RewardAmt) {
		return true
	}

	return false
}

// SetRewardAmt gets a reference to the given string and assigns it to the RewardAmt field.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) SetRewardAmt(v string) {
	o.RewardAmt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) SetStatus(v string) {
	o.Status = &v
}

func (o GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetOnChainYieldsLockedRedemptionRecordResponseRowsInner := _GetOnChainYieldsLockedRedemptionRecordResponseRowsInner{}

	err = json.Unmarshal(data, &varGetOnChainYieldsLockedRedemptionRecordResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetOnChainYieldsLockedRedemptionRecordResponseRowsInner(varGetOnChainYieldsLockedRedemptionRecordResponseRowsInner)

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
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetOnChainYieldsLockedRedemptionRecordResponseRowsInner struct {
	value *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner
	isSet bool
}

func (v NullableGetOnChainYieldsLockedRedemptionRecordResponseRowsInner) Get() *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner {
	return v.value
}

func (v *NullableGetOnChainYieldsLockedRedemptionRecordResponseRowsInner) Set(val *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOnChainYieldsLockedRedemptionRecordResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOnChainYieldsLockedRedemptionRecordResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOnChainYieldsLockedRedemptionRecordResponseRowsInner(val *GetOnChainYieldsLockedRedemptionRecordResponseRowsInner) *NullableGetOnChainYieldsLockedRedemptionRecordResponseRowsInner {
	return &NullableGetOnChainYieldsLockedRedemptionRecordResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetOnChainYieldsLockedRedemptionRecordResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOnChainYieldsLockedRedemptionRecordResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
