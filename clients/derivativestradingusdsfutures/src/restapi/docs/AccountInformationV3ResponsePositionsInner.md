# AccountInformationV3ResponsePositionsInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**PositionSide** | Pointer to **string** |  | [optional] 
**PositionAmt** | Pointer to **string** |  | [optional] 
**UnrealizedProfit** | Pointer to **string** |  | [optional] 
**IsolatedMargin** | Pointer to **string** |  | [optional] 
**Notional** | Pointer to **string** |  | [optional] 
**IsolatedWallet** | Pointer to **string** |  | [optional] 
**InitialMargin** | Pointer to **string** |  | [optional] 
**MaintMargin** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewAccountInformationV3ResponsePositionsInner

`func NewAccountInformationV3ResponsePositionsInner() *AccountInformationV3ResponsePositionsInner`

NewAccountInformationV3ResponsePositionsInner instantiates a new AccountInformationV3ResponsePositionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountInformationV3ResponsePositionsInnerWithDefaults

`func NewAccountInformationV3ResponsePositionsInnerWithDefaults() *AccountInformationV3ResponsePositionsInner`

NewAccountInformationV3ResponsePositionsInnerWithDefaults instantiates a new AccountInformationV3ResponsePositionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *AccountInformationV3ResponsePositionsInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *AccountInformationV3ResponsePositionsInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *AccountInformationV3ResponsePositionsInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *AccountInformationV3ResponsePositionsInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetPositionSide

`func (o *AccountInformationV3ResponsePositionsInner) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *AccountInformationV3ResponsePositionsInner) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *AccountInformationV3ResponsePositionsInner) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *AccountInformationV3ResponsePositionsInner) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetPositionAmt

`func (o *AccountInformationV3ResponsePositionsInner) GetPositionAmt() string`

GetPositionAmt returns the PositionAmt field if non-nil, zero value otherwise.

### GetPositionAmtOk

`func (o *AccountInformationV3ResponsePositionsInner) GetPositionAmtOk() (*string, bool)`

GetPositionAmtOk returns a tuple with the PositionAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionAmt

`func (o *AccountInformationV3ResponsePositionsInner) SetPositionAmt(v string)`

SetPositionAmt sets PositionAmt field to given value.

### HasPositionAmt

`func (o *AccountInformationV3ResponsePositionsInner) HasPositionAmt() bool`

HasPositionAmt returns a boolean if a field has been set.

### GetUnrealizedProfit

`func (o *AccountInformationV3ResponsePositionsInner) GetUnrealizedProfit() string`

GetUnrealizedProfit returns the UnrealizedProfit field if non-nil, zero value otherwise.

### GetUnrealizedProfitOk

`func (o *AccountInformationV3ResponsePositionsInner) GetUnrealizedProfitOk() (*string, bool)`

GetUnrealizedProfitOk returns a tuple with the UnrealizedProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfit

`func (o *AccountInformationV3ResponsePositionsInner) SetUnrealizedProfit(v string)`

SetUnrealizedProfit sets UnrealizedProfit field to given value.

### HasUnrealizedProfit

`func (o *AccountInformationV3ResponsePositionsInner) HasUnrealizedProfit() bool`

HasUnrealizedProfit returns a boolean if a field has been set.

### GetIsolatedMargin

`func (o *AccountInformationV3ResponsePositionsInner) GetIsolatedMargin() string`

GetIsolatedMargin returns the IsolatedMargin field if non-nil, zero value otherwise.

### GetIsolatedMarginOk

`func (o *AccountInformationV3ResponsePositionsInner) GetIsolatedMarginOk() (*string, bool)`

GetIsolatedMarginOk returns a tuple with the IsolatedMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolatedMargin

`func (o *AccountInformationV3ResponsePositionsInner) SetIsolatedMargin(v string)`

SetIsolatedMargin sets IsolatedMargin field to given value.

### HasIsolatedMargin

`func (o *AccountInformationV3ResponsePositionsInner) HasIsolatedMargin() bool`

HasIsolatedMargin returns a boolean if a field has been set.

### GetNotional

`func (o *AccountInformationV3ResponsePositionsInner) GetNotional() string`

GetNotional returns the Notional field if non-nil, zero value otherwise.

### GetNotionalOk

`func (o *AccountInformationV3ResponsePositionsInner) GetNotionalOk() (*string, bool)`

GetNotionalOk returns a tuple with the Notional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotional

`func (o *AccountInformationV3ResponsePositionsInner) SetNotional(v string)`

SetNotional sets Notional field to given value.

### HasNotional

`func (o *AccountInformationV3ResponsePositionsInner) HasNotional() bool`

HasNotional returns a boolean if a field has been set.

### GetIsolatedWallet

`func (o *AccountInformationV3ResponsePositionsInner) GetIsolatedWallet() string`

GetIsolatedWallet returns the IsolatedWallet field if non-nil, zero value otherwise.

### GetIsolatedWalletOk

`func (o *AccountInformationV3ResponsePositionsInner) GetIsolatedWalletOk() (*string, bool)`

GetIsolatedWalletOk returns a tuple with the IsolatedWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolatedWallet

`func (o *AccountInformationV3ResponsePositionsInner) SetIsolatedWallet(v string)`

SetIsolatedWallet sets IsolatedWallet field to given value.

### HasIsolatedWallet

`func (o *AccountInformationV3ResponsePositionsInner) HasIsolatedWallet() bool`

HasIsolatedWallet returns a boolean if a field has been set.

### GetInitialMargin

`func (o *AccountInformationV3ResponsePositionsInner) GetInitialMargin() string`

GetInitialMargin returns the InitialMargin field if non-nil, zero value otherwise.

### GetInitialMarginOk

`func (o *AccountInformationV3ResponsePositionsInner) GetInitialMarginOk() (*string, bool)`

GetInitialMarginOk returns a tuple with the InitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialMargin

`func (o *AccountInformationV3ResponsePositionsInner) SetInitialMargin(v string)`

SetInitialMargin sets InitialMargin field to given value.

### HasInitialMargin

`func (o *AccountInformationV3ResponsePositionsInner) HasInitialMargin() bool`

HasInitialMargin returns a boolean if a field has been set.

### GetMaintMargin

`func (o *AccountInformationV3ResponsePositionsInner) GetMaintMargin() string`

GetMaintMargin returns the MaintMargin field if non-nil, zero value otherwise.

### GetMaintMarginOk

`func (o *AccountInformationV3ResponsePositionsInner) GetMaintMarginOk() (*string, bool)`

GetMaintMarginOk returns a tuple with the MaintMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintMargin

`func (o *AccountInformationV3ResponsePositionsInner) SetMaintMargin(v string)`

SetMaintMargin sets MaintMargin field to given value.

### HasMaintMargin

`func (o *AccountInformationV3ResponsePositionsInner) HasMaintMargin() bool`

HasMaintMargin returns a boolean if a field has been set.

### GetUpdateTime

`func (o *AccountInformationV3ResponsePositionsInner) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *AccountInformationV3ResponsePositionsInner) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *AccountInformationV3ResponsePositionsInner) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *AccountInformationV3ResponsePositionsInner) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


