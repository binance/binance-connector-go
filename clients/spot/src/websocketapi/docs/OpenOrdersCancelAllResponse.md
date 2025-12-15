# OpenOrdersCancelAllResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**[]OpenOrdersCancelAllResponseResultInner**](OpenOrdersCancelAllResponseResultInner.md) |  | [optional] 
**RateLimits** | Pointer to [**[]RateLimits**](RateLimits.md) |  | [optional] 

## Methods

### NewOpenOrdersCancelAllResponse

`func NewOpenOrdersCancelAllResponse() *OpenOrdersCancelAllResponse`

NewOpenOrdersCancelAllResponse instantiates a new OpenOrdersCancelAllResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenOrdersCancelAllResponseWithDefaults

`func NewOpenOrdersCancelAllResponseWithDefaults() *OpenOrdersCancelAllResponse`

NewOpenOrdersCancelAllResponseWithDefaults instantiates a new OpenOrdersCancelAllResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpenOrdersCancelAllResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpenOrdersCancelAllResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpenOrdersCancelAllResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OpenOrdersCancelAllResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *OpenOrdersCancelAllResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OpenOrdersCancelAllResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OpenOrdersCancelAllResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OpenOrdersCancelAllResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *OpenOrdersCancelAllResponse) GetResult() []OpenOrdersCancelAllResponseResultInner`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OpenOrdersCancelAllResponse) GetResultOk() (*[]OpenOrdersCancelAllResponseResultInner, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OpenOrdersCancelAllResponse) SetResult(v []OpenOrdersCancelAllResponseResultInner)`

SetResult sets Result field to given value.

### HasResult

`func (o *OpenOrdersCancelAllResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *OpenOrdersCancelAllResponse) GetRateLimits() []RateLimits`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *OpenOrdersCancelAllResponse) GetRateLimitsOk() (*[]RateLimits, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *OpenOrdersCancelAllResponse) SetRateLimits(v []RateLimits)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *OpenOrdersCancelAllResponse) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


