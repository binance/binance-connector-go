# ExchangeInformationResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**ExchangeFilters** | Pointer to **[]string** |  | [optional] 
**RateLimits** | Pointer to [**[]ExchangeInformationResponseRateLimitsInner**](ExchangeInformationResponseRateLimitsInner.md) |  | [optional] 
**ServerTime** | Pointer to **int64** |  | [optional] 
**Assets** | Pointer to [**[]ExchangeInformationResponseAssetsInner**](ExchangeInformationResponseAssetsInner.md) |  | [optional] 
**Symbols** | Pointer to [**[]ExchangeInformationResponseSymbolsInner**](ExchangeInformationResponseSymbolsInner.md) |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 

## Methods

### NewExchangeInformationResponse

`func NewExchangeInformationResponse() *ExchangeInformationResponse`

NewExchangeInformationResponse instantiates a new ExchangeInformationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeInformationResponseWithDefaults

`func NewExchangeInformationResponseWithDefaults() *ExchangeInformationResponse`

NewExchangeInformationResponseWithDefaults instantiates a new ExchangeInformationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchangeFilters

`func (o *ExchangeInformationResponse) GetExchangeFilters() []string`

GetExchangeFilters returns the ExchangeFilters field if non-nil, zero value otherwise.

### GetExchangeFiltersOk

`func (o *ExchangeInformationResponse) GetExchangeFiltersOk() (*[]string, bool)`

GetExchangeFiltersOk returns a tuple with the ExchangeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeFilters

`func (o *ExchangeInformationResponse) SetExchangeFilters(v []string)`

SetExchangeFilters sets ExchangeFilters field to given value.

### HasExchangeFilters

`func (o *ExchangeInformationResponse) HasExchangeFilters() bool`

HasExchangeFilters returns a boolean if a field has been set.

### GetRateLimits

`func (o *ExchangeInformationResponse) GetRateLimits() []ExchangeInformationResponseRateLimitsInner`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *ExchangeInformationResponse) GetRateLimitsOk() (*[]ExchangeInformationResponseRateLimitsInner, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *ExchangeInformationResponse) SetRateLimits(v []ExchangeInformationResponseRateLimitsInner)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *ExchangeInformationResponse) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.

### GetServerTime

`func (o *ExchangeInformationResponse) GetServerTime() int64`

GetServerTime returns the ServerTime field if non-nil, zero value otherwise.

### GetServerTimeOk

`func (o *ExchangeInformationResponse) GetServerTimeOk() (*int64, bool)`

GetServerTimeOk returns a tuple with the ServerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTime

`func (o *ExchangeInformationResponse) SetServerTime(v int64)`

SetServerTime sets ServerTime field to given value.

### HasServerTime

`func (o *ExchangeInformationResponse) HasServerTime() bool`

HasServerTime returns a boolean if a field has been set.

### GetAssets

`func (o *ExchangeInformationResponse) GetAssets() []ExchangeInformationResponseAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *ExchangeInformationResponse) GetAssetsOk() (*[]ExchangeInformationResponseAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *ExchangeInformationResponse) SetAssets(v []ExchangeInformationResponseAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *ExchangeInformationResponse) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetSymbols

`func (o *ExchangeInformationResponse) GetSymbols() []ExchangeInformationResponseSymbolsInner`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *ExchangeInformationResponse) GetSymbolsOk() (*[]ExchangeInformationResponseSymbolsInner, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *ExchangeInformationResponse) SetSymbols(v []ExchangeInformationResponseSymbolsInner)`

SetSymbols sets Symbols field to given value.

### HasSymbols

`func (o *ExchangeInformationResponse) HasSymbols() bool`

HasSymbols returns a boolean if a field has been set.

### GetTimezone

`func (o *ExchangeInformationResponse) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ExchangeInformationResponse) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ExchangeInformationResponse) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ExchangeInformationResponse) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to README]](../README.md)


