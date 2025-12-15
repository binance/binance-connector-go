# AllOrdersResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**[]AllOrdersResponseResultInner**](AllOrdersResponseResultInner.md) |  | [optional] 
**RateLimits** | Pointer to [**[]RateLimits**](RateLimits.md) |  | [optional] 

## Methods

### NewAllOrdersResponse

`func NewAllOrdersResponse() *AllOrdersResponse`

NewAllOrdersResponse instantiates a new AllOrdersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllOrdersResponseWithDefaults

`func NewAllOrdersResponseWithDefaults() *AllOrdersResponse`

NewAllOrdersResponseWithDefaults instantiates a new AllOrdersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AllOrdersResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AllOrdersResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AllOrdersResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AllOrdersResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *AllOrdersResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AllOrdersResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AllOrdersResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AllOrdersResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *AllOrdersResponse) GetResult() []AllOrdersResponseResultInner`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AllOrdersResponse) GetResultOk() (*[]AllOrdersResponseResultInner, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AllOrdersResponse) SetResult(v []AllOrdersResponseResultInner)`

SetResult sets Result field to given value.

### HasResult

`func (o *AllOrdersResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *AllOrdersResponse) GetRateLimits() []RateLimits`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *AllOrdersResponse) GetRateLimitsOk() (*[]RateLimits, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *AllOrdersResponse) SetRateLimits(v []RateLimits)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *AllOrdersResponse) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


