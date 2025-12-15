# DailyAccountSnapshotResponseSnapshotVosInnerData

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Balances** | Pointer to [**[]DailyAccountSnapshotResponseSnapshotVosInnerDataBalancesInner**](DailyAccountSnapshotResponseSnapshotVosInnerDataBalancesInner.md) |  | [optional] 
**TotalAssetOfBtc** | Pointer to **string** |  | [optional] 
**MarginLevel** | Pointer to **string** |  | [optional] 
**TotalLiabilityOfBtc** | Pointer to **string** |  | [optional] 
**TotalNetAssetOfBtc** | Pointer to **string** |  | [optional] 
**UserAssets** | Pointer to [**[]DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner**](DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner.md) |  | [optional] 
**Assets** | Pointer to [**[]DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner**](DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner.md) |  | [optional] 
**Position** | Pointer to [**[]DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner**](DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner.md) |  | [optional] 

## Methods

### NewDailyAccountSnapshotResponseSnapshotVosInnerData

`func NewDailyAccountSnapshotResponseSnapshotVosInnerData() *DailyAccountSnapshotResponseSnapshotVosInnerData`

NewDailyAccountSnapshotResponseSnapshotVosInnerData instantiates a new DailyAccountSnapshotResponseSnapshotVosInnerData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailyAccountSnapshotResponseSnapshotVosInnerDataWithDefaults

`func NewDailyAccountSnapshotResponseSnapshotVosInnerDataWithDefaults() *DailyAccountSnapshotResponseSnapshotVosInnerData`

NewDailyAccountSnapshotResponseSnapshotVosInnerDataWithDefaults instantiates a new DailyAccountSnapshotResponseSnapshotVosInnerData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalances

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetBalances() []DailyAccountSnapshotResponseSnapshotVosInnerDataBalancesInner`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetBalancesOk() (*[]DailyAccountSnapshotResponseSnapshotVosInnerDataBalancesInner, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) SetBalances(v []DailyAccountSnapshotResponseSnapshotVosInnerDataBalancesInner)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetTotalAssetOfBtc

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetTotalAssetOfBtc() string`

GetTotalAssetOfBtc returns the TotalAssetOfBtc field if non-nil, zero value otherwise.

### GetTotalAssetOfBtcOk

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetTotalAssetOfBtcOk() (*string, bool)`

GetTotalAssetOfBtcOk returns a tuple with the TotalAssetOfBtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAssetOfBtc

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) SetTotalAssetOfBtc(v string)`

SetTotalAssetOfBtc sets TotalAssetOfBtc field to given value.

### HasTotalAssetOfBtc

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) HasTotalAssetOfBtc() bool`

HasTotalAssetOfBtc returns a boolean if a field has been set.

### GetMarginLevel

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetMarginLevel() string`

GetMarginLevel returns the MarginLevel field if non-nil, zero value otherwise.

### GetMarginLevelOk

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetMarginLevelOk() (*string, bool)`

GetMarginLevelOk returns a tuple with the MarginLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginLevel

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) SetMarginLevel(v string)`

SetMarginLevel sets MarginLevel field to given value.

### HasMarginLevel

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) HasMarginLevel() bool`

HasMarginLevel returns a boolean if a field has been set.

### GetTotalLiabilityOfBtc

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetTotalLiabilityOfBtc() string`

GetTotalLiabilityOfBtc returns the TotalLiabilityOfBtc field if non-nil, zero value otherwise.

### GetTotalLiabilityOfBtcOk

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetTotalLiabilityOfBtcOk() (*string, bool)`

GetTotalLiabilityOfBtcOk returns a tuple with the TotalLiabilityOfBtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLiabilityOfBtc

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) SetTotalLiabilityOfBtc(v string)`

SetTotalLiabilityOfBtc sets TotalLiabilityOfBtc field to given value.

### HasTotalLiabilityOfBtc

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) HasTotalLiabilityOfBtc() bool`

HasTotalLiabilityOfBtc returns a boolean if a field has been set.

### GetTotalNetAssetOfBtc

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetTotalNetAssetOfBtc() string`

GetTotalNetAssetOfBtc returns the TotalNetAssetOfBtc field if non-nil, zero value otherwise.

### GetTotalNetAssetOfBtcOk

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetTotalNetAssetOfBtcOk() (*string, bool)`

GetTotalNetAssetOfBtcOk returns a tuple with the TotalNetAssetOfBtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNetAssetOfBtc

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) SetTotalNetAssetOfBtc(v string)`

SetTotalNetAssetOfBtc sets TotalNetAssetOfBtc field to given value.

### HasTotalNetAssetOfBtc

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) HasTotalNetAssetOfBtc() bool`

HasTotalNetAssetOfBtc returns a boolean if a field has been set.

### GetUserAssets

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetUserAssets() []DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner`

GetUserAssets returns the UserAssets field if non-nil, zero value otherwise.

### GetUserAssetsOk

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetUserAssetsOk() (*[]DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner, bool)`

GetUserAssetsOk returns a tuple with the UserAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAssets

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) SetUserAssets(v []DailyAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner)`

SetUserAssets sets UserAssets field to given value.

### HasUserAssets

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) HasUserAssets() bool`

HasUserAssets returns a boolean if a field has been set.

### GetAssets

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetAssets() []DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetAssetsOk() (*[]DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) SetAssets(v []DailyAccountSnapshotResponseSnapshotVosInnerDataAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetPosition

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetPosition() []DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) GetPositionOk() (*[]DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) SetPosition(v []DailyAccountSnapshotResponseSnapshotVosInnerDataPositionInner)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *DailyAccountSnapshotResponseSnapshotVosInnerData) HasPosition() bool`

HasPosition returns a boolean if a field has been set.


[[Back to README]](../README.md)


