# CloseUserDataStreamResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to **map[string]interface{}** |  | [optional] 
**RateLimits** | Pointer to [**[]CloseUserDataStreamResponseRateLimitsInner**](CloseUserDataStreamResponseRateLimitsInner.md) |  | [optional] 

## Methods

### NewCloseUserDataStreamResponse

`func NewCloseUserDataStreamResponse() *CloseUserDataStreamResponse`

NewCloseUserDataStreamResponse instantiates a new CloseUserDataStreamResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloseUserDataStreamResponseWithDefaults

`func NewCloseUserDataStreamResponseWithDefaults() *CloseUserDataStreamResponse`

NewCloseUserDataStreamResponseWithDefaults instantiates a new CloseUserDataStreamResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloseUserDataStreamResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloseUserDataStreamResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloseUserDataStreamResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloseUserDataStreamResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *CloseUserDataStreamResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloseUserDataStreamResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloseUserDataStreamResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloseUserDataStreamResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *CloseUserDataStreamResponse) GetResult() map[string]interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CloseUserDataStreamResponse) GetResultOk() (*map[string]interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CloseUserDataStreamResponse) SetResult(v map[string]interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *CloseUserDataStreamResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *CloseUserDataStreamResponse) GetRateLimits() []CloseUserDataStreamResponseRateLimitsInner`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *CloseUserDataStreamResponse) GetRateLimitsOk() (*[]CloseUserDataStreamResponseRateLimitsInner, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *CloseUserDataStreamResponse) SetRateLimits(v []CloseUserDataStreamResponseRateLimitsInner)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *CloseUserDataStreamResponse) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


