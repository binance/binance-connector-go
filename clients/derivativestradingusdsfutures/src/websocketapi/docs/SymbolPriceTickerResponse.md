# SymbolPriceTickerResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**[]SymbolPriceTickerResponse1Result**](SymbolPriceTickerResponse1Result.md) |  | [optional] 
**RateLimits** | Pointer to [**[]SymbolOrderBookTickerResponse1RateLimitsInner**](SymbolOrderBookTickerResponse1RateLimitsInner.md) |  | [optional] 

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

### GetId

`func (o *SymbolPriceTickerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SymbolPriceTickerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SymbolPriceTickerResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SymbolPriceTickerResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *SymbolPriceTickerResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SymbolPriceTickerResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SymbolPriceTickerResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SymbolPriceTickerResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *SymbolPriceTickerResponse) GetResult() []SymbolPriceTickerResponse1Result`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SymbolPriceTickerResponse) GetResultOk() (*[]SymbolPriceTickerResponse1Result, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SymbolPriceTickerResponse) SetResult(v []SymbolPriceTickerResponse1Result)`

SetResult sets Result field to given value.

### HasResult

`func (o *SymbolPriceTickerResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *SymbolPriceTickerResponse) GetRateLimits() []SymbolOrderBookTickerResponse1RateLimitsInner`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *SymbolPriceTickerResponse) GetRateLimitsOk() (*[]SymbolOrderBookTickerResponse1RateLimitsInner, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *SymbolPriceTickerResponse) SetRateLimits(v []SymbolOrderBookTickerResponse1RateLimitsInner)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *SymbolPriceTickerResponse) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


