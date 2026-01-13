# KeepaliveUserDataStreamResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**KeepaliveUserDataStreamResponseResult**](KeepaliveUserDataStreamResponseResult.md) |  | [optional] 
**RateLimits** | Pointer to [**[]CloseUserDataStreamResponseRateLimitsInner**](CloseUserDataStreamResponseRateLimitsInner.md) |  | [optional] 

## Methods

### NewKeepaliveUserDataStreamResponse

`func NewKeepaliveUserDataStreamResponse() *KeepaliveUserDataStreamResponse`

NewKeepaliveUserDataStreamResponse instantiates a new KeepaliveUserDataStreamResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeepaliveUserDataStreamResponseWithDefaults

`func NewKeepaliveUserDataStreamResponseWithDefaults() *KeepaliveUserDataStreamResponse`

NewKeepaliveUserDataStreamResponseWithDefaults instantiates a new KeepaliveUserDataStreamResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeepaliveUserDataStreamResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeepaliveUserDataStreamResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeepaliveUserDataStreamResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KeepaliveUserDataStreamResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *KeepaliveUserDataStreamResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KeepaliveUserDataStreamResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KeepaliveUserDataStreamResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KeepaliveUserDataStreamResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *KeepaliveUserDataStreamResponse) GetResult() KeepaliveUserDataStreamResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *KeepaliveUserDataStreamResponse) GetResultOk() (*KeepaliveUserDataStreamResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *KeepaliveUserDataStreamResponse) SetResult(v KeepaliveUserDataStreamResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *KeepaliveUserDataStreamResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *KeepaliveUserDataStreamResponse) GetRateLimits() []CloseUserDataStreamResponseRateLimitsInner`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *KeepaliveUserDataStreamResponse) GetRateLimitsOk() (*[]CloseUserDataStreamResponseRateLimitsInner, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *KeepaliveUserDataStreamResponse) SetRateLimits(v []CloseUserDataStreamResponseRateLimitsInner)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *KeepaliveUserDataStreamResponse) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


