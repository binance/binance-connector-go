# TickerBookResponse1

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**TickerBookResponse1Result**](TickerBookResponse1Result.md) |  | [optional] 
**RateLimits** | Pointer to [**[]RateLimits**](RateLimits.md) |  | [optional] 

## Methods

### NewTickerBookResponse1

`func NewTickerBookResponse1() *TickerBookResponse1`

NewTickerBookResponse1 instantiates a new TickerBookResponse1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTickerBookResponse1WithDefaults

`func NewTickerBookResponse1WithDefaults() *TickerBookResponse1`

NewTickerBookResponse1WithDefaults instantiates a new TickerBookResponse1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TickerBookResponse1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TickerBookResponse1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TickerBookResponse1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TickerBookResponse1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *TickerBookResponse1) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TickerBookResponse1) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TickerBookResponse1) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TickerBookResponse1) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *TickerBookResponse1) GetResult() TickerBookResponse1Result`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TickerBookResponse1) GetResultOk() (*TickerBookResponse1Result, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TickerBookResponse1) SetResult(v TickerBookResponse1Result)`

SetResult sets Result field to given value.

### HasResult

`func (o *TickerBookResponse1) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *TickerBookResponse1) GetRateLimits() []RateLimits`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *TickerBookResponse1) GetRateLimitsOk() (*[]RateLimits, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *TickerBookResponse1) SetRateLimits(v []RateLimits)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *TickerBookResponse1) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


