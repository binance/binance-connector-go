# SystemStatusResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int64** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 

## Methods

### NewSystemStatusResponse

`func NewSystemStatusResponse() *SystemStatusResponse`

NewSystemStatusResponse instantiates a new SystemStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemStatusResponseWithDefaults

`func NewSystemStatusResponseWithDefaults() *SystemStatusResponse`

NewSystemStatusResponseWithDefaults instantiates a new SystemStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SystemStatusResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SystemStatusResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SystemStatusResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SystemStatusResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMsg

`func (o *SystemStatusResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *SystemStatusResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *SystemStatusResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *SystemStatusResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to README]](../README.md)


