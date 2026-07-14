# ExchangeInfoResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Timezone** | Pointer to **string** |  | [optional] 
**ServerTime** | Pointer to **int64** |  | [optional] 
**RateLimits** | Pointer to [**[]MyFiltersResponseRateLimitsInner**](MyFiltersResponseRateLimitsInner.md) |  | [optional] 
**ExchangeFilters** | Pointer to [**[]MyFiltersResponseExchangeFiltersInner**](MyFiltersResponseExchangeFiltersInner.md) |  | [optional] 
**Symbols** | Pointer to [**[]ExchangeInfoResponseSymbolsInner**](ExchangeInfoResponseSymbolsInner.md) |  | [optional] 
**Sors** | Pointer to [**[]ExchangeInfoResponseSorsInner**](ExchangeInfoResponseSorsInner.md) | Optional. Present only when SOR is available. | [optional] 

## Methods

### NewExchangeInfoResponse

`func NewExchangeInfoResponse() *ExchangeInfoResponse`

NewExchangeInfoResponse instantiates a new ExchangeInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeInfoResponseWithDefaults

`func NewExchangeInfoResponseWithDefaults() *ExchangeInfoResponse`

NewExchangeInfoResponseWithDefaults instantiates a new ExchangeInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimezone

`func (o *ExchangeInfoResponse) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ExchangeInfoResponse) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ExchangeInfoResponse) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ExchangeInfoResponse) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetServerTime

`func (o *ExchangeInfoResponse) GetServerTime() int64`

GetServerTime returns the ServerTime field if non-nil, zero value otherwise.

### GetServerTimeOk

`func (o *ExchangeInfoResponse) GetServerTimeOk() (*int64, bool)`

GetServerTimeOk returns a tuple with the ServerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTime

`func (o *ExchangeInfoResponse) SetServerTime(v int64)`

SetServerTime sets ServerTime field to given value.

### HasServerTime

`func (o *ExchangeInfoResponse) HasServerTime() bool`

HasServerTime returns a boolean if a field has been set.

### GetRateLimits

`func (o *ExchangeInfoResponse) GetRateLimits() []MyFiltersResponseRateLimitsInner`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *ExchangeInfoResponse) GetRateLimitsOk() (*[]MyFiltersResponseRateLimitsInner, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *ExchangeInfoResponse) SetRateLimits(v []MyFiltersResponseRateLimitsInner)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *ExchangeInfoResponse) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.

### GetExchangeFilters

`func (o *ExchangeInfoResponse) GetExchangeFilters() []MyFiltersResponseExchangeFiltersInner`

GetExchangeFilters returns the ExchangeFilters field if non-nil, zero value otherwise.

### GetExchangeFiltersOk

`func (o *ExchangeInfoResponse) GetExchangeFiltersOk() (*[]MyFiltersResponseExchangeFiltersInner, bool)`

GetExchangeFiltersOk returns a tuple with the ExchangeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeFilters

`func (o *ExchangeInfoResponse) SetExchangeFilters(v []MyFiltersResponseExchangeFiltersInner)`

SetExchangeFilters sets ExchangeFilters field to given value.

### HasExchangeFilters

`func (o *ExchangeInfoResponse) HasExchangeFilters() bool`

HasExchangeFilters returns a boolean if a field has been set.

### GetSymbols

`func (o *ExchangeInfoResponse) GetSymbols() []ExchangeInfoResponseSymbolsInner`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *ExchangeInfoResponse) GetSymbolsOk() (*[]ExchangeInfoResponseSymbolsInner, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *ExchangeInfoResponse) SetSymbols(v []ExchangeInfoResponseSymbolsInner)`

SetSymbols sets Symbols field to given value.

### HasSymbols

`func (o *ExchangeInfoResponse) HasSymbols() bool`

HasSymbols returns a boolean if a field has been set.

### GetSors

`func (o *ExchangeInfoResponse) GetSors() []ExchangeInfoResponseSorsInner`

GetSors returns the Sors field if non-nil, zero value otherwise.

### GetSorsOk

`func (o *ExchangeInfoResponse) GetSorsOk() (*[]ExchangeInfoResponseSorsInner, bool)`

GetSorsOk returns a tuple with the Sors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSors

`func (o *ExchangeInfoResponse) SetSors(v []ExchangeInfoResponseSorsInner)`

SetSors sets Sors field to given value.

### HasSors

`func (o *ExchangeInfoResponse) HasSors() bool`

HasSors returns a boolean if a field has been set.


[[Back to README]](../README.md)


