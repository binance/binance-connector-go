# TradesHistoricalResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**[]TradesHistoricalResponseResultInner**](TradesHistoricalResponseResultInner.md) |  | [optional] 
**RateLimits** | Pointer to [**[]RateLimits**](RateLimits.md) |  | [optional] 

## Methods

### NewTradesHistoricalResponse

`func NewTradesHistoricalResponse() *TradesHistoricalResponse`

NewTradesHistoricalResponse instantiates a new TradesHistoricalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTradesHistoricalResponseWithDefaults

`func NewTradesHistoricalResponseWithDefaults() *TradesHistoricalResponse`

NewTradesHistoricalResponseWithDefaults instantiates a new TradesHistoricalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TradesHistoricalResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TradesHistoricalResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TradesHistoricalResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TradesHistoricalResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *TradesHistoricalResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TradesHistoricalResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TradesHistoricalResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TradesHistoricalResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *TradesHistoricalResponse) GetResult() []TradesHistoricalResponseResultInner`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TradesHistoricalResponse) GetResultOk() (*[]TradesHistoricalResponseResultInner, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TradesHistoricalResponse) SetResult(v []TradesHistoricalResponseResultInner)`

SetResult sets Result field to given value.

### HasResult

`func (o *TradesHistoricalResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *TradesHistoricalResponse) GetRateLimits() []RateLimits`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *TradesHistoricalResponse) GetRateLimitsOk() (*[]RateLimits, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *TradesHistoricalResponse) SetRateLimits(v []RateLimits)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *TradesHistoricalResponse) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


