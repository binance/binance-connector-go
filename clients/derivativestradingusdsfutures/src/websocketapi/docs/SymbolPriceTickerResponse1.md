# SymbolPriceTickerResponse1

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**SymbolPriceTickerResponse1Result**](SymbolPriceTickerResponse1Result.md) |  | [optional] 
**RateLimits** | Pointer to [**[]SymbolOrderBookTickerResponse1RateLimitsInner**](SymbolOrderBookTickerResponse1RateLimitsInner.md) |  | [optional] 

## Methods

### NewSymbolPriceTickerResponse1

`func NewSymbolPriceTickerResponse1() *SymbolPriceTickerResponse1`

NewSymbolPriceTickerResponse1 instantiates a new SymbolPriceTickerResponse1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSymbolPriceTickerResponse1WithDefaults

`func NewSymbolPriceTickerResponse1WithDefaults() *SymbolPriceTickerResponse1`

NewSymbolPriceTickerResponse1WithDefaults instantiates a new SymbolPriceTickerResponse1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SymbolPriceTickerResponse1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SymbolPriceTickerResponse1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SymbolPriceTickerResponse1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SymbolPriceTickerResponse1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *SymbolPriceTickerResponse1) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SymbolPriceTickerResponse1) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SymbolPriceTickerResponse1) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SymbolPriceTickerResponse1) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *SymbolPriceTickerResponse1) GetResult() SymbolPriceTickerResponse1Result`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SymbolPriceTickerResponse1) GetResultOk() (*SymbolPriceTickerResponse1Result, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SymbolPriceTickerResponse1) SetResult(v SymbolPriceTickerResponse1Result)`

SetResult sets Result field to given value.

### HasResult

`func (o *SymbolPriceTickerResponse1) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *SymbolPriceTickerResponse1) GetRateLimits() []SymbolOrderBookTickerResponse1RateLimitsInner`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *SymbolPriceTickerResponse1) GetRateLimitsOk() (*[]SymbolOrderBookTickerResponse1RateLimitsInner, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *SymbolPriceTickerResponse1) SetRateLimits(v []SymbolOrderBookTickerResponse1RateLimitsInner)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *SymbolPriceTickerResponse1) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


