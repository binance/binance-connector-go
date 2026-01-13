# OptionPositionInformationResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**EntryPrice** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **string** |  | [optional] 
**ReducibleQty** | Pointer to **string** |  | [optional] 
**MarkValue** | Pointer to **string** |  | [optional] 
**Ror** | Pointer to **string** |  | [optional] 
**UnrealizedPNL** | Pointer to **string** |  | [optional] 
**MarkPrice** | Pointer to **string** |  | [optional] 
**StrikePrice** | Pointer to **string** |  | [optional] 
**PositionCost** | Pointer to **string** |  | [optional] 
**ExpiryDate** | Pointer to **int64** |  | [optional] 
**PriceScale** | Pointer to **int64** |  | [optional] 
**QuantityScale** | Pointer to **int64** |  | [optional] 
**OptionSide** | Pointer to **string** |  | [optional] 
**QuoteAsset** | Pointer to **string** |  | [optional] 

## Methods

### NewOptionPositionInformationResponseInner

`func NewOptionPositionInformationResponseInner() *OptionPositionInformationResponseInner`

NewOptionPositionInformationResponseInner instantiates a new OptionPositionInformationResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionPositionInformationResponseInnerWithDefaults

`func NewOptionPositionInformationResponseInnerWithDefaults() *OptionPositionInformationResponseInner`

NewOptionPositionInformationResponseInnerWithDefaults instantiates a new OptionPositionInformationResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryPrice

`func (o *OptionPositionInformationResponseInner) GetEntryPrice() string`

GetEntryPrice returns the EntryPrice field if non-nil, zero value otherwise.

### GetEntryPriceOk

`func (o *OptionPositionInformationResponseInner) GetEntryPriceOk() (*string, bool)`

GetEntryPriceOk returns a tuple with the EntryPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryPrice

`func (o *OptionPositionInformationResponseInner) SetEntryPrice(v string)`

SetEntryPrice sets EntryPrice field to given value.

### HasEntryPrice

`func (o *OptionPositionInformationResponseInner) HasEntryPrice() bool`

HasEntryPrice returns a boolean if a field has been set.

### GetSymbol

`func (o *OptionPositionInformationResponseInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OptionPositionInformationResponseInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OptionPositionInformationResponseInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *OptionPositionInformationResponseInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetSide

`func (o *OptionPositionInformationResponseInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *OptionPositionInformationResponseInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *OptionPositionInformationResponseInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *OptionPositionInformationResponseInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetQuantity

`func (o *OptionPositionInformationResponseInner) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OptionPositionInformationResponseInner) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OptionPositionInformationResponseInner) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OptionPositionInformationResponseInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetReducibleQty

`func (o *OptionPositionInformationResponseInner) GetReducibleQty() string`

GetReducibleQty returns the ReducibleQty field if non-nil, zero value otherwise.

### GetReducibleQtyOk

`func (o *OptionPositionInformationResponseInner) GetReducibleQtyOk() (*string, bool)`

GetReducibleQtyOk returns a tuple with the ReducibleQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReducibleQty

`func (o *OptionPositionInformationResponseInner) SetReducibleQty(v string)`

SetReducibleQty sets ReducibleQty field to given value.

### HasReducibleQty

`func (o *OptionPositionInformationResponseInner) HasReducibleQty() bool`

HasReducibleQty returns a boolean if a field has been set.

### GetMarkValue

`func (o *OptionPositionInformationResponseInner) GetMarkValue() string`

GetMarkValue returns the MarkValue field if non-nil, zero value otherwise.

### GetMarkValueOk

`func (o *OptionPositionInformationResponseInner) GetMarkValueOk() (*string, bool)`

GetMarkValueOk returns a tuple with the MarkValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkValue

`func (o *OptionPositionInformationResponseInner) SetMarkValue(v string)`

SetMarkValue sets MarkValue field to given value.

### HasMarkValue

`func (o *OptionPositionInformationResponseInner) HasMarkValue() bool`

HasMarkValue returns a boolean if a field has been set.

### GetRor

`func (o *OptionPositionInformationResponseInner) GetRor() string`

GetRor returns the Ror field if non-nil, zero value otherwise.

### GetRorOk

`func (o *OptionPositionInformationResponseInner) GetRorOk() (*string, bool)`

GetRorOk returns a tuple with the Ror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRor

`func (o *OptionPositionInformationResponseInner) SetRor(v string)`

SetRor sets Ror field to given value.

### HasRor

`func (o *OptionPositionInformationResponseInner) HasRor() bool`

HasRor returns a boolean if a field has been set.

### GetUnrealizedPNL

`func (o *OptionPositionInformationResponseInner) GetUnrealizedPNL() string`

GetUnrealizedPNL returns the UnrealizedPNL field if non-nil, zero value otherwise.

### GetUnrealizedPNLOk

`func (o *OptionPositionInformationResponseInner) GetUnrealizedPNLOk() (*string, bool)`

GetUnrealizedPNLOk returns a tuple with the UnrealizedPNL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedPNL

`func (o *OptionPositionInformationResponseInner) SetUnrealizedPNL(v string)`

SetUnrealizedPNL sets UnrealizedPNL field to given value.

### HasUnrealizedPNL

`func (o *OptionPositionInformationResponseInner) HasUnrealizedPNL() bool`

HasUnrealizedPNL returns a boolean if a field has been set.

### GetMarkPrice

`func (o *OptionPositionInformationResponseInner) GetMarkPrice() string`

GetMarkPrice returns the MarkPrice field if non-nil, zero value otherwise.

### GetMarkPriceOk

`func (o *OptionPositionInformationResponseInner) GetMarkPriceOk() (*string, bool)`

GetMarkPriceOk returns a tuple with the MarkPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkPrice

`func (o *OptionPositionInformationResponseInner) SetMarkPrice(v string)`

SetMarkPrice sets MarkPrice field to given value.

### HasMarkPrice

`func (o *OptionPositionInformationResponseInner) HasMarkPrice() bool`

HasMarkPrice returns a boolean if a field has been set.

### GetStrikePrice

`func (o *OptionPositionInformationResponseInner) GetStrikePrice() string`

GetStrikePrice returns the StrikePrice field if non-nil, zero value otherwise.

### GetStrikePriceOk

`func (o *OptionPositionInformationResponseInner) GetStrikePriceOk() (*string, bool)`

GetStrikePriceOk returns a tuple with the StrikePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrikePrice

`func (o *OptionPositionInformationResponseInner) SetStrikePrice(v string)`

SetStrikePrice sets StrikePrice field to given value.

### HasStrikePrice

`func (o *OptionPositionInformationResponseInner) HasStrikePrice() bool`

HasStrikePrice returns a boolean if a field has been set.

### GetPositionCost

`func (o *OptionPositionInformationResponseInner) GetPositionCost() string`

GetPositionCost returns the PositionCost field if non-nil, zero value otherwise.

### GetPositionCostOk

`func (o *OptionPositionInformationResponseInner) GetPositionCostOk() (*string, bool)`

GetPositionCostOk returns a tuple with the PositionCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionCost

`func (o *OptionPositionInformationResponseInner) SetPositionCost(v string)`

SetPositionCost sets PositionCost field to given value.

### HasPositionCost

`func (o *OptionPositionInformationResponseInner) HasPositionCost() bool`

HasPositionCost returns a boolean if a field has been set.

### GetExpiryDate

`func (o *OptionPositionInformationResponseInner) GetExpiryDate() int64`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *OptionPositionInformationResponseInner) GetExpiryDateOk() (*int64, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *OptionPositionInformationResponseInner) SetExpiryDate(v int64)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *OptionPositionInformationResponseInner) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetPriceScale

`func (o *OptionPositionInformationResponseInner) GetPriceScale() int64`

GetPriceScale returns the PriceScale field if non-nil, zero value otherwise.

### GetPriceScaleOk

`func (o *OptionPositionInformationResponseInner) GetPriceScaleOk() (*int64, bool)`

GetPriceScaleOk returns a tuple with the PriceScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceScale

`func (o *OptionPositionInformationResponseInner) SetPriceScale(v int64)`

SetPriceScale sets PriceScale field to given value.

### HasPriceScale

`func (o *OptionPositionInformationResponseInner) HasPriceScale() bool`

HasPriceScale returns a boolean if a field has been set.

### GetQuantityScale

`func (o *OptionPositionInformationResponseInner) GetQuantityScale() int64`

GetQuantityScale returns the QuantityScale field if non-nil, zero value otherwise.

### GetQuantityScaleOk

`func (o *OptionPositionInformationResponseInner) GetQuantityScaleOk() (*int64, bool)`

GetQuantityScaleOk returns a tuple with the QuantityScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityScale

`func (o *OptionPositionInformationResponseInner) SetQuantityScale(v int64)`

SetQuantityScale sets QuantityScale field to given value.

### HasQuantityScale

`func (o *OptionPositionInformationResponseInner) HasQuantityScale() bool`

HasQuantityScale returns a boolean if a field has been set.

### GetOptionSide

`func (o *OptionPositionInformationResponseInner) GetOptionSide() string`

GetOptionSide returns the OptionSide field if non-nil, zero value otherwise.

### GetOptionSideOk

`func (o *OptionPositionInformationResponseInner) GetOptionSideOk() (*string, bool)`

GetOptionSideOk returns a tuple with the OptionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSide

`func (o *OptionPositionInformationResponseInner) SetOptionSide(v string)`

SetOptionSide sets OptionSide field to given value.

### HasOptionSide

`func (o *OptionPositionInformationResponseInner) HasOptionSide() bool`

HasOptionSide returns a boolean if a field has been set.

### GetQuoteAsset

`func (o *OptionPositionInformationResponseInner) GetQuoteAsset() string`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *OptionPositionInformationResponseInner) GetQuoteAssetOk() (*string, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *OptionPositionInformationResponseInner) SetQuoteAsset(v string)`

SetQuoteAsset sets QuoteAsset field to given value.

### HasQuoteAsset

`func (o *OptionPositionInformationResponseInner) HasQuoteAsset() bool`

HasQuoteAsset returns a boolean if a field has been set.


[[Back to README]](../README.md)


