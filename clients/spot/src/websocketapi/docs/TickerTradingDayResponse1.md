# TickerTradingDayResponse1

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**TickerTradingDayResponse1Result**](TickerTradingDayResponse1Result.md) |  | [optional] 
**RateLimits** | Pointer to [**[]OrderAmendmentsResponseRateLimitsInner**](OrderAmendmentsResponseRateLimitsInner.md) |  | [optional] 

## Methods

### NewTickerTradingDayResponse1

`func NewTickerTradingDayResponse1() *TickerTradingDayResponse1`

NewTickerTradingDayResponse1 instantiates a new TickerTradingDayResponse1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTickerTradingDayResponse1WithDefaults

`func NewTickerTradingDayResponse1WithDefaults() *TickerTradingDayResponse1`

NewTickerTradingDayResponse1WithDefaults instantiates a new TickerTradingDayResponse1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TickerTradingDayResponse1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TickerTradingDayResponse1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TickerTradingDayResponse1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TickerTradingDayResponse1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *TickerTradingDayResponse1) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TickerTradingDayResponse1) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TickerTradingDayResponse1) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TickerTradingDayResponse1) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *TickerTradingDayResponse1) GetResult() TickerTradingDayResponse1Result`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TickerTradingDayResponse1) GetResultOk() (*TickerTradingDayResponse1Result, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TickerTradingDayResponse1) SetResult(v TickerTradingDayResponse1Result)`

SetResult sets Result field to given value.

### HasResult

`func (o *TickerTradingDayResponse1) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *TickerTradingDayResponse1) GetRateLimits() []OrderAmendmentsResponseRateLimitsInner`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *TickerTradingDayResponse1) GetRateLimitsOk() (*[]OrderAmendmentsResponseRateLimitsInner, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *TickerTradingDayResponse1) SetRateLimits(v []OrderAmendmentsResponseRateLimitsInner)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *TickerTradingDayResponse1) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


