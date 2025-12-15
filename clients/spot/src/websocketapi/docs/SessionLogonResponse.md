# SessionLogonResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**SessionLogonResponseResult**](SessionLogonResponseResult.md) |  | [optional] 

## Methods

### NewSessionLogonResponse

`func NewSessionLogonResponse() *SessionLogonResponse`

NewSessionLogonResponse instantiates a new SessionLogonResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionLogonResponseWithDefaults

`func NewSessionLogonResponseWithDefaults() *SessionLogonResponse`

NewSessionLogonResponseWithDefaults instantiates a new SessionLogonResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SessionLogonResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SessionLogonResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SessionLogonResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SessionLogonResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *SessionLogonResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SessionLogonResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SessionLogonResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SessionLogonResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *SessionLogonResponse) GetResult() SessionLogonResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SessionLogonResponse) GetResultOk() (*SessionLogonResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SessionLogonResponse) SetResult(v SessionLogonResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *SessionLogonResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to README]](../README.md)


