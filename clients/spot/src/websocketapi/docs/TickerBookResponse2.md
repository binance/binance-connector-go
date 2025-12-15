# TickerBookResponse2

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**[]TickerBookResponse2ResultInner**](TickerBookResponse2ResultInner.md) |  | [optional] 
**RateLimits** | Pointer to [**[]RateLimits**](RateLimits.md) |  | [optional] 

## Methods

### NewTickerBookResponse2

`func NewTickerBookResponse2() *TickerBookResponse2`

NewTickerBookResponse2 instantiates a new TickerBookResponse2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTickerBookResponse2WithDefaults

`func NewTickerBookResponse2WithDefaults() *TickerBookResponse2`

NewTickerBookResponse2WithDefaults instantiates a new TickerBookResponse2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TickerBookResponse2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TickerBookResponse2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TickerBookResponse2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TickerBookResponse2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *TickerBookResponse2) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TickerBookResponse2) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TickerBookResponse2) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TickerBookResponse2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *TickerBookResponse2) GetResult() []TickerBookResponse2ResultInner`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TickerBookResponse2) GetResultOk() (*[]TickerBookResponse2ResultInner, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TickerBookResponse2) SetResult(v []TickerBookResponse2ResultInner)`

SetResult sets Result field to given value.

### HasResult

`func (o *TickerBookResponse2) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *TickerBookResponse2) GetRateLimits() []RateLimits`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *TickerBookResponse2) GetRateLimitsOk() (*[]RateLimits, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *TickerBookResponse2) SetRateLimits(v []RateLimits)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *TickerBookResponse2) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


