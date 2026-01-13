# PositionInformationV3ResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**PositionSide** | Pointer to **string** |  | [optional] 
**PositionAmt** | Pointer to **string** |  | [optional] 
**EntryPrice** | Pointer to **string** |  | [optional] 
**BreakEvenPrice** | Pointer to **string** |  | [optional] 
**MarkPrice** | Pointer to **string** |  | [optional] 
**UnRealizedProfit** | Pointer to **string** |  | [optional] 
**LiquidationPrice** | Pointer to **string** |  | [optional] 
**IsolatedMargin** | Pointer to **string** |  | [optional] 
**Notional** | Pointer to **string** |  | [optional] 
**MarginAsset** | Pointer to **string** |  | [optional] 
**IsolatedWallet** | Pointer to **string** |  | [optional] 
**InitialMargin** | Pointer to **string** |  | [optional] 
**MaintMargin** | Pointer to **string** |  | [optional] 
**PositionInitialMargin** | Pointer to **string** |  | [optional] 
**OpenOrderInitialMargin** | Pointer to **string** |  | [optional] 
**Adl** | Pointer to **int64** |  | [optional] 
**BidNotional** | Pointer to **string** |  | [optional] 
**AskNotional** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewPositionInformationV3ResponseInner

`func NewPositionInformationV3ResponseInner() *PositionInformationV3ResponseInner`

NewPositionInformationV3ResponseInner instantiates a new PositionInformationV3ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPositionInformationV3ResponseInnerWithDefaults

`func NewPositionInformationV3ResponseInnerWithDefaults() *PositionInformationV3ResponseInner`

NewPositionInformationV3ResponseInnerWithDefaults instantiates a new PositionInformationV3ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *PositionInformationV3ResponseInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *PositionInformationV3ResponseInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *PositionInformationV3ResponseInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *PositionInformationV3ResponseInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetPositionSide

`func (o *PositionInformationV3ResponseInner) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *PositionInformationV3ResponseInner) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *PositionInformationV3ResponseInner) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *PositionInformationV3ResponseInner) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetPositionAmt

`func (o *PositionInformationV3ResponseInner) GetPositionAmt() string`

GetPositionAmt returns the PositionAmt field if non-nil, zero value otherwise.

### GetPositionAmtOk

`func (o *PositionInformationV3ResponseInner) GetPositionAmtOk() (*string, bool)`

GetPositionAmtOk returns a tuple with the PositionAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionAmt

`func (o *PositionInformationV3ResponseInner) SetPositionAmt(v string)`

SetPositionAmt sets PositionAmt field to given value.

### HasPositionAmt

`func (o *PositionInformationV3ResponseInner) HasPositionAmt() bool`

HasPositionAmt returns a boolean if a field has been set.

### GetEntryPrice

`func (o *PositionInformationV3ResponseInner) GetEntryPrice() string`

GetEntryPrice returns the EntryPrice field if non-nil, zero value otherwise.

### GetEntryPriceOk

`func (o *PositionInformationV3ResponseInner) GetEntryPriceOk() (*string, bool)`

GetEntryPriceOk returns a tuple with the EntryPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryPrice

`func (o *PositionInformationV3ResponseInner) SetEntryPrice(v string)`

SetEntryPrice sets EntryPrice field to given value.

### HasEntryPrice

`func (o *PositionInformationV3ResponseInner) HasEntryPrice() bool`

HasEntryPrice returns a boolean if a field has been set.

### GetBreakEvenPrice

`func (o *PositionInformationV3ResponseInner) GetBreakEvenPrice() string`

GetBreakEvenPrice returns the BreakEvenPrice field if non-nil, zero value otherwise.

### GetBreakEvenPriceOk

`func (o *PositionInformationV3ResponseInner) GetBreakEvenPriceOk() (*string, bool)`

GetBreakEvenPriceOk returns a tuple with the BreakEvenPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakEvenPrice

`func (o *PositionInformationV3ResponseInner) SetBreakEvenPrice(v string)`

SetBreakEvenPrice sets BreakEvenPrice field to given value.

### HasBreakEvenPrice

`func (o *PositionInformationV3ResponseInner) HasBreakEvenPrice() bool`

HasBreakEvenPrice returns a boolean if a field has been set.

### GetMarkPrice

`func (o *PositionInformationV3ResponseInner) GetMarkPrice() string`

GetMarkPrice returns the MarkPrice field if non-nil, zero value otherwise.

### GetMarkPriceOk

`func (o *PositionInformationV3ResponseInner) GetMarkPriceOk() (*string, bool)`

GetMarkPriceOk returns a tuple with the MarkPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkPrice

`func (o *PositionInformationV3ResponseInner) SetMarkPrice(v string)`

SetMarkPrice sets MarkPrice field to given value.

### HasMarkPrice

`func (o *PositionInformationV3ResponseInner) HasMarkPrice() bool`

HasMarkPrice returns a boolean if a field has been set.

### GetUnRealizedProfit

`func (o *PositionInformationV3ResponseInner) GetUnRealizedProfit() string`

GetUnRealizedProfit returns the UnRealizedProfit field if non-nil, zero value otherwise.

### GetUnRealizedProfitOk

`func (o *PositionInformationV3ResponseInner) GetUnRealizedProfitOk() (*string, bool)`

GetUnRealizedProfitOk returns a tuple with the UnRealizedProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnRealizedProfit

`func (o *PositionInformationV3ResponseInner) SetUnRealizedProfit(v string)`

SetUnRealizedProfit sets UnRealizedProfit field to given value.

### HasUnRealizedProfit

`func (o *PositionInformationV3ResponseInner) HasUnRealizedProfit() bool`

HasUnRealizedProfit returns a boolean if a field has been set.

### GetLiquidationPrice

`func (o *PositionInformationV3ResponseInner) GetLiquidationPrice() string`

GetLiquidationPrice returns the LiquidationPrice field if non-nil, zero value otherwise.

### GetLiquidationPriceOk

`func (o *PositionInformationV3ResponseInner) GetLiquidationPriceOk() (*string, bool)`

GetLiquidationPriceOk returns a tuple with the LiquidationPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidationPrice

`func (o *PositionInformationV3ResponseInner) SetLiquidationPrice(v string)`

SetLiquidationPrice sets LiquidationPrice field to given value.

### HasLiquidationPrice

`func (o *PositionInformationV3ResponseInner) HasLiquidationPrice() bool`

HasLiquidationPrice returns a boolean if a field has been set.

### GetIsolatedMargin

`func (o *PositionInformationV3ResponseInner) GetIsolatedMargin() string`

GetIsolatedMargin returns the IsolatedMargin field if non-nil, zero value otherwise.

### GetIsolatedMarginOk

`func (o *PositionInformationV3ResponseInner) GetIsolatedMarginOk() (*string, bool)`

GetIsolatedMarginOk returns a tuple with the IsolatedMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolatedMargin

`func (o *PositionInformationV3ResponseInner) SetIsolatedMargin(v string)`

SetIsolatedMargin sets IsolatedMargin field to given value.

### HasIsolatedMargin

`func (o *PositionInformationV3ResponseInner) HasIsolatedMargin() bool`

HasIsolatedMargin returns a boolean if a field has been set.

### GetNotional

`func (o *PositionInformationV3ResponseInner) GetNotional() string`

GetNotional returns the Notional field if non-nil, zero value otherwise.

### GetNotionalOk

`func (o *PositionInformationV3ResponseInner) GetNotionalOk() (*string, bool)`

GetNotionalOk returns a tuple with the Notional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotional

`func (o *PositionInformationV3ResponseInner) SetNotional(v string)`

SetNotional sets Notional field to given value.

### HasNotional

`func (o *PositionInformationV3ResponseInner) HasNotional() bool`

HasNotional returns a boolean if a field has been set.

### GetMarginAsset

`func (o *PositionInformationV3ResponseInner) GetMarginAsset() string`

GetMarginAsset returns the MarginAsset field if non-nil, zero value otherwise.

### GetMarginAssetOk

`func (o *PositionInformationV3ResponseInner) GetMarginAssetOk() (*string, bool)`

GetMarginAssetOk returns a tuple with the MarginAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginAsset

`func (o *PositionInformationV3ResponseInner) SetMarginAsset(v string)`

SetMarginAsset sets MarginAsset field to given value.

### HasMarginAsset

`func (o *PositionInformationV3ResponseInner) HasMarginAsset() bool`

HasMarginAsset returns a boolean if a field has been set.

### GetIsolatedWallet

`func (o *PositionInformationV3ResponseInner) GetIsolatedWallet() string`

GetIsolatedWallet returns the IsolatedWallet field if non-nil, zero value otherwise.

### GetIsolatedWalletOk

`func (o *PositionInformationV3ResponseInner) GetIsolatedWalletOk() (*string, bool)`

GetIsolatedWalletOk returns a tuple with the IsolatedWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolatedWallet

`func (o *PositionInformationV3ResponseInner) SetIsolatedWallet(v string)`

SetIsolatedWallet sets IsolatedWallet field to given value.

### HasIsolatedWallet

`func (o *PositionInformationV3ResponseInner) HasIsolatedWallet() bool`

HasIsolatedWallet returns a boolean if a field has been set.

### GetInitialMargin

`func (o *PositionInformationV3ResponseInner) GetInitialMargin() string`

GetInitialMargin returns the InitialMargin field if non-nil, zero value otherwise.

### GetInitialMarginOk

`func (o *PositionInformationV3ResponseInner) GetInitialMarginOk() (*string, bool)`

GetInitialMarginOk returns a tuple with the InitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialMargin

`func (o *PositionInformationV3ResponseInner) SetInitialMargin(v string)`

SetInitialMargin sets InitialMargin field to given value.

### HasInitialMargin

`func (o *PositionInformationV3ResponseInner) HasInitialMargin() bool`

HasInitialMargin returns a boolean if a field has been set.

### GetMaintMargin

`func (o *PositionInformationV3ResponseInner) GetMaintMargin() string`

GetMaintMargin returns the MaintMargin field if non-nil, zero value otherwise.

### GetMaintMarginOk

`func (o *PositionInformationV3ResponseInner) GetMaintMarginOk() (*string, bool)`

GetMaintMarginOk returns a tuple with the MaintMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintMargin

`func (o *PositionInformationV3ResponseInner) SetMaintMargin(v string)`

SetMaintMargin sets MaintMargin field to given value.

### HasMaintMargin

`func (o *PositionInformationV3ResponseInner) HasMaintMargin() bool`

HasMaintMargin returns a boolean if a field has been set.

### GetPositionInitialMargin

`func (o *PositionInformationV3ResponseInner) GetPositionInitialMargin() string`

GetPositionInitialMargin returns the PositionInitialMargin field if non-nil, zero value otherwise.

### GetPositionInitialMarginOk

`func (o *PositionInformationV3ResponseInner) GetPositionInitialMarginOk() (*string, bool)`

GetPositionInitialMarginOk returns a tuple with the PositionInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionInitialMargin

`func (o *PositionInformationV3ResponseInner) SetPositionInitialMargin(v string)`

SetPositionInitialMargin sets PositionInitialMargin field to given value.

### HasPositionInitialMargin

`func (o *PositionInformationV3ResponseInner) HasPositionInitialMargin() bool`

HasPositionInitialMargin returns a boolean if a field has been set.

### GetOpenOrderInitialMargin

`func (o *PositionInformationV3ResponseInner) GetOpenOrderInitialMargin() string`

GetOpenOrderInitialMargin returns the OpenOrderInitialMargin field if non-nil, zero value otherwise.

### GetOpenOrderInitialMarginOk

`func (o *PositionInformationV3ResponseInner) GetOpenOrderInitialMarginOk() (*string, bool)`

GetOpenOrderInitialMarginOk returns a tuple with the OpenOrderInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenOrderInitialMargin

`func (o *PositionInformationV3ResponseInner) SetOpenOrderInitialMargin(v string)`

SetOpenOrderInitialMargin sets OpenOrderInitialMargin field to given value.

### HasOpenOrderInitialMargin

`func (o *PositionInformationV3ResponseInner) HasOpenOrderInitialMargin() bool`

HasOpenOrderInitialMargin returns a boolean if a field has been set.

### GetAdl

`func (o *PositionInformationV3ResponseInner) GetAdl() int64`

GetAdl returns the Adl field if non-nil, zero value otherwise.

### GetAdlOk

`func (o *PositionInformationV3ResponseInner) GetAdlOk() (*int64, bool)`

GetAdlOk returns a tuple with the Adl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdl

`func (o *PositionInformationV3ResponseInner) SetAdl(v int64)`

SetAdl sets Adl field to given value.

### HasAdl

`func (o *PositionInformationV3ResponseInner) HasAdl() bool`

HasAdl returns a boolean if a field has been set.

### GetBidNotional

`func (o *PositionInformationV3ResponseInner) GetBidNotional() string`

GetBidNotional returns the BidNotional field if non-nil, zero value otherwise.

### GetBidNotionalOk

`func (o *PositionInformationV3ResponseInner) GetBidNotionalOk() (*string, bool)`

GetBidNotionalOk returns a tuple with the BidNotional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidNotional

`func (o *PositionInformationV3ResponseInner) SetBidNotional(v string)`

SetBidNotional sets BidNotional field to given value.

### HasBidNotional

`func (o *PositionInformationV3ResponseInner) HasBidNotional() bool`

HasBidNotional returns a boolean if a field has been set.

### GetAskNotional

`func (o *PositionInformationV3ResponseInner) GetAskNotional() string`

GetAskNotional returns the AskNotional field if non-nil, zero value otherwise.

### GetAskNotionalOk

`func (o *PositionInformationV3ResponseInner) GetAskNotionalOk() (*string, bool)`

GetAskNotionalOk returns a tuple with the AskNotional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskNotional

`func (o *PositionInformationV3ResponseInner) SetAskNotional(v string)`

SetAskNotional sets AskNotional field to given value.

### HasAskNotional

`func (o *PositionInformationV3ResponseInner) HasAskNotional() bool`

HasAskNotional returns a boolean if a field has been set.

### GetUpdateTime

`func (o *PositionInformationV3ResponseInner) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *PositionInformationV3ResponseInner) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *PositionInformationV3ResponseInner) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *PositionInformationV3ResponseInner) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


