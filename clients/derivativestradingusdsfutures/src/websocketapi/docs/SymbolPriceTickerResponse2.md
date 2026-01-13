# SymbolPriceTickerResponse2

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**[]SymbolPriceTickerResponse1Result**](SymbolPriceTickerResponse1Result.md) |  | [optional] 
**RateLimits** | Pointer to [**[]SymbolOrderBookTickerResponse1RateLimitsInner**](SymbolOrderBookTickerResponse1RateLimitsInner.md) |  | [optional] 

## Methods

### NewSymbolPriceTickerResponse2

`func NewSymbolPriceTickerResponse2() *SymbolPriceTickerResponse2`

NewSymbolPriceTickerResponse2 instantiates a new SymbolPriceTickerResponse2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSymbolPriceTickerResponse2WithDefaults

`func NewSymbolPriceTickerResponse2WithDefaults() *SymbolPriceTickerResponse2`

NewSymbolPriceTickerResponse2WithDefaults instantiates a new SymbolPriceTickerResponse2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SymbolPriceTickerResponse2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SymbolPriceTickerResponse2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SymbolPriceTickerResponse2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SymbolPriceTickerResponse2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *SymbolPriceTickerResponse2) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SymbolPriceTickerResponse2) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SymbolPriceTickerResponse2) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SymbolPriceTickerResponse2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *SymbolPriceTickerResponse2) GetResult() []SymbolPriceTickerResponse1Result`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SymbolPriceTickerResponse2) GetResultOk() (*[]SymbolPriceTickerResponse1Result, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SymbolPriceTickerResponse2) SetResult(v []SymbolPriceTickerResponse1Result)`

SetResult sets Result field to given value.

### HasResult

`func (o *SymbolPriceTickerResponse2) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *SymbolPriceTickerResponse2) GetRateLimits() []SymbolOrderBookTickerResponse1RateLimitsInner`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *SymbolPriceTickerResponse2) GetRateLimitsOk() (*[]SymbolOrderBookTickerResponse1RateLimitsInner, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *SymbolPriceTickerResponse2) SetRateLimits(v []SymbolOrderBookTickerResponse1RateLimitsInner)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *SymbolPriceTickerResponse2) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


