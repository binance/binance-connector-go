# ExchangeInfoResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Timezone** | Pointer to **string** | Always &#x60;UTC&#x60;. | [optional] 
**Symbols** | Pointer to [**[]ExchangeInfoResponseSymbolsInner**](ExchangeInfoResponseSymbolsInner.md) | List of tradable symbols with trading rules. | [optional] 

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


[[Back to README]](../README.md)


