# SymbolPriceTickerResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 

## Methods

### NewSymbolPriceTickerResponse

`func NewSymbolPriceTickerResponse() *SymbolPriceTickerResponse`

NewSymbolPriceTickerResponse instantiates a new SymbolPriceTickerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSymbolPriceTickerResponseWithDefaults

`func NewSymbolPriceTickerResponseWithDefaults() *SymbolPriceTickerResponse`

NewSymbolPriceTickerResponseWithDefaults instantiates a new SymbolPriceTickerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *SymbolPriceTickerResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *SymbolPriceTickerResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *SymbolPriceTickerResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *SymbolPriceTickerResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetPrice

`func (o *SymbolPriceTickerResponse) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SymbolPriceTickerResponse) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SymbolPriceTickerResponse) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *SymbolPriceTickerResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetTime

`func (o *SymbolPriceTickerResponse) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *SymbolPriceTickerResponse) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *SymbolPriceTickerResponse) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *SymbolPriceTickerResponse) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


