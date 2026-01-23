# GetExchangeInfoResponseDataSymbolsInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**BaseAsset** | Pointer to **string** |  | [optional] 
**QuoteAsset** | Pointer to **string** |  | [optional] 
**PricePrecision** | Pointer to **int64** |  | [optional] 
**QuantityPrecision** | Pointer to **int64** |  | [optional] 
**BaseAssetPrecision** | Pointer to **int64** |  | [optional] 
**QuotePrecision** | Pointer to **int64** |  | [optional] 
**Filters** | Pointer to [**[]GetExchangeInfoResponseDataSymbolsInnerFiltersInner**](GetExchangeInfoResponseDataSymbolsInnerFiltersInner.md) |  | [optional] 
**OrderTypes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetExchangeInfoResponseDataSymbolsInner

`func NewGetExchangeInfoResponseDataSymbolsInner() *GetExchangeInfoResponseDataSymbolsInner`

NewGetExchangeInfoResponseDataSymbolsInner instantiates a new GetExchangeInfoResponseDataSymbolsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetExchangeInfoResponseDataSymbolsInnerWithDefaults

`func NewGetExchangeInfoResponseDataSymbolsInnerWithDefaults() *GetExchangeInfoResponseDataSymbolsInner`

NewGetExchangeInfoResponseDataSymbolsInnerWithDefaults instantiates a new GetExchangeInfoResponseDataSymbolsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *GetExchangeInfoResponseDataSymbolsInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetExchangeInfoResponseDataSymbolsInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetExchangeInfoResponseDataSymbolsInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetExchangeInfoResponseDataSymbolsInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetStatus

`func (o *GetExchangeInfoResponseDataSymbolsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetExchangeInfoResponseDataSymbolsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetExchangeInfoResponseDataSymbolsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetExchangeInfoResponseDataSymbolsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBaseAsset

`func (o *GetExchangeInfoResponseDataSymbolsInner) GetBaseAsset() string`

GetBaseAsset returns the BaseAsset field if non-nil, zero value otherwise.

### GetBaseAssetOk

`func (o *GetExchangeInfoResponseDataSymbolsInner) GetBaseAssetOk() (*string, bool)`

GetBaseAssetOk returns a tuple with the BaseAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAsset

`func (o *GetExchangeInfoResponseDataSymbolsInner) SetBaseAsset(v string)`

SetBaseAsset sets BaseAsset field to given value.

### HasBaseAsset

`func (o *GetExchangeInfoResponseDataSymbolsInner) HasBaseAsset() bool`

HasBaseAsset returns a boolean if a field has been set.

### GetQuoteAsset

`func (o *GetExchangeInfoResponseDataSymbolsInner) GetQuoteAsset() string`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *GetExchangeInfoResponseDataSymbolsInner) GetQuoteAssetOk() (*string, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *GetExchangeInfoResponseDataSymbolsInner) SetQuoteAsset(v string)`

SetQuoteAsset sets QuoteAsset field to given value.

### HasQuoteAsset

`func (o *GetExchangeInfoResponseDataSymbolsInner) HasQuoteAsset() bool`

HasQuoteAsset returns a boolean if a field has been set.

### GetPricePrecision

`func (o *GetExchangeInfoResponseDataSymbolsInner) GetPricePrecision() int64`

GetPricePrecision returns the PricePrecision field if non-nil, zero value otherwise.

### GetPricePrecisionOk

`func (o *GetExchangeInfoResponseDataSymbolsInner) GetPricePrecisionOk() (*int64, bool)`

GetPricePrecisionOk returns a tuple with the PricePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePrecision

`func (o *GetExchangeInfoResponseDataSymbolsInner) SetPricePrecision(v int64)`

SetPricePrecision sets PricePrecision field to given value.

### HasPricePrecision

`func (o *GetExchangeInfoResponseDataSymbolsInner) HasPricePrecision() bool`

HasPricePrecision returns a boolean if a field has been set.

### GetQuantityPrecision

`func (o *GetExchangeInfoResponseDataSymbolsInner) GetQuantityPrecision() int64`

GetQuantityPrecision returns the QuantityPrecision field if non-nil, zero value otherwise.

### GetQuantityPrecisionOk

`func (o *GetExchangeInfoResponseDataSymbolsInner) GetQuantityPrecisionOk() (*int64, bool)`

GetQuantityPrecisionOk returns a tuple with the QuantityPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityPrecision

`func (o *GetExchangeInfoResponseDataSymbolsInner) SetQuantityPrecision(v int64)`

SetQuantityPrecision sets QuantityPrecision field to given value.

### HasQuantityPrecision

`func (o *GetExchangeInfoResponseDataSymbolsInner) HasQuantityPrecision() bool`

HasQuantityPrecision returns a boolean if a field has been set.

### GetBaseAssetPrecision

`func (o *GetExchangeInfoResponseDataSymbolsInner) GetBaseAssetPrecision() int64`

GetBaseAssetPrecision returns the BaseAssetPrecision field if non-nil, zero value otherwise.

### GetBaseAssetPrecisionOk

`func (o *GetExchangeInfoResponseDataSymbolsInner) GetBaseAssetPrecisionOk() (*int64, bool)`

GetBaseAssetPrecisionOk returns a tuple with the BaseAssetPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAssetPrecision

`func (o *GetExchangeInfoResponseDataSymbolsInner) SetBaseAssetPrecision(v int64)`

SetBaseAssetPrecision sets BaseAssetPrecision field to given value.

### HasBaseAssetPrecision

`func (o *GetExchangeInfoResponseDataSymbolsInner) HasBaseAssetPrecision() bool`

HasBaseAssetPrecision returns a boolean if a field has been set.

### GetQuotePrecision

`func (o *GetExchangeInfoResponseDataSymbolsInner) GetQuotePrecision() int64`

GetQuotePrecision returns the QuotePrecision field if non-nil, zero value otherwise.

### GetQuotePrecisionOk

`func (o *GetExchangeInfoResponseDataSymbolsInner) GetQuotePrecisionOk() (*int64, bool)`

GetQuotePrecisionOk returns a tuple with the QuotePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotePrecision

`func (o *GetExchangeInfoResponseDataSymbolsInner) SetQuotePrecision(v int64)`

SetQuotePrecision sets QuotePrecision field to given value.

### HasQuotePrecision

`func (o *GetExchangeInfoResponseDataSymbolsInner) HasQuotePrecision() bool`

HasQuotePrecision returns a boolean if a field has been set.

### GetFilters

`func (o *GetExchangeInfoResponseDataSymbolsInner) GetFilters() []GetExchangeInfoResponseDataSymbolsInnerFiltersInner`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *GetExchangeInfoResponseDataSymbolsInner) GetFiltersOk() (*[]GetExchangeInfoResponseDataSymbolsInnerFiltersInner, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *GetExchangeInfoResponseDataSymbolsInner) SetFilters(v []GetExchangeInfoResponseDataSymbolsInnerFiltersInner)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *GetExchangeInfoResponseDataSymbolsInner) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetOrderTypes

`func (o *GetExchangeInfoResponseDataSymbolsInner) GetOrderTypes() []string`

GetOrderTypes returns the OrderTypes field if non-nil, zero value otherwise.

### GetOrderTypesOk

`func (o *GetExchangeInfoResponseDataSymbolsInner) GetOrderTypesOk() (*[]string, bool)`

GetOrderTypesOk returns a tuple with the OrderTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTypes

`func (o *GetExchangeInfoResponseDataSymbolsInner) SetOrderTypes(v []string)`

SetOrderTypes sets OrderTypes field to given value.

### HasOrderTypes

`func (o *GetExchangeInfoResponseDataSymbolsInner) HasOrderTypes() bool`

HasOrderTypes returns a boolean if a field has been set.


[[Back to README]](../README.md)


