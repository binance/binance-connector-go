# GetDetailOnSubAccountsFuturesAccountResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Asset** | Pointer to **string** |  | [optional] 
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

### NewGetDetailOnSubAccountsFuturesAccountResponse

`func NewGetDetailOnSubAccountsFuturesAccountResponse() *GetDetailOnSubAccountsFuturesAccountResponse`

NewGetDetailOnSubAccountsFuturesAccountResponse instantiates a new GetDetailOnSubAccountsFuturesAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDetailOnSubAccountsFuturesAccountResponseWithDefaults

`func NewGetDetailOnSubAccountsFuturesAccountResponseWithDefaults() *GetDetailOnSubAccountsFuturesAccountResponse`

NewGetDetailOnSubAccountsFuturesAccountResponseWithDefaults instantiates a new GetDetailOnSubAccountsFuturesAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAsset

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetAssets

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetAssets() []GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetAssetsOk() (*[]GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetAssets(v []GetDetailOnSubAccountsFuturesAccountV2ResponseFutureAccountRespAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetCanDeposit

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetCanDeposit() bool`

GetCanDeposit returns the CanDeposit field if non-nil, zero value otherwise.

### GetCanDepositOk

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetCanDepositOk() (*bool, bool)`

GetCanDepositOk returns a tuple with the CanDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeposit

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetCanDeposit(v bool)`

SetCanDeposit sets CanDeposit field to given value.

### HasCanDeposit

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasCanDeposit() bool`

HasCanDeposit returns a boolean if a field has been set.

### GetCanTrade

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetCanTrade() bool`

GetCanTrade returns the CanTrade field if non-nil, zero value otherwise.

### GetCanTradeOk

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetCanTradeOk() (*bool, bool)`

GetCanTradeOk returns a tuple with the CanTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTrade

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetCanTrade(v bool)`

SetCanTrade sets CanTrade field to given value.

### HasCanTrade

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasCanTrade() bool`

HasCanTrade returns a boolean if a field has been set.

### GetCanWithdraw

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetCanWithdraw() bool`

GetCanWithdraw returns the CanWithdraw field if non-nil, zero value otherwise.

### GetCanWithdrawOk

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetCanWithdrawOk() (*bool, bool)`

GetCanWithdrawOk returns a tuple with the CanWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWithdraw

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetCanWithdraw(v bool)`

SetCanWithdraw sets CanWithdraw field to given value.

### HasCanWithdraw

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasCanWithdraw() bool`

HasCanWithdraw returns a boolean if a field has been set.

### GetFeeTier

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetFeeTier() int64`

GetFeeTier returns the FeeTier field if non-nil, zero value otherwise.

### GetFeeTierOk

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetFeeTierOk() (*int64, bool)`

GetFeeTierOk returns a tuple with the FeeTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTier

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetFeeTier(v int64)`

SetFeeTier sets FeeTier field to given value.

### HasFeeTier

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasFeeTier() bool`

HasFeeTier returns a boolean if a field has been set.

### GetMaxWithdrawAmount

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetMaxWithdrawAmount() string`

GetMaxWithdrawAmount returns the MaxWithdrawAmount field if non-nil, zero value otherwise.

### GetMaxWithdrawAmountOk

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetMaxWithdrawAmountOk() (*string, bool)`

GetMaxWithdrawAmountOk returns a tuple with the MaxWithdrawAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWithdrawAmount

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetMaxWithdrawAmount(v string)`

SetMaxWithdrawAmount sets MaxWithdrawAmount field to given value.

### HasMaxWithdrawAmount

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasMaxWithdrawAmount() bool`

HasMaxWithdrawAmount returns a boolean if a field has been set.

### GetTotalInitialMargin

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetTotalInitialMargin() string`

GetTotalInitialMargin returns the TotalInitialMargin field if non-nil, zero value otherwise.

### GetTotalInitialMarginOk

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetTotalInitialMarginOk() (*string, bool)`

GetTotalInitialMarginOk returns a tuple with the TotalInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInitialMargin

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetTotalInitialMargin(v string)`

SetTotalInitialMargin sets TotalInitialMargin field to given value.

### HasTotalInitialMargin

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasTotalInitialMargin() bool`

HasTotalInitialMargin returns a boolean if a field has been set.

### GetTotalMaintenanceMargin

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetTotalMaintenanceMargin() string`

GetTotalMaintenanceMargin returns the TotalMaintenanceMargin field if non-nil, zero value otherwise.

### GetTotalMaintenanceMarginOk

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetTotalMaintenanceMarginOk() (*string, bool)`

GetTotalMaintenanceMarginOk returns a tuple with the TotalMaintenanceMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMaintenanceMargin

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetTotalMaintenanceMargin(v string)`

SetTotalMaintenanceMargin sets TotalMaintenanceMargin field to given value.

### HasTotalMaintenanceMargin

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasTotalMaintenanceMargin() bool`

HasTotalMaintenanceMargin returns a boolean if a field has been set.

### GetTotalMarginBalance

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetTotalMarginBalance() string`

GetTotalMarginBalance returns the TotalMarginBalance field if non-nil, zero value otherwise.

### GetTotalMarginBalanceOk

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetTotalMarginBalanceOk() (*string, bool)`

GetTotalMarginBalanceOk returns a tuple with the TotalMarginBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMarginBalance

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetTotalMarginBalance(v string)`

SetTotalMarginBalance sets TotalMarginBalance field to given value.

### HasTotalMarginBalance

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasTotalMarginBalance() bool`

HasTotalMarginBalance returns a boolean if a field has been set.

### GetTotalOpenOrderInitialMargin

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetTotalOpenOrderInitialMargin() string`

GetTotalOpenOrderInitialMargin returns the TotalOpenOrderInitialMargin field if non-nil, zero value otherwise.

### GetTotalOpenOrderInitialMarginOk

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetTotalOpenOrderInitialMarginOk() (*string, bool)`

GetTotalOpenOrderInitialMarginOk returns a tuple with the TotalOpenOrderInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOpenOrderInitialMargin

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetTotalOpenOrderInitialMargin(v string)`

SetTotalOpenOrderInitialMargin sets TotalOpenOrderInitialMargin field to given value.

### HasTotalOpenOrderInitialMargin

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasTotalOpenOrderInitialMargin() bool`

HasTotalOpenOrderInitialMargin returns a boolean if a field has been set.

### GetTotalPositionInitialMargin

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetTotalPositionInitialMargin() string`

GetTotalPositionInitialMargin returns the TotalPositionInitialMargin field if non-nil, zero value otherwise.

### GetTotalPositionInitialMarginOk

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetTotalPositionInitialMarginOk() (*string, bool)`

GetTotalPositionInitialMarginOk returns a tuple with the TotalPositionInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPositionInitialMargin

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetTotalPositionInitialMargin(v string)`

SetTotalPositionInitialMargin sets TotalPositionInitialMargin field to given value.

### HasTotalPositionInitialMargin

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasTotalPositionInitialMargin() bool`

HasTotalPositionInitialMargin returns a boolean if a field has been set.

### GetTotalUnrealizedProfit

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetTotalUnrealizedProfit() string`

GetTotalUnrealizedProfit returns the TotalUnrealizedProfit field if non-nil, zero value otherwise.

### GetTotalUnrealizedProfitOk

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetTotalUnrealizedProfitOk() (*string, bool)`

GetTotalUnrealizedProfitOk returns a tuple with the TotalUnrealizedProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUnrealizedProfit

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetTotalUnrealizedProfit(v string)`

SetTotalUnrealizedProfit sets TotalUnrealizedProfit field to given value.

### HasTotalUnrealizedProfit

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasTotalUnrealizedProfit() bool`

HasTotalUnrealizedProfit returns a boolean if a field has been set.

### GetTotalWalletBalance

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetTotalWalletBalance() string`

GetTotalWalletBalance returns the TotalWalletBalance field if non-nil, zero value otherwise.

### GetTotalWalletBalanceOk

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetTotalWalletBalanceOk() (*string, bool)`

GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWalletBalance

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetTotalWalletBalance(v string)`

SetTotalWalletBalance sets TotalWalletBalance field to given value.

### HasTotalWalletBalance

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasTotalWalletBalance() bool`

HasTotalWalletBalance returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GetDetailOnSubAccountsFuturesAccountResponse) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


