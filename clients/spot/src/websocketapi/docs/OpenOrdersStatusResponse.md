# OpenOrdersStatusResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**[]OpenOrdersStatusResponseResultInner**](OpenOrdersStatusResponseResultInner.md) |  | [optional] 
**RateLimits** | Pointer to [**[]RateLimits**](RateLimits.md) |  | [optional] 

## Methods

### NewOpenOrdersStatusResponse

`func NewOpenOrdersStatusResponse() *OpenOrdersStatusResponse`

NewOpenOrdersStatusResponse instantiates a new OpenOrdersStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenOrdersStatusResponseWithDefaults

`func NewOpenOrdersStatusResponseWithDefaults() *OpenOrdersStatusResponse`

NewOpenOrdersStatusResponseWithDefaults instantiates a new OpenOrdersStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpenOrdersStatusResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpenOrdersStatusResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpenOrdersStatusResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OpenOrdersStatusResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *OpenOrdersStatusResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OpenOrdersStatusResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OpenOrdersStatusResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OpenOrdersStatusResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *OpenOrdersStatusResponse) GetResult() []OpenOrdersStatusResponseResultInner`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OpenOrdersStatusResponse) GetResultOk() (*[]OpenOrdersStatusResponseResultInner, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OpenOrdersStatusResponse) SetResult(v []OpenOrdersStatusResponseResultInner)`

SetResult sets Result field to given value.

### HasResult

`func (o *OpenOrdersStatusResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *OpenOrdersStatusResponse) GetRateLimits() []RateLimits`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *OpenOrdersStatusResponse) GetRateLimitsOk() (*[]RateLimits, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *OpenOrdersStatusResponse) SetRateLimits(v []RateLimits)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *OpenOrdersStatusResponse) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


