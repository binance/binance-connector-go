# SymbolOrderBookTickerResponse2

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**[]SymbolOrderBookTickerResponse1Result**](SymbolOrderBookTickerResponse1Result.md) |  | [optional] 
**RateLimits** | Pointer to [**[]SymbolOrderBookTickerResponse1RateLimitsInner**](SymbolOrderBookTickerResponse1RateLimitsInner.md) |  | [optional] 

## Methods

### NewSymbolOrderBookTickerResponse2

`func NewSymbolOrderBookTickerResponse2() *SymbolOrderBookTickerResponse2`

NewSymbolOrderBookTickerResponse2 instantiates a new SymbolOrderBookTickerResponse2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSymbolOrderBookTickerResponse2WithDefaults

`func NewSymbolOrderBookTickerResponse2WithDefaults() *SymbolOrderBookTickerResponse2`

NewSymbolOrderBookTickerResponse2WithDefaults instantiates a new SymbolOrderBookTickerResponse2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SymbolOrderBookTickerResponse2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SymbolOrderBookTickerResponse2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SymbolOrderBookTickerResponse2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SymbolOrderBookTickerResponse2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *SymbolOrderBookTickerResponse2) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SymbolOrderBookTickerResponse2) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SymbolOrderBookTickerResponse2) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SymbolOrderBookTickerResponse2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *SymbolOrderBookTickerResponse2) GetResult() []SymbolOrderBookTickerResponse1Result`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SymbolOrderBookTickerResponse2) GetResultOk() (*[]SymbolOrderBookTickerResponse1Result, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SymbolOrderBookTickerResponse2) SetResult(v []SymbolOrderBookTickerResponse1Result)`

SetResult sets Result field to given value.

### HasResult

`func (o *SymbolOrderBookTickerResponse2) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *SymbolOrderBookTickerResponse2) GetRateLimits() []SymbolOrderBookTickerResponse1RateLimitsInner`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *SymbolOrderBookTickerResponse2) GetRateLimitsOk() (*[]SymbolOrderBookTickerResponse1RateLimitsInner, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *SymbolOrderBookTickerResponse2) SetRateLimits(v []SymbolOrderBookTickerResponse1RateLimitsInner)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *SymbolOrderBookTickerResponse2) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


