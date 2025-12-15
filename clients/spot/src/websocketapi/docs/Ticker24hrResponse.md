# Ticker24hrResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**[]Ticker24hrResponse2ResultInner**](Ticker24hrResponse2ResultInner.md) |  | [optional] 
**RateLimits** | Pointer to [**[]RateLimits**](RateLimits.md) |  | [optional] 

## Methods

### NewTicker24hrResponse

`func NewTicker24hrResponse() *Ticker24hrResponse`

NewTicker24hrResponse instantiates a new Ticker24hrResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicker24hrResponseWithDefaults

`func NewTicker24hrResponseWithDefaults() *Ticker24hrResponse`

NewTicker24hrResponseWithDefaults instantiates a new Ticker24hrResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Ticker24hrResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ticker24hrResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Ticker24hrResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Ticker24hrResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *Ticker24hrResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Ticker24hrResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Ticker24hrResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Ticker24hrResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *Ticker24hrResponse) GetResult() []Ticker24hrResponse2ResultInner`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Ticker24hrResponse) GetResultOk() (*[]Ticker24hrResponse2ResultInner, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Ticker24hrResponse) SetResult(v []Ticker24hrResponse2ResultInner)`

SetResult sets Result field to given value.

### HasResult

`func (o *Ticker24hrResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *Ticker24hrResponse) GetRateLimits() []RateLimits`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *Ticker24hrResponse) GetRateLimitsOk() (*[]RateLimits, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *Ticker24hrResponse) SetRateLimits(v []RateLimits)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *Ticker24hrResponse) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


