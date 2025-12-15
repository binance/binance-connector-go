# QueryManagedSubAccountSnapshotResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int64** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**SnapshotVos** | Pointer to [**[]QueryManagedSubAccountSnapshotResponseSnapshotVosInner**](QueryManagedSubAccountSnapshotResponseSnapshotVosInner.md) |  | [optional] 

## Methods

### NewQueryManagedSubAccountSnapshotResponse

`func NewQueryManagedSubAccountSnapshotResponse() *QueryManagedSubAccountSnapshotResponse`

NewQueryManagedSubAccountSnapshotResponse instantiates a new QueryManagedSubAccountSnapshotResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryManagedSubAccountSnapshotResponseWithDefaults

`func NewQueryManagedSubAccountSnapshotResponseWithDefaults() *QueryManagedSubAccountSnapshotResponse`

NewQueryManagedSubAccountSnapshotResponseWithDefaults instantiates a new QueryManagedSubAccountSnapshotResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *QueryManagedSubAccountSnapshotResponse) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *QueryManagedSubAccountSnapshotResponse) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *QueryManagedSubAccountSnapshotResponse) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *QueryManagedSubAccountSnapshotResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMsg

`func (o *QueryManagedSubAccountSnapshotResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *QueryManagedSubAccountSnapshotResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *QueryManagedSubAccountSnapshotResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *QueryManagedSubAccountSnapshotResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetSnapshotVos

`func (o *QueryManagedSubAccountSnapshotResponse) GetSnapshotVos() []QueryManagedSubAccountSnapshotResponseSnapshotVosInner`

GetSnapshotVos returns the SnapshotVos field if non-nil, zero value otherwise.

### GetSnapshotVosOk

`func (o *QueryManagedSubAccountSnapshotResponse) GetSnapshotVosOk() (*[]QueryManagedSubAccountSnapshotResponseSnapshotVosInner, bool)`

GetSnapshotVosOk returns a tuple with the SnapshotVos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotVos

`func (o *QueryManagedSubAccountSnapshotResponse) SetSnapshotVos(v []QueryManagedSubAccountSnapshotResponseSnapshotVosInner)`

SetSnapshotVos sets SnapshotVos field to given value.

### HasSnapshotVos

`func (o *QueryManagedSubAccountSnapshotResponse) HasSnapshotVos() bool`

HasSnapshotVos returns a boolean if a field has been set.


[[Back to README]](../README.md)


