# Ticker24hrResponse2

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**[]Ticker24hrResponse2ResultInner**](Ticker24hrResponse2ResultInner.md) |  | [optional] 
**RateLimits** | Pointer to [**[]RateLimits**](RateLimits.md) |  | [optional] 

## Methods

### NewTicker24hrResponse2

`func NewTicker24hrResponse2() *Ticker24hrResponse2`

NewTicker24hrResponse2 instantiates a new Ticker24hrResponse2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicker24hrResponse2WithDefaults

`func NewTicker24hrResponse2WithDefaults() *Ticker24hrResponse2`

NewTicker24hrResponse2WithDefaults instantiates a new Ticker24hrResponse2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Ticker24hrResponse2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ticker24hrResponse2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Ticker24hrResponse2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Ticker24hrResponse2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *Ticker24hrResponse2) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Ticker24hrResponse2) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Ticker24hrResponse2) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Ticker24hrResponse2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *Ticker24hrResponse2) GetResult() []Ticker24hrResponse2ResultInner`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Ticker24hrResponse2) GetResultOk() (*[]Ticker24hrResponse2ResultInner, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Ticker24hrResponse2) SetResult(v []Ticker24hrResponse2ResultInner)`

SetResult sets Result field to given value.

### HasResult

`func (o *Ticker24hrResponse2) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *Ticker24hrResponse2) GetRateLimits() []RateLimits`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *Ticker24hrResponse2) GetRateLimitsOk() (*[]RateLimits, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *Ticker24hrResponse2) SetRateLimits(v []RateLimits)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *Ticker24hrResponse2) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


