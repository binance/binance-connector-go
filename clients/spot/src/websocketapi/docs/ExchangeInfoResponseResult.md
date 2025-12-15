# ExchangeInfoResponseResult

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Timezone** | Pointer to **string** |  | [optional] 
**ServerTime** | Pointer to **int64** |  | [optional] 
**RateLimits** | Pointer to [**[]RateLimits**](RateLimits.md) |  | [optional] 
**ExchangeFilters** | Pointer to [**[]ExchangeFilters**](ExchangeFilters.md) |  | [optional] 
**Symbols** | Pointer to [**[]ExchangeInfoResponseResultSymbolsInner**](ExchangeInfoResponseResultSymbolsInner.md) |  | [optional] 
**Sors** | Pointer to [**[]ExchangeInfoResponseResultSorsInner**](ExchangeInfoResponseResultSorsInner.md) |  | [optional] 

## Methods

### NewExchangeInfoResponseResult

`func NewExchangeInfoResponseResult() *ExchangeInfoResponseResult`

NewExchangeInfoResponseResult instantiates a new ExchangeInfoResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeInfoResponseResultWithDefaults

`func NewExchangeInfoResponseResultWithDefaults() *ExchangeInfoResponseResult`

NewExchangeInfoResponseResultWithDefaults instantiates a new ExchangeInfoResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimezone

`func (o *ExchangeInfoResponseResult) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ExchangeInfoResponseResult) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ExchangeInfoResponseResult) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ExchangeInfoResponseResult) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetServerTime

`func (o *ExchangeInfoResponseResult) GetServerTime() int64`

GetServerTime returns the ServerTime field if non-nil, zero value otherwise.

### GetServerTimeOk

`func (o *ExchangeInfoResponseResult) GetServerTimeOk() (*int64, bool)`

GetServerTimeOk returns a tuple with the ServerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTime

`func (o *ExchangeInfoResponseResult) SetServerTime(v int64)`

SetServerTime sets ServerTime field to given value.

### HasServerTime

`func (o *ExchangeInfoResponseResult) HasServerTime() bool`

HasServerTime returns a boolean if a field has been set.

### GetRateLimits

`func (o *ExchangeInfoResponseResult) GetRateLimits() []RateLimits`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *ExchangeInfoResponseResult) GetRateLimitsOk() (*[]RateLimits, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *ExchangeInfoResponseResult) SetRateLimits(v []RateLimits)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *ExchangeInfoResponseResult) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.

### GetExchangeFilters

`func (o *ExchangeInfoResponseResult) GetExchangeFilters() []ExchangeFilters`

GetExchangeFilters returns the ExchangeFilters field if non-nil, zero value otherwise.

### GetExchangeFiltersOk

`func (o *ExchangeInfoResponseResult) GetExchangeFiltersOk() (*[]ExchangeFilters, bool)`

GetExchangeFiltersOk returns a tuple with the ExchangeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeFilters

`func (o *ExchangeInfoResponseResult) SetExchangeFilters(v []ExchangeFilters)`

SetExchangeFilters sets ExchangeFilters field to given value.

### HasExchangeFilters

`func (o *ExchangeInfoResponseResult) HasExchangeFilters() bool`

HasExchangeFilters returns a boolean if a field has been set.

### GetSymbols

`func (o *ExchangeInfoResponseResult) GetSymbols() []ExchangeInfoResponseResultSymbolsInner`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *ExchangeInfoResponseResult) GetSymbolsOk() (*[]ExchangeInfoResponseResultSymbolsInner, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *ExchangeInfoResponseResult) SetSymbols(v []ExchangeInfoResponseResultSymbolsInner)`

SetSymbols sets Symbols field to given value.

### HasSymbols

`func (o *ExchangeInfoResponseResult) HasSymbols() bool`

HasSymbols returns a boolean if a field has been set.

### GetSors

`func (o *ExchangeInfoResponseResult) GetSors() []ExchangeInfoResponseResultSorsInner`

GetSors returns the Sors field if non-nil, zero value otherwise.

### GetSorsOk

`func (o *ExchangeInfoResponseResult) GetSorsOk() (*[]ExchangeInfoResponseResultSorsInner, bool)`

GetSorsOk returns a tuple with the Sors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSors

`func (o *ExchangeInfoResponseResult) SetSors(v []ExchangeInfoResponseResultSorsInner)`

SetSors sets Sors field to given value.

### HasSors

`func (o *ExchangeInfoResponseResult) HasSors() bool`

HasSors returns a boolean if a field has been set.


[[Back to README]](../README.md)


