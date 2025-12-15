# QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Balances** | Pointer to [**[]QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner**](QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner.md) |  | [optional] 
**TotalAssetOfBtc** | Pointer to **string** |  | [optional] 
**MarginLevel** | Pointer to **string** |  | [optional] 
**TotalLiabilityOfBtc** | Pointer to **string** |  | [optional] 
**TotalNetAssetOfBtc** | Pointer to **string** |  | [optional] 
**UserAssets** | Pointer to [**[]QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner**](QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner.md) |  | [optional] 
**Assets** | Pointer to [**[]QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner**](QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner.md) |  | [optional] 
**Position** | Pointer to [**[]QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataPositionInner**](QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataPositionInner.md) |  | [optional] 

## Methods

### NewQueryManagedSubAccountSnapshotResponseSnapshotVosInnerData

`func NewQueryManagedSubAccountSnapshotResponseSnapshotVosInnerData() *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData`

NewQueryManagedSubAccountSnapshotResponseSnapshotVosInnerData instantiates a new QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataWithDefaults

`func NewQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataWithDefaults() *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData`

NewQueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataWithDefaults instantiates a new QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalances

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetBalances() []QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetBalancesOk() (*[]QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) SetBalances(v []QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataBalancesInner)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetTotalAssetOfBtc

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetTotalAssetOfBtc() string`

GetTotalAssetOfBtc returns the TotalAssetOfBtc field if non-nil, zero value otherwise.

### GetTotalAssetOfBtcOk

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetTotalAssetOfBtcOk() (*string, bool)`

GetTotalAssetOfBtcOk returns a tuple with the TotalAssetOfBtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAssetOfBtc

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) SetTotalAssetOfBtc(v string)`

SetTotalAssetOfBtc sets TotalAssetOfBtc field to given value.

### HasTotalAssetOfBtc

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) HasTotalAssetOfBtc() bool`

HasTotalAssetOfBtc returns a boolean if a field has been set.

### GetMarginLevel

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetMarginLevel() string`

GetMarginLevel returns the MarginLevel field if non-nil, zero value otherwise.

### GetMarginLevelOk

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetMarginLevelOk() (*string, bool)`

GetMarginLevelOk returns a tuple with the MarginLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginLevel

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) SetMarginLevel(v string)`

SetMarginLevel sets MarginLevel field to given value.

### HasMarginLevel

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) HasMarginLevel() bool`

HasMarginLevel returns a boolean if a field has been set.

### GetTotalLiabilityOfBtc

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetTotalLiabilityOfBtc() string`

GetTotalLiabilityOfBtc returns the TotalLiabilityOfBtc field if non-nil, zero value otherwise.

### GetTotalLiabilityOfBtcOk

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetTotalLiabilityOfBtcOk() (*string, bool)`

GetTotalLiabilityOfBtcOk returns a tuple with the TotalLiabilityOfBtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLiabilityOfBtc

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) SetTotalLiabilityOfBtc(v string)`

SetTotalLiabilityOfBtc sets TotalLiabilityOfBtc field to given value.

### HasTotalLiabilityOfBtc

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) HasTotalLiabilityOfBtc() bool`

HasTotalLiabilityOfBtc returns a boolean if a field has been set.

### GetTotalNetAssetOfBtc

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetTotalNetAssetOfBtc() string`

GetTotalNetAssetOfBtc returns the TotalNetAssetOfBtc field if non-nil, zero value otherwise.

### GetTotalNetAssetOfBtcOk

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetTotalNetAssetOfBtcOk() (*string, bool)`

GetTotalNetAssetOfBtcOk returns a tuple with the TotalNetAssetOfBtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNetAssetOfBtc

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) SetTotalNetAssetOfBtc(v string)`

SetTotalNetAssetOfBtc sets TotalNetAssetOfBtc field to given value.

### HasTotalNetAssetOfBtc

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) HasTotalNetAssetOfBtc() bool`

HasTotalNetAssetOfBtc returns a boolean if a field has been set.

### GetUserAssets

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetUserAssets() []QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner`

GetUserAssets returns the UserAssets field if non-nil, zero value otherwise.

### GetUserAssetsOk

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetUserAssetsOk() (*[]QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner, bool)`

GetUserAssetsOk returns a tuple with the UserAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAssets

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) SetUserAssets(v []QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataUserAssetsInner)`

SetUserAssets sets UserAssets field to given value.

### HasUserAssets

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) HasUserAssets() bool`

HasUserAssets returns a boolean if a field has been set.

### GetAssets

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetAssets() []QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetAssetsOk() (*[]QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) SetAssets(v []QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetPosition

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetPosition() []QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataPositionInner`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) GetPositionOk() (*[]QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataPositionInner, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) SetPosition(v []QueryManagedSubAccountSnapshotResponseSnapshotVosInnerDataPositionInner)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *QueryManagedSubAccountSnapshotResponseSnapshotVosInnerData) HasPosition() bool`

HasPosition returns a boolean if a field has been set.


[[Back to README]](../README.md)


