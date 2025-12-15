# UserDataStreamSubscribeResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**UserDataStreamSubscribeResponseResult**](UserDataStreamSubscribeResponseResult.md) |  | [optional] 

## Methods

### NewUserDataStreamSubscribeResponse

`func NewUserDataStreamSubscribeResponse() *UserDataStreamSubscribeResponse`

NewUserDataStreamSubscribeResponse instantiates a new UserDataStreamSubscribeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDataStreamSubscribeResponseWithDefaults

`func NewUserDataStreamSubscribeResponseWithDefaults() *UserDataStreamSubscribeResponse`

NewUserDataStreamSubscribeResponseWithDefaults instantiates a new UserDataStreamSubscribeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserDataStreamSubscribeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserDataStreamSubscribeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserDataStreamSubscribeResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserDataStreamSubscribeResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *UserDataStreamSubscribeResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserDataStreamSubscribeResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserDataStreamSubscribeResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserDataStreamSubscribeResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *UserDataStreamSubscribeResponse) GetResult() UserDataStreamSubscribeResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *UserDataStreamSubscribeResponse) GetResultOk() (*UserDataStreamSubscribeResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *UserDataStreamSubscribeResponse) SetResult(v UserDataStreamSubscribeResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *UserDataStreamSubscribeResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to README]](../README.md)


