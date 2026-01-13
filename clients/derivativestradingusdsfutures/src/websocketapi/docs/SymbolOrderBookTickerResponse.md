# SymbolOrderBookTickerResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**[]SymbolOrderBookTickerResponse1Result**](SymbolOrderBookTickerResponse1Result.md) |  | [optional] 
**RateLimits** | Pointer to [**[]SymbolOrderBookTickerResponse1RateLimitsInner**](SymbolOrderBookTickerResponse1RateLimitsInner.md) |  | [optional] 

## Methods

### NewSymbolOrderBookTickerResponse

`func NewSymbolOrderBookTickerResponse() *SymbolOrderBookTickerResponse`

NewSymbolOrderBookTickerResponse instantiates a new SymbolOrderBookTickerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSymbolOrderBookTickerResponseWithDefaults

`func NewSymbolOrderBookTickerResponseWithDefaults() *SymbolOrderBookTickerResponse`

NewSymbolOrderBookTickerResponseWithDefaults instantiates a new SymbolOrderBookTickerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SymbolOrderBookTickerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SymbolOrderBookTickerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SymbolOrderBookTickerResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SymbolOrderBookTickerResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *SymbolOrderBookTickerResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SymbolOrderBookTickerResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SymbolOrderBookTickerResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SymbolOrderBookTickerResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *SymbolOrderBookTickerResponse) GetResult() []SymbolOrderBookTickerResponse1Result`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SymbolOrderBookTickerResponse) GetResultOk() (*[]SymbolOrderBookTickerResponse1Result, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SymbolOrderBookTickerResponse) SetResult(v []SymbolOrderBookTickerResponse1Result)`

SetResult sets Result field to given value.

### HasResult

`func (o *SymbolOrderBookTickerResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *SymbolOrderBookTickerResponse) GetRateLimits() []SymbolOrderBookTickerResponse1RateLimitsInner`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *SymbolOrderBookTickerResponse) GetRateLimitsOk() (*[]SymbolOrderBookTickerResponse1RateLimitsInner, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *SymbolOrderBookTickerResponse) SetRateLimits(v []SymbolOrderBookTickerResponse1RateLimitsInner)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *SymbolOrderBookTickerResponse) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


