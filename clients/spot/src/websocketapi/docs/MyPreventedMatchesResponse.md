# MyPreventedMatchesResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**[]MyPreventedMatchesResponseResultInner**](MyPreventedMatchesResponseResultInner.md) |  | [optional] 
**RateLimits** | Pointer to [**[]RateLimits**](RateLimits.md) |  | [optional] 

## Methods

### NewMyPreventedMatchesResponse

`func NewMyPreventedMatchesResponse() *MyPreventedMatchesResponse`

NewMyPreventedMatchesResponse instantiates a new MyPreventedMatchesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMyPreventedMatchesResponseWithDefaults

`func NewMyPreventedMatchesResponseWithDefaults() *MyPreventedMatchesResponse`

NewMyPreventedMatchesResponseWithDefaults instantiates a new MyPreventedMatchesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MyPreventedMatchesResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MyPreventedMatchesResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MyPreventedMatchesResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MyPreventedMatchesResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *MyPreventedMatchesResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MyPreventedMatchesResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MyPreventedMatchesResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MyPreventedMatchesResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *MyPreventedMatchesResponse) GetResult() []MyPreventedMatchesResponseResultInner`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *MyPreventedMatchesResponse) GetResultOk() (*[]MyPreventedMatchesResponseResultInner, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *MyPreventedMatchesResponse) SetResult(v []MyPreventedMatchesResponseResultInner)`

SetResult sets Result field to given value.

### HasResult

`func (o *MyPreventedMatchesResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *MyPreventedMatchesResponse) GetRateLimits() []RateLimits`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *MyPreventedMatchesResponse) GetRateLimitsOk() (*[]RateLimits, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *MyPreventedMatchesResponse) SetRateLimits(v []RateLimits)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *MyPreventedMatchesResponse) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


