# MyFiltersResponseResult

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**ExchangeFilters** | Pointer to [**[]ExchangeFilters**](ExchangeFilters.md) |  | [optional] 
**SymbolFilters** | Pointer to [**[]SymbolFilters**](SymbolFilters.md) |  | [optional] 
**AssetFilters** | Pointer to [**[]AssetFilters**](AssetFilters.md) |  | [optional] 

## Methods

### NewMyFiltersResponseResult

`func NewMyFiltersResponseResult() *MyFiltersResponseResult`

NewMyFiltersResponseResult instantiates a new MyFiltersResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMyFiltersResponseResultWithDefaults

`func NewMyFiltersResponseResultWithDefaults() *MyFiltersResponseResult`

NewMyFiltersResponseResultWithDefaults instantiates a new MyFiltersResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchangeFilters

`func (o *MyFiltersResponseResult) GetExchangeFilters() []ExchangeFilters`

GetExchangeFilters returns the ExchangeFilters field if non-nil, zero value otherwise.

### GetExchangeFiltersOk

`func (o *MyFiltersResponseResult) GetExchangeFiltersOk() (*[]ExchangeFilters, bool)`

GetExchangeFiltersOk returns a tuple with the ExchangeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeFilters

`func (o *MyFiltersResponseResult) SetExchangeFilters(v []ExchangeFilters)`

SetExchangeFilters sets ExchangeFilters field to given value.

### HasExchangeFilters

`func (o *MyFiltersResponseResult) HasExchangeFilters() bool`

HasExchangeFilters returns a boolean if a field has been set.

### GetSymbolFilters

`func (o *MyFiltersResponseResult) GetSymbolFilters() []SymbolFilters`

GetSymbolFilters returns the SymbolFilters field if non-nil, zero value otherwise.

### GetSymbolFiltersOk

`func (o *MyFiltersResponseResult) GetSymbolFiltersOk() (*[]SymbolFilters, bool)`

GetSymbolFiltersOk returns a tuple with the SymbolFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbolFilters

`func (o *MyFiltersResponseResult) SetSymbolFilters(v []SymbolFilters)`

SetSymbolFilters sets SymbolFilters field to given value.

### HasSymbolFilters

`func (o *MyFiltersResponseResult) HasSymbolFilters() bool`

HasSymbolFilters returns a boolean if a field has been set.

### GetAssetFilters

`func (o *MyFiltersResponseResult) GetAssetFilters() []AssetFilters`

GetAssetFilters returns the AssetFilters field if non-nil, zero value otherwise.

### GetAssetFiltersOk

`func (o *MyFiltersResponseResult) GetAssetFiltersOk() (*[]AssetFilters, bool)`

GetAssetFiltersOk returns a tuple with the AssetFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetFilters

`func (o *MyFiltersResponseResult) SetAssetFilters(v []AssetFilters)`

SetAssetFilters sets AssetFilters field to given value.

### HasAssetFilters

`func (o *MyFiltersResponseResult) HasAssetFilters() bool`

HasAssetFilters returns a boolean if a field has been set.


[[Back to README]](../README.md)


