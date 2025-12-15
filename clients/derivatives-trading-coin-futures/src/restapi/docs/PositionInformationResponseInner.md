# PositionInformationResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**PositionAmt** | Pointer to **string** |  | [optional] 
**EntryPrice** | Pointer to **string** |  | [optional] 
**BreakEvenPrice** | Pointer to **string** |  | [optional] 
**MarkPrice** | Pointer to **string** |  | [optional] 
**UnRealizedProfit** | Pointer to **string** |  | [optional] 
**LiquidationPrice** | Pointer to **string** |  | [optional] 
**Leverage** | Pointer to **string** |  | [optional] 
**MaxQty** | Pointer to **string** |  | [optional] 
**MarginType** | Pointer to **string** |  | [optional] 
**IsolatedMargin** | Pointer to **string** |  | [optional] 
**IsAutoAddMargin** | Pointer to **string** |  | [optional] 
**PositionSide** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewPositionInformationResponseInner

`func NewPositionInformationResponseInner() *PositionInformationResponseInner`

NewPositionInformationResponseInner instantiates a new PositionInformationResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPositionInformationResponseInnerWithDefaults

`func NewPositionInformationResponseInnerWithDefaults() *PositionInformationResponseInner`

NewPositionInformationResponseInnerWithDefaults instantiates a new PositionInformationResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *PositionInformationResponseInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *PositionInformationResponseInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *PositionInformationResponseInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *PositionInformationResponseInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetPositionAmt

`func (o *PositionInformationResponseInner) GetPositionAmt() string`

GetPositionAmt returns the PositionAmt field if non-nil, zero value otherwise.

### GetPositionAmtOk

`func (o *PositionInformationResponseInner) GetPositionAmtOk() (*string, bool)`

GetPositionAmtOk returns a tuple with the PositionAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionAmt

`func (o *PositionInformationResponseInner) SetPositionAmt(v string)`

SetPositionAmt sets PositionAmt field to given value.

### HasPositionAmt

`func (o *PositionInformationResponseInner) HasPositionAmt() bool`

HasPositionAmt returns a boolean if a field has been set.

### GetEntryPrice

`func (o *PositionInformationResponseInner) GetEntryPrice() string`

GetEntryPrice returns the EntryPrice field if non-nil, zero value otherwise.

### GetEntryPriceOk

`func (o *PositionInformationResponseInner) GetEntryPriceOk() (*string, bool)`

GetEntryPriceOk returns a tuple with the EntryPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryPrice

`func (o *PositionInformationResponseInner) SetEntryPrice(v string)`

SetEntryPrice sets EntryPrice field to given value.

### HasEntryPrice

`func (o *PositionInformationResponseInner) HasEntryPrice() bool`

HasEntryPrice returns a boolean if a field has been set.

### GetBreakEvenPrice

`func (o *PositionInformationResponseInner) GetBreakEvenPrice() string`

GetBreakEvenPrice returns the BreakEvenPrice field if non-nil, zero value otherwise.

### GetBreakEvenPriceOk

`func (o *PositionInformationResponseInner) GetBreakEvenPriceOk() (*string, bool)`

GetBreakEvenPriceOk returns a tuple with the BreakEvenPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakEvenPrice

`func (o *PositionInformationResponseInner) SetBreakEvenPrice(v string)`

SetBreakEvenPrice sets BreakEvenPrice field to given value.

### HasBreakEvenPrice

`func (o *PositionInformationResponseInner) HasBreakEvenPrice() bool`

HasBreakEvenPrice returns a boolean if a field has been set.

### GetMarkPrice

`func (o *PositionInformationResponseInner) GetMarkPrice() string`

GetMarkPrice returns the MarkPrice field if non-nil, zero value otherwise.

### GetMarkPriceOk

`func (o *PositionInformationResponseInner) GetMarkPriceOk() (*string, bool)`

GetMarkPriceOk returns a tuple with the MarkPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkPrice

`func (o *PositionInformationResponseInner) SetMarkPrice(v string)`

SetMarkPrice sets MarkPrice field to given value.

### HasMarkPrice

`func (o *PositionInformationResponseInner) HasMarkPrice() bool`

HasMarkPrice returns a boolean if a field has been set.

### GetUnRealizedProfit

`func (o *PositionInformationResponseInner) GetUnRealizedProfit() string`

GetUnRealizedProfit returns the UnRealizedProfit field if non-nil, zero value otherwise.

### GetUnRealizedProfitOk

`func (o *PositionInformationResponseInner) GetUnRealizedProfitOk() (*string, bool)`

GetUnRealizedProfitOk returns a tuple with the UnRealizedProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnRealizedProfit

`func (o *PositionInformationResponseInner) SetUnRealizedProfit(v string)`

SetUnRealizedProfit sets UnRealizedProfit field to given value.

### HasUnRealizedProfit

`func (o *PositionInformationResponseInner) HasUnRealizedProfit() bool`

HasUnRealizedProfit returns a boolean if a field has been set.

### GetLiquidationPrice

`func (o *PositionInformationResponseInner) GetLiquidationPrice() string`

GetLiquidationPrice returns the LiquidationPrice field if non-nil, zero value otherwise.

### GetLiquidationPriceOk

`func (o *PositionInformationResponseInner) GetLiquidationPriceOk() (*string, bool)`

GetLiquidationPriceOk returns a tuple with the LiquidationPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidationPrice

`func (o *PositionInformationResponseInner) SetLiquidationPrice(v string)`

SetLiquidationPrice sets LiquidationPrice field to given value.

### HasLiquidationPrice

`func (o *PositionInformationResponseInner) HasLiquidationPrice() bool`

HasLiquidationPrice returns a boolean if a field has been set.

### GetLeverage

`func (o *PositionInformationResponseInner) GetLeverage() string`

GetLeverage returns the Leverage field if non-nil, zero value otherwise.

### GetLeverageOk

`func (o *PositionInformationResponseInner) GetLeverageOk() (*string, bool)`

GetLeverageOk returns a tuple with the Leverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverage

`func (o *PositionInformationResponseInner) SetLeverage(v string)`

SetLeverage sets Leverage field to given value.

### HasLeverage

`func (o *PositionInformationResponseInner) HasLeverage() bool`

HasLeverage returns a boolean if a field has been set.

### GetMaxQty

`func (o *PositionInformationResponseInner) GetMaxQty() string`

GetMaxQty returns the MaxQty field if non-nil, zero value otherwise.

### GetMaxQtyOk

`func (o *PositionInformationResponseInner) GetMaxQtyOk() (*string, bool)`

GetMaxQtyOk returns a tuple with the MaxQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQty

`func (o *PositionInformationResponseInner) SetMaxQty(v string)`

SetMaxQty sets MaxQty field to given value.

### HasMaxQty

`func (o *PositionInformationResponseInner) HasMaxQty() bool`

HasMaxQty returns a boolean if a field has been set.

### GetMarginType

`func (o *PositionInformationResponseInner) GetMarginType() string`

GetMarginType returns the MarginType field if non-nil, zero value otherwise.

### GetMarginTypeOk

`func (o *PositionInformationResponseInner) GetMarginTypeOk() (*string, bool)`

GetMarginTypeOk returns a tuple with the MarginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginType

`func (o *PositionInformationResponseInner) SetMarginType(v string)`

SetMarginType sets MarginType field to given value.

### HasMarginType

`func (o *PositionInformationResponseInner) HasMarginType() bool`

HasMarginType returns a boolean if a field has been set.

### GetIsolatedMargin

`func (o *PositionInformationResponseInner) GetIsolatedMargin() string`

GetIsolatedMargin returns the IsolatedMargin field if non-nil, zero value otherwise.

### GetIsolatedMarginOk

`func (o *PositionInformationResponseInner) GetIsolatedMarginOk() (*string, bool)`

GetIsolatedMarginOk returns a tuple with the IsolatedMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolatedMargin

`func (o *PositionInformationResponseInner) SetIsolatedMargin(v string)`

SetIsolatedMargin sets IsolatedMargin field to given value.

### HasIsolatedMargin

`func (o *PositionInformationResponseInner) HasIsolatedMargin() bool`

HasIsolatedMargin returns a boolean if a field has been set.

### GetIsAutoAddMargin

`func (o *PositionInformationResponseInner) GetIsAutoAddMargin() string`

GetIsAutoAddMargin returns the IsAutoAddMargin field if non-nil, zero value otherwise.

### GetIsAutoAddMarginOk

`func (o *PositionInformationResponseInner) GetIsAutoAddMarginOk() (*string, bool)`

GetIsAutoAddMarginOk returns a tuple with the IsAutoAddMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoAddMargin

`func (o *PositionInformationResponseInner) SetIsAutoAddMargin(v string)`

SetIsAutoAddMargin sets IsAutoAddMargin field to given value.

### HasIsAutoAddMargin

`func (o *PositionInformationResponseInner) HasIsAutoAddMargin() bool`

HasIsAutoAddMargin returns a boolean if a field has been set.

### GetPositionSide

`func (o *PositionInformationResponseInner) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *PositionInformationResponseInner) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *PositionInformationResponseInner) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *PositionInformationResponseInner) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetUpdateTime

`func (o *PositionInformationResponseInner) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *PositionInformationResponseInner) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *PositionInformationResponseInner) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *PositionInformationResponseInner) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


