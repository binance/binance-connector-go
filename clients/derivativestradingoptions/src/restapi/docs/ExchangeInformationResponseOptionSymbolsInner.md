# ExchangeInformationResponseOptionSymbolsInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**ExpiryDate** | Pointer to **int64** |  | [optional] 
**Filters** | Pointer to [**[]ExchangeInformationResponseOptionSymbolsInnerFiltersInner**](ExchangeInformationResponseOptionSymbolsInnerFiltersInner.md) |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**StrikePrice** | Pointer to **string** |  | [optional] 
**Underlying** | Pointer to **string** |  | [optional] 
**Unit** | Pointer to **int64** |  | [optional] 
**MakerFeeRate** | Pointer to **string** |  | [optional] 
**TakerFeeRate** | Pointer to **string** |  | [optional] 
**LiquidationFeeRate** | Pointer to **string** |  | [optional] 
**MinQty** | Pointer to **string** |  | [optional] 
**MaxQty** | Pointer to **string** |  | [optional] 
**InitialMargin** | Pointer to **string** |  | [optional] 
**MaintenanceMargin** | Pointer to **string** |  | [optional] 
**MinInitialMargin** | Pointer to **string** |  | [optional] 
**MinMaintenanceMargin** | Pointer to **string** |  | [optional] 
**PriceScale** | Pointer to **int64** |  | [optional] 
**QuantityScale** | Pointer to **int64** |  | [optional] 
**QuoteAsset** | Pointer to **string** |  | [optional] 

## Methods

### NewExchangeInformationResponseOptionSymbolsInner

`func NewExchangeInformationResponseOptionSymbolsInner() *ExchangeInformationResponseOptionSymbolsInner`

NewExchangeInformationResponseOptionSymbolsInner instantiates a new ExchangeInformationResponseOptionSymbolsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeInformationResponseOptionSymbolsInnerWithDefaults

`func NewExchangeInformationResponseOptionSymbolsInnerWithDefaults() *ExchangeInformationResponseOptionSymbolsInner`

NewExchangeInformationResponseOptionSymbolsInnerWithDefaults instantiates a new ExchangeInformationResponseOptionSymbolsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiryDate

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetExpiryDate() int64`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetExpiryDateOk() (*int64, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *ExchangeInformationResponseOptionSymbolsInner) SetExpiryDate(v int64)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *ExchangeInformationResponseOptionSymbolsInner) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetFilters

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetFilters() []ExchangeInformationResponseOptionSymbolsInnerFiltersInner`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetFiltersOk() (*[]ExchangeInformationResponseOptionSymbolsInnerFiltersInner, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ExchangeInformationResponseOptionSymbolsInner) SetFilters(v []ExchangeInformationResponseOptionSymbolsInnerFiltersInner)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ExchangeInformationResponseOptionSymbolsInner) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSymbol

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ExchangeInformationResponseOptionSymbolsInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *ExchangeInformationResponseOptionSymbolsInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetSide

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *ExchangeInformationResponseOptionSymbolsInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *ExchangeInformationResponseOptionSymbolsInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetStrikePrice

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetStrikePrice() string`

GetStrikePrice returns the StrikePrice field if non-nil, zero value otherwise.

### GetStrikePriceOk

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetStrikePriceOk() (*string, bool)`

GetStrikePriceOk returns a tuple with the StrikePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrikePrice

`func (o *ExchangeInformationResponseOptionSymbolsInner) SetStrikePrice(v string)`

SetStrikePrice sets StrikePrice field to given value.

### HasStrikePrice

`func (o *ExchangeInformationResponseOptionSymbolsInner) HasStrikePrice() bool`

HasStrikePrice returns a boolean if a field has been set.

### GetUnderlying

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetUnderlying() string`

GetUnderlying returns the Underlying field if non-nil, zero value otherwise.

### GetUnderlyingOk

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetUnderlyingOk() (*string, bool)`

GetUnderlyingOk returns a tuple with the Underlying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlying

`func (o *ExchangeInformationResponseOptionSymbolsInner) SetUnderlying(v string)`

SetUnderlying sets Underlying field to given value.

### HasUnderlying

`func (o *ExchangeInformationResponseOptionSymbolsInner) HasUnderlying() bool`

HasUnderlying returns a boolean if a field has been set.

### GetUnit

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetUnit() int64`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetUnitOk() (*int64, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ExchangeInformationResponseOptionSymbolsInner) SetUnit(v int64)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *ExchangeInformationResponseOptionSymbolsInner) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetMakerFeeRate

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetMakerFeeRate() string`

GetMakerFeeRate returns the MakerFeeRate field if non-nil, zero value otherwise.

### GetMakerFeeRateOk

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetMakerFeeRateOk() (*string, bool)`

GetMakerFeeRateOk returns a tuple with the MakerFeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerFeeRate

`func (o *ExchangeInformationResponseOptionSymbolsInner) SetMakerFeeRate(v string)`

SetMakerFeeRate sets MakerFeeRate field to given value.

### HasMakerFeeRate

`func (o *ExchangeInformationResponseOptionSymbolsInner) HasMakerFeeRate() bool`

HasMakerFeeRate returns a boolean if a field has been set.

### GetTakerFeeRate

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetTakerFeeRate() string`

GetTakerFeeRate returns the TakerFeeRate field if non-nil, zero value otherwise.

### GetTakerFeeRateOk

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetTakerFeeRateOk() (*string, bool)`

GetTakerFeeRateOk returns a tuple with the TakerFeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakerFeeRate

`func (o *ExchangeInformationResponseOptionSymbolsInner) SetTakerFeeRate(v string)`

SetTakerFeeRate sets TakerFeeRate field to given value.

### HasTakerFeeRate

`func (o *ExchangeInformationResponseOptionSymbolsInner) HasTakerFeeRate() bool`

HasTakerFeeRate returns a boolean if a field has been set.

### GetLiquidationFeeRate

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetLiquidationFeeRate() string`

GetLiquidationFeeRate returns the LiquidationFeeRate field if non-nil, zero value otherwise.

### GetLiquidationFeeRateOk

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetLiquidationFeeRateOk() (*string, bool)`

GetLiquidationFeeRateOk returns a tuple with the LiquidationFeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidationFeeRate

`func (o *ExchangeInformationResponseOptionSymbolsInner) SetLiquidationFeeRate(v string)`

SetLiquidationFeeRate sets LiquidationFeeRate field to given value.

### HasLiquidationFeeRate

`func (o *ExchangeInformationResponseOptionSymbolsInner) HasLiquidationFeeRate() bool`

HasLiquidationFeeRate returns a boolean if a field has been set.

### GetMinQty

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetMinQty() string`

GetMinQty returns the MinQty field if non-nil, zero value otherwise.

### GetMinQtyOk

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetMinQtyOk() (*string, bool)`

GetMinQtyOk returns a tuple with the MinQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQty

`func (o *ExchangeInformationResponseOptionSymbolsInner) SetMinQty(v string)`

SetMinQty sets MinQty field to given value.

### HasMinQty

`func (o *ExchangeInformationResponseOptionSymbolsInner) HasMinQty() bool`

HasMinQty returns a boolean if a field has been set.

### GetMaxQty

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetMaxQty() string`

GetMaxQty returns the MaxQty field if non-nil, zero value otherwise.

### GetMaxQtyOk

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetMaxQtyOk() (*string, bool)`

GetMaxQtyOk returns a tuple with the MaxQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQty

`func (o *ExchangeInformationResponseOptionSymbolsInner) SetMaxQty(v string)`

SetMaxQty sets MaxQty field to given value.

### HasMaxQty

`func (o *ExchangeInformationResponseOptionSymbolsInner) HasMaxQty() bool`

HasMaxQty returns a boolean if a field has been set.

### GetInitialMargin

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetInitialMargin() string`

GetInitialMargin returns the InitialMargin field if non-nil, zero value otherwise.

### GetInitialMarginOk

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetInitialMarginOk() (*string, bool)`

GetInitialMarginOk returns a tuple with the InitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialMargin

`func (o *ExchangeInformationResponseOptionSymbolsInner) SetInitialMargin(v string)`

SetInitialMargin sets InitialMargin field to given value.

### HasInitialMargin

`func (o *ExchangeInformationResponseOptionSymbolsInner) HasInitialMargin() bool`

HasInitialMargin returns a boolean if a field has been set.

### GetMaintenanceMargin

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetMaintenanceMargin() string`

GetMaintenanceMargin returns the MaintenanceMargin field if non-nil, zero value otherwise.

### GetMaintenanceMarginOk

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetMaintenanceMarginOk() (*string, bool)`

GetMaintenanceMarginOk returns a tuple with the MaintenanceMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceMargin

`func (o *ExchangeInformationResponseOptionSymbolsInner) SetMaintenanceMargin(v string)`

SetMaintenanceMargin sets MaintenanceMargin field to given value.

### HasMaintenanceMargin

`func (o *ExchangeInformationResponseOptionSymbolsInner) HasMaintenanceMargin() bool`

HasMaintenanceMargin returns a boolean if a field has been set.

### GetMinInitialMargin

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetMinInitialMargin() string`

GetMinInitialMargin returns the MinInitialMargin field if non-nil, zero value otherwise.

### GetMinInitialMarginOk

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetMinInitialMarginOk() (*string, bool)`

GetMinInitialMarginOk returns a tuple with the MinInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInitialMargin

`func (o *ExchangeInformationResponseOptionSymbolsInner) SetMinInitialMargin(v string)`

SetMinInitialMargin sets MinInitialMargin field to given value.

### HasMinInitialMargin

`func (o *ExchangeInformationResponseOptionSymbolsInner) HasMinInitialMargin() bool`

HasMinInitialMargin returns a boolean if a field has been set.

### GetMinMaintenanceMargin

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetMinMaintenanceMargin() string`

GetMinMaintenanceMargin returns the MinMaintenanceMargin field if non-nil, zero value otherwise.

### GetMinMaintenanceMarginOk

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetMinMaintenanceMarginOk() (*string, bool)`

GetMinMaintenanceMarginOk returns a tuple with the MinMaintenanceMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMaintenanceMargin

`func (o *ExchangeInformationResponseOptionSymbolsInner) SetMinMaintenanceMargin(v string)`

SetMinMaintenanceMargin sets MinMaintenanceMargin field to given value.

### HasMinMaintenanceMargin

`func (o *ExchangeInformationResponseOptionSymbolsInner) HasMinMaintenanceMargin() bool`

HasMinMaintenanceMargin returns a boolean if a field has been set.

### GetPriceScale

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetPriceScale() int64`

GetPriceScale returns the PriceScale field if non-nil, zero value otherwise.

### GetPriceScaleOk

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetPriceScaleOk() (*int64, bool)`

GetPriceScaleOk returns a tuple with the PriceScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceScale

`func (o *ExchangeInformationResponseOptionSymbolsInner) SetPriceScale(v int64)`

SetPriceScale sets PriceScale field to given value.

### HasPriceScale

`func (o *ExchangeInformationResponseOptionSymbolsInner) HasPriceScale() bool`

HasPriceScale returns a boolean if a field has been set.

### GetQuantityScale

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetQuantityScale() int64`

GetQuantityScale returns the QuantityScale field if non-nil, zero value otherwise.

### GetQuantityScaleOk

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetQuantityScaleOk() (*int64, bool)`

GetQuantityScaleOk returns a tuple with the QuantityScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityScale

`func (o *ExchangeInformationResponseOptionSymbolsInner) SetQuantityScale(v int64)`

SetQuantityScale sets QuantityScale field to given value.

### HasQuantityScale

`func (o *ExchangeInformationResponseOptionSymbolsInner) HasQuantityScale() bool`

HasQuantityScale returns a boolean if a field has been set.

### GetQuoteAsset

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetQuoteAsset() string`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *ExchangeInformationResponseOptionSymbolsInner) GetQuoteAssetOk() (*string, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *ExchangeInformationResponseOptionSymbolsInner) SetQuoteAsset(v string)`

SetQuoteAsset sets QuoteAsset field to given value.

### HasQuoteAsset

`func (o *ExchangeInformationResponseOptionSymbolsInner) HasQuoteAsset() bool`

HasQuoteAsset returns a boolean if a field has been set.


[[Back to README]](../README.md)


