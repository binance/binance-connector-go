# OpenOrderListsStatusResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**[]OpenOrderListsStatusResponseResultInner**](OpenOrderListsStatusResponseResultInner.md) |  | [optional] 
**RateLimits** | Pointer to [**[]RateLimits**](RateLimits.md) |  | [optional] 

## Methods

### NewOpenOrderListsStatusResponse

`func NewOpenOrderListsStatusResponse() *OpenOrderListsStatusResponse`

NewOpenOrderListsStatusResponse instantiates a new OpenOrderListsStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenOrderListsStatusResponseWithDefaults

`func NewOpenOrderListsStatusResponseWithDefaults() *OpenOrderListsStatusResponse`

NewOpenOrderListsStatusResponseWithDefaults instantiates a new OpenOrderListsStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpenOrderListsStatusResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpenOrderListsStatusResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpenOrderListsStatusResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OpenOrderListsStatusResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *OpenOrderListsStatusResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OpenOrderListsStatusResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OpenOrderListsStatusResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OpenOrderListsStatusResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *OpenOrderListsStatusResponse) GetResult() []OpenOrderListsStatusResponseResultInner`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OpenOrderListsStatusResponse) GetResultOk() (*[]OpenOrderListsStatusResponseResultInner, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OpenOrderListsStatusResponse) SetResult(v []OpenOrderListsStatusResponseResultInner)`

SetResult sets Result field to given value.

### HasResult

`func (o *OpenOrderListsStatusResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *OpenOrderListsStatusResponse) GetRateLimits() []RateLimits`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *OpenOrderListsStatusResponse) GetRateLimitsOk() (*[]RateLimits, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *OpenOrderListsStatusResponse) SetRateLimits(v []RateLimits)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *OpenOrderListsStatusResponse) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


