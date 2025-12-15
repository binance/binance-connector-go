# ExchangeInformationResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Timezone** | Pointer to **string** |  | [optional] 
**ServerTime** | Pointer to **int64** |  | [optional] 
**OptionContracts** | Pointer to [**[]ExchangeInformationResponseOptionContractsInner**](ExchangeInformationResponseOptionContractsInner.md) |  | [optional] 
**OptionAssets** | Pointer to [**[]ExchangeInformationResponseOptionAssetsInner**](ExchangeInformationResponseOptionAssetsInner.md) |  | [optional] 
**OptionSymbols** | Pointer to [**[]ExchangeInformationResponseOptionSymbolsInner**](ExchangeInformationResponseOptionSymbolsInner.md) |  | [optional] 
**RateLimits** | Pointer to [**[]ExchangeInformationResponseRateLimitsInner**](ExchangeInformationResponseRateLimitsInner.md) |  | [optional] 

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

### GetOptionContracts

`func (o *ExchangeInformationResponse) GetOptionContracts() []ExchangeInformationResponseOptionContractsInner`

GetOptionContracts returns the OptionContracts field if non-nil, zero value otherwise.

### GetOptionContractsOk

`func (o *ExchangeInformationResponse) GetOptionContractsOk() (*[]ExchangeInformationResponseOptionContractsInner, bool)`

GetOptionContractsOk returns a tuple with the OptionContracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionContracts

`func (o *ExchangeInformationResponse) SetOptionContracts(v []ExchangeInformationResponseOptionContractsInner)`

SetOptionContracts sets OptionContracts field to given value.

### HasOptionContracts

`func (o *ExchangeInformationResponse) HasOptionContracts() bool`

HasOptionContracts returns a boolean if a field has been set.

### GetOptionAssets

`func (o *ExchangeInformationResponse) GetOptionAssets() []ExchangeInformationResponseOptionAssetsInner`

GetOptionAssets returns the OptionAssets field if non-nil, zero value otherwise.

### GetOptionAssetsOk

`func (o *ExchangeInformationResponse) GetOptionAssetsOk() (*[]ExchangeInformationResponseOptionAssetsInner, bool)`

GetOptionAssetsOk returns a tuple with the OptionAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionAssets

`func (o *ExchangeInformationResponse) SetOptionAssets(v []ExchangeInformationResponseOptionAssetsInner)`

SetOptionAssets sets OptionAssets field to given value.

### HasOptionAssets

`func (o *ExchangeInformationResponse) HasOptionAssets() bool`

HasOptionAssets returns a boolean if a field has been set.

### GetOptionSymbols

`func (o *ExchangeInformationResponse) GetOptionSymbols() []ExchangeInformationResponseOptionSymbolsInner`

GetOptionSymbols returns the OptionSymbols field if non-nil, zero value otherwise.

### GetOptionSymbolsOk

`func (o *ExchangeInformationResponse) GetOptionSymbolsOk() (*[]ExchangeInformationResponseOptionSymbolsInner, bool)`

GetOptionSymbolsOk returns a tuple with the OptionSymbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSymbols

`func (o *ExchangeInformationResponse) SetOptionSymbols(v []ExchangeInformationResponseOptionSymbolsInner)`

SetOptionSymbols sets OptionSymbols field to given value.

### HasOptionSymbols

`func (o *ExchangeInformationResponse) HasOptionSymbols() bool`

HasOptionSymbols returns a boolean if a field has been set.

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


[[Back to README]](../README.md)


