# SessionStatusResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**SessionStatusResponseResult**](SessionStatusResponseResult.md) |  | [optional] 

## Methods

### NewSessionStatusResponse

`func NewSessionStatusResponse() *SessionStatusResponse`

NewSessionStatusResponse instantiates a new SessionStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionStatusResponseWithDefaults

`func NewSessionStatusResponseWithDefaults() *SessionStatusResponse`

NewSessionStatusResponseWithDefaults instantiates a new SessionStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SessionStatusResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SessionStatusResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SessionStatusResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SessionStatusResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *SessionStatusResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SessionStatusResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SessionStatusResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SessionStatusResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *SessionStatusResponse) GetResult() SessionStatusResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SessionStatusResponse) GetResultOk() (*SessionStatusResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SessionStatusResponse) SetResult(v SessionStatusResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *SessionStatusResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to README]](../README.md)


