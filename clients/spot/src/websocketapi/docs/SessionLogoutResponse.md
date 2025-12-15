# SessionLogoutResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**SessionLogoutResponseResult**](SessionLogoutResponseResult.md) |  | [optional] 

## Methods

### NewSessionLogoutResponse

`func NewSessionLogoutResponse() *SessionLogoutResponse`

NewSessionLogoutResponse instantiates a new SessionLogoutResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionLogoutResponseWithDefaults

`func NewSessionLogoutResponseWithDefaults() *SessionLogoutResponse`

NewSessionLogoutResponseWithDefaults instantiates a new SessionLogoutResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SessionLogoutResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SessionLogoutResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SessionLogoutResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SessionLogoutResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *SessionLogoutResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SessionLogoutResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SessionLogoutResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SessionLogoutResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *SessionLogoutResponse) GetResult() SessionLogoutResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SessionLogoutResponse) GetResultOk() (*SessionLogoutResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SessionLogoutResponse) SetResult(v SessionLogoutResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *SessionLogoutResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to README]](../README.md)


