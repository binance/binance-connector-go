# GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Assets** | Pointer to [**[]GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountRespAssetsInner**](GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountRespAssetsInner.md) |  | [optional] 
**CanDeposit** | Pointer to **bool** |  | [optional] 
**CanTrade** | Pointer to **bool** |  | [optional] 
**CanWithdraw** | Pointer to **bool** |  | [optional] 
**FeeTier** | Pointer to **int64** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp

`func NewGetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp() *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp`

NewGetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp instantiates a new GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountRespWithDefaults

`func NewGetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountRespWithDefaults() *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp`

NewGetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountRespWithDefaults instantiates a new GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAssets

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) GetAssets() []GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountRespAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) GetAssetsOk() (*[]GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountRespAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) SetAssets(v []GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountRespAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetCanDeposit

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) GetCanDeposit() bool`

GetCanDeposit returns the CanDeposit field if non-nil, zero value otherwise.

### GetCanDepositOk

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) GetCanDepositOk() (*bool, bool)`

GetCanDepositOk returns a tuple with the CanDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeposit

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) SetCanDeposit(v bool)`

SetCanDeposit sets CanDeposit field to given value.

### HasCanDeposit

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) HasCanDeposit() bool`

HasCanDeposit returns a boolean if a field has been set.

### GetCanTrade

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) GetCanTrade() bool`

GetCanTrade returns the CanTrade field if non-nil, zero value otherwise.

### GetCanTradeOk

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) GetCanTradeOk() (*bool, bool)`

GetCanTradeOk returns a tuple with the CanTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTrade

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) SetCanTrade(v bool)`

SetCanTrade sets CanTrade field to given value.

### HasCanTrade

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) HasCanTrade() bool`

HasCanTrade returns a boolean if a field has been set.

### GetCanWithdraw

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) GetCanWithdraw() bool`

GetCanWithdraw returns the CanWithdraw field if non-nil, zero value otherwise.

### GetCanWithdrawOk

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) GetCanWithdrawOk() (*bool, bool)`

GetCanWithdrawOk returns a tuple with the CanWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWithdraw

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) SetCanWithdraw(v bool)`

SetCanWithdraw sets CanWithdraw field to given value.

### HasCanWithdraw

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) HasCanWithdraw() bool`

HasCanWithdraw returns a boolean if a field has been set.

### GetFeeTier

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) GetFeeTier() int64`

GetFeeTier returns the FeeTier field if non-nil, zero value otherwise.

### GetFeeTierOk

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) GetFeeTierOk() (*int64, bool)`

GetFeeTierOk returns a tuple with the FeeTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTier

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) SetFeeTier(v int64)`

SetFeeTier sets FeeTier field to given value.

### HasFeeTier

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) HasFeeTier() bool`

HasFeeTier returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseDeliveryAccountResp) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


