# GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Assets** | Pointer to [**[]GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner**](GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner.md) |  | [optional] 
**CanDeposit** | Pointer to **bool** |  | [optional] 
**CanTrade** | Pointer to **bool** |  | [optional] 
**CanWithdraw** | Pointer to **bool** |  | [optional] 
**FeeTier** | Pointer to **int64** |  | [optional] 
**MaxWithdrawAmount** | Pointer to **string** |  | [optional] 
**TotalInitialMargin** | Pointer to **string** |  | [optional] 
**TotalMaintenanceMargin** | Pointer to **string** |  | [optional] 
**TotalMarginBalance** | Pointer to **string** |  | [optional] 
**TotalOpenOrderInitialMargin** | Pointer to **string** |  | [optional] 
**TotalPositionInitialMargin** | Pointer to **string** |  | [optional] 
**TotalUnrealizedProfit** | Pointer to **string** |  | [optional] 
**TotalWalletBalance** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp

`func NewGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp() *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp`

NewGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp instantiates a new GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespWithDefaults

`func NewGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespWithDefaults() *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp`

NewGetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespWithDefaults instantiates a new GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAssets

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetAssets() []GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetAssetsOk() (*[]GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) SetAssets(v []GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetCanDeposit

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetCanDeposit() bool`

GetCanDeposit returns the CanDeposit field if non-nil, zero value otherwise.

### GetCanDepositOk

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetCanDepositOk() (*bool, bool)`

GetCanDepositOk returns a tuple with the CanDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeposit

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) SetCanDeposit(v bool)`

SetCanDeposit sets CanDeposit field to given value.

### HasCanDeposit

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) HasCanDeposit() bool`

HasCanDeposit returns a boolean if a field has been set.

### GetCanTrade

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetCanTrade() bool`

GetCanTrade returns the CanTrade field if non-nil, zero value otherwise.

### GetCanTradeOk

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetCanTradeOk() (*bool, bool)`

GetCanTradeOk returns a tuple with the CanTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTrade

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) SetCanTrade(v bool)`

SetCanTrade sets CanTrade field to given value.

### HasCanTrade

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) HasCanTrade() bool`

HasCanTrade returns a boolean if a field has been set.

### GetCanWithdraw

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetCanWithdraw() bool`

GetCanWithdraw returns the CanWithdraw field if non-nil, zero value otherwise.

### GetCanWithdrawOk

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetCanWithdrawOk() (*bool, bool)`

GetCanWithdrawOk returns a tuple with the CanWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWithdraw

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) SetCanWithdraw(v bool)`

SetCanWithdraw sets CanWithdraw field to given value.

### HasCanWithdraw

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) HasCanWithdraw() bool`

HasCanWithdraw returns a boolean if a field has been set.

### GetFeeTier

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetFeeTier() int64`

GetFeeTier returns the FeeTier field if non-nil, zero value otherwise.

### GetFeeTierOk

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetFeeTierOk() (*int64, bool)`

GetFeeTierOk returns a tuple with the FeeTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTier

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) SetFeeTier(v int64)`

SetFeeTier sets FeeTier field to given value.

### HasFeeTier

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) HasFeeTier() bool`

HasFeeTier returns a boolean if a field has been set.

### GetMaxWithdrawAmount

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetMaxWithdrawAmount() string`

GetMaxWithdrawAmount returns the MaxWithdrawAmount field if non-nil, zero value otherwise.

### GetMaxWithdrawAmountOk

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetMaxWithdrawAmountOk() (*string, bool)`

GetMaxWithdrawAmountOk returns a tuple with the MaxWithdrawAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWithdrawAmount

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) SetMaxWithdrawAmount(v string)`

SetMaxWithdrawAmount sets MaxWithdrawAmount field to given value.

### HasMaxWithdrawAmount

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) HasMaxWithdrawAmount() bool`

HasMaxWithdrawAmount returns a boolean if a field has been set.

### GetTotalInitialMargin

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetTotalInitialMargin() string`

GetTotalInitialMargin returns the TotalInitialMargin field if non-nil, zero value otherwise.

### GetTotalInitialMarginOk

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetTotalInitialMarginOk() (*string, bool)`

GetTotalInitialMarginOk returns a tuple with the TotalInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInitialMargin

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) SetTotalInitialMargin(v string)`

SetTotalInitialMargin sets TotalInitialMargin field to given value.

### HasTotalInitialMargin

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) HasTotalInitialMargin() bool`

HasTotalInitialMargin returns a boolean if a field has been set.

### GetTotalMaintenanceMargin

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetTotalMaintenanceMargin() string`

GetTotalMaintenanceMargin returns the TotalMaintenanceMargin field if non-nil, zero value otherwise.

### GetTotalMaintenanceMarginOk

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetTotalMaintenanceMarginOk() (*string, bool)`

GetTotalMaintenanceMarginOk returns a tuple with the TotalMaintenanceMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMaintenanceMargin

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) SetTotalMaintenanceMargin(v string)`

SetTotalMaintenanceMargin sets TotalMaintenanceMargin field to given value.

### HasTotalMaintenanceMargin

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) HasTotalMaintenanceMargin() bool`

HasTotalMaintenanceMargin returns a boolean if a field has been set.

### GetTotalMarginBalance

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetTotalMarginBalance() string`

GetTotalMarginBalance returns the TotalMarginBalance field if non-nil, zero value otherwise.

### GetTotalMarginBalanceOk

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetTotalMarginBalanceOk() (*string, bool)`

GetTotalMarginBalanceOk returns a tuple with the TotalMarginBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMarginBalance

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) SetTotalMarginBalance(v string)`

SetTotalMarginBalance sets TotalMarginBalance field to given value.

### HasTotalMarginBalance

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) HasTotalMarginBalance() bool`

HasTotalMarginBalance returns a boolean if a field has been set.

### GetTotalOpenOrderInitialMargin

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetTotalOpenOrderInitialMargin() string`

GetTotalOpenOrderInitialMargin returns the TotalOpenOrderInitialMargin field if non-nil, zero value otherwise.

### GetTotalOpenOrderInitialMarginOk

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetTotalOpenOrderInitialMarginOk() (*string, bool)`

GetTotalOpenOrderInitialMarginOk returns a tuple with the TotalOpenOrderInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOpenOrderInitialMargin

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) SetTotalOpenOrderInitialMargin(v string)`

SetTotalOpenOrderInitialMargin sets TotalOpenOrderInitialMargin field to given value.

### HasTotalOpenOrderInitialMargin

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) HasTotalOpenOrderInitialMargin() bool`

HasTotalOpenOrderInitialMargin returns a boolean if a field has been set.

### GetTotalPositionInitialMargin

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetTotalPositionInitialMargin() string`

GetTotalPositionInitialMargin returns the TotalPositionInitialMargin field if non-nil, zero value otherwise.

### GetTotalPositionInitialMarginOk

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetTotalPositionInitialMarginOk() (*string, bool)`

GetTotalPositionInitialMarginOk returns a tuple with the TotalPositionInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPositionInitialMargin

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) SetTotalPositionInitialMargin(v string)`

SetTotalPositionInitialMargin sets TotalPositionInitialMargin field to given value.

### HasTotalPositionInitialMargin

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) HasTotalPositionInitialMargin() bool`

HasTotalPositionInitialMargin returns a boolean if a field has been set.

### GetTotalUnrealizedProfit

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetTotalUnrealizedProfit() string`

GetTotalUnrealizedProfit returns the TotalUnrealizedProfit field if non-nil, zero value otherwise.

### GetTotalUnrealizedProfitOk

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetTotalUnrealizedProfitOk() (*string, bool)`

GetTotalUnrealizedProfitOk returns a tuple with the TotalUnrealizedProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUnrealizedProfit

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) SetTotalUnrealizedProfit(v string)`

SetTotalUnrealizedProfit sets TotalUnrealizedProfit field to given value.

### HasTotalUnrealizedProfit

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) HasTotalUnrealizedProfit() bool`

HasTotalUnrealizedProfit returns a boolean if a field has been set.

### GetTotalWalletBalance

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetTotalWalletBalance() string`

GetTotalWalletBalance returns the TotalWalletBalance field if non-nil, zero value otherwise.

### GetTotalWalletBalanceOk

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetTotalWalletBalanceOk() (*string, bool)`

GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWalletBalance

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) SetTotalWalletBalance(v string)`

SetTotalWalletBalance sets TotalWalletBalance field to given value.

### HasTotalWalletBalance

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) HasTotalWalletBalance() bool`

HasTotalWalletBalance returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountResp) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


