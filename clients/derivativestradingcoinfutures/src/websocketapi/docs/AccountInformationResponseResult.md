# AccountInformationResponseResult

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**FeeTier** | Pointer to **int64** |  | [optional] 
**CanTrade** | Pointer to **bool** |  | [optional] 
**CanDeposit** | Pointer to **bool** |  | [optional] 
**CanWithdraw** | Pointer to **bool** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**Assets** | Pointer to [**[]AccountInformationResponseResultAssetsInner**](AccountInformationResponseResultAssetsInner.md) |  | [optional] 
**Positions** | Pointer to [**[]AccountInformationResponseResultPositionsInner**](AccountInformationResponseResultPositionsInner.md) |  | [optional] 

## Methods

### NewAccountInformationResponseResult

`func NewAccountInformationResponseResult() *AccountInformationResponseResult`

NewAccountInformationResponseResult instantiates a new AccountInformationResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountInformationResponseResultWithDefaults

`func NewAccountInformationResponseResultWithDefaults() *AccountInformationResponseResult`

NewAccountInformationResponseResultWithDefaults instantiates a new AccountInformationResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeTier

`func (o *AccountInformationResponseResult) GetFeeTier() int64`

GetFeeTier returns the FeeTier field if non-nil, zero value otherwise.

### GetFeeTierOk

`func (o *AccountInformationResponseResult) GetFeeTierOk() (*int64, bool)`

GetFeeTierOk returns a tuple with the FeeTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTier

`func (o *AccountInformationResponseResult) SetFeeTier(v int64)`

SetFeeTier sets FeeTier field to given value.

### HasFeeTier

`func (o *AccountInformationResponseResult) HasFeeTier() bool`

HasFeeTier returns a boolean if a field has been set.

### GetCanTrade

`func (o *AccountInformationResponseResult) GetCanTrade() bool`

GetCanTrade returns the CanTrade field if non-nil, zero value otherwise.

### GetCanTradeOk

`func (o *AccountInformationResponseResult) GetCanTradeOk() (*bool, bool)`

GetCanTradeOk returns a tuple with the CanTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTrade

`func (o *AccountInformationResponseResult) SetCanTrade(v bool)`

SetCanTrade sets CanTrade field to given value.

### HasCanTrade

`func (o *AccountInformationResponseResult) HasCanTrade() bool`

HasCanTrade returns a boolean if a field has been set.

### GetCanDeposit

`func (o *AccountInformationResponseResult) GetCanDeposit() bool`

GetCanDeposit returns the CanDeposit field if non-nil, zero value otherwise.

### GetCanDepositOk

`func (o *AccountInformationResponseResult) GetCanDepositOk() (*bool, bool)`

GetCanDepositOk returns a tuple with the CanDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeposit

`func (o *AccountInformationResponseResult) SetCanDeposit(v bool)`

SetCanDeposit sets CanDeposit field to given value.

### HasCanDeposit

`func (o *AccountInformationResponseResult) HasCanDeposit() bool`

HasCanDeposit returns a boolean if a field has been set.

### GetCanWithdraw

`func (o *AccountInformationResponseResult) GetCanWithdraw() bool`

GetCanWithdraw returns the CanWithdraw field if non-nil, zero value otherwise.

### GetCanWithdrawOk

`func (o *AccountInformationResponseResult) GetCanWithdrawOk() (*bool, bool)`

GetCanWithdrawOk returns a tuple with the CanWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWithdraw

`func (o *AccountInformationResponseResult) SetCanWithdraw(v bool)`

SetCanWithdraw sets CanWithdraw field to given value.

### HasCanWithdraw

`func (o *AccountInformationResponseResult) HasCanWithdraw() bool`

HasCanWithdraw returns a boolean if a field has been set.

### GetUpdateTime

`func (o *AccountInformationResponseResult) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *AccountInformationResponseResult) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *AccountInformationResponseResult) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *AccountInformationResponseResult) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetAssets

`func (o *AccountInformationResponseResult) GetAssets() []AccountInformationResponseResultAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *AccountInformationResponseResult) GetAssetsOk() (*[]AccountInformationResponseResultAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *AccountInformationResponseResult) SetAssets(v []AccountInformationResponseResultAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *AccountInformationResponseResult) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetPositions

`func (o *AccountInformationResponseResult) GetPositions() []AccountInformationResponseResultPositionsInner`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *AccountInformationResponseResult) GetPositionsOk() (*[]AccountInformationResponseResultPositionsInner, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *AccountInformationResponseResult) SetPositions(v []AccountInformationResponseResultPositionsInner)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *AccountInformationResponseResult) HasPositions() bool`

HasPositions returns a boolean if a field has been set.


[[Back to README]](../README.md)


