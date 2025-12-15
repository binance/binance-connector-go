# MyFiltersResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**ExchangeFilters** | Pointer to [**[]ExchangeFilters**](ExchangeFilters.md) |  | [optional] 
**SymbolFilters** | Pointer to [**[]SymbolFilters**](SymbolFilters.md) |  | [optional] 
**AssetFilters** | Pointer to [**[]AssetFilters**](AssetFilters.md) |  | [optional] 
**RateLimits** | Pointer to [**[]RateLimits**](RateLimits.md) |  | [optional] 

## Methods

### NewMyFiltersResponse

`func NewMyFiltersResponse() *MyFiltersResponse`

NewMyFiltersResponse instantiates a new MyFiltersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMyFiltersResponseWithDefaults

`func NewMyFiltersResponseWithDefaults() *MyFiltersResponse`

NewMyFiltersResponseWithDefaults instantiates a new MyFiltersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchangeFilters

`func (o *MyFiltersResponse) GetExchangeFilters() []ExchangeFilters`

GetExchangeFilters returns the ExchangeFilters field if non-nil, zero value otherwise.

### GetExchangeFiltersOk

`func (o *MyFiltersResponse) GetExchangeFiltersOk() (*[]ExchangeFilters, bool)`

GetExchangeFiltersOk returns a tuple with the ExchangeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeFilters

`func (o *MyFiltersResponse) SetExchangeFilters(v []ExchangeFilters)`

SetExchangeFilters sets ExchangeFilters field to given value.

### HasExchangeFilters

`func (o *MyFiltersResponse) HasExchangeFilters() bool`

HasExchangeFilters returns a boolean if a field has been set.

### GetSymbolFilters

`func (o *MyFiltersResponse) GetSymbolFilters() []SymbolFilters`

GetSymbolFilters returns the SymbolFilters field if non-nil, zero value otherwise.

### GetSymbolFiltersOk

`func (o *MyFiltersResponse) GetSymbolFiltersOk() (*[]SymbolFilters, bool)`

GetSymbolFiltersOk returns a tuple with the SymbolFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbolFilters

`func (o *MyFiltersResponse) SetSymbolFilters(v []SymbolFilters)`

SetSymbolFilters sets SymbolFilters field to given value.

### HasSymbolFilters

`func (o *MyFiltersResponse) HasSymbolFilters() bool`

HasSymbolFilters returns a boolean if a field has been set.

### GetAssetFilters

`func (o *MyFiltersResponse) GetAssetFilters() []AssetFilters`

GetAssetFilters returns the AssetFilters field if non-nil, zero value otherwise.

### GetAssetFiltersOk

`func (o *MyFiltersResponse) GetAssetFiltersOk() (*[]AssetFilters, bool)`

GetAssetFiltersOk returns a tuple with the AssetFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetFilters

`func (o *MyFiltersResponse) SetAssetFilters(v []AssetFilters)`

SetAssetFilters sets AssetFilters field to given value.

### HasAssetFilters

`func (o *MyFiltersResponse) HasAssetFilters() bool`

HasAssetFilters returns a boolean if a field has been set.

### GetRateLimits

`func (o *MyFiltersResponse) GetRateLimits() []RateLimits`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *MyFiltersResponse) GetRateLimitsOk() (*[]RateLimits, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *MyFiltersResponse) SetRateLimits(v []RateLimits)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *MyFiltersResponse) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


