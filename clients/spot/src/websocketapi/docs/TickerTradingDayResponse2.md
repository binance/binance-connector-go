# TickerTradingDayResponse2

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**[]TickerTradingDayResponse2ResultInner**](TickerTradingDayResponse2ResultInner.md) |  | [optional] 
**RateLimits** | Pointer to [**[]TickerResponse2RateLimitsInner**](TickerResponse2RateLimitsInner.md) |  | [optional] 

## Methods

### NewTickerTradingDayResponse2

`func NewTickerTradingDayResponse2() *TickerTradingDayResponse2`

NewTickerTradingDayResponse2 instantiates a new TickerTradingDayResponse2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTickerTradingDayResponse2WithDefaults

`func NewTickerTradingDayResponse2WithDefaults() *TickerTradingDayResponse2`

NewTickerTradingDayResponse2WithDefaults instantiates a new TickerTradingDayResponse2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TickerTradingDayResponse2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TickerTradingDayResponse2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TickerTradingDayResponse2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TickerTradingDayResponse2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *TickerTradingDayResponse2) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TickerTradingDayResponse2) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TickerTradingDayResponse2) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TickerTradingDayResponse2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *TickerTradingDayResponse2) GetResult() []TickerTradingDayResponse2ResultInner`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TickerTradingDayResponse2) GetResultOk() (*[]TickerTradingDayResponse2ResultInner, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TickerTradingDayResponse2) SetResult(v []TickerTradingDayResponse2ResultInner)`

SetResult sets Result field to given value.

### HasResult

`func (o *TickerTradingDayResponse2) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *TickerTradingDayResponse2) GetRateLimits() []TickerResponse2RateLimitsInner`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *TickerTradingDayResponse2) GetRateLimitsOk() (*[]TickerResponse2RateLimitsInner, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *TickerTradingDayResponse2) SetRateLimits(v []TickerResponse2RateLimitsInner)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *TickerTradingDayResponse2) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


