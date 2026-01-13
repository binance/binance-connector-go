# QueryIsolatedMarginAccountInfoResponseAssetsInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**BaseAsset** | Pointer to [**QueryIsolatedMarginAccountInfoResponseAssetsInnerBaseAsset**](QueryIsolatedMarginAccountInfoResponseAssetsInnerBaseAsset.md) |  | [optional] 
**QuoteAsset** | Pointer to [**QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset**](QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset.md) |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**IsolatedCreated** | Pointer to **bool** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**MarginLevel** | Pointer to **string** |  | [optional] 
**MarginLevelStatus** | Pointer to **string** |  | [optional] 
**MarginRatio** | Pointer to **string** |  | [optional] 
**IndexPrice** | Pointer to **string** |  | [optional] 
**LiquidatePrice** | Pointer to **string** |  | [optional] 
**LiquidateRate** | Pointer to **string** |  | [optional] 
**TradeEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewQueryIsolatedMarginAccountInfoResponseAssetsInner

`func NewQueryIsolatedMarginAccountInfoResponseAssetsInner() *QueryIsolatedMarginAccountInfoResponseAssetsInner`

NewQueryIsolatedMarginAccountInfoResponseAssetsInner instantiates a new QueryIsolatedMarginAccountInfoResponseAssetsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryIsolatedMarginAccountInfoResponseAssetsInnerWithDefaults

`func NewQueryIsolatedMarginAccountInfoResponseAssetsInnerWithDefaults() *QueryIsolatedMarginAccountInfoResponseAssetsInner`

NewQueryIsolatedMarginAccountInfoResponseAssetsInnerWithDefaults instantiates a new QueryIsolatedMarginAccountInfoResponseAssetsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseAsset

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetBaseAsset() QueryIsolatedMarginAccountInfoResponseAssetsInnerBaseAsset`

GetBaseAsset returns the BaseAsset field if non-nil, zero value otherwise.

### GetBaseAssetOk

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetBaseAssetOk() (*QueryIsolatedMarginAccountInfoResponseAssetsInnerBaseAsset, bool)`

GetBaseAssetOk returns a tuple with the BaseAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAsset

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) SetBaseAsset(v QueryIsolatedMarginAccountInfoResponseAssetsInnerBaseAsset)`

SetBaseAsset sets BaseAsset field to given value.

### HasBaseAsset

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) HasBaseAsset() bool`

HasBaseAsset returns a boolean if a field has been set.

### GetQuoteAsset

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetQuoteAsset() QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetQuoteAssetOk() (*QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) SetQuoteAsset(v QueryIsolatedMarginAccountInfoResponseAssetsInnerQuoteAsset)`

SetQuoteAsset sets QuoteAsset field to given value.

### HasQuoteAsset

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) HasQuoteAsset() bool`

HasQuoteAsset returns a boolean if a field has been set.

### GetSymbol

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetIsolatedCreated

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetIsolatedCreated() bool`

GetIsolatedCreated returns the IsolatedCreated field if non-nil, zero value otherwise.

### GetIsolatedCreatedOk

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetIsolatedCreatedOk() (*bool, bool)`

GetIsolatedCreatedOk returns a tuple with the IsolatedCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolatedCreated

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) SetIsolatedCreated(v bool)`

SetIsolatedCreated sets IsolatedCreated field to given value.

### HasIsolatedCreated

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) HasIsolatedCreated() bool`

HasIsolatedCreated returns a boolean if a field has been set.

### GetEnabled

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMarginLevel

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetMarginLevel() string`

GetMarginLevel returns the MarginLevel field if non-nil, zero value otherwise.

### GetMarginLevelOk

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetMarginLevelOk() (*string, bool)`

GetMarginLevelOk returns a tuple with the MarginLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginLevel

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) SetMarginLevel(v string)`

SetMarginLevel sets MarginLevel field to given value.

### HasMarginLevel

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) HasMarginLevel() bool`

HasMarginLevel returns a boolean if a field has been set.

### GetMarginLevelStatus

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetMarginLevelStatus() string`

GetMarginLevelStatus returns the MarginLevelStatus field if non-nil, zero value otherwise.

### GetMarginLevelStatusOk

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetMarginLevelStatusOk() (*string, bool)`

GetMarginLevelStatusOk returns a tuple with the MarginLevelStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginLevelStatus

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) SetMarginLevelStatus(v string)`

SetMarginLevelStatus sets MarginLevelStatus field to given value.

### HasMarginLevelStatus

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) HasMarginLevelStatus() bool`

HasMarginLevelStatus returns a boolean if a field has been set.

### GetMarginRatio

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetMarginRatio() string`

GetMarginRatio returns the MarginRatio field if non-nil, zero value otherwise.

### GetMarginRatioOk

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetMarginRatioOk() (*string, bool)`

GetMarginRatioOk returns a tuple with the MarginRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginRatio

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) SetMarginRatio(v string)`

SetMarginRatio sets MarginRatio field to given value.

### HasMarginRatio

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) HasMarginRatio() bool`

HasMarginRatio returns a boolean if a field has been set.

### GetIndexPrice

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetIndexPrice() string`

GetIndexPrice returns the IndexPrice field if non-nil, zero value otherwise.

### GetIndexPriceOk

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetIndexPriceOk() (*string, bool)`

GetIndexPriceOk returns a tuple with the IndexPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexPrice

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) SetIndexPrice(v string)`

SetIndexPrice sets IndexPrice field to given value.

### HasIndexPrice

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) HasIndexPrice() bool`

HasIndexPrice returns a boolean if a field has been set.

### GetLiquidatePrice

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetLiquidatePrice() string`

GetLiquidatePrice returns the LiquidatePrice field if non-nil, zero value otherwise.

### GetLiquidatePriceOk

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetLiquidatePriceOk() (*string, bool)`

GetLiquidatePriceOk returns a tuple with the LiquidatePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidatePrice

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) SetLiquidatePrice(v string)`

SetLiquidatePrice sets LiquidatePrice field to given value.

### HasLiquidatePrice

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) HasLiquidatePrice() bool`

HasLiquidatePrice returns a boolean if a field has been set.

### GetLiquidateRate

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetLiquidateRate() string`

GetLiquidateRate returns the LiquidateRate field if non-nil, zero value otherwise.

### GetLiquidateRateOk

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetLiquidateRateOk() (*string, bool)`

GetLiquidateRateOk returns a tuple with the LiquidateRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidateRate

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) SetLiquidateRate(v string)`

SetLiquidateRate sets LiquidateRate field to given value.

### HasLiquidateRate

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) HasLiquidateRate() bool`

HasLiquidateRate returns a boolean if a field has been set.

### GetTradeEnabled

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetTradeEnabled() bool`

GetTradeEnabled returns the TradeEnabled field if non-nil, zero value otherwise.

### GetTradeEnabledOk

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) GetTradeEnabledOk() (*bool, bool)`

GetTradeEnabledOk returns a tuple with the TradeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeEnabled

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) SetTradeEnabled(v bool)`

SetTradeEnabled sets TradeEnabled field to given value.

### HasTradeEnabled

`func (o *QueryIsolatedMarginAccountInfoResponseAssetsInner) HasTradeEnabled() bool`

HasTradeEnabled returns a boolean if a field has been set.


[[Back to README]](../README.md)


