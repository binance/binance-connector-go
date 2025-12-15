# AccountInformationV2Response

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**FeeTier** | Pointer to **int64** |  | [optional] 
**FeeBurn** | Pointer to **bool** |  | [optional] 
**CanDeposit** | Pointer to **bool** |  | [optional] 
**CanWithdraw** | Pointer to **bool** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**MultiAssetsMargin** | Pointer to **bool** |  | [optional] 
**TradeGroupId** | Pointer to **int64** |  | [optional] 
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
**Assets** | Pointer to [**[]AccountInformationV2ResponseAssetsInner**](AccountInformationV2ResponseAssetsInner.md) |  | [optional] 
**Positions** | Pointer to [**[]AccountInformationV2ResponsePositionsInner**](AccountInformationV2ResponsePositionsInner.md) |  | [optional] 
**CanTrade** | Pointer to **bool** |  | [optional] 

## Methods

### NewAccountInformationV2Response

`func NewAccountInformationV2Response() *AccountInformationV2Response`

NewAccountInformationV2Response instantiates a new AccountInformationV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountInformationV2ResponseWithDefaults

`func NewAccountInformationV2ResponseWithDefaults() *AccountInformationV2Response`

NewAccountInformationV2ResponseWithDefaults instantiates a new AccountInformationV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeTier

`func (o *AccountInformationV2Response) GetFeeTier() int64`

GetFeeTier returns the FeeTier field if non-nil, zero value otherwise.

### GetFeeTierOk

`func (o *AccountInformationV2Response) GetFeeTierOk() (*int64, bool)`

GetFeeTierOk returns a tuple with the FeeTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTier

`func (o *AccountInformationV2Response) SetFeeTier(v int64)`

SetFeeTier sets FeeTier field to given value.

### HasFeeTier

`func (o *AccountInformationV2Response) HasFeeTier() bool`

HasFeeTier returns a boolean if a field has been set.

### GetFeeBurn

`func (o *AccountInformationV2Response) GetFeeBurn() bool`

GetFeeBurn returns the FeeBurn field if non-nil, zero value otherwise.

### GetFeeBurnOk

`func (o *AccountInformationV2Response) GetFeeBurnOk() (*bool, bool)`

GetFeeBurnOk returns a tuple with the FeeBurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeBurn

`func (o *AccountInformationV2Response) SetFeeBurn(v bool)`

SetFeeBurn sets FeeBurn field to given value.

### HasFeeBurn

`func (o *AccountInformationV2Response) HasFeeBurn() bool`

HasFeeBurn returns a boolean if a field has been set.

### GetCanDeposit

`func (o *AccountInformationV2Response) GetCanDeposit() bool`

GetCanDeposit returns the CanDeposit field if non-nil, zero value otherwise.

### GetCanDepositOk

`func (o *AccountInformationV2Response) GetCanDepositOk() (*bool, bool)`

GetCanDepositOk returns a tuple with the CanDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeposit

`func (o *AccountInformationV2Response) SetCanDeposit(v bool)`

SetCanDeposit sets CanDeposit field to given value.

### HasCanDeposit

`func (o *AccountInformationV2Response) HasCanDeposit() bool`

HasCanDeposit returns a boolean if a field has been set.

### GetCanWithdraw

`func (o *AccountInformationV2Response) GetCanWithdraw() bool`

GetCanWithdraw returns the CanWithdraw field if non-nil, zero value otherwise.

### GetCanWithdrawOk

`func (o *AccountInformationV2Response) GetCanWithdrawOk() (*bool, bool)`

GetCanWithdrawOk returns a tuple with the CanWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWithdraw

`func (o *AccountInformationV2Response) SetCanWithdraw(v bool)`

SetCanWithdraw sets CanWithdraw field to given value.

### HasCanWithdraw

`func (o *AccountInformationV2Response) HasCanWithdraw() bool`

HasCanWithdraw returns a boolean if a field has been set.

### GetUpdateTime

`func (o *AccountInformationV2Response) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *AccountInformationV2Response) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *AccountInformationV2Response) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *AccountInformationV2Response) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetMultiAssetsMargin

`func (o *AccountInformationV2Response) GetMultiAssetsMargin() bool`

GetMultiAssetsMargin returns the MultiAssetsMargin field if non-nil, zero value otherwise.

### GetMultiAssetsMarginOk

`func (o *AccountInformationV2Response) GetMultiAssetsMarginOk() (*bool, bool)`

GetMultiAssetsMarginOk returns a tuple with the MultiAssetsMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiAssetsMargin

`func (o *AccountInformationV2Response) SetMultiAssetsMargin(v bool)`

SetMultiAssetsMargin sets MultiAssetsMargin field to given value.

### HasMultiAssetsMargin

`func (o *AccountInformationV2Response) HasMultiAssetsMargin() bool`

HasMultiAssetsMargin returns a boolean if a field has been set.

### GetTradeGroupId

`func (o *AccountInformationV2Response) GetTradeGroupId() int64`

GetTradeGroupId returns the TradeGroupId field if non-nil, zero value otherwise.

### GetTradeGroupIdOk

`func (o *AccountInformationV2Response) GetTradeGroupIdOk() (*int64, bool)`

GetTradeGroupIdOk returns a tuple with the TradeGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeGroupId

`func (o *AccountInformationV2Response) SetTradeGroupId(v int64)`

SetTradeGroupId sets TradeGroupId field to given value.

### HasTradeGroupId

`func (o *AccountInformationV2Response) HasTradeGroupId() bool`

HasTradeGroupId returns a boolean if a field has been set.

### GetTotalInitialMargin

`func (o *AccountInformationV2Response) GetTotalInitialMargin() string`

GetTotalInitialMargin returns the TotalInitialMargin field if non-nil, zero value otherwise.

### GetTotalInitialMarginOk

`func (o *AccountInformationV2Response) GetTotalInitialMarginOk() (*string, bool)`

GetTotalInitialMarginOk returns a tuple with the TotalInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInitialMargin

`func (o *AccountInformationV2Response) SetTotalInitialMargin(v string)`

SetTotalInitialMargin sets TotalInitialMargin field to given value.

### HasTotalInitialMargin

`func (o *AccountInformationV2Response) HasTotalInitialMargin() bool`

HasTotalInitialMargin returns a boolean if a field has been set.

### GetTotalMaintMargin

`func (o *AccountInformationV2Response) GetTotalMaintMargin() string`

GetTotalMaintMargin returns the TotalMaintMargin field if non-nil, zero value otherwise.

### GetTotalMaintMarginOk

`func (o *AccountInformationV2Response) GetTotalMaintMarginOk() (*string, bool)`

GetTotalMaintMarginOk returns a tuple with the TotalMaintMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMaintMargin

`func (o *AccountInformationV2Response) SetTotalMaintMargin(v string)`

SetTotalMaintMargin sets TotalMaintMargin field to given value.

### HasTotalMaintMargin

`func (o *AccountInformationV2Response) HasTotalMaintMargin() bool`

HasTotalMaintMargin returns a boolean if a field has been set.

### GetTotalWalletBalance

`func (o *AccountInformationV2Response) GetTotalWalletBalance() string`

GetTotalWalletBalance returns the TotalWalletBalance field if non-nil, zero value otherwise.

### GetTotalWalletBalanceOk

`func (o *AccountInformationV2Response) GetTotalWalletBalanceOk() (*string, bool)`

GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWalletBalance

`func (o *AccountInformationV2Response) SetTotalWalletBalance(v string)`

SetTotalWalletBalance sets TotalWalletBalance field to given value.

### HasTotalWalletBalance

`func (o *AccountInformationV2Response) HasTotalWalletBalance() bool`

HasTotalWalletBalance returns a boolean if a field has been set.

### GetTotalUnrealizedProfit

`func (o *AccountInformationV2Response) GetTotalUnrealizedProfit() string`

GetTotalUnrealizedProfit returns the TotalUnrealizedProfit field if non-nil, zero value otherwise.

### GetTotalUnrealizedProfitOk

`func (o *AccountInformationV2Response) GetTotalUnrealizedProfitOk() (*string, bool)`

GetTotalUnrealizedProfitOk returns a tuple with the TotalUnrealizedProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUnrealizedProfit

`func (o *AccountInformationV2Response) SetTotalUnrealizedProfit(v string)`

SetTotalUnrealizedProfit sets TotalUnrealizedProfit field to given value.

### HasTotalUnrealizedProfit

`func (o *AccountInformationV2Response) HasTotalUnrealizedProfit() bool`

HasTotalUnrealizedProfit returns a boolean if a field has been set.

### GetTotalMarginBalance

`func (o *AccountInformationV2Response) GetTotalMarginBalance() string`

GetTotalMarginBalance returns the TotalMarginBalance field if non-nil, zero value otherwise.

### GetTotalMarginBalanceOk

`func (o *AccountInformationV2Response) GetTotalMarginBalanceOk() (*string, bool)`

GetTotalMarginBalanceOk returns a tuple with the TotalMarginBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMarginBalance

`func (o *AccountInformationV2Response) SetTotalMarginBalance(v string)`

SetTotalMarginBalance sets TotalMarginBalance field to given value.

### HasTotalMarginBalance

`func (o *AccountInformationV2Response) HasTotalMarginBalance() bool`

HasTotalMarginBalance returns a boolean if a field has been set.

### GetTotalPositionInitialMargin

`func (o *AccountInformationV2Response) GetTotalPositionInitialMargin() string`

GetTotalPositionInitialMargin returns the TotalPositionInitialMargin field if non-nil, zero value otherwise.

### GetTotalPositionInitialMarginOk

`func (o *AccountInformationV2Response) GetTotalPositionInitialMarginOk() (*string, bool)`

GetTotalPositionInitialMarginOk returns a tuple with the TotalPositionInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPositionInitialMargin

`func (o *AccountInformationV2Response) SetTotalPositionInitialMargin(v string)`

SetTotalPositionInitialMargin sets TotalPositionInitialMargin field to given value.

### HasTotalPositionInitialMargin

`func (o *AccountInformationV2Response) HasTotalPositionInitialMargin() bool`

HasTotalPositionInitialMargin returns a boolean if a field has been set.

### GetTotalOpenOrderInitialMargin

`func (o *AccountInformationV2Response) GetTotalOpenOrderInitialMargin() string`

GetTotalOpenOrderInitialMargin returns the TotalOpenOrderInitialMargin field if non-nil, zero value otherwise.

### GetTotalOpenOrderInitialMarginOk

`func (o *AccountInformationV2Response) GetTotalOpenOrderInitialMarginOk() (*string, bool)`

GetTotalOpenOrderInitialMarginOk returns a tuple with the TotalOpenOrderInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOpenOrderInitialMargin

`func (o *AccountInformationV2Response) SetTotalOpenOrderInitialMargin(v string)`

SetTotalOpenOrderInitialMargin sets TotalOpenOrderInitialMargin field to given value.

### HasTotalOpenOrderInitialMargin

`func (o *AccountInformationV2Response) HasTotalOpenOrderInitialMargin() bool`

HasTotalOpenOrderInitialMargin returns a boolean if a field has been set.

### GetTotalCrossWalletBalance

`func (o *AccountInformationV2Response) GetTotalCrossWalletBalance() string`

GetTotalCrossWalletBalance returns the TotalCrossWalletBalance field if non-nil, zero value otherwise.

### GetTotalCrossWalletBalanceOk

`func (o *AccountInformationV2Response) GetTotalCrossWalletBalanceOk() (*string, bool)`

GetTotalCrossWalletBalanceOk returns a tuple with the TotalCrossWalletBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCrossWalletBalance

`func (o *AccountInformationV2Response) SetTotalCrossWalletBalance(v string)`

SetTotalCrossWalletBalance sets TotalCrossWalletBalance field to given value.

### HasTotalCrossWalletBalance

`func (o *AccountInformationV2Response) HasTotalCrossWalletBalance() bool`

HasTotalCrossWalletBalance returns a boolean if a field has been set.

### GetTotalCrossUnPnl

`func (o *AccountInformationV2Response) GetTotalCrossUnPnl() string`

GetTotalCrossUnPnl returns the TotalCrossUnPnl field if non-nil, zero value otherwise.

### GetTotalCrossUnPnlOk

`func (o *AccountInformationV2Response) GetTotalCrossUnPnlOk() (*string, bool)`

GetTotalCrossUnPnlOk returns a tuple with the TotalCrossUnPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCrossUnPnl

`func (o *AccountInformationV2Response) SetTotalCrossUnPnl(v string)`

SetTotalCrossUnPnl sets TotalCrossUnPnl field to given value.

### HasTotalCrossUnPnl

`func (o *AccountInformationV2Response) HasTotalCrossUnPnl() bool`

HasTotalCrossUnPnl returns a boolean if a field has been set.

### GetAvailableBalance

`func (o *AccountInformationV2Response) GetAvailableBalance() string`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *AccountInformationV2Response) GetAvailableBalanceOk() (*string, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *AccountInformationV2Response) SetAvailableBalance(v string)`

SetAvailableBalance sets AvailableBalance field to given value.

### HasAvailableBalance

`func (o *AccountInformationV2Response) HasAvailableBalance() bool`

HasAvailableBalance returns a boolean if a field has been set.

### GetMaxWithdrawAmount

`func (o *AccountInformationV2Response) GetMaxWithdrawAmount() string`

GetMaxWithdrawAmount returns the MaxWithdrawAmount field if non-nil, zero value otherwise.

### GetMaxWithdrawAmountOk

`func (o *AccountInformationV2Response) GetMaxWithdrawAmountOk() (*string, bool)`

GetMaxWithdrawAmountOk returns a tuple with the MaxWithdrawAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWithdrawAmount

`func (o *AccountInformationV2Response) SetMaxWithdrawAmount(v string)`

SetMaxWithdrawAmount sets MaxWithdrawAmount field to given value.

### HasMaxWithdrawAmount

`func (o *AccountInformationV2Response) HasMaxWithdrawAmount() bool`

HasMaxWithdrawAmount returns a boolean if a field has been set.

### GetAssets

`func (o *AccountInformationV2Response) GetAssets() []AccountInformationV2ResponseAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *AccountInformationV2Response) GetAssetsOk() (*[]AccountInformationV2ResponseAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *AccountInformationV2Response) SetAssets(v []AccountInformationV2ResponseAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *AccountInformationV2Response) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetPositions

`func (o *AccountInformationV2Response) GetPositions() []AccountInformationV2ResponsePositionsInner`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *AccountInformationV2Response) GetPositionsOk() (*[]AccountInformationV2ResponsePositionsInner, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *AccountInformationV2Response) SetPositions(v []AccountInformationV2ResponsePositionsInner)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *AccountInformationV2Response) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetCanTrade

`func (o *AccountInformationV2Response) GetCanTrade() bool`

GetCanTrade returns the CanTrade field if non-nil, zero value otherwise.

### GetCanTradeOk

`func (o *AccountInformationV2Response) GetCanTradeOk() (*bool, bool)`

GetCanTradeOk returns a tuple with the CanTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTrade

`func (o *AccountInformationV2Response) SetCanTrade(v bool)`

SetCanTrade sets CanTrade field to given value.

### HasCanTrade

`func (o *AccountInformationV2Response) HasCanTrade() bool`

HasCanTrade returns a boolean if a field has been set.


[[Back to README]](../README.md)


