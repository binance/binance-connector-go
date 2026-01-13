/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetLockedProductPositionResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetLockedProductPositionResponseRowsInner{}

// GetLockedProductPositionResponseRowsInner struct for GetLockedProductPositionResponseRowsInner
type GetLockedProductPositionResponseRowsInner struct {
	PositionId            *int64  `json:"positionId,omitempty"`
	ParentPositionId      *int64  `json:"parentPositionId,omitempty"`
	ProjectId             *string `json:"projectId,omitempty"`
	Asset                 *string `json:"asset,omitempty"`
	Amount                *string `json:"amount,omitempty"`
	PurchaseTime          *string `json:"purchaseTime,omitempty"`
	Duration              *string `json:"duration,omitempty"`
	AccrualDays           *string `json:"accrualDays,omitempty"`
	RewardAsset           *string `json:"rewardAsset,omitempty"`
	APY                   *string `json:"APY,omitempty"`
	RewardAmt             *string `json:"rewardAmt,omitempty"`
	ExtraRewardAsset      *string `json:"extraRewardAsset,omitempty"`
	ExtraRewardAPR        *string `json:"extraRewardAPR,omitempty"`
	EstExtraRewardAmt     *string `json:"estExtraRewardAmt,omitempty"`
	BoostRewardAsset      *string `json:"boostRewardAsset,omitempty"`
	BoostApr              *string `json:"boostApr,omitempty"`
	TotalBoostRewardAmt   *string `json:"totalBoostRewardAmt,omitempty"`
	NextPay               *string `json:"nextPay,omitempty"`
	NextPayDate           *string `json:"nextPayDate,omitempty"`
	PayPeriod             *string `json:"payPeriod,omitempty"`
	RedeemAmountEarly     *string `json:"redeemAmountEarly,omitempty"`
	RewardsEndDate        *string `json:"rewardsEndDate,omitempty"`
	DeliverDate           *string `json:"deliverDate,omitempty"`
	RedeemPeriod          *string `json:"redeemPeriod,omitempty"`
	RedeemingAmt          *string `json:"redeemingAmt,omitempty"`
	RedeemTo              *string `json:"redeemTo,omitempty"`
	PartialAmtDeliverDate *string `json:"partialAmtDeliverDate,omitempty"`
	CanRedeemEarly        *bool   `json:"canRedeemEarly,omitempty"`
	CanFastRedemption     *bool   `json:"canFastRedemption,omitempty"`
	AutoSubscribe         *bool   `json:"autoSubscribe,omitempty"`
	Type                  *string `json:"type,omitempty"`
	Status                *string `json:"status,omitempty"`
	CanReStake            *bool   `json:"canReStake,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _GetLockedProductPositionResponseRowsInner GetLockedProductPositionResponseRowsInner

// NewGetLockedProductPositionResponseRowsInner instantiates a new GetLockedProductPositionResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLockedProductPositionResponseRowsInner() *GetLockedProductPositionResponseRowsInner {
	this := GetLockedProductPositionResponseRowsInner{}
	return &this
}

// NewGetLockedProductPositionResponseRowsInnerWithDefaults instantiates a new GetLockedProductPositionResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLockedProductPositionResponseRowsInnerWithDefaults() *GetLockedProductPositionResponseRowsInner {
	this := GetLockedProductPositionResponseRowsInner{}
	return &this
}

// GetPositionId returns the PositionId field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetPositionId() int64 {
	if o == nil || common.IsNil(o.PositionId) {
		var ret int64
		return ret
	}
	return *o.PositionId
}

// GetPositionIdOk returns a tuple with the PositionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetPositionIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PositionId) {
		return nil, false
	}
	return o.PositionId, true
}

// HasPositionId returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasPositionId() bool {
	if o != nil && !common.IsNil(o.PositionId) {
		return true
	}

	return false
}

// SetPositionId gets a reference to the given int64 and assigns it to the PositionId field.
func (o *GetLockedProductPositionResponseRowsInner) SetPositionId(v int64) {
	o.PositionId = &v
}

// GetParentPositionId returns the ParentPositionId field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetParentPositionId() int64 {
	if o == nil || common.IsNil(o.ParentPositionId) {
		var ret int64
		return ret
	}
	return *o.ParentPositionId
}

// GetParentPositionIdOk returns a tuple with the ParentPositionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetParentPositionIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ParentPositionId) {
		return nil, false
	}
	return o.ParentPositionId, true
}

// HasParentPositionId returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasParentPositionId() bool {
	if o != nil && !common.IsNil(o.ParentPositionId) {
		return true
	}

	return false
}

// SetParentPositionId gets a reference to the given int64 and assigns it to the ParentPositionId field.
func (o *GetLockedProductPositionResponseRowsInner) SetParentPositionId(v int64) {
	o.ParentPositionId = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetProjectId() string {
	if o == nil || common.IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetProjectIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasProjectId() bool {
	if o != nil && !common.IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *GetLockedProductPositionResponseRowsInner) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetLockedProductPositionResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetLockedProductPositionResponseRowsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetPurchaseTime returns the PurchaseTime field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetPurchaseTime() string {
	if o == nil || common.IsNil(o.PurchaseTime) {
		var ret string
		return ret
	}
	return *o.PurchaseTime
}

// GetPurchaseTimeOk returns a tuple with the PurchaseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetPurchaseTimeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PurchaseTime) {
		return nil, false
	}
	return o.PurchaseTime, true
}

// HasPurchaseTime returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasPurchaseTime() bool {
	if o != nil && !common.IsNil(o.PurchaseTime) {
		return true
	}

	return false
}

// SetPurchaseTime gets a reference to the given string and assigns it to the PurchaseTime field.
func (o *GetLockedProductPositionResponseRowsInner) SetPurchaseTime(v string) {
	o.PurchaseTime = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetDuration() string {
	if o == nil || common.IsNil(o.Duration) {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetDurationOk() (*string, bool) {
	if o == nil || common.IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasDuration() bool {
	if o != nil && !common.IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *GetLockedProductPositionResponseRowsInner) SetDuration(v string) {
	o.Duration = &v
}

// GetAccrualDays returns the AccrualDays field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetAccrualDays() string {
	if o == nil || common.IsNil(o.AccrualDays) {
		var ret string
		return ret
	}
	return *o.AccrualDays
}

// GetAccrualDaysOk returns a tuple with the AccrualDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetAccrualDaysOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccrualDays) {
		return nil, false
	}
	return o.AccrualDays, true
}

// HasAccrualDays returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasAccrualDays() bool {
	if o != nil && !common.IsNil(o.AccrualDays) {
		return true
	}

	return false
}

// SetAccrualDays gets a reference to the given string and assigns it to the AccrualDays field.
func (o *GetLockedProductPositionResponseRowsInner) SetAccrualDays(v string) {
	o.AccrualDays = &v
}

// GetRewardAsset returns the RewardAsset field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetRewardAsset() string {
	if o == nil || common.IsNil(o.RewardAsset) {
		var ret string
		return ret
	}
	return *o.RewardAsset
}

// GetRewardAssetOk returns a tuple with the RewardAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetRewardAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.RewardAsset) {
		return nil, false
	}
	return o.RewardAsset, true
}

// HasRewardAsset returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasRewardAsset() bool {
	if o != nil && !common.IsNil(o.RewardAsset) {
		return true
	}

	return false
}

// SetRewardAsset gets a reference to the given string and assigns it to the RewardAsset field.
func (o *GetLockedProductPositionResponseRowsInner) SetRewardAsset(v string) {
	o.RewardAsset = &v
}

// GetAPY returns the APY field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetAPY() string {
	if o == nil || common.IsNil(o.APY) {
		var ret string
		return ret
	}
	return *o.APY
}

// GetAPYOk returns a tuple with the APY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetAPYOk() (*string, bool) {
	if o == nil || common.IsNil(o.APY) {
		return nil, false
	}
	return o.APY, true
}

// HasAPY returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasAPY() bool {
	if o != nil && !common.IsNil(o.APY) {
		return true
	}

	return false
}

// SetAPY gets a reference to the given string and assigns it to the APY field.
func (o *GetLockedProductPositionResponseRowsInner) SetAPY(v string) {
	o.APY = &v
}

// GetRewardAmt returns the RewardAmt field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetRewardAmt() string {
	if o == nil || common.IsNil(o.RewardAmt) {
		var ret string
		return ret
	}
	return *o.RewardAmt
}

// GetRewardAmtOk returns a tuple with the RewardAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetRewardAmtOk() (*string, bool) {
	if o == nil || common.IsNil(o.RewardAmt) {
		return nil, false
	}
	return o.RewardAmt, true
}

// HasRewardAmt returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasRewardAmt() bool {
	if o != nil && !common.IsNil(o.RewardAmt) {
		return true
	}

	return false
}

// SetRewardAmt gets a reference to the given string and assigns it to the RewardAmt field.
func (o *GetLockedProductPositionResponseRowsInner) SetRewardAmt(v string) {
	o.RewardAmt = &v
}

// GetExtraRewardAsset returns the ExtraRewardAsset field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetExtraRewardAsset() string {
	if o == nil || common.IsNil(o.ExtraRewardAsset) {
		var ret string
		return ret
	}
	return *o.ExtraRewardAsset
}

// GetExtraRewardAssetOk returns a tuple with the ExtraRewardAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetExtraRewardAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExtraRewardAsset) {
		return nil, false
	}
	return o.ExtraRewardAsset, true
}

// HasExtraRewardAsset returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasExtraRewardAsset() bool {
	if o != nil && !common.IsNil(o.ExtraRewardAsset) {
		return true
	}

	return false
}

// SetExtraRewardAsset gets a reference to the given string and assigns it to the ExtraRewardAsset field.
func (o *GetLockedProductPositionResponseRowsInner) SetExtraRewardAsset(v string) {
	o.ExtraRewardAsset = &v
}

// GetExtraRewardAPR returns the ExtraRewardAPR field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetExtraRewardAPR() string {
	if o == nil || common.IsNil(o.ExtraRewardAPR) {
		var ret string
		return ret
	}
	return *o.ExtraRewardAPR
}

// GetExtraRewardAPROk returns a tuple with the ExtraRewardAPR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetExtraRewardAPROk() (*string, bool) {
	if o == nil || common.IsNil(o.ExtraRewardAPR) {
		return nil, false
	}
	return o.ExtraRewardAPR, true
}

// HasExtraRewardAPR returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasExtraRewardAPR() bool {
	if o != nil && !common.IsNil(o.ExtraRewardAPR) {
		return true
	}

	return false
}

// SetExtraRewardAPR gets a reference to the given string and assigns it to the ExtraRewardAPR field.
func (o *GetLockedProductPositionResponseRowsInner) SetExtraRewardAPR(v string) {
	o.ExtraRewardAPR = &v
}

// GetEstExtraRewardAmt returns the EstExtraRewardAmt field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetEstExtraRewardAmt() string {
	if o == nil || common.IsNil(o.EstExtraRewardAmt) {
		var ret string
		return ret
	}
	return *o.EstExtraRewardAmt
}

// GetEstExtraRewardAmtOk returns a tuple with the EstExtraRewardAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetEstExtraRewardAmtOk() (*string, bool) {
	if o == nil || common.IsNil(o.EstExtraRewardAmt) {
		return nil, false
	}
	return o.EstExtraRewardAmt, true
}

// HasEstExtraRewardAmt returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasEstExtraRewardAmt() bool {
	if o != nil && !common.IsNil(o.EstExtraRewardAmt) {
		return true
	}

	return false
}

// SetEstExtraRewardAmt gets a reference to the given string and assigns it to the EstExtraRewardAmt field.
func (o *GetLockedProductPositionResponseRowsInner) SetEstExtraRewardAmt(v string) {
	o.EstExtraRewardAmt = &v
}

// GetBoostRewardAsset returns the BoostRewardAsset field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetBoostRewardAsset() string {
	if o == nil || common.IsNil(o.BoostRewardAsset) {
		var ret string
		return ret
	}
	return *o.BoostRewardAsset
}

// GetBoostRewardAssetOk returns a tuple with the BoostRewardAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetBoostRewardAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.BoostRewardAsset) {
		return nil, false
	}
	return o.BoostRewardAsset, true
}

// HasBoostRewardAsset returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasBoostRewardAsset() bool {
	if o != nil && !common.IsNil(o.BoostRewardAsset) {
		return true
	}

	return false
}

// SetBoostRewardAsset gets a reference to the given string and assigns it to the BoostRewardAsset field.
func (o *GetLockedProductPositionResponseRowsInner) SetBoostRewardAsset(v string) {
	o.BoostRewardAsset = &v
}

// GetBoostApr returns the BoostApr field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetBoostApr() string {
	if o == nil || common.IsNil(o.BoostApr) {
		var ret string
		return ret
	}
	return *o.BoostApr
}

// GetBoostAprOk returns a tuple with the BoostApr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetBoostAprOk() (*string, bool) {
	if o == nil || common.IsNil(o.BoostApr) {
		return nil, false
	}
	return o.BoostApr, true
}

// HasBoostApr returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasBoostApr() bool {
	if o != nil && !common.IsNil(o.BoostApr) {
		return true
	}

	return false
}

// SetBoostApr gets a reference to the given string and assigns it to the BoostApr field.
func (o *GetLockedProductPositionResponseRowsInner) SetBoostApr(v string) {
	o.BoostApr = &v
}

// GetTotalBoostRewardAmt returns the TotalBoostRewardAmt field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetTotalBoostRewardAmt() string {
	if o == nil || common.IsNil(o.TotalBoostRewardAmt) {
		var ret string
		return ret
	}
	return *o.TotalBoostRewardAmt
}

// GetTotalBoostRewardAmtOk returns a tuple with the TotalBoostRewardAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetTotalBoostRewardAmtOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalBoostRewardAmt) {
		return nil, false
	}
	return o.TotalBoostRewardAmt, true
}

// HasTotalBoostRewardAmt returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasTotalBoostRewardAmt() bool {
	if o != nil && !common.IsNil(o.TotalBoostRewardAmt) {
		return true
	}

	return false
}

// SetTotalBoostRewardAmt gets a reference to the given string and assigns it to the TotalBoostRewardAmt field.
func (o *GetLockedProductPositionResponseRowsInner) SetTotalBoostRewardAmt(v string) {
	o.TotalBoostRewardAmt = &v
}

// GetNextPay returns the NextPay field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetNextPay() string {
	if o == nil || common.IsNil(o.NextPay) {
		var ret string
		return ret
	}
	return *o.NextPay
}

// GetNextPayOk returns a tuple with the NextPay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetNextPayOk() (*string, bool) {
	if o == nil || common.IsNil(o.NextPay) {
		return nil, false
	}
	return o.NextPay, true
}

// HasNextPay returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasNextPay() bool {
	if o != nil && !common.IsNil(o.NextPay) {
		return true
	}

	return false
}

// SetNextPay gets a reference to the given string and assigns it to the NextPay field.
func (o *GetLockedProductPositionResponseRowsInner) SetNextPay(v string) {
	o.NextPay = &v
}

// GetNextPayDate returns the NextPayDate field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetNextPayDate() string {
	if o == nil || common.IsNil(o.NextPayDate) {
		var ret string
		return ret
	}
	return *o.NextPayDate
}

// GetNextPayDateOk returns a tuple with the NextPayDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetNextPayDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.NextPayDate) {
		return nil, false
	}
	return o.NextPayDate, true
}

// HasNextPayDate returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasNextPayDate() bool {
	if o != nil && !common.IsNil(o.NextPayDate) {
		return true
	}

	return false
}

// SetNextPayDate gets a reference to the given string and assigns it to the NextPayDate field.
func (o *GetLockedProductPositionResponseRowsInner) SetNextPayDate(v string) {
	o.NextPayDate = &v
}

// GetPayPeriod returns the PayPeriod field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetPayPeriod() string {
	if o == nil || common.IsNil(o.PayPeriod) {
		var ret string
		return ret
	}
	return *o.PayPeriod
}

// GetPayPeriodOk returns a tuple with the PayPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetPayPeriodOk() (*string, bool) {
	if o == nil || common.IsNil(o.PayPeriod) {
		return nil, false
	}
	return o.PayPeriod, true
}

// HasPayPeriod returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasPayPeriod() bool {
	if o != nil && !common.IsNil(o.PayPeriod) {
		return true
	}

	return false
}

// SetPayPeriod gets a reference to the given string and assigns it to the PayPeriod field.
func (o *GetLockedProductPositionResponseRowsInner) SetPayPeriod(v string) {
	o.PayPeriod = &v
}

// GetRedeemAmountEarly returns the RedeemAmountEarly field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetRedeemAmountEarly() string {
	if o == nil || common.IsNil(o.RedeemAmountEarly) {
		var ret string
		return ret
	}
	return *o.RedeemAmountEarly
}

// GetRedeemAmountEarlyOk returns a tuple with the RedeemAmountEarly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetRedeemAmountEarlyOk() (*string, bool) {
	if o == nil || common.IsNil(o.RedeemAmountEarly) {
		return nil, false
	}
	return o.RedeemAmountEarly, true
}

// HasRedeemAmountEarly returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasRedeemAmountEarly() bool {
	if o != nil && !common.IsNil(o.RedeemAmountEarly) {
		return true
	}

	return false
}

// SetRedeemAmountEarly gets a reference to the given string and assigns it to the RedeemAmountEarly field.
func (o *GetLockedProductPositionResponseRowsInner) SetRedeemAmountEarly(v string) {
	o.RedeemAmountEarly = &v
}

// GetRewardsEndDate returns the RewardsEndDate field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetRewardsEndDate() string {
	if o == nil || common.IsNil(o.RewardsEndDate) {
		var ret string
		return ret
	}
	return *o.RewardsEndDate
}

// GetRewardsEndDateOk returns a tuple with the RewardsEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetRewardsEndDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.RewardsEndDate) {
		return nil, false
	}
	return o.RewardsEndDate, true
}

// HasRewardsEndDate returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasRewardsEndDate() bool {
	if o != nil && !common.IsNil(o.RewardsEndDate) {
		return true
	}

	return false
}

// SetRewardsEndDate gets a reference to the given string and assigns it to the RewardsEndDate field.
func (o *GetLockedProductPositionResponseRowsInner) SetRewardsEndDate(v string) {
	o.RewardsEndDate = &v
}

// GetDeliverDate returns the DeliverDate field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetDeliverDate() string {
	if o == nil || common.IsNil(o.DeliverDate) {
		var ret string
		return ret
	}
	return *o.DeliverDate
}

// GetDeliverDateOk returns a tuple with the DeliverDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetDeliverDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.DeliverDate) {
		return nil, false
	}
	return o.DeliverDate, true
}

// HasDeliverDate returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasDeliverDate() bool {
	if o != nil && !common.IsNil(o.DeliverDate) {
		return true
	}

	return false
}

// SetDeliverDate gets a reference to the given string and assigns it to the DeliverDate field.
func (o *GetLockedProductPositionResponseRowsInner) SetDeliverDate(v string) {
	o.DeliverDate = &v
}

// GetRedeemPeriod returns the RedeemPeriod field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetRedeemPeriod() string {
	if o == nil || common.IsNil(o.RedeemPeriod) {
		var ret string
		return ret
	}
	return *o.RedeemPeriod
}

// GetRedeemPeriodOk returns a tuple with the RedeemPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetRedeemPeriodOk() (*string, bool) {
	if o == nil || common.IsNil(o.RedeemPeriod) {
		return nil, false
	}
	return o.RedeemPeriod, true
}

// HasRedeemPeriod returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasRedeemPeriod() bool {
	if o != nil && !common.IsNil(o.RedeemPeriod) {
		return true
	}

	return false
}

// SetRedeemPeriod gets a reference to the given string and assigns it to the RedeemPeriod field.
func (o *GetLockedProductPositionResponseRowsInner) SetRedeemPeriod(v string) {
	o.RedeemPeriod = &v
}

// GetRedeemingAmt returns the RedeemingAmt field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetRedeemingAmt() string {
	if o == nil || common.IsNil(o.RedeemingAmt) {
		var ret string
		return ret
	}
	return *o.RedeemingAmt
}

// GetRedeemingAmtOk returns a tuple with the RedeemingAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetRedeemingAmtOk() (*string, bool) {
	if o == nil || common.IsNil(o.RedeemingAmt) {
		return nil, false
	}
	return o.RedeemingAmt, true
}

// HasRedeemingAmt returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasRedeemingAmt() bool {
	if o != nil && !common.IsNil(o.RedeemingAmt) {
		return true
	}

	return false
}

// SetRedeemingAmt gets a reference to the given string and assigns it to the RedeemingAmt field.
func (o *GetLockedProductPositionResponseRowsInner) SetRedeemingAmt(v string) {
	o.RedeemingAmt = &v
}

// GetRedeemTo returns the RedeemTo field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetRedeemTo() string {
	if o == nil || common.IsNil(o.RedeemTo) {
		var ret string
		return ret
	}
	return *o.RedeemTo
}

// GetRedeemToOk returns a tuple with the RedeemTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetRedeemToOk() (*string, bool) {
	if o == nil || common.IsNil(o.RedeemTo) {
		return nil, false
	}
	return o.RedeemTo, true
}

// HasRedeemTo returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasRedeemTo() bool {
	if o != nil && !common.IsNil(o.RedeemTo) {
		return true
	}

	return false
}

// SetRedeemTo gets a reference to the given string and assigns it to the RedeemTo field.
func (o *GetLockedProductPositionResponseRowsInner) SetRedeemTo(v string) {
	o.RedeemTo = &v
}

// GetPartialAmtDeliverDate returns the PartialAmtDeliverDate field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetPartialAmtDeliverDate() string {
	if o == nil || common.IsNil(o.PartialAmtDeliverDate) {
		var ret string
		return ret
	}
	return *o.PartialAmtDeliverDate
}

// GetPartialAmtDeliverDateOk returns a tuple with the PartialAmtDeliverDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetPartialAmtDeliverDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.PartialAmtDeliverDate) {
		return nil, false
	}
	return o.PartialAmtDeliverDate, true
}

// HasPartialAmtDeliverDate returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasPartialAmtDeliverDate() bool {
	if o != nil && !common.IsNil(o.PartialAmtDeliverDate) {
		return true
	}

	return false
}

// SetPartialAmtDeliverDate gets a reference to the given string and assigns it to the PartialAmtDeliverDate field.
func (o *GetLockedProductPositionResponseRowsInner) SetPartialAmtDeliverDate(v string) {
	o.PartialAmtDeliverDate = &v
}

// GetCanRedeemEarly returns the CanRedeemEarly field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetCanRedeemEarly() bool {
	if o == nil || common.IsNil(o.CanRedeemEarly) {
		var ret bool
		return ret
	}
	return *o.CanRedeemEarly
}

// GetCanRedeemEarlyOk returns a tuple with the CanRedeemEarly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetCanRedeemEarlyOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanRedeemEarly) {
		return nil, false
	}
	return o.CanRedeemEarly, true
}

// HasCanRedeemEarly returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasCanRedeemEarly() bool {
	if o != nil && !common.IsNil(o.CanRedeemEarly) {
		return true
	}

	return false
}

// SetCanRedeemEarly gets a reference to the given bool and assigns it to the CanRedeemEarly field.
func (o *GetLockedProductPositionResponseRowsInner) SetCanRedeemEarly(v bool) {
	o.CanRedeemEarly = &v
}

// GetCanFastRedemption returns the CanFastRedemption field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetCanFastRedemption() bool {
	if o == nil || common.IsNil(o.CanFastRedemption) {
		var ret bool
		return ret
	}
	return *o.CanFastRedemption
}

// GetCanFastRedemptionOk returns a tuple with the CanFastRedemption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetCanFastRedemptionOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanFastRedemption) {
		return nil, false
	}
	return o.CanFastRedemption, true
}

// HasCanFastRedemption returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasCanFastRedemption() bool {
	if o != nil && !common.IsNil(o.CanFastRedemption) {
		return true
	}

	return false
}

// SetCanFastRedemption gets a reference to the given bool and assigns it to the CanFastRedemption field.
func (o *GetLockedProductPositionResponseRowsInner) SetCanFastRedemption(v bool) {
	o.CanFastRedemption = &v
}

// GetAutoSubscribe returns the AutoSubscribe field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetAutoSubscribe() bool {
	if o == nil || common.IsNil(o.AutoSubscribe) {
		var ret bool
		return ret
	}
	return *o.AutoSubscribe
}

// GetAutoSubscribeOk returns a tuple with the AutoSubscribe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetAutoSubscribeOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AutoSubscribe) {
		return nil, false
	}
	return o.AutoSubscribe, true
}

// HasAutoSubscribe returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasAutoSubscribe() bool {
	if o != nil && !common.IsNil(o.AutoSubscribe) {
		return true
	}

	return false
}

// SetAutoSubscribe gets a reference to the given bool and assigns it to the AutoSubscribe field.
func (o *GetLockedProductPositionResponseRowsInner) SetAutoSubscribe(v bool) {
	o.AutoSubscribe = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetLockedProductPositionResponseRowsInner) SetType(v string) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetLockedProductPositionResponseRowsInner) SetStatus(v string) {
	o.Status = &v
}

// GetCanReStake returns the CanReStake field value if set, zero value otherwise.
func (o *GetLockedProductPositionResponseRowsInner) GetCanReStake() bool {
	if o == nil || common.IsNil(o.CanReStake) {
		var ret bool
		return ret
	}
	return *o.CanReStake
}

// GetCanReStakeOk returns a tuple with the CanReStake field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLockedProductPositionResponseRowsInner) GetCanReStakeOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanReStake) {
		return nil, false
	}
	return o.CanReStake, true
}

// HasCanReStake returns a boolean if a field has been set.
func (o *GetLockedProductPositionResponseRowsInner) HasCanReStake() bool {
	if o != nil && !common.IsNil(o.CanReStake) {
		return true
	}

	return false
}

// SetCanReStake gets a reference to the given bool and assigns it to the CanReStake field.
func (o *GetLockedProductPositionResponseRowsInner) SetCanReStake(v bool) {
	o.CanReStake = &v
}

func (o GetLockedProductPositionResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLockedProductPositionResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.PositionId) {
		toSerialize["positionId"] = o.PositionId
	}
	if !common.IsNil(o.ParentPositionId) {
		toSerialize["parentPositionId"] = o.ParentPositionId
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
	if !common.IsNil(o.ExtraRewardAsset) {
		toSerialize["extraRewardAsset"] = o.ExtraRewardAsset
	}
	if !common.IsNil(o.ExtraRewardAPR) {
		toSerialize["extraRewardAPR"] = o.ExtraRewardAPR
	}
	if !common.IsNil(o.EstExtraRewardAmt) {
		toSerialize["estExtraRewardAmt"] = o.EstExtraRewardAmt
	}
	if !common.IsNil(o.BoostRewardAsset) {
		toSerialize["boostRewardAsset"] = o.BoostRewardAsset
	}
	if !common.IsNil(o.BoostApr) {
		toSerialize["boostApr"] = o.BoostApr
	}
	if !common.IsNil(o.TotalBoostRewardAmt) {
		toSerialize["totalBoostRewardAmt"] = o.TotalBoostRewardAmt
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
	if !common.IsNil(o.RedeemAmountEarly) {
		toSerialize["redeemAmountEarly"] = o.RedeemAmountEarly
	}
	if !common.IsNil(o.RewardsEndDate) {
		toSerialize["rewardsEndDate"] = o.RewardsEndDate
	}
	if !common.IsNil(o.DeliverDate) {
		toSerialize["deliverDate"] = o.DeliverDate
	}
	if !common.IsNil(o.RedeemPeriod) {
		toSerialize["redeemPeriod"] = o.RedeemPeriod
	}
	if !common.IsNil(o.RedeemingAmt) {
		toSerialize["redeemingAmt"] = o.RedeemingAmt
	}
	if !common.IsNil(o.RedeemTo) {
		toSerialize["redeemTo"] = o.RedeemTo
	}
	if !common.IsNil(o.PartialAmtDeliverDate) {
		toSerialize["partialAmtDeliverDate"] = o.PartialAmtDeliverDate
	}
	if !common.IsNil(o.CanRedeemEarly) {
		toSerialize["canRedeemEarly"] = o.CanRedeemEarly
	}
	if !common.IsNil(o.CanFastRedemption) {
		toSerialize["canFastRedemption"] = o.CanFastRedemption
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
	if !common.IsNil(o.CanReStake) {
		toSerialize["canReStake"] = o.CanReStake
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetLockedProductPositionResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetLockedProductPositionResponseRowsInner := _GetLockedProductPositionResponseRowsInner{}

	err = json.Unmarshal(data, &varGetLockedProductPositionResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetLockedProductPositionResponseRowsInner(varGetLockedProductPositionResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "positionId")
		delete(additionalProperties, "parentPositionId")
		delete(additionalProperties, "projectId")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "purchaseTime")
		delete(additionalProperties, "duration")
		delete(additionalProperties, "accrualDays")
		delete(additionalProperties, "rewardAsset")
		delete(additionalProperties, "APY")
		delete(additionalProperties, "rewardAmt")
		delete(additionalProperties, "extraRewardAsset")
		delete(additionalProperties, "extraRewardAPR")
		delete(additionalProperties, "estExtraRewardAmt")
		delete(additionalProperties, "boostRewardAsset")
		delete(additionalProperties, "boostApr")
		delete(additionalProperties, "totalBoostRewardAmt")
		delete(additionalProperties, "nextPay")
		delete(additionalProperties, "nextPayDate")
		delete(additionalProperties, "payPeriod")
		delete(additionalProperties, "redeemAmountEarly")
		delete(additionalProperties, "rewardsEndDate")
		delete(additionalProperties, "deliverDate")
		delete(additionalProperties, "redeemPeriod")
		delete(additionalProperties, "redeemingAmt")
		delete(additionalProperties, "redeemTo")
		delete(additionalProperties, "partialAmtDeliverDate")
		delete(additionalProperties, "canRedeemEarly")
		delete(additionalProperties, "canFastRedemption")
		delete(additionalProperties, "autoSubscribe")
		delete(additionalProperties, "type")
		delete(additionalProperties, "status")
		delete(additionalProperties, "canReStake")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetLockedProductPositionResponseRowsInner struct {
	value *GetLockedProductPositionResponseRowsInner
	isSet bool
}

func (v NullableGetLockedProductPositionResponseRowsInner) Get() *GetLockedProductPositionResponseRowsInner {
	return v.value
}

func (v *NullableGetLockedProductPositionResponseRowsInner) Set(val *GetLockedProductPositionResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLockedProductPositionResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLockedProductPositionResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLockedProductPositionResponseRowsInner(val *GetLockedProductPositionResponseRowsInner) *NullableGetLockedProductPositionResponseRowsInner {
	return &NullableGetLockedProductPositionResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetLockedProductPositionResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLockedProductPositionResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
