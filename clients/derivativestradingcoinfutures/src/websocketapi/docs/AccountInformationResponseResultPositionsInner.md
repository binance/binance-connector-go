# AccountInformationResponseResultPositionsInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**InitialMargin** | Pointer to **string** |  | [optional] 
**MaintMargin** | Pointer to **string** |  | [optional] 
**UnrealizedProfit** | Pointer to **string** |  | [optional] 
**PositionInitialMargin** | Pointer to **string** |  | [optional] 
**OpenOrderInitialMargin** | Pointer to **string** |  | [optional] 
**Leverage** | Pointer to **string** |  | [optional] 
**Isolated** | Pointer to **bool** |  | [optional] 
**PositionSide** | Pointer to **string** |  | [optional] 
**EntryPrice** | Pointer to **string** |  | [optional] 
**MaxQty** | Pointer to **string** |  | [optional] 
**NotionalValue** | Pointer to **string** |  | [optional] 
**IsolatedWallet** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**PositionAmt** | Pointer to **string** |  | [optional] 
**BreakEvenPrice** | Pointer to **string** |  | [optional] 

## Methods

### NewAccountInformationResponseResultPositionsInner

`func NewAccountInformationResponseResultPositionsInner() *AccountInformationResponseResultPositionsInner`

NewAccountInformationResponseResultPositionsInner instantiates a new AccountInformationResponseResultPositionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountInformationResponseResultPositionsInnerWithDefaults

`func NewAccountInformationResponseResultPositionsInnerWithDefaults() *AccountInformationResponseResultPositionsInner`

NewAccountInformationResponseResultPositionsInnerWithDefaults instantiates a new AccountInformationResponseResultPositionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *AccountInformationResponseResultPositionsInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *AccountInformationResponseResultPositionsInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *AccountInformationResponseResultPositionsInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *AccountInformationResponseResultPositionsInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetInitialMargin

`func (o *AccountInformationResponseResultPositionsInner) GetInitialMargin() string`

GetInitialMargin returns the InitialMargin field if non-nil, zero value otherwise.

### GetInitialMarginOk

`func (o *AccountInformationResponseResultPositionsInner) GetInitialMarginOk() (*string, bool)`

GetInitialMarginOk returns a tuple with the InitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialMargin

`func (o *AccountInformationResponseResultPositionsInner) SetInitialMargin(v string)`

SetInitialMargin sets InitialMargin field to given value.

### HasInitialMargin

`func (o *AccountInformationResponseResultPositionsInner) HasInitialMargin() bool`

HasInitialMargin returns a boolean if a field has been set.

### GetMaintMargin

`func (o *AccountInformationResponseResultPositionsInner) GetMaintMargin() string`

GetMaintMargin returns the MaintMargin field if non-nil, zero value otherwise.

### GetMaintMarginOk

`func (o *AccountInformationResponseResultPositionsInner) GetMaintMarginOk() (*string, bool)`

GetMaintMarginOk returns a tuple with the MaintMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintMargin

`func (o *AccountInformationResponseResultPositionsInner) SetMaintMargin(v string)`

SetMaintMargin sets MaintMargin field to given value.

### HasMaintMargin

`func (o *AccountInformationResponseResultPositionsInner) HasMaintMargin() bool`

HasMaintMargin returns a boolean if a field has been set.

### GetUnrealizedProfit

`func (o *AccountInformationResponseResultPositionsInner) GetUnrealizedProfit() string`

GetUnrealizedProfit returns the UnrealizedProfit field if non-nil, zero value otherwise.

### GetUnrealizedProfitOk

`func (o *AccountInformationResponseResultPositionsInner) GetUnrealizedProfitOk() (*string, bool)`

GetUnrealizedProfitOk returns a tuple with the UnrealizedProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfit

`func (o *AccountInformationResponseResultPositionsInner) SetUnrealizedProfit(v string)`

SetUnrealizedProfit sets UnrealizedProfit field to given value.

### HasUnrealizedProfit

`func (o *AccountInformationResponseResultPositionsInner) HasUnrealizedProfit() bool`

HasUnrealizedProfit returns a boolean if a field has been set.

### GetPositionInitialMargin

`func (o *AccountInformationResponseResultPositionsInner) GetPositionInitialMargin() string`

GetPositionInitialMargin returns the PositionInitialMargin field if non-nil, zero value otherwise.

### GetPositionInitialMarginOk

`func (o *AccountInformationResponseResultPositionsInner) GetPositionInitialMarginOk() (*string, bool)`

GetPositionInitialMarginOk returns a tuple with the PositionInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionInitialMargin

`func (o *AccountInformationResponseResultPositionsInner) SetPositionInitialMargin(v string)`

SetPositionInitialMargin sets PositionInitialMargin field to given value.

### HasPositionInitialMargin

`func (o *AccountInformationResponseResultPositionsInner) HasPositionInitialMargin() bool`

HasPositionInitialMargin returns a boolean if a field has been set.

### GetOpenOrderInitialMargin

`func (o *AccountInformationResponseResultPositionsInner) GetOpenOrderInitialMargin() string`

GetOpenOrderInitialMargin returns the OpenOrderInitialMargin field if non-nil, zero value otherwise.

### GetOpenOrderInitialMarginOk

`func (o *AccountInformationResponseResultPositionsInner) GetOpenOrderInitialMarginOk() (*string, bool)`

GetOpenOrderInitialMarginOk returns a tuple with the OpenOrderInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenOrderInitialMargin

`func (o *AccountInformationResponseResultPositionsInner) SetOpenOrderInitialMargin(v string)`

SetOpenOrderInitialMargin sets OpenOrderInitialMargin field to given value.

### HasOpenOrderInitialMargin

`func (o *AccountInformationResponseResultPositionsInner) HasOpenOrderInitialMargin() bool`

HasOpenOrderInitialMargin returns a boolean if a field has been set.

### GetLeverage

`func (o *AccountInformationResponseResultPositionsInner) GetLeverage() string`

GetLeverage returns the Leverage field if non-nil, zero value otherwise.

### GetLeverageOk

`func (o *AccountInformationResponseResultPositionsInner) GetLeverageOk() (*string, bool)`

GetLeverageOk returns a tuple with the Leverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverage

`func (o *AccountInformationResponseResultPositionsInner) SetLeverage(v string)`

SetLeverage sets Leverage field to given value.

### HasLeverage

`func (o *AccountInformationResponseResultPositionsInner) HasLeverage() bool`

HasLeverage returns a boolean if a field has been set.

### GetIsolated

`func (o *AccountInformationResponseResultPositionsInner) GetIsolated() bool`

GetIsolated returns the Isolated field if non-nil, zero value otherwise.

### GetIsolatedOk

`func (o *AccountInformationResponseResultPositionsInner) GetIsolatedOk() (*bool, bool)`

GetIsolatedOk returns a tuple with the Isolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolated

`func (o *AccountInformationResponseResultPositionsInner) SetIsolated(v bool)`

SetIsolated sets Isolated field to given value.

### HasIsolated

`func (o *AccountInformationResponseResultPositionsInner) HasIsolated() bool`

HasIsolated returns a boolean if a field has been set.

### GetPositionSide

`func (o *AccountInformationResponseResultPositionsInner) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *AccountInformationResponseResultPositionsInner) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *AccountInformationResponseResultPositionsInner) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *AccountInformationResponseResultPositionsInner) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetEntryPrice

`func (o *AccountInformationResponseResultPositionsInner) GetEntryPrice() string`

GetEntryPrice returns the EntryPrice field if non-nil, zero value otherwise.

### GetEntryPriceOk

`func (o *AccountInformationResponseResultPositionsInner) GetEntryPriceOk() (*string, bool)`

GetEntryPriceOk returns a tuple with the EntryPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryPrice

`func (o *AccountInformationResponseResultPositionsInner) SetEntryPrice(v string)`

SetEntryPrice sets EntryPrice field to given value.

### HasEntryPrice

`func (o *AccountInformationResponseResultPositionsInner) HasEntryPrice() bool`

HasEntryPrice returns a boolean if a field has been set.

### GetMaxQty

`func (o *AccountInformationResponseResultPositionsInner) GetMaxQty() string`

GetMaxQty returns the MaxQty field if non-nil, zero value otherwise.

### GetMaxQtyOk

`func (o *AccountInformationResponseResultPositionsInner) GetMaxQtyOk() (*string, bool)`

GetMaxQtyOk returns a tuple with the MaxQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQty

`func (o *AccountInformationResponseResultPositionsInner) SetMaxQty(v string)`

SetMaxQty sets MaxQty field to given value.

### HasMaxQty

`func (o *AccountInformationResponseResultPositionsInner) HasMaxQty() bool`

HasMaxQty returns a boolean if a field has been set.

### GetNotionalValue

`func (o *AccountInformationResponseResultPositionsInner) GetNotionalValue() string`

GetNotionalValue returns the NotionalValue field if non-nil, zero value otherwise.

### GetNotionalValueOk

`func (o *AccountInformationResponseResultPositionsInner) GetNotionalValueOk() (*string, bool)`

GetNotionalValueOk returns a tuple with the NotionalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalValue

`func (o *AccountInformationResponseResultPositionsInner) SetNotionalValue(v string)`

SetNotionalValue sets NotionalValue field to given value.

### HasNotionalValue

`func (o *AccountInformationResponseResultPositionsInner) HasNotionalValue() bool`

HasNotionalValue returns a boolean if a field has been set.

### GetIsolatedWallet

`func (o *AccountInformationResponseResultPositionsInner) GetIsolatedWallet() string`

GetIsolatedWallet returns the IsolatedWallet field if non-nil, zero value otherwise.

### GetIsolatedWalletOk

`func (o *AccountInformationResponseResultPositionsInner) GetIsolatedWalletOk() (*string, bool)`

GetIsolatedWalletOk returns a tuple with the IsolatedWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolatedWallet

`func (o *AccountInformationResponseResultPositionsInner) SetIsolatedWallet(v string)`

SetIsolatedWallet sets IsolatedWallet field to given value.

### HasIsolatedWallet

`func (o *AccountInformationResponseResultPositionsInner) HasIsolatedWallet() bool`

HasIsolatedWallet returns a boolean if a field has been set.

### GetUpdateTime

`func (o *AccountInformationResponseResultPositionsInner) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *AccountInformationResponseResultPositionsInner) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *AccountInformationResponseResultPositionsInner) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *AccountInformationResponseResultPositionsInner) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetPositionAmt

`func (o *AccountInformationResponseResultPositionsInner) GetPositionAmt() string`

GetPositionAmt returns the PositionAmt field if non-nil, zero value otherwise.

### GetPositionAmtOk

`func (o *AccountInformationResponseResultPositionsInner) GetPositionAmtOk() (*string, bool)`

GetPositionAmtOk returns a tuple with the PositionAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionAmt

`func (o *AccountInformationResponseResultPositionsInner) SetPositionAmt(v string)`

SetPositionAmt sets PositionAmt field to given value.

### HasPositionAmt

`func (o *AccountInformationResponseResultPositionsInner) HasPositionAmt() bool`

HasPositionAmt returns a boolean if a field has been set.

### GetBreakEvenPrice

`func (o *AccountInformationResponseResultPositionsInner) GetBreakEvenPrice() string`

GetBreakEvenPrice returns the BreakEvenPrice field if non-nil, zero value otherwise.

### GetBreakEvenPriceOk

`func (o *AccountInformationResponseResultPositionsInner) GetBreakEvenPriceOk() (*string, bool)`

GetBreakEvenPriceOk returns a tuple with the BreakEvenPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakEvenPrice

`func (o *AccountInformationResponseResultPositionsInner) SetBreakEvenPrice(v string)`

SetBreakEvenPrice sets BreakEvenPrice field to given value.

### HasBreakEvenPrice

`func (o *AccountInformationResponseResultPositionsInner) HasBreakEvenPrice() bool`

HasBreakEvenPrice returns a boolean if a field has been set.


[[Back to README]](../README.md)


