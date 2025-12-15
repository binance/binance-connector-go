# QueryUserUniversalTransferHistoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** |  | [optional] 
**Rows** | Pointer to [**[]QueryUserUniversalTransferHistoryResponseRowsInner**](QueryUserUniversalTransferHistoryResponseRowsInner.md) |  | [optional] 

## Methods

### NewQueryUserUniversalTransferHistoryResponse

`func NewQueryUserUniversalTransferHistoryResponse() *QueryUserUniversalTransferHistoryResponse`

NewQueryUserUniversalTransferHistoryResponse instantiates a new QueryUserUniversalTransferHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryUserUniversalTransferHistoryResponseWithDefaults

`func NewQueryUserUniversalTransferHistoryResponseWithDefaults() *QueryUserUniversalTransferHistoryResponse`

NewQueryUserUniversalTransferHistoryResponseWithDefaults instantiates a new QueryUserUniversalTransferHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *QueryUserUniversalTransferHistoryResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *QueryUserUniversalTransferHistoryResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *QueryUserUniversalTransferHistoryResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *QueryUserUniversalTransferHistoryResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetRows

`func (o *QueryUserUniversalTransferHistoryResponse) GetRows() []QueryUserUniversalTransferHistoryResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *QueryUserUniversalTransferHistoryResponse) GetRowsOk() (*[]QueryUserUniversalTransferHistoryResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *QueryUserUniversalTransferHistoryResponse) SetRows(v []QueryUserUniversalTransferHistoryResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *QueryUserUniversalTransferHistoryResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to README]](../README.md)


