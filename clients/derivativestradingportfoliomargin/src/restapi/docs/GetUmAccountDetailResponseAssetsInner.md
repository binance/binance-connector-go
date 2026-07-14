# GetUmAccountDetailResponseAssetsInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Asset** | Pointer to **string** | asset name | [optional] 
**CrossWalletBalance** | Pointer to **string** | wallet balance | [optional] 
**CrossUnPnl** | Pointer to **string** | unrealized profit | [optional] 
**MaintMargin** | Pointer to **string** | maintenance margin required | [optional] 
**InitialMargin** | Pointer to **string** | total initial margin required with current mark price | [optional] 
**PositionInitialMargin** | Pointer to **string** | initial margin required for positions with current mark price | [optional] 
**OpenOrderInitialMargin** | Pointer to **string** | initial margin required for open orders with current mark price | [optional] 
**UpdateTime** | Pointer to **int64** | last update time | [optional] 

## Methods

### NewGetUmAccountDetailResponseAssetsInner

`func NewGetUmAccountDetailResponseAssetsInner() *GetUmAccountDetailResponseAssetsInner`

NewGetUmAccountDetailResponseAssetsInner instantiates a new GetUmAccountDetailResponseAssetsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUmAccountDetailResponseAssetsInnerWithDefaults

`func NewGetUmAccountDetailResponseAssetsInnerWithDefaults() *GetUmAccountDetailResponseAssetsInner`

NewGetUmAccountDetailResponseAssetsInnerWithDefaults instantiates a new GetUmAccountDetailResponseAssetsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *GetUmAccountDetailResponseAssetsInner) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *GetUmAccountDetailResponseAssetsInner) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *GetUmAccountDetailResponseAssetsInner) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *GetUmAccountDetailResponseAssetsInner) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetCrossWalletBalance

`func (o *GetUmAccountDetailResponseAssetsInner) GetCrossWalletBalance() string`

GetCrossWalletBalance returns the CrossWalletBalance field if non-nil, zero value otherwise.

### GetCrossWalletBalanceOk

`func (o *GetUmAccountDetailResponseAssetsInner) GetCrossWalletBalanceOk() (*string, bool)`

GetCrossWalletBalanceOk returns a tuple with the CrossWalletBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossWalletBalance

`func (o *GetUmAccountDetailResponseAssetsInner) SetCrossWalletBalance(v string)`

SetCrossWalletBalance sets CrossWalletBalance field to given value.

### HasCrossWalletBalance

`func (o *GetUmAccountDetailResponseAssetsInner) HasCrossWalletBalance() bool`

HasCrossWalletBalance returns a boolean if a field has been set.

### GetCrossUnPnl

`func (o *GetUmAccountDetailResponseAssetsInner) GetCrossUnPnl() string`

GetCrossUnPnl returns the CrossUnPnl field if non-nil, zero value otherwise.

### GetCrossUnPnlOk

`func (o *GetUmAccountDetailResponseAssetsInner) GetCrossUnPnlOk() (*string, bool)`

GetCrossUnPnlOk returns a tuple with the CrossUnPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossUnPnl

`func (o *GetUmAccountDetailResponseAssetsInner) SetCrossUnPnl(v string)`

SetCrossUnPnl sets CrossUnPnl field to given value.

### HasCrossUnPnl

`func (o *GetUmAccountDetailResponseAssetsInner) HasCrossUnPnl() bool`

HasCrossUnPnl returns a boolean if a field has been set.

### GetMaintMargin

`func (o *GetUmAccountDetailResponseAssetsInner) GetMaintMargin() string`

GetMaintMargin returns the MaintMargin field if non-nil, zero value otherwise.

### GetMaintMarginOk

`func (o *GetUmAccountDetailResponseAssetsInner) GetMaintMarginOk() (*string, bool)`

GetMaintMarginOk returns a tuple with the MaintMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintMargin

`func (o *GetUmAccountDetailResponseAssetsInner) SetMaintMargin(v string)`

SetMaintMargin sets MaintMargin field to given value.

### HasMaintMargin

`func (o *GetUmAccountDetailResponseAssetsInner) HasMaintMargin() bool`

HasMaintMargin returns a boolean if a field has been set.

### GetInitialMargin

`func (o *GetUmAccountDetailResponseAssetsInner) GetInitialMargin() string`

GetInitialMargin returns the InitialMargin field if non-nil, zero value otherwise.

### GetInitialMarginOk

`func (o *GetUmAccountDetailResponseAssetsInner) GetInitialMarginOk() (*string, bool)`

GetInitialMarginOk returns a tuple with the InitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialMargin

`func (o *GetUmAccountDetailResponseAssetsInner) SetInitialMargin(v string)`

SetInitialMargin sets InitialMargin field to given value.

### HasInitialMargin

`func (o *GetUmAccountDetailResponseAssetsInner) HasInitialMargin() bool`

HasInitialMargin returns a boolean if a field has been set.

### GetPositionInitialMargin

`func (o *GetUmAccountDetailResponseAssetsInner) GetPositionInitialMargin() string`

GetPositionInitialMargin returns the PositionInitialMargin field if non-nil, zero value otherwise.

### GetPositionInitialMarginOk

`func (o *GetUmAccountDetailResponseAssetsInner) GetPositionInitialMarginOk() (*string, bool)`

GetPositionInitialMarginOk returns a tuple with the PositionInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionInitialMargin

`func (o *GetUmAccountDetailResponseAssetsInner) SetPositionInitialMargin(v string)`

SetPositionInitialMargin sets PositionInitialMargin field to given value.

### HasPositionInitialMargin

`func (o *GetUmAccountDetailResponseAssetsInner) HasPositionInitialMargin() bool`

HasPositionInitialMargin returns a boolean if a field has been set.

### GetOpenOrderInitialMargin

`func (o *GetUmAccountDetailResponseAssetsInner) GetOpenOrderInitialMargin() string`

GetOpenOrderInitialMargin returns the OpenOrderInitialMargin field if non-nil, zero value otherwise.

### GetOpenOrderInitialMarginOk

`func (o *GetUmAccountDetailResponseAssetsInner) GetOpenOrderInitialMarginOk() (*string, bool)`

GetOpenOrderInitialMarginOk returns a tuple with the OpenOrderInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenOrderInitialMargin

`func (o *GetUmAccountDetailResponseAssetsInner) SetOpenOrderInitialMargin(v string)`

SetOpenOrderInitialMargin sets OpenOrderInitialMargin field to given value.

### HasOpenOrderInitialMargin

`func (o *GetUmAccountDetailResponseAssetsInner) HasOpenOrderInitialMargin() bool`

HasOpenOrderInitialMargin returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GetUmAccountDetailResponseAssetsInner) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GetUmAccountDetailResponseAssetsInner) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GetUmAccountDetailResponseAssetsInner) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GetUmAccountDetailResponseAssetsInner) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


