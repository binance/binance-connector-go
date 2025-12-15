# GetLockedRedemptionRecordResponseRowsInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**PositionId** | Pointer to **int64** |  | [optional] 
**RedeemId** | Pointer to **int64** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 
**Asset** | Pointer to **string** |  | [optional] 
**LockPeriod** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **string** |  | [optional] 
**OriginalAmount** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**DeliverDate** | Pointer to **string** |  | [optional] 
**LossAmount** | Pointer to **string** |  | [optional] 
**IsComplete** | Pointer to **bool** |  | [optional] 
**RewardAsset** | Pointer to **string** |  | [optional] 
**RewardAmt** | Pointer to **string** |  | [optional] 
**ExtraRewardAsset** | Pointer to **string** |  | [optional] 
**EstExtraRewardAmt** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewGetLockedRedemptionRecordResponseRowsInner

`func NewGetLockedRedemptionRecordResponseRowsInner() *GetLockedRedemptionRecordResponseRowsInner`

NewGetLockedRedemptionRecordResponseRowsInner instantiates a new GetLockedRedemptionRecordResponseRowsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLockedRedemptionRecordResponseRowsInnerWithDefaults

`func NewGetLockedRedemptionRecordResponseRowsInnerWithDefaults() *GetLockedRedemptionRecordResponseRowsInner`

NewGetLockedRedemptionRecordResponseRowsInnerWithDefaults instantiates a new GetLockedRedemptionRecordResponseRowsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPositionId

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetPositionId() int64`

GetPositionId returns the PositionId field if non-nil, zero value otherwise.

### GetPositionIdOk

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetPositionIdOk() (*int64, bool)`

GetPositionIdOk returns a tuple with the PositionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionId

`func (o *GetLockedRedemptionRecordResponseRowsInner) SetPositionId(v int64)`

SetPositionId sets PositionId field to given value.

### HasPositionId

`func (o *GetLockedRedemptionRecordResponseRowsInner) HasPositionId() bool`

HasPositionId returns a boolean if a field has been set.

### GetRedeemId

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetRedeemId() int64`

GetRedeemId returns the RedeemId field if non-nil, zero value otherwise.

### GetRedeemIdOk

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetRedeemIdOk() (*int64, bool)`

GetRedeemIdOk returns a tuple with the RedeemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemId

`func (o *GetLockedRedemptionRecordResponseRowsInner) SetRedeemId(v int64)`

SetRedeemId sets RedeemId field to given value.

### HasRedeemId

`func (o *GetLockedRedemptionRecordResponseRowsInner) HasRedeemId() bool`

HasRedeemId returns a boolean if a field has been set.

### GetTime

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *GetLockedRedemptionRecordResponseRowsInner) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *GetLockedRedemptionRecordResponseRowsInner) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetAsset

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *GetLockedRedemptionRecordResponseRowsInner) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *GetLockedRedemptionRecordResponseRowsInner) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetLockPeriod

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetLockPeriod() string`

GetLockPeriod returns the LockPeriod field if non-nil, zero value otherwise.

### GetLockPeriodOk

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetLockPeriodOk() (*string, bool)`

GetLockPeriodOk returns a tuple with the LockPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockPeriod

`func (o *GetLockedRedemptionRecordResponseRowsInner) SetLockPeriod(v string)`

SetLockPeriod sets LockPeriod field to given value.

### HasLockPeriod

`func (o *GetLockedRedemptionRecordResponseRowsInner) HasLockPeriod() bool`

HasLockPeriod returns a boolean if a field has been set.

### GetAmount

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetLockedRedemptionRecordResponseRowsInner) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GetLockedRedemptionRecordResponseRowsInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetOriginalAmount

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetOriginalAmount() string`

GetOriginalAmount returns the OriginalAmount field if non-nil, zero value otherwise.

### GetOriginalAmountOk

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetOriginalAmountOk() (*string, bool)`

GetOriginalAmountOk returns a tuple with the OriginalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalAmount

`func (o *GetLockedRedemptionRecordResponseRowsInner) SetOriginalAmount(v string)`

SetOriginalAmount sets OriginalAmount field to given value.

### HasOriginalAmount

`func (o *GetLockedRedemptionRecordResponseRowsInner) HasOriginalAmount() bool`

HasOriginalAmount returns a boolean if a field has been set.

### GetType

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetLockedRedemptionRecordResponseRowsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetLockedRedemptionRecordResponseRowsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDeliverDate

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetDeliverDate() string`

GetDeliverDate returns the DeliverDate field if non-nil, zero value otherwise.

### GetDeliverDateOk

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetDeliverDateOk() (*string, bool)`

GetDeliverDateOk returns a tuple with the DeliverDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverDate

`func (o *GetLockedRedemptionRecordResponseRowsInner) SetDeliverDate(v string)`

SetDeliverDate sets DeliverDate field to given value.

### HasDeliverDate

`func (o *GetLockedRedemptionRecordResponseRowsInner) HasDeliverDate() bool`

HasDeliverDate returns a boolean if a field has been set.

### GetLossAmount

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetLossAmount() string`

GetLossAmount returns the LossAmount field if non-nil, zero value otherwise.

### GetLossAmountOk

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetLossAmountOk() (*string, bool)`

GetLossAmountOk returns a tuple with the LossAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLossAmount

`func (o *GetLockedRedemptionRecordResponseRowsInner) SetLossAmount(v string)`

SetLossAmount sets LossAmount field to given value.

### HasLossAmount

`func (o *GetLockedRedemptionRecordResponseRowsInner) HasLossAmount() bool`

HasLossAmount returns a boolean if a field has been set.

### GetIsComplete

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetIsComplete() bool`

GetIsComplete returns the IsComplete field if non-nil, zero value otherwise.

### GetIsCompleteOk

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetIsCompleteOk() (*bool, bool)`

GetIsCompleteOk returns a tuple with the IsComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsComplete

`func (o *GetLockedRedemptionRecordResponseRowsInner) SetIsComplete(v bool)`

SetIsComplete sets IsComplete field to given value.

### HasIsComplete

`func (o *GetLockedRedemptionRecordResponseRowsInner) HasIsComplete() bool`

HasIsComplete returns a boolean if a field has been set.

### GetRewardAsset

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetRewardAsset() string`

GetRewardAsset returns the RewardAsset field if non-nil, zero value otherwise.

### GetRewardAssetOk

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetRewardAssetOk() (*string, bool)`

GetRewardAssetOk returns a tuple with the RewardAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardAsset

`func (o *GetLockedRedemptionRecordResponseRowsInner) SetRewardAsset(v string)`

SetRewardAsset sets RewardAsset field to given value.

### HasRewardAsset

`func (o *GetLockedRedemptionRecordResponseRowsInner) HasRewardAsset() bool`

HasRewardAsset returns a boolean if a field has been set.

### GetRewardAmt

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetRewardAmt() string`

GetRewardAmt returns the RewardAmt field if non-nil, zero value otherwise.

### GetRewardAmtOk

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetRewardAmtOk() (*string, bool)`

GetRewardAmtOk returns a tuple with the RewardAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardAmt

`func (o *GetLockedRedemptionRecordResponseRowsInner) SetRewardAmt(v string)`

SetRewardAmt sets RewardAmt field to given value.

### HasRewardAmt

`func (o *GetLockedRedemptionRecordResponseRowsInner) HasRewardAmt() bool`

HasRewardAmt returns a boolean if a field has been set.

### GetExtraRewardAsset

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetExtraRewardAsset() string`

GetExtraRewardAsset returns the ExtraRewardAsset field if non-nil, zero value otherwise.

### GetExtraRewardAssetOk

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetExtraRewardAssetOk() (*string, bool)`

GetExtraRewardAssetOk returns a tuple with the ExtraRewardAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraRewardAsset

`func (o *GetLockedRedemptionRecordResponseRowsInner) SetExtraRewardAsset(v string)`

SetExtraRewardAsset sets ExtraRewardAsset field to given value.

### HasExtraRewardAsset

`func (o *GetLockedRedemptionRecordResponseRowsInner) HasExtraRewardAsset() bool`

HasExtraRewardAsset returns a boolean if a field has been set.

### GetEstExtraRewardAmt

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetEstExtraRewardAmt() string`

GetEstExtraRewardAmt returns the EstExtraRewardAmt field if non-nil, zero value otherwise.

### GetEstExtraRewardAmtOk

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetEstExtraRewardAmtOk() (*string, bool)`

GetEstExtraRewardAmtOk returns a tuple with the EstExtraRewardAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstExtraRewardAmt

`func (o *GetLockedRedemptionRecordResponseRowsInner) SetEstExtraRewardAmt(v string)`

SetEstExtraRewardAmt sets EstExtraRewardAmt field to given value.

### HasEstExtraRewardAmt

`func (o *GetLockedRedemptionRecordResponseRowsInner) HasEstExtraRewardAmt() bool`

HasEstExtraRewardAmt returns a boolean if a field has been set.

### GetStatus

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetLockedRedemptionRecordResponseRowsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetLockedRedemptionRecordResponseRowsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetLockedRedemptionRecordResponseRowsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to README]](../README.md)


