# GetLockedSubscriptionRecordResponseRowsInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**PositionId** | Pointer to **int64** |  | [optional] 
**PurchaseId** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 
**Asset** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **string** |  | [optional] 
**LockPeriod** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**SourceAccount** | Pointer to **string** |  | [optional] 
**AmtFromSpot** | Pointer to **string** |  | [optional] 
**AmtFromFunding** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewGetLockedSubscriptionRecordResponseRowsInner

`func NewGetLockedSubscriptionRecordResponseRowsInner() *GetLockedSubscriptionRecordResponseRowsInner`

NewGetLockedSubscriptionRecordResponseRowsInner instantiates a new GetLockedSubscriptionRecordResponseRowsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLockedSubscriptionRecordResponseRowsInnerWithDefaults

`func NewGetLockedSubscriptionRecordResponseRowsInnerWithDefaults() *GetLockedSubscriptionRecordResponseRowsInner`

NewGetLockedSubscriptionRecordResponseRowsInnerWithDefaults instantiates a new GetLockedSubscriptionRecordResponseRowsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPositionId

`func (o *GetLockedSubscriptionRecordResponseRowsInner) GetPositionId() int64`

GetPositionId returns the PositionId field if non-nil, zero value otherwise.

### GetPositionIdOk

`func (o *GetLockedSubscriptionRecordResponseRowsInner) GetPositionIdOk() (*int64, bool)`

GetPositionIdOk returns a tuple with the PositionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionId

`func (o *GetLockedSubscriptionRecordResponseRowsInner) SetPositionId(v int64)`

SetPositionId sets PositionId field to given value.

### HasPositionId

`func (o *GetLockedSubscriptionRecordResponseRowsInner) HasPositionId() bool`

HasPositionId returns a boolean if a field has been set.

### GetPurchaseId

`func (o *GetLockedSubscriptionRecordResponseRowsInner) GetPurchaseId() string`

GetPurchaseId returns the PurchaseId field if non-nil, zero value otherwise.

### GetPurchaseIdOk

`func (o *GetLockedSubscriptionRecordResponseRowsInner) GetPurchaseIdOk() (*string, bool)`

GetPurchaseIdOk returns a tuple with the PurchaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseId

`func (o *GetLockedSubscriptionRecordResponseRowsInner) SetPurchaseId(v string)`

SetPurchaseId sets PurchaseId field to given value.

### HasPurchaseId

`func (o *GetLockedSubscriptionRecordResponseRowsInner) HasPurchaseId() bool`

HasPurchaseId returns a boolean if a field has been set.

### GetProjectId

`func (o *GetLockedSubscriptionRecordResponseRowsInner) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GetLockedSubscriptionRecordResponseRowsInner) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GetLockedSubscriptionRecordResponseRowsInner) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *GetLockedSubscriptionRecordResponseRowsInner) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetTime

`func (o *GetLockedSubscriptionRecordResponseRowsInner) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *GetLockedSubscriptionRecordResponseRowsInner) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *GetLockedSubscriptionRecordResponseRowsInner) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *GetLockedSubscriptionRecordResponseRowsInner) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetAsset

`func (o *GetLockedSubscriptionRecordResponseRowsInner) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *GetLockedSubscriptionRecordResponseRowsInner) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *GetLockedSubscriptionRecordResponseRowsInner) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *GetLockedSubscriptionRecordResponseRowsInner) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetAmount

`func (o *GetLockedSubscriptionRecordResponseRowsInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetLockedSubscriptionRecordResponseRowsInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetLockedSubscriptionRecordResponseRowsInner) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GetLockedSubscriptionRecordResponseRowsInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetLockPeriod

`func (o *GetLockedSubscriptionRecordResponseRowsInner) GetLockPeriod() string`

GetLockPeriod returns the LockPeriod field if non-nil, zero value otherwise.

### GetLockPeriodOk

`func (o *GetLockedSubscriptionRecordResponseRowsInner) GetLockPeriodOk() (*string, bool)`

GetLockPeriodOk returns a tuple with the LockPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockPeriod

`func (o *GetLockedSubscriptionRecordResponseRowsInner) SetLockPeriod(v string)`

SetLockPeriod sets LockPeriod field to given value.

### HasLockPeriod

`func (o *GetLockedSubscriptionRecordResponseRowsInner) HasLockPeriod() bool`

HasLockPeriod returns a boolean if a field has been set.

### GetType

`func (o *GetLockedSubscriptionRecordResponseRowsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetLockedSubscriptionRecordResponseRowsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetLockedSubscriptionRecordResponseRowsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetLockedSubscriptionRecordResponseRowsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSourceAccount

`func (o *GetLockedSubscriptionRecordResponseRowsInner) GetSourceAccount() string`

GetSourceAccount returns the SourceAccount field if non-nil, zero value otherwise.

### GetSourceAccountOk

`func (o *GetLockedSubscriptionRecordResponseRowsInner) GetSourceAccountOk() (*string, bool)`

GetSourceAccountOk returns a tuple with the SourceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccount

`func (o *GetLockedSubscriptionRecordResponseRowsInner) SetSourceAccount(v string)`

SetSourceAccount sets SourceAccount field to given value.

### HasSourceAccount

`func (o *GetLockedSubscriptionRecordResponseRowsInner) HasSourceAccount() bool`

HasSourceAccount returns a boolean if a field has been set.

### GetAmtFromSpot

`func (o *GetLockedSubscriptionRecordResponseRowsInner) GetAmtFromSpot() string`

GetAmtFromSpot returns the AmtFromSpot field if non-nil, zero value otherwise.

### GetAmtFromSpotOk

`func (o *GetLockedSubscriptionRecordResponseRowsInner) GetAmtFromSpotOk() (*string, bool)`

GetAmtFromSpotOk returns a tuple with the AmtFromSpot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmtFromSpot

`func (o *GetLockedSubscriptionRecordResponseRowsInner) SetAmtFromSpot(v string)`

SetAmtFromSpot sets AmtFromSpot field to given value.

### HasAmtFromSpot

`func (o *GetLockedSubscriptionRecordResponseRowsInner) HasAmtFromSpot() bool`

HasAmtFromSpot returns a boolean if a field has been set.

### GetAmtFromFunding

`func (o *GetLockedSubscriptionRecordResponseRowsInner) GetAmtFromFunding() string`

GetAmtFromFunding returns the AmtFromFunding field if non-nil, zero value otherwise.

### GetAmtFromFundingOk

`func (o *GetLockedSubscriptionRecordResponseRowsInner) GetAmtFromFundingOk() (*string, bool)`

GetAmtFromFundingOk returns a tuple with the AmtFromFunding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmtFromFunding

`func (o *GetLockedSubscriptionRecordResponseRowsInner) SetAmtFromFunding(v string)`

SetAmtFromFunding sets AmtFromFunding field to given value.

### HasAmtFromFunding

`func (o *GetLockedSubscriptionRecordResponseRowsInner) HasAmtFromFunding() bool`

HasAmtFromFunding returns a boolean if a field has been set.

### GetStatus

`func (o *GetLockedSubscriptionRecordResponseRowsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetLockedSubscriptionRecordResponseRowsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetLockedSubscriptionRecordResponseRowsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetLockedSubscriptionRecordResponseRowsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to README]](../README.md)


