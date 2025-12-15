# GetLockedProductPositionResponseRowsInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**PositionId** | Pointer to **int64** |  | [optional] 
**ParentPositionId** | Pointer to **int64** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**Asset** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **string** |  | [optional] 
**PurchaseTime** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **string** |  | [optional] 
**AccrualDays** | Pointer to **string** |  | [optional] 
**RewardAsset** | Pointer to **string** |  | [optional] 
**APY** | Pointer to **string** |  | [optional] 
**RewardAmt** | Pointer to **string** |  | [optional] 
**ExtraRewardAsset** | Pointer to **string** |  | [optional] 
**ExtraRewardAPR** | Pointer to **string** |  | [optional] 
**EstExtraRewardAmt** | Pointer to **string** |  | [optional] 
**BoostRewardAsset** | Pointer to **string** |  | [optional] 
**BoostApr** | Pointer to **string** |  | [optional] 
**TotalBoostRewardAmt** | Pointer to **string** |  | [optional] 
**NextPay** | Pointer to **string** |  | [optional] 
**NextPayDate** | Pointer to **string** |  | [optional] 
**PayPeriod** | Pointer to **string** |  | [optional] 
**RedeemAmountEarly** | Pointer to **string** |  | [optional] 
**RewardsEndDate** | Pointer to **string** |  | [optional] 
**DeliverDate** | Pointer to **string** |  | [optional] 
**RedeemPeriod** | Pointer to **string** |  | [optional] 
**RedeemingAmt** | Pointer to **string** |  | [optional] 
**RedeemTo** | Pointer to **string** |  | [optional] 
**PartialAmtDeliverDate** | Pointer to **string** |  | [optional] 
**CanRedeemEarly** | Pointer to **bool** |  | [optional] 
**CanFastRedemption** | Pointer to **bool** |  | [optional] 
**AutoSubscribe** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CanReStake** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetLockedProductPositionResponseRowsInner

`func NewGetLockedProductPositionResponseRowsInner() *GetLockedProductPositionResponseRowsInner`

NewGetLockedProductPositionResponseRowsInner instantiates a new GetLockedProductPositionResponseRowsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLockedProductPositionResponseRowsInnerWithDefaults

`func NewGetLockedProductPositionResponseRowsInnerWithDefaults() *GetLockedProductPositionResponseRowsInner`

NewGetLockedProductPositionResponseRowsInnerWithDefaults instantiates a new GetLockedProductPositionResponseRowsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPositionId

`func (o *GetLockedProductPositionResponseRowsInner) GetPositionId() int64`

GetPositionId returns the PositionId field if non-nil, zero value otherwise.

### GetPositionIdOk

`func (o *GetLockedProductPositionResponseRowsInner) GetPositionIdOk() (*int64, bool)`

GetPositionIdOk returns a tuple with the PositionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionId

`func (o *GetLockedProductPositionResponseRowsInner) SetPositionId(v int64)`

SetPositionId sets PositionId field to given value.

### HasPositionId

`func (o *GetLockedProductPositionResponseRowsInner) HasPositionId() bool`

HasPositionId returns a boolean if a field has been set.

### GetParentPositionId

`func (o *GetLockedProductPositionResponseRowsInner) GetParentPositionId() int64`

GetParentPositionId returns the ParentPositionId field if non-nil, zero value otherwise.

### GetParentPositionIdOk

`func (o *GetLockedProductPositionResponseRowsInner) GetParentPositionIdOk() (*int64, bool)`

GetParentPositionIdOk returns a tuple with the ParentPositionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPositionId

`func (o *GetLockedProductPositionResponseRowsInner) SetParentPositionId(v int64)`

SetParentPositionId sets ParentPositionId field to given value.

### HasParentPositionId

`func (o *GetLockedProductPositionResponseRowsInner) HasParentPositionId() bool`

HasParentPositionId returns a boolean if a field has been set.

### GetProjectId

`func (o *GetLockedProductPositionResponseRowsInner) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GetLockedProductPositionResponseRowsInner) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GetLockedProductPositionResponseRowsInner) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *GetLockedProductPositionResponseRowsInner) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetAsset

`func (o *GetLockedProductPositionResponseRowsInner) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *GetLockedProductPositionResponseRowsInner) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *GetLockedProductPositionResponseRowsInner) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *GetLockedProductPositionResponseRowsInner) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetAmount

`func (o *GetLockedProductPositionResponseRowsInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetLockedProductPositionResponseRowsInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetLockedProductPositionResponseRowsInner) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GetLockedProductPositionResponseRowsInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetPurchaseTime

`func (o *GetLockedProductPositionResponseRowsInner) GetPurchaseTime() string`

GetPurchaseTime returns the PurchaseTime field if non-nil, zero value otherwise.

### GetPurchaseTimeOk

`func (o *GetLockedProductPositionResponseRowsInner) GetPurchaseTimeOk() (*string, bool)`

GetPurchaseTimeOk returns a tuple with the PurchaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseTime

`func (o *GetLockedProductPositionResponseRowsInner) SetPurchaseTime(v string)`

SetPurchaseTime sets PurchaseTime field to given value.

### HasPurchaseTime

`func (o *GetLockedProductPositionResponseRowsInner) HasPurchaseTime() bool`

HasPurchaseTime returns a boolean if a field has been set.

### GetDuration

`func (o *GetLockedProductPositionResponseRowsInner) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GetLockedProductPositionResponseRowsInner) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GetLockedProductPositionResponseRowsInner) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *GetLockedProductPositionResponseRowsInner) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetAccrualDays

`func (o *GetLockedProductPositionResponseRowsInner) GetAccrualDays() string`

GetAccrualDays returns the AccrualDays field if non-nil, zero value otherwise.

### GetAccrualDaysOk

`func (o *GetLockedProductPositionResponseRowsInner) GetAccrualDaysOk() (*string, bool)`

GetAccrualDaysOk returns a tuple with the AccrualDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrualDays

`func (o *GetLockedProductPositionResponseRowsInner) SetAccrualDays(v string)`

SetAccrualDays sets AccrualDays field to given value.

### HasAccrualDays

`func (o *GetLockedProductPositionResponseRowsInner) HasAccrualDays() bool`

HasAccrualDays returns a boolean if a field has been set.

### GetRewardAsset

`func (o *GetLockedProductPositionResponseRowsInner) GetRewardAsset() string`

GetRewardAsset returns the RewardAsset field if non-nil, zero value otherwise.

### GetRewardAssetOk

`func (o *GetLockedProductPositionResponseRowsInner) GetRewardAssetOk() (*string, bool)`

GetRewardAssetOk returns a tuple with the RewardAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardAsset

`func (o *GetLockedProductPositionResponseRowsInner) SetRewardAsset(v string)`

SetRewardAsset sets RewardAsset field to given value.

### HasRewardAsset

`func (o *GetLockedProductPositionResponseRowsInner) HasRewardAsset() bool`

HasRewardAsset returns a boolean if a field has been set.

### GetAPY

`func (o *GetLockedProductPositionResponseRowsInner) GetAPY() string`

GetAPY returns the APY field if non-nil, zero value otherwise.

### GetAPYOk

`func (o *GetLockedProductPositionResponseRowsInner) GetAPYOk() (*string, bool)`

GetAPYOk returns a tuple with the APY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPY

`func (o *GetLockedProductPositionResponseRowsInner) SetAPY(v string)`

SetAPY sets APY field to given value.

### HasAPY

`func (o *GetLockedProductPositionResponseRowsInner) HasAPY() bool`

HasAPY returns a boolean if a field has been set.

### GetRewardAmt

`func (o *GetLockedProductPositionResponseRowsInner) GetRewardAmt() string`

GetRewardAmt returns the RewardAmt field if non-nil, zero value otherwise.

### GetRewardAmtOk

`func (o *GetLockedProductPositionResponseRowsInner) GetRewardAmtOk() (*string, bool)`

GetRewardAmtOk returns a tuple with the RewardAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardAmt

`func (o *GetLockedProductPositionResponseRowsInner) SetRewardAmt(v string)`

SetRewardAmt sets RewardAmt field to given value.

### HasRewardAmt

`func (o *GetLockedProductPositionResponseRowsInner) HasRewardAmt() bool`

HasRewardAmt returns a boolean if a field has been set.

### GetExtraRewardAsset

`func (o *GetLockedProductPositionResponseRowsInner) GetExtraRewardAsset() string`

GetExtraRewardAsset returns the ExtraRewardAsset field if non-nil, zero value otherwise.

### GetExtraRewardAssetOk

`func (o *GetLockedProductPositionResponseRowsInner) GetExtraRewardAssetOk() (*string, bool)`

GetExtraRewardAssetOk returns a tuple with the ExtraRewardAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraRewardAsset

`func (o *GetLockedProductPositionResponseRowsInner) SetExtraRewardAsset(v string)`

SetExtraRewardAsset sets ExtraRewardAsset field to given value.

### HasExtraRewardAsset

`func (o *GetLockedProductPositionResponseRowsInner) HasExtraRewardAsset() bool`

HasExtraRewardAsset returns a boolean if a field has been set.

### GetExtraRewardAPR

`func (o *GetLockedProductPositionResponseRowsInner) GetExtraRewardAPR() string`

GetExtraRewardAPR returns the ExtraRewardAPR field if non-nil, zero value otherwise.

### GetExtraRewardAPROk

`func (o *GetLockedProductPositionResponseRowsInner) GetExtraRewardAPROk() (*string, bool)`

GetExtraRewardAPROk returns a tuple with the ExtraRewardAPR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraRewardAPR

`func (o *GetLockedProductPositionResponseRowsInner) SetExtraRewardAPR(v string)`

SetExtraRewardAPR sets ExtraRewardAPR field to given value.

### HasExtraRewardAPR

`func (o *GetLockedProductPositionResponseRowsInner) HasExtraRewardAPR() bool`

HasExtraRewardAPR returns a boolean if a field has been set.

### GetEstExtraRewardAmt

`func (o *GetLockedProductPositionResponseRowsInner) GetEstExtraRewardAmt() string`

GetEstExtraRewardAmt returns the EstExtraRewardAmt field if non-nil, zero value otherwise.

### GetEstExtraRewardAmtOk

`func (o *GetLockedProductPositionResponseRowsInner) GetEstExtraRewardAmtOk() (*string, bool)`

GetEstExtraRewardAmtOk returns a tuple with the EstExtraRewardAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstExtraRewardAmt

`func (o *GetLockedProductPositionResponseRowsInner) SetEstExtraRewardAmt(v string)`

SetEstExtraRewardAmt sets EstExtraRewardAmt field to given value.

### HasEstExtraRewardAmt

`func (o *GetLockedProductPositionResponseRowsInner) HasEstExtraRewardAmt() bool`

HasEstExtraRewardAmt returns a boolean if a field has been set.

### GetBoostRewardAsset

`func (o *GetLockedProductPositionResponseRowsInner) GetBoostRewardAsset() string`

GetBoostRewardAsset returns the BoostRewardAsset field if non-nil, zero value otherwise.

### GetBoostRewardAssetOk

`func (o *GetLockedProductPositionResponseRowsInner) GetBoostRewardAssetOk() (*string, bool)`

GetBoostRewardAssetOk returns a tuple with the BoostRewardAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoostRewardAsset

`func (o *GetLockedProductPositionResponseRowsInner) SetBoostRewardAsset(v string)`

SetBoostRewardAsset sets BoostRewardAsset field to given value.

### HasBoostRewardAsset

`func (o *GetLockedProductPositionResponseRowsInner) HasBoostRewardAsset() bool`

HasBoostRewardAsset returns a boolean if a field has been set.

### GetBoostApr

`func (o *GetLockedProductPositionResponseRowsInner) GetBoostApr() string`

GetBoostApr returns the BoostApr field if non-nil, zero value otherwise.

### GetBoostAprOk

`func (o *GetLockedProductPositionResponseRowsInner) GetBoostAprOk() (*string, bool)`

GetBoostAprOk returns a tuple with the BoostApr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoostApr

`func (o *GetLockedProductPositionResponseRowsInner) SetBoostApr(v string)`

SetBoostApr sets BoostApr field to given value.

### HasBoostApr

`func (o *GetLockedProductPositionResponseRowsInner) HasBoostApr() bool`

HasBoostApr returns a boolean if a field has been set.

### GetTotalBoostRewardAmt

`func (o *GetLockedProductPositionResponseRowsInner) GetTotalBoostRewardAmt() string`

GetTotalBoostRewardAmt returns the TotalBoostRewardAmt field if non-nil, zero value otherwise.

### GetTotalBoostRewardAmtOk

`func (o *GetLockedProductPositionResponseRowsInner) GetTotalBoostRewardAmtOk() (*string, bool)`

GetTotalBoostRewardAmtOk returns a tuple with the TotalBoostRewardAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBoostRewardAmt

`func (o *GetLockedProductPositionResponseRowsInner) SetTotalBoostRewardAmt(v string)`

SetTotalBoostRewardAmt sets TotalBoostRewardAmt field to given value.

### HasTotalBoostRewardAmt

`func (o *GetLockedProductPositionResponseRowsInner) HasTotalBoostRewardAmt() bool`

HasTotalBoostRewardAmt returns a boolean if a field has been set.

### GetNextPay

`func (o *GetLockedProductPositionResponseRowsInner) GetNextPay() string`

GetNextPay returns the NextPay field if non-nil, zero value otherwise.

### GetNextPayOk

`func (o *GetLockedProductPositionResponseRowsInner) GetNextPayOk() (*string, bool)`

GetNextPayOk returns a tuple with the NextPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPay

`func (o *GetLockedProductPositionResponseRowsInner) SetNextPay(v string)`

SetNextPay sets NextPay field to given value.

### HasNextPay

`func (o *GetLockedProductPositionResponseRowsInner) HasNextPay() bool`

HasNextPay returns a boolean if a field has been set.

### GetNextPayDate

`func (o *GetLockedProductPositionResponseRowsInner) GetNextPayDate() string`

GetNextPayDate returns the NextPayDate field if non-nil, zero value otherwise.

### GetNextPayDateOk

`func (o *GetLockedProductPositionResponseRowsInner) GetNextPayDateOk() (*string, bool)`

GetNextPayDateOk returns a tuple with the NextPayDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPayDate

`func (o *GetLockedProductPositionResponseRowsInner) SetNextPayDate(v string)`

SetNextPayDate sets NextPayDate field to given value.

### HasNextPayDate

`func (o *GetLockedProductPositionResponseRowsInner) HasNextPayDate() bool`

HasNextPayDate returns a boolean if a field has been set.

### GetPayPeriod

`func (o *GetLockedProductPositionResponseRowsInner) GetPayPeriod() string`

GetPayPeriod returns the PayPeriod field if non-nil, zero value otherwise.

### GetPayPeriodOk

`func (o *GetLockedProductPositionResponseRowsInner) GetPayPeriodOk() (*string, bool)`

GetPayPeriodOk returns a tuple with the PayPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayPeriod

`func (o *GetLockedProductPositionResponseRowsInner) SetPayPeriod(v string)`

SetPayPeriod sets PayPeriod field to given value.

### HasPayPeriod

`func (o *GetLockedProductPositionResponseRowsInner) HasPayPeriod() bool`

HasPayPeriod returns a boolean if a field has been set.

### GetRedeemAmountEarly

`func (o *GetLockedProductPositionResponseRowsInner) GetRedeemAmountEarly() string`

GetRedeemAmountEarly returns the RedeemAmountEarly field if non-nil, zero value otherwise.

### GetRedeemAmountEarlyOk

`func (o *GetLockedProductPositionResponseRowsInner) GetRedeemAmountEarlyOk() (*string, bool)`

GetRedeemAmountEarlyOk returns a tuple with the RedeemAmountEarly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemAmountEarly

`func (o *GetLockedProductPositionResponseRowsInner) SetRedeemAmountEarly(v string)`

SetRedeemAmountEarly sets RedeemAmountEarly field to given value.

### HasRedeemAmountEarly

`func (o *GetLockedProductPositionResponseRowsInner) HasRedeemAmountEarly() bool`

HasRedeemAmountEarly returns a boolean if a field has been set.

### GetRewardsEndDate

`func (o *GetLockedProductPositionResponseRowsInner) GetRewardsEndDate() string`

GetRewardsEndDate returns the RewardsEndDate field if non-nil, zero value otherwise.

### GetRewardsEndDateOk

`func (o *GetLockedProductPositionResponseRowsInner) GetRewardsEndDateOk() (*string, bool)`

GetRewardsEndDateOk returns a tuple with the RewardsEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardsEndDate

`func (o *GetLockedProductPositionResponseRowsInner) SetRewardsEndDate(v string)`

SetRewardsEndDate sets RewardsEndDate field to given value.

### HasRewardsEndDate

`func (o *GetLockedProductPositionResponseRowsInner) HasRewardsEndDate() bool`

HasRewardsEndDate returns a boolean if a field has been set.

### GetDeliverDate

`func (o *GetLockedProductPositionResponseRowsInner) GetDeliverDate() string`

GetDeliverDate returns the DeliverDate field if non-nil, zero value otherwise.

### GetDeliverDateOk

`func (o *GetLockedProductPositionResponseRowsInner) GetDeliverDateOk() (*string, bool)`

GetDeliverDateOk returns a tuple with the DeliverDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverDate

`func (o *GetLockedProductPositionResponseRowsInner) SetDeliverDate(v string)`

SetDeliverDate sets DeliverDate field to given value.

### HasDeliverDate

`func (o *GetLockedProductPositionResponseRowsInner) HasDeliverDate() bool`

HasDeliverDate returns a boolean if a field has been set.

### GetRedeemPeriod

`func (o *GetLockedProductPositionResponseRowsInner) GetRedeemPeriod() string`

GetRedeemPeriod returns the RedeemPeriod field if non-nil, zero value otherwise.

### GetRedeemPeriodOk

`func (o *GetLockedProductPositionResponseRowsInner) GetRedeemPeriodOk() (*string, bool)`

GetRedeemPeriodOk returns a tuple with the RedeemPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemPeriod

`func (o *GetLockedProductPositionResponseRowsInner) SetRedeemPeriod(v string)`

SetRedeemPeriod sets RedeemPeriod field to given value.

### HasRedeemPeriod

`func (o *GetLockedProductPositionResponseRowsInner) HasRedeemPeriod() bool`

HasRedeemPeriod returns a boolean if a field has been set.

### GetRedeemingAmt

`func (o *GetLockedProductPositionResponseRowsInner) GetRedeemingAmt() string`

GetRedeemingAmt returns the RedeemingAmt field if non-nil, zero value otherwise.

### GetRedeemingAmtOk

`func (o *GetLockedProductPositionResponseRowsInner) GetRedeemingAmtOk() (*string, bool)`

GetRedeemingAmtOk returns a tuple with the RedeemingAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemingAmt

`func (o *GetLockedProductPositionResponseRowsInner) SetRedeemingAmt(v string)`

SetRedeemingAmt sets RedeemingAmt field to given value.

### HasRedeemingAmt

`func (o *GetLockedProductPositionResponseRowsInner) HasRedeemingAmt() bool`

HasRedeemingAmt returns a boolean if a field has been set.

### GetRedeemTo

`func (o *GetLockedProductPositionResponseRowsInner) GetRedeemTo() string`

GetRedeemTo returns the RedeemTo field if non-nil, zero value otherwise.

### GetRedeemToOk

`func (o *GetLockedProductPositionResponseRowsInner) GetRedeemToOk() (*string, bool)`

GetRedeemToOk returns a tuple with the RedeemTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemTo

`func (o *GetLockedProductPositionResponseRowsInner) SetRedeemTo(v string)`

SetRedeemTo sets RedeemTo field to given value.

### HasRedeemTo

`func (o *GetLockedProductPositionResponseRowsInner) HasRedeemTo() bool`

HasRedeemTo returns a boolean if a field has been set.

### GetPartialAmtDeliverDate

`func (o *GetLockedProductPositionResponseRowsInner) GetPartialAmtDeliverDate() string`

GetPartialAmtDeliverDate returns the PartialAmtDeliverDate field if non-nil, zero value otherwise.

### GetPartialAmtDeliverDateOk

`func (o *GetLockedProductPositionResponseRowsInner) GetPartialAmtDeliverDateOk() (*string, bool)`

GetPartialAmtDeliverDateOk returns a tuple with the PartialAmtDeliverDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialAmtDeliverDate

`func (o *GetLockedProductPositionResponseRowsInner) SetPartialAmtDeliverDate(v string)`

SetPartialAmtDeliverDate sets PartialAmtDeliverDate field to given value.

### HasPartialAmtDeliverDate

`func (o *GetLockedProductPositionResponseRowsInner) HasPartialAmtDeliverDate() bool`

HasPartialAmtDeliverDate returns a boolean if a field has been set.

### GetCanRedeemEarly

`func (o *GetLockedProductPositionResponseRowsInner) GetCanRedeemEarly() bool`

GetCanRedeemEarly returns the CanRedeemEarly field if non-nil, zero value otherwise.

### GetCanRedeemEarlyOk

`func (o *GetLockedProductPositionResponseRowsInner) GetCanRedeemEarlyOk() (*bool, bool)`

GetCanRedeemEarlyOk returns a tuple with the CanRedeemEarly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRedeemEarly

`func (o *GetLockedProductPositionResponseRowsInner) SetCanRedeemEarly(v bool)`

SetCanRedeemEarly sets CanRedeemEarly field to given value.

### HasCanRedeemEarly

`func (o *GetLockedProductPositionResponseRowsInner) HasCanRedeemEarly() bool`

HasCanRedeemEarly returns a boolean if a field has been set.

### GetCanFastRedemption

`func (o *GetLockedProductPositionResponseRowsInner) GetCanFastRedemption() bool`

GetCanFastRedemption returns the CanFastRedemption field if non-nil, zero value otherwise.

### GetCanFastRedemptionOk

`func (o *GetLockedProductPositionResponseRowsInner) GetCanFastRedemptionOk() (*bool, bool)`

GetCanFastRedemptionOk returns a tuple with the CanFastRedemption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanFastRedemption

`func (o *GetLockedProductPositionResponseRowsInner) SetCanFastRedemption(v bool)`

SetCanFastRedemption sets CanFastRedemption field to given value.

### HasCanFastRedemption

`func (o *GetLockedProductPositionResponseRowsInner) HasCanFastRedemption() bool`

HasCanFastRedemption returns a boolean if a field has been set.

### GetAutoSubscribe

`func (o *GetLockedProductPositionResponseRowsInner) GetAutoSubscribe() bool`

GetAutoSubscribe returns the AutoSubscribe field if non-nil, zero value otherwise.

### GetAutoSubscribeOk

`func (o *GetLockedProductPositionResponseRowsInner) GetAutoSubscribeOk() (*bool, bool)`

GetAutoSubscribeOk returns a tuple with the AutoSubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSubscribe

`func (o *GetLockedProductPositionResponseRowsInner) SetAutoSubscribe(v bool)`

SetAutoSubscribe sets AutoSubscribe field to given value.

### HasAutoSubscribe

`func (o *GetLockedProductPositionResponseRowsInner) HasAutoSubscribe() bool`

HasAutoSubscribe returns a boolean if a field has been set.

### GetType

`func (o *GetLockedProductPositionResponseRowsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetLockedProductPositionResponseRowsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetLockedProductPositionResponseRowsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetLockedProductPositionResponseRowsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *GetLockedProductPositionResponseRowsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetLockedProductPositionResponseRowsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetLockedProductPositionResponseRowsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetLockedProductPositionResponseRowsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCanReStake

`func (o *GetLockedProductPositionResponseRowsInner) GetCanReStake() bool`

GetCanReStake returns the CanReStake field if non-nil, zero value otherwise.

### GetCanReStakeOk

`func (o *GetLockedProductPositionResponseRowsInner) GetCanReStakeOk() (*bool, bool)`

GetCanReStakeOk returns a tuple with the CanReStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanReStake

`func (o *GetLockedProductPositionResponseRowsInner) SetCanReStake(v bool)`

SetCanReStake sets CanReStake field to given value.

### HasCanReStake

`func (o *GetLockedProductPositionResponseRowsInner) HasCanReStake() bool`

HasCanReStake returns a boolean if a field has been set.


[[Back to README]](../README.md)


