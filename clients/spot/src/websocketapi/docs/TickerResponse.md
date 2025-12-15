# TickerResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**[]TickerResponse2ResultInner**](TickerResponse2ResultInner.md) |  | [optional] 
**RateLimits** | Pointer to [**[]RateLimits**](RateLimits.md) |  | [optional] 

## Methods

### NewTickerResponse

`func NewTickerResponse() *TickerResponse`

NewTickerResponse instantiates a new TickerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTickerResponseWithDefaults

`func NewTickerResponseWithDefaults() *TickerResponse`

NewTickerResponseWithDefaults instantiates a new TickerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TickerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TickerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TickerResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TickerResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *TickerResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TickerResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TickerResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TickerResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *TickerResponse) GetResult() []TickerResponse2ResultInner`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TickerResponse) GetResultOk() (*[]TickerResponse2ResultInner, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TickerResponse) SetResult(v []TickerResponse2ResultInner)`

SetResult sets Result field to given value.

### HasResult

`func (o *TickerResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *TickerResponse) GetRateLimits() []RateLimits`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *TickerResponse) GetRateLimitsOk() (*[]RateLimits, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *TickerResponse) SetRateLimits(v []RateLimits)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *TickerResponse) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


