/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetOnChainYieldsLockedProductPositionResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetOnChainYieldsLockedProductPositionResponseRowsInner{}

// GetOnChainYieldsLockedProductPositionResponseRowsInner struct for GetOnChainYieldsLockedProductPositionResponseRowsInner
type GetOnChainYieldsLockedProductPositionResponseRowsInner struct {
	PositionId           *string `json:"positionId,omitempty"`
	ProjectId            *string `json:"projectId,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	PurchaseTime         *string `json:"purchaseTime,omitempty"`
	Duration             *string `json:"duration,omitempty"`
	AccrualDays          *string `json:"accrualDays,omitempty"`
	RewardAsset          *string `json:"rewardAsset,omitempty"`
	APY                  *string `json:"APY,omitempty"`
	RewardAmt            *string `json:"rewardAmt,omitempty"`
	NextPay              *string `json:"nextPay,omitempty"`
	NextPayDate          *string `json:"nextPayDate,omitempty"`
	PayPeriod            *string `json:"payPeriod,omitempty"`
	RewardsPayDate       *string `json:"rewardsPayDate,omitempty"`
	RewardsEndDate       *string `json:"rewardsEndDate,omitempty"`
	DeliverDate          *string `json:"deliverDate,omitempty"`
	NextSubscriptionDate *string `json:"nextSubscriptionDate,omitempty"`
	RedeemingAmt         *string `json:"redeemingAmt,omitempty"`
	RedeemTo             *string `json:"redeemTo,omitempty"`
	CanRedeemEarly       *bool   `json:"canRedeemEarly,omitempty"`
	AutoSubscribe        *bool   `json:"autoSubscribe,omitempty"`
	Type                 *string `json:"type,omitempty"`
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetOnChainYieldsLockedProductPositionResponseRowsInner GetOnChainYieldsLockedProductPositionResponseRowsInner

// NewGetOnChainYieldsLockedProductPositionResponseRowsInner instantiates a new GetOnChainYieldsLockedProductPositionResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOnChainYieldsLockedProductPositionResponseRowsInner() *GetOnChainYieldsLockedProductPositionResponseRowsInner {
	this := GetOnChainYieldsLockedProductPositionResponseRowsInner{}
	return &this
}

// NewGetOnChainYieldsLockedProductPositionResponseRowsInnerWithDefaults instantiates a new GetOnChainYieldsLockedProductPositionResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOnChainYieldsLockedProductPositionResponseRowsInnerWithDefaults() *GetOnChainYieldsLockedProductPositionResponseRowsInner {
	this := GetOnChainYieldsLockedProductPositionResponseRowsInner{}
	return &this
}

// GetPositionId returns the PositionId field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetPositionId() string {
	if o == nil || common.IsNil(o.PositionId) {
		var ret string
		return ret
	}
	return *o.PositionId
}

// GetPositionIdOk returns a tuple with the PositionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetPositionIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionId) {
		return nil, false
	}
	return o.PositionId, true
}

// HasPositionId returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) HasPositionId() bool {
	if o != nil && !common.IsNil(o.PositionId) {
		return true
	}

	return false
}

// SetPositionId gets a reference to the given string and assigns it to the PositionId field.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) SetPositionId(v string) {
	o.PositionId = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetProjectId() string {
	if o == nil || common.IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetProjectIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) HasProjectId() bool {
	if o != nil && !common.IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetPurchaseTime returns the PurchaseTime field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetPurchaseTime() string {
	if o == nil || common.IsNil(o.PurchaseTime) {
		var ret string
		return ret
	}
	return *o.PurchaseTime
}

// GetPurchaseTimeOk returns a tuple with the PurchaseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetPurchaseTimeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PurchaseTime) {
		return nil, false
	}
	return o.PurchaseTime, true
}

// HasPurchaseTime returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) HasPurchaseTime() bool {
	if o != nil && !common.IsNil(o.PurchaseTime) {
		return true
	}

	return false
}

// SetPurchaseTime gets a reference to the given string and assigns it to the PurchaseTime field.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) SetPurchaseTime(v string) {
	o.PurchaseTime = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetDuration() string {
	if o == nil || common.IsNil(o.Duration) {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetDurationOk() (*string, bool) {
	if o == nil || common.IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) HasDuration() bool {
	if o != nil && !common.IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) SetDuration(v string) {
	o.Duration = &v
}

// GetAccrualDays returns the AccrualDays field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetAccrualDays() string {
	if o == nil || common.IsNil(o.AccrualDays) {
		var ret string
		return ret
	}
	return *o.AccrualDays
}

// GetAccrualDaysOk returns a tuple with the AccrualDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetAccrualDaysOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccrualDays) {
		return nil, false
	}
	return o.AccrualDays, true
}

// HasAccrualDays returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) HasAccrualDays() bool {
	if o != nil && !common.IsNil(o.AccrualDays) {
		return true
	}

	return false
}

// SetAccrualDays gets a reference to the given string and assigns it to the AccrualDays field.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) SetAccrualDays(v string) {
	o.AccrualDays = &v
}

// GetRewardAsset returns the RewardAsset field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetRewardAsset() string {
	if o == nil || common.IsNil(o.RewardAsset) {
		var ret string
		return ret
	}
	return *o.RewardAsset
}

// GetRewardAssetOk returns a tuple with the RewardAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetRewardAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.RewardAsset) {
		return nil, false
	}
	return o.RewardAsset, true
}

// HasRewardAsset returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) HasRewardAsset() bool {
	if o != nil && !common.IsNil(o.RewardAsset) {
		return true
	}

	return false
}

// SetRewardAsset gets a reference to the given string and assigns it to the RewardAsset field.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) SetRewardAsset(v string) {
	o.RewardAsset = &v
}

// GetAPY returns the APY field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetAPY() string {
	if o == nil || common.IsNil(o.APY) {
		var ret string
		return ret
	}
	return *o.APY
}

// GetAPYOk returns a tuple with the APY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetAPYOk() (*string, bool) {
	if o == nil || common.IsNil(o.APY) {
		return nil, false
	}
	return o.APY, true
}

// HasAPY returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) HasAPY() bool {
	if o != nil && !common.IsNil(o.APY) {
		return true
	}

	return false
}

// SetAPY gets a reference to the given string and assigns it to the APY field.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) SetAPY(v string) {
	o.APY = &v
}

// GetRewardAmt returns the RewardAmt field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetRewardAmt() string {
	if o == nil || common.IsNil(o.RewardAmt) {
		var ret string
		return ret
	}
	return *o.RewardAmt
}

// GetRewardAmtOk returns a tuple with the RewardAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetRewardAmtOk() (*string, bool) {
	if o == nil || common.IsNil(o.RewardAmt) {
		return nil, false
	}
	return o.RewardAmt, true
}

// HasRewardAmt returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) HasRewardAmt() bool {
	if o != nil && !common.IsNil(o.RewardAmt) {
		return true
	}

	return false
}

// SetRewardAmt gets a reference to the given string and assigns it to the RewardAmt field.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) SetRewardAmt(v string) {
	o.RewardAmt = &v
}

// GetNextPay returns the NextPay field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetNextPay() string {
	if o == nil || common.IsNil(o.NextPay) {
		var ret string
		return ret
	}
	return *o.NextPay
}

// GetNextPayOk returns a tuple with the NextPay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetNextPayOk() (*string, bool) {
	if o == nil || common.IsNil(o.NextPay) {
		return nil, false
	}
	return o.NextPay, true
}

// HasNextPay returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) HasNextPay() bool {
	if o != nil && !common.IsNil(o.NextPay) {
		return true
	}

	return false
}

// SetNextPay gets a reference to the given string and assigns it to the NextPay field.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) SetNextPay(v string) {
	o.NextPay = &v
}

// GetNextPayDate returns the NextPayDate field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetNextPayDate() string {
	if o == nil || common.IsNil(o.NextPayDate) {
		var ret string
		return ret
	}
	return *o.NextPayDate
}

// GetNextPayDateOk returns a tuple with the NextPayDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetNextPayDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.NextPayDate) {
		return nil, false
	}
	return o.NextPayDate, true
}

// HasNextPayDate returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) HasNextPayDate() bool {
	if o != nil && !common.IsNil(o.NextPayDate) {
		return true
	}

	return false
}

// SetNextPayDate gets a reference to the given string and assigns it to the NextPayDate field.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) SetNextPayDate(v string) {
	o.NextPayDate = &v
}

// GetPayPeriod returns the PayPeriod field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetPayPeriod() string {
	if o == nil || common.IsNil(o.PayPeriod) {
		var ret string
		return ret
	}
	return *o.PayPeriod
}

// GetPayPeriodOk returns a tuple with the PayPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetPayPeriodOk() (*string, bool) {
	if o == nil || common.IsNil(o.PayPeriod) {
		return nil, false
	}
	return o.PayPeriod, true
}

// HasPayPeriod returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) HasPayPeriod() bool {
	if o != nil && !common.IsNil(o.PayPeriod) {
		return true
	}

	return false
}

// SetPayPeriod gets a reference to the given string and assigns it to the PayPeriod field.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) SetPayPeriod(v string) {
	o.PayPeriod = &v
}

// GetRewardsPayDate returns the RewardsPayDate field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetRewardsPayDate() string {
	if o == nil || common.IsNil(o.RewardsPayDate) {
		var ret string
		return ret
	}
	return *o.RewardsPayDate
}

// GetRewardsPayDateOk returns a tuple with the RewardsPayDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetRewardsPayDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.RewardsPayDate) {
		return nil, false
	}
	return o.RewardsPayDate, true
}

// HasRewardsPayDate returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) HasRewardsPayDate() bool {
	if o != nil && !common.IsNil(o.RewardsPayDate) {
		return true
	}

	return false
}

// SetRewardsPayDate gets a reference to the given string and assigns it to the RewardsPayDate field.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) SetRewardsPayDate(v string) {
	o.RewardsPayDate = &v
}

// GetRewardsEndDate returns the RewardsEndDate field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetRewardsEndDate() string {
	if o == nil || common.IsNil(o.RewardsEndDate) {
		var ret string
		return ret
	}
	return *o.RewardsEndDate
}

// GetRewardsEndDateOk returns a tuple with the RewardsEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetRewardsEndDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.RewardsEndDate) {
		return nil, false
	}
	return o.RewardsEndDate, true
}

// HasRewardsEndDate returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) HasRewardsEndDate() bool {
	if o != nil && !common.IsNil(o.RewardsEndDate) {
		return true
	}

	return false
}

// SetRewardsEndDate gets a reference to the given string and assigns it to the RewardsEndDate field.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) SetRewardsEndDate(v string) {
	o.RewardsEndDate = &v
}

// GetDeliverDate returns the DeliverDate field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetDeliverDate() string {
	if o == nil || common.IsNil(o.DeliverDate) {
		var ret string
		return ret
	}
	return *o.DeliverDate
}

// GetDeliverDateOk returns a tuple with the DeliverDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetDeliverDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.DeliverDate) {
		return nil, false
	}
	return o.DeliverDate, true
}

// HasDeliverDate returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) HasDeliverDate() bool {
	if o != nil && !common.IsNil(o.DeliverDate) {
		return true
	}

	return false
}

// SetDeliverDate gets a reference to the given string and assigns it to the DeliverDate field.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) SetDeliverDate(v string) {
	o.DeliverDate = &v
}

// GetNextSubscriptionDate returns the NextSubscriptionDate field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetNextSubscriptionDate() string {
	if o == nil || common.IsNil(o.NextSubscriptionDate) {
		var ret string
		return ret
	}
	return *o.NextSubscriptionDate
}

// GetNextSubscriptionDateOk returns a tuple with the NextSubscriptionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetNextSubscriptionDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.NextSubscriptionDate) {
		return nil, false
	}
	return o.NextSubscriptionDate, true
}

// HasNextSubscriptionDate returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) HasNextSubscriptionDate() bool {
	if o != nil && !common.IsNil(o.NextSubscriptionDate) {
		return true
	}

	return false
}

// SetNextSubscriptionDate gets a reference to the given string and assigns it to the NextSubscriptionDate field.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) SetNextSubscriptionDate(v string) {
	o.NextSubscriptionDate = &v
}

// GetRedeemingAmt returns the RedeemingAmt field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetRedeemingAmt() string {
	if o == nil || common.IsNil(o.RedeemingAmt) {
		var ret string
		return ret
	}
	return *o.RedeemingAmt
}

// GetRedeemingAmtOk returns a tuple with the RedeemingAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetRedeemingAmtOk() (*string, bool) {
	if o == nil || common.IsNil(o.RedeemingAmt) {
		return nil, false
	}
	return o.RedeemingAmt, true
}

// HasRedeemingAmt returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) HasRedeemingAmt() bool {
	if o != nil && !common.IsNil(o.RedeemingAmt) {
		return true
	}

	return false
}

// SetRedeemingAmt gets a reference to the given string and assigns it to the RedeemingAmt field.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) SetRedeemingAmt(v string) {
	o.RedeemingAmt = &v
}

// GetRedeemTo returns the RedeemTo field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetRedeemTo() string {
	if o == nil || common.IsNil(o.RedeemTo) {
		var ret string
		return ret
	}
	return *o.RedeemTo
}

// GetRedeemToOk returns a tuple with the RedeemTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetRedeemToOk() (*string, bool) {
	if o == nil || common.IsNil(o.RedeemTo) {
		return nil, false
	}
	return o.RedeemTo, true
}

// HasRedeemTo returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) HasRedeemTo() bool {
	if o != nil && !common.IsNil(o.RedeemTo) {
		return true
	}

	return false
}

// SetRedeemTo gets a reference to the given string and assigns it to the RedeemTo field.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) SetRedeemTo(v string) {
	o.RedeemTo = &v
}

// GetCanRedeemEarly returns the CanRedeemEarly field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetCanRedeemEarly() bool {
	if o == nil || common.IsNil(o.CanRedeemEarly) {
		var ret bool
		return ret
	}
	return *o.CanRedeemEarly
}

// GetCanRedeemEarlyOk returns a tuple with the CanRedeemEarly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetCanRedeemEarlyOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanRedeemEarly) {
		return nil, false
	}
	return o.CanRedeemEarly, true
}

// HasCanRedeemEarly returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) HasCanRedeemEarly() bool {
	if o != nil && !common.IsNil(o.CanRedeemEarly) {
		return true
	}

	return false
}

// SetCanRedeemEarly gets a reference to the given bool and assigns it to the CanRedeemEarly field.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) SetCanRedeemEarly(v bool) {
	o.CanRedeemEarly = &v
}

// GetAutoSubscribe returns the AutoSubscribe field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetAutoSubscribe() bool {
	if o == nil || common.IsNil(o.AutoSubscribe) {
		var ret bool
		return ret
	}
	return *o.AutoSubscribe
}

// GetAutoSubscribeOk returns a tuple with the AutoSubscribe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetAutoSubscribeOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AutoSubscribe) {
		return nil, false
	}
	return o.AutoSubscribe, true
}

// HasAutoSubscribe returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) HasAutoSubscribe() bool {
	if o != nil && !common.IsNil(o.AutoSubscribe) {
		return true
	}

	return false
}

// SetAutoSubscribe gets a reference to the given bool and assigns it to the AutoSubscribe field.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) SetAutoSubscribe(v bool) {
	o.AutoSubscribe = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) SetType(v string) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) SetStatus(v string) {
	o.Status = &v
}

func (o GetOnChainYieldsLockedProductPositionResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOnChainYieldsLockedProductPositionResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.PositionId) {
		toSerialize["positionId"] = o.PositionId
	}
	if !common.IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.PurchaseTime) {
		toSerialize["purchaseTime"] = o.PurchaseTime
	}
	if !common.IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !common.IsNil(o.AccrualDays) {
		toSerialize["accrualDays"] = o.AccrualDays
	}
	if !common.IsNil(o.RewardAsset) {
		toSerialize["rewardAsset"] = o.RewardAsset
	}
	if !common.IsNil(o.APY) {
		toSerialize["APY"] = o.APY
	}
	if !common.IsNil(o.RewardAmt) {
		toSerialize["rewardAmt"] = o.RewardAmt
	}
	if !common.IsNil(o.NextPay) {
		toSerialize["nextPay"] = o.NextPay
	}
	if !common.IsNil(o.NextPayDate) {
		toSerialize["nextPayDate"] = o.NextPayDate
	}
	if !common.IsNil(o.PayPeriod) {
		toSerialize["payPeriod"] = o.PayPeriod
	}
	if !common.IsNil(o.RewardsPayDate) {
		toSerialize["rewardsPayDate"] = o.RewardsPayDate
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
	if !common.IsNil(o.RedeemingAmt) {
		toSerialize["redeemingAmt"] = o.RedeemingAmt
	}
	if !common.IsNil(o.RedeemTo) {
		toSerialize["redeemTo"] = o.RedeemTo
	}
	if !common.IsNil(o.CanRedeemEarly) {
		toSerialize["canRedeemEarly"] = o.CanRedeemEarly
	}
	if !common.IsNil(o.AutoSubscribe) {
		toSerialize["autoSubscribe"] = o.AutoSubscribe
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetOnChainYieldsLockedProductPositionResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetOnChainYieldsLockedProductPositionResponseRowsInner := _GetOnChainYieldsLockedProductPositionResponseRowsInner{}

	err = json.Unmarshal(data, &varGetOnChainYieldsLockedProductPositionResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetOnChainYieldsLockedProductPositionResponseRowsInner(varGetOnChainYieldsLockedProductPositionResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "positionId")
		delete(additionalProperties, "projectId")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "purchaseTime")
		delete(additionalProperties, "duration")
		delete(additionalProperties, "accrualDays")
		delete(additionalProperties, "rewardAsset")
		delete(additionalProperties, "APY")
		delete(additionalProperties, "rewardAmt")
		delete(additionalProperties, "nextPay")
		delete(additionalProperties, "nextPayDate")
		delete(additionalProperties, "payPeriod")
		delete(additionalProperties, "rewardsPayDate")
		delete(additionalProperties, "rewardsEndDate")
		delete(additionalProperties, "deliverDate")
		delete(additionalProperties, "nextSubscriptionDate")
		delete(additionalProperties, "redeemingAmt")
		delete(additionalProperties, "redeemTo")
		delete(additionalProperties, "canRedeemEarly")
		delete(additionalProperties, "autoSubscribe")
		delete(additionalProperties, "type")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetOnChainYieldsLockedProductPositionResponseRowsInner struct {
	value *GetOnChainYieldsLockedProductPositionResponseRowsInner
	isSet bool
}

func (v NullableGetOnChainYieldsLockedProductPositionResponseRowsInner) Get() *GetOnChainYieldsLockedProductPositionResponseRowsInner {
	return v.value
}

func (v *NullableGetOnChainYieldsLockedProductPositionResponseRowsInner) Set(val *GetOnChainYieldsLockedProductPositionResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOnChainYieldsLockedProductPositionResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOnChainYieldsLockedProductPositionResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOnChainYieldsLockedProductPositionResponseRowsInner(val *GetOnChainYieldsLockedProductPositionResponseRowsInner) *NullableGetOnChainYieldsLockedProductPositionResponseRowsInner {
	return &NullableGetOnChainYieldsLockedProductPositionResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetOnChainYieldsLockedProductPositionResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOnChainYieldsLockedProductPositionResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
