# UserDataStreamUnsubscribeResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUserDataStreamUnsubscribeResponse

`func NewUserDataStreamUnsubscribeResponse() *UserDataStreamUnsubscribeResponse`

NewUserDataStreamUnsubscribeResponse instantiates a new UserDataStreamUnsubscribeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDataStreamUnsubscribeResponseWithDefaults

`func NewUserDataStreamUnsubscribeResponseWithDefaults() *UserDataStreamUnsubscribeResponse`

NewUserDataStreamUnsubscribeResponseWithDefaults instantiates a new UserDataStreamUnsubscribeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserDataStreamUnsubscribeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserDataStreamUnsubscribeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserDataStreamUnsubscribeResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserDataStreamUnsubscribeResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *UserDataStreamUnsubscribeResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserDataStreamUnsubscribeResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserDataStreamUnsubscribeResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserDataStreamUnsubscribeResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *UserDataStreamUnsubscribeResponse) GetResult() map[string]interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *UserDataStreamUnsubscribeResponse) GetResultOk() (*map[string]interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *UserDataStreamUnsubscribeResponse) SetResult(v map[string]interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *UserDataStreamUnsubscribeResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to README]](../README.md)


