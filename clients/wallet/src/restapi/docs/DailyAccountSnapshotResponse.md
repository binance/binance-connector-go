# DailyAccountSnapshotResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int64** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**SnapshotVos** | Pointer to [**[]DailyAccountSnapshotResponseSnapshotVosInner**](DailyAccountSnapshotResponseSnapshotVosInner.md) |  | [optional] 

## Methods

### NewDailyAccountSnapshotResponse

`func NewDailyAccountSnapshotResponse() *DailyAccountSnapshotResponse`

NewDailyAccountSnapshotResponse instantiates a new DailyAccountSnapshotResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailyAccountSnapshotResponseWithDefaults

`func NewDailyAccountSnapshotResponseWithDefaults() *DailyAccountSnapshotResponse`

NewDailyAccountSnapshotResponseWithDefaults instantiates a new DailyAccountSnapshotResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *DailyAccountSnapshotResponse) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DailyAccountSnapshotResponse) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DailyAccountSnapshotResponse) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *DailyAccountSnapshotResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMsg

`func (o *DailyAccountSnapshotResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *DailyAccountSnapshotResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *DailyAccountSnapshotResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *DailyAccountSnapshotResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetSnapshotVos

`func (o *DailyAccountSnapshotResponse) GetSnapshotVos() []DailyAccountSnapshotResponseSnapshotVosInner`

GetSnapshotVos returns the SnapshotVos field if non-nil, zero value otherwise.

### GetSnapshotVosOk

`func (o *DailyAccountSnapshotResponse) GetSnapshotVosOk() (*[]DailyAccountSnapshotResponseSnapshotVosInner, bool)`

GetSnapshotVosOk returns a tuple with the SnapshotVos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotVos

`func (o *DailyAccountSnapshotResponse) SetSnapshotVos(v []DailyAccountSnapshotResponseSnapshotVosInner)`

SetSnapshotVos sets SnapshotVos field to given value.

### HasSnapshotVos

`func (o *DailyAccountSnapshotResponse) HasSnapshotVos() bool`

HasSnapshotVos returns a boolean if a field has been set.


[[Back to README]](../README.md)


