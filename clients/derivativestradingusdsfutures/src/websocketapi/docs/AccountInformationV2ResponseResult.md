# AccountInformationV2ResponseResult

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**TotalInitialMargin** | Pointer to **string** |  | [optional] 
**TotalMaintMargin** | Pointer to **string** |  | [optional] 
**TotalWalletBalance** | Pointer to **string** |  | [optional] 
**TotalUnrealizedProfit** | Pointer to **string** |  | [optional] 
**TotalMarginBalance** | Pointer to **string** |  | [optional] 
**TotalPositionInitialMargin** | Pointer to **string** |  | [optional] 
**TotalOpenOrderInitialMargin** | Pointer to **string** |  | [optional] 
**TotalCrossWalletBalance** | Pointer to **string** |  | [optional] 
**TotalCrossUnPnl** | Pointer to **string** |  | [optional] 
**AvailableBalance** | Pointer to **string** |  | [optional] 
**MaxWithdrawAmount** | Pointer to **string** |  | [optional] 
**Assets** | Pointer to [**[]AccountInformationV2ResponseResultAssetsInner**](AccountInformationV2ResponseResultAssetsInner.md) |  | [optional] 
**Positions** | Pointer to [**[]AccountInformationV2ResponseResultPositionsInner**](AccountInformationV2ResponseResultPositionsInner.md) |  | [optional] 

## Methods

### NewAccountInformationV2ResponseResult

`func NewAccountInformationV2ResponseResult() *AccountInformationV2ResponseResult`

NewAccountInformationV2ResponseResult instantiates a new AccountInformationV2ResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountInformationV2ResponseResultWithDefaults

`func NewAccountInformationV2ResponseResultWithDefaults() *AccountInformationV2ResponseResult`

NewAccountInformationV2ResponseResultWithDefaults instantiates a new AccountInformationV2ResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalInitialMargin

`func (o *AccountInformationV2ResponseResult) GetTotalInitialMargin() string`

GetTotalInitialMargin returns the TotalInitialMargin field if non-nil, zero value otherwise.

### GetTotalInitialMarginOk

`func (o *AccountInformationV2ResponseResult) GetTotalInitialMarginOk() (*string, bool)`

GetTotalInitialMarginOk returns a tuple with the TotalInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInitialMargin

`func (o *AccountInformationV2ResponseResult) SetTotalInitialMargin(v string)`

SetTotalInitialMargin sets TotalInitialMargin field to given value.

### HasTotalInitialMargin

`func (o *AccountInformationV2ResponseResult) HasTotalInitialMargin() bool`

HasTotalInitialMargin returns a boolean if a field has been set.

### GetTotalMaintMargin

`func (o *AccountInformationV2ResponseResult) GetTotalMaintMargin() string`

GetTotalMaintMargin returns the TotalMaintMargin field if non-nil, zero value otherwise.

### GetTotalMaintMarginOk

`func (o *AccountInformationV2ResponseResult) GetTotalMaintMarginOk() (*string, bool)`

GetTotalMaintMarginOk returns a tuple with the TotalMaintMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMaintMargin

`func (o *AccountInformationV2ResponseResult) SetTotalMaintMargin(v string)`

SetTotalMaintMargin sets TotalMaintMargin field to given value.

### HasTotalMaintMargin

`func (o *AccountInformationV2ResponseResult) HasTotalMaintMargin() bool`

HasTotalMaintMargin returns a boolean if a field has been set.

### GetTotalWalletBalance

`func (o *AccountInformationV2ResponseResult) GetTotalWalletBalance() string`

GetTotalWalletBalance returns the TotalWalletBalance field if non-nil, zero value otherwise.

### GetTotalWalletBalanceOk

`func (o *AccountInformationV2ResponseResult) GetTotalWalletBalanceOk() (*string, bool)`

GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWalletBalance

`func (o *AccountInformationV2ResponseResult) SetTotalWalletBalance(v string)`

SetTotalWalletBalance sets TotalWalletBalance field to given value.

### HasTotalWalletBalance

`func (o *AccountInformationV2ResponseResult) HasTotalWalletBalance() bool`

HasTotalWalletBalance returns a boolean if a field has been set.

### GetTotalUnrealizedProfit

`func (o *AccountInformationV2ResponseResult) GetTotalUnrealizedProfit() string`

GetTotalUnrealizedProfit returns the TotalUnrealizedProfit field if non-nil, zero value otherwise.

### GetTotalUnrealizedProfitOk

`func (o *AccountInformationV2ResponseResult) GetTotalUnrealizedProfitOk() (*string, bool)`

GetTotalUnrealizedProfitOk returns a tuple with the TotalUnrealizedProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUnrealizedProfit

`func (o *AccountInformationV2ResponseResult) SetTotalUnrealizedProfit(v string)`

SetTotalUnrealizedProfit sets TotalUnrealizedProfit field to given value.

### HasTotalUnrealizedProfit

`func (o *AccountInformationV2ResponseResult) HasTotalUnrealizedProfit() bool`

HasTotalUnrealizedProfit returns a boolean if a field has been set.

### GetTotalMarginBalance

`func (o *AccountInformationV2ResponseResult) GetTotalMarginBalance() string`

GetTotalMarginBalance returns the TotalMarginBalance field if non-nil, zero value otherwise.

### GetTotalMarginBalanceOk

`func (o *AccountInformationV2ResponseResult) GetTotalMarginBalanceOk() (*string, bool)`

GetTotalMarginBalanceOk returns a tuple with the TotalMarginBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMarginBalance

`func (o *AccountInformationV2ResponseResult) SetTotalMarginBalance(v string)`

SetTotalMarginBalance sets TotalMarginBalance field to given value.

### HasTotalMarginBalance

`func (o *AccountInformationV2ResponseResult) HasTotalMarginBalance() bool`

HasTotalMarginBalance returns a boolean if a field has been set.

### GetTotalPositionInitialMargin

`func (o *AccountInformationV2ResponseResult) GetTotalPositionInitialMargin() string`

GetTotalPositionInitialMargin returns the TotalPositionInitialMargin field if non-nil, zero value otherwise.

### GetTotalPositionInitialMarginOk

`func (o *AccountInformationV2ResponseResult) GetTotalPositionInitialMarginOk() (*string, bool)`

GetTotalPositionInitialMarginOk returns a tuple with the TotalPositionInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPositionInitialMargin

`func (o *AccountInformationV2ResponseResult) SetTotalPositionInitialMargin(v string)`

SetTotalPositionInitialMargin sets TotalPositionInitialMargin field to given value.

### HasTotalPositionInitialMargin

`func (o *AccountInformationV2ResponseResult) HasTotalPositionInitialMargin() bool`

HasTotalPositionInitialMargin returns a boolean if a field has been set.

### GetTotalOpenOrderInitialMargin

`func (o *AccountInformationV2ResponseResult) GetTotalOpenOrderInitialMargin() string`

GetTotalOpenOrderInitialMargin returns the TotalOpenOrderInitialMargin field if non-nil, zero value otherwise.

### GetTotalOpenOrderInitialMarginOk

`func (o *AccountInformationV2ResponseResult) GetTotalOpenOrderInitialMarginOk() (*string, bool)`

GetTotalOpenOrderInitialMarginOk returns a tuple with the TotalOpenOrderInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOpenOrderInitialMargin

`func (o *AccountInformationV2ResponseResult) SetTotalOpenOrderInitialMargin(v string)`

SetTotalOpenOrderInitialMargin sets TotalOpenOrderInitialMargin field to given value.

### HasTotalOpenOrderInitialMargin

`func (o *AccountInformationV2ResponseResult) HasTotalOpenOrderInitialMargin() bool`

HasTotalOpenOrderInitialMargin returns a boolean if a field has been set.

### GetTotalCrossWalletBalance

`func (o *AccountInformationV2ResponseResult) GetTotalCrossWalletBalance() string`

GetTotalCrossWalletBalance returns the TotalCrossWalletBalance field if non-nil, zero value otherwise.

### GetTotalCrossWalletBalanceOk

`func (o *AccountInformationV2ResponseResult) GetTotalCrossWalletBalanceOk() (*string, bool)`

GetTotalCrossWalletBalanceOk returns a tuple with the TotalCrossWalletBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCrossWalletBalance

`func (o *AccountInformationV2ResponseResult) SetTotalCrossWalletBalance(v string)`

SetTotalCrossWalletBalance sets TotalCrossWalletBalance field to given value.

### HasTotalCrossWalletBalance

`func (o *AccountInformationV2ResponseResult) HasTotalCrossWalletBalance() bool`

HasTotalCrossWalletBalance returns a boolean if a field has been set.

### GetTotalCrossUnPnl

`func (o *AccountInformationV2ResponseResult) GetTotalCrossUnPnl() string`

GetTotalCrossUnPnl returns the TotalCrossUnPnl field if non-nil, zero value otherwise.

### GetTotalCrossUnPnlOk

`func (o *AccountInformationV2ResponseResult) GetTotalCrossUnPnlOk() (*string, bool)`

GetTotalCrossUnPnlOk returns a tuple with the TotalCrossUnPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCrossUnPnl

`func (o *AccountInformationV2ResponseResult) SetTotalCrossUnPnl(v string)`

SetTotalCrossUnPnl sets TotalCrossUnPnl field to given value.

### HasTotalCrossUnPnl

`func (o *AccountInformationV2ResponseResult) HasTotalCrossUnPnl() bool`

HasTotalCrossUnPnl returns a boolean if a field has been set.

### GetAvailableBalance

`func (o *AccountInformationV2ResponseResult) GetAvailableBalance() string`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *AccountInformationV2ResponseResult) GetAvailableBalanceOk() (*string, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *AccountInformationV2ResponseResult) SetAvailableBalance(v string)`

SetAvailableBalance sets AvailableBalance field to given value.

### HasAvailableBalance

`func (o *AccountInformationV2ResponseResult) HasAvailableBalance() bool`

HasAvailableBalance returns a boolean if a field has been set.

### GetMaxWithdrawAmount

`func (o *AccountInformationV2ResponseResult) GetMaxWithdrawAmount() string`

GetMaxWithdrawAmount returns the MaxWithdrawAmount field if non-nil, zero value otherwise.

### GetMaxWithdrawAmountOk

`func (o *AccountInformationV2ResponseResult) GetMaxWithdrawAmountOk() (*string, bool)`

GetMaxWithdrawAmountOk returns a tuple with the MaxWithdrawAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWithdrawAmount

`func (o *AccountInformationV2ResponseResult) SetMaxWithdrawAmount(v string)`

SetMaxWithdrawAmount sets MaxWithdrawAmount field to given value.

### HasMaxWithdrawAmount

`func (o *AccountInformationV2ResponseResult) HasMaxWithdrawAmount() bool`

HasMaxWithdrawAmount returns a boolean if a field has been set.

### GetAssets

`func (o *AccountInformationV2ResponseResult) GetAssets() []AccountInformationV2ResponseResultAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *AccountInformationV2ResponseResult) GetAssetsOk() (*[]AccountInformationV2ResponseResultAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *AccountInformationV2ResponseResult) SetAssets(v []AccountInformationV2ResponseResultAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *AccountInformationV2ResponseResult) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetPositions

`func (o *AccountInformationV2ResponseResult) GetPositions() []AccountInformationV2ResponseResultPositionsInner`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *AccountInformationV2ResponseResult) GetPositionsOk() (*[]AccountInformationV2ResponseResultPositionsInner, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *AccountInformationV2ResponseResult) SetPositions(v []AccountInformationV2ResponseResultPositionsInner)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *AccountInformationV2ResponseResult) HasPositions() bool`

HasPositions returns a boolean if a field has been set.


[[Back to README]](../README.md)


