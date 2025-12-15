# KlinesResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**[]KlinesItem**](KlinesItem.md) |  | [optional] 
**RateLimits** | Pointer to [**[]RateLimits**](RateLimits.md) |  | [optional] 

## Methods

### NewKlinesResponse

`func NewKlinesResponse() *KlinesResponse`

NewKlinesResponse instantiates a new KlinesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKlinesResponseWithDefaults

`func NewKlinesResponseWithDefaults() *KlinesResponse`

NewKlinesResponseWithDefaults instantiates a new KlinesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KlinesResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KlinesResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KlinesResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KlinesResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *KlinesResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KlinesResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KlinesResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KlinesResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *KlinesResponse) GetResult() []KlinesItem`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *KlinesResponse) GetResultOk() (*[]KlinesItem, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *KlinesResponse) SetResult(v []KlinesItem)`

SetResult sets Result field to given value.

### HasResult

`func (o *KlinesResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *KlinesResponse) GetRateLimits() []RateLimits`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *KlinesResponse) GetRateLimitsOk() (*[]RateLimits, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *KlinesResponse) SetRateLimits(v []RateLimits)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *KlinesResponse) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


